package handler

import "github.com/google/wire"

// ProviderSet 是 handler 层的 Provider 集合
var ProviderSet = wire.NewSet(
	NewUserHandler,
	NewProductHandler,
)
