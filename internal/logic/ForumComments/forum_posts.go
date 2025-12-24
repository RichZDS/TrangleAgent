package ForumComments

import (
	"context"
	"errors"
	"leke/api/forum/v1"
	"leke/internal/dao"
	"leke/internal/model/entity"
	"leke/internal/service"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
)

type sForumPosts struct{}

func init() {
	service.RegisterForumPosts(&sForumPosts{})
}

// ForumPostsCreate 创建帖子
func (s *sForumPosts) Create(ctx context.Context, req *v1.ForumPostsCreateReq) (res *v1.ForumPostsCreateRes, err error) {
	data := &entity.ForumPosts{
		UserId:       req.UserId,
		Title:        req.Title,
		Content:      req.Content,
		CoverImage:   req.CoverImage,
		Status:       req.Status,
		ViewCount:    0,
		LikeCount:    0,
		CommentCount: 0,
	}

	result, err := dao.ForumPosts.Ctx(ctx).Data(data).OmitEmpty().Insert()
	if err != nil {
		return nil, gerror.Wrap(err, "创建帖子失败")
	}

	id, err := result.LastInsertId()
	if err != nil {
		return nil, gerror.Wrap(err, "获取帖子ID失败")
	}

	return &v1.ForumPostsCreateRes{
		Id: uint64(id),
	}, nil
}

// ForumPostsUpdate 更新帖子
func (s *sForumPosts) Update(ctx context.Context, req *v1.ForumPostsUpdateReq) (res *v1.ForumPostsUpdateRes, err error) {
	// 检查帖子是否存在
	one, err := dao.ForumPosts.Ctx(ctx).WherePri(req.Id).One()
	if err != nil {
		return nil, gerror.Wrap(err, "查询帖子失败")
	}
	if one.IsEmpty() {
		return nil, errors.New("帖子不存在")
	}

	// 更新帖子
	_, err = dao.ForumPosts.Ctx(ctx).WherePri(req.Id).Data(req).OmitEmpty().Update()
	if err != nil {
		return nil, gerror.Wrap(err, "更新帖子失败")
	}

	return &v1.ForumPostsUpdateRes{
		Id: req.Id,
	}, nil
}

// ForumPostsView 查看帖子
func (s *sForumPosts) View(ctx context.Context, req *v1.ForumPostsViewReq) (res *v1.ForumPostsViewRes, err error) {
	record, err := dao.ForumPosts.Ctx(ctx).WherePri(req.Id).One()
	if err != nil {
		return nil, gerror.Wrap(err, "查询帖子失败")
	}
	if record.IsEmpty() {
		return nil, errors.New("帖子不存在")
	}

	// 更新浏览量
	_, err = dao.ForumPosts.Ctx(ctx).WherePri(req.Id).OnDuplicate("view_count", "view_count+1").Update()
	if err != nil {
		g.Log().Warning(ctx, "更新浏览量失败:", err)
	}

	var entity *entity.ForumPosts
	err = record.Struct(&entity)
	if err != nil {
		return nil, gerror.Wrap(err, "解析帖子数据失败")
	}

	res = &v1.ForumPostsViewRes{
		Id:           entity.Id,
		UserId:       entity.UserId,
		Title:        entity.Title,
		Content:      entity.Content,
		CoverImage:   entity.CoverImage,
		Status:       entity.Status,
		ViewCount:    entity.ViewCount,
		LikeCount:    entity.LikeCount,
		CommentCount: entity.CommentCount,
		CreatedAt:    entity.CreatedAt,
		UpdatedAt:    entity.UpdatedAt,
	}

	return res, nil
}

// ForumPostsList 帖子列表
func (s *sForumPosts) List(ctx context.Context, req *v1.ForumPostsListReq) (res *v1.ForumPostsListRes, err error) {
	query := dao.ForumPosts.Ctx(ctx).OmitEmpty()

	if req.UserId > 0 {
		query = query.Where("user_id", req.UserId)
	}
	if req.Title != "" {
		query = query.WhereLike("title", "%"+req.Title+"%")
	}
	if req.Status != "" {
		query = query.Where("status", req.Status)
	}

	// 分页查询
	pageResult, err := query.Page(req.Page, req.PageSize).Order("created_at DESC").All()
	if err != nil {
		return nil, gerror.Wrap(err, "查询帖子列表失败")
	}

	var list []*v1.ForumPostsViewRes
	for _, record := range pageResult {
		var entity *entity.ForumPosts
		err = record.Struct(&entity)
		if err != nil {
			return nil, gerror.Wrap(err, "解析帖子数据失败")
		}

		item := &v1.ForumPostsViewRes{
			Id:           entity.Id,
			UserId:       entity.UserId,
			Title:        entity.Title,
			Content:      entity.Content,
			CoverImage:   entity.CoverImage,
			Status:       entity.Status,
			ViewCount:    entity.ViewCount,
			LikeCount:    entity.LikeCount,
			CommentCount: entity.CommentCount,
			CreatedAt:    entity.CreatedAt,
			UpdatedAt:    entity.UpdatedAt,
		}
		list = append(list, item)
	}

	total, err := query.Count()
	if err != nil {
		return nil, gerror.Wrap(err, "统计帖子总数失败")
	}

	res = &v1.ForumPostsListRes{}
	res.List = list
	res.PageResult.Page = req.Page
	res.PageResult.PageSize = req.PageSize
	res.PageResult.Total = total

	return res, nil
}

// ForumPostsDelete 删除帖子
func (s *sForumPosts) Delete(ctx context.Context, req *v1.ForumPostsDeleteReq) (res *v1.ForumPostsDeleteRes, err error) {
	// 检查帖子是否存在
	one, err := dao.ForumPosts.Ctx(ctx).WherePri(req.Id).One()
	if err != nil {
		return nil, gerror.Wrap(err, "查询帖子失败")
	}
	if one.IsEmpty() {
		return nil, errors.New("帖子不存在")
	}

	// 软删除帖子
	_, err = dao.ForumPosts.Ctx(ctx).WherePri(req.Id).Delete()
	if err != nil {
		return nil, gerror.Wrap(err, "删除帖子失败")
	}

	return &v1.ForumPostsDeleteRes{}, nil
}
