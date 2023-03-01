package route
 
import (
	"fmt"
	"reflect"
	"github.com/gin-gonic/gin"
	"strings"
)
//路由结构体
type Route struct {
	path string //url路径
	httpMethod string //http方法 get post
	Method reflect.Value //方法路由
	Args   []reflect.Type //参数类型
}
//路由集合
var Routes =[]Route{}
func InitRouter() *gin.Engine  {
	//初始化路由
	r := gin.Default()
	//绑定基本路由，访问路径：/User/List
	Bind(r)
	return r
}
 
//注册控制器
func Register(controller interface{}) bool{
	ctrlName:=reflect.TypeOf(controller).String()
	fmt.Println("ctrlName=",ctrlName)
	module := ctrlName
	if strings.Contains(ctrlName, ".") {
		module = ctrlName[strings.Index(ctrlName, ".") + 1:]
	}
	fmt.Println("module=",module)
	v := reflect.ValueOf(controller)
	//遍历方法
	for i:= 0; i < v.NumMethod(); i++ {
		method := v.Method(i)
		action := v.Type().Method(i).Name
		path := "/"+module+"/"+action
		//遍历参数
		params := make([]reflect.Type, 0, v.NumMethod())
		httpMethod :="POST";
		if( strings.Contains(action,"_get") ){
			httpMethod = "GET"
		}
		for j := 0; j < method.Type().NumIn(); j++ {
			params = append(params, method.Type().In(j))
			fmt.Println("params-name=",method.Type().In(j))
		}
		fmt.Println("params=",params)
		fmt.Println("action=",action)
		//if Routes[module] == nil {
		//	Routes[module] = make(map[string]Route)
		//}
		//Routes[module][action] = Route{method,params}
		route := Route{ path:path, Method:method,Args:params,httpMethod:httpMethod}
		Routes =append(Routes,route)
	}
	fmt.Println("Routes=",Routes)
	return true
}
//绑定路由 m是方法GET POST等
//绑定基本路由
func Bind(e *gin.Engine) {
	//pathInit()
	for _, route := range Routes {
		//e.POST(path, match(path))
		if(route.httpMethod=="GET"){
			e.GET(route.path,match(route.path,route))
		}
		if(route.httpMethod=="POST"){
			e.POST(route.path,match(route.path,route))
		}
 
	}
}
//根据path匹配对应的方法
func match(path string,route Route) gin.HandlerFunc {
	return func(c *gin.Context) {
		fields := strings.Split(path, "/")
		fmt.Println("fields,len(fields)=",fields,len(fields))
		if len(fields) < 3 {
			return
		}
 
		if len(Routes)>0 {
			arguments := make([]reflect.Value, 1)
			arguments[0] = reflect.ValueOf(c) // *gin.Context
			//reflect.ValueOf(method).Method(0).Call(arguments)
			route.Method.Call(arguments)
		}
	}
}