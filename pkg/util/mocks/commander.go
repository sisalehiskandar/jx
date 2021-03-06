// Code generated by pegomock. DO NOT EDIT.
// Source: github.com/jenkins-x/jx/pkg/util (interfaces: Commander)

package util_test

import (
	backoff "github.com/cenkalti/backoff"
	pegomock "github.com/petergtz/pegomock"
	"reflect"
	"time"
)

type MockCommander struct {
	fail func(message string, callerSkip ...int)
}

func NewMockCommander() *MockCommander {
	return &MockCommander{fail: pegomock.GlobalFailHandler}
}

func (mock *MockCommander) CurrentArgs() []string {
	if mock == nil {
		panic("mock must not be nil. Use myMock := NewMockCommander().")
	}
	params := []pegomock.Param{}
	result := pegomock.GetGenericMockFrom(mock).Invoke("CurrentArgs", params, []reflect.Type{reflect.TypeOf((*[]string)(nil)).Elem()})
	var ret0 []string
	if len(result) != 0 {
		if result[0] != nil {
			ret0 = result[0].([]string)
		}
	}
	return ret0
}

func (mock *MockCommander) CurrentDir() string {
	if mock == nil {
		panic("mock must not be nil. Use myMock := NewMockCommander().")
	}
	params := []pegomock.Param{}
	result := pegomock.GetGenericMockFrom(mock).Invoke("CurrentDir", params, []reflect.Type{reflect.TypeOf((*string)(nil)).Elem()})
	var ret0 string
	if len(result) != 0 {
		if result[0] != nil {
			ret0 = result[0].(string)
		}
	}
	return ret0
}

func (mock *MockCommander) CurrentEnv() map[string]string {
	if mock == nil {
		panic("mock must not be nil. Use myMock := NewMockCommander().")
	}
	params := []pegomock.Param{}
	result := pegomock.GetGenericMockFrom(mock).Invoke("CurrentEnv", params, []reflect.Type{reflect.TypeOf((*map[string]string)(nil)).Elem()})
	var ret0 map[string]string
	if len(result) != 0 {
		if result[0] != nil {
			ret0 = result[0].(map[string]string)
		}
	}
	return ret0
}

func (mock *MockCommander) CurrentName() string {
	if mock == nil {
		panic("mock must not be nil. Use myMock := NewMockCommander().")
	}
	params := []pegomock.Param{}
	result := pegomock.GetGenericMockFrom(mock).Invoke("CurrentName", params, []reflect.Type{reflect.TypeOf((*string)(nil)).Elem()})
	var ret0 string
	if len(result) != 0 {
		if result[0] != nil {
			ret0 = result[0].(string)
		}
	}
	return ret0
}

func (mock *MockCommander) DidError() bool {
	if mock == nil {
		panic("mock must not be nil. Use myMock := NewMockCommander().")
	}
	params := []pegomock.Param{}
	result := pegomock.GetGenericMockFrom(mock).Invoke("DidError", params, []reflect.Type{reflect.TypeOf((*bool)(nil)).Elem()})
	var ret0 bool
	if len(result) != 0 {
		if result[0] != nil {
			ret0 = result[0].(bool)
		}
	}
	return ret0
}

func (mock *MockCommander) DidFail() bool {
	if mock == nil {
		panic("mock must not be nil. Use myMock := NewMockCommander().")
	}
	params := []pegomock.Param{}
	result := pegomock.GetGenericMockFrom(mock).Invoke("DidFail", params, []reflect.Type{reflect.TypeOf((*bool)(nil)).Elem()})
	var ret0 bool
	if len(result) != 0 {
		if result[0] != nil {
			ret0 = result[0].(bool)
		}
	}
	return ret0
}

func (mock *MockCommander) Error() error {
	if mock == nil {
		panic("mock must not be nil. Use myMock := NewMockCommander().")
	}
	params := []pegomock.Param{}
	result := pegomock.GetGenericMockFrom(mock).Invoke("Error", params, []reflect.Type{reflect.TypeOf((*error)(nil)).Elem()})
	var ret0 error
	if len(result) != 0 {
		if result[0] != nil {
			ret0 = result[0].(error)
		}
	}
	return ret0
}

func (mock *MockCommander) Run() (string, error) {
	if mock == nil {
		panic("mock must not be nil. Use myMock := NewMockCommander().")
	}
	params := []pegomock.Param{}
	result := pegomock.GetGenericMockFrom(mock).Invoke("Run", params, []reflect.Type{reflect.TypeOf((*string)(nil)).Elem(), reflect.TypeOf((*error)(nil)).Elem()})
	var ret0 string
	var ret1 error
	if len(result) != 0 {
		if result[0] != nil {
			ret0 = result[0].(string)
		}
		if result[1] != nil {
			ret1 = result[1].(error)
		}
	}
	return ret0, ret1
}

func (mock *MockCommander) RunWithoutRetry() (string, error) {
	if mock == nil {
		panic("mock must not be nil. Use myMock := NewMockCommander().")
	}
	params := []pegomock.Param{}
	result := pegomock.GetGenericMockFrom(mock).Invoke("RunWithoutRetry", params, []reflect.Type{reflect.TypeOf((*string)(nil)).Elem(), reflect.TypeOf((*error)(nil)).Elem()})
	var ret0 string
	var ret1 error
	if len(result) != 0 {
		if result[0] != nil {
			ret0 = result[0].(string)
		}
		if result[1] != nil {
			ret1 = result[1].(error)
		}
	}
	return ret0, ret1
}

func (mock *MockCommander) SetArgs(_param0 []string) {
	if mock == nil {
		panic("mock must not be nil. Use myMock := NewMockCommander().")
	}
	params := []pegomock.Param{_param0}
	pegomock.GetGenericMockFrom(mock).Invoke("SetArgs", params, []reflect.Type{})
}

func (mock *MockCommander) SetDir(_param0 string) {
	if mock == nil {
		panic("mock must not be nil. Use myMock := NewMockCommander().")
	}
	params := []pegomock.Param{_param0}
	pegomock.GetGenericMockFrom(mock).Invoke("SetDir", params, []reflect.Type{})
}

func (mock *MockCommander) SetEnv(_param0 map[string]string) {
	if mock == nil {
		panic("mock must not be nil. Use myMock := NewMockCommander().")
	}
	params := []pegomock.Param{_param0}
	pegomock.GetGenericMockFrom(mock).Invoke("SetEnv", params, []reflect.Type{})
}

func (mock *MockCommander) SetEnvVariable(_param0 string, _param1 string) {
	if mock == nil {
		panic("mock must not be nil. Use myMock := NewMockCommander().")
	}
	params := []pegomock.Param{_param0, _param1}
	pegomock.GetGenericMockFrom(mock).Invoke("SetEnvVariable", params, []reflect.Type{})
}

func (mock *MockCommander) SetExponentialBackOff(_param0 *backoff.ExponentialBackOff) {
	if mock == nil {
		panic("mock must not be nil. Use myMock := NewMockCommander().")
	}
	params := []pegomock.Param{_param0}
	pegomock.GetGenericMockFrom(mock).Invoke("SetExponentialBackOff", params, []reflect.Type{})
}

func (mock *MockCommander) SetName(_param0 string) {
	if mock == nil {
		panic("mock must not be nil. Use myMock := NewMockCommander().")
	}
	params := []pegomock.Param{_param0}
	pegomock.GetGenericMockFrom(mock).Invoke("SetName", params, []reflect.Type{})
}

func (mock *MockCommander) SetTimeout(_param0 time.Duration) {
	if mock == nil {
		panic("mock must not be nil. Use myMock := NewMockCommander().")
	}
	params := []pegomock.Param{_param0}
	pegomock.GetGenericMockFrom(mock).Invoke("SetTimeout", params, []reflect.Type{})
}

func (mock *MockCommander) VerifyWasCalledOnce() *VerifierCommander {
	return &VerifierCommander{
		mock:                   mock,
		invocationCountMatcher: pegomock.Times(1),
	}
}

func (mock *MockCommander) VerifyWasCalled(invocationCountMatcher pegomock.Matcher) *VerifierCommander {
	return &VerifierCommander{
		mock:                   mock,
		invocationCountMatcher: invocationCountMatcher,
	}
}

func (mock *MockCommander) VerifyWasCalledInOrder(invocationCountMatcher pegomock.Matcher, inOrderContext *pegomock.InOrderContext) *VerifierCommander {
	return &VerifierCommander{
		mock:                   mock,
		invocationCountMatcher: invocationCountMatcher,
		inOrderContext:         inOrderContext,
	}
}

func (mock *MockCommander) VerifyWasCalledEventually(invocationCountMatcher pegomock.Matcher, timeout time.Duration) *VerifierCommander {
	return &VerifierCommander{
		mock:                   mock,
		invocationCountMatcher: invocationCountMatcher,
		timeout:                timeout,
	}
}

type VerifierCommander struct {
	mock                   *MockCommander
	invocationCountMatcher pegomock.Matcher
	inOrderContext         *pegomock.InOrderContext
	timeout                time.Duration
}

func (verifier *VerifierCommander) CurrentArgs() *Commander_CurrentArgs_OngoingVerification {
	params := []pegomock.Param{}
	methodInvocations := pegomock.GetGenericMockFrom(verifier.mock).Verify(verifier.inOrderContext, verifier.invocationCountMatcher, "CurrentArgs", params, verifier.timeout)
	return &Commander_CurrentArgs_OngoingVerification{mock: verifier.mock, methodInvocations: methodInvocations}
}

type Commander_CurrentArgs_OngoingVerification struct {
	mock              *MockCommander
	methodInvocations []pegomock.MethodInvocation
}

func (c *Commander_CurrentArgs_OngoingVerification) GetCapturedArguments() {
}

func (c *Commander_CurrentArgs_OngoingVerification) GetAllCapturedArguments() {
}

func (verifier *VerifierCommander) CurrentDir() *Commander_CurrentDir_OngoingVerification {
	params := []pegomock.Param{}
	methodInvocations := pegomock.GetGenericMockFrom(verifier.mock).Verify(verifier.inOrderContext, verifier.invocationCountMatcher, "CurrentDir", params, verifier.timeout)
	return &Commander_CurrentDir_OngoingVerification{mock: verifier.mock, methodInvocations: methodInvocations}
}

type Commander_CurrentDir_OngoingVerification struct {
	mock              *MockCommander
	methodInvocations []pegomock.MethodInvocation
}

func (c *Commander_CurrentDir_OngoingVerification) GetCapturedArguments() {
}

func (c *Commander_CurrentDir_OngoingVerification) GetAllCapturedArguments() {
}

func (verifier *VerifierCommander) CurrentEnv() *Commander_CurrentEnv_OngoingVerification {
	params := []pegomock.Param{}
	methodInvocations := pegomock.GetGenericMockFrom(verifier.mock).Verify(verifier.inOrderContext, verifier.invocationCountMatcher, "CurrentEnv", params, verifier.timeout)
	return &Commander_CurrentEnv_OngoingVerification{mock: verifier.mock, methodInvocations: methodInvocations}
}

type Commander_CurrentEnv_OngoingVerification struct {
	mock              *MockCommander
	methodInvocations []pegomock.MethodInvocation
}

func (c *Commander_CurrentEnv_OngoingVerification) GetCapturedArguments() {
}

func (c *Commander_CurrentEnv_OngoingVerification) GetAllCapturedArguments() {
}

func (verifier *VerifierCommander) CurrentName() *Commander_CurrentName_OngoingVerification {
	params := []pegomock.Param{}
	methodInvocations := pegomock.GetGenericMockFrom(verifier.mock).Verify(verifier.inOrderContext, verifier.invocationCountMatcher, "CurrentName", params, verifier.timeout)
	return &Commander_CurrentName_OngoingVerification{mock: verifier.mock, methodInvocations: methodInvocations}
}

type Commander_CurrentName_OngoingVerification struct {
	mock              *MockCommander
	methodInvocations []pegomock.MethodInvocation
}

func (c *Commander_CurrentName_OngoingVerification) GetCapturedArguments() {
}

func (c *Commander_CurrentName_OngoingVerification) GetAllCapturedArguments() {
}

func (verifier *VerifierCommander) DidError() *Commander_DidError_OngoingVerification {
	params := []pegomock.Param{}
	methodInvocations := pegomock.GetGenericMockFrom(verifier.mock).Verify(verifier.inOrderContext, verifier.invocationCountMatcher, "DidError", params, verifier.timeout)
	return &Commander_DidError_OngoingVerification{mock: verifier.mock, methodInvocations: methodInvocations}
}

type Commander_DidError_OngoingVerification struct {
	mock              *MockCommander
	methodInvocations []pegomock.MethodInvocation
}

func (c *Commander_DidError_OngoingVerification) GetCapturedArguments() {
}

func (c *Commander_DidError_OngoingVerification) GetAllCapturedArguments() {
}

func (verifier *VerifierCommander) DidFail() *Commander_DidFail_OngoingVerification {
	params := []pegomock.Param{}
	methodInvocations := pegomock.GetGenericMockFrom(verifier.mock).Verify(verifier.inOrderContext, verifier.invocationCountMatcher, "DidFail", params, verifier.timeout)
	return &Commander_DidFail_OngoingVerification{mock: verifier.mock, methodInvocations: methodInvocations}
}

type Commander_DidFail_OngoingVerification struct {
	mock              *MockCommander
	methodInvocations []pegomock.MethodInvocation
}

func (c *Commander_DidFail_OngoingVerification) GetCapturedArguments() {
}

func (c *Commander_DidFail_OngoingVerification) GetAllCapturedArguments() {
}

func (verifier *VerifierCommander) Error() *Commander_Error_OngoingVerification {
	params := []pegomock.Param{}
	methodInvocations := pegomock.GetGenericMockFrom(verifier.mock).Verify(verifier.inOrderContext, verifier.invocationCountMatcher, "Error", params, verifier.timeout)
	return &Commander_Error_OngoingVerification{mock: verifier.mock, methodInvocations: methodInvocations}
}

type Commander_Error_OngoingVerification struct {
	mock              *MockCommander
	methodInvocations []pegomock.MethodInvocation
}

func (c *Commander_Error_OngoingVerification) GetCapturedArguments() {
}

func (c *Commander_Error_OngoingVerification) GetAllCapturedArguments() {
}

func (verifier *VerifierCommander) Run() *Commander_Run_OngoingVerification {
	params := []pegomock.Param{}
	methodInvocations := pegomock.GetGenericMockFrom(verifier.mock).Verify(verifier.inOrderContext, verifier.invocationCountMatcher, "Run", params, verifier.timeout)
	return &Commander_Run_OngoingVerification{mock: verifier.mock, methodInvocations: methodInvocations}
}

type Commander_Run_OngoingVerification struct {
	mock              *MockCommander
	methodInvocations []pegomock.MethodInvocation
}

func (c *Commander_Run_OngoingVerification) GetCapturedArguments() {
}

func (c *Commander_Run_OngoingVerification) GetAllCapturedArguments() {
}

func (verifier *VerifierCommander) RunWithoutRetry() *Commander_RunWithoutRetry_OngoingVerification {
	params := []pegomock.Param{}
	methodInvocations := pegomock.GetGenericMockFrom(verifier.mock).Verify(verifier.inOrderContext, verifier.invocationCountMatcher, "RunWithoutRetry", params, verifier.timeout)
	return &Commander_RunWithoutRetry_OngoingVerification{mock: verifier.mock, methodInvocations: methodInvocations}
}

type Commander_RunWithoutRetry_OngoingVerification struct {
	mock              *MockCommander
	methodInvocations []pegomock.MethodInvocation
}

func (c *Commander_RunWithoutRetry_OngoingVerification) GetCapturedArguments() {
}

func (c *Commander_RunWithoutRetry_OngoingVerification) GetAllCapturedArguments() {
}

func (verifier *VerifierCommander) SetArgs(_param0 []string) *Commander_SetArgs_OngoingVerification {
	params := []pegomock.Param{_param0}
	methodInvocations := pegomock.GetGenericMockFrom(verifier.mock).Verify(verifier.inOrderContext, verifier.invocationCountMatcher, "SetArgs", params, verifier.timeout)
	return &Commander_SetArgs_OngoingVerification{mock: verifier.mock, methodInvocations: methodInvocations}
}

type Commander_SetArgs_OngoingVerification struct {
	mock              *MockCommander
	methodInvocations []pegomock.MethodInvocation
}

func (c *Commander_SetArgs_OngoingVerification) GetCapturedArguments() []string {
	_param0 := c.GetAllCapturedArguments()
	return _param0[len(_param0)-1]
}

func (c *Commander_SetArgs_OngoingVerification) GetAllCapturedArguments() (_param0 [][]string) {
	params := pegomock.GetGenericMockFrom(c.mock).GetInvocationParams(c.methodInvocations)
	if len(params) > 0 {
		_param0 = make([][]string, len(params[0]))
		for u, param := range params[0] {
			_param0[u] = param.([]string)
		}
	}
	return
}

func (verifier *VerifierCommander) SetDir(_param0 string) *Commander_SetDir_OngoingVerification {
	params := []pegomock.Param{_param0}
	methodInvocations := pegomock.GetGenericMockFrom(verifier.mock).Verify(verifier.inOrderContext, verifier.invocationCountMatcher, "SetDir", params, verifier.timeout)
	return &Commander_SetDir_OngoingVerification{mock: verifier.mock, methodInvocations: methodInvocations}
}

type Commander_SetDir_OngoingVerification struct {
	mock              *MockCommander
	methodInvocations []pegomock.MethodInvocation
}

func (c *Commander_SetDir_OngoingVerification) GetCapturedArguments() string {
	_param0 := c.GetAllCapturedArguments()
	return _param0[len(_param0)-1]
}

func (c *Commander_SetDir_OngoingVerification) GetAllCapturedArguments() (_param0 []string) {
	params := pegomock.GetGenericMockFrom(c.mock).GetInvocationParams(c.methodInvocations)
	if len(params) > 0 {
		_param0 = make([]string, len(params[0]))
		for u, param := range params[0] {
			_param0[u] = param.(string)
		}
	}
	return
}

func (verifier *VerifierCommander) SetEnv(_param0 map[string]string) *Commander_SetEnv_OngoingVerification {
	params := []pegomock.Param{_param0}
	methodInvocations := pegomock.GetGenericMockFrom(verifier.mock).Verify(verifier.inOrderContext, verifier.invocationCountMatcher, "SetEnv", params, verifier.timeout)
	return &Commander_SetEnv_OngoingVerification{mock: verifier.mock, methodInvocations: methodInvocations}
}

type Commander_SetEnv_OngoingVerification struct {
	mock              *MockCommander
	methodInvocations []pegomock.MethodInvocation
}

func (c *Commander_SetEnv_OngoingVerification) GetCapturedArguments() map[string]string {
	_param0 := c.GetAllCapturedArguments()
	return _param0[len(_param0)-1]
}

func (c *Commander_SetEnv_OngoingVerification) GetAllCapturedArguments() (_param0 []map[string]string) {
	params := pegomock.GetGenericMockFrom(c.mock).GetInvocationParams(c.methodInvocations)
	if len(params) > 0 {
		_param0 = make([]map[string]string, len(params[0]))
		for u, param := range params[0] {
			_param0[u] = param.(map[string]string)
		}
	}
	return
}

func (verifier *VerifierCommander) SetEnvVariable(_param0 string, _param1 string) *Commander_SetEnvVariable_OngoingVerification {
	params := []pegomock.Param{_param0, _param1}
	methodInvocations := pegomock.GetGenericMockFrom(verifier.mock).Verify(verifier.inOrderContext, verifier.invocationCountMatcher, "SetEnvVariable", params, verifier.timeout)
	return &Commander_SetEnvVariable_OngoingVerification{mock: verifier.mock, methodInvocations: methodInvocations}
}

type Commander_SetEnvVariable_OngoingVerification struct {
	mock              *MockCommander
	methodInvocations []pegomock.MethodInvocation
}

func (c *Commander_SetEnvVariable_OngoingVerification) GetCapturedArguments() (string, string) {
	_param0, _param1 := c.GetAllCapturedArguments()
	return _param0[len(_param0)-1], _param1[len(_param1)-1]
}

func (c *Commander_SetEnvVariable_OngoingVerification) GetAllCapturedArguments() (_param0 []string, _param1 []string) {
	params := pegomock.GetGenericMockFrom(c.mock).GetInvocationParams(c.methodInvocations)
	if len(params) > 0 {
		_param0 = make([]string, len(params[0]))
		for u, param := range params[0] {
			_param0[u] = param.(string)
		}
		_param1 = make([]string, len(params[1]))
		for u, param := range params[1] {
			_param1[u] = param.(string)
		}
	}
	return
}

func (verifier *VerifierCommander) SetExponentialBackOff(_param0 *backoff.ExponentialBackOff) *Commander_SetExponentialBackOff_OngoingVerification {
	params := []pegomock.Param{_param0}
	methodInvocations := pegomock.GetGenericMockFrom(verifier.mock).Verify(verifier.inOrderContext, verifier.invocationCountMatcher, "SetExponentialBackOff", params, verifier.timeout)
	return &Commander_SetExponentialBackOff_OngoingVerification{mock: verifier.mock, methodInvocations: methodInvocations}
}

type Commander_SetExponentialBackOff_OngoingVerification struct {
	mock              *MockCommander
	methodInvocations []pegomock.MethodInvocation
}

func (c *Commander_SetExponentialBackOff_OngoingVerification) GetCapturedArguments() *backoff.ExponentialBackOff {
	_param0 := c.GetAllCapturedArguments()
	return _param0[len(_param0)-1]
}

func (c *Commander_SetExponentialBackOff_OngoingVerification) GetAllCapturedArguments() (_param0 []*backoff.ExponentialBackOff) {
	params := pegomock.GetGenericMockFrom(c.mock).GetInvocationParams(c.methodInvocations)
	if len(params) > 0 {
		_param0 = make([]*backoff.ExponentialBackOff, len(params[0]))
		for u, param := range params[0] {
			_param0[u] = param.(*backoff.ExponentialBackOff)
		}
	}
	return
}

func (verifier *VerifierCommander) SetName(_param0 string) *Commander_SetName_OngoingVerification {
	params := []pegomock.Param{_param0}
	methodInvocations := pegomock.GetGenericMockFrom(verifier.mock).Verify(verifier.inOrderContext, verifier.invocationCountMatcher, "SetName", params, verifier.timeout)
	return &Commander_SetName_OngoingVerification{mock: verifier.mock, methodInvocations: methodInvocations}
}

type Commander_SetName_OngoingVerification struct {
	mock              *MockCommander
	methodInvocations []pegomock.MethodInvocation
}

func (c *Commander_SetName_OngoingVerification) GetCapturedArguments() string {
	_param0 := c.GetAllCapturedArguments()
	return _param0[len(_param0)-1]
}

func (c *Commander_SetName_OngoingVerification) GetAllCapturedArguments() (_param0 []string) {
	params := pegomock.GetGenericMockFrom(c.mock).GetInvocationParams(c.methodInvocations)
	if len(params) > 0 {
		_param0 = make([]string, len(params[0]))
		for u, param := range params[0] {
			_param0[u] = param.(string)
		}
	}
	return
}

func (verifier *VerifierCommander) SetTimeout(_param0 time.Duration) *Commander_SetTimeout_OngoingVerification {
	params := []pegomock.Param{_param0}
	methodInvocations := pegomock.GetGenericMockFrom(verifier.mock).Verify(verifier.inOrderContext, verifier.invocationCountMatcher, "SetTimeout", params, verifier.timeout)
	return &Commander_SetTimeout_OngoingVerification{mock: verifier.mock, methodInvocations: methodInvocations}
}

type Commander_SetTimeout_OngoingVerification struct {
	mock              *MockCommander
	methodInvocations []pegomock.MethodInvocation
}

func (c *Commander_SetTimeout_OngoingVerification) GetCapturedArguments() time.Duration {
	_param0 := c.GetAllCapturedArguments()
	return _param0[len(_param0)-1]
}

func (c *Commander_SetTimeout_OngoingVerification) GetAllCapturedArguments() (_param0 []time.Duration) {
	params := pegomock.GetGenericMockFrom(c.mock).GetInvocationParams(c.methodInvocations)
	if len(params) > 0 {
		_param0 = make([]time.Duration, len(params[0]))
		for u, param := range params[0] {
			_param0[u] = param.(time.Duration)
		}
	}
	return
}
