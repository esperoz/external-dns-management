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
	//	"github.com/joeig/go-powerdns/v2"
	pdns "github.com/joeig/go-powerdns/v2"

	"github.com/gardener/external-dns-management/pkg/dns/provider/raw"
)

type Record struct {
	rset *pdns.RRset
	raw.Record
}

func NewRecordFromRecordset(rset *pdns.RRset) *Record {
	record := &Record{
		rset: rset,
	}
	return record
}

func (r *Record) GetType() string    { return string(*r.rset.Type) }
func (r *Record) GetId() string      { return *r.rset.Name }
func (r *Record) GetDNSName() string { return *r.rset.Name }
func (r *Record) GetValue() string {
	var val string

	for _, rec := range r.rset.Records {
		if *r.rset.Type == pdns.RRTypeTXT {
			val = raw.EnsureQuotedText(*rec.Content)
		} else {
			val = *rec.Content
		}
	}

	return val
}
func (r *Record) GetTTL() int      { return int(*r.rset.TTL) }
func (r *Record) SetTTL(ttl int)   { *r.rset.TTL = uint32(ttl) }
func (r *Record) Copy() raw.Record { n := *r; return &n }
