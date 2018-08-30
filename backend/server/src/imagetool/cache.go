package imagetool

import "time"

type memo struct{
	img []byte
	hitTime time.Time
}
var path2byte map[string][]byte
func CacheInit() {
	path2byte = make(map[string][]byte,conf.CacheSize)
}

func GetCache(imgpath string) (bool, []byte) {


	if _, ok := path2byte[imgpath]; ok {
		//存在
		return true, path2byte[imgpath]
	}
	return false, nil
}

func AddCache(imgpath string) bool {

	return true
}
