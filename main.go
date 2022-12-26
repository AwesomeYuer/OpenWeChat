package main

import (
	"fmt"

	"github.com/eatmoreapple/openwechat"
)

func main() {
	messageHandler := func(msg *openwechat.Message) {
		//fmt.Println('===' + msg + "===")

		if msg.IsText() {
			fmt.Println(msg.Content)
			msg.ReplyText("pong:" + msg.Content)
		}

	}
	//bot := openwechat.DefaultBot()
	bot := openwechat.DefaultBot(openwechat.Desktop)

	// 注册消息处理函数
	bot.MessageHandler = messageHandler
	// 设置默认的登录回调
	// 可以设置通过该uuid获取到登录的二维码
	bot.UUIDCallback = openwechat.PrintlnQrcodeUrl
	// 登录
	if err := bot.Login(); err != nil {
		fmt.Println(err)
		return
	}
	// 阻塞主程序,直到用户退出或发生异常
	bot.Block()
}
