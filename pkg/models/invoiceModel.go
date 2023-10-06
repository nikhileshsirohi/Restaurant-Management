package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Invoice struct {
	ID               primitive.ObjectID `bson:"_id"`
	Invoice_id       string             `bson:"invoice_id"`
	Order_id         string             `bson:"order_id"`
	Payment_method   *string            `bson:"payment_method" validate:"eq=CARD|eq=CASH|eq="`
	Payment_status   *string            `bson:"payment_status" validate:"eq=PENDING|eq=PAID"`
	Payment_due_date time.Time          `bson:"payment_due_date"`
	Created_at       time.Time          `bson:"created_at"`
	Updated_at       time.Time          `bson:"updated_at"`
}
