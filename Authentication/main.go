package main

import (
	auth "github.com/princeprabhat/Authentication/Auth"
	controllers "github.com/princeprabhat/Authentication/Controllers"
	database "github.com/princeprabhat/Authentication/Database"
	models "github.com/princeprabhat/Authentication/Models"
)

func main() {
	db := database.InitDatabase()
	newUser := models.StartUser(db)

	authController := auth.NewAuthController(db)
	r := controllers.InitRouter()

	r.GET("/AllUsers", newUser.GetUsers)

	r.POST("/Signup", authController.SignUp)
	r.GET("/SingleUser/:id", newUser.GetOneUser)
	r.POST("/CreateUser", newUser.CreateUser)
	r.GET("/UpdateUser/:id", newUser.UpdateUser)
	r.DELETE("/DeleteUser/:id", newUser.DeleteUser)
	r.POST("/Login", authController.Login)
	r.POST("/Verify", authController.VerifyUser)
	r.GET("/Welcome", authController.Welcome)
	r.GET("/Logout", authController.Logout)
	r.Run()
}
