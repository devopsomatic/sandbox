package v1

import (
	"github.com/stackfoundation/core/pkg/workflows/errors"
	"github.com/stackfoundation/core/pkg/workflows/properties"
)

func expandStepOptions(step *StepOptions, variables *properties.Properties) error {
	name, err := variables.Expand(step.Name)
	step.Name = name
	return err
}

func expandRunStepOptions(run *RunStepOptions, variables *properties.Properties) error {
	composite := errors.NewCompositeError()

	composite.Append(expandScriptStepOptions(&run.ScriptStepOptions, variables))

	cache, err := variables.Expand(run.Cache)
	run.Cache = cache
	composite.Append(err)

	parallel, err := variables.Expand(run.Parallel)
	run.Parallel = parallel
	composite.Append(err)

	return composite.OrNilIfEmpty()
}

// ExpandStep Expand any placeholders in this step that haven't been expanded yet
func ExpandStep(step *WorkflowStep, variables *properties.Properties) error {
	if step.Run != nil {
		return expandRunStepOptions(step.Run, variables)
	} else if step.Generator != nil {
		return expandGeneratorStepOptions(step.Generator, variables)
	} else if step.External != nil {
		return expandExternalStepOptions(step.External, variables)
	} else if step.Service != nil {
		return expandServiceStepOptions(step.Service, variables)
	}

	return nil
}
