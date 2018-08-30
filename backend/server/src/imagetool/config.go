package imagetool

import (
	"os"
	"log"
	"encoding/json"
	"fmt"
)

type Config struct {
	Storage string
	CacheSize uint
}

var conf Config

func LoadConf(){
	r, err := os.Open("/root/pickme/backend/server/src/imagetool/config.json")
    if err != nil {
        log.Fatalln(err)
    }
    decoder := json.NewDecoder(r)
    err = decoder.Decode(&conf)
    if err != nil {
        log.Fatalln(err)
    }
    fmt.Println(conf.Storage)
}
