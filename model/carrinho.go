package model

type Carrinho struct {
	IdUsuario       string
	ValorTotal      float64
	Id              string
	InfosProduto    []InfosProduto
	QuantidadeTotal int
}

type InfosProduto struct {
	IdProduto         string
	QuantidadeProduto string
}
