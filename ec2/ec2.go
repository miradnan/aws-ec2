package ec2

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

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
	return GetMetaData("instance-id")
}

// GetInstanceType ...
func GetInstanceType() string {
	return GetMetaData("instance-type")
}

// GetPublicHostname ...
func GetPublicHostname() string {
	return GetMetaData("public-hostname")
}

// GetLocalHostname ...
func GetLocalHostname() string {
	return GetMetaData("local-hostname")
}

// GetPublicIP ...
func GetPublicIP() string {
	return GetMetaData("public-ipv4")
}

// GetLocalIP ...
func GetLocalIP() string {
	return GetMetaData("local-ipv4")
}

// GetMetaData ...
func GetMetaData(Key string) string {
	URL := fmt.Sprintf("http://169.254.169.254/latest/meta-data/%s", Key)

	req, _ := http.NewRequest("GET", URL, nil)

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	return string(body)
}
