package nethgateapi

import (
	"time"

	"github.com/google/uuid"
)

////////////////////////////////

//

////////////////////////////////

// NethGateService defines the interface for the NethGate API operations.
type NethGateServiceInterface interface {
	CreateUser(login string, password string) (User, error)
	SearchUser(login string) (User, error)
	UpdateUser(userID uuid.UUID, updates UserUpdates) (User, error)
	DeleteUser(userID uuid.UUID) error

	CreateProduct(name string, isFreeTier bool, rateLimit int, hostNameRegex string) (Product, error)
	GetProduct(productID uuid.UUID) (Product, error)
	UpdateProduct(productID uuid.UUID, updates ProductUpdates) (Product, error)

	CreateSubscription(userID uuid.UUID, productID uuid.UUID, isHardStop bool) (Subscription, error)
	GetSubscriptionsForUser(userID uuid.UUID) ([]Subscription, error)
	UpdateSubscription(subscriptionID uuid.UUID, updates SubscriptionUpdates) (Subscription, error)
	DeleteSubscription(subscriptionID uuid.UUID) error

	CreateAPIKey(subscriptionID uuid.UUID) (APIKey, error)
	UpdateAPIKey(apiKeyID uuid.UUID, isActive bool) (APIKey, error)
	DeleteAPIKey(apiKeyID uuid.UUID) error

	UpdateUsageStats(apiKeyID uuid.UUID, date time.Time, successCount int, failureCount int) (DailyUsage, error)
	GetUsageStats(apiKeyID uuid.UUID, startDate time.Time, endDate time.Time) ([]DailyUsage, error)
}

// UserUpdates defines the fields that can be updated for a user.
type UserUpdates struct {
	Login    string
	Password string
}

// ProductUpdates defines the fields that can be updated for a product.
type ProductUpdates struct {
	Name          string
	IsFreeTier    bool
	RateLimit     int
	HostNameRegex string
}

// SubscriptionUpdates defines the fields that can be updated for a subscription.
type SubscriptionUpdates struct {
	IsHardStop bool
}

// User represents a user in the system.
type User struct {
	UserID      uuid.UUID `gorm:"type:uuid;primaryKey"`
	Login       string
	CreatedDate time.Time
	UpdatedDate time.Time
}

// Product represents a product offered in the system.
type Product struct {
	ProductID     uuid.UUID `gorm:"type:uuid;primaryKey"`
	Name          string
	IsFreeTier    bool
	RateLimit     int
	HostNameRegex string
	Quotas        []Quota
}

// Quota represents usage limits on a product.
type Quota struct {
	QuotaID   uuid.UUID `gorm:"type:uuid;primaryKey"`
	ProductID uuid.UUID `gorm:"type:uuid;foreignKey"`
	Period    time.Duration
	Limit     int
}

// Subscription represents a user's subscription to a product.
type Subscription struct {
	SubscriptionID uuid.UUID `gorm:"type:uuid;primaryKey"`
	UserID         uuid.UUID `gorm:"type:uuid;foreignKey"`
	ProductID      uuid.UUID `gorm:"type:uuid;foreignKey"`
	IsHardStop     bool
	CreatedDate    time.Time
	UpdatedDate    time.Time
	APIKeys        []APIKey
}

// APIKey represents an API key associated with a subscription.
type APIKey struct {
	APIKeyID       uuid.UUID `gorm:"type:uuid;primaryKey"`
	SubscriptionID uuid.UUID `gorm:"type:uuid;foreignKey"`
	Key            string
	IsActive       bool
	DailyUsages    []DailyUsage
}

// DailyUsage represents usage stats for an API key.
type DailyUsage struct {
	DailyUsageID     uuid.UUID `gorm:"type:uuid;primaryKey"`
	APIKeyID         uuid.UUID `gorm:"type:uuid;foreignKey"`
	UsageDate        time.Time
	SuccRequestCount int
	FailRequestCount int
	UpdatedDate      time.Time
}
