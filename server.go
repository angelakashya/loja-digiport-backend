package main

import (
	"encoding/json"
	"net/http"
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

/*func addEmployee(w http.ResponseWriter, r *http.Request) {
	var employee model.Employee
	json.NewDecoder(r.Body).Decode(&employee)

	registerEmployee(employee)

	w.WriteHeader(http.StatusCreated)
}*/
