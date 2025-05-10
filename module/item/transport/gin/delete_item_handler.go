package ginitem

import (
	"net/http"
	"strconv"
	"todo-api/common"
	"todo-api/module/item/biz"
	"todo-api/module/item/storage"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func DeleteItem(db *gorm.DB) func(c *gin.Context) {
	return func(c *gin.Context) {

		//get param default retturn string need parse int
		id, err := strconv.Atoi(c.Param(("id")))

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		// data.Id = id // method 1
		store := storage.NewSQLStore(db)
		business := biz.NewDeleteItemBiz(store)
		if err := business.DeleteItemId(c.Request.Context(), id); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		}
		c.JSON(http.StatusOK, common.SimpleSuccessResponse(true))

	}

}
