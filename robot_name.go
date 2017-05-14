package robotname

import (
	"math/rand"
	"strconv"
	"time"
)

const testVersion = 1

type Robot struct {
	name     string
	oldNames []string
}

func randomString(l int) string {
	bytes := make([]byte, l)
	for i := 0; i < l; i++ {
		bytes[i] = byte(randomInt(65, 90))
	}
	return string(bytes)
}

func randomInt(min int, max int) int {
	return min + rand.Intn(max-min)
}
func (r *Robot) ExistsName(name string) bool {
	for _, n := range r.oldNames {
		if n == name {
			return true
		}
	}
	return false
}
func (r *Robot) Name() string {
	if r.name != "" {
		return r.name
	}
	for {
		rand.Seed(time.Now().UTC().UnixNano())
		n := randomString(2) + strconv.Itoa(randomInt(100, 999))
		if !r.ExistsName(n) {
			r.oldNames = append(r.oldNames, n)
			r.name = n
			break
		}
	}
	return r.name
}
func (r *Robot) Reset() {
	r.name = ""
	r.oldNames = []string{}
}
