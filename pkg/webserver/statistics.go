package webserver

import (
	"os"
	"sync"
	"time"
	"strings"
	"strconv"
	"runtime"
	"io/ioutil"
	"encoding/json"
	"net/http"
	"net/url"
)

type StatisticsFormat struct {
	Time struct {
		Start int64 `json:"start"`
		Latest int64 `json:"latest"`
	} `json:"time"`
	Tables struct {
		Hits struct {
			Daily map[string]int `json:"daily"`
			Total int `json:"total"`
		} `json:"hits"`
		Requests struct {
			Counters map[string]int `json:"counters"`
		} `json:"requests"`
		Referrers struct {
			Counters map[string]int `json:"counters"`
		} `json:"referrers"`
		StatusCodes struct {
			Counters map[string]int `json:"counters"`
		} `json:"statusCodes"`
		RemoteAddresses struct {
			Counters map[string]int `json:"counters"`
		} `json:"remoteAddresses"`
	} `json:"tables"`
	Site struct {
		Size int `json:"size"`
		Pages int `json:"pages"`
	} `json:"site"`
	TLS struct {
		Issued int64 `json:"date"`
		Expires int64 `json:"expires"`
		Domains []string `json:"domains"`
		Authority string `json:"authority"`
	} `json:"tls"`
	Data struct {
		Total int64 `json:"total"`
		NoCompression int64 `json:"noCompression"`
		Ratio float64 `json:"ratio"`
	} `json:"data"`
	Server struct {
		Hostname string `json:"hostname"`
		OperatingSystem string `json:"operatingSystem"`
		Architecture string `json:"architecture"`
		Address string `json:"address"`
		CPU struct {
			Cores int `json:"cores"`
		} `json:"cpu"`
		Memory struct {
			Total int64 `json:"total"`
			Available int64 `json:"available"`
		} `json:"memory"`
		Uptime struct {
			Server int64 `json:"server"`
			System int64 `json:"system"`
		} `json:"uptime"`
	} `json:"server"`
}

type statistics struct {
	file string
	pages []string
	mutex sync.Mutex
	stats StatisticsFormat
	startTime int64
	stopChan chan bool
}

func (s *statistics) record(status int, r *http.Request) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	s.stats.Tables.Hits.Daily[time.Now().Format("2006-01-02")]++
	s.stats.Tables.StatusCodes.Counters[strconv.Itoa(status)]++
	if status == http.StatusOK {
		route := r.URL.Path
		if route == "" { route = "/" }
		for _, page := range s.pages {
			if route == page {
				s.stats.Tables.Requests.Counters[route]++
			}
		}
	}
	if u, err := url.Parse(r.Referer()); err == nil {
		hostname := u.Hostname()
		if len(hostname) != 0 && hostname != r.Host {
			s.stats.Tables.Referrers.Counters[hostname]++
		}
	}
	if len(r.RemoteAddr) != 0 {
		s.stats.Tables.RemoteAddresses.Counters[r.RemoteAddr]++;
	}
}

func (s *statistics) start(errChan chan<- error, delay time.Duration) {
	s.stopChan = make(chan bool)
	if err := s.load(); err != nil { errChan <- err }
	s.initEmpty()
	s.startTime = time.Now().Unix()
	t := time.NewTicker(delay)
	for {
		select {
		case <-t.C:
			if err := s.save(); err != nil {
				errChan <- err
			}
		case <-s.stopChan:
			s.save();
			t.Stop()
			return
		}
	}
}

func (s *statistics) initEmpty() {
	if s.stats.Tables.Hits.Daily == nil { s.stats.Tables.Hits.Daily = make(map[string]int) }
	if s.stats.Tables.Requests.Counters == nil { s.stats.Tables.Requests.Counters = make(map[string]int) }
	if s.stats.Tables.Referrers.Counters == nil { s.stats.Tables.Referrers.Counters = make(map[string]int) }
	if s.stats.Tables.StatusCodes.Counters == nil { s.stats.Tables.StatusCodes.Counters = make(map[string]int) }
	if s.stats.Tables.RemoteAddresses.Counters == nil { s.stats.Tables.RemoteAddresses.Counters = make(map[string]int) }
}

func (s *statistics) save() error {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	now := time.Now().Unix()

	// Time
	s.stats.Time.Latest = now
	if s.stats.Time.Start == 0 {
		s.stats.Time.Start = now
	}

	// Tables
	s.stats.Tables.Hits.Total = 0
	for _, count := range s.stats.Tables.Hits.Daily {
		s.stats.Tables.Hits.Total += count
	}

	// Site TODO
	s.stats.Site.Size = 0 // size of every file referenced in sitefile
	s.stats.Site.Pages = 0 // number of pages listed in sitefile

	// TLS TODO
	s.stats.TLS.Issued = 0 // not before time
	s.stats.TLS.Expires = 0 // not after time
	s.stats.TLS.Domains = nil
	s.stats.TLS.Authority = "TODO"

	// Data
	if s.stats.Data.NoCompression != 0 {
		s.stats.Data.Ratio = float64(s.stats.Data.Total) / float64(s.stats.Data.NoCompression)
	} else {
		s.stats.Data.Ratio = 0
	}

	// Server
	hostname, err := os.Hostname()
	if err != nil { hostname = "<unknown>" }
	s.stats.Server.Hostname = hostname
	s.stats.Server.OperatingSystem = runtime.GOOS
	s.stats.Server.Architecture = runtime.GOARCH
	s.stats.Server.Address = "TODO"
	s.stats.Server.CPU.Cores = runtime.NumCPU()
	meminfoBytes, err := ioutil.ReadFile("/proc/meminfo")
	if err != nil { return err }
	meminfoFields := strings.Fields(string(meminfoBytes))
	s.stats.Server.Memory.Total, err = strconv.ParseInt(meminfoFields[1], 10, 64)
	if err != nil { return err }
	s.stats.Server.Memory.Available, err = strconv.ParseInt(meminfoFields[7], 10, 64)
	if err != nil { return err }
	s.stats.Server.Uptime.Server = now - s.startTime
	uptimeBytes, err := ioutil.ReadFile("/proc/uptime")
	if err != nil { return err }
	uptimeFloat, err := strconv.ParseFloat(strings.Fields(string(uptimeBytes))[0], 64)
	if err != nil { return err }
	s.stats.Server.Uptime.System = int64(uptimeFloat)
	if err != nil { return err }

	bytes, err := json.Marshal(s.stats)
	if err != nil { return err }
	var mode os.FileMode
	if info, err := os.Stat(s.file); err == nil {
		mode = info.Mode()
	} else {
		mode = 0777
	}
	return ioutil.WriteFile(s.file, bytes, mode)
}

func (s *statistics) load() error {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	bytes, err := ioutil.ReadFile(s.file)
	if err != nil { return err }
	return json.Unmarshal(bytes, &s.stats)
}

