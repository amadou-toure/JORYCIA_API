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
	"github.com/stripe/stripe-go/v82"
	"github.com/stripe/stripe-go/v82/price"
	"github.com/stripe/stripe-go/v82/product"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func CreateStripeProduct(c *fiber.Ctx, p *models.Product) error {
	
	var imagePointers []*string
	stripe.Key = os.Getenv("STRIPE_SECRET_KEY")
	for _, img := range p.Images {
		imgCopy := img
		imagePointers = append(imagePointers, &imgCopy)
	}

	params := &stripe.ProductParams{
		Name:        stripe.String(p.Name),
		Description: stripe.String(p.Description),
		Images:      imagePointers,
		Metadata:    p.Metadata,
	}

	prod, err := product.New(params)
	if err != nil {
		return c.Status(HTTP_CODE.Server_error).SendString(err.Error())
	}
	newProductPrice:=int64(p.Price*100)
	priceParams := stripe.PriceParams{
		Product:    stripe.String(prod.ID),
		UnitAmount: stripe.Int64(newProductPrice),
		Currency:   stripe.String("cad"),
		Recurring:  nil,
	}
	returnedStripePriceId, err := price.New(&priceParams)
	if err != nil {
		return c.Status(HTTP_CODE.Server_error).SendString(err.Error())
	}
	p.StripePriceID=returnedStripePriceId.ID
	p.StripeProductID=prod.ID
	return nil
}


func AddProduct(c *fiber.Ctx)error{
	var newProduct models.Product
	err := c.BodyParser(&newProduct)
	if err != nil {
		return c.Status(HTTP_CODE.Server_error).SendString(err.Error())
	}
	// Check if product with same name already exists
	existsFilter := bson.D{{Key: "name", Value: newProduct.Name}}
	count, err := Database.Mg.Db.Collection("products").CountDocuments(c.Context(), existsFilter)
	if err != nil {
		return c.Status(HTTP_CODE.Server_error).SendString(err.Error())
	}
	if count > 0 {
		return c.Status(HTTP_CODE.Bad_request).SendString("Product already exists")
	}
	//todo:add a condition to reject request when the image already exist
	for i, item := range newProduct.Images {

		path:= "./Files/Images/"
		filename := strings.ReplaceAll(fmt.Sprintf("%s_%d_%d.webp", newProduct.Name,time.Now().Unix(),i), " ", "")
		fullpath := filepath.Join(path, filename)
		err:=DecodeBase64ToWebP(item,fullpath)
		if err != nil {
			return err  
		}
		newProduct.Images[i] = fmt.Sprintf("%s:%s/image/%s", os.Getenv("API_URL"), os.Getenv("PORT"), filename)
		
	}
	err = CreateStripeProduct(c, &newProduct)
	if err != nil {
		return c.Status(HTTP_CODE.Server_error).SendString(err.Error())
	}
	//todo:add a condition to reject request when the product already exist
	result,err:=Database.Mg.Db.Collection("products").InsertOne(c.Context(),newProduct)
	if err != nil {
		return c.Status(HTTP_CODE.Server_error).SendString(err.Error())
	}
	
	return c.Status(HTTP_CODE.Created).SendString(newProduct.Name + " created with id " + result.InsertedID.(primitive.ObjectID).Hex())
 
}

func GetProducts(c *fiber.Ctx)error{
	filter:=bson.D{{}}
	var products []models.Product
	 result,err := Database.Mg.Db.Collection("products").Find(c.Context(),filter)
	 if err != nil {
		if err == mongo.ErrNoDocuments {
			return c.Status(HTTP_CODE.Not_found).SendString("items not found")
		}
		return c.Status(HTTP_CODE.Server_error).SendString(err.Error())
	 }
	 err = result.All(c.Context(),&products)
	 if err != nil {
		return c.Status(HTTP_CODE.Server_error).SendString(err.Error())
	 }
	 return c.Status(HTTP_CODE.Ok).JSON(products)

}

func GetOneProduct(c *fiber.Ctx) error{
	id:= c.Params("id")
	ID,err:= primitive.ObjectIDFromHex(id)
	if err != nil{
		return c.Status(HTTP_CODE.Bad_request).SendString("unvalid id")
	}
	filter:= bson.D{{Key: "_id",Value: ID}}
	product:=new(models.Product)
	query:= Database.Mg.Db.Collection("products").FindOne(c.Context(),filter)
	if query.Err() != nil{
		if query.Err() == mongo.ErrNoDocuments{
		return c.Status(HTTP_CODE.Not_found).SendString("Products not found")
		}
		return c.Status(HTTP_CODE.Server_error).SendString(query.Err().Error())
	}
	err=query.Decode(product)
	if err != nil{
		if err == mongo.ErrNoDocuments{
			return c.Status(HTTP_CODE.Not_found).SendString("Products not found 2")
		}
	}
	return c.Status(HTTP_CODE.Ok).JSON(product)

}





func DeleteProduct(c *fiber.Ctx)error{
	id:=c.Params("id")
	productId,err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return c.Status(HTTP_CODE.Bad_request).SendString("invalid id")
	}
	filter:=bson.D{{Key:"_id",Value: productId}}
	err = Database.Mg.Db.Collection("products").FindOneAndDelete(c.Context(),filter).Err()
	if err != nil{
		if err == mongo.ErrNoDocuments{
			return c.Status(HTTP_CODE.Not_found).SendString("product not found")
		}
		return c.Status(HTTP_CODE.Server_error).SendString(err.Error())
	}
	return c.Status(HTTP_CODE.Ok).SendString("user deleted")


}