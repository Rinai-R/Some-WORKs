package Add

import (
	"Golang/12December/20241202/MySQL/dao"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func Add(c *gin.Context) {
	id := c.PostForm("id")
	name := c.PostForm("name")
	email := c.PostForm("email")
	phone := c.PostForm("phone")
	hire_date := c.PostForm("hire_date")
	salary := c.PostForm("salary")
	department := c.PostForm("department")
	manager_id := c.PostForm("manager_id")
	status := c.PostForm("status")
	created_at := c.PostForm("created_at")
	updated_at := c.PostForm("updated_at")

	//flag := dao.Search(id)
	//if flag {
	//	c.JSON(http.StatusInternalServerError, gin.H{
	//		"code":    200,
	//		"message": "id重复！！！",
	//	})
	//	return
	//}
	H, err := time.Parse("2006-01-02", hire_date)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "hire_date 格式错误",
		})
		return
	}

	C, err2 := time.Parse("2006-01-02", created_at)
	if err2 != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "created_at 格式错误",
		})
		return
	}

	U, err3 := time.Parse("2006-01-02", updated_at)
	if err3 != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "updated_at 格式错误",
		})
		return
	}
	dao.AddUser(id, name, email, phone, H, salary, department, manager_id, status, C, U)
	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "success",
	})
}
