package domain

import (
	"github.com/tendermint/tendermint/abci/types"
)

type Application struct {
	types.BaseApplication

	hashCount int
	txCount   int
	serial    bool
}

func NewApplication(serial bool) *Application {
	return &Application{serial: serial}
}

func (app *Application) SetOption(req types.RequestSetOption) types.ResponseSetOption {
	key, value := req.Key, req.Value
	if key == "serial" && value == "on" {
		app.serial = true
	} else {
		/*
			TODO Panic and have the ABCI server pass an exception.
			The client can call SetOptionSync() and get an `error`.
			return types.ResponseSetOption{
				Error: fmt.Sprintf("Unknown key (%s) or value (%s)", key, value),
			}
		*/
		return types.ResponseSetOption{}
	}

	return types.ResponseSetOption{}
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
