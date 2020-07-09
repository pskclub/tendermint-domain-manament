package domain

import (
	"fmt"
	"github.com/pskclub/tendermint-domain-manament/app/domain/models"
	"github.com/pskclub/tendermint-domain-manament/app/utils"
	"github.com/tendermint/tendermint/abci/example/code"
	"github.com/tendermint/tendermint/abci/types"
	"github.com/tendermint/tendermint/libs/common"
	"strconv"
)

func deliverTX(app *Application, req types.RequestDeliverTx) types.ResponseDeliverTx {
	fmt.Println("deliverTX...")
	tx := &models.Tx{}
	utils.JSONParse(req.Tx, tx)
	events := make([]types.Event, 0)
	events = append(events, types.Event{
		Type: "app",
		Attributes: []common.KVPair{
			{
				Key:   []byte("operation"),
				Value: []byte(strconv.Itoa(tx.Operation)),
			},
			{
				Key:   []byte("domain_name"),
				Value: []byte(tx.DomainName),
			},
			{
				Key:   []byte("by"),
				Value: []byte(tx.By),
			},
			{
				Key:   []byte("nonce"),
				Value: []byte(tx.Nonce),
			},
		},
	})
	return types.ResponseDeliverTx{Code: code.CodeTypeOK, Events: events}
}
