package domain

import (
	"fmt"
	"github.com/tendermint/tendermint/abci/example/code"
	"github.com/tendermint/tendermint/abci/types"
)

func deliverTX(app *Application, req types.RequestDeliverTx) types.ResponseDeliverTx {
	fmt.Println("deliverTX...")
	return types.ResponseDeliverTx{Code: code.CodeTypeOK}
}
