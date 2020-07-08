package domain

import (
	"encoding/binary"
	"fmt"
	"github.com/tendermint/tendermint/abci/example/code"
	"github.com/tendermint/tendermint/abci/types"
)

func checkTX(app *Application, req types.RequestCheckTx) types.ResponseCheckTx {
	fmt.Println(fmt.Sprintf("Invalid nonce. Expected >= %v, got %v", app.txCount, string(req.Tx)))
	fmt.Println(req.Type)
	if app.serial {
		if len(req.Tx) > 8 {
			return types.ResponseCheckTx{
				Code: code.CodeTypeEncodingError,
				Log:  fmt.Sprintf("Max tx size is 8 bytes, got %d", len(req.Tx))}
		}
		tx8 := make([]byte, 8)
		copy(tx8[len(tx8)-len(req.Tx):], req.Tx)
		txValue := binary.BigEndian.Uint64(tx8)
		if txValue < uint64(app.txCount) {
			return types.ResponseCheckTx{
				Code: code.CodeTypeBadNonce,
				Log:  fmt.Sprintf("Invalid nonce. Expected >= %v, got %v", app.txCount, txValue)}
		}
	}
	return types.ResponseCheckTx{Code: code.CodeTypeOK}
}
