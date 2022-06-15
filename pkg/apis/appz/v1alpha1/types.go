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
	"context"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/apiserver/pkg/endpoints/request"
	"sigs.k8s.io/apiserver-runtime/pkg/builder/resource"

	"github.com/kubevela/prism/pkg/util/singleton"
)

// Appz is an extension model for Application
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
type Appz struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:pruning:PreserveUnknownFields
	Spec runtime.RawExtension `json:"spec,omitempty"`
}

// AppzList list for Appz
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
type AppzList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`

	Items []Appz `json:"items"`
}

var _ resource.Object = &Appz{}

// GetObjectMeta returns the object meta reference.
func (in *Appz) GetObjectMeta() *metav1.ObjectMeta {
	return &in.ObjectMeta
}

// NamespaceScoped returns if the object must be in a namespace.
func (in *Appz) NamespaceScoped() bool {
	return true
}

// New returns a new instance of the resource
func (in *Appz) New() runtime.Object {
	return &Appz{}
}

// NewList return a new list instance of the resource
func (in *Appz) NewList() runtime.Object {
	return &AppzList{}
}

// GetGroupVersionResource returns the GroupVersionResource for this resource.
func (in *Appz) GetGroupVersionResource() schema.GroupVersionResource {
	return GroupVersion.WithResource(AppzResource)
}

// IsStorageVersion returns true if the object is also the internal version
func (in *Appz) IsStorageVersion() bool {
	return true
}

// ShortNames delivers a list of short names for a resource.
func (in *Appz) ShortNames() []string {
	return []string{"Appz"}
}

var _ resource.ObjectWithArbitrarySubResource = &Appz{}

func (in *Appz) GetArbitrarySubResources() []resource.ArbitrarySubResource {
	return []resource.ArbitrarySubResource{
		&AppzRevisions{},
	}
}

func (in *Appz) Get(ctx context.Context, name string, options *metav1.GetOptions) (runtime.Object, error) {
	un := &unstructured.Unstructured{}
	un.SetGroupVersionKind(ApplicationGroupVersionKind)
	ns := request.NamespaceValue(ctx)
	if err := singleton.GetKubeClient().Get(ctx, types.NamespacedName{Namespace: ns, Name: name}, un); err != nil {
		return nil, err
	}
	appz := &Appz{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(un.Object, appz); err != nil {
		return nil, err
	}
	appz.SetGroupVersionKind(AppzGroupVersionKind)
	return appz, nil
}
