package logs

import (
	"context"
	"github.com/JingruiLea/ad_boost/lark"
	"github.com/sirupsen/logrus"
)

// AlertHook is a logrus hook which sends alerts on error logs.
type AlertHook struct {
	// Add any required fields, like channels, auth keys, etc.
}

// Levels implement Hook interface to return the levels that this hook is interested in.
func (hook *AlertHook) Levels() []logrus.Level {
	return []logrus.Level{logrus.ErrorLevel}
}

// Fire implement Hook interface to send an alert when an error log occurs.
func (hook *AlertHook) Fire(entry *logrus.Entry) error {
	// Implement your alerting logic here. For example:
	alertMessage := entry.Message
	// Send the alertMessage to your alerting system (e.g., email, Slack, pager, etc.)
	SendAlert(entry.Context, alertMessage)
	return nil
}

// SendAlert is a dummy function representing an alert sending operation.
func SendAlert(ctx context.Context, message string) {
	lark.SendAlertMessage(ctx, message)
}
