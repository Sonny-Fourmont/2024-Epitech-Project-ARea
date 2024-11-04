package controllers

import (
	"area/config"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func AboutHandler(c *gin.Context) (int, map[string]interface{}) {
	response := map[string]interface{}{
		"client": map[string]string{
			"host": c.ClientIP(),
		},
		"server": map[string]interface{}{
			"current_time": time.Now().Unix(),
			"services":     buildServices(),
		},
	}

	return http.StatusOK, response
}

func buildServices() []map[string]interface{} {
	config.LoadServices()
	var services []map[string]interface{}

	for _, service := range config.AllServices.If {
		services = append(services, map[string]interface{}{
			"name": service.TokenName,
			"actions": []map[string]string{
				{
					"name":        service.Type,
					"description": service.Description,
				},
			},
			"reactions": []map[string]string{},
		})
	}

	for _, service := range config.AllServices.That {
		found := false
		for i, srv := range services {
			if srv["name"] == service.TokenName {
				services[i]["reactions"] = append(
					services[i]["reactions"].([]map[string]string),
					map[string]string{
						"name":        service.Type,
						"description": service.Description,
					},
				)
				found = true
				break
			}
		}
		if !found {
			services = append(services, map[string]interface{}{
				"name":      service.TokenName,
				"actions":   []map[string]string{},
				"reactions": []map[string]string{{"name": service.Type, "description": service.Description}},
			})
		}
	}

	return services
}
