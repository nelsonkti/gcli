package logger

import (
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
	"time"
)

var Sugar *zap.SugaredLogger

type Option func(*option)

type option struct {
	env      string
	path     string
	fileName string
}

func NewEncoderConfig() zapcore.EncoderConfig {
	return zapcore.EncoderConfig{
		// Keys can be anything except the empty string.
		TimeKey:        "T",
		LevelKey:       "L",
		NameKey:        "N",
		CallerKey:      "C",
		MessageKey:     "M",
		StacktraceKey:  "S",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.CapitalLevelEncoder,
		EncodeTime:     TimeEncoder,
		EncodeDuration: zapcore.StringDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}
}

func TimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format("2006-01-02 15:04:05.000"))
}

func Base(env string, path string) Option {
	return func(o *option) {
		o.env = env
		o.path = path
	}
}

func FileName(name string) Option {
	return func(o *option) {
		o.fileName = name
	}
}

func defaultLocalPath() string {
	p, _ := os.Getwd()
	return fmt.Sprintf("%s/log", p)
}

func defaultPath() string {
	return "/wwwlogs"
}

func defaultFileName() string {
	return "logic"
}

func setPath(o *option) {
	if o.env == "local" || o.path == "" {
		o.path = defaultLocalPath()
	}

	if o.fileName == "" {
		o.fileName = defaultFileName()
	}

	o.path = fmt.Sprintf("%s/%s.log", o.path, o.fileName)
}

func New(opts ...Option) {

	o := option{env: "local"}

	for _, opt := range opts {
		opt(&o)
	}

	setPath(&o)

	w := zapcore.AddSync(&lumberjack.Logger{
		Filename:   o.path,
		MaxSize:    500, // megabytes
		MaxBackups: 3,
		MaxAge:     28, // days
	})

	core := zapcore.NewCore(
		zapcore.NewConsoleEncoder(NewEncoderConfig()),
		zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout),
			w),
		zap.DebugLevel,
	)

	Sugar = zap.New(core, zap.AddCaller()).Sugar()
}
