package repo


type UserRepo struct {}

func NewUserRepo() *UserRepo {
	return &UserRepo{}
}

func (ur *UserRepo) GetInfoUserByID() string {
	return "anhduong"
}