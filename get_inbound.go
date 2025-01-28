/* Copyright 2025 İrem Kuyucu <irem@digilol.net>
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
	"fmt"
	"net/http"
)

type GetInboundResponse struct {
	Success bool    `json:"success"`
	Msg     string  `json:"msg"`
	Obj     Inbound `json:"obj"`
}

func (c *Client) GetInbound(ctx context.Context, inbound_id uint) (*GetInboundResponse, error) {
	resp := &GetInboundResponse{}
	err := c.Do(ctx, http.MethodGet, fmt.Sprintf("/panel/api/inbounds/get/%d", inbound_id), nil, resp)
	if err != nil {
		return nil, err
	}
	if !resp.Success {
		return resp, fmt.Errorf(resp.Msg)
	}
	return resp, err
}
