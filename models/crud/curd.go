package crud

import (
	"gin-bubble/dao"
	"gin-bubble/models"
)

// Todo CURD
// CreateATodo

func CreateATodo(todo *models.Todo) (err error){
	if err = dao.DB.Create(&todo).Error; err != nil {
		return err
	}
	return
}

func UpdateATodo(id string, todo *models.Todo) (err error) {
	if err = dao.DB.Where("id=?", id).First(&todo).Error; err != nil {
		return err
	}
	return
}

func UpdateATodoSave(todo *models.Todo) (err error) {
	if err = dao.DB.Save(&todo).Error; err != nil {
		return err
	}
	return
}

func DeleteATodo(id string, todo *models.Todo) (err error) {
	if err = dao.DB.Where("id=?", id).Delete(&todo).Error; err != nil {
		return err
	}
	return
}

func GetTodoList(todoList *[]models.Todo) (err error) {
	if err = dao.DB.Find(&todoList).Error; err != nil {
		return err
	}
	return
}
