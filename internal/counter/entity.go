package counter

import (
	"sync/atomic"
)

type ListUserCount map[string]*UserCount

func newListUserCount() ListUserCount {
	return make(map[string]*UserCount)
}

func (l ListUserCount) getUserCount(userId string) *UserCount {
	if el, ok := l[userId]; ok {
		return el
	}
	return nil
}

type UserCount struct {
	Count int32
}

func newUserCount() *UserCount {
	return &UserCount{Count: 1}
}

func (u *UserCount) incr() {
	atomic.AddInt32(&u.Count, 1)
}
