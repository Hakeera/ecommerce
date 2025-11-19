package main

import (
	"erp/config"
	"erp/internal/controller"
	"erp/internal/repository"
	"erp/internal/routes"
	"erp/internal/service"
	"html/template"
	"log"

	"github.com/labstack/echo/v4"
)

func main() {

	// Carrega variáveis de ambiente e inicializa o banco
	config.LoadEnv()
	config.InitDB()

	// Teste de conexão com o banco
	db := config.GetDB()
	if db != nil {
		log.Println("✅ Banco de dados conectado com sucesso!")
		if sqlDB, err := db.DB(); err == nil {
			if err := sqlDB.Ping(); err == nil {
				log.Println("✅ Ping no banco OK!")
			} else {
				log.Printf("❌ Erro no ping: %v", err)
			}
		}
	} else {
		log.Println("❌ Banco de dados é nil!")
	}

	// dependencies
	userRepo := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepo)
	userController := &controller.UserController{
		UserService: userService,
	}

	tmpl := template.New("").Funcs(config.TemplateFunctions)
	tmpl = template.Must(tmpl.ParseGlob("view/**/*.html"))

	renderer := &config.TemplateRenderer{
		Templates: tmpl,
	}

	// Inicialização do servidor Echo
	e := echo.New()
	e.Renderer = renderer
	e.Static("/static", "view/static")

	// Configuração das rotas da aplicação
	routes.SetUpRoutes(e, userController)

	log.Println("🚀 Servidor iniciando na porta :8080")
	e.Logger.Fatal(e.Start("127.0.0.1:8080"))
}
