package login

import(
	"fmt"
	"strconv"
	"net/http"
	"encoding/json"
	"io/ioutil"
	"crypto/tls"
)

func Login(Username string, Password string, Address string, Port int)(result bool){
	urlString := "https://" + Address + ":" + strconv.Itoa(Port) + "/devmgr/v2/storage-systems"
	//отключение проверки безопасности для client
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}

	request, err := http.NewRequest("GET", urlString, nil)
	if err!=nil {
		fmt.Println(err)
	}
	//Username = "ghjf"
	request.SetBasicAuth(Username, Password)
	resp, err := client.Do(request)
	if err!=nil {
		fmt.Println(err)
	}

	defer resp.Body.Close()

	var buf []byte
	buf, err = ioutil.ReadAll(resp.Body)
	var raw interface{}
	json.Unmarshal(buf, &raw)
	if raw==nil{
		//fmt.Println("Error Unauthorized")
		result = true
	}
	return
	//fmt.Println(string(buf))
	//fmt.Println(raw.([]interface{})[0].(map[string]interface{})["id"])

}
