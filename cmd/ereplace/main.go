package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"sync"
)

func modifyPayload(payload string) string {
	payload = strings.Trim(payload)
	equalIndex := strings.Index(payload, "=")
	if equalIndex != -1 {
		url := payload[:equalIndex+1]
		payload = fmt.Sprintf("%s%s", url, payload[equalIndex+1:])
	}
	return payload
}

func worker(input <-chan string, output chan<- string, wg *sync.WaitGroup) {
	defer wg.Done()
	for payload := range input {
		modifiedPayload := modifyPayload(payload)
		output <- modifiedPayload
	}
}

func main() {
	inputChan := make(chan string)
	outputChan := make(chan string)
	var wg sync.WaitGroup

	// Parse flags
	urlFlag := flag.String("u", "", "Single payload to modify. Example: -u '/../../etc/passwd'")
	payloadFileFlag := flag.String("pf", "", "File containing multiple payloads to modify. One payload per line.")
	multiplePayloadFilesFlag := flag.String("mpf", "", "Files containing multiple payloads to modify. Can specify multiple files.")
	saveFlag := flag.String("s", "", "Output file to save the modified payloads. Example: -s output.txt")
	versionFlag := flag.Bool("version", false, "Print version information")

	flag.Parse()

	if *versionFlag {
		fmt.Println("Version:", getProjectVersion())
		os.Exit(0)
	}

	// Start worker goroutines
	numWorkers := min(os.cpuCount(), len(*multiplePayloadFilesFlag)*4) // Adjust this number based on your system's capabilities
	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go worker(inputChan, outputChan, &wg)
	}

	// Read inputs
	if *urlFlag != "" {
		inputChan <- *urlFlag
	} else if *payloadFileFlag != "" {
		data, err := ioutil.ReadFile(*payloadFileFlag)
		if err != nil {
			fmt.Printf("Error reading file '%s': %v\n", *payloadFileFlag, err)
			os.Exit(1)
		}
		for _, line := range strings.Split(string(data), "\n") {
			inputChan <- line
		}
	} else if *multiplePayloadFilesFlag != "" {
		for _, file := range strings.Split(*multiplePayloadFilesFlag, " ") {
			data, err := ioutil.ReadFile(file)
			if err != nil {
				fmt.Printf("Error reading file '%s': %v\n", file, err)
				continue
			}
			for _, line := range strings.Split(string(data), "\n") {
				inputChan <- line
			}
		}
	}

	close(inputChan)
	wg.Wait()
	close(outputChan)

	// Write output to file
	outputFile, err := os.Create(*saveFlag)
	if err != nil {
		fmt.Printf("Error creating output file '%s': %v\n", *saveFlag, err)
		os.Exit(1)
	}
	defer outputFile.Close()

	for payload := range outputChan {
		fmt.Fprintln(outputFile, payload)
	}
}

func getProjectVersion() string {
	v, err := os.ReadFile("VERSION")
	if err != nil {
		return "Unknown"
	}
	return string(v)
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}
