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
	for true {
		err = StartTask()
		if err != nil {
			logx.Error("Error while doing task: ", err)
		}
		time.Sleep(2 * time.Second)
	}
	// 	s, err := GetLabelsFromChatGPT(`{
	//     "spam": 0,
	//     "fails_task": 0,
	//     "lang_mismatch": 0,
	//     "not_appropriate": 0,
	//     "pii": 0,
	//     "hate_speech": 0,
	//     "sexual_content": 0,
	//     "quality": 0.75,
	//     "helpfulness": 1,
	//     "creativity": 0.5,
	//     "humor": 0,
	//     "toxicity": 0,
	//     "violence": 0
	// }`)
	// 	if err != nil {
	// 		panic(err)
	// 	}
	// 	for l, v := range s {
	// 		println(l, " ", v)
	// 	}
}
