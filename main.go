package main

import (
	"github.com/zeromicro/go-zero/core/logx"
	"open-assistant-helper-go/model"
	"time"
)

func main() {
	_ = logx.SetUp(logx.LogConf{
		Mode:     "console",
		Encoding: "plain",
	})
	err := model.LoadConfig()
	if err != nil {
		panic(err)
	}
	InitChatGPTClient(model.Conf.ApiKey)
	i := 0
	for true {
		err = StartTask()
		if err != nil {
			logx.Error("Error while doing task: ", err)
		}
		if i++; i == 10 {
			i = 0
			err = RefreshCookie()
			if err != nil {
				logx.Error("Error while refreshing cookie: ", err)
			}
		}
		time.Sleep(2 * time.Second)
	}
}
