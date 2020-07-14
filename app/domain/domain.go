package domain

import (
	"bytes"
	"fmt"
	"github.com/imroc/req"
	"github.com/jinzhu/gorm"
	"github.com/pskclub/tendermint-domain-manament/app/utils"
	"github.com/tendermint/tendermint/abci/example/code"
	"github.com/tendermint/tendermint/abci/types"
	"github.com/tendermint/tendermint/crypto/ed25519"
)

type Application struct {
	types.BaseApplication

	ValUpdates         []types.ValidatorUpdate
	valAddrToPubKeyMap map[string]types.PubKey
	Size               int64  `json:"size"`
	Hash               []byte `json:"hash"`
	Height             int64  `json:"height"`
	DB                 *gorm.DB
	Requester          *req.Req
}

func NewApplication(db *gorm.DB) *Application {
	return &Application{DB: db, Requester: req.New()}
}

func (app *Application) SetOption(req types.RequestSetOption) types.ResponseSetOption {
	return types.ResponseSetOption{}
}

func (app *Application) InitChain(req types.RequestInitChain) types.ResponseInitChain {
	return initChain(app, req)
}

func (app *Application) Info(req types.RequestInfo) types.ResponseInfo {
	return info(app, req)
}

func (app *Application) BeginBlock(req types.RequestBeginBlock) types.ResponseBeginBlock {
	app.ValUpdates = make([]types.ValidatorUpdate, 0)

	utils.LogStruct(req.ByzantineValidators)
	for _, ev := range req.ByzantineValidators {
		if ev.Type == "duplicate/vote" {
			// decrease voting power by 1
			if ev.TotalVotingPower == 0 {
				continue
			}
			app.updateValidator(types.ValidatorUpdate{
				PubKey: app.valAddrToPubKeyMap[string(ev.Validator.Address)],
				Power:  ev.TotalVotingPower - 1,
			})
		}
	}
	return types.ResponseBeginBlock{}
}

func (app *Application) DeliverTx(req types.RequestDeliverTx) types.ResponseDeliverTx {
	return deliverTX(app, req)
}

func (app *Application) CheckTx(req types.RequestCheckTx) types.ResponseCheckTx {
	return checkTX(app, req)
}

func (app *Application) Commit() (resp types.ResponseCommit) {
	return commit(app)
}

func (app *Application) Query(req types.RequestQuery) types.ResponseQuery {
	return queryTX(app, req)
}

func (app *Application) EndBlock(req types.RequestEndBlock) types.ResponseEndBlock {
	return endBlock(app, req)
}

// add, update, or remove a validator
func (app *Application) updateValidator(v types.ValidatorUpdate) types.ResponseDeliverTx {
	pubkey := ed25519.PubKeyEd25519{}
	copy(pubkey[:], v.PubKey.Data)

	value := bytes.NewBuffer(make([]byte, 0))
	if err := types.WriteMessage(&v, value); err != nil {
		return types.ResponseDeliverTx{
			Code: code.CodeTypeEncodingError,
			Log:  fmt.Sprintf("Error encoding validator: %v", err)}
	}
	app.valAddrToPubKeyMap[string(pubkey.Address())] = v.PubKey

	// we only update the changes array if we successfully updated the tree
	app.ValUpdates = append(app.ValUpdates, v)

	return types.ResponseDeliverTx{Code: code.CodeTypeOK}
}
