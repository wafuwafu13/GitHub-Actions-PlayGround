package main

import (
	"fmt"
	"net"
	"os"
	"runtime"
	"strconv"

	"github.com/miekg/dns"
	// "syscall"
	// "unsafe"
	// "os"
	// "golang.org/x/sys/windows"
)

// func adapterAddresses() ([]*windows.IpAdapterAddresses, error) {
// 	var b []byte
// 	l := uint32(15000) // recommended initial size
// 	for {
// 		b = make([]byte, l)
// 		err := windows.GetAdaptersAddresses(syscall.AF_UNSPEC, windows.GAA_FLAG_INCLUDE_PREFIX, 0, (*windows.IpAdapterAddresses)(unsafe.Pointer(&b[0])), &l)
// 		if err == nil {
// 			if l == 0 {
// 				return nil, nil
// 			}
// 			break
// 		}
// 		if err.(syscall.Errno) != syscall.ERROR_BUFFER_OVERFLOW {
// 			return nil, os.NewSyscallError("getadaptersaddresses", err)
// 		}
// 		if l <= uint32(len(b)) {
// 			return nil, os.NewSyscallError("getadaptersaddresses", err)
// 		}
// 	}
// 	var aas []*windows.IpAdapterAddresses
// 	for aa := (*windows.IpAdapterAddresses)(unsafe.Pointer(&b[0])); aa != nil; aa = aa.Next {
// 		aas = append(aas, aa)
// 	}
// 	return aas, nil
// }

func main() {
	var nameserver string
	if runtime.GOOS == "windows" {
		// servers, err := adapterAddresses()
		// if err != nil {
		// 	fmt.Printf("err: %+v \n", err)
		// }
		// for _, v := range servers {
		// 	fmt.Printf("%+v \n", v.FirstDnsServerAddress.Address.IP().String())
		// 	// 168.63.129.16 
    //   // fec0:0:0:ffff::1 
    //   // fec0:0:0:ffff::1 
		// }
		// nameserver = servers[0].FirstDnsServerAddress.Address.IP().String()
	} else {
		conf, err := dns.ClientConfigFromFile("/etc/resolv.conf")
	  if err != nil {
	  	fmt.Printf("err: %+v \n", err)
	  }
		nameserver = conf.Servers[0]
	}
	nameserver = "8.8.8.8"
	nameserver = net.JoinHostPort(nameserver, strconv.Itoa(53))
	fmt.Printf("env -> %v \n", os.Getenv("RUN_TEST_ON_GITHUB_ACTIONS"))
	fmt.Printf("nameserver -> %s \n", nameserver)

	c := new(dns.Client)
	m := &dns.Msg{
		MsgHdr: dns.MsgHdr{
			Opcode:            dns.OpcodeQuery,
		},
		Question: []dns.Question{{Name: dns.Fqdn("exampleee.com"), Qtype: dns.TypeA, Qclass: uint16(dns.ClassINET)}},
	}
	r, _, err := c.Exchange(m, nameserver)

	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("r -> %v", r)
	fmt.Println("=========================")
	fmt.Printf("r.Answer -> %+v \n", r.Answer)
}
