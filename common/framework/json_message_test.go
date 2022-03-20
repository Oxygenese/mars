package framework

import "testing"
import (
	"encoding/base64"
	"math/rand"
	"time"
)

const (
	testRepeat = 10
)

var generator = rand.New(rand.NewSource(time.Now().UnixNano()))

var checkConsistency = func(t *testing.T, f func(msg *JsonMessage)) {
	for i := 0; i < testRepeat; i++ {
		origin, err := generateMessage()
		if err != nil {
			t.Fatalf("generate message fail: %s", err.Error())
		}
		f(origin)
		if err != nil {
			t.Fatalf("parse message fail: %s", err.Error())
		}
		t.Logf("%dth try success", i+1)
	}
}

func Test_BaseMember(t *testing.T) {
	var empty = func(msg *JsonMessage) {}
	checkConsistency(t, empty)
}

func Test_BoolParam(t *testing.T) {
	const paramCount = 5
	var prepare = func(msg *JsonMessage) {
		generateBoolParam(msg, paramCount)
	}
	checkConsistency(t, prepare)
}

func Test_StringParam(t *testing.T) {
	const paramCount = 5
	var prepare = func(msg *JsonMessage) {
		generateStringParam(msg, paramCount)
	}
	checkConsistency(t, prepare)
}

func Test_UIntParam(t *testing.T) {
	const paramCount = 5
	var prepare = func(msg *JsonMessage) {
		generateUIntParam(msg, paramCount)
	}
	checkConsistency(t, prepare)
}

func Test_IntParam(t *testing.T) {
	const paramCount = 5
	var prepare = func(msg *JsonMessage) {
		generateIntParam(msg, paramCount)
	}
	checkConsistency(t, prepare)
}

func Test_FloatParam(t *testing.T) {
	const paramCount = 5
	var prepare = func(msg *JsonMessage) {
		generateFloatParam(msg, paramCount)
	}
	checkConsistency(t, prepare)
}

func Test_UIntArrayParam(t *testing.T) {
	const paramCount = 5
	var prepare = func(msg *JsonMessage) {
		generateUIntArrayParam(msg, paramCount)
	}
	checkConsistency(t, prepare)
}

func Test_StringArrayParam(t *testing.T) {
	const paramCount = 5
	var prepare = func(msg *JsonMessage) {
		generateStringArrayParam(msg, paramCount)
	}
	checkConsistency(t, prepare)
}

func Test_MixedParam(t *testing.T) {
	const paramCount = 3
	var prepare = func(msg *JsonMessage) {
		generateBoolParam(msg, paramCount)
		generateUIntParam(msg, paramCount)
		generateIntParam(msg, paramCount)
		generateStringParam(msg, paramCount)
		generateFloatParam(msg, paramCount)
		generateUIntArrayParam(msg, paramCount)
		generateStringArrayParam(msg, paramCount)
	}
	checkConsistency(t, prepare)
}

func Test_CloneMessage(t *testing.T) {
	const (
		paramCount = 3
	)
	for i := 0; i < testRepeat; i++ {
		origin, err := generateMessage()
		if err != nil {
			t.Fatalf("generate message fail: %s", err.Error())
		}
		generateBoolParam(origin, paramCount)
		generateUIntParam(origin, paramCount)
		generateIntParam(origin, paramCount)
		generateStringParam(origin, paramCount)
		generateFloatParam(origin, paramCount)
		generateUIntArrayParam(origin, paramCount)
		generateStringArrayParam(origin, paramCount)
		if err != nil {
			t.Fatalf("clone message fail: %s", err.Error())
		}
		if err != nil {
			t.Fatalf("compare %d clone fail: %s", i, err.Error())
		}
		t.Logf("%dth clone is identical", i)
	}

}

func generateMessage() (msg *JsonMessage, err error) {
	const (
		ErrorLength = 32
	)
	msg, err = NewJsonMessage(generator.Uint32())
	msg.SetTransactionID(generator.Uint32())
	if generator.Intn(2) > 0 {
		msg.SetSuccess(true)
	} else {
		var buf = make([]byte, ErrorLength)
		_, err = generator.Read(buf)
		if err != nil {
			return
		}
		msg.SetError(base64.StdEncoding.EncodeToString(buf))
		msg.SetSuccess(false)
	}
	return msg, nil
}

func generateBoolParam(msg Message, count int) {
	for i := 0; i < count; i++ {
		var key = generator.Uint32()
		if generator.Intn(2) > 0 {
			msg.SetBoolean(key, true)
		} else {
			msg.SetBoolean(key, false)
		}
	}
}

func generateStringParam(msg Message, count int) {
	const (
		MinStringLength = 6
		MaxStringLength = 20
	)
	var bufSize = MinStringLength + generator.Intn(MaxStringLength-MinStringLength)
	var buf = make([]byte, bufSize)
	for i := 0; i < count; i++ {
		generator.Read(buf)
		var key = generator.Uint32()
		msg.SetString(key, base64.StdEncoding.EncodeToString(buf))
	}
}

func generateUIntParam(msg Message, count int) {
	for i := 0; i < count; i++ {
		var key = generator.Uint32()
		var value = generator.Uint32()
		msg.SetUInt(key, value)
	}
}

func generateIntParam(msg Message, count int) {
	for i := 0; i < count; i++ {
		var key = generator.Uint32()
		var value = generator.Int63()
		msg.SetInt(key, value)
	}
}

func generateFloatParam(msg Message, count int) {
	for i := 0; i < count; i++ {
		var key = generator.Uint32()
		var value = generator.Float32()
		msg.SetFloat(key, value)
	}
}

func generateStringArrayParam(msg Message, count int) {
	const (
		MinArrayLength  = 1
		MaxArrayLength  = 10
		MinStringLength = 6
		MaxStringLength = 20
	)
	var arrayLength = MinArrayLength + generator.Intn(MaxArrayLength-MinArrayLength)
	var bufSize = MinStringLength + generator.Intn(MaxStringLength-MinStringLength)
	var buf = make([]byte, bufSize)
	for i := 0; i < count; i++ {
		var value = make([]string, arrayLength)
		for j := 0; j < arrayLength; j++ {
			generator.Read(buf)
			value[j] = base64.StdEncoding.EncodeToString(buf)
		}
		var key = generator.Uint32()
		msg.SetStringArray(key, value)
	}
}

func generateUIntArrayParam(msg Message, count int) {
	const (
		MinArrayLength = 1
		MaxArrayLength = 10
	)
	var size = MinArrayLength + generator.Intn(MaxArrayLength-MinArrayLength)
	for i := 0; i < count; i++ {
		var value = make([]uint64, size)
		for j := 0; j < size; j++ {
			value[j] = generator.Uint64()
		}
		var key = generator.Uint32()
		msg.SetUIntArray(key, value)
	}
}
