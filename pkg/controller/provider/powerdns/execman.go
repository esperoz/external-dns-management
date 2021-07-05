/*
 * This file is licensed under the Apache Software License, v. 2 except as noted otherwise in the LICENSE file
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

package powerdns

import (
	pdns "github.com/joeig/go-powerdns/v2"

	"github.com/gardener/controller-manager-library/pkg/logger"
	"github.com/gardener/external-dns-management/pkg/dns"
	"github.com/gardener/external-dns-management/pkg/dns/provider"
	"github.com/gardener/external-dns-management/pkg/dns/provider/raw"
)

//  type Executor interface {
// 	CreateRecord(r Record, zone provider.DNSHostedZone) error
// 	UpdateRecord(r Record, zone provider.DNSHostedZone) error
// 	DeleteRecord(r Record, zone provider.DNSHostedZone) error

// 	NewRecord(fqdn, rtype, value string, zone provider.DNSHostedZone, ttl int64) Record
// }
var _ raw.Executor = (*PowerDNSExecMan)(nil)

type PowerDNSExecMan struct {
	logger   logger.LogContext
	base     string
	basezone *pdns.Zone
	client   *pdns.Client
	metrics  provider.Metrics
}

func NewExecutor(logger logger.LogContext, cfg *PowerDNSConfig, metrics provider.Metrics) *PowerDNSExecMan {
	execman := &PowerDNSExecMan{
		logger:  logger,
		base:    *cfg.Basedomain,
		client:  pdns.NewClient(*cfg.Server, *cfg.VHost, map[string]string{"X-API-Key": *cfg.APIKey}, nil),
		metrics: metrics,
	}

	bz, err := execman.client.Zones.Get(*cfg.Basedomain)
	if err != nil {
		return nil
	}

	execman.basezone = bz
	//pdns.RRTypeA
	return execman
}

func (exec *PowerDNSExecMan) CreateRecord(r raw.Record, zone provider.DNSHostedZone) error {
	//this.metrics.AddZoneRequests(zone.Id(), provider.M_CREATERECORDS, 1)
	exec.logger.Infof("PowerDNS createRecord %s with ip %s type %s at zone %s", r.GetDNSName(), r.GetValue(), r.GetType(), zone.Id())

	// really create record
	return nil
}

func (exec *PowerDNSExecMan) UpdateRecord(r raw.Record, zone provider.DNSHostedZone) error {
	exec.logger.Infof("PowerDNS updateRecord %s with ip %s type %s at zone %s", r.GetDNSName(), r.GetValue(), r.GetType(), zone.Id())
	// really update record
	return nil
}

func (exec *PowerDNSExecMan) DeleteRecord(r raw.Record, zone provider.DNSHostedZone) error {
	exec.logger.Infof("PowerDNS deleteRecord %s with ip %s type %s at zone %s", r.GetDNSName(), r.GetValue(), r.GetType(), zone.Id())
	// really delete record
	return nil
}

func (exec *PowerDNSExecMan) NewRecord(fqdn, rtype, value string, zone provider.DNSHostedZone, ttl int64) (newrecord raw.Record) {
	switch rtype {
	case dns.RS_A:

	case dns.RS_CNAME:

	case dns.RS_TXT:

	}

	exec.logger.Infof("Newrecord - %s - fqdn: %s - val: %s - zone: %s", rtype, fqdn, value, zone.Domain())

	return
}
