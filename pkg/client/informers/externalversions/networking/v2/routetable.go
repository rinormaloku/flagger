/*
Copyright 2020 The Flux authors

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

package v2

import (
	"context"
	time "time"

	networkingv2 "github.com/fluxcd/flagger/pkg/apis/gloo/networking/v2"
	versioned "github.com/fluxcd/flagger/pkg/client/clientset/versioned"
	internalinterfaces "github.com/fluxcd/flagger/pkg/client/informers/externalversions/internalinterfaces"
	v2 "github.com/fluxcd/flagger/pkg/client/listers/networking/v2"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// RouteTableInformer provides access to a shared informer and lister for
// RouteTables.
type RouteTableInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v2.RouteTableLister
}

type routeTableInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewRouteTableInformer constructs a new informer for RouteTable type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewRouteTableInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredRouteTableInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredRouteTableInformer constructs a new informer for RouteTable type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredRouteTableInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.GloomeshnetworkingV2().RouteTables(namespace).List(context.TODO(), options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.GloomeshnetworkingV2().RouteTables(namespace).Watch(context.TODO(), options)
			},
		},
		&networkingv2.RouteTable{},
		resyncPeriod,
		indexers,
	)
}

func (f *routeTableInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredRouteTableInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *routeTableInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&networkingv2.RouteTable{}, f.defaultInformer)
}

func (f *routeTableInformer) Lister() v2.RouteTableLister {
	return v2.NewRouteTableLister(f.Informer().GetIndexer())
}
