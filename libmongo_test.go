package libmongo

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
	"go.mongodb.org/mongo-driver/bson"
)

func TestConnect(t *testing.T) {
	config := DBConfig{
		Host:   "localhost",
		Port:   27017,
		DBName: "jihar",
	}
	db, err := Connect(config)
	require.NoError(t, err)
	name := db.Name()
	t.Log(name)

	coll := db.Collection("post")
	type User struct {
		Name string `json:"name" bson:"name"`
	}
	var data User
	data.Name = "jihar"

	resp, err := coll.InsertOne(context.TODO(), bson.D{{"name", "Alice"}})
	require.NoError(t, err)
	t.Log(resp)
	var result bson.M

	respData := coll.FindOne(context.TODO(), bson.D{})
	errA := respData.Decode(result)
	require.NoError(t, errA)
	t.Log(result)
}
