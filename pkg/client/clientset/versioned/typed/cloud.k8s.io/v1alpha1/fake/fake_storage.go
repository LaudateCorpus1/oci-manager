/*
Copyright 2018 Oracle and/or its affiliates. All rights reserved.

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
package fake

import (
	v1alpha1 "github.com/oracle/oci-manager/pkg/apis/cloud.k8s.io/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeStorages implements StorageInterface
type FakeStorages struct {
	Fake *FakeCloudV1alpha1
	ns   string
}

var storagesResource = schema.GroupVersionResource{Group: "cloud.k8s.io", Version: "v1alpha1", Resource: "storages"}

var storagesKind = schema.GroupVersionKind{Group: "cloud.k8s.io", Version: "v1alpha1", Kind: "Storage"}

// Get takes name of the storage, and returns the corresponding storage object, and an error if there is any.
func (c *FakeStorages) Get(name string, options v1.GetOptions) (result *v1alpha1.Storage, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(storagesResource, c.ns, name), &v1alpha1.Storage{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Storage), err
}

// List takes label and field selectors, and returns the list of Storages that match those selectors.
func (c *FakeStorages) List(opts v1.ListOptions) (result *v1alpha1.StorageList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(storagesResource, storagesKind, c.ns, opts), &v1alpha1.StorageList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.StorageList{ListMeta: obj.(*v1alpha1.StorageList).ListMeta}
	for _, item := range obj.(*v1alpha1.StorageList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested storages.
func (c *FakeStorages) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(storagesResource, c.ns, opts))

}

// Create takes the representation of a storage and creates it.  Returns the server's representation of the storage, and an error, if there is any.
func (c *FakeStorages) Create(storage *v1alpha1.Storage) (result *v1alpha1.Storage, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(storagesResource, c.ns, storage), &v1alpha1.Storage{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Storage), err
}

// Update takes the representation of a storage and updates it. Returns the server's representation of the storage, and an error, if there is any.
func (c *FakeStorages) Update(storage *v1alpha1.Storage) (result *v1alpha1.Storage, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(storagesResource, c.ns, storage), &v1alpha1.Storage{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Storage), err
}

// Delete takes name of the storage and deletes it. Returns an error if one occurs.
func (c *FakeStorages) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(storagesResource, c.ns, name), &v1alpha1.Storage{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeStorages) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(storagesResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.StorageList{})
	return err
}

// Patch applies the patch and returns the patched storage.
func (c *FakeStorages) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.Storage, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(storagesResource, c.ns, name, data, subresources...), &v1alpha1.Storage{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Storage), err
}
