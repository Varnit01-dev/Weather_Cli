package main

import (
	// "flag"
	"fmt"
	// "io"
	"encoding/json"
	"net/http"
)

func main() {

	// cityPtr := flag.String("city", "New York", "an integer")
	// fmt.Scanf("%s", &cityPtr)
	// flag.Parse()

	url := "https://weatherapi-com.p.rapidapi.com/forecast.json?q=New%20%20Delhi&days=3"

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("X-RapidAPI-Key", "d595689a79msh0ba5899a5cb0525p11a42bjsn6a7fa8a37025")
	req.Header.Add("X-RapidAPI-Host", "weatherapi-com.p.rapidapi.com")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	// fmt.Println(res)
	// fmt.Println(string(body))

	var weatherData map[string]interface{}
	if err := json.NewDecoder(res.Body).Decode(&weatherData); err != nil {
		fmt.Println("Error decoding JSON response:", err)
		return
	}
	location, ok := weatherData["location"].(map[string]interface{})
	if ok {
		city, ok := location["name"].(string)
		if ok {
			fmt.Printf("City: %s\n", city)
		}
	}

	current, ok := weatherData["current"].(map[string]interface{})
	if ok {
		tempC, ok := current["temp_c"].(float64)
		if ok {
			fmt.Printf("Temperature (Celsius): %.1f\n", tempC)
		}
	}
	condition, ok := current["condition"].(map[string]interface{})
	if ok {
		text, ok := condition["text"].(string)
		if ok {
			fmt.Printf("Condition: %s\n", text)
		}
		windkph, ok := condition["windkph"].(float64)
		if ok {
			fmt.Printf("Wind (kph): %.1f\n", windkph)
		}
	}
}
