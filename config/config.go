package config

import (
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	GRPCServer     GRPCServerConfig     `validate:"required"`
	GRPCApi        GRPCApiConfig        `validate:"required"`
	GRPCClient     GRPCClientConfig     `validate:"required"`
	Infrastructure InfrastructureConfig `validate:"required"`
}

type GRPCServerConfig struct {
	Address              string        `env:"GRPC_SERVER_ADDRESS" validate:"required"`
	MaxRecvMsgSize       int           `env:"GRPC_SERVER_MAX_RECV_MSG_SIZE" validate:"gte=0"`
	MaxSendMsgSize       int           `env:"GRPC_SERVER_MAX_SEND_MSG_SIZE" validate:"gte=0"`
	EnableReflection     bool          `env:"GRPC_SERVER_ENABLE_REFLECTION" validate:"-"`
	TLSCertFile          string        `env:"GRPC_SERVER_TLS_CERT_FILE" validate:"omitempty,file"`
	TLSKeyFile           string        `env:"GRPC_SERVER_TLS_KEY_FILE" validate:"omitempty,file"`
	ReadTimeout          time.Duration `env:"GRPC_SERVER_READ_TIMEOUT" validate:"gte=0"`
	WriteTimeout         time.Duration `env:"GRPC_SERVER_WRITE_TIMEOUT" validate:"gte=0"`
	EnablePrometheus     bool          `env:"GRPC_SERVER_ENABLE_PROMETHEUS" validate:"-"`
	PrometheusListenAddr string        `env:"GRPC_SERVER_PROMETHEUS_LISTEN_ADDR" validate:"required_with=EnablePrometheus,omitempty"`
}

type GRPCClientConfig struct {
	ConnectTimeout    time.Duration `env:"GRPC_CLIENT_CONNECT_TIMEOUT" validate:"gte=0"`
	MaxBackoffDelay   time.Duration `env:"GRPC_CLIENT_MAX_BACKOFF_DELAY" validate:"gte=0"`
	BaseBackoffDelay  time.Duration `env:"GRPC_CLIENT_BASE_BACKOFF_DELAY" validate:"gte=0"`
	BackoffMultiplier float64       `env:"GRPC_CLIENT_BACKOFF_MULTIPLIER" validate:"gte=1"`
	BackoffJitter     float64       `env:"GRPC_CLIENT_BACKOFF_JITTER" validate:"gte=0"`
}

type InfrastructureConfig struct {
	RedisConfig RedisConfig `validate:"required"`
}

type GRPCApiConfig struct {
	SpotInstrumentServiceHost string `env:"GRPC_API_SPOT_INSTRUMENT_SERVICE_HOST" validate:"required"`
}

type RedisConfig struct {
	Host         string        `env:"REDIS_HOST" validate:"required,hostname|ip"`
	Port         string        `env:"REDIS_PORT" validate:"required,numeric"`
	Password     string        `env:"REDIS_PASSWORD" validate:"-"`
	DB           int           `env:"REDIS_DB" validate:"gte=0"`
	MinIdleConns int           `env:"REDIS_MIN_IDLE_CONNS" validate:"gte=0"`
	PoolSize     int           `env:"REDIS_POOL_SIZE" validate:"gte=0"`
	PoolTimeout  time.Duration `env:"REDIS_POOL_TIMEOUT" validate:"gte=0"`
}

func New() (*Config, error) {
	cfg := &Config{}
	if err := cleanenv.ReadEnv(cfg); err != nil {
		return nil, err
	}

	v := validator.New()
	if err := v.Struct(cfg); err != nil {
		panic(err)
	}
	return cfg, nil
}
