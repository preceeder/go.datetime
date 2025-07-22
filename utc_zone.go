package datetime

import (
	"errors"
	"fmt"
	"regexp"
	"slices"
	"strconv"
	"sync"
	"time"
)

// zone 目前只支持 utc
func timeZoneHandler(zone ...string) *time.Location {
	if len(zone) > 0 {
		l, _ := ParseUtcTimeOffset(zone[0])
		return l
	}
	return time.Local
}

// var timeZoneMap = make(map[string]*time.Location)
var timeZoneMap = sync.Map{}

// ParseUtcTimeOffset
// TUC pares   "UTC+12"  -> 12*60*60 -> time.FixedZone("UTC+12", 12*60*60)
func ParseUtcTimeOffset(offsetStr string) (location *time.Location, err error) {
	if lz, ok := timeZoneMap.Load(offsetStr); ok {
		return lz.(*time.Location), nil
	}
	mcp := regexp.MustCompile(`UTC$|UTC([-+])(\d+)$|UTC([-+])(\d+):(\d+)$`)
	fin := mcp.FindStringSubmatch(offsetStr)
	var offset int = 0
	sign := 1
	if len(fin) > 0 {
		validData := slices.DeleteFunc(fin, func(e string) bool {
			if e == "" {
				return true
			}
			return false
		})
		for index, d := range validData {
			if d == offsetStr {
				continue
			}
			switch index {
			case 1:
				if d == "-" {
					sign = -1
				}
			case 2:
				hour, err := strconv.Atoi(d)
				if err != nil {
					return nil, err
				}
				offset += hour * 60 * 60
			case 3:
				minute, err := strconv.Atoi(d)
				if err != nil {
					return nil, err
				}
				offset += minute * 60
			}
		}
		offset *= sign
		location = time.FixedZone(offsetStr, offset)
		timeZoneMap.Store(offsetStr, location)
		timeZoneMap.Store(offset, location) //偏差值作为key
		return
	}

	return nil, errors.New(fmt.Sprintf("TimeZoneName `%s` is not UTC", offsetStr))
}
