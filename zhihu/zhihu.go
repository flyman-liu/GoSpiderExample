/*
Copyright 2017 by GoSpider author.
Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at
    http://www.apache.org/licenses/LICENSE-2.0
Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License
*/

// 知乎系列爬虫
package zhihu

import (
// 第一步：引入库
)
import (
	"github.com/hunterhug/GoSpider/spider"
	"github.com/hunterhug/GoSpider/util"
	"strings"
)

var (
	// 一只小爬虫
	Baba  *spider.Spider
	Debug = "info"
)

func init() {
	// 第一步：可选设置全局
	spider.SetLogLevel(Debug)   // 设置全局爬虫日志，可不设置，设置debug可打印出http请求轨迹
	spider.SetGlobalTimeout(10) // 爬虫超时时间，可不设置，默认超长时间

	// 第二步： 新建一个爬虫对象，nil表示不使用代理IP，可选代理
	spiders, err := spider.NewSpider(nil) // 也可以使用boss.New(nil),同名函数

	if err != nil {
		panic(err)
	}
	Baba = spiders
}

func Login(email, password string) ([]byte, error) {
	if strings.Contains(email, "@") {
		Baba.SetUrl("https://www.zhihu.com/login/email").SetRefer("https://www.zhihu.com/").SetUa(spider.RandomUa())
		Baba.SetFormParm("email", email).SetFormParm("password", password)
	} else {
		Baba.SetUrl("https://www.zhihu.com/login/phone_num").SetRefer("https://www.zhihu.com/").SetUa(spider.RandomUa())
		Baba.SetFormParm("phone_num", email).SetFormParm("password", password)
	}
	body, err := Baba.Post()

	// 清除Post的数据，方便下次使用
	Baba.Clear()

	if err != nil {
		return []byte("网路错误..."), err
	}
	return util.JsonBack(body)
}