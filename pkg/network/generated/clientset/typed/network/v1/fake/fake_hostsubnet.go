package fake

import (
	network_v1 "github.com/openshift/origin/pkg/network/apis/network/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeHostSubnets implements HostSubnetInterface
type FakeHostSubnets struct {
	Fake *FakeNetworkV1
}

var hostsubnetsResource = schema.GroupVersionResource{Group: "network.openshift.io", Version: "v1", Resource: "hostsubnets"}

var hostsubnetsKind = schema.GroupVersionKind{Group: "network.openshift.io", Version: "v1", Kind: "HostSubnet"}

// Get takes name of the hostSubnet, and returns the corresponding hostSubnet object, and an error if there is any.
func (c *FakeHostSubnets) Get(name string, options v1.GetOptions) (result *network_v1.HostSubnet, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(hostsubnetsResource, name), &network_v1.HostSubnet{})
	if obj == nil {
		return nil, err
	}
	return obj.(*network_v1.HostSubnet), err
}

// List takes label and field selectors, and returns the list of HostSubnets that match those selectors.
func (c *FakeHostSubnets) List(opts v1.ListOptions) (result *network_v1.HostSubnetList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(hostsubnetsResource, hostsubnetsKind, opts), &network_v1.HostSubnetList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &network_v1.HostSubnetList{}
	for _, item := range obj.(*network_v1.HostSubnetList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested hostSubnets.
func (c *FakeHostSubnets) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(hostsubnetsResource, opts))
}

// Create takes the representation of a hostSubnet and creates it.  Returns the server's representation of the hostSubnet, and an error, if there is any.
func (c *FakeHostSubnets) Create(hostSubnet *network_v1.HostSubnet) (result *network_v1.HostSubnet, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(hostsubnetsResource, hostSubnet), &network_v1.HostSubnet{})
	if obj == nil {
		return nil, err
	}
	return obj.(*network_v1.HostSubnet), err
}

// Update takes the representation of a hostSubnet and updates it. Returns the server's representation of the hostSubnet, and an error, if there is any.
func (c *FakeHostSubnets) Update(hostSubnet *network_v1.HostSubnet) (result *network_v1.HostSubnet, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(hostsubnetsResource, hostSubnet), &network_v1.HostSubnet{})
	if obj == nil {
		return nil, err
	}
	return obj.(*network_v1.HostSubnet), err
}

// Delete takes name of the hostSubnet and deletes it. Returns an error if one occurs.
func (c *FakeHostSubnets) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(hostsubnetsResource, name), &network_v1.HostSubnet{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeHostSubnets) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(hostsubnetsResource, listOptions)

	_, err := c.Fake.Invokes(action, &network_v1.HostSubnetList{})
	return err
}

// Patch applies the patch and returns the patched hostSubnet.
func (c *FakeHostSubnets) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *network_v1.HostSubnet, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(hostsubnetsResource, name, data, subresources...), &network_v1.HostSubnet{})
	if obj == nil {
		return nil, err
	}
	return obj.(*network_v1.HostSubnet), err
}
