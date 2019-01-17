/*
Copyright (c) 2018 SAP SE or an SAP affiliate company. All rights reserved. This file is licensed under the Apache Software License, v. 2 except as noted otherwise in the LICENSE file

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

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	v1alpha1 "github.com/gardener/dns-controller-manager/pkg/apis/dns/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeDNSEntries implements DNSEntryInterface
type FakeDNSEntries struct {
	Fake *FakeKracV1alpha1
	ns   string
}

var dnsentriesResource = schema.GroupVersionResource{Group: "krac", Version: "v1alpha1", Resource: "dnsentries"}

var dnsentriesKind = schema.GroupVersionKind{Group: "krac", Version: "v1alpha1", Kind: "DNSEntry"}

// Get takes name of the dNSEntry, and returns the corresponding dNSEntry object, and an error if there is any.
func (c *FakeDNSEntries) Get(name string, options v1.GetOptions) (result *v1alpha1.DNSEntry, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(dnsentriesResource, c.ns, name), &v1alpha1.DNSEntry{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DNSEntry), err
}

// List takes label and field selectors, and returns the list of DNSEntries that match those selectors.
func (c *FakeDNSEntries) List(opts v1.ListOptions) (result *v1alpha1.DNSEntryList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(dnsentriesResource, dnsentriesKind, c.ns, opts), &v1alpha1.DNSEntryList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.DNSEntryList{ListMeta: obj.(*v1alpha1.DNSEntryList).ListMeta}
	for _, item := range obj.(*v1alpha1.DNSEntryList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested dNSEntries.
func (c *FakeDNSEntries) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(dnsentriesResource, c.ns, opts))

}

// Create takes the representation of a dNSEntry and creates it.  Returns the server's representation of the dNSEntry, and an error, if there is any.
func (c *FakeDNSEntries) Create(dNSEntry *v1alpha1.DNSEntry) (result *v1alpha1.DNSEntry, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(dnsentriesResource, c.ns, dNSEntry), &v1alpha1.DNSEntry{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DNSEntry), err
}

// Update takes the representation of a dNSEntry and updates it. Returns the server's representation of the dNSEntry, and an error, if there is any.
func (c *FakeDNSEntries) Update(dNSEntry *v1alpha1.DNSEntry) (result *v1alpha1.DNSEntry, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(dnsentriesResource, c.ns, dNSEntry), &v1alpha1.DNSEntry{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DNSEntry), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeDNSEntries) UpdateStatus(dNSEntry *v1alpha1.DNSEntry) (*v1alpha1.DNSEntry, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(dnsentriesResource, "status", c.ns, dNSEntry), &v1alpha1.DNSEntry{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DNSEntry), err
}

// Delete takes name of the dNSEntry and deletes it. Returns an error if one occurs.
func (c *FakeDNSEntries) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(dnsentriesResource, c.ns, name), &v1alpha1.DNSEntry{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeDNSEntries) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(dnsentriesResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.DNSEntryList{})
	return err
}

// Patch applies the patch and returns the patched dNSEntry.
func (c *FakeDNSEntries) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.DNSEntry, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(dnsentriesResource, c.ns, name, data, subresources...), &v1alpha1.DNSEntry{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DNSEntry), err
}
