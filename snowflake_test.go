package gotils_test

import (
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

		a := gotils.SnowflakeID("BATMAN", now)
		b := gotils.SnowflakeID("BATMAN", now)

		Expect(a).NotTo(BeEquivalentTo(b))
	})
}
