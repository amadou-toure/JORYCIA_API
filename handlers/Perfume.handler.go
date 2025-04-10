package handlers

import (
	"fmt"
	"jorycia_api/Database"
	"jorycia_api/HTTP_CODE"
	"jorycia_api/models"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func AddPerfume(c *fiber.Ctx)error{
	var newPerfume models.Perfume
	err := c.BodyParser(&newPerfume)
	if err != nil {
		return c.Status(HTTP_CODE.Server_error).SendString(err.Error())
	}
	//todo:add a condition to reject request when the image already exist
	for i, item := range newPerfume.Image {

		path:= "./Files/Images/"
		filename := strings.ReplaceAll(fmt.Sprintf("%s %d %d", newPerfume.Name,time.Now().Unix(),i), " ", "")
		fullpath:= filepath.Join(path,filename)
		err:=DecodeBase64ToWebP(item,fullpath)
		// err := os.WriteFile(fullpath, []byte(item), 0644)
		if err != nil {
			return err  
		}
		newPerfume.Image[i]=os.Getenv("api_url")+"/image/"+filename+".webp"
		
	}

	result,err:=Database.Mg.Db.Collection("perfumes").InsertOne(c.Context(),newPerfume)
	if err != nil {
		return c.Status(HTTP_CODE.Server_error).SendString(err.Error())
	}
	
	return c.Status(HTTP_CODE.Created).SendString(newPerfume.Name + " created with id " + result.InsertedID.(primitive.ObjectID).Hex())
 
}

func GetPerfumes(c *fiber.Ctx)error{
	filter:=bson.D{{}}
	var perfumes []models.Perfume
	 result,err := Database.Mg.Db.Collection("perfumes").Find(c.Context(),filter)
	 if err != nil {
		if err == mongo.ErrNoDocuments {
			return c.Status(HTTP_CODE.Not_found).SendString("items not found")
		}
		return c.Status(HTTP_CODE.Server_error).SendString(err.Error())
	 }
	 err = result.All(c.Context(),&perfumes)
	 if err != nil {
		return c.Status(HTTP_CODE.Server_error).SendString(err.Error())
	 }
	 return c.Status(HTTP_CODE.Ok).JSON(perfumes)

}

func GetOnePerfume(c *fiber.Ctx) error{
	id:= c.Params("id")
	ID,err:= primitive.ObjectIDFromHex(id)
	if err != nil{
		return c.Status(HTTP_CODE.Bad_request).SendString("unvalid id")
	}
	filter:= bson.D{{Key: "_id",Value: ID}}
	perfume:=new(models.Perfume)
	query:= Database.Mg.Db.Collection("perfumes").FindOne(c.Context(),filter)
	if query.Err() != nil{
		if query.Err() == mongo.ErrNoDocuments{
		return c.Status(HTTP_CODE.Not_found).SendString("Perfumes not found")
		}
		return c.Status(HTTP_CODE.Server_error).SendString(query.Err().Error())
	}
	err=query.Decode(perfume)
	if err != nil{
		if err == mongo.ErrNoDocuments{
			return c.Status(HTTP_CODE.Not_found).SendString("Perfumes not found 2")
		}
	}
	return c.Status(HTTP_CODE.Ok).JSON(perfume)

}



func DeletePerfume(c *fiber.Ctx)error{
	id:=c.Params("id")
	perfumeId,err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return c.Status(HTTP_CODE.Bad_request).SendString("invalid id")
	}
	filter:=bson.D{{Key:"_id",Value: perfumeId}}
	err = Database.Mg.Db.Collection("perfumes").FindOneAndDelete(c.Context(),filter).Err()
	if err != nil{
		if err == mongo.ErrNoDocuments{
			return c.Status(HTTP_CODE.Not_found).SendString("perfume not foound")
		}
		return c.Status(HTTP_CODE.Server_error).SendString(err.Error())
	}
	return c.Status(HTTP_CODE.Ok).SendString("user deleted")


}