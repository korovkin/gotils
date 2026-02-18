package gotils

import (
	"errors"
	"fmt"
	"log"
	"regexp"
	"strings"
	"sync/atomic"
	"time"
)

var ErrorInvalidId = errors.New("invalid Snowflake ID")

var snowflakeMonoCount = uint32(0)
var snowflakeMachineID = IPBasedMachineID()
var reOnlyChars *regexp.Regexp = regexp.MustCompile(`[a-zA-Z0-9]+`)

func SnowflakeID(idType string, nowLocal time.Time) string {
	uniqueID := ""
	nowUTC := nowLocal.UTC()

	if reOnlyChars.MatchString(idType) {
		var uniqueC = (atomic.AddUint32(&snowflakeMonoCount, 1)) % 0xFFFF
		uniqueID = fmt.Sprintf("%s_%04d%02d%02d%02d%02d%02d_%12s%04x",
			idType,
			nowUTC.Year(),
			nowUTC.Month(),
			nowUTC.Day(),
			nowUTC.Hour(),
			nowUTC.Minute(),
			nowUTC.Second(),
			snowflakeMachineID,
			uniqueC)
	} else {
		log.Fatalln("Invalid ID:", idType)
		CheckFatal(ErrorInvalidId)
	}

	return uniqueID
}

func SnowflakeIDWithGroup(idType string, nowLocal time.Time) (groupID string, uniqueID string) {
	nowUTC := nowLocal.UTC()
	if reOnlyChars.MatchString(idType) {
		var uniqueC = (atomic.AddUint32(&snowflakeMonoCount, 1)) % 0xFFFF
		groupID = fmt.Sprintf("%04d%02d%02d",
			nowUTC.Year(),
			nowUTC.Month(),
			nowUTC.Day())
		uniqueID = fmt.Sprintf("%s_%04d%02d%02d%02d%02d%02d_%12s%04x",

			idType,
			nowUTC.Year(),
			nowUTC.Month(),
			nowUTC.Day(),
			nowUTC.Hour(),
			nowUTC.Minute(),
			nowUTC.Second(),
			snowflakeMachineID,
			uniqueC)
	} else {
		log.Fatalln("Invalid ID:", idType)
		CheckFatal(ErrorInvalidId)
	}

	return groupID, uniqueID
}

func SnowflakeExtractGroup(id string, idType string) string {
	groupID := ""

	components := strings.Split(id, idType)
	if len(components) > 1 {
		groupID = components[1][1:9]
	}

	return groupID
}
