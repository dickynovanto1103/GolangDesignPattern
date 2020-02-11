package iterator

import "GolangDesignPattern/behavioural/iterator/user"

type UserIterator struct {
	index int
	Users []*user.User
}

func (u *UserIterator) HasNext() bool {
	if u.index == len(u.Users) {
		return false
	}

	return true
}

func (u *UserIterator) GetNext() *user.User {
	user := u.Users[u.index]
	u.index++
	return user
}
