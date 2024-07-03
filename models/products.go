package models

import (
	"context"
	"log"
	"time"

	"github.com/anhhuy1010/customer-menu/database"
	"go.mongodb.org/mongo-driver/mongo"

	//"go.mongodb.org/mongo-driver/bson"

	"github.com/anhhuy1010/customer-menu/constant"
)

type Products struct {
	Uuid        string    `json:"uuid,omitempty" bson:"uuid"`
	Price       int       `json:"price,omitempty" bson:"price"`
	Image       string    `json:"image" bson:"image"`
	Name        string    `json:"name,omitempty" bson:"name"`
	Sequence    int       `json:"sequence" bson:"sequence"`
	Quantity    int       `json:"quantity" bson:"quantity"`
	Description string    `json:"description" bson:"description"`
	Gallery     []string  `json:"gallery" bson:"gallery"`
	IsActive    int       `json:"is_active" bson:"is_active"`
	IsDelete    int       `json:"is_delete" bson:"is_delete"`
	StartDate   time.Time `json:"start_date" bson:"start_date"`
	EndDate     time.Time `json:"end_date" bson:"end_date  "`
	CreatedAt   time.Time `json:"created_at" bson:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" bson:"updated_at"`
}

func (u *Products) Model() *mongo.Collection {
	db := database.GetInstance()
	return db.Collection("products")
}
func (u *Products) Pagination(ctx context.Context, conditions map[string]interface{}, modelOptions ...ModelOption) ([]*Products, error) {
	coll := u.Model()

	conditions["is_delete"] = constant.UNDELETE

	modelOpt := ModelOption{}
	findOptions := modelOpt.GetOption(modelOptions)
	cursor, err := coll.Find(context.TODO(), conditions, findOptions)
	if err != nil {
		return nil, err
	}

	var product []*Products
	for cursor.Next(context.TODO()) {
		var elem Products
		err := cursor.Decode(&elem)
		if err != nil {
			log.Println("[Decode] PopularCuisine:", err)
			log.Println("-> #", elem.Uuid)
			continue
		}

		product = append(product, &elem)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}
	_ = cursor.Close(context.TODO())

	return product, nil
}

func (u *Products) FindOne(conditions map[string]interface{}) (*Products, error) {
	coll := u.Model()

	conditions["is_delete"] = constant.UNDELETE
	err := coll.FindOne(context.TODO(), conditions).Decode(&u)
	if err != nil {
		return nil, err
	}

	return u, nil
}
func (u *Products) Count(ctx context.Context, condition map[string]interface{}) (int64, error) {
	coll := u.Model()

	condition["is_delete"] = constant.UNDELETE

	total, err := coll.CountDocuments(ctx, condition)
	if err != nil {
		return 0, err
	}

	return total, nil
}
