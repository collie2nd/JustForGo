package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type RawStruct struct {
	Id   string
	Text string
}

type RawQuery struct {
	Query string `json:"query" uri:"query" form:"query"`
}

type PostFormStruct struct {
	GetPostForm string `json:"getPostForm" uri:"getPostForm" form:"getPostForm"`
}

type PostxwwwFormStruct struct {
	PostFormString string `json:"PostFormString" uri:"PostFormString" form:"PostFormString"`
	PostFormInt    int    `json:"PostFormInt" uri:"PostFormInt" form:"PostFormInt"`
}

type User struct {
	Name string `uri:"name"`
	Age  int    `uri:"age"`
}

var HttpContentType = map[string]string{
	".avi":  "video/avi",
	".mp3":  "audio/mp3",
	".mp4":  "video/mp4",
	".wmv":  "video/x-ms-wmv",
	".asf":  "video/x-ms-asf",
	".rm":   "application/vnd.rn-realmedia",
	".rmvb": "application/vnd.rn-realmedia-vbr",
	".mov":  "video/quicktime",
	".m4v":  "video/mp4",
	".flv":  "video/x-flv",
	".jpg":  "image/jpeg",
	".png":  "image/png",
}

func main() {
	e := gin.Default()

	// GET
	// query类型 请求参数
	// http://localhost:8080/paramtype/query/this is param?query=this is query&defaultQuery=this is defaultQuery&getQuery=this is getQuery&queryArray=this&queryArray=is&queryArray=queryArray&getQueryArray=this&getQueryArray=is&getQueryArray=getQueryArray&queryMap[1]=this&queryMap[2]=is&queryMap[3]=queryMap&getQueryMap[1]=this&getQueryMap[2]=is&getQueryMap[3]=getQueryMap

	e.GET("/paramtype/query/:param", func(c *gin.Context) {
		//需定义合规结构体
		rawQuery := &RawQuery{}
		c.ShouldBindQuery(rawQuery)
		fmt.Printf("%+v \n", rawQuery)

		param := c.Param("param") // Param函数取路由中的参数“:param”
		fmt.Println(param)

		query := c.Query("query")
		fmt.Println(query)

		defaultQuery := c.DefaultQuery("defaultQuery", "no") // Default*类函数，当该param不存在时，会自动给一个默认值
		fmt.Println(defaultQuery)

		getQuery, res := c.GetQuery("getQuery") // Get*类函数，会多返回一个指向该param是否存在的bool类型
		fmt.Println(getQuery, res)

		// query 中的参数均为字符串类型，所以 QueryArray 为 []string, QueryMap 为 map[string]string

		queryArray := c.QueryArray("queryArray")
		fmt.Println(queryArray)

		getQueryArray, res := c.GetQueryArray("getQueryArray") // Get*类函数，会多返回一个指向该param是否存在的bool类型
		fmt.Println(getQueryArray, res)

		queryMap := c.QueryMap("queryMap")
		fmt.Println(queryMap)

		getQueryMap, res := c.GetQueryMap("getQueryMap") // Get*类函数，会多返回一个指向该param是否存在的bool类型
		fmt.Println(getQueryMap, res)

		c.String(http.StatusOK, "ok")
	})

	// POST
	// form-data
	// form-data的获取正规正矩。
	// form-data就是http请求中的multipart/form-data，它会将表单的数据处理为一条消息，以标签为单元，用分隔符分开。
	// 既可以上传键值对，也可以上传文件。当上传的字段是文件时，会有Content-Type来说明文件类型；content-disposition，用来说明字段的一些信息；
	// 如果是文件，可以通过FormFile或者MultipartForm获取文件内容，FormFile获取一个，MultipartForm获取多个。使用SaveUploadedFile存储文件。
	e.POST("/paramtype/postformdata", func(c *gin.Context) {
		postFormStruct := &PostFormStruct{}
		c.ShouldBind(postFormStruct)
		fmt.Printf("%+v \n", postFormStruct)

		// form-data 表单的 body 类型会把所有类型都识别为 string ,如果要传递其他类型，可以用映射，或者再自行转换
		postForm := c.PostForm("postForm")
		fmt.Println(postForm)

		defaultPostForm := c.DefaultPostForm("defaultPostForm", "no")
		fmt.Println(defaultPostForm)

		getPostForm, res := c.GetPostForm("getPostForm")
		fmt.Println(getPostForm, res)

		postFormArray := c.PostFormArray("postFormArray")
		fmt.Println(postFormArray)

		getPostFormArray, res := c.GetPostFormArray("getPostFormArray")
		fmt.Println(getPostFormArray, res)

		postFormMap := c.PostFormMap("postFormMap")
		fmt.Println(postFormMap)

		getPostFormMap, res := c.GetPostFormMap("getPostFormMap")
		fmt.Println(getPostFormMap, res)

		postFormFile, _ := c.FormFile("file") // 同一个key，甚至可以挂载不同类型的数据
		fmt.Println(postFormFile.Filename)

		fileStr := c.PostForm("file") // 同一个key，甚至可以挂载不同类型的数据
		fmt.Println(fileStr)

		// 此处可以看清，表单类型的本质，就是两种数据类型
		// 	Value map[string][]string
		//	File  map[string][]*FileHeader
		multipartForms, _ := c.MultipartForm()
		fmt.Println(multipartForms)
		// 可以通过 multipartForms 获取到文件列表
		for key, files := range multipartForms.File {
			fmt.Println(key)
			for _, file := range files {
				fmt.Println(file.Filename)
			}
		}

		c.String(http.StatusOK, "ok")
	})

	// POST
	// x-www-form-urlencoded
	// x-www-form-urlencoded是application/x-www-from-urlencoded，将表单内的数据转换为键值对，&分隔。
	// 当form表单的action为get时，浏览器用x-www-form-urlencoded的编码方式，将表单数据编码为 (name1=value1&name2=value2…)，然后把这个字符串append到url后面，用?分隔，跳转到这个新的url。
	// 当form表单的action为post时，浏览器将form数据封装到http body中，然后发到server。该格式不能提交文件。
	// 与 form-data 的主要区别在于不能提交文件
	e.POST("/paramtype/postxwwwformurlencoded", func(c *gin.Context) {
		postFormStruct := &PostxwwwFormStruct{}
		c.ShouldBind(postFormStruct)
		fmt.Printf("%+v \n", postFormStruct)

		// form-data 表单的 body 类型会把所有类型都识别为 string ,如果要传递其他类型，可以用映射，或者再自行转换
		postForm := c.PostForm("postForm")
		fmt.Println(postForm)

		defaultPostForm := c.DefaultPostForm("defaultPostForm", "no")
		fmt.Println(defaultPostForm)

		getPostForm, res := c.GetPostForm("getPostForm")
		fmt.Println(getPostForm, res)

		postFormArray := c.PostFormArray("postFormArray")
		fmt.Println(postFormArray)

		getPostFormArray, res := c.GetPostFormArray("getPostFormArray")
		fmt.Println(getPostFormArray, res)

		postFormMap := c.PostFormMap("postFormMap")
		fmt.Println(postFormMap)

		getPostFormMap, res := c.GetPostFormMap("getPostFormMap")
		fmt.Println(getPostFormMap, res)

		c.String(http.StatusOK, "ok")
	})

	// POST
	// raw
	// 可以上传任意格式的文本，包括text、json、xml、html等。
	// raw一般使用Bind和ShouldBind系列函数来获取数据，
	// 两者的区别是如果输入数据无效，Bind会将返回状态设置为400，并且中断，ShouldBind就没有这么狠。
	// Bind和ShouldBind根据Content-Type来判断是json/xml/yaml格式，做对应解析。
	// ShouldBindUri和ShouldBindQuery使用起来有些技巧，代码中对此有标注。
	// {"id":"1","text":"I am raw json"}
	// Content-Type: application/json
	// http://localhost:8080/paramtype/raw/collie/27?query=I am query
	e.POST("/paramtype/raw/:name/:age", func(c *gin.Context) {
		raw := &RawStruct{}
		c.ShouldBind(&raw) // 需要根据 Content-Type 来判断如何解析
		//c.Bind(raw) // bind 如果检测对应的键值类型错误，会返回400，但是如果对应键值只是找不到，则不会
		c.ShouldBindJSON(&raw) // 指定获取到的raw文本为json类型，即便Content-Type不为application/json也能直接解析
		fmt.Printf("%+v \n", raw)

		// 需定义合规结构体
		user := &User{}
		c.ShouldBindUri(user)
		fmt.Printf("%+v \n", user)

		// 需定义合规结构体
		rawQuery := &RawQuery{}
		c.ShouldBindQuery(rawQuery)
		fmt.Printf("%+v \n", rawQuery)

		c.String(http.StatusOK, "ok")
	})

	// POST
	// binary
	// 相当于Content-Type:application/octet-stream，从字面意思得知，只可以上传二进制数据，通常用来上传文件，由于没有键值，所以，一次只能上传一个文件,且没有后缀，路径等设置，故不推荐

	// Get的使用Query、Param系列
	// Post的form-data和x-www-form-urlencoded使用PostForm系列
	// Get和Post的数据都可以用Bind、ShouldBind系列，不过结构体的tag需要写正确
	// 文件上传使用form-data，通过函数FormFile或者MultipartForm

	e.Run()

}
