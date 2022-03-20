package framework

import (
	"encoding/json"
	"fmt"
	"github.com/mars-projects/mars/api/chief"
)

type JsonMessage struct {
	*chief.Message
}

func NewTransJsonMessage(message *chief.Message) *JsonMessage {
	return &JsonMessage{message}
}

func NewJsonMessage(msg uint32) (*JsonMessage, error) {
	return &JsonMessage{
		Message: &chief.Message{Id: msg},
	}, nil
}

func CloneJsonMessage(origin Message) (clone *JsonMessage) {
	clone, _ = NewJsonMessage(origin.GetID())
	clone.SetSuccess(origin.IsSuccess())
	clone.SetTransactionID(origin.GetTransactionID())
	if "" != origin.GetError() {
		clone.SetError(origin.GetError())
	}
	//clone params
	if 0 != len(origin.GetAllBoolean()) {
		for key, value := range origin.GetAllBoolean() {
			clone.SetBoolean(key, value)
		}
	}
	if 0 != len(origin.GetAllString()) {
		for key, value := range origin.GetAllString() {
			clone.SetString(key, value)
		}
	}
	if 0 != len(origin.GetAllUInt()) {
		for key, value := range origin.GetAllUInt() {
			clone.SetUInt(key, value)
		}
	}
	if 0 != len(origin.GetAllInt()) {
		for key, value := range origin.GetAllInt() {
			clone.SetInt(key, value)
		}
	}
	if 0 != len(origin.GetAllFloat()) {
		for key, value := range origin.GetAllFloat() {
			clone.SetFloat(key, value)
		}
	}
	if 0 != len(origin.GetAllUIntArray()) {
		for key, value := range origin.GetAllUIntArray() {
			clone.SetUIntArray(key, value)
		}
	}
	if 0 != len(origin.GetAllStringArray()) {
		for key, value := range origin.GetAllStringArray() {
			clone.SetStringArray(key, value)
		}
	}
	return clone
}

func MessageFromJson(data []byte) (*JsonMessage, error) {
	var msg JsonMessage
	var err = json.Unmarshal(data, &msg)
	return &msg, err
}

func (msg *JsonMessage) GetID() uint32 {
	return msg.Id
}

func (msg *JsonMessage) SetID(id uint32) {
	msg.Id = id
}

func (msg *JsonMessage) IsSuccess() bool {
	return msg.Success
}

func (msg *JsonMessage) SetSuccess(flag bool) {
	msg.Success = flag
}

func (msg *JsonMessage) SetSender(value string) {
	msg.Sender = value
}

func (msg *JsonMessage) GetSender() string {
	return msg.Sender
}

func (msg *JsonMessage) SetTransactionID(id uint32) {
	msg.Transaction = id
}
func (msg *JsonMessage) GetTransactionID() uint32 {
	return msg.Transaction
}

func (msg *JsonMessage) SetError(err string) {
	msg.Error = err
}
func (msg *JsonMessage) GetError() string {
	return msg.Error
}

func (msg *JsonMessage) GetString(key uint32) (string, error) {
	if msg.StringParams != nil {
		if value, exists := msg.StringParams[key]; exists {
			return value, nil
		}
	}
	return "", fmt.Errorf("no string param for key %d", key)
}

func (msg *JsonMessage) GetUInt(key uint32) (uint, error) {
	if msg.UintParams != nil {
		if value, exists := msg.UintParams[key]; exists {
			return uint(value), nil
		}
	}
	return 0, fmt.Errorf("no uint param for key %d", key)
}

func (msg *JsonMessage) GetInt(key uint32) (int, error) {
	if msg.IntParams != nil {
		if value, exists := msg.IntParams[key]; exists {
			return int(value), nil
		}
	}
	return 0, fmt.Errorf("no int param for key %d", key)
}

func (msg *JsonMessage) GetFloat(key uint32) (float64, error) {
	if msg.FloatParams != nil {
		if value, exists := msg.FloatParams[key]; exists {
			return float64(value), nil
		}
	}
	return 0.0, fmt.Errorf("no float param for key %d", key)
}

func (msg *JsonMessage) GetBoolean(key uint32) (bool, error) {
	if msg.BoolParams != nil {
		if value, exists := msg.BoolParams[key]; exists {
			return value, nil
		}
	}
	return false, fmt.Errorf("no bool param for key %d", key)
}

func (msg *JsonMessage) SetString(key uint32, value string) {
	if msg.StringParams != nil {
		msg.StringParams[key] = value
	} else {
		msg.StringParams = map[uint32]string{key: value}
	}
}

func (msg *JsonMessage) SetUInt(key uint32, value uint32) {
	if msg.UintParams != nil {
		msg.UintParams[key] = value
	} else {
		msg.UintParams = map[uint32]uint32{key: value}
	}
}

func (msg *JsonMessage) SetInt(key uint32, value int64) {
	if msg.IntParams != nil {
		msg.IntParams[key] = value
	} else {
		msg.IntParams = map[uint32]int64{key: value}
	}
}

func (msg *JsonMessage) SetFloat(key uint32, value float32) {
	if msg.FloatParams != nil {
		msg.FloatParams[key] = value
	} else {
		msg.FloatParams = map[uint32]float32{key: value}
	}
}

func (msg *JsonMessage) SetBoolean(key uint32, value bool) {
	if msg.BoolParams != nil {
		msg.BoolParams[key] = value
	} else {
		msg.BoolParams = map[uint32]bool{key: value}
	}
}

func (msg *JsonMessage) GetFromSession() uint32 {
	return msg.From
}

func (msg *JsonMessage) SetFromSession(session uint32) {
	msg.From = session
}

func (msg *JsonMessage) GetToSession() uint32 {
	return msg.To
}

func (msg *JsonMessage) SetToSession(session uint32) {
	msg.To = session
}

func (msg *JsonMessage) SetUIntArray(key uint32, value []uint64) {
	if msg.ArrayParams == nil {
		msg.ArrayParams = make(map[uint32]*chief.ArrayParams, 0)
	}
	msg.ArrayParams[key] = &chief.ArrayParams{UintArray: value}
}

func (msg *JsonMessage) SetStringArray(key uint32, value []string) {
	if msg.ArrayParams == nil {
		msg.ArrayParams = make(map[uint32]*chief.ArrayParams, 0)
	}
	msg.ArrayParams[key] = &chief.ArrayParams{StringArray: value}
}

func (msg *JsonMessage) GetUIntArray(key uint32) ([]uint64, error) {
	if msg.ArrayParams == nil {
		return nil, fmt.Errorf("no uint array for key %d", key)
	}
	if msg.ArrayParams[key].UintArray != nil {
		if value, exists := msg.ArrayParams[key]; exists {
			return value.GetUintArray(), nil
		}
	}
	return nil, fmt.Errorf("no uint array for key %d", key)
}

func (msg *JsonMessage) GetStringArray(key uint32) ([]string, error) {
	if msg.ArrayParams == nil {
		return nil, fmt.Errorf("no string array for key %d", key)
	}
	if msg.ArrayParams[key].StringArray != nil {
		if value, exists := msg.ArrayParams[key]; exists {
			return value.StringArray, nil
		}
	}
	return nil, fmt.Errorf("no string array for key %d", key)
}

func (msg *JsonMessage) Serialize() ([]byte, error) {
	return json.Marshal(&msg)
}

var (
	emptyFloatMap       = map[uint32]float32{}
	emptyStringMap      = map[uint32]string{}
	emptyUIntMap        = map[uint32]uint32{}
	emptyIntMap         = map[uint32]int64{}
	emptyBooleanMap     = map[uint32]bool{}
	emptyStringArrayMap = map[uint32][]string{}
	emptyUIntArrayMap   = map[uint32][]uint64{}
)

func (msg *JsonMessage) GetAllString() map[uint32]string {
	if msg.StringParams != nil {
		return msg.StringParams
	}
	return emptyStringMap
}
func (msg *JsonMessage) GetAllUInt() map[uint32]uint32 {
	if msg.UintParams != nil {
		return msg.UintParams
	}
	return emptyUIntMap
}

func (msg *JsonMessage) GetAllInt() map[uint32]int64 {
	if msg.IntParams != nil {
		return msg.IntParams
	}
	return emptyIntMap
}

func (msg *JsonMessage) GetAllFloat() map[uint32]float32 {
	if msg.FloatParams != nil {
		return msg.FloatParams
	}
	return emptyFloatMap
}

func (msg *JsonMessage) GetAllBoolean() map[uint32]bool {
	if msg.BoolParams != nil {
		return msg.BoolParams
	}
	return emptyBooleanMap
}

func (msg *JsonMessage) GetAllUIntArray() map[uint32][]uint64 {
	if msg.ArrayParams == nil {
		m := make(map[uint32][]uint64)
		for k, v := range msg.ArrayParams {
			m[k] = v.UintArray
		}
		return m
	}
	return emptyUIntArrayMap
}

func (msg *JsonMessage) GetAllStringArray() map[uint32][]string {
	if msg.ArrayParams != nil {
		m := make(map[uint32][]string)
		for k, v := range msg.ArrayParams {
			m[k] = v.StringArray
		}
		return m
	}
	return emptyStringArrayMap
}
