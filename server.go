package main

import (
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/GenkiHirano/gin-gorm-docker-gqlgen-crud/graph"
	"github.com/GenkiHirano/gin-gorm-docker-gqlgen-crud/graph/generated"
	"github.com/GenkiHirano/gin-gorm-docker-gqlgen-crud/infra"
	"github.com/gin-gonic/gin"
)

func graphqlHandler() gin.HandlerFunc {
	h := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))

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

func main() {
	r := gin.Default()
	r.POST("/query", graphqlHandler())
	r.GET("/", playgroundHandler())
	db := infra.NewDB()
	infra.CloseDB(db)
	r.Run()
}
