package handlers

import (
	"errors"
	"fmt"
	"weixinchat/config"
	"weixinchat/gpt"
	"weixinchat/pkg/logger"
	"weixinchat/service"

	"github.com/eatmoreapple/openwechat"
	"strings"
)

var _ MessageHandlerInterface = (*GroupMessageHandler)(nil)


type GroupMessageHandler struct {

	self *openwechat.Self
	
	group *openwechat.Group
	
	msg *openwechat.Message
	
	sender *openwechat.User
	
	service service.UserServiceInterface
}

func GroupMessageContextHandler() func(ctx *openwechat.MessageContext) {
	return func(ctx *openwechat.MessageContext) {
		msg := ctx.Message
		
		handler, err := NewGroupMessageHandler(msg)
		if err != nil {
			logger.Warning(fmt.Sprintf("init group message handler error: %s", err))
			return
		}

		
		err = handler.handle()
		if err != nil {
			logger.Warning(fmt.Sprintf("handle group message error: %s", err))
		}
	}
}


func NewGroupMessageHandler(msg *openwechat.Message) (MessageHandlerInterface, error) {
	sender, err := msg.Sender()
	if err != nil {
		return nil, err
	}
	group := &openwechat.Group{User: sender}
	groupSender, err := msg.SenderInGroup()
	if err != nil {
		return nil, err
	}

	userService := service.NewUserService(c, groupSender)
	handler := &GroupMessageHandler{
		self:    sender.Self,
		msg:     msg,
		group:   group,
		sender:  groupSender,
		service: userService,
	}
	return handler, nil

}


func (g *GroupMessageHandler) handle() error {
	if g.msg.IsText() {
		return g.ReplyText()
	}
	return nil
}


func (g *GroupMessageHandler) ReplyText() error {
	reply, err := gpt.Text("")
	logger.Info(fmt.Sprintf("Received Group %v Text Msg : %v", g.group.NickName, g.msg.Content))

	
	if !g.msg.IsAt() {
		return nil
	}
	g.msg.IsJoinGroup()


	requestText := g.getRequestText()
	if requestText == "" {
		reply, err = gpt.Text("艾特小鸿干嘛！能不能说点话")
		return nil
	}


	t, e := config.FData(requestText)
	if e {
		reply, err = gpt.Text(t)
	} else {
		reply, err = gpt.Completions(requestText)
	}

	if err != nil {
		
		errMsg := fmt.Sprintf("gpt request error: %v", err)
		_, err = g.msg.ReplyText(errMsg)
		if err != nil {
			return errors.New(fmt.Sprintf("response group error: %v ", err))
		}
		return err
	}

	
	g.service.SetUserSessionContext(requestText, reply)
	_, err = g.msg.ReplyText(g.buildReplyText(reply))
	if err != nil {
		return errors.New(fmt.Sprintf("response user error: %v ", err))
	}


	return err
}


func (g *GroupMessageHandler) getRequestText() string {
	
	requestText := strings.TrimSpace(g.msg.Content)
	requestText = strings.Trim(g.msg.Content, "\n")

	replaceText := "@" + g.self.NickName
	requestText = strings.TrimSpace(strings.ReplaceAll(g.msg.Content, replaceText, ""))
	if requestText == "" {
		return ""
	}


	sessionText := g.service.GetUserSessionContext()
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
		requestText = requestText + "。" 
	}


	return requestText
}


func (g *GroupMessageHandler) buildReplyText(reply string) string {

	atText := "@" + g.sender.NickName
	textSplit := strings.Split(reply, "\n\n")
	if len(textSplit) > 1 {
		trimText := textSplit[0]
		reply = strings.Trim(reply, trimText)
	}
	reply = strings.TrimSpace(reply)
	if reply == "" {
		return atText + " 请求得不到任何有意义的回复，请具体提出问题。"
	}

	replaceText := "@" + g.self.NickName
	question := strings.TrimSpace(strings.ReplaceAll(g.msg.Content, replaceText, ""))
	reply = atText + "\n" + question + "\n --------------------------------\n" + reply + openwechat.Emoji.Doge
	reply = strings.Trim(reply, "\n")

	return reply
}
