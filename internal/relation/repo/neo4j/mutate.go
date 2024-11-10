package neo4j

import (
	"context"

	"github.com/neo4j/neo4j-go-driver/v5/neo4j"

	rerr "github.com/Karzoug/meower-relation-service/internal/relation/repo"
)

func (r repo) Follow(ctx context.Context, reqUserID, targetUserID string) error {
	res, err := neo4j.ExecuteQuery(ctx, r.driver,
		`MATCH (u1:User{id: $ruser})
         MATCH (u2:User{id: $tuser})
         CREATE (u1)-[:FOLLOWS {start: date()}]->(u2)`,
		map[string]any{
			"ruser": reqUserID,
			"tuser": targetUserID,
		}, neo4j.EagerResultTransformer,
		neo4j.ExecuteQueryWithDatabase(r.cfg.DBName))
	if err != nil {
		return err
	}

	if res.Summary.Counters().RelationshipsCreated() == 0 {
		return rerr.ErrNoAffected
	}

	return nil
}

func (r repo) Unfollow(ctx context.Context, reqUserID, targetUserID string) error {
	res, err := neo4j.ExecuteQuery(ctx, r.driver,
		`MATCH (u1:User{id: $ruser})-[f:FOLLOWS]->(u2:User{id: $tuser})
         DELETE f`,
		map[string]any{
			"ruser": reqUserID,
			"tuser": targetUserID,
		}, neo4j.EagerResultTransformer,
		neo4j.ExecuteQueryWithDatabase(r.cfg.DBName))
	if err != nil {
		return err
	}

	if res.Summary.Counters().RelationshipsDeleted() == 0 {
		return rerr.ErrNoAffected
	}

	return nil
}

func (r repo) Hide(ctx context.Context, reqUserID, targetUserID string) error {
	res, err := neo4j.ExecuteQuery(ctx, r.driver,
		`MATCH (u1:User{id: $ruser})
         MATCH (u2:User{id: $tuser})
         CREATE (u1)-[:HIDES]->(u2)`,
		map[string]any{
			"ruser": reqUserID,
			"tuser": targetUserID,
		}, neo4j.EagerResultTransformer,
		neo4j.ExecuteQueryWithDatabase(r.cfg.DBName))
	if err != nil {
		return err
	}

	if res.Summary.Counters().RelationshipsCreated() == 0 {
		return rerr.ErrNoAffected
	}

	return nil
}

func (r repo) Unhide(ctx context.Context, reqUserID, targetUserID string) error {
	res, err := neo4j.ExecuteQuery(ctx, r.driver,
		`MATCH (u1:User{id: $ruser})-[h:HIDES]->(u2:User{id: $tuser})
         DELETE h`,
		map[string]any{
			"ruser": reqUserID,
			"tuser": targetUserID,
		}, neo4j.EagerResultTransformer,
		neo4j.ExecuteQueryWithDatabase(r.cfg.DBName))
	if err != nil {
		return err
	}

	if res.Summary.Counters().RelationshipsDeleted() == 0 {
		return rerr.ErrNoAffected
	}

	return nil
}
