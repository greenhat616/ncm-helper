package commander

import (
	flag "github.com/spf13/pflag"
)
// Parse params
func Parse () {
	// Register Flag mapping
	registerVersionFlag()
	registerHelpFlag()
	registerDebugFlag()
	registerConfigFlag()

	// Parse Flag
	flag.Parse()

	// Handle Flag event
	handleVersionFlag()
	handleHelpFlag()
}
