package handler

import (
	"net/http"

	"github.com/GenkiHirano/gin-practice/repository"
	"github.com/gin-gonic/gin"
)

type requestUser struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Comment string `json:"comment"`
}

type responseUser struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Comment string `json:"comment"`
}

func CreateUser(c *gin.Context) {

	var req requestUser
	if err := c.Bind(&req); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
	}

	createdUser, err := repository.Create(req.ID, req.Name, req.Comment)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
	}

	res := responseUser{
		ID:      createdUser.ID,
		Name:    createdUser.Name,
		Comment: createdUser.Comment,
	}

	c.JSON(http.StatusCreated, res)
}
