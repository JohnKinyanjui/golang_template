package blogger

import (
	"errors"
	"log"
)

const (
	AUTH_SERVICE_LOG         = "auth-service-log"
	USER_SERVICE_LOG         = "user-service-log"
	NOTIFICATION_SERVICE_LOG = "notification-service-log"
	DEPLOYMENT_SERVICE_LOG   = "deployment-service-log"
	FEED_SERVICE_LOG         = "feed-service-log"
	WS_SERVICE_LOG           = "user-service-log"
	WORK_SERVICE_LOG         = "work-service-log"
	FILE_SERVICE_LOG         = "feed-service-log"
	CAMPAIGN_SERVICE_LOG     = "campaign-service-log"
	ORDER_SERVICE_LOG        = "order-service-log"
	PAYMENT_SERVICE_LOG      = "payment-service-log"
)

type Log struct {
	Service             string
	UnderstandbleFormat string
	Error               error
}

func Logger(service string, err error, text string) *Log {
	return &Log{
		Service:             service,
		UnderstandbleFormat: text,
		Error:               err,
	}
}

func (l *Log) Log() error {

	log.Printf("[ %s ] %s \n", l.Service, l.Error.Error())

	return errors.New(l.UnderstandbleFormat)
}
