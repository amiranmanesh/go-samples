package main

import "sync"

type SuccessFunc func(string)
type FailFunc func(error)
type ExecuteStringFunc func() (string, error)

type MaybeString struct {
	successFunc SuccessFunc
	failFunc    FailFunc
}

func (m *MaybeString) Success(f SuccessFunc) *MaybeString {
	m.successFunc = f
	return m
}

func (m *MaybeString) Fail(f FailFunc) *MaybeString {
	m.failFunc = f
	return m
}

func (m *MaybeString) Execute(f ExecuteStringFunc) {
	go func(s *MaybeString) {
		str, err := f()
		if err != nil {
			s.failFunc(err)
		} else {
			s.successFunc(str)
		}
	}(m)
}

func main() {
	var wait sync.WaitGroup
	wait.Add(1)
	futures := &MaybeString{}
	futures.Success(func(s string) {
		println(s)
		wait.Done()
	})
	futures.Fail(func(err error) {
		println(err)
		wait.Done()
	})
	futures.Execute(func() (string, error) {
		return "Hello, world!", nil
	})
	wait.Wait()
}
