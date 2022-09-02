package main

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io/ioutil"
	"path/filepath"

	"github.com/gocql/gocql"
)

func main() {
	var _cqlshrc_host = "" // host from config.json in the SCB
	var _cqlshrc_port = "" // cql_port from config.json in the SCB
	var _username = "" // Astra ClientID, or the word "token"
	var _password = "" // Astra Client Secret, or Astra token if "token" was used as username

	// put the absolute paths of the following files in the unzipped SCB
	_certPath, _ := filepath.Abs("...<cluster_unzipped_scb>/cert")
	_keyPath, _ := filepath.Abs("...<cluster_unzipped_scb>/key")
	_caPath, _ := filepath.Abs("...<cluster_unzipped_scb>/ca.crt")
	_cert, err := tls.LoadX509KeyPair(_certPath, _keyPath)
	if err != nil {
		fmt.Printf("Certificate could not be created from the provided cert and key files due to %v", err)
		return
	}
	_caCert, err := ioutil.ReadFile(_caPath)
	if err != nil {
		fmt.Printf("Certificate could not be created from the provided cert and key files due to %v", err)
		return
	}
	_caCertPool := x509.NewCertPool()
	ok := _caCertPool.AppendCertsFromPEM(_caCert)
	if !ok {
		fmt.Println("the provided CA cert could not be added to the rootCAs")
		return
	}

	_tlsConfig := getTlsConfigFromParsedCerts(_caCertPool, _cert, _cqlshrc_host)

	_cluster := gocql.NewCluster(_cqlshrc_host)
	_cluster.SslOpts = &gocql.SslOptions{
		Config:                 _tlsConfig,
		EnableHostVerification: false,
	}
	_cluster.Authenticator = gocql.PasswordAuthenticator{
		Username: _username,
		Password: _password,
	}
	_cluster.Hosts = []string{_cqlshrc_host + ":" + _cqlshrc_port}

	session, err := _cluster.CreateSession()
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	var _rank int
	var _city string
	var _country string
	var _query = "SELECT rank, city, country FROM community.cities_by_rank WHERE rank IN ( 1, 2, 3, 4, 5 )"

	fmt.Println("According to independent.co.uk, the top 5 most liveable cities in 2019 were:")
	iter := session.Query(_query).Iter()
	for iter.Scan(&_rank, &_city, &_country) {
		fmt.Printf("\tRank %d: %s, %s\n", _rank, _city, _country)
	}
}

func getTlsConfigFromParsedCerts(rootCAs *x509.CertPool, clientCert tls.Certificate, serverName string) *tls.Config {

	tlsConfig := &tls.Config{
		RootCAs:            rootCAs,
		Certificates:       []tls.Certificate{clientCert},
		ServerName:         serverName,
		InsecureSkipVerify: false,
		VerifyConnection:   getVerifyConnectionCallback(serverName, rootCAs),
	}

	return tlsConfig
}

func getVerifyConnectionCallback(certificateDnsName string, rootCAs *x509.CertPool) func(cs tls.ConnectionState) error {
	return func(cs tls.ConnectionState) error {
		dnsName := cs.ServerName
		if certificateDnsName != "" {
			dnsName = certificateDnsName
		}
		opts := x509.VerifyOptions{
			DNSName:       dnsName,
			Roots:         rootCAs,
		}
		if len(cs.PeerCertificates) > 0 {
			opts.Intermediates = x509.NewCertPool()
			for _, cert := range cs.PeerCertificates[1:] {
				opts.Intermediates.AddCert(cert)
			}
		}
		_, err := cs.PeerCertificates[0].Verify(opts)
		return err
	}
}
