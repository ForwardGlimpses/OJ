package judge

import "github.com/ForwardGlimpses/OJ/server/pkg/config"

func Init() {
	// 初始化的时候从 config 中读取地址相关信息传进去
	if config.C.Judge.UseHTTP {
		judge = newHTTPClient()
	}
}

type judgeInterface interface {
	Submit(req Request) (Response, error)
}

var judge judgeInterface

type Request struct {
	ID     int    // 提交的唯一标识
	Code   string // 提交代码
	Input  string // 输入
	Output string // 输出
}

type Response struct {
	Status  string // 判定结果
	Memory  uint64 //  使用内存
	Time    uint64 // 判断时间？
	RunTime uint64 // 耗时？
}

func Submit(req Request) (Response, error) {
	return judge.Submit(req)
}
