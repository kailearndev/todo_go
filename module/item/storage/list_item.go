package storage

import (
	"context"
	"todo-api/common"
	"todo-api/module/item/model"
)

func (s *sqlStore) GetListItem(ctx context.Context,
	filter *model.Filter,
	paging *common.Paging,
	moreKey ...string) ([]model.TodoItem, error) {

	var result []model.TodoItem

	db := s.db.Where("status <> ?", "deleted")

	if f := filter; f != nil {
		if v := f.Status; v != "" {
			db = db.Where("status = ?", v)
		}
	}

	if err := db.Table(model.TodoItem{}.TableName()).Count(&paging.Total).Error; err != nil {

		return nil, err
	}

	if err := db.Order("id desc").
		Offset((paging.Page - 1) * paging.Limit).
		Limit(paging.Limit).
		Find(&result).Error; err != nil {

		return nil, err

	}
	return result, nil

}
