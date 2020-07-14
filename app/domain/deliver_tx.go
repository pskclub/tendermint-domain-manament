package domain

import (
	"fmt"
	"github.com/pskclub/tendermint-domain-manament/app/domain/models"
	"github.com/pskclub/tendermint-domain-manament/app/utils"
	"github.com/tendermint/tendermint/abci/example/code"
	"github.com/tendermint/tendermint/abci/types"
	"github.com/tendermint/tendermint/libs/common"
	"time"
)

func deliverTX(app *Application, req types.RequestDeliverTx) types.ResponseDeliverTx {
	fmt.Println("deliverTX...")
	tx := &models.Tx{}
	utils.JSONParse(req.Tx, tx)
	if tx.DomainName == "kuy.com" {
		return types.ResponseDeliverTx{Code: code.CodeTypeUnknownError}
	}
	events := []types.Event{
		{
			Type: "app",
			Attributes: []common.KVPair{
				{
					Key:   []byte("operation"),
					Value: []byte(tx.Operation),
				},
				{
					Key:   []byte("domain_name"),
					Value: []byte(tx.DomainName),
				},
				{
					Key:   []byte("owner"),
					Value: []byte(tx.Owner),
				},
				{
					Key:   []byte("receiver"),
					Value: []byte(utils.GetString(tx.Receiver)),
				},
				{
					Key:   []byte("nonce"),
					Value: []byte(tx.Nonce),
				},
			},
		},
	}

	t := new(time.Time)
	tx.Timestamp = *t
	app.DB.Create(tx)
	utils.LogStruct(tx)

	app.Size++
	return types.ResponseDeliverTx{Code: code.CodeTypeOK, Events: events}
}
