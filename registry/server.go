package registry

import (
	"encoding/json"
	"log"
	"net/http"
	"sync"
)

const ServicePort = ":3000"
const ServicesURL = "http://localhost" + ServicePort + "/services"

type registry struct {
	registrations []Registration
	mutex         *sync.Mutex
}

func (r *registry) add(reg Registration) error {
	r.mutex.Lock()
	r.registrations = append(r.registrations, reg)
	r.mutex.Unlock()
	return nil
}

// 建立一个包级的registry变量
var reg = registry{
	registrations: make([]Registration, 0),
	mutex:         new(sync.Mutex),
}

type RegistryService struct{}

func (s RegistryService) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Println("Request received")
	switch r.Method {
	case http.MethodPost:
		dec := json.NewDecoder(r.Body)
		var r Registration
		err := dec.Decode(&r)
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		log.Printf("Adding service: %v with URL: %s\n", r.ServiceName, r.ServiceURL)
		err = reg.add(r)
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}
