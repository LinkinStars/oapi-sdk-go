// Code generated by lark suite oapi sdk gen
package v1

import (
	"github.com/larksuite/oapi-sdk-go/api"
	"github.com/larksuite/oapi-sdk-go/api/core/request"
	"github.com/larksuite/oapi-sdk-go/core"
	"github.com/larksuite/oapi-sdk-go/core/config"
)

type Service struct {
	conf  *config.Config
	Texts *TextService
}

func NewService(conf *config.Config) *Service {
	s := &Service{
		conf: conf,
	}
	s.Texts = newTextService(s)
	return s
}

type TextService struct {
	service *Service
}

func newTextService(service *Service) *TextService {
	return &TextService{
		service: service,
	}
}

type TextTranslateReqCall struct {
	ctx    *core.Context
	texts  *TextService
	body   *TextTranslateReqBody
	optFns []request.OptFn
}

func (rc *TextTranslateReqCall) Do() (*TextTranslateResult, error) {
	var result = &TextTranslateResult{}
	req := request.NewRequest("/open-apis/translation/v1/text/translate", "POST",
		[]request.AccessTokenType{request.AccessTokenTypeTenant}, rc.body, result, rc.optFns...)
	err := api.Send(rc.ctx, rc.texts.service.conf, req)
	return result, err
}

func (texts *TextService) Translate(ctx *core.Context, body *TextTranslateReqBody, optFns ...request.OptFn) *TextTranslateReqCall {
	return &TextTranslateReqCall{
		ctx:    ctx,
		texts:  texts,
		body:   body,
		optFns: optFns,
	}
}

type TextDetectReqCall struct {
	ctx    *core.Context
	texts  *TextService
	body   *TextDetectReqBody
	optFns []request.OptFn
}

func (rc *TextDetectReqCall) Do() (*TextDetectResult, error) {
	var result = &TextDetectResult{}
	req := request.NewRequest("/open-apis/translation/v1/text/detect", "POST",
		[]request.AccessTokenType{request.AccessTokenTypeTenant}, rc.body, result, rc.optFns...)
	err := api.Send(rc.ctx, rc.texts.service.conf, req)
	return result, err
}

func (texts *TextService) Detect(ctx *core.Context, body *TextDetectReqBody, optFns ...request.OptFn) *TextDetectReqCall {
	return &TextDetectReqCall{
		ctx:    ctx,
		texts:  texts,
		body:   body,
		optFns: optFns,
	}
}
