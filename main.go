package main

import (
	"fmt"
	"github.com/gocolly/colly"
	"github.com/robfig/cron"
	"ji_sign/util"
	"log"
	"time"
)

func init()  {
	//获取执行文件路径
	util.GetExecutePath()
	//加载配置文件
	util.LoadConfig()
	util.OpenLogFile()

}

func main() {
	signall()
	c := cron.New()
	c.AddFunc("0 * * * * ?", func() {
		timeNow:=time.Now()
		timeNowStr:=timeNow.Format("2006-01-02 15:04:05")
		fmt.Print(timeNowStr +": test cron ,every hour \n")
	})
	c.AddFunc("0 0 9 * * ?", func() {
		fmt.Print("每天9:00签到")
		signall()
	})
	c.Start()
	select {}
}

func signall(){
	if (len(util.AppConfig.GetString("email1") != 0){
		sign(util.AppConfig.GetString("email1"), util.AppConfig.GetString("passwd1"))
	}
	if (len(util.AppConfig.GetString("email2") != 0){
		sign(util.AppConfig.GetString("email2"), util.AppConfig.GetString("passwd2"))
	}
	if (len(util.AppConfig.GetString("email3") != 0){
		sign(util.AppConfig.GetString("email3"), util.AppConfig.GetString("passwd3"))
	}
}

//登录并签到
func sign(a_email, a_passwd string)  {
	// create a new collector
	v_url := util.AppConfig.GetString("url")
	c := colly.NewCollector(
		colly.AllowedDomains("j05.space"),
	)

	// authenticate
	err := c.Post(v_url + "/signin", map[string]string{"email": a_email, "passwd": a_passwd})
	if err != nil {
		log.Fatal(err)
		util.Log(err.Error())
	}

	c.OnResponse(func(r *colly.Response) {
		log.Println("response revice", string(r.Body))
		util.Log("response revice :"+string(r.Body))

	})
	c.Visit(v_url + "/xiaoma/get_user")
	//签到
	err = c.Post(v_url + "/user/checkin", map[string]string{})
	if err != nil {
		log.Fatal(err)
		util.Log(err.Error())
	}
}
