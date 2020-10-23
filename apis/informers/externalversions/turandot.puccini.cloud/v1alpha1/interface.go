// Code generated by informer-gen. DO NOT EDIT.

package v1alpha1

import (
	internalinterfaces "github.com/tliron/turandot/apis/informers/externalversions/internalinterfaces"
)

// Interface provides access to all the informers in this group version.
type Interface interface {
	// Repositories returns a RepositoryInformer.
	Repositories() RepositoryInformer
	// Services returns a ServiceInformer.
	Services() ServiceInformer
}

type version struct {
	factory          internalinterfaces.SharedInformerFactory
	namespace        string
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// New returns a new Interface.
func New(f internalinterfaces.SharedInformerFactory, namespace string, tweakListOptions internalinterfaces.TweakListOptionsFunc) Interface {
	return &version{factory: f, namespace: namespace, tweakListOptions: tweakListOptions}
}

// Repositories returns a RepositoryInformer.
func (v *version) Repositories() RepositoryInformer {
	return &repositoryInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// Services returns a ServiceInformer.
func (v *version) Services() ServiceInformer {
	return &serviceInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}
