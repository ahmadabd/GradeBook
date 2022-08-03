package register

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"sync"
)

const ServicePort = ":3000"
const ServiceURL = "http://localhost" + ServicePort + "/services"

type registery struct {
	registration []Registration
	mutex        *sync.RWMutex
}

func (r *registery) add(reg Registration) error {
	r.mutex.Lock()
	r.registration = append(r.registration, reg)
	r.mutex.Unlock()
	err := r.sendRequiredServices(reg)
	if err != nil {
		return err
	}
	return nil
}

func (r registery) sendRequiredServices(reg Registration) error {
	r.mutex.RLock()
	defer r.mutex.RUnlock()

	var p patch
	for _, serviceReg := range r.registration {
		for _, reqService := range reg.RequiredServices {
			if serviceReg.ServiceName == reqService {
				p.Added = append(p.Added, patchEntry{
					Name: serviceReg.ServiceName,
					URL: serviceReg.ServiceURL,
				})
			}
		}
	}

	err := r.sendPatch(p, reg.ServiceUpdateURL)
	if err != nil {
		return err
	}

	return nil
}

func (r registery) sendPatch(p patch, url string) error {
	d, err := json.Marshal(p)
	if err != nil {
		return err
	}
	_, err = http.Post(url, "application/json", bytes.NewBuffer(d))
	if err != nil {
		return err
	}

	return nil
}

func (r *registery) remove(url string) error {
	for i := range r.registration {
		if r.registration[i].ServiceURL == url {
			r.mutex.Lock()
			r.registration = append(r.registration[:i], r.registration[i+1:]...)
			r.mutex.Unlock()
			return nil
		}
	}

	return fmt.Errorf("service not found url to remove: %v", url)
}

var reg = registery{
	registration: make([]Registration, 0),
	mutex:        new(sync.RWMutex),
}

type RegisteryService struct{}

func (s RegisteryService) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Println("Request Recived")

	switch r.Method {
	case http.MethodPost:

		var regist Registration
		err := json.NewDecoder(r.Body).Decode(&regist)

		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		err = reg.add(regist)
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		log.Printf("Adding service %v with url %v\n", regist.ServiceName, regist.ServiceURL)

	case http.MethodDelete:
		payload, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		url := string(payload)
		log.Println("Removing service at url: ", url)
		err = reg.remove(url)
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
}
