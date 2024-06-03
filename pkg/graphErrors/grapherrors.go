package grapherrors

import (
	"strings"

	"github.com/vektah/gqlparser/v2/gqlerror"
)



func TransformError(inputErr error) error {
	errList := gqlerror.List{}
	
	for _, err := range strings.Split(inputErr.Error(), "\n") {
		errList = append(errList, gqlerror.Errorf("%s", err))
	}

	return errList
}
