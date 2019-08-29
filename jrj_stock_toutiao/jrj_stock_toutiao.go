package jrj_stock_toutiao

import (
	"context"
	"fmt"
	"github.com/gocolly/colly"
	"github.com/yidezhangfei/gocrawler/mongodb"
	"github.com/yidezhangfei/gocrawler/util"
	"go.mongodb.org/mongo-driver/bson"
	"log"
	"os"
	"time"
)

var ListSelector = "div[class='mo04 mt15']>div[class='list-s1 mt15']"
var ContentSelector = "div[class='titmain']"
var MongoUri = "mongodb://localhost:27017"

type article struct {
	title   string
	content string
}

func Init(mongoUri string) {
	MongoUri = mongoUri
}

func FindListCallback(e *colly.HTMLElement) {
	var href = "a[href]"
	e.ForEach(href, func(i int, e *colly.HTMLElement) {
		//text := e.Text
		link := e.Attr("href")
		//fmt.Printf("text: %v;\n link: %v\n\n", text, link)
		e.Request.Visit(link)
	})
}

func GetArticleCallback(e *colly.HTMLElement) {
	title := e.ChildText("h1")
	text := e.ChildText("div[class='texttit_m1']")
	var item = article{title: title, content: text}
	//fmt.Printf("item: %v", item)
	//saveToFile("toutiao.txt", item)
	saveToDB(item)
}

func saveToFile(filename string, item article) {
	file, err := os.OpenFile(filename, os.O_CREATE|os.O_APPEND, os.ModeAppend)
	defer file.Close()
	if err == nil {
		file.WriteString(item.title)
		file.WriteString("\n")
		file.WriteString(item.content)
		file.WriteString("\n\n")
	} else {
		log.Printf("save to file err, err: %v", err)
	}
}

var initDB = false

func saveToDB(item article) {
	if initDB != true {
		err := mongodb.MongoInit(MongoUri)
		if err == nil {
			initDB = true
		}
	}
	if initDB == true {
		document, md5String := itemToDocument(item)
		fmt.Printf("title: %v", item.title)
		if document != nil {
			filter := bson.M{"md5": md5String}
			cursor := mongodb.Collection("stock", "jrj_stock_toutiao").FindOne(context.TODO(), filter)
			var result = bson.M{}
			err := cursor.Decode(&result)
			if err == nil && result == nil {
				_, err := mongodb.Collection("stock", "jrj_stock_toutiao").InsertOne(context.TODO(), document)
				if err != nil {
					log.Fatalf("insert err: %v \n", err)
				}
			}
		}
	}
}

func itemToDocument(item article) (doc bson.M, md5String string) {
	now := time.Now()
	year, month, day := now.Date()
	var currentDate = fmt.Sprintf("%d-%d-%d", year, month, day)
	md5string := util.Md5String(item.title)
	if md5string == "" {
		log.Fatal("md5 is null")
		return nil, ""
	}
	document := bson.M{"date": currentDate, "title": item.title, "content": item.content, "md5": md5string}
	return document, md5string
}
