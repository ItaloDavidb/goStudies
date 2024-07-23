package models

import (
	"log"

	db "github.com/goStudies/infra"
)

type Product struct {
	ID        int // Seu banco de dados provavelmente gerenciará o ID
	Nome      string
	Descricao string
	Preco     float64
}

func CreateProduct(produto *Product) error {
	db := db.ConnectToDb()
	_, err := db.Exec("INSERT INTO produtos (nome, descricao, preco) VALUES ($1, $2, $3)", produto.Nome, produto.Descricao, produto.Preco)
	if err != nil {
		log.Println("Erro ao inserir produto no banco de dados:", err)
		return err
	}
	defer db.Close()
	return nil
}
func GetAllProducts() ([]Product, error) {
	db := db.ConnectToDb()
	defer db.Close()

	rows, err := db.Query("SELECT * FROM produtos")
	if err != nil {
		log.Println("Erro ao buscar produtos no banco de dados:", err)
		return nil, err
	}
	defer rows.Close()

	var produtos []Product

	for rows.Next() {
		var p Product
		err := rows.Scan(&p.ID, &p.Nome, &p.Descricao, &p.Preco)
		if err != nil {
			log.Println("Erro ao ler os dados do produto:", err)
			return nil, err
		}
		produtos = append(produtos, p)
	}

	if err = rows.Err(); err != nil {
		log.Println("Erro no resultado dos produtos:", err)
		return nil, err
	}

	return produtos, nil
}
