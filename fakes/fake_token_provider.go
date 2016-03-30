// This file was generated by counterfeiter
package fakes

import (
	"sync"

	"github.wdf.sap.corp/cloudfoundry/goauth_handlers"
	"golang.org/x/oauth2"
)

type FakeTokenProvider struct {
	RequestTokenStub        func(code string) (*oauth2.Token, error)
	requestTokenMutex       sync.RWMutex
	requestTokenArgsForCall []struct {
		code string
	}
	requestTokenReturns struct {
		result1 *oauth2.Token
		result2 error
	}
	LoginURLStub        func(state string) string
	loginURLMutex       sync.RWMutex
	loginURLArgsForCall []struct {
		state string
	}
	loginURLReturns struct {
		result1 string
	}
}

func (fake *FakeTokenProvider) RequestToken(code string) (*oauth2.Token, error) {
	fake.requestTokenMutex.Lock()
	fake.requestTokenArgsForCall = append(fake.requestTokenArgsForCall, struct {
		code string
	}{code})
	fake.requestTokenMutex.Unlock()
	if fake.RequestTokenStub != nil {
		return fake.RequestTokenStub(code)
	} else {
		return fake.requestTokenReturns.result1, fake.requestTokenReturns.result2
	}
}

func (fake *FakeTokenProvider) RequestTokenCallCount() int {
	fake.requestTokenMutex.RLock()
	defer fake.requestTokenMutex.RUnlock()
	return len(fake.requestTokenArgsForCall)
}

func (fake *FakeTokenProvider) RequestTokenArgsForCall(i int) string {
	fake.requestTokenMutex.RLock()
	defer fake.requestTokenMutex.RUnlock()
	return fake.requestTokenArgsForCall[i].code
}

func (fake *FakeTokenProvider) RequestTokenReturns(result1 *oauth2.Token, result2 error) {
	fake.RequestTokenStub = nil
	fake.requestTokenReturns = struct {
		result1 *oauth2.Token
		result2 error
	}{result1, result2}
}

func (fake *FakeTokenProvider) LoginURL(state string) string {
	fake.loginURLMutex.Lock()
	fake.loginURLArgsForCall = append(fake.loginURLArgsForCall, struct {
		state string
	}{state})
	fake.loginURLMutex.Unlock()
	if fake.LoginURLStub != nil {
		return fake.LoginURLStub(state)
	} else {
		return fake.loginURLReturns.result1
	}
}

func (fake *FakeTokenProvider) LoginURLCallCount() int {
	fake.loginURLMutex.RLock()
	defer fake.loginURLMutex.RUnlock()
	return len(fake.loginURLArgsForCall)
}

func (fake *FakeTokenProvider) LoginURLArgsForCall(i int) string {
	fake.loginURLMutex.RLock()
	defer fake.loginURLMutex.RUnlock()
	return fake.loginURLArgsForCall[i].state
}

func (fake *FakeTokenProvider) LoginURLReturns(result1 string) {
	fake.LoginURLStub = nil
	fake.loginURLReturns = struct {
		result1 string
	}{result1}
}

var _ goauth_handlers.TokenProvider = new(FakeTokenProvider)