package neo4j

import (
	"context"

	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
	"github.com/rs/xid"

	rerr "github.com/Karzoug/meower-relation-service/internal/relation/repo"
)

func (r repo) CreateUser(ctx context.Context, id xid.ID) error {
	res, err := neo4j.ExecuteQuery(ctx, r.driver,
		`CREATE (u:User {id: $id})`,
		map[string]any{
			"id": id.String(),
		}, neo4j.EagerResultTransformer,
		neo4j.ExecuteQueryWithDatabase(r.cfg.DBName))

	if res.Summary.Counters().NodesCreated() == 0 {
		return rerr.ErrAlreadyExists
	}

	return err
}

func (r repo) DeleteUser(ctx context.Context, id xid.ID) error {
	_, err := neo4j.ExecuteQuery(ctx, r.driver,
		`MATCH (u:User {id: $id})
		 DETACH DELETE u`,
		map[string]any{
			"id": id.String(),
		}, neo4j.EagerResultTransformer,
		neo4j.ExecuteQueryWithDatabase(r.cfg.DBName))

	return err
}

func (r repo) Follow(ctx context.Context, sourceUserID, targetUserID xid.ID) error {
	res, err := neo4j.ExecuteQuery(ctx, r.driver,
		`MATCH (u1:User{id: $ruser})
         MATCH (u2:User{id: $tuser})
         CREATE (u1)-[:FOLLOWS {start: date()}]->(u2)`,
		map[string]any{
			"ruser": sourceUserID.String(),
			"tuser": targetUserID.String(),
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

func (r repo) Unfollow(ctx context.Context, sourceUserID, targetUserID xid.ID) error {
	res, err := neo4j.ExecuteQuery(ctx, r.driver,
		`MATCH (u1:User{id: $ruser})-[f:FOLLOWS]->(u2:User{id: $tuser})
         DELETE f`,
		map[string]any{
			"ruser": sourceUserID.String(),
			"tuser": targetUserID.String(),
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

func (r repo) Mute(ctx context.Context, sourceUserID, targetUserID xid.ID) error {
	res, err := neo4j.ExecuteQuery(ctx, r.driver,
		`MATCH (u1:User{id: $ruser})
         MATCH (u2:User{id: $tuser})
         CREATE (u1)-[:MUTES]->(u2)`,
		map[string]any{
			"ruser": sourceUserID.String(),
			"tuser": targetUserID.String(),
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

func (r repo) Unmute(ctx context.Context, sourceUserID, targetUserID xid.ID) error {
	res, err := neo4j.ExecuteQuery(ctx, r.driver,
		`MATCH (u1:User{id: $ruser})-[h:MUTES]->(u2:User{id: $tuser})
         DELETE h`,
		map[string]any{
			"ruser": sourceUserID.String(),
			"tuser": targetUserID.String(),
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
