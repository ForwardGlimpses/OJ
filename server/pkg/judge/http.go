package judge

type httpClient struct {
	// 需要存储 judge 服务地址
}

func newHTTPClient() judgeInterface {
	return &httpClient{}
}

func (a *httpClient) Submit(req Request) (Response, error) {
	// 提交 + 判断逻辑

	return Response{}, nil
}
