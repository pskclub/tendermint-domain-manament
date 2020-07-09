package domain

import (
	"fmt"
	"github.com/tendermint/tendermint/abci/types"
)

func info(app *Application, req types.RequestInfo) types.ResponseInfo {
	fmt.Println("info...")
	return types.ResponseInfo{Data: fmt.Sprintf("{\"hashes\":%v,\"txs\":%v}", app.hashCount, 0)}
}
