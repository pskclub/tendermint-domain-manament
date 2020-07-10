package domain

import (
	"fmt"
	"github.com/tendermint/tendermint/abci/types"
)

func info(app *Application, req types.RequestInfo) types.ResponseInfo {
	fmt.Println("info...")
	return types.ResponseInfo{
		Data:             fmt.Sprintf("{\"size\":%v}", app.Size),
		LastBlockHeight:  app.Height,
		LastBlockAppHash: app.Hash,
	}
}
