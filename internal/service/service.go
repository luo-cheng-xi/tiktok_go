package service

import "github.com/google/wire"

var ProviderSet = wire.NewSet(NewUserService, NewVideoService, NewRelationService, NewFavoriteService, NewCommentService)
