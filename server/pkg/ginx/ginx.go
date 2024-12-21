package ginx

import (
	"encoding/json"
	"net/http"

	"github.com/ForwardGlimpses/OJ/server/pkg/errors"
	"github.com/ForwardGlimpses/OJ/server/pkg/schema"
	"github.com/gin-gonic/gin"
)

func ResJSON(c *gin.Context, status int, v interface{}) {
	buf, err := json.Marshal(v)
	if err != nil {
		panic(err)
	}

	c.Data(status, "application/json; charset=utf-8", buf)
	c.Abort()
}

func ResSuccess(c *gin.Context, v interface{}) {
	ResJSON(c, http.StatusOK, schema.Response{
		Success: true,
		Data:    v,
	})
}

func ResOK(c *gin.Context) {
	ResJSON(c, http.StatusOK, schema.Response{
		Success: true,
	})
}

func ResError(c *gin.Context, err error) {
	var merr *errors.Error
	if e, ok := errors.As(err); ok {
		merr = e
	} else {
		merr = errors.InternalServer("internal server error: %v", err)
	}
	ResJSON(c, merr.Code, merr)
}
