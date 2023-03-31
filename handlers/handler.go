package handlers

import (
	"fmt"
	"github.com/eatmoreapple/openwechat"
	"github.com/patrickmn/go-cache"
	"github.com/skip2/go-qrcode"
	"log"
	"runtime"
	"strings"
	"time"
	"weixinchat/config"
	"weixinchat/pkg/logger"
)

var c = cache.New(config.LoadConfig().SessionTimeout, time.Minute*5)


type MessageHandlerInterface interface {
	handle() error
	ReplyText() error
}


func QrCodeCallBack(uuid string) {
	if runtime.GOOS == "windows" {
	
		openwechat.PrintlnQrcodeUrl(uuid)
	} else {
		log.Println("login in linux")
		url := "https://login.weixin.qq.com/l/" + uuid
		log.Printf("如果二维码无法扫描，请缩小控制台尺寸，或更换命令行工具，缩小二维码像素")
		q, _ := qrcode.New(url, qrcode.High)
		fmt.Println(q.ToSmallString(true))
	}
}

func NewHandler() (msgFunc func(msg *openwechat.Message), err error) {
	dispatcher := openwechat.NewMessageMatchDispatcher()

	
	dispatcher.RegisterHandler(func(message *openwechat.Message) bool {
		return strings.Contains(message.Content, config.LoadConfig().SessionClearToken)
	}, TokenMessageContextHandler())

	
	dispatcher.RegisterHandler(func(message *openwechat.Message) bool {
		return message.IsSendByGroup()
	}, GroupMessageContextHandler())

	dispatcher.RegisterHandler(func(message *openwechat.Message) bool {
		return message.IsFriendAdd()
	}, func(ctx *openwechat.MessageContext) {
		msg := ctx.Message
		if config.LoadConfig().AutoPass {
			_, err := msg.Agree("")
			if err != nil {
				logger.Warning(fmt.Sprintf("add friend agree error : %v", err))
				return
			}
		}
	})


	dispatcher.RegisterHandler(func(message *openwechat.Message) bool {
		return !(strings.Contains(message.Content, config.LoadConfig().SessionClearToken) || message.IsSendByGroup() || message.IsFriendAdd())
	}, UserMessageContextHandler())
	return openwechat.DispatchMessage(dispatcher), nil
}
