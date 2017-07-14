package context

import (
	"bufio"
	"encoding/csv"
	"os"
	"path/filepath"
	"sort"
)

type TesterRecord []string
type DeviceRecord []string
type Testers []string
type Devices []string
type Countries []string

const (
	bugs_data = "applause/data/bugs.csv"
	devices_data = "applause/data/devices.csv"
	testers_data = "applause/data/testers.csv"
	tester_device_data = "applause/data/tester_device.csv"
)

var (
	data_files = []string {testers_data, devices_data, tester_device_data, bugs_data}

	// dictionary of testers by id
	TesterMap = make(map[string]TesterRecord)

	// dictionary of devices by id
	DeviceMap = make(map[string]DeviceRecord)

	// dictionary of testers by country
	CountryTestersMap = make(map[string]Testers)

	// dictionary of testers by device
	DeviceTestersMap = make(map[string]Testers)

	// dictionary of testers by country_device pair
	CountryDeviceTestersMap = make(map[string]Testers)

	// dictionary of bug count by tester_device pair
	TesterDeviceBugCountMap = make(map[string]int)

	CountryList = Countries{}
	DeviceList = Devices{}
)

// Initializes the context by fetching all data records into various maps.
func InitContext() error {
	for _, path := range data_files {
		records, err := fetch(path)
		if err != nil {
			return err
		}
		switch path {
		case testers_data:
			for _, record := range records[1:] {
				id := record[0]
				country := record[3]
				TesterMap[id] = record
				CountryTestersMap[country] = append(CountryTestersMap[country], id)
			}
			break
		case devices_data:
			for _, record := range records[1:] {
				id := record[0]
				DeviceMap[id] = record
				DeviceList = append(DeviceList, id)
			}
			break
		case tester_device_data:
			for _, record := range records[1:] {
				tester_id := record[0]
				device_id := record[1]
				DeviceTestersMap[device_id] = append(DeviceTestersMap[device_id], tester_id)
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
	// Initialize CountryDeviceTestersMap
	for country, country_testers := range CountryTestersMap {
		CountryList = append(CountryList, country)
		for device, device_testers := range DeviceTestersMap {
			country_device := country + "_" + device
			for _, country_tester := range country_testers {
				for _, device_tester := range device_testers {
					if country_tester == device_tester {
						CountryDeviceTestersMap[country_device] =
							append(CountryDeviceTestersMap[country_device], country_tester)
					}
				}
			}
		}
	}
	return nil
}

// Fetches data records from the given file path
func fetch(path string) (records [][]string, err error) {
	fp, _ := filepath.Abs(path)
	fh, err := os.Open(fp)
	if err != nil {
		return
	}
	defer fh.Close()
	data := csv.NewReader(bufio.NewReader(fh))
	records, err = data.ReadAll()
	return
}

// Returns the unique testers for the given countries and devices
func TestersByCountryDevice(countries Countries, devices Devices) Testers {
	testers := Testers{}
	if len(countries) == 0 {
		countries = CountryList
	} else {
		countries = Unique(countries)
	}
	if len(devices) == 0 {
		devices = DeviceList
	} else {
		devices = Unique(devices)
	}
	for _, country := range countries {
		for _, device := range devices {
			country_device := country + "_" + device
			testers = append(testers, CountryDeviceTestersMap[country_device]...)
		}
	}
	return Unique(testers)
}

func Unique(testers []string) []string {
	uniqueTesters := []string{}
	found := make(map[string]bool)
	for i, x := range testers {
		if !found[x] {
			found[x] = true
			uniqueTesters = append(uniqueTesters, testers[i])
		}
	}
	return uniqueTesters
}

func RankTesters(testers Testers, devices Devices) map[string]int {
	testerRankMap := make(map[string]int)
	if len(devices) == 0 {
		devices = DeviceList
	} else {
		devices = Unique(devices)
	}
	for _, tester := range testers {
		testerRankMap[tester] += 0
		for _, device := range devices {
			tester_device := tester + "_" + device
			bug_count := TesterDeviceBugCountMap[tester_device]
			testerRankMap[tester] += bug_count
		}
	}
	return testerRankMap
}

func SortByRank(testerRankMap map[string]int) PairList {
	return sortMapByValue(testerRankMap)
}

// A data structure to hold a key/value pair.
type Pair struct {
	Key string
	Value int
}

// A slice of Pairs that implements sort.Interface to sort by Value.
type PairList []Pair
func (p PairList) Swap(i, j int) { p[i], p[j] = p[j], p[i] }
func (p PairList) Len() int { return len(p) }
func (p PairList) Less(i, j int) bool { return p[i].Value < p[j].Value }

// A function to turn a map into a PairList, then sort and return it.
func sortMapByValue(m map[string]int) PairList {
	p := make(PairList, len(m))
	i := 0
	for k, v := range m {
		p[i] = Pair{k, v}
		i++
	}
	sort.Sort(sort.Reverse(p))
	return p
}
