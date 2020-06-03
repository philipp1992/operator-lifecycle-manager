// Code generated by counterfeiter. DO NOT EDIT.
package fakes

import (
	"context"
	"sync"
	"time"

	"github.com/operator-framework/operator-lifecycle-manager/pkg/controller/registry"
	"github.com/operator-framework/operator-registry/pkg/api"
)

type FakeClientInterface struct {
	CloseStub        func() error
	closeMutex       sync.RWMutex
	closeArgsForCall []struct {
	}
	closeReturns struct {
		result1 error
	}
	closeReturnsOnCall map[int]struct {
		result1 error
	}
	FindBundleThatProvidesStub        func(context.Context, string, string, string, string) (*api.Bundle, error)
	findBundleThatProvidesMutex       sync.RWMutex
	findBundleThatProvidesArgsForCall []struct {
		arg1 context.Context
		arg2 string
		arg3 string
		arg4 string
		arg5 string
	}
	findBundleThatProvidesReturns struct {
		result1 *api.Bundle
		result2 error
	}
	findBundleThatProvidesReturnsOnCall map[int]struct {
		result1 *api.Bundle
		result2 error
	}
	GetBundleStub        func(context.Context, string, string, string) (*api.Bundle, error)
	getBundleMutex       sync.RWMutex
	getBundleArgsForCall []struct {
		arg1 context.Context
		arg2 string
		arg3 string
		arg4 string
	}
	getBundleReturns struct {
		result1 *api.Bundle
		result2 error
	}
	getBundleReturnsOnCall map[int]struct {
		result1 *api.Bundle
		result2 error
	}
	GetBundleInPackageChannelStub        func(context.Context, string, string) (*api.Bundle, error)
	getBundleInPackageChannelMutex       sync.RWMutex
	getBundleInPackageChannelArgsForCall []struct {
		arg1 context.Context
		arg2 string
		arg3 string
	}
	getBundleInPackageChannelReturns struct {
		result1 *api.Bundle
		result2 error
	}
	getBundleInPackageChannelReturnsOnCall map[int]struct {
		result1 *api.Bundle
		result2 error
	}
	GetBundleThatProvidesStub        func(context.Context, string, string, string) (*api.Bundle, error)
	getBundleThatProvidesMutex       sync.RWMutex
	getBundleThatProvidesArgsForCall []struct {
		arg1 context.Context
		arg2 string
		arg3 string
		arg4 string
	}
	getBundleThatProvidesReturns struct {
		result1 *api.Bundle
		result2 error
	}
	getBundleThatProvidesReturnsOnCall map[int]struct {
		result1 *api.Bundle
		result2 error
	}
	GetLatestChannelEntriesThatProvideStub        func(context.Context, string, string, string) (*registry.ChannelEntryIterator, error)
	getLatestChannelEntriesThatProvideMutex       sync.RWMutex
	getLatestChannelEntriesThatProvideArgsForCall []struct {
		arg1 context.Context
		arg2 string
		arg3 string
		arg4 string
	}
	getLatestChannelEntriesThatProvideReturns struct {
		result1 *registry.ChannelEntryIterator
		result2 error
	}
	getLatestChannelEntriesThatProvideReturnsOnCall map[int]struct {
		result1 *registry.ChannelEntryIterator
		result2 error
	}
	GetReplacementBundleInPackageChannelStub        func(context.Context, string, string, string) (*api.Bundle, error)
	getReplacementBundleInPackageChannelMutex       sync.RWMutex
	getReplacementBundleInPackageChannelArgsForCall []struct {
		arg1 context.Context
		arg2 string
		arg3 string
		arg4 string
	}
	getReplacementBundleInPackageChannelReturns struct {
		result1 *api.Bundle
		result2 error
	}
	getReplacementBundleInPackageChannelReturnsOnCall map[int]struct {
		result1 *api.Bundle
		result2 error
	}
	HealthCheckStub        func(context.Context, time.Duration) (bool, error)
	healthCheckMutex       sync.RWMutex
	healthCheckArgsForCall []struct {
		arg1 context.Context
		arg2 time.Duration
	}
	healthCheckReturns struct {
		result1 bool
		result2 error
	}
	healthCheckReturnsOnCall map[int]struct {
		result1 bool
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeClientInterface) Close() error {
	fake.closeMutex.Lock()
	ret, specificReturn := fake.closeReturnsOnCall[len(fake.closeArgsForCall)]
	fake.closeArgsForCall = append(fake.closeArgsForCall, struct {
	}{})
	fake.recordInvocation("Close", []interface{}{})
	fake.closeMutex.Unlock()
	if fake.CloseStub != nil {
		return fake.CloseStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.closeReturns
	return fakeReturns.result1
}

func (fake *FakeClientInterface) CloseCallCount() int {
	fake.closeMutex.RLock()
	defer fake.closeMutex.RUnlock()
	return len(fake.closeArgsForCall)
}

func (fake *FakeClientInterface) CloseCalls(stub func() error) {
	fake.closeMutex.Lock()
	defer fake.closeMutex.Unlock()
	fake.CloseStub = stub
}

func (fake *FakeClientInterface) CloseReturns(result1 error) {
	fake.closeMutex.Lock()
	defer fake.closeMutex.Unlock()
	fake.CloseStub = nil
	fake.closeReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeClientInterface) CloseReturnsOnCall(i int, result1 error) {
	fake.closeMutex.Lock()
	defer fake.closeMutex.Unlock()
	fake.CloseStub = nil
	if fake.closeReturnsOnCall == nil {
		fake.closeReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.closeReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeClientInterface) FindBundleThatProvides(arg1 context.Context, arg2 string, arg3 string, arg4 string, arg5 string) (*api.Bundle, error) {
	fake.findBundleThatProvidesMutex.Lock()
	ret, specificReturn := fake.findBundleThatProvidesReturnsOnCall[len(fake.findBundleThatProvidesArgsForCall)]
	fake.findBundleThatProvidesArgsForCall = append(fake.findBundleThatProvidesArgsForCall, struct {
		arg1 context.Context
		arg2 string
		arg3 string
		arg4 string
		arg5 string
	}{arg1, arg2, arg3, arg4, arg5})
	fake.recordInvocation("FindBundleThatProvides", []interface{}{arg1, arg2, arg3, arg4, arg5})
	fake.findBundleThatProvidesMutex.Unlock()
	if fake.FindBundleThatProvidesStub != nil {
		return fake.FindBundleThatProvidesStub(arg1, arg2, arg3, arg4, arg5)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.findBundleThatProvidesReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeClientInterface) FindBundleThatProvidesCallCount() int {
	fake.findBundleThatProvidesMutex.RLock()
	defer fake.findBundleThatProvidesMutex.RUnlock()
	return len(fake.findBundleThatProvidesArgsForCall)
}

func (fake *FakeClientInterface) FindBundleThatProvidesCalls(stub func(context.Context, string, string, string, string) (*api.Bundle, error)) {
	fake.findBundleThatProvidesMutex.Lock()
	defer fake.findBundleThatProvidesMutex.Unlock()
	fake.FindBundleThatProvidesStub = stub
}

func (fake *FakeClientInterface) FindBundleThatProvidesArgsForCall(i int) (context.Context, string, string, string, string) {
	fake.findBundleThatProvidesMutex.RLock()
	defer fake.findBundleThatProvidesMutex.RUnlock()
	argsForCall := fake.findBundleThatProvidesArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3, argsForCall.arg4, argsForCall.arg5
}

func (fake *FakeClientInterface) FindBundleThatProvidesReturns(result1 *api.Bundle, result2 error) {
	fake.findBundleThatProvidesMutex.Lock()
	defer fake.findBundleThatProvidesMutex.Unlock()
	fake.FindBundleThatProvidesStub = nil
	fake.findBundleThatProvidesReturns = struct {
		result1 *api.Bundle
		result2 error
	}{result1, result2}
}

func (fake *FakeClientInterface) FindBundleThatProvidesReturnsOnCall(i int, result1 *api.Bundle, result2 error) {
	fake.findBundleThatProvidesMutex.Lock()
	defer fake.findBundleThatProvidesMutex.Unlock()
	fake.FindBundleThatProvidesStub = nil
	if fake.findBundleThatProvidesReturnsOnCall == nil {
		fake.findBundleThatProvidesReturnsOnCall = make(map[int]struct {
			result1 *api.Bundle
			result2 error
		})
	}
	fake.findBundleThatProvidesReturnsOnCall[i] = struct {
		result1 *api.Bundle
		result2 error
	}{result1, result2}
}

func (fake *FakeClientInterface) GetBundle(arg1 context.Context, arg2 string, arg3 string, arg4 string) (*api.Bundle, error) {
	fake.getBundleMutex.Lock()
	ret, specificReturn := fake.getBundleReturnsOnCall[len(fake.getBundleArgsForCall)]
	fake.getBundleArgsForCall = append(fake.getBundleArgsForCall, struct {
		arg1 context.Context
		arg2 string
		arg3 string
		arg4 string
	}{arg1, arg2, arg3, arg4})
	fake.recordInvocation("GetBundle", []interface{}{arg1, arg2, arg3, arg4})
	fake.getBundleMutex.Unlock()
	if fake.GetBundleStub != nil {
		return fake.GetBundleStub(arg1, arg2, arg3, arg4)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.getBundleReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeClientInterface) GetBundleCallCount() int {
	fake.getBundleMutex.RLock()
	defer fake.getBundleMutex.RUnlock()
	return len(fake.getBundleArgsForCall)
}

func (fake *FakeClientInterface) GetBundleCalls(stub func(context.Context, string, string, string) (*api.Bundle, error)) {
	fake.getBundleMutex.Lock()
	defer fake.getBundleMutex.Unlock()
	fake.GetBundleStub = stub
}

func (fake *FakeClientInterface) GetBundleArgsForCall(i int) (context.Context, string, string, string) {
	fake.getBundleMutex.RLock()
	defer fake.getBundleMutex.RUnlock()
	argsForCall := fake.getBundleArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3, argsForCall.arg4
}

func (fake *FakeClientInterface) GetBundleReturns(result1 *api.Bundle, result2 error) {
	fake.getBundleMutex.Lock()
	defer fake.getBundleMutex.Unlock()
	fake.GetBundleStub = nil
	fake.getBundleReturns = struct {
		result1 *api.Bundle
		result2 error
	}{result1, result2}
}

func (fake *FakeClientInterface) GetBundleReturnsOnCall(i int, result1 *api.Bundle, result2 error) {
	fake.getBundleMutex.Lock()
	defer fake.getBundleMutex.Unlock()
	fake.GetBundleStub = nil
	if fake.getBundleReturnsOnCall == nil {
		fake.getBundleReturnsOnCall = make(map[int]struct {
			result1 *api.Bundle
			result2 error
		})
	}
	fake.getBundleReturnsOnCall[i] = struct {
		result1 *api.Bundle
		result2 error
	}{result1, result2}
}

func (fake *FakeClientInterface) GetBundleInPackageChannel(arg1 context.Context, arg2 string, arg3 string) (*api.Bundle, error) {
	fake.getBundleInPackageChannelMutex.Lock()
	ret, specificReturn := fake.getBundleInPackageChannelReturnsOnCall[len(fake.getBundleInPackageChannelArgsForCall)]
	fake.getBundleInPackageChannelArgsForCall = append(fake.getBundleInPackageChannelArgsForCall, struct {
		arg1 context.Context
		arg2 string
		arg3 string
	}{arg1, arg2, arg3})
	fake.recordInvocation("GetBundleInPackageChannel", []interface{}{arg1, arg2, arg3})
	fake.getBundleInPackageChannelMutex.Unlock()
	if fake.GetBundleInPackageChannelStub != nil {
		return fake.GetBundleInPackageChannelStub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.getBundleInPackageChannelReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeClientInterface) GetBundleInPackageChannelCallCount() int {
	fake.getBundleInPackageChannelMutex.RLock()
	defer fake.getBundleInPackageChannelMutex.RUnlock()
	return len(fake.getBundleInPackageChannelArgsForCall)
}

func (fake *FakeClientInterface) GetBundleInPackageChannelCalls(stub func(context.Context, string, string) (*api.Bundle, error)) {
	fake.getBundleInPackageChannelMutex.Lock()
	defer fake.getBundleInPackageChannelMutex.Unlock()
	fake.GetBundleInPackageChannelStub = stub
}

func (fake *FakeClientInterface) GetBundleInPackageChannelArgsForCall(i int) (context.Context, string, string) {
	fake.getBundleInPackageChannelMutex.RLock()
	defer fake.getBundleInPackageChannelMutex.RUnlock()
	argsForCall := fake.getBundleInPackageChannelArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakeClientInterface) GetBundleInPackageChannelReturns(result1 *api.Bundle, result2 error) {
	fake.getBundleInPackageChannelMutex.Lock()
	defer fake.getBundleInPackageChannelMutex.Unlock()
	fake.GetBundleInPackageChannelStub = nil
	fake.getBundleInPackageChannelReturns = struct {
		result1 *api.Bundle
		result2 error
	}{result1, result2}
}

func (fake *FakeClientInterface) GetBundleInPackageChannelReturnsOnCall(i int, result1 *api.Bundle, result2 error) {
	fake.getBundleInPackageChannelMutex.Lock()
	defer fake.getBundleInPackageChannelMutex.Unlock()
	fake.GetBundleInPackageChannelStub = nil
	if fake.getBundleInPackageChannelReturnsOnCall == nil {
		fake.getBundleInPackageChannelReturnsOnCall = make(map[int]struct {
			result1 *api.Bundle
			result2 error
		})
	}
	fake.getBundleInPackageChannelReturnsOnCall[i] = struct {
		result1 *api.Bundle
		result2 error
	}{result1, result2}
}

func (fake *FakeClientInterface) GetBundleThatProvides(arg1 context.Context, arg2 string, arg3 string, arg4 string) (*api.Bundle, error) {
	fake.getBundleThatProvidesMutex.Lock()
	ret, specificReturn := fake.getBundleThatProvidesReturnsOnCall[len(fake.getBundleThatProvidesArgsForCall)]
	fake.getBundleThatProvidesArgsForCall = append(fake.getBundleThatProvidesArgsForCall, struct {
		arg1 context.Context
		arg2 string
		arg3 string
		arg4 string
	}{arg1, arg2, arg3, arg4})
	fake.recordInvocation("GetBundleThatProvides", []interface{}{arg1, arg2, arg3, arg4})
	fake.getBundleThatProvidesMutex.Unlock()
	if fake.GetBundleThatProvidesStub != nil {
		return fake.GetBundleThatProvidesStub(arg1, arg2, arg3, arg4)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.getBundleThatProvidesReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeClientInterface) GetBundleThatProvidesCallCount() int {
	fake.getBundleThatProvidesMutex.RLock()
	defer fake.getBundleThatProvidesMutex.RUnlock()
	return len(fake.getBundleThatProvidesArgsForCall)
}

func (fake *FakeClientInterface) GetBundleThatProvidesCalls(stub func(context.Context, string, string, string) (*api.Bundle, error)) {
	fake.getBundleThatProvidesMutex.Lock()
	defer fake.getBundleThatProvidesMutex.Unlock()
	fake.GetBundleThatProvidesStub = stub
}

func (fake *FakeClientInterface) GetBundleThatProvidesArgsForCall(i int) (context.Context, string, string, string) {
	fake.getBundleThatProvidesMutex.RLock()
	defer fake.getBundleThatProvidesMutex.RUnlock()
	argsForCall := fake.getBundleThatProvidesArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3, argsForCall.arg4
}

func (fake *FakeClientInterface) GetBundleThatProvidesReturns(result1 *api.Bundle, result2 error) {
	fake.getBundleThatProvidesMutex.Lock()
	defer fake.getBundleThatProvidesMutex.Unlock()
	fake.GetBundleThatProvidesStub = nil
	fake.getBundleThatProvidesReturns = struct {
		result1 *api.Bundle
		result2 error
	}{result1, result2}
}

func (fake *FakeClientInterface) GetBundleThatProvidesReturnsOnCall(i int, result1 *api.Bundle, result2 error) {
	fake.getBundleThatProvidesMutex.Lock()
	defer fake.getBundleThatProvidesMutex.Unlock()
	fake.GetBundleThatProvidesStub = nil
	if fake.getBundleThatProvidesReturnsOnCall == nil {
		fake.getBundleThatProvidesReturnsOnCall = make(map[int]struct {
			result1 *api.Bundle
			result2 error
		})
	}
	fake.getBundleThatProvidesReturnsOnCall[i] = struct {
		result1 *api.Bundle
		result2 error
	}{result1, result2}
}

func (fake *FakeClientInterface) GetLatestChannelEntriesThatProvide(arg1 context.Context, arg2 string, arg3 string, arg4 string) (*registry.ChannelEntryIterator, error) {
	fake.getLatestChannelEntriesThatProvideMutex.Lock()
	ret, specificReturn := fake.getLatestChannelEntriesThatProvideReturnsOnCall[len(fake.getLatestChannelEntriesThatProvideArgsForCall)]
	fake.getLatestChannelEntriesThatProvideArgsForCall = append(fake.getLatestChannelEntriesThatProvideArgsForCall, struct {
		arg1 context.Context
		arg2 string
		arg3 string
		arg4 string
	}{arg1, arg2, arg3, arg4})
	fake.recordInvocation("GetLatestChannelEntriesThatProvide", []interface{}{arg1, arg2, arg3, arg4})
	fake.getLatestChannelEntriesThatProvideMutex.Unlock()
	if fake.GetLatestChannelEntriesThatProvideStub != nil {
		return fake.GetLatestChannelEntriesThatProvideStub(arg1, arg2, arg3, arg4)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.getLatestChannelEntriesThatProvideReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeClientInterface) GetLatestChannelEntriesThatProvideCallCount() int {
	fake.getLatestChannelEntriesThatProvideMutex.RLock()
	defer fake.getLatestChannelEntriesThatProvideMutex.RUnlock()
	return len(fake.getLatestChannelEntriesThatProvideArgsForCall)
}

func (fake *FakeClientInterface) GetLatestChannelEntriesThatProvideCalls(stub func(context.Context, string, string, string) (*registry.ChannelEntryIterator, error)) {
	fake.getLatestChannelEntriesThatProvideMutex.Lock()
	defer fake.getLatestChannelEntriesThatProvideMutex.Unlock()
	fake.GetLatestChannelEntriesThatProvideStub = stub
}

func (fake *FakeClientInterface) GetLatestChannelEntriesThatProvideArgsForCall(i int) (context.Context, string, string, string) {
	fake.getLatestChannelEntriesThatProvideMutex.RLock()
	defer fake.getLatestChannelEntriesThatProvideMutex.RUnlock()
	argsForCall := fake.getLatestChannelEntriesThatProvideArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3, argsForCall.arg4
}

func (fake *FakeClientInterface) GetLatestChannelEntriesThatProvideReturns(result1 *registry.ChannelEntryIterator, result2 error) {
	fake.getLatestChannelEntriesThatProvideMutex.Lock()
	defer fake.getLatestChannelEntriesThatProvideMutex.Unlock()
	fake.GetLatestChannelEntriesThatProvideStub = nil
	fake.getLatestChannelEntriesThatProvideReturns = struct {
		result1 *registry.ChannelEntryIterator
		result2 error
	}{result1, result2}
}

func (fake *FakeClientInterface) GetLatestChannelEntriesThatProvideReturnsOnCall(i int, result1 *registry.ChannelEntryIterator, result2 error) {
	fake.getLatestChannelEntriesThatProvideMutex.Lock()
	defer fake.getLatestChannelEntriesThatProvideMutex.Unlock()
	fake.GetLatestChannelEntriesThatProvideStub = nil
	if fake.getLatestChannelEntriesThatProvideReturnsOnCall == nil {
		fake.getLatestChannelEntriesThatProvideReturnsOnCall = make(map[int]struct {
			result1 *registry.ChannelEntryIterator
			result2 error
		})
	}
	fake.getLatestChannelEntriesThatProvideReturnsOnCall[i] = struct {
		result1 *registry.ChannelEntryIterator
		result2 error
	}{result1, result2}
}

func (fake *FakeClientInterface) GetReplacementBundleInPackageChannel(arg1 context.Context, arg2 string, arg3 string, arg4 string) (*api.Bundle, error) {
	fake.getReplacementBundleInPackageChannelMutex.Lock()
	ret, specificReturn := fake.getReplacementBundleInPackageChannelReturnsOnCall[len(fake.getReplacementBundleInPackageChannelArgsForCall)]
	fake.getReplacementBundleInPackageChannelArgsForCall = append(fake.getReplacementBundleInPackageChannelArgsForCall, struct {
		arg1 context.Context
		arg2 string
		arg3 string
		arg4 string
	}{arg1, arg2, arg3, arg4})
	fake.recordInvocation("GetReplacementBundleInPackageChannel", []interface{}{arg1, arg2, arg3, arg4})
	fake.getReplacementBundleInPackageChannelMutex.Unlock()
	if fake.GetReplacementBundleInPackageChannelStub != nil {
		return fake.GetReplacementBundleInPackageChannelStub(arg1, arg2, arg3, arg4)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.getReplacementBundleInPackageChannelReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeClientInterface) GetReplacementBundleInPackageChannelCallCount() int {
	fake.getReplacementBundleInPackageChannelMutex.RLock()
	defer fake.getReplacementBundleInPackageChannelMutex.RUnlock()
	return len(fake.getReplacementBundleInPackageChannelArgsForCall)
}

func (fake *FakeClientInterface) GetReplacementBundleInPackageChannelCalls(stub func(context.Context, string, string, string) (*api.Bundle, error)) {
	fake.getReplacementBundleInPackageChannelMutex.Lock()
	defer fake.getReplacementBundleInPackageChannelMutex.Unlock()
	fake.GetReplacementBundleInPackageChannelStub = stub
}

func (fake *FakeClientInterface) GetReplacementBundleInPackageChannelArgsForCall(i int) (context.Context, string, string, string) {
	fake.getReplacementBundleInPackageChannelMutex.RLock()
	defer fake.getReplacementBundleInPackageChannelMutex.RUnlock()
	argsForCall := fake.getReplacementBundleInPackageChannelArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3, argsForCall.arg4
}

func (fake *FakeClientInterface) GetReplacementBundleInPackageChannelReturns(result1 *api.Bundle, result2 error) {
	fake.getReplacementBundleInPackageChannelMutex.Lock()
	defer fake.getReplacementBundleInPackageChannelMutex.Unlock()
	fake.GetReplacementBundleInPackageChannelStub = nil
	fake.getReplacementBundleInPackageChannelReturns = struct {
		result1 *api.Bundle
		result2 error
	}{result1, result2}
}

func (fake *FakeClientInterface) GetReplacementBundleInPackageChannelReturnsOnCall(i int, result1 *api.Bundle, result2 error) {
	fake.getReplacementBundleInPackageChannelMutex.Lock()
	defer fake.getReplacementBundleInPackageChannelMutex.Unlock()
	fake.GetReplacementBundleInPackageChannelStub = nil
	if fake.getReplacementBundleInPackageChannelReturnsOnCall == nil {
		fake.getReplacementBundleInPackageChannelReturnsOnCall = make(map[int]struct {
			result1 *api.Bundle
			result2 error
		})
	}
	fake.getReplacementBundleInPackageChannelReturnsOnCall[i] = struct {
		result1 *api.Bundle
		result2 error
	}{result1, result2}
}

func (fake *FakeClientInterface) HealthCheck(arg1 context.Context, arg2 time.Duration) (bool, error) {
	fake.healthCheckMutex.Lock()
	ret, specificReturn := fake.healthCheckReturnsOnCall[len(fake.healthCheckArgsForCall)]
	fake.healthCheckArgsForCall = append(fake.healthCheckArgsForCall, struct {
		arg1 context.Context
		arg2 time.Duration
	}{arg1, arg2})
	fake.recordInvocation("HealthCheck", []interface{}{arg1, arg2})
	fake.healthCheckMutex.Unlock()
	if fake.HealthCheckStub != nil {
		return fake.HealthCheckStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.healthCheckReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeClientInterface) HealthCheckCallCount() int {
	fake.healthCheckMutex.RLock()
	defer fake.healthCheckMutex.RUnlock()
	return len(fake.healthCheckArgsForCall)
}

func (fake *FakeClientInterface) HealthCheckCalls(stub func(context.Context, time.Duration) (bool, error)) {
	fake.healthCheckMutex.Lock()
	defer fake.healthCheckMutex.Unlock()
	fake.HealthCheckStub = stub
}

func (fake *FakeClientInterface) HealthCheckArgsForCall(i int) (context.Context, time.Duration) {
	fake.healthCheckMutex.RLock()
	defer fake.healthCheckMutex.RUnlock()
	argsForCall := fake.healthCheckArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeClientInterface) HealthCheckReturns(result1 bool, result2 error) {
	fake.healthCheckMutex.Lock()
	defer fake.healthCheckMutex.Unlock()
	fake.HealthCheckStub = nil
	fake.healthCheckReturns = struct {
		result1 bool
		result2 error
	}{result1, result2}
}

func (fake *FakeClientInterface) HealthCheckReturnsOnCall(i int, result1 bool, result2 error) {
	fake.healthCheckMutex.Lock()
	defer fake.healthCheckMutex.Unlock()
	fake.HealthCheckStub = nil
	if fake.healthCheckReturnsOnCall == nil {
		fake.healthCheckReturnsOnCall = make(map[int]struct {
			result1 bool
			result2 error
		})
	}
	fake.healthCheckReturnsOnCall[i] = struct {
		result1 bool
		result2 error
	}{result1, result2}
}

func (fake *FakeClientInterface) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.closeMutex.RLock()
	defer fake.closeMutex.RUnlock()
	fake.findBundleThatProvidesMutex.RLock()
	defer fake.findBundleThatProvidesMutex.RUnlock()
	fake.getBundleMutex.RLock()
	defer fake.getBundleMutex.RUnlock()
	fake.getBundleInPackageChannelMutex.RLock()
	defer fake.getBundleInPackageChannelMutex.RUnlock()
	fake.getBundleThatProvidesMutex.RLock()
	defer fake.getBundleThatProvidesMutex.RUnlock()
	fake.getLatestChannelEntriesThatProvideMutex.RLock()
	defer fake.getLatestChannelEntriesThatProvideMutex.RUnlock()
	fake.getReplacementBundleInPackageChannelMutex.RLock()
	defer fake.getReplacementBundleInPackageChannelMutex.RUnlock()
	fake.healthCheckMutex.RLock()
	defer fake.healthCheckMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeClientInterface) recordInvocation(key string, args []interface{}) {
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

var _ registry.ClientInterface = new(FakeClientInterface)
