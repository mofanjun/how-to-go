# 项目说明

## simple
一个简单的爬虫，演示如何使用go来爬取网站首页

## singleCrawler
一个单任务爬虫，演示了爬虫架构

## simpleScheduler
一个并发爬虫，演示了如何使用 `goroutine` 来提升爬虫效率

## queueScheduler
在 scheduler版本的基础上，演示如何增加对`goroutione`的控制力

## finalCrawler
fix `Scheduler Interface`，让 simpleScheduler 和 queueScheduler 都能用。