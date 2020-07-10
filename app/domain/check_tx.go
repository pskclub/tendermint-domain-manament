package domain

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/pskclub/tendermint-domain-manament/app/domain/models"
	"github.com/pskclub/tendermint-domain-manament/app/utils"
	"github.com/tendermint/tendermint/abci/example/code"
	"github.com/tendermint/tendermint/abci/types"
)

func validateInput() {

}

func checkTX(app *Application, req types.RequestCheckTx) types.ResponseCheckTx {
	fmt.Println("checkTX...")
	data := &models.Tx{}
	err := utils.JSONParse(req.Tx, data)
	if err != nil {
		return types.ResponseCheckTx{Code: code.CodeTypeEncodingError, Log: err.Error()}
	}

	if utils.IsEmpty(data.Owner) {
		return types.ResponseCheckTx{Code: code.CodeTypeEncodingError, Log: "owner is required"}
	}

	if utils.IsEmpty(data.DomainName) {
		return types.ResponseCheckTx{Code: code.CodeTypeEncodingError, Log: "domain_name is required"}
	}

	if utils.IsEmpty(data.Nonce) {
		return types.ResponseCheckTx{Code: code.CodeTypeEncodingError, Log: "nonce is required"}
	}

	if utils.IsEmpty(data.Operation) {
		return types.ResponseCheckTx{Code: code.CodeTypeEncodingError, Log: "operation is required"}
	} else {
		if _, err := utils.IsStrIn(&data.Operation, "reserve|release|transfer", "operation"); err != nil {
			return types.ResponseCheckTx{Code: code.CodeTypeEncodingError, Log: err.Error()}
		}
	}

	checkNonce := &models.Tx{}
	result := app.DB.Where("nonce = ?", data.Nonce).Find(checkNonce)
	if result.Error == nil {
		return types.ResponseCheckTx{Code: code.CodeTypeBadNonce, Log: "nonce already exits"}
	} else {
		if !gorm.IsRecordNotFoundError(result.Error) {
			return types.ResponseCheckTx{Code: code.CodeTypeUnknownError, Log: result.Error.Error()}
		}
	}

	return types.ResponseCheckTx{Code: code.CodeTypeOK, Log: utils.StructToString(data)}
}
