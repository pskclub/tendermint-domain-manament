package domain

import (
	"fmt"
	"github.com/pskclub/tendermint-domain-manament/app/domain/models"
	"github.com/tendermint/tendermint/abci/types"
)

func initChain(app *Application, req types.RequestInitChain) types.ResponseInitChain {
	fmt.Println("initChain....")
	app.DB.AutoMigrate(&models.Tx{})
	app.DB.Delete(&models.Tx{})
	return types.ResponseInitChain{}
}
