package servermanager

import (
	"net/http"

	"github.com/sirupsen/logrus"
	"gopkg.in/yaml.v3"
)

type TrafficConfigHandler struct {
	*BaseHandler

	store Store
}

type trafficConfigTemplateVar struct {
	BaseTemplateVars
	Config string
}

type TrafficConfig struct {
	Config string
}

func (tch *TrafficConfigHandler) view(w http.ResponseWriter, r *http.Request) {

	config := "hehe"

	tch.viewRenderer.MustLoadTemplate(w, r, "traffic-config/index.html", &trafficConfigTemplateVar{
		Config: config,
	})
}

func (tch *TrafficConfigHandler) save(w http.ResponseWriter, r *http.Request) {
	config := TrafficConfig{
		"",
	}
	err := DecodeFormData(config, r)

	if err != nil {
		logrus.WithError(err).Error("couldn't fetch traffic config")
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	err = config.saveConfig()

	if err != nil {
		logrus.WithError(err).Errorf("couldn't save config")
		AddErrorFlash(w, r, "Failed to save config options")
	} else {
		AddFlash(w, r, "Traffic config successfully saved!")
	}

	tch.viewRenderer.MustLoadTemplate(w, r, "traffic-config/index.html", &trafficConfigTemplateVar{
		Config: config.Config,
	})
}

func (tc *TrafficConfig) saveConfig() error {
	var t 
	error := yaml.Unmarshal([]byte(tc.Config), &t)
	return
}

func NewTrafficConfigHandler(baseHandler *BaseHandler, store Store) *TrafficConfigHandler {
	return &TrafficConfigHandler{
		BaseHandler: baseHandler,
		store:       store,
	}
}
