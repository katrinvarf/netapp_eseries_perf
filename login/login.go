package login

import(
	"strconv"
	"net/http"
	"encoding/json"
	"io/ioutil"
	"crypto/tls"
	"errors"
	"github.com/sirupsen/logrus"
	"time"
)

func Login(log *logrus.Logger, Username string, Password string, Address string, Port int, Timeout int) error{
	urlString := "https://" + Address + ":" + strconv.Itoa(Port) + "/devmgr/v2/storage-systems"
	//отключение проверки безопасности для client
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr, Timeout: time.Duration(Timeout) * time.Second}
	request, err := http.NewRequest("GET", urlString, nil)
	if err!=nil {
		log.Warning("Failed to create new http request: Error: ", err)
		return err
	}

	request.SetBasicAuth(Username, Password)
	resp, err := client.Do(request)
	if err!=nil {
		log.Warning("Failed to do client request: Error: ", err)
		return err
	}

	defer resp.Body.Close()

	var buf []byte
	buf, err = ioutil.ReadAll(resp.Body)
	if err!=nil{
		log.Warning("Failed to read response body: Error: ", err)
		return err
	}
	var raw interface{}
	json.Unmarshal(buf, &raw)
	if raw==nil{
		err = errors.New("login: wrong username or password")
		log.Warning("Failed to authorize user: Error: ", err)
		return err
	}
	log.Debug("Successful login; address: ", Address)
	return nil
}
