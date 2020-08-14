package graph

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.
import (
	"yourJube_API/graph/model"
)

//Resolver -> ini Static DB
type Resolver struct {
	videos   []*model.Video
	user     model.User
	users    []*model.User
	comments []*model.Comment
}
