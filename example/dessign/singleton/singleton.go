package singleton

import "sync"

//使用懒惰模式的单例模式，使用双重检查加锁保证线程安全

type Singleton interface {
	foo()
}

type singleton struct{}

func (s singleton) foo() {}

var (
	instance *singleton
	once sync.Once
)

func GetInstance() Singleton {
	once.Do(func() {
		instance = &singleton{}
	})
	return instance
}