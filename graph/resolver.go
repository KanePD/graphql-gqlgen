package graph
import (
	"github.com/kanePD/gqlgen-todos/graph/model"
)
// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct{
	todos []*model.Todo
	articles []*model.Article
}
