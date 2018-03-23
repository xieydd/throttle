/*
Copyright The Kubernetes Authors.

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

package v1alpha1

import (
	v1alpha1 "github.com/xychu/throttle/pkg/apis/throttlecontroller/v1alpha1"
	scheme "github.com/xychu/throttle/pkg/client/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// GPUQuotasGetter has a method to return a GPUQuotaInterface.
// A group's client should implement this interface.
type GPUQuotasGetter interface {
	GPUQuotas(namespace string) GPUQuotaInterface
}

// GPUQuotaInterface has methods to work with GPUQuota resources.
type GPUQuotaInterface interface {
	Create(*v1alpha1.GPUQuota) (*v1alpha1.GPUQuota, error)
	Update(*v1alpha1.GPUQuota) (*v1alpha1.GPUQuota, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.GPUQuota, error)
	List(opts v1.ListOptions) (*v1alpha1.GPUQuotaList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.GPUQuota, err error)
	GPUQuotaExpansion
}

// gPUQuotas implements GPUQuotaInterface
type gPUQuotas struct {
	client rest.Interface
	ns     string
}

// newGPUQuotas returns a GPUQuotas
func newGPUQuotas(c *ThrottlecontrollerV1alpha1Client, namespace string) *gPUQuotas {
	return &gPUQuotas{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the gPUQuota, and returns the corresponding gPUQuota object, and an error if there is any.
func (c *gPUQuotas) Get(name string, options v1.GetOptions) (result *v1alpha1.GPUQuota, err error) {
	result = &v1alpha1.GPUQuota{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("gpuquotas").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of GPUQuotas that match those selectors.
func (c *gPUQuotas) List(opts v1.ListOptions) (result *v1alpha1.GPUQuotaList, err error) {
	result = &v1alpha1.GPUQuotaList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("gpuquotas").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested gPUQuotas.
func (c *gPUQuotas) Watch(opts v1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("gpuquotas").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch()
}

// Create takes the representation of a gPUQuota and creates it.  Returns the server's representation of the gPUQuota, and an error, if there is any.
func (c *gPUQuotas) Create(gPUQuota *v1alpha1.GPUQuota) (result *v1alpha1.GPUQuota, err error) {
	result = &v1alpha1.GPUQuota{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("gpuquotas").
		Body(gPUQuota).
		Do().
		Into(result)
	return
}

// Update takes the representation of a gPUQuota and updates it. Returns the server's representation of the gPUQuota, and an error, if there is any.
func (c *gPUQuotas) Update(gPUQuota *v1alpha1.GPUQuota) (result *v1alpha1.GPUQuota, err error) {
	result = &v1alpha1.GPUQuota{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("gpuquotas").
		Name(gPUQuota.Name).
		Body(gPUQuota).
		Do().
		Into(result)
	return
}

// Delete takes name of the gPUQuota and deletes it. Returns an error if one occurs.
func (c *gPUQuotas) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("gpuquotas").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *gPUQuotas) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("gpuquotas").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched gPUQuota.
func (c *gPUQuotas) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.GPUQuota, err error) {
	result = &v1alpha1.GPUQuota{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("gpuquotas").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
