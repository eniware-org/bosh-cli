// Code generated by counterfeiter. DO NOT EDIT.
package uaafakes

import (
	"sync"

	"github.com/cloudfoundry/bosh-cli/uaa"
)

type FakeRefreshableAccessToken struct {
	TypeStub        func() string
	typeMutex       sync.RWMutex
	typeArgsForCall []struct{}
	typeReturns     struct {
		result1 string
	}
	typeReturnsOnCall map[int]struct {
		result1 string
	}
	ValueStub        func() string
	valueMutex       sync.RWMutex
	valueArgsForCall []struct{}
	valueReturns     struct {
		result1 string
	}
	valueReturnsOnCall map[int]struct {
		result1 string
	}
	IsValidStub        func() bool
	isValidMutex       sync.RWMutex
	isValidArgsForCall []struct{}
	isValidReturns     struct {
		result1 bool
	}
	isValidReturnsOnCall map[int]struct {
		result1 bool
	}
	RefreshValueStub        func() string
	refreshValueMutex       sync.RWMutex
	refreshValueArgsForCall []struct{}
	refreshValueReturns     struct {
		result1 string
	}
	refreshValueReturnsOnCall map[int]struct {
		result1 string
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeRefreshableAccessToken) Type() string {
	fake.typeMutex.Lock()
	ret, specificReturn := fake.typeReturnsOnCall[len(fake.typeArgsForCall)]
	fake.typeArgsForCall = append(fake.typeArgsForCall, struct{}{})
	fake.recordInvocation("Type", []interface{}{})
	fake.typeMutex.Unlock()
	if fake.TypeStub != nil {
		return fake.TypeStub()
	}
	if specificReturn {
		return ret.result1
	}
	return fake.typeReturns.result1
}

func (fake *FakeRefreshableAccessToken) TypeCallCount() int {
	fake.typeMutex.RLock()
	defer fake.typeMutex.RUnlock()
	return len(fake.typeArgsForCall)
}

func (fake *FakeRefreshableAccessToken) TypeReturns(result1 string) {
	fake.TypeStub = nil
	fake.typeReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeRefreshableAccessToken) TypeReturnsOnCall(i int, result1 string) {
	fake.TypeStub = nil
	if fake.typeReturnsOnCall == nil {
		fake.typeReturnsOnCall = make(map[int]struct {
			result1 string
		})
	}
	fake.typeReturnsOnCall[i] = struct {
		result1 string
	}{result1}
}

func (fake *FakeRefreshableAccessToken) Value() string {
	fake.valueMutex.Lock()
	ret, specificReturn := fake.valueReturnsOnCall[len(fake.valueArgsForCall)]
	fake.valueArgsForCall = append(fake.valueArgsForCall, struct{}{})
	fake.recordInvocation("Value", []interface{}{})
	fake.valueMutex.Unlock()
	if fake.ValueStub != nil {
		return fake.ValueStub()
	}
	if specificReturn {
		return ret.result1
	}
	return fake.valueReturns.result1
}

func (fake *FakeRefreshableAccessToken) ValueCallCount() int {
	fake.valueMutex.RLock()
	defer fake.valueMutex.RUnlock()
	return len(fake.valueArgsForCall)
}

func (fake *FakeRefreshableAccessToken) ValueReturns(result1 string) {
	fake.ValueStub = nil
	fake.valueReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeRefreshableAccessToken) ValueReturnsOnCall(i int, result1 string) {
	fake.ValueStub = nil
	if fake.valueReturnsOnCall == nil {
		fake.valueReturnsOnCall = make(map[int]struct {
			result1 string
		})
	}
	fake.valueReturnsOnCall[i] = struct {
		result1 string
	}{result1}
}

func (fake *FakeRefreshableAccessToken) IsValid() bool {
	fake.isValidMutex.Lock()
	ret, specificReturn := fake.isValidReturnsOnCall[len(fake.isValidArgsForCall)]
	fake.isValidArgsForCall = append(fake.isValidArgsForCall, struct{}{})
	fake.recordInvocation("IsValid", []interface{}{})
	fake.isValidMutex.Unlock()
	if fake.IsValidStub != nil {
		return fake.IsValidStub()
	}
	if specificReturn {
		return ret.result1
	}
	return fake.isValidReturns.result1
}

func (fake *FakeRefreshableAccessToken) IsValidCallCount() int {
	fake.isValidMutex.RLock()
	defer fake.isValidMutex.RUnlock()
	return len(fake.isValidArgsForCall)
}

func (fake *FakeRefreshableAccessToken) IsValidReturns(result1 bool) {
	fake.IsValidStub = nil
	fake.isValidReturns = struct {
		result1 bool
	}{result1}
}

func (fake *FakeRefreshableAccessToken) IsValidReturnsOnCall(i int, result1 bool) {
	fake.IsValidStub = nil
	if fake.isValidReturnsOnCall == nil {
		fake.isValidReturnsOnCall = make(map[int]struct {
			result1 bool
		})
	}
	fake.isValidReturnsOnCall[i] = struct {
		result1 bool
	}{result1}
}

func (fake *FakeRefreshableAccessToken) RefreshValue() string {
	fake.refreshValueMutex.Lock()
	ret, specificReturn := fake.refreshValueReturnsOnCall[len(fake.refreshValueArgsForCall)]
	fake.refreshValueArgsForCall = append(fake.refreshValueArgsForCall, struct{}{})
	fake.recordInvocation("RefreshValue", []interface{}{})
	fake.refreshValueMutex.Unlock()
	if fake.RefreshValueStub != nil {
		return fake.RefreshValueStub()
	}
	if specificReturn {
		return ret.result1
	}
	return fake.refreshValueReturns.result1
}

func (fake *FakeRefreshableAccessToken) RefreshValueCallCount() int {
	fake.refreshValueMutex.RLock()
	defer fake.refreshValueMutex.RUnlock()
	return len(fake.refreshValueArgsForCall)
}

func (fake *FakeRefreshableAccessToken) RefreshValueReturns(result1 string) {
	fake.RefreshValueStub = nil
	fake.refreshValueReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeRefreshableAccessToken) RefreshValueReturnsOnCall(i int, result1 string) {
	fake.RefreshValueStub = nil
	if fake.refreshValueReturnsOnCall == nil {
		fake.refreshValueReturnsOnCall = make(map[int]struct {
			result1 string
		})
	}
	fake.refreshValueReturnsOnCall[i] = struct {
		result1 string
	}{result1}
}

func (fake *FakeRefreshableAccessToken) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.typeMutex.RLock()
	defer fake.typeMutex.RUnlock()
	fake.valueMutex.RLock()
	defer fake.valueMutex.RUnlock()
	fake.isValidMutex.RLock()
	defer fake.isValidMutex.RUnlock()
	fake.refreshValueMutex.RLock()
	defer fake.refreshValueMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeRefreshableAccessToken) recordInvocation(key string, args []interface{}) {
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

var _ uaa.RefreshableAccessToken = new(FakeRefreshableAccessToken)