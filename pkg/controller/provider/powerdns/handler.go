/*
 * 	This file is licensed under the Apache Software License, v. 2 except as noted otherwise in the LICENSE file
 *
 *  Licensed under the Apache License, Version 2.0 (the "License");
 *  you may not use this file except in compliance with the License.
 *  You may obtain a copy of the License at
 *
 *       http://www.apache.org/licenses/LICENSE-2.0
 *
 *  Unless required by applicable law or agreed to in writing, software
 *  distributed under the License is distributed on an "AS IS" BASIS,
 *  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 *  See the License for the specific language governing permissions and
 *  limitations under the License.
 *
 */

// to be removed - helm install dns-controller charts/external-dns-management --namespace=pdns --set configuration.identifier=ocp

package powerdns

import (
	"context"
	//	"fmt"

	//	pdns "github.com/joeig/go-powerdns/v2"
	//	"github.com/netlify/open-api/go/models"

	"github.com/gardener/controller-manager-library/pkg/logger"
	"github.com/gardener/external-dns-management/pkg/dns"
	"github.com/gardener/external-dns-management/pkg/dns/provider"
	"github.com/gardener/external-dns-management/pkg/dns/provider/raw"
)

type Handler struct {
	cache provider.ZoneCache
	provider.DefaultDNSHandler
	config  provider.DNSHandlerConfig
	execman *PowerDNSExecMan
	ctx     context.Context
}

type PowerDNSConfig struct {
	Server     *string `json:"host,omitempty"`
	VHost      *string `json:"vhost,omitempty"`
	APIKey     *string `json:"apikey,omitempty"`
	Basedomain *string `json:"basedomain,omitempty"`
}

var _ provider.DNSHandler = &Handler{}

func NewHandler(config *provider.DNSHandlerConfig) (provider.DNSHandler, error) {

	//get provider config
	pdnsconfig := &PowerDNSConfig{}

	var err error

	if err := config.FillRequiredProperty(&pdnsconfig.Server, "SERVER", "server"); err != nil {
		return nil, err
	}

	if err := config.FillRequiredProperty(&pdnsconfig.VHost, "VHOST", "vhost"); err != nil {
		return nil, err
	}

	if err := config.FillRequiredProperty(&pdnsconfig.APIKey, "APIKEY", "apikey"); err != nil {
		return nil, err
	}

	if err := config.FillRequiredProperty(&pdnsconfig.Basedomain, "BASEDOMAIN", "basedomain"); err != nil {
		return nil, err
	}

	// init handler
	h := &Handler{
		DefaultDNSHandler: provider.NewDefaultDNSHandler(TYPE_CODE),
		config:            *config,
		execman:           NewExecutor(config.Logger, pdnsconfig, config.Metrics),
		ctx:               config.Context,
	}

	// TODO: check zone existence and create new if any

	forwardedDomains := provider.NewForwardedDomainsHandlerData()
	h.cache, err = provider.NewZoneCache(config.CacheConfig, config.Metrics, forwardedDomains, h.getZones, h.getZoneState)
	if err != nil {
		return nil, err
	}

	logger.Infof("PowerDNS handler for %s", *pdnsconfig.Basedomain)

	return h, nil
}

func (h *Handler) getZones(cache provider.ZoneCache) (provider.DNSHostedZones, error) {

	h.config.RateLimiter.Accept()
	h.config.Metrics.AddGenericRequests(provider.M_LISTZONES, 1)

	zonelist, err := h.execman.client.Zones.List()
	if err != nil {
		return nil, err
	}

	zones := provider.DNSHostedZones{}

	for _, z := range zonelist {
		logger.Infof("PowerDNS got zone %s", *z.Name)
		forwarded := []string{}
		hostedZone := provider.NewDNSHostedZone(h.ProviderType(), *z.ID, dns.NormalizeHostname(*z.Name), *z.Name, forwarded, false)
		zones = append(zones, hostedZone)
	}

	return zones, nil
}

func (h *Handler) getZoneState(zone provider.DNSHostedZone, cache provider.ZoneCache) (provider.DNSZoneState, error) {
	h.config.Metrics.AddZoneRequests(zone.Id(), provider.M_LISTRECORDS, 1)
	//get zones from PowerDNS
	pdnszone, err := h.execman.client.Zones.Get(zone.Domain())
	if err != nil {
		return nil, err
	}

	state := raw.NewState()
	for _, rrset := range pdnszone.RRsets {
		// add state record from PowerDNS zone record
		state.AddRecord(NewRecordFromRecordset(&rrset).Copy())
	}

	// generate DNSSets from records
	state.CalculateDNSSets()
	return state, nil
}

func (h *Handler) ExecuteRequests(logger logger.LogContext, zone provider.DNSHostedZone, state provider.DNSZoneState, reqs []*provider.ChangeRequest) error {
	err := raw.ExecuteRequests(logger, &h.config, h.execman, zone, state, reqs)
	h.cache.ApplyRequests(logger, err, zone, reqs)
	return err
}

func (h *Handler) Release() {
	h.cache.Release()
}

func (h *Handler) GetZones() (provider.DNSHostedZones, error) {
	return h.cache.GetZones()
}

func (h *Handler) GetZoneState(zone provider.DNSHostedZone) (provider.DNSZoneState, error) {
	return h.cache.GetZoneState(zone)
}

func (h *Handler) ReportZoneStateConflict(zone provider.DNSHostedZone, err error) bool {
	return h.cache.ReportZoneStateConflict(zone, err)
}
