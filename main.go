package main

import (
	"fmt"
	"log"

	"github.com/eatmoreapple/openwechat"
	"github.com/skip2/go-qrcode"
)

func main() {
	messageHandler := func(msg *openwechat.Message) {
		//fmt.Println('===' + msg + "===")

		if msg.IsText() {
			fmt.Println(msg.Content)
			msg.ReplyText("pong: " + msg.Content)
		}

	}
	//bot := openwechat.DefaultBot()
	bot := openwechat.DefaultBot(openwechat.Desktop)

	// 注册消息处理函数
	bot.MessageHandler = messageHandler
	// 设置默认的登录回调
	// 可以设置通过该uuid获取到登录的二维码

	bot.UUIDCallback = QrCodeCallBack
	// 登录
	if err := bot.Login(); err != nil {
		fmt.Println(err)
		return
	}
	// 阻塞主程序,直到用户退出或发生异常
	bot.Block()
}

// QrCodeCallBack 登录扫码回调，
func QrCodeCallBack(uuid string) {
	log.Println("login in")
	q, _ := qrcode.New("https://login.weixin.qq.com/l/"+uuid, qrcode.Highest)
	fmt.Println(q.ToString(true))
}
