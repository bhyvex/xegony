package cases

import (
	"fmt"

	"github.com/xackery/xegony/model"
	"github.com/xackery/xegony/storage"
	"github.com/xeipuuv/gojsonschema"
)

type UserRepository struct {
	stor storage.Storage
}

func (g *UserRepository) Initialize(stor storage.Storage) (err error) {
	if stor == nil {
		err = fmt.Errorf("Invalid storage type")
		return
	}
	g.stor = stor
	return
}

func (g *UserRepository) Get(userId int64) (user *model.User, err error) {
	if userId == 0 {
		err = fmt.Errorf("Invalid User ID")
		return
	}
	user, err = g.stor.GetUser(userId)
	return
}

func (g *UserRepository) Create(user *model.User) (err error) {
	if user == nil {
		err = fmt.Errorf("Empty user")
		return
	}
	schema, err := user.NewSchema([]string{"name", "password", "email"}, nil)
	if err != nil {
		return
	}
	user.Id = 0 //strip ID
	result, err := schema.Validate(gojsonschema.NewGoLoader(user))
	if err != nil {
		return
	}
	if !result.Valid() {
		vErr := &model.ErrValidation{
			Message: "invalid",
		}
		vErr.Reasons = map[string]string{}
		for _, res := range result.Errors() {
			vErr.Reasons[res.Field()] = res.Description()
		}
		err = vErr
		return
	}
	err = g.stor.CreateUser(user)
	if err != nil {
		return
	}
	return
}

func (g *UserRepository) Login(username string, password string) (user *model.User, err error) {
	user = &model.User{
		Name:     username,
		Password: password,
	}

	schema, err := user.NewSchema([]string{"name", "password"}, nil)
	if err != nil {
		return
	}
	result, err := schema.Validate(gojsonschema.NewGoLoader(user))
	if err != nil {
		return
	}
	if !result.Valid() {
		vErr := &model.ErrValidation{
			Message: "invalid",
		}
		vErr.Reasons = map[string]string{}
		for _, res := range result.Errors() {
			vErr.Reasons[res.Field()] = res.Description()
		}
		err = vErr
		return
	}
	user, err = g.stor.LoginUser(username, password)
	if err != nil {
		return
	}
	return
}

func (g *UserRepository) Edit(userId int64, user *model.User) (err error) {
	schema, err := user.NewSchema([]string{"name"}, nil)
	if err != nil {
		return
	}

	result, err := schema.Validate(gojsonschema.NewGoLoader(user))
	if err != nil {
		return
	}
	if !result.Valid() {
		vErr := &model.ErrValidation{
			Message: "invalid",
		}
		vErr.Reasons = map[string]string{}
		for _, res := range result.Errors() {
			vErr.Reasons[res.Field()] = res.Description()
		}
		err = vErr
		return
	}

	err = g.stor.EditUser(userId, user)
	if err != nil {
		return
	}
	return
}

func (g *UserRepository) Delete(userId int64) (err error) {
	err = g.stor.DeleteUser(userId)
	if err != nil {
		return
	}
	return
}

func (g *UserRepository) List() (users []*model.User, err error) {
	users, err = g.stor.ListUser()
	if err != nil {
		return
	}
	return
}
