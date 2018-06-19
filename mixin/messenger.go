package mixin

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"time"

	"github.com/ibigbug/caoliu/config"
	"github.com/mixinmessenger/bot-api-go-client"
)

func init() {
	rand.Seed(time.Now().Unix())
}

type Messenger struct{}

func (m Messenger) OnMessage(ctx context.Context, mc *bot.MessageContext, msg bot.MessageView, uid string) error {
	files := ioutil.ReadDir(config.ImageDir)
	n := rand.Intn(len(files))
	var image, err = os.Open(files[n].Name())
	if err != nil {
		panic(err)
	}
	bytes, err := ioutil.ReadAll(image)
	if err != nil {
		panic(err)
	}
	image.Close()
	rv := UploadAttachment(bytes)
	param := SendImageRequest{
		AttachmentId: rv.Data.AttachmentId,
		MimeType:     "image/jpeg",
		Size:         int64(len(bytes)),
		Width:        10,
		Height:       10,
	}
	paramBytes, _ := json.Marshal(param)
	fmt.Println(string(paramBytes))
	paramBase64 := base64.StdEncoding.EncodeToString(paramBytes)
	if msg.Category != bot.MessageCategorySystemAccountSnapshot && msg.Category != bot.MessageCategorySystemConversation && msg.ConversationId == bot.UniqueConversationId(config.ClientID, msg.UserId) {
		if msg.Category == PLAIN_TEXT {
			data, err := base64.StdEncoding.DecodeString(msg.Data)
			if err != nil {
				return bot.BlazeServerError(ctx, err)
			}
			if string(data) == "caoliu" {
				fmt.Println(bot.SendImage(ctx, mc, msg, paramBase64))
			} else {
				if err := bot.SendPlainText(ctx, mc, msg, "please send me caoliu"); err != nil {
					return bot.BlazeServerError(ctx, err)
				}
			}
		}
	}
	return nil
}
