package model

type Produto struct {
	Id                  string  `json:"id"`
	Nome                string  `json:"nome`
	Preco               float64 `json:"preco"`
	Descricao           string  `json:"descricao"`
	Imagem              string  `json:"imagem"`
	QuantidadeEmEstoque int     `json: "quantidadeEmEstoque"`
}
