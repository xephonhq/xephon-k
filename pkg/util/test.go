package util

import "sync/atomic"

var dummyVar = "dummy"
var mockID int32
var mockedVars map[int32]interface{}
var mockedVarsValue map[int32]interface{}

// FIXME: these won't work
//func MockStringVar(v *interface{}, val interface{}) int32{
//func RecoverMockedVar(mockID int32) {
//	*mockedVars[mockID] = mockedVarsValue[mockID]
//}

func MockStringVar(v *string, val string) int32 {
	atomic.AddInt32(&mockID, 1)
	mockedVars[mockID] = v
	mockedVarsValue[mockID] = *v
	*v = val
	return mockID
}

func RecoverMockedStringVar(mockID int32) {
	//log.Info(mockedVarsValue[mockID])
	*mockedVars[mockID].(*string) = mockedVarsValue[mockID].(string)
}

// TODO: use bench/generator to generate fake series for testing

func init() {
	mockedVars = make(map[int32]interface{}, 5)
	mockedVarsValue = make(map[int32]interface{}, 5)
}
