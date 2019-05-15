/*
Copyright The Kubernetes Authors.

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

// Code generated by client-gen. DO NOT EDIT.

package v1

import (
	v1 "github.com/rook/rook/pkg/apis/ceph.rook.io/v1"
	scheme "github.com/rook/rook/pkg/client/clientset/versioned/scheme"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// CephClustersGetter has a method to return a CephClusterInterface.
// A group's client should implement this interface.
type CephClustersGetter interface {
	CephClusters(namespace string) CephClusterInterface
}

// CephClusterInterface has methods to work with CephCluster resources.
type CephClusterInterface interface {
	Create(*v1.CephCluster) (*v1.CephCluster, error)
	Update(*v1.CephCluster) (*v1.CephCluster, error)
	Delete(name string, options *metav1.DeleteOptions) error
	DeleteCollection(options *metav1.DeleteOptions, listOptions metav1.ListOptions) error
	Get(name string, options metav1.GetOptions) (*v1.CephCluster, error)
	List(opts metav1.ListOptions) (*v1.CephClusterList, error)
	Watch(opts metav1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.CephCluster, err error)
	CephClusterExpansion
}

// cephClusters implements CephClusterInterface
type cephClusters struct {
	client rest.Interface
	ns     string
}

// newCephClusters returns a CephClusters
func newCephClusters(c *CephV1Client, namespace string) *cephClusters {
	return &cephClusters{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the cephCluster, and returns the corresponding cephCluster object, and an error if there is any.
func (c *cephClusters) Get(name string, options metav1.GetOptions) (result *v1.CephCluster, err error) {
	result = &v1.CephCluster{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("cephclusters").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of CephClusters that match those selectors.
func (c *cephClusters) List(opts metav1.ListOptions) (result *v1.CephClusterList, err error) {
	result = &v1.CephClusterList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("cephclusters").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested cephClusters.
func (c *cephClusters) Watch(opts metav1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("cephclusters").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch()
}

// Create takes the representation of a cephCluster and creates it.  Returns the server's representation of the cephCluster, and an error, if there is any.
func (c *cephClusters) Create(cephCluster *v1.CephCluster) (result *v1.CephCluster, err error) {
	result = &v1.CephCluster{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("cephclusters").
		Body(cephCluster).
		Do().
		Into(result)
	return
}

// Update takes the representation of a cephCluster and updates it. Returns the server's representation of the cephCluster, and an error, if there is any.
func (c *cephClusters) Update(cephCluster *v1.CephCluster) (result *v1.CephCluster, err error) {
	result = &v1.CephCluster{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("cephclusters").
		Name(cephCluster.Name).
		Body(cephCluster).
		Do().
		Into(result)
	return
}

// Delete takes name of the cephCluster and deletes it. Returns an error if one occurs.
func (c *cephClusters) Delete(name string, options *metav1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("cephclusters").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *cephClusters) DeleteCollection(options *metav1.DeleteOptions, listOptions metav1.ListOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("cephclusters").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched cephCluster.
func (c *cephClusters) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.CephCluster, err error) {
	result = &v1.CephCluster{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("cephclusters").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
