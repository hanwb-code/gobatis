package mapper

import (
	"github.com/gobatis/gobatis/example/entity"
)

var (
	User = &userMapper{}
)

type userMapper struct {
	AddUser func(user *entity.User) (rows int, err error)
	List    func(param map[string]interface{}) ([]*entity.User, error)
}
