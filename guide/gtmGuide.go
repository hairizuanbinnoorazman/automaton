package main

import (
	"bufio"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"text/template"
	"time"

	"github.com/olekukonko/tablewriter"
)

func main() {
	//fmt.Println(string(NewGuideConfig()))
	GenerateGuide("zzz.json")
}

func csvTests() {
	// data := [][]string{
	// 	[]string{"A", "The Good", "500"},
	// 	[]string{"B", "The Very very Bad Man", "288"},
	// 	[]string{"C", "The Ugly", "120"},
	// 	[]string{"D", "The Gopher", "800"},
	// }

	randomFile, _ := os.Create("testing.md")
	randomWriter := bufio.NewWriter(randomFile)

	nyaaz, _ := os.Open("trial.csv")
	data2, errz := csv.NewReader(nyaaz).ReadAll()
	if errz != nil {
		fmt.Println(errz.Error())
	}

	table := tablewriter.NewWriter(randomWriter)
	table.SetHeader([]string{"Name", "Sign", "Rating"})

	for _, v := range data2 {
		table.Append(v)
	}
	table.Render()

	randomWriter.Flush()
}

// GuideConfig Struct to define all required properties in a GTM Document
type GuideConfig struct {
	OutputFile     string       `json:"outputFile"`
	GtmContainerID string       `json:"gtmContainerID"`
	InitialSetup   initialSetup `json:"initialSetup"`
	Events         events       `json:"events"`
}

type initialSetup struct {
	Include  bool   `json:"include"`
	Template string `json:"template"`
}

type events struct {
	Include       bool   `json:"include"`
	TitleTemplate string `json:"titleTemplate"`
	Template      string `json:"template"`
}

// NewGuideConfig Need to define and initialize the initial guide json values
func NewGuideConfig() []byte {
	newConfig := GuideConfig{
		OutputFile:     fmt.Sprintf("%v.md", time.Now().Format("20060102150405")),
		GtmContainerID: fmt.Sprintf("GTM-%v", "1234567"),
		InitialSetup: initialSetup{
			Include:  true,
			Template: "templates/gtm/initialSetup.md",
		},
		Events: events{
			Include:       true,
			TitleTemplate: "templates/gtm/eventsTitleTemplate.md",
			Template:      "templates/gtm/eventsTemplate.md",
		},
	}
	output, err := json.MarshalIndent(newConfig, "", "\t")
	if err != nil {
		fmt.Println("Issue in parsing the initial config file")
		fmt.Println(err.Error())
	}
	return output
}

// GenerateGuide Function to generate out the markdown file required for the whole GTM Implementation Guide
func GenerateGuide(configFile string) {
	guideData, err := ioutil.ReadFile(configFile)
	if err != nil {
		fmt.Println("Error in parsing the config file for initial users")
		fmt.Println(err.Error())
	}
	var config GuideConfig
	err = json.Unmarshal(guideData, &config)
	if err != nil {
		fmt.Println("Error with reading the struct")
		fmt.Println(err.Error())
	}

	file, err := os.Create(config.OutputFile)
	if err != nil {
		fmt.Println("Error in creation of new file to store output")
		fmt.Println(err.Error())
	}
	bufferedFile := bufio.NewWriter(file)

	generateInitialSetupDoc(bufferedFile, config)

	bufferedFile.Flush()
}

func generateInitialSetupDoc(w io.Writer, config GuideConfig) {
	t := template.Must(template.New("initialSetup.md").ParseFiles(config.InitialSetup.Template))
	err := t.Execute(w, config)
	if err != nil {
		fmt.Println("Issue in trying to generate the initial setup section of the GTM Guide")
		fmt.Println(err.Error())
	}
}
