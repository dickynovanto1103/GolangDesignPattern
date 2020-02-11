package collection

import (
	"GolangDesignPattern/behavioural/iterator/iterator"
	"GolangDesignPattern/behavioural/iterator/user"
)

type UserCollection struct {
	Users []*user.User
}

func (u *UserCollection) CreateIterator() iterator.Iterator {
	return &iterator.UserIterator{
		Users: u.Users,
	}
}