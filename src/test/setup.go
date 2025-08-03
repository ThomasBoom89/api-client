package test

import (
	"bytes"
	"context"
	"crypto/ed25519"
	"crypto/tls"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"math/big"
	"net"
	"os"
	"time"
)

type UserDir struct {
	Dir string
}

func (U *UserDir) GetConfigPath() string {
	return U.Dir
}
func (U *UserDir) GetDataPath() string {
	return U.Dir
}

func (U *UserDir) Cleanup() {
	err := os.RemoveAll(U.Dir)
	if err != nil {
		panic(err)
	}
}

type Event struct {
	receive chan []interface{}
}

func NewEvent(receive chan []interface{}) *Event {
	return &Event{receive: receive}
}

func (E *Event) EventsEmit(ctx context.Context, eventName string, optionalData ...interface{}) {
	E.receive <- optionalData
}

func CreateTLSCertificate() []tls.Certificate {
	// rootca
	capub, capriv, err := ed25519.GenerateKey(nil)
	if err != nil {
		panic(err)
	}
	ca := &x509.Certificate{
		SerialNumber: big.NewInt(1234),
		Subject: pkix.Name{
			Country:            nil,
			Organization:       nil,
			OrganizationalUnit: nil,
			Locality:           nil,
			Province:           nil,
			StreetAddress:      nil,
			PostalCode:         nil,
			SerialNumber:       "",
			CommonName:         "",
			Names:              nil,
			ExtraNames:         nil,
		},
		NotBefore:             time.Now(),
		NotAfter:              time.Now().Add(time.Hour),
		IsCA:                  true,
		ExtKeyUsage:           []x509.ExtKeyUsage{x509.ExtKeyUsageClientAuth, x509.ExtKeyUsageServerAuth},
		KeyUsage:              x509.KeyUsageCertSign | x509.KeyUsageDigitalSignature,
		BasicConstraintsValid: true,
	}

	_, err = x509.CreateCertificate(nil, ca, ca, capub, capriv)
	if err != nil {
		panic(err)
	}

	// ssl cert
	sslpub, sslpriv, err := ed25519.GenerateKey(nil)
	if err != nil {
		panic(err)
	}
	ssl := &x509.Certificate{
		SerialNumber: big.NewInt(12334),
		Subject: pkix.Name{
			Country:            nil,
			Organization:       nil,
			OrganizationalUnit: nil,
			Locality:           nil,
			Province:           nil,
			StreetAddress:      nil,
			PostalCode:         nil,
			SerialNumber:       "",
			CommonName:         "",
			Names:              nil,
			ExtraNames:         nil,
		},
		NotBefore:    time.Now(),
		NotAfter:     time.Now().Add(time.Hour),
		IPAddresses:  []net.IP{net.IPv4(127, 0, 0, 1)},
		SubjectKeyId: []byte{1, 2, 3, 4, 6},
		ExtKeyUsage:  []x509.ExtKeyUsage{x509.ExtKeyUsageClientAuth, x509.ExtKeyUsageServerAuth},
		KeyUsage:     x509.KeyUsageDigitalSignature,
	}

	sslBytes, err := x509.CreateCertificate(nil, ssl, ca, sslpub, capriv)
	if err != nil {
		panic(err)
	}
	certPEM := new(bytes.Buffer)
	pem.Encode(certPEM, &pem.Block{
		Type:  "CERTIFICATE",
		Bytes: sslBytes,
	})

	privBytes, err := x509.MarshalPKCS8PrivateKey(sslpriv)
	if err != nil {
		panic(err)
	}
	certPrivKeyPEM := new(bytes.Buffer)
	pem.Encode(certPrivKeyPEM, &pem.Block{
		Type:  "ED25519 PRIVATE KEY",
		Bytes: privBytes,
	})

	serverCert, err := tls.X509KeyPair(certPEM.Bytes(), certPrivKeyPEM.Bytes())
	if err != nil {
		panic(err)
	}

	return []tls.Certificate{serverCert}
}
