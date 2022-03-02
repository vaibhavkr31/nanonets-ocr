package config

import (
	"fmt"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

var Logger *zap.Logger
var atomicLevel zap.AtomicLevel

func init() {
	InitLogger("config")
	viper.AutomaticEnv()
	viper.AddConfigPath(".")
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")

	err := viper.ReadInConfig()
	if err != nil {
		Logger.Error(fmt.Sprintf("Error While reading the config : %s", err))
		panic(err)
	}
}

func InitLogger(packageName string) *zap.Logger {
	s := fmt.Sprintf("NANONETS:OCR:%s", packageName)
	zapLevel := returnLogLevel(viper.GetString("logLevel"))
	atomicLevel = zap.NewAtomicLevel()
	encoderCfg := zap.NewProductionEncoderConfig()
	Logger = zap.New(zapcore.NewCore(
		zapcore.NewJSONEncoder(encoderCfg),
		zapcore.Lock(os.Stdout),
		atomicLevel,
	)).Named(s)
	atomicLevel.SetLevel(zapLevel)
	return Logger
}

/**
Valid Values TRACE, DEBUG, INFO, WARN, ERROR, FATAL
*/
func returnLogLevel(logLevel string) zapcore.Level {
	switch logLevel {
	case "FATAL":
		return zap.FatalLevel
	case "ERROR":
		return zap.ErrorLevel
	case "WARN":
		return zap.WarnLevel
	case "INFO":
		return zap.InfoLevel
	case "DEBUG":
		return zap.DebugLevel
	default:
		return zap.InfoLevel
	}
}
