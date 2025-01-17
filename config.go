package gobincode

import "encoding/binary"

type Config struct {
	binary.ByteOrder
	VariantIntEncoding bool
}

func DefaultConfig() *Config {
	return &Config{
		ByteOrder:          binary.LittleEndian,
		VariantIntEncoding: true,
	}
}
