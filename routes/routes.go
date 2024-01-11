package routes

import (
	"B51-MICRO-FEATURE-GO/controller"
	"B51-MICRO-FEATURE-GO/pkg/mysql"
	"B51-MICRO-FEATURE-GO/services"

	"github.com/labstack/echo/v4"
)

func Routes(e *echo.Group) {
	//Article
	r := services.ArticleRepository(mysql.DB)
	cr := controller.ArticleController(r)

	e.GET("/articles", cr.FindArticles)
	e.GET("/article/:id", cr.GetArticle)

	//Paslon
	rp := services.PaslonRepository(mysql.DB)
	cp := controller.PaslonController(rp)

	e.GET("/paslons", cp.FindPaslons)
	e.GET("/paslon/:id", cp.GetPaslon)

	// Partai
	rpt := services.PartaiRepository(mysql.DB)
	cpt := controller.PartaiController(rpt)

	e.GET("/partais", cpt.FindPartais)
	e.GET("/partai/:id", cpt.GetPartai)

	// User
	ru := services.UserRepository(mysql.DB)
	cu := controller.UserController(ru)

	e.GET("/users", cu.GetUsers)
	e.GET("/user/:id", cu.GetUser)
}
