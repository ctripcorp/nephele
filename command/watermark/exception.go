package watermark

import (
	"strings"
)

// exceptionType returns gmagick's error type, gmagick returns error with format: error_type: error_message
func exceptionType(err error) string {
	i := strings.Index(err.Error(), ":")
	if i == -1 {
		return "UnknownError"
	}
	return err.Error()[:i]
}
