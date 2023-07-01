// Tencent is pleased to support the open source community by making Polaris available.
//
// Copyright (C) 2019 THL A29 Limited, a Tencent company. All rights reserved.
//
// Licensed under the BSD 3-Clause License (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// https://opensource.org/licenses/BSD-3-Clause
//
// Unless required by applicable law or agreed to in writing, software distributed
// under the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR
// CONDITIONS OF ANY KIND, either express or implied. See the License for the
// specific language governing permissions and limitations under the License.

package polarisapi

// Response
type Response struct {
	Code      uint32     `json:"code"`
	Info      string     `json:"info"`
	Amount    uint32     `json:"amount"`
	Size      uint32     `json:"size"`
	Instances []Instance `json:"instances"`
	User      User       `json:"user"`
	LogicSet  string     `json:"logic_set"`
	Ctime     bool       `json:"ctime"`
	Mtime     bool       `json:"mtime"`
}

// Instance
type Instance struct {
	ID                string            `json:"id,omitempty"`
	ServiceToken      string            `json:"service_token,omitempty"`
	Service           string            `json:"service,omitempty"`
	Namespace         string            `json:"namespace,omitempty"`
	VpcID             string            `json:"vpc_id,omitempty"`
	Host              string            `json:"host,omitempty"`
	Port              *int              `json:"port,omitempty"`
	Protocol          string            `json:"protocol,omitempty"`
	Version           string            `json:"version,omitempty"`
	Priority          int               `json:"priority,omitempty"`
	Weight            *int              `json:"weight,omitempty"`
	EnableHealthCheck *bool             `json:"enableHealthCheck,omitempty"`
	HealthCheck       *HealthCheck      `json:"healthCheck,omitempty"`
	Healthy           *bool             `json:"healthy,omitempty"`
	Isolate           *bool             `json:"isolate,omitempty"`
	Metadata          map[string]string `json:"metadata,omitempty"`
	Revision          string            `json:"revision,omitempty"`
}

type SimpleInstance struct {
	ID           string            `json:"id,omitempty"`
	ServiceToken string            `json:"service_token,omitempty"`
	Service      string            `json:"service,omitempty"`
	Namespace    string            `json:"namespace,omitempty"`
	VpcID        string            `json:"vpc_id,omitempty"`
	Host         string            `json:"host,omitempty"`
	Port         *int              `json:"port,omitempty"`
	Protocol     string            `json:"protocol,omitempty"`
	Version      string            `json:"version,omitempty"`
	Priority     int               `json:"priority,omitempty"`
	Weight       *int              `json:"weight,omitempty"`
	Healthy      *bool             `json:"healthy,omitempty"`
	Isolate      *bool             `json:"isolate,omitempty"`
	Metadata     map[string]string `json:"metadata,omitempty"`
	Revision     string            `json:"revision,omitempty"`
}

// HealthCheck
type HealthCheck struct {
	Type      *int      `json:"type,omitempty"`
	Heartbeat Heartbeat `json:"heartbeat"`
}

// Heartbeat
type Heartbeat struct {
	TTL int `json:"ttl"`
}

// AddResponse
type AddResponse struct {
	Code      uint32             `json:"code"`
	Info      string             `json:"info"`
	Size      uint32             `json:"size"`
	Responses []InstanceResponse `json:"responses"`
}

// InstanceResponse
type InstanceResponse struct {
	Code         uint32         `json:"code"`
	Info         string         `json:"info"`
	Instance     SimpleInstance `json:"instance"`
	ServiceToken string         `json:"service_token"`
}

// GetServiceResponse
type GetServiceResponse struct {
	Code     uint32    `json:"code"`
	Info     string    `json:"info"`
	Amount   uint32    `json:"amount"`
	Size     uint32    `json:"size"`
	Services []Service `json:"services"`
}

// Service
type Service struct {
	Name       string            `json:"name"`
	Token      string            `json:"token,omitempty"`
	Namespace  string            `json:"namespace,omitempty"`
	Metadata   map[string]string `json:"metadata,omitempty"`
	Ports      string            `json:"ports,omitempty"`
	Business   string            `json:"business,omitempty"`
	Department string            `json:"department,omitempty"`
	CmdbMod1   string            `json:"cmdb_mod1,omitempty"`
	CmdbMod2   string            `json:"cmdb_mod2,omitempty"`
	CmdbMod3   string            `json:"cmdb_mod3,omitempty"`
	Owners     string            `json:"owners,omitempty"`
}

// PutServicesResponse
type PutServicesResponse struct {
	Code     uint32               `json:"code"`
	Info     string               `json:"info"`
	Size     uint32               `json:"size"`
	Response []PutServiceResponse `json:"response"`
}

// PutServiceResponse
type PutServiceResponse struct {
	Code    uint32  `json:"code"`
	Info    string  `json:"info"`
	Token   string  `json:"token"`
	Size    uint32  `json:"size"`
	Service Service `json:"service"`
}

// PutServicesResponse
type CreateServicesResponse struct {
	Code     uint32                  `json:"code"`
	Info     string                  `json:"info"`
	Size     uint32                  `json:"size"`
	Response []CreateServiceResponse `json:"response"`
}

// PutServicesResponse
type CreateServiceResponse struct {
	Code    uint32  `json:"code"`
	Info    string  `json:"info"`
	Service Service `json:"service"`
}

// CreateServiceRequest 创建北极星 service 请求
type CreateServiceRequest struct {
	Name       string            `json:"name"`
	Namespace  string            `json:"namespace"`
	Owners     string            `json:"owners,omitempty"`
	Business   string            `json:"business,omitempty"`
	Department string            `json:"department,omitempty"`
	Comment    string            `json:"comment,omitempty"`
	Metadata   map[string]string `json:"metadata,omitempty"`
	Ports      string            `json:"ports,omitempty"`
}

// CreateServiceAliasRequest create service alias request
type CreateServiceAliasRequest struct {
	Service        string `json:"service"`
	Namespace      string `json:"namespace"`
	Owners         string `json:"owners,omitempty"`
	Alias          string `json:"alias"`
	AliasNamespace string `json:"alias_namespace"`
}

// CreateServiceAliasResponse create service alias response
type CreateServiceAliasResponse struct {
	Code  uint32 `json:"code"`
	Info  string `json:"info"`
	Alias Alias  `json:"alias"`
}

// Alias service alias
type Alias struct {
	Service        string `json:"service"`
	Namespace      string `json:"namespace"`
	Alias          string `json:"alias"`
	AliasNamespace string `json:"alias_namespace"`
}

// GetNamespacesResponse 获取命名空间的返回
type GetNamespacesResponse struct {
	Code       uint32              `json:"code"`
	Info       string              `json:"info"`
	Amount     uint32              `json:"amount"`
	Size       uint32              `json:"size"`
	Namespaces []NamespaceResponse `json:"namespaces"`
}

// NamespaceResponse
type NamespaceResponse struct {
	Name    string `json:"name"`
	Comment string `json:"comment"`
	Owners  string `json:"owners"`
}

// CreateNamespacesRequest
type CreateNamespacesRequest struct {
	Name    string `json:"name"`
	Owners  string `json:"owners,omitempty"`
	Comment string `json:"comment,omitempty"`
}

// CreateNamespacesResponse
type CreateNamespacesResponse struct {
	Code      uint32                    `json:"code"`
	Info      string                    `json:"info"`
	Size      uint32                    `json:"size"`
	Responses []CreateNamespaceResponse `json:"responses"`
}

// CreateNamespaceResponse
type CreateNamespaceResponse struct {
	Code      uint32    `json:"code"`
	Info      string    `json:"info"`
	Namespace Namespace `json:"namespace"`
}

// Namespace
type Namespace struct {
	Name   string `json:"name"`
	Owners string `json:"owners"`
}

type User struct {
	Id        string `json:"id,omitempty"`
	Name      string `json:"name,omitempty"`
	AuthToken string `json:"auth_token,omitempty"`
}

type ConfigFileTag struct {
	Key   string `json:"key,omitempty"`
	Value string `json:"value,omitempty"`
}

type ConfigFile struct {
	Id          uint64           `json:"id,omitempty"`
	Name        string           `json:"name,omitempty"`
	Namespace   string           `json:"namespace,omitempty"`
	Group       string           `json:"group,omitempty"`
	Content     string           `json:"content,omitempty"`
	Format      string           `json:"format,omitempty"`
	Comment     string           `json:"comment,omitempty"`
	Status      string           `json:"status,omitempty"`
	Tags        []*ConfigFileTag `json:"tags,omitempty"`
	CreateTime  string           `json:"createTime,omitempty"`
	CreateBy    string           `json:"createBy,omitempty"`
	ModifyTime  string           `json:"modifyTime,omitempty"`
	ModifyBy    string           `json:"modifyBy,omitempty"`
	ReleaseTime string           `json:"releaseTime,omitempty"`
	ReleaseBy   string           `json:"releaseBy,omitempty"`
	// 是否为加密配置文件
	Encrypted bool `json:"encrypted,omitempty"`
	// 加密算法
	EncryptAlgo string `json:"encryptAlgo,omitempty"`
}

type ConfigFileRelease struct {
	Id         uint64 `json:"id,omitempty"`
	Name       string `json:"name,omitempty"`
	Namespace  string `json:"namespace,omitempty"`
	Group      string `json:"group,omitempty"`
	FileName   string `json:"file_name,omitempty"`
	Content    string `json:"content,omitempty"`
	Comment    string `json:"comment,omitempty"`
	Md5        string `json:"md5,omitempty"`
	Version    uint64 `json:"version,omitempty"`
	CreateTime string `json:"create_time,omitempty"`
	CreateBy   string `json:"create_by,omitempty"`
	ModifyTime string `json:"modify_time,omitempty"`
	ModifyBy   string `json:"modify_by,omitempty"`
}

type ConfigResponse struct {
	Code              uint32             `json:"code"`
	Info              string             `json:"info"`
	ConfigFile        *ConfigFile        `json:"configFile,omitempty"`
	ConfigFileRelease *ConfigFileRelease `json:"configFileRelease,omitempty"`
}
