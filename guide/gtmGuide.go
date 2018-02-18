package main

import (
	"bufio"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path"
	"strconv"
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
	EventList     string `json:"eventList"`
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
			TitleTemplate: "templates/gtm/eventsHeader.md",
			Template:      "templates/gtm/events.md",
			EventList:     "templates/gtm/events.csv",
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
	generateEventDoc(bufferedFile, config)

	bufferedFile.Flush()
}

func generateInitialSetupDoc(w io.Writer, config GuideConfig) {
	_, filename := path.Split(config.InitialSetup.Template)
	t := template.Must(template.New(filename).ParseFiles(config.InitialSetup.Template))
	err := t.Execute(w, config)
	if err != nil {
		fmt.Println("Issue in trying to generate the initial setup section of the GTM Guide")
		fmt.Println(err.Error())
	}
	// Add a few new lines
	io.WriteString(w, "\n\n")
}

func generateEventDoc(w io.Writer, config GuideConfig) {
	_, titleTemplateFilename := path.Split(config.Events.TitleTemplate)
	t := template.Must(template.New(titleTemplateFilename).ParseFiles(config.Events.TitleTemplate))
	err := t.Execute(w, config)
	if err != nil {
		fmt.Println("Issue in trying to generate the event header setup section of the GTM Guide")
		fmt.Println(err.Error())
	}

	// Add a few new lines
	io.WriteString(w, "\n\n")

	// Generate the event table in markdown
	eventListFile, err := os.Open(config.Events.EventList)
	if err != nil {
		fmt.Println("Error in trying to find event list file")
		fmt.Println(err.Error())
	}
	eventList, err := csv.NewReader(eventListFile).ReadAll()
	if err != nil {
		fmt.Println("Error when reading event list file")
		fmt.Println(err.Error())
	}
	eventList = eventList[1:len(eventList)]
	var filteredEventList [][]string
	for _, row := range eventList {
		filteredEventList = append(filteredEventList, row[1:len(row)])
	}
	eventListTable := tablewriter.NewWriter(w)
	eventListTable.SetHeader([]string{"Event Name", "Event Description", "Event Category", "Event Action", "Event Label", "Event Value"})
	eventListTable.SetBorders(tablewriter.Border{Left: true, Top: false, Right: true, Bottom: false})
	eventListTable.SetCenterSeparator("|")
	eventListTable.AppendBulk(filteredEventList)
	eventListTable.Render()

	// Add a few new lines
	io.WriteString(w, "\n\n")

	// Create event pages
	type eventDetails struct {
		Image       string
		Name        string
		Description string
		Category    string
		Action      string
		Label       string
		Value       int
	}
	for _, row := range eventList {
		eventValue, err := strconv.Atoi(row[6])
		if err != nil {
			fmt.Println("Error in parsing a string to a number")
			fmt.Println(err.Error())
		}
		eventData := eventDetails{
			Image:       row[0],
			Name:        row[1],
			Description: row[2],
			Category:    row[3],
			Action:      row[4],
			Label:       row[5],
			Value:       eventValue,
		}

		_, titleTemplateFilename := path.Split(config.Events.Template)
		t := template.Must(template.New(titleTemplateFilename).ParseFiles(config.Events.Template))
		err = t.Execute(w, eventData)
		if err != nil {
			fmt.Println("Issue in trying to generate the event section of the GTM Guide")
			fmt.Println(err.Error())
		}

		// Add a few new lines
		io.WriteString(w, "\n\n")
	}
}
