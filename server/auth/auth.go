package auth

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/yusrililhm/to-do-app/server/database"
	"github.com/yusrililhm/to-do-app/server/models"
	"go.mongodb.org/mongo-driver/bson"
	"golang.org/x/crypto/bcrypt"
)

func AddUser(w http.ResponseWriter, r *http.Request)  {
	db, err := database.ConnectMongo()
	if err != nil {
		log.Fatal(err.Error())
	}
	username := "timeng" // r.FormValue("username") 
	password := "timeng123" // r.FormValue("password")
	
	securePwd, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Fatal(err.Error())
	}

	_, err = db.Collection("user").InsertOne(context.Background(), models.User{Username: username, Password: string(securePwd)})
	if err != nil {
		log.Fatal(err.Error())
	}
	http.Redirect(w, r, "/show", http.StatusSeeOther)
}

func ShowUser(w http.ResponseWriter, r *http.Request)  {
	db, err := database.ConnectMongo()
	if err != nil {
		log.Fatal(err.Error())
	}

	docs, err := db.Collection("user").Find(context.Background(), bson.D{{}})
	if err != nil {
		log.Fatal(err.Error())
	}

	var results []models.User
	if docs.All(context.Background(), &results); err != nil {
		log.Fatal(err.Error())
	}

	for _, result := range results {
		docs.Decode(&result)
		ouput, err := json.MarshalIndent(result, "", "	")
		if err != nil {
			log.Fatal(err)
		}
		fmt.Fprintf(w, "%s\n", ouput)
	}
}

func SignIn(w http.ResponseWriter, r *http.Request)  {
	db, err := database.ConnectMongo()
	if err != nil {
		log.Fatal(err.Error())
	}
	username := "timeng" // r.FormValue("username")
	password := "$2a$10$MlI3ps/HsPxvV7fhW9Nst..Ex1eSmYHiQpHwMqyfH.SzdsPE2XkKW" // r.FormValue("password")

	// var (userName, pwd string)
	
	_, err = db.Collection("user").Find(context.Background(), models.User{Username: username, Password: password})
	if err != nil {
		log.Fatal(err.Error())
	}

	// err = bcrypt.CompareHashAndPassword([]byte(pwd), []byte(password))

	// check password incorrect
	if err != nil {
		// http.Redirect(w, r, "/", http.StatusSeeOther)
		fmt.Fprintln(w, "Password Incorrect")
		return
	}

	// check user not found
	if err != nil {
		// http.Redirect(w, r, "/", http.StatusSeeOther)
		fmt.Fprintln(w, "User Not Found")
		return
	}

	fmt.Fprintln(w,"Login Successfull")
	// http.Redirect(w, r, "/", http.StatusSeeOther)
}