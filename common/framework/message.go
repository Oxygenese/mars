package framework

type Message interface {
	GetID() uint32
	SetID(id uint32)
	IsSuccess() bool
	SetSuccess(flag bool)
	SetSender(string)
	GetSender() string
	SetTransactionID(id uint32)
	GetTransactionID() uint32

	SetError(msg string)
	GetError() string

	GetString(key uint32) (string, error)
	GetUInt(key uint32) (uint, error)
	GetInt(key uint32) (int, error)
	GetFloat(key uint32) (float64, error)
	GetBoolean(key uint32) (bool, error)
	SetString(key uint32, value string)
	SetUInt(key uint32, value uint32)
	SetInt(key uint32, value int64)
	SetFloat(key uint32, value float32)
	SetBoolean(key uint32, value bool)

	SetUIntArray(key uint32, value []uint64)
	GetUIntArray(key uint32) ([]uint64, error)

	SetStringArray(key uint32, value []string)
	GetStringArray(key uint32) ([]string, error)

	Serialize() ([]byte, error)

	GetFromSession() uint32
	SetFromSession(session uint32)
	GetToSession() uint32
	SetToSession(session uint32)

	GetAllString() map[uint32]string
	GetAllUInt() map[uint32]uint32
	GetAllInt() map[uint32]int64
	GetAllFloat() map[uint32]float32
	GetAllBoolean() map[uint32]bool
	GetAllUIntArray() map[uint32][]uint64
	GetAllStringArray() map[uint32][]string
}
