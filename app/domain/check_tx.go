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

	find := &models.Tx{}
	result = app.DB.Where("domain_name = ?", data.DomainName).Last(find)

	if data.Operation == OperationReserve {
		if gorm.IsRecordNotFoundError(result.Error) {
			return types.ResponseCheckTx{Code: code.CodeTypeOK, Log: utils.StructToString(data)}
		}

		if result.Error != nil {
			return types.ResponseCheckTx{Code: code.CodeTypeUnknownError, Log: result.Error.Error()}
		}

		if find.Operation == OperationRelease {
			return types.ResponseCheckTx{Code: code.CodeTypeOK, Log: utils.StructToString(data)}
		}

	}

	if data.Operation == OperationTransfer {
		if utils.IsEmpty(data.Receiver) {
			return types.ResponseCheckTx{Code: code.CodeTypeEncodingError, Log: "receiver is required"}
		}

		if gorm.IsRecordNotFoundError(result.Error) {
			return types.ResponseCheckTx{Code: code.CodeTypeUnauthorized, Log: result.Error.Error()}
		}

		if result.Error != nil {
			return types.ResponseCheckTx{Code: code.CodeTypeUnknownError, Log: result.Error.Error()}
		}

		if data.Owner == find.Owner && data.Owner != utils.GetString(find.Receiver) && (find.Operation == OperationTransfer || find.Operation == OperationReserve) {
			return types.ResponseCheckTx{Code: code.CodeTypeOK, Log: utils.StructToString(data)}
		}

	}

	if data.Operation == OperationRelease {
		if gorm.IsRecordNotFoundError(result.Error) {
			return types.ResponseCheckTx{Code: code.CodeTypeUnauthorized, Log: result.Error.Error()}
		}

		if result.Error != nil {
			return types.ResponseCheckTx{Code: code.CodeTypeUnknownError, Log: result.Error.Error()}
		}

		if find.Operation == OperationTransfer && data.Owner == utils.GetString(find.Receiver) {
			return types.ResponseCheckTx{Code: code.CodeTypeOK, Log: utils.StructToString(data)}
		}

		if find.Operation == OperationReserve && data.Owner == find.Owner {
			return types.ResponseCheckTx{Code: code.CodeTypeOK, Log: utils.StructToString(data)}
		}
	}

	return types.ResponseCheckTx{Code: code.CodeTypeUnauthorized, Log: utils.StructToString(data)}
}
