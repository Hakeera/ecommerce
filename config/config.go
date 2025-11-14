// config/config.go
// Pacote responsável pelo carregamento de variáveis de ambiente da aplicação.

package config

import (
	"log"

	"github.com/joho/godotenv"
)

// LoadEnv carrega variáveis do arquivo .env, caso exista.
// Se o arquivo não for encontrado, utiliza variáveis de ambiente do sistema.
func LoadEnv() {
	if err := godotenv.Load(); err != nil {
		log.Println("⚠️  Aviso: .env não encontrado, usando variáveis de ambiente do sistema.")
	}
}
