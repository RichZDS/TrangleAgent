package ForumComments

import (
	"context"
	"errors"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"leke/api/forum/v1"
	"leke/internal/dao"
	"leke/internal/model/entity"
	"leke/internal/service"
)

// ForumComments  评论相关方法
type sForumComments struct{}

func init() {
	service.RegisterForumComments(&sForumComments{})
}

// ForumCommentsCreate 创建评论
func (s *sForumComments) Create(ctx context.Context, req *v1.ForumCommentsCreateReq) (res *v1.ForumCommentsCreateRes, err error) {
	data := &entity.ForumComments{
		UserId:     req.UserId,
		PostId:     req.PostId,
		ParentId:   req.ParentId,
		Content:    req.Content,
		Status:     req.Status,
		LikeCount:  0,
		ReplyCount: 0,
	}

	result, err := dao.ForumComments.Ctx(ctx).Data(data).OmitEmpty().Insert()
	if err != nil {
		return nil, gerror.Wrap(err, "创建评论失败")
	}

	id, err := result.LastInsertId()
	if err != nil {
		return nil, gerror.Wrap(err, "获取评论ID失败")
	}

	// 更新帖子的评论数
	if req.PostId > 0 {
		_, err = dao.ForumPosts.Ctx(ctx).WherePri(req.PostId).OnDuplicate("comment_count", 1).Update()
		if err != nil {
			g.Log().Warning(ctx, "更新帖子评论数失败:", err)
		}
	}

	// 如果是回复评论，则更新父评论的回复数
	if req.ParentId > 0 {
		_, err = dao.ForumComments.Ctx(ctx).WherePri(req.ParentId).OnDuplicate("reply_count", 1).Update()
		if err != nil {
			g.Log().Warning(ctx, "更新父评论回复数失败:", err)
		}
	}

	return &v1.ForumCommentsCreateRes{
		Id: uint64(id),
	}, nil
}

// ForumCommentsUpdate 更新评论
func (s *sForumComments) Update(ctx context.Context, req *v1.ForumCommentsUpdateReq) (res *v1.ForumCommentsUpdateRes, err error) {
	// 检查评论是否存在
	one, err := dao.ForumComments.Ctx(ctx).WherePri(req.Id).One()
	if err != nil {
		return nil, gerror.Wrap(err, "查询评论失败")
	}
	if one.IsEmpty() {
		return nil, errors.New("评论不存在")
	}

	// 更新评论
	_, err = dao.ForumComments.Ctx(ctx).WherePri(req.Id).Data(req).OmitEmpty().Update()
	if err != nil {
		return nil, gerror.Wrap(err, "更新评论失败")
	}

	return &v1.ForumCommentsUpdateRes{
		Id: req.Id,
	}, nil
}

// ForumCommentsView 查看评论
func (s *sForumComments) View(ctx context.Context, req *v1.ForumCommentsViewReq) (res *v1.ForumCommentsViewRes, err error) {
	record, err := dao.ForumComments.Ctx(ctx).WherePri(req.Id).One()
	if err != nil {
		return nil, gerror.Wrap(err, "查询评论失败")
	}
	if record.IsEmpty() {
		return nil, errors.New("评论不存在")
	}

	var entity *entity.ForumComments
	err = record.Struct(&entity)
	if err != nil {
		return nil, gerror.Wrap(err, "解析评论数据失败")
	}

	res = &v1.ForumCommentsViewRes{
		Id:         entity.Id,
		UserId:     entity.UserId,
		PostId:     entity.PostId,
		ParentId:   entity.ParentId,
		Content:    entity.Content,
		Status:     entity.Status,
		LikeCount:  entity.LikeCount,
		ReplyCount: entity.ReplyCount,
		CreatedAt:  entity.CreatedAt,
		UpdatedAt:  entity.UpdatedAt,
	}

	return res, nil
}

// ForumCommentsList 评论列表
func (s *sForumComments) List(ctx context.Context, req *v1.ForumCommentsListReq) (res *v1.ForumCommentsListRes, err error) {
	query := dao.ForumComments.Ctx(ctx).OmitEmpty()

	if req.UserId > 0 {
		query = query.Where("user_id", req.UserId)
	}
	if req.PostId > 0 {
		query = query.Where("post_id", req.PostId)
	}
	if req.ParentId > 0 {
		query = query.Where("parent_id", req.ParentId)
	}
	if req.Status != "" {
		query = query.Where("status", req.Status)
	}

	// 分页查询
	pageResult, err := query.Page(req.Page, req.PageSize).Order("created_at DESC").All()
	if err != nil {
		return nil, gerror.Wrap(err, "查询评论列表失败")
	}

	var list []*v1.ForumCommentsViewRes
	for _, record := range pageResult {
		var entity *entity.ForumComments
		err = record.Struct(&entity)
		if err != nil {
			return nil, gerror.Wrap(err, "解析评论数据失败")
		}

		item := &v1.ForumCommentsViewRes{
			Id:         entity.Id,
			UserId:     entity.UserId,
			PostId:     entity.PostId,
			ParentId:   entity.ParentId,
			Content:    entity.Content,
			Status:     entity.Status,
			LikeCount:  entity.LikeCount,
			ReplyCount: entity.ReplyCount,
			CreatedAt:  entity.CreatedAt,
			UpdatedAt:  entity.UpdatedAt,
		}
		list = append(list, item)
	}

	total, err := query.Count()
	if err != nil {
		return nil, gerror.Wrap(err, "统计评论总数失败")
	}

	res = &v1.ForumCommentsListRes{}
	res.List = list
	res.Page = req.Page
	res.PageSize = req.PageSize
	res.Total = total

	return res, nil
}

// ForumCommentsDelete 删除评论
func (s *sForumComments) Delete(ctx context.Context, req *v1.ForumCommentsDeleteReq) (res *v1.ForumCommentsDeleteRes, err error) {
	// 检查评论是否存在
	one, err := dao.ForumComments.Ctx(ctx).WherePri(req.Id).One()
	if err != nil {
		return nil, gerror.Wrap(err, "查询评论失败")
	}
	if one.IsEmpty() {
		return nil, errors.New("评论不存在")
	}

	var comment *entity.ForumComments
	err = one.Struct(&comment)
	if err != nil {
		return nil, gerror.Wrap(err, "解析评论数据失败")
	}

	// 软删除评论
	_, err = dao.ForumComments.Ctx(ctx).WherePri(req.Id).Delete()
	if err != nil {
		return nil, gerror.Wrap(err, "删除评论失败")
	}

	// 更新帖子的评论数
	if comment.PostId > 0 {
		_, err = dao.ForumPosts.Ctx(ctx).WherePri(comment.PostId).OnDuplicate("comment_count", 1).Update()
		if err != nil {
			g.Log().Warning(ctx, "更新帖子评论数失败:", err)
		}
	}

	// 如果是回复评论，则更新父评论的回复数
	if comment.ParentId > 0 {
		_, err = dao.ForumComments.Ctx(ctx).WherePri(comment.ParentId).OnDuplicate("reply_count", 1).Update()
		if err != nil {
			g.Log().Warning(ctx, "更新父评论回复数失败:", err)
		}
	}

	return &v1.ForumCommentsDeleteRes{}, nil
}
