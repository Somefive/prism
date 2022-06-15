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
	"sort"
	"strconv"
	"strings"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/util/json"
	"k8s.io/apiserver/pkg/endpoints/request"
	"k8s.io/apiserver/pkg/registry/rest"
	"sigs.k8s.io/apiserver-runtime/pkg/builder/resource"
	"sigs.k8s.io/controller-runtime/pkg/client"

	"github.com/kubevela/prism/pkg/util/singleton"
)

var _ resource.ArbitrarySubResource = &AppzRevisions{}
var _ rest.Getter = &AppzRevisions{}

var (
	// AppzRevisionsKind kind name for AppzRevisions
	AppzRevisionsKind = "AppzRevisions"
	// AppzRevisionsGroupVersionKind GroupVersionKind for AppzRevisions
	AppzRevisionsGroupVersionKind = GroupVersion.WithKind(AppzRevisionsKind)
)

// AppzRevisions is an extension model for ApplicationRevisions
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
type AppzRevisions struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:pruning:PreserveUnknownFields
	Items []runtime.RawExtension `json:"items,omitempty"`
}

func (in *AppzRevisions) New() runtime.Object {
	return &AppzRevisions{}
}

func (in *AppzRevisions) SubResourceName() string {
	return "revisions"
}

func extractRevisionNum(appRevision string) int {
	splits := strings.Split(appRevision, "-")
	if len(splits) > 1 {
		n, err := strconv.Atoi(strings.TrimPrefix(splits[len(splits)-1], "v"))
		if err == nil {
			return n
		}
	}
	return 0
}

func (in *AppzRevisions) Get(ctx context.Context, name string, options *metav1.GetOptions) (runtime.Object, error) {
	uns := &unstructured.UnstructuredList{}
	uns.SetGroupVersionKind(ApplicationRevisionGroupVersionKind)
	ns := request.NamespaceValue(ctx)
	listOpts := []client.ListOption{
		client.InNamespace(ns),
		client.MatchingLabels{labelAppName: name},
	}
	if err := singleton.GetKubeClient().List(ctx, uns, listOpts...); err != nil {
		return nil, err
	}
	sort.Slice(uns.Items, func(i, j int) bool {
		return extractRevisionNum(uns.Items[i].GetName()) < extractRevisionNum(uns.Items[j].GetName())
	})
	revs := &AppzRevisions{ObjectMeta: metav1.ObjectMeta{Namespace: ns, Name: name}}
	revs.SetGroupVersionKind(AppzRevisionsGroupVersionKind)
	revs.Items = make([]runtime.RawExtension, len(uns.Items))
	var err error
	for i := range uns.Items {
		if revs.Items[i].Raw, err = json.Marshal(uns.Items[i].Object); err != nil {
			return nil, err
		}
	}
	return revs, nil
}
