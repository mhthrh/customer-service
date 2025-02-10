package control

import (
	"context"
	"fmt"
	cMerror "github.com/mhthrh/GoNest/model/error"
	l "github.com/mhthrh/GoNest/model/loader"
	mPool "github.com/mhthrh/GoNest/model/pool"
	cPool "github.com/mhthrh/GoNest/pkg/pool/postgres"
	"time"
)

var (
	req        chan mPool.Request
	res        chan mPool.Response
	reqRefresh chan struct{}
	resRefresh chan *cMerror.XError
	reqManage  chan mPool.ManageRequest
	resManage  chan *mPool.Connection
)

func init() {
	req = make(chan mPool.Request)
	res = make(chan mPool.Response)
	reqRefresh = make(chan struct{})
	resRefresh = make(chan *cMerror.XError)
	reqManage = make(chan mPool.ManageRequest)
	resManage = make(chan *mPool.Connection)

}
func Run(ctx context.Context, config l.Config, e chan<- *cMerror.XError) {
	var r mPool.Response
	pool, err := cPool.New(config.DataBase)
	if err != nil {
		e <- err
	}
	go pool.Maker(req, res)
	req <- mPool.Request{
		Count: uint(config.DataBase.PoolSize),
		Type:  mPool.Types(1),
		Stop:  false,
	}
	select {
	case r = <-res:
		if r.Error != nil {
			e <- r.Error
		}
		if r.Total != uint(config.DataBase.PoolSize) {
			e <- mPool.SizeUnexpected(nil)
		}
	case <-time.After(time.Second * 10):
		e <- mPool.TimeOut(nil)
	case <-ctx.Done():
		e <- mPool.TimeOut(nil)
	}

	go pool.Refresh(reqRefresh, resRefresh)
	go pool.Manager(reqManage, resManage)
	for {
		select {
		case <-ctx.Done():
			e <- mPool.TerminateByMain(nil)
		case <-time.After(time.Second * time.Duration(config.DataBase.RefreshTime)):
			reqRefresh <- struct{}{}
		case f := <-resRefresh:
			fmt.Println(f) // should be change
		}
	}
}
