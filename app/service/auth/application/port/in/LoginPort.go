package in

type LoginPort interface {
	Login(email, password string) (string, error)
}
