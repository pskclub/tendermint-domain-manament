package domain

import (
	"encoding/json"
	"fmt"
	"github.com/tendermint/tendermint/abci/example/code"
	"github.com/tendermint/tendermint/abci/types"
)

type Payload struct {
	Name interface{}
}

func checkTX(app *Application, req types.RequestCheckTx) types.ResponseCheckTx {
	data := &Payload{}
	json.Unmarshal(req.Tx, data)
	fmt.Println(fmt.Sprintf("Invalid nonce. Expected >= %v", data.Name))
	return types.ResponseCheckTx{Code: code.CodeTypeOK}
}
