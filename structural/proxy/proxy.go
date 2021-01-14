package proxy

import "errors"

//remote proxy: rpc net proxy
//virtual proxy: big image database
//protection proxy: restricted access object
type UserFinder interface {
	FindUser(id int32) (User, error)
}

type User struct {
	ID int32
}

//db
type UserList []User

type UserListProxy struct {
	SomeDatabase           *UserList
	StackCache             UserList
	StackCapacity          int
	DidLastSearchUsedCache bool
}

func (u *UserListProxy) FindUser(id int32) (User, error) {
	return User{}, errors.New("not implemented yet")
}
