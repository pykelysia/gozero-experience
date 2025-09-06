package svc

import (
	"gozero-exp/rpc/add/internal/config"
	"gozero-exp/rpc/model"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config config.Config
	Model model.BookModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		Model: model.NewBookModel(sqlx.NewMysql(c.DataSource), c.Cache),
	}
}
