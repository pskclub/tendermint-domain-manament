package domain

import (
	"fmt"
	"github.com/pskclub/tendermint-domain-manament/app/utils"
	"github.com/tendermint/tendermint/abci/types"
	"strconv"
)

func queryTX(app *Application, req types.RequestQuery) types.ResponseQuery {
	fmt.Println("queryTX...")
	utils.LogStruct(req)

	switch req.Path {
	case "size":
		return types.ResponseQuery{
			Value: []byte(fmt.Sprintf("%v", app.Size)),
			Log:   strconv.FormatInt(app.Size, 10),
		}
	case "height":
		return types.ResponseQuery{
			Value: []byte(fmt.Sprintf("%v", app.Height)),
			Log:   strconv.FormatInt(app.Height, 10),
		}
	case "owner":
		return types.ResponseQuery{
			Value: []byte(fmt.Sprintf("%v", app.Height)),
			Log:   strconv.FormatInt(app.Height, 10),
		}
	case "domain":
		return types.ResponseQuery{
			Value: []byte(fmt.Sprintf("%v", app.Height)),
			Log:   strconv.FormatInt(app.Height, 10),
		}
	default:
		return types.ResponseQuery{Log: fmt.Sprintf("Invalid query path. Expected Hash or tx, got %v", req.Path)}
	}
}
