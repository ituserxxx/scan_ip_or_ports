package main

import (
	"fmt"
	"golang.org/x/net/icmp"
	"golang.org/x/net/ipv4"
	"net"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"
)

var succIp = make(chan string)
var errIp = make(chan string)
var wg sync.WaitGroup
var chanlNum = make(chan int, 20)

func main() {
	var iptmp = "*.0.1.1"
	go printSucc()
	go printErr()
	for _, ipstr := range getCheckIpList(iptmp) {
		println(ipstr)
		//wg.Add(1)
		//go addWork(ipstr, &wg)
	}
	wg.Wait()

}
func printSucc() {
	for {
		select {
		case v := <-succIp:
			fmt.Printf("\n succ ip=%s", v)
		}
	}
}
func printErr() {
	for {
		select {
		case v := <-errIp:
			fmt.Printf("\n err ip=%s", v)

		}
	}
}
func addWork(ipstr string, wg *sync.WaitGroup) {
	defer func() {
		wg.Done()
		<-chanlNum
	}()
	var find = make(chan int)
	ipAddr := net.ParseIP(ipstr)
	if ipAddr == nil {
		fmt.Printf("\nInvalid IP address err:%s", ipstr)
		return
	}
	var AheadExitAll = make(chan int)
	go func() {
		select {
		case <-AheadExitAll:
			return
		default:
			err := pingIp(ipAddr)
			if err != nil {
				fmt.Printf("\nping IP err:%s", err.Error())
				return
			}
			find <- 1
			return
		}
	}()
	select {
	case <-find:
		go func() { succIp <- ipstr }()
		return
	case <-time.After(time.Duration(3) * time.Second):
		go func() {
			errIp <- ipstr
			AheadExitAll <- 1
		}()
		return
	}
}

func getCheckIpList(ipTmp string) []string {
	ips := make([]string, 0)
	//if strings.Count(ipTmp, "*") == 0 {
	//	ips = append(ips, ipTmp)
	//} else {
	generateIP(&ips, "", strings.Split(ipTmp, "."))
	//}

	return ips
}

func generateIP(ips *[]string, currentIP string, parts []string) {
	if len(parts) == 0 {
		*ips = append(*ips, currentIP)
		return
	}
	part := parts[0]
	nextParts := parts[1:]
	if part == "*" {
		for i := 0; i <= 10; i++ {
			generateIP(ips, currentIP+strconv.Itoa(i)+".", nextParts)
		}
	} else {
		generateIP(ips, currentIP+part+".", nextParts)
	}
}

func pingIp(ipAddress net.IP) error {

	conn, err := net.Dial("ip4:icmp", ipAddress.String())
	if err != nil {
		println(1111)
		return err
	}
	defer conn.Close()
	// 构建 ICMP 报文
	msg := icmp.Message{
		Type: ipv4.ICMPTypeEcho,
		Code: 0,
		Body: &icmp.Echo{
			ID:   os.Getpid() & 0xffff,
			Seq:  1,
			Data: []byte("hello"),
		},
	}
	msgBytes, err := msg.Marshal(nil)
	if err != nil {
		println(2222)
		return err
	}
	// 发送 ICMP 报文
	_, err = conn.Write(msgBytes)
	if err != nil {
		println(3333)
		return err
	}
	// 接收 ICMP 回复报文
	reply := make([]byte, 1500)
	_, err = conn.Read(reply)
	if err != nil {
		println(44444)
		return err
	}

	// 解析 ICMP 回复报文
	_, err = icmp.ParseMessage(ipv4.ICMPTypeEchoReply.Protocol(), reply)
	if err != nil {
		println(5555)
		return err
	}
	//if replyMsg.Type != ipv4.ICMPTypeEchoReply {
	//	println(6666)
	//	return fmt.Errorf("unexpected ICMP message type %v", replyMsg.Type)
	//}
	//fmt.Printf("9999 ->%s", ipAddress.String())
	return nil
}
