package biz

import (
	"context"
	"todo-api/module/item/model"
)

type DeleteItemStorage interface {
	GetItem(ctx context.Context, cond map[string]interface{}) (*model.TodoItem, error)
	DeleteItem(ctx context.Context, cond map[string]interface{}) error
}

type deleteItemBiz struct {
	store DeleteItemStorage
}

func NewDeleteItemBiz(store DeleteItemStorage) *deleteItemBiz {
	return &deleteItemBiz{store}

}
func (biz *deleteItemBiz) DeleteItemId(ctx context.Context, id int) error {

	//check title not null

	data, err := biz.store.GetItem(ctx, map[string]interface{}{"id": id})

	if err != nil {
		return err
	}
	if data.Status != nil && *data.Status == model.ItemsStatusDeleted {
		return model.ErrTitleIsBlank
	}
	if err := biz.store.DeleteItem(ctx, map[string]interface{}{"id": id}); err != nil {
		return err
	}
	// check if item delete => return err

	return nil
}
