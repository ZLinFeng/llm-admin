package service

type ServiceGroup struct {
	UserService
	AuthService
}

var ServiceApp = new(ServiceGroup)
