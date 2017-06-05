// Code generated by counterfeiter. DO NOT EDIT.
package dbfakes

import (
	"sync"

	"github.com/concourse/atc/db"
)

type FakeContainerOwner struct {
	SetMapStub        func() map[string]interface{}
	setMapMutex       sync.RWMutex
	setMapArgsForCall []struct{}
	setMapReturns     struct {
		result1 map[string]interface{}
	}
	setMapReturnsOnCall map[int]struct {
		result1 map[string]interface{}
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeContainerOwner) SetMap() map[string]interface{} {
	fake.setMapMutex.Lock()
	ret, specificReturn := fake.setMapReturnsOnCall[len(fake.setMapArgsForCall)]
	fake.setMapArgsForCall = append(fake.setMapArgsForCall, struct{}{})
	fake.recordInvocation("SetMap", []interface{}{})
	fake.setMapMutex.Unlock()
	if fake.SetMapStub != nil {
		return fake.SetMapStub()
	}
	if specificReturn {
		return ret.result1
	}
	return fake.setMapReturns.result1
}

func (fake *FakeContainerOwner) SetMapCallCount() int {
	fake.setMapMutex.RLock()
	defer fake.setMapMutex.RUnlock()
	return len(fake.setMapArgsForCall)
}

func (fake *FakeContainerOwner) SetMapReturns(result1 map[string]interface{}) {
	fake.SetMapStub = nil
	fake.setMapReturns = struct {
		result1 map[string]interface{}
	}{result1}
}

func (fake *FakeContainerOwner) SetMapReturnsOnCall(i int, result1 map[string]interface{}) {
	fake.SetMapStub = nil
	if fake.setMapReturnsOnCall == nil {
		fake.setMapReturnsOnCall = make(map[int]struct {
			result1 map[string]interface{}
		})
	}
	fake.setMapReturnsOnCall[i] = struct {
		result1 map[string]interface{}
	}{result1}
}

func (fake *FakeContainerOwner) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.setMapMutex.RLock()
	defer fake.setMapMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeContainerOwner) recordInvocation(key string, args []interface{}) {
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

var _ db.ContainerOwner = new(FakeContainerOwner)
