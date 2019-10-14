package exec

import (
	"context"

	"github.com/concourse/concourse/atc/tracing"
)

// OnSuccessStep will run one step, and then a second step if the first step
// succeeds.
type OnSuccessStep struct {
	step Step
	hook Step
}

// OnSuccess constructs an OnSuccessStep factory.
func OnSuccess(firstStep Step, secondStep Step) Step {
	return OnSuccessStep{
		step: firstStep,
		hook: secondStep,
	}
}

// Run will call Run on the first step and wait for it to complete. If the
// first step errors, Run returns the error. OnSuccessStep is ready as soon as
// the first step is ready.
//
// If the first step succeeds (that is, its Success result is true), the second
// step is executed. If the second step errors, its error is returned.
func (o OnSuccessStep) Run(ctx context.Context, state RunState) error {
	// [cc] wrap the step in a span
	//
	ctx, span := tracing.StartSpan(ctx, "on-success", nil)
	defer span.End()

	stepRunErr := o.step.Run(ctx, state)
	if stepRunErr != nil {
		return stepRunErr
	}

	success := o.step.Succeeded()
	if !success {
		return nil
	}

	return o.hook.Run(ctx, state)
}

// Succeeded is true if the first step completed and the second
// step completed successfully.
func (o OnSuccessStep) Succeeded() bool {
	return o.step.Succeeded() && o.hook.Succeeded()
}
