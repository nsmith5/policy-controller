// Copyright 2022 The Sigstore Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by informer-gen. DO NOT EDIT.

package v1alpha1

import (
	"context"
	time "time"

	policyv1alpha1 "github.com/sigstore/policy-controller/pkg/apis/policy/v1alpha1"
	versioned "github.com/sigstore/policy-controller/pkg/client/clientset/versioned"
	internalinterfaces "github.com/sigstore/policy-controller/pkg/client/informers/externalversions/internalinterfaces"
	v1alpha1 "github.com/sigstore/policy-controller/pkg/client/listers/policy/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// ClusterImagePolicyInformer provides access to a shared informer and lister for
// ClusterImagePolicies.
type ClusterImagePolicyInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.ClusterImagePolicyLister
}

type clusterImagePolicyInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// NewClusterImagePolicyInformer constructs a new informer for ClusterImagePolicy type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewClusterImagePolicyInformer(client versioned.Interface, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredClusterImagePolicyInformer(client, resyncPeriod, indexers, nil)
}

// NewFilteredClusterImagePolicyInformer constructs a new informer for ClusterImagePolicy type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredClusterImagePolicyInformer(client versioned.Interface, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.PolicyV1alpha1().ClusterImagePolicies().List(context.TODO(), options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.PolicyV1alpha1().ClusterImagePolicies().Watch(context.TODO(), options)
			},
		},
		&policyv1alpha1.ClusterImagePolicy{},
		resyncPeriod,
		indexers,
	)
}

func (f *clusterImagePolicyInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredClusterImagePolicyInformer(client, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *clusterImagePolicyInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&policyv1alpha1.ClusterImagePolicy{}, f.defaultInformer)
}

func (f *clusterImagePolicyInformer) Lister() v1alpha1.ClusterImagePolicyLister {
	return v1alpha1.NewClusterImagePolicyLister(f.Informer().GetIndexer())
}
