package idl

import (
	"errors"
	"github.com/huangwei021230/api-gateway/hertz-server/biz/model/gateway"
	"os"
	"sync"
)

var (
	serviceMap = make(map[string]gateway.Service)
	mapMutex   = &sync.Mutex{}
	ioMutex    = &sync.Mutex{}
)

func GetService(serviceName string) (*gateway.Service, error) {
	// Acquire lock
	mapMutex.Lock()
	defer mapMutex.Unlock()
	service, ok := serviceMap[serviceName]
	if !ok {
		return nil, errors.New("service not found")
	}
	return &service, nil
}

func GetIdlPath(serviceName string) string {
	service, err := GetService(serviceName)
	if err != nil {
		// Handle the error, e.g., log it or return a default path
		return ""
	}
	return service.Idl
}

// GetIdlContent involves file read/write operations, so it requires locking
func GetIdlContent(serviceName string) (string, error) {
	service, err := GetService(serviceName)
	if err != nil {
		return "", err
	}
	ioMutex.Lock()
	defer ioMutex.Unlock()

	content, err := os.ReadFile(service.Idl)
	if err != nil {
		return "", err
	}
	return string(content), nil
}

func AddService(service gateway.Service) {
	mapMutex.Lock() // Acquire lock
	defer mapMutex.Unlock()
	serviceMap[service.Name] = service
}

func DelService(serviceName string) {
	mapMutex.Lock() // Acquire lock
	defer mapMutex.Unlock()
	delete(serviceMap, serviceName)
}

func EditService(service gateway.Service) {
	mapMutex.Lock() // Acquire lock
	defer mapMutex.Unlock()
	serviceMap[service.Name] = service
}

func ListAllService() []*gateway.Service {
	var services []*gateway.Service
	mapMutex.Lock() // Acquire lock
	defer mapMutex.Unlock()
	for _, service := range serviceMap {
		s := service
		services = append(services, &s)
	}
	return services
}
