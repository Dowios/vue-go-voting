package main

import (
	"fmt"
	_ "net/http"
	"./handlers"

  "database/sql"
  _ "github.com/lib/pq"
	"github.com/gin-gonic/gin"
)

var router *gin.Engine
var db *sql.DB
func main() {
	// Set Gin to production mode
	gin.SetMode(gin.ReleaseMode)

	// Set the router as the default one provided by Gin
	router = gin.Default()
	db = initSql()
	// Process the templates at the start so that they don't have to be loaded
	// from the disk again. This makes serving HTML pages very fast.
	router.LoadHTMLGlob("public/*")

	// Initialize the routes
	router.GET("/", func(c *gin.Context) {
    c.HTML(200, "index.html", gin.H{})
  })

	poll := router.Group("/polls")

	poll.Use(DatabaseMiddleware(db))
	{
		poll.GET("/", handlers.GetPolls)
		poll.PUT("/:index", handlers.UpdatePoll)
	}

	// Start serving the application
	router.Run()
}
const (
    c_host     = "localhost"
    c_port     = 5432
    c_user     = "gouser"
    c_password = "justgo"
    c_dbname   = "voting"
)
func initSql() *sql.DB{
  psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
      "password=%s dbname=%s sslmode=disable",
      c_host, c_port, c_user, c_password, c_dbname)
  db, err := sql.Open("postgres", psqlInfo)
  if err != nil {
      panic(err)
  }

  err = db.Ping()
  if err != nil {
      panic(err)
  }

  fmt.Println("Successfully connected!")
	return db
}

func DatabaseMiddleware(db *sql.DB) gin.HandlerFunc {
  // Do some initialization logic here
  // Foo()
  return func(c *gin.Context) {
		c.Set("db", db)
    c.Next()
  }
}
