package repositories

import (
	"context"
	"errors"
	"github.com/luizarnoldch/neo4j-Golang/domain/entities"

	// "github.com/luizarnoldch/neo4j-Golang/infraestructure/neo4jClient"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

type ItemRepository interface {
	FindAllPersonas() ([]entities.Item, error)
	//FindAllPersonas() ([]entities.Item, error)

	SavePersonas(id int64, name string) (any, error)
	//SavePersonas(id int64, name string) (entities.Item, error)
}

/*
type ItemDatabaseNeo4j struct {
	client neo4jClient.GraphRepository
}
*/

type ItemDatabaseNeo4j struct {
	driver neo4j.DriverWithContext
}

func NewItemDatabaseNeo4j(driver neo4j.DriverWithContext) ItemDatabaseNeo4j {
	return ItemDatabaseNeo4j{driver}
}

func (db ItemDatabaseNeo4j) FindAllPersonas() ([]entities.Item, error) {
	/*
		query := "MATCH (n:Item) RETURN n LIMIT 100"
		mapString := map[string]any{}
		records, errQuery := db.client.ReadCypherQuery(query, mapString)
		if errQuery != nil {
			return nil, errors.New("unexpected database error")
		}
	*/

	res := make([]entities.Item, 0)

	ctx := context.Background()
	session := db.driver.NewSession(ctx, neo4j.SessionConfig{DatabaseName: "neo4j"})
	defer session.Close(ctx)
	_, errExecute := session.ExecuteRead(ctx, func(tx neo4j.ManagedTransaction) (any, error) {
		query := "MATCH (n:Item) RETURN n.id, n.name LIMIT 100"
		mapString := map[string]any{}
		result, err := tx.Run(ctx, query, mapString)
		if err != nil {
			return nil, errors.New("unexpected query error")
		}
		for result.Next(ctx) {
			id := result.Record().Values[0]
			name := result.Record().Values[1]
			res = append(res, entities.NewItem(id.(int64), name.(string)))
		}
		return nil, result.Err()
	})
	if errExecute != nil {
		return nil, errors.New("unexpected database error")
	}

	return res, nil
}

func (db ItemDatabaseNeo4j) SavePersonas(id int64, name string) (any, error) {
	/*
		query := "CREATE (n:Item { id: $id, name: $name }) RETURN n.id, n.name"
		mapString := map[string]any{
			"id":   id,
			"name": name,
		}
		records, errQuery := db.client.WriteCypherQuery(query, mapString)
		if errQuery != nil {
			return nil, errors.New("unexpected database error")
		}

	*/
	return "records", nil
}
