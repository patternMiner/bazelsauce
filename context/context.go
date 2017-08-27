package context

import (
	"github.com/patternMiner/async"
	"sync"
	"log"
)

const (
	bugs_data = "data/bugs.csv"
	devices_data = "data/devices.csv"
	testers_data = "data/testers.csv"
	tester_device_data = "data/tester_device.csv"
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

	// data fetcher task synchronization lock
	wg sync.WaitGroup
)


// Initializes the context by fetching all data records into various maps.
func InitContext() error {
	async.StartDispatcher(4)
	wg.Add(len(data_files))
	for _, path := range data_files {
		async.EventQueue <- DataFetchEvent{path}
	}
	wg.Wait()

	// Initialize CountryDeviceTestersMap
	for country, country_testers := range CountryTestersMap {
		CountryList.append(country)
		log.Printf("Country: %s\n", country)
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
