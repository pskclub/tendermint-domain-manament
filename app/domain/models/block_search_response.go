package models

type TxSearchResponse struct {
	Jsonrpc string `json:"jsonrpc"`
	ID      string `json:"id"`
	Result  struct {
		Txs []struct {
			Hash     string `json:"hash"`
			Height   string `json:"height"`
			Index    int    `json:"index"`
			TxResult struct {
				Code      int         `json:"code"`
				Data      interface{} `json:"data"`
				Log       string      `json:"log"`
				Info      string      `json:"info"`
				GasWanted string      `json:"gasWanted"`
				GasUsed   string      `json:"gasUsed"`
				Events    []struct {
					Type       string `json:"type"`
					Attributes []struct {
						Key   string `json:"key"`
						Value string `json:"value"`
					} `json:"attributes"`
				} `json:"events"`
				Codespace string `json:"codespace"`
			} `json:"tx_result"`
			Tx string `json:"tx"`
		} `json:"txs"`
		TotalCount string `json:"total_count"`
	} `json:"result"`
}
