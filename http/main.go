
package main  
  
import (  
	"bytes"
	"os"	
	// "encoding/json"  
	//"mime/multipart"	
    "log"  
    "net/http"  
	"fmt"  
	//"io"
	"flag"
	
)  

var tag ="latest"
var registry="hub.opshub.sh"
var namespace ="containerops"
var api_url ="https://build.opshub.sh/assembling/build?"
func main() {
	//url := "https://build.opshub.sh/assembling/build?image=test-java-gradle-testng&tag=latest&registry=hub.opshub.sh&namespace=containerops"


    //port := flag.String("port", ":8080", "http listen port")
    var image string
    flag.StringVar(&image, "image", "test_image", "image name")
 
    flag.Parse()
 
  //  fmt.Println("port:", *port)
	fmt.Println("image:", image)
	 imagename:= image
	 url := api_url
     buf := bytes.NewBufferString(url)
	 buf.Write([]byte("image="))
	 buf.Write([]byte(imagename))
	 buf.Write([]byte("&tag="))
	 buf.Write([]byte(tag))
	 buf.Write([]byte("&tag="))
	 buf.Write([]byte(tag))
	 buf.Write([]byte("&registry="))
	 buf.Write([]byte(registry))
	 buf.Write([]byte("&namespace="))
	 buf.Write([]byte(namespace))
	 fmt.Println(buf.String()) //hello roc
		UploadBinaryFile("./coala.tar",buf.String())
	}
// Upload binary file to the Dockyard service.
func UploadBinaryFile(filePath, url string) error {
		
	if f, err := os.Open(filePath); err != nil {
		return err
	} else {
		defer f.Close()
		if req, err := http.NewRequest(http.MethodPost,
			url, f); err != nil {
			return err
		} else {
			req.Header.Set("Content-Type", "text/plain")

			client := &http.Client{}
			if resp, err := client.Do(req); err != nil {
				return err
			} else {
				defer resp.Body.Close()

				switch resp.StatusCode {
				case http.StatusOK:
					{
						body := &bytes.Buffer{}
						_, err := body.ReadFrom(resp.Body)
					 	 if err != nil {
							log.Fatal(err)
						}
					 	resp.Body.Close()
						fmt.Println(resp.StatusCode)
						fmt.Println(resp.Header)
						fmt.Println(body)
						return nil						
					}
				case http.StatusBadRequest:
					return fmt.Errorf("Binary upload failed.")
				case http.StatusUnauthorized:
					return fmt.Errorf("Action unauthorized.")
				default:
					return fmt.Errorf("Unknown error.")
				}
			
			}
		}
	}

	return nil
}
