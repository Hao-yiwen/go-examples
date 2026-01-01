package service

import "github.com/google/wire"

// ProviderSet 是 service 层的 Provider 集合
var ProviderSet = wire.NewSet(
	NewUserService,
	NewProductService,
)
