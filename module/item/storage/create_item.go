package storage

import (
	"context"
	"todo-api/module/item/model"
)

func (s *sqlStore) CreateItem(ctx context.Context, data *model.TodoItemCreate) error {
	if err := s.db.Create(&data).Error; err != nil {

		return err
	}

	return nil
}
