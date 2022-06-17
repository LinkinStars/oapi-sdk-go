// Package task code generated by oapi sdk gen
package task

import (
	"context"
	"net/http"

	"github.com/larksuite/oapi-sdk-go/core"
)

/**
构建业务域服务实例
**/
func NewService(httpClient *http.Client, config *core.Config) *TaskService {
	t := &TaskService{httpClient: httpClient, config: config}
	t.Task = &task{service: t}
	t.TaskCollaborator = &taskCollaborator{service: t}
	t.TaskComment = &taskComment{service: t}
	t.TaskFollower = &taskFollower{service: t}
	t.TaskReminder = &taskReminder{service: t}
	return t
}

/**
业务域服务定义
**/
type TaskService struct {
	httpClient       *http.Client
	config           *core.Config
	Task             *task
	TaskCollaborator *taskCollaborator
	TaskComment      *taskComment
	TaskFollower     *taskFollower
	TaskReminder     *taskReminder
}

/**
资源服务定义
**/
type task struct {
	service *TaskService
}
type taskCollaborator struct {
	service *TaskService
}
type taskComment struct {
	service *TaskService
}
type taskFollower struct {
	service *TaskService
}
type taskReminder struct {
	service *TaskService
}

/**
资源服务方法定义
**/
func (t *task) Complete(ctx context.Context, req *CompleteTaskReq, options ...core.RequestOptionFunc) (*CompleteTaskResp, error) {
	// 发起请求
	rawResp, err := core.SendRequest(ctx, t.service.config, http.MethodPost,
		"/open-apis/task/v1/tasks/:task_id/complete", []core.AccessTokenType{core.AccessTokenTypeTenant, core.AccessTokenTypeUser}, req, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &CompleteTaskResp{RawResponse: rawResp}
	err = rawResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (t *task) Create(ctx context.Context, req *CreateTaskReq, options ...core.RequestOptionFunc) (*CreateTaskResp, error) {
	// 发起请求
	rawResp, err := core.SendRequest(ctx, t.service.config, http.MethodPost,
		"/open-apis/task/v1/tasks", []core.AccessTokenType{core.AccessTokenTypeTenant, core.AccessTokenTypeUser}, req, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &CreateTaskResp{RawResponse: rawResp}
	err = rawResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (t *task) Delete(ctx context.Context, req *DeleteTaskReq, options ...core.RequestOptionFunc) (*DeleteTaskResp, error) {
	// 发起请求
	rawResp, err := core.SendRequest(ctx, t.service.config, http.MethodDelete,
		"/open-apis/task/v1/tasks/:task_id", []core.AccessTokenType{core.AccessTokenTypeTenant, core.AccessTokenTypeUser}, req, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &DeleteTaskResp{RawResponse: rawResp}
	err = rawResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (t *task) Get(ctx context.Context, req *GetTaskReq, options ...core.RequestOptionFunc) (*GetTaskResp, error) {
	// 发起请求
	rawResp, err := core.SendRequest(ctx, t.service.config, http.MethodGet,
		"/open-apis/task/v1/tasks/:task_id", []core.AccessTokenType{core.AccessTokenTypeTenant, core.AccessTokenTypeUser}, req, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &GetTaskResp{RawResponse: rawResp}
	err = rawResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (t *task) List(ctx context.Context, req *ListTaskReq, options ...core.RequestOptionFunc) (*ListTaskResp, error) {
	// 发起请求
	rawResp, err := core.SendRequest(ctx, t.service.config, http.MethodGet,
		"/open-apis/task/v1/tasks", []core.AccessTokenType{core.AccessTokenTypeUser, core.AccessTokenTypeTenant}, req, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &ListTaskResp{RawResponse: rawResp}
	err = rawResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}

/**如果是分页查询，则添加迭代器函数**/
func (t *task) ListTask(ctx context.Context, req *ListTaskReq, options ...core.RequestOptionFunc) (*ListTaskIterator, error) {
	return &ListTaskIterator{
		ctx:      ctx,
		req:      req,
		listFunc: t.List,
		options:  options,
		limit:    req.Limit}, nil
}
func (t *task) Patch(ctx context.Context, req *PatchTaskReq, options ...core.RequestOptionFunc) (*PatchTaskResp, error) {
	// 发起请求
	rawResp, err := core.SendRequest(ctx, t.service.config, http.MethodPatch,
		"/open-apis/task/v1/tasks/:task_id", []core.AccessTokenType{core.AccessTokenTypeTenant, core.AccessTokenTypeUser}, req, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &PatchTaskResp{RawResponse: rawResp}
	err = rawResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (t *task) Uncomplete(ctx context.Context, req *UncompleteTaskReq, options ...core.RequestOptionFunc) (*UncompleteTaskResp, error) {
	// 发起请求
	rawResp, err := core.SendRequest(ctx, t.service.config, http.MethodPost,
		"/open-apis/task/v1/tasks/:task_id/uncomplete", []core.AccessTokenType{core.AccessTokenTypeTenant, core.AccessTokenTypeUser}, req, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &UncompleteTaskResp{RawResponse: rawResp}
	err = rawResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (t *taskCollaborator) Create(ctx context.Context, req *CreateTaskCollaboratorReq, options ...core.RequestOptionFunc) (*CreateTaskCollaboratorResp, error) {
	// 发起请求
	rawResp, err := core.SendRequest(ctx, t.service.config, http.MethodPost,
		"/open-apis/task/v1/tasks/:task_id/collaborators", []core.AccessTokenType{core.AccessTokenTypeTenant, core.AccessTokenTypeUser}, req, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &CreateTaskCollaboratorResp{RawResponse: rawResp}
	err = rawResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (t *taskCollaborator) Delete(ctx context.Context, req *DeleteTaskCollaboratorReq, options ...core.RequestOptionFunc) (*DeleteTaskCollaboratorResp, error) {
	// 发起请求
	rawResp, err := core.SendRequest(ctx, t.service.config, http.MethodDelete,
		"/open-apis/task/v1/tasks/:task_id/collaborators/:collaborator_id", []core.AccessTokenType{core.AccessTokenTypeTenant, core.AccessTokenTypeUser}, req, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &DeleteTaskCollaboratorResp{RawResponse: rawResp}
	err = rawResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (t *taskCollaborator) List(ctx context.Context, req *ListTaskCollaboratorReq, options ...core.RequestOptionFunc) (*ListTaskCollaboratorResp, error) {
	// 发起请求
	rawResp, err := core.SendRequest(ctx, t.service.config, http.MethodGet,
		"/open-apis/task/v1/tasks/:task_id/collaborators", []core.AccessTokenType{core.AccessTokenTypeTenant, core.AccessTokenTypeUser}, req, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &ListTaskCollaboratorResp{RawResponse: rawResp}
	err = rawResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}

/**如果是分页查询，则添加迭代器函数**/
func (t *taskCollaborator) ListTaskCollaborator(ctx context.Context, req *ListTaskCollaboratorReq, options ...core.RequestOptionFunc) (*ListTaskCollaboratorIterator, error) {
	return &ListTaskCollaboratorIterator{
		ctx:      ctx,
		req:      req,
		listFunc: t.List,
		options:  options,
		limit:    req.Limit}, nil
}
func (t *taskComment) Create(ctx context.Context, req *CreateTaskCommentReq, options ...core.RequestOptionFunc) (*CreateTaskCommentResp, error) {
	// 发起请求
	rawResp, err := core.SendRequest(ctx, t.service.config, http.MethodPost,
		"/open-apis/task/v1/tasks/:task_id/comments", []core.AccessTokenType{core.AccessTokenTypeTenant, core.AccessTokenTypeUser}, req, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &CreateTaskCommentResp{RawResponse: rawResp}
	err = rawResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (t *taskComment) Delete(ctx context.Context, req *DeleteTaskCommentReq, options ...core.RequestOptionFunc) (*DeleteTaskCommentResp, error) {
	// 发起请求
	rawResp, err := core.SendRequest(ctx, t.service.config, http.MethodDelete,
		"/open-apis/task/v1/tasks/:task_id/comments/:comment_id", []core.AccessTokenType{core.AccessTokenTypeTenant, core.AccessTokenTypeUser}, req, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &DeleteTaskCommentResp{RawResponse: rawResp}
	err = rawResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (t *taskComment) Get(ctx context.Context, req *GetTaskCommentReq, options ...core.RequestOptionFunc) (*GetTaskCommentResp, error) {
	// 发起请求
	rawResp, err := core.SendRequest(ctx, t.service.config, http.MethodGet,
		"/open-apis/task/v1/tasks/:task_id/comments/:comment_id", []core.AccessTokenType{core.AccessTokenTypeTenant, core.AccessTokenTypeUser}, req, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &GetTaskCommentResp{RawResponse: rawResp}
	err = rawResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (t *taskComment) List(ctx context.Context, req *ListTaskCommentReq, options ...core.RequestOptionFunc) (*ListTaskCommentResp, error) {
	// 发起请求
	rawResp, err := core.SendRequest(ctx, t.service.config, http.MethodGet,
		"/open-apis/task/v1/tasks/:task_id/comments", []core.AccessTokenType{core.AccessTokenTypeTenant, core.AccessTokenTypeUser}, req, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &ListTaskCommentResp{RawResponse: rawResp}
	err = rawResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}

/**如果是分页查询，则添加迭代器函数**/
func (t *taskComment) ListTaskComment(ctx context.Context, req *ListTaskCommentReq, options ...core.RequestOptionFunc) (*ListTaskCommentIterator, error) {
	return &ListTaskCommentIterator{
		ctx:      ctx,
		req:      req,
		listFunc: t.List,
		options:  options,
		limit:    req.Limit}, nil
}
func (t *taskComment) Update(ctx context.Context, req *UpdateTaskCommentReq, options ...core.RequestOptionFunc) (*UpdateTaskCommentResp, error) {
	// 发起请求
	rawResp, err := core.SendRequest(ctx, t.service.config, http.MethodPut,
		"/open-apis/task/v1/tasks/:task_id/comments/:comment_id", []core.AccessTokenType{core.AccessTokenTypeTenant, core.AccessTokenTypeUser}, req, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &UpdateTaskCommentResp{RawResponse: rawResp}
	err = rawResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (t *taskFollower) Create(ctx context.Context, req *CreateTaskFollowerReq, options ...core.RequestOptionFunc) (*CreateTaskFollowerResp, error) {
	// 发起请求
	rawResp, err := core.SendRequest(ctx, t.service.config, http.MethodPost,
		"/open-apis/task/v1/tasks/:task_id/followers", []core.AccessTokenType{core.AccessTokenTypeTenant, core.AccessTokenTypeUser}, req, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &CreateTaskFollowerResp{RawResponse: rawResp}
	err = rawResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (t *taskFollower) Delete(ctx context.Context, req *DeleteTaskFollowerReq, options ...core.RequestOptionFunc) (*DeleteTaskFollowerResp, error) {
	// 发起请求
	rawResp, err := core.SendRequest(ctx, t.service.config, http.MethodDelete,
		"/open-apis/task/v1/tasks/:task_id/followers/:follower_id", []core.AccessTokenType{core.AccessTokenTypeTenant, core.AccessTokenTypeUser}, req, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &DeleteTaskFollowerResp{RawResponse: rawResp}
	err = rawResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (t *taskFollower) List(ctx context.Context, req *ListTaskFollowerReq, options ...core.RequestOptionFunc) (*ListTaskFollowerResp, error) {
	// 发起请求
	rawResp, err := core.SendRequest(ctx, t.service.config, http.MethodGet,
		"/open-apis/task/v1/tasks/:task_id/followers", []core.AccessTokenType{core.AccessTokenTypeTenant, core.AccessTokenTypeUser}, req, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &ListTaskFollowerResp{RawResponse: rawResp}
	err = rawResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}

/**如果是分页查询，则添加迭代器函数**/
func (t *taskFollower) ListTaskFollower(ctx context.Context, req *ListTaskFollowerReq, options ...core.RequestOptionFunc) (*ListTaskFollowerIterator, error) {
	return &ListTaskFollowerIterator{
		ctx:      ctx,
		req:      req,
		listFunc: t.List,
		options:  options,
		limit:    req.Limit}, nil
}
func (t *taskReminder) Create(ctx context.Context, req *CreateTaskReminderReq, options ...core.RequestOptionFunc) (*CreateTaskReminderResp, error) {
	// 发起请求
	rawResp, err := core.SendRequest(ctx, t.service.config, http.MethodPost,
		"/open-apis/task/v1/tasks/:task_id/reminders", []core.AccessTokenType{core.AccessTokenTypeTenant, core.AccessTokenTypeUser}, req, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &CreateTaskReminderResp{RawResponse: rawResp}
	err = rawResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (t *taskReminder) Delete(ctx context.Context, req *DeleteTaskReminderReq, options ...core.RequestOptionFunc) (*DeleteTaskReminderResp, error) {
	// 发起请求
	rawResp, err := core.SendRequest(ctx, t.service.config, http.MethodDelete,
		"/open-apis/task/v1/tasks/:task_id/reminders/:reminder_id", []core.AccessTokenType{core.AccessTokenTypeTenant, core.AccessTokenTypeUser}, req, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &DeleteTaskReminderResp{RawResponse: rawResp}
	err = rawResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (t *taskReminder) List(ctx context.Context, req *ListTaskReminderReq, options ...core.RequestOptionFunc) (*ListTaskReminderResp, error) {
	// 发起请求
	rawResp, err := core.SendRequest(ctx, t.service.config, http.MethodGet,
		"/open-apis/task/v1/tasks/:task_id/reminders", []core.AccessTokenType{core.AccessTokenTypeTenant, core.AccessTokenTypeUser}, req, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &ListTaskReminderResp{RawResponse: rawResp}
	err = rawResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}

/**如果是分页查询，则添加迭代器函数**/
func (t *taskReminder) ListTaskReminder(ctx context.Context, req *ListTaskReminderReq, options ...core.RequestOptionFunc) (*ListTaskReminderIterator, error) {
	return &ListTaskReminderIterator{
		ctx:      ctx,
		req:      req,
		listFunc: t.List,
		options:  options,
		limit:    req.Limit}, nil
}
