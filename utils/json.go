package utils

import (
	"encoding/json"

	"github.com/gin-gonic/gin"
)

func ReturnJSONStruct(c *gin.Context, genericStruct interface{}) {
	c.Writer.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(c.Writer).Encode(genericStruct)
	if err != nil {
		return
	}
}
