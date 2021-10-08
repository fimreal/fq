package main

// Forward from local port to remote port
// 代理远端端口到本地
// 例如 实现代理服务器上 clash HTTP 内网代理端口 7890 到本地

import (
	"io"
	"log"
	"net"

	"golang.org/x/crypto/ssh"
)

// 需要手动修改配置，然后 make 编译
const (
	username         = ""
	password         = ""
	serverAddrString = "ip:22"
	localAddrString  = "127.0.0.1:7890"
	remoteAddrString = "172.17.0.1:7890"
)

func forward(localConn net.Conn, config *ssh.ClientConfig) {
	// Setup sshClientConn (type *ssh.ClientConn)
	sshClientConn, err := ssh.Dial("tcp", serverAddrString, config)
	if err != nil {
		log.Fatalf("ssh.Dial failed: %s", err)
	}

	// Setup sshConn (type net.Conn)
	sshConn, err := sshClientConn.Dial("tcp", remoteAddrString)

	// Copy localConn.Reader to sshConn.Writer
	go func() {
		_, err = io.Copy(sshConn, localConn)
		if err != nil {
			log.Fatalf("io.Copy failed: %v", err)
		}
	}()

	// Copy sshConn.Reader to localConn.Writer
	go func() {
		_, err = io.Copy(localConn, sshConn)
		if err != nil {
			log.Fatalf("io.Copy failed: %v", err)
		}
	}()
}

func main() {
	// switch runtime.GOOS {
	// case "windows":
	// 	setProxy("127.0.0.1:7890")
	// }

	log.Println("代理地址：http://" + localAddrString)
	// Setup SSH config (type *ssh.ClientConfig)
	config := &ssh.ClientConfig{
		User: username,
		Auth: []ssh.AuthMethod{
			ssh.Password(password),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}

	// Setup localListener (type net.Listener)
	localListener, err := net.Listen("tcp", localAddrString)
	if err != nil {
		log.Fatalf("net.Listen failed: %v", err)
	}

	for {
		// Setup localConn (type net.Conn)
		localConn, err := localListener.Accept()
		if err != nil {
			log.Fatalf("listen.Accept failed: %v", err)
		}
		go forward(localConn, config)
	}
}

// go get -u github.com/Trisia/gosysproxy
// func setProxy(proxyAddr string) {
// 	// 设置全局代理
// 	err := gosysproxy.SetGlobalProxy(proxyAddr)
// 	if err {
// 		panic(err)
// 	}

// 	defer func() {
// 		err := gosysproxy.Off()
// 		if err {
// 			panic(err)
// 		}
// 	}()
// }

