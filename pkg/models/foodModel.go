package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Food struct {
	ID         primitive.ObjectID `bson:"_id"`
	Name       *string            `bson:"name" validate:"required, min=2, max=50"`
	Price      *float64           `bson:"price" validate:"required"`
	Food_image *string            `bson:"food_image" validate:"required"`
	Created_at time.Time          `bson:"created_at"`
	Updated_at time.Time          `bson:"updated_at"`
	Food_id    string             `bson:"food_id"`
	Menu_id    *string            `bson:"menu_id"`
}
