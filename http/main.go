
package main  
  
import (  
	"bytes"
	"os"	
	// "encoding/json"  
	"mime/multipart"	
    "log"  
    "net/http"  
	"fmt"  
	"io"
//	"io/ioutil"

)  


// Upload binary file to the Dockyard service.
//func UploadBinaryFile(filePath, domain, namespace, repository, tag string) error {
func UploadBinaryFile(filePath, url string) error {
		
	if f, err := os.Open(filePath); err != nil {
		return err
	} else {
		defer f.Close()

		// if req, err := http.NewRequest(http.MethodPost,
		// 	fmt.Sprintf("https://%s/binary/v1/%s/%s/binary/%s/%s",
		// 		domain, namespace, repository, tag, filepath.Base(filePath)), f); err != nil {
		// 	return err
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

	//init()
	func init() {
	
	}
	
	//main()
	func main() {
	url := "https://build.opshub.sh/assembling/build?image=test-java-gradle-testng&tag=latest&registry=hub.opshub.sh&namespace=containerops"
	UploadBinaryFile("./checkstyle.tar",url)
		/////////////////////////////////
		// if err := cmd.RootCmd.Execute(); err != nil {
		// 	fmt.Fprintf(os.Stderr, err.Error())
		// 	os.Exit(1)
		// }
//https://build.opshub.sh/assembling/build?image=test-java-gradle-testng&tag=latest&registry=hub.opshub.sh&namespace=containerops
	// extraParams := map[string]string{
	// 	"image":       "testtestng",
	// 	"tag":      "latest",
	// 	"registry": "hub.opshub.sh",
	// 	"namespace":"containerops",
	// }
	//request, err := newfileUploadRequest("https://google.com/upload", extraParams, "file", "/tmp/doc.pdf")	
	//url := "https://build.opshub.sh/assembling/build"
	//url := "https://build.opshub.sh/assembling/build?image=test-java-gradle-testng&tag=latest&registry=hub.opshub.sh&namespace=containerops"
	
	// resp ,err :=newfileUploadRequest(url,extraParams,"file", "checkstyle.tar")

	// defer resp.Body.Close()
	// 		b, err := ioutil.ReadAll(resp.Body)
	// 		if err != nil {
	// 			log.Println("http.Do failed,[err=%s][url=%s]", err, url,b)
	// 		}

///////////////////////////////////////////////////////
	// extraParams := map[string]string{
	// 	"image":       "testtestng",
	// 	"tag":      "latest",
	// 	"registry": "hub.opshub.sh",
	// 	"namespace":"containerops",
	// }
	// 	UploadBinaryFile(extraParams);

		/////////////////
			// request, err := newfileUploadRequest(url, extraParams, "file","./cpd.tar")
			// if err != nil {
			// 	log.Fatal(err)
			// }
			// client := &http.Client{}
			// resp, err := client.Do(request)
			// if err != nil {
			// 	log.Fatal(err)
			// } else {
			// 	body := &bytes.Buffer{}
			// 	_, err := body.ReadFrom(resp.Body)
			//   if err != nil {
			// 		log.Fatal(err)
			// 	}
			//   resp.Body.Close()
			// 	fmt.Println(resp.StatusCode)
			// 	fmt.Println(resp.Header)
			// 	fmt.Println(body)
			// }
	}


	func newfileUploadhttp(params map[string]string){
		
		target_url := "https://build.opshub.sh/assembling/build"
		
	//	target_url := "http://localhost:9200/prod/aws"
		body_buf := bytes.NewBufferString("")
		body_writer := multipart.NewWriter(body_buf)
		jsonfile := "cpd.tar"
		file_writer, err := body_writer.CreateFormFile("upfile", jsonfile)

		for key, val := range params {
			_ = body_writer.WriteField(key, val)
		}

		if err != nil {
		fmt.Println("error writing to buffer")
		return
		}
		fh, err := os.Open(jsonfile)
		if err != nil {
		fmt.Println("error opening file")
		return
		}
		io.Copy(file_writer, fh)
		body_writer.Close()
		//http.Post(target_url, "application/json", body_buf)
		
		resp, err := http.Post(target_url, "application/x-tar", body_buf)
		if err != nil {
			log.Fatal(err)
		} else {
			body := &bytes.Buffer{}
			_, err := body.ReadFrom(resp.Body)
		  if err != nil {
				log.Fatal(err)
			}
		  resp.Body.Close()
			fmt.Println(resp.StatusCode)
			fmt.Println(resp.Header)
			fmt.Println(body)
		}

		return 
	}
	

// func DoBytesPost(url string, data []byte) ([]byte, error) {
	
// 		body := bytes.NewReader(data)
// 		request, err := http.NewRequest(POST_METHOD, url, body)
// 		if err != nil {
// 			log.Println("http.NewRequest,[err=%s][url=%s]", err, url)
// 			return []byte(""), err
// 		}
// 		request.Header.Set("Connection", "Keep-Alive")
// 		var resp *http.Response
// 		resp, err = http.DefaultClient.Do(request)
// 		if err != nil {
// 			log.Println("http.Do failed,[err=%s][url=%s]", err, url)
// 			return []byte(""), err
// 		}
// 		defer resp.Body.Close()
// 		b, err := ioutil.ReadAll(resp.Body)
// 		if err != nil {
// 			log.Println("http.Do failed,[err=%s][url=%s]", err, url)
// 		}
// 		return b, err
// 	}
func newfileUploadRequest(uri string, params map[string]string, paramName, path string) (*http.Request, error) {
    file, err := os.Open(path)
    if err != nil {
        return nil, err
    }
    defer file.Close()

    body := &bytes.Buffer{}
    writer := multipart.NewWriter(body)
    part, err := writer.CreateFormFile(paramName, path)
    if err != nil {
        return nil, err
    }
    _, err = io.Copy(part, file)

    for key, val := range params {
        _ = writer.WriteField(key, val)
    }
    err = writer.Close()
    if err != nil {
        return nil, err
    }
    request, err := http.NewRequest("POST", uri, body)
	//request.Header.Set("Content-Type", writer.FormDataContentType())
	// request.Header.Set("Content-Type", writer.FormDataContentType())
	
	
    return request, err
}
