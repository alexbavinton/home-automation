package service

import (
	"embed"
	"html/template"
	"log"
	"net/http"
)

// type getDeviceser interface {
// 	GetDevices() ([]client.Device, error)
// }

//go:embed devices.html
var devicesHTML embed.FS

func getDevicesPage() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tmpl, err := template.ParseFS(devicesHTML, "devices.html")
		if err != nil {
			log.Print(err)
		}
		tmpl.ExecuteTemplate(w, "devices.html", nil)
		w.Header().Set("Content-Type", "text/html")
		w.WriteHeader(http.StatusOK)
	}
}
