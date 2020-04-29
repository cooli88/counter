package counter

import (
	"sync"
	"sync/atomic"
	"time"
)

const CountRequestIsRobot = 100

type UserCounter struct {
	listUserCount ListUserCount
	count         int32
	mux           sync.Mutex
}

func newUserCounter() *UserCounter {
	var userCounter = &UserCounter{}
	userCounter.reset()
	startResetSchedulerUserCounter(userCounter)
	return userCounter
}

//startResetSchedulerUserCounter запуск сброса счетчика по рассписанию (каждую минуту)
func startResetSchedulerUserCounter(userCounter *UserCounter) {
	ticker := time.NewTicker(time.Minute * 1)
	go func() {
		for _ = range ticker.C {
			userCounter.reset()
		}
	}()
}

//getRobotCount  получить кол-во роботов за последнюю минуту
func (u *UserCounter) getRobotCount() int32 {
	return u.count
}

//incrUser увеличить счетчик для пользователя
func (u *UserCounter) incrUser(userId string) {
	u.mux.Lock()
	userCount := u.listUserCount.getUserCount(userId)
	if userCount == nil {
		u.listUserCount[userId] = newUserCount()
		u.mux.Unlock()
		return
	}
	u.mux.Unlock()
	if userCount.Count >= CountRequestIsRobot {
		return
	}
	userCount.incr()
	if userCount.Count == CountRequestIsRobot {
		u.inc()
	}
}

//inc увеличить счетчик роботов
func (u *UserCounter) inc() {
	atomic.AddInt32(&u.count, 1)
}

//reset сбросить записи счетчика
func (u *UserCounter) reset() {
	u.mux.Lock()
	defer u.mux.Unlock()
	u.listUserCount = newListUserCount()
	u.count = 0
}
