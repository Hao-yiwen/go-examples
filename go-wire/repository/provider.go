package repository

import "github.com/google/wire"

// ProviderSet 是 repository 层的 Provider 集合
var ProviderSet = wire.NewSet(
	NewUserRepository,
	NewProductRepository,
)
