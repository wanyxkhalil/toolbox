package ip

import (
	"errors"
	"fmt"
	"github.com/gocolly/colly/v2"
	"net"
	"strings"
)

const website = "https://www.ipaddress.my/"

func Local() {
	ip, country := internet()
	fmt.Printf("%s\n%s\n", ip, country)

	in := intranet()
	fmt.Printf("Intranet:\t\t%s\n", in)
}

func internet() (ip, country string) {
	c := colly.NewCollector()

	c.OnHTML(".table tbody", func(t *colly.HTMLElement) {
		if len(ip) > 0 && len(country) > 0 {
			return
		}
		t.ForEachWithBreak("tr", func(i int, tr *colly.HTMLElement) bool {
			if i == 0 {
				s := tr.ChildText("td")
				if strings.HasPrefix(s, "IP Address:") {
					ip = s
				}
			}
			if i == 5 {
				s := tr.ChildText("td")
				if strings.HasPrefix(s, "Country:") {
					country = s
					return false
				}
			}
			return true
		})
	})

	err := c.Visit(website)
	if err != nil {
		panic(err)
	}

	ip = strings.Replace(ip, ":", ":\t\t", 1)
	country = strings.Replace(country, ":", ":\t\t", 1)
	return
}

func intranet() string {
	netInterfaces, err := net.Interfaces()
	if err != nil {
		panic(err)
	}

	for i := 0; i < len(netInterfaces); i++ {
		if (netInterfaces[i].Flags & net.FlagUp) != 0 {
			addrs, _ := netInterfaces[i].Addrs()

			for _, address := range addrs {
				if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
					if ipnet.IP.To4() != nil {
						return ipnet.IP.String()
					}
				}
			}
		}
	}

	panic(errors.New("Not Found "))
}
