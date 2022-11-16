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

package singleton

import (
	"k8s.io/apiserver/pkg/server"
	"sigs.k8s.io/apiserver-runtime/pkg/builder"
)

var apiserver *builder.GenericAPIServer

func GetGenericAPIServer() *builder.GenericAPIServer {
	return apiserver
}

func InitGenericAPIServer(server *builder.GenericAPIServer) *builder.GenericAPIServer {
	apiserver = server
	return server
}

var serverConfig *server.Config

func GetServerConfig() *server.Config {
	return serverConfig
}

func InitServerConfig(config *server.RecommendedConfig) *server.RecommendedConfig {
	serverConfig = &server.Config{}
	*serverConfig = config.Config
	return config
}
