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
	//	"strconv"
	//	"strings"
	pdns "github.com/joeig/go-powerdns/v2"

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

type PowerDNSExecMan struct {
	base     string
	basezone *pdns.Zone
	client   *pdns.Client
}

func NewExecutor(basedomain string, srv string, vhost string, apikey string) *PowerDNSExecMan {
	execman := &PowerDNSExecMan{
		base:   basedomain,
		client: pdns.NewClient(srv, vhost, map[string]string{"X-API-Key": apikey}, nil),
	}

	bz, err := execman.client.Zones.Get(basedomain)
	if err != nil {
		return nil
	}

	execman.basezone = bz

	return execman
}

func (this *PowerDNSExecMan) CreateRecord(r raw.Record, zone provider.DNSHostedZone) error {

	return nil
}

func (this *PowerDNSExecMan) UpdateRecord(r raw.Record, zone provider.DNSHostedZone) error {

	return nil
}

func (this *PowerDNSExecMan) DeleteRecord(r raw.Record, zone provider.DNSHostedZone) error {

	return nil
}

func (this *PowerDNSExecMan) NewRecord(fqdn, rtype, value string, zone provider.DNSHostedZone, ttl int64) (newrecord raw.Record) {
	switch rtype {
	case dns.RS_A:

	case dns.RS_CNAME:

	case dns.RS_TXT:

	}

	return
}
