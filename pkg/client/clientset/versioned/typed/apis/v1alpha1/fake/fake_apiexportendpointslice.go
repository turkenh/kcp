/*
Copyright The KCP Authors.

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

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"

	v1alpha1 "github.com/kcp-dev/kcp/pkg/apis/apis/v1alpha1"
)

// FakeAPIExportEndpointSlices implements APIExportEndpointSliceInterface
type FakeAPIExportEndpointSlices struct {
	Fake *FakeApisV1alpha1
}

var apiexportendpointslicesResource = schema.GroupVersionResource{Group: "apis.kcp.dev", Version: "v1alpha1", Resource: "apiexportendpointslices"}

var apiexportendpointslicesKind = schema.GroupVersionKind{Group: "apis.kcp.dev", Version: "v1alpha1", Kind: "APIExportEndpointSlice"}

// Get takes name of the aPIExportEndpointSlice, and returns the corresponding aPIExportEndpointSlice object, and an error if there is any.
func (c *FakeAPIExportEndpointSlices) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.APIExportEndpointSlice, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(apiexportendpointslicesResource, name), &v1alpha1.APIExportEndpointSlice{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.APIExportEndpointSlice), err
}

// List takes label and field selectors, and returns the list of APIExportEndpointSlices that match those selectors.
func (c *FakeAPIExportEndpointSlices) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.APIExportEndpointSliceList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(apiexportendpointslicesResource, apiexportendpointslicesKind, opts), &v1alpha1.APIExportEndpointSliceList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.APIExportEndpointSliceList{ListMeta: obj.(*v1alpha1.APIExportEndpointSliceList).ListMeta}
	for _, item := range obj.(*v1alpha1.APIExportEndpointSliceList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested aPIExportEndpointSlices.
func (c *FakeAPIExportEndpointSlices) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(apiexportendpointslicesResource, opts))
}

// Create takes the representation of a aPIExportEndpointSlice and creates it.  Returns the server's representation of the aPIExportEndpointSlice, and an error, if there is any.
func (c *FakeAPIExportEndpointSlices) Create(ctx context.Context, aPIExportEndpointSlice *v1alpha1.APIExportEndpointSlice, opts v1.CreateOptions) (result *v1alpha1.APIExportEndpointSlice, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(apiexportendpointslicesResource, aPIExportEndpointSlice), &v1alpha1.APIExportEndpointSlice{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.APIExportEndpointSlice), err
}

// Update takes the representation of a aPIExportEndpointSlice and updates it. Returns the server's representation of the aPIExportEndpointSlice, and an error, if there is any.
func (c *FakeAPIExportEndpointSlices) Update(ctx context.Context, aPIExportEndpointSlice *v1alpha1.APIExportEndpointSlice, opts v1.UpdateOptions) (result *v1alpha1.APIExportEndpointSlice, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(apiexportendpointslicesResource, aPIExportEndpointSlice), &v1alpha1.APIExportEndpointSlice{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.APIExportEndpointSlice), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeAPIExportEndpointSlices) UpdateStatus(ctx context.Context, aPIExportEndpointSlice *v1alpha1.APIExportEndpointSlice, opts v1.UpdateOptions) (*v1alpha1.APIExportEndpointSlice, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(apiexportendpointslicesResource, "status", aPIExportEndpointSlice), &v1alpha1.APIExportEndpointSlice{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.APIExportEndpointSlice), err
}

// Delete takes name of the aPIExportEndpointSlice and deletes it. Returns an error if one occurs.
func (c *FakeAPIExportEndpointSlices) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteActionWithOptions(apiexportendpointslicesResource, name, opts), &v1alpha1.APIExportEndpointSlice{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeAPIExportEndpointSlices) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(apiexportendpointslicesResource, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.APIExportEndpointSliceList{})
	return err
}

// Patch applies the patch and returns the patched aPIExportEndpointSlice.
func (c *FakeAPIExportEndpointSlices) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.APIExportEndpointSlice, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(apiexportendpointslicesResource, name, pt, data, subresources...), &v1alpha1.APIExportEndpointSlice{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.APIExportEndpointSlice), err
}
