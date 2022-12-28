package gotils_test

import (
	"bytes"
	"encoding/gob"
	"net/http"
	"os"
	"testing"

	. "github.com/onsi/gomega"

	gotils "github.com/korovkin/gotils"
)

type Error struct {
	ErrorStr       string `json:"err_str"`
	HttpStatusCode int    `json:"err_http_status_code"`
	ClientErrCode  int    `json:"err_client_code"`
}

// TODO:: This is where generics are missing:

func ErrorToGOB(o *Error) []byte {
	var network bytes.Buffer
	enc := gob.NewEncoder(&network)
	err := enc.Encode(*o)
	gotils.CheckNotFatal(err)
	return network.Bytes()
}

func ErrorFromGOB(network []byte) *Error {
	var s Error

	dec := gob.NewDecoder(bytes.NewBuffer(network))
	err := dec.Decode(&s)
	gotils.CheckNotFatal(err)

	return &s
}

func TestGOB(t *testing.T) {
	RegisterTestingT(t)

	t.Run("gob", func(_ *testing.T) {
		e := &Error{
			ErrorStr:       os.ErrNotExist.Error(),
			HttpStatusCode: http.StatusNotFound,
			ClientErrCode:  0,
		}

		eCopy := ErrorFromGOB(ErrorToGOB(e))
		Expect(eCopy).NotTo(BeNil())
		// log.Println("=> e:", gotils.ToJSONString(e))
		// log.Println("=> eCopy:", gotils.ToJSONString(eCopy))
		Expect(eCopy.ErrorStr).To(BeEquivalentTo(os.ErrNotExist.Error()))
	})
}
