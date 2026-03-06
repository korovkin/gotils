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

func TestImplode(t *testing.T) {
	RegisterTestingT(t)

	s := gotils.Implode("superman", "batman", "spiderman", "wonderwoman")
	l := gotils.Explode(s)

	Expect(l).To(HaveLen(4))
	Expect(l).To(Equal([]string{"superman", "batman", "spiderman", "wonderwoman"}))
}

func ExampleImplode() {
	s := gotils.Implode("one", "two", "three")
	fmt.Println(s)
	// Output: one:::two:::three
}

func ExampleExplode() {
	l := gotils.Explode("one:::two:::three")
	fmt.Println(l)
	// Output: [one two three]
}

func TestImplode2(t *testing.T) {
	RegisterTestingT(t)

	superman, batman := gotils.Explode2(gotils.Implode2("superman", "batman"))
	Expect(superman).To(Equal("superman"))
	Expect(batman).To(Equal("batman"))
}

func TestImplode3(t *testing.T) {
	RegisterTestingT(t)

	superman, batman, spiderman := gotils.Explode3(gotils.Implode3("superman", "batman", "spiderman"))
	Expect(superman).To(Equal("superman"))
	Expect(batman).To(Equal("batman"))
	Expect(spiderman).To(Equal("spiderman"))
}

func TestUniqueID(t *testing.T) {
	RegisterTestingT(t)

	superman, batman, spiderman := gotils.Explode3(gotils.Implode3("superman", "batman", "spiderman"))
	Expect(superman).To(Equal("superman"))
	Expect(batman).To(Equal("batman"))
	Expect(spiderman).To(Equal("spiderman"))
}

func TestUniqueIDs(t *testing.T) {
	RegisterTestingT(t)

	t.Run("ids", func(_ *testing.T) {
		now := time.Now()
		a, aa := gotils.GenerateUniqueID("test", now)
		log.Println("=> a:", a, aa)
		Expect(a).NotTo(BeEmpty())
		Expect(aa).NotTo(BeEmpty())
	})
}

func TestUniqueIDIP(t *testing.T) {
	RegisterTestingT(t)
	machineid := gotils.PrivateIPV4GetLower32OrDie()
	log.Printf("=> machineid: %08x\n", machineid)
	Expect(machineid).NotTo(BeZero())
}
