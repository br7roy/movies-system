package initRouter

import (
	"github.com/gin-gonic/gin"
	"movie-system/app/handler"
)

func SetupRouter(engine *gin.Engine) *gin.Engine {

	/*	engine.GET("/", func(context *gin.Context) {
		context.JSON(200, gin.H{
			"hello": "https://takfu.tk/movie",
			"blog":  "https://takblog.netlify.com",
		})
	})*/

	uGroup := engine.Group("/user")
	{
		uGroup.Any("", handler.UserNameAgeSave)
		//uGroup.Any("/:name", handler.UserSave)
		uGroup.GET("/rj", handler.UserRegisterJump)
		uGroup.POST("/register", handler.UserRegister)
	}

	mGoup := engine.Group("")
	{
		mGoup.GET("", handler.MovieIndex)

		mGoup.GET("/show", handler.ShowAllMovies)
	}

	return engine
}