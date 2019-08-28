package test

import (
	"bytes"
	"encoding/json"
	"flag"
	"github.com/sirupsen/logrus"
	"github.com/ypapax/logrus_conf"
	"io/ioutil"
	"net/http"
	"os"
	"testing"
	"time"
)

var serviceAddr string
var dockerComposeConfigFile string

const reqTimeout = time.Second

func TestMain(m *testing.M) {
	if err := logrus_conf.Files("comment_test", logrus.TraceLevel); err != nil {
		logrus.Error(err)
		os.Exit(1)
	}
	flag.StringVar(&serviceAddr, "api", "http://localhost:3002", "address of comment api web service")
	flag.StringVar(&dockerComposeConfigFile, "docker-compose", "", "docker compose config file")
	flag.Parse()

	ret := m.Run()
	os.Exit(ret)
}

func pathReq(path string, method string, bodyStruct interface{}) (int, []byte, time.Duration, error) {
	return req(serviceAddr+path, method, bodyStruct)
}

func req(u, method string, bodyStruct interface{}) (int, []byte, time.Duration, error) {
	b, err := json.Marshal(bodyStruct)
	logrus.Printf("requesting: curl -X%+v  -d'%+v' %+v", method, string(b), u)
	req, err := http.NewRequest(method, u, bytes.NewBuffer(b))
	if err != nil {
		logrus.Error(err)
		return 0, nil, 0, err
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	t1 := time.Now()
	resp, err := client.Do(req)
	if err != nil {
		logrus.Error(err)
		return 0, nil, 0, err
	}
	respTime := time.Since(t1)
	defer resp.Body.Close()
	b, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return resp.StatusCode, nil, respTime, err
	}
	logrus.Tracef("resp: %+v", string(b))
	return resp.StatusCode, b, respTime, nil
}
