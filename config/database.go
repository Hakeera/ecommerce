// config/database.go
// Inicializa e mantém a conexão com o banco de dados PostgreSQL via GORM.

package config

import (
	"erp/internal/model"
	"fmt"
	"log"
	"os"
	"sync"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	// DB é a instância global de conexão com o banco.
	DB *gorm.DB

	// once garante inicialização única e thread-safe.
	once sync.Once
)

// InitDB configura e conecta ao banco PostgreSQL.
// Usa variáveis de ambiente para montar a string de conexão (DSN).
// Executa migrações automáticas dos modelos principais.
func InitDB() {
	once.Do(func() {
		dbUser := os.Getenv("DB_USER")
		dbPassword := os.Getenv("DB_PASSWORD")
		dbName := os.Getenv("DB_NAME")
		dbHost := os.Getenv("DB_HOST")
		dbPort := os.Getenv("DB_PORT")
		sslmode := os.Getenv("DB_SSLMODE")

		dsn := fmt.Sprintf(
			"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
			dbHost, dbUser, dbPassword, dbName, dbPort, sslmode,
		)

		var err error
		DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err != nil {
			log.Fatalf("❌ Erro ao conectar ao banco: %v", err)
		}
		log.Println("✅ Banco de dados conectado com sucesso!")

		// Migração automática de tabelas principais
		if err := DB.AutoMigrate(
			&model.User{},
			&model.Product{},
		); err != nil {
			log.Fatalf("❌ Erro na migração: %v", err)
		}
	})
}

// GetDB retorna a instância ativa do banco.
// Encerra a aplicação se o banco não tiver sido inicializado.
func GetDB() *gorm.DB {
	if DB == nil {
		log.Fatal("❌ Banco de dados não inicializado!")
	}
	return DB
}
