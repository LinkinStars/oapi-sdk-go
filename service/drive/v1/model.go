// Code generated by lark suite oapi sdk gen
package v1

import (
	"github.com/larksuite/oapi-sdk-go/api/core/tools"
)

type Collaborator struct {
	MemberType      string   `json:"member_type,omitempty"`
	MemberOpenId    string   `json:"member_open_id,omitempty"`
	MemberUserId    string   `json:"member_user_id,omitempty"`
	Perm            string   `json:"perm,omitempty"`
	ForceSendFields []string `json:"-"`
}

func (s *Collaborator) MarshalJSON() ([]byte, error) {
	type cp Collaborator
	raw := cp(*s)
	return tools.MarshalJSON(raw, s.ForceSendFields)
}

type DocsLink struct {
	Url             string   `json:"url,omitempty"`
	ForceSendFields []string `json:"-"`
}

func (s *DocsLink) MarshalJSON() ([]byte, error) {
	type cp DocsLink
	raw := cp(*s)
	return tools.MarshalJSON(raw, s.ForceSendFields)
}

type File struct {
	FileToken       string   `json:"file_token,omitempty"`
	FileName        string   `json:"file_name,omitempty"`
	Size            int      `json:"size,omitempty"`
	MimeType        string   `json:"mime_type,omitempty"`
	ForceSendFields []string `json:"-"`
}

func (s *File) MarshalJSON() ([]byte, error) {
	type cp File
	raw := cp(*s)
	return tools.MarshalJSON(raw, s.ForceSendFields)
}

type FileComment struct {
	CommentId       string     `json:"comment_id,omitempty"`
	UserId          string     `json:"user_id,omitempty"`
	CreateTime      int        `json:"create_time,omitempty"`
	UpdateTime      int        `json:"update_time,omitempty"`
	IsSolved        bool       `json:"is_solved,omitempty"`
	SolvedTime      int        `json:"solved_time,omitempty"`
	SolverUserId    string     `json:"solver_user_id,omitempty"`
	ReplyList       *ReplyList `json:"reply_list,omitempty"`
	ForceSendFields []string   `json:"-"`
}

func (s *FileComment) MarshalJSON() ([]byte, error) {
	type cp FileComment
	raw := cp(*s)
	return tools.MarshalJSON(raw, s.ForceSendFields)
}

type FileCommentReply struct {
	ReplyId         string        `json:"reply_id,omitempty"`
	UserId          string        `json:"user_id,omitempty"`
	CreateTime      int           `json:"create_time,omitempty"`
	UpdateTime      int           `json:"update_time,omitempty"`
	Content         *ReplyContent `json:"content,omitempty"`
	ForceSendFields []string      `json:"-"`
}

func (s *FileCommentReply) MarshalJSON() ([]byte, error) {
	type cp FileCommentReply
	raw := cp(*s)
	return tools.MarshalJSON(raw, s.ForceSendFields)
}

type Media struct {
	FileToken       string   `json:"file_token,omitempty"`
	FileName        string   `json:"file_name,omitempty"`
	Size            int      `json:"size,omitempty"`
	MimeType        string   `json:"mime_type,omitempty"`
	ForceSendFields []string `json:"-"`
}

func (s *Media) MarshalJSON() ([]byte, error) {
	type cp Media
	raw := cp(*s)
	return tools.MarshalJSON(raw, s.ForceSendFields)
}

type Member struct {
	MemberType      string   `json:"member_type,omitempty"`
	MemberId        string   `json:"member_id,omitempty"`
	Perm            string   `json:"perm,omitempty"`
	ForceSendFields []string `json:"-"`
}

func (s *Member) MarshalJSON() ([]byte, error) {
	type cp Member
	raw := cp(*s)
	return tools.MarshalJSON(raw, s.ForceSendFields)
}

type Owner struct {
	MemberType      string   `json:"member_type,omitempty"`
	MemberId        string   `json:"member_id,omitempty"`
	ForceSendFields []string `json:"-"`
}

func (s *Owner) MarshalJSON() ([]byte, error) {
	type cp Owner
	raw := cp(*s)
	return tools.MarshalJSON(raw, s.ForceSendFields)
}

type Person struct {
	UserId          string   `json:"user_id,omitempty"`
	ForceSendFields []string `json:"-"`
}

func (s *Person) MarshalJSON() ([]byte, error) {
	type cp Person
	raw := cp(*s)
	return tools.MarshalJSON(raw, s.ForceSendFields)
}

type ReplyContent struct {
	Elements        []*ReplyElement `json:"elements,omitempty"`
	ForceSendFields []string        `json:"-"`
}

func (s *ReplyContent) MarshalJSON() ([]byte, error) {
	type cp ReplyContent
	raw := cp(*s)
	return tools.MarshalJSON(raw, s.ForceSendFields)
}

type ReplyElement struct {
	Type            string    `json:"type,omitempty"`
	TextRun         *TextRun  `json:"text_run,omitempty"`
	DocsLink        *DocsLink `json:"docs_link,omitempty"`
	Person          *Person   `json:"person,omitempty"`
	ForceSendFields []string  `json:"-"`
}

func (s *ReplyElement) MarshalJSON() ([]byte, error) {
	type cp ReplyElement
	raw := cp(*s)
	return tools.MarshalJSON(raw, s.ForceSendFields)
}

type ReplyList struct {
	Replies         []*FileCommentReply `json:"replies,omitempty"`
	ForceSendFields []string            `json:"-"`
}

func (s *ReplyList) MarshalJSON() ([]byte, error) {
	type cp ReplyList
	raw := cp(*s)
	return tools.MarshalJSON(raw, s.ForceSendFields)
}

type TextRun struct {
	Text            string   `json:"text,omitempty"`
	ForceSendFields []string `json:"-"`
}

func (s *TextRun) MarshalJSON() ([]byte, error) {
	type cp TextRun
	raw := cp(*s)
	return tools.MarshalJSON(raw, s.ForceSendFields)
}

type TmpDownloadUrl struct {
	FileToken       string   `json:"file_token,omitempty"`
	TmpDownloadUrl  string   `json:"tmp_download_url,omitempty"`
	ForceSendFields []string `json:"-"`
}

func (s *TmpDownloadUrl) MarshalJSON() ([]byte, error) {
	type cp TmpDownloadUrl
	raw := cp(*s)
	return tools.MarshalJSON(raw, s.ForceSendFields)
}

type TokenType struct {
	Token           string   `json:"token,omitempty"`
	Type            string   `json:"type,omitempty"`
	ForceSendFields []string `json:"-"`
}

func (s *TokenType) MarshalJSON() ([]byte, error) {
	type cp TokenType
	raw := cp(*s)
	return tools.MarshalJSON(raw, s.ForceSendFields)
}

type FileStatistics struct {
	Uv              int      `json:"uv,omitempty"`
	Pv              int      `json:"pv,omitempty"`
	LikeCount       int      `json:"like_count,omitempty"`
	Timestamp       int      `json:"timestamp,omitempty"`
	ForceSendFields []string `json:"-"`
}

func (s *FileStatistics) MarshalJSON() ([]byte, error) {
	type cp FileStatistics
	raw := cp(*s)
	return tools.MarshalJSON(raw, s.ForceSendFields)
}

type FileUploadInfo struct {
	FileName        string   `json:"file_name,omitempty"`
	ParentType      string   `json:"parent_type,omitempty"`
	ParentNode      string   `json:"parent_node,omitempty"`
	Size            int      `json:"size,omitempty"`
	ForceSendFields []string `json:"-"`
}

func (s *FileUploadInfo) MarshalJSON() ([]byte, error) {
	type cp FileUploadInfo
	raw := cp(*s)
	return tools.MarshalJSON(raw, s.ForceSendFields)
}

type ImportTask struct {
	Ticket          string                `json:"ticket,omitempty"`
	FileExtension   string                `json:"file_extension,omitempty"`
	FileToken       string                `json:"file_token,omitempty"`
	Type            string                `json:"type,omitempty"`
	FileName        string                `json:"file_name,omitempty"`
	Point           *ImportTaskMountPoint `json:"point,omitempty"`
	JobStatus       int                   `json:"job_status,omitempty"`
	JobErrorMsg     string                `json:"job_error_msg,omitempty"`
	Token           string                `json:"token,omitempty"`
	Url             string                `json:"url,omitempty"`
	Extra           []string              `json:"extra,omitempty"`
	ForceSendFields []string              `json:"-"`
}

func (s *ImportTask) MarshalJSON() ([]byte, error) {
	type cp ImportTask
	raw := cp(*s)
	return tools.MarshalJSON(raw, s.ForceSendFields)
}

type ImportTaskMountPoint struct {
	MountType       int      `json:"mount_type,omitempty"`
	MountKey        string   `json:"mount_key,omitempty"`
	ForceSendFields []string `json:"-"`
}

func (s *ImportTaskMountPoint) MarshalJSON() ([]byte, error) {
	type cp ImportTaskMountPoint
	raw := cp(*s)
	return tools.MarshalJSON(raw, s.ForceSendFields)
}

type MediaUploadInfo struct {
	FileName        string   `json:"file_name,omitempty"`
	ParentType      string   `json:"parent_type,omitempty"`
	ParentNode      string   `json:"parent_node,omitempty"`
	Size            int      `json:"size,omitempty"`
	Extra           string   `json:"extra,omitempty"`
	ForceSendFields []string `json:"-"`
}

func (s *MediaUploadInfo) MarshalJSON() ([]byte, error) {
	type cp MediaUploadInfo
	raw := cp(*s)
	return tools.MarshalJSON(raw, s.ForceSendFields)
}

type PermissionPublic struct {
	ExternalAccess  bool     `json:"external_access,omitempty"`
	SecurityEntity  string   `json:"security_entity,omitempty"`
	CommentEntity   string   `json:"comment_entity,omitempty"`
	ShareEntity     string   `json:"share_entity,omitempty"`
	LinkShareEntity string   `json:"link_share_entity,omitempty"`
	InviteExternal  bool     `json:"invite_external,omitempty"`
	ForceSendFields []string `json:"-"`
}

func (s *PermissionPublic) MarshalJSON() ([]byte, error) {
	type cp PermissionPublic
	raw := cp(*s)
	return tools.MarshalJSON(raw, s.ForceSendFields)
}

type UploadInfo struct {
	FileName        string   `json:"file_name,omitempty"`
	ParentType      string   `json:"parent_type,omitempty"`
	ParentNode      string   `json:"parent_node,omitempty"`
	Size            int      `json:"size,omitempty"`
	ForceSendFields []string `json:"-"`
}

func (s *UploadInfo) MarshalJSON() ([]byte, error) {
	type cp UploadInfo
	raw := cp(*s)
	return tools.MarshalJSON(raw, s.ForceSendFields)
}

type FileUploadFinishReqBody struct {
	UploadId        string   `json:"upload_id,omitempty"`
	BlockNum        int      `json:"block_num,omitempty"`
	ForceSendFields []string `json:"-"`
}

func (s *FileUploadFinishReqBody) MarshalJSON() ([]byte, error) {
	type cp FileUploadFinishReqBody
	raw := cp(*s)
	return tools.MarshalJSON(raw, s.ForceSendFields)
}

type FileUploadFinishResult struct {
	FileToken string `json:"file_token,omitempty"`
}

type FileUploadPrepareResult struct {
	UploadId  string `json:"upload_id,omitempty"`
	BlockSize int    `json:"block_size,omitempty"`
	BlockNum  int    `json:"block_num,omitempty"`
}

type MediaUploadAllResult struct {
	FileToken string `json:"file_token,omitempty"`
}

type MediaUploadFinishReqBody struct {
	UploadId        string   `json:"upload_id,omitempty"`
	BlockNum        int      `json:"block_num,omitempty"`
	ForceSendFields []string `json:"-"`
}

func (s *MediaUploadFinishReqBody) MarshalJSON() ([]byte, error) {
	type cp MediaUploadFinishReqBody
	raw := cp(*s)
	return tools.MarshalJSON(raw, s.ForceSendFields)
}

type MediaUploadFinishResult struct {
	FileToken string `json:"file_token,omitempty"`
}

type FileUploadAllResult struct {
	FileToken string `json:"file_token,omitempty"`
}

type MediaUploadPrepareResult struct {
	UploadId  string `json:"upload_id,omitempty"`
	BlockSize int    `json:"block_size,omitempty"`
	BlockNum  int    `json:"block_num,omitempty"`
}

type MediaBatchGetTmpDownloadUrlResult struct {
	TmpDownloadUrls []*TmpDownloadUrl `json:"tmp_download_urls,omitempty"`
}

type FileCommentReplyUpdateReqBody struct {
	Content         *ReplyContent `json:"content,omitempty"`
	ForceSendFields []string      `json:"-"`
}

func (s *FileCommentReplyUpdateReqBody) MarshalJSON() ([]byte, error) {
	type cp FileCommentReplyUpdateReqBody
	raw := cp(*s)
	return tools.MarshalJSON(raw, s.ForceSendFields)
}

type FileCommentListResult struct {
	HasMore   bool           `json:"has_more,omitempty"`
	PageToken string         `json:"page_token,omitempty"`
	Items     []*FileComment `json:"items,omitempty"`
}

type FileCommentPatchReqBody struct {
	IsSolved        bool     `json:"is_solved,omitempty"`
	ForceSendFields []string `json:"-"`
}

func (s *FileCommentPatchReqBody) MarshalJSON() ([]byte, error) {
	type cp FileCommentPatchReqBody
	raw := cp(*s)
	return tools.MarshalJSON(raw, s.ForceSendFields)
}

type PermissionMemberCreateResult struct {
	Member *Member `json:"member,omitempty"`
}

type PermissionPublicPatchResult struct {
	PermissionPublic *PermissionPublic `json:"permission_public,omitempty"`
}

type ImportTaskCreateResult struct {
	Ticket string `json:"ticket,omitempty"`
}

type ImportTaskGetResult struct {
	Result *ImportTask `json:"result,omitempty"`
}

type FileStatisticsGetResult struct {
	FileToken  string          `json:"file_token,omitempty"`
	FileType   string          `json:"file_type,omitempty"`
	Statistics *FileStatistics `json:"statistics,omitempty"`
}

type PermissionMemberUpdateResult struct {
	Member *Member `json:"member,omitempty"`
}
