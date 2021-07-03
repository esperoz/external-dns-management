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

package powerdns

import (
	"context"
	//	"encoding/json"
	//	"fmt"
	//	"io/ioutil"
	//	"net/url"
	//	"os"
	//	"strconv"

	//	"github.com/aws/aws-sdk-go/aws/client"
	"github.com/gardener/controller-manager-library/pkg/logger"
	//	"github.com/pkg/errors"

	//	"github.com/gardener/external-dns-management/pkg/dns"
	"github.com/gardener/external-dns-management/pkg/dns/provider"
	"github.com/gardener/external-dns-management/pkg/dns/provider/raw"
)

type Handler struct {
	provider.ZoneCache
	provider.DefaultDNSHandler
	config  provider.DNSHandlerConfig
	execman *PowerDNSExecMan //to be replaced
	ctx     context.Context
}

// type InfobloxConfig struct {
// 	Host            *string `json:"host,omitempty"`
// 	Port            *int    `json:"port,omitempty"`
// 	SSLVerify       *bool   `json:"sslVerify,omitempty"`
// 	Version         *string `json:"version,omitempty"`
// 	View            *string `json:"view,omitempty"`
// 	PoolConnections *int    `json:"httpPoolConnections,omitempty"`
// 	RequestTimeout  *int    `json:"httpRequestTimeout,omitempty"`
// 	CaCert          *string `json:"caCert,omitempty"`
// 	MaxResults      int     `json:"maxResults,omitempty"`
// 	ProxyURL        *string `json:"proxyUrl,omitempty"`
// }

var _ provider.DNSHandler = &Handler{}

func NewHandler(config *provider.DNSHandlerConfig) (provider.DNSHandler, error) {

	//NewExecutor()

	return nil, nil
}

func (h *Handler) getZones(cache provider.ZoneCache) (provider.DNSHostedZones, error) {

	return nil, nil
}

func (h *Handler) getZoneState(zone provider.DNSHostedZone, cache provider.ZoneCache) (provider.DNSZoneState, error) {

	return nil, nil
}

func (h *Handler) ExecuteRequests(logger logger.LogContext, zone provider.DNSHostedZone, state provider.DNSZoneState, reqs []*provider.ChangeRequest) error {
	err := raw.ExecuteRequests(logger, &h.config, h.execman, zone, state, reqs)
	h.ApplyRequests(logger, err, zone, reqs)
	return err
}
