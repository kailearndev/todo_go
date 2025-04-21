package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type TodoItem struct {

	// Image
	// json javascript object notasion
	Id          int        `json:"id"`
	Title       string     `json:"title"`
	Description string     `json:"description"`
	Status      string     `json:"status"`
	CreatedAt   *time.Time `json:"created_at"`
	UpdatedAt   *time.Time `json:"updated_at"`
}

func (TodoItem) TableName() string {
	return "todo_items"
}

type TodoItemCreate struct {

	// Image
	// json javascript object notasion
	Id          int    `json:"-" gorm:"column:id;"`
	Title       string `json:"title" gorm:"column:title;"`
	Description string `json:"description" gorm:"column:description;"`
}

func (TodoItemCreate) TableName() string {
	return TodoItem{}.TableName()
}

// update
type TodoItemUpdate struct {

	// Image
	// json javascript object notasion

	Title       *string `json:"title" gorm:"column:title;"`
	Description *string `json:"description" gorm:"column:description;"`
	Status      *string `json:"status"`
}

func (TodoItemUpdate) TableName() string {
	return TodoItem{}.TableName()
}

type Paging struct {
	Page  int   `json:"page" form:"page"`
	Limit int   `json:"limit"  form:"limit"`
	Total int64 `json:"total"  form:"-"`
}

func (p *Paging) Process() {
	if p.Page <= 0 {
		p.Page = 1
	}
	if p.Limit <= 0 || p.Limit >= 100 {
		p.Limit = 10
	}

}

// /
func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("❌ Error loading .env file")
	}

	port := os.Getenv("PORT")
	dsn := os.Getenv("DB_URL")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("connected_ok", db)

	r := gin.Default()

	v1 := r.Group("/v1")

	{
		items := v1.Group("/items")
		{
			items.POST("", CreateItem(db))
			items.GET("", GetListItem(db))
			items.GET("/:id", GetItem(db))
			items.PATCH("/:id", UpdateItem(db))
			items.DELETE("/:id", DeleteItem(db))
		}

	}

	r.Run(":" + port)
}

// function
func CreateItem(db *gorm.DB) func(c *gin.Context) {

	return func(c *gin.Context) {
		var data TodoItemCreate

		//
		if err := c.ShouldBind(&data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		if err := db.Create(&data).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"data": data.Id,
		})
	}

}

func GetItem(db *gorm.DB) func(c *gin.Context) {
	return func(c *gin.Context) {
		var data TodoItem

		//get param default retturn string need parse int
		id, err := strconv.Atoi(c.Param(("id")))

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		// data.Id = id // method 1
		if err := db.Where("id = ?", id).First(&data).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return

		}
		c.JSON(http.StatusOK, gin.H{
			"data": data,
		})
	}

}
func UpdateItem(db *gorm.DB) func(c *gin.Context) {
	return func(c *gin.Context) {
		var data TodoItemUpdate

		//get param default retturn string need parse int
		id, err := strconv.Atoi(c.Param(("id")))

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		if err := c.ShouldBind(&data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		// data.Id = id // method 1
		if err := db.Where("id = ?", id).Updates(&data).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return

		}
		c.JSON(http.StatusOK, gin.H{
			"data": true,
		})
	}

}
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
		if err := db.Table(TodoItem{}.TableName()).Where("id = ?", id).Updates(map[string]interface{}{
			"status": "deleted",
		}).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return

		}
		c.JSON(http.StatusOK, gin.H{
			"data": true,
		})
	}

}
func GetListItem(db *gorm.DB) func(c *gin.Context) {
	return func(c *gin.Context) {

		//poaging
		var paging Paging

		if err := c.ShouldBind(&paging); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		paging.Process()

		//
		var result []TodoItem
		db = db.Where("status <> ?", "deleted")
		if err := db.Table(TodoItem{}.TableName()).
			Count(&paging.Total).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		if err := db.Order("id desc").
			Offset((paging.Page - 1) * paging.Limit).
			Limit(paging.Limit).
			Find(&result).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return

		}
		c.JSON(http.StatusOK, gin.H{
			"data":   result,
			"paging": paging,
		})
	}

}
