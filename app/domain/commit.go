package domain

import (
	"encoding/binary"
	"github.com/tendermint/tendermint/abci/types"
)

func commit(app *Application) types.ResponseCommit {
	app.hashCount++
	if app.txCount == 0 {
		return types.ResponseCommit{}
	}
	hash := make([]byte, 8)
	binary.BigEndian.PutUint64(hash, uint64(app.txCount))
	return types.ResponseCommit{Data: hash}
}
