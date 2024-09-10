package db

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

var DB *gorm.DB

func OpenConnection() (*gorm.DB, error) {
	dns := "host=go_db port=5432 dbname='PostgreSQL 16' user=postgres password=12345678 connect_timeout=10 sslmode=prefer"
	db, err := gorm.Open(postgres.Open(dns), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return db, nil
}

func InitGorm() {
	var err error
	DB, err = OpenConnection()
	if err != nil {
		log.Fatalf("Erro ao conectar ao banco de dados: %v", err)
	}

	if DB == nil {
		log.Fatalf("A instância do banco de dados é nil após a inicialização")
	}

	log.Println("Conexão com o banco de dados estabelecida com sucesso")
}
