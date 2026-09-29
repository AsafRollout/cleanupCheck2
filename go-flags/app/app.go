package app

import (
	"go-flags/doit"
)

type Doer interface {
	Do(arg, message string)
}

func Run() {
	userId := "me@example.org"

	var doer Doer
	doer = doit.NewEnhancedDoer()
	doer.Do(userId, "Hello world!")
	doer.Do(userId, "Hello world again!")

	doer = doit.NewEnhancedDoer()
	doer.Do(userId, "Hello world!")
	doer.Do(userId, "Hello world again!")

	doer = doit.NewEnhancedDoer()

	doer.Do(userId, "Hello world!")
	doer.Do(userId, "Hello world again!")
}
