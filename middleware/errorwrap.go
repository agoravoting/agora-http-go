package middleware

import (
	"fmt"
	"github.com/agoravoting/agora-http-go/util"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"runtime/debug"
	"log"
)

// Ravenable is used to get the RavenClient from an object
type Ravenable interface {
	RavenClient() RavenClientIface
}

// HandledError is the type of managed error that can happen in app views
type HandledError struct {
	Err          error
	Code         int
	Message      string
	CodedMessage string
}

type handledErrorJson struct {
	Message      string `json:"error"`
	CodedMessage string `json:"error_code"`
}

// ErrorHandler is the signature of an app view that handles errors
type ErrorHandler func(rw http.ResponseWriter, r *http.Request, p httprouter.Params) *HandledError

// ErrorWrap is a struct used to create an instance of this middleware.
type ErrorWrap struct {
	Raven RavenClientIface
	Logger *log.Logger
}

// NewErrorWrap instantiates the ErrorWrap middleware
func NewErrorWrap(r Ravenable, l *log.Logger) *ErrorWrap {
	return &ErrorWrap{Raven: r.RavenClient(), Logger: l}
}

// ErrorWrap handles errors nicely and returns an standard httprouter.Handle.
func (ew *ErrorWrap) Do(handle ErrorHandler) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		if err := handle(w, r, p); err != nil {
			// record strange or internal errors
			if err.Code >= http.StatusRequestTimeout && !util.IsNil(ew.Raven) {
				msg := fmt.Sprintf("Error: code=%d, message='%s', err=%v, stack=%v", err.Code, err.Message, err.Err, debug.Stack())
				ew.Raven.CaptureMessage(msg)
			}
			ew.Logger.Printf("Error: code=%d, message='%s', err=%v", err.Code, err.Message, err.Err)

			content, err2 := util.JsonSortedMarshal(&handledErrorJson{err.Message, err.CodedMessage})
			if err2 != nil {
				panic(err2)
			}
			http.Error(w, string(content), err.Code)
		}
	}
}
