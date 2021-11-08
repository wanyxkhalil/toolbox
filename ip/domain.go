package ip

import (
	"fmt"
	"net"
)

func Domain(s string) {
	ip, err := net.ResolveIPAddr("ip", s)
	if err != nil {
		panic(err)
	}
	fmt.Println(ip.String())
}
