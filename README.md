# Use it in your project

Your main.go looks like this: usual endpoints, then the doc conf (use .env or something to make it clear what env it is).

* Create Gin endpoints
* Add part between // START OF MICROAPIDOC & // END
* When service starts, it collects annotations from the controllers and reads all pathes

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

	products := r.Group("/products")
	{
		products.GET("", rest.GetProducts)
		products.POST("create", rest.PostNewProducts)
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
	microApiDocConf := microapidoc.GeneralDoc{
		BuildTag:            GitTag,
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
	docController := microapidoc.NewMicroapidoc(microApiDocConf)
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

And your controller's annotation look like this:

```
package controllers

import (
	"apidocexample/internal/models"
	"fmt"

	"github.com/gin-gonic/gin"
)

// #Group GetAuth
// #Summary GetAuth controller to controll auth
// #GoodResponse string
// #BadResponse string
// #QueryParameters id uuid.UUID
// #QueryParameters name string
// #HeaderParameters xxx-rate uuid.UUID
// #HeaderParameters yyy-date string
// #Label production, development
func (c *Controller) GetAuth(ctx *gin.Context) {
	fmt.Println(ctx.GetHeader("Bearer"))
	fmt.Println(ctx.GetHeader("Token"))
	fmt.Println(ctx.GetHeader("xxx-rate"))
	fmt.Println(ctx.GetHeader("yyy-date"))
	fmt.Println(ctx.Query("id"))
	fmt.Println(ctx.Query("name"))
	ctx.Data(200, "text/plain", []byte("Auth controller to controll auth"))
	return
}
```

