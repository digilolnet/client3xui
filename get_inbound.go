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
	err := c.Do(ctx, http.MethodPost, fmt.Sprintf("/panel/api/inbounds/get/%d", inbound_id), nil, resp)
	if err != nil {
		return nil, err
	}
	if !resp.Success {
		return resp, fmt.Errorf(resp.Msg)
	}
	return resp, err
}
