package router

import (
	"bili-favorites/app/logic"
	"bili-favorites/app/middleware"
	"github.com/gin-gonic/gin"
)

func InitRouter() {
	router := gin.Default()

	router.Use(middleware.Headers)

	router.GET("/", logic.Index)

	fav := router.Group("/x/v3/fav")
	{
		fav.GET("/resource/list", logic.ResourceList)

		folder := fav.Group("/folder")
		{
			folder.GET("/created/list-all", logic.CreatedListAll)
			folder.GET("/collected/list", logic.CollectedList)
		}
	}

	if err := router.Run(":3128"); err != nil {
		panic(err)
	}
}
