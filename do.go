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
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) Do(ctx context.Context, method, path string, in, out interface{}) error {
	b := new(bytes.Buffer)
	err := json.NewEncoder(b).Encode(in)
	if err != nil {
		return err
	}
	req, err := http.NewRequestWithContext(ctx, method, c.url+path, b)
	if err != nil {
		return err
	}
	err = c.loginIfNoCookie(ctx)
	if err != nil {
		return err
	}
	req.AddCookie(c.sessionCookie)
	req.Header.Set("Content-Type", "application/json")
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return err
	}
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("http status %v", resp.StatusCode)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	return json.Unmarshal(body, out)
}

func (c *Client) DoRaw(ctx context.Context, method, baseurl, path, contentType string, body []byte) ([]byte, error) {
	req, err := http.NewRequestWithContext(ctx, method, baseurl+path, bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}
	err = c.loginIfNoCookie(ctx)
	if err != nil {
		return nil, err
	}
	req.AddCookie(c.sessionCookie)
	req.Header.Set("Content-Type", contentType)
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("http status %v", resp.StatusCode)
	}
	defer resp.Body.Close()
	return io.ReadAll(resp.Body)
}
