package doit

import (
	"fmt"

	"go-flags/fm"

	"github.com/mitchellh/colorstring"
	roxcontext "github.com/rollout/rox-go/v6/core/context"
)

type SimpleDoer struct{}

func (s *SimpleDoer) Do(arg, message string) {
	fmt.Printf("Hello %s; you are using the old feature.\n", arg)
	fmt.Printf("The message is: %s.\n", message)
}

type enhancedDoer struct {
	messageHistory []int
}

func NewEnhancedDoer() *enhancedDoer {
	// Get the value of the 'MessageHistorySize' integer number flag, using the dynamic API:
	messageHistorySize := fm.DynamicAPI().GetInt(fm.FlagName_MessageHistorySize, 3, []int{}, nil)
	fmt.Printf("Previous %d messages will be displayed.\n", messageHistorySize)
	return &enhancedDoer{
		messageHistory: []int{},
	}
}

func (e *enhancedDoer) Do(arg, message string) {
	// Get the value of the 'ux.TextColour' string flag using the static API:
	textColour := fm.UxFlags.TextColour.GetValue(roxcontext.NewContext(map[string]interface{}{
		"userId": arg,
	}))
	if _, err := colorstring.Printf("["+textColour+"]Hello %s; you are using the new feature.[reset]\n", arg); err != nil {
		// fallback to boring text without colour.
		fmt.Printf("Hello %s; you are using the new feature.\n", arg)
	}
	fmt.Printf("The message is: %s.\n", message)
	// Get the value of the 'ux.MessageHistorySize' integer number flag using the static API:
	messageHistorySize := fm.UxFlags.MessageHistorySize.GetValue(nil)
	fmt.Printf("Previous %d messages:\n", messageHistorySize)
	for i := 0; i < messageHistorySize && i < len(e.messageHistory); i++ {
		fmt.Printf("\tMessage %d: %s\n", i, e.messageHistory[i])
	}
	// Get the value of the 'ux.ThresholdValue' floating-point flag using the static API:
	thresholdValue := fm.UxFlags.ThresholdValue.GetValue(nil)
	fmt.Printf("Threshold: %.2f\n", thresholdValue)
}
