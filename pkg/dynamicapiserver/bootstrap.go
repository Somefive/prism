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

package apiserver

import (
	"k8s.io/apiserver/pkg/server"

	"github.com/kubevela/prism/pkg/util/singleton"
)

// StartDefaultDynamicAPIServer run default dynamic apiserver in backend
func StartDefaultDynamicAPIServer(ctx server.PostStartHookContext) error {
	DefaultDynamicAPIServer = NewDynamicAPIServer(
		singleton.GenericAPIServer.Get(),
		singleton.APIServerConfig.Get())
	go StartDynamicResourceFactoryWithConfigMapInformer(ctx.StopCh)
	//go StartDynamicResourceFactoryWithConfigMapInformer(ctx.StopCh)
	return nil
}