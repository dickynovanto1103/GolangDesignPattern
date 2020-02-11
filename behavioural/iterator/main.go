package main

import (
	"GolangDesignPattern/behavioural/iterator/collection"
	"GolangDesignPattern/behavioural/iterator/user"
	"fmt"
)

func main() {
	//create an array of user
	user1 := &user.User{
		Name: "dicky",
		Age: 22,
	}
	user2 := &user.User{
		Name: "angel",
		Age:  20,
	}
	users := []*user.User{user1, user2}

	//wrap it inside collection
	userCollection := &collection.UserCollection{Users: users}

	//get iterator
	userIterator := userCollection.CreateIterator()

	//iterate using the iterator
	for userIterator.HasNext() {
		fmt.Println(userIterator.GetNext())
	}
}
