package gin

import "github.com/hubertme/gin-demo/gin/handlers"

func initGinHandlerGroups() {
	// Test Endpoints
	registerGetGinHandler()
	registerPostGinHandler()
	//registerTestDataGinHandler()

	// App Endpoints
	registerUserGinHandler()
}

func registerGetGinHandler() {
	getRoute := router.Group("/get")

	getRoute.GET("/ping", handlers.PingHandler)
	getRoute.GET("/pong", handlers.PongHandler)
	getRoute.GET("/devError", handlers.DevErrorHandler)
}

func registerPostGinHandler() {
	postRoute := router.Group("/post")

	postRoute.POST("/submit", handlers.SubmitHandler)
}

func registerUserGinHandler() {
	userRoute := router.Group("/user")

	userRoute.GET("/fetch", handlers.GetAllUsers)
	userRoute.GET("/fetch/:id", handlers.GetUserById)
	userRoute.POST("/add", handlers.PostNewUser)
	userRoute.PUT("/update/:id", handlers.UpdateUserById)
	userRoute.DELETE("/delete", handlers.DeleteAllUsers)
	userRoute.DELETE("/delete/:id", handlers.DeleteUserById)
}

func registerTestDataGinHandler() {
	dataTestRoute := router.Group("/data")

	dataTestRoute.GET("/fetch", handlers.GetAllTestUsers)
	dataTestRoute.GET("/fetch/:id", handlers.GetTestUserById)
	dataTestRoute.POST("/add", handlers.PostNewTestUser)
	dataTestRoute.PUT("/update/:id", handlers.UpdateTestUserById)
	dataTestRoute.DELETE("/delete", handlers.DeleteAllTestUsers)
	dataTestRoute.DELETE("/delete/:id", handlers.DeleteTestUserById)
}
