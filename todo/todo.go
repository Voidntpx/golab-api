package todo

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

var db = []Todo{}
var running uint = 0

type Todo struct {
	ID        int       `json:"id"`
	Title     string    `json:"title"`
	Done      bool      `json:"done"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `json:"deleted_at"`
}

func List(c *gin.Context) {
	c.JSON(http.StatusOK, db)
}

func NewTask(c *gin.Context) {
	var t Todo
	if err := c.Bind(&t); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
	}

	t.ID = int(running)
	running++
	t.CreatedAt = time.Now()
	db = append(db, t)
	c.Status(http.StatusCreated)
}
