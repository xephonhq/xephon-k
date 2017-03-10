package middleware

import (
	"time"

	dlog "github.com/dyweb/gommon/log"
	"github.com/xephonhq/xephon-k/pkg/common"
	"github.com/xephonhq/xephon-k/pkg/server/service"
	"github.com/xephonhq/xephon-k/pkg/util"
)

var logger = util.Logger.NewEntryWithPkg("k.s.m")

type LoggingInfoServiceMiddleware struct {
	service.InfoService
	logger *dlog.Entry
}

func NewLoggingInfoServiceMiddleware(service service.InfoService) LoggingInfoServiceMiddleware {
	return LoggingInfoServiceMiddleware{InfoService: service, logger: logger}
}

// FIXME: the naming here is misleading, the info actually return all the info, more than just version
// and how to hand things like info/version in go-kit
func (mw LoggingInfoServiceMiddleware) Version() string {
	defer func(begin time.Time) {
		// TODO: human readable time format, what's the number, ms, ns?
		mw.logger.Infof("GET /info %d", time.Since(begin))
	}(time.Now())
	return mw.InfoService.Version()
}

type LoggingWriteServiceMiddleware struct {
	service.WriteService
	logger *dlog.Entry
}

func NewLoggingWriteServiceMiddleware(service service.WriteService) LoggingWriteServiceMiddleware {
	return LoggingWriteServiceMiddleware{WriteService: service, logger: logger}
}

func (mw LoggingWriteServiceMiddleware) WriteInt(series []common.IntSeries) error {
	defer func(begin time.Time) {
		// TODO: human readable time format, what's the number, ms, ns?
		mw.logger.Infof("POST /write %d", time.Since(begin))
	}(time.Now())
	return mw.WriteService.WriteInt(series)
}

type LoggingReadServiceMiddleware struct {
	service.ReadService
	logger *dlog.Entry
}

func NewLoggingReadServiceMiddleware(service service.ReadService) LoggingReadServiceMiddleware {
	return LoggingReadServiceMiddleware{ReadService: service, logger: logger}
}
