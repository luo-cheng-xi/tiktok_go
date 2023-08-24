package util

import "github.com/google/wire"

var ProviderSet = wire.NewSet(NewJwtUtil, NewOssUtil, NewVoUtil)
