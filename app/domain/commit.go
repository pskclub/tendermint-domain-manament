package domain

import (
	"fmt"
	"github.com/tendermint/tendermint/abci/types"
)

func commit(app *Application) types.ResponseCommit {
	fmt.Println("commit...")
	app.Height++
	hash := make([]byte, 8)
	app.Hash = hash
	return types.ResponseCommit{Data: app.Hash}
}
