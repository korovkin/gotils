package gotils_test

import (
	"net/http"
	"os"
	"testing"

	. "github.com/onsi/gomega"

	gotils "github.com/korovkin/gotils"
)

type MyObject struct {
	ErrorStr       string            `json:"err_str"`
	HttpStatusCode int               `json:"err_http_status_code"`
	ClientErrCode  int               `json:"err_client_code"`
	M              map[string]string `json:"m"`
	A              []string          `json:"a"`
}

func TestGOB(t *testing.T) {
	RegisterTestingT(t)

	t.Run("gob", func(_ *testing.T) {
		e := &MyObject{
			ErrorStr:       os.ErrNotExist.Error(),
			HttpStatusCode: http.StatusNotFound,
			ClientErrCode:  0,
			M: map[string]string{
				"a": "aa",
			},
			A: []string{
				"a", "aa", "aaa",
			},
		}

		eCopy := &MyObject{}
		gotils.FromGOB(gotils.ToGOB(e), eCopy)
		Expect(eCopy).NotTo(BeNil())

		// log.Println("=> e:", gotils.ToJSONString(e))
		// log.Println("=> eCopy:", gotils.ToJSONString(eCopy))

		Expect(eCopy.ErrorStr).To(BeEquivalentTo(e.ErrorStr))
		Expect(eCopy.HttpStatusCode).To(BeEquivalentTo(e.HttpStatusCode))
		Expect(eCopy.ClientErrCode).To(BeEquivalentTo(e.ClientErrCode))
	})
}
