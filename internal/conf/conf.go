package conf

import (
	"github.com/google/wire"
	"path"
	"runtime"
)

var ProviderSet = wire.NewSet(GetData, GetJwtConf, GetOSSConf, GetTiktokConf)

func getIniPath() string {
	_, filename, _, _ := runtime.Caller(1)
	str := path.Dir(filename)
	return getAncientDirectory(str, 2) + "/configs/conf.ini"
}
func getAncientDirectory(path string, level int) string {
	var i = len(path) - 1
	for ; i >= 0 && level > 0; i-- {
		if path[i] == '/' {
			level--
		}
	}
	return path[0 : i+1]
}
