package biz

import (
	"context"
	"todo-api/module/item/model"
)

type UpdateItemStorage interface {
	GetItem(ctx context.Context, cond map[string]interface{}) (*model.TodoItem, error)
	UpdateItem(ctx context.Context, cond map[string]interface{}, dataUpdate *model.TodoItemUpdate) error
}

type updateItemBiz struct {
	store UpdateItemStorage
}

func NewUpdateItemBiz(store UpdateItemStorage) *updateItemBiz {
	return &updateItemBiz{store}

}
func (biz *updateItemBiz) UpdateItemId(ctx context.Context, id int, dataUpdate *model.TodoItemUpdate) error {

	//check title not null

	data, err := biz.store.GetItem(ctx, map[string]interface{}{"id": id})

	if err != nil {
		return err
	}
	if data.Status != nil && *data.Status == model.ItemsStatusDeleted {
		return model.ErrTitleIsBlank
	}
	if err := biz.store.UpdateItem(ctx, map[string]interface{}{"id": id}, dataUpdate); err != nil {
		return err
	}
	// check if item delete => return err

	return nil
}
