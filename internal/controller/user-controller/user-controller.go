package usercontroller

type userProvider interface{}

type UserController struct {
	user userProvider
}

func NewUserController(user userProvider) *UserController {
	return &UserController{
		user: user,
	}
}
