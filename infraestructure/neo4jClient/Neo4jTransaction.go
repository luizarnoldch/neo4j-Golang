package neo4jClient

import (
	"context"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

type GraphRepository interface {
	WriteCypherQuery(query string, params map[string]any) (any, error)
	ReadCypherQuery(query string, params map[string]any) (any, error)
}

type DefaultNeo4jTransaction struct {
	s   neo4j.SessionWithContext
	ctx context.Context
}

func NewNeo4jTransaction(s neo4j.SessionWithContext, ctx context.Context) DefaultNeo4jTransaction {
	return DefaultNeo4jTransaction{s, ctx}
}

func (t DefaultNeo4jTransaction) WriteCypherQuery(query string, params map[string]any) (any, error) {
	records, err := t.s.ExecuteWrite(t.ctx, func(tx neo4j.ManagedTransaction) (any, error) {
		records, err := tx.Run(t.ctx, query, params)
		if err != nil {
			return nil, err
		}
		return records, nil
	})
	if err != nil {
		return nil, err
	}
	return records, err
}

func (t DefaultNeo4jTransaction) ReadCypherQuery(query string, params map[string]any) (any, error) {
	records, err := t.s.ExecuteRead(t.ctx, func(tx neo4j.ManagedTransaction) (any, error) {
		records, err := tx.Run(t.ctx, query, params)
		if err != nil {
			return nil, err
		}
		return records.Collect(t.ctx)
	})
	if err != nil {
		return nil, err
	}
	return records, err
}
