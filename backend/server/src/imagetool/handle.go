package imagetool

import (
	"fmt"
	"log"
	"os"
	"io"
	"net/http"
	"github.com/kataras/iris"
	"path"
)

func UploadHandler(ctx iris.Context) {

	//上传参数为uploadfile
	ctx.Request().ParseMultipartForm(32 << 20)
	file, _, err := ctx.Request().FormFile("uploadfile")
	if err != nil {
		log.Println(err)
		ctx.JSON(iris.Map{
			"state": "Error:Upload Error. - request",
		})
		return
	}
	defer file.Close()

	//检测文件类型
	buff := make([]byte, 512)
	_, err = file.Read(buff)
	if err != nil {
		log.Println(err, "readfile")
		fmt.Println("readfile")
		ctx.JSON(iris.Map{
			"state": "Error:Upload Error - readfile.",
		})
		return
	}
	filetype := http.DetectContentType(buff)
	fmt.Println(filetype)
	var suffix string
	if filetype =="image/jpeg"{
		suffix = "jpg"
	} else if filetype=="image/png" {
		suffix = "png"
	} else {
		ctx.JSON(iris.Map{
			"state": "Error:Not JPEG OR PNG",
		})
		return
	}

	//随机生成一个不存在的fileid
	var imgid string
	for{
		imgid=MakeImageID()
		if !FileExist(ImageID2Path(imgid,suffix)){
			break
		}
	}

	//回绕文件指针
	log.Println(filetype)
	if  _, err = file.Seek(0, 0); err!=nil{
		log.Println(err)
	}
	//提前创建整棵存储树
	if err=BuildTree(imgid); err!=nil{
		log.Println(err)
	}
	//log.Println(ImageID2Path(imgid))
	f, err := os.OpenFile(ImageID2Path(imgid,suffix), os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		log.Println(err)
		ctx.JSON(iris.Map{
			"state": "Error:Save Error.",
		})
		return
	}
	defer f.Close()
	io.Copy(f, file)

	ctx.JSON(iris.Map{
		"state": "success",
		"data": iris.Map{
			"imgid":imgid + "." + suffix,
		},
	})
}

func DownloadHandler(ctx iris.Context) {
	//vars := mux.Vars(r)
	//imageid := vars["imgid"]

	imgid := ctx.Params().Get("imgid")
	suffix := path.Ext(imgid)[1:]
	imageid := imgid[:len(imgid)-len(suffix)-1]

	fmt.Println(imageid)
	fmt.Println(suffix)
	if len([]rune(imageid)) != 16 {
		ctx.JSON(iris.Map{
			"state": "Error:ImageID incorrect.",
		})
		return
	}
	imgpath := ImageID2Path(imageid,suffix)
	fmt.Println(imgpath)
	if !FileExist(imgpath) {
		ctx.JSON(iris.Map{
			"state": "Error:Image Not Found.",
		})
		return
	}
	http.ServeFile(ctx.ResponseWriter(), ctx.Request(), imgpath)
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<html><body><center><h1>It Works!</h1></center><hr><center>Quick Image Server</center></body></html>"))
}
