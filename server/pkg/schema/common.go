package schema

type ID struct {
	ID int `json:"id" form:"id"`
}

type BaseFiled struct{
	
}

type Response struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data,omitempty"`
	Error   error       `json:"error,omitempty"`
}
