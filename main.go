package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {

}

// Metadata ... Ec2 Machines Metadata
type Metadata struct {
	Port           int
	AmiID          string
	LocalHostname  string
	ProductCodes   string
	ReservationID  string
	PublicHostname string
	PublicIPV4     string
}

// GetInstanceID ...
func GetInstanceID() string {
	return Get("instance-id")
}

// GetType ...
func GetType() string {
	return Get("instance-type")
}

// GetPublicHostname ...
func GetPublicHostname() string {
	return Get("public-hostname")
}

// GetLocalHostname ...
func GetLocalHostname() string {
	return Get("local-hostname")
}

// GetPublicIP ...
func GetPublicIP() string {
	return Get("public-ipv4")
}

// GetLocalIP ...
func GetLocalIP() string {
	return Get("local-ipv4")
}

// Get ...
func Get(Key string) string {
	URL := []string{"http://169.254.169.254/latest/meta-data/", Key}
	log.Panicln("Link:", fmt.Sprintf("%s", URL))
	return GetMetaData(fmt.Sprintf("%s", URL))
}

// GetMetaData ...
func GetMetaData(URL string) string {
	req, _ := http.NewRequest("GET", URL, nil)

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	return string(body)
}
