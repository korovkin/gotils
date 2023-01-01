package gotils_test

import (
	"log"
	"testing"

	"github.com/korovkin/gotils"
	. "github.com/onsi/gomega"
)

type TestStruct struct {
	TSStart gotils.Time `json:"ts_start_sec,omitempty"`
	TSEnd   gotils.Time `json:"ts_end_sec,omitempty"`
}

func init() {
	log.SetFlags(log.Ltime | log.Lshortfile | log.Lmicroseconds | log.Ldate)
}

func TestTime(t *testing.T) {
	RegisterTestingT(t)

	t.Run("zero_time", func(_ *testing.T) {
		// 2022-12-29 23:40:50 +0000 UTC
		// 2022-12-29 15:40:50 -0800 PST
		var ts gotils.Time
		log.Println("=> ts: zero:", ts)
		log.Println("=> ts: zero: json:", gotils.ToJSONString(ts))

		response := &TestStruct{}
		responseMap := map[string]interface{}{}
		gotils.FromJSONString(gotils.ToJSONString(response), &responseMap)
		Expect(responseMap["ts_start_sec"].(float64)).To(BeEquivalentTo(0))
		Expect(responseMap["ts_end_sec"].(float64)).To(BeEquivalentTo(0))
	})

	t.Run("time", func(_ *testing.T) {
		// 2022-12-29 23:40:50 +0000 UTC
		// 2022-12-29 15:40:50 -0800 PST
		ts := gotils.FromUnix(1672357250)
		Expect(ts.UTC().Hour()).To(BeEquivalentTo(23))
		Expect(ts.UTC().Minute()).To(BeEquivalentTo(40))
		Expect(ts.UTC().Second()).To(BeEquivalentTo(50))

		// serialize to JSON and back
		{
			response := &TestStruct{
				TSStart: ts,
				TSEnd:   ts,
			}
			responseMap := map[string]interface{}{}
			gotils.FromJSONString(gotils.ToJSONString(response), &responseMap)

			// the value is in UNIX time:
			Expect(responseMap["ts_start_sec"].(float64)).To(BeEquivalentTo(1672357250))
			Expect(responseMap["ts_end_sec"].(float64)).To(BeEquivalentTo(1672357250))

			// parse it:
			responseParsed := &TestStruct{}
			gotils.FromJSONString(gotils.ToJSONString(response), responseParsed)
			// log.Println("=> responseParsed: ", gotils.ToJSONString(responseParsed))
		}

	})

}
