package gobincode

import (
	"encoding/binary"
	"errors"
	"io"
)

var (
	ErrInvalidType = errors.New("invalid type: unable to encode/decode")
)

type Encoder struct {
	*Config
	io.Writer
}

func (e *Encoder) EncodeInt(v *interface{}) error {
	return binary.Write(e.Writer, e.Config.ByteOrder, *v)
}
func (e *Encoder) EncodeElement(v *interface{}) error {
	switch (*v).(type) {
	case uint8:
	case uint16:
	case uint32:
	case uint64:
	case uint:
	case int8:
	case int16:
	case int32:
	case int64:
	case int:
		return e.EncodeInt(v)
	default:
		return ErrInvalidType
	}
	return nil
}

func (e *Encoder) Encode(v []interface{}) error {
	for _, el := range v {
		err := e.EncodeElement(&el)
		if err != nil {
			return err
		}
	}
	return nil
}
