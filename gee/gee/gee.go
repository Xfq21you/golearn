package gee

import (
	"net/http"
)

// HandlerFunc 定义 gee 使用的请求处理程序
type HandlerFunc func(*Context) //定义路由映射的处理方法

// Engine 实现了 ServeHTTP 的接口
type Engine struct {
	router *router
}

// New 是 gee. Engine 的构造函数
func New() *Engine {
	return &Engine{router: newRouter()}
}

func (engine *Engine) addRoute(method string, pattern string, handler HandlerFunc) { //添加了一张路由映射表router
	//key := method + "-" + pattern //key 由请求方法和静态路由地址构成
	//engine.router[key] = handler  //请求方法不同,可以映射不同的处理方法(Handler)，value 是用户映射的处理方法
	engine.router.addRoute(method, pattern, handler)
}

// GET 定义了添加 GET 请求的方法
func (engine *Engine) GET(pattern string, handler HandlerFunc) {
	engine.addRoute("GET", pattern, handler) //将路由和处理方法注册到映射表 router 中
}

// POST 定义了添加 POST 请求的方法
func (engine *Engine) POST(pattern string, handler HandlerFunc) {
	engine.addRoute("POST", pattern, handler)
}

// Run 定义启动 http 服务器的方法
func (engine *Engine) Run(addr string) (err error) {
	return http.ListenAndServe(addr, engine) //ListenAndServe 的包装
}

func (engine *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	// key := req.Method + "-" + req.URL.Path     //解析请求的路径，查找路由映射表
	// if handler, ok := engine.router[key]; ok { //如果查到，就执行注册的处理方法
	// 	handler(w, req)
	// } else {
	// 	fmt.Fprintf(w, "404 NOT FOUND: %s\n", req.URL) //如果查不到，就返回 404 NOT FOUND
	// }
	c := newContext(w, req)
	engine.router.handle(c)
}
