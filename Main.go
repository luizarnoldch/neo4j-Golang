package main

import "github.com/luizarnoldch/neo4j-Golang/application"

func main() {
	application.Start()
}

/*
func main() {

	//NEO4J_URI := "neo4j+s://29be56aa.databases.neo4j.io"
	//NEO4J_USERNAME := "neo4j"
	//NEO4J_PASSWORD := "ugyIT258-KLMmC2i2Uir5GN4A2x_tO743U8BBUwqW6c"
	//AURA_INSTANCENAME := "Instance01"

	//Conection_URL := 29be56aa.databases.neo4j.io:7687

	// Aura requires you to use "neo4j+s" protocol
	// (You need to replace your connection details, username and password)
	uri := "neo4j+s://29be56aa.databases.neo4j.io:7687"
	auth := neo4j.BasicAuth("neo4j", "ugyIT258-KLMmC2i2Uir5GN4A2x_tO743U8BBUwqW6c", "")
	// You typically have one driver instance for the entire application. The
	// driver maintains a pool of database connections to be used by the sessions.
	// The driver is thread safe.
	driver, err := neo4j.NewDriverWithContext(uri, auth)
	if err != nil {
		panic(err)
	}
	// You can specify custom contexts to control the execution of the driver operations
	// This one never cancels, never times out and is used in subsequent calls
	// You can instead specify different contexts for different operations
	// Read more about contexts in https://pkg.go.dev/context
	ctx := context.Background()
	// Don't forget to close the driver connection when you are finished with it
	defer driver.Close(ctx)
	// Create a session to run transactions in. Sessions are lightweight to
	// create and use. Sessions are NOT thread safe.
	session := driver.NewSession(ctx, neo4j.SessionConfig{DatabaseName: "neo4j"})
	defer session.Close(ctx)
	// WriteTransaction retries the operation in case of transient errors by
	// invoking the anonymous function multiple times until it succeeds.
	records, err := session.ExecuteWrite(ctx,
		func(tx neo4j.ManagedTransaction) (any, error) {
			// To learn more about the Cypher syntax, see https://neo4j.com/docs/cypher-manual/current/
			// The Reference Card is also a good resource for keywords https://neo4j.com/docs/cypher-refcard/current/
			createRelationshipBetweenPeopleQuery := `
				MERGE (p1:Person { name: $person1_name })
				MERGE (p2:Person { name: $person2_name })
				MERGE (p1)-[:KNOWS]->(p2)
				RETURN p1, p2`
			result, err := tx.Run(ctx, createRelationshipBetweenPeopleQuery, map[string]any{
				"person1_name": "Alice",
				"person2_name": "David",
			})
			if err != nil {
				// Return the error received from driver here to indicate rollback,
				// the error is analyzed by the driver to determine if it should try again.
				return nil, err
			}
			// Collects all records and commits the transaction (as long as
			// Collect doesn't return an error).
			// Beware that Collect will buffer the records in memory.
			return result.Collect(ctx)
		})
	if err != nil {
		panic(err)
	}
	for _, record := range records.([]*neo4j.Record) {
		firstPerson := record.Values[0].(neo4j.Node)
		fmt.Printf("First: '%s'\n", firstPerson.Props["name"].(string))
		secondPerson := record.Values[1].(neo4j.Node)
		fmt.Printf("Second: '%s'\n", secondPerson.Props["name"].(string))
	}

	// Now read the created persons. By using ExecuteRead method a connection
	// to a read replica can be used which reduces load on writer nodes in cluster.
	_, err = session.ExecuteRead(ctx, func(tx neo4j.ManagedTransaction) (any, error) {
		// Code within this function might be invoked more than once in case of
		// transient errors.

		readPersonByName := `
				MATCH (n:Item) RETURN n.id, n.name LIMIT 100`
		result, err := tx.Run(ctx, readPersonByName, map[string]any{})


			readPersonByName := `
				MATCH (p:Person)
				WHERE p.name = $person_name
				RETURN p.name AS name`
			result, err := tx.Run(ctx, readPersonByName, map[string]any{
				"person_name": "Alice",
			})


		if err != nil {
			return nil, err
		}

		// Iterate over the result within the transaction instead of using
		// Collect (just to show how it looks...). Result.Next returns true
		// while a record could be retrieved, in case of error result.Err()
		// will return the error.

		for result.Next(ctx) {
			//record1 := result.Record()
			record2 := result.Record().Values[1]

			//fmt.Println(record)
			//fmt.Println(record1)
			fmt.Println(record2)
			//fmt.Println(record2.(string))
			//fmt.Printf("Person name: '%s' \n", name)
			//fmt.Printf("Person name: '%s' \n", result.Record().Values[0].(string))
		}


			for result.Next(ctx) {
				fmt.Println(result)
				fmt.Println()
				fmt.Println(result.Record())
				fmt.Println()
				fmt.Println(result.Record().Values[0])
				fmt.Println()
				//fmt.Printf("Person name: '%s' \n", result.Record().Values[0].(string))
			}


		// Again, return any error back to driver to indicate rollback and
		// retry in case of transient error.
		return nil, result.Err()
	})
	if err != nil {
		panic(err)
	}
}
*/

/*
func main() {

	//NEO4J_URI := "neo4j+s://29be56aa.databases.neo4j.io"
	//NEO4J_USERNAME := "neo4j"
	//NEO4J_PASSWORD := "ugyIT258-KLMmC2i2Uir5GN4A2x_tO743U8BBUwqW6c"
	//AURA_INSTANCENAME := "Instance01"

	//Conection_URL := 29be56aa.databases.neo4j.io:7687

	// Neo4j 4.0, defaults to no TLS therefore use bolt:// or neo4j://
	// Neo4j 3.5, defaults to self-signed certificates, TLS on, therefore use bolt+ssc:// or neo4j+ssc://
	dbUri := "neo4j+s://29be56aa.databases.neo4j.io:7687"
	driver, err := neo4j.NewDriverWithContext(dbUri, neo4j.BasicAuth("neo4j", "ugyIT258-KLMmC2i2Uir5GN4A2x_tO743U8BBUwqW6c", ""))
	if err != nil {
		panic(err)
	}
	// Starting with 5.0, you can control the execution of most driver APIs
	// To keep things simple, we create here a never-cancelling context
	// Read https://pkg.go.dev/context to learn more about contexts
	ctx := context.Background()
	// Handle driver lifetime based on your application lifetime requirements  driver's lifetime is usually
	// bound by the application lifetime, which usually implies one driver instance per application
	// Make sure to handle errors during deferred calls
	defer driver.Close(ctx)
	item, err := insertItem(ctx, driver)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%v\n", *item)
}

func insertItem(ctx context.Context, driver neo4j.DriverWithContext) (*Item, error) {
	// Sessions are short-lived, cheap to create and NOT thread safe. Typically create one or more sessions
	// per request in your web application. Make sure to call Close on the session when done.
	// For multi-database support, set sessionConfig.DatabaseName to requested database
	// Session config will default to write mode, if only reads are to be used configure session for
	// read mode.
	session := driver.NewSession(ctx, neo4j.SessionConfig{})
	defer session.Close(ctx)
	result, err := session.ExecuteWrite(ctx, createItemFn(ctx))
	if err != nil {
		return nil, err
	}
	return result.(*Item), nil
}

func createItemFn(ctx context.Context) neo4j.ManagedTransactionWork {
	return func(tx neo4j.ManagedTransaction) (any, error) {
		records, err := tx.Run(ctx, "CREATE (n:Item { id: $id, name: $name }) RETURN n.id, n.name", map[string]any{
			"id":   1,
			"name": "Item 1",
		})
		// In face of driver native errors, make sure to return them directly.
		// Depending on the error, the driver may try to execute the function again.
		if err != nil {
			return nil, err
		}
		record, err := records.Single(ctx)
		if err != nil {
			return nil, err
		}
		// You can also retrieve values by name, with e.g. `id, found := record.Get("n.id")`
		return &Item{
			Id:   record.Values[0].(int64),
			Name: record.Values[1].(string),
		}, nil
	}
}

type Item struct {
	Id   int64
	Name string
}
*/
