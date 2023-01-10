package gotils

import (
	"bytes"
	"compress/gzip"
	"crypto/md5"
	"encoding/gob"
	"encoding/hex"
	"encoding/json"
	"encoding/xml"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"net/url"
	"os"
	"runtime/debug"
	"strconv"
	"strings"
	"sync/atomic"
	"time"
)

const (
	sep       = ":::"
	httpOkMin = 200
	httpOkmax = 300
)

var (
	EmptyJSONBytes  = []byte("{}")
	EmptyJSONString = string(EmptyJSONBytes)
)

// FromJSONBytes parses JSON bytes buffer into o
func FromJSONBytes(buf []byte, o interface{}) error {
	err := json.Unmarshal(buf, &o)
	return err
}

// FromJSONString parses JSON string into o
func FromJSONString(buf string, o interface{}) error {
	err := json.Unmarshal([]byte(buf), &o)
	return err
}

// ToJSONBytes serializes v into a buffer of JSON bytes
func ToJSONBytes(v interface{}) []byte {
	bytes, err := json.MarshalIndent(v, " ", " ")
	CheckNotFatal(err)

	if err == nil {
		return bytes
	}
	return EmptyJSONBytes
}

// ToJSONBytesOrError serializes v into a buffer of JSON bytes
func ToJSONBytesOrError(v interface{}) ([]byte, error) {
	bytes, err := json.MarshalIndent(v, " ", " ")
	CheckNotFatal(err)

	if err == nil {
		return bytes, err
	}
	return EmptyJSONBytes, err
}

// ToJSONBytesNoIndent serializes v into a buffer of JSON bytes
func ToJSONBytesNoIndent(v interface{}) []byte {
	bytes, err := json.Marshal(v)
	CheckNotFatal(err)
	if err == nil {
		return bytes
	}
	return EmptyJSONBytes
}

// ToJSONBytesNoIndent serializes v into a buffer of JSON bytes
func ToJSONBytesNoIndentOrError(v interface{}) ([]byte, error) {
	bytes, err := json.Marshal(v)
	CheckNotFatal(err)
	if err == nil {
		return bytes, err
	}
	return EmptyJSONBytes, err
}

// ToJSONString serializes v into a JSON string
func ToJSONString(v interface{}) string {
	bytes := ToJSONBytes(v)
	return string(bytes)
}

// ToJSONStringOrError serializes v into a JSON string
func ToJSONStringOrError(v interface{}) (string, error) {
	bytes, err := ToJSONBytesOrError(v)
	return string(bytes), err
}

// ToXMLString serializes v into an XML string
func ToXMLString(v interface{}) string {
	bytes, err := xml.MarshalIndent(v, " ", " ")
	CheckNotFatal(err)
	if err == nil {
		return string(bytes)
	}
	return ""
}

// ToJSONStringNoIndent serialize v into a JSON string
func ToJSONStringNoIndent(v interface{}) string {
	bytes := ToJSONBytesNoIndent(v)
	return string(bytes)
}

// KeepShort keeps the string short
func KeepShort(s *string, maxLenght int) *string {
	if s == nil {
		NIL := "<nil>"
		return &NIL
	}

	if len(*s) > maxLenght {
		ss := (*s)[0:maxLenght] + " ..."
		return &ss
	}

	return s
}

// CheckFatal checks the error and panics if it's not nil
func CheckFatal(e error) error {
	if e != nil {
		debug.PrintStack()
		log.Println("CHECK: FATAL: ERROR:", e)
		panic(e)
	}

	return e
}

// CheckNotFatal checks the error, logs it if not nil, and moves on
func CheckNotFatal(e error) error {
	if e != nil {
		debug.PrintStack()
		log.Println("CHECK: NOT FATAL: ERROR:", e, e.Error())
	}
	return e
}

// ReadJSONFile reads a JSON file into o
func ReadJSONFile(filename string, o interface{}) error {
	file, err := ioutil.ReadFile(filename)
	CheckNotFatal(err)
	if err != nil {
		return err
	}
	err = json.Unmarshal(file, o)
	CheckNotFatal(err)
	return err
}

// WriteJSONFile writes the object 'o' as JSON into the given file
func WriteJSONFile(filename string, o interface{}) error {
	log.Println("WriteJSONFile:", filename)

	err := ioutil.WriteFile(filename, []byte(ToJSONString(o)), 0666)
	CheckNotFatal(err)
	if err != nil {
		return err
	}
	return err
}

// WriteJSONFileCompressed write o as a compressed JSON file (gzip)
func WriteJSONFileCompressed(filename string, o interface{}) error {
	log.Println("WriteJSONFileCompressed:", filename)

	var b bytes.Buffer
	w, err := gzip.NewWriterLevel(&b, gzip.BestCompression)
	CheckNotFatal(err)

	if err != nil {
		return err
	}
	defer w.Close()

	_, err = w.Write(ToJSONBytes(o))
	CheckNotFatal(err)
	if err != nil {
		return err
	}

	w.Close()

	err = ioutil.WriteFile(filename, b.Bytes(), 0666)
	CheckNotFatal(err)
	if err != nil {
		return err
	}

	return err
}

// ReadJSONFileCompressed read compressed (gzip) JSON file
func ReadJSONFileCompressed(filename string, o interface{}) error {
	f, err := os.Open(filename)
	CheckNotFatal(err)

	w, err := gzip.NewReader(f)
	CheckNotFatal(err)
	if err != nil {
		return err
	}
	defer w.Close()

	b, err := ioutil.ReadAll(w)
	CheckNotFatal(err)
	if err != nil {
		return err
	}

	err = FromJSONBytes(b, o)
	CheckNotFatal(err)
	if err != nil {
		return err
	}
	return err
}

// ReadXMLFile reads an XML file into o
func ReadXMLFile(filename string, o interface{}) error {
	file, err := ioutil.ReadFile(filename)
	CheckNotFatal(err)
	if err != nil {
		return err
	}
	err = xml.Unmarshal(file, o)
	CheckNotFatal(err)
	return err
}

// TouchFile creates a filename if it doesn't exist yet
func TouchFile(filename string) error {
	file, err := os.Create(filename)
	CheckNotFatal(err)
	if file != nil {
		file.Close()
	}
	return err
}

// HTTPJSONGet getches JSON from the given url
func HTTPJSONGet(url string, o interface{}, username string, password string) error {
	var err error
	body := []byte{}

	req, err := http.NewRequest("GET", url, nil)
	CheckNotFatal(err)
	if err != nil {
		return err
	}

	if username != "" || password != "" {
		req.SetBasicAuth(username, password)
	}

	client := http.Client{}
	resp, err := client.Do(req)
	CheckNotFatal(err)
	if err != nil {
		return err
	}

	if resp.StatusCode < httpOkMin || resp.StatusCode >= httpOkmax {
		err = fmt.Errorf("ERROR: HTTP: CODE: %d", resp.StatusCode)
		CheckNotFatal(err)
		return err
	}

	defer resp.Body.Close()
	body, err = ioutil.ReadAll(resp.Body)
	CheckNotFatal(err)
	if err != nil {
		return err
	}

	json.Unmarshal(body, o)
	CheckNotFatal(err)
	if err != nil {
		return err
	}

	return nil
}

// HTTPXMLGet fetches XML payload and parses it into o
func HTTPXMLGet(url string, o interface{}) error {
	resp, err := http.Get(url)
	CheckNotFatal(err)
	if err != nil {
		return err
	}

	if resp.StatusCode < httpOkMin || resp.StatusCode >= httpOkmax {
		return fmt.Errorf("ERROR: HTTP: CODE: %d", resp.StatusCode)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	CheckNotFatal(err)
	if err != nil {
		return err
	}

	xml.Unmarshal(body, o)
	CheckNotFatal(err)
	if err != nil {
		return err
	}

	return err
}

// HTTPBytesGet fetches payload into a buffer
func HTTPBytesGet(url string) ([]byte, *http.Header, error) {
	var body []byte
	var headers *http.Header

	req, err := http.NewRequest("GET", url, nil)
	CheckNotFatal(err)

	if err != nil {
		return body, headers, err
	}

	const username = ""
	const password = ""

	if username != "" || password != "" {
		req.SetBasicAuth(username, password)
	}

	client := http.Client{}
	resp, err := client.Do(req)
	CheckNotFatal(err)
	if err != nil {
		return body, headers, err
	}

	if resp.StatusCode < httpOkMin || resp.StatusCode >= httpOkmax {
		err = fmt.Errorf("ERROR: HTTP: CODE: %d", resp.StatusCode)
		CheckNotFatal(err)
		return body, &resp.Header, err
	}

	defer resp.Body.Close()
	body, err = ioutil.ReadAll(resp.Body)
	CheckNotFatal(err)
	if err != nil {
		return body, &resp.Header, err
	}

	return body, &resp.Header, err
}

// MD5Hash computes MD5 hash on text
func MD5Hash(text string) string {
	hasher := md5.New()
	hasher.Write([]byte(text))
	return hex.EncodeToString(hasher.Sum(nil))
}

// URLIsValid check if a given string is a valid URL
func URLIsValid(urlString string) error {
	_, err := url.Parse(urlString)
	CheckNotFatal(err)
	return err
}

// Implode2 cancatenates two strings with a seperator 'sep'
func Implode2(a, b string) string {
	return strings.Join([]string{a, b}, sep)
}

// Implode3 cancatenates two strings with a seperator 'sep'
func Implode3(a, b, c string) string {
	return strings.Join([]string{a, b, c}, sep)
}

// Explode2 splits the string by a seperator 'sep'
func Explode2(x string) (a, b string) {
	l := strings.Split(x, sep)
	if len(l) > 1 {
		return l[0], l[1]
	} else if len(l) > 0 {
		return l[0], ""
	}
	return "", ""
}

// Explode3 splits the string by a seperator 'sep'
func Explode3(x string) (a, b, c string) {
	l := strings.Split(x, sep)
	if len(l) > 2 {
		return l[0], l[1], l[2]
	} else if len(l) > 1 {
		return l[0], l[1], ""
	} else if len(l) > 0 {
		return l[0], "", ""
	}
	return "", "", ""
}

// PrivateIPV4Get gets the private IP of the current AWS/EC2 server
func PrivateIPV4Get() (net.IP, error) {
	as, err := net.InterfaceAddrs()
	if err != nil {
		return nil, err
	}

	for _, a := range as {
		ipnet, ok := a.(*net.IPNet)
		if !ok || ipnet.IP.IsLoopback() {
			continue
		}

		ip := ipnet.IP.To4()
		if IsPrivateIPV4Get(ip) {
			return ip, err
		}
	}
	return nil, errors.New("no private ip address")
}

// IsPrivateIPV4Get checks if the given IP is private (AWS/EC2)
func IsPrivateIPV4Get(ip net.IP) bool {
	return ip != nil &&
		(ip[0] == 10 || ip[0] == 172 && (ip[1] >= 16 && ip[1] < 32) || ip[0] == 192 && ip[1] == 168)
}

// PrivateIPV4GetLower16 gets the lower 16 bits of the IP address
func PrivateIPV4GetLower16() (uint16, error) {
	ip, err := PrivateIPV4Get()
	if err != nil {
		return 0, err
	}

	return uint16(ip[2])<<8 + uint16(ip[3]), nil
}

// PrivateIPV4GetLower16OrDie gets the lower 16 bits of the IP address
func PrivateIPV4GetLower16OrDie() uint16 {
	ip, err := PrivateIPV4Get()
	if err != nil {
		log.Fatalln("PrivateIPV4GetLower16OrDie: failed:", err.Error())
		return 0
	}
	return uint16(ip[2])<<8 + uint16(ip[3])
}

func PrivateIPV4GetLower32OrDie() uint32 {
	ip, err := PrivateIPV4Get()
	if err != nil {
		log.Fatalln("PrivateIPV4GetLower32OrDie: failed:", err.Error())
		return 0
	}
	return uint32(ip[0])<<24 + uint32(ip[1])<<16 + uint32(ip[2])<<8 + uint32(ip[3])
}

// HTTPParamGetInt get an integer param from an HTTP request
func HTTPParamGetInt(req *http.Request, key string, defaultValue int) int {
	vv := req.URL.Query().Get(key)
	v := defaultValue
	if vv != "" {
		v, _ = strconv.Atoi(vv)
	}
	return v
}

// HTTPParamGetString get a string param from an HTTP request
func HTTPParamGetString(req *http.Request, key string, defaultValue string) string {
	vv := req.URL.Query().Get(key)
	v := defaultValue
	if vv != "" {
		v = vv
	}
	return v
}

var monotonicUniqueCounter = uint32(0)
var machineID = PrivateIPV4GetLower32OrDie()

// GenerateUniqueID sorted by time
func GenerateUniqueID(idType string, nowLocal time.Time) (string, string) {
	var i = (atomic.AddUint32(&monotonicUniqueCounter, 1)) % 0xFFFF

	// UTC as some machines can be in different time zones:
	now := nowLocal.UTC()

	dailyFolderID := fmt.Sprintf(
		"%04d%02d%02d",
		now.Year(),
		now.Month(),
		now.Day())

	uniqueID := fmt.Sprintf("%s_%04d%02d%02d%02d%02d%02d%08x%02x",
		idType,
		now.Year(),
		now.Month(),
		now.Day(),
		now.Hour(),
		now.Minute(),
		now.Second(),
		machineID,
		i)

	return dailyFolderID, uniqueID
}

// ToGOB serialize an interface into a GOB
func ToGOB(o interface{}) []byte {
	var network bytes.Buffer
	enc := gob.NewEncoder(&network)
	err := enc.Encode(o)
	CheckNotFatal(err)

	if err != nil {
		return nil
	}

	return network.Bytes()
}

// FromGOB parse GOB blob into an interface object
func FromGOB(network []byte, o interface{}) error {
	dec := gob.NewDecoder(bytes.NewBuffer(network))
	err := dec.Decode(o)
	CheckNotFatal(err)
	return err
}
