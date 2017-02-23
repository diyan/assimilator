package tsdb

import (
	"sort"
	"time"
)

type TSDBModel int

// TODO consider add Model prefix, so usage would be:
//   tsdb.ModelInternal, tsdb.ModelProjectTagKey, tsdb.ModelFrequentProjectsByOrganization
const (
	Internal TSDBModel = 0

	// number of events seen specific to grouping
	Project         TSDBModel = 1
	ProjectTagKey   TSDBModel = 2
	ProjectTagValue TSDBModel = 3
	Group           TSDBModel = 4
	GroupTagKey     TSDBModel = 5
	GroupTagValue   TSDBModel = 6
	Release         TSDBModel = 7

	// the number of events sent to the server
	ProjectTotalReceived TSDBModel = 100
	// the number of events rejected due to rate limiting
	ProjectTotalRejected TSDBModel = 101
	// the number of operations
	ProjectOperations TSDBModel = 102
	// the number of operations with an error state
	ProjectOperationErrors TSDBModel = 103
	// the number of events blocked due to being blacklisted
	ProjectTotalBlacklisted TSDBModel = 104

	// the number of events sent to the server
	OrganizationTotalReceived TSDBModel = 200
	// the number of events rejected due to rate limiting
	OrganizationTotalTejected TSDBModel = 201
	// the number of events blocked due to being blacklisted
	OrganizationTotalBlacklisted TSDBModel = 202

	// distinct count of users that have been affected by an event in a group
	UsersAffectedByGroup TSDBModel = 300
	// distinct count of users that have been affected by an event in a project
	UsersAffectedByProject TSDBModel = 301

	// FrequentOrganizationReceivedBySystem TSDBModel = 400
	// FrequentOrganizationRejectedBySystem TSDBModel = 401
	// FrequentOrganizationBlacklistedBySystem TSDBModel = 402
	// FrequentValuesByIssueTag TSDBModel = 405

	// number of events seen for a project, by organization
	FrequentProjectsByOrganization TSDBModel = 403
	// number of issues seen for a project, by project
	FrequentIssuesByProject TSDBModel = 404
	// number of events seen for a release, by issue
	// FrequentReleasesByGroup TSDBModel = 406  // DEPRECATED
	// number of events seen for a release, by issue
	FrequentReleasesByGroup TSDBModel = 407
	// number of events seen for an environment, by issue
	FrequentEnvironmentsByGroup TSDBModel = 408
)

// Int64Slice attaches the methods of Interface to []int64, sorting in increasing order.
// TODO move this type to the generic package
type Int64Slice []int64

func (s Int64Slice) Len() int           { return len(s) }
func (s Int64Slice) Less(i, j int) bool { return s[i] < s[j] }
func (s Int64Slice) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }

type TimeSeries struct {
	UnixTime int64
	Value    int
}

// timeSeriesSlice attaches the methods of Interface to []TimeSeries, sorting by UnixTime in increasing order.
type timeSeriesSlice []TimeSeries

func (s timeSeriesSlice) Len() int           { return len(s) }
func (s timeSeriesSlice) Less(i, j int) bool { return s[i].UnixTime < s[j].UnixTime }
func (s timeSeriesSlice) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }

// TODO BaseTSDB class holds state in self.rollups and self.__legacy_rollups
//   self.rollups is an OrderedDict

// TODO InMemoryTSDB class also holds a state in:
//   model => key => timestamp = count
//     self.data
//   self.sets[model][key][rollup] = set of elements
//   i.e. nested defaultdict that holds set
//     self.sets
//   self.frequencies[model][key][rollup] = Counter()
//   i.e. nested defaultdict that holds Counter
//     self.frequencies

type intSet map[int]bool

// TODO is counter should be map[int64]float64 ?
type counter map[int64]int

type InMemoryTSDB struct {
	data        map[TSDBModel]map[int]map[int64]int
	sets        map[TSDBModel]map[int]map[int64]intSet
	frequencies map[TSDBModel]map[int]map[int64]counter
}

func New() InMemoryTSDB {
	tsdb := InMemoryTSDB{}
	// TODO init data, sets, frequencies
	return tsdb
}

// TODO is GetCount a good func name?
//   Will it conflict with getter for tsdb.frequencies that holds counters?
func (tsdb InMemoryTSDB) getRecord(model TSDBModel, key int, epoch int64) int {
	value := 0
	if byModel, ok := tsdb.data[model]; ok {
		if byModelAndKey, ok := byModel[key]; ok {
			if byModelAndKeyAndTime, ok := byModelAndKey[epoch]; ok {
				return byModelAndKeyAndTime
			}
		}
	}
	return value
}

// GetRange returns a slice of time-series for specified model, keys and interval
// TODO reformat code from Python into Golang
//   To get a range of data for group ID=[1, 2, 3]:
//   Returns a mapping of key => [(timestamp, count), ...].
//   >>> now = timezone.now()
//   >>> get_range([TSDBModel.group], [1, 2, 3],
//   >>>           start=now - timedelta(days=1),
//   >>>           end=now)
// TODO GetRange implemented in Redis- InMemory- but not in BaseTSDB
func (tsdb InMemoryTSDB) GetRange(model TSDBModel, keys []int, start, end time.Time, rollup int) map[int][]TimeSeries {
	// TODO rollup arg is optional in Python, None by default
	rollup, series := getOptimalRollupSeries(start, end, rollup)
	results := map[int][]TimeSeries{}
	for _, epoch := range series {
		normEpoch := normalizeToRollup(time.Unix(epoch, 0), rollup)
		for _, key := range keys {
			value := tsdb.getRecord(model, key, normEpoch)
			if ts, ok := results[key]; ok {
				results[key] = append(ts, TimeSeries{UnixTime: epoch, Value: value})
			} else {
				results[key] = []TimeSeries{}
			}
		}
	}
	for _, series := range results {
		sort.Sort(timeSeriesSlice(series))
	}
	return results
}

func getOptimalRollupSeries(start, end time.Time, rollup int) (int, []int64) {
	// TODO rollup arg is optional in Python, None by default
	// TODO end is optional in Python, utc now by default
	// TODO if rollup is None, we should call rollup := GetOptimalRollup(start, end)

	// This attempts to create a range with a duration as close as possible
	// to the requested interval using the requested (or inferred) rollup
	// resolution. This result always includes the ``end`` timestamp, but
	// may not include the ``start`` timestamp.
	series := []int64{}
	for timestamp := end; timestamp.Before(start); timestamp = timestamp.Add(time.Duration(-1*rollup) * time.Second) {
		series = append(series, normalizeToEpoch(timestamp, rollup))
	}
	sort.Sort(Int64Slice(series))
	return rollup, series
}

// Given a ``timestamp`` (datetime object) normalize to an epoch timestamp.
//   i.e. if the rollup is minutes, the resulting timestamp would have
//   the seconds and microseconds rounded down.
func normalizeToEpoch(timestamp time.Time, seconds int) int64 {
	epoch := timestamp.Unix()
	return epoch - (epoch % int64(seconds))
	// TODO consider
	//return timestamp.Round(time.Duration(seconds) * time.Second).Unix()
}

// Given a ``epoch`` normalize to an epoch rollup.
func normalizeTsToEpoch(epoch int64, seconds int) int64 {
	return epoch - (epoch % int64(seconds))
}

// Given a ``timestamp`` (datetime object) normalize to an epoch rollup.
func normalizeToRollup(timestamp time.Time, seconds int) int64 {
	return timestamp.Unix() / int64(seconds)
}
