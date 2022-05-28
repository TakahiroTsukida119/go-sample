package exception

import (
	"errors"
	"fmt"
	"net/http"
)

type Exception struct {
	StatusCode int
	Err        error
}

func (e *Exception) Error() string {
	return e.Err.Error()
}

func NewRecordNotFound(table string) *Exception {
	return &Exception{
		StatusCode: http.StatusNotFound,
		Err:        errors.New(fmt.Sprintf("%sテーブルに、該当するデータが見つかりませんでした。", table)),
	}
}
