package handler

import (
	"net/http"
	"time"

	"github.com/GenkiHirano/gin-practice/repository"
	"github.com/gin-gonic/gin"
)

type requestUser struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Comment string `json:"comment" binding:"required,max=10"`
}

type responseUser struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Comment   string    `json:"comment"`
	CreatedAt time.Time `json:"createdAt"`
}

func CreateUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req requestUser
		if err := c.Bind(&req); err != nil {
			c.JSON(http.StatusBadRequest, err.Error())
		}

		createdUser, err := repository.Create(req.ID, req.Name, req.Comment)
		if err != nil {
			c.JSON(http.StatusBadRequest, err.Error())
		}

		res := responseUser{
			ID:        createdUser.ID,
			Name:      createdUser.Name,
			Comment:   createdUser.Comment,
			CreatedAt: createdUser.CreatedAt,
		}

		c.JSON(http.StatusCreated, res)
	}
}
