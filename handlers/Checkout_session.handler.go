package handlers

import (
	"fmt"
	"jorycia_api/models"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/stripe/stripe-go/v82"
	"github.com/stripe/stripe-go/v82/checkout/session"
)

func CreateCheckoutSession(c *fiber.Ctx) error {
	var newCart  []models.CartItem
	err := c.BodyParser(&newCart)
	if err!= nil {
		
		return c.Status(400).SendString(  "Invalid request body: " + err.Error())
	}
	
	stripe.Key = os.Getenv("STRIPE_SECRET_KEY")

	params := &stripe.CheckoutSessionParams{
		PaymentMethodTypes: stripe.StringSlice([]string{"card"}), 
		Mode:               stripe.String(string(stripe.CheckoutSessionModePayment)),
		SuccessURL:         stripe.String(os.Getenv("FRONTEND_URL") + "/success?session_id={CHECKOUT_SESSION_ID}"),
		CancelURL:          stripe.String(os.Getenv("FRONTEND_URL") + "/cancel"),
	}

	for _, item := range newCart {
		params.LineItems = append(params.LineItems, &stripe.CheckoutSessionLineItemParams{
			Price:    stripe.String(item.Product.StripePriceID),
			Quantity: stripe.Int64(int64(item.Quantity)),
		})
	}
	

	// Create the Checkout Session
	sess, err := session.New(params)
	if err != nil {
		fmt.Println("Stripe session creation failed: ", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Session creation failed"})
	}

	return c.JSON(sess.ID)
}

