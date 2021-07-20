package account

type (
	mockRepository struct {
		InsertNewUserFn      func(user *user) *errRepository
		InsertNewUserInvoked bool
	}
)

func (r *mockRepository) InsertNewUser(user *user) *errRepository {
	r.InsertNewUserInvoked = true
	return r.InsertNewUserFn(user)
}
