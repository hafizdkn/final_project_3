package main

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"

	"final_project_3/auth"
	"final_project_3/category"
	"final_project_3/database"
	"final_project_3/handler"
	"final_project_3/middleware"
	"final_project_3/task"
	"final_project_3/user"
)

func main() {
	db, err := database.ConnectionDB()
	if err != nil {
		panic(err)
	}

	authService := auth.NewJwtService

	userRepository := user.NewUserRepository(db)
	userService := user.NewUserService(userRepository, authService)
	userHandler := handler.NewUserHandler(userService)

	categoryRepository := category.NewCategoryRepository(db)
	categoryService := category.NewCategoryService(categoryRepository)
	categoryHandler := handler.NewCategoryHandler(categoryService)

	taskRepository := task.NewCategoryRepository(db)
	taskService := task.NewCategoryServie(taskRepository, categoryService)
	taskHandler := handler.NewTaskHandler(taskService)

	app := gin.Default()
	users := app.Group("/users")
	{
		users.POST("/register", userHandler.UserRegister)
		users.POST("/login", userHandler.UserLogin)
		users.GET("/", userHandler.GetAllUsers)
		users.PUT("/update-account", middleware.AuthMiddleware(authService, userService), userHandler.UserUpdate)
		users.DELETE("/delete-account", middleware.AuthMiddleware(authService, userService), userHandler.DeleteUser)
	}

	categories := app.Group("/categories")
	{
		categories.POST("/", middleware.AuthMiddleware(authService, userService), categoryHandler.CreateCategory)
		categories.GET("/", middleware.AuthMiddleware(authService, userService), categoryHandler.GetCategorys)
		categories.PATCH("/:categoryId", middleware.AuthMiddleware(authService, userService), categoryHandler.UpdateCategory)
		categories.DELETE("/:categoryId", middleware.AuthMiddleware(authService, userService), categoryHandler.DeleteCategory)
	}

	tasks := app.Group("/tasks")
	{
		tasks.POST("/", middleware.AuthMiddleware(authService, userService), taskHandler.CreateTask)
		tasks.GET("/", middleware.AuthMiddleware(authService, userService), taskHandler.GetTasks)
		tasks.PUT("/:taksId", middleware.AuthMiddleware(authService, userService), taskHandler.UpdateTask)
		tasks.PATCH("/update-status/:taksId", middleware.AuthMiddleware(authService, userService), taskHandler.UpdateStatusTask)
		tasks.PATCH("/update-category/:taksId", middleware.AuthMiddleware(authService, userService), taskHandler.UpdateStatusCategory)
		tasks.DELETE("/:taksId", middleware.AuthMiddleware(authService, userService), taskHandler.DeleteTask)
	}

	app.Run(fmt.Sprintf(":%s", os.Getenv("PORT")))

	// app.Run(":8080")
}
