# Dating Module API

## Overview

The dating module provides user discovery, matching, and story viewing. All endpoints require authentication and use sqlc-generated result structs (no custom DTOs). Feed and stories default to a maximum of 15 items; subscription/daily usage caps may further reduce the number returned.

## Endpoints

### GET /dating/feed

- **Auth**: Required
- **Query params**: none (`limit` is not accepted)
- **Behavior**:
  - Returns up to 15 profiles (default), respecting the user’s subscription/daily view caps.
  - Subscription users with higher caps still get at most 15 due to the SQL default.
  - Free users get up to 15 or their remaining daily views, whichever is lower.
  - Each view increments daily usage.
- **Response**: sqlc-generated feed rows (ListFeed result struct) wrapped in the standard API response envelope.

### GET /dating/stories

- **Auth**: Required
- **Query params**: none (`limit` is not accepted)
- **Behavior**:
  - Returns up to 15 active stories from other users (default).
  - Stories that have expired are filtered out.
- **Response**: sqlc-generated story rows (ListActiveStories result struct) wrapped in the standard API response envelope.

### POST /dating/like

- **Auth**: Required
- **Body**:

  ```json
  { "to_user_id": "<uuid>" }
  ```

- **Behavior**:
  - Records a like from the authenticated user to `to_user_id`.
  - If the other user has already liked back, a match is created.
- **Response**: sqlc-generated result row (LikeAndMaybeMatch result struct) wrapped in the standard API response envelope.

## Models

### ListFeed row (GET /dating/feed response items)

```go
type ListFeedRow struct {
    UserID      uuid.UUID `json:"user_id"`
    Username    string    `json:"username"`
    Picture     string    `json:"picture"`
    About       string    `json:"about"`
    Location    string    `json:"location"`
    LookingFor  string    `json:"looking_for"`
    BirthDate   time.Time `json:"birth_date"`
    PhotoUrl    string    `json:"photo_url"`
    Phone       string    `json:"phone"`
}
```

### ListActiveStories row (GET /dating/stories response items)

```go
type ListActiveStoriesRow struct {
    ID        uuid.UUID `json:"id"`
    UserID    uuid.UUID `json:"user_id"`
    MediaUrl  string    `json:"media_url"`
    MediaType string    `json:"media_type"`
    ExpiresAt time.Time `json:"expires_at"`
    CreatedAt time.Time `json:"created_at"`
}
```

### LikeAndMaybeMatch row (POST /dating/like response)

```go
type LikeAndMaybeMatchRow struct {
    Matched bool `json:"matched"`
}
```

## Implementation notes

- Services use only sqlc-generated queries; no raw SQL.
- Return types are sqlc-generated structs (no custom DTOs).
- Feed respects subscription/daily usage limits by truncating results if necessary, even though the underlying query uses `limit 15`.
- Stories do not enforce daily caps; they simply return up to 15 active stories.