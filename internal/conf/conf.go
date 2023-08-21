package conf

import "github.com/google/wire"

var ProviderSet = wire.NewSet(GetData, GetJwtConf, GetOSSConf, GetTiktokConf)
