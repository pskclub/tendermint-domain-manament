package domain

import (
	"github.com/jinzhu/gorm"
	"github.com/tendermint/tendermint/abci/types"
)

type Application struct {
	types.BaseApplication

	Size   int64  `json:"size"`
	Hash   []byte `json:"hash"`
	Height int64  `json:"height"`
	DB     *gorm.DB
}

func NewApplication(db *gorm.DB) *Application {
	return &Application{DB: db}
}

func (app *Application) SetOption(req types.RequestSetOption) types.ResponseSetOption {
	return types.ResponseSetOption{}
}

func (app *Application) InitChain(req types.RequestInitChain) types.ResponseInitChain {
	return initChain(app, req)
}

func (app *Application) Info(req types.RequestInfo) types.ResponseInfo {
	return info(app, req)
}

func (app *Application) DeliverTx(req types.RequestDeliverTx) types.ResponseDeliverTx {
	return deliverTX(app, req)
}

func (app *Application) CheckTx(req types.RequestCheckTx) types.ResponseCheckTx {
	return checkTX(app, req)
}

func (app *Application) Commit() (resp types.ResponseCommit) {
	return commit(app)
}

func (app *Application) Query(req types.RequestQuery) types.ResponseQuery {
	return queryTX(app, req)
}
