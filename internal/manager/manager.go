package manager

import (
	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(NewFavoriteManager, NewVideoManager, NewRelationManager, NewUserManager, NewCommentManager)
