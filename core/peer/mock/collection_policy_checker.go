// Code generated by counterfeiter. DO NOT EDIT.
package mock

import (
	"sync"

	"github.com/hyperledger/fabric/core/ledger"
	"github.com/hyperledger/fabric/msp"
	"github.com/hyperledger/fabric/protoutil"
)

type CollectionPolicyChecker struct {
	CheckCollectionPolicyStub        func(blockNum uint64, ccName string, collName string, cfgHistoryRetriever ledger.ConfigHistoryRetriever, deserializer msp.IdentityDeserializer, signedData *protoutil.SignedData) (bool, error)
	checkCollectionPolicyMutex       sync.RWMutex
	checkCollectionPolicyArgsForCall []struct {
		blockNum            uint64
		ccName              string
		collName            string
		cfgHistoryRetriever ledger.ConfigHistoryRetriever
		deserializer        msp.IdentityDeserializer
		signedData          *protoutil.SignedData
	}
	checkCollectionPolicyReturns struct {
		result1 bool
		result2 error
	}
	checkCollectionPolicyReturnsOnCall map[int]struct {
		result1 bool
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *CollectionPolicyChecker) CheckCollectionPolicy(blockNum uint64, ccName string, collName string, cfgHistoryRetriever ledger.ConfigHistoryRetriever, deserializer msp.IdentityDeserializer, signedData *protoutil.SignedData) (bool, error) {
	fake.checkCollectionPolicyMutex.Lock()
	ret, specificReturn := fake.checkCollectionPolicyReturnsOnCall[len(fake.checkCollectionPolicyArgsForCall)]
	fake.checkCollectionPolicyArgsForCall = append(fake.checkCollectionPolicyArgsForCall, struct {
		blockNum            uint64
		ccName              string
		collName            string
		cfgHistoryRetriever ledger.ConfigHistoryRetriever
		deserializer        msp.IdentityDeserializer
		signedData          *protoutil.SignedData
	}{blockNum, ccName, collName, cfgHistoryRetriever, deserializer, signedData})
	fake.recordInvocation("CheckCollectionPolicy", []interface{}{blockNum, ccName, collName, cfgHistoryRetriever, deserializer, signedData})
	fake.checkCollectionPolicyMutex.Unlock()
	if fake.CheckCollectionPolicyStub != nil {
		return fake.CheckCollectionPolicyStub(blockNum, ccName, collName, cfgHistoryRetriever, deserializer, signedData)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.checkCollectionPolicyReturns.result1, fake.checkCollectionPolicyReturns.result2
}

func (fake *CollectionPolicyChecker) CheckCollectionPolicyCallCount() int {
	fake.checkCollectionPolicyMutex.RLock()
	defer fake.checkCollectionPolicyMutex.RUnlock()
	return len(fake.checkCollectionPolicyArgsForCall)
}

func (fake *CollectionPolicyChecker) CheckCollectionPolicyArgsForCall(i int) (uint64, string, string, ledger.ConfigHistoryRetriever, msp.IdentityDeserializer, *protoutil.SignedData) {
	fake.checkCollectionPolicyMutex.RLock()
	defer fake.checkCollectionPolicyMutex.RUnlock()
	return fake.checkCollectionPolicyArgsForCall[i].blockNum, fake.checkCollectionPolicyArgsForCall[i].ccName, fake.checkCollectionPolicyArgsForCall[i].collName, fake.checkCollectionPolicyArgsForCall[i].cfgHistoryRetriever, fake.checkCollectionPolicyArgsForCall[i].deserializer, fake.checkCollectionPolicyArgsForCall[i].signedData
}

func (fake *CollectionPolicyChecker) CheckCollectionPolicyReturns(result1 bool, result2 error) {
	fake.CheckCollectionPolicyStub = nil
	fake.checkCollectionPolicyReturns = struct {
		result1 bool
		result2 error
	}{result1, result2}
}

func (fake *CollectionPolicyChecker) CheckCollectionPolicyReturnsOnCall(i int, result1 bool, result2 error) {
	fake.CheckCollectionPolicyStub = nil
	if fake.checkCollectionPolicyReturnsOnCall == nil {
		fake.checkCollectionPolicyReturnsOnCall = make(map[int]struct {
			result1 bool
			result2 error
		})
	}
	fake.checkCollectionPolicyReturnsOnCall[i] = struct {
		result1 bool
		result2 error
	}{result1, result2}
}

func (fake *CollectionPolicyChecker) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.checkCollectionPolicyMutex.RLock()
	defer fake.checkCollectionPolicyMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *CollectionPolicyChecker) recordInvocation(key string, args []interface{}) {
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
