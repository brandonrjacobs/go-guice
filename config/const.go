package config

type FileType int

const (

	FileTypeYaml FileType = iota
	FileTypeJson
	FileTypeToml

)

// ToViper converts our internal FileType values to spf13/viper formats
func (f FileType) ToViper() string {
	switch f {
	case 0:
		return "yaml"
	case 1:
		return "json"
	case 2:
		return "toml"
	default:
		return "yaml"
	}
}

// Default Config Values

var (
	Development Key = "development"
	LogLevel Key = "log.level"
	LogFormat Key = "log.format"
	LogSamplingRate Key = "log.samplingRate"
	LogSampleEvery Key = "log.sampleEvery"
)
