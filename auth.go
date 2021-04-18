package main

import (
  "fmt"
  "strings"
  "net/http"
  "io/ioutil"
  "crypto/tls"
  "crypto/x509"
)

const(
	client_id = ""
	client_secret = ""
)

func main() {

  url := "https://api-pix.gerencianet.com.br/oauth/token"
  method := "POST"

  payload := strings.NewReader(`{"grant_type": "client_credentials"}`)

  caCert, _ := ioutil.ReadFile("")
  caCertPool := x509.NewCertPool()
  caCertPool.AppendCertsFromPEM(caCert)

  cert, _ := tls.LoadX509KeyPair("", "")

  client := &http.Client{
    Transport: &http.Transport{
        TLSClientConfig: &tls.Config{
            RootCAs: caCertPool,
            Certificates: []tls.Certificate{cert},
			InsecureSkipVerify : true,
        },
    },
}

  

  req, err := http.NewRequest(method, url, payload)

  if err != nil {
    fmt.Println(err)
    return
  }
  req.SetBasicAuth(client_id, client_secret)
  req.Header.Add("Content-Type", "application/json")

  res, err := client.Do(req)
  if err != nil {
    fmt.Println(err)
    return
  }
  defer res.Body.Close()

  body, err := ioutil.ReadAll(res.Body)
  if err != nil {
    fmt.Println(err)
    return
  }
  fmt.Println(string(body))
}