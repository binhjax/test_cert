package main

import (
	"io/ioutil"
	// "net/http"
	"fmt"
	"log"
	"crypto/x509"
	"crypto/tls"
	// "github.com/kevinburke/handlers"
)

func main() {
			fmt.Println("Start go client ")
      rootPEM, err := ioutil.ReadFile("cert/root.pem")
      if err != nil {
      	log.Fatal(err)
      }
      roots := x509.NewCertPool()
      ok := roots.AppendCertsFromPEM(rootPEM)
      if !ok {
      	panic("failed to parse root certificate")
      }

      // Use the tls.Config here in http.Transport.TLSClientConfig
      conn, err := tls.Dial("tcp", "localhost:7252", &tls.Config{
          RootCAs: roots,
      })
      if err != nil {
          panic("failed to connect: " + err.Error())
      }
      defer conn.Close()

			n, err := conn.Write([]byte("hello\n"))
			if err != nil {
					 log.Println(n, err)
					 return
			 }

			 buf := make([]byte, 100)
			 n, err = conn.Read(buf)
			 if err != nil {
					 log.Println(n, err)
					 return
			 }

			 println(string(buf[:n]))
}
