package sdk

import (
	"context"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

type AddOfflineTaskURIsResp = struct {
	State    bool   `json:"state"`     // 链接任务添加状态，成功true；失败false
	Code     int64  `json:"code"`      // 链接任务状态码，成功返回0
	Message  string `json:"message"`   // 链接任务状态描述，成功返回空字符串
	InfoHash string `json:"info_hash"` // 链接任务sha1，只有任务成功的时候才会返回
	URL      string `json:"url"`       // 链接任务url
}

// AddOfflineTaskURIs  https://www.yuque.com/115yun/open/zkyfq2499gdn3mty
func (c *Client) AddOfflineTaskURIs(ctx context.Context, uris []string, saveDirID string) ([]string, error) {
	var resp []AddOfflineTaskURIsResp

	if len(uris) == 0 {
		return nil, fmt.Errorf("uris is empty")
	}
	urlsStr := strings.Join(uris, "\n")

	_, err := c.AuthRequest(ctx, ApiAddOffline, http.MethodPost, &resp, ReqWithForm(Form{
		"urls":       urlsStr,
		"wp_path_id": saveDirID,
	}))
	if err != nil {
		return nil, err
	}

	var hashes []string
	for _, item := range resp {
		if item.State && item.InfoHash != "" {
			hashes = append(hashes, item.InfoHash)
		}
	}

	return hashes, err
}

// DeleteOfflineTask  https://www.yuque.com/115yun/open/pmgwc86lpcy238nw
func (c *Client) DeleteOfflineTask(ctx context.Context, infoHash string, deleteFiles bool) error {
	var resp []string

	form := Form{
		"info_hash":       infoHash,
		"del_source_file": "0",
	}
	if deleteFiles {
		form["del_source_file"] = "1"
	}

	_, err := c.AuthRequest(ctx, ApiDeleteOffline, http.MethodPost, &resp, ReqWithForm(form))
	return err
}

type OfflineTaskListResp struct {
	Page      int           `json:"page"`       // 当前第几页
	PageCount int           `json:"page_count"` // 总页数
	Count     int           `json:"count"`      // 总数量
	Tasks     []OfflineTask `json:"tasks"`      // 云下载任务列表
}

type OfflineTask struct {
	InfoHash     string `json:"info_hash"`      // 任务 sha1
	AddTime      int64  `json:"add_time"`       // 添加时间戳
	PercentDone  int    `json:"percentDone"`    // 下载进度
	Size         int64  `json:"size"`           // 总大小（字节）
	Name         string `json:"name"`           // 任务名
	LastUpdate   int64  `json:"last_update"`    // 最后更新时间戳
	FileID       string `json:"file_id"`        // 文件或文件夹 ID
	DeleteFileID string `json:"delete_file_id"` // 删除源文件时需传递的 ID
	Status       int    `json:"status"`         // 任务状态（-1失败，0分配中，1下载中，2成功）
	URL          string `json:"url"`            // 链接 URL
	WpPathID     string `json:"wp_path_id"`     // 所在父文件夹 ID
	Def2         int    `json:"def2"`           // 视频清晰度（1~5, 100）
	PlayLong     int    `json:"play_long"`      // 视频时长
	CanAppeal    int    `json:"can_appeal"`     // 是否可申诉
}

func (t *OfflineTask) IsTodo() bool {
	return t.Status == 0
}

func (t *OfflineTask) IsRunning() bool {
	return t.Status == 1
}

func (t *OfflineTask) IsDone() bool {
	return t.Status == 2
}

func (t *OfflineTask) IsFailed() bool {
	return t.Status == -1
}

func (t *OfflineTask) GetStatus() string {
	if t.IsTodo() {
		return "准备开始离线下载"
	}
	if t.IsDone() {
		return "离线下载完成"
	}
	if t.IsFailed() {
		return "离线下载失败"
	}
	if t.IsRunning() {
		return "离线任务下载中"
	}
	return fmt.Sprintf("未知状态: %d", t.Status)
}

// OfflineTaskList  https://www.yuque.com/115yun/open/av2mluz7uwigz74k
func (c *Client) OfflineTaskList(ctx context.Context, page int64) (*OfflineTaskListResp, error) {
	var resp OfflineTaskListResp

	_, err := c.AuthRequest(ctx, ApiOfflineList, http.MethodGet, &resp, ReqWithForm(Form{
		"page": strconv.FormatInt(page, 10),
	}))
	return &resp, err
}
