package routers

import (
	"github.com/gin-gonic/gin"
)

func (r *DefaultRouter) smsRouter(gr *gin.RouterGroup) {
	ptexter := gr.Group("/sms")
	{
		ptexter.POST("", r.handler.SMSStoreLufft)
	}
}
