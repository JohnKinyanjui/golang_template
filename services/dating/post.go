package dating_service

import (
	"context"
	"golang_template/internal/db"
	query "golang_template/internal/db/generated"
	"golang_template/pkg/helpers"

	"github.com/google/uuid"
)

func LikeUser(fromUserID string, toUserID string) (any, *helpers.ResultError) {
	ctx := context.Background()
	fromID := uuid.MustParse(fromUserID)
	toID := uuid.MustParse(toUserID)

	tx, err := db.PgConn.Begin(ctx)
	if err != nil {
		return nil, helpers.Error("unable to like user", err)
	}
	defer func() {
		_ = tx.Rollback(ctx)
	}()

	qtx := db.Query.WithTx(tx)

	row, err := qtx.LikeAndMaybeMatch(ctx, query.LikeAndMaybeMatchParams{
		FromUserID: fromID,
		ToUserID:   toID,
	})
	if err != nil {
		return nil, helpers.Error("unable to like user", err)
	}

	_, err = qtx.UpsertDatingInteraction(ctx, query.UpsertDatingInteractionParams{
		UserID:       fromID,
		TargetUserID: toID,
		Action:       query.DatingInteractionActionLike,
	})
	if err != nil {
		return nil, helpers.Error("unable to like user", err)
	}

	if err := tx.Commit(ctx); err != nil {
		return nil, helpers.Error("unable to like user", err)
	}

	return row, nil
}

func RejectUser(fromUserID string, toUserID string) (any, *helpers.ResultError) {
	ctx := context.Background()
	fromID := uuid.MustParse(fromUserID)
	toID := uuid.MustParse(toUserID)

	row, err := db.Query.UpsertDatingInteraction(ctx, query.UpsertDatingInteractionParams{
		UserID:       fromID,
		TargetUserID: toID,
		Action:       query.DatingInteractionActionNope,
	})
	if err != nil {
		return nil, helpers.Error("unable to reject user", err)
	}

	return row, nil
}
