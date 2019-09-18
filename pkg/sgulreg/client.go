package sgulreg

// import (
// 	"bytes"
// 	"context"
// 	"encoding/json"
// 	"io/ioutil"
// 	"net"
// 	"net/http"
// 	"time"

// 	"go.uber.org/zap"
// )

// // ServiceRegistryClient .
// type ServiceRegistryClient interface {
// 	Register(ctx context.Context, request sgulreg.ServiceRegistrationRequest) (sgulreg.ServiceRegistrationResponse, error)
// }

// type serviceRegistryClient struct {
// 	httpClient   *http.Client
// 	logger       *zap.SugaredLogger
// 	registryHost string
// }

// // NewServiceRegistryClient .
// func NewServiceRegistryClient(registryHost string, l *zap.SugaredLogger) ServiceRegistryClient {
// 	return &serviceRegistryClient{
// 		registryHost: registryHost,
// 		logger:       l,
// 		httpClient: &http.Client{
// 			Transport: &http.Transport{
// 				DialContext: (&net.Dialer{
// 					Timeout: 3 * time.Second,
// 				}).DialContext,
// 				TLSHandshakeTimeout:   10 * time.Second,
// 				ExpectContinueTimeout: 4 * time.Second,
// 				ResponseHeaderTimeout: 10 * time.Second,
// 			},
// 			Timeout: 10 * time.Minute,
// 		},
// 	}
// }

// /*
// TO PASS HEADERS:
// url := "http://restapi3.apiary.io/notes"
//     fmt.Println("URL:>", url)

//     var jsonStr = []byte(`{"title":"Buy cheese and bread for breakfast."}`)
//     req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
//     req.Header.Set("X-Custom-Header", "myvalue")
//     req.Header.Set("Content-Type", "application/json")

//     client := &http.Client{}
//     resp, err := client.Do(req)
//     if err != nil {
//         panic(err)
//     }
//     defer resp.Body.Close()

//     fmt.Println("response Status:", resp.Status)
//     fmt.Println("response Headers:", resp.Header)
//     body, _ := ioutil.ReadAll(resp.Body)
//     fmt.Println("response Body:", string(body))
// */
// // Register .
// func (c *serviceRegistryClient) Register(ctx context.Context, request ServiceRegistrationRequest) (ServiceRegistrationResponse, error) {
// 	jsonRequest, _ := json.Marshal(request)
// 	res, err := c.httpClient.Post("http://"+c.registryHost+"/services", "application/json", bytes.NewBuffer(jsonRequest))
// 	if err != nil {
// 		c.logger.Fatalf("Error making HTTP request: %s", err.Error())
// 		return ServiceRegistrationResponse{}, err
// 	}
// 	defer res.Body.Close()
// 	c.logger.Infof("service registration response Status:", res.Status)
// 	c.logger.Debugf("service registration response Headers:", res.Header)
// 	body, _ := ioutil.ReadAll(res.Body)
// 	var srr ServiceRegistrationResponse
// 	json.Unmarshal([]byte(body), &srr)
// 	return srr, nil
// }
