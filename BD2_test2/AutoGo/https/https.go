package https

// Get 发送 GET 请求并返回响应状态码和数据。
// 参数:
//
//	url: 请求的 URL。
//	timeout: 请求的超时时间（毫秒），如果为 0 则不设置超时。
//
// 返回:
//
//	code: 响应的状态码。
//	data: 响应的数据。
func Get(url string, timeout int) (code int, data []byte) {
	return 0, nil
}

// PostMultipart 发送带有文件的 POST 请求并返回响应状态码和数据。
// 参数:
//
//	url: 请求的 URL。
//	fileName: 文件名。
//	fileData: 文件数据。
//	timeout: 请求的超时时间（毫秒），如果为 0 则不设置超时。
//
// 返回:
//
//	code: 响应的状态码。
//	data: 响应的数据。
func PostMultipart(url string, fileName string, fileData []byte, timeout int) (code int, data []byte) {
	return 0, nil
}

// Post 发送 POST 请求并返回响应状态码和数据。
// 参数:
//
//	url: 请求的 URL。
//	data: 请求体数据（如 JSON 序列化后的字节数组）。
//	headers: 自定义请求头，如果为 nil 或未设置 Content-Type，默认使用 application/json。
//	timeout: 请求的超时时间（毫秒），如果为 0 则不设置超时。
//
// 返回:
//
//	code: 响应的状态码。
//	body: 响应的数据。
func Post(url string, data []byte, headers map[string]string, timeout int) (code int, body []byte) {
	return 0, nil
}
