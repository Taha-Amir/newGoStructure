package routes

import (
	addtostate "V3/packages/api/addToState"
	"V3/packages/api/registration"

	"github.com/gin-gonic/gin"
)

func Routes(incommingRoutes *gin.Engine) {

	incommingRoutes.POST("user/signup", registration.SignUp())
	incommingRoutes.GET("user/login", registration.Login())
	incommingRoutes.POST("/addOrder", addtostate.AddToState())

}
