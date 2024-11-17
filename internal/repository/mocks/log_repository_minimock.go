// Code generated by http://github.com/gojuno/minimock (v3.4.2). DO NOT EDIT.

package mocks

import (
	"context"
	"sync"
	mm_atomic "sync/atomic"
	mm_time "time"

	"github.com/8thgencore/microservice-auth/internal/model"
	"github.com/gojuno/minimock/v3"
)

// LogRepositoryMock implements repository.LogRepository
type LogRepositoryMock struct {
	t          minimock.Tester
	finishOnce sync.Once

	funcLog          func(ctx context.Context, log *model.Log) (err error)
	inspectFuncLog   func(ctx context.Context, log *model.Log)
	afterLogCounter  uint64
	beforeLogCounter uint64
	LogMock          mLogRepositoryMockLog
}

// NewLogRepositoryMock returns a mock for repository.LogRepository
func NewLogRepositoryMock(t minimock.Tester) *LogRepositoryMock {
	m := &LogRepositoryMock{t: t}

	if controller, ok := t.(minimock.MockController); ok {
		controller.RegisterMocker(m)
	}

	m.LogMock = mLogRepositoryMockLog{mock: m}
	m.LogMock.callArgs = []*LogRepositoryMockLogParams{}

	t.Cleanup(m.MinimockFinish)

	return m
}

type mLogRepositoryMockLog struct {
	optional           bool
	mock               *LogRepositoryMock
	defaultExpectation *LogRepositoryMockLogExpectation
	expectations       []*LogRepositoryMockLogExpectation

	callArgs []*LogRepositoryMockLogParams
	mutex    sync.RWMutex

	expectedInvocations uint64
}

// LogRepositoryMockLogExpectation specifies expectation struct of the LogRepository.Log
type LogRepositoryMockLogExpectation struct {
	mock      *LogRepositoryMock
	params    *LogRepositoryMockLogParams
	paramPtrs *LogRepositoryMockLogParamPtrs
	results   *LogRepositoryMockLogResults
	Counter   uint64
}

// LogRepositoryMockLogParams contains parameters of the LogRepository.Log
type LogRepositoryMockLogParams struct {
	ctx context.Context
	log *model.Log
}

// LogRepositoryMockLogParamPtrs contains pointers to parameters of the LogRepository.Log
type LogRepositoryMockLogParamPtrs struct {
	ctx *context.Context
	log **model.Log
}

// LogRepositoryMockLogResults contains results of the LogRepository.Log
type LogRepositoryMockLogResults struct {
	err error
}

// Marks this method to be optional. The default behavior of any method with Return() is '1 or more', meaning
// the test will fail minimock's automatic final call check if the mocked method was not called at least once.
// Optional() makes method check to work in '0 or more' mode.
// It is NOT RECOMMENDED to use this option unless you really need it, as default behaviour helps to
// catch the problems when the expected method call is totally skipped during test run.
func (mmLog *mLogRepositoryMockLog) Optional() *mLogRepositoryMockLog {
	mmLog.optional = true
	return mmLog
}

// Expect sets up expected params for LogRepository.Log
func (mmLog *mLogRepositoryMockLog) Expect(ctx context.Context, log *model.Log) *mLogRepositoryMockLog {
	if mmLog.mock.funcLog != nil {
		mmLog.mock.t.Fatalf("LogRepositoryMock.Log mock is already set by Set")
	}

	if mmLog.defaultExpectation == nil {
		mmLog.defaultExpectation = &LogRepositoryMockLogExpectation{}
	}

	if mmLog.defaultExpectation.paramPtrs != nil {
		mmLog.mock.t.Fatalf("LogRepositoryMock.Log mock is already set by ExpectParams functions")
	}

	mmLog.defaultExpectation.params = &LogRepositoryMockLogParams{ctx, log}
	for _, e := range mmLog.expectations {
		if minimock.Equal(e.params, mmLog.defaultExpectation.params) {
			mmLog.mock.t.Fatalf("Expectation set by When has same params: %#v", *mmLog.defaultExpectation.params)
		}
	}

	return mmLog
}

// ExpectCtxParam1 sets up expected param ctx for LogRepository.Log
func (mmLog *mLogRepositoryMockLog) ExpectCtxParam1(ctx context.Context) *mLogRepositoryMockLog {
	if mmLog.mock.funcLog != nil {
		mmLog.mock.t.Fatalf("LogRepositoryMock.Log mock is already set by Set")
	}

	if mmLog.defaultExpectation == nil {
		mmLog.defaultExpectation = &LogRepositoryMockLogExpectation{}
	}

	if mmLog.defaultExpectation.params != nil {
		mmLog.mock.t.Fatalf("LogRepositoryMock.Log mock is already set by Expect")
	}

	if mmLog.defaultExpectation.paramPtrs == nil {
		mmLog.defaultExpectation.paramPtrs = &LogRepositoryMockLogParamPtrs{}
	}
	mmLog.defaultExpectation.paramPtrs.ctx = &ctx

	return mmLog
}

// ExpectLogParam2 sets up expected param log for LogRepository.Log
func (mmLog *mLogRepositoryMockLog) ExpectLogParam2(log *model.Log) *mLogRepositoryMockLog {
	if mmLog.mock.funcLog != nil {
		mmLog.mock.t.Fatalf("LogRepositoryMock.Log mock is already set by Set")
	}

	if mmLog.defaultExpectation == nil {
		mmLog.defaultExpectation = &LogRepositoryMockLogExpectation{}
	}

	if mmLog.defaultExpectation.params != nil {
		mmLog.mock.t.Fatalf("LogRepositoryMock.Log mock is already set by Expect")
	}

	if mmLog.defaultExpectation.paramPtrs == nil {
		mmLog.defaultExpectation.paramPtrs = &LogRepositoryMockLogParamPtrs{}
	}
	mmLog.defaultExpectation.paramPtrs.log = &log

	return mmLog
}

// Inspect accepts an inspector function that has same arguments as the LogRepository.Log
func (mmLog *mLogRepositoryMockLog) Inspect(f func(ctx context.Context, log *model.Log)) *mLogRepositoryMockLog {
	if mmLog.mock.inspectFuncLog != nil {
		mmLog.mock.t.Fatalf("Inspect function is already set for LogRepositoryMock.Log")
	}

	mmLog.mock.inspectFuncLog = f

	return mmLog
}

// Return sets up results that will be returned by LogRepository.Log
func (mmLog *mLogRepositoryMockLog) Return(err error) *LogRepositoryMock {
	if mmLog.mock.funcLog != nil {
		mmLog.mock.t.Fatalf("LogRepositoryMock.Log mock is already set by Set")
	}

	if mmLog.defaultExpectation == nil {
		mmLog.defaultExpectation = &LogRepositoryMockLogExpectation{mock: mmLog.mock}
	}
	mmLog.defaultExpectation.results = &LogRepositoryMockLogResults{err}
	return mmLog.mock
}

// Set uses given function f to mock the LogRepository.Log method
func (mmLog *mLogRepositoryMockLog) Set(f func(ctx context.Context, log *model.Log) (err error)) *LogRepositoryMock {
	if mmLog.defaultExpectation != nil {
		mmLog.mock.t.Fatalf("Default expectation is already set for the LogRepository.Log method")
	}

	if len(mmLog.expectations) > 0 {
		mmLog.mock.t.Fatalf("Some expectations are already set for the LogRepository.Log method")
	}

	mmLog.mock.funcLog = f
	return mmLog.mock
}

// When sets expectation for the LogRepository.Log which will trigger the result defined by the following
// Then helper
func (mmLog *mLogRepositoryMockLog) When(ctx context.Context, log *model.Log) *LogRepositoryMockLogExpectation {
	if mmLog.mock.funcLog != nil {
		mmLog.mock.t.Fatalf("LogRepositoryMock.Log mock is already set by Set")
	}

	expectation := &LogRepositoryMockLogExpectation{
		mock:   mmLog.mock,
		params: &LogRepositoryMockLogParams{ctx, log},
	}
	mmLog.expectations = append(mmLog.expectations, expectation)
	return expectation
}

// Then sets up LogRepository.Log return parameters for the expectation previously defined by the When method
func (e *LogRepositoryMockLogExpectation) Then(err error) *LogRepositoryMock {
	e.results = &LogRepositoryMockLogResults{err}
	return e.mock
}

// Times sets number of times LogRepository.Log should be invoked
func (mmLog *mLogRepositoryMockLog) Times(n uint64) *mLogRepositoryMockLog {
	if n == 0 {
		mmLog.mock.t.Fatalf("Times of LogRepositoryMock.Log mock can not be zero")
	}
	mm_atomic.StoreUint64(&mmLog.expectedInvocations, n)
	return mmLog
}

func (mmLog *mLogRepositoryMockLog) invocationsDone() bool {
	if len(mmLog.expectations) == 0 && mmLog.defaultExpectation == nil && mmLog.mock.funcLog == nil {
		return true
	}

	totalInvocations := mm_atomic.LoadUint64(&mmLog.mock.afterLogCounter)
	expectedInvocations := mm_atomic.LoadUint64(&mmLog.expectedInvocations)

	return totalInvocations > 0 && (expectedInvocations == 0 || expectedInvocations == totalInvocations)
}

// Log implements repository.LogRepository
func (mmLog *LogRepositoryMock) Log(ctx context.Context, log *model.Log) (err error) {
	mm_atomic.AddUint64(&mmLog.beforeLogCounter, 1)
	defer mm_atomic.AddUint64(&mmLog.afterLogCounter, 1)

	if mmLog.inspectFuncLog != nil {
		mmLog.inspectFuncLog(ctx, log)
	}

	mm_params := LogRepositoryMockLogParams{ctx, log}

	// Record call args
	mmLog.LogMock.mutex.Lock()
	mmLog.LogMock.callArgs = append(mmLog.LogMock.callArgs, &mm_params)
	mmLog.LogMock.mutex.Unlock()

	for _, e := range mmLog.LogMock.expectations {
		if minimock.Equal(*e.params, mm_params) {
			mm_atomic.AddUint64(&e.Counter, 1)
			return e.results.err
		}
	}

	if mmLog.LogMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmLog.LogMock.defaultExpectation.Counter, 1)
		mm_want := mmLog.LogMock.defaultExpectation.params
		mm_want_ptrs := mmLog.LogMock.defaultExpectation.paramPtrs

		mm_got := LogRepositoryMockLogParams{ctx, log}

		if mm_want_ptrs != nil {

			if mm_want_ptrs.ctx != nil && !minimock.Equal(*mm_want_ptrs.ctx, mm_got.ctx) {
				mmLog.t.Errorf("LogRepositoryMock.Log got unexpected parameter ctx, want: %#v, got: %#v%s\n", *mm_want_ptrs.ctx, mm_got.ctx, minimock.Diff(*mm_want_ptrs.ctx, mm_got.ctx))
			}

			if mm_want_ptrs.log != nil && !minimock.Equal(*mm_want_ptrs.log, mm_got.log) {
				mmLog.t.Errorf("LogRepositoryMock.Log got unexpected parameter log, want: %#v, got: %#v%s\n", *mm_want_ptrs.log, mm_got.log, minimock.Diff(*mm_want_ptrs.log, mm_got.log))
			}

		} else if mm_want != nil && !minimock.Equal(*mm_want, mm_got) {
			mmLog.t.Errorf("LogRepositoryMock.Log got unexpected parameters, want: %#v, got: %#v%s\n", *mm_want, mm_got, minimock.Diff(*mm_want, mm_got))
		}

		mm_results := mmLog.LogMock.defaultExpectation.results
		if mm_results == nil {
			mmLog.t.Fatal("No results are set for the LogRepositoryMock.Log")
		}
		return (*mm_results).err
	}
	if mmLog.funcLog != nil {
		return mmLog.funcLog(ctx, log)
	}
	mmLog.t.Fatalf("Unexpected call to LogRepositoryMock.Log. %v %v", ctx, log)
	return
}

// LogAfterCounter returns a count of finished LogRepositoryMock.Log invocations
func (mmLog *LogRepositoryMock) LogAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmLog.afterLogCounter)
}

// LogBeforeCounter returns a count of LogRepositoryMock.Log invocations
func (mmLog *LogRepositoryMock) LogBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmLog.beforeLogCounter)
}

// Calls returns a list of arguments used in each call to LogRepositoryMock.Log.
// The list is in the same order as the calls were made (i.e. recent calls have a higher index)
func (mmLog *mLogRepositoryMockLog) Calls() []*LogRepositoryMockLogParams {
	mmLog.mutex.RLock()

	argCopy := make([]*LogRepositoryMockLogParams, len(mmLog.callArgs))
	copy(argCopy, mmLog.callArgs)

	mmLog.mutex.RUnlock()

	return argCopy
}

// MinimockLogDone returns true if the count of the Log invocations corresponds
// the number of defined expectations
func (m *LogRepositoryMock) MinimockLogDone() bool {
	if m.LogMock.optional {
		// Optional methods provide '0 or more' call count restriction.
		return true
	}

	for _, e := range m.LogMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	return m.LogMock.invocationsDone()
}

// MinimockLogInspect logs each unmet expectation
func (m *LogRepositoryMock) MinimockLogInspect() {
	for _, e := range m.LogMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Errorf("Expected call to LogRepositoryMock.Log with params: %#v", *e.params)
		}
	}

	afterLogCounter := mm_atomic.LoadUint64(&m.afterLogCounter)
	// if default expectation was set then invocations count should be greater than zero
	if m.LogMock.defaultExpectation != nil && afterLogCounter < 1 {
		if m.LogMock.defaultExpectation.params == nil {
			m.t.Error("Expected call to LogRepositoryMock.Log")
		} else {
			m.t.Errorf("Expected call to LogRepositoryMock.Log with params: %#v", *m.LogMock.defaultExpectation.params)
		}
	}
	// if func was set then invocations count should be greater than zero
	if m.funcLog != nil && afterLogCounter < 1 {
		m.t.Error("Expected call to LogRepositoryMock.Log")
	}

	if !m.LogMock.invocationsDone() && afterLogCounter > 0 {
		m.t.Errorf("Expected %d calls to LogRepositoryMock.Log but found %d calls",
			mm_atomic.LoadUint64(&m.LogMock.expectedInvocations), afterLogCounter)
	}
}

// MinimockFinish checks that all mocked methods have been called the expected number of times
func (m *LogRepositoryMock) MinimockFinish() {
	m.finishOnce.Do(func() {
		if !m.minimockDone() {
			m.MinimockLogInspect()
		}
	})
}

// MinimockWait waits for all mocked methods to be called the expected number of times
func (m *LogRepositoryMock) MinimockWait(timeout mm_time.Duration) {
	timeoutCh := mm_time.After(timeout)
	for {
		if m.minimockDone() {
			return
		}
		select {
		case <-timeoutCh:
			m.MinimockFinish()
			return
		case <-mm_time.After(10 * mm_time.Millisecond):
		}
	}
}

func (m *LogRepositoryMock) minimockDone() bool {
	done := true
	return done &&
		m.MinimockLogDone()
}
