/* Copyright 2025 İrem Kuyucu <irem@digilol.net>
 * Copyright 2025 Laurynas Četyrkinas <laurynas@digilol.net>
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
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

type InboundSetting struct {
	Up, Down, Total, Remark, Enable, ExpiryTime, Listen, Port, Protocol string
}

type ClientOptions struct {
	ID         string `json:"id"`
	Flow       string `json:"flow"`
	Email      string `json:"email"`
	LimitIp    int    `json:"limitIp"`
	TotalGb    int    `json:"totalGB"`
	ExpiryTime int    `json:"expiryTime"`
	Enable     bool   `json:"enable"`
	TgId       string `json:"tgId"`
	SubId      string `json:"subId"`
	Reset      int    `json:"reset"`
}

type HeaderSetting struct {
	Type string `json:"type"`
}

type SniffingSetting struct {
	Enabled      bool     `json:"enabled"`
	DestOverride []string `json:"destOverride"`
	MetadataOnly bool     `json:"metadataOnly"`
	RouteOnly    bool     `json:"routeOnly"`
}

type FallbackOptions struct {
	Name string `json:"name"`
	Alpn string `json:"alpn"`
	Path string `json:"path"`
	Dest string `json:"dest"`
	Xver int    `json:"xver"`
}

type VlessSetting struct {
	Clients    []ClientOptions   `json:"clients"`
	Decryption string            `json:"decryption"`
	Fallbacks  []FallbackOptions `json:"fallbacks"`
}

type VmessSetting struct {
	Clients []ClientOptions `json:"clients"`
}

func AddInbound[T VlessSetting | VmessSetting, K TcpStreamSetting | QuicStreamSetting | XhttpStreamSetting](ctx context.Context, c *Client, inOpt InboundSetting, pOpt T, strOpt K, sniOpt SniffingSetting) (*ApiResponse, error) {
	form := url.Values{}

	protoSettings, err := json.Marshal(pOpt)
	if err != nil {
		return nil, err
	}
	form.Add("settings", string(protoSettings))

	streamSettings, err := json.Marshal(strOpt)
	if err != nil {
		return nil, err
	}
	form.Add("streamSettings", string(streamSettings))

	sniffingSettings, err := json.Marshal(sniOpt)
	if err != nil {
		return nil, err
	}
	form.Add("sniffing", string(sniffingSettings))

	form.Add("up", inOpt.Up)
	form.Add("down", inOpt.Down)
	form.Add("total", inOpt.Total)
	form.Add("remark", inOpt.Remark)
	form.Add("enable", inOpt.Enable)
	form.Add("expiryTime", inOpt.ExpiryTime)
	form.Add("listen", inOpt.Listen)
	form.Add("port", inOpt.Port)
	form.Add("protocol", inOpt.Protocol)

	resp := &ApiResponse{}
	err = c.DoForm(ctx, http.MethodPost, "/panel/api/inbounds/add", form, resp)
	if err != nil {
		return nil, err
	}
	if !resp.Success {
		return resp, fmt.Errorf(resp.Msg)
	}
	return resp, err
}
