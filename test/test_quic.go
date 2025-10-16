package main

import (
	"context"
	"crypto/rand"
	"crypto/rsa"
	"crypto/tls"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"fmt"
	"math/big"
	"net"
	"strconv"
	"time"

	"github.com/quic-go/quic-go"
)

// 生成自签名证书
func generateTLSConfig() *tls.Config {
	key, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		panic(err)
	}
	template := x509.Certificate{
		SerialNumber: big.NewInt(1),
		Subject: pkix.Name{
			Organization:  []string{"Test"},
			Country:       []string{"US"},
			Province:      []string{""},
			Locality:      []string{"San Francisco"},
			StreetAddress: []string{""},
			PostalCode:    []string{""},
		},
		NotBefore:   time.Now(),
		NotAfter:    time.Now().Add(365 * 24 * time.Hour),
		KeyUsage:    x509.KeyUsageKeyEncipherment | x509.KeyUsageDigitalSignature,
		ExtKeyUsage: []x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth},
		IPAddresses: []net.IP{net.IPv4(127, 0, 0, 1)},
	}
	certDER, err := x509.CreateCertificate(rand.Reader, &template, &template, &key.PublicKey, key)
	if err != nil {
		panic(err)
	}
	keyPEM := pem.EncodeToMemory(&pem.Block{Type: "RSA PRIVATE KEY", Bytes: x509.MarshalPKCS1PrivateKey(key)})
	certPEM := pem.EncodeToMemory(&pem.Block{Type: "CERTIFICATE", Bytes: certDER})

	tlsCert, err := tls.X509KeyPair(certPEM, keyPEM)
	if err != nil {
		panic(err)
	}
	return &tls.Config{
		Certificates: []tls.Certificate{tlsCert},
	}
}

type TestQUIC struct {
	name     string
	listener *quic.Listener
	conn     *quic.Conn
	stream   *quic.Stream
}

func (t *TestQUIC) Start(name string, port int) error {
	t.name = name
	var err error
	addr := net.JoinHostPort("0.0.0.0", strconv.Itoa(port))
	tlsConfig := generateTLSConfig()
	t.listener, err = quic.ListenAddr(addr, tlsConfig, nil)
	if err != nil {
		panic(err)
	}
	for {
		conn, err := t.listener.Accept(context.Background())
		if err != nil {
			fmt.Println("--------------------------------")
			fmt.Println(name, "accept conn error", err)
			fmt.Println("--------------------------------")
			continue
		}
		t.conn = conn
		for {
			t.stream, err = t.conn.AcceptStream(context.Background())
			if err != nil {
				fmt.Println("--------------------------------")
				fmt.Println(name, "accept stream error", err)
				fmt.Println("--------------------------------")
				continue
			}
			go t.handleStream(t.stream)
		}
	}
}

func (t *TestQUIC) handleStream(stream *quic.Stream) {
	defer stream.Close()
	for {
		// 读取buffer 打印
		buf := make([]byte, 1024)
		_, err := stream.Read(buf)
		if err != nil {
			fmt.Println(t.name, "read buffer error", err)
			return
		}
		fmt.Println(t.name, "read buffer", string(buf))
	}
}

type QuicClient struct {
	conn   *quic.Conn
	stream *quic.Stream
}

func StartQuicClient(port int) {
	addr := net.JoinHostPort("127.0.0.1", strconv.Itoa(port))
	tlsConfig := &tls.Config{
		InsecureSkipVerify: true, // 跳过证书验证，仅用于测试
	}
	conn, err := quic.DialAddr(context.Background(), addr, tlsConfig, nil)
	if err != nil {
		panic("22222")
	}
	// open一个stream
	stream, err := conn.OpenStreamSync(context.Background())
	if err != nil {
		panic("33333")
	}
	go func() {
		for {
			// 读取buffer 打印
			buf := make([]byte, 1024)
			_, err := stream.Read(buf)
			if err != nil {
				fmt.Println("read buffer error", err)
				return
			}
			fmt.Println("read buffer", string(buf))
		}
	}()
	go func() {
		time.Sleep(5 * time.Second)
		//stream.Close()
		conn.CloseWithError(0, "client close")
	}()
	for {
		time.Sleep(1 * time.Second)
		// 格式化当前时间
		_, err := stream.Write([]byte(time.Now().Format("2006-01-02 15:04:05")))
		if err != nil {
			fmt.Println("write buffer error", err)
		}
	}
}

func main() {
	s := TestQUIC{}
	go s.Start("server", 34434)
	time.Sleep(1 * time.Second)
	StartQuicClient(34434)

	fmt.Println("123")
}
