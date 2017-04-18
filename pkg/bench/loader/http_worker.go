package loader

import (
	"context"
	"github.com/xephonhq/xephon-k/pkg/bench"
	"github.com/xephonhq/xephon-k/pkg/bench/generator"
	"github.com/xephonhq/xephon-k/pkg/bench/serialize"
	"github.com/xephonhq/xephon-k/pkg/common"
	"io"
	"io/ioutil"
	"net/http"
	"time"
)

type HTTPWorker struct {
	config      Config
	ctx         context.Context
	tr          *http.Transport
	baseRequest *http.Request
	metricChan  chan *bench.RequestMetric
}

func NewHTTPWorker(config Config, ctx context.Context, baseReq *http.Request, tr *http.Transport, c chan *bench.RequestMetric) *HTTPWorker {
	return &HTTPWorker{
		config:      config,
		ctx:         ctx,
		baseRequest: baseReq,
		tr:          tr,
		metricChan:  c,
	}
}

func (worker *HTTPWorker) work() {
	client := http.Client{Transport: worker.tr, Timeout: worker.config.Timeout}
	gen := generator.NewConstantGenerator()
	var serializer serialize.Serializer
	switch worker.config.TargetDB {
	case bench.DBInfluxDB:
		serializer = &serialize.InfluxDBSerialize{}
	case bench.DBXephonK:
		serializer = &serialize.XephonKSerialize{}
	case bench.DBKairosDB:
		serializer = &serialize.KairosDBSerialize{}
	default:
		log.Panic("unsupported database, not serailizer avaliable")
		return
	}
	tags := make(map[string]string)
	tags["agent"] = "xephon-k-loader"

	for {
		select {
		case <-worker.ctx.Done():
			log.Info("worker finished")
			return
		default:
			// generate the series based on batch
			series := common.IntSeries{Name: "xephon", Tags: tags}
			// TODO: pre allocate the batch size slice should be more efficient
			for i := 0; i < worker.config.BatchSize; i++ {
				series.Points = append(series.Points, gen.NextIntPoint())
			}

			serializer.Start()
			serializer.WriteInt(series)
			serializer.End()

			result := &bench.RequestMetric{
				Start: time.Now(),
				// FIXME: len return int, does that mean golang can't have array larger than int32?
				RequestSize:  int64(serializer.DataLen()),
				ResponseSize: 0,
			}
			req := new(http.Request)
			// copy base request
			*req = *worker.baseRequest
			req.Body = serializer.ReadCloser()
			// do the request
			res, err := client.Do(req)
			// reset the serializer
			serializer.Reset()
			if err != nil {
				log.Warn(err)
			} else {
				io.Copy(ioutil.Discard, res.Body)
				res.Body.Close()
			}
			if res != nil {
				result.Code = res.StatusCode
				result.ResponseSize = res.ContentLength
			}
			result.Err = err
			result.End = time.Now()
			worker.metricChan <- result
		}
	}
}
