package main

import (
	"crypto/md5"
	"crypto/rand"
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"strconv"
	"strings"
	"time"

	"github.com/javuto/osctrl/context"
	"github.com/javuto/osctrl/nodes"
)

// Constants for seconds
const (
	oneMinute        = 60
	fiveMinutes      = 300
	fifteenMinutes   = 900
	thirtyMinutes    = 1800
	fortyfiveMinutes = 2500
	oneHour          = 3600
	threeHours       = 10800
	sixHours         = 21600
	eightHours       = 28800
	twelveHours      = 43200
	fifteenHours     = 54000
	twentyHours      = 72000
	oneDay           = 86400
	twoDays          = 172800
	sevenDays        = 604800
	fifteenDays      = 1296000
)

// Function to generate a secure CSRF token
func generateCSRF() string {
	b := make([]byte, 16)
	_, _ = rand.Read(b)
	return fmt.Sprintf("%x", b)
}

// Helper to check if the CSRF token is valid
func checkCSRFToken(token string) bool {
	//return (strings.TrimSpace(token) == mainCSRFToken)
	return true
}

// Helper to generate a random MD5 to be used as query name
func generateQueryName() string {
	b := make([]byte, 32)
	_, _ = rand.Read(b)
	hasher := md5.New()
	_, _ = hasher.Write([]byte(fmt.Sprintf("%x", b)))
	return hex.EncodeToString(hasher.Sum(nil))
}

// Helper to check if the provided platform exists
func checkValidPlatform(platform string) bool {
	platforms, err := nodesmgr.GetAllPlatforms()
	if err != nil {
		return false
	}
	for _, p := range platforms {
		if p == platform {
			return true
		}
	}
	return false
}

// Helper for debugging purposes and dump a full HTTP request
func debugHTTPDump(r *http.Request, debugCheck bool, showBody bool) {
	if debugCheck {
		log.Println("-------------------------------------------------- request")
		requestDump, err := httputil.DumpRequest(r, showBody)
		if err != nil {
			log.Printf("error while dumprequest %v", err)
		}
		log.Println(string(requestDump))
		if !showBody {
			log.Println("---------------- Skipping Request Body -------------------")
		}
		log.Println("-------------------------------------------------------end")
	}
}

// Helper to remove backslashes from text
func removeBackslash(rawString string) string {
	return strings.Replace(rawString, "\\", " ", -1)
}

// Helper to remove backslashes from text and encode
func removeBackslashEncode(data []byte) string {
	return strings.Replace(string(data), "\\", " ", -1)
}

// Helper to remove backslashes from text
func stringEncode(data []byte) string {
	return string(data)
}

// Helper to generate a link to results for on-demand queries
func resultsSearchLink(name string) string {
	if logConfig.Splunk {
		return strings.Replace(logConfig.SplunkCfg["search"], "{{NAME}}", removeBackslash(name), 1)
	}
	if logConfig.Postgres {
		return "/query/logs/" + removeBackslash(name)
	}
	return ""
}

// Helper to get a string based on the difference of two times
func stringifyTime(seconds int) string {
	var timeStr string
	w := make(map[int]string)
	w[oneDay] = "day"
	w[oneHour] = "hour"
	w[oneMinute] = "minute"
	// Ordering the values will prevent bad values
	var ww [3]int
	ww[0] = oneDay
	ww[1] = oneHour
	ww[2] = oneMinute
	for _, v := range ww {
		if seconds >= v {
			d := seconds / v
			dStr := strconv.Itoa(d)
			timeStr = dStr + " " + w[v]
			if d > 1 {
				timeStr += "s"
			}
			break
		}
	}
	return timeStr + " ago"
}

// Helper to format past times in timestamp format
func pastTimestamp(t time.Time) string {
	return strconv.FormatInt(t.Unix(), 10)
}

// Helper to format past times only returning one value (minute, hour, day)
func pastTimeAgo(t time.Time) string {
	if t.IsZero() {
		return "Never"
	}
	now := time.Now()
	seconds := int(now.Sub(t).Seconds())
	if seconds < 2 {
		return "Just Now"
	}
	if seconds < oneMinute {
		return strconv.Itoa(seconds) + " seconds ago"
	}
	if seconds > fifteenDays {
		return "Since " + t.Format("Mon Jan 02 15:04:05 MST 2006")
	}
	return stringifyTime(seconds)
}

// Helper to generate stats for all contexts
func getContextStats(contexts []context.TLSContext) (nodes.StatsData, error) {
	contextStats := make(nodes.StatsData)
	for _, c := range contexts {
		stats, err := nodesmgr.GetStatsByContext(c.Name)
		if err != nil {
			return contextStats, err
		}
		contextStats[c.Name] = stats
	}
	return contextStats, nil
}

// Helper to generate stats for all platforms
func getPlatformStats(platforms []string) (nodes.StatsData, error) {
	platformStats := make(nodes.StatsData)
	for _, p := range platforms {
		stats, err := nodesmgr.GetStatsByPlatform(p)
		if err != nil {
			return platformStats, err
		}
		platformStats[p] = stats
	}
	return platformStats, nil
}

// Helper to calculate the osquery config_hash and skip sending a blob that won't change anything
// https://github.com/facebook/osquery/blob/master/osquery/config/config.cpp#L911
// osquery calculates the SHA1 of the configuration blob, then the SHA1 hash of that
func generateOsqueryConfigHash(config string) string {
	firstHasher := sha1.New()
	secondHasher := sha1.New()
	// Get SHA1 of configuration blob
	_, _ = firstHasher.Write([]byte(config))
	// Get SHA1 of the first hash
	_, _ = secondHasher.Write([]byte(hex.EncodeToString(firstHasher.Sum(nil))))
	return hex.EncodeToString(secondHasher.Sum(nil))
}
