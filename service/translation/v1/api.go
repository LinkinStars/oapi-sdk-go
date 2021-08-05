// Code generated by lark suite oapi sdk gen
package v1

import (
	"github.com/larksuite/oapi-sdk-go"
	"github.com/larksuite/oapi-sdk-go/api"
	"github.com/larksuite/oapi-sdk-go/api/core/request"
)

type Service struct {
	conf  lark.Config
	Texts *TextService
}

func NewService(conf lark.Config) *Service {
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
	ctx   *lark.Context
	texts *TextService
	body  *TextTranslateReqBody
	opts  []lark.APIRequestOpt
}

func (rc *TextTranslateReqCall) Do() (*TextTranslateResult, error) {
	var result = &TextTranslateResult{}
	req := request.NewRequest("/open-apis/translation/v1/text/translate", "POST",
		[]request.AccessTokenType{request.AccessTokenTypeTenant}, rc.body, result, rc.opts...)
	err := api.Send(rc.ctx, rc.texts.service.conf, req)
	return result, err
}

func (texts *TextService) Translate(ctx *lark.Context, body *TextTranslateReqBody, opts ...lark.APIRequestOpt) *TextTranslateReqCall {
	return &TextTranslateReqCall{
		ctx:   ctx,
		texts: texts,
		body:  body,
		opts:  opts,
	}
}

type TextDetectReqCall struct {
	ctx   *lark.Context
	texts *TextService
	body  *TextDetectReqBody
	opts  []lark.APIRequestOpt
}

func (rc *TextDetectReqCall) Do() (*TextDetectResult, error) {
	var result = &TextDetectResult{}
	req := request.NewRequest("/open-apis/translation/v1/text/detect", "POST",
		[]request.AccessTokenType{request.AccessTokenTypeTenant}, rc.body, result, rc.opts...)
	err := api.Send(rc.ctx, rc.texts.service.conf, req)
	return result, err
}

func (texts *TextService) Detect(ctx *lark.Context, body *TextDetectReqBody, opts ...lark.APIRequestOpt) *TextDetectReqCall {
	return &TextDetectReqCall{
		ctx:   ctx,
		texts: texts,
		body:  body,
		opts:  opts,
	}
}
