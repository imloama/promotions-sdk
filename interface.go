package sdk

//Request 请求接口
type Request interface {
}

//Response 返回结果
type Response interface {
}

// Server 服务端请求接口
type Server interface {
	// 执行请求
	execute(req *Request) (resp Response, err error)
}
