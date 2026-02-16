package sdk

import (
	"context"
	"net/http"
)

type VipQrURLReq struct {
	DefaultProductID string `json:"default_product_id"`
	OpenDevice       string `json:"open_device"`
	HideTitle        string `json:"hide_title"`
}

type VipQrURLResp struct {
	QrcodeURL string `json:"qrcode_url"`
}

// VipQrURL https://www.yuque.com/115yun/open/cguk6qshgapwg4qn#oByvI
func (c *Client) VipQrURL(ctx context.Context, req *VipQrURLReq) (*VipQrURLResp, error) {
	var resp VipQrURLResp
	_, err := c.AuthRequest(ctx, ApiVipQrURL, http.MethodGet, &resp, ReqWithQuery(Form{
		"default_product_id": req.DefaultProductID,
		"open_device":        req.OpenDevice,
		"hide_title":         req.HideTitle,
	}))
	if err != nil {
		return nil, err
	}
	return &resp, err
}
