package record

import (
	"aquarium/config"
	"context"
	"encoding/json"
	"log"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type collection struct {
	Measurements []config.Record
}

func saveMeasurementRecords(req *http.Request) error {
	bs := make([]byte, req.ContentLength)
	req.Body.Read(bs)

	var records collection

	if err := json.Unmarshal(bs, &records); err != nil {
		return err
	}

	if err := config.MassInsertRecord("water", records.Measurements); err != nil {
		log.Println("ERROR: could not save the records in Influx", err)
		return err
	}
	log.Println("Records are inserted succesfully")

	return nil
}

type event struct {
	Name     string
	Date     string
	Note     string
	Category uint8
}

func saveEvent(req *http.Request) error {
	bs := make([]byte, req.ContentLength)
	req.Body.Read(bs)

	var e event

	if err := json.Unmarshal(bs, &e); err != nil {
		log.Println("Unmarshal event body failed")
		return err
	}

	log.Println("Getting the event collection")
	eventsCollection := config.MongoDB.Database("aquarium_db").Collection("events")

	log.Println("Insert event to MongodDB")

	_, err := eventsCollection.InsertOne(context.TODO(), e)
	if err != nil {
		log.Println("Inserting event to MongoDB failed")
		return err
	}

	return nil
}

func findEvents() ([]*event, error) {
	eventsCollection := config.MongoDB.Database("aquarium_db").Collection("events")
	cur, err := eventsCollection.Find(context.TODO(), bson.D{{}}, options.Find())
	if err != nil {
		log.Println("Failed to execute Find operation in MongoDB events collection")
		return nil, err
	}

	defer cur.Close(context.TODO())

	var snaps []*event

	for cur.Next(context.TODO()) {
		var elem event
		err := cur.Decode(&elem)
		if err != nil {
			log.Println("Decode Fail")
			return nil, err
		}

		snaps = append(snaps, &elem)
	}

	return snaps, nil
}
