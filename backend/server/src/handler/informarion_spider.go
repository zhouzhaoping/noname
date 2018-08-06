package handler

import (
	"time"
	"os/exec"
	"fmt"
)

func Refresh(star_id int) {
	lasttime := findLastTime(star_id)
	if lasttime == nil{
		cmd := exec.Command("ls")
		buf, err := cmd.Output()
		fmt.Printf("%s\n%s",buf,err)
	}else{

	}
}
func findLastTime(star_id int) (lasttime *time.Time){

	return nil
}

