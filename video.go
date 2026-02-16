package sdk

import (
	"context"
	"net/http"
	"strconv"
)

type VideoPlayURL struct {
	URL        string `json:"url"`
	Definition int    `json:"definition"`
	Desc       string `json:"desc"`
}

type VideoPlayResp struct {
	FileID   string         `json:"file_id"`
	FileName string         `json:"file_name"`
	FileSize int64          `json:"file_size"`
	Duration int64          `json:"duration"`
	Width    int            `json:"width"`
	Height   int            `json:"height"`
	VideoURL []VideoPlayURL `json:"video_url"`
}

// VideoPlay https://www.yuque.com/115yun/open/hqglxv3cedi3p9dz
func (c *Client) VideoPlay(ctx context.Context, pickCode string) (*VideoPlayResp, error) {
	var resp VideoPlayResp
	_, err := c.AuthRequest(ctx, ApiVideoPlay, http.MethodGet, &resp, ReqWithQuery(Form{
		"pick_code": pickCode,
	}))
	if err != nil {
		return nil, err
	}
	return &resp, err
}

type VideoHistoryResp struct {
	AddTime  int64  `json:"add_time"`
	FileID   string `json:"file_id"`
	FileName string `json:"file_name"`
	Hash     string `json:"hash"`
	PickCode string `json:"pick_code"`
	Time     string `json:"time"`
}

// GetVideoHistory https://www.yuque.com/115yun/open/gssqdrsq6vfqigag
func (c *Client) GetVideoHistory(ctx context.Context, pickCode string) (*VideoHistoryResp, error) {
	var resp VideoHistoryResp
	_, err := c.AuthRequest(ctx, ApiVideoHistory, http.MethodGet, &resp, ReqWithQuery(Form{
		"pick_code": pickCode,
	}))
	if err != nil {
		return nil, err
	}
	return &resp, err
}

type SetVideoHistoryReq struct {
	PickCode string `json:"pick_code"`
	Time     int    `json:"time"`
	WatchEnd int    `json:"watch_end"`
}

// SetVideoHistory https://www.yuque.com/115yun/open/bshagbxv1gzqglg4
func (c *Client) SetVideoHistory(ctx context.Context, req *SetVideoHistoryReq) error {
	var resp any
	_, err := c.AuthRequest(ctx, ApiVideoHistory, http.MethodPost, &resp, ReqWithForm(Form{
		"pick_code": req.PickCode,
		"time":      strconv.Itoa(req.Time),
		"watch_end": strconv.Itoa(req.WatchEnd),
	}))
	return err
}

type SubtitleItem struct {
	SID          string `json:"sid"`
	Language     string `json:"language"`
	Title        string `json:"title"`
	URL          string `json:"url"`
	Type         string `json:"type"`
	Sha1         string `json:"sha1"`
	FileID       string `json:"file_id"`
	FileName     string `json:"file_name"`
	PickCode     string `json:"pick_code"`
	CaptionMapID string `json:"caption_map_id"`
	IsCaptionMap int    `json:"is_caption_map"`
	SyncTime     int    `json:"sync_time"`
}

type SubtitleAutoload struct {
	SID      string `json:"sid"`
	Language string `json:"language"`
	Title    string `json:"title"`
	URL      string `json:"url"`
	Type     string `json:"type"`
}

type VideoSubtitleResp struct {
	Autoload *SubtitleAutoload `json:"autoload"`
	List     []SubtitleItem    `json:"list"`
}

// VideoSubtitle https://www.yuque.com/115yun/open/nx076h3glapoyh7u
func (c *Client) VideoSubtitle(ctx context.Context, pickCode string) (*VideoSubtitleResp, error) {
	var resp VideoSubtitleResp
	_, err := c.AuthRequest(ctx, ApiVideoSubtitle, http.MethodGet, &resp, ReqWithQuery(Form{
		"pick_code": pickCode,
	}))
	if err != nil {
		return nil, err
	}
	return &resp, err
}

// VideoPush https://www.yuque.com/115yun/open/nxt8r1qcktmg3oan
func (c *Client) VideoPush(ctx context.Context, pickCode, op string) error {
	var resp any
	_, err := c.AuthRequest(ctx, ApiVideoPush, http.MethodPost, &resp, ReqWithForm(Form{
		"pick_code": pickCode,
		"op":        op,
	}))
	return err
}
