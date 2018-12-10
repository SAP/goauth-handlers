// Code generated by counterfeiter. DO NOT EDIT.
package sessionfakes

import (
	sync "sync"

	session "github.com/SAP/goauth-handlers/session"
)

type FakeSession struct {
	ClearStub        func()
	clearMutex       sync.RWMutex
	clearArgsForCall []struct {
	}
	NameStub        func() string
	nameMutex       sync.RWMutex
	nameArgsForCall []struct {
	}
	nameReturns struct {
		result1 string
	}
	nameReturnsOnCall map[int]struct {
		result1 string
	}
	ValuesStub        func() map[string]string
	valuesMutex       sync.RWMutex
	valuesArgsForCall []struct {
	}
	valuesReturns struct {
		result1 map[string]string
	}
	valuesReturnsOnCall map[int]struct {
		result1 map[string]string
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeSession) Clear() {
	fake.clearMutex.Lock()
	fake.clearArgsForCall = append(fake.clearArgsForCall, struct {
	}{})
	fake.recordInvocation("Clear", []interface{}{})
	fake.clearMutex.Unlock()
	if fake.ClearStub != nil {
		fake.ClearStub()
	}
}

func (fake *FakeSession) ClearCallCount() int {
	fake.clearMutex.RLock()
	defer fake.clearMutex.RUnlock()
	return len(fake.clearArgsForCall)
}

func (fake *FakeSession) ClearCalls(stub func()) {
	fake.clearMutex.Lock()
	defer fake.clearMutex.Unlock()
	fake.ClearStub = stub
}

func (fake *FakeSession) Name() string {
	fake.nameMutex.Lock()
	ret, specificReturn := fake.nameReturnsOnCall[len(fake.nameArgsForCall)]
	fake.nameArgsForCall = append(fake.nameArgsForCall, struct {
	}{})
	fake.recordInvocation("Name", []interface{}{})
	fake.nameMutex.Unlock()
	if fake.NameStub != nil {
		return fake.NameStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.nameReturns
	return fakeReturns.result1
}

func (fake *FakeSession) NameCallCount() int {
	fake.nameMutex.RLock()
	defer fake.nameMutex.RUnlock()
	return len(fake.nameArgsForCall)
}

func (fake *FakeSession) NameCalls(stub func() string) {
	fake.nameMutex.Lock()
	defer fake.nameMutex.Unlock()
	fake.NameStub = stub
}

func (fake *FakeSession) NameReturns(result1 string) {
	fake.nameMutex.Lock()
	defer fake.nameMutex.Unlock()
	fake.NameStub = nil
	fake.nameReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeSession) NameReturnsOnCall(i int, result1 string) {
	fake.nameMutex.Lock()
	defer fake.nameMutex.Unlock()
	fake.NameStub = nil
	if fake.nameReturnsOnCall == nil {
		fake.nameReturnsOnCall = make(map[int]struct {
			result1 string
		})
	}
	fake.nameReturnsOnCall[i] = struct {
		result1 string
	}{result1}
}

func (fake *FakeSession) Values() map[string]string {
	fake.valuesMutex.Lock()
	ret, specificReturn := fake.valuesReturnsOnCall[len(fake.valuesArgsForCall)]
	fake.valuesArgsForCall = append(fake.valuesArgsForCall, struct {
	}{})
	fake.recordInvocation("Values", []interface{}{})
	fake.valuesMutex.Unlock()
	if fake.ValuesStub != nil {
		return fake.ValuesStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.valuesReturns
	return fakeReturns.result1
}

func (fake *FakeSession) ValuesCallCount() int {
	fake.valuesMutex.RLock()
	defer fake.valuesMutex.RUnlock()
	return len(fake.valuesArgsForCall)
}

func (fake *FakeSession) ValuesCalls(stub func() map[string]string) {
	fake.valuesMutex.Lock()
	defer fake.valuesMutex.Unlock()
	fake.ValuesStub = stub
}

func (fake *FakeSession) ValuesReturns(result1 map[string]string) {
	fake.valuesMutex.Lock()
	defer fake.valuesMutex.Unlock()
	fake.ValuesStub = nil
	fake.valuesReturns = struct {
		result1 map[string]string
	}{result1}
}

func (fake *FakeSession) ValuesReturnsOnCall(i int, result1 map[string]string) {
	fake.valuesMutex.Lock()
	defer fake.valuesMutex.Unlock()
	fake.ValuesStub = nil
	if fake.valuesReturnsOnCall == nil {
		fake.valuesReturnsOnCall = make(map[int]struct {
			result1 map[string]string
		})
	}
	fake.valuesReturnsOnCall[i] = struct {
		result1 map[string]string
	}{result1}
}

func (fake *FakeSession) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.clearMutex.RLock()
	defer fake.clearMutex.RUnlock()
	fake.nameMutex.RLock()
	defer fake.nameMutex.RUnlock()
	fake.valuesMutex.RLock()
	defer fake.valuesMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeSession) recordInvocation(key string, args []interface{}) {
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

var _ session.Session = new(FakeSession)
