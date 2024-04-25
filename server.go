package main

import (
	"encoding/json"
	"net/http"

	"github.com/angelakashya/loja-digiport-backend/model"
)

func StartServer() {
	http.HandleFunc("/produtos", produtosHandler)

	http.ListenAndServe(":8080", nil)
}

func produtosHandler(w http.ResponseWriter, r *http.Request) {
	queryNome := r.URL.Query().Get("nome")
	if queryNome != "" {
		produtosPorNome := buscaPorNome(queryNome)
		json.NewEncoder(w).Encode(produtosPorNome)
	} else {
		produtos := ListaProdutos
		json.NewEncoder(w).Encode(produtos)
	}
}

func addProduto(w http.ResponseWriter, r *http.Request) {
	var produto []model.Produto
	json.NewDecoder(r.Body).Decode(&produto)
	criaEstoque()

	w.WriteHeader(http.StatusCreated)

}
