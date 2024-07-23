package models

import (
	"database/sql"
	"log"
)

type Product struct {
	ID        int // Seu banco de dados provavelmente gerenciará o ID
	Nome      string
	Descricao string
	Preco     float64
}

func CreateProduct(db *sql.DB, produto *Product) error {
	// Aqui você implementa a lógica para inserir o produto no banco de dados
	_, err := db.Exec("INSERT INTO produtos (nome, descricao, preco) VALUES ($1, $2, $3)", produto.Nome, produto.Descricao, produto.Preco)
	if err != nil {
		log.Println("Erro ao inserir produto no banco de dados:", err)
		return err
	}
	return nil
}
