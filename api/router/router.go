package router

import (
	"context"
	"errors"
	"time"
	"todo_graphql/api/service"
	"todo_graphql/api/utils"
	"todo_graphql/constants"
	"todo_graphql/graph"

	// lru "github.com/hashicorp/golang-lru"

	"github.com/99designs/gqlgen/graphql"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
	nice "github.com/ekyoung/gin-nice-recovery"
	"github.com/gin-gonic/gin"
)

func graphqlHandler(dependencies service.Services, introspectionEnabled bool) gin.HandlerFunc {
	h := handler.New(
		// generated.NewExecutableSchema(
		// 	generated.Config{
		// 		Resolvers: &graph.Resolver{
		// 			Services: dependencies,
		// 		},
		// 	},
		// ),
		generated.NewExecutableSchema(
			generated.Config{
				Resolvers: &graph.Resolver{
					Services: dependencies,
				},
			},
		),
	)	

	h.SetRecoverFunc(func(ctx context.Context, err interface{}) error {
		e, ok := err.(error)
		if !ok {
			errString := err.(string)
			return utils.HandleError(ctx, constants.InternalServerError, errors.New(errString))
		}
		return utils.HandleError(ctx, constants.InternalServerError, e)
	})
	h.AddTransport(transport.Websocket{
		KeepAlivePingInterval: 10 * time.Second,
	})
	h.AddTransport(transport.Options{})
	h.AddTransport(transport.GET{})
	h.AddTransport(transport.POST{})
	h.AddTransport(transport.MultipartForm{})

	h.SetQueryCache(graphql.MapCache{})
	//h.SetQueryCache(lru.New(1000))

	if introspectionEnabled {
		h.Use(extension.Introspection{})
	}
	h.Use(extension.AutomaticPersistedQuery{
		//Cache: lru.New(100),
		Cache: graphql.MapCache{},
	})

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

func playgroundHandler() gin.HandlerFunc {
	h := playground.Handler("GraphQL", "/query")

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}


// Init sets up the route for the REST API
func Init(dependencies service.Services) *gin.Engine {
	router := gin.Default()

	// panic recovery
	router.Use(nice.Recovery(func(c *gin.Context, err interface{}) {
		var e error
		if err == nil {
			e = nil
		} else {
			e = err.(error)
		}
		utils.HandleError(c, constants.InternalServerError, e)
	}))

	//introspectionEnabled := true

	router.NoRoute(func(c *gin.Context) {
		utils.HandleError(c, constants.NotFound, errors.New("not found"))
	})

	introspectionEnabled := true
	router.POST("/query", graphqlHandler(dependencies, introspectionEnabled))

	router.GET("/", playgroundHandler())
	
	return router
}
