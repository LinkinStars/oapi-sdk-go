// Package drive code generated by oapi sdk gen
package drive

import (
	"bytes"
	"context"
	"net/http"

	"github.com/larksuite/oapi-sdk-go/core"
)

/**
构建业务域服务实例
**/
func NewService(httpClient *http.Client, config *core.Config) *DriveService {
	d := &DriveService{httpClient: httpClient, config: config}
	d.ExportTask = &exportTask{service: d}
	d.File = &file{service: d}
	d.FileComment = &fileComment{service: d}
	d.FileCommentReply = &fileCommentReply{service: d}
	d.FileStatistics = &fileStatistics{service: d}
	d.FileSubscription = &fileSubscription{service: d}
	d.ImportTask = &importTask{service: d}
	d.Media = &media{service: d}
	d.Meta = &meta{service: d}
	d.PermissionMember = &permissionMember{service: d}
	d.PermissionPublic = &permissionPublic{service: d}
	return d
}

/**
业务域服务定义
**/
type DriveService struct {
	httpClient       *http.Client
	config           *core.Config
	ExportTask       *exportTask
	File             *file
	FileComment      *fileComment
	FileCommentReply *fileCommentReply
	FileStatistics   *fileStatistics
	FileSubscription *fileSubscription
	ImportTask       *importTask
	Media            *media
	Meta             *meta
	PermissionMember *permissionMember
	PermissionPublic *permissionPublic
}

/**
资源服务定义
**/
type exportTask struct {
	service *DriveService
}
type file struct {
	service *DriveService
}
type fileComment struct {
	service *DriveService
}
type fileCommentReply struct {
	service *DriveService
}
type fileStatistics struct {
	service *DriveService
}
type fileSubscription struct {
	service *DriveService
}
type importTask struct {
	service *DriveService
}
type media struct {
	service *DriveService
}
type meta struct {
	service *DriveService
}
type permissionMember struct {
	service *DriveService
}
type permissionPublic struct {
	service *DriveService
}

/**
资源服务方法定义
**/
func (e *exportTask) Create(ctx context.Context, req *CreateExportTaskReq, options ...core.RequestOptionFunc) (*CreateExportTaskResp, error) {
	// 发起请求
	rawResp, err := core.SendRequest(ctx, e.service.config, http.MethodPost,
		"/open-apis/drive/v1/export_tasks", []core.AccessTokenType{core.AccessTokenTypeTenant, core.AccessTokenTypeUser}, req, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &CreateExportTaskResp{RawResponse: rawResp}
	err = rawResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (e *exportTask) Download(ctx context.Context, req *DownloadExportTaskReq, options ...core.RequestOptionFunc) (*DownloadExportTaskResp, error) {
	// 发起请求
	rawResp, err := core.SendRequest(ctx, e.service.config, http.MethodGet,
		"/open-apis/drive/v1/export_tasks/file/:file_token/download", []core.AccessTokenType{core.AccessTokenTypeTenant, core.AccessTokenTypeUser}, req, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &DownloadExportTaskResp{RawResponse: rawResp}
	// 如果是下载，则设置响应结果
	if rawResp.StatusCode == http.StatusOK {
		resp.File = bytes.NewBuffer(rawResp.RawBody)
		resp.FileName = core.FileNameByHeader(rawResp.Header)
		return resp, err
	}
	err = rawResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (e *exportTask) Get(ctx context.Context, req *GetExportTaskReq, options ...core.RequestOptionFunc) (*GetExportTaskResp, error) {
	// 发起请求
	rawResp, err := core.SendRequest(ctx, e.service.config, http.MethodGet,
		"/open-apis/drive/v1/export_tasks/:ticket", []core.AccessTokenType{core.AccessTokenTypeTenant, core.AccessTokenTypeUser}, req, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &GetExportTaskResp{RawResponse: rawResp}
	err = rawResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (f *file) Copy(ctx context.Context, req *CopyFileReq, options ...core.RequestOptionFunc) (*CopyFileResp, error) {
	// 发起请求
	rawResp, err := core.SendRequest(ctx, f.service.config, http.MethodPost,
		"/open-apis/drive/v1/files/:file_token/copy", []core.AccessTokenType{core.AccessTokenTypeTenant, core.AccessTokenTypeUser}, req, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &CopyFileResp{RawResponse: rawResp}
	err = rawResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (f *file) CreateFolder(ctx context.Context, req *CreateFolderFileReq, options ...core.RequestOptionFunc) (*CreateFolderFileResp, error) {
	// 发起请求
	rawResp, err := core.SendRequest(ctx, f.service.config, http.MethodPost,
		"/open-apis/drive/v1/files/create_folder", []core.AccessTokenType{core.AccessTokenTypeTenant, core.AccessTokenTypeUser}, req, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &CreateFolderFileResp{RawResponse: rawResp}
	err = rawResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (f *file) Delete(ctx context.Context, req *DeleteFileReq, options ...core.RequestOptionFunc) (*DeleteFileResp, error) {
	// 发起请求
	rawResp, err := core.SendRequest(ctx, f.service.config, http.MethodDelete,
		"/open-apis/drive/v1/files/:file_token", []core.AccessTokenType{core.AccessTokenTypeTenant, core.AccessTokenTypeUser}, req, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &DeleteFileResp{RawResponse: rawResp}
	err = rawResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (f *file) Download(ctx context.Context, req *DownloadFileReq, options ...core.RequestOptionFunc) (*DownloadFileResp, error) {
	// 发起请求
	rawResp, err := core.SendRequest(ctx, f.service.config, http.MethodGet,
		"/open-apis/drive/v1/files/:file_token/download", []core.AccessTokenType{core.AccessTokenTypeUser, core.AccessTokenTypeTenant}, req, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &DownloadFileResp{RawResponse: rawResp}
	// 如果是下载，则设置响应结果
	if rawResp.StatusCode == http.StatusOK {
		resp.File = bytes.NewBuffer(rawResp.RawBody)
		resp.FileName = core.FileNameByHeader(rawResp.Header)
		return resp, err
	}
	err = rawResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (f *file) List(ctx context.Context, req *ListFileReq, options ...core.RequestOptionFunc) (*ListFileResp, error) {
	// 发起请求
	rawResp, err := core.SendRequest(ctx, f.service.config, http.MethodGet,
		"/open-apis/drive/v1/files", []core.AccessTokenType{core.AccessTokenTypeTenant, core.AccessTokenTypeUser}, req, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &ListFileResp{RawResponse: rawResp}
	err = rawResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (f *file) Move(ctx context.Context, req *MoveFileReq, options ...core.RequestOptionFunc) (*MoveFileResp, error) {
	// 发起请求
	rawResp, err := core.SendRequest(ctx, f.service.config, http.MethodPost,
		"/open-apis/drive/v1/files/:file_token/move", []core.AccessTokenType{core.AccessTokenTypeTenant, core.AccessTokenTypeUser}, req, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &MoveFileResp{RawResponse: rawResp}
	err = rawResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (f *file) Subscribe(ctx context.Context, req *SubscribeFileReq, options ...core.RequestOptionFunc) (*SubscribeFileResp, error) {
	// 发起请求
	rawResp, err := core.SendRequest(ctx, f.service.config, http.MethodPost,
		"/open-apis/drive/v1/files/:file_token/subscribe", []core.AccessTokenType{core.AccessTokenTypeUser, core.AccessTokenTypeTenant}, req, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &SubscribeFileResp{RawResponse: rawResp}
	err = rawResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (f *file) TaskCheck(ctx context.Context, req *TaskCheckFileReq, options ...core.RequestOptionFunc) (*TaskCheckFileResp, error) {
	// 发起请求
	rawResp, err := core.SendRequest(ctx, f.service.config, http.MethodGet,
		"/open-apis/drive/v1/files/task_check", []core.AccessTokenType{core.AccessTokenTypeTenant, core.AccessTokenTypeUser}, req, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &TaskCheckFileResp{RawResponse: rawResp}
	err = rawResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (f *file) UploadAll(ctx context.Context, req *UploadAllFileReq, options ...core.RequestOptionFunc) (*UploadAllFileResp, error) {
	options = append(options, core.WithFileUpload())
	// 发起请求
	rawResp, err := core.SendRequest(ctx, f.service.config, http.MethodPost,
		"/open-apis/drive/v1/files/upload_all", []core.AccessTokenType{core.AccessTokenTypeUser, core.AccessTokenTypeTenant}, req, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &UploadAllFileResp{RawResponse: rawResp}
	err = rawResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (f *file) UploadFinish(ctx context.Context, req *UploadFinishFileReq, options ...core.RequestOptionFunc) (*UploadFinishFileResp, error) {
	// 发起请求
	rawResp, err := core.SendRequest(ctx, f.service.config, http.MethodPost,
		"/open-apis/drive/v1/files/upload_finish", []core.AccessTokenType{core.AccessTokenTypeUser, core.AccessTokenTypeTenant}, req, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &UploadFinishFileResp{RawResponse: rawResp}
	err = rawResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (f *file) UploadPart(ctx context.Context, req *UploadPartFileReq, options ...core.RequestOptionFunc) (*UploadPartFileResp, error) {
	options = append(options, core.WithFileUpload())
	// 发起请求
	rawResp, err := core.SendRequest(ctx, f.service.config, http.MethodPost,
		"/open-apis/drive/v1/files/upload_part", []core.AccessTokenType{core.AccessTokenTypeUser, core.AccessTokenTypeTenant}, req, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &UploadPartFileResp{RawResponse: rawResp}
	err = rawResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (f *file) UploadPrepare(ctx context.Context, req *UploadPrepareFileReq, options ...core.RequestOptionFunc) (*UploadPrepareFileResp, error) {
	// 发起请求
	rawResp, err := core.SendRequest(ctx, f.service.config, http.MethodPost,
		"/open-apis/drive/v1/files/upload_prepare", []core.AccessTokenType{core.AccessTokenTypeUser, core.AccessTokenTypeTenant}, req, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &UploadPrepareFileResp{RawResponse: rawResp}
	err = rawResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (f *fileComment) Create(ctx context.Context, req *CreateFileCommentReq, options ...core.RequestOptionFunc) (*CreateFileCommentResp, error) {
	// 发起请求
	rawResp, err := core.SendRequest(ctx, f.service.config, http.MethodPost,
		"/open-apis/drive/v1/files/:file_token/comments", []core.AccessTokenType{core.AccessTokenTypeUser, core.AccessTokenTypeTenant}, req, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &CreateFileCommentResp{RawResponse: rawResp}
	err = rawResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (f *fileComment) Get(ctx context.Context, req *GetFileCommentReq, options ...core.RequestOptionFunc) (*GetFileCommentResp, error) {
	// 发起请求
	rawResp, err := core.SendRequest(ctx, f.service.config, http.MethodGet,
		"/open-apis/drive/v1/files/:file_token/comments/:comment_id", []core.AccessTokenType{core.AccessTokenTypeUser, core.AccessTokenTypeTenant}, req, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &GetFileCommentResp{RawResponse: rawResp}
	err = rawResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (f *fileComment) List(ctx context.Context, req *ListFileCommentReq, options ...core.RequestOptionFunc) (*ListFileCommentResp, error) {
	// 发起请求
	rawResp, err := core.SendRequest(ctx, f.service.config, http.MethodGet,
		"/open-apis/drive/v1/files/:file_token/comments", []core.AccessTokenType{core.AccessTokenTypeUser, core.AccessTokenTypeTenant}, req, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &ListFileCommentResp{RawResponse: rawResp}
	err = rawResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}

/**如果是分页查询，则添加迭代器函数**/
func (f *fileComment) ListFileComment(ctx context.Context, req *ListFileCommentReq, options ...core.RequestOptionFunc) (*ListFileCommentIterator, error) {
	return &ListFileCommentIterator{
		ctx:      ctx,
		req:      req,
		listFunc: f.List,
		options:  options,
		limit:    req.Limit}, nil
}
func (f *fileComment) Patch(ctx context.Context, req *PatchFileCommentReq, options ...core.RequestOptionFunc) (*PatchFileCommentResp, error) {
	// 发起请求
	rawResp, err := core.SendRequest(ctx, f.service.config, http.MethodPatch,
		"/open-apis/drive/v1/files/:file_token/comments/:comment_id", []core.AccessTokenType{core.AccessTokenTypeUser, core.AccessTokenTypeTenant}, req, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &PatchFileCommentResp{RawResponse: rawResp}
	err = rawResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (f *fileCommentReply) Delete(ctx context.Context, req *DeleteFileCommentReplyReq, options ...core.RequestOptionFunc) (*DeleteFileCommentReplyResp, error) {
	// 发起请求
	rawResp, err := core.SendRequest(ctx, f.service.config, http.MethodDelete,
		"/open-apis/drive/v1/files/:file_token/comments/:comment_id/replies/:reply_id", []core.AccessTokenType{core.AccessTokenTypeUser, core.AccessTokenTypeTenant}, req, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &DeleteFileCommentReplyResp{RawResponse: rawResp}
	err = rawResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (f *fileCommentReply) Update(ctx context.Context, req *UpdateFileCommentReplyReq, options ...core.RequestOptionFunc) (*UpdateFileCommentReplyResp, error) {
	// 发起请求
	rawResp, err := core.SendRequest(ctx, f.service.config, http.MethodPut,
		"/open-apis/drive/v1/files/:file_token/comments/:comment_id/replies/:reply_id", []core.AccessTokenType{core.AccessTokenTypeUser, core.AccessTokenTypeTenant}, req, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &UpdateFileCommentReplyResp{RawResponse: rawResp}
	err = rawResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (f *fileStatistics) Get(ctx context.Context, req *GetFileStatisticsReq, options ...core.RequestOptionFunc) (*GetFileStatisticsResp, error) {
	// 发起请求
	rawResp, err := core.SendRequest(ctx, f.service.config, http.MethodGet,
		"/open-apis/drive/v1/files/:file_token/statistics", []core.AccessTokenType{core.AccessTokenTypeTenant, core.AccessTokenTypeUser}, req, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &GetFileStatisticsResp{RawResponse: rawResp}
	err = rawResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (f *fileSubscription) Create(ctx context.Context, req *CreateFileSubscriptionReq, options ...core.RequestOptionFunc) (*CreateFileSubscriptionResp, error) {
	// 发起请求
	rawResp, err := core.SendRequest(ctx, f.service.config, http.MethodPost,
		"/open-apis/drive/v1/files/:file_token/subscriptions", []core.AccessTokenType{core.AccessTokenTypeUser}, req, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &CreateFileSubscriptionResp{RawResponse: rawResp}
	err = rawResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (f *fileSubscription) Get(ctx context.Context, req *GetFileSubscriptionReq, options ...core.RequestOptionFunc) (*GetFileSubscriptionResp, error) {
	// 发起请求
	rawResp, err := core.SendRequest(ctx, f.service.config, http.MethodGet,
		"/open-apis/drive/v1/files/:file_token/subscriptions/:subscription_id", []core.AccessTokenType{core.AccessTokenTypeUser}, req, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &GetFileSubscriptionResp{RawResponse: rawResp}
	err = rawResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (f *fileSubscription) Patch(ctx context.Context, req *PatchFileSubscriptionReq, options ...core.RequestOptionFunc) (*PatchFileSubscriptionResp, error) {
	// 发起请求
	rawResp, err := core.SendRequest(ctx, f.service.config, http.MethodPatch,
		"/open-apis/drive/v1/files/:file_token/subscriptions/:subscription_id", []core.AccessTokenType{core.AccessTokenTypeUser}, req, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &PatchFileSubscriptionResp{RawResponse: rawResp}
	err = rawResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (i *importTask) Create(ctx context.Context, req *CreateImportTaskReq, options ...core.RequestOptionFunc) (*CreateImportTaskResp, error) {
	// 发起请求
	rawResp, err := core.SendRequest(ctx, i.service.config, http.MethodPost,
		"/open-apis/drive/v1/import_tasks", []core.AccessTokenType{core.AccessTokenTypeUser, core.AccessTokenTypeTenant}, req, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &CreateImportTaskResp{RawResponse: rawResp}
	err = rawResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (i *importTask) Get(ctx context.Context, req *GetImportTaskReq, options ...core.RequestOptionFunc) (*GetImportTaskResp, error) {
	// 发起请求
	rawResp, err := core.SendRequest(ctx, i.service.config, http.MethodGet,
		"/open-apis/drive/v1/import_tasks/:ticket", []core.AccessTokenType{core.AccessTokenTypeUser, core.AccessTokenTypeTenant}, req, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &GetImportTaskResp{RawResponse: rawResp}
	err = rawResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (m *media) BatchGetTmpDownloadUrl(ctx context.Context, req *BatchGetTmpDownloadUrlMediaReq, options ...core.RequestOptionFunc) (*BatchGetTmpDownloadUrlMediaResp, error) {
	// 发起请求
	rawResp, err := core.SendRequest(ctx, m.service.config, http.MethodGet,
		"/open-apis/drive/v1/medias/batch_get_tmp_download_url", []core.AccessTokenType{core.AccessTokenTypeUser, core.AccessTokenTypeTenant}, req, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &BatchGetTmpDownloadUrlMediaResp{RawResponse: rawResp}
	err = rawResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (m *media) Download(ctx context.Context, req *DownloadMediaReq, options ...core.RequestOptionFunc) (*DownloadMediaResp, error) {
	// 发起请求
	rawResp, err := core.SendRequest(ctx, m.service.config, http.MethodGet,
		"/open-apis/drive/v1/medias/:file_token/download", []core.AccessTokenType{core.AccessTokenTypeUser, core.AccessTokenTypeTenant}, req, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &DownloadMediaResp{RawResponse: rawResp}
	// 如果是下载，则设置响应结果
	if rawResp.StatusCode == http.StatusOK {
		resp.File = bytes.NewBuffer(rawResp.RawBody)
		resp.FileName = core.FileNameByHeader(rawResp.Header)
		return resp, err
	}
	err = rawResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (m *media) UploadAll(ctx context.Context, req *UploadAllMediaReq, options ...core.RequestOptionFunc) (*UploadAllMediaResp, error) {
	options = append(options, core.WithFileUpload())
	// 发起请求
	rawResp, err := core.SendRequest(ctx, m.service.config, http.MethodPost,
		"/open-apis/drive/v1/medias/upload_all", []core.AccessTokenType{core.AccessTokenTypeUser, core.AccessTokenTypeTenant}, req, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &UploadAllMediaResp{RawResponse: rawResp}
	err = rawResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (m *media) UploadFinish(ctx context.Context, req *UploadFinishMediaReq, options ...core.RequestOptionFunc) (*UploadFinishMediaResp, error) {
	// 发起请求
	rawResp, err := core.SendRequest(ctx, m.service.config, http.MethodPost,
		"/open-apis/drive/v1/medias/upload_finish", []core.AccessTokenType{core.AccessTokenTypeUser, core.AccessTokenTypeTenant}, req, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &UploadFinishMediaResp{RawResponse: rawResp}
	err = rawResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (m *media) UploadPart(ctx context.Context, req *UploadPartMediaReq, options ...core.RequestOptionFunc) (*UploadPartMediaResp, error) {
	options = append(options, core.WithFileUpload())
	// 发起请求
	rawResp, err := core.SendRequest(ctx, m.service.config, http.MethodPost,
		"/open-apis/drive/v1/medias/upload_part", []core.AccessTokenType{core.AccessTokenTypeUser, core.AccessTokenTypeTenant}, req, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &UploadPartMediaResp{RawResponse: rawResp}
	err = rawResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (m *media) UploadPrepare(ctx context.Context, req *UploadPrepareMediaReq, options ...core.RequestOptionFunc) (*UploadPrepareMediaResp, error) {
	// 发起请求
	rawResp, err := core.SendRequest(ctx, m.service.config, http.MethodPost,
		"/open-apis/drive/v1/medias/upload_prepare", []core.AccessTokenType{core.AccessTokenTypeUser, core.AccessTokenTypeTenant}, req, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &UploadPrepareMediaResp{RawResponse: rawResp}
	err = rawResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (m *meta) BatchQuery(ctx context.Context, req *BatchQueryMetaReq, options ...core.RequestOptionFunc) (*BatchQueryMetaResp, error) {
	// 发起请求
	rawResp, err := core.SendRequest(ctx, m.service.config, http.MethodPost,
		"/open-apis/drive/v1/metas/batch_query", []core.AccessTokenType{core.AccessTokenTypeTenant, core.AccessTokenTypeUser}, req, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &BatchQueryMetaResp{RawResponse: rawResp}
	err = rawResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (p *permissionMember) Create(ctx context.Context, req *CreatePermissionMemberReq, options ...core.RequestOptionFunc) (*CreatePermissionMemberResp, error) {
	// 发起请求
	rawResp, err := core.SendRequest(ctx, p.service.config, http.MethodPost,
		"/open-apis/drive/v1/permissions/:token/members", []core.AccessTokenType{core.AccessTokenTypeTenant, core.AccessTokenTypeUser}, req, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &CreatePermissionMemberResp{RawResponse: rawResp}
	err = rawResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (p *permissionMember) Delete(ctx context.Context, req *DeletePermissionMemberReq, options ...core.RequestOptionFunc) (*DeletePermissionMemberResp, error) {
	// 发起请求
	rawResp, err := core.SendRequest(ctx, p.service.config, http.MethodDelete,
		"/open-apis/drive/v1/permissions/:token/members/:member_id", []core.AccessTokenType{core.AccessTokenTypeTenant, core.AccessTokenTypeUser}, req, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &DeletePermissionMemberResp{RawResponse: rawResp}
	err = rawResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (p *permissionMember) Update(ctx context.Context, req *UpdatePermissionMemberReq, options ...core.RequestOptionFunc) (*UpdatePermissionMemberResp, error) {
	// 发起请求
	rawResp, err := core.SendRequest(ctx, p.service.config, http.MethodPut,
		"/open-apis/drive/v1/permissions/:token/members/:member_id", []core.AccessTokenType{core.AccessTokenTypeTenant, core.AccessTokenTypeUser}, req, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &UpdatePermissionMemberResp{RawResponse: rawResp}
	err = rawResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (p *permissionPublic) Get(ctx context.Context, req *GetPermissionPublicReq, options ...core.RequestOptionFunc) (*GetPermissionPublicResp, error) {
	// 发起请求
	rawResp, err := core.SendRequest(ctx, p.service.config, http.MethodGet,
		"/open-apis/drive/v1/permissions/:token/public", []core.AccessTokenType{core.AccessTokenTypeTenant, core.AccessTokenTypeUser}, req, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &GetPermissionPublicResp{RawResponse: rawResp}
	err = rawResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (p *permissionPublic) Patch(ctx context.Context, req *PatchPermissionPublicReq, options ...core.RequestOptionFunc) (*PatchPermissionPublicResp, error) {
	// 发起请求
	rawResp, err := core.SendRequest(ctx, p.service.config, http.MethodPatch,
		"/open-apis/drive/v1/permissions/:token/public", []core.AccessTokenType{core.AccessTokenTypeTenant, core.AccessTokenTypeUser}, req, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &PatchPermissionPublicResp{RawResponse: rawResp}
	err = rawResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
