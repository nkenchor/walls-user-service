package ports

import (
	"context"
	"walls-user-service/internal/core/domain/entity"
)

type UserRepository interface {
	// USER MANAGEMENT
	//--------------------------------------------------------------------------

	// CRUD Operations on User
	CreateUser(ctx context.Context, user entity.User) (interface{}, error)
	UpdateUser(ctx context.Context, user_reference string, user entity.User) (interface{}, error)

	// Retrieval of Users by various identifiers
	GetUserByReference(ctx context.Context, user_reference string) (interface{}, error)
	GetUserByPhone(ctx context.Context, phone string) (interface{}, error)
	GetUserByWallsTag(ctx context.Context, wallsTag string) (interface{}, error)
	GetUserByWallsBadgeReference(ctx context.Context, wallsBadgeReference string) (interface{}, error)
	GetUserByDevice(ctx context.Context, device entity.Device) (interface{}, error)
	GetUserDefaultWallsBadge(ctx context.Context, userReference string) (interface{}, error)
}
