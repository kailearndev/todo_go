package storage

import (
	"context"
	"todo-api/module/item/model"
)

func (s *sqlStore) DeleteItem(ctx context.Context, cond map[string]interface{}) error {

	if err := s.db.Table(model.TodoItem{}.TableName()).Where(cond).Updates(map[string]interface{}{
		"status": "deleted",
	}).Error; err != nil {

		return err

	}
	return nil
}
