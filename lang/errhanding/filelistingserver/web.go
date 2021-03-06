package main

import (
	"log"
	"net/http"
	"os"
	_ "net/http/pprof"
	"richie.com/richie/learn_golang/lang/errhanding/filelistingserver/filelisting"
)

type appHandle func(writer http.ResponseWriter,
	request *http.Request) error

func errWrapper(
	handler appHandle) func(
	http.ResponseWriter, *http.Request) {

	return func(writer http.ResponseWriter,
		request *http.Request) {
		// panic
		defer func() {
			if r := recover(); r != nil {
				log.Printf("Panic: %v", r)
				http.Error(writer,
					http.StatusText(http.StatusInternalServerError),
					http.StatusInternalServerError)
			}
		}()

		err := handler(writer, request)
		//log.Warn()
		if err != nil {
			log.Printf("Error occurred" +
				" handling request: %s", err.Error())
			// user error
			if userErr, ok := err.(userError); ok{
				http.Error(writer, userErr.Message(), http.StatusBadRequest)
				return
			}
			// system error
			code := http.StatusOK
			switch {
			case os.IsNotExist(err):
				code = http.StatusNotFound
			case os.IsPermission(err):
				code = http.StatusForbidden
			default:
				code = http.StatusInternalServerError
			}
			http.Error(writer, http.StatusText(code), code)
		}

	}
}

type userError interface {
	error
	Message() string
}

func main() {
	http.HandleFunc("/",
		errWrapper(filelisting.HandleFileList))
	err := http.ListenAndServe(":8888", nil)
	if err != nil {
		panic(err)
	}
}
