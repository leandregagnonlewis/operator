/*


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
// Code generated by client-gen-v0.30. DO NOT EDIT.

package v1beta1

import (
	"context"
	"time"

	scheme "github.com/VictoriaMetrics/operator/api/client/versioned/scheme"
	v1beta1 "github.com/VictoriaMetrics/operator/api/operator/v1beta1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// VMStaticScrapesGetter has a method to return a VMStaticScrapeInterface.
// A group's client should implement this interface.
type VMStaticScrapesGetter interface {
	VMStaticScrapes(namespace string) VMStaticScrapeInterface
}

// VMStaticScrapeInterface has methods to work with VMStaticScrape resources.
type VMStaticScrapeInterface interface {
	Create(ctx context.Context, vMStaticScrape *v1beta1.VMStaticScrape, opts v1.CreateOptions) (*v1beta1.VMStaticScrape, error)
	Update(ctx context.Context, vMStaticScrape *v1beta1.VMStaticScrape, opts v1.UpdateOptions) (*v1beta1.VMStaticScrape, error)
	UpdateStatus(ctx context.Context, vMStaticScrape *v1beta1.VMStaticScrape, opts v1.UpdateOptions) (*v1beta1.VMStaticScrape, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1beta1.VMStaticScrape, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1beta1.VMStaticScrapeList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta1.VMStaticScrape, err error)
	VMStaticScrapeExpansion
}

// vMStaticScrapes implements VMStaticScrapeInterface
type vMStaticScrapes struct {
	client rest.Interface
	ns     string
}

// newVMStaticScrapes returns a VMStaticScrapes
func newVMStaticScrapes(c *OperatorV1beta1Client, namespace string) *vMStaticScrapes {
	return &vMStaticScrapes{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the vMStaticScrape, and returns the corresponding vMStaticScrape object, and an error if there is any.
func (c *vMStaticScrapes) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1beta1.VMStaticScrape, err error) {
	result = &v1beta1.VMStaticScrape{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("vmstaticscrapes").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of VMStaticScrapes that match those selectors.
func (c *vMStaticScrapes) List(ctx context.Context, opts v1.ListOptions) (result *v1beta1.VMStaticScrapeList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1beta1.VMStaticScrapeList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("vmstaticscrapes").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested vMStaticScrapes.
func (c *vMStaticScrapes) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("vmstaticscrapes").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a vMStaticScrape and creates it.  Returns the server's representation of the vMStaticScrape, and an error, if there is any.
func (c *vMStaticScrapes) Create(ctx context.Context, vMStaticScrape *v1beta1.VMStaticScrape, opts v1.CreateOptions) (result *v1beta1.VMStaticScrape, err error) {
	result = &v1beta1.VMStaticScrape{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("vmstaticscrapes").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(vMStaticScrape).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a vMStaticScrape and updates it. Returns the server's representation of the vMStaticScrape, and an error, if there is any.
func (c *vMStaticScrapes) Update(ctx context.Context, vMStaticScrape *v1beta1.VMStaticScrape, opts v1.UpdateOptions) (result *v1beta1.VMStaticScrape, err error) {
	result = &v1beta1.VMStaticScrape{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("vmstaticscrapes").
		Name(vMStaticScrape.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(vMStaticScrape).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *vMStaticScrapes) UpdateStatus(ctx context.Context, vMStaticScrape *v1beta1.VMStaticScrape, opts v1.UpdateOptions) (result *v1beta1.VMStaticScrape, err error) {
	result = &v1beta1.VMStaticScrape{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("vmstaticscrapes").
		Name(vMStaticScrape.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(vMStaticScrape).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the vMStaticScrape and deletes it. Returns an error if one occurs.
func (c *vMStaticScrapes) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("vmstaticscrapes").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *vMStaticScrapes) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("vmstaticscrapes").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched vMStaticScrape.
func (c *vMStaticScrapes) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta1.VMStaticScrape, err error) {
	result = &v1beta1.VMStaticScrape{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("vmstaticscrapes").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}