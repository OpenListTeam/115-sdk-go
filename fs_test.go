package sdk

import (
	"encoding/json"
	"errors"
	"testing"
)

func TestGetFolderInfoRespUnmarshalObject(t *testing.T) {
	var resp GetFolderInfoResp
	err := json.Unmarshal([]byte(`{"file_id":"1","file_name":"dir","file_category":"0"}`), &resp)
	if err != nil {
		t.Fatal(err)
	}
	if resp.FileID != "1" || resp.FileName != "dir" || resp.FileCategory != "0" {
		t.Fatalf("unexpected response: %+v", resp)
	}
}

func TestGetFolderInfoRespUnmarshalArray(t *testing.T) {
	var resp GetFolderInfoResp
	err := json.Unmarshal([]byte(`[{"file_id":"1","file_name":"dir","file_category":"0"}]`), &resp)
	if err != nil {
		t.Fatal(err)
	}
	if resp.FileID != "1" || resp.FileName != "dir" || resp.FileCategory != "0" {
		t.Fatalf("unexpected response: %+v", resp)
	}
}

func TestGetFolderInfoRespUnmarshalEmptyArray(t *testing.T) {
	var resp GetFolderInfoResp
	err := json.Unmarshal([]byte(`[]`), &resp)
	if !errors.Is(err, ErrObjectNotFound) {
		t.Fatalf("expected ErrObjectNotFound, got %v", err)
	}
}
