# Realtor Backend API

## Authentication Endpoints
- `POST /api/v1/auth/google` - Google OAuth login
- `POST /api/v1/auth/github` - GitHub OAuth login
- `POST /api/v1/auth/logout` - Logout user

## Account Endpoints
- `GET /api/v1/account/my` - Get current user account info

## Properties Endpoints

### Private (Requires Authentication)
- `POST /api/v1/properties` - Create a new property
- `GET /api/v1/properties` - Get user's properties
- `PUT /api/v1/properties/:id` - Update a property
- `DELETE /api/v1/properties/:id` - Delete a property

### Public
- `GET /api/v1/properties/:id` - Get property details

## Subscription Endpoints

### Private (Requires Authentication)
- `GET /api/v1/subscription` - Get current user subscription
- `PUT /api/v1/subscription` - Update subscription tier

### Public
- `GET /api/v1/subscription/tiers` - Get available subscription tiers

## Subscription Tiers

### Free
- 1 property maximum
- Basic property listing
- Standard support
- Price: $0/month

### Basic
- 5 properties maximum
- Priority support
- Property analytics
- Price: $9.99/month

### Professional
- 50 properties maximum
- Premium support
- Advanced analytics
- Featured listings
- Price: $29.99/month

## Property Creation Example

```json
{
  "title": "Beautiful 3-Bedroom House",
  "description": "Spacious family home with modern amenities",
  "price": 450000.00,
  "address": "123 Main Street",
  "city": "New York",
  "state": "NY",
  "zip_code": "10001",
  "bedrooms": 3,
  "bathrooms": 2,
  "square_feet": 1800,
  "property_type": "house",
  "images": ["url1", "url2"]
}
```

## Subscription Update Example

```json
{
  "tier": "basic"
}
``` 