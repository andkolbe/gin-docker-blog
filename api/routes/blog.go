package routes

import (
	"github.com/andkolbe/gin-docker-blog/api/controller"
	"github.com/andkolbe/gin-docker-blog/infrastructure"
)

// Routes for question module
type PostRoute struct {
	Controller controller.PostController
	Handler infrastructure.GinRouter
}

// initializes new choice routes
func NewPostRoute(
	controller controller.PostController,
	handler infrastructure.GinRouter,
) PostRoute {
	return PostRoute{
		Controller: controller,
		Handler: handler,
	}
}

// setups new choice Routes
func (p PostRoute) Setup() {
	post := p.Handler.Gin.Group("/posts") // router group
	{
		post.GET("/", p.Controller.GetPosts)
        post.POST("/", p.Controller.AddPost)
        post.GET("/:id", p.Controller.GetPost)
        post.DELETE("/:id", p.Controller.DeletePost)
        post.PUT("/:id", p.Controller.UpdatePost)
	}
}