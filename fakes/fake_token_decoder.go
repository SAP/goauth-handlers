// This file was generated by counterfeiter
package fakes

import (
	"sync"

	"github.wdf.sap.corp/cloudfoundry/goauth_handlers"
	"github.wdf.sap.corp/cloudfoundry/goauth_handlers/token"
	"golang.org/x/oauth2"
)

type FakeTokenDecoder struct {
	DecodeStub        func(*oauth2.Token) (token.Info, error)
	decodeMutex       sync.RWMutex
	decodeArgsForCall []struct {
		arg1 *oauth2.Token
	}
	decodeReturns struct {
		result1 token.Info
		result2 error
	}
}

func (fake *FakeTokenDecoder) Decode(arg1 *oauth2.Token) (token.Info, error) {
	fake.decodeMutex.Lock()
	fake.decodeArgsForCall = append(fake.decodeArgsForCall, struct {
		arg1 *oauth2.Token
	}{arg1})
	fake.decodeMutex.Unlock()
	if fake.DecodeStub != nil {
		return fake.DecodeStub(arg1)
	} else {
		return fake.decodeReturns.result1, fake.decodeReturns.result2
	}
}

func (fake *FakeTokenDecoder) DecodeCallCount() int {
	fake.decodeMutex.RLock()
	defer fake.decodeMutex.RUnlock()
	return len(fake.decodeArgsForCall)
}

func (fake *FakeTokenDecoder) DecodeArgsForCall(i int) *oauth2.Token {
	fake.decodeMutex.RLock()
	defer fake.decodeMutex.RUnlock()
	return fake.decodeArgsForCall[i].arg1
}

func (fake *FakeTokenDecoder) DecodeReturns(result1 token.Info, result2 error) {
	fake.DecodeStub = nil
	fake.decodeReturns = struct {
		result1 token.Info
		result2 error
	}{result1, result2}
}

var _ goauth_handlers.TokenDecoder = new(FakeTokenDecoder)
