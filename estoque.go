package main

import (
	"errors"

	"github.com/angelakashya/loja-digiport-backend/model"
)

var ListaProdutos []model.Produto = []model.Produto{}

func criaEstoque() {
	produtos := []model.Produto{
		{
			Id:                  "1",
			Nome:                "Cabo Vela Uno",
			Preco:               69.90,
			Descricao:           "Cabo Vela Uno 1985 86 87 88 1989 1.3 1.5 8v SCT42",
			Imagem:              "blob:https://web.whatsapp.com/574d7341-f19a-409d-be93-becf32542f61",
			QuantidadeEmEstoque: 5,
		},

		{
			Id:                  "2",
			Nome:                "Cabo Vela Palio/Siena/Idea/Punto/Uno",
			Preco:               69.90,
			Descricao:           "Cabo Vela Palio Siena Idea Punto Uno 2004 A 2012 Fire 8v CVMT0902",
			Imagem:              "blob:https://web.whatsapp.com/3ed57519-0039-4bb4-833b-704d37174e29",
			QuantidadeEmEstoque: 10,
		},

		{
			Id:                  "3",
			Nome:                "Bobina Iginicao Sportage 2005 A 2006 2.0 16v",
			Preco:               150.00,
			Descricao:           "Bobina Iginicao Sportage 2005 A 2006 2.0 16v",
			Imagem:              "blob:https://web.whatsapp.com/55c59e90-d51c-481d-9d42-a42ea9da2198",
			QuantidadeEmEstoque: 8,
		},
	}

	ListaProdutos = produtos
}

func buscaPorNome(nome string) []model.Produto {

	resultado := []model.Produto{}

	for _, produto := range ListaProdutos {
		if produto.Nome == nome {
			resultado = append(resultado, produto)
		}
	}
	return resultado
}

func adicionaProduto(produtoNovo model.Produto) error {
	for _, produto := range ListaProdutos {
		if produtoNovo.Id == produto.Id {
			return errors.New("Id de produto já existe")
		}
	}
	ListaProdutos = append(ListaProdutos, produtoNovo)
	return nil
}
