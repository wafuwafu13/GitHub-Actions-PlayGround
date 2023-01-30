package main

import (
	"fmt"
	"github.com/miekg/dns"
	"net"
	"strconv"
)

func main() {
	conf, err := dns.ClientConfigFromFile("/etc/resolv.conf")
	if err != nil {
		fmt.Printf("%+v \n", err)
	}
	fmt.Printf("%+v \n", conf)
	fmt.Println("=========================")

	nameserver := conf.Servers[0]
	nameserver = net.JoinHostPort(nameserver, strconv.Itoa(53))

	c := new(dns.Client)
	m := &dns.Msg{
		MsgHdr: dns.MsgHdr{
			Opcode:            dns.OpcodeQuery,
		},
		Question: []dns.Question{{Name: dns.Fqdn("hatenablog.com"), Qtype: dns.TypeA, Qclass: uint16(dns.ClassINET)}},
	}
	r, _, err := c.Exchange(m, nameserver)

	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%v", r)
	fmt.Println("=========================")
	fmt.Printf("%+v \n", r.Answer)
}
