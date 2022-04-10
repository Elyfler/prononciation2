package stores

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/prononciation2/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoCityRepository struct {
	Cities *mongo.Collection
}

type mongoCity struct {
	ID            primitive.ObjectID `bson:"_id,omitempty"`
	Name          string             `bson:"name"`
	PostCode      string             `bson:"post_code"`
	Prononciation string             `bson:"prononciation"`
}

func toMongoCity(c models.City) (mc mongoCity, err error) {
	if c.Exists() {
		mc.ID, err = primitive.ObjectIDFromHex(c.ID)
		if err != nil {
			return mongoCity{}, err
		}
	}
	mc.Name = c.Name
	mc.PostCode = c.PostCode
	mc.Prononciation = c.Prononciation
	return mc, nil
}

func (mc *mongoCity) toCity() models.City {
	return models.City{
		ID:            mc.ID.Hex(),
		Name:          mc.Name,
		PostCode:      mc.PostCode,
		Prononciation: mc.Prononciation,
	}
}

func NewMongoDB(dbName string) (*mongo.Database, error) {
	var uri string
	_, mongoExists := os.LookupEnv("MONGO_URI")
	if mongoExists {
		uri = os.Getenv("MONGO_URI")
	} else {
		uri = "mongodb://127.0.0.1:27017"
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		log.Print(err)
		return nil, err
	}
	db := client.Database(dbName)
	return db, nil
}

// Gros doute sur une utilisation avec plusieurs repo qui essaient de taper la DB en même temps
// Probablement passer une db en paramètre de NewMongoCityRepo pour pallier ce souci
// NewMongoCityRepo ...
func NewMongoCityRepo(db *mongo.Database) *MongoCityRepository {
	return &MongoCityRepository{Cities: db.Collection("cities")}
}

func (r MongoCityRepository) CreateCity(ctx context.Context, c models.City) (models.City, error) {
	mongoCity, err := toMongoCity(c)
	if err != nil {
		return models.City{}, err
	}
	res, err := r.Cities.InsertOne(ctx, c)
	if err != nil {
		return models.City{}, err
	}
	mongoCity.ID = res.InsertedID.(primitive.ObjectID)
	return mongoCity.toCity(), nil
}

func (r MongoCityRepository) GetCityByID(ctx context.Context, id string) (models.City, error) {
	var mongoCity mongoCity
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return models.City{}, err
	}
	filter := bson.M{"_id": objectID}
	err = r.Cities.FindOne(ctx, filter).Decode(&mongoCity)
	if err != nil {
		return models.City{}, err
	}
	return mongoCity.toCity(), nil
}

func (r MongoCityRepository) GetCities(ctx context.Context) ([]models.City, error) {
	var mongoCities []mongoCity
	var cities []models.City
	cur, err := r.Cities.Find(ctx, nil)
	if err != nil {
		return nil, err
	}
	defer cur.Close(ctx)

	if err := cur.All(ctx, &mongoCities); err != nil {
		return nil, err
	}
	for _, mongoCity := range mongoCities {
		cities = append(cities, mongoCity.toCity())
	}
	return cities, nil
}

func (r MongoCityRepository) DeleteCity(ctx context.Context, id string) error {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	filter := bson.M{"_id": objectID}

	_, err = r.Cities.DeleteOne(ctx, filter)
	if err != nil {
		return err
	}
	return nil
}

func (r MongoCityRepository) UpdateCity(ctx context.Context, c models.City) (models.City, error) {
	mongoCity, err := toMongoCity(c)
	if err != nil {
		return models.City{}, err
	}
	filter := bson.M{"_id": mongoCity.ID}

	res, err := r.Cities.UpdateOne(ctx, filter, &mongoCity)
	if err != nil {
		return models.City{}, err
	}
	mongoCity.ID = res.UpsertedID.(primitive.ObjectID)
	return mongoCity.toCity(), nil
}
