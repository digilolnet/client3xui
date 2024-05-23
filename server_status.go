/* Copyright 2024 İrem Kuyucu <irem@digilol.net>
 * Copyright 2024 Laurynas Četyrkinas <laurynas@digilol.net>
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package client3xui

import (
	"context"
	"net/http"
)

type ProcessState string

const (
	Running ProcessState = "running"
	Stop    ProcessState = "stop"
	Error   ProcessState = "error"
)

type ServerStatusResponse struct {
	Success bool   `json:"success"`
	Msg     string `json:"msg"`
	Obj     *struct {
		Cpu         float64 `json:"cpu"`
		CpuCores    int     `json:"cpuCores"`
		CpuSpeedMhz float64 `json:"cpuSpeedMhz"`
		Mem         struct {
			Current uint64 `json:"current"`
			Total   uint64 `json:"total"`
		} `json:"mem"`
		Swap struct {
			Current uint64 `json:"current"`
			Total   uint64 `json:"total"`
		} `json:"swap"`
		Disk struct {
			Current uint64 `json:"current"`
			Total   uint64 `json:"total"`
		} `json:"disk"`
		Xray struct {
			State    ProcessState `json:"state"`
			ErrorMsg string       `json:"errorMsg"`
			Version  string       `json:"version"`
		} `json:"xray"`
		Uptime   uint64    `json:"uptime"`
		Loads    []float64 `json:"loads"`
		TcpCount int       `json:"tcpCount"`
		UdpCount int       `json:"udpCount"`
		NetIO    struct {
			Up   uint64 `json:"up"`
			Down uint64 `json:"down"`
		} `json:"netIO"`
		NetTraffic struct {
			Sent uint64 `json:"sent"`
			Recv uint64 `json:"recv"`
		} `json:"netTraffic"`
		PublicIP struct {
			IPv4 string `json:"ipv4"`
			IPv6 string `json:"ipv6"`
		} `json:"publicIP"`
		AppStats struct {
			Threads uint32 `json:"threads"`
			Mem     uint64 `json:"mem"`
			Uptime  uint64 `json:"uptime"`
		} `json:"appStats"`
	} `json:"obj"`
}

func (c *Client) ServerStatus(ctx context.Context) (*ServerStatusResponse, error) {
	resp := &ServerStatusResponse{}
	err := c.Do(ctx, http.MethodPost, "/server/status", nil, resp)
	return resp, err
}
