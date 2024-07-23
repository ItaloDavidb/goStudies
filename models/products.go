package models

import (
	"log"

	db "github.com/goStudies/infra"
)

type Product struct {
	ID        int 
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

	rows, err := db.Query("SELECT * FROM produtos")
	if err != nil {
		log.Println("Erro ao buscar produtos no banco de dados:", err)
		return nil, err
	}

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
	defer rows.Close()
	defer db.Close()
	return produtos, nil
}
func UpdateProduct(id string, attProduto *Product) error {
	db := db.ConnectToDb()
	_, err := db.Exec("UPDATE produtos SET nome = $1, descricao = $2, preco = $3 WHERE id = $4", attProduto.Nome, attProduto.Descricao, attProduto.Preco, id)
	if err != nil {
		log.Println("Erro ao Atualizar produto no banco de dados:", err)
		return err
	}
	defer db.Close()
	return nil
}
func RemoveProduct(id string) error {
	db := db.ConnectToDb()
	_, err := db.Exec("DELETE FROM produtos where id = $1", id)
	if err != nil {
		log.Println("Erro ao Remover produto no banco de dados:", err)
		return err
	}
	defer db.Close()
	return nil
}
