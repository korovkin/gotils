package gotils_test

import (
	"fmt"
	"log"
	"testing"
	"time"

	"github.com/korovkin/gotils"
	. "github.com/onsi/gomega"
)

func init() {
	log.SetFlags(log.Ltime | log.Lshortfile | log.Lmicroseconds | log.Ldate)
	log.Println("=> init")
}

func TestSnowflake(t *testing.T) {
	RegisterTestingT(t)

	t.Run("snowflake", func(_ *testing.T) {
		now := time.Now()
		const idType = "BATMAN"

		a := gotils.SnowflakeID(idType, now)
		b := gotils.SnowflakeID(idType, now)
		Expect(a).NotTo(BeEquivalentTo(b))
		log.Println("a:", a)
		log.Println("b:", b)

		cGroup, c := gotils.SnowflakeIDWithGroup(idType, now)
		log.Println("c:", c, "cGroup:", cGroup)

		aGroup := gotils.SnowflakeExtractGroup(a, idType)
		group := fmt.Sprintf("%04d%02d%02d",
			now.Year(),
			now.Month(),
			now.Day())
		Expect(aGroup).To(BeEquivalentTo(group))
		Expect(cGroup).To(BeEquivalentTo(group))

	})
}
