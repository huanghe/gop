package main

import "io/ioutil"

func main() {

	//构造出真实的网站url集合
	list := []string
	//生成total行日志内容，院子上面的集合
	for i = 0; i < *total; i++ {
		ioutil.WriteFile(*filePath, []byte(logStr), 0644)
	}
}
