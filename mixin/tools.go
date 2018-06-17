package mixin

import (
	"bytes"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/ibigbug/caoliu/config"
	"github.com/mixinmessenger/bot-api-go-client"
)

type RequestAttachmentResponse struct {
	Data struct {
		AttachmentId string `json:"attachment_id"`
		UploadUrl    string `json:"upload_url"`
		ViewUrl      string `json:"view_url"`
	}
}

type SendImageRequest struct {
	AttachmentId string `json:"attachment_id,omitempty"`
	MimeType     string `json:"mime_type,omitempty"`
	Width        int    `json:"width,omitempty"`
	Height       int    `json:"height,omitempty"`
	Size         int64  `json:"size,omitempty"`
	Thumbnail    string `json:"thumbnail,omitempty"`
}

func UploadAttachment(content []byte) RequestAttachmentResponse {
	token, err := bot.SignAuthenticationToken(config.ClientID, config.SessionID, config.PrivateKey, "POST", "/attachments", "")
	if err != nil {
		panic(err)
	}
	req, err := http.NewRequest("POST", "https://api.mixin.one/attachments", nil)
	if err != nil {
		panic(err)
	}
	req.Header.Set("Authorization", "Bearer "+token)
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()
	var rv RequestAttachmentResponse
	json.NewDecoder(res.Body).Decode(&rv)

	req, _ = http.NewRequest("PUT", rv.Data.UploadUrl, bytes.NewReader(content))
	req.Header.Set("Content-Type", "application/octet-stream")
	req.Header.Set("x-amz-acl", "public-read")
	req.Header.Set("Content-Length", strconv.Itoa(len(content)))
	res, err = http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()
	return rv
}
