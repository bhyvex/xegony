package api

import (
	"database/sql"
	"net/http"

	"github.com/pkg/errors"
	"github.com/xackery/xegony/cases"
	"github.com/xackery/xegony/model"
)

// UserRequest is a list of parameters used for user
// swagger:parameters deleteUser editUser getUser
type UserRequest struct {
	// ID to get information about
	// in: path
	// example: 74887
	ID int64 `json:"ID"`
}

// UserResponse is what endpoints respond with
// swagger:response
type UserResponse struct {
	User *model.User `json:"user,omitempty"`
}

// UserCreateRequest is the body parameters for creating an user
// swagger:parameters createUser
type UserCreateRequest struct {
	// User details to create
	// in: body
	User *model.User `json:"user"`
}

// UserLoginRequest is the body parameters for creating an user
// swagger:parameters postUserLogin
type UserLoginRequest struct {
	// User details to create
	// in: body
	User *model.User `json:"user"`
}

// UserLoginResponse returns a token along with user details
// swagger:response
type UserLoginResponse struct {
	User  *model.User `json:"user,omitempty"`
	Token string      `json:"token"`
}

// UserEditRequest is the body parameters for editing an user
// swagger:parameters editUser
type UserEditRequest struct {
	// ID to get information about
	// in: path
	// example: 74887
	ID int64 `json:"ID"`
	// User details to edit
	// in: body
	User *model.User `json:"user"`
}

// UserLinkRequest is the body parameters for linking an user
// swagger:parameters getUserLink
type UserLinkRequest struct {
	// link to get information about
	// in: path
	// example: hashCode
	Link string `json:"link"`
}

// UserLinkResponse is a general response to a request
// swagger:response
type UserLinkResponse struct {
	Account   *model.Account   `json:"account,omitempty"`
	Character *model.Character `json:"character,omitempty"`
}

// UsersRequest is a list of parameters used for user
// swagger:parameters listUser
type UsersRequest struct {
	// Offset is pagination, offset*limit
	// example: 0
	// in: query
	Offset int64 `json:"offset"`
	// Limit to how many items per page
	// example: 10
	// in: query
	Limit int64 `json:"limit"`
	// OrderBy is which field to order a page by
	// example: short_name
	// in: query
	OrderBy string `json:"orderBy"`
	// IsDescending will change sort order when true
	// example: 0
	// in: query
	IsDescending int64 `json:"isDescending"`
}

// UsersResponse is a general response to a request
// swagger:response
type UsersResponse struct {
	Page  *model.Page `json:"page,omitempty"`
	Users model.Users `json:"users,omitempty"`
}

// UsersBySearchRequest is a list of parameters used for user
// swagger:parameters listUserBySearch
type UsersBySearchRequest struct {
	// ShortName is which user to get information about
	// example: xackery
	// in: query
	ShortName string `json:"shortName"`
	// Offset is pagination, offset*limit
	// example: 0
	// in: query
	Offset int64 `json:"offset"`
	// Limit to how many items per page
	// example: 10
	// in: query
	Limit int64 `json:"limit"`
	// OrderBy is which field to order a page by
	// example: short_name
	// in: query
	OrderBy string `json:"orderBy"`
	// IsDescending will change sort order when true
	// example: 0
	// in: query
	IsDescending int64 `json:"isDescending"`
}

// UsersBySearchResponse is a general response to a request
// swagger:response
type UsersBySearchResponse struct {
	Search *model.User `json:"search,omitempty"`
	Page   *model.Page `json:"page,omitempty"`
	Users  model.Users `json:"users,omitempty"`
}

func userRoutes() (routes []*route) {

	routes = []*route{
		// swagger:route GET /user user listUser
		//
		// Lists users
		//
		// This will show all available users by default.
		//
		//     Consumes:
		//     - application/json
		//
		//     Produces:
		//     - application/json
		//     - application/xml
		//     - application/yaml
		//
		//
		//     Responses:
		//       default: ErrInternal
		//       200: UsersResponse
		//       400: ErrValidation
		//		 401: ErrPermission
		{
			"ListUser",
			"GET",
			"/user",
			listUser,
		},
		// swagger:route GET /user/search user listUserBySearch
		//
		// Search users by name
		//
		// This will show all available users by default.
		//
		//     Consumes:
		//     - application/json
		//
		//     Produces:
		//     - application/json
		//     - application/xml
		//     - application/yaml
		//
		//
		//     Responses:
		//       default: ErrInternal
		//       200: UsersBySearchResponse
		//       400: ErrValidation
		//		 401: ErrPermission
		{
			"ListUserBySearch",
			"GET",
			"/user/search",
			listUserBySearch,
		},
		// swagger:route POST /user user createUser
		//
		// Create an user
		//
		// This will create a user
		//
		//     Security:
		//       apiKey:
		//
		//     Responses:
		//       default: ErrInternal
		//		 200: UserResponse
		//       204: ErrNoContent
		//       400: ErrValidation
		//		 401: ErrPermission
		{
			"CreateUser",
			"POST",
			"/user",
			createUser,
		},
		// swagger:route GET /user/{ID} user getUser
		//
		// Get an user
		//
		// This will get an individual user available users by default.
		//
		//     Responses:
		//       default: ErrInternal
		//       200: UserResponse
		//       400: ErrValidation
		//		 401: ErrPermission
		{
			"GetUser",
			"GET",
			"/user/{ID:[0-9]+}",
			getUser,
		},
		// swagger:route PUT /user/{ID} user editUser
		//
		// Edit an user
		//
		// This will edit an user
		//
		//     Security:
		//       apiKey:
		//
		//     Responses:
		//       default: ErrInternal
		//		 200: UserResponse
		//       204: ErrNoContent
		//       400: ErrValidation
		//		 401: ErrPermission
		{
			"EditUser",
			"PUT",
			"/user/{ID:[0-9]+}",
			editUser,
		},
		// swagger:route DELETE /user/{ID} user deleteUser
		//
		// Delete an user
		//
		// This will delete an user
		//
		//     Security:
		//       apiKey:
		//
		//     Responses:
		//       default: ErrInternal
		//       204: ErrNoContent
		//       400: ErrValidation
		//		 401: ErrPermission
		{
			"DeleteUser",
			"DELETE",
			"/user/{ID:[0-9]+}",
			deleteUser,
		},
		// swagger:route GET /user/link/{link} user getUserLink
		//
		// Get a user link
		//
		// Used for linking a user to an account
		//
		//     Responses:
		//       default: ErrInternal
		//       200: UserResponse
		//       400: ErrValidation
		//		 401: ErrPermission
		{
			"GetUserLink",
			"GET",
			"/user/link/{link:[a-zA-Z0-9]+}",
			getUserLink,
		},
		// swagger:route POST /user/login user postUserLogin
		//
		// Logs in a user
		//
		// Logs in a user with provided data
		//
		//     Consumes:
		//     - application/json
		//
		//     Produces:
		//     - application/json
		//     - application/xml
		//     - application/yaml
		//
		//
		//     Responses:
		//       default: ErrInternal
		//       200: UsersResponse
		//       400: ErrValidation
		//		 401: ErrPermission
		{
			"PostUserLogin",
			"POST",
			"/user/login",
			postUserLogin,
		},
	}
	return
}

func getUser(w http.ResponseWriter, r *http.Request, user *model.User, statusCode int) (content interface{}, err error) {
	request := &UserRequest{
		ID: getIntVar(r, "ID"),
	}

	focusUser := &model.User{
		ID: request.ID,
	}

	err = cases.GetUser(focusUser, user)
	if err != nil {
		if err == sql.ErrNoRows {
			return
		}
		err = errors.Wrap(err, "Request error")
		return
	}
	response := &UserResponse{
		User: focusUser,
	}
	content = response
	return
}

func createUser(w http.ResponseWriter, r *http.Request, user *model.User, statusCode int) (content interface{}, err error) {
	focusUser := &model.User{}
	err = decodeBody(r, focusUser)
	if err != nil {
		return
	}
	err = cases.CreateUser(focusUser, user)
	if err != nil {
		return
	}
	response := &UserResponse{
		User: focusUser,
	}
	content = response
	return
}

func deleteUser(w http.ResponseWriter, r *http.Request, user *model.User, statusCode int) (content interface{}, err error) {
	request := &UserRequest{
		ID: getIntVar(r, "ID"),
	}

	focusUser := &model.User{
		ID: request.ID,
	}

	err = cases.DeleteUser(focusUser, user)
	if err != nil {
		switch errors.Cause(err).(type) {
		case *model.ErrNoContent:
			return
		default:
			err = errors.Wrap(err, "Request failed")
		}
	}
	err = &model.ErrNoContent{}
	return
}

func editUser(w http.ResponseWriter, r *http.Request, user *model.User, statusCode int) (content interface{}, err error) {
	request := &UserEditRequest{
		ID: getIntVar(r, "ID"),
	}

	focusUser := &model.User{
		ID: request.ID,
	}

	err = decodeBody(r, focusUser)
	if err != nil {
		err = errors.Wrap(err, "Request error")
		return
	}

	err = cases.EditUser(user, user)
	if err != nil {
		return
	}
	response := &UserResponse{
		User: focusUser,
	}
	content = response
	return
}

func listUser(w http.ResponseWriter, r *http.Request, user *model.User, statusCode int) (content interface{}, err error) {

	page := &model.Page{
		Offset:       getIntQuery(r, "offset"),
		Limit:        getIntQuery(r, "limit"),
		OrderBy:      getQuery(r, "orderBy"),
		IsDescending: getIntQuery(r, "isDescending"),
	}
	users, err := cases.ListUser(page, user)
	if err != nil {
		err = errors.Wrap(err, "Request error")
		return
	}

	response := &UsersResponse{
		Page:  page,
		Users: users,
	}
	content = response
	return
}

func listUserBySearch(w http.ResponseWriter, r *http.Request, user *model.User, statusCode int) (content interface{}, err error) {

	page := &model.Page{
		Offset:       getIntQuery(r, "offset"),
		Limit:        getIntQuery(r, "limit"),
		OrderBy:      getQuery(r, "orderBy"),
		IsDescending: getIntQuery(r, "isDescending"),
	}
	focusUser := &model.User{}

	user.DisplayName = getQuery(r, "displayName")
	users, err := cases.ListUserBySearch(page, focusUser, user)
	if err != nil {
		err = errors.Wrap(err, "Request error")
		return
	}

	response := &UsersBySearchResponse{
		Page:   page,
		Users:  users,
		Search: focusUser,
	}
	content = response
	return
}

func getUserLink(w http.ResponseWriter, r *http.Request, user *model.User, statusCode int) (content interface{}, err error) {
	request := &UserLinkRequest{
		Link: getVar(r, "link"),
	}

	userLink := &model.UserLink{
		Link: request.Link,
	}

	err = cases.GetUserLink(userLink, user)
	if err != nil {
		if err == sql.ErrNoRows {
			return
		}
		err = errors.Wrap(err, "Request error")
		return
	}

	response := &UserLinkResponse{
		Character: userLink.Character,
		Account:   userLink.Account,
	}
	content = response
	return
}

func postUserLogin(w http.ResponseWriter, r *http.Request, user *model.User, statusCode int) (content interface{}, err error) {
	focusUser := &model.User{}
	err = decodeBody(r, focusUser)
	if err != nil {
		return
	}
	err = cases.LoginUser(focusUser)
	if err != nil {
		return
	}
	token, err := generateToken(focusUser)
	if err != nil {
		return
	}
	response := &UserLoginResponse{
		User:  focusUser,
		Token: token,
	}
	content = response
	return
}
