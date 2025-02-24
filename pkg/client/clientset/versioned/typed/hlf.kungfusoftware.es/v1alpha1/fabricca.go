/*
 * Copyright Kungfusoftware.es. All Rights Reserved.
 *
 * SPDX-License-Identifier: Apache-2.0
 */

package v1alpha1

import (
	"context"
	json "encoding/json"
	"fmt"
	"time"

	v1alpha1 "github.com/kfsoftware/hlf-operator/api/hlf.kungfusoftware.es/v1alpha1"
	hlfkungfusoftwareesv1alpha1 "github.com/kfsoftware/hlf-operator/pkg/client/applyconfiguration/hlf.kungfusoftware.es/v1alpha1"
	scheme "github.com/kfsoftware/hlf-operator/pkg/client/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// FabricCAsGetter has a method to return a FabricCAInterface.
// A group's client should implement this interface.
type FabricCAsGetter interface {
	FabricCAs(namespace string) FabricCAInterface
}

// FabricCAInterface has methods to work with FabricCA resources.
type FabricCAInterface interface {
	Create(ctx context.Context, fabricCA *v1alpha1.FabricCA, opts v1.CreateOptions) (*v1alpha1.FabricCA, error)
	Update(ctx context.Context, fabricCA *v1alpha1.FabricCA, opts v1.UpdateOptions) (*v1alpha1.FabricCA, error)
	UpdateStatus(ctx context.Context, fabricCA *v1alpha1.FabricCA, opts v1.UpdateOptions) (*v1alpha1.FabricCA, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1alpha1.FabricCA, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1alpha1.FabricCAList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.FabricCA, err error)
	Apply(ctx context.Context, fabricCA *hlfkungfusoftwareesv1alpha1.FabricCAApplyConfiguration, opts v1.ApplyOptions) (result *v1alpha1.FabricCA, err error)
	ApplyStatus(ctx context.Context, fabricCA *hlfkungfusoftwareesv1alpha1.FabricCAApplyConfiguration, opts v1.ApplyOptions) (result *v1alpha1.FabricCA, err error)
	FabricCAExpansion
}

// fabricCAs implements FabricCAInterface
type fabricCAs struct {
	client rest.Interface
	ns     string
}

// newFabricCAs returns a FabricCAs
func newFabricCAs(c *HlfV1alpha1Client, namespace string) *fabricCAs {
	return &fabricCAs{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the fabricCA, and returns the corresponding fabricCA object, and an error if there is any.
func (c *fabricCAs) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.FabricCA, err error) {
	result = &v1alpha1.FabricCA{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("fabriccas").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of FabricCAs that match those selectors.
func (c *fabricCAs) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.FabricCAList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.FabricCAList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("fabriccas").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested fabricCAs.
func (c *fabricCAs) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("fabriccas").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a fabricCA and creates it.  Returns the server's representation of the fabricCA, and an error, if there is any.
func (c *fabricCAs) Create(ctx context.Context, fabricCA *v1alpha1.FabricCA, opts v1.CreateOptions) (result *v1alpha1.FabricCA, err error) {
	result = &v1alpha1.FabricCA{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("fabriccas").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(fabricCA).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a fabricCA and updates it. Returns the server's representation of the fabricCA, and an error, if there is any.
func (c *fabricCAs) Update(ctx context.Context, fabricCA *v1alpha1.FabricCA, opts v1.UpdateOptions) (result *v1alpha1.FabricCA, err error) {
	result = &v1alpha1.FabricCA{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("fabriccas").
		Name(fabricCA.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(fabricCA).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *fabricCAs) UpdateStatus(ctx context.Context, fabricCA *v1alpha1.FabricCA, opts v1.UpdateOptions) (result *v1alpha1.FabricCA, err error) {
	result = &v1alpha1.FabricCA{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("fabriccas").
		Name(fabricCA.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(fabricCA).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the fabricCA and deletes it. Returns an error if one occurs.
func (c *fabricCAs) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("fabriccas").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *fabricCAs) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("fabriccas").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched fabricCA.
func (c *fabricCAs) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.FabricCA, err error) {
	result = &v1alpha1.FabricCA{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("fabriccas").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}

// Apply takes the given apply declarative configuration, applies it and returns the applied fabricCA.
func (c *fabricCAs) Apply(ctx context.Context, fabricCA *hlfkungfusoftwareesv1alpha1.FabricCAApplyConfiguration, opts v1.ApplyOptions) (result *v1alpha1.FabricCA, err error) {
	if fabricCA == nil {
		return nil, fmt.Errorf("fabricCA provided to Apply must not be nil")
	}
	patchOpts := opts.ToPatchOptions()
	data, err := json.Marshal(fabricCA)
	if err != nil {
		return nil, err
	}
	name := fabricCA.Name
	if name == nil {
		return nil, fmt.Errorf("fabricCA.Name must be provided to Apply")
	}
	result = &v1alpha1.FabricCA{}
	err = c.client.Patch(types.ApplyPatchType).
		Namespace(c.ns).
		Resource("fabriccas").
		Name(*name).
		VersionedParams(&patchOpts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}

// ApplyStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating ApplyStatus().
func (c *fabricCAs) ApplyStatus(ctx context.Context, fabricCA *hlfkungfusoftwareesv1alpha1.FabricCAApplyConfiguration, opts v1.ApplyOptions) (result *v1alpha1.FabricCA, err error) {
	if fabricCA == nil {
		return nil, fmt.Errorf("fabricCA provided to Apply must not be nil")
	}
	patchOpts := opts.ToPatchOptions()
	data, err := json.Marshal(fabricCA)
	if err != nil {
		return nil, err
	}

	name := fabricCA.Name
	if name == nil {
		return nil, fmt.Errorf("fabricCA.Name must be provided to Apply")
	}

	result = &v1alpha1.FabricCA{}
	err = c.client.Patch(types.ApplyPatchType).
		Namespace(c.ns).
		Resource("fabriccas").
		Name(*name).
		SubResource("status").
		VersionedParams(&patchOpts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
