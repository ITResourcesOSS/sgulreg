package sgulreg

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net"
	"net/http"
	"time"

	"go.uber.org/zap"
)

// ServiceRegistryClient .
type ServiceRegistryClient interface {
	Register(request ServiceRegistrationRequest) (ServiceRegistrationResponse, error)
}

type serviceRegistryClient struct {
	httpClient   *http.Client
	logger       *zap.SugaredLogger
	registryHost string
}

// NewServiceRegistryClient .
func NewServiceRegistryClient(registryHost string, l *zap.SugaredLogger) ServiceRegistryClient {
	return &serviceRegistryClient{
		registryHost: registryHost,
		logger:       l,
		httpClient: &http.Client{
			Transport: &http.Transport{
				DialContext: (&net.Dialer{
					Timeout: 3 * time.Second,
				}).DialContext,
				TLSHandshakeTimeout:   10 * time.Second,
				ExpectContinueTimeout: 4 * time.Second,
				ResponseHeaderTimeout: 10 * time.Second,
			},
			Timeout: 10 * time.Minute,
		},
	}
}

func (c *serviceRegistryClient) Register(request ServiceRegistrationRequest) (ServiceRegistrationResponse, error) {
	jsonRequest, _ := json.Marshal(request)
	res, err := c.httpClient.Post("http://"+c.registryHost+"/services", "application/json", bytes.NewBuffer(jsonRequest))
	if err != nil {
		c.logger.Fatalf("Error making HTTP request: %s", err.Error())
		return ServiceRegistrationResponse{}, err
	}
	defer res.Body.Close()
	c.logger.Infof("service registration response Status:", res.Status)
	c.logger.Debugf("service registration response Headers:", res.Header)
	body, _ := ioutil.ReadAll(res.Body)
	var srr ServiceRegistrationResponse
	json.Unmarshal([]byte(body), &srr)
	return srr, nil
}
