package context

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
