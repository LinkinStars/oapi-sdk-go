// Code generated by lark suite oapi sdk gen
package v1

import (
	"github.com/larksuite/oapi-sdk-go"
	"github.com/larksuite/oapi-sdk-go/event"
)

type RoomCreatedEventHandler struct {
	Fn func(*lark.Context, *RoomCreatedEvent) error
}

func (h *RoomCreatedEventHandler) GetEvent() interface{} {
	return &RoomCreatedEvent{}
}

func (h *RoomCreatedEventHandler) Handle(ctx *lark.Context, event interface{}) error {
	return h.Fn(ctx, event.(*RoomCreatedEvent))
}

func SetRoomCreatedEventHandler(conf lark.Config, fn func(ctx *lark.Context, event *RoomCreatedEvent) error) {
	event.SetTypeHandler(conf, "meeting_room.room.created_v1", &RoomCreatedEventHandler{Fn: fn})
}

type RoomDeletedEventHandler struct {
	Fn func(*lark.Context, *RoomDeletedEvent) error
}

func (h *RoomDeletedEventHandler) GetEvent() interface{} {
	return &RoomDeletedEvent{}
}

func (h *RoomDeletedEventHandler) Handle(ctx *lark.Context, event interface{}) error {
	return h.Fn(ctx, event.(*RoomDeletedEvent))
}

func SetRoomDeletedEventHandler(conf lark.Config, fn func(ctx *lark.Context, event *RoomDeletedEvent) error) {
	event.SetTypeHandler(conf, "meeting_room.room.deleted_v1", &RoomDeletedEventHandler{Fn: fn})
}

type RoomStatusChangedEventHandler struct {
	Fn func(*lark.Context, *RoomStatusChangedEvent) error
}

func (h *RoomStatusChangedEventHandler) GetEvent() interface{} {
	return &RoomStatusChangedEvent{}
}

func (h *RoomStatusChangedEventHandler) Handle(ctx *lark.Context, event interface{}) error {
	return h.Fn(ctx, event.(*RoomStatusChangedEvent))
}

func SetRoomStatusChangedEventHandler(conf lark.Config, fn func(ctx *lark.Context, event *RoomStatusChangedEvent) error) {
	event.SetTypeHandler(conf, "meeting_room.room.status_changed_v1", &RoomStatusChangedEventHandler{Fn: fn})
}

type RoomUpdatedEventHandler struct {
	Fn func(*lark.Context, *RoomUpdatedEvent) error
}

func (h *RoomUpdatedEventHandler) GetEvent() interface{} {
	return &RoomUpdatedEvent{}
}

func (h *RoomUpdatedEventHandler) Handle(ctx *lark.Context, event interface{}) error {
	return h.Fn(ctx, event.(*RoomUpdatedEvent))
}

func SetRoomUpdatedEventHandler(conf lark.Config, fn func(ctx *lark.Context, event *RoomUpdatedEvent) error) {
	event.SetTypeHandler(conf, "meeting_room.room.updated_v1", &RoomUpdatedEventHandler{Fn: fn})
}

type MeetingRoomStatusChangedEventHandler struct {
	Fn func(*lark.Context, *MeetingRoomStatusChangedEvent) error
}

func (h *MeetingRoomStatusChangedEventHandler) GetEvent() interface{} {
	return &MeetingRoomStatusChangedEvent{}
}

func (h *MeetingRoomStatusChangedEventHandler) Handle(ctx *lark.Context, event interface{}) error {
	return h.Fn(ctx, event.(*MeetingRoomStatusChangedEvent))
}

func SetMeetingRoomStatusChangedEventHandler(conf lark.Config, fn func(ctx *lark.Context, event *MeetingRoomStatusChangedEvent) error) {
	event.SetTypeHandler(conf, "meeting_room.meeting_room.status_changed_v1", &MeetingRoomStatusChangedEventHandler{Fn: fn})
}

type MeetingRoomCreatedEventHandler struct {
	Fn func(*lark.Context, *MeetingRoomCreatedEvent) error
}

func (h *MeetingRoomCreatedEventHandler) GetEvent() interface{} {
	return &MeetingRoomCreatedEvent{}
}

func (h *MeetingRoomCreatedEventHandler) Handle(ctx *lark.Context, event interface{}) error {
	return h.Fn(ctx, event.(*MeetingRoomCreatedEvent))
}

func SetMeetingRoomCreatedEventHandler(conf lark.Config, fn func(ctx *lark.Context, event *MeetingRoomCreatedEvent) error) {
	event.SetTypeHandler(conf, "meeting_room.meeting_room.created_v1", &MeetingRoomCreatedEventHandler{Fn: fn})
}

type MeetingRoomDeletedEventHandler struct {
	Fn func(*lark.Context, *MeetingRoomDeletedEvent) error
}

func (h *MeetingRoomDeletedEventHandler) GetEvent() interface{} {
	return &MeetingRoomDeletedEvent{}
}

func (h *MeetingRoomDeletedEventHandler) Handle(ctx *lark.Context, event interface{}) error {
	return h.Fn(ctx, event.(*MeetingRoomDeletedEvent))
}

func SetMeetingRoomDeletedEventHandler(conf lark.Config, fn func(ctx *lark.Context, event *MeetingRoomDeletedEvent) error) {
	event.SetTypeHandler(conf, "meeting_room.meeting_room.deleted_v1", &MeetingRoomDeletedEventHandler{Fn: fn})
}

type MeetingRoomUpdatedEventHandler struct {
	Fn func(*lark.Context, *MeetingRoomUpdatedEvent) error
}

func (h *MeetingRoomUpdatedEventHandler) GetEvent() interface{} {
	return &MeetingRoomUpdatedEvent{}
}

func (h *MeetingRoomUpdatedEventHandler) Handle(ctx *lark.Context, event interface{}) error {
	return h.Fn(ctx, event.(*MeetingRoomUpdatedEvent))
}

func SetMeetingRoomUpdatedEventHandler(conf lark.Config, fn func(ctx *lark.Context, event *MeetingRoomUpdatedEvent) error) {
	event.SetTypeHandler(conf, "meeting_room.meeting_room.updated_v1", &MeetingRoomUpdatedEventHandler{Fn: fn})
}
