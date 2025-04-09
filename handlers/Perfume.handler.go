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

	for i, item := range newPerfume.Image {

		path:= "./Files/Images/"
		filename := strings.ReplaceAll(fmt.Sprintf("%s %d %d"+".png", newPerfume.Name,time.Now().Unix(),i), " ", "")
		fullpath:= filepath.Join(path,filename)
		err := os.WriteFile(fullpath, []byte(item), 0644)
		if err != nil {
			return err  
		}
		newPerfume.Image[i]=fullpath
	}

	result,err:=Database.Mg.Db.Collection("perfumes").InsertOne(c.Context(),newPerfume)
	if err != nil {
		return c.Status(HTTP_CODE.Server_error).SendString(err.Error())
	}
	
	return c.Status(HTTP_CODE.Created).SendString(newPerfume.Name + " created with id " + result.InsertedID.(primitive.ObjectID).Hex())
 
}


// func UploadOneFile(c *fiber.Ctx, File_name string,File_type string) error {
    
//     File, err := c.FormFile(File_type)
// 	if err != nil {
//     	return c.Status(HTTP_CODE.Server_error).SendString("no file uploaded:" + err.Error())
// 	}
// 	File.Filename=File_name
// 	err = c.SaveFile(File,"./Files/"+File_type+"/"+File.Filename)
// 	if err != nil {
// 		return c.Status(HTTP_CODE.Server_error).SendString(err.Error())
// 	}
//     return nil
// }



func DeletePerfume(c *fiber.Ctx)error{
	id:=c.Params("id")
	perfumeId,err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return c.Status(HTTP_CODE.Bad_request).SendString("invalid id")
	}
	filter:=bson.D{{Key:"_id",Value: perfumeId}}
	err = Database.Mg.Db.Collection("Perfumes").FindOneAndDelete(c.Context(),filter).Err()
	if err != nil{
		if err == mongo.ErrNoDocuments{
			return c.Status(HTTP_CODE.Not_found).SendString("perfume not foound")
		}
		return c.Status(HTTP_CODE.Server_error).SendString(err.Error())
	}
	return c.Status(HTTP_CODE.Ok).SendString("user deleted")


}