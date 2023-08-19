package form

import (
	"github.com/joaomarcuscardoso/formValidator/internal/config"
)

var (
	logger *config.Logger
)

func InitializarHandler() {
	logger = config.GetLogger("handler")
}
