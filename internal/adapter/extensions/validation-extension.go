package extensions

import (
	"fmt"
	validation "walls-user-service/internal/core/helper/validation-helper"

	"github.com/gin-gonic/gin"
)

func ValidateBody(c *gin.Context, body interface{}) bool {
	err := validation.Validate(body)
	fmt.Println("body error:", err)
	// if err != nil {
	// 	c.AbortWithStatusJSON(400, err)
	// 	return false
	// }
	// return true
	return err == nil
}

func ValidateHeaders(c *gin.Context, currentUser interface{}) bool {
	err := validation.Validate(currentUser)
	fmt.Println("header error:", err)
	// if err != nil {
	// 	c.AbortWithStatusJSON(400, err)
	// 	return false
	// }
	// return true
	return err == nil
}
