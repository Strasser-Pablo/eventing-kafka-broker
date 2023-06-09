/*
Copyright 2021 The Knative Authors

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
	"context"
	"time"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
	v1 "knative.dev/eventing/pkg/apis/flows/v1"
	scheme "knative.dev/eventing/pkg/client/clientset/versioned/scheme"
)

// SequencesGetter has a method to return a SequenceInterface.
// A group's client should implement this interface.
type SequencesGetter interface {
	Sequences(namespace string) SequenceInterface
}

// SequenceInterface has methods to work with Sequence resources.
type SequenceInterface interface {
	Create(ctx context.Context, sequence *v1.Sequence, opts metav1.CreateOptions) (*v1.Sequence, error)
	Update(ctx context.Context, sequence *v1.Sequence, opts metav1.UpdateOptions) (*v1.Sequence, error)
	UpdateStatus(ctx context.Context, sequence *v1.Sequence, opts metav1.UpdateOptions) (*v1.Sequence, error)
	Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error
	Get(ctx context.Context, name string, opts metav1.GetOptions) (*v1.Sequence, error)
	List(ctx context.Context, opts metav1.ListOptions) (*v1.SequenceList, error)
	Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.Sequence, err error)
	SequenceExpansion
}

// sequences implements SequenceInterface
type sequences struct {
	client rest.Interface
	ns     string
}

// newSequences returns a Sequences
func newSequences(c *FlowsV1Client, namespace string) *sequences {
	return &sequences{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the sequence, and returns the corresponding sequence object, and an error if there is any.
func (c *sequences) Get(ctx context.Context, name string, options metav1.GetOptions) (result *v1.Sequence, err error) {
	result = &v1.Sequence{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("sequences").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of Sequences that match those selectors.
func (c *sequences) List(ctx context.Context, opts metav1.ListOptions) (result *v1.SequenceList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1.SequenceList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("sequences").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested sequences.
func (c *sequences) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("sequences").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a sequence and creates it.  Returns the server's representation of the sequence, and an error, if there is any.
func (c *sequences) Create(ctx context.Context, sequence *v1.Sequence, opts metav1.CreateOptions) (result *v1.Sequence, err error) {
	result = &v1.Sequence{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("sequences").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(sequence).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a sequence and updates it. Returns the server's representation of the sequence, and an error, if there is any.
func (c *sequences) Update(ctx context.Context, sequence *v1.Sequence, opts metav1.UpdateOptions) (result *v1.Sequence, err error) {
	result = &v1.Sequence{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("sequences").
		Name(sequence.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(sequence).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *sequences) UpdateStatus(ctx context.Context, sequence *v1.Sequence, opts metav1.UpdateOptions) (result *v1.Sequence, err error) {
	result = &v1.Sequence{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("sequences").
		Name(sequence.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(sequence).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the sequence and deletes it. Returns an error if one occurs.
func (c *sequences) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("sequences").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *sequences) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("sequences").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched sequence.
func (c *sequences) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.Sequence, err error) {
	result = &v1.Sequence{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("sequences").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
