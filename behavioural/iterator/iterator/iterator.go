package iterator

import "GolangDesignPattern/behavioural/iterator/user"

type Iterator interface {
	HasNext() bool
	GetNext() *user.User
}
