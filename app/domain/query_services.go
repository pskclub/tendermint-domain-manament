package domain

import (
	"fmt"
	"github.com/pskclub/tendermint-domain-manament/app/domain/models"
	"github.com/pskclub/tendermint-domain-manament/app/utils"
	"github.com/tendermint/tendermint/abci/types"
)

func getDomainHistory(app *Application, req types.RequestQuery) types.ResponseQuery {
	txs := make([]models.Tx, 0)
	app.DB.Where("domain_name=?", req.Data).Find(&txs)
	return types.ResponseQuery{
		Value: []byte(utils.StructToString(txs)),
		Info:  fmt.Sprintf("%v", len(txs)),
		Log:   utils.StructToString(txs),
	}
}

func getUserHistory(app *Application, req types.RequestQuery) types.ResponseQuery {
	txs := make([]models.Tx, 0)
	app.DB.Where("owner=? OR receiver=?", req.Data, req.Data).Find(&txs)
	return types.ResponseQuery{
		Value: []byte(utils.StructToString(txs)),
		Info:  fmt.Sprintf("%v", len(txs)),
		Log:   utils.StructToString(txs),
	}
}
