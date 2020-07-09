package models

type Tx struct {
	Operation  int    `json:"operation"`
	DomainName string `json:"domain_name"`
	By         string `json:"by"`
	Nonce      string `json:"nonce"`
}
