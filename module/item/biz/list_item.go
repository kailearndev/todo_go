package biz

import (
	"context"
	"todo-api/common"
	"todo-api/module/item/model"
)

type GetListItemStorage interface {
	GetListItem(ctx context.Context,
		filter *model.Filter,
		paging *common.Paging,
		moreKey ...string) ([]model.TodoItem, error)
}

type getListItemBiz struct {
	store GetListItemStorage
}

func NewGetListItemBiz(store GetListItemStorage) *getListItemBiz {
	return &getListItemBiz{store}

}
func (biz *getListItemBiz) GetListItem(ctx context.Context,
	filter *model.Filter,
	paging *common.Paging) ([]model.TodoItem, error) {

	//check title not null

	data, err := biz.store.GetListItem(ctx, filter, paging)

	if err != nil {
		return nil, err
	}

	return data, nil
}
