// Code generated by lark suite oapi sdk gen
package v1

import (
	"github.com/larksuite/oapi-sdk-go"
	"github.com/larksuite/oapi-sdk-go/api"
	"github.com/larksuite/oapi-sdk-go/api/core/request"
)

type Service struct {
	conf    lark.Config
	Authens *AuthenService
}

func NewService(conf lark.Config) *Service {
	s := &Service{
		conf: conf,
	}
	s.Authens = newAuthenService(s)
	return s
}

type AuthenService struct {
	service *Service
}

func newAuthenService(service *Service) *AuthenService {
	return &AuthenService{
		service: service,
	}
}

type AuthenAccessTokenReqCall struct {
	ctx     *lark.Context
	authens *AuthenService
	body    *AuthenAccessTokenReqBody
	opts    []lark.APIRequestOpt
}

func (rc *AuthenAccessTokenReqCall) Do() (*UserAccessTokenInfo, error) {
	var result = &UserAccessTokenInfo{}
	req := request.NewRequest("/open-apis/authen/v1/access_token", "POST",
		[]request.AccessTokenType{request.AccessTokenTypeApp}, rc.body, result, rc.opts...)
	err := api.Send(rc.ctx, rc.authens.service.conf, req)
	return result, err
}

func (authens *AuthenService) AccessToken(ctx *lark.Context, body *AuthenAccessTokenReqBody, opts ...lark.APIRequestOpt) *AuthenAccessTokenReqCall {
	return &AuthenAccessTokenReqCall{
		ctx:     ctx,
		authens: authens,
		body:    body,
		opts:    opts,
	}
}

type AuthenRefreshAccessTokenReqCall struct {
	ctx     *lark.Context
	authens *AuthenService
	body    *AuthenRefreshAccessTokenReqBody
	opts    []lark.APIRequestOpt
}

func (rc *AuthenRefreshAccessTokenReqCall) Do() (*UserAccessTokenInfo, error) {
	var result = &UserAccessTokenInfo{}
	req := request.NewRequest("/open-apis/authen/v1/refresh_access_token", "POST",
		[]request.AccessTokenType{request.AccessTokenTypeApp}, rc.body, result, rc.opts...)
	err := api.Send(rc.ctx, rc.authens.service.conf, req)
	return result, err
}

func (authens *AuthenService) RefreshAccessToken(ctx *lark.Context, body *AuthenRefreshAccessTokenReqBody, opts ...lark.APIRequestOpt) *AuthenRefreshAccessTokenReqCall {
	return &AuthenRefreshAccessTokenReqCall{
		ctx:     ctx,
		authens: authens,
		body:    body,
		opts:    opts,
	}
}

type AuthenUserInfoReqCall struct {
	ctx     *lark.Context
	authens *AuthenService
	opts    []lark.APIRequestOpt
}

func (rc *AuthenUserInfoReqCall) Do() (*UserInfo, error) {
	var result = &UserInfo{}
	req := request.NewRequest("/open-apis/authen/v1/user_info", "GET",
		[]request.AccessTokenType{request.AccessTokenTypeUser}, nil, result, rc.opts...)
	err := api.Send(rc.ctx, rc.authens.service.conf, req)
	return result, err
}

func (authens *AuthenService) UserInfo(ctx *lark.Context, opts ...lark.APIRequestOpt) *AuthenUserInfoReqCall {
	return &AuthenUserInfoReqCall{
		ctx:     ctx,
		authens: authens,
		opts:    opts,
	}
}
