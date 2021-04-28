package models

import (
	"time"

	"github.com/elmagician/godog/colors"
	"github.com/elmagician/godog/internal/utils"
)

// TestRunStarted ...
type TestRunStarted struct {
	StartedAt time.Time
}

// PickleResult ...
type PickleResult struct {
	PickleID  string
	StartedAt time.Time
}

// PickleStepResult ...
type PickleStepResult struct {
	Status     StepResultStatus
	FinishedAt time.Time
	Err        error

	PickleID     string
	PickleStepID string

	Def *StepDefinition
}

// NewStepResult ...
func NewStepResult(pickleID, pickleStepID string, match *StepDefinition) PickleStepResult {
	return PickleStepResult{FinishedAt: utils.TimeNowFunc(), PickleID: pickleID, PickleStepID: pickleStepID, Def: match}
}

// StepResultStatus ...
type StepResultStatus int

const (
	// Passed ...
	Passed StepResultStatus = iota
	// Failed ...
	Failed
	// Skipped ...
	Skipped
	// Undefined ...
	Undefined
	// Pending ...
	Pending
)

// Color ...
func (st StepResultStatus) Color() colors.ColorFunc {
	switch st {
	case Passed:
		return colors.Green
	case Failed:
		return colors.Red
	case Skipped:
		return colors.Cyan
	default:
		return colors.Yellow
	}
}

// String ...
func (st StepResultStatus) String() string {
	switch st {
	case Passed:
		return "passed"
	case Failed:
		return "failed"
	case Skipped:
		return "skipped"
	case Undefined:
		return "undefined"
	case Pending:
		return "pending"
	default:
		return "unknown"
	}
}
