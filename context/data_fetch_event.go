package context

import (
	"encoding/csv"
	"bufio"
	"log"
	"path/filepath"
	"fmt"
	"os"
)

type DataFetchEvent struct {
	path string
}

func (t DataFetchEvent) Process() {
	defer func() {
		log.Printf("Done reading: %s\n", t.path)
		wg.Done()
	}()
	records, err := fetch(t.path)
	if err != nil {
		log.Printf("Error fetching %s: %s\n", t.path, err)
		return
	}
	switch t.path {
	case testers_data:
		for _, record := range records[1:] {
			id := record[0]
			country := record[3]
			TesterMap[id] = record
			CountryTestersMap.stringSet(country).append(id)
		}
		break
	case devices_data:
		for _, record := range records[1:] {
			id := record[0]
			DeviceMap[id] = record
			DeviceList.append(id)
		}
		break
	case tester_device_data:
		for _, record := range records[1:] {
			tester_id := record[0]
			device_id := record[1]
			DeviceTestersMap.stringSet(device_id).append(tester_id)
		}
		break
	case bugs_data:
		for _, record := range records[1:] {
			tester_device := record[2] + "_" + record[1]
			TesterDeviceBugCountMap[tester_device]++
		}
		break
	}
}

// Fetches data records from the given file path
func fetch(path string) (records [][]string, err error) {
	fp, _ := filepath.Abs(path)
	fmt.Printf("%s: %s\n", path, fp)
	fh, err := os.Open(fp)
	if err != nil {
		return
	}
	defer fh.Close()
	data := csv.NewReader(bufio.NewReader(fh))
	records, err = data.ReadAll()
	return
}
