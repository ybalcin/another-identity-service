package account

type (
	mockRepository struct {
		InsertNewUserFn      func(user *user) *ErrUserRepository
		InsertNewUserInvoked bool
	}
)

func (r *mockRepository) InsertNewUser(user *user) *ErrUserRepository {
	r.InsertNewUserInvoked = true
	return r.InsertNewUserFn(user)
}
