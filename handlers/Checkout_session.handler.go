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
		//AutomaticTax: &stripe.CheckoutSessionAutomaticTaxParams{
        //Enabled: stripe.Bool(true),},
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

type StripeCheckoutSessionStatus struct {
	Status stripe.CheckoutSessionStatus `json:"status"`
	PaymentStatus stripe.PaymentIntentStatus `json:"payment_status"`
}

func GetCheckoutSession(c *fiber.Ctx) error {
	fmt.Println("Geting CheckoutSession ...")
	stripe.Key = os.Getenv("STRIPE_SECRET_KEY")
	id := c.Params("session_id")
	fmt.Println("id:",id)
	// Retrieve the Checkout Session
	sess, err := session.Get(id, nil)
	if err != nil {
		fmt.Println("Stripe session retrieval failed: ", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Session retrieval failed"})
	}
	fmt.Println(stripe.CheckoutSessionStatus(sess.Status))
	fmt.Println(stripe.CheckoutSessionStatus(sess.PaymentStatus))
	return c.JSON(StripeCheckoutSessionStatus{
		Status: stripe.CheckoutSessionStatus(sess.Status),
		PaymentStatus: stripe.PaymentIntentStatus(sess.PaymentStatus),
	})
}
