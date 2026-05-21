// Licensed to the Apache Software Foundation (ASF) under one
// or more contributor license agreements.  See the NOTICE file
// distributed with this work for additional information
// regarding copyright ownership.  The ASF licenses this file
// to you under the Apache License, Version 2.0 (the
// "License"); you may not use this file except in compliance
// with the License.  You may obtain a copy of the License at
//
//	http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License
package base

import "fmt"

type Backend struct {
	Id       int64
	Host     string
	BePort   uint16
	HttpPort uint16
	BrpcPort uint16

	// External ports for Kubernetes NodePort/LoadBalancer scenarios
	ExternalHttpPort uint16 `json:"external_http_port,omitempty"`
	ExternalBePort   uint16 `json:"external_be_port,omitempty"`
}

// Backend Stringer
func (b *Backend) String() string {
	if b.ExternalHttpPort != 0 || b.ExternalBePort != 0 {
		return fmt.Sprintf("Backend: {Id: %d, Host: %s, BePort: %d, HttpPort: %d, ExternalBePort: %d, ExternalHttpPort: %d, BrpcPort: %d}",
			b.Id, b.Host, b.BePort, b.HttpPort, b.ExternalBePort, b.ExternalHttpPort, b.BrpcPort)
	}
	return fmt.Sprintf("Backend: {Id: %d, Host: %s, BePort: %d, HttpPort: %d, BrpcPort: %d}",
		b.Id, b.Host, b.BePort, b.HttpPort, b.BrpcPort)
}

func (b *Backend) GetHttpPortStr() string {
	// Use external port if available, otherwise use internal port
	if b.ExternalHttpPort != 0 {
		log.Debugf("[GetHttpPortStr] Using external HttpPort: %d (internal: %d) for backend %d on host %s",
			b.ExternalHttpPort, b.HttpPort, b.Id, b.Host)
		return fmt.Sprintf("%d", b.ExternalHttpPort)
	}
	log.Tracef("[GetHttpPortStr] Using internal HttpPort: %d for backend %d on host %s",
		b.HttpPort, b.Id, b.Host)
	return fmt.Sprintf("%d", b.HttpPort)
}

// GetBePort returns the external BePort if available, otherwise returns internal BePort
func (b *Backend) GetBePort() uint16 {
	if b.ExternalBePort != 0 {
		log.Debugf("[GetBePort] Using external BePort: %d (internal: %d) for backend %d on host %s",
			b.ExternalBePort, b.BePort, b.Id, b.Host)
		return b.ExternalBePort
	}
	log.Tracef("[GetBePort] Using internal BePort: %d for backend %d on host %s",
		b.BePort, b.Id, b.Host)
	return b.BePort
}
