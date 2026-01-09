package dating_service

import (
	"context"
	"errors"
	"golang_template/internal/db"
	query "golang_template/internal/db/generated"
	"golang_template/pkg/helpers"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
)

func GetFeed(userID string) (any, *helpers.ResultError) {
	ctx := context.Background()
	uid := uuid.MustParse(userID)

	limit := 10000

	// maxViewsPerDay, canViewPhone, err := getEntitlements(ctx, uid)
	// if err != nil {
	// 	return nil, helpers.Error("unable to load subscription", err)
	// }

	periodStart := time.Now().UTC().Truncate(24 * time.Hour)
	// used, err := getDailyUsage(ctx, uid, "profile_view", periodStart)
	// if err != nil {
	// 	return nil, helpers.Error("unable to load usage", err)
	// }

	remaining := limit
	// if maxViewsPerDay != nil {
	// 	if used >= *maxViewsPerDay {
	// 		return []any{}, helpers.Error("daily limit reached", errors.New("limit reached"))
	// 	}
	// 	allowed := *maxViewsPerDay - used
	// 	if remaining > allowed {
	// 		remaining = allowed
	// 	}
	// }

	items, err := db.Query.ListFeed(ctx, query.ListFeedParams{
		UserID:       uid,
		CanViewPhone: true,
	})
	if err != nil {
		return nil, helpers.Error("unable to load feed", err)
	}

	if remaining < len(items) {
		items = items[:remaining]
	}

	if len(items) > 0 {
		periodDate := pgtype.Date{Time: periodStart, Valid: true}
		_, _ = db.Query.UpsertDailyUsage(ctx, query.UpsertDailyUsageParams{
			UserID:      uid,
			FeatureCode: "profile_view",
			PeriodStart: periodDate,
			UsedCount:   int32(len(items)),
		})
	}

	return items, nil
}

func getEntitlements(ctx context.Context, userID uuid.UUID) (*int, bool, error) {
	row, err := db.Query.GetActiveSubscriptionForUser(ctx, userID)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			free := 20
			return &free, false, nil
		}
		return nil, false, err
	}

	canViewPhone := row.CanViewPhone
	if !row.MaxProfileViewsPerDay.Valid {
		return nil, canViewPhone, nil
	}
	max := int(row.MaxProfileViewsPerDay.Int32)
	return &max, canViewPhone, nil
}

func getDailyUsage(ctx context.Context, userID uuid.UUID, featureCode string, periodStart time.Time) (int, error) {
	periodDate := pgtype.Date{Time: periodStart, Valid: true}
	row, err := db.Query.GetDailyUsage(ctx, query.GetDailyUsageParams{
		UserID:      userID,
		FeatureCode: featureCode,
		PeriodStart: periodDate,
	})
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return 0, nil
		}
		return 0, err
	}
	return int(row.UsedCount), nil
}

func GetStories(userID string) (any, *helpers.ResultError) {
	ctx := context.Background()
	uid := uuid.MustParse(userID)

	rows, err := db.Query.ListActiveStories(ctx, uid)
	if err != nil {
		return nil, helpers.Error("unable to load stories", err)
	}

	return rows, nil
}
