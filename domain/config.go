package domain

import (
	"log"
	"os"
	"strings"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Config struct {
	LogLevel   string
	Log        *zap.SugaredLogger
	ConfigFile map[string]string
	Mongo      *MongoConfig
	GrpcPort   string
}

type MongoConfig struct {
	Hosts          []string // coma separated list of host:port
	DBName         string   `envconfig:"MONGO_DB_NAME"`
	ReadPref       readpref.Mode
	ReplicaSetName string `envconfig:"MONGO_REPLICA_SET_NAME"`
	Options        string
}

// GetReadPref gets from environment ReadPref value for mongo, if not set get from config file.
func GetReadPref(configFile map[string]string, configKey string, log *zap.SugaredLogger) readpref.Mode {
	var readPref readpref.Mode

	var err error

	val, ok := os.LookupEnv(configKey)
	if ok {
		readPref, err = readpref.ModeFromString(val)
	} else {
		readPref, err = readpref.ModeFromString(configFile[configKey])
	}

	if err != nil {
		log.Info("invalid or not set", configKey)
	}

	return readPref
}

// GetSliceOfStrings gets from environment variable values separated by comma, if not set get from config file.
func GetSliceOfStrings(configFile map[string]string, configKey string, log *zap.SugaredLogger) []string {
	var result []string

	value, ok := os.LookupEnv(configKey)

	switch {
	case ok:
		result = strings.Split(value, ",")
	case configFile != nil:
		result = strings.Split(configFile[configKey], ",")
	case log != nil:
		log.Warn("configFile invalid or not set")
	}

	if result == nil {
		if configFile != nil {
			if log != nil {
				log.Warn("invalid or not set: ", configKey)
			}
		} else if log != nil {
			log.Warn("configFile invalid or not set")
		}
	}

	return result
}

// GetStringValue get string value from environment variable, if not set get from config file.
func GetStringValue(configFile map[string]string, configKey string, log *zap.SugaredLogger) string {
	var result string

	value, ok := os.LookupEnv(configKey)

	switch {
	case ok:
		result = value
	case configFile != nil:
		result = configFile[configKey]
	case log != nil:
		log.Warn("invalid or not set")
	}

	if result == "" {
		if configFile != nil {
			if log != nil {
				log.Info("invalid or not set", configKey)
			}
		} else if log != nil {
			log.Warn("configFile invalid or not set")
		}
	}

	return result
}

// GetLogLevel from environment variable, if not set get from config file.
func GetLogLevel(configFile map[string]string, configKey string) string {
	var result string

	defaultLogLevel := "info"

	value, ok := os.LookupEnv(configKey)
	switch {
	case ok:
		result = value
	case configFile != nil:
		result = configFile[configKey]
	default:
		log.Println("configFile invalid or not set")
	}

	if result == "" {
		log.Println("invalid or not set", configKey)

		result = defaultLogLevel
	}

	return result
}

func NewConfig(cfgFile string) *Config {
	if cfgFile == "" {
		cfgFile = os.Getenv("CONFIG_FILE")
		if cfgFile == "" {
			log.Println("CONFIG_FILE not set,load .env.dev")
			cfgFile = "./.env.dev"
		}
	}

	log.Printf("CONFIG_FILE load %s\n", cfgFile)

	env, err := godotenv.Read(cfgFile)
	if err != nil {
		log.Panicf("failed to read %s, %s", cfgFile, err.Error())
	}

	logLevelString := GetLogLevel(env, "LOG_LEVEL")

	logLevel, err := zapcore.ParseLevel(logLevelString)
	if err != nil {
		log.Printf("invalid logLevel : %s", logLevelString)
	}

	cfgLog := zap.Config{
		Encoding:         "console",
		Level:            zap.NewAtomicLevelAt(logLevel),
		OutputPaths:      []string{"stderr"},
		ErrorOutputPaths: []string{"stderr"},
		EncoderConfig:    zap.NewDevelopmentEncoderConfig(),
	}

	logger, err := cfgLog.Build()
	if err != nil {
		panic(err)
	}

	sugaredLogger := logger.Sugar()
	sugaredLogger = sugaredLogger.Named("config")

	cfg := &Config{
		Log:        sugaredLogger,
		LogLevel:   logLevelString,
		ConfigFile: env,
		GrpcPort:   GetStringValue(env, "GRPC_PORT", sugaredLogger),
		Mongo: &MongoConfig{
			Hosts:          GetSliceOfStrings(env, "MONGO_HOSTS", sugaredLogger),
			DBName:         GetStringValue(env, "MONGO_DB_NAME", sugaredLogger),
			ReadPref:       GetReadPref(env, "MONGO_READ_PREF", sugaredLogger),
			ReplicaSetName: GetStringValue(env, "MONGO_REPLICA_SET_NAME", sugaredLogger),
		},
	}

	return cfg
}
