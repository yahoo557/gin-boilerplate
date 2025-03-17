package user

import "github.com/yahoo557/gin-boilerplate/internal/common"

type Order struct {
	ID    string
	Email string
	Name  int
	common.BaseEntity
}
