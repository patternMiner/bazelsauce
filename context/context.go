package context

import (
	"bufio"
	"encoding/csv"
	"os"
	"path/filepath"
)

const (
	bugs_data = "github.com/patternMiner/applause/data/bugs.csv"
	devices_data = "github.com/patternMiner/applause/data/devices.csv"
	testers_data = "github.com/patternMiner/applause/data/testers.csv"
	tester_device_data = "github.com/patternMiner/applause/data/tester_device.csv"
)

var (
	data_files = []string {testers_data, devices_data, tester_device_data, bugs_data}

	// dictionary of testers by id
	TesterMap = make(StringSliceMap)

	// dictionary of devices by id
	DeviceMap = make(StringSliceMap)

	// dictionary of testers by country
	CountryTestersMap = make(StringSetMap)

	// dictionary of testers by device
	DeviceTestersMap = make(StringSetMap)

	// dictionary of testers by country_device pair
	CountryDeviceTestersMap = make(StringSetMap)

	// dictionary of bug count by tester_device pair
	TesterDeviceBugCountMap = make(RankMap)

	CountryList = make(StringSet)
	DeviceList = make(StringSet)
)

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
	// Initialize CountryDeviceTestersMap
	for country, country_testers := range CountryTestersMap {
		CountryList.append(country)
		for device, device_testers := range DeviceTestersMap {
			country_device := country + "_" + device
			for country_tester := range country_testers {
				for device_tester := range device_testers {
					if country_tester == device_tester {
						CountryDeviceTestersMap.stringSet(country_device).append(country_tester)
					}
				}
			}
		}
	}
	return nil
}

// Returns the testerSet for the given countrySet and deviceSet
func TesterSetByCountryDevice(countrySet, deviceSet StringSet) StringSet {
	testers := make(StringSet)
	for country := range countrySet {
		for device := range deviceSet {
			country_device := country + "_" + device
			testers.merge(CountryDeviceTestersMap[country_device])
		}
	}
	return testers
}

// Returns the RankMap of the given testerSet for the given deviceSet
func TesterRankMap(testerSet, deviceSet StringSet) RankMap {
	testerRankMap := make(RankMap)
	for tester := range testerSet {
		testerRankMap[tester] += 0
		for device := range deviceSet {
			tester_device := tester + "_" + device
			bug_count := TesterDeviceBugCountMap[tester_device]
			testerRankMap[tester] += bug_count
		}
	}
	return testerRankMap
}

func MatchTesters(countries, devices StringSlice) PairList {
	countrySet := countries.makeStringSet()
	deviceSet := devices.makeStringSet()
	if len(countrySet) == 0 {
		countrySet = CountryList
	}
	if len(deviceSet) == 0 {
		deviceSet = DeviceList
	}
	testerSet := TesterSetByCountryDevice(countrySet, deviceSet)
	testerRankMap := TesterRankMap(testerSet, deviceSet)
	// Make a PairList out of testerRankMap, each Pair containing the tester_id as key and bug_count as value.
	testerList := makePairList(testerRankMap)
	// Sort the PairList by rank
	testerList.SortByValue()
	return testerList
}
