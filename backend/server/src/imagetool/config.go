package imagetool

type Config struct {
	Storage string
}

var conf Config

func LoadConf(){
	/*r, err := os.Open("../backend/server/src/imagetool/config.json")
    if err != nil {
        log.Fatalln(err)
    }
    decoder := json.NewDecoder(r)
    err = decoder.Decode(&conf)
    if err != nil {
        log.Fatalln(err)
    }*/
    conf.Storage = "D:"
}
