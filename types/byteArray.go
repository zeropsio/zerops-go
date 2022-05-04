// Generated ZEROPS sdk

package types

type ByteArray []byte

func NewByteArray(value []byte) ByteArray {
	return ByteArray(value)
}

func NewByteArrayFromString(value string) (out ByteArray, err error) {
	return ByteArray(value), nil
}
func (parameter ByteArray) Native() []byte {
	return []byte(parameter)
}
