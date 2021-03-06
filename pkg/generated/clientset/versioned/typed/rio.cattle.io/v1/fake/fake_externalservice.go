/*
Copyright 2019 Rancher Labs.

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

// Code generated by main. DO NOT EDIT.

package fake

import (
	riocattleiov1 "github.com/rancher/rio/pkg/apis/rio.cattle.io/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeExternalServices implements ExternalServiceInterface
type FakeExternalServices struct {
	Fake *FakeRioV1
	ns   string
}

var externalservicesResource = schema.GroupVersionResource{Group: "rio.cattle.io", Version: "v1", Resource: "externalservices"}

var externalservicesKind = schema.GroupVersionKind{Group: "rio.cattle.io", Version: "v1", Kind: "ExternalService"}

// Get takes name of the externalService, and returns the corresponding externalService object, and an error if there is any.
func (c *FakeExternalServices) Get(name string, options v1.GetOptions) (result *riocattleiov1.ExternalService, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(externalservicesResource, c.ns, name), &riocattleiov1.ExternalService{})

	if obj == nil {
		return nil, err
	}
	return obj.(*riocattleiov1.ExternalService), err
}

// List takes label and field selectors, and returns the list of ExternalServices that match those selectors.
func (c *FakeExternalServices) List(opts v1.ListOptions) (result *riocattleiov1.ExternalServiceList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(externalservicesResource, externalservicesKind, c.ns, opts), &riocattleiov1.ExternalServiceList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &riocattleiov1.ExternalServiceList{ListMeta: obj.(*riocattleiov1.ExternalServiceList).ListMeta}
	for _, item := range obj.(*riocattleiov1.ExternalServiceList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested externalServices.
func (c *FakeExternalServices) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(externalservicesResource, c.ns, opts))

}

// Create takes the representation of a externalService and creates it.  Returns the server's representation of the externalService, and an error, if there is any.
func (c *FakeExternalServices) Create(externalService *riocattleiov1.ExternalService) (result *riocattleiov1.ExternalService, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(externalservicesResource, c.ns, externalService), &riocattleiov1.ExternalService{})

	if obj == nil {
		return nil, err
	}
	return obj.(*riocattleiov1.ExternalService), err
}

// Update takes the representation of a externalService and updates it. Returns the server's representation of the externalService, and an error, if there is any.
func (c *FakeExternalServices) Update(externalService *riocattleiov1.ExternalService) (result *riocattleiov1.ExternalService, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(externalservicesResource, c.ns, externalService), &riocattleiov1.ExternalService{})

	if obj == nil {
		return nil, err
	}
	return obj.(*riocattleiov1.ExternalService), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeExternalServices) UpdateStatus(externalService *riocattleiov1.ExternalService) (*riocattleiov1.ExternalService, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(externalservicesResource, "status", c.ns, externalService), &riocattleiov1.ExternalService{})

	if obj == nil {
		return nil, err
	}
	return obj.(*riocattleiov1.ExternalService), err
}

// Delete takes name of the externalService and deletes it. Returns an error if one occurs.
func (c *FakeExternalServices) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(externalservicesResource, c.ns, name), &riocattleiov1.ExternalService{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeExternalServices) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(externalservicesResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &riocattleiov1.ExternalServiceList{})
	return err
}

// Patch applies the patch and returns the patched externalService.
func (c *FakeExternalServices) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *riocattleiov1.ExternalService, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(externalservicesResource, c.ns, name, pt, data, subresources...), &riocattleiov1.ExternalService{})

	if obj == nil {
		return nil, err
	}
	return obj.(*riocattleiov1.ExternalService), err
}
