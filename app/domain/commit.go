package domain

import (
	"fmt"
	"github.com/tendermint/tendermint/abci/types"
)

func commit(app *Application) types.ResponseCommit {
	fmt.Println("commit...")
	app.hashCount++
	hash := make([]byte, 8)
	return types.ResponseCommit{Data: hash}
}
