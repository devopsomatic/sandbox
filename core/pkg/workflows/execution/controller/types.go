package controller

import (
	"context"

	executioncontext "github.com/stackfoundation/sandbox/core/pkg/workflows/execution/context"
	"github.com/stackfoundation/sandbox/core/pkg/workflows/execution/coordinator"
	"github.com/stackfoundation/sandbox/core/pkg/workflows/v1"
)

// Controller A controller used to execute workflows
type Controller interface {
	Execute(context context.Context, workflow *v1.Workflow)
}

type executionController struct {
	coordinator        coordinator.Coordinator
	pendingTransitions chan pendingTransition
}

type pendingTransition struct {
	context    *executioncontext.StepContext
	transition func(*executioncontext.StepContext)
}

type runListener struct {
	controller *executionController
}
