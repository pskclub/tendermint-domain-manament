package domain

import (
	"fmt"
	"github.com/pskclub/tendermint-domain-manament/app/utils"
	"github.com/tendermint/tendermint/abci/types"
)

func queryTX(app *Application, req types.RequestQuery) types.ResponseQuery {
	fmt.Println("queryTX...")
	utils.LogStruct(req)
	switch req.Path {
	case "size":
		return types.ResponseQuery{Value: []byte(fmt.Sprintf("%v", app.Size))}
	case "height":
		return types.ResponseQuery{Value: []byte(fmt.Sprintf("%v", app.Height))}
	default:
		return types.ResponseQuery{Log: fmt.Sprintf("Invalid query path. Expected Hash or tx, got %v", req.Path)}
	}
}
