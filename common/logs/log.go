package logs

import (
	"github.com/JingruiLea/ad_boost/common/envs"
	"github.com/JingruiLea/ad_boost/config"
	"github.com/JingruiLea/ad_boost/utils"
	"io"
	"os"
	"time"

	"golang.org/x/net/context"

	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/sirupsen/logrus"
)

var log = logrus.New()

func init() {
	logPath := config.GetLogPath()
	writer, _ := rotatelogs.New(
		logPath+".%Y%m%d%H%M",
		rotatelogs.WithClock(rotatelogs.Local),
		rotatelogs.WithLinkName(logPath),
		rotatelogs.WithMaxAge(time.Duration(180)*time.Hour*24),
		rotatelogs.WithRotationTime(time.Duration(60)*time.Minute),
	)
	if envs.IsDev() {
		mw := io.MultiWriter(os.Stdout, writer)
		log.SetOutput(mw)
	} else {
		log.SetOutput(writer)
	}
	log.SetLevel(logrus.DebugLevel)
	log.AddHook(&truncateHook{})
}

const DefaultXRequestIDKey = "x-request-id"

func entry(ctx context.Context) *logrus.Entry {
	entry := logrus.NewEntry(log)
	entry = entry.WithContext(ctx)
	reqID, _ := utils.GetReqID(ctx)
	entry = entry.WithField(DefaultXRequestIDKey, reqID)
	return entry
}

const maxLogLength = 1000

func truncateToLongStr(str string) string {
	if len(str) > maxLogLength {
		return str[:maxLogLength]
	}
	return str
}

type truncateHook struct {
}

func (t truncateHook) Levels() []logrus.Level {
	return logrus.AllLevels
}

func (t truncateHook) Fire(e *logrus.Entry) error {
	e.Message = truncateToLongStr(e.Message)
	return nil
}

func CtxInfof(ctx context.Context, format string, args ...interface{}) {
	entry(ctx).Infof(format, args...)
}

func CtxDebugf(ctx context.Context, format string, args ...interface{}) {
	entry(ctx).Debugf(format, args...)
}

func CtxWarnf(ctx context.Context, format string, args ...interface{}) {
	entry(ctx).Warnf(format, args...)
}

func CtxErrorf(ctx context.Context, format string, args ...interface{}) {
	entry(ctx).Errorf(format, args...)
}

func Fatal(args ...interface{}) {
	log.Fatal(args...)
}

func Infof(format string, args ...interface{}) {
	log.Infof(format, args...)
}

func Debugf(format string, args ...interface{}) {
	log.Debugf(format, args...)
}

func Tracef(format string, args ...interface{}) {
	log.Tracef(format, args...)
}

func CtxTracef(ctx context.Context, format string, args ...interface{}) {
	entry(ctx).Tracef(format, args...)
}

func Logger() *logrus.Logger {
	return log
}
