package biz

import (
	"context"
	"strings"
	"todo-api/module/item/model"
)

type CreateItemStorage interface {
	CreateItem(ctx context.Context, data *model.TodoItemCreate) error
}

type createItemBiz struct {
	store CreateItemStorage
}

func NewCreateItemBiz(storage CreateItemStorage) *createItemBiz {
	return &createItemBiz{store: storage}

}
func (biz *createItemBiz) CreateItem(ctx context.Context, data *model.TodoItemCreate) error {

	//check title not null

	title := strings.TrimSpace(data.Title)
	if title == "" {
		return model.ErrTitleIsBlank //return erorrs
	}
	if err := biz.store.CreateItem(ctx, data); err != nil {
		return err
	}
	return nil
}
