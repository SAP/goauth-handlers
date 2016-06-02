// This file was generated by counterfeiter
package sessionfakes

import (
	"net/http"
	"sync"

	"github.infra.hana.ondemand.com/cloudfoundry/goauth_handlers/session"
)

type FakeStore struct {
	GetStub        func(req *http.Request, name string) (session.Session, error)
	getMutex       sync.RWMutex
	getArgsForCall []struct {
		req  *http.Request
		name string
	}
	getReturns struct {
		result1 session.Session
		result2 error
	}
	SaveStub        func(resp http.ResponseWriter, s session.Session) error
	saveMutex       sync.RWMutex
	saveArgsForCall []struct {
		resp http.ResponseWriter
		s    session.Session
	}
	saveReturns struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeStore) Get(req *http.Request, name string) (session.Session, error) {
	fake.getMutex.Lock()
	fake.getArgsForCall = append(fake.getArgsForCall, struct {
		req  *http.Request
		name string
	}{req, name})
	fake.recordInvocation("Get", []interface{}{req, name})
	fake.getMutex.Unlock()
	if fake.GetStub != nil {
		return fake.GetStub(req, name)
	} else {
		return fake.getReturns.result1, fake.getReturns.result2
	}
}

func (fake *FakeStore) GetCallCount() int {
	fake.getMutex.RLock()
	defer fake.getMutex.RUnlock()
	return len(fake.getArgsForCall)
}

func (fake *FakeStore) GetArgsForCall(i int) (*http.Request, string) {
	fake.getMutex.RLock()
	defer fake.getMutex.RUnlock()
	return fake.getArgsForCall[i].req, fake.getArgsForCall[i].name
}

func (fake *FakeStore) GetReturns(result1 session.Session, result2 error) {
	fake.GetStub = nil
	fake.getReturns = struct {
		result1 session.Session
		result2 error
	}{result1, result2}
}

func (fake *FakeStore) Save(resp http.ResponseWriter, s session.Session) error {
	fake.saveMutex.Lock()
	fake.saveArgsForCall = append(fake.saveArgsForCall, struct {
		resp http.ResponseWriter
		s    session.Session
	}{resp, s})
	fake.recordInvocation("Save", []interface{}{resp, s})
	fake.saveMutex.Unlock()
	if fake.SaveStub != nil {
		return fake.SaveStub(resp, s)
	} else {
		return fake.saveReturns.result1
	}
}

func (fake *FakeStore) SaveCallCount() int {
	fake.saveMutex.RLock()
	defer fake.saveMutex.RUnlock()
	return len(fake.saveArgsForCall)
}

func (fake *FakeStore) SaveArgsForCall(i int) (http.ResponseWriter, session.Session) {
	fake.saveMutex.RLock()
	defer fake.saveMutex.RUnlock()
	return fake.saveArgsForCall[i].resp, fake.saveArgsForCall[i].s
}

func (fake *FakeStore) SaveReturns(result1 error) {
	fake.SaveStub = nil
	fake.saveReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeStore) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.getMutex.RLock()
	defer fake.getMutex.RUnlock()
	fake.saveMutex.RLock()
	defer fake.saveMutex.RUnlock()
	return fake.invocations
}

func (fake *FakeStore) recordInvocation(key string, args []interface{}) {
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

var _ session.Store = new(FakeStore)
