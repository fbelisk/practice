package main

import "sync"

type a struct {

}

var testLock *sync.Mutex
var test *a

func GetInstance1() *a{
	if test != nil {
		return test
	}
	testLock.Lock()
	defer func() {
		testLock.Unlock()
	}()
	if test == nil {
		test = &a{}
	}
	return test
}

var testOnce *sync.Once

func GetInstance2() *a{
	testOnce.Do(func() {
		test = &a{}
	})
	return test
}

func init() {
	test = &a{}
}

func GetInstance3() *a{
	return test
}

