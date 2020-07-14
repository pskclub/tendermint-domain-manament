package domain

import (
	"github.com/pskclub/tendermint-domain-manament/app/utils"
	"github.com/tendermint/tendermint/abci/types"
)

func endBlock(app *Application, req types.RequestEndBlock) types.ResponseEndBlock {
	/*validator := []types.ValidatorUpdate{
		{
			PubKey: types.PubKey{
				Type: "tendermint/PubKeyEd25519",
				Data: []byte("8/GplILf/qdHowmC7AuPeXKECB+HnBdkpriMaL6yaQ0="),
			},
			Power: 2,
		},
	}*/
	utils.LogStruct(app.ValUpdates)
	return types.ResponseEndBlock{ValidatorUpdates: app.ValUpdates}
}
