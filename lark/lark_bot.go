package lark

import (
	"context"
	"github.com/JingruiLea/ad_boost/dal/redis_dal"
	"github.com/chyroc/lark"
	"github.com/gin-gonic/gin"
	jsoniter "github.com/json-iterator/go"
	"github.com/sirupsen/logrus"
	"net/http"
)

// 定义自己的logger, 避免循环引用
var logs Logger

type Logger struct {
}

func (*Logger) CtxInfof(ctx context.Context, format string, args ...interface{}) {
	logrus.WithContext(ctx).Infof(format, args)
}

func (*Logger) CtxErrorf(ctx context.Context, format string, args ...interface{}) {
	logrus.WithContext(ctx).Errorf(format, args)
}

const AppID = "cli_a457b42cbbf8d00e"
const AppSec = "HBKzMGefcvsk2HGn0q2lCbNSV5xtm5Sr"
const EncryptKey = "1iJcXP40z2rOPBYf3VqD6codBZItkbRO"
const VerificationToken = "rGCakotMu664p648Quqqkh6r5zjWThWE"

const AdBoostChatID = "oc_586a7dec27b417492a5b0ae7ba11a75e"
const AlertChatID = "oc_9634312c93f56e516d6501e2425d8f32"

var Bot *lark.Lark

func Init() {
	Bot = lark.New(
		lark.WithAppCredential(AppID, AppSec),
		lark.WithEventCallbackVerify(EncryptKey, VerificationToken),
	)
	Bot.EventCallback.HandlerEventV2IMMessageReceiveV1(func(ctx context.Context, cli *lark.Lark, schema string, header *lark.EventHeaderV2, event *lark.EventV2IMMessageReceiveV1) (string, error) {
		logs.CtxInfof(ctx, "receive message, msg_id: %s, content: %s", event.Message.MessageID, event.Message.Content)
		dup, err := redis_dal.CheckDuplicateRequest(ctx, event.Message.MessageID)
		if err != nil {
			logs.CtxErrorf(ctx, "check duplicate request failed, err: %v", err)
		}
		if dup {
			logs.CtxInfof(ctx, "duplicate request, ignore:%s", event.Message.Content)
			return "", nil
		}
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

func SendAlertMessage(ctx context.Context, msg string) {
	_, _, err := Bot.Message.Send().ToChatID(AlertChatID).SendText(ctx, msg)
	if err != nil {
		logs.CtxErrorf(ctx, "send message failed, err: %v", err)
	}
}

var handlers = make([]func(string) string, 0)

func HandleTextMsg(ctx context.Context, msg string, cli *lark.Lark) {
	var text = make(map[string]string)
	err := jsoniter.UnmarshalFromString(msg, &text)
	if err != nil {
		logs.CtxErrorf(ctx, "unmarshal msg failed, err: %v", err)
		return
	}
	msg = text["text"]
	for _, h := range handlers {
		hCopy := h
		resp := hCopy(msg)
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
