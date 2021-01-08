package godog

import "github.com/tankyouoss/godog/internal/flags"

// Options are suite run options
// flags are mapped to these options.
//
// It can also be used together with godog.RunWithOptions
// to run test suite from go source directly
//
// See the flags for more details
type Options = flags.Options
