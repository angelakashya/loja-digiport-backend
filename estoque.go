package main

import (
	"github.com/angelakashya/loja-digiport-backend/model"
)

func criaEstoque() []model.Produto {
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

	return produtos
}
