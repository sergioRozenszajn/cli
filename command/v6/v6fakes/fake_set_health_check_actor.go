// Code generated by counterfeiter. DO NOT EDIT.
package v6fakes

import (
	"sync"

	"code.cloudfoundry.org/cli/actor/v2action"
	"code.cloudfoundry.org/cli/api/cloudcontroller/ccv2/constant"
	v6 "code.cloudfoundry.org/cli/command/v6"
)

type FakeSetHealthCheckActor struct {
	SetApplicationHealthCheckTypeByNameAndSpaceStub        func(string, string, constant.ApplicationHealthCheckType, string) (v2action.Application, v2action.Warnings, error)
	setApplicationHealthCheckTypeByNameAndSpaceMutex       sync.RWMutex
	setApplicationHealthCheckTypeByNameAndSpaceArgsForCall []struct {
		arg1 string
		arg2 string
		arg3 constant.ApplicationHealthCheckType
		arg4 string
	}
	setApplicationHealthCheckTypeByNameAndSpaceReturns struct {
		result1 v2action.Application
		result2 v2action.Warnings
		result3 error
	}
	setApplicationHealthCheckTypeByNameAndSpaceReturnsOnCall map[int]struct {
		result1 v2action.Application
		result2 v2action.Warnings
		result3 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeSetHealthCheckActor) SetApplicationHealthCheckTypeByNameAndSpace(arg1 string, arg2 string, arg3 constant.ApplicationHealthCheckType, arg4 string) (v2action.Application, v2action.Warnings, error) {
	fake.setApplicationHealthCheckTypeByNameAndSpaceMutex.Lock()
	ret, specificReturn := fake.setApplicationHealthCheckTypeByNameAndSpaceReturnsOnCall[len(fake.setApplicationHealthCheckTypeByNameAndSpaceArgsForCall)]
	fake.setApplicationHealthCheckTypeByNameAndSpaceArgsForCall = append(fake.setApplicationHealthCheckTypeByNameAndSpaceArgsForCall, struct {
		arg1 string
		arg2 string
		arg3 constant.ApplicationHealthCheckType
		arg4 string
	}{arg1, arg2, arg3, arg4})
	fake.recordInvocation("SetApplicationHealthCheckTypeByNameAndSpace", []interface{}{arg1, arg2, arg3, arg4})
	fake.setApplicationHealthCheckTypeByNameAndSpaceMutex.Unlock()
	if fake.SetApplicationHealthCheckTypeByNameAndSpaceStub != nil {
		return fake.SetApplicationHealthCheckTypeByNameAndSpaceStub(arg1, arg2, arg3, arg4)
	}
	if specificReturn {
		return ret.result1, ret.result2, ret.result3
	}
	fakeReturns := fake.setApplicationHealthCheckTypeByNameAndSpaceReturns
	return fakeReturns.result1, fakeReturns.result2, fakeReturns.result3
}

func (fake *FakeSetHealthCheckActor) SetApplicationHealthCheckTypeByNameAndSpaceCallCount() int {
	fake.setApplicationHealthCheckTypeByNameAndSpaceMutex.RLock()
	defer fake.setApplicationHealthCheckTypeByNameAndSpaceMutex.RUnlock()
	return len(fake.setApplicationHealthCheckTypeByNameAndSpaceArgsForCall)
}

func (fake *FakeSetHealthCheckActor) SetApplicationHealthCheckTypeByNameAndSpaceCalls(stub func(string, string, constant.ApplicationHealthCheckType, string) (v2action.Application, v2action.Warnings, error)) {
	fake.setApplicationHealthCheckTypeByNameAndSpaceMutex.Lock()
	defer fake.setApplicationHealthCheckTypeByNameAndSpaceMutex.Unlock()
	fake.SetApplicationHealthCheckTypeByNameAndSpaceStub = stub
}

func (fake *FakeSetHealthCheckActor) SetApplicationHealthCheckTypeByNameAndSpaceArgsForCall(i int) (string, string, constant.ApplicationHealthCheckType, string) {
	fake.setApplicationHealthCheckTypeByNameAndSpaceMutex.RLock()
	defer fake.setApplicationHealthCheckTypeByNameAndSpaceMutex.RUnlock()
	argsForCall := fake.setApplicationHealthCheckTypeByNameAndSpaceArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3, argsForCall.arg4
}

func (fake *FakeSetHealthCheckActor) SetApplicationHealthCheckTypeByNameAndSpaceReturns(result1 v2action.Application, result2 v2action.Warnings, result3 error) {
	fake.setApplicationHealthCheckTypeByNameAndSpaceMutex.Lock()
	defer fake.setApplicationHealthCheckTypeByNameAndSpaceMutex.Unlock()
	fake.SetApplicationHealthCheckTypeByNameAndSpaceStub = nil
	fake.setApplicationHealthCheckTypeByNameAndSpaceReturns = struct {
		result1 v2action.Application
		result2 v2action.Warnings
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeSetHealthCheckActor) SetApplicationHealthCheckTypeByNameAndSpaceReturnsOnCall(i int, result1 v2action.Application, result2 v2action.Warnings, result3 error) {
	fake.setApplicationHealthCheckTypeByNameAndSpaceMutex.Lock()
	defer fake.setApplicationHealthCheckTypeByNameAndSpaceMutex.Unlock()
	fake.SetApplicationHealthCheckTypeByNameAndSpaceStub = nil
	if fake.setApplicationHealthCheckTypeByNameAndSpaceReturnsOnCall == nil {
		fake.setApplicationHealthCheckTypeByNameAndSpaceReturnsOnCall = make(map[int]struct {
			result1 v2action.Application
			result2 v2action.Warnings
			result3 error
		})
	}
	fake.setApplicationHealthCheckTypeByNameAndSpaceReturnsOnCall[i] = struct {
		result1 v2action.Application
		result2 v2action.Warnings
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeSetHealthCheckActor) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.setApplicationHealthCheckTypeByNameAndSpaceMutex.RLock()
	defer fake.setApplicationHealthCheckTypeByNameAndSpaceMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeSetHealthCheckActor) recordInvocation(key string, args []interface{}) {
	fake.invocationsMutex.Lock()
	defer fake.invocationsMutex.Unlock()
	if fake.invocations == nil {
		fake.invocations = map[string][][]interface{}{}
	}
	if fake.invocations[key] == nil {
		fake.invocations[key] = [][]interface{}{}
	}
	fake.invocations[key] = append(fake.invocations[key], args)
}

var _ v6.SetHealthCheckActor = new(FakeSetHealthCheckActor)
