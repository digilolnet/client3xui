# client3xui

[![Go report](https://goreportcard.com/badge/github.com/digilolnet/client3xui)](https://goreportcard.com/report/github.com/digilolnet/client3xui)
[![GoDoc](https://godoc.org/github.com/digilolnet/client3xui?status.svg)](https://godoc.org/github.com/digilolnet/client3xui)
[![License](https://img.shields.io/github/license/digilolnet/client3xui.svg)](https://github.com/digilolnet/client3xui/blob/master/LICENSE.md)

[3X-UI](https://github.com/MHSanaei/3x-ui) API wrapper in Go brought to you by:

[<img alt="Digilol offers managed hosting and software development" src="https://www.digilol.net/banner-hosting-development.png" width="500" />](https://www.digilol.net)

## Examples

### VMESS + TCP

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

### VLESS + REALITY + XHTTP

```go
package main

import (
        "context"
        "fmt"
        "log"

        "github.com/google/uuid"
        "github.com/digilolnet/client3xui"
)

func main() {
        server := client3xui.New(client3xui.Config{
                Url:      "https://xrayserver.tld:8843/panelpath",
                Username: "digilol",
                Password: "secr3t",
        })


        // Get server status
        status, err := server.ServerStatus(context.Background())
        if err != nil {
                log.Fatal(err)
        }
        fmt.Println(status)

        // Get new X25519 certificate for Reality
        cert, err := server.GetNewX25519Cert(context.Background())
        if err != nil {
                log.Fatal(err)
        }

        // Add new inbound
        inbound := client3xui.InboundSetting{
                Up:         "0",
                Down:       "0",
                Total:      "0",
                Remark:     "",
                Enable:     "true",
                ExpiryTime: "0",
                Listen:     "",
                Port:       "10600",
                Protocol:   "vless",
        }

        proto := client3xui.VlessSetting{
                Clients: []client3xui.ClientOptions{
                        {
                                ID:     uuid.NewString(),
                                Flow:   "",
                                Email:  "realityclient",
                                Enable: true,
                                SubId:  "81h6cr12m5w0wm6i",
                        },
                },
                Decryption: "none",
                Fallbacks:  []client3xui.FallbackOptions{},
        }

        stream := client3xui.XhttpStreamSetting{
                Network:       "xhttp",
                Security:     "reality",
                ExternalProxy: []string{},
                RealitySettings: client3xui.RealitySettings{
                        Show:        false,
                        Xver:        0,
                        Dest:        "yahoo.com:443",
                        ServerNames: []string{"yahoo.com", "www.yahoo.com"},
                        PrivateKey:  cert.Obj.PrivateKey,
                        MinClient:   "",
                        MaxClient:   "",
                        MaxTimediff: 0,
                        ShortIds:    []string{client3xui.GenerateShortId(14), client3xui.GenerateShortId(2), client3xui.GenerateShortId(16)},
                        Settings: client3xui.RealitySettingsInner{
                                PublicKey:   cert.Obj.PublicKey,
                                Fingerprint: "chrome",
                                ServerName:  "",
                                SpiderX:     "/",
                        },
                },
                XhttpSettings: client3xui.XhttpSetting{
                        Path:               "/",
                        Host:               "",
                        Headers:            map[string]string{},
                        ScMaxBufferedPosts: 30,
                        ScMaxEachPostBytes: "1000000",
                        NoSSEHeader:        false,
                        XPaddingBytes:      "100-1000",
                        Mode:               "auto",
                },
        }

        snif := client3xui.SniffingSetting{
                Enabled:      false,
                DestOverride: []string{"http", "tls", "quic", "fakedns"},
                MetadataOnly: false,
                RouteOnly:    false,
        }

        ret, err := client3xui.AddInbound[client3xui.VlessSetting, client3xui.XhttpStreamSetting](
                context.Background(),
                server,
                inbound,
                proto,
                stream,
                snif,
        )
        if err != nil {
                log.Fatal(err)
        }
        log.Printf("%v", ret)
}
```
### Development process
| Method | Path                               | Action                                      | Done |
| :----: | ---------------------------------- | ------------------------------------------- | ---- |
| `GET`  | `"/list"`                          | Get all inbounds                            |  ✅  |
| `GET`  | `"/get/:id"`                       | Get inbound with inbound.id                 |  ✅  |
| `GET`  | `"/getClientTraffics/:email"`      | Get Client Traffics with email              |  ⛔️  |
| `GET`  | `"/getClientTrafficsById/:id"`     | Get client's traffic By ID                  |  ⛔️  |
| `GET`  | `"/createbackup"`                  | Telegram bot sends backup to admins         |  ⛔️  |
| `POST` | `"/add"`                           | Add inbound                                 |  ✅  |
| `POST` | `"/del/:id"`                       | Delete Inbound                              |  ✅  |
| `POST` | `"/update/:id"`                    | Update Inbound                              |  ⛔️  |
| `POST` | `"/clientIps/:email"`              | Client Ip address                           |  ⛔️  |
| `POST` | `"/clearClientIps/:email"`         | Clear Client Ip address                     |  ⛔️  |
| `POST` | `"/addClient"`                     | Add Client to inbound                       |  ✅  |
| `POST` | `"/:id/delClient/:clientId"`       | Delete Client by clientId\*                 |  ✅  |
| `POST` | `"/updateClient/:clientId"`        | Update Client by clientId\*                 |  ✅  |
| `POST` | `"/:id/resetClientTraffic/:email"` | Reset Client's Traffic                      |  ⛔️  |
| `POST` | `"/resetAllTraffics"`              | Reset traffics of all inbounds              |  ⛔️  |
| `POST` | `"/resetAllClientTraffics/:id"`    | Reset traffics of all clients in an inbound |  ⛔️  |
| `POST` | `"/delDepletedClients/:id"`        | Delete inbound depleted clients (-1: all)   |  ⛔️  |
| `POST` | `"/onlines"`                       | Get Online users ( list of emails )         |  ✅  |
