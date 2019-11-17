// Code generated by counterfeiter. DO NOT EDIT.
package libcontainerdfakes

import (
	"context"
	"sync"

	"github.com/concourse/concourse/worker/backend/libcontainerd"
	"github.com/containerd/containerd"
)

type FakeClient struct {
	InitStub        func() error
	initMutex       sync.RWMutex
	initArgsForCall []struct {
	}
	initReturns struct {
		result1 error
	}
	initReturnsOnCall map[int]struct {
		result1 error
	}
	NewContainerStub        func(context.Context, string, ...containerd.NewContainerOpts) (containerd.Container, error)
	newContainerMutex       sync.RWMutex
	newContainerArgsForCall []struct {
		arg1 context.Context
		arg2 string
		arg3 []containerd.NewContainerOpts
	}
	newContainerReturns struct {
		result1 containerd.Container
		result2 error
	}
	newContainerReturnsOnCall map[int]struct {
		result1 containerd.Container
		result2 error
	}
	StopStub        func() error
	stopMutex       sync.RWMutex
	stopArgsForCall []struct {
	}
	stopReturns struct {
		result1 error
	}
	stopReturnsOnCall map[int]struct {
		result1 error
	}
	VersionStub        func(context.Context) error
	versionMutex       sync.RWMutex
	versionArgsForCall []struct {
		arg1 context.Context
	}
	versionReturns struct {
		result1 error
	}
	versionReturnsOnCall map[int]struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeClient) Init() error {
	fake.initMutex.Lock()
	ret, specificReturn := fake.initReturnsOnCall[len(fake.initArgsForCall)]
	fake.initArgsForCall = append(fake.initArgsForCall, struct {
	}{})
	fake.recordInvocation("Init", []interface{}{})
	fake.initMutex.Unlock()
	if fake.InitStub != nil {
		return fake.InitStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.initReturns
	return fakeReturns.result1
}

func (fake *FakeClient) InitCallCount() int {
	fake.initMutex.RLock()
	defer fake.initMutex.RUnlock()
	return len(fake.initArgsForCall)
}

func (fake *FakeClient) InitCalls(stub func() error) {
	fake.initMutex.Lock()
	defer fake.initMutex.Unlock()
	fake.InitStub = stub
}

func (fake *FakeClient) InitReturns(result1 error) {
	fake.initMutex.Lock()
	defer fake.initMutex.Unlock()
	fake.InitStub = nil
	fake.initReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeClient) InitReturnsOnCall(i int, result1 error) {
	fake.initMutex.Lock()
	defer fake.initMutex.Unlock()
	fake.InitStub = nil
	if fake.initReturnsOnCall == nil {
		fake.initReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.initReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeClient) NewContainer(arg1 context.Context, arg2 string, arg3 ...containerd.NewContainerOpts) (containerd.Container, error) {
	fake.newContainerMutex.Lock()
	ret, specificReturn := fake.newContainerReturnsOnCall[len(fake.newContainerArgsForCall)]
	fake.newContainerArgsForCall = append(fake.newContainerArgsForCall, struct {
		arg1 context.Context
		arg2 string
		arg3 []containerd.NewContainerOpts
	}{arg1, arg2, arg3})
	fake.recordInvocation("NewContainer", []interface{}{arg1, arg2, arg3})
	fake.newContainerMutex.Unlock()
	if fake.NewContainerStub != nil {
		return fake.NewContainerStub(arg1, arg2, arg3...)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.newContainerReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeClient) NewContainerCallCount() int {
	fake.newContainerMutex.RLock()
	defer fake.newContainerMutex.RUnlock()
	return len(fake.newContainerArgsForCall)
}

func (fake *FakeClient) NewContainerCalls(stub func(context.Context, string, ...containerd.NewContainerOpts) (containerd.Container, error)) {
	fake.newContainerMutex.Lock()
	defer fake.newContainerMutex.Unlock()
	fake.NewContainerStub = stub
}

func (fake *FakeClient) NewContainerArgsForCall(i int) (context.Context, string, []containerd.NewContainerOpts) {
	fake.newContainerMutex.RLock()
	defer fake.newContainerMutex.RUnlock()
	argsForCall := fake.newContainerArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakeClient) NewContainerReturns(result1 containerd.Container, result2 error) {
	fake.newContainerMutex.Lock()
	defer fake.newContainerMutex.Unlock()
	fake.NewContainerStub = nil
	fake.newContainerReturns = struct {
		result1 containerd.Container
		result2 error
	}{result1, result2}
}

func (fake *FakeClient) NewContainerReturnsOnCall(i int, result1 containerd.Container, result2 error) {
	fake.newContainerMutex.Lock()
	defer fake.newContainerMutex.Unlock()
	fake.NewContainerStub = nil
	if fake.newContainerReturnsOnCall == nil {
		fake.newContainerReturnsOnCall = make(map[int]struct {
			result1 containerd.Container
			result2 error
		})
	}
	fake.newContainerReturnsOnCall[i] = struct {
		result1 containerd.Container
		result2 error
	}{result1, result2}
}

func (fake *FakeClient) Stop() error {
	fake.stopMutex.Lock()
	ret, specificReturn := fake.stopReturnsOnCall[len(fake.stopArgsForCall)]
	fake.stopArgsForCall = append(fake.stopArgsForCall, struct {
	}{})
	fake.recordInvocation("Stop", []interface{}{})
	fake.stopMutex.Unlock()
	if fake.StopStub != nil {
		return fake.StopStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.stopReturns
	return fakeReturns.result1
}

func (fake *FakeClient) StopCallCount() int {
	fake.stopMutex.RLock()
	defer fake.stopMutex.RUnlock()
	return len(fake.stopArgsForCall)
}

func (fake *FakeClient) StopCalls(stub func() error) {
	fake.stopMutex.Lock()
	defer fake.stopMutex.Unlock()
	fake.StopStub = stub
}

func (fake *FakeClient) StopReturns(result1 error) {
	fake.stopMutex.Lock()
	defer fake.stopMutex.Unlock()
	fake.StopStub = nil
	fake.stopReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeClient) StopReturnsOnCall(i int, result1 error) {
	fake.stopMutex.Lock()
	defer fake.stopMutex.Unlock()
	fake.StopStub = nil
	if fake.stopReturnsOnCall == nil {
		fake.stopReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.stopReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeClient) Version(arg1 context.Context) error {
	fake.versionMutex.Lock()
	ret, specificReturn := fake.versionReturnsOnCall[len(fake.versionArgsForCall)]
	fake.versionArgsForCall = append(fake.versionArgsForCall, struct {
		arg1 context.Context
	}{arg1})
	fake.recordInvocation("Version", []interface{}{arg1})
	fake.versionMutex.Unlock()
	if fake.VersionStub != nil {
		return fake.VersionStub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.versionReturns
	return fakeReturns.result1
}

func (fake *FakeClient) VersionCallCount() int {
	fake.versionMutex.RLock()
	defer fake.versionMutex.RUnlock()
	return len(fake.versionArgsForCall)
}

func (fake *FakeClient) VersionCalls(stub func(context.Context) error) {
	fake.versionMutex.Lock()
	defer fake.versionMutex.Unlock()
	fake.VersionStub = stub
}

func (fake *FakeClient) VersionArgsForCall(i int) context.Context {
	fake.versionMutex.RLock()
	defer fake.versionMutex.RUnlock()
	argsForCall := fake.versionArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeClient) VersionReturns(result1 error) {
	fake.versionMutex.Lock()
	defer fake.versionMutex.Unlock()
	fake.VersionStub = nil
	fake.versionReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeClient) VersionReturnsOnCall(i int, result1 error) {
	fake.versionMutex.Lock()
	defer fake.versionMutex.Unlock()
	fake.VersionStub = nil
	if fake.versionReturnsOnCall == nil {
		fake.versionReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.versionReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeClient) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.initMutex.RLock()
	defer fake.initMutex.RUnlock()
	fake.newContainerMutex.RLock()
	defer fake.newContainerMutex.RUnlock()
	fake.stopMutex.RLock()
	defer fake.stopMutex.RUnlock()
	fake.versionMutex.RLock()
	defer fake.versionMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeClient) recordInvocation(key string, args []interface{}) {
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

var _ libcontainerd.Client = new(FakeClient)
