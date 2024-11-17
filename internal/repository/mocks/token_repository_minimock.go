// Code generated by http://github.com/gojuno/minimock (v3.4.2). DO NOT EDIT.

package mocks

import (
	"context"
	"sync"
	mm_atomic "sync/atomic"
	mm_time "time"

	"github.com/gojuno/minimock/v3"
)

// TokenRepositoryMock implements repository.TokenRepository
type TokenRepositoryMock struct {
	t          minimock.Tester
	finishOnce sync.Once

	funcAddRevokedToken          func(ctx context.Context, refreshToken string) (err error)
	inspectFuncAddRevokedToken   func(ctx context.Context, refreshToken string)
	afterAddRevokedTokenCounter  uint64
	beforeAddRevokedTokenCounter uint64
	AddRevokedTokenMock          mTokenRepositoryMockAddRevokedToken

	funcIsTokenRevoked          func(ctx context.Context, refreshToken string) (b1 bool, err error)
	inspectFuncIsTokenRevoked   func(ctx context.Context, refreshToken string)
	afterIsTokenRevokedCounter  uint64
	beforeIsTokenRevokedCounter uint64
	IsTokenRevokedMock          mTokenRepositoryMockIsTokenRevoked
}

// NewTokenRepositoryMock returns a mock for repository.TokenRepository
func NewTokenRepositoryMock(t minimock.Tester) *TokenRepositoryMock {
	m := &TokenRepositoryMock{t: t}

	if controller, ok := t.(minimock.MockController); ok {
		controller.RegisterMocker(m)
	}

	m.AddRevokedTokenMock = mTokenRepositoryMockAddRevokedToken{mock: m}
	m.AddRevokedTokenMock.callArgs = []*TokenRepositoryMockAddRevokedTokenParams{}

	m.IsTokenRevokedMock = mTokenRepositoryMockIsTokenRevoked{mock: m}
	m.IsTokenRevokedMock.callArgs = []*TokenRepositoryMockIsTokenRevokedParams{}

	t.Cleanup(m.MinimockFinish)

	return m
}

type mTokenRepositoryMockAddRevokedToken struct {
	optional           bool
	mock               *TokenRepositoryMock
	defaultExpectation *TokenRepositoryMockAddRevokedTokenExpectation
	expectations       []*TokenRepositoryMockAddRevokedTokenExpectation

	callArgs []*TokenRepositoryMockAddRevokedTokenParams
	mutex    sync.RWMutex

	expectedInvocations uint64
}

// TokenRepositoryMockAddRevokedTokenExpectation specifies expectation struct of the TokenRepository.AddRevokedToken
type TokenRepositoryMockAddRevokedTokenExpectation struct {
	mock      *TokenRepositoryMock
	params    *TokenRepositoryMockAddRevokedTokenParams
	paramPtrs *TokenRepositoryMockAddRevokedTokenParamPtrs
	results   *TokenRepositoryMockAddRevokedTokenResults
	Counter   uint64
}

// TokenRepositoryMockAddRevokedTokenParams contains parameters of the TokenRepository.AddRevokedToken
type TokenRepositoryMockAddRevokedTokenParams struct {
	ctx          context.Context
	refreshToken string
}

// TokenRepositoryMockAddRevokedTokenParamPtrs contains pointers to parameters of the TokenRepository.AddRevokedToken
type TokenRepositoryMockAddRevokedTokenParamPtrs struct {
	ctx          *context.Context
	refreshToken *string
}

// TokenRepositoryMockAddRevokedTokenResults contains results of the TokenRepository.AddRevokedToken
type TokenRepositoryMockAddRevokedTokenResults struct {
	err error
}

// Marks this method to be optional. The default behavior of any method with Return() is '1 or more', meaning
// the test will fail minimock's automatic final call check if the mocked method was not called at least once.
// Optional() makes method check to work in '0 or more' mode.
// It is NOT RECOMMENDED to use this option unless you really need it, as default behaviour helps to
// catch the problems when the expected method call is totally skipped during test run.
func (mmAddRevokedToken *mTokenRepositoryMockAddRevokedToken) Optional() *mTokenRepositoryMockAddRevokedToken {
	mmAddRevokedToken.optional = true
	return mmAddRevokedToken
}

// Expect sets up expected params for TokenRepository.AddRevokedToken
func (mmAddRevokedToken *mTokenRepositoryMockAddRevokedToken) Expect(ctx context.Context, refreshToken string) *mTokenRepositoryMockAddRevokedToken {
	if mmAddRevokedToken.mock.funcAddRevokedToken != nil {
		mmAddRevokedToken.mock.t.Fatalf("TokenRepositoryMock.AddRevokedToken mock is already set by Set")
	}

	if mmAddRevokedToken.defaultExpectation == nil {
		mmAddRevokedToken.defaultExpectation = &TokenRepositoryMockAddRevokedTokenExpectation{}
	}

	if mmAddRevokedToken.defaultExpectation.paramPtrs != nil {
		mmAddRevokedToken.mock.t.Fatalf("TokenRepositoryMock.AddRevokedToken mock is already set by ExpectParams functions")
	}

	mmAddRevokedToken.defaultExpectation.params = &TokenRepositoryMockAddRevokedTokenParams{ctx, refreshToken}
	for _, e := range mmAddRevokedToken.expectations {
		if minimock.Equal(e.params, mmAddRevokedToken.defaultExpectation.params) {
			mmAddRevokedToken.mock.t.Fatalf("Expectation set by When has same params: %#v", *mmAddRevokedToken.defaultExpectation.params)
		}
	}

	return mmAddRevokedToken
}

// ExpectCtxParam1 sets up expected param ctx for TokenRepository.AddRevokedToken
func (mmAddRevokedToken *mTokenRepositoryMockAddRevokedToken) ExpectCtxParam1(ctx context.Context) *mTokenRepositoryMockAddRevokedToken {
	if mmAddRevokedToken.mock.funcAddRevokedToken != nil {
		mmAddRevokedToken.mock.t.Fatalf("TokenRepositoryMock.AddRevokedToken mock is already set by Set")
	}

	if mmAddRevokedToken.defaultExpectation == nil {
		mmAddRevokedToken.defaultExpectation = &TokenRepositoryMockAddRevokedTokenExpectation{}
	}

	if mmAddRevokedToken.defaultExpectation.params != nil {
		mmAddRevokedToken.mock.t.Fatalf("TokenRepositoryMock.AddRevokedToken mock is already set by Expect")
	}

	if mmAddRevokedToken.defaultExpectation.paramPtrs == nil {
		mmAddRevokedToken.defaultExpectation.paramPtrs = &TokenRepositoryMockAddRevokedTokenParamPtrs{}
	}
	mmAddRevokedToken.defaultExpectation.paramPtrs.ctx = &ctx

	return mmAddRevokedToken
}

// ExpectRefreshTokenParam2 sets up expected param refreshToken for TokenRepository.AddRevokedToken
func (mmAddRevokedToken *mTokenRepositoryMockAddRevokedToken) ExpectRefreshTokenParam2(refreshToken string) *mTokenRepositoryMockAddRevokedToken {
	if mmAddRevokedToken.mock.funcAddRevokedToken != nil {
		mmAddRevokedToken.mock.t.Fatalf("TokenRepositoryMock.AddRevokedToken mock is already set by Set")
	}

	if mmAddRevokedToken.defaultExpectation == nil {
		mmAddRevokedToken.defaultExpectation = &TokenRepositoryMockAddRevokedTokenExpectation{}
	}

	if mmAddRevokedToken.defaultExpectation.params != nil {
		mmAddRevokedToken.mock.t.Fatalf("TokenRepositoryMock.AddRevokedToken mock is already set by Expect")
	}

	if mmAddRevokedToken.defaultExpectation.paramPtrs == nil {
		mmAddRevokedToken.defaultExpectation.paramPtrs = &TokenRepositoryMockAddRevokedTokenParamPtrs{}
	}
	mmAddRevokedToken.defaultExpectation.paramPtrs.refreshToken = &refreshToken

	return mmAddRevokedToken
}

// Inspect accepts an inspector function that has same arguments as the TokenRepository.AddRevokedToken
func (mmAddRevokedToken *mTokenRepositoryMockAddRevokedToken) Inspect(f func(ctx context.Context, refreshToken string)) *mTokenRepositoryMockAddRevokedToken {
	if mmAddRevokedToken.mock.inspectFuncAddRevokedToken != nil {
		mmAddRevokedToken.mock.t.Fatalf("Inspect function is already set for TokenRepositoryMock.AddRevokedToken")
	}

	mmAddRevokedToken.mock.inspectFuncAddRevokedToken = f

	return mmAddRevokedToken
}

// Return sets up results that will be returned by TokenRepository.AddRevokedToken
func (mmAddRevokedToken *mTokenRepositoryMockAddRevokedToken) Return(err error) *TokenRepositoryMock {
	if mmAddRevokedToken.mock.funcAddRevokedToken != nil {
		mmAddRevokedToken.mock.t.Fatalf("TokenRepositoryMock.AddRevokedToken mock is already set by Set")
	}

	if mmAddRevokedToken.defaultExpectation == nil {
		mmAddRevokedToken.defaultExpectation = &TokenRepositoryMockAddRevokedTokenExpectation{mock: mmAddRevokedToken.mock}
	}
	mmAddRevokedToken.defaultExpectation.results = &TokenRepositoryMockAddRevokedTokenResults{err}
	return mmAddRevokedToken.mock
}

// Set uses given function f to mock the TokenRepository.AddRevokedToken method
func (mmAddRevokedToken *mTokenRepositoryMockAddRevokedToken) Set(f func(ctx context.Context, refreshToken string) (err error)) *TokenRepositoryMock {
	if mmAddRevokedToken.defaultExpectation != nil {
		mmAddRevokedToken.mock.t.Fatalf("Default expectation is already set for the TokenRepository.AddRevokedToken method")
	}

	if len(mmAddRevokedToken.expectations) > 0 {
		mmAddRevokedToken.mock.t.Fatalf("Some expectations are already set for the TokenRepository.AddRevokedToken method")
	}

	mmAddRevokedToken.mock.funcAddRevokedToken = f
	return mmAddRevokedToken.mock
}

// When sets expectation for the TokenRepository.AddRevokedToken which will trigger the result defined by the following
// Then helper
func (mmAddRevokedToken *mTokenRepositoryMockAddRevokedToken) When(ctx context.Context, refreshToken string) *TokenRepositoryMockAddRevokedTokenExpectation {
	if mmAddRevokedToken.mock.funcAddRevokedToken != nil {
		mmAddRevokedToken.mock.t.Fatalf("TokenRepositoryMock.AddRevokedToken mock is already set by Set")
	}

	expectation := &TokenRepositoryMockAddRevokedTokenExpectation{
		mock:   mmAddRevokedToken.mock,
		params: &TokenRepositoryMockAddRevokedTokenParams{ctx, refreshToken},
	}
	mmAddRevokedToken.expectations = append(mmAddRevokedToken.expectations, expectation)
	return expectation
}

// Then sets up TokenRepository.AddRevokedToken return parameters for the expectation previously defined by the When method
func (e *TokenRepositoryMockAddRevokedTokenExpectation) Then(err error) *TokenRepositoryMock {
	e.results = &TokenRepositoryMockAddRevokedTokenResults{err}
	return e.mock
}

// Times sets number of times TokenRepository.AddRevokedToken should be invoked
func (mmAddRevokedToken *mTokenRepositoryMockAddRevokedToken) Times(n uint64) *mTokenRepositoryMockAddRevokedToken {
	if n == 0 {
		mmAddRevokedToken.mock.t.Fatalf("Times of TokenRepositoryMock.AddRevokedToken mock can not be zero")
	}
	mm_atomic.StoreUint64(&mmAddRevokedToken.expectedInvocations, n)
	return mmAddRevokedToken
}

func (mmAddRevokedToken *mTokenRepositoryMockAddRevokedToken) invocationsDone() bool {
	if len(mmAddRevokedToken.expectations) == 0 && mmAddRevokedToken.defaultExpectation == nil && mmAddRevokedToken.mock.funcAddRevokedToken == nil {
		return true
	}

	totalInvocations := mm_atomic.LoadUint64(&mmAddRevokedToken.mock.afterAddRevokedTokenCounter)
	expectedInvocations := mm_atomic.LoadUint64(&mmAddRevokedToken.expectedInvocations)

	return totalInvocations > 0 && (expectedInvocations == 0 || expectedInvocations == totalInvocations)
}

// AddRevokedToken implements repository.TokenRepository
func (mmAddRevokedToken *TokenRepositoryMock) AddRevokedToken(ctx context.Context, refreshToken string) (err error) {
	mm_atomic.AddUint64(&mmAddRevokedToken.beforeAddRevokedTokenCounter, 1)
	defer mm_atomic.AddUint64(&mmAddRevokedToken.afterAddRevokedTokenCounter, 1)

	if mmAddRevokedToken.inspectFuncAddRevokedToken != nil {
		mmAddRevokedToken.inspectFuncAddRevokedToken(ctx, refreshToken)
	}

	mm_params := TokenRepositoryMockAddRevokedTokenParams{ctx, refreshToken}

	// Record call args
	mmAddRevokedToken.AddRevokedTokenMock.mutex.Lock()
	mmAddRevokedToken.AddRevokedTokenMock.callArgs = append(mmAddRevokedToken.AddRevokedTokenMock.callArgs, &mm_params)
	mmAddRevokedToken.AddRevokedTokenMock.mutex.Unlock()

	for _, e := range mmAddRevokedToken.AddRevokedTokenMock.expectations {
		if minimock.Equal(*e.params, mm_params) {
			mm_atomic.AddUint64(&e.Counter, 1)
			return e.results.err
		}
	}

	if mmAddRevokedToken.AddRevokedTokenMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmAddRevokedToken.AddRevokedTokenMock.defaultExpectation.Counter, 1)
		mm_want := mmAddRevokedToken.AddRevokedTokenMock.defaultExpectation.params
		mm_want_ptrs := mmAddRevokedToken.AddRevokedTokenMock.defaultExpectation.paramPtrs

		mm_got := TokenRepositoryMockAddRevokedTokenParams{ctx, refreshToken}

		if mm_want_ptrs != nil {

			if mm_want_ptrs.ctx != nil && !minimock.Equal(*mm_want_ptrs.ctx, mm_got.ctx) {
				mmAddRevokedToken.t.Errorf("TokenRepositoryMock.AddRevokedToken got unexpected parameter ctx, want: %#v, got: %#v%s\n", *mm_want_ptrs.ctx, mm_got.ctx, minimock.Diff(*mm_want_ptrs.ctx, mm_got.ctx))
			}

			if mm_want_ptrs.refreshToken != nil && !minimock.Equal(*mm_want_ptrs.refreshToken, mm_got.refreshToken) {
				mmAddRevokedToken.t.Errorf("TokenRepositoryMock.AddRevokedToken got unexpected parameter refreshToken, want: %#v, got: %#v%s\n", *mm_want_ptrs.refreshToken, mm_got.refreshToken, minimock.Diff(*mm_want_ptrs.refreshToken, mm_got.refreshToken))
			}

		} else if mm_want != nil && !minimock.Equal(*mm_want, mm_got) {
			mmAddRevokedToken.t.Errorf("TokenRepositoryMock.AddRevokedToken got unexpected parameters, want: %#v, got: %#v%s\n", *mm_want, mm_got, minimock.Diff(*mm_want, mm_got))
		}

		mm_results := mmAddRevokedToken.AddRevokedTokenMock.defaultExpectation.results
		if mm_results == nil {
			mmAddRevokedToken.t.Fatal("No results are set for the TokenRepositoryMock.AddRevokedToken")
		}
		return (*mm_results).err
	}
	if mmAddRevokedToken.funcAddRevokedToken != nil {
		return mmAddRevokedToken.funcAddRevokedToken(ctx, refreshToken)
	}
	mmAddRevokedToken.t.Fatalf("Unexpected call to TokenRepositoryMock.AddRevokedToken. %v %v", ctx, refreshToken)
	return
}

// AddRevokedTokenAfterCounter returns a count of finished TokenRepositoryMock.AddRevokedToken invocations
func (mmAddRevokedToken *TokenRepositoryMock) AddRevokedTokenAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmAddRevokedToken.afterAddRevokedTokenCounter)
}

// AddRevokedTokenBeforeCounter returns a count of TokenRepositoryMock.AddRevokedToken invocations
func (mmAddRevokedToken *TokenRepositoryMock) AddRevokedTokenBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmAddRevokedToken.beforeAddRevokedTokenCounter)
}

// Calls returns a list of arguments used in each call to TokenRepositoryMock.AddRevokedToken.
// The list is in the same order as the calls were made (i.e. recent calls have a higher index)
func (mmAddRevokedToken *mTokenRepositoryMockAddRevokedToken) Calls() []*TokenRepositoryMockAddRevokedTokenParams {
	mmAddRevokedToken.mutex.RLock()

	argCopy := make([]*TokenRepositoryMockAddRevokedTokenParams, len(mmAddRevokedToken.callArgs))
	copy(argCopy, mmAddRevokedToken.callArgs)

	mmAddRevokedToken.mutex.RUnlock()

	return argCopy
}

// MinimockAddRevokedTokenDone returns true if the count of the AddRevokedToken invocations corresponds
// the number of defined expectations
func (m *TokenRepositoryMock) MinimockAddRevokedTokenDone() bool {
	if m.AddRevokedTokenMock.optional {
		// Optional methods provide '0 or more' call count restriction.
		return true
	}

	for _, e := range m.AddRevokedTokenMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	return m.AddRevokedTokenMock.invocationsDone()
}

// MinimockAddRevokedTokenInspect logs each unmet expectation
func (m *TokenRepositoryMock) MinimockAddRevokedTokenInspect() {
	for _, e := range m.AddRevokedTokenMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Errorf("Expected call to TokenRepositoryMock.AddRevokedToken with params: %#v", *e.params)
		}
	}

	afterAddRevokedTokenCounter := mm_atomic.LoadUint64(&m.afterAddRevokedTokenCounter)
	// if default expectation was set then invocations count should be greater than zero
	if m.AddRevokedTokenMock.defaultExpectation != nil && afterAddRevokedTokenCounter < 1 {
		if m.AddRevokedTokenMock.defaultExpectation.params == nil {
			m.t.Error("Expected call to TokenRepositoryMock.AddRevokedToken")
		} else {
			m.t.Errorf("Expected call to TokenRepositoryMock.AddRevokedToken with params: %#v", *m.AddRevokedTokenMock.defaultExpectation.params)
		}
	}
	// if func was set then invocations count should be greater than zero
	if m.funcAddRevokedToken != nil && afterAddRevokedTokenCounter < 1 {
		m.t.Error("Expected call to TokenRepositoryMock.AddRevokedToken")
	}

	if !m.AddRevokedTokenMock.invocationsDone() && afterAddRevokedTokenCounter > 0 {
		m.t.Errorf("Expected %d calls to TokenRepositoryMock.AddRevokedToken but found %d calls",
			mm_atomic.LoadUint64(&m.AddRevokedTokenMock.expectedInvocations), afterAddRevokedTokenCounter)
	}
}

type mTokenRepositoryMockIsTokenRevoked struct {
	optional           bool
	mock               *TokenRepositoryMock
	defaultExpectation *TokenRepositoryMockIsTokenRevokedExpectation
	expectations       []*TokenRepositoryMockIsTokenRevokedExpectation

	callArgs []*TokenRepositoryMockIsTokenRevokedParams
	mutex    sync.RWMutex

	expectedInvocations uint64
}

// TokenRepositoryMockIsTokenRevokedExpectation specifies expectation struct of the TokenRepository.IsTokenRevoked
type TokenRepositoryMockIsTokenRevokedExpectation struct {
	mock      *TokenRepositoryMock
	params    *TokenRepositoryMockIsTokenRevokedParams
	paramPtrs *TokenRepositoryMockIsTokenRevokedParamPtrs
	results   *TokenRepositoryMockIsTokenRevokedResults
	Counter   uint64
}

// TokenRepositoryMockIsTokenRevokedParams contains parameters of the TokenRepository.IsTokenRevoked
type TokenRepositoryMockIsTokenRevokedParams struct {
	ctx          context.Context
	refreshToken string
}

// TokenRepositoryMockIsTokenRevokedParamPtrs contains pointers to parameters of the TokenRepository.IsTokenRevoked
type TokenRepositoryMockIsTokenRevokedParamPtrs struct {
	ctx          *context.Context
	refreshToken *string
}

// TokenRepositoryMockIsTokenRevokedResults contains results of the TokenRepository.IsTokenRevoked
type TokenRepositoryMockIsTokenRevokedResults struct {
	b1  bool
	err error
}

// Marks this method to be optional. The default behavior of any method with Return() is '1 or more', meaning
// the test will fail minimock's automatic final call check if the mocked method was not called at least once.
// Optional() makes method check to work in '0 or more' mode.
// It is NOT RECOMMENDED to use this option unless you really need it, as default behaviour helps to
// catch the problems when the expected method call is totally skipped during test run.
func (mmIsTokenRevoked *mTokenRepositoryMockIsTokenRevoked) Optional() *mTokenRepositoryMockIsTokenRevoked {
	mmIsTokenRevoked.optional = true
	return mmIsTokenRevoked
}

// Expect sets up expected params for TokenRepository.IsTokenRevoked
func (mmIsTokenRevoked *mTokenRepositoryMockIsTokenRevoked) Expect(ctx context.Context, refreshToken string) *mTokenRepositoryMockIsTokenRevoked {
	if mmIsTokenRevoked.mock.funcIsTokenRevoked != nil {
		mmIsTokenRevoked.mock.t.Fatalf("TokenRepositoryMock.IsTokenRevoked mock is already set by Set")
	}

	if mmIsTokenRevoked.defaultExpectation == nil {
		mmIsTokenRevoked.defaultExpectation = &TokenRepositoryMockIsTokenRevokedExpectation{}
	}

	if mmIsTokenRevoked.defaultExpectation.paramPtrs != nil {
		mmIsTokenRevoked.mock.t.Fatalf("TokenRepositoryMock.IsTokenRevoked mock is already set by ExpectParams functions")
	}

	mmIsTokenRevoked.defaultExpectation.params = &TokenRepositoryMockIsTokenRevokedParams{ctx, refreshToken}
	for _, e := range mmIsTokenRevoked.expectations {
		if minimock.Equal(e.params, mmIsTokenRevoked.defaultExpectation.params) {
			mmIsTokenRevoked.mock.t.Fatalf("Expectation set by When has same params: %#v", *mmIsTokenRevoked.defaultExpectation.params)
		}
	}

	return mmIsTokenRevoked
}

// ExpectCtxParam1 sets up expected param ctx for TokenRepository.IsTokenRevoked
func (mmIsTokenRevoked *mTokenRepositoryMockIsTokenRevoked) ExpectCtxParam1(ctx context.Context) *mTokenRepositoryMockIsTokenRevoked {
	if mmIsTokenRevoked.mock.funcIsTokenRevoked != nil {
		mmIsTokenRevoked.mock.t.Fatalf("TokenRepositoryMock.IsTokenRevoked mock is already set by Set")
	}

	if mmIsTokenRevoked.defaultExpectation == nil {
		mmIsTokenRevoked.defaultExpectation = &TokenRepositoryMockIsTokenRevokedExpectation{}
	}

	if mmIsTokenRevoked.defaultExpectation.params != nil {
		mmIsTokenRevoked.mock.t.Fatalf("TokenRepositoryMock.IsTokenRevoked mock is already set by Expect")
	}

	if mmIsTokenRevoked.defaultExpectation.paramPtrs == nil {
		mmIsTokenRevoked.defaultExpectation.paramPtrs = &TokenRepositoryMockIsTokenRevokedParamPtrs{}
	}
	mmIsTokenRevoked.defaultExpectation.paramPtrs.ctx = &ctx

	return mmIsTokenRevoked
}

// ExpectRefreshTokenParam2 sets up expected param refreshToken for TokenRepository.IsTokenRevoked
func (mmIsTokenRevoked *mTokenRepositoryMockIsTokenRevoked) ExpectRefreshTokenParam2(refreshToken string) *mTokenRepositoryMockIsTokenRevoked {
	if mmIsTokenRevoked.mock.funcIsTokenRevoked != nil {
		mmIsTokenRevoked.mock.t.Fatalf("TokenRepositoryMock.IsTokenRevoked mock is already set by Set")
	}

	if mmIsTokenRevoked.defaultExpectation == nil {
		mmIsTokenRevoked.defaultExpectation = &TokenRepositoryMockIsTokenRevokedExpectation{}
	}

	if mmIsTokenRevoked.defaultExpectation.params != nil {
		mmIsTokenRevoked.mock.t.Fatalf("TokenRepositoryMock.IsTokenRevoked mock is already set by Expect")
	}

	if mmIsTokenRevoked.defaultExpectation.paramPtrs == nil {
		mmIsTokenRevoked.defaultExpectation.paramPtrs = &TokenRepositoryMockIsTokenRevokedParamPtrs{}
	}
	mmIsTokenRevoked.defaultExpectation.paramPtrs.refreshToken = &refreshToken

	return mmIsTokenRevoked
}

// Inspect accepts an inspector function that has same arguments as the TokenRepository.IsTokenRevoked
func (mmIsTokenRevoked *mTokenRepositoryMockIsTokenRevoked) Inspect(f func(ctx context.Context, refreshToken string)) *mTokenRepositoryMockIsTokenRevoked {
	if mmIsTokenRevoked.mock.inspectFuncIsTokenRevoked != nil {
		mmIsTokenRevoked.mock.t.Fatalf("Inspect function is already set for TokenRepositoryMock.IsTokenRevoked")
	}

	mmIsTokenRevoked.mock.inspectFuncIsTokenRevoked = f

	return mmIsTokenRevoked
}

// Return sets up results that will be returned by TokenRepository.IsTokenRevoked
func (mmIsTokenRevoked *mTokenRepositoryMockIsTokenRevoked) Return(b1 bool, err error) *TokenRepositoryMock {
	if mmIsTokenRevoked.mock.funcIsTokenRevoked != nil {
		mmIsTokenRevoked.mock.t.Fatalf("TokenRepositoryMock.IsTokenRevoked mock is already set by Set")
	}

	if mmIsTokenRevoked.defaultExpectation == nil {
		mmIsTokenRevoked.defaultExpectation = &TokenRepositoryMockIsTokenRevokedExpectation{mock: mmIsTokenRevoked.mock}
	}
	mmIsTokenRevoked.defaultExpectation.results = &TokenRepositoryMockIsTokenRevokedResults{b1, err}
	return mmIsTokenRevoked.mock
}

// Set uses given function f to mock the TokenRepository.IsTokenRevoked method
func (mmIsTokenRevoked *mTokenRepositoryMockIsTokenRevoked) Set(f func(ctx context.Context, refreshToken string) (b1 bool, err error)) *TokenRepositoryMock {
	if mmIsTokenRevoked.defaultExpectation != nil {
		mmIsTokenRevoked.mock.t.Fatalf("Default expectation is already set for the TokenRepository.IsTokenRevoked method")
	}

	if len(mmIsTokenRevoked.expectations) > 0 {
		mmIsTokenRevoked.mock.t.Fatalf("Some expectations are already set for the TokenRepository.IsTokenRevoked method")
	}

	mmIsTokenRevoked.mock.funcIsTokenRevoked = f
	return mmIsTokenRevoked.mock
}

// When sets expectation for the TokenRepository.IsTokenRevoked which will trigger the result defined by the following
// Then helper
func (mmIsTokenRevoked *mTokenRepositoryMockIsTokenRevoked) When(ctx context.Context, refreshToken string) *TokenRepositoryMockIsTokenRevokedExpectation {
	if mmIsTokenRevoked.mock.funcIsTokenRevoked != nil {
		mmIsTokenRevoked.mock.t.Fatalf("TokenRepositoryMock.IsTokenRevoked mock is already set by Set")
	}

	expectation := &TokenRepositoryMockIsTokenRevokedExpectation{
		mock:   mmIsTokenRevoked.mock,
		params: &TokenRepositoryMockIsTokenRevokedParams{ctx, refreshToken},
	}
	mmIsTokenRevoked.expectations = append(mmIsTokenRevoked.expectations, expectation)
	return expectation
}

// Then sets up TokenRepository.IsTokenRevoked return parameters for the expectation previously defined by the When method
func (e *TokenRepositoryMockIsTokenRevokedExpectation) Then(b1 bool, err error) *TokenRepositoryMock {
	e.results = &TokenRepositoryMockIsTokenRevokedResults{b1, err}
	return e.mock
}

// Times sets number of times TokenRepository.IsTokenRevoked should be invoked
func (mmIsTokenRevoked *mTokenRepositoryMockIsTokenRevoked) Times(n uint64) *mTokenRepositoryMockIsTokenRevoked {
	if n == 0 {
		mmIsTokenRevoked.mock.t.Fatalf("Times of TokenRepositoryMock.IsTokenRevoked mock can not be zero")
	}
	mm_atomic.StoreUint64(&mmIsTokenRevoked.expectedInvocations, n)
	return mmIsTokenRevoked
}

func (mmIsTokenRevoked *mTokenRepositoryMockIsTokenRevoked) invocationsDone() bool {
	if len(mmIsTokenRevoked.expectations) == 0 && mmIsTokenRevoked.defaultExpectation == nil && mmIsTokenRevoked.mock.funcIsTokenRevoked == nil {
		return true
	}

	totalInvocations := mm_atomic.LoadUint64(&mmIsTokenRevoked.mock.afterIsTokenRevokedCounter)
	expectedInvocations := mm_atomic.LoadUint64(&mmIsTokenRevoked.expectedInvocations)

	return totalInvocations > 0 && (expectedInvocations == 0 || expectedInvocations == totalInvocations)
}

// IsTokenRevoked implements repository.TokenRepository
func (mmIsTokenRevoked *TokenRepositoryMock) IsTokenRevoked(ctx context.Context, refreshToken string) (b1 bool, err error) {
	mm_atomic.AddUint64(&mmIsTokenRevoked.beforeIsTokenRevokedCounter, 1)
	defer mm_atomic.AddUint64(&mmIsTokenRevoked.afterIsTokenRevokedCounter, 1)

	if mmIsTokenRevoked.inspectFuncIsTokenRevoked != nil {
		mmIsTokenRevoked.inspectFuncIsTokenRevoked(ctx, refreshToken)
	}

	mm_params := TokenRepositoryMockIsTokenRevokedParams{ctx, refreshToken}

	// Record call args
	mmIsTokenRevoked.IsTokenRevokedMock.mutex.Lock()
	mmIsTokenRevoked.IsTokenRevokedMock.callArgs = append(mmIsTokenRevoked.IsTokenRevokedMock.callArgs, &mm_params)
	mmIsTokenRevoked.IsTokenRevokedMock.mutex.Unlock()

	for _, e := range mmIsTokenRevoked.IsTokenRevokedMock.expectations {
		if minimock.Equal(*e.params, mm_params) {
			mm_atomic.AddUint64(&e.Counter, 1)
			return e.results.b1, e.results.err
		}
	}

	if mmIsTokenRevoked.IsTokenRevokedMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmIsTokenRevoked.IsTokenRevokedMock.defaultExpectation.Counter, 1)
		mm_want := mmIsTokenRevoked.IsTokenRevokedMock.defaultExpectation.params
		mm_want_ptrs := mmIsTokenRevoked.IsTokenRevokedMock.defaultExpectation.paramPtrs

		mm_got := TokenRepositoryMockIsTokenRevokedParams{ctx, refreshToken}

		if mm_want_ptrs != nil {

			if mm_want_ptrs.ctx != nil && !minimock.Equal(*mm_want_ptrs.ctx, mm_got.ctx) {
				mmIsTokenRevoked.t.Errorf("TokenRepositoryMock.IsTokenRevoked got unexpected parameter ctx, want: %#v, got: %#v%s\n", *mm_want_ptrs.ctx, mm_got.ctx, minimock.Diff(*mm_want_ptrs.ctx, mm_got.ctx))
			}

			if mm_want_ptrs.refreshToken != nil && !minimock.Equal(*mm_want_ptrs.refreshToken, mm_got.refreshToken) {
				mmIsTokenRevoked.t.Errorf("TokenRepositoryMock.IsTokenRevoked got unexpected parameter refreshToken, want: %#v, got: %#v%s\n", *mm_want_ptrs.refreshToken, mm_got.refreshToken, minimock.Diff(*mm_want_ptrs.refreshToken, mm_got.refreshToken))
			}

		} else if mm_want != nil && !minimock.Equal(*mm_want, mm_got) {
			mmIsTokenRevoked.t.Errorf("TokenRepositoryMock.IsTokenRevoked got unexpected parameters, want: %#v, got: %#v%s\n", *mm_want, mm_got, minimock.Diff(*mm_want, mm_got))
		}

		mm_results := mmIsTokenRevoked.IsTokenRevokedMock.defaultExpectation.results
		if mm_results == nil {
			mmIsTokenRevoked.t.Fatal("No results are set for the TokenRepositoryMock.IsTokenRevoked")
		}
		return (*mm_results).b1, (*mm_results).err
	}
	if mmIsTokenRevoked.funcIsTokenRevoked != nil {
		return mmIsTokenRevoked.funcIsTokenRevoked(ctx, refreshToken)
	}
	mmIsTokenRevoked.t.Fatalf("Unexpected call to TokenRepositoryMock.IsTokenRevoked. %v %v", ctx, refreshToken)
	return
}

// IsTokenRevokedAfterCounter returns a count of finished TokenRepositoryMock.IsTokenRevoked invocations
func (mmIsTokenRevoked *TokenRepositoryMock) IsTokenRevokedAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmIsTokenRevoked.afterIsTokenRevokedCounter)
}

// IsTokenRevokedBeforeCounter returns a count of TokenRepositoryMock.IsTokenRevoked invocations
func (mmIsTokenRevoked *TokenRepositoryMock) IsTokenRevokedBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmIsTokenRevoked.beforeIsTokenRevokedCounter)
}

// Calls returns a list of arguments used in each call to TokenRepositoryMock.IsTokenRevoked.
// The list is in the same order as the calls were made (i.e. recent calls have a higher index)
func (mmIsTokenRevoked *mTokenRepositoryMockIsTokenRevoked) Calls() []*TokenRepositoryMockIsTokenRevokedParams {
	mmIsTokenRevoked.mutex.RLock()

	argCopy := make([]*TokenRepositoryMockIsTokenRevokedParams, len(mmIsTokenRevoked.callArgs))
	copy(argCopy, mmIsTokenRevoked.callArgs)

	mmIsTokenRevoked.mutex.RUnlock()

	return argCopy
}

// MinimockIsTokenRevokedDone returns true if the count of the IsTokenRevoked invocations corresponds
// the number of defined expectations
func (m *TokenRepositoryMock) MinimockIsTokenRevokedDone() bool {
	if m.IsTokenRevokedMock.optional {
		// Optional methods provide '0 or more' call count restriction.
		return true
	}

	for _, e := range m.IsTokenRevokedMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	return m.IsTokenRevokedMock.invocationsDone()
}

// MinimockIsTokenRevokedInspect logs each unmet expectation
func (m *TokenRepositoryMock) MinimockIsTokenRevokedInspect() {
	for _, e := range m.IsTokenRevokedMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Errorf("Expected call to TokenRepositoryMock.IsTokenRevoked with params: %#v", *e.params)
		}
	}

	afterIsTokenRevokedCounter := mm_atomic.LoadUint64(&m.afterIsTokenRevokedCounter)
	// if default expectation was set then invocations count should be greater than zero
	if m.IsTokenRevokedMock.defaultExpectation != nil && afterIsTokenRevokedCounter < 1 {
		if m.IsTokenRevokedMock.defaultExpectation.params == nil {
			m.t.Error("Expected call to TokenRepositoryMock.IsTokenRevoked")
		} else {
			m.t.Errorf("Expected call to TokenRepositoryMock.IsTokenRevoked with params: %#v", *m.IsTokenRevokedMock.defaultExpectation.params)
		}
	}
	// if func was set then invocations count should be greater than zero
	if m.funcIsTokenRevoked != nil && afterIsTokenRevokedCounter < 1 {
		m.t.Error("Expected call to TokenRepositoryMock.IsTokenRevoked")
	}

	if !m.IsTokenRevokedMock.invocationsDone() && afterIsTokenRevokedCounter > 0 {
		m.t.Errorf("Expected %d calls to TokenRepositoryMock.IsTokenRevoked but found %d calls",
			mm_atomic.LoadUint64(&m.IsTokenRevokedMock.expectedInvocations), afterIsTokenRevokedCounter)
	}
}

// MinimockFinish checks that all mocked methods have been called the expected number of times
func (m *TokenRepositoryMock) MinimockFinish() {
	m.finishOnce.Do(func() {
		if !m.minimockDone() {
			m.MinimockAddRevokedTokenInspect()

			m.MinimockIsTokenRevokedInspect()
		}
	})
}

// MinimockWait waits for all mocked methods to be called the expected number of times
func (m *TokenRepositoryMock) MinimockWait(timeout mm_time.Duration) {
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

func (m *TokenRepositoryMock) minimockDone() bool {
	done := true
	return done &&
		m.MinimockAddRevokedTokenDone() &&
		m.MinimockIsTokenRevokedDone()
}
