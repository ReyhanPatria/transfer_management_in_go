package entity

type Transaction struct {
	BcaId  string `gorm:"column:BCAID" json:"bcaId"`
	Type   string `gorm:"column:TYPE" json:"type"`
	Key    string `gorm:"column:KEY" json:"key"`
	Amount int    `gorm:"column:AMOUNT " json:"amount"`
	Action string `gorm:"column:ACTION" json:"action"`
}
