package main

// Como leer un json v1: Agregamos los paquetes: encoding/json; os; io/ioutil

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

// Como leer un json v1: Declaración de estructuras de acuerdo al archivo de ejemplo "users.json"

// Users struct which contains an array of users
type Users struct {
	Users []User `json:"users"`
}

// Estructura de Base de Datos, que contiene un array de BDs
type Dblists struct {
	Dblists []Dblist `json:"db_list"`
}

// User struct which contains a name a type and a list of social links
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

// Social struct which contains a list of links
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

}
