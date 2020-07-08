package domain

import (
	"fmt"
	"github.com/tendermint/tendermint/abci/types"
)

func queryTX(app *Application, req types.RequestQuery) types.ResponseQuery {
	switch req.Path {
	case "hash":
		return types.ResponseQuery{Value: []byte(fmt.Sprintf("%v", app.hashCount))}
	case "tx":
		return types.ResponseQuery{Value: []byte(fmt.Sprintf("%v", app.txCount))}
	default:
		return types.ResponseQuery{Log: fmt.Sprintf("Invalid query path. Expected hash or tx, got %v", req.Path)}
	}
}
