// This file was generated by counterfeiter
package goauth_handlersfakes

import (
	"sync"

	"github.com/SAP/goauth-handlers"
	"github.com/SAP/goauth-handlers/token"
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
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeTokenDecoder) Decode(arg1 *oauth2.Token) (token.Info, error) {
	fake.decodeMutex.Lock()
	fake.decodeArgsForCall = append(fake.decodeArgsForCall, struct {
		arg1 *oauth2.Token
	}{arg1})
	fake.recordInvocation("Decode", []interface{}{arg1})
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

func (fake *FakeTokenDecoder) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.decodeMutex.RLock()
	defer fake.decodeMutex.RUnlock()
	return fake.invocations
}

func (fake *FakeTokenDecoder) recordInvocation(key string, args []interface{}) {
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

var _ goauth_handlers.TokenDecoder = new(FakeTokenDecoder)
