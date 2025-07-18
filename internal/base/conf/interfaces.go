package conf

import (
	"io"
	"os"
	"strings"
	"time"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/afero"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

// Config is an interface for processing configuration file(s).
type Config interface {
	AddConfigPath(in string)
	AddRemoteProvider(provider, endpoint, path string) error
	AddSecureRemoteProvider(provider, endpoint, path, secretkeyring string) error
	AllKeys() []string
	AllSettings() map[string]any
	AllowEmptyEnv(allowEmptyEnv bool)
	AutomaticEnv()
	BindEnv(input ...string) error
	BindPFlag(key string, flag *pflag.Flag) error
	BindPFlags(flags *pflag.FlagSet) error
	ConfigFileUsed() string
	Debug()
	DebugTo(w io.Writer)
	Get(key string) any
	GetBool(key string) bool
	GetDuration(key string) time.Duration
	GetEnvPrefix() string
	GetFloat64(key string) float64
	GetInt(key string) int
	GetInt32(key string) int32
	GetInt64(key string) int64
	GetIntSlice(key string) []int
	GetSizeInBytes(key string) uint
	GetString(key string) string
	GetStringMap(key string) map[string]any
	GetStringMapString(key string) map[string]string
	GetStringMapStringSlice(key string) map[string][]string
	GetStringSlice(key string) []string
	GetTime(key string) time.Time
	GetUint(key string) uint
	GetUint16(key string) uint16
	GetUint32(key string) uint32
	GetUint64(key string) uint64
	GetUint8(key string) uint8
	InConfig(key string) bool
	IsSet(key string) bool
	MergeConfig(in io.Reader) error
	MergeConfigMap(cfg map[string]any) error
	MergeInConfig() error
	MustBindEnv(input ...string)
	OnConfigChange(run func(in fsnotify.Event))
	ReadConfig(in io.Reader) error
	ReadInConfig() error
	ReadRemoteConfig() error
	RegisterAlias(alias, key string)
	SafeWriteConfig() error
	SafeWriteConfigAs(filename string) error
	Set(key string, value any)
	SetConfigFile(in string)
	SetConfigName(in string)
	SetConfigPermissions(perm os.FileMode)
	SetConfigType(in string)
	SetDefault(key string, value any)
	SetEnvKeyReplacer(r *strings.Replacer)
	SetEnvPrefix(in string)
	SetFs(fs afero.Fs)
	SetTypeByDefaultValue(enable bool)
	Unmarshal(rawVal any, opts ...viper.DecoderConfigOption) error
	UnmarshalExact(rawVal any, opts ...viper.DecoderConfigOption) error
	UnmarshalKey(key string, rawVal any, opts ...viper.DecoderConfigOption) error
	WatchConfig()
	WatchRemoteConfig() error
	WatchRemoteConfigOnChannel() error
	WriteConfig() error
	WriteConfigAs(filename string) error
	WriteConfigTo(w io.Writer) error
}
