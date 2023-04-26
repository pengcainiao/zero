package ssh

//
//import (
//	"bytes"
//	"io/ioutil"
//	"net"
//	"sync"
//
//	"golang.org/x/crypto/ssh"
//)
//
//// SSHClient 定义ssh客户端
//type SSHClient struct {
//	client *ssh.Client
//	m      sync.Mutex
//
//	Addr    string
//	User    string
//	Keyfile string
//}
//
//// connect 连接ssh客户端
//func (s *SSHClient) connect() error {
//	if s.client == nil {
//		s.m.Lock()
//		defer s.m.Unlock()
//
//		key, err := ioutil.ReadFile(s.Keyfile)
//		if err != nil {
//			return err
//		}
//
//		signer, err := ssh.ParsePrivateKey(key)
//		if err != nil {
//			return err
//		}
//
//		config := &ssh.ClientConfig{
//			User: s.User,
//			Auth: []ssh.AuthMethod{
//				ssh.PublicKeys(signer),
//			},
//			HostKeyCallback: ssh.HostKeyCallback(func(hostname string, remote net.Addr, key ssh.PublicKey) error { return nil }),
//		}
//
//		client, err := ssh.Dial("tcp", s.Addr, config)
//		if err != nil {
//			return err
//		}
//		s.client = client
//	}
//
//	return nil
//}
//
//// Command ssh执行命令
//func (s SSHClient) Command(command string) (string, error) {
//	if err := s.connect(); err != nil {
//		return "", err
//	}
//
//	session, err := s.client.NewSession()
//	if err != nil {
//		return "", err
//	}
//	defer session.Close()
//
//	var stdout bytes.Buffer
//	session.Stdout = &stdout
//	if err := session.Run(command); err != nil {
//		return "", err
//	}
//	return stdout.String(), nil
//}
