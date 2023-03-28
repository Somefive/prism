/*
Copyright 2022 The KubeVela Authors.

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

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	kruntime "k8s.io/apimachinery/pkg/util/runtime"
	"k8s.io/client-go/kubernetes/scheme"
)

const (
	// Group the group for the apiextensions
	Group = "prism.oam.dev"
	// Version the version for the v1alpha1 apiextensions
	Version = "v1alpha1"
)

func init() {
	kruntime.Must(AddToScheme(scheme.Scheme))
}

// AddToScheme add DynamicAPIDefinition scheme
var AddToScheme = func(scheme *runtime.Scheme) error {
	metav1.AddToGroupVersion(scheme, GroupVersion)
	// +kubebuilder:scaffold:install
	scheme.AddKnownTypes(GroupVersion)
	return nil
}

// GroupVersion the apiextensions v1alpha1 group version
var GroupVersion = schema.GroupVersion{Group: Group, Version: Version}

var (
	// DynamicAPIDefinitionResource resource name for DynamicAPIDefinition
	DynamicAPIDefinitionResource = "dynamicapidefinition"
	// DynamicAPIDefinitionKind kind name for DynamicAPIDefinition
	DynamicAPIDefinitionKind = "DynamicAPIDefinition"
	// DynamicAPIDefinitionGroupResource GroupResource for DynamicAPIDefinition
	DynamicAPIDefinitionGroupResource = schema.GroupResource{Group: Group, Resource: DynamicAPIDefinitionResource}
	// DynamicAPIDefinitionGroupVersionKind GroupVersionKind for DynamicAPIDefinition
	DynamicAPIDefinitionGroupVersionKind = GroupVersion.WithKind(DynamicAPIDefinitionKind)
)
