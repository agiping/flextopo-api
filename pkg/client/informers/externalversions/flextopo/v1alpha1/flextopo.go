/*
Copyright 2022 The Katalyst Authors.

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

// Code generated by informer-gen. DO NOT EDIT.

package v1alpha1

import (
	"context"
	time "time"

	flextopov1alpha1 "github.com/agiping/flextopo-api/pkg/apis/flextopo/v1alpha1"
	versioned "github.com/agiping/flextopo-api/pkg/client/clientset/versioned"
	internalinterfaces "github.com/agiping/flextopo-api/pkg/client/informers/externalversions/internalinterfaces"
	v1alpha1 "github.com/agiping/flextopo-api/pkg/client/listers/flextopo/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// FlexTopoInformer provides access to a shared informer and lister for
// FlexTopos.
type FlexTopoInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.FlexTopoLister
}

type flexTopoInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// NewFlexTopoInformer constructs a new informer for FlexTopo type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFlexTopoInformer(client versioned.Interface, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredFlexTopoInformer(client, resyncPeriod, indexers, nil)
}

// NewFilteredFlexTopoInformer constructs a new informer for FlexTopo type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredFlexTopoInformer(client versioned.Interface, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.FlextopoV1alpha1().FlexTopos().List(context.TODO(), options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.FlextopoV1alpha1().FlexTopos().Watch(context.TODO(), options)
			},
		},
		&flextopov1alpha1.FlexTopo{},
		resyncPeriod,
		indexers,
	)
}

func (f *flexTopoInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredFlexTopoInformer(client, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *flexTopoInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&flextopov1alpha1.FlexTopo{}, f.defaultInformer)
}

func (f *flexTopoInformer) Lister() v1alpha1.FlexTopoLister {
	return v1alpha1.NewFlexTopoLister(f.Informer().GetIndexer())
}
