package main

import (
	//"fmt"
	"fmt"
	"runtime"
	//"reflect"
	"log"
	"os/exec"
)
func main(){

	openbrowser("http://www.google.com")

}

func openbrowser(url string){
var err error
const r string = runtime.GOOS 
	if  err == nil{
		if r == "darwin"{
			err = exec.Command("open", url).Start()
		}else if r == " windows"{
			err = exec.Command("rundll32", 
"url.dll,FileProtocolHandler", url).Start()
		} else if r == "linux"{
			err = exec.Command("xdg-open", url).Start()
		}else {
			fmt.Println("operating system is not applicable for program")
		}
		return
	} else{
		log.Fatal(err)

		return
	}
}
