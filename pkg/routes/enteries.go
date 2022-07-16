package routes

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"www.github.com/iZarrios/calorie-tracker-api/pkg/db"
	"www.github.com/iZarrios/calorie-tracker-api/pkg/models"
)

var entryCollection *mongo.Collection = db.OpenCollection(db.Client, "calories")

func Ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})

}

func AddEntry(c *gin.Context) {
	var ctx, cancel = context.WithTimeout(context.Background(), time.Second*100)
	defer cancel()

	var entry models.Entry

	err := c.BindJSON(&entry)
	if err != nil {
		c.JSON(http.StatusInternalServerError, bson.M{
			"error": err.Error(),
		})
		fmt.Println(err)
		return
	}
	validate := validator.New()
	errValidator := validate.Struct(entry)

	if errValidator != nil {
		c.JSON(http.StatusInternalServerError, bson.M{
			"error": err,
		})
		fmt.Println(errValidator)
		return

	}
	entry.ID = primitive.NewObjectID()
	res, insertErr := entryCollection.InsertOne(ctx, entry)
	if insertErr != nil {
		c.JSON(http.StatusInternalServerError, bson.M{
			"error": "entry not created",
		})
		fmt.Println(errValidator)
		return
	}
	fmt.Print(res)

	c.JSON(http.StatusOK, res)

}
func GetEntries(c *gin.Context) {
	var ctx, cancel = context.WithTimeout(context.Background(), time.Second*100)
	defer cancel()

	var enteries []bson.M

	cursor, err := entryCollection.Find(ctx, bson.M{})

	if err != nil {
		c.JSON(http.StatusInternalServerError, bson.M{
			"error": err.Error(),
		})
		fmt.Println(err)
		return

	}
	if err = cursor.All(ctx, &enteries); err != nil {
		c.JSON(http.StatusInternalServerError, bson.M{
			"error": err.Error(),
			"hi":    "here",
		})
		fmt.Println(err)
		return
	}
	fmt.Println(enteries)
	c.JSON(http.StatusOK, enteries)

}

func DeleteEntry(c *gin.Context) {

	entryID := c.Params.ByName("id")
	docID, _ := primitive.ObjectIDFromHex(entryID)

	var ctx, cancel = context.WithTimeout(context.Background(), time.Second*100)

	defer cancel()

	res, err := entryCollection.DeleteOne(ctx, bson.M{"_id": docID})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		fmt.Println(err)
		return
	}
	c.JSON(http.StatusOK, res.DeletedCount)
}

func GetEntryByID(c *gin.Context) {
	var entry bson.M

	entryID := c.Params.ByName("id")

	docID, _ := primitive.ObjectIDFromHex(entryID)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*100)
	defer cancel()

	err := entryCollection.FindOne(ctx, bson.M{"_id": docID}).Decode(&entry)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		fmt.Println(err)
		return
	}

	fmt.Println(entry)
	c.JSON(http.StatusOK, entry)
	return
}

func UpdateEntry(c *gin.Context) {
	entryID := c.Params.ByName("id")
	docID, _ := primitive.ObjectIDFromHex(entryID)
	var ctx, cancel = context.WithTimeout(context.Background(), time.Second*100)
	defer cancel()

	var entry models.Entry

	err := c.BindJSON(&entry)
	if err != nil {
		c.JSON(http.StatusInternalServerError, bson.M{
			"error": err.Error(),
		})
		fmt.Println(err)
		return
	}
	validate := validator.New()
	errValidator := validate.Struct(entry)

	if errValidator != nil {
		c.JSON(http.StatusInternalServerError, bson.M{
			"error": err,
		})
		fmt.Println(errValidator)
		return

	}
	res, replaceErr := entryCollection.ReplaceOne(
		ctx,
		bson.M{"_id": docID},
		// ? entry,
		bson.M{
			"dish":        entry.Dish,
			"fat":         entry.Fat,
			"ingredients": entry.Ingredients,
			"calories":    entry.Calories,
		},
	)
	if replaceErr != nil {
		c.JSON(http.StatusInternalServerError, bson.M{
			"error": replaceErr.Error(),
		})
		fmt.Println(errValidator)
		return
	}
	fmt.Print(res)

	c.JSON(http.StatusOK, res)

}
func UpdateIngredient(c *gin.Context) {
	entryID := c.Params.ByName("id")
	docID, _ := primitive.ObjectIDFromHex(entryID)
	var ctx, cancel = context.WithTimeout(context.Background(), time.Second*100)
	defer cancel()

	type m struct {
		me *string `json:"ingredient"`
	}

	var model m
	err := c.BindJSON(&model)
	if err != nil {
		c.JSON(http.StatusInternalServerError, bson.M{
			"error": err.Error(),
		})
		fmt.Println(err)
		return
	}
	validate := validator.New()
	errValidator := validate.Struct(model)

	if errValidator != nil {
		c.JSON(http.StatusInternalServerError, bson.M{
			"error": err,
		})
		fmt.Println(errValidator)
		return

	}
	res, updateErr := entryCollection.UpdateOne(
		ctx,
		bson.M{"_id": docID},
		bson.D{{"$set", bson.D{{"ingredient", model.me}}}},
	)
	if updateErr != nil {
		c.JSON(http.StatusInternalServerError, bson.M{
			"error": updateErr.Error(),
		})
		fmt.Println(errValidator)
		return
	}
	fmt.Print(res)

	c.JSON(http.StatusOK, res.ModifiedCount)

}
func GetEntryByIngredient(c *gin.Context) {
	ingredient := c.Params.ByName("id")

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*100)
	defer cancel()

	var enteries []bson.M

	cursor, err := entryCollection.Find(ctx, bson.M{"ingredient": ingredient})

	if err != nil {
		c.JSON(http.StatusInternalServerError, bson.M{
			"error": err.Error(),
		})
		fmt.Println(err)
		return
	}
	if err = cursor.All(ctx, &enteries); err != nil {
		c.JSON(http.StatusInternalServerError, bson.M{
			"error": err.Error(),
		})
		fmt.Println(err)
		return

	}
	fmt.Print(enteries)
	c.JSON(http.StatusOK, enteries)

}
