package repo

import "sync"

type Confirm struct {
	Users map[string]string
	Mx *sync.Mutex
}

func InitConfirm() *Confirm{
	return &Confirm{
		Users: make(map[string]string),
		Mx : &sync.Mutex{},
	}
}