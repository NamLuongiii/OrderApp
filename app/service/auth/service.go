package auth

type Service interface {
	getUser() string
}

type ServiceImpl struct {
}

func NewService() Service {
	return &ServiceImpl{}
}
