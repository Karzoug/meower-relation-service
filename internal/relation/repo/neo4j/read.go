package neo4j

import (
	"context"
	"fmt"

	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j/db"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j/dbtype"
	"github.com/rs/xid"

	"github.com/Karzoug/meower-relation-service/internal/relation/entity"
)

func (r repo) ListFollowers(ctx context.Context, userID, afterID xid.ID, limit int) ([]entity.User, xid.ID, error) {
	const q = `MATCH (u:User{id: $tuser})<-[f:FOLLOWS]-(followers:User %s)
OPTIONAL MATCH (u)<-[h:MUTES]-(followers)
WITH followers.id as id, f.start as follow_start, h IS NOT NULL as muted
ORDER BY id
RETURN id, follow_start, muted
%s`

	pms := map[string]any{
		"tuser": userID.String(),
	}

	var limitQ, afterQ string
	if !afterID.IsNil() {
		afterQ = "WHERE followers.id>$after_id"
		pms["after_id"] = afterID.String()
	}
	if limit != -1 {
		limitQ = "LIMIT $limit" //nolint:goconst
		pms["limit"] = limit
	}

	query := fmt.Sprintf(q, afterQ, limitQ)

	result, err := neo4j.ExecuteQuery(ctx,
		r.driver,
		query,
		pms,
		neo4j.EagerResultTransformer,
		neo4j.ExecuteQueryWithDatabase(r.cfg.DBName),
		neo4j.ExecuteQueryWithReadersRouting(),
	)
	if err != nil {
		return nil, xid.NilID(), err
	}

	users := make([]entity.User, len(result.Records))
	for i := 0; i < len(result.Records); i++ {
		users[i] = toUser(result.Records[i])
	}

	if limit != -1 && len(users) == limit {
		nextID := users[len(users)-1].ID
		return users, nextID, nil
	}

	return users, xid.NilID(), nil
}

func (r repo) ListFollowings(ctx context.Context, userID, afterID xid.ID, limit int) ([]entity.User, xid.ID, error) {
	const q = `
MATCH (u:User{id: $tuser})-[f:FOLLOWS]->(followings:User %s)
OPTIONAL MATCH (u)-[h:MUTES]->(followings)
WITH followings.id as id, f.start as follow_start, h IS NOT NULL as muted
ORDER BY id
RETURN id, follow_start, muted
%s`

	pms := map[string]any{
		"tuser": userID.String(),
	}

	var limitQ, afterQ string
	if !afterID.IsNil() {
		afterQ = "WHERE followings.id>$after_id"
		pms["after_id"] = afterID.String()
	}
	if limit != -1 {
		limitQ = "LIMIT $limit"
		pms["limit"] = limit
	}

	query := fmt.Sprintf(q, afterQ, limitQ)

	result, err := neo4j.ExecuteQuery(ctx,
		r.driver,
		query,
		pms,
		neo4j.EagerResultTransformer,
		neo4j.ExecuteQueryWithDatabase(r.cfg.DBName),
		neo4j.ExecuteQueryWithReadersRouting(),
	)
	if err != nil {
		return nil, xid.NilID(), err
	}

	users := make([]entity.User, len(result.Records))
	for i := 0; i < len(result.Records); i++ {
		users[i] = toUser(result.Records[i])
	}

	if limit != -1 && len(users) == limit {
		nextID := users[len(users)-1].ID
		return users, nextID, nil
	}

	return users, xid.NilID(), nil
}

// toUser converts a neo4j.Record to entity.User, returning an error if the record is not valid.
func toUser(record *db.Record) entity.User {
	idStr, _ := record.Get("id")
	id, _ := xid.FromString(idStr.(string))
	muted, _ := record.Get("muted")
	t, _ := record.Get("follow_start")
	followStart := t.(dbtype.Date).Time()

	return entity.User{
		ID:          id,
		Muted:       muted.(bool),
		FollowStart: followStart,
	}
}
