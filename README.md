# Use it in your project

should look something like this: first do your endpoints then, the part that generates doc and creates 2 endpoints

```
package main 

import (
	"apidocexample/cmd/controllers"

	"github.com/triiberg/microapidoc"
	"github.com/gin-gonic/gin"
)


var GitTag string = "dev"

func main() {

	r := gin.Default()

	rest := controllers.NewController()

	auth := r.Group("/auth")
	{
		auth.GET("", rest.GetAccountData)
		auth.GET("/:id", rest.GetAccountData)
		auth.GET("/get-header-auth", rest.GetAuth)
	}

	apiV2 := r.Group("/products")
	{
		apiV2.GET("", rest.GetProducts)
		apiV2.POST("create", rest.PostNewProducts)
	}

	// START OF MICROAPIDOC
	// read routes
	var routes []microapidoc.RouteInfo
	for _, route := range r.Routes() {
		routes = append(routes, microapidoc.RouteInfo{
			Method:      route.Method,
			Path:        route.Path,
			HandlerFunc: route.Handler,
		})
	}
	// generic settings
	luunjaHandler := microapidoc.GeneralDoc{
		BuildTag:            "v1.0.0",
		SearchControllersIn: "./cmd/controllers",
		AllRoutes:           routes,
		Name:                "API Doc test project",
		BaseUrl:             "http://localhost:8083",
		HeaderColor:         "blue",
		AuthHeaderDefaultOn: true,
		AuthHeaderNames:     []string{"Bearer", "Token"},
		HighlightResponseHeaders: []string{
			"x-header:ok:green",
			"x-header:error:red",
		},
	}
	// set handlers
	docController := microapidoc.NewMicroapidoc(luunjaHandler)
	doc := r.Group("/microapidocs")
	{
		doc.GET("/doc.json", docController.DocHAndler)
		doc.GET("/", docController.DocIndexHAndler)
	}
	// END OF MICROAPIDOC

	// now run it all
	r.Run(":8083")
}

```