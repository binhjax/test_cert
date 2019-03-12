package main

import (
    "crypto/rand"
    "crypto/rsa"
    "crypto/x509"
    "crypto/x509/pkix"
    "encoding/asn1"
    "encoding/pem"
    "os"
    "log"
)

var oidEmailAddress = asn1.ObjectIdentifier{1, 2, 840, 113549, 1, 9, 1}

func main() {
    priv, _ := rsa.GenerateKey(rand.Reader, 2048)
    //pub := &priv.PublicKey

    os.MkdirAll("cert", os.ModePerm);

    // Private key
  	keyOut, _ := os.OpenFile("cert/private.key", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0600)
  	pem.Encode(keyOut, &pem.Block{Type: "RSA PRIVATE KEY", Bytes: x509.MarshalPKCS1PrivateKey(priv)})
  	keyOut.Close()
  	log.Print("written key.pem\n")

    val, err := asn1.Marshal(basicConstraints{true, 0})
  	if err != nil {
  		return nil, err
  	}

    emailAddress := "test@example.com"
    subj := pkix.Name{
        CommonName:         "localhost",
        Country:            []string{"VN"},
        Province:           []string{"Hanoi"},
        Locality:           []string{"Hanoi"},
        Organization:       []string{"Vnpay"},
        OrganizationalUnit: []string{"IT"},
        ExtraNames: []pkix.AttributeTypeAndValue{
            {
                Type:  oidEmailAddress,
                Value: asn1.RawValue{
                    Tag:   asn1.TagIA5String,
                    Bytes: []byte(emailAddress),
                },
            },
        },
    }

    csrTemplate := x509.CertificateRequest{
        Subject:            subj,
        SignatureAlgorithm: x509.SHA256WithRSA,
        ExtraExtensions: []pkix.Extension{
                    			{
                    				Id:       asn1.ObjectIdentifier{2, 5, 29, 19},
                    				Value:    val,
                    				Critical: true,
                    			},
      		},
    }

    csrBytes, _ := x509.CreateCertificateRequest(rand.Reader, &csrTemplate, priv)

    // Public key
  	certOut, _ := os.Create("cert/ca.csr")
  	pem.Encode(certOut, &pem.Block{Type: "CERTIFICATE REQUEST", Bytes: csrBytes})
  	certOut.Close()
  	log.Print("written ca.csr\n")
}
