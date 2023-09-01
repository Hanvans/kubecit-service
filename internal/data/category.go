package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"kubecit-service/internal/biz"
)

type categoryRepo struct {
	data *Data
	log  *log.Helper
}

// NewUserRepo 用户数据仓库构造方法
func NewCategoryRepo(data *Data, logger log.Logger) biz.CategoryRepo {
	return &categoryRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (c *categoryRepo) ListAll(ctx context.Context) ([]*biz.Category, error) {
	categories, err := c.data.db.Category.Query().All(ctx)
	if err != nil {
		return nil, err
	}

	var categoryResult []*biz.Category

	for _, v := range categories {
		categoryResult = append(categoryResult, &biz.Category{
			CategoryName: v.Name,
			Id:           v.Name,
		})
	}
	return categoryResult, nil
}
