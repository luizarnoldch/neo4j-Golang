package application

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"github.com/luizarnoldch/neo4j-Golang/domain/repositories"
	"github.com/luizarnoldch/neo4j-Golang/domain/services"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
	"log"
	"os"
)

func Start() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error al cargar archivo .env ")
	}
	app := fiber.New()

	// -- DB
	// neo4jSession, neo4jContext := NewNeo4jClient()
	neo4jSession := NewNeo4jDriver()
	//client := neo4jClient.NewNeo4jTransaction(neo4jSession, neo4jContext)
	//neo4jTransactions :=

	// -- Repositorys
	//itemDB := repositories.NewItemDatabaseNeo4j(client)
	itemDB := repositories.NewItemDatabaseNeo4j(neo4jSession)

	// -- Service
	itemService := services.NewItemService(itemDB)

	// -- Handler
	itemHandler := ItemHandler{itemService}

	// API
	apiItem := app.Group("/item")
	apiItem.Get("/", itemHandler.GetAllItems)
	apiItem.Post("/", itemHandler.SaveItem)

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	PORT := os.Getenv("PORT")
	app.Listen(":" + PORT)
}

// func NewNeo4jClient() (neo4j.SessionWithContext, context.Context) {
func NewNeo4jDriver() neo4j.DriverWithContext {
	uri := os.Getenv("NEO4J_URI")
	port := os.Getenv("NEO4J_PORT")
	username := os.Getenv("NEO4J_USERNAME")
	password := os.Getenv("NEO4J_PASSWORD")
	//instance := os.Getenv("AURA_INSTANCENAME")

	auth := neo4j.BasicAuth(username, password, "")
	uriPort := fmt.Sprintf("%s:%s", uri, port)
	driver, err := neo4j.NewDriverWithContext(uriPort, auth)
	if err != nil {
		panic(err)
	}
	// ctx := context.Background()
	// defer driver.Close(ctx)
	// session := driver.NewSession(ctx, neo4j.SessionConfig{DatabaseName: "neo4j"})
	// defer session.Close(ctx)
	// return session, ctx
	return driver
}
