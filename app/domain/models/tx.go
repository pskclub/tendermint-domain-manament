package models

import "time"

type Tx struct {
	ID         int64     `json:"id" gorm:"AUTO_INCREMENT;primary_key"`
	Operation  string    `json:"operation"`
	DomainName string    `json:"domain_name"`
	Owner      string    `json:"owner"`
	Receiver   *string   `json:"receiver"`
	Nonce      string    `json:"nonce"`
	CreatedAt  time.Time `json:"created_at"`
}

func (Tx) TableName() string {
	return "tx"
}
