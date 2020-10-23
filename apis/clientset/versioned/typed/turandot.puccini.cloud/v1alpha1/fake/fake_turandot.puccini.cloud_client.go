// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	v1alpha1 "github.com/tliron/turandot/apis/clientset/versioned/typed/turandot.puccini.cloud/v1alpha1"
	rest "k8s.io/client-go/rest"
	testing "k8s.io/client-go/testing"
)

type FakeTurandotV1alpha1 struct {
	*testing.Fake
}

func (c *FakeTurandotV1alpha1) Repositories(namespace string) v1alpha1.RepositoryInterface {
	return &FakeRepositories{c, namespace}
}

func (c *FakeTurandotV1alpha1) Services(namespace string) v1alpha1.ServiceInterface {
	return &FakeServices{c, namespace}
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *FakeTurandotV1alpha1) RESTClient() rest.Interface {
	var ret *rest.RESTClient
	return ret
}
