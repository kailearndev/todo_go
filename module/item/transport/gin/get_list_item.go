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

func GetListItem(db *gorm.DB) func(c *gin.Context) {
	return func(c *gin.Context) {

		//poaging
		var paging common.Paging

		if err := c.ShouldBind(&paging); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		paging.Process()
		var filter model.Filter
		if err := c.ShouldBind(&filter); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		store := storage.NewSQLStore(db)
		bussiness := biz.GetListItemStorage(store)
		result, err := bussiness.GetListItem(c.Request.Context(), &filter, &paging)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		}

		c.JSON(http.StatusOK, common.NewSuccessResponse(result, paging, filter))
	}

}
