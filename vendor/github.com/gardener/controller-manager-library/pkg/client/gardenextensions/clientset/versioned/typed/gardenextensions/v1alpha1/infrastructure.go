// Code generated by client-gen. DO NOT EDIT.

package v1alpha1

import (
	v1alpha1 "github.com/gardener/controller-manager-library/pkg/apis/gardenextensions/v1alpha1"
	scheme "github.com/gardener/controller-manager-library/pkg/client/gardenextensions/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// InfrastructuresGetter has a method to return a InfrastructureInterface.
// A group's client should implement this interface.
type InfrastructuresGetter interface {
	Infrastructures(namespace string) InfrastructureInterface
}

// InfrastructureInterface has methods to work with Infrastructure resources.
type InfrastructureInterface interface {
	Create(*v1alpha1.Infrastructure) (*v1alpha1.Infrastructure, error)
	Update(*v1alpha1.Infrastructure) (*v1alpha1.Infrastructure, error)
	UpdateStatus(*v1alpha1.Infrastructure) (*v1alpha1.Infrastructure, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.Infrastructure, error)
	List(opts v1.ListOptions) (*v1alpha1.InfrastructureList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.Infrastructure, err error)
	InfrastructureExpansion
}

// infrastructures implements InfrastructureInterface
type infrastructures struct {
	client rest.Interface
	ns     string
}

// newInfrastructures returns a Infrastructures
func newInfrastructures(c *GardenextensionsV1alpha1Client, namespace string) *infrastructures {
	return &infrastructures{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Data takes name of the infrastructure, and returns the corresponding infrastructure object, and an error if there is any.
func (c *infrastructures) Get(name string, options v1.GetOptions) (result *v1alpha1.Infrastructure, err error) {
	result = &v1alpha1.Infrastructure{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("infrastructures").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of Infrastructures that match those selectors.
func (c *infrastructures) List(opts v1.ListOptions) (result *v1alpha1.InfrastructureList, err error) {
	result = &v1alpha1.InfrastructureList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("infrastructures").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested infrastructures.
func (c *infrastructures) Watch(opts v1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("infrastructures").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch()
}

// Create takes the representation of a infrastructure and creates it.  Returns the server's representation of the infrastructure, and an error, if there is any.
func (c *infrastructures) Create(infrastructure *v1alpha1.Infrastructure) (result *v1alpha1.Infrastructure, err error) {
	result = &v1alpha1.Infrastructure{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("infrastructures").
		Body(infrastructure).
		Do().
		Into(result)
	return
}

// Update takes the representation of a infrastructure and updates it. Returns the server's representation of the infrastructure, and an error, if there is any.
func (c *infrastructures) Update(infrastructure *v1alpha1.Infrastructure) (result *v1alpha1.Infrastructure, err error) {
	result = &v1alpha1.Infrastructure{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("infrastructures").
		Name(infrastructure.Name).
		Body(infrastructure).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *infrastructures) UpdateStatus(infrastructure *v1alpha1.Infrastructure) (result *v1alpha1.Infrastructure, err error) {
	result = &v1alpha1.Infrastructure{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("infrastructures").
		Name(infrastructure.Name).
		SubResource("status").
		Body(infrastructure).
		Do().
		Into(result)
	return
}

// Delete takes name of the infrastructure and deletes it. Returns an error if one occurs.
func (c *infrastructures) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("infrastructures").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *infrastructures) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("infrastructures").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched infrastructure.
func (c *infrastructures) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.Infrastructure, err error) {
	result = &v1alpha1.Infrastructure{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("infrastructures").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
