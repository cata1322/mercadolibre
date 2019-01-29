package main

// Como leer un json v1: Agregamos los paquetes: encoding/json; os; io/ioutil

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

// Como leer un json v1: Declaración de estructuras de acuerdo al archivo de ejemplo "users.json"

// Example Users struct which contains an array of users
type Users struct {
	Users []User `json:"users"`
}

// Estructura de Base de Datos, que contiene un array de BDs
type Dblists struct {
	Dblists []Dblist `json:"db_list"`
}

// Example User struct which contains a name a type and a list of social links
type User struct {
	Name   string `json:"name"`
	Type   string `json:"type"`
	Age    int    `json:"Age"`
	Social Social `json:"social"`
}

//MERCADOLIBRE Un base de datos contiene el campo Nombre,
//y una lista para los campos Clasificación, Owner y Tiempo-
type Dblist struct {
	Name            string           `json:"dn_name"`
	Classifications []Classification `json:"classification"`
	Owners          []Owner          `json:"owner"`
	Time            string           `json:"time_stamp"`
}

// Example Social struct which contains a list of links
type Social struct {
	Facebook string `json:"facebook"`
	Twitter  string `json:"twitter"`
}

// MERCADO LIBRE Declaración de la estructura Clasificacion
type Classification struct {
	Confidentiality string `json:"confidentiality"`
	Integrity       string `json:"integrity"`
	Availability    string `json:"availability"`
}

// MERCADO LIBRE - Declaración de la estructura Owner
type Owner struct {
	Name           string `json:"name"`
	Identification string `json:"uid"`
	Email          string `json:"email"`
}

func main() {

	fmt.Printf("Hola mundo cruel_Parate_1\n")

	// Open our jsonFile
	jsonFile, err := os.Open("dblist.json")
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
		fmt.Printf("Error_Atencion_1\n")
	}

	fmt.Printf("Hola mundo cruel_Parate_2\n")

	fmt.Println("Successfully Opened deblist.json")
	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var result map[string]interface{}

	json.Unmarshal([]byte(byteValue), &result)

	fmt.Println(result["db_list"])

	fmt.Printf("Hola mundo cruel_Parate_3\n")

	// ---------Mercado Libre Lectura mejorada => Error, no funciona------------

	// var dblists Dblists
	// json.Unmarshal(byteValue, &dblists)
	// for i := 0; i < len(dblists.Dblists); i++ {
	//fmt.Println("DB_Name: " + dblists.Dblists[i].Name)
	//fmt.Println("DB_Time: " + dblists.Dblists[i].Time)
	//fmt.Println("Confidentiality: " + dblists.Dblists[i].Classifications[i].Confidentiality)
	//fmt.Println("Integrity: " + dblists.Dblists[i].Classifications[i].Integrity)
	//fmt.Println("Availability: " + dblists.Dblists[i].Classifications[i].Availability)
	//fmt.Println("Owners Name: " + dblists.Dblists[i].Owners[i].Name)
	//fmt.Println("Owners Identification: " + dblists.Dblists[i].Owners[i].Identification)
	//fmt.Println("Owners Email: " + dblists.Dblists[i].Owners[i].Email)
	// ---------Mercado Libre Lectura mejorada => Error, no funciona -----------

	// ---------Mercado Libre - Conexión a la base de datos --------------

	fmt.Println("Establecienco conección MySQL - Base: mercadolibre")

	// Open up our database connection.
	// I've set up a database on my local machine using phpmyadmin.
	// The database is called mercadolibre
	db, err := sql.Open("mysql", "root:admin0000@tcp(127.0.0.1:3306)/mercadolibre")

	// if there is an error opening the connection, handle it
	if err != nil {
		panic(err.Error())
		fmt.Println("ERROR en conección MySQL - Base: mercadolibre")
	}

	// defer the close till after the main function has finished
	// executing
	fmt.Println("Conección MySQL exitosa")

	defer db.Close()

	// perform a db.Query insert
	insert, err := db.Query("INSERT INTO Test VALUES ('ConexionTest')")

	// if there is an error inserting, handle it
	if err != nil {
		panic(err.Error())
	}
	// be careful deferring Queries if you are using transactions
	fmt.Println("Texto insertado correctamente")
	defer insert.Close()

}
