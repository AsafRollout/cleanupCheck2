package app

import (
	"go-flags/doit"
)

type Doer interface {
	Do(arg, message string)
}

func Run() {
	userId := "me@example.org"

	// Reference to 'EnableNewFeature.IsEnabled()' removed & replaced with flag's 'True' case:
	doer := doit.NewEnhancedDoer()
	doer.Do(userId, "Hello world!")
	doer.Do(userId, "Hello world again!")

	// Reference to 'EnableNewFeature.IsEnabled()' & context removed & replaced with flag's 'True' case:
	doer = doit.NewEnhancedDoer()
	doer.Do(userId, "Hello world!")
	doer.Do(userId, "Hello world again!")

	// Reference to 'EnableNewFeature.Enabled()' removed & replaced with flag's 'True' case:
	doer = doit.NewEnhancedDoer()

	// Reference to 'EnableNewFeature.Disabled()' removed completely as it is unreachable when replaced with flag's 'True' case:
	doer.Do(userId, "Hello world!")
	doer.Do(userId, "Hello world again!")
}
