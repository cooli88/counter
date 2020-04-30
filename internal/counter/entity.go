package counter

import (
	"sync/atomic"
)

type listUserCount map[string]*userCount

func newListUserCount() listUserCount {
	return make(map[string]*userCount)
}

func (l listUserCount) userCount(userId string) *userCount {
	if el, ok := l[userId]; ok {
		return el
	}
	return nil
}

type userCount struct {
	Count int32
}

func newUserCount() *userCount {
	return &userCount{Count: 1}
}

func (u *userCount) incr() {
	atomic.AddInt32(&u.Count, 1)
}
