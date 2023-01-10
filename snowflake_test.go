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
}

func TestSnowflake(t *testing.T) {
	RegisterTestingT(t)

	t.Run("snowflake", func(_ *testing.T) {
		now := time.Now()
		const idType = "BATMAN"

		a := gotils.SnowflakeID(idType, now)
		b := gotils.SnowflakeID(idType, now)
		log.Println("a:", a)
		log.Println("b:", b)
		Expect(a).NotTo(BeEquivalentTo(b))

		cGroup, c := gotils.SnowflakeIDWithGroup(idType, now)
		Expect(c).To(HavePrefix(idType))

		aGroup := gotils.SnowflakeExtractGroup(a, idType)
		group := fmt.Sprintf("%04d%02d%02d",
			now.UTC().Year(),
			now.UTC().Month(),
			now.UTC().Day())
		Expect(aGroup).To(BeEquivalentTo(group))
		Expect(cGroup).To(BeEquivalentTo(group))

	})
}
