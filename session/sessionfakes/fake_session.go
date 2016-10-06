// This file was generated by counterfeiter
package sessionfakes

import (
	"sync"

	"github.com/SAP/goauth_handlers/session"
)

type FakeSession struct {
	NameStub        func() string
	nameMutex       sync.RWMutex
	nameArgsForCall []struct{}
	nameReturns     struct {
		result1 string
	}
	ValuesStub        func() map[string]string
	valuesMutex       sync.RWMutex
	valuesArgsForCall []struct{}
	valuesReturns     struct {
		result1 map[string]string
	}
	ClearStub        func()
	clearMutex       sync.RWMutex
	clearArgsForCall []struct{}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeSession) Name() string {
	fake.nameMutex.Lock()
	fake.nameArgsForCall = append(fake.nameArgsForCall, struct{}{})
	fake.recordInvocation("Name", []interface{}{})
	fake.nameMutex.Unlock()
	if fake.NameStub != nil {
		return fake.NameStub()
	} else {
		return fake.nameReturns.result1
	}
}

func (fake *FakeSession) NameCallCount() int {
	fake.nameMutex.RLock()
	defer fake.nameMutex.RUnlock()
	return len(fake.nameArgsForCall)
}

func (fake *FakeSession) NameReturns(result1 string) {
	fake.NameStub = nil
	fake.nameReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeSession) Values() map[string]string {
	fake.valuesMutex.Lock()
	fake.valuesArgsForCall = append(fake.valuesArgsForCall, struct{}{})
	fake.recordInvocation("Values", []interface{}{})
	fake.valuesMutex.Unlock()
	if fake.ValuesStub != nil {
		return fake.ValuesStub()
	} else {
		return fake.valuesReturns.result1
	}
}

func (fake *FakeSession) ValuesCallCount() int {
	fake.valuesMutex.RLock()
	defer fake.valuesMutex.RUnlock()
	return len(fake.valuesArgsForCall)
}

func (fake *FakeSession) ValuesReturns(result1 map[string]string) {
	fake.ValuesStub = nil
	fake.valuesReturns = struct {
		result1 map[string]string
	}{result1}
}

func (fake *FakeSession) Clear() {
	fake.clearMutex.Lock()
	fake.clearArgsForCall = append(fake.clearArgsForCall, struct{}{})
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

func (fake *FakeSession) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.nameMutex.RLock()
	defer fake.nameMutex.RUnlock()
	fake.valuesMutex.RLock()
	defer fake.valuesMutex.RUnlock()
	fake.clearMutex.RLock()
	defer fake.clearMutex.RUnlock()
	return fake.invocations
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
