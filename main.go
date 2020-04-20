package main

import (
	"strconv"
	"fmt"
	"sort"
	"strings"
	"time"
)

type photo struct {
	id          int
	sequence    int
	name        string
	city        string
	photoTime   time.Time
	updatedName string
	extension   string
	paddingWidth int
}

type ByDate []photo
type ById []photo

var photos []photo

func (p photo) createName() {

	p.updatedName = fmt.Sprintf("%s%0*d.%s", p.city,p.paddingWidth, p.sequence, p.extension)
	
}

func (p ByDate) Len() int {
	return len(p)
}

func (p ByDate) Less(i, j int) bool {
	return p[i].photoTime.Before(p[j].photoTime)
}

func (p ByDate) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func (p ById) Len() int {
	return len(p)
}

func (p ById) Less(i, j int) bool {
	return p[i].id < p[j].id
}

func (p ById) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func createDateTime(s string) time.Time {
	t, err := time.Parse("2006-01-02 15:04:05", s)
	if err != nil {
		fmt.Printf("error parsing time: %v", err)
	}
	return t
}
func main() {

	// map cityname to slice of photos for that city (sorted by date time)

	s := `photo.jpg, Warsaw, 2013-09-05 14:08:15
john.png, London, 2015-06-20 15:13:22
myFriends.png, Warsaw, 2013-09-05 14:07:13
Eiffel.jpg, Paris, 2015-07-23 08:03:02
pisatower.jpg, Paris, 2015-07-22 23:59:59
BOB.jpg, London, 2015-08-05 00:02:03
notredame.png, Paris, 2015-09-01 12:00:00
me.jpg, Warsaw, 2013-09-06 15:40:22
a.png, Warsaw, 2016-02-13 13:33:50
b.jpg, Warsaw, 2016-01-02 15:12:22
c.jpg, Warsaw, 2016-01-02 14:34:30
d.jpg, Warsaw, 2016-01-02 15:15:01
e.png, Warsaw, 2016-01-02 09:49:09
f.png, Warsaw, 2016-01-02 10:55:32
g.jpg, Warsaw, 2016-02-29 22:13:11`

	fmt.Println(Solution(s))

}

func Solution(S string) string {
	// write your code in Go 1.4
	citymap := make(map[string][]photo)

	sSlice := strings.Split(S, "\n")
	i := 0
	for _, line := range sSlice {
		photoSlice := strings.Split(line, ",")
		p := photo{
			id:        i,
			name:      strings.Trim(photoSlice[0], " "),
			city:      strings.Trim(photoSlice[1], " "),
			photoTime: createDateTime(strings.Trim(photoSlice[2], " ")),
			extension: strings.Split(strings.Trim(photoSlice[0], " "), ".")[1],
		}
		photos = append(photos, p)
		i++
		citymap[p.city] = append(citymap[p.city], p)

	}
	cityKeys := make([]string, 0, len(citymap))
	for k := range citymap {
		cityKeys = append(cityKeys, k)
	}


	for _, key := range cityKeys {
		sort.Sort(ByDate(citymap[key]))

	}
	
	var photofiles []photo

	for _, u := range citymap {
		x:=1
		for _, p := range u {
			pWidth := len(strconv.Itoa(len(u)))
			p.sequence = x
			p.paddingWidth = pWidth
		
			
			x++

			photofiles = append(photofiles, p)
		}

	}
	var finalNames []string
	sort.Sort(ById(photofiles))
	for _, ph := range photofiles {
		ph.updatedName = fmt.Sprintf("%s%0*d.%s", ph.city,ph.paddingWidth, ph.sequence, ph.extension)
		finalNames = append(finalNames,ph.updatedName)
		
	}
	return strings.Join(finalNames,"\n")
}
