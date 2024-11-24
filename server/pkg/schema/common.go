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
