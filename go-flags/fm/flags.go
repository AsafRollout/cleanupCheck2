package fm

import (
	"fmt"
	"os"

	"go-flags/fmsync"

	"github.com/rollout/rox-go/v6/core/model"
	"github.com/rollout/rox-go/v6/server"
)

// The fully qualified name of the MessageHistorySize flag, for use with the dynamic API:
const FlagName_MessageHistorySize = "ux.MessageHistorySize"

// Type for 'Flags' feature flag container.
type flags struct {
	// EnableNewFeature controls whether the new feature is enabled for the current user.
	EnableNewFeature server.RoxFlag
	rox              *server.Rox
}

// Type for 'UxFlags' feature flag container.
type uxFlags struct {
	TextColour         server.RoxString
	MessageHistorySize server.RoxInt
	ThresholdValue     server.RoxDouble
}

// Type for 'SvcFlags' feature flag container. Registered through the
// CloudBees `fmsync` wrapper (not a raw Rox.Register call) — the shape
// used by every calculi-corp Go service.
type svcFlags struct {
	// ApplyStrictMode is namespaced 'go-flags' via the fmsync wrapper.
	// Its fully-qualified key is 'go-flags.ApplyStrictMode', but every
	// reference in source uses the bare field name — the namespace prefix
	// exists only at registration time.
	ApplyStrictMode server.RoxFlag
}

// Flags container that will be registered that will be registered without a namespace.
var Flags = &flags{
	EnableNewFeature: server.NewRoxFlag(false),
	rox:              server.NewRox(),
}

// UxFlags container that will be registered under the 'ux' namespace.
var UxFlags = &uxFlags{
	TextColour:         server.NewRoxString("red", []string{"red", "green", "blue"}),
	MessageHistorySize: server.NewRoxInt(3, []int{}),
	ThresholdValue:     server.NewRoxDouble(1.5, []float64{0.5, 1.5, 2.5}),
}

// SvcFlags is registered under the 'go-flags' namespace via the fmsync
// wrapper in InitService below.
var SvcFlags = &svcFlags{
	ApplyStrictMode: server.NewRoxFlag(false),
}

func Init() {
	sdkKey := os.Getenv("FM_SDK_KEY")
	if sdkKey != "" {
		Flags.rox = server.NewRox()
		// Register Flags without a namespace; flag names will not be prefixed.
		Flags.rox.RegisterWithEmptyNamespace(Flags)
		// Register UxFlags under the 'ux' namespace; flag names will be prefixed 'ux.'.
		Flags.rox.Register("ux", UxFlags)
		options := server.NewRoxOptions(server.RoxOptionsBuilder{})
		// Initialise the CloudBees Unify feature management SDK.
		<-Flags.rox.Setup(sdkKey, options)
	} else {
		fmt.Println("FM_SDK_KEY environment variable not set. Feature flags not initialised.")
	}
}

// InitService registers SvcFlags under the 'go-flags' namespace through
// the fmsync wrapper — the standard calculi-corp registration path. A
// qualified key 'go-flags.ApplyStrictMode' must resolve here even though
// no literal Rox.Register("go-flags", …) call exists.
func InitService() {
	fmsync.NewFeatureFlags("go-flags", map[string]any{}, SvcFlags)
}

// ApplyStrictMode reports the current value of the namespaced flag. The
// reference is the bare field name, never the fully-qualified key.
func ApplyStrictMode() bool {
	return SvcFlags.ApplyStrictMode.IsEnabled(nil)
}

func DynamicAPI() model.DynamicAPI {
	return Flags.rox.DynamicAPI()
}
