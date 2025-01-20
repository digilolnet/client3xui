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

type XhttpSetting struct {
	Path               string            `json:"path"`
	Host               string            `json:"host"`
	Headers            map[string]string `json:"headers"`
	ScMaxBufferedPosts int               `json:"scMaxBufferedPosts"`
	ScMaxEachPostBytes string            `json:"scMaxEachPostBytes"`
	NoSSEHeader        bool              `json:"noSSEHeader"`
	XPaddingBytes      string            `json:"xPaddingBytes"`
	Mode               string            `json:"mode"`
}

type XhttpStreamSetting struct {
	Network         string          `json:"network"`
	Security        string          `json:"security"`
	ExternalProxy   []string        `json:"externalProxy"`
	RealitySettings RealitySettings `json:"realitySettings"`
	XhttpSettings   XhttpSetting    `json:"xhttpSettings"`
}
