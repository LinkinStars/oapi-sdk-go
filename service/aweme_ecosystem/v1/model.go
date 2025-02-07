// Code generated by lark suite oapi sdk gen
package v1

import (
	"github.com/larksuite/oapi-sdk-go/api/core/tools"
	"github.com/larksuite/oapi-sdk-go/event/core/model"
)

type UserId struct {
	UserId          string   `json:"user_id,omitempty"`
	OpenId          string   `json:"open_id,omitempty"`
	UnionId         string   `json:"union_id,omitempty"`
	ForceSendFields []string `json:"-"`
}

func (s *UserId) MarshalJSON() ([]byte, error) {
	type cp UserId
	raw := cp(*s)
	return tools.MarshalJSON(raw, s.ForceSendFields)
}

type AwemeUser struct {
	AwemeUserId     int64    `json:"aweme_user_id,omitempty,string"`
	UserId          string   `json:"user_id,omitempty"`
	IsBinded        bool     `json:"is_binded,omitempty"`
	ForceSendFields []string `json:"-"`
}

func (s *AwemeUser) MarshalJSON() ([]byte, error) {
	type cp AwemeUser
	raw := cp(*s)
	return tools.MarshalJSON(raw, s.ForceSendFields)
}

type EventAwemeUser struct {
	UserId          *UserId  `json:"user_id,omitempty"`
	AwemeUserId     string   `json:"aweme_user_id,omitempty"`
	IsBinded        bool     `json:"is_binded,omitempty"`
	ForceSendFields []string `json:"-"`
}

func (s *EventAwemeUser) MarshalJSON() ([]byte, error) {
	type cp EventAwemeUser
	raw := cp(*s)
	return tools.MarshalJSON(raw, s.ForceSendFields)
}

type AwemeUserGetBindInfoResult struct {
	AwemeUser *AwemeUser `json:"aweme_user,omitempty"`
}

type AwemeUserBindedAccountEventData struct {
	EventAwemeUser *EventAwemeUser `json:"event_aweme_user,omitempty"`
}

type AwemeUserBindedAccountEvent struct {
	*model.BaseEventV2
	Event *AwemeUserBindedAccountEventData `json:"event"`
}
