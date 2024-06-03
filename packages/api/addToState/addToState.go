package addtostate

import (
	"V3/packages/api/types"
	"V3/packages/database"
	"context"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func AddToState() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		var reqBody map[string]interface{}
		if err := ctx.ShouldBindJSON(&reqBody); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
			return
		}
		pledgeBaseAmount := reqBody["orderAmount"].(string)
		pledgeQuoteAmount := reqBody["totalPrice"].(string)

		limitOrder := types.Order{
			ID:          primitive.NewObjectID(),
			OrderAmount: pledgeBaseAmount,
			TotalPrice:  pledgeQuoteAmount,
		}
		limitOrderCollection := database.OrderData(database.Client, "limitOrders")
		result, err := limitOrderCollection.InsertOne(context.Background(), limitOrder)
		if err != nil {
			log.Println(err)
			ctx.JSON(http.StatusBadGateway, gin.H{"error": "Invalid request body"})
			return
		}
		limitOrderId := result.InsertedID.(primitive.ObjectID).Hex()
		log.Println(limitOrderId)
		ctx.JSON(http.StatusOK, result)
	}

}
