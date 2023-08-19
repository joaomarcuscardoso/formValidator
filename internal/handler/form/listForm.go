package form

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/pdfcrowd/pdfcrowd-go"
)

func GetFormHandler(ctx *gin.Context) {

	logger.Infof("Read Form")
	fmt.Fprintf(ctx.Writer, "Read Form")

	client := pdfcrowd.NewPdfToHtmlClient("demo", "ce544b6ea52a5621fb9d55f8b542d14d")

	err := client.ConvertUrlToFile("https://pdfcrowd.com/static/pdf/apisamples/invoice.pdf", "templates/pdf.html")

	// check for the conversion errors
	if err != nil {
		why, ok := err.(pdfcrowd.Error)
		if ok {
			logger.Errorf("Pdfcrowd Error: %s\n", why)
		} else {
			logger.Errorf("Generic Error: %s\n", err)
		}
		panic(err.Error())
	}
}
