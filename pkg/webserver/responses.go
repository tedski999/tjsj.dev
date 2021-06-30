package webserver

import (
	"net/http"
	"html/template"
	"fmt"
	"errors"
)

type postMetadata struct {
	Date string
	Title string
	Tags []string
	ID string
}

type homeResponseData struct {
	SplashText string
	ProjectsList []template.HTML
	RecentPostsList []template.HTML
}

type postsResponseData struct {
	SearchQuery, MonthQuery, YearQuery string
	TagList map[string]bool
	YearList map[string][][2]string
	PostList []postMetadata
}
type postsResponseDataYearList struct {
	Name string
	MonthList []string
}


type statsResponseData struct {
	StatsLists [2][]string
	StatsStartDatetime string
}

// General method for executing HTML templates by name
func (server *Server) executeHTMLTemplate(w http.ResponseWriter, templateName string, data interface{}) {
	template := server.content.GetHTMLTemplate(templateName)
	if template == nil {
		server.errChan <- errors.New("HTML template '" + templateName + "' not found!")
		return
	}

	template.Execute(w, data)
}

// Respond with the HTML template "home.html"
func (server *Server) homeResponse(w http.ResponseWriter, r *http.Request) {
	// TODO: get list of recent posts metadata
	server.executeHTMLTemplate(w, "home.html", homeResponseData {
		server.content.GetRandomSplashText(),
		nil,
		nil,
	})
}

// Respond with a list of projects in the HTML template "projects.html"
func (server *Server) projectsResponse(w http.ResponseWriter, r *http.Request) {
	// TODO: get list of projects metadata
	//server.executeHTMLTemplate(w, "projects.html", nil)
	server.errorResponse(w, r, http.StatusNotFound)
}

// Respond with the project page of the id given in the URL
func (server *Server) projectResponse(w http.ResponseWriter, r *http.Request) {
	//vars := mux.Vars(r)
	//id := vars["id"]
	// TODO: find the project
	//server.executeHTMLTemplate(w, "project.html", nil)
	server.errorResponse(w, r, http.StatusNotFound)
}

// Respond with a list of posts in the HTML template "posts.html"
func (server *Server) postsResponse(w http.ResponseWriter, r *http.Request) {

	// Retrieve search query
	searchQuery := r.URL.Query().Get("s")
	monthQuery := r.URL.Query().Get("m")
	yearQuery := r.URL.Query().Get("y")
	tagsQuery := r.URL.Query()["t"]

	// TODO: get list of posts metadata filtered by queries
	// NOTE: below is an example response
	postList := []postMetadata {
		{ "2021-06-28", "Web Dev Shenanigans",                    []string {"Project","Web Dev","Go" }, "web-dev-shenanigans" },
		{ "2021-06-22", "a",                                      []string {"Game Dev","C/C++" },       "a" },
		{ "2021-06-04", "Another post title which is quite long", nil,                  "another-post-tite-which-is-quite-long" },
		{ "2021-06-04", "This title is obnoxiously long just to show off what happens to those who defy me",
		   nil,"this-title-is-obnoxiously-long-just-to-show-off-what-happens-to-those-who-defy-me" },
		{ "2021-05-11", "Hello World",                            []string {"Casual"},                  "hello-world" },
	}

	// TODO: get all the possible search queries
	// NOTE: below is an example response
	allTags := []string {"Project","Web Dev","Game Dev","Go","C/C++","Python","Java"}
	allDates := map[string][][2]string {
		"2021": {{"January","1"},{"March","3"},{"April","4"},{"June","8"}},
		"2020": {{"May","5"},{"June","6"},{"August","8"},{"November","11"},{"December","12"}},
	}

	// Mark every tag that is part of the query
	tagMap := make(map[string]bool)
	for _, tag := range allTags {
		tagMap[tag] = isStringInSlice(tag, tagsQuery)
	}

	server.executeHTMLTemplate(w, "posts.html", postsResponseData {
		searchQuery, monthQuery, yearQuery,
		tagMap, allDates, postList,
	})
}

// Respond with the post page of the id given in the URL
func (server *Server) postResponse(w http.ResponseWriter, r *http.Request) {
	//vars := mux.Vars(r)
	//id := vars["id"]
	// TODO: find the post
	//server.executeHTMLTemplate(w, "post.html", nil)
	server.errorResponse(w, r, http.StatusNotFound)
}

// Respond with the statistics page
func (server *Server) statsResponse(w http.ResponseWriter, r *http.Request) {
	statsLists := [2][]string{}

	// Total data transferred
	totalCompressedDataTransferred, totalUncompressedDataTransferred := server.stats.GetTotalDataTransferred()
	compressionRatio := float64(totalCompressedDataTransferred) / float64(totalUncompressedDataTransferred)

	// Hit counters
	hitCounters, hitCountersOrder := server.stats.GetHitCounters()
	totalHitCounts := 0
	topURLsList := []string{}
	for _, url := range hitCountersOrder {
		totalHitCounts += hitCounters[url]
		if len(topURLsList) < 10 {
			topURLsList = append(
				topURLsList,
				fmt.Sprintf("| %d hits: %s\n", hitCounters[url], url))
		}
	}

	// Referrer counters
	referrerCounters, referrerCountersOrder := server.stats.GetReferrerCounters()
	topReferrersList := []string{}
	for _, referrer := range referrerCountersOrder {
		if len(topReferrersList) < 10 {
			topReferrersList = append(
				topReferrersList,
				fmt.Sprintf("| %d referees: %s\n", referrerCounters[referrer], referrer))
		}
	}

	// Response code counters
	responseCodeCounters, responseCodeCountersOrder := server.stats.GetResponseCodeCounters()
	topResponseCodesList := []string{}
	for _, responseCode := range responseCodeCountersOrder {
		if len(topResponseCodesList) < 10 {
			topResponseCodesList = append(
				topResponseCodesList,
				fmt.Sprintf(
					"| %d responses: %d - %s\n",
					responseCodeCounters[responseCode],
					responseCode,
					http.StatusText(responseCode)))
		}
	}

	// Request stats list
	statsLists[0] = make([]string, 0, 7 + len(topURLsList) + len(topReferrersList) + len(topResponseCodesList))
	statsLists[0] = append(statsLists[0], fmt.Sprintf("Total website hits: %d", totalHitCounts))
	statsLists[0] = append(statsLists[0], "")
	statsLists[0] = append(statsLists[0], fmt.Sprintf("Top %d URLs:", len(topURLsList)))
	statsLists[0] = append(statsLists[0], topURLsList...)
	statsLists[0] = append(statsLists[0], "")
	statsLists[0] = append(statsLists[0], fmt.Sprintf("Top %d referrers:", len(topReferrersList)))
	statsLists[0] = append(statsLists[0], topReferrersList...)
	statsLists[0] = append(statsLists[0], "")
	statsLists[0] = append(statsLists[0], fmt.Sprintf("Top %d response codes:", len(topResponseCodesList)))
	statsLists[0] = append(statsLists[0], topResponseCodesList...)

	// System stats list
	sysstats := server.stats.GetSystemStats()
	statsLists[1] = []string {
		"Server: " + sysstats.Hostname,
		"Platform: " + sysstats.OS + "/" + sysstats.Arch,
		"",
		fmt.Sprintf("CPU usage: %d%% (%d cores)", sysstats.CPUUsage, sysstats.CPUCount),
		fmt.Sprintf("RAM usage: %s (%s available)",
			bytesToHumanReadable(sysstats.RAMUsage),
			bytesToHumanReadable(sysstats.RAMAvailable)),
		fmt.Sprintf("Active Goroutines: %d", sysstats.GoroutineCount),
		"",
		fmt.Sprintf("Total transferred: %s", bytesToHumanReadable(totalCompressedDataTransferred)),
		fmt.Sprintf("Before compression: %s", bytesToHumanReadable(totalUncompressedDataTransferred)),
		fmt.Sprintf("Avg. compression ratio: %f", compressionRatio),
		"",
		"Server Uptime: " + server.stats.GetUptime(),
		"System Uptime: " + sysstats.Uptime,
	}

	// Execute HTML template response
	server.executeHTMLTemplate(w, "stats.html", statsResponseData {
		statsLists,
		server.stats.GetStatsStartDatetime(),
	})
}

// Respond with the error page with an appropriate message
func (server *Server) errorResponse(w http.ResponseWriter, r *http.Request, code int) {
	w.WriteHeader(code)
	data := struct { Code int; Message string } { code, http.StatusText(code) }
	server.executeHTMLTemplate(w, "error.html", data)
}

// Convert bytes to an approprate magnitude
func bytesToHumanReadable(bytes uint64) string {
	if bytes < 1000 {
		return fmt.Sprintf("%dB", bytes);
	}

	magnitude := 0
	prefixes := []byte{'k','M','G','T','P','E'}
	for bytes >= 999950 && magnitude < len(prefixes) {
		bytes /= 1000;
		magnitude++
	}

	return fmt.Sprintf("%.1f%cB", float64(bytes) / 1000.0, prefixes[magnitude]);
}

// Determine if a string is within a slice
func isStringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

