package handlers

import (
	"errors"
	"fmt"
	"github.com/eatmoreapple/openwechat"
	"strings"
	"weixinchat/config"
	"weixinchat/gpt"
	"weixinchat/pkg/logger"
	"weixinchat/service"
)

var _ MessageHandlerInterface = (*UserMessageHandler)(nil)


type UserMessageHandler struct {
	
	msg *openwechat.Message
	
	sender *openwechat.User

	service service.UserServiceInterface
}

func UserMessageContextHandler() func(ctx *openwechat.MessageContext) {
	return func(ctx *openwechat.MessageContext) {
		msg := ctx.Message
		handler, err := NewUserMessageHandler(msg)
		if err != nil {
			logger.Warning(fmt.Sprintf("init user message handler error: %s", err))
		}


		err = handler.handle()
		if err != nil {
			logger.Warning(fmt.Sprintf("handle user message error: %s", err))
		}
	}
}


func NewUserMessageHandler(message *openwechat.Message) (MessageHandlerInterface, error) {
	sender, err := message.Sender()
	if err != nil {
		return nil, err
	}
	userService := service.NewUserService(c, sender)
	handler := &UserMessageHandler{
		msg:     message,
		sender:  sender,
		service: userService,
	}

	return handler, nil
}


func (h *UserMessageHandler) handle() error {
	if h.msg.IsText() {
		return h.ReplyText()
	}
	return nil
}


func (h *UserMessageHandler) ReplyText() error {
	logger.Info(fmt.Sprintf("Received User %v Text Msg : %v", h.sender.NickName, h.msg.Content))
	
	requestText := h.getRequestText()
	if requestText == "" {
		logger.Info("user message is null")
		return nil
	}



	reply, err := gpt.Text("")
	t, e := config.FData(requestText)
	println(t, e)
	if e {
		reply, err = gpt.Text(t)
	} else {
		reply, err = gpt.Completions(requestText)
	}
	if err != nil {

		errMsg := fmt.Sprintf("gpt request error: %v", err)
		_, err = h.msg.ReplyText(errMsg)
		if err != nil {
			return errors.New(fmt.Sprintf("response user error: %v ", err))
		}
		return err
	}


	h.service.SetUserSessionContext(requestText, reply)
	_, err = h.msg.ReplyText(buildUserReply(reply))
	if err != nil {
		return errors.New(fmt.Sprintf("response user error: %v ", err))
	}


	return err
}

func (h *UserMessageHandler) getRequestText() string {

	requestText := strings.TrimSpace(h.msg.Content)
	requestText = strings.Trim(h.msg.Content, "\n")

	
	sessionText := h.service.GetUserSessionContext()
	if sessionText != "" {
		requestText = sessionText + "\n" + requestText
	}
	if len(requestText) >= 4000 {
		requestText = requestText[:4000]
	}


	punctuation := ",.;!?，。！？、…"
	runeRequestText := []rune(requestText)
	lastChar := string(runeRequestText[len(runeRequestText)-1:])
	if strings.Index(punctuation, lastChar) < 0 {
		requestText = requestText + "？" // 判断最后字符是否加了标点，没有的话加上句号，避免openai自动补齐引起混乱。
	}


	return requestText
}

func buildUserReply(reply string) string {

	textSplit := strings.Split(reply, "\n\n")
	if len(textSplit) > 1 {
		trimText := textSplit[0]
		reply = strings.Trim(reply, trimText)
	}
	reply = strings.TrimSpace(reply)

	reply = strings.TrimSpace(reply)
	if reply == "" {
		return "请求得不到任何有意义的回复，请具体提出问题。"
	}


	reply = config.LoadConfig().ReplyPrefix + "\n" + reply + openwechat.Emoji.Doge
	reply = strings.Trim(reply, "\n")


	return reply
}
