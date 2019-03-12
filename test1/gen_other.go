package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/tls"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"log"
	"math/big"
	"os"
	"time"
)

func main() {

	// Load CA
	catls, err := tls.LoadX509KeyPair("cert/root.pem", "cert/root.key")
	if err != nil {
		panic(err)
	}
	ca, err := x509.ParseCertificate(catls.Certificate[0])
	if err != nil {
		panic(err)
	}

	// Prepare certificate
	cert := &x509.Certificate{
		SerialNumber: big.NewInt(1658),
		Subject: pkix.Name{
			Organization:  []string{"Vnpay"},
			Country:       []string{"VN"},
			Province:      []string{"Hanoi"},
			Locality:      []string{"Hanoi"},
			StreetAddress: []string{"22 Lang ha "},
			PostalCode:    []string{"10000"},
		},
		NotBefore:    time.Now(),
		NotAfter:     time.Now().AddDate(10, 0, 0),
		SubjectKeyId: []byte{1, 2, 3, 4, 6},
		ExtKeyUsage:  []x509.ExtKeyUsage{x509.ExtKeyUsageClientAuth, x509.ExtKeyUsageServerAuth},
		KeyUsage:     x509.KeyUsageDigitalSignature,
	}
	priv, _ := rsa.GenerateKey(rand.Reader, 2048)
	pub := &priv.PublicKey

	// Sign the certificate
	cert_b, err := x509.CreateCertificate(rand.Reader, cert, ca, pub, catls.PrivateKey)

	// Public key
	certOut, err := os.Create("cert/node1.pem")
	pem.Encode(certOut, &pem.Block{Type: "CERTIFICATE", Bytes: cert_b})
	certOut.Close()
	log.Print("written node1.pem\n")

	// Private key
	keyOut, err := os.OpenFile("cert/node1.key", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0600)
	pem.Encode(keyOut, &pem.Block{Type: "RSA PRIVATE KEY", Bytes: x509.MarshalPKCS1PrivateKey(priv)})
	keyOut.Close()
	log.Print("written node1.pem\n")

}
