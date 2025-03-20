package user

import "github.com/yahoo557/gin-boilerplate/internal/common"

type User struct {
	ID    string
	Email string
	Name  int
	common.BaseEntity
}
