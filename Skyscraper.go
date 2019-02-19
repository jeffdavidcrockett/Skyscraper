package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	. "github.com/logrusorgru/aurora"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

var url = "http://api.openweathermap.org/data/2.5/weather?"

type weatherAll struct {
	Weather Weather
	Main Main
	Wind Wind
	Name string `json:"name"`
}

type Weather []struct {
	Description string `json:"description"`
}

type Main struct {
	Temp float64 `json:"temp"`
	Humidity float64 `json:"humidity"`
}

type Wind struct {
	Speed float64 `json:"speed"`
}

func checkErr(e error) {
	if e != nil {
		panic(e)
	}
}

func removeNewLine(s string) string {
	return strings.TrimRight(s, "\n")
}

func byZipcode(key string, zip string) {
	zipurl := url + "zip=" + zip + key + "&units=imperial"
	getWeather(zipurl)
}

func byCity(key string, city string) {
	cityurl := url + "q=" + city + key + "&units=imperial"
	getWeather(cityurl)
}

func getWeather(fullurl string) {
	weatherClient := http.Client {
		Timeout: time.Second * 2,
	}

	req, err4 := http.NewRequest(http.MethodGet, fullurl, nil)
	if err4 != nil {
		log.Fatal(err4)
	}

	res, getErr := weatherClient.Do(req)
	if getErr != nil {
		log.Fatal(getErr)
	}

	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}

	weatherData := weatherAll{}
	jsonErr := json.Unmarshal(body, &weatherData)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	fmt.Println(Cyan("\nCity:"), weatherData.Name)
	fmt.Println(Cyan("Current conditions:"), weatherData.Weather[0].Description)
	fmt.Println(Cyan("Temperature:"), weatherData.Main.Temp, "F")
	fmt.Println(Cyan("Humidity:"), weatherData.Main.Humidity,"%")
	fmt.Println(Cyan("Wind speed:"), weatherData.Wind.Speed, "MPH")
}

func main() {
	fmt.Println("\n")
	fmt.Println("           .-~~~-.")
	fmt.Println("   .- ~ ~-(       )_ _")
    fmt.Println(" /                     ~ -.")
	fmt.Print("|        ")
	fmt.Print(Cyan("Skyscraper"))
	fmt.Println("         \\")
	fmt.Println(" \\                         .'")
	fmt.Println("   ~- . _____________ . -~")

	reader := bufio.NewReader(os.Stdin)
	apiKeyByte, err := ioutil.ReadFile("api_key.txt")
	checkErr(err)
	apiKey := "&appid=" + string(apiKeyByte)

	for {
		fmt.Println("\nMain Menu")
		fmt.Println("----------\n")
		fmt.Println(`1) By City
2) By Zipcode`)
		fmt.Print("> ")
		rawChoice, err := reader.ReadString('\n')
		checkErr(err)
		choice := removeNewLine(rawChoice)

		if choice == "1" {
			fmt.Print("\nCity: ")
			rawCity, err := reader.ReadString('\n')
			checkErr(err)
			city := removeNewLine(rawCity)
			byCity(apiKey, city)
		} else if choice == "2" {
			fmt.Print("\nZipcode: ")
			rawZip, err := reader.ReadString('\n')
			checkErr(err)
			zip := removeNewLine(rawZip)
			byZipcode(apiKey, zip)
		}
	}
}
