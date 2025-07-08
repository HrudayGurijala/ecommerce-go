package token

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/HrudayGurijala/ecommerce-go/database"
	"github.com/golang-jwt/jwt/v5"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type SignedDetails struct {
	Email      string
	First_Name string
	Last_Name  string
	Uid        string
	jwt.RegisteredClaims
}

var UserData *mongo.Collection = database.UserData(database.Client, "Users")
var SECRET_KEY = os.Getenv("SECRET_LOVE")

func TokenGenerator(email, firstname, lastname, uid string) (signedtoken string, signedrefreshtoken string, err error) {
	claims := &SignedDetails{
		Email:      email,
		First_Name: firstname,
		Last_Name:  lastname,
		Uid:        uid,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
		},
	}

	refreshclaims := &SignedDetails{
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(168 * time.Hour)), // 7 days
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedtoken, err = token.SignedString([]byte(SECRET_KEY))
	if err != nil {
		return "", "", err
	}

	refreshtoken := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshclaims)
	signedrefreshtoken, err = refreshtoken.SignedString([]byte(SECRET_KEY))
	if err != nil {
		log.Panicln(err)
		return
	}

	return signedtoken, signedrefreshtoken, nil
}

func ValidateToken(signedtoken string) (claims *SignedDetails, msg string) {
	token, err := jwt.ParseWithClaims(signedtoken, &SignedDetails{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(SECRET_KEY), nil
	})

	if err != nil {
		msg = err.Error()
		return nil, msg
	}

	claims, ok := token.Claims.(*SignedDetails)
	if !ok || !token.Valid {
		msg = "the token is invalid"
		return nil, msg
	}

	if claims.ExpiresAt.Time.Before(time.Now()) {
		msg = "token is expired"
		return nil, msg
	}

	return claims, ""
}

func UpdateAllTokens(signedtoken, signedrefreshtoken, userid string) {
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()

	updateobj := bson.D{
		{Key: "token", Value: signedtoken},
		{Key: "refresh_token", Value: signedrefreshtoken},
		{Key: "updatedat", Value: time.Now()},
	}

	filter := bson.M{"user_id": userid}
	upsert := true //update/insert = upsert
	opt := options.UpdateOptions{Upsert: &upsert}

	_, err := UserData.UpdateOne(ctx, filter, bson.D{
		{Key: "$set", Value: updateobj},
	}, &opt)

	if err != nil {
		log.Panic(err)
	}
}
