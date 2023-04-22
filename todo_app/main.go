package main

import (
	"fmt"
	"todo_app/app/models"
)

func main() {
	fmt.Println(models.Db)
	/*
		fmt.Println(config.Config.Port)
		fmt.Println(config.Config.SQLDriver)
		fmt.Println(config.Config.DbName)
		fmt.Println(config.Config.LogFile)

		log.Println("test")
	*/

	// u := &models.User{}
	// u.Name = "test2"
	// u.Email = "test2@com"
	// u.Password = "test"
	// fmt.Println(u)

	// u.CreateUser()

	// u, _ := models.GetUser(1)
	// fmt.Println(u)

	// u.Name = "tets2"
	// u.Email = "test2@gmail.com"
	// u.UpdateUser()
	// u, _ = models.GetUser(1)
	// fmt.Println(u)

	// u.DeleteUser()
	// u, _ = models.GetUser(1)
	// fmt.Println(u)

	/*
		user, _ := models.GetUser(2)
		user.CreateTodo("first todo")
	*/
	/*
		user, _ := models.GetUser(2)
		// todo := &models.Todo{
			Content: "first todo",
			UserId:  user.ID,
		}
		// todo.CreateTodo(todo.Content)
	*/
	/*
	   t, _ := models.GetTodo(1)
	   fmt.Println(t)
	*/
	// user, _ := models.GetUser(0)
	// user.CreateTodo("third todo by user_id")

	// todos, _ := models.GetTodos()
	// for _, v := range todos {
	// 	fmt.Println(v)
	// }
	user2, _ := models.GetUser(0)
	todos, _ := user2.GetTodosByUser()
	for _, v := range todos {
		fmt.Println(v)
	}

}
