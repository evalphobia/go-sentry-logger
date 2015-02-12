// send data to sentry

package log

import (
	"github.com/getsentry/raven-go"
	"net/http"
)

const (
	sentryConfigSectionName = "sentry"
)

var (
	sentryDSN string
)

// send data to sentry
func SendSentry(args map[string]string, label string, req *http.Request, e error) {
	if sentryDSN == "" {
		sentryDSN = config.GetConfigValue(sentryConfigSectionName, "dsn")
		if sentryDSN == "" {
			ErrorWithoutTrack(NewLogData("cannot find sentry dsn", 2))
			return
		}
	}
	client, err := raven.NewClient(sentryDSN, args)
	if err != nil {
		ErrorWithoutTrack(NewLogData(err, 2))
	}

	interfaces := []raven.Interface{}
	if req != nil {
		interfaces = append(interfaces, raven.NewException(e, trace()), raven.NewHttp(req))
	} else {
		interfaces = append(interfaces, raven.NewException(e, trace()))
	}
	packet := &raven.Packet{
		Message:    label,
		Interfaces: interfaces,
	}
	_, ch := client.Capture(packet, nil)
	if err = <-ch; err != nil {
		ErrorWithoutTrack(NewLogData(err, 2))
	}
}

func SendSentryMapData(args map[string]string, req *http.Request, e error) {
	SendSentry(args, args["label"], req, e)
}

// get stack trace
func trace() *raven.Stacktrace {
	return raven.NewStacktrace(0, 2, nil)
}
