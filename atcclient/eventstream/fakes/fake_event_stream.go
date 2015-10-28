// This file was generated by counterfeiter
package fakes

import (
	"sync"

	"github.com/concourse/atc"
	"github.com/concourse/fly/atcclient/eventstream"
)

type FakeEventStream struct {
	NextEventStub        func() (atc.Event, error)
	nextEventMutex       sync.RWMutex
	nextEventArgsForCall []struct{}
	nextEventReturns     struct {
		result1 atc.Event
		result2 error
	}
	CloseStub        func() error
	closeMutex       sync.RWMutex
	closeArgsForCall []struct{}
	closeReturns     struct {
		result1 error
	}
}

func (fake *FakeEventStream) NextEvent() (atc.Event, error) {
	fake.nextEventMutex.Lock()
	fake.nextEventArgsForCall = append(fake.nextEventArgsForCall, struct{}{})
	fake.nextEventMutex.Unlock()
	if fake.NextEventStub != nil {
		return fake.NextEventStub()
	} else {
		return fake.nextEventReturns.result1, fake.nextEventReturns.result2
	}
}

func (fake *FakeEventStream) NextEventCallCount() int {
	fake.nextEventMutex.RLock()
	defer fake.nextEventMutex.RUnlock()
	return len(fake.nextEventArgsForCall)
}

func (fake *FakeEventStream) NextEventReturns(result1 atc.Event, result2 error) {
	fake.NextEventStub = nil
	fake.nextEventReturns = struct {
		result1 atc.Event
		result2 error
	}{result1, result2}
}

func (fake *FakeEventStream) Close() error {
	fake.closeMutex.Lock()
	fake.closeArgsForCall = append(fake.closeArgsForCall, struct{}{})
	fake.closeMutex.Unlock()
	if fake.CloseStub != nil {
		return fake.CloseStub()
	} else {
		return fake.closeReturns.result1
	}
}

func (fake *FakeEventStream) CloseCallCount() int {
	fake.closeMutex.RLock()
	defer fake.closeMutex.RUnlock()
	return len(fake.closeArgsForCall)
}

func (fake *FakeEventStream) CloseReturns(result1 error) {
	fake.CloseStub = nil
	fake.closeReturns = struct {
		result1 error
	}{result1}
}

var _ eventstream.EventStream = new(FakeEventStream)
