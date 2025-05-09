package handlers

import (
	"jorycia_api/HTTP_CODE"

	"github.com/gofiber/fiber/v2"
)

//function to download the images from the file name
func GetImage(c *fiber.Ctx) error {
    fileName := c.Params("fileName")
    if fileName == ""{
        return c.Status(HTTP_CODE.Bad_request).SendString("fileName parameter is required")
    }
    filePath := "./Files/Images/"+fileName
    err := c.SendFile(filePath)
    if err != nil {
        if err == fiber.ErrNotFound {
            return c.Status(HTTP_CODE.Not_found).SendString("image not found")
        }
        return c.Status(HTTP_CODE.Server_error).SendString(err.Error())
    }
    return nil
}

