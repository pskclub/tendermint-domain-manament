package domain

import (
	"fmt"
	"github.com/tendermint/tendermint/abci/types"
)

func endBlock(app *Application, req types.RequestEndBlock) types.ResponseEndBlock {
	fmt.Println("end block...")
	return types.ResponseEndBlock{}
}
