package main

import (
	"fmt"
	"github.com/miekg/dns"
	"net"
	"strconv"
	"runtime"
	"syscall"
	"unsafe"
	"os"
	"golang.org/x/sys/windows"
)

func adapterAddresses() ([]*windows.IpAdapterAddresses, error) {
	var b []byte
	l := uint32(15000) // recommended initial size
	for {
		b = make([]byte, l)
		err := windows.GetAdaptersAddresses(syscall.AF_UNSPEC, windows.GAA_FLAG_INCLUDE_PREFIX, 0, (*windows.IpAdapterAddresses)(unsafe.Pointer(&b[0])), &l)
		if err == nil {
			if l == 0 {
				return nil, nil
			}
			break
		}
		if err.(syscall.Errno) != syscall.ERROR_BUFFER_OVERFLOW {
			return nil, os.NewSyscallError("getadaptersaddresses", err)
		}
		if l <= uint32(len(b)) {
			return nil, os.NewSyscallError("getadaptersaddresses", err)
		}
	}
	var aas []*windows.IpAdapterAddresses
	for aa := (*windows.IpAdapterAddresses)(unsafe.Pointer(&b[0])); aa != nil; aa = aa.Next {
		aas = append(aas, aa)
	}
	return aas, nil
}

func main() {
	conf, err := dns.ClientConfigFromFile("/etc/resolv.conf")
	if err != nil {
		fmt.Printf("err: %+v \n", err)
	}
	fmt.Printf("%+v \n", conf)
	fmt.Println("=========================")

	if runtime.GOOS == "windows" {
		servers, err := adapterAddresses()
		if err != nil {
			fmt.Printf("err: %+v \n", err)
		}
		for _, v := range servers {
			fmt.Printf("%+v \n", v.FirstDnsServerAddress.Address.IP().String())
			// 168.63.129.16 
      // fec0:0:0:ffff::1 
      // fec0:0:0:ffff::1 
		}
	}

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
