package credentials

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"runtime"
)

const credentialsPath = "credentials.json"

// Credentials of Spotify application
type Credentials struct {
	Identifier string `json:"SPOTIFY_ID"`
	Secret     string `json:"SPOTIFY_SECRET"`
}

func getCurrentPath() string {
	_, filename, _, _ := runtime.Caller(1)
	return path.Dir(filename)
}

// Get retrieves application credentials
func Get() Credentials {
	fmt.Println(getCurrentPath() + credentialsPath)
	jsonFile, err := os.Open(getCurrentPath() + "/" + credentialsPath)
	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)
	var credentials Credentials
	json.Unmarshal(byteValue, &credentials)
	return credentials
}
