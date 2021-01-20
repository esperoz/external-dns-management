/*
Copyright (c) 2021 SAP SE or an SAP affiliate company. All rights reserved. This file is licensed under the Apache Software License, v. 2 except as noted otherwise in the LICENSE file

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

     http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by lister-gen. DO NOT EDIT.

package v1alpha1

import (
	v1alpha1 "github.com/gardener/external-dns-management/pkg/apis/dns/v1alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// DNSOwnerLister helps list DNSOwners.
type DNSOwnerLister interface {
	// List lists all DNSOwners in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.DNSOwner, err error)
	// DNSOwners returns an object that can list and get DNSOwners.
	DNSOwners(namespace string) DNSOwnerNamespaceLister
	DNSOwnerListerExpansion
}

// dNSOwnerLister implements the DNSOwnerLister interface.
type dNSOwnerLister struct {
	indexer cache.Indexer
}

// NewDNSOwnerLister returns a new DNSOwnerLister.
func NewDNSOwnerLister(indexer cache.Indexer) DNSOwnerLister {
	return &dNSOwnerLister{indexer: indexer}
}

// List lists all DNSOwners in the indexer.
func (s *dNSOwnerLister) List(selector labels.Selector) (ret []*v1alpha1.DNSOwner, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.DNSOwner))
	})
	return ret, err
}

// DNSOwners returns an object that can list and get DNSOwners.
func (s *dNSOwnerLister) DNSOwners(namespace string) DNSOwnerNamespaceLister {
	return dNSOwnerNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// DNSOwnerNamespaceLister helps list and get DNSOwners.
type DNSOwnerNamespaceLister interface {
	// List lists all DNSOwners in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.DNSOwner, err error)
	// Get retrieves the DNSOwner from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.DNSOwner, error)
	DNSOwnerNamespaceListerExpansion
}

// dNSOwnerNamespaceLister implements the DNSOwnerNamespaceLister
// interface.
type dNSOwnerNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all DNSOwners in the indexer for a given namespace.
func (s dNSOwnerNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.DNSOwner, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.DNSOwner))
	})
	return ret, err
}

// Get retrieves the DNSOwner from the indexer for a given namespace and name.
func (s dNSOwnerNamespaceLister) Get(name string) (*v1alpha1.DNSOwner, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("dnsowner"), name)
	}
	return obj.(*v1alpha1.DNSOwner), nil
}
