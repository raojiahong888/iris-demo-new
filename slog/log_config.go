package slog

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/middleware/pprof"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"io"
	"iris-demo-new/config"
	"os"
	"path"
	"strings"
	"time"
)

var Log *zap.SugaredLogger

func InitLogger(logConfig config.LogConfig, app *iris.Application)  {
	encoder := getEncoder()
	infoLevel := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl < zapcore.ErrorLevel
	})
	infoWriter := getLogWriter(logConfig.Path, logConfig.FileName + "-info")
	errorLevel := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl >= zapcore.ErrorLevel
	})
	errorWriter := getLogWriter(logConfig.Path, logConfig.FileName + "-error")
	core := zapcore.NewTee(
		zapcore.NewCore(encoder, zapcore.AddSync(os.Stdout), zap.DebugLevel),
		zapcore.NewCore(encoder, zapcore.AddSync(infoWriter), infoLevel),
		zapcore.NewCore(encoder, zapcore.AddSync(errorWriter), errorLevel),
	)
	//develop mode
	caller := zap.AddCaller()
	//open the code line
	development := zap.Development()
	logger := zap.New(core, caller, development, zap.AddCallerSkip(1))
	Log = logger.Sugar()
	//set iris log level
	if strings.EqualFold(logConfig.Level, "INFO") {
		app.Logger().SetLevel("info")
		app.Logger().SetOutput(infoWriter)
	}
	if strings.EqualFold(logConfig.Level, "Debug") {
		app.Logger().SetLevel("debug")
		app.Logger().SetOutput(os.Stdout)
		p := pprof.New()
		app.Get("/debug/pprof", p)
		app.Get("/debug/pprof/{action:path}", p)
	}
}

func getEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = customTimeEncoder
	encoderConfig.EncodeLevel = zapcore.LowercaseLevelEncoder
	encoderConfig.LineEnding = zapcore.DefaultLineEnding
	return zapcore.NewConsoleEncoder(encoderConfig)
}

func customTimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format("["+config.Setting.Iris.TimeFormat+"]"))
}

func getLogWriter(logPath, level string) io.Writer {
	logFullPath := path.Join(logPath, level)
	hook, err := rotatelogs.New(logFullPath+"-%Y%m%d%H"+".txt", rotatelogs.WithRotationTime(24*time.Hour))
	if err != nil {
		panic(err)
	}
	return hook
}

func Info(args ...interface{})  {
	Log.Info(args...)
}

func Infof(template string, args ...interface{})  {
	Log.Infof(template, args...)
}

func Debug(args ...interface{})  {
	Log.Debug(args...)
}

func Debugf(template string, args ...interface{})  {
	Log.Debugf(template, args...)
}

func Warn(args ...interface{})  {
	Log.Warn(args...)
}

func Warnf(template string, args ...interface{})  {
	Log.Warnf(template, args...)
}

func Error(args ...interface{})  {
	Log.Error(args...)
}

func Errorf(template string, args ...interface{})  {
	Log.Errorf(template, args...)
}
