package cache

import "github.com/fsena92/golang-mux-api/entity"

type PostCache interface {
	Set(key string, value *entity.Post)
	Get(key string) *entity.Post
}
