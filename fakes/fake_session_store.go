// This file was generated by counterfeiter
package fakes

import (
	"net/http"
	"sync"

	"github.com/gorilla/sessions"
	"github.infra.hana.ondemand.com/cloudfoundry/goauth_handlers"
)

type FakeSessionStore struct {
	GetStub        func(r *http.Request, name string) (*sessions.Session, error)
	getMutex       sync.RWMutex
	getArgsForCall []struct {
		r    *http.Request
		name string
	}
	getReturns struct {
		result1 *sessions.Session
		result2 error
	}
	NewStub        func(r *http.Request, name string) (*sessions.Session, error)
	newMutex       sync.RWMutex
	newArgsForCall []struct {
		r    *http.Request
		name string
	}
	newReturns struct {
		result1 *sessions.Session
		result2 error
	}
	SaveStub        func(r *http.Request, w http.ResponseWriter, s *sessions.Session) error
	saveMutex       sync.RWMutex
	saveArgsForCall []struct {
		r *http.Request
		w http.ResponseWriter
		s *sessions.Session
	}
	saveReturns struct {
		result1 error
	}
}

func (fake *FakeSessionStore) Get(r *http.Request, name string) (*sessions.Session, error) {
	fake.getMutex.Lock()
	fake.getArgsForCall = append(fake.getArgsForCall, struct {
		r    *http.Request
		name string
	}{r, name})
	fake.getMutex.Unlock()
	if fake.GetStub != nil {
		return fake.GetStub(r, name)
	} else {
		return fake.getReturns.result1, fake.getReturns.result2
	}
}

func (fake *FakeSessionStore) GetCallCount() int {
	fake.getMutex.RLock()
	defer fake.getMutex.RUnlock()
	return len(fake.getArgsForCall)
}

func (fake *FakeSessionStore) GetArgsForCall(i int) (*http.Request, string) {
	fake.getMutex.RLock()
	defer fake.getMutex.RUnlock()
	return fake.getArgsForCall[i].r, fake.getArgsForCall[i].name
}

func (fake *FakeSessionStore) GetReturns(result1 *sessions.Session, result2 error) {
	fake.GetStub = nil
	fake.getReturns = struct {
		result1 *sessions.Session
		result2 error
	}{result1, result2}
}

func (fake *FakeSessionStore) New(r *http.Request, name string) (*sessions.Session, error) {
	fake.newMutex.Lock()
	fake.newArgsForCall = append(fake.newArgsForCall, struct {
		r    *http.Request
		name string
	}{r, name})
	fake.newMutex.Unlock()
	if fake.NewStub != nil {
		return fake.NewStub(r, name)
	} else {
		return fake.newReturns.result1, fake.newReturns.result2
	}
}

func (fake *FakeSessionStore) NewCallCount() int {
	fake.newMutex.RLock()
	defer fake.newMutex.RUnlock()
	return len(fake.newArgsForCall)
}

func (fake *FakeSessionStore) NewArgsForCall(i int) (*http.Request, string) {
	fake.newMutex.RLock()
	defer fake.newMutex.RUnlock()
	return fake.newArgsForCall[i].r, fake.newArgsForCall[i].name
}

func (fake *FakeSessionStore) NewReturns(result1 *sessions.Session, result2 error) {
	fake.NewStub = nil
	fake.newReturns = struct {
		result1 *sessions.Session
		result2 error
	}{result1, result2}
}

func (fake *FakeSessionStore) Save(r *http.Request, w http.ResponseWriter, s *sessions.Session) error {
	fake.saveMutex.Lock()
	fake.saveArgsForCall = append(fake.saveArgsForCall, struct {
		r *http.Request
		w http.ResponseWriter
		s *sessions.Session
	}{r, w, s})
	fake.saveMutex.Unlock()
	if fake.SaveStub != nil {
		return fake.SaveStub(r, w, s)
	} else {
		return fake.saveReturns.result1
	}
}

func (fake *FakeSessionStore) SaveCallCount() int {
	fake.saveMutex.RLock()
	defer fake.saveMutex.RUnlock()
	return len(fake.saveArgsForCall)
}

func (fake *FakeSessionStore) SaveArgsForCall(i int) (*http.Request, http.ResponseWriter, *sessions.Session) {
	fake.saveMutex.RLock()
	defer fake.saveMutex.RUnlock()
	return fake.saveArgsForCall[i].r, fake.saveArgsForCall[i].w, fake.saveArgsForCall[i].s
}

func (fake *FakeSessionStore) SaveReturns(result1 error) {
	fake.SaveStub = nil
	fake.saveReturns = struct {
		result1 error
	}{result1}
}

var _ goauth_handlers.SessionStore = new(FakeSessionStore)
