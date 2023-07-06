package test

import (
	"fmt"
	"math"
	"time"

	"github.com/kubeshop/tracetest/server/environment"
	"github.com/kubeshop/tracetest/server/model"
	"github.com/kubeshop/tracetest/server/pkg/id"
	"github.com/kubeshop/tracetest/server/pkg/maps"
	"github.com/kubeshop/tracetest/server/test/trigger"
)

var (
	Now = func() time.Time {
		return time.Now().UTC()
	}

	IDGen = id.NewRandGenerator()
)

func NewRun() Run {
	return Run{
		ID:        0,
		TraceID:   IDGen.TraceID(),
		SpanID:    IDGen.SpanID(),
		State:     RunStateCreated,
		CreatedAt: Now(),
	}
}

func (r Run) ResourceID() string {
	return fmt.Sprintf("test/%s/run/%d", r.TestID, r.ID)
}

func (r Run) TransactionStepResourceID() string {
	return fmt.Sprintf("transaction_step/%s/run/%d", r.TestID, r.ID)
}

func (r Run) Copy() Run {
	r.ID = 0
	r.Results = nil
	r.CreatedAt = Now()

	return r
}

func (r Run) ExecutionTime() int {
	return durationInSeconds(
		timeDiff(r.CreatedAt, r.CompletedAt),
	)
}

func (r Run) TriggerTime() int {
	return durationInMillieconds(
		timeDiff(r.ServiceTriggeredAt, r.ServiceTriggerCompletedAt),
	)
}

func timeDiff(start, end time.Time) time.Duration {
	var endDate time.Time
	if !dateIsZero(end) {
		endDate = end
	} else {
		endDate = Now()
	}
	return endDate.Sub(start)
}

func durationInMillieconds(d time.Duration) int {
	return int(d.Milliseconds())
}

func durationInSeconds(d time.Duration) int {
	return int(math.Ceil(d.Seconds()))
}

func dateIsZero(in time.Time) bool {
	return in.IsZero() || in.Unix() == 0
}

func (r Run) Start() Run {
	r.State = RunStateExecuting
	r.ServiceTriggeredAt = Now()
	return r
}

func (r Run) TriggerCompleted(tr trigger.TriggerResult) Run {
	r.ServiceTriggerCompletedAt = Now()
	r.TriggerResult = tr
	return r
}

func (r Run) SuccessfullyTriggered() Run {
	r.State = RunStateAwaitingTrace
	return r
}

func (r Run) SuccessfullyPolledTraces(t *model.Trace) Run {
	r.State = RunStateAnalyzingTrace
	r.Trace = t
	r.ObtainedTraceAt = time.Now()
	return r
}

func (r Run) SuccessfullyAsserted(
	outputs maps.Ordered[string, RunOutput],
	environment environment.Environment,
	res maps.Ordered[SpanQuery, []AssertionResult],
	allPassed bool,
) Run {
	r.Outputs = outputs
	r.Environment = environment
	r.Results = &RunResults{
		AllPassed: allPassed,
		Results:   res,
	}
	r.State = RunStateFinished
	return r.Finish()
}

func (r Run) Finish() Run {
	r.CompletedAt = Now()
	return r
}

func (r Run) Stopped() Run {
	r.State = RunStateStopped
	return r.Finish()
}

func (r Run) TriggerFailed(err error) Run {
	r.State = RunStateTriggerFailed
	r.LastError = err
	return r.Finish()
}

func (r Run) TraceFailed(err error) Run {
	r.State = RunStateTraceFailed
	r.LastError = err
	return r.Finish()
}

func (r Run) AssertionFailed(err error) Run {
	r.State = RunStateAssertionFailed
	r.LastError = err
	return r.Finish()
}

func (r Run) LinterError(err error) Run {
	r.State = RunStateAnalyzingError
	r.LastError = err
	return r.Finish()
}

func (r Run) SuccessfulLinterExecution(linter model.LinterResult) Run {
	r.State = RunStateAwaitingTestResults
	r.Linter = linter

	return r
}

func NewTracetestRootSpan(run Run) model.Span {
	return model.AugmentRootSpan(model.Span{
		ID:         id.NewRandGenerator().SpanID(),
		Name:       model.TriggerSpanName,
		StartTime:  run.ServiceTriggeredAt,
		EndTime:    run.ServiceTriggerCompletedAt,
		Attributes: model.Attributes{},
		Children:   []*model.Span{},
	}, run.TriggerResult)
}