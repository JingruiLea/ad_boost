package lark

import (
	"context"
	"github.com/JingruiLea/ad_boost/common/logs"
	"github.com/chyroc/lark"
	"github.com/gin-gonic/gin"
	"net/http"
)

const AppID = "cli_a457b42cbbf8d00e"
const AppSec = "HBKzMGefcvsk2HGn0q2lCbNSV5xtm5Sr"
const EncryptKey = "1iJcXP40z2rOPBYf3VqD6codBZItkbRO"
const VerificationToken = "rGCakotMu664p648Quqqkh6r5zjWThWE"

const AdBoostChatID = "oc_37e74ef798859f1b40bf25a63777ab61"

var Bot *lark.Lark

func init() {
	Bot = lark.New(
		lark.WithAppCredential(AppID, AppSec),
		lark.WithEventCallbackVerify(EncryptKey, VerificationToken),
	)
	Bot.EventCallback.HandlerEventV2IMMessageReceiveV1(func(ctx context.Context, cli *lark.Lark, schema string, header *lark.EventHeaderV2, event *lark.EventV2IMMessageReceiveV1) (string, error) {
		logs.CtxInfof(ctx, "receive message, user_id: %s, content: %s", event.Message.ChatID, event.Message.Content)
		HandleTextMsg(ctx, event.Message.Content, cli)
		return "", nil
	})
}

func Register(ctx context.Context, w gin.ResponseWriter, r *http.Request) {
	Bot.EventCallback.ListenCallback(ctx, r.Body, w)
}

func SendRoomMessage(ctx context.Context, msg string) {
	_, _, err := Bot.Message.Send().ToChatID(AdBoostChatID).SendText(ctx, msg)
	if err != nil {
		logs.CtxErrorf(ctx, "send message failed, err: %v", err)
	}
}

var handlers = make([]func(string) string, 0)

func HandleTextMsg(ctx context.Context, msg string, cli *lark.Lark) {
	for _, h := range handlers {
		resp := h(msg)
		if resp != "" {
			_, _, err := cli.Message.Send().ToChatID(AdBoostChatID).SendText(ctx, resp)
			if err != nil {
				logs.CtxErrorf(ctx, "send message failed, err: %v", err)
			}
		}
	}
}

func RegisterTextHandler(ctx context.Context, h func(string) string) {
	handlers = append(handlers, h)
}
