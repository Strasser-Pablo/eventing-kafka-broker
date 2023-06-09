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

// Code generated by informer-gen. DO NOT EDIT.

package v1

import (
	internalinterfaces "knative.dev/eventing/pkg/client/informers/externalversions/internalinterfaces"
)

// Interface provides access to all the informers in this group version.
type Interface interface {
	// ApiServerSources returns a ApiServerSourceInformer.
	ApiServerSources() ApiServerSourceInformer
	// ContainerSources returns a ContainerSourceInformer.
	ContainerSources() ContainerSourceInformer
	// PingSources returns a PingSourceInformer.
	PingSources() PingSourceInformer
	// SinkBindings returns a SinkBindingInformer.
	SinkBindings() SinkBindingInformer
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

// ApiServerSources returns a ApiServerSourceInformer.
func (v *version) ApiServerSources() ApiServerSourceInformer {
	return &apiServerSourceInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// ContainerSources returns a ContainerSourceInformer.
func (v *version) ContainerSources() ContainerSourceInformer {
	return &containerSourceInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// PingSources returns a PingSourceInformer.
func (v *version) PingSources() PingSourceInformer {
	return &pingSourceInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// SinkBindings returns a SinkBindingInformer.
func (v *version) SinkBindings() SinkBindingInformer {
	return &sinkBindingInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}
