package registration

import (
	"V3/packages/api/types"
	"V3/packages/config"
	"V3/packages/database"
	staking "V3/packages/genratedAbis"
	jwttokengenrator "V3/packages/jwtTokenGenrator"
	"context"
	"log"
	"net/http"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

const (
	userAlreadyExists        = "email already linked with another account "
	phoneNumberAlreadyExists = "contact number already linked with an account"
	signedUp                 = "Success! Account Created"
	signIn                   = "Sucessfully signed In"
)

func HashPassword(password string) string {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		log.Panic(err)
	}
	return string(bytes)
}

func VerifyPassword(userPassword string, givenPassword string) (bool, string) {
	err := bcrypt.CompareHashAndPassword([]byte(givenPassword), []byte(userPassword))
	valid := true
	msg := ""
	if err != nil {
		msg = "login or Password is incorrect"
		valid = false
	}
	return valid, msg
}

var Validate = validator.New()

func SignUp() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()
		var user types.User
		if err := c.BindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		validationErr := Validate.Struct(user)
		if validationErr != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": validationErr})
			return
		}
		userCollection := database.UserData(database.Client, "users")
		count, err := userCollection.CountDocuments(ctx, bson.M{"email": user.Email})
		if err != nil {
			log.Println(err)
			c.JSON(http.StatusInternalServerError, gin.H{
				"code":    http.StatusBadRequest,
				"message": err.Error(),
			})
			return
		}
		if count > 0 {
			c.JSON(http.StatusBadRequest, gin.H{
				"code":    http.StatusBadRequest,
				"message": userAlreadyExists,
			})
			return
		}
		count, err = userCollection.CountDocuments(ctx, bson.M{"phone": user.Phone})
		defer cancel()
		if err != nil {
			log.Println(err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": err})
			return
		}
		if count > 0 {
			c.JSON(http.StatusBadRequest, gin.H{
				"code":    http.StatusBadRequest,
				"message": phoneNumberAlreadyExists,
			})
			return
		}

		contract, err := staking.NewStakingCaller(common.HexToAddress(config.StakingContractAddress), config.EthClient)
		if err != nil {
			log.Println("The error", err)
		}

		userStatus, err := contract.UserProfile(&bind.CallOpts{}, common.HexToAddress(user.UserWalletAddress))
		if err != nil {
			log.Println("The error1", err)
		}

		if common.IsHexAddress(user.UserWalletAddress) && userStatus.HasDeposited {
			count, err = userCollection.CountDocuments(ctx, bson.M{"user_wallet_address": user.UserWalletAddress})
			if err != nil {
				log.Println(err)
				c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Wallet address or is not staked"})
				return
			}
			if count > 0 {
				c.JSON(http.StatusBadRequest, gin.H{
					"code":    http.StatusBadRequest,
					"message": "Invalid Wallet address",
				})
				return
			}

		} else {
			c.JSON(http.StatusBadRequest, gin.H{
				"code":    http.StatusBadRequest,
				"message": "Invalid Wallet address Or Is not staked",
			})
			return
		}
		password := HashPassword(user.Password)
		user.Password = password
		user.CreatedAt, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		user.UpdatedAt, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		user.ID = primitive.NewObjectID()
		user.UserID = user.ID.Hex()
		token, refreshtoken, _ := jwttokengenrator.TokenGenerator(user.Email, user.UserID, user.UserName)
		//user.Token = token
		log.Println(token)
		user.Refresh_Token = refreshtoken

		_, inserterr := userCollection.InsertOne(ctx, user)
		if inserterr != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "not created"})
			return
		}

		c.JSON(http.StatusCreated, gin.H{
			"code":    http.StatusOK,
			"message": signedUp,
		})
	}
}

func Login() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()
		var user types.User
		var foundUser types.User
		if err := c.BindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err})
			return
		}
		userCollection := database.UserData(database.Client, "users")
		err := userCollection.FindOne(ctx, bson.M{"email": user.Email}).Decode(&foundUser)
		defer cancel()
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "login Or Password incorrect"})
			return
		}
		passwordIsValid, msg := VerifyPassword(user.Password, foundUser.Password)
		defer cancel()
		if !passwordIsValid {
			c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
			return
		}
		token, refreshToken, _ := jwttokengenrator.TokenGenerator(foundUser.Email, foundUser.UserName, foundUser.UserID)
		defer cancel()
		jwttokengenrator.UpdateAllTokens(token, refreshToken, foundUser.UserID)
		c.JSON(http.StatusFound, foundUser)
	}

}
