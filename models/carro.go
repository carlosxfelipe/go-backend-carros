package models

// d_carro
type Carro struct {
	ID          int    `gorm:"column:id"`
	Tipo        string `gorm:"column:tipo"`
	Ano         int    `gorm:"column:ano"`
	Marca       string `gorm:"column:marca"`
	Modelo      string `gorm:"column:modelo"`
	Combustivel string `gorm:"column:combustivel"`
}

func (Carro) TableName() string {
	return "d_carro"
}

// f_carrodetalhado
type CarroDetalhado struct {
	ID             int     `gorm:"column:id"`
	DataReferencia string  `gorm:"column:data_referencia"`
	Preco          float64 `gorm:"column:preco"`
}

func (CarroDetalhado) TableName() string {
	return "f_carrodetalhado"
}

// f_carrovariacao
type CarroVariacao struct {
	ID             int     `gorm:"column:id"`
	DataReferencia string  `gorm:"column:data_referencia"`
	Preco          float64 `gorm:"column:preco"`
}

func (CarroVariacao) TableName() string {
	return "f_carrovariacao"
}
