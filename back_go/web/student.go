package web

import (
	"fmt"
	"net/http"

	"back_go/model"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
)

func Update(c *gin.Context) {
	student := model.Student{}
	if err := c.BindJSON(&student); err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, -1)
		return
	}
	if err := model.DefaultStudentDb.Update(student); err != nil {
		c.JSON(http.StatusInternalServerError, -1)
		return
	}
	c.JSON(http.StatusOK, 1)
}

func Remove(c *gin.Context) {
	id := c.Param("id")
	if err := model.DefaultStudentDb.Remove(cast.ToInt32(id)); err != nil {
		c.JSON(http.StatusInternalServerError, -1)
		return
	}
	c.JSON(http.StatusOK, 1)
}

func FindById(c *gin.Context) {
	id := c.Param("id")
	student, err := model.DefaultStudentDb.GetOne(cast.ToInt32(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, -1)
		return
	}
	c.JSON(http.StatusOK, student)
}

func FindAll(c *gin.Context) {
	students, err := model.DefaultStudentDb.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, -1)
		return
	}
	c.JSON(http.StatusOK, students)
}

func Save(c *gin.Context) {
	student := model.Student{}
	if err := c.BindJSON(&student); err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, -1)
		return
	}
	if err := model.DefaultStudentDb.Save(student); err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, -1)
		return
	}
	c.JSON(http.StatusOK, 1)
}
