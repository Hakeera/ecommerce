package dependency

import (
	"ecommerce/config"
	"ecommerce/internal/controller"
	"ecommerce/internal/repository"
	"ecommerce/internal/service"
	"html/template"
)

type Container struct {
	UserController    *controller.UserController
	ProductController *controller.ProductController
	Renderer          *config.TemplateRenderer
}

func NewContainer() *Container {
	// env + db
	config.LoadEnv()
	config.InitDB()
	db := config.GetDB()

	// user
	userRepo := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepo)
	userController := &controller.UserController{
		UserService: userService,
	}

	// product
	productRepo := repository.NewProductRepository(db)
	productService := service.NewProductService(productRepo)
	productController := &controller.ProductController{
		ProductService: productService,
	}

	// templates
	tmpl := template.New("").Funcs(config.TemplateFunctions)
	tmpl = template.Must(tmpl.ParseGlob("view/**/*.html"))
	renderer := &config.TemplateRenderer{Templates: tmpl}

	return &Container{
		UserController:    userController,
		ProductController: productController,
		Renderer:          renderer,
	}
}
