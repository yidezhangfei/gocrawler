package main

import (
	"github.com/gocolly/colly"
	"github.com/yidezhangfei/gocrawler/jrj_stock_toutiao"
)

var gMainUrl = "http://www.jrj.com.cn"

func main() {
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
