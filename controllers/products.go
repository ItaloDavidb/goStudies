package controllers

import (
	"encoding/json"
	"html/template"
	"log"
	"net/http"

	"github.com/goStudies/models"
	"github.com/gorilla/mux"
)

var temp = template.Must(template.ParseGlob("templates/*.html"))

func Index(w http.ResponseWriter, r *http.Request) {
	temp.ExecuteTemplate(w, "Index", nil)
}
func GetAllProducts(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		produtos, err := models.GetAllProducts()
		if err != nil {
			http.Error(w, "Erro ao buscar produtos", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(produtos)
	}
}
func RemoveProduct(w http.ResponseWriter, r *http.Request) {
	if r.Method == "DELETE" {
		vars := mux.Vars(r)
		id := vars["id"]

		err := models.RemoveProduct(id)
		if err != nil {
			http.Error(w, "Erro ao buscar produtos", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode("Produto Removido com Sucesso")
	}
}

func UpdateProduct(w http.ResponseWriter, r *http.Request) {
	if r.Method == "PATCH" {
		vars := mux.Vars(r)
		id := vars["id"]

		// Decodifica o JSON do corpo da requisição para o struct Product
		var attProduto models.Product
		err := json.NewDecoder(r.Body).Decode(&attProduto)
		if err != nil {
			http.Error(w, "Erro ao decodificar JSON", http.StatusBadRequest)
			log.Println("Erro ao decodificar JSON:", err)
			return
		}

		if attProduto.Nome == "" || attProduto.Descricao == "" || attProduto.Preco <= 0 {
			http.Error(w, "Todos os campos devem ser preenchidos", http.StatusBadRequest)
			return
		}

		// Atualizar o produto no banco de dados
		err = models.UpdateProduct(id, &attProduto)
		if err != nil {
			http.Error(w, "Erro ao atualizar produto", http.StatusInternalServerError)
			return
		}

		// Log para depuração
		log.Printf("Produto Atualizado: ID=%s, Nome=%s, Descrição=%s, Preço=%.2f", id, attProduto.Nome, attProduto.Descricao, attProduto.Preco)

		// Retornar uma resposta de sucesso com o produto atualizado
		w.WriteHeader(http.StatusOK) // Status 200 OK
		json.NewEncoder(w).Encode(attProduto)
	}
}
func CreateProduct(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		// Decodifica o JSON do corpo da requisição para o struct Product
		var novoProduto models.Product
		err := json.NewDecoder(r.Body).Decode(&novoProduto)
		if err != nil {
			http.Error(w, "Erro ao decodificar JSON", http.StatusBadRequest)
			log.Println("Erro ao decodificar JSON:", err)
			return
		}

		// Validação simples
		if novoProduto.Nome == "" || novoProduto.Descricao == "" || novoProduto.Preco <= 0 {
			http.Error(w, "Todos os campos devem ser preenchidos", http.StatusBadRequest)
			return
		}

		// Chama a função para criar o produto no banco de dados

		err = models.CreateProduct(&novoProduto)
		if err != nil {
			http.Error(w, "Erro ao criar produto", http.StatusInternalServerError)
			return
		}

		// Log para depuração
		log.Printf("Novo produto criado: ID=%d, Nome=%s, Descrição=%s, Preço=%.2f", novoProduto.ID, novoProduto.Nome, novoProduto.Descricao, novoProduto.Preco)

		// Retornar uma resposta de sucesso
		w.WriteHeader(http.StatusCreated) // Status 201 Created
		json.NewEncoder(w).Encode(novoProduto)
	} else {
		http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
	}
}
