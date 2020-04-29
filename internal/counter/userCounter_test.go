package counter

import (
	"math/rand"
	"sync"
	"testing"
	"time"
)

func BenchmarkSample(b *testing.B) {
	userCounter := newUserCounter()
	rand.Seed(time.Now().UnixNano())
	sliceUsersIds := []string{"test1", "test2", "test3", "test4", "test5", "test6", "test7", "test8", "test9"}
	for i := 0; i < b.N; i++ {
		userCounter.incrUser(sliceUsersIds[rand.Intn(8)])
	}
}

func TestCounter1(t *testing.T) {
	userCounter := newUserCounter()

	var wg sync.WaitGroup
	wg.Add(500)
	go func() {
		userId := "test1"
		for i := 0; i < 100; i++ {
			go func() {
				userCounter.incrUser(userId)
				defer wg.Done()
			}()
		}
	}()
	go func() {
		userId := "test2"
		for i := 0; i < 100; i++ {
			go func() {
				userCounter.incrUser(userId)
				defer wg.Done()
			}()
		}
	}()
	go func() {
		userId := "test3"
		for i := 0; i < 100; i++ {
			go func() {
				userCounter.incrUser(userId)
				defer wg.Done()
			}()
		}
	}()
	go func() {
		userId := "test4"
		for i := 0; i < 100; i++ {
			go func() {
				userCounter.incrUser(userId)
				defer wg.Done()
			}()
		}
	}()
	go func() {
		userId := "test5"
		for i := 0; i < 100; i++ {
			go func() {
				userCounter.incrUser(userId)
				defer wg.Done()
			}()
		}
	}()

	wg.Wait()

	if userCounter.getRobotCount() != 5 {
		t.Error("Expected 5, got ", userCounter.getRobotCount())
	}
}

func TestCounter3(t *testing.T) {
	userCounter := newUserCounter()

	for i := 0; i < 100; i++ {
		userCounter.incrUser("test1")
	}

	if userCounter.getRobotCount() != 1 {
		t.Error("Expected 1, got ", userCounter.getRobotCount())
	}

	for i := 0; i < 100; i++ {
		userCounter.incrUser("test2")
	}

	if userCounter.getRobotCount() != 2 {
		t.Error("Expected 1, got ", userCounter.getRobotCount())
	}

	userCounter.reset()

	if userCounter.getRobotCount() != 0 {
		t.Error("Expected 0, got ", userCounter.getRobotCount())
	}

	for i := 0; i < 100; i++ {
		userCounter.incrUser("test2")
	}

	if userCounter.getRobotCount() != 1 {
		t.Error("Expected 1, got ", userCounter.getRobotCount())
	}
}
