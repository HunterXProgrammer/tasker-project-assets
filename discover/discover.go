package main

import (
    "bufio"
    "encoding/json"
    "fmt"
    "net"
    "net/http"
    "net/url"
    "os"
    "strings"
    "time"
)

type SSDPResponse struct {
    Location string `json:"location"`
}

type ServiceInfo struct {
    Location []string `json:"location"`
    Host     []string `json:"host"`
    Port     []string `json:"port"`
}

func Discover(service string, timeout time.Duration, retries int, mx int) ServiceInfo {
    group := &net.UDPAddr{IP: net.ParseIP("239.255.255.250"), Port: 1900}
    message := fmt.Sprintf(
        "M-SEARCH * HTTP/1.1\r\n"+
            "HOST: %s:%d\r\n"+
            "MAN: \"ssdp:discover\"\r\n"+
            "ST: %s\r\n"+
            "MX: %d\r\n"+
            "\r\n",
        group.IP.String(), group.Port, service, mx,
    )
    responses := []SSDPResponse{}
    for i := 0; i < retries; i++ {
        addr, err := net.ResolveUDPAddr("udp", ":0")
        if err != nil {
            continue
        }
        conn, err := net.ListenUDP("udp", addr)
        if err != nil {
            continue
        }
        if err := conn.SetDeadline(time.Now().Add(timeout)); err != nil {
            continue
        }
        _, err = conn.WriteTo([]byte(message), group)
        if err != nil {
            continue
        }
        for {
            b := make([]byte, 2048)
            if err := conn.SetDeadline(time.Now().Add(timeout)); err != nil {
                break
             }
            n, _, err := conn.ReadFrom(b)
            if err != nil {
                break
            }
            resp, err := http.ReadResponse(bufio.NewReader(strings.NewReader(string(b[:n]))), nil)
            if err != nil {
                continue
            }
            if !strings.EqualFold(resp.Header.Get("ST"), service) {
                continue
            }
            responses = append(responses, SSDPResponse{Location: resp.Header.Get("Location")})
        }
    }
    // Extract scheme, host, and port from location URLs
    locations := []string{}
    schemes := []string{}
    hosts := []string{}
    ports := []string{}
    for _, resp := range responses {
        u, err := url.Parse(resp.Location)
        if err != nil {
            continue
        }
        locations = append(locations, resp.Location)
        schemes = append(schemes, u.Scheme)
        hosts = append(hosts, u.Hostname())
        ports = append(ports, u.Port())
    }
    return ServiceInfo{
        Location: locations,
        Scheme:   schemes,
        Host:     hosts,
        Port:     ports,
    }
}

func main() {
	if len(os.Args) != 2 {
		fmt.Printf("Usage: %s <service> (eg:- ssdp:all)\n", os.Args[0])
		os.Exit(1)
	}
	service := os.Args[1]
	info := Discover(service, 5*time.Second, 1, 3)
	data, err := json.Marshal(info)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(string(data))
}
