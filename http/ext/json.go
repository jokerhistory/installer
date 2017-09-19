package ext

import (
	"encoding/json"
	"fmt"
	"os"
)

type Endpoint struct {
	endpoint string
}

//{"endpoint":"hub.opshub.sh/containerops/test-java-gradle-test:latest"}
func Json2String(endpoint string) {
	s, err := json.Marshal(endpoint)
	if err != nil {
		fmt.Println("error:", err)
	}
	os.Stdout.Write(s)
}

func createYML(url string){

}
