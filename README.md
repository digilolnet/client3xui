# client3xui

[![Go report](https://goreportcard.com/badge/github.com/digilolnet/client3xui)](https://goreportcard.com/report/github.com/digilolnet/client3xui)
[![GoDoc](https://godoc.org/github.com/digilolnet/client3xui?status.svg)](https://godoc.org/github.com/digilolnet/client3xui)
[![License](https://img.shields.io/github/license/digilolnet/client3xui.svg)](https://github.com/digilolnet/client3xui/blob/master/LICENSE.md)

[3X-UI](https://github.com/MHSanaei/3x-ui) API wrapper in Go.

[![Digilol offers managed hosting and software development](https://www.digilol.net/banner-hosting-development.png)](https://www.digilol.net)

## Usage

```go
package main

import (
        "context"
        "fmt"
        "log"

        "github.com/digilolnet/client3xui"
)

func main() {
        server := client3xui.New(client3xui.Config{
                Url:      "https://xrayserver.tld:8843",
                Username: "digilol",
                Password: "secr3t",
        })

        // Get server status
        status, err := server.ServerStatus(context.Background())
        if err != nil {
                log.Fatal(err)
        }
        fmt.Println(status)

        //Add new inbound
        inbound := client3xui.InboundSetting{
                Up:         "0",
                Down:       "0",
                Total:      "0",
                Remark:     "",
                Enable:     "true",
                ExpiryTime: "0",
                Listen:     "",
                Port:       "13337",
                Protocol:   "vmess",
        }

        proto := client3xui.VmessSetting{
                Clients: []client3xui.ClientOptions{
                        client3xui.ClientOptions{
                                ID:     uuid.NewString(),
                                Email:  "niceclient",
                                Enable: true,
                                SubId:  "dhgsyf6384j9u889hd89edhlj",
                        },
                },
        }

        tcp := client3xui.TcpStreamSetting{
                Network:  "tcp",
                Security: "none",
                TcpSettings: client3xui.TcpSetting{
                        Header: client3xui.HeaderSetting{
                                Type: "none",
                        },
                },
        }

        snif := client3xui.SniffingSetting{
                Enabled:      true,
                DestOverride: []string{"http", "tls", "quic", "fakedns"},
        }

        ret, err := client3xui.AddInbound(context.Background(), server, inbound, proto, tcp, snif)
        if err != nil {
                log.Fatal(err)
        }

        // Add new client
        clis := []client3xui.XrayClient{
                {ID: "fab5a8c0-89b4-43a8-9871-82fe6e2c8c8a",
                Email:  "fab5a8c0-89b4-43a8-9871-82fe6e2c8c8a",
                Enable: true},
        }
        resp, err := server.AddClient(context.Background(), 1, clis)
        if err != nil {
                log.Fatal(err)
        }
        fmt.Println(*resp)
}
```
