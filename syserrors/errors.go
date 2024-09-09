package syerrors

import "fmt"

func Newf(format string, args ...any) error {
	return fmt.Errorf(format, args...)
}

type jsonPayload struct {
	Message *string `json:"message,omitempty"`
}

func JSONPayload(err error) any {
	if err == nil {
		return nil
	}

	message := err.Error()

	return &jsonPayload{
		Message: &message,
	}
}

func payloadCode(err error) string {
	return ""
}

func snakeCase(s string) string {
	return s
}
