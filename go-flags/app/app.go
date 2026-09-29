package app

import (
	"go-flags/doit"
	"go-flags/fm"

	roxcontext "github.com/rollout/rox-go/v6/core/context"
)

type Doer interface {
	Do(arg, message string)
}

func Run() {
	userId := "me@example.org"

	var doer Doer
	// Get the value of the 'EnableNewFeature' Boolean flag using the static API:
	if fm.Flags.EnableNewFeature.IsEnabled(nil) {
		doer = doit.NewEnhancedDoer()
	} else {
		doer = &doit.SimpleDoer{}
	}
	doer.Do(userId, "Hello world!")
	doer.Do(userId, "Hello world again!")

	// Get the value of the 'EnableNewFeature' Boolean flag using the static API and using a context of property values:
	rCtx := roxcontext.NewContext(map[string]interface{}{
		"userId": userId,
	})
	if fm.Flags.EnableNewFeature.IsEnabled(rCtx) {
		doer = doit.NewEnhancedDoer()
	} else {
		doer = &doit.SimpleDoer{}
	}
	doer.Do(userId, "Hello world!")
	doer.Do(userId, "Hello world again!")

	// Get the value of the 'EnableNewFeature' Boolean flag using the static API, and the alternate 'Enabled' function to execute a function only if the flag evaluates to True:
	fm.Flags.EnableNewFeature.Enabled(nil, func() {
		// This code will only run if the flag is enabled.
		doer = doit.NewEnhancedDoer()

	})

	// Get the value of the 'EnableNewFeature' Boolean flag using the static API, and the alternate 'Disabled' function to execute a function only if the flag evaluates to False:
	fm.Flags.EnableNewFeature.Disabled(nil, func() {
		// This code will only run if the flag is disabled.
		doer = &doit.SimpleDoer{}
	})
	doer.Do(userId, "Hello world!")
	doer.Do(userId, "Hello world again!")
}
