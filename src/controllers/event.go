package controllers

import (
	"net/http"
	"testProj/src/service"

	"github.com/sirupsen/logrus"
)

// @Summary handle event
// @Description handles and saves event
// @ID HandleEvent
// @Param Event body models.Event true "event"
// @Success 200
// @Router /api/event [post]
func HandleEvent(w http.ResponseWriter, r *http.Request) {
	if err := service.HandleEvent(r.Body); err != nil {
		logrus.Error(err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	} else {
		w.WriteHeader(http.StatusOK)
	}
}
