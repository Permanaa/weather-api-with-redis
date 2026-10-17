package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/permanaa/weather-api-with-redis/internal/helper"
	"github.com/permanaa/weather-api-with-redis/internal/model"
)

func (h *Handler) GetTimeline(w http.ResponseWriter, r *http.Request) {
	location := r.PathValue("location")

	client := &http.Client{
		Timeout: 5 * time.Second,
	}

	url := fmt.Sprintf("%s/timeline/%s?unitGroup=us&key=%s&include=days,hours", h.cfg.BaseUrl, location, h.cfg.ApiKey)

	request, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		helper.JsonError(w, http.StatusInternalServerError, err.Error())
		return
	}

	res, err := client.Do(request)
	if err != nil {
		helper.JsonError(w, http.StatusInternalServerError, err.Error())
		return
	}
	defer res.Body.Close()

	var timeline model.Timeline
	if err := json.NewDecoder(res.Body).Decode(&timeline); err != nil {
		helper.JsonError(w, http.StatusInternalServerError, err.Error())
	}

	helper.JsonSuccess(w, http.StatusOK, timeline)
}
