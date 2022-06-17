// Package wiki code generated by oapi sdk gen
package wiki

import (
	"context"
	"errors"
	"fmt"

	"github.com/larksuite/oapi-sdk-go/core"
)

/**生成枚举值 **/

const (
	ObjTypeObjTypeDoc      string = "doc"
	ObjTypeObjTypeSheet    string = "sheet"
	ObjTypeObjTypeMindNote string = "mindnote"
	ObjTypeObjTypeBitable  string = "bitable"
	ObjTypeObjTypeFile     string = "file"
	ObjTypeObjTypeDocx     string = "docx"
)

const (
	NodeTypeNodeTypeEntity   string = "origin"
	NodeTypeNodeTypeShortCut string = "shortcut"
)

const (
	MoveDocsToWikiObjTypeObjTypeDoc      string = "doc"
	MoveDocsToWikiObjTypeObjTypeSheet    string = "sheet"
	MoveDocsToWikiObjTypeObjTypeBitable  string = "bitable"
	MoveDocsToWikiObjTypeObjTypeMindNote string = "mindnote"
	MoveDocsToWikiObjTypeObjTypeDocx     string = "docx"
	MoveDocsToWikiObjTypeObjTypeFile     string = "file"
)

const (
	TaskTypeMove string = "move"
)

/**生成数据类型 **/

type Member struct {
	MemberType *string `json:"member_type,omitempty"`
	MemberId   *string `json:"member_id,omitempty"`
	MemberRole *string `json:"member_role,omitempty"`
}

type MoveResult struct {
	Node      *Node   `json:"node,omitempty"`
	Status    *int    `json:"status,omitempty"`
	StatusMsg *string `json:"status_msg,omitempty"`
}

type Node struct {
	SpaceId         *int64  `json:"space_id,omitempty,string"`
	NodeToken       *string `json:"node_token,omitempty"`
	ObjToken        *string `json:"obj_token,omitempty"`
	ObjType         *string `json:"obj_type,omitempty"`
	ParentNodeToken *string `json:"parent_node_token,omitempty"`
	NodeType        *string `json:"node_type,omitempty"`
	OriginNodeToken *string `json:"origin_node_token,omitempty"`
	OriginSpaceId   *int64  `json:"origin_space_id,omitempty,string"`
	HasChild        *bool   `json:"has_child,omitempty"`
	Title           *string `json:"title,omitempty"`
	ObjCreateTime   *int64  `json:"obj_create_time,omitempty,string"`
	ObjEditTime     *int64  `json:"obj_edit_time,omitempty,string"`
	NodeCreateTime  *int64  `json:"node_create_time,omitempty,string"`
	Creator         *string `json:"creator,omitempty"`
	Owner           *string `json:"owner,omitempty"`
}

type NodeSearch struct {
	NodeId   *string  `json:"node_id,omitempty"`
	SpaceId  *string  `json:"space_id,omitempty"`
	ParentId *string  `json:"parent_id,omitempty"`
	ObjType  *int     `json:"obj_type,omitempty"`
	Title    *string  `json:"title,omitempty"`
	Url      *string  `json:"url,omitempty"`
	Icon     *string  `json:"icon,omitempty"`
	AreaId   *string  `json:"area_id,omitempty"`
	SortId   *float64 `json:"sort_id,omitempty"`
	Domain   *string  `json:"domain,omitempty"`
	ObjToken *string  `json:"obj_token,omitempty"`
}

type Setting struct {
	CreateSetting   *string `json:"create_setting,omitempty"`
	SecuritySetting *string `json:"security_setting,omitempty"`
	CommentSetting  *string `json:"comment_setting,omitempty"`
}

type Space struct {
	Name        *string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
	SpaceId     *string `json:"space_id,omitempty"`
	SpaceType   *string `json:"space_type,omitempty"`
	Visibility  *string `json:"visibility,omitempty"`
}

type TaskResult struct {
	TaskId     *string       `json:"task_id,omitempty"`
	MoveResult []*MoveResult `json:"move_result,omitempty"`
}

/**生成请求和响应结果类型，以及请求对象的Builder构造器 **/

/*1.4 生成请求的builder结构体*/
type CreateSpaceReqBuilder struct {
	space     *Space
	spaceFlag bool
}

// 生成请求的New构造器
func NewCreateSpaceReqBuilder() *CreateSpaceReqBuilder {
	builder := &CreateSpaceReqBuilder{}
	return builder
}

/*1.5 生成请求的builder属性方法*/
func (builder *CreateSpaceReqBuilder) Space(space *Space) *CreateSpaceReqBuilder {
	builder.space = space
	builder.spaceFlag = true
	return builder
}

/*1.5 生成请求的builder的build方法*/
func (builder *CreateSpaceReqBuilder) Build() *CreateSpaceReq {
	req := &CreateSpaceReq{}
	if builder.spaceFlag {
		req.Space = builder.space
	}
	return req
}

type CreateSpaceReq struct {
	Space *Space `body:""`
}

type CreateSpaceRespData struct {
	Space *Space `json:"space,omitempty"`
}

type CreateSpaceResp struct {
	*core.RawResponse `json:"-"`
	core.CodeError
	Data *CreateSpaceRespData `json:"data"`
}

func (resp *CreateSpaceResp) Success() bool {
	return resp.Code == 0
}

/*1.4 生成请求的builder结构体*/
type GetSpaceReqBuilder struct {
	spaceId     string
	spaceIdFlag bool
}

// 生成请求的New构造器
func NewGetSpaceReqBuilder() *GetSpaceReqBuilder {
	builder := &GetSpaceReqBuilder{}
	return builder
}

/*1.5 生成请求的builder属性方法*/
func (builder *GetSpaceReqBuilder) SpaceId(spaceId string) *GetSpaceReqBuilder {
	builder.spaceId = spaceId
	builder.spaceIdFlag = true
	return builder
}

/*1.5 生成请求的builder的build方法*/
func (builder *GetSpaceReqBuilder) Build() *GetSpaceReq {
	req := &GetSpaceReq{}
	if builder.spaceIdFlag {
		req.SpaceId = builder.spaceId
	}
	return req
}

type GetSpaceReq struct {
	SpaceId string `path:"space_id"`
}

type GetSpaceRespData struct {
	Space *Space `json:"space,omitempty"`
}

type GetSpaceResp struct {
	*core.RawResponse `json:"-"`
	core.CodeError
	Data *GetSpaceRespData `json:"data"`
}

func (resp *GetSpaceResp) Success() bool {
	return resp.Code == 0
}

/*1.4 生成请求的builder结构体*/
type GetNodeSpaceReqBuilder struct {
	token     string
	tokenFlag bool
}

// 生成请求的New构造器
func NewGetNodeSpaceReqBuilder() *GetNodeSpaceReqBuilder {
	builder := &GetNodeSpaceReqBuilder{}
	return builder
}

/*1.5 生成请求的builder属性方法*/
func (builder *GetNodeSpaceReqBuilder) Token(token string) *GetNodeSpaceReqBuilder {
	builder.token = token
	builder.tokenFlag = true
	return builder
}

/*1.5 生成请求的builder的build方法*/
func (builder *GetNodeSpaceReqBuilder) Build() *GetNodeSpaceReq {
	req := &GetNodeSpaceReq{}
	if builder.tokenFlag {
		req.Token = &builder.token
	}
	return req
}

type GetNodeSpaceReq struct {
	Token *string `query:"token"`
}

type GetNodeSpaceRespData struct {
	Node *Node `json:"node,omitempty"`
}

type GetNodeSpaceResp struct {
	*core.RawResponse `json:"-"`
	core.CodeError
	Data *GetNodeSpaceRespData `json:"data"`
}

func (resp *GetNodeSpaceResp) Success() bool {
	return resp.Code == 0
}

/*1.4 生成请求的builder结构体*/
type ListSpaceReqBuilder struct {
	pageSize      int
	pageSizeFlag  bool
	pageToken     string
	pageTokenFlag bool
	limit         int
}

// 生成请求的New构造器
func NewListSpaceReqBuilder() *ListSpaceReqBuilder {
	builder := &ListSpaceReqBuilder{}
	return builder
}

/*1.5 生成请求的builder属性方法*/
func (builder *ListSpaceReqBuilder) Limit(limit int) *ListSpaceReqBuilder {
	builder.limit = limit
	return builder
}
func (builder *ListSpaceReqBuilder) PageSize(pageSize int) *ListSpaceReqBuilder {
	builder.pageSize = pageSize
	builder.pageSizeFlag = true
	return builder
}
func (builder *ListSpaceReqBuilder) PageToken(pageToken string) *ListSpaceReqBuilder {
	builder.pageToken = pageToken
	builder.pageTokenFlag = true
	return builder
}

/*1.5 生成请求的builder的build方法*/
func (builder *ListSpaceReqBuilder) Build() *ListSpaceReq {
	req := &ListSpaceReq{}
	req.Limit = builder.limit
	if builder.pageSizeFlag {
		req.PageSize = &builder.pageSize
	}
	if builder.pageTokenFlag {
		req.PageToken = &builder.pageToken
	}
	return req
}

type ListSpaceReq struct {
	PageSize  *int    `query:"page_size"`
	PageToken *string `query:"page_token"`
	Limit     int
}

type ListSpaceRespData struct {
	Items     []*Space `json:"items,omitempty"`
	PageToken *string  `json:"page_token,omitempty"`
	HasMore   *bool    `json:"has_more,omitempty"`
}

type ListSpaceResp struct {
	*core.RawResponse `json:"-"`
	core.CodeError
	Data *ListSpaceRespData `json:"data"`
}

func (resp *ListSpaceResp) Success() bool {
	return resp.Code == 0
}

/*1.4 生成请求的builder结构体*/
type CreateSpaceMemberReqBuilder struct {
	spaceId              string
	spaceIdFlag          bool
	needNotification     bool
	needNotificationFlag bool
	member               *Member
	memberFlag           bool
}

// 生成请求的New构造器
func NewCreateSpaceMemberReqBuilder() *CreateSpaceMemberReqBuilder {
	builder := &CreateSpaceMemberReqBuilder{}
	return builder
}

/*1.5 生成请求的builder属性方法*/
func (builder *CreateSpaceMemberReqBuilder) SpaceId(spaceId string) *CreateSpaceMemberReqBuilder {
	builder.spaceId = spaceId
	builder.spaceIdFlag = true
	return builder
}
func (builder *CreateSpaceMemberReqBuilder) NeedNotification(needNotification bool) *CreateSpaceMemberReqBuilder {
	builder.needNotification = needNotification
	builder.needNotificationFlag = true
	return builder
}
func (builder *CreateSpaceMemberReqBuilder) Member(member *Member) *CreateSpaceMemberReqBuilder {
	builder.member = member
	builder.memberFlag = true
	return builder
}

/*1.5 生成请求的builder的build方法*/
func (builder *CreateSpaceMemberReqBuilder) Build() *CreateSpaceMemberReq {
	req := &CreateSpaceMemberReq{}
	if builder.spaceIdFlag {
		req.SpaceId = builder.spaceId
	}
	if builder.needNotificationFlag {
		req.NeedNotification = &builder.needNotification
	}
	if builder.memberFlag {
		req.Member = builder.member
	}
	return req
}

type CreateSpaceMemberReq struct {
	SpaceId          string  `path:"space_id"`
	NeedNotification *bool   `query:"need_notification"`
	Member           *Member `body:""`
}

type CreateSpaceMemberRespData struct {
	Member *Member `json:"member,omitempty"`
}

type CreateSpaceMemberResp struct {
	*core.RawResponse `json:"-"`
	core.CodeError
	Data *CreateSpaceMemberRespData `json:"data"`
}

func (resp *CreateSpaceMemberResp) Success() bool {
	return resp.Code == 0
}

/*1.4 生成请求的builder结构体*/
type DeleteSpaceMemberReqBuilder struct {
	spaceId      string
	spaceIdFlag  bool
	memberId     string
	memberIdFlag bool
	member       *Member
	memberFlag   bool
}

// 生成请求的New构造器
func NewDeleteSpaceMemberReqBuilder() *DeleteSpaceMemberReqBuilder {
	builder := &DeleteSpaceMemberReqBuilder{}
	return builder
}

/*1.5 生成请求的builder属性方法*/
func (builder *DeleteSpaceMemberReqBuilder) SpaceId(spaceId string) *DeleteSpaceMemberReqBuilder {
	builder.spaceId = spaceId
	builder.spaceIdFlag = true
	return builder
}
func (builder *DeleteSpaceMemberReqBuilder) MemberId(memberId string) *DeleteSpaceMemberReqBuilder {
	builder.memberId = memberId
	builder.memberIdFlag = true
	return builder
}
func (builder *DeleteSpaceMemberReqBuilder) Member(member *Member) *DeleteSpaceMemberReqBuilder {
	builder.member = member
	builder.memberFlag = true
	return builder
}

/*1.5 生成请求的builder的build方法*/
func (builder *DeleteSpaceMemberReqBuilder) Build() *DeleteSpaceMemberReq {
	req := &DeleteSpaceMemberReq{}
	if builder.spaceIdFlag {
		req.SpaceId = builder.spaceId
	}
	if builder.memberIdFlag {
		req.MemberId = builder.memberId
	}
	if builder.memberFlag {
		req.Member = builder.member
	}
	return req
}

type DeleteSpaceMemberReq struct {
	SpaceId  string  `path:"space_id"`
	MemberId string  `path:"member_id"`
	Member   *Member `body:""`
}

type DeleteSpaceMemberRespData struct {
	Member *Member `json:"member,omitempty"`
}

type DeleteSpaceMemberResp struct {
	*core.RawResponse `json:"-"`
	core.CodeError
	Data *DeleteSpaceMemberRespData `json:"data"`
}

func (resp *DeleteSpaceMemberResp) Success() bool {
	return resp.Code == 0
}

type CopySpaceNodeReqBodyBuilder struct {
	targetParentToken     string
	targetParentTokenFlag bool
	targetSpaceId         int64
	targetSpaceIdFlag     bool
	title                 string
	titleFlag             bool
}

// 生成body的New构造器
func NewCopySpaceNodeReqBodyBuilder() *CopySpaceNodeReqBodyBuilder {
	builder := &CopySpaceNodeReqBodyBuilder{}
	return builder
}

/*1.2 生成body的builder属性方法*/
func (builder *CopySpaceNodeReqBodyBuilder) TargetParentToken(targetParentToken string) *CopySpaceNodeReqBodyBuilder {
	builder.targetParentToken = targetParentToken
	builder.targetParentTokenFlag = true
	return builder
}
func (builder *CopySpaceNodeReqBodyBuilder) TargetSpaceId(targetSpaceId int64) *CopySpaceNodeReqBodyBuilder {
	builder.targetSpaceId = targetSpaceId
	builder.targetSpaceIdFlag = true
	return builder
}
func (builder *CopySpaceNodeReqBodyBuilder) Title(title string) *CopySpaceNodeReqBodyBuilder {
	builder.title = title
	builder.titleFlag = true
	return builder
}

/*1.3 生成body的build方法*/
func (builder *CopySpaceNodeReqBodyBuilder) Build() *CopySpaceNodeReqBody {
	req := &CopySpaceNodeReqBody{}
	if builder.targetParentTokenFlag {
		req.TargetParentToken = &builder.targetParentToken

	}
	if builder.targetSpaceIdFlag {
		req.TargetSpaceId = &builder.targetSpaceId

	}
	if builder.titleFlag {
		req.Title = &builder.title

	}
	return req
}

/**上传文件path开始**/
type CopySpaceNodePathReqBodyBuilder struct {
	targetParentToken     string
	targetParentTokenFlag bool
	targetSpaceId         int64
	targetSpaceIdFlag     bool
	title                 string
	titleFlag             bool
}

func NewCopySpaceNodePathReqBodyBuilder() *CopySpaceNodePathReqBodyBuilder {
	builder := &CopySpaceNodePathReqBodyBuilder{}
	return builder
}
func (builder *CopySpaceNodePathReqBodyBuilder) TargetParentToken(targetParentToken string) *CopySpaceNodePathReqBodyBuilder {
	builder.targetParentToken = targetParentToken
	builder.targetParentTokenFlag = true
	return builder
}
func (builder *CopySpaceNodePathReqBodyBuilder) TargetSpaceId(targetSpaceId int64) *CopySpaceNodePathReqBodyBuilder {
	builder.targetSpaceId = targetSpaceId
	builder.targetSpaceIdFlag = true
	return builder
}
func (builder *CopySpaceNodePathReqBodyBuilder) Title(title string) *CopySpaceNodePathReqBodyBuilder {
	builder.title = title
	builder.titleFlag = true
	return builder
}

func (builder *CopySpaceNodePathReqBodyBuilder) Build() (*CopySpaceNodeReqBody, error) {
	req := &CopySpaceNodeReqBody{}
	if builder.targetParentTokenFlag {
		req.TargetParentToken = &builder.targetParentToken

	}
	if builder.targetSpaceIdFlag {
		req.TargetSpaceId = &builder.targetSpaceId

	}
	if builder.titleFlag {
		req.Title = &builder.title

	}
	return req, nil
}

/**上传文件path结束**/

/*1.4 生成请求的builder结构体*/
type CopySpaceNodeReqBuilder struct {
	spaceId       int64
	spaceIdFlag   bool
	nodeToken     string
	nodeTokenFlag bool
	body          *CopySpaceNodeReqBody
	bodyFlag      bool
}

// 生成请求的New构造器
func NewCopySpaceNodeReqBuilder() *CopySpaceNodeReqBuilder {
	builder := &CopySpaceNodeReqBuilder{}
	return builder
}

/*1.5 生成请求的builder属性方法*/
func (builder *CopySpaceNodeReqBuilder) SpaceId(spaceId int64) *CopySpaceNodeReqBuilder {
	builder.spaceId = spaceId
	builder.spaceIdFlag = true
	return builder
}
func (builder *CopySpaceNodeReqBuilder) NodeToken(nodeToken string) *CopySpaceNodeReqBuilder {
	builder.nodeToken = nodeToken
	builder.nodeTokenFlag = true
	return builder
}
func (builder *CopySpaceNodeReqBuilder) Body(body *CopySpaceNodeReqBody) *CopySpaceNodeReqBuilder {
	builder.body = body
	builder.bodyFlag = true
	return builder
}

/*1.5 生成请求的builder的build方法*/
func (builder *CopySpaceNodeReqBuilder) Build() *CopySpaceNodeReq {
	req := &CopySpaceNodeReq{}
	if builder.spaceIdFlag {
		req.SpaceId = builder.spaceId
	}
	if builder.nodeTokenFlag {
		req.NodeToken = builder.nodeToken
	}
	if builder.bodyFlag {
		req.Body = builder.body
	}
	return req
}

type CopySpaceNodeReqBody struct {
	TargetParentToken *string `json:"target_parent_token,omitempty"`
	TargetSpaceId     *int64  `json:"target_space_id,omitempty,string"`
	Title             *string `json:"title,omitempty"`
}

type CopySpaceNodeReq struct {
	SpaceId   int64                 `path:"space_id"`
	NodeToken string                `path:"node_token"`
	Body      *CopySpaceNodeReqBody `body:""`
}

type CopySpaceNodeRespData struct {
	Node *Node `json:"node,omitempty"`
}

type CopySpaceNodeResp struct {
	*core.RawResponse `json:"-"`
	core.CodeError
	Data *CopySpaceNodeRespData `json:"data"`
}

func (resp *CopySpaceNodeResp) Success() bool {
	return resp.Code == 0
}

/*1.4 生成请求的builder结构体*/
type CreateSpaceNodeReqBuilder struct {
	spaceId     string
	spaceIdFlag bool
	node        *Node
	nodeFlag    bool
}

// 生成请求的New构造器
func NewCreateSpaceNodeReqBuilder() *CreateSpaceNodeReqBuilder {
	builder := &CreateSpaceNodeReqBuilder{}
	return builder
}

/*1.5 生成请求的builder属性方法*/
func (builder *CreateSpaceNodeReqBuilder) SpaceId(spaceId string) *CreateSpaceNodeReqBuilder {
	builder.spaceId = spaceId
	builder.spaceIdFlag = true
	return builder
}
func (builder *CreateSpaceNodeReqBuilder) Node(node *Node) *CreateSpaceNodeReqBuilder {
	builder.node = node
	builder.nodeFlag = true
	return builder
}

/*1.5 生成请求的builder的build方法*/
func (builder *CreateSpaceNodeReqBuilder) Build() *CreateSpaceNodeReq {
	req := &CreateSpaceNodeReq{}
	if builder.spaceIdFlag {
		req.SpaceId = builder.spaceId
	}
	if builder.nodeFlag {
		req.Node = builder.node
	}
	return req
}

type CreateSpaceNodeReq struct {
	SpaceId string `path:"space_id"`
	Node    *Node  `body:""`
}

type CreateSpaceNodeRespData struct {
	Node *Node `json:"node,omitempty"`
}

type CreateSpaceNodeResp struct {
	*core.RawResponse `json:"-"`
	core.CodeError
	Data *CreateSpaceNodeRespData `json:"data"`
}

func (resp *CreateSpaceNodeResp) Success() bool {
	return resp.Code == 0
}

/*1.4 生成请求的builder结构体*/
type ListSpaceNodeReqBuilder struct {
	spaceId             string
	spaceIdFlag         bool
	pageSize            int
	pageSizeFlag        bool
	pageToken           string
	pageTokenFlag       bool
	parentNodeToken     string
	parentNodeTokenFlag bool
	limit               int
}

// 生成请求的New构造器
func NewListSpaceNodeReqBuilder() *ListSpaceNodeReqBuilder {
	builder := &ListSpaceNodeReqBuilder{}
	return builder
}

/*1.5 生成请求的builder属性方法*/
func (builder *ListSpaceNodeReqBuilder) Limit(limit int) *ListSpaceNodeReqBuilder {
	builder.limit = limit
	return builder
}
func (builder *ListSpaceNodeReqBuilder) SpaceId(spaceId string) *ListSpaceNodeReqBuilder {
	builder.spaceId = spaceId
	builder.spaceIdFlag = true
	return builder
}
func (builder *ListSpaceNodeReqBuilder) PageSize(pageSize int) *ListSpaceNodeReqBuilder {
	builder.pageSize = pageSize
	builder.pageSizeFlag = true
	return builder
}
func (builder *ListSpaceNodeReqBuilder) PageToken(pageToken string) *ListSpaceNodeReqBuilder {
	builder.pageToken = pageToken
	builder.pageTokenFlag = true
	return builder
}
func (builder *ListSpaceNodeReqBuilder) ParentNodeToken(parentNodeToken string) *ListSpaceNodeReqBuilder {
	builder.parentNodeToken = parentNodeToken
	builder.parentNodeTokenFlag = true
	return builder
}

/*1.5 生成请求的builder的build方法*/
func (builder *ListSpaceNodeReqBuilder) Build() *ListSpaceNodeReq {
	req := &ListSpaceNodeReq{}
	req.Limit = builder.limit
	if builder.spaceIdFlag {
		req.SpaceId = builder.spaceId
	}
	if builder.pageSizeFlag {
		req.PageSize = &builder.pageSize
	}
	if builder.pageTokenFlag {
		req.PageToken = &builder.pageToken
	}
	if builder.parentNodeTokenFlag {
		req.ParentNodeToken = &builder.parentNodeToken
	}
	return req
}

type ListSpaceNodeReq struct {
	SpaceId         string  `path:"space_id"`
	PageSize        *int    `query:"page_size"`
	PageToken       *string `query:"page_token"`
	ParentNodeToken *string `query:"parent_node_token"`
	Limit           int
}

type ListSpaceNodeRespData struct {
	Items     []*Node `json:"items,omitempty"`
	PageToken *string `json:"page_token,omitempty"`
	HasMore   *bool   `json:"has_more,omitempty"`
}

type ListSpaceNodeResp struct {
	*core.RawResponse `json:"-"`
	core.CodeError
	Data *ListSpaceNodeRespData `json:"data"`
}

func (resp *ListSpaceNodeResp) Success() bool {
	return resp.Code == 0
}

type MoveSpaceNodeReqBodyBuilder struct {
	targetParentToken     string
	targetParentTokenFlag bool
	targetSpaceId         string
	targetSpaceIdFlag     bool
}

// 生成body的New构造器
func NewMoveSpaceNodeReqBodyBuilder() *MoveSpaceNodeReqBodyBuilder {
	builder := &MoveSpaceNodeReqBodyBuilder{}
	return builder
}

/*1.2 生成body的builder属性方法*/
func (builder *MoveSpaceNodeReqBodyBuilder) TargetParentToken(targetParentToken string) *MoveSpaceNodeReqBodyBuilder {
	builder.targetParentToken = targetParentToken
	builder.targetParentTokenFlag = true
	return builder
}
func (builder *MoveSpaceNodeReqBodyBuilder) TargetSpaceId(targetSpaceId string) *MoveSpaceNodeReqBodyBuilder {
	builder.targetSpaceId = targetSpaceId
	builder.targetSpaceIdFlag = true
	return builder
}

/*1.3 生成body的build方法*/
func (builder *MoveSpaceNodeReqBodyBuilder) Build() *MoveSpaceNodeReqBody {
	req := &MoveSpaceNodeReqBody{}
	if builder.targetParentTokenFlag {
		req.TargetParentToken = &builder.targetParentToken

	}
	if builder.targetSpaceIdFlag {
		req.TargetSpaceId = &builder.targetSpaceId

	}
	return req
}

/**上传文件path开始**/
type MoveSpaceNodePathReqBodyBuilder struct {
	targetParentToken     string
	targetParentTokenFlag bool
	targetSpaceId         string
	targetSpaceIdFlag     bool
}

func NewMoveSpaceNodePathReqBodyBuilder() *MoveSpaceNodePathReqBodyBuilder {
	builder := &MoveSpaceNodePathReqBodyBuilder{}
	return builder
}
func (builder *MoveSpaceNodePathReqBodyBuilder) TargetParentToken(targetParentToken string) *MoveSpaceNodePathReqBodyBuilder {
	builder.targetParentToken = targetParentToken
	builder.targetParentTokenFlag = true
	return builder
}
func (builder *MoveSpaceNodePathReqBodyBuilder) TargetSpaceId(targetSpaceId string) *MoveSpaceNodePathReqBodyBuilder {
	builder.targetSpaceId = targetSpaceId
	builder.targetSpaceIdFlag = true
	return builder
}

func (builder *MoveSpaceNodePathReqBodyBuilder) Build() (*MoveSpaceNodeReqBody, error) {
	req := &MoveSpaceNodeReqBody{}
	if builder.targetParentTokenFlag {
		req.TargetParentToken = &builder.targetParentToken

	}
	if builder.targetSpaceIdFlag {
		req.TargetSpaceId = &builder.targetSpaceId

	}
	return req, nil
}

/**上传文件path结束**/

/*1.4 生成请求的builder结构体*/
type MoveSpaceNodeReqBuilder struct {
	spaceId       string
	spaceIdFlag   bool
	nodeToken     string
	nodeTokenFlag bool
	body          *MoveSpaceNodeReqBody
	bodyFlag      bool
}

// 生成请求的New构造器
func NewMoveSpaceNodeReqBuilder() *MoveSpaceNodeReqBuilder {
	builder := &MoveSpaceNodeReqBuilder{}
	return builder
}

/*1.5 生成请求的builder属性方法*/
func (builder *MoveSpaceNodeReqBuilder) SpaceId(spaceId string) *MoveSpaceNodeReqBuilder {
	builder.spaceId = spaceId
	builder.spaceIdFlag = true
	return builder
}
func (builder *MoveSpaceNodeReqBuilder) NodeToken(nodeToken string) *MoveSpaceNodeReqBuilder {
	builder.nodeToken = nodeToken
	builder.nodeTokenFlag = true
	return builder
}
func (builder *MoveSpaceNodeReqBuilder) Body(body *MoveSpaceNodeReqBody) *MoveSpaceNodeReqBuilder {
	builder.body = body
	builder.bodyFlag = true
	return builder
}

/*1.5 生成请求的builder的build方法*/
func (builder *MoveSpaceNodeReqBuilder) Build() *MoveSpaceNodeReq {
	req := &MoveSpaceNodeReq{}
	if builder.spaceIdFlag {
		req.SpaceId = builder.spaceId
	}
	if builder.nodeTokenFlag {
		req.NodeToken = builder.nodeToken
	}
	if builder.bodyFlag {
		req.Body = builder.body
	}
	return req
}

type MoveSpaceNodeReqBody struct {
	TargetParentToken *string `json:"target_parent_token,omitempty"`
	TargetSpaceId     *string `json:"target_space_id,omitempty"`
}

type MoveSpaceNodeReq struct {
	SpaceId   string                `path:"space_id"`
	NodeToken string                `path:"node_token"`
	Body      *MoveSpaceNodeReqBody `body:""`
}

type MoveSpaceNodeRespData struct {
	Node *Node `json:"node,omitempty"`
}

type MoveSpaceNodeResp struct {
	*core.RawResponse `json:"-"`
	core.CodeError
	Data *MoveSpaceNodeRespData `json:"data"`
}

func (resp *MoveSpaceNodeResp) Success() bool {
	return resp.Code == 0
}

type MoveDocsToWikiSpaceNodeReqBodyBuilder struct {
	parentWikiToken     string
	parentWikiTokenFlag bool
	objType             string
	objTypeFlag         bool
	objToken            string
	objTokenFlag        bool
	apply               bool
	applyFlag           bool
}

// 生成body的New构造器
func NewMoveDocsToWikiSpaceNodeReqBodyBuilder() *MoveDocsToWikiSpaceNodeReqBodyBuilder {
	builder := &MoveDocsToWikiSpaceNodeReqBodyBuilder{}
	return builder
}

/*1.2 生成body的builder属性方法*/
func (builder *MoveDocsToWikiSpaceNodeReqBodyBuilder) ParentWikiToken(parentWikiToken string) *MoveDocsToWikiSpaceNodeReqBodyBuilder {
	builder.parentWikiToken = parentWikiToken
	builder.parentWikiTokenFlag = true
	return builder
}
func (builder *MoveDocsToWikiSpaceNodeReqBodyBuilder) ObjType(objType string) *MoveDocsToWikiSpaceNodeReqBodyBuilder {
	builder.objType = objType
	builder.objTypeFlag = true
	return builder
}
func (builder *MoveDocsToWikiSpaceNodeReqBodyBuilder) ObjToken(objToken string) *MoveDocsToWikiSpaceNodeReqBodyBuilder {
	builder.objToken = objToken
	builder.objTokenFlag = true
	return builder
}
func (builder *MoveDocsToWikiSpaceNodeReqBodyBuilder) Apply(apply bool) *MoveDocsToWikiSpaceNodeReqBodyBuilder {
	builder.apply = apply
	builder.applyFlag = true
	return builder
}

/*1.3 生成body的build方法*/
func (builder *MoveDocsToWikiSpaceNodeReqBodyBuilder) Build() *MoveDocsToWikiSpaceNodeReqBody {
	req := &MoveDocsToWikiSpaceNodeReqBody{}
	if builder.parentWikiTokenFlag {
		req.ParentWikiToken = &builder.parentWikiToken

	}
	if builder.objTypeFlag {
		req.ObjType = &builder.objType

	}
	if builder.objTokenFlag {
		req.ObjToken = &builder.objToken

	}
	if builder.applyFlag {
		req.Apply = &builder.apply

	}
	return req
}

/**上传文件path开始**/
type MoveDocsToWikiSpaceNodePathReqBodyBuilder struct {
	parentWikiToken     string
	parentWikiTokenFlag bool
	objType             string
	objTypeFlag         bool
	objToken            string
	objTokenFlag        bool
	apply               bool
	applyFlag           bool
}

func NewMoveDocsToWikiSpaceNodePathReqBodyBuilder() *MoveDocsToWikiSpaceNodePathReqBodyBuilder {
	builder := &MoveDocsToWikiSpaceNodePathReqBodyBuilder{}
	return builder
}
func (builder *MoveDocsToWikiSpaceNodePathReqBodyBuilder) ParentWikiToken(parentWikiToken string) *MoveDocsToWikiSpaceNodePathReqBodyBuilder {
	builder.parentWikiToken = parentWikiToken
	builder.parentWikiTokenFlag = true
	return builder
}
func (builder *MoveDocsToWikiSpaceNodePathReqBodyBuilder) ObjType(objType string) *MoveDocsToWikiSpaceNodePathReqBodyBuilder {
	builder.objType = objType
	builder.objTypeFlag = true
	return builder
}
func (builder *MoveDocsToWikiSpaceNodePathReqBodyBuilder) ObjToken(objToken string) *MoveDocsToWikiSpaceNodePathReqBodyBuilder {
	builder.objToken = objToken
	builder.objTokenFlag = true
	return builder
}
func (builder *MoveDocsToWikiSpaceNodePathReqBodyBuilder) Apply(apply bool) *MoveDocsToWikiSpaceNodePathReqBodyBuilder {
	builder.apply = apply
	builder.applyFlag = true
	return builder
}

func (builder *MoveDocsToWikiSpaceNodePathReqBodyBuilder) Build() (*MoveDocsToWikiSpaceNodeReqBody, error) {
	req := &MoveDocsToWikiSpaceNodeReqBody{}
	if builder.parentWikiTokenFlag {
		req.ParentWikiToken = &builder.parentWikiToken

	}
	if builder.objTypeFlag {
		req.ObjType = &builder.objType

	}
	if builder.objTokenFlag {
		req.ObjToken = &builder.objToken

	}
	if builder.applyFlag {
		req.Apply = &builder.apply

	}
	return req, nil
}

/**上传文件path结束**/

/*1.4 生成请求的builder结构体*/
type MoveDocsToWikiSpaceNodeReqBuilder struct {
	spaceId     int64
	spaceIdFlag bool
	body        *MoveDocsToWikiSpaceNodeReqBody
	bodyFlag    bool
}

// 生成请求的New构造器
func NewMoveDocsToWikiSpaceNodeReqBuilder() *MoveDocsToWikiSpaceNodeReqBuilder {
	builder := &MoveDocsToWikiSpaceNodeReqBuilder{}
	return builder
}

/*1.5 生成请求的builder属性方法*/
func (builder *MoveDocsToWikiSpaceNodeReqBuilder) SpaceId(spaceId int64) *MoveDocsToWikiSpaceNodeReqBuilder {
	builder.spaceId = spaceId
	builder.spaceIdFlag = true
	return builder
}
func (builder *MoveDocsToWikiSpaceNodeReqBuilder) Body(body *MoveDocsToWikiSpaceNodeReqBody) *MoveDocsToWikiSpaceNodeReqBuilder {
	builder.body = body
	builder.bodyFlag = true
	return builder
}

/*1.5 生成请求的builder的build方法*/
func (builder *MoveDocsToWikiSpaceNodeReqBuilder) Build() *MoveDocsToWikiSpaceNodeReq {
	req := &MoveDocsToWikiSpaceNodeReq{}
	if builder.spaceIdFlag {
		req.SpaceId = builder.spaceId
	}
	if builder.bodyFlag {
		req.Body = builder.body
	}
	return req
}

type MoveDocsToWikiSpaceNodeReqBody struct {
	ParentWikiToken *string `json:"parent_wiki_token,omitempty"`
	ObjType         *string `json:"obj_type,omitempty"`
	ObjToken        *string `json:"obj_token,omitempty"`
	Apply           *bool   `json:"apply,omitempty"`
}

type MoveDocsToWikiSpaceNodeReq struct {
	SpaceId int64                           `path:"space_id"`
	Body    *MoveDocsToWikiSpaceNodeReqBody `body:""`
}

type MoveDocsToWikiSpaceNodeRespData struct {
	WikiToken *string `json:"wiki_token,omitempty"`
	TaskId    *string `json:"task_id,omitempty"`
	Applied   *bool   `json:"applied,omitempty"`
}

type MoveDocsToWikiSpaceNodeResp struct {
	*core.RawResponse `json:"-"`
	core.CodeError
	Data *MoveDocsToWikiSpaceNodeRespData `json:"data"`
}

func (resp *MoveDocsToWikiSpaceNodeResp) Success() bool {
	return resp.Code == 0
}

type UpdateTitleSpaceNodeReqBodyBuilder struct {
	title     string
	titleFlag bool
}

// 生成body的New构造器
func NewUpdateTitleSpaceNodeReqBodyBuilder() *UpdateTitleSpaceNodeReqBodyBuilder {
	builder := &UpdateTitleSpaceNodeReqBodyBuilder{}
	return builder
}

/*1.2 生成body的builder属性方法*/
func (builder *UpdateTitleSpaceNodeReqBodyBuilder) Title(title string) *UpdateTitleSpaceNodeReqBodyBuilder {
	builder.title = title
	builder.titleFlag = true
	return builder
}

/*1.3 生成body的build方法*/
func (builder *UpdateTitleSpaceNodeReqBodyBuilder) Build() *UpdateTitleSpaceNodeReqBody {
	req := &UpdateTitleSpaceNodeReqBody{}
	if builder.titleFlag {
		req.Title = &builder.title

	}
	return req
}

/**上传文件path开始**/
type UpdateTitleSpaceNodePathReqBodyBuilder struct {
	title     string
	titleFlag bool
}

func NewUpdateTitleSpaceNodePathReqBodyBuilder() *UpdateTitleSpaceNodePathReqBodyBuilder {
	builder := &UpdateTitleSpaceNodePathReqBodyBuilder{}
	return builder
}
func (builder *UpdateTitleSpaceNodePathReqBodyBuilder) Title(title string) *UpdateTitleSpaceNodePathReqBodyBuilder {
	builder.title = title
	builder.titleFlag = true
	return builder
}

func (builder *UpdateTitleSpaceNodePathReqBodyBuilder) Build() (*UpdateTitleSpaceNodeReqBody, error) {
	req := &UpdateTitleSpaceNodeReqBody{}
	if builder.titleFlag {
		req.Title = &builder.title

	}
	return req, nil
}

/**上传文件path结束**/

/*1.4 生成请求的builder结构体*/
type UpdateTitleSpaceNodeReqBuilder struct {
	spaceId       int64
	spaceIdFlag   bool
	nodeToken     string
	nodeTokenFlag bool
	body          *UpdateTitleSpaceNodeReqBody
	bodyFlag      bool
}

// 生成请求的New构造器
func NewUpdateTitleSpaceNodeReqBuilder() *UpdateTitleSpaceNodeReqBuilder {
	builder := &UpdateTitleSpaceNodeReqBuilder{}
	return builder
}

/*1.5 生成请求的builder属性方法*/
func (builder *UpdateTitleSpaceNodeReqBuilder) SpaceId(spaceId int64) *UpdateTitleSpaceNodeReqBuilder {
	builder.spaceId = spaceId
	builder.spaceIdFlag = true
	return builder
}
func (builder *UpdateTitleSpaceNodeReqBuilder) NodeToken(nodeToken string) *UpdateTitleSpaceNodeReqBuilder {
	builder.nodeToken = nodeToken
	builder.nodeTokenFlag = true
	return builder
}
func (builder *UpdateTitleSpaceNodeReqBuilder) Body(body *UpdateTitleSpaceNodeReqBody) *UpdateTitleSpaceNodeReqBuilder {
	builder.body = body
	builder.bodyFlag = true
	return builder
}

/*1.5 生成请求的builder的build方法*/
func (builder *UpdateTitleSpaceNodeReqBuilder) Build() *UpdateTitleSpaceNodeReq {
	req := &UpdateTitleSpaceNodeReq{}
	if builder.spaceIdFlag {
		req.SpaceId = builder.spaceId
	}
	if builder.nodeTokenFlag {
		req.NodeToken = builder.nodeToken
	}
	if builder.bodyFlag {
		req.Body = builder.body
	}
	return req
}

type UpdateTitleSpaceNodeReqBody struct {
	Title *string `json:"title,omitempty"`
}

type UpdateTitleSpaceNodeReq struct {
	SpaceId   int64                        `path:"space_id"`
	NodeToken string                       `path:"node_token"`
	Body      *UpdateTitleSpaceNodeReqBody `body:""`
}

type UpdateTitleSpaceNodeResp struct {
	*core.RawResponse `json:"-"`
	core.CodeError
}

func (resp *UpdateTitleSpaceNodeResp) Success() bool {
	return resp.Code == 0
}

/*1.4 生成请求的builder结构体*/
type UpdateSpaceSettingReqBuilder struct {
	spaceId     string
	spaceIdFlag bool
	setting     *Setting
	settingFlag bool
}

// 生成请求的New构造器
func NewUpdateSpaceSettingReqBuilder() *UpdateSpaceSettingReqBuilder {
	builder := &UpdateSpaceSettingReqBuilder{}
	return builder
}

/*1.5 生成请求的builder属性方法*/
func (builder *UpdateSpaceSettingReqBuilder) SpaceId(spaceId string) *UpdateSpaceSettingReqBuilder {
	builder.spaceId = spaceId
	builder.spaceIdFlag = true
	return builder
}
func (builder *UpdateSpaceSettingReqBuilder) Setting(setting *Setting) *UpdateSpaceSettingReqBuilder {
	builder.setting = setting
	builder.settingFlag = true
	return builder
}

/*1.5 生成请求的builder的build方法*/
func (builder *UpdateSpaceSettingReqBuilder) Build() *UpdateSpaceSettingReq {
	req := &UpdateSpaceSettingReq{}
	if builder.spaceIdFlag {
		req.SpaceId = builder.spaceId
	}
	if builder.settingFlag {
		req.Setting = builder.setting
	}
	return req
}

type UpdateSpaceSettingReq struct {
	SpaceId string   `path:"space_id"`
	Setting *Setting `body:""`
}

type UpdateSpaceSettingRespData struct {
	Setting *Setting `json:"setting,omitempty"`
}

type UpdateSpaceSettingResp struct {
	*core.RawResponse `json:"-"`
	core.CodeError
	Data *UpdateSpaceSettingRespData `json:"data"`
}

func (resp *UpdateSpaceSettingResp) Success() bool {
	return resp.Code == 0
}

/*1.4 生成请求的builder结构体*/
type GetTaskReqBuilder struct {
	taskId       string
	taskIdFlag   bool
	taskType     string
	taskTypeFlag bool
}

// 生成请求的New构造器
func NewGetTaskReqBuilder() *GetTaskReqBuilder {
	builder := &GetTaskReqBuilder{}
	return builder
}

/*1.5 生成请求的builder属性方法*/
func (builder *GetTaskReqBuilder) TaskId(taskId string) *GetTaskReqBuilder {
	builder.taskId = taskId
	builder.taskIdFlag = true
	return builder
}
func (builder *GetTaskReqBuilder) TaskType(taskType string) *GetTaskReqBuilder {
	builder.taskType = taskType
	builder.taskTypeFlag = true
	return builder
}

/*1.5 生成请求的builder的build方法*/
func (builder *GetTaskReqBuilder) Build() *GetTaskReq {
	req := &GetTaskReq{}
	if builder.taskIdFlag {
		req.TaskId = builder.taskId
	}
	if builder.taskTypeFlag {
		req.TaskType = &builder.taskType
	}
	return req
}

type GetTaskReq struct {
	TaskId   string  `path:"task_id"`
	TaskType *string `query:"task_type"`
}

type GetTaskRespData struct {
	Task *TaskResult `json:"task,omitempty"`
}

type GetTaskResp struct {
	*core.RawResponse `json:"-"`
	core.CodeError
	Data *GetTaskRespData `json:"data"`
}

func (resp *GetTaskResp) Success() bool {
	return resp.Code == 0
}

/**生成消息事件结构体 **/

/* 生成请求的builder构造器*/
/*1.1 生成body的builder结构体*/
/**如果是分页查询，则添加迭代器定义**/
type ListSpaceIterator struct {
	nextPageToken *string
	items         []*Space
	index         int
	limit         int
	ctx           context.Context
	req           *ListSpaceReq
	listFunc      func(ctx context.Context, req *ListSpaceReq, options ...core.RequestOptionFunc) (*ListSpaceResp, error)
	options       []core.RequestOptionFunc
	curlNum       int
}

func (iterator *ListSpaceIterator) Next() (bool, *Space, error) {
	// 达到最大量，则返回
	if iterator.limit > 0 && iterator.curlNum > iterator.limit {
		return false, nil, nil
	}

	// 为0则拉取数据
	if iterator.index == 0 || iterator.index >= len(iterator.items) {
		if iterator.index != 0 && iterator.nextPageToken == nil {
			return false, nil, nil
		}
		if iterator.nextPageToken != nil {
			iterator.req.PageToken = iterator.nextPageToken
		}
		resp, err := iterator.listFunc(iterator.ctx, iterator.req, iterator.options...)
		if err != nil {
			return false, nil, err
		}

		if resp.Code != 0 {
			return false, nil, errors.New(fmt.Sprintf("Code:%d,Msg:%s", resp.Code, resp.Msg))
		}

		if len(resp.Data.Items) == 0 {
			return false, nil, nil
		}

		iterator.nextPageToken = resp.Data.PageToken
		iterator.items = resp.Data.Items
		iterator.index = 0
	}

	block := iterator.items[iterator.index]
	iterator.index++
	iterator.curlNum++
	return true, block, nil
}

func (iterator *ListSpaceIterator) NextPageToken() *string {
	return iterator.nextPageToken
}

/**如果是分页查询，则添加迭代器定义**/
type ListSpaceNodeIterator struct {
	nextPageToken *string
	items         []*Node
	index         int
	limit         int
	ctx           context.Context
	req           *ListSpaceNodeReq
	listFunc      func(ctx context.Context, req *ListSpaceNodeReq, options ...core.RequestOptionFunc) (*ListSpaceNodeResp, error)
	options       []core.RequestOptionFunc
	curlNum       int
}

func (iterator *ListSpaceNodeIterator) Next() (bool, *Node, error) {
	// 达到最大量，则返回
	if iterator.limit > 0 && iterator.curlNum > iterator.limit {
		return false, nil, nil
	}

	// 为0则拉取数据
	if iterator.index == 0 || iterator.index >= len(iterator.items) {
		if iterator.index != 0 && iterator.nextPageToken == nil {
			return false, nil, nil
		}
		if iterator.nextPageToken != nil {
			iterator.req.PageToken = iterator.nextPageToken
		}
		resp, err := iterator.listFunc(iterator.ctx, iterator.req, iterator.options...)
		if err != nil {
			return false, nil, err
		}

		if resp.Code != 0 {
			return false, nil, errors.New(fmt.Sprintf("Code:%d,Msg:%s", resp.Code, resp.Msg))
		}

		if len(resp.Data.Items) == 0 {
			return false, nil, nil
		}

		iterator.nextPageToken = resp.Data.PageToken
		iterator.items = resp.Data.Items
		iterator.index = 0
	}

	block := iterator.items[iterator.index]
	iterator.index++
	iterator.curlNum++
	return true, block, nil
}

func (iterator *ListSpaceNodeIterator) NextPageToken() *string {
	return iterator.nextPageToken
}
