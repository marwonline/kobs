// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	v1 "github.com/kobsio/kobs/pkg/cluster/kubernetes/clients/team/clientset/versioned/typed/team/v1"
	rest "k8s.io/client-go/rest"
	testing "k8s.io/client-go/testing"
)

type FakeKobsV1 struct {
	*testing.Fake
}

func (c *FakeKobsV1) Teams(namespace string) v1.TeamInterface {
	return &FakeTeams{c, namespace}
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *FakeKobsV1) RESTClient() rest.Interface {
	var ret *rest.RESTClient
	return ret
}