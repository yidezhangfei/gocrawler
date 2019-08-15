package main

import (
	"encoding/xml"
	"fmt"
	"github.com/gocolly/colly"
	"github.com/yidezhangfei/gocrawler/jrj_stock_toutiao"
	"io/ioutil"
	"log"
	"os"
)

var gMainUrl = ""

type SConfig struct {
	XMLName  xml.Name `xml:"config"`
	MainUrl  string   `xml:"mainUrl"`
	MongoUri string   `xml:"mongodbUri"`
}

func initConfig() bool {
	configFile, err := os.OpenFile("conf.xml", os.O_RDONLY, os.ModePerm)
	if err != nil {
		fmt.Printf("err: %v \n", err)
		return false
	}
	defer configFile.Close()
	data, err := ioutil.ReadAll(configFile)
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return false
	}
	config := SConfig{}
	err = xml.Unmarshal(data, &config)
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return false
	}
	gMainUrl = config.MainUrl
	mongoUri := config.MongoUri
	jrj_stock_toutiao.Init(mongoUri)
	return true
}

func main() {
	init := initConfig()
	if init != true {
		log.Fatal("init config err")
		return
	}
	crawler := colly.NewCollector(
		colly.MaxDepth(2),
		colly.Async(true))

	crawler.Limit(&colly.LimitRule{DomainGlob: "*", Parallelism: 2})

	// On every  a element which has href attribute call back
	crawler.OnHTML(jrj_stock_toutiao.ListSelector, jrj_stock_toutiao.FindListCallback)

	crawler.OnHTML(jrj_stock_toutiao.ContentSelector, jrj_stock_toutiao.GetArticleCallback)

	// Starting scrapping
	crawler.Visit(gMainUrl)
	crawler.Wait()
}
