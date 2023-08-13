package form

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func GetFormHandler(ctx *gin.Context) {
	fmt.Fprintf(ctx.Writer, "Read Form")
}
