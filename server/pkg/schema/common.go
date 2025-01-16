package schema

type ID struct {
	//ID int `json:"id" form:"id"`
	ID int `uri:"id" binding:"required"`
}

type BaseFiled struct {
}

type Response struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data,omitempty"`
	Error   error       `json:"error,omitempty"`
}

type P struct {
	Page     int
	PageSize int
}

type QueryResult[T any] struct {
	Items      T     // 当前页的数据
	TotalCount int64 // 总记录数
	Page       int   // 当前页码
	PageSize   int   // 每页条数
}
