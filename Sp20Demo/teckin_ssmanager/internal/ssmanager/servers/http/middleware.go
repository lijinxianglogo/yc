package http

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"net/http/httputil"
	"reflect"
	"regexp"
	"runtime"
	"strings"
	"teckin_ssmanager/pkg/utils/sort"
	"time"
)

var skipParamsForSQLInjectMap = map[string]int{
	"openid":           1,
}

var skipSQLInjectUrl = []string{
}

//跨域配置中间件
func crossDomain(c *gin.Context){
	method := c.Request.Method
	c.Header("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token, Authorization, Token")
	c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, PATCH, DELETE")
	c.Header("Access-Control-Expose-Headers", "Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
	c.Header("Access-Control-Allow-Credentials", "true")
	c.Header("Access-Control-Allow-Origin", "*")
	// 放行所有OPTIONS方法，因为有的模板是要请求两次的
	if method == "OPTIONS" {
		c.AbortWithStatus(http.StatusNoContent)
	}
	// 处理请求
	c.Next()
}

//当遇到panic直接拦截返回500
func recoverHandler(c *gin.Context) {
	defer func() {
		if err := recover(); err != nil {
			const size = 64 << 10
			buf := make([]byte, size)
			buf = buf[:runtime.Stack(buf, false)]
			httprequest, _ := httputil.DumpRequest(c.Request, false)
			pnc := fmt.Sprintf("[Recovery] %s panic recovered:\n%s\n%s\n%s", time.Now().Format("2006-01-02 15:04:05"), string(httprequest), err, buf)
			fmt.Print(pnc)
			c.AbortWithStatus(500)
		}
	}()
	c.Next()
}

//检查sql注入中间件
func checkSQLInject(c *gin.Context){
	for _,url := range skipSQLInjectUrl{
		if strings.Contains(c.Request.URL.Path, url){
			c.Next()
			return
		}
	}

	content,_ := httputil.DumpRequest(c.Request,true)
	if bytes.Index(content,[]byte("application/json")) != -1{
		content = content[bytes.Index(content, []byte("User-Agent")):]
		index := bytes.Index(content, []byte("{"))
		if index != -1 {
			content = content[index:]
		} else {
			content = nil
		}

		if len(content) != 0 { //check request json parameter
			fmt.Println(string(content))
			m := make(map[string]interface{})
			err := json.Unmarshal(content, &m)
			if err != nil {
				err = fmt.Errorf("parse request body failed, err: %s", err.Error())
				panic(err)
			}
			for k, v := range m {
				if skipParamsForSQLInjectMap[k] == 1 {
					continue
				}
				if reflect.TypeOf(v).String() == "string" {
					if filteredSQLInject(v.(string)) {
						err := fmt.Errorf("sql注入攻击 %s", v)
						panic(err)
					}
				}
			}
		}
	}
	if c.Request.Form == nil { //check method post and get
		c.Request.ParseMultipartForm(32 << 20)
	}
	for k, arr := range c.Request.Form {
		if c.Request.Method != http.MethodGet {
			fmt.Printf("%s=%v&", k, arr)
		}
		if skipParamsForSQLInjectMap[k] == 1 {
			continue
		}
		for _, v := range arr {
			if filteredSQLInject(v) {
				err := fmt.Errorf("sql注入攻击 %s", v)
				panic(err)
			}

		}
	}
	c.Next()
}

func filteredSQLInject(tomatch_str string) bool{
	str := `(?:')|(?:--)|(/\\*(?:.|[\\n\\r])*?\\*/)|(\b(select|update|and|or|delete|insert|trancate|char|chr|into|substr|ascii|declare|exec|count|master|into|drop|execute)\b)`

	re, err := regexp.Compile(str)
	if err != nil {
		fmt.Println(err.Error())
		return true
	}

	return re.MatchString(tomatch_str)
}
//```
//参数名称				|类型		|出现要求	|描述
//:----				|:---		|:------	|:---
//token				|string		|R			|用户登录后token，没有登录则为空字符串
//vendor              |String     |R          |APP类型，teckin,teckinBeta
//version				|string		|R			|版本号，如1.1.36
//apptype			    |string	    |R			|app系统平台，iOS、 Android
//TimeStamp			|bigint		|R			|当前UNIX时间戳（秒级）
//nonce				|string		|R			|随机流水号（防止重复提交）
//i18n			    |string     |O			|多语言，如zh_CN，简体中文
//sign			    |string     |R			|请求签名，全大写
//```

//部分接口必须携带token，需单独写中间件校验
func headerTokenVerify(c *gin.Context){
	token := c.GetHeader("token")
	if token == ""{
		c.Abort()
		responseFail(c, 400, "token为空，校验失败")
	}
	c.Next()
}

//签名生成的通用步骤如下：
//
//第一步，设所有发送或者接收到的数据为集合M，将集合M内非空参数值的参数按照参数名ASCII码从小到大排序（字典序），
//使用URL键值对的格式（即key1=value1&key2=value2…）拼接成字符串params。
//
//特别注意以下重要规则：
//
//- 参数名ASCII码从小到大排序（字典序）；
//- 如果参数的值为空不参与签名；
//- 参数名区分大小写；
//- 传送的sign参数不参与签名，将生成的签名与该sign值作校验。
//- 接口可能增加字段，验证签名时必须支持增加的扩展字段
//第二步，按照顺序key+TimeStamp+nonce+params进行MD5运算，再将得到的字符串所有字符转换为大写，得到sign值signValue。

//头校验中间件
func headerVerify(c *gin.Context){
	var header Header
	if err := c.BindHeader(&header);err != nil{
		c.Abort()
		responseFail(c, 400, "头信息解析错误，或缺少必须字段")
	}
	if err := checkHeader(header);err != nil{
		c.Abort()
		responseFail(c, 400, err.Error())
	}

	c.Next()
}

//校验头信息
func checkHeader(header Header) error {
	if header.Token != ""{ //若token不为空，代表为登录态，校验key需要查找
		//TODO 登录态方法未添加
	}
	headerParamsInit(header)
	fmt.Println(fmt.Sprintf("%v", header))

	return nil
}


//初始化用于校验的params字段
func headerParamsInit(header Header){
	headerParamList := headerListInit(header)
	headerParamList, _ = sort.StringListPositiveOrder(headerParamList)
	strings.Join(headerParamList, "&")
}

//初始化已有参数列表，用于排序，之后只需判断token及i18n字段即可
func headerListInit(header Header) []string {
	headerParamList := []string{"vendor", "version", "apptype", "TimeStamp", "nonce"}
	headerParamList = append(headerParamList, fmt.Sprintf("vendor=%s", header.Vendor))
	headerParamList = append(headerParamList, fmt.Sprintf("version=%s", header.Version))
	headerParamList = append(headerParamList, fmt.Sprintf("apptype=%s", header.Apptype))
	headerParamList = append(headerParamList, fmt.Sprintf("TimeStamp=%s", header.TimeStamp))
	headerParamList = append(headerParamList, fmt.Sprintf("nonce=%s", header.Nonce))
	if header.Token != ""{
		headerParamList = append(headerParamList, fmt.Sprintf("token=%s", header.Token))
	}
	if header.I18n != ""{
		headerParamList = append(headerParamList, fmt.Sprintf("i18n=%s", header.I18n))
	}

	return headerParamList
}

func signGenerate(str string){

}