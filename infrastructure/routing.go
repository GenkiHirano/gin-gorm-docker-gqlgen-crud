package infrastructure

import "github.com/gin-gonic/gin"

type Routing struct {
	DB *DB
	Gin *gin.Engine
	Port string
}

func NewRouting(db *DB) *Routing {
    r := &Routing{
        DB: db,
        Gin: gin.Default(),
		// 環境変数に切り出す
        Port: "8080",
    }
    r.setRouting()
    return r
}

func (r *Routing) setRouting() {
	// NewUsersController実装予定
    usersController := controllers.NewUsersController(r.DB)
    r.Gin.GET("/users/:id", func (c *gin.Context) { usersController.Get(c) })
}

func (r *Routing) Run() {
    r.Gin.Run(r.Port)
}
