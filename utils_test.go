package gotils_test

import (
	"fmt"
	"log"
	"testing"
	"time"

	. "github.com/onsi/gomega"
	"github.com/stretchr/testify/assert"

	gotils "github.com/korovkin/gotils"
)

func init() {
	log.SetFlags(log.Ltime | log.Lshortfile | log.Lmicroseconds | log.Ldate)
}

func TestImplode(t *testing.T) {
	s := gotils.Implode("superman", "batman", "spiderman", "wonderwoman")
	l := gotils.Explode(s)
	assert.Equal(t, 4, len(l))
	assert.Equal(t, "superman", l[0])
	assert.Equal(t, "batman", l[1])
	assert.Equal(t, "spiderman", l[2])
	assert.Equal(t, "wonderwoman", l[3])
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
	superman, batman := gotils.Explode2(gotils.Implode2("superman", "batman"))
	assert.Equal(t, superman, "superman", "")
	assert.Equal(t, batman, "batman", "")
}

func TestImplode3(t *testing.T) {
	superman, batman, spiderman := gotils.Explode3(gotils.Implode3("superman", "batman", "spiderman"))
	assert.Equal(t, superman, "superman", "")
	assert.Equal(t, batman, "batman", "")
	assert.Equal(t, spiderman, "spiderman", "")
}

func TestUniqueID(t *testing.T) {
	superman, batman, spiderman := gotils.Explode3(gotils.Implode3("superman", "batman", "spiderman"))
	assert.Equal(t, superman, "superman", "")
	assert.Equal(t, batman, "batman", "")
	assert.Equal(t, spiderman, "spiderman", "")
}

func TestUniqueIDs(t *testing.T) {
	RegisterTestingT(t)

	t.Run("ids", func(_ *testing.T) {
		now := time.Now()
		a, aa := gotils.GenerateUniqueID("test", now)
		log.Println("=> a:", a, aa)

	})
}

func TestUniqueIDIP(t *testing.T) {
	machineid := gotils.IPBasedMachineID()
	log.Println("=> machineid:", machineid)
}
