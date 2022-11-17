package main

import (
	"fmt"
	"log"
	"os/exec"
	"reflect"
	"time"
)
func main(){
	video :=[]string{"https://variety.com/2021/digital/news/rick-astley-never-gonna-give-you-up-1-billion-youtube-views-1235030404/","https://www.google.com/search?q=never+gonna+give+you+up&oq=never+gonna+give+you+up&aqs=chrome.0.0i271j46i433i512j0i512l5j69i60.4278j0j4&sourceid=chrome&ie=UTF-8#wptab=si:AC1wQDAmhH4WxhqkLyzXLNCgm7uMlL-cwjtGo7YoiLsJI42k1Zt3T5h9PClNLNtt0Z7jv-9rAWci7j6ZMsPAo7Gyn8gq81Qa7NPUIPM7QqFB3BzmUKtdv8f4OYoMQY2ot3_-2YFA2xD-LgK7eFAFptkcJaB8wJebkE3UETnybratKkcMGrPQ9mQqCY3MKWHKyvsTMZL57Zaw", "https://www.youtube.com/watch?v=LLFhKaqnWwk", "https://en.wikipedia.org/wiki/Never_Gonna_Give_You_Up", "https://americansongwriter.com/meaning-never-gonna-give-you-up-rick-astley-song-lyrics/", "https://open.spotify.com/track/4uLU6hMCjMI75M1A2tKUQC", "https://www.youtube.com/watch?v=pUvwleVGVmA", "https://www.youtube.com/watch?v=IAJeZwXk3cU"}
	pictures:= []string{"https://hackster.imgix.net/uploads/attachments/1317693/_Ek101jDIJo.blob?auto=compress%2Cformat&w=900&h=675&fit=min", "https://media-cldnry.s-nbcnews.com/image/upload/t_fit-1240w,f_auto,q_auto:best/streams/2012/June/120601/398909-trollface.jpg", "https://www.ebay.com/itm/163776448187", " https://media-cldnry.s-nbcnews.com/image/upload/t_fit-1240w,f_auto,q_auto:best/streams/2012/June/120601/398909-trollface.jpg", "https://www.vanityfair.com/hollywood/2019/01/fyre-festival-documentary-netflix-chris-smith-interview", "https://media-cldnry.s-nbcnews.com/image/upload/t_fit-1240w,f_auto,q_auto:best/streams/2012/June/120601/398909-trollface.jpg","https://www.foxbusiness.com/technology/elizabeth-holmes-sentencing-look-where-she-could-serve-time", "https://www.cnbc.com/2022/10/19/heres-what-the-wells-fargo-cross-selling-scandal-means-for-the-bank.html" }
	fmt.Println(len(pictures))
	for i := 0; (i+1) <= len(video); i++{
		index := (i +1)
		if i < len(video){
			time.Sleep(3 * time.Second)
			fmt.Printf("\nindex (%d)\n", index)
			browser(video[i])
			time.Sleep(2 * time.Second)
			browser(pictures[i])

		}
	}

}

func browser(url string){
var err error
	cmd := exec.Command("open", url).Start()
	if err != nil{
		log.Fatal(err)
	} 
	fmt.Println(reflect.TypeOf(cmd))
	fmt.Printf("Hey Ricky: %s", cmd)
	return
}