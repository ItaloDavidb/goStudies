package controllers

import (
	"encoding/json"
	"html/template"
	"log"
	"net/http"

	"github.com/goStudies/models"
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
