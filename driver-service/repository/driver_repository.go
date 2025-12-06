package repository

import (
	"context"
	"time"

	"github.com/bekiryildirimcode/driver-service/model"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

type mongoRepository struct {
	collection *mongo.Collection
}

func NewMongoRepository(db *mongo.Database) model.DriverRepository {
	return &mongoRepository{
		collection: db.Collection("drivers"),
	}
}

func (r *mongoRepository) Create(driver *model.Driver) (string, error) {
	driver.CreatedAt = time.Now()
	driver.UpdatedAt = time.Now()
	res, err := r.collection.InsertOne(context.Background(), driver)
	if err != nil {
		return "", err
	}

	return res.InsertedID.(bson.ObjectID).Hex(), nil
}

func (r *mongoRepository) Update(id string, updates map[string]interface{}) error {
	oid, _ := bson.ObjectIDFromHex(id)
	updates["updatedAt"] = time.Now()
	filter := bson.M{"_id": oid}
	update := bson.M{"$set": updates}
	_, err := r.collection.UpdateOne(context.Background(), filter, update)
	return err
}

func (r *mongoRepository) List(page, pageSize int) ([]model.Driver, error) {
	skip := int64((page - 1) * pageSize)
	limit := int64(pageSize)
	opts := options.Find().SetSkip(skip).SetLimit(limit)

	cursor, err := r.collection.Find(context.Background(), bson.M{}, opts)
	if err != nil {
		return nil, err
	}
	var drivers []model.Driver
	if err = cursor.All(context.Background(), &drivers); err != nil {
		return nil, err
	}
	return drivers, nil
}

func (r *mongoRepository) FindAllByType(taksiType string) ([]model.Driver, error) {
	filter := bson.M{"taksiType": taksiType}
	cursor, err := r.collection.Find(context.Background(), filter)
	if err != nil {
		return nil, err
	}
	var drivers []model.Driver
	err = cursor.All(context.Background(), &drivers)
	return drivers, err
}
