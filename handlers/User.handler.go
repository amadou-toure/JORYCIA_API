package handlers

import (
	"jorycia_api/Database"
	"jorycia_api/HTTP_CODE"
	"jorycia_api/models"
	"jorycia_api/utils"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/stripe/stripe-go/v82"
	"github.com/stripe/stripe-go/v82/customer"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func CreateUser(c *fiber.Ctx)error{

	var newUser models.User
	err := c.BodyParser(&newUser)
	if err != nil {
		return c.Status(HTTP_CODE.Server_error).SendString(err.Error())
	}
	hashPassword,err := utils.HashPassword(*newUser.Password)
	if err != nil{
		return c.Status(HTTP_CODE.Server_error).SendString("error crypting the password")
	}
	newUser.Password = &hashPassword
	newUser.ID = ""
	result,err:=Database.Mg.Db.Collection("Users").InsertOne(c.Context(),newUser)
	if err != nil {
		return c.Status(HTTP_CODE.Server_error).SendString(err.Error())
	}
	if newUser.Role == "user"{
		stripe.Key = os.Getenv("STRIPE_SECRET_KEY")
		customerParams := &stripe.CustomerParams{
			Email: stripe.String(newUser.Email),
		}
		customer,err := customer.New(customerParams)
		if err != nil {
			return c.Status(HTTP_CODE.Server_error).SendString("error creating stripe customer")
		}
		newUser.StripeCustomerID = &customer.ID
	}
	
	
	return c.Status(HTTP_CODE.Created).SendString("user " + newUser.FirstName + " created with id " + result.InsertedID.(primitive.ObjectID).Hex())
 
}

func Login(c *fiber.Ctx) error {
	body := new(models.User)
	user := new(models.User)
	err:=c.BodyParser(body)
	if err != nil {
		return c.Status(HTTP_CODE.Bad_request).SendString(err.Error())
	}
	filter:=bson.D{{Key:"email",Value: body.Email}}
	err= Database.Mg.Db.Collection("Users").FindOne(c.Context(),filter).Decode(&user)
	
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return c.Status(HTTP_CODE.Not_found).JSON("no user found with this email")
		}
		return c.Status(HTTP_CODE.Server_error).SendString(err.Error())
	}
 	passwordIsCorrect:=utils.CompareHashedPassword(*body.Password, *user.Password)
	if !passwordIsCorrect {
    return c.Status(HTTP_CODE.Bad_request).SendString("Wrong password!, try again")
}
	token,err:=utils.Token.GenerateToken(user.ID)
	if err != nil {
		return c.Status(HTTP_CODE.Server_error).SendString("Erreur génération token")
	}
	return c.Status(HTTP_CODE.Accepted).JSON(fiber.Map{
		"user": user,
		"token": token,
	})
}

func GetUsers(c *fiber.Ctx)error{
	filter:=bson.D{{}}
	var users []models.User
	 result,err := Database.Mg.Db.Collection("Users").Find(c.Context(),filter)
	 if err != nil {
		if err == mongo.ErrNoDocuments {
			return c.Status(HTTP_CODE.Not_found).SendString("user not found")
		}
		return c.Status(HTTP_CODE.Server_error).SendString(err.Error())
	 }
	 err = result.All(c.Context(),&users)
	 if err != nil {
		return c.Status(HTTP_CODE.Server_error).SendString(err.Error())
	 }
	 return c.Status(HTTP_CODE.Ok).JSON(users)

}
func GetOneUser(c *fiber.Ctx) error{
	id:= c.Params("id")
	UserID,err:= primitive.ObjectIDFromHex(id)
	if err != nil{
		return c.Status(HTTP_CODE.Bad_request).SendString("unvalid User id")
	}
	filter:= bson.D{{Key: "_id",Value: UserID,}}
	user:=new(models.User)
	query:= Database.Mg.Db.Collection("Users").FindOne(c.Context(),filter)
	if query.Err() != nil{
		if query.Err() == mongo.ErrNoDocuments{
		return c.Status(HTTP_CODE.Not_found).SendString("User not found")
		}
		return c.Status(HTTP_CODE.Server_error).SendString(query.Err().Error())
	}
	err=query.Decode(user)
	if err != nil{
		if err == mongo.ErrNoDocuments{
			return c.Status(HTTP_CODE.Not_found).SendString("User not found 2")
		}
	}
	return c.Status(HTTP_CODE.Ok).JSON(user)

}

func UpdateUser(c *fiber.Ctx)error{
id:=c.Params("id")
userID,err:= primitive.ObjectIDFromHex(id)
if err != nil {
	return c.Status(HTTP_CODE.Bad_request).SendString("invalid id")
}
user:= new(models.User)
err = c.BodyParser(user)
if err != nil {
	return c.Status(HTTP_CODE.Bad_request).SendString(err.Error())
}
filter:=bson.D{{Key:"_id",Value: userID}}
update:=bson.D{{Key:"$set",Value: bson.D{
	{Key: "FirstName", Value: user.FirstName},
{Key: "LastName", Value: user.LastName},
{Key: "Email", Value: user.Email},
{Key: "Password", Value: user.Password},
{Key: "Phone", Value: user.Phone},
{Key: "Address", Value: user.Address},
{Key: "CreatedAt", Value: user.CreatedAt},
{Key: "UpdatedAt", Value: user.UpdatedAt},

}}}
err = Database.Mg.Db.Collection("Users").FindOneAndUpdate(c.Context(),filter,update).Err()
if err != nil {
	if err == mongo.ErrNoDocuments {
		return c.Status(HTTP_CODE.Not_found).SendString("user not found")
	}
	return c.Status(HTTP_CODE.Server_error).SendString(err.Error())
}
return c.Status(HTTP_CODE.Ok).JSON(user)
}

func DeleteUser(c *fiber.Ctx)error{
	id:=c.Params("id")
	userId,err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return c.Status(HTTP_CODE.Bad_request).SendString("invalid id")
	}
	filter:=bson.D{{Key:"_id",Value: userId}}
	err = Database.Mg.Db.Collection("Users").FindOneAndDelete(c.Context(),filter).Err()
	if err != nil{
		if err == mongo.ErrNoDocuments{
			return c.Status(HTTP_CODE.Not_found).SendString("user not foound")
		}
		return c.Status(HTTP_CODE.Server_error).SendString(err.Error())
	}
	return c.Status(HTTP_CODE.Ok).SendString("user deleted")


}

// SaveOrUpdateOAuthToken inserts or updates an OAuthToken in the MongoDB collection.
func SaveOrUpdateOAuthToken(c *fiber.Ctx, token models.OAuthToken) error {
	filter := bson.M{"user_id": token.UserID, "provider": token.Provider}
	update := bson.M{
		"$set": bson.M{
			"access_token":  token.AccessToken,
			"refresh_token": token.RefreshToken,
			"expiry_date":   token.ExpiryDate,
			"updated_at":    time.Now(),
		},
		"$setOnInsert": bson.M{
			"created_at": time.Now(),
		},
	}
	opts := options.Update().SetUpsert(true)

	_, err := Database.Mg.Db.Collection("OAuthTokens").UpdateOne(c.Context(), filter, update, opts)
	if err != nil {
		return c.Status(HTTP_CODE.Server_error).SendString("failed to store OAuth token: " + err.Error())
	}

	return nil
}