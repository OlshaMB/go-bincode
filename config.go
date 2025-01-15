package gobincode;

type Config struct {
	LittleEndian bool
	VariantIntEncoding bool
}

func DefaultConfig() *Config {
	return &Config{
		LittleEndian: true,
		VariantIntEncoding: true,
	}
}