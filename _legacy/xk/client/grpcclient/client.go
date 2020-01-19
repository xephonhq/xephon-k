// Package grpcclient is grpc client of Xephon-K, it's in separated package to avoid import grpc
// for clients that don't use grpc
package grpcclient

import (
	"context"
	"net/url"

	"github.com/dyweb/gommon/errors"
	"google.golang.org/grpc"

	"github.com/libtsdb/libtsdb-go/libtsdb"
	pb "github.com/libtsdb/libtsdb-go/libtsdb/libtsdbpb"

	xkc "github.com/xephonhq/xephon-k/xk/client"
	"github.com/xephonhq/xephon-k/xk/config"
	rpc "github.com/xephonhq/xephon-k/xk/transport/grpc"
)

var _ libtsdb.WriteClient = (*Client)(nil)

// TODO: support prepare and columnar format
type Client struct {
	cfg    config.ClientConfig
	client rpc.XephonkClient
	conn   *grpc.ClientConn

	pointsInt    []pb.PointIntTagged
	pointsDouble []pb.PointDoubleTagged
	seriesInt    []pb.SeriesIntTagged
	seriesDouble []pb.SeriesDoubleTagged
}

func New(cfg config.ClientConfig) (*Client, error) {
	_, err := url.Parse(cfg.Addr)
	if err != nil {
		return nil, errors.Wrap(err, "can't parse server address")
	}
	conn, err := grpc.Dial(cfg.Addr, grpc.WithInsecure())
	if err != nil {
		return nil, errors.Wrapf(err, "grpc dial failed %s", cfg.Addr)
	}
	client := rpc.NewClient(conn)
	return &Client{
		cfg:    cfg,
		client: client,
		conn:   conn,
	}, nil
}

func (c *Client) Meta() libtsdb.Meta {
	return xkc.Meta()
}

func (c *Client) Close() error {
	if err := c.conn.Close(); err != nil {
		return errors.Wrap(err, "can't close grpc client connection")
	}
	return nil
}

func (c *Client) WriteIntPoint(p *pb.PointIntTagged) {
	// TODO: deal with prepare and columnar
	c.pointsInt = append(c.pointsInt, *p)
}

func (c *Client) WriteDoublePoint(p *pb.PointDoubleTagged) {
	// TODO: deal with prepare and columnar
	c.pointsDouble = append(c.pointsDouble, *p)
}

func (c *Client) WriteSeriesIntTagged(p *pb.SeriesIntTagged) {
	// TODO: deal with prepare and columnar
	c.seriesInt = append(c.seriesInt, *p)
}

func (c *Client) WriteSeriesDoubleTagged(p *pb.SeriesDoubleTagged) {
	// TODO: deal with prepare and columnar
	c.seriesDouble = append(c.seriesDouble, *p)
}

func (c *Client) Flush() error {
	return c.send()
}

func (c *Client) send() error {
	merr := errors.NewMultiErr()
	// NOTE: normally we assume user only use one methods, so we just use one go routine
	if len(c.pointsInt) != 0 || len(c.pointsDouble) != 0 {
		req := rpc.WritePointsReq{
			Int:    c.pointsInt,
			Double: c.pointsDouble,
		}
		_, err := c.client.WritePoints(context.Background(), &req)
		if err != nil {
			merr.Append(err)
		}
		c.pointsInt = c.pointsInt[:0]
		c.pointsDouble = c.pointsDouble[:0]
	}
	if len(c.seriesInt) != 0 || len(c.seriesDouble) != 0 {
		req := rpc.WriteSeriesReq{
			Int:    c.seriesInt,
			Double: c.seriesDouble,
		}
		_, err := c.client.WriteSeries(context.Background(), &req)
		if err != nil {
			merr.Append(err)
		}
		c.seriesInt = c.seriesInt[:0]
		c.seriesDouble = c.seriesDouble[:0]
	}
	return merr.ErrorOrNil()
}
