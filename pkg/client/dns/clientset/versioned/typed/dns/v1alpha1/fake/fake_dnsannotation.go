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

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"

	v1alpha1 "github.com/gardener/external-dns-management/pkg/apis/dns/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeDNSAnnotations implements DNSAnnotationInterface
type FakeDNSAnnotations struct {
	Fake *FakeDnsV1alpha1
	ns   string
}

var dnsannotationsResource = schema.GroupVersionResource{Group: "dns.gardener.cloud", Version: "v1alpha1", Resource: "dnsannotations"}

var dnsannotationsKind = schema.GroupVersionKind{Group: "dns.gardener.cloud", Version: "v1alpha1", Kind: "DNSAnnotation"}

// Get takes name of the dNSAnnotation, and returns the corresponding dNSAnnotation object, and an error if there is any.
func (c *FakeDNSAnnotations) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.DNSAnnotation, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(dnsannotationsResource, c.ns, name), &v1alpha1.DNSAnnotation{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DNSAnnotation), err
}

// List takes label and field selectors, and returns the list of DNSAnnotations that match those selectors.
func (c *FakeDNSAnnotations) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.DNSAnnotationList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(dnsannotationsResource, dnsannotationsKind, c.ns, opts), &v1alpha1.DNSAnnotationList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.DNSAnnotationList{ListMeta: obj.(*v1alpha1.DNSAnnotationList).ListMeta}
	for _, item := range obj.(*v1alpha1.DNSAnnotationList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested dNSAnnotations.
func (c *FakeDNSAnnotations) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(dnsannotationsResource, c.ns, opts))

}

// Create takes the representation of a dNSAnnotation and creates it.  Returns the server's representation of the dNSAnnotation, and an error, if there is any.
func (c *FakeDNSAnnotations) Create(ctx context.Context, dNSAnnotation *v1alpha1.DNSAnnotation, opts v1.CreateOptions) (result *v1alpha1.DNSAnnotation, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(dnsannotationsResource, c.ns, dNSAnnotation), &v1alpha1.DNSAnnotation{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DNSAnnotation), err
}

// Update takes the representation of a dNSAnnotation and updates it. Returns the server's representation of the dNSAnnotation, and an error, if there is any.
func (c *FakeDNSAnnotations) Update(ctx context.Context, dNSAnnotation *v1alpha1.DNSAnnotation, opts v1.UpdateOptions) (result *v1alpha1.DNSAnnotation, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(dnsannotationsResource, c.ns, dNSAnnotation), &v1alpha1.DNSAnnotation{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DNSAnnotation), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeDNSAnnotations) UpdateStatus(ctx context.Context, dNSAnnotation *v1alpha1.DNSAnnotation, opts v1.UpdateOptions) (*v1alpha1.DNSAnnotation, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(dnsannotationsResource, "status", c.ns, dNSAnnotation), &v1alpha1.DNSAnnotation{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DNSAnnotation), err
}

// Delete takes name of the dNSAnnotation and deletes it. Returns an error if one occurs.
func (c *FakeDNSAnnotations) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(dnsannotationsResource, c.ns, name), &v1alpha1.DNSAnnotation{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeDNSAnnotations) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(dnsannotationsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.DNSAnnotationList{})
	return err
}

// Patch applies the patch and returns the patched dNSAnnotation.
func (c *FakeDNSAnnotations) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.DNSAnnotation, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(dnsannotationsResource, c.ns, name, pt, data, subresources...), &v1alpha1.DNSAnnotation{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DNSAnnotation), err
}
