package errs

import "fmt"

var (
	ERROR_CONFIG_INVALID error = fmt.Errorf("[ERROR]: config file path is invalid")
)
