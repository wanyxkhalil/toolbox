package httpsexpired

import (
	"crypto/tls"
	"fmt"
	"time"
)

func Run(host string) {
	notBefore, notAfter := callHost(host)
	fmt.Printf("notBefore=%s\nnoAfter=%s\n", notBefore, notAfter)
}

func callHost(host string) (notBefore, notAfter time.Time) {
	conn, err := tls.Dial("tcp", host+":443", nil)
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	chains := conn.ConnectionState().VerifiedChains
	for _, chain := range chains {
		for _, cert := range chain {
			notBefore = cert.NotBefore
			notAfter = cert.NotAfter
			return
		}
	}
	return
}
