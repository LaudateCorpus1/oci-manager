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
package v1alpha1

import (
	ocilb_oracle_com_v1alpha1 "github.com/oracle/oci-manager/pkg/apis/ocilb.oracle.com/v1alpha1"
	versioned "github.com/oracle/oci-manager/pkg/client/clientset/versioned"
	internalinterfaces "github.com/oracle/oci-manager/pkg/client/informers/externalversions/internalinterfaces"
	v1alpha1 "github.com/oracle/oci-manager/pkg/client/listers/ocilb.oracle.com/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
	time "time"
)

// BackendSetInformer provides access to a shared informer and lister for
// BackendSets.
type BackendSetInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.BackendSetLister
}

type backendSetInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewBackendSetInformer constructs a new informer for BackendSet type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewBackendSetInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredBackendSetInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredBackendSetInformer constructs a new informer for BackendSet type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredBackendSetInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.OcilbV1alpha1().BackendSets(namespace).List(options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.OcilbV1alpha1().BackendSets(namespace).Watch(options)
			},
		},
		&ocilb_oracle_com_v1alpha1.BackendSet{},
		resyncPeriod,
		indexers,
	)
}

func (f *backendSetInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredBackendSetInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *backendSetInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&ocilb_oracle_com_v1alpha1.BackendSet{}, f.defaultInformer)
}

func (f *backendSetInformer) Lister() v1alpha1.BackendSetLister {
	return v1alpha1.NewBackendSetLister(f.Informer().GetIndexer())
}
