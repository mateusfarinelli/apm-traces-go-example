package handles

import (
	"apm-trace-with-gorilla-mux-example/env"
	"apm-trace-with-gorilla-mux-example/errs"
	"encoding/json"
	"log/slog"
	"net/http"
)

const (
	name        string = "teste-go"
	description string = "Api criada com o intuito de exemplificar a utilização de libs datadog para trace"
)

type ResponseBodyLayout struct {
	AppInfo AppInfo `json:"app_info"`
	Data    any     `json: "data"`
}

type AppInfo struct {
	Env         string `json:"env"`
	Name        string `json:"name"`
	Version     string `json:"version"`
	Description string `json:"description"`
}

func Success(w http.ResponseWriter, statusCode int, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	
	if statusCode != http.StatusNoContent{		
		responseBody := ResponseBodyLayout{
			Data: data,
			AppInfo: AppInfo{
				Env: env.Environment,
				Name: name,
				Version: env.Version,
				Description: description,
			},
		}
		
		if err := json.NewEncoder(w).Encode(responseBody); err != nil {
			slog.Error("Failed to encode response", slog.String("error", err.Error()))
			Error(w, errs.InternalServerError(err.Error()))
		}
	}
}

func Error(w http.ResponseWriter, err *errs.Err){
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(err.StatusCode)

	if err := json.NewEncoder(w).Encode(err); err != nil {
		slog.Error("Failed to encode response", slog.String("error", err.Error()))
		json.NewEncoder(w).Encode(err.Error())
}
}