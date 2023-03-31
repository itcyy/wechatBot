package handlers

import (
	"fmt"
	"github.com/eatmoreapple/openwechat"
	"weixinchat/pkg/logger"
	"weixinchat/service"
)

var _ MessageHandlerInterface = (*TokenMessageHandler)(nil)

type TokenMessageHandler struct {
	
	msg *openwechat.Message

	sender *openwechat.User

	service service.UserServiceInterface
}

func TokenMessageContextHandler() func(ctx *openwechat.MessageContext) {
	return func(ctx *openwechat.MessageContext) {
		msg := ctx.Message
	
		handler, err := NewTokenMessageHandler(msg)
		if err != nil {
			logger.Warning(fmt.Sprintf("init token message handler error: %s", err))
		}

		err = handler.handle()
		if err != nil {
			logger.Warning(fmt.Sprintf("handle token message error: %s", err))
		}

	}
}

func NewTokenMessageHandler(msg *openwechat.Message) (MessageHandlerInterface, error) {
	sender, err := msg.Sender()
	if err != nil {
		return nil, err
	}
	if msg.IsComeFromGroup() {
		sender, err = msg.SenderInGroup()
	}
	userService := service.NewUserService(c, sender)
	handler := &TokenMessageHandler{
		msg:     msg,
		sender:  sender,
		service: userService,
	}

	return handler, nil
}


func (t *TokenMessageHandler) handle() error {
	return t.ReplyText()
}


func (t *TokenMessageHandler) ReplyText() error {
	logger.Info("user clear token")
	t.service.ClearUserSessionContext()
	var err error
	if t.msg.IsComeFromGroup() {
		if !t.msg.IsAt() {
			return err
		}
		atText := "@" + t.sender.NickName + "上下文已经清空，请问下一个问题。"
		_, err = t.msg.ReplyText(atText)
	} else {
		_, err = t.msg.ReplyText("上下文已经清空，请问下一个问题。")
	}
	return err
}
