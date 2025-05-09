package handlers

import (
	"jorycia_api/Database"
	"jorycia_api/HTTP_CODE"
	"jorycia_api/models"
	"jorycia_api/utils"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func CreateUser(c *fiber.Ctx)error{
	var newUser models.User
	err := c.BodyParser(&newUser)
	if err != nil {
		return c.Status(HTTP_CODE.Server_error).SendString(err.Error())
	}
	hashPassword,err := utils.HashPassword(newUser.Password)
	if err != nil{
		return c.Status(HTTP_CODE.Server_error).SendString("error crypting the password")
	}
	newUser.Password=hashPassword
	newUser.ID = ""
	result,err:=Database.Mg.Db.Collection("Users").InsertOne(c.Context(),newUser)
	if err != nil {
		return c.Status(HTTP_CODE.Server_error).SendString(err.Error())
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
 	passwordIsCorrect:=utils.CompareHashedPassword(body.Password,user.Password)
	if !passwordIsCorrect {
    return c.Status(HTTP_CODE.Bad_request).SendString("Wrong password!, try again")
}

	return c.Status(HTTP_CODE.Accepted).JSON(user)

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
{Key: "Password", Value: user.Password}, // Note: The password should typically not be included in such mappings due to security concerns
{Key: "Phone", Value: user.Phone},
{Key: "Address", Value: user.Address},
{Key: "Orders", Value: user.Orders},
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