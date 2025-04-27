package ginitem

import (
	"net/http"
	"todo-api/common"
	"todo-api/module/item/biz"
	"todo-api/module/item/model"
	"todo-api/module/item/storage"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CreateItem(db *gorm.DB) func(c *gin.Context) {

	return func(c *gin.Context) {
		var data model.TodoItemCreate

		//
		if err := c.ShouldBind(&data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		store := storage.NewSQLStore(db)
		bizz := biz.NewCreateItemBiz(store)

		if err := bizz.CreateItem(c.Request.Context(), &data); err != nil {
			if err := c.ShouldBind(&data); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"error": err.Error(),
				})
				return
			}
		}
		c.JSON(http.StatusOK, common.SimpleSuccessResponse(data.Id))
	}

}
