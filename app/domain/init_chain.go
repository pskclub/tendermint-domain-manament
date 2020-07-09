package domain

import (
	"fmt"
	"github.com/tendermint/tendermint/abci/types"
)

func initChain(app *Application, req types.RequestInitChain) types.ResponseInitChain {
	fmt.Println("initChain....")
	return types.ResponseInitChain{}
}
