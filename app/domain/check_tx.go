package domain

import (
	"fmt"
	"github.com/pskclub/tendermint-domain-manament/app/domain/models"
	"github.com/pskclub/tendermint-domain-manament/app/utils"
	"github.com/tendermint/tendermint/abci/example/code"
	"github.com/tendermint/tendermint/abci/types"
)

func checkTX(app *Application, req types.RequestCheckTx) types.ResponseCheckTx {
	fmt.Println("checkTX...")
	data := &models.Tx{}
	err := utils.JSONParse(req.Tx, data)
	if err != nil {
		fmt.Println(err.Error())
		return types.ResponseCheckTx{Code: code.CodeTypeEncodingError, Log: err.Error()}
	}

	if utils.IsEmpty(data.By) {
		return types.ResponseCheckTx{Code: code.CodeTypeEncodingError, Log: "by is required"}
	}

	if utils.IsEmpty(data.DomainName) {
		return types.ResponseCheckTx{Code: code.CodeTypeEncodingError, Log: "domain_name is required"}
	}

	if utils.IsEmpty(data.Nonce) {
		return types.ResponseCheckTx{Code: code.CodeTypeEncodingError, Log: "nonce is required"}
	}

	if utils.IsEmpty(data.Operation) {
		return types.ResponseCheckTx{Code: code.CodeTypeEncodingError, Log: "Operation is required"}
	}

	return types.ResponseCheckTx{Code: code.CodeTypeOK, Log: fmt.Sprintf("%v", data)}
}
