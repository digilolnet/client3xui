package client3xui

import (
	"context"
	"fmt"
	"net/http"
)

type GetInboundsResponse struct {
	Success bool      `json:"success"`
	Msg     string    `json:"msg"`
	Obj     []Inbound `json:"obj"`
}

type Inbound struct {
	ID             int          `json:"id"`
	Up             int          `json:"up"`
	Down           int          `json:"down"`
	Total          int          `json:"total"`
	Remark         string       `json:"remark"`
	Enable         bool         `json:"enable"`
	ExpiryTime     int          `json:"expiryTime"`
	ClientStats    []ClientStat `json:"clientStats"`
	Listen         string       `json:"listen"`
	Port           int          `json:"port"`
	Protocol       string       `json:"protocol"`
	Settings       string       `json:"settings"`
	StreamSettings string       `json:"streamSettings"`
	Tag            string       `json:"tag"`
	Sniffing       string       `json:"sniffing"`
}

type ClientStat struct {
	ID         int    `json:"id"`
	InboundID  int    `json:"inboundId"`
	Enable     bool   `json:"enable"`
	Email      string `json:"email"`
	Up         int    `json:"up"`
	Down       int    `json:"down"`
	ExpiryTime int    `json:"expiryTime"`
	Total      int    `json:"total"`
	Reset      int    `json:"reset"`
}

func (c *Client) GetInbounds(ctx context.Context) (*GetInboundsResponse, error) {
	resp := &GetInboundsResponse{}
	err := c.Do(ctx, http.MethodPost, "/panel/api/inbounds/list", nil, resp)
	if err != nil {
		return nil, err
	}
	if !resp.Success {
		return resp, fmt.Errorf(resp.Msg)
	}
	return resp, err
}
