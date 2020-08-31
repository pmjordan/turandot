// Code generated by lister-gen. DO NOT EDIT.

package v1alpha1

import (
	v1alpha1 "github.com/tliron/turandot/resources/turandot.puccini.cloud/v1alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// ServiceLister helps list Services.
// All objects returned here must be treated as read-only.
type ServiceLister interface {
	// List lists all Services in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.Service, err error)
	// Services returns an object that can list and get Services.
	Services(namespace string) ServiceNamespaceLister
	ServiceListerExpansion
}

// serviceLister implements the ServiceLister interface.
type serviceLister struct {
	indexer cache.Indexer
}

// NewServiceLister returns a new ServiceLister.
func NewServiceLister(indexer cache.Indexer) ServiceLister {
	return &serviceLister{indexer: indexer}
}

// List lists all Services in the indexer.
func (s *serviceLister) List(selector labels.Selector) (ret []*v1alpha1.Service, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.Service))
	})
	return ret, err
}

// Services returns an object that can list and get Services.
func (s *serviceLister) Services(namespace string) ServiceNamespaceLister {
	return serviceNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// ServiceNamespaceLister helps list and get Services.
// All objects returned here must be treated as read-only.
type ServiceNamespaceLister interface {
	// List lists all Services in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.Service, err error)
	// Get retrieves the Service from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.Service, error)
	ServiceNamespaceListerExpansion
}

// serviceNamespaceLister implements the ServiceNamespaceLister
// interface.
type serviceNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all Services in the indexer for a given namespace.
func (s serviceNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.Service, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.Service))
	})
	return ret, err
}

// Get retrieves the Service from the indexer for a given namespace and name.
func (s serviceNamespaceLister) Get(name string) (*v1alpha1.Service, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("service"), name)
	}
	return obj.(*v1alpha1.Service), nil
}
