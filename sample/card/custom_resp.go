package main

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/larksuite/oapi-sdk-go/card"
	"github.com/larksuite/oapi-sdk-go/card/model"
	"github.com/larksuite/oapi-sdk-go/core"
	"github.com/larksuite/oapi-sdk-go/core/tools"
	"github.com/larksuite/oapi-sdk-go/sample/configs"
)

func main() {
	// for redis store and logrus
	// var conf = configs.TestConfigWithLogrusAndRedisStore(core.DomainFeiShu)
	// var conf = configs.TestConfig("https://open.feishu.cn")
	var conf = configs.TestConfig(core.DomainFeiShu)

	card.SetHandler(conf, func(coreCtx *core.Context, card *model.Card) (interface{}, error) {
		fmt.Println(coreCtx.GetRequestID())
		fmt.Println(tools.Prettify(card.Action))

		body := make(map[string]interface{})
		body["content"] = "hello"

		i18n := make(map[string]string)
		i18n["zh_cn"] = "你好"
		i18n["en_us"] = "hello"
		i18n["ja_jp"] = "こんにちは"
		body["i18n"] = i18n

		b, err := json.Marshal(body)
		if err != nil {
			return nil, err
		}

		return &model.CustomResp{
			StatusCode: 400,
			Body:       b,
		}, nil
	})

	header := make(map[string][]string)
	// from http request header
	// and "X-Lark-Request-Timestamp"/"X-Lark-Request-Nonce"/"X-Lark-Signature" validate request, Required!

	// header["X-Request-Id"] = []string{"63278309j-yuewuyeu-7828389"}
	// header["X-Lark-Request-Timestamp"] = []string{"Monday, 09-Nov-20 23:33:53 CST"}
	// header["X-Lark-Request-Nonce"] = []string{"0404f57f-102e-4c91-bb32-a501ad0b7600"}
	// header["X-Lark-Signature"] = []string{"26cb59f4f5a91c4147d0xxxxxxxxxc4a36fb2c"}
	// header["X-Refresh-Token"] = []string{"acc4d5f2-4bc6-4394-a9d4-45e168fcde97"}

	req := &core.OapiRequest{
		Ctx:    context.Background(),
		Header: core.NewOapiHeader(header),
		Body:   "{\"token\":\"sss\",\"type\":\"event_callback\"}", // from http request body
	}
	resp := card.Handle(conf, req)
	fmt.Println(tools.Prettify(resp))
}
