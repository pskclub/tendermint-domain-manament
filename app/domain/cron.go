package domain

import (
	"errors"
	"fmt"
	"github.com/pskclub/tendermint-domain-manament/app/domain/models"
)

func saveState(app *Application) {
	txs := make([]models.Tx, 0)
	app.DB.Where("timestamp=''").Find(&txs)
	for _, item := range txs {
		txSearch := &models.TxSearchResponse{}
		r, err := app.Requester.Get(fmt.Sprintf(
			"http://tendermint_core:26657/tx_search?query=\"app.nonce='%v'\"", item.Nonce))
		if err != nil || r == nil {
			fmt.Println(errors.New("req error"))
		}

		r.ToJSON(txSearch)

		for _, tx := range txSearch.Result.Txs {
			blockSearch := &models.BlockSearchResponse{}
			r2, err := app.Requester.Get(fmt.Sprintf(
				"http://tendermint_core:26657/block?height=%v", tx.Height))
			if err != nil || r2 == nil {
				fmt.Println(errors.New("req2 error"))
			}

			r2.ToJSON(blockSearch)
			item.Timestamp = blockSearch.Result.Block.Header.Time
			app.DB.Save(item)
		}
	}

}

func AddTimeStamp(app *Application) {
	fmt.Println("Start Cron Every 10s")
	go saveState(app)
}
