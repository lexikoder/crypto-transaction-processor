package utils
import "github.com/go-playground/validator/v10"

func FormatValidationErrors(err error) map[string]string {
    res := make(map[string]string)
    if errs, ok := err.(validator.ValidationErrors); ok {
        for _, e := range errs {
            res[e.Field()] = e.Error() // default error or custom formatting
        }
    }
    return res
}