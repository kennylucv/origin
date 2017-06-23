// This file was automatically generated by informer-gen

package internalversion

import (
	api "github.com/openshift/origin/pkg/security/api"
	internalinterfaces "github.com/openshift/origin/pkg/security/generated/informers/internalversion/internalinterfaces"
	internalclientset "github.com/openshift/origin/pkg/security/generated/internalclientset"
	internalversion "github.com/openshift/origin/pkg/security/generated/listers/security/internalversion"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
	time "time"
)

// SecurityContextConstraintsInformer provides access to a shared informer and lister for
// SecurityContextConstraints.
type SecurityContextConstraintsInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() internalversion.SecurityContextConstraintsLister
}

type securityContextConstraintsInformer struct {
	factory internalinterfaces.SharedInformerFactory
}

func newSecurityContextConstraintsInformer(client internalclientset.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	sharedIndexInformer := cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				return client.Security().SecurityContextConstraints().List(options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				return client.Security().SecurityContextConstraints().Watch(options)
			},
		},
		&api.SecurityContextConstraints{},
		resyncPeriod,
		cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc},
	)

	return sharedIndexInformer
}

func (f *securityContextConstraintsInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&api.SecurityContextConstraints{}, newSecurityContextConstraintsInformer)
}

func (f *securityContextConstraintsInformer) Lister() internalversion.SecurityContextConstraintsLister {
	return internalversion.NewSecurityContextConstraintsLister(f.Informer().GetIndexer())
}
