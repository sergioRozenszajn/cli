// Code generated by counterfeiter. DO NOT EDIT.
package v7fakes

import (
	"sync"

	"code.cloudfoundry.org/cli/actor/v7action"
	v7 "code.cloudfoundry.org/cli/command/v7"
)

type FakeRoutesActor struct {
	GetRoutesBySpaceStub        func(string) ([]v7action.Route, v7action.Warnings, error)
	getRoutesBySpaceMutex       sync.RWMutex
	getRoutesBySpaceArgsForCall []struct {
		arg1 string
	}
	getRoutesBySpaceReturns struct {
		result1 []v7action.Route
		result2 v7action.Warnings
		result3 error
	}
	getRoutesBySpaceReturnsOnCall map[int]struct {
		result1 []v7action.Route
		result2 v7action.Warnings
		result3 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeRoutesActor) GetRoutesBySpace(arg1 string) ([]v7action.Route, v7action.Warnings, error) {
	fake.getRoutesBySpaceMutex.Lock()
	ret, specificReturn := fake.getRoutesBySpaceReturnsOnCall[len(fake.getRoutesBySpaceArgsForCall)]
	fake.getRoutesBySpaceArgsForCall = append(fake.getRoutesBySpaceArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.recordInvocation("GetRoutesBySpace", []interface{}{arg1})
	fake.getRoutesBySpaceMutex.Unlock()
	if fake.GetRoutesBySpaceStub != nil {
		return fake.GetRoutesBySpaceStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2, ret.result3
	}
	fakeReturns := fake.getRoutesBySpaceReturns
	return fakeReturns.result1, fakeReturns.result2, fakeReturns.result3
}

func (fake *FakeRoutesActor) GetRoutesBySpaceCallCount() int {
	fake.getRoutesBySpaceMutex.RLock()
	defer fake.getRoutesBySpaceMutex.RUnlock()
	return len(fake.getRoutesBySpaceArgsForCall)
}

func (fake *FakeRoutesActor) GetRoutesBySpaceCalls(stub func(string) ([]v7action.Route, v7action.Warnings, error)) {
	fake.getRoutesBySpaceMutex.Lock()
	defer fake.getRoutesBySpaceMutex.Unlock()
	fake.GetRoutesBySpaceStub = stub
}

func (fake *FakeRoutesActor) GetRoutesBySpaceArgsForCall(i int) string {
	fake.getRoutesBySpaceMutex.RLock()
	defer fake.getRoutesBySpaceMutex.RUnlock()
	argsForCall := fake.getRoutesBySpaceArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeRoutesActor) GetRoutesBySpaceReturns(result1 []v7action.Route, result2 v7action.Warnings, result3 error) {
	fake.getRoutesBySpaceMutex.Lock()
	defer fake.getRoutesBySpaceMutex.Unlock()
	fake.GetRoutesBySpaceStub = nil
	fake.getRoutesBySpaceReturns = struct {
		result1 []v7action.Route
		result2 v7action.Warnings
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeRoutesActor) GetRoutesBySpaceReturnsOnCall(i int, result1 []v7action.Route, result2 v7action.Warnings, result3 error) {
	fake.getRoutesBySpaceMutex.Lock()
	defer fake.getRoutesBySpaceMutex.Unlock()
	fake.GetRoutesBySpaceStub = nil
	if fake.getRoutesBySpaceReturnsOnCall == nil {
		fake.getRoutesBySpaceReturnsOnCall = make(map[int]struct {
			result1 []v7action.Route
			result2 v7action.Warnings
			result3 error
		})
	}
	fake.getRoutesBySpaceReturnsOnCall[i] = struct {
		result1 []v7action.Route
		result2 v7action.Warnings
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeRoutesActor) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.getRoutesBySpaceMutex.RLock()
	defer fake.getRoutesBySpaceMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeRoutesActor) recordInvocation(key string, args []interface{}) {
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

var _ v7.RoutesActor = new(FakeRoutesActor)
