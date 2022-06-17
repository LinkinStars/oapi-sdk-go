// Package translation code generated by oapi sdk gen
package translation

import (
	"context"
	"net/http"

	"github.com/larksuite/oapi-sdk-go/core"
)

/**
构建业务域服务实例
**/
func NewService(httpClient *http.Client, config *core.Config) *TranslationService {
	t := &TranslationService{httpClient: httpClient, config: config}
	t.Text = &text{service: t}
	return t
}

/**
业务域服务定义
**/
type TranslationService struct {
	httpClient *http.Client
	config     *core.Config
	Text       *text
}

/**
资源服务定义
**/
type text struct {
	service *TranslationService
}

/**
资源服务方法定义
**/
func (t *text) Detect(ctx context.Context, req *DetectTextReq, options ...core.RequestOptionFunc) (*DetectTextResp, error) {
	// 发起请求
	rawResp, err := core.SendRequest(ctx, t.service.config, http.MethodPost,
		"/open-apis/translation/v1/text/detect", []core.AccessTokenType{core.AccessTokenTypeTenant}, req, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &DetectTextResp{RawResponse: rawResp}
	err = rawResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (t *text) Translate(ctx context.Context, req *TranslateTextReq, options ...core.RequestOptionFunc) (*TranslateTextResp, error) {
	// 发起请求
	rawResp, err := core.SendRequest(ctx, t.service.config, http.MethodPost,
		"/open-apis/translation/v1/text/translate", []core.AccessTokenType{core.AccessTokenTypeTenant}, req, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &TranslateTextResp{RawResponse: rawResp}
	err = rawResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
