package main

import (
	"fmt"
	"net/http"

	"github.com/fsena92/golang-mux-api/cache"
	"github.com/fsena92/golang-mux-api/controller"
	router "github.com/fsena92/golang-mux-api/http"
	"github.com/fsena92/golang-mux-api/repository"
	"github.com/fsena92/golang-mux-api/service"
)

var (
	httpRouter     router.Router             = router.NewMuxRouter()
	postRepository repository.PostRepository = repository.NewFirestoreRepository()
	postService    service.PostService       = service.NewPostService(postRepository)
	postCache      cache.PostCache           = cache.NewRedisCache("localhost;6379", 1, 10)
	postController controller.PostController = controller.NewPostController(postService, postCache)
)

func main() {
	const port string = ":8000"

	httpRouter.GET("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Up and running...")
	})
	httpRouter.GET("/posts", postController.GetPosts)
	httpRouter.GET("posts/{id}", postController.GetPostByID)
	httpRouter.POST("/posts", postController.AddPost)

	httpRouter.SERVE(port)
}
