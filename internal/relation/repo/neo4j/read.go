package neo4j

import (
	"context"
	"errors"
	"fmt"

	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j/db"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j/dbtype"

	"github.com/Karzoug/meower-relation-service/internal/relation/entity"
)

var errBadLimit = errors.New("limit must be positive or -1")

func (r repo) ListFollowers(ctx context.Context,
	_, targetUserID, afterID string, limit int,
) ([]entity.User, *string, error) {
	const q = `MATCH (u:User{id: $tuser})<-[f:FOLLOWS]-(followers:User %s)
OPTIONAL MATCH (u)<-[h:HIDES]-(followers)
WITH followers.id as id, f.start as follow_start, h IS NOT NULL as hidden
ORDER BY id
RETURN id, follow_start, hidden
%s`

	if limit == 0 {
		return nil, nil, errBadLimit
	}

	pms := map[string]any{
		"tuser": targetUserID,
	}

	var limitQ, afterQ string
	if afterID != "" {
		afterQ = "WHERE followers.id>$after_id"
		pms["after_id"] = afterID
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
		return nil, nil, err
	}

	users := make([]entity.User, len(result.Records))
	for i := 0; i < len(result.Records); i++ {
		users[i] = toUser(result.Records[i])
	}

	if limit != -1 && len(users) == limit {
		nextID := users[len(users)-1].ID
		return users, &nextID, nil
	}

	return users, nil, nil
}

func (r repo) ListFollowings(ctx context.Context, _, targetUserID, afterID string, limit int) ([]entity.User, *string, error) {
	const q = `
MATCH (u:User{id: $tuser})-[f:FOLLOWS]->(followings:User %s)
OPTIONAL MATCH (u)-[h:HIDES]->(followings)
WITH followings.id as id, f.start as follow_start, h IS NOT NULL as hidden
ORDER BY id
RETURN id, follow_start, hidden
%s`

	pms := map[string]any{
		"tuser": targetUserID,
	}

	var limitQ, afterQ string
	if afterID != "" {
		afterQ = "WHERE followings.id>$after_id"
		pms["after_id"] = afterID
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
		return nil, nil, err
	}

	users := make([]entity.User, len(result.Records))
	for i := 0; i < len(result.Records); i++ {
		users[i] = toUser(result.Records[i])
	}

	if limit != -1 && len(users) == limit {
		nextID := users[len(users)-1].ID
		return users, &nextID, nil
	}

	return users, nil, nil
}

// toUser converts a neo4j.Record to entity.User, returning an error if the record is not valid.
func toUser(record *db.Record) entity.User {
	id, _ := record.Get("id")
	hidden, _ := record.Get("hidden")
	t, _ := record.Get("follow_start")
	followStart := t.(dbtype.Date).Time()

	return entity.User{
		ID:          id.(string),
		Hidden:      hidden.(bool),
		FollowStart: followStart,
	}
}
