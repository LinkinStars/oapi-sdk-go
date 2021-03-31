// Code generated by lark suite oapi sdk gen
package v1

import (
	"github.com/larksuite/oapi-sdk-go/api/core/tools"
	"github.com/larksuite/oapi-sdk-go/event/core/model"
)

type Meeting struct {
	Id               int64                 `json:"id,omitempty,string"`
	Topic            string                `json:"topic,omitempty"`
	Url              string                `json:"url,omitempty"`
	CreateTime       int64                 `json:"create_time,omitempty,string"`
	StartTime        int64                 `json:"start_time,omitempty,string"`
	EndTime          int64                 `json:"end_time,omitempty,string"`
	HostUser         *MeetingUser          `json:"host_user,omitempty"`
	Status           int                   `json:"status,omitempty"`
	ParticipantCount int64                 `json:"participant_count,omitempty,string"`
	Participants     []*MeetingParticipant `json:"participants,omitempty"`
	Ability          *MeetingAbility       `json:"ability,omitempty"`
	ForceSendFields  []string              `json:"-"`
}

func (s *Meeting) MarshalJSON() ([]byte, error) {
	type cp Meeting
	raw := cp(*s)
	return tools.MarshalJSON(raw, s.ForceSendFields)
}

type MeetingRecording struct {
	Id              int64    `json:"id,omitempty,string"`
	MeetingId       int64    `json:"meeting_id,omitempty,string"`
	Url             string   `json:"url,omitempty"`
	Duration        int64    `json:"duration,omitempty,string"`
	ForceSendFields []string `json:"-"`
}

func (s *MeetingRecording) MarshalJSON() ([]byte, error) {
	type cp MeetingRecording
	raw := cp(*s)
	return tools.MarshalJSON(raw, s.ForceSendFields)
}

type MeetingAbility struct {
	UseVideo        bool     `json:"use_video,omitempty"`
	UseAudio        bool     `json:"use_audio,omitempty"`
	UseShareScreen  bool     `json:"use_share_screen,omitempty"`
	UseFollowScreen bool     `json:"use_follow_screen,omitempty"`
	UseRecording    bool     `json:"use_recording,omitempty"`
	UsePstn         bool     `json:"use_pstn,omitempty"`
	ForceSendFields []string `json:"-"`
}

func (s *MeetingAbility) MarshalJSON() ([]byte, error) {
	type cp MeetingAbility
	raw := cp(*s)
	return tools.MarshalJSON(raw, s.ForceSendFields)
}

type MeetingEventMeeting struct {
	Id              int64             `json:"id,omitempty,string"`
	Topic           string            `json:"topic,omitempty"`
	MeetingNo       string            `json:"meeting_no,omitempty"`
	StartTime       int64             `json:"start_time,omitempty,string"`
	EndTime         int64             `json:"end_time,omitempty,string"`
	HostUser        *MeetingEventUser `json:"host_user,omitempty"`
	Owner           *MeetingEventUser `json:"owner,omitempty"`
	ForceSendFields []string          `json:"-"`
}

func (s *MeetingEventMeeting) MarshalJSON() ([]byte, error) {
	type cp MeetingEventMeeting
	raw := cp(*s)
	return tools.MarshalJSON(raw, s.ForceSendFields)
}

type MeetingEventUser struct {
	Id              *UserId  `json:"id,omitempty"`
	UserRole        int      `json:"user_role,omitempty"`
	UserType        int      `json:"user_type,omitempty"`
	ForceSendFields []string `json:"-"`
}

func (s *MeetingEventUser) MarshalJSON() ([]byte, error) {
	type cp MeetingEventUser
	raw := cp(*s)
	return tools.MarshalJSON(raw, s.ForceSendFields)
}

type MeetingInviteStatus struct {
	Id              string   `json:"id,omitempty"`
	UserType        int      `json:"user_type,omitempty"`
	Status          int      `json:"status,omitempty"`
	ForceSendFields []string `json:"-"`
}

func (s *MeetingInviteStatus) MarshalJSON() ([]byte, error) {
	type cp MeetingInviteStatus
	raw := cp(*s)
	return tools.MarshalJSON(raw, s.ForceSendFields)
}

type MeetingParticipant struct {
	Id              string   `json:"id,omitempty"`
	UserType        int      `json:"user_type,omitempty"`
	IsHost          bool     `json:"is_host,omitempty"`
	IsCohost        bool     `json:"is_cohost,omitempty"`
	IsExternal      bool     `json:"is_external,omitempty"`
	Status          int      `json:"status,omitempty"`
	ForceSendFields []string `json:"-"`
}

func (s *MeetingParticipant) MarshalJSON() ([]byte, error) {
	type cp MeetingParticipant
	raw := cp(*s)
	return tools.MarshalJSON(raw, s.ForceSendFields)
}

type MeetingParticipantResult struct {
	Id              string   `json:"id,omitempty"`
	UserType        int      `json:"user_type,omitempty"`
	Result          int      `json:"result,omitempty"`
	ForceSendFields []string `json:"-"`
}

func (s *MeetingParticipantResult) MarshalJSON() ([]byte, error) {
	type cp MeetingParticipantResult
	raw := cp(*s)
	return tools.MarshalJSON(raw, s.ForceSendFields)
}

type MeetingUser struct {
	Id              string   `json:"id,omitempty"`
	UserType        int      `json:"user_type,omitempty"`
	ForceSendFields []string `json:"-"`
}

func (s *MeetingUser) MarshalJSON() ([]byte, error) {
	type cp MeetingUser
	raw := cp(*s)
	return tools.MarshalJSON(raw, s.ForceSendFields)
}

type RecordingPermissionObject struct {
	Id              string   `json:"id,omitempty"`
	Type            int      `json:"type,omitempty"`
	Permission      int      `json:"permission,omitempty"`
	ForceSendFields []string `json:"-"`
}

func (s *RecordingPermissionObject) MarshalJSON() ([]byte, error) {
	type cp RecordingPermissionObject
	raw := cp(*s)
	return tools.MarshalJSON(raw, s.ForceSendFields)
}

type Report struct {
	TotalMeetingCount     int64                 `json:"total_meeting_count,omitempty,string"`
	TotalMeetingDuration  int64                 `json:"total_meeting_duration,omitempty,string"`
	TotalParticipantCount int64                 `json:"total_participant_count,omitempty,string"`
	DailyReport           []*ReportMeetingDaily `json:"daily_report,omitempty"`
	ForceSendFields       []string              `json:"-"`
}

func (s *Report) MarshalJSON() ([]byte, error) {
	type cp Report
	raw := cp(*s)
	return tools.MarshalJSON(raw, s.ForceSendFields)
}

type ReportMeetingDaily struct {
	Date             int64    `json:"date,omitempty,string"`
	MeetingCount     int64    `json:"meeting_count,omitempty,string"`
	MeetingDuration  int64    `json:"meeting_duration,omitempty,string"`
	ParticipantCount int64    `json:"participant_count,omitempty,string"`
	ForceSendFields  []string `json:"-"`
}

func (s *ReportMeetingDaily) MarshalJSON() ([]byte, error) {
	type cp ReportMeetingDaily
	raw := cp(*s)
	return tools.MarshalJSON(raw, s.ForceSendFields)
}

type ReportTopUser struct {
	Id              int64    `json:"id,omitempty,string"`
	Name            string   `json:"name,omitempty"`
	UserType        int      `json:"user_type,omitempty"`
	MeetingCount    int64    `json:"meeting_count,omitempty,string"`
	MeetingDuration int64    `json:"meeting_duration,omitempty,string"`
	ForceSendFields []string `json:"-"`
}

func (s *ReportTopUser) MarshalJSON() ([]byte, error) {
	type cp ReportTopUser
	raw := cp(*s)
	return tools.MarshalJSON(raw, s.ForceSendFields)
}

type Reserve struct {
	Id              int64                  `json:"id,omitempty,string"`
	MeetingNo       string                 `json:"meeting_no,omitempty"`
	Url             string                 `json:"url,omitempty"`
	AppLink         string                 `json:"app_link,omitempty"`
	EndTime         int64                  `json:"end_time,omitempty,string"`
	ExpireStatus    int                    `json:"expire_status,omitempty"`
	ReserveUserId   string                 `json:"reserve_user_id,omitempty"`
	MeetingSettings *ReserveMeetingSetting `json:"meeting_settings,omitempty"`
	ForceSendFields []string               `json:"-"`
}

func (s *Reserve) MarshalJSON() ([]byte, error) {
	type cp Reserve
	raw := cp(*s)
	return tools.MarshalJSON(raw, s.ForceSendFields)
}

type ReserveActionPermission struct {
	Permission         int                         `json:"permission,omitempty"`
	PermissionCheckers []*ReservePermissionChecker `json:"permission_checkers,omitempty"`
	ForceSendFields    []string                    `json:"-"`
}

func (s *ReserveActionPermission) MarshalJSON() ([]byte, error) {
	type cp ReserveActionPermission
	raw := cp(*s)
	return tools.MarshalJSON(raw, s.ForceSendFields)
}

type ReserveMeetingSetting struct {
	Topic             string                     `json:"topic,omitempty"`
	ActionPermissions []*ReserveActionPermission `json:"action_permissions,omitempty"`
	ForceSendFields   []string                   `json:"-"`
}

func (s *ReserveMeetingSetting) MarshalJSON() ([]byte, error) {
	type cp ReserveMeetingSetting
	raw := cp(*s)
	return tools.MarshalJSON(raw, s.ForceSendFields)
}

type ReservePermissionChecker struct {
	CheckField      int      `json:"check_field,omitempty"`
	CheckMode       int      `json:"check_mode,omitempty"`
	CheckList       []string `json:"check_list,omitempty"`
	ForceSendFields []string `json:"-"`
}

func (s *ReservePermissionChecker) MarshalJSON() ([]byte, error) {
	type cp ReservePermissionChecker
	raw := cp(*s)
	return tools.MarshalJSON(raw, s.ForceSendFields)
}

type RoomConfig struct {
	RoomBackground    string              `json:"room_background,omitempty"`
	DisplayBackground string              `json:"display_background,omitempty"`
	DigitalSignage    *RoomDigitalSignage `json:"digital_signage,omitempty"`
	ForceSendFields   []string            `json:"-"`
}

func (s *RoomConfig) MarshalJSON() ([]byte, error) {
	type cp RoomConfig
	raw := cp(*s)
	return tools.MarshalJSON(raw, s.ForceSendFields)
}

type RoomDigitalSignage struct {
	Enable          bool                          `json:"enable,omitempty"`
	Mute            bool                          `json:"mute,omitempty"`
	StartDisplay    int                           `json:"start_display,omitempty"`
	StopDisplay     int                           `json:"stop_display,omitempty"`
	Materials       []*RoomDigitalSignageMaterial `json:"materials,omitempty"`
	ForceSendFields []string                      `json:"-"`
}

func (s *RoomDigitalSignage) MarshalJSON() ([]byte, error) {
	type cp RoomDigitalSignage
	raw := cp(*s)
	return tools.MarshalJSON(raw, s.ForceSendFields)
}

type RoomDigitalSignageMaterial struct {
	Id              string   `json:"id,omitempty"`
	Name            string   `json:"name,omitempty"`
	MaterialType    int      `json:"material_type,omitempty"`
	Url             string   `json:"url,omitempty"`
	Duration        int      `json:"duration,omitempty"`
	Cover           string   `json:"cover,omitempty"`
	Md5             string   `json:"md5,omitempty"`
	ForceSendFields []string `json:"-"`
}

func (s *RoomDigitalSignageMaterial) MarshalJSON() ([]byte, error) {
	type cp RoomDigitalSignageMaterial
	raw := cp(*s)
	return tools.MarshalJSON(raw, s.ForceSendFields)
}

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

type MeetingInviteReqBody struct {
	Invitees        []*MeetingUser `json:"invitees,omitempty"`
	ForceSendFields []string       `json:"-"`
}

func (s *MeetingInviteReqBody) MarshalJSON() ([]byte, error) {
	type cp MeetingInviteReqBody
	raw := cp(*s)
	return tools.MarshalJSON(raw, s.ForceSendFields)
}

type MeetingInviteResult struct {
	InviteResults []*MeetingInviteStatus `json:"invite_results,omitempty"`
}

type MeetingListResult struct {
	HasMore   bool       `json:"has_more,omitempty"`
	PageToken string     `json:"page_token,omitempty"`
	Meetings  []*Meeting `json:"meetings,omitempty"`
}

type ReportGetTopUserResult struct {
	TopUserReport []*ReportTopUser `json:"top_user_report,omitempty"`
}

type MeetingSetHostReqBody struct {
	HostUser        *MeetingUser `json:"host_user,omitempty"`
	OldHostUser     *MeetingUser `json:"old_host_user,omitempty"`
	ForceSendFields []string     `json:"-"`
}

func (s *MeetingSetHostReqBody) MarshalJSON() ([]byte, error) {
	type cp MeetingSetHostReqBody
	raw := cp(*s)
	return tools.MarshalJSON(raw, s.ForceSendFields)
}

type MeetingSetHostResult struct {
	HostUser *MeetingUser `json:"host_user,omitempty"`
}

type MeetingRecordingGetResult struct {
	Recording *MeetingRecording `json:"recording,omitempty"`
}

type ReportGetDailyResult struct {
	MeetingReport *Report `json:"meeting_report,omitempty"`
}

type MeetingGetResult struct {
	Meeting *Meeting `json:"meeting,omitempty"`
}

type RoomConfigSetReqBody struct {
	Scope           int         `json:"scope,omitempty"`
	CountryId       int64       `json:"country_id,omitempty,string"`
	DistrictId      int64       `json:"district_id,omitempty,string"`
	BuildingId      int64       `json:"building_id,omitempty,string"`
	FloorName       string      `json:"floor_name,omitempty"`
	RoomId          int64       `json:"room_id,omitempty,string"`
	RoomConfig      *RoomConfig `json:"room_config,omitempty"`
	ForceSendFields []string    `json:"-"`
}

func (s *RoomConfigSetReqBody) MarshalJSON() ([]byte, error) {
	type cp RoomConfigSetReqBody
	raw := cp(*s)
	return tools.MarshalJSON(raw, s.ForceSendFields)
}

type MeetingRecordingSetPermissionReqBody struct {
	PermissionObjects []*RecordingPermissionObject `json:"permission_objects,omitempty"`
	ForceSendFields   []string                     `json:"-"`
}

func (s *MeetingRecordingSetPermissionReqBody) MarshalJSON() ([]byte, error) {
	type cp MeetingRecordingSetPermissionReqBody
	raw := cp(*s)
	return tools.MarshalJSON(raw, s.ForceSendFields)
}

type MeetingRecordingStartReqBody struct {
	Timezone        int      `json:"timezone,omitempty"`
	ForceSendFields []string `json:"-"`
}

func (s *MeetingRecordingStartReqBody) MarshalJSON() ([]byte, error) {
	type cp MeetingRecordingStartReqBody
	raw := cp(*s)
	return tools.MarshalJSON(raw, s.ForceSendFields)
}

type ReserveUpdateReqBody struct {
	EndTime         int64                  `json:"end_time,omitempty,string"`
	MeetingSettings *ReserveMeetingSetting `json:"meeting_settings,omitempty"`
	ForceSendFields []string               `json:"-"`
}

func (s *ReserveUpdateReqBody) MarshalJSON() ([]byte, error) {
	type cp ReserveUpdateReqBody
	raw := cp(*s)
	return tools.MarshalJSON(raw, s.ForceSendFields)
}

type ReserveUpdateResult struct {
	Reserve *Reserve `json:"reserve,omitempty"`
}

type ReserveApplyReqBody struct {
	EndTime         int64                  `json:"end_time,omitempty,string"`
	MeetingSettings *ReserveMeetingSetting `json:"meeting_settings,omitempty"`
	ForceSendFields []string               `json:"-"`
}

func (s *ReserveApplyReqBody) MarshalJSON() ([]byte, error) {
	type cp ReserveApplyReqBody
	raw := cp(*s)
	return tools.MarshalJSON(raw, s.ForceSendFields)
}

type ReserveApplyResult struct {
	Reserve *Reserve `json:"reserve,omitempty"`
}

type ReserveGetResult struct {
	Reserve *Reserve `json:"reserve,omitempty"`
}

type ReserveGetActiveMeetingResult struct {
	Meeting *Meeting `json:"meeting,omitempty"`
}

type MeetingLeaveMeetingEventData struct {
	Meeting     *MeetingEventMeeting `json:"meeting,omitempty"`
	Operator    *MeetingEventUser    `json:"operator,omitempty"`
	LeaveReason int                  `json:"leave_reason,omitempty"`
}

type MeetingLeaveMeetingEvent struct {
	*model.BaseEventV2
	Event *MeetingLeaveMeetingEventData `json:"event"`
}

type MeetingMeetingEndedEventData struct {
	Meeting  *MeetingEventMeeting `json:"meeting,omitempty"`
	Operator *MeetingEventUser    `json:"operator,omitempty"`
}

type MeetingMeetingEndedEvent struct {
	*model.BaseEventV2
	Event *MeetingMeetingEndedEventData `json:"event"`
}

type MeetingMeetingStartedEventData struct {
	Meeting  *MeetingEventMeeting `json:"meeting,omitempty"`
	Operator *MeetingEventUser    `json:"operator,omitempty"`
}

type MeetingMeetingStartedEvent struct {
	*model.BaseEventV2
	Event *MeetingMeetingStartedEventData `json:"event"`
}

type MeetingRecordingEndedEventData struct {
	Meeting  *MeetingEventMeeting `json:"meeting,omitempty"`
	Operator *MeetingEventUser    `json:"operator,omitempty"`
}

type MeetingRecordingEndedEvent struct {
	*model.BaseEventV2
	Event *MeetingRecordingEndedEventData `json:"event"`
}

type MeetingShareEndedEventData struct {
	Meeting  *MeetingEventMeeting `json:"meeting,omitempty"`
	Operator *MeetingEventUser    `json:"operator,omitempty"`
}

type MeetingShareEndedEvent struct {
	*model.BaseEventV2
	Event *MeetingShareEndedEventData `json:"event"`
}

type MeetingJoinMeetingEventData struct {
	Meeting  *MeetingEventMeeting `json:"meeting,omitempty"`
	Operator *MeetingEventUser    `json:"operator,omitempty"`
}

type MeetingJoinMeetingEvent struct {
	*model.BaseEventV2
	Event *MeetingJoinMeetingEventData `json:"event"`
}

type MeetingRecordingStartedEventData struct {
	Meeting  *MeetingEventMeeting `json:"meeting,omitempty"`
	Operator *MeetingEventUser    `json:"operator,omitempty"`
}

type MeetingRecordingStartedEvent struct {
	*model.BaseEventV2
	Event *MeetingRecordingStartedEventData `json:"event"`
}

type MeetingSendMeetingImEventData struct {
	Meeting  *MeetingEventMeeting `json:"meeting,omitempty"`
	Operator *MeetingEventUser    `json:"operator,omitempty"`
	Content  string               `json:"content,omitempty"`
}

type MeetingSendMeetingImEvent struct {
	*model.BaseEventV2
	Event *MeetingSendMeetingImEventData `json:"event"`
}

type MeetingShareStartedEventData struct {
	Meeting  *MeetingEventMeeting `json:"meeting,omitempty"`
	Operator *MeetingEventUser    `json:"operator,omitempty"`
}

type MeetingShareStartedEvent struct {
	*model.BaseEventV2
	Event *MeetingShareStartedEventData `json:"event"`
}

type MeetingRecordingReadyEventData struct {
	Meeting  *MeetingEventMeeting `json:"meeting,omitempty"`
	Url      string               `json:"url,omitempty"`
	Duration int64                `json:"duration,omitempty,string"`
}

type MeetingRecordingReadyEvent struct {
	*model.BaseEventV2
	Event *MeetingRecordingReadyEventData `json:"event"`
}
