package handler

import (
	"session-14/service"
)

type TemplateHandler struct {
	Service service.ServiceTemplate
}

func NewTemplateHandler(service service.ServiceTemplate) *TemplateHandler {
	return &TemplateHandler{
		Service: service,
	}
}
