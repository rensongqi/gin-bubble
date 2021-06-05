// Control层只调用具体的逻辑
/*
url --> controller --> loginc --> model
请求		控制器		   业务逻辑     模型层的增删改查

url --> controller --> service --> dao --> model
 */

package controller

import (
	"gin-bubble/models"
	"gin-bubble/models/crud"
	"github.com/gin-gonic/gin"
	"net/http"
)

func IndexHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", nil)
}

func CreateATodo(c *gin.Context) {
	// 前端页面填写待办事项 点击提交 会发请求到这里
	// 1. 从请求中把数据拿出来
	var todo models.Todo
	c.BindJSON(&todo)
	// 2 存入数据库
	err := crud.CreateATodo(&todo)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, todo)
	}
	// 3 返回响应
}

func GetTodoList(c *gin.Context) {
	var todoList []models.Todo
	err := crud.GetTodoList(&todoList)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, todoList)
	}
}

func UpdateATodo(c *gin.Context) {
	id, ok := c.Params.Get("id")
	if !ok {
		c.JSON(http.StatusOK, gin.H{"error": "id dose not exist."})
		return
	}
	var todo models.Todo
	err := crud.UpdateATodo(id, &todo)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
		return
	}
	c.BindJSON(&todo)
	//if err = DB.Model(&todo).Where("id=?", id).Update("status", 1).Error; err != nil {
	//	c.JSON(http.StatusOK, gin.H{"error": err.Error()})
	//} else {
	//	c.JSON(http.StatusOK, todo)
	//}
	err = crud.UpdateATodoSave(&todo)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, todo)
	}
}

func DeleteATodo(c *gin.Context) {
	id, ok := c.Params.Get("id")
	if !ok {
		c.JSON(http.StatusOK, gin.H{"error": "id dose not exist."})
		return
	}
	var todo models.Todo
	err := crud.DeleteATodo(id, &todo)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, todo)
	}
}

