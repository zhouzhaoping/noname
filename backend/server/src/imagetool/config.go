package imagetool

type Config struct {
	Storage string
}

var conf Config

func LoadConf(){
	conf.Storage = "D:/下载"
}
