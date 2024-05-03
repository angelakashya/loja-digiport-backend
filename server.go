package main

/*import (
	"encoding/json"
	"net/http"

	"github.com/angelakashya/loja-digiport-backend/model"
)

func StartServer() {
	http.HandleFunc("/produtos", produtosHandler)

	http.ListenAndServe(":8080", nil)
}

func produtosHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		getProdutos(w, r)
	} else if r.Method == "POST" {
		addProdutos(w, r)
	}
}
func addProdutos(w http.ResponseWriter, r *http.Request) {
	var produtos model.Produto
	json.NewDecoder(r.Body).Decode(&produtos)

	err := adicionaProduto(produtos)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(model.Erro{})
	}
	w.WriteHeader(http.StatusCreated)
}

func getProdutos(w http.ResponseWriter, r *http.Request) {
	queryNome := r.URL.Query().Get("nome")
	if queryNome != "" {
		produtosPorNome := buscaPorNome(queryNome)
		json.NewEncoder(w).Encode(produtosPorNome)
	} else {
		produtos := ListaProdutos
		json.NewEncoder(w).Encode(produtos)
	}
}
*/
