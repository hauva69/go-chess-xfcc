// Package xfcc provides methods for communicating correspondence chess sites with XFCC protocol.
//
// POST /xfccbasic.asmx HTTP/1.1
// Host: www.iccf.com
// Content-Type: text/xml; charset=utf-8
// Content-Length: length
// SOAPAction: "http://www.bennedik.com/webservices/XfccBasic/GetMyGames"
//
// <?xml version="1.0" encoding="utf-8"?>
// <soap:Envelope xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xmlns:xsd="http://www.w3.org/2001/XMLSchema" xmlns:soap="http://schemas.xmlsoap.org/soap/envelope/">
//   <soap:Body>
//     <GetMyGames xmlns="http://www.bennedik.com/webservices/XfccBasic">
//       <username>string</username>
//       <password>string</password>
//     </GetMyGames>
//   </soap:Body>
// </soap:Envelope>
package xfcc

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
)

// FewPiecesGame2b is something to be fixed
const FewPiecesGame2b = `1.d4 {[%ccsnt 2018.10.02,12:56]} d5 {[%ccsnt 2018.10.02,13:34]} 
2.c4 {[%ccsnt 2018.10.03,04:20]} c6 {[%ccsnt 2018.10.06,04:23]} 
3.cxd5 {[%ccsnt 2018.10.06,11:00]} cxd5 {[%ccsnt 2018.10.06,11:08]} 
4.Bf4 {[%ccsnt 2018.10.06,20:54]} Nc6 {[%ccsnt 2018.10.11,19:56]} 
5.e3 {[%ccsnt 2018.10.13,07:15]} Qb6 {[%ccsnt 2018.10.14,10:28]} 
6.Nc3 {[%ccsnt 2018.10.14,11:13]} Nf6 {[%ccsnt 2018.10.23,10:42]} 
7.Bd3 {[%ccsnt 2018.10.23,17:47]} a6 {[%ccsnt 2018.10.30,15:40]} 
8.Nge2 {[%ccsnt 2018.10.31,17:26]} e6 {[%ccsnt 2018.11.23,04:22]} 
9.O-O {[%ccsnt 2018.11.23,08:24]} Be7 {[%ccsnt 2018.11.25,05:35]} 
10.Na4 {[%ccsnt 2018.11.25,17:00]} Qd8 {[%ccsnt 2018.12.04,20:01]} 
11.h3 {[%ccsnt 2018.12.05,07:04]} Nb4 {[%ccsnt 2018.12.26,13:43]} 
12.Bb1 {[%ccsnt 2018.12.27,07:57]} Nc6 {[%ccsnt 2018.12.27,15:33]} 
13.Bd3 {[%ccsnt 2018.12.27,17:07]} Nb4 {[%ccsnt 2018.12.27,20:54]} 
14.Bb1 {[%ccsnt 2018.12.28,08:09]} Nc6 {[%ccsnt 2018.12.28,09:00]} 
15.a3 {[%ccsnt 2018.12.28,17:47]} Na5 {[%ccsnt 2018.12.29,12:17]} 
16.Qc2 {[%ccsnt 2018.12.29,19:36]} Nc4 {[%ccsnt 2018.12.29,22:50]} 
17.b3 {[%ccsnt 2019.01.01,13:36]} Nd6 {[%ccsnt 2019.01.03,16:59]} 
18.Rc1 {[%ccsnt 2019.01.03,20:03]} Bd7 {[%ccsnt 2019.01.03,21:15]} 
19.Nc5 {[%ccsnt 2019.01.04,07:16]} Bc6 {[%ccsnt 2019.01.04,12:17]} 
20.Qb2 {[%ccsnt 2019.01.04,18:02]} O-O {[%ccsnt 2019.01.04,19:52]} 
21.Nd3 {[%ccsnt 2019.01.04,21:24]} Qb6 {[%ccsnt 2019.01.05,00:38]} 
22.a4 {[%ccsnt 2019.01.06,11:12]} a5 {[%ccsnt 2019.01.06,20:42]} 
23.g4 {[%ccsnt 2019.01.07,09:28]} Rfc8 {[%ccsnt 2019.01.07,10:07]} 
24.f3 {[%ccsnt 2019.01.09,06:01]} Be8 {[%ccsnt 2019.01.09,12:02]} 
25.Nc5 {[%ccsnt 2019.01.10,20:16]} e5 {[%ccsnt 2019.01.11,10:43]} 
26.dxe5 {[%ccsnt 2019.01.11,10:55]} Nc4 {[%ccsnt 2019.01.11,11:14]} 
27.Qc3 {[%ccsnt 2019.01.11,16:54]} Nxe5 {[%ccsnt 2019.01.11,19:27]} 
28.Bxe5 {[%ccsnt 2019.01.12,19:37]} Bxc5 {[%ccsnt 2019.01.12,19:56]} 
29.Bd4 {[%ccsnt 2019.01.13,17:29]} Bxd4 {[%ccsnt 2019.01.13,18:02]} 
30.Qxd4 {[%ccsnt 2019.01.14,20:21]} Qxb3 {[%ccsnt 2019.01.14,20:35]} 
31.Rxc8 {[%ccsnt 2019.01.15,09:55]} Rxc8 {[%ccsnt 2019.01.15,11:12]} 
32.g5 {[%ccsnt 2019.01.16,06:35]} Nd7 {[%ccsnt 2019.01.16,06:38]} 
33.Ba2 {[%ccsnt 2019.01.16,07:59]} Qc2 {[%ccsnt 2019.01.16,21:05]} 
34.Nf4 {[%ccsnt 2019.01.17,07:47]} Qc7 {[%ccsnt 2019.01.17,09:58]} 
35.Kg2 {[%ccsnt 2019.01.18,07:04]} Kf8 {[%ccsnt 2019.01.18,15:08]} 
36.Bxd5 {[%ccsnt 2019.01.18,16:30]} Nb6 {[%ccsnt 2019.01.18,17:11]} 
37.Be4 {[%ccsnt 2019.01.18,18:03]} Rd8 {[%ccsnt 2019.01.19,03:24]} 
38.Rc1 {[%ccsnt 2019.01.20,20:26]} Qxc1 {[%ccsnt 2019.01.21,05:40]} 
39.Qxd8 {[%ccsnt 2019.01.22,16:56]} Nc4 {[%ccsnt 2019.01.22,19:37]} 
40.Nd5 {[%ccsnt 2019.01.23,07:02]} Qb2+ {[%ccsnt 2019.01.23,08:20]} 
41.Kg3 {[%ccsnt 2019.01.24,18:41]} Qe5+ {[%ccsnt 2019.01.24,19:36]} 
42.Kg2 {[%ccsnt 2019.01.25,07:55]} Qd6 {[%ccsnt 2019.01.25,13:59]} 
43.Qxd6+ {[%ccsnt 2019.01.26,16:55]} Nxd6 {[%ccsnt 2019.01.26,17:14]} 
44.Bc2 {[%ccsnt 2019.01.28,06:50]} b5 {[%ccsnt 2019.01.28,18:53]} 
45.Nb6 {[%ccsnt 2019.01.29,10:17]} bxa4 {[%ccsnt 2019.02.03,08:53]} 
46.Nxa4 {[%ccsnt 2019.02.04,15:23]} Nc4 {[%ccsnt 2019.02.04,17:29]} 
47.Kf2 {[%ccsnt 2019.02.05,11:35]} Bd7 {[%ccsnt 2019.02.05,19:26]} 
48.h4 {[%ccsnt 2019.02.06,16:38]} Bxa4 {[%ccsnt 2019.02.06,18:14]} 
49.Bxa4 {[%ccsnt 2019.02.07,18:09]} Nb6 {[%ccsnt 2019.02.07,18:24]} 
50.Bb3 {[%ccsnt 2019.02.08,07:29]} a4 {[%ccsnt 2019.02.08,08:59]} 
51.Ba2 {[%ccsnt 2019.02.08,16:03]} Nc8 {[%ccsnt 2019.02.08,19:15]} 
52.Ke1 {[%ccsnt 2019.02.10,16:13]} Nd6 {[%ccsnt 2019.02.10,17:44]} 
53.Bb1 {[%ccsnt 2019.02.11,15:10]} a3 {[%ccsnt 2019.02.11,15:29]} 
54.Kd2 {[%ccsnt 2019.02.12,06:49]} Nf5 {[%ccsnt 2019.02.12,09:54]} 
55.h5 {[%ccsnt 2019.02.12,11:23]} h6 {[%ccsnt 2019.02.12,12:07]} 
56.f4 {[%ccsnt 2019.02.13,12:15]} Ng3 {[%ccsnt 2019.02.13,13:03]} 
57.Kc3 {[%ccsnt 2019.02.14,13:51]} Nxh5 {[%ccsnt 2019.02.14,15:15]} 
58.Bf5 {[%ccsnt 2019.02.14,19:37]} Ng3 {[%ccsnt 2019.02.14,20:00]} 
59.gxh6 {[%ccsnt 2019.02.15,06:54]} gxh6 {[%ccsnt 2019.02.15,06:54]} 
60.Bh3 {[%ccsnt 2019.02.15,06:56]} Ne4+ {[%ccsnt 2019.02.15,07:51]} 
61.Kb3 {[%ccsnt 2019.02.15,11:24]} h5 {[%ccsnt 2019.02.15,11:49]} 
62.Kxa3 {[%ccsnt 2019.02.16,07:05]} Nf2 {[%ccsnt 2019.02.16,08:35]} 
63.Bf1 {[%ccsnt 2019.02.16,14:41]} h4 {[%ccsnt 2019.02.16,14:54]} 
64.Kb3 {[%ccsnt 2019.02.16,18:00]} h3 {[%ccsnt 2019.02.16,20:07]} 
65.Bxh3 {[%ccsnt 2019.02.17,17:10]} Nxh3 {[%ccsnt 2019.02.17,17:33]} 
66.Kc2 {[%ccsnt 2019.02.18,16:26]} Ke7 {[%ccsnt 2019.02.18,18:09]} 
0-1`

// Envelope is just for parsing the XML.
type Envelope struct {
	XMLName xml.Name
	Body    Body
}

// Body is just for parsing the XML.
type Body struct {
	XMLName  xml.Name
	Response GetMyGamesResponse `xml:"GetMyGamesResponse"`
}

// GetMyGamesResponse models the result of GetMyGames SOAP query.
type GetMyGamesResponse struct {
	XMLName xml.Name         `xml:"GetMyGamesResponse"`
	Result  GetMyGamesResult `xml:"GetMyGamesResult"`
}

// GetMyGamesResult models the result containg the games.
type GetMyGamesResult struct {
	XMLName xml.Name `xml:"GetMyGamesResult"`
	Games   []Game   `xml:"XfccGame"`
}

// GetMyGamesXML returns the XML of your games.
func GetMyGamesXML(url, MIMEType, username, password string) ([]byte, error) {
	data := fmt.Sprintf(GetMyGamesSOAPXML, username, password)
	buffer := bytes.NewBufferString(data)
	resp, err := http.Post(url, MIMEType, buffer)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

// GetMyGames returns your games.
func GetMyGames(url, MIMEType, username, password string) ([]Game, error) {
	body, err := GetMyGamesXML(url, MIMEType, username, password)
	if err != nil {
		return nil, err
	}

	var result Envelope
	if err = xml.Unmarshal(body, &result); err != nil {
		return nil, err
	}

	// TODO this is ugly. Very, very ugly. Could it be implemented with a sensible XML tag?
	// Get rid of all the extra data types.
	return result.Body.Response.Result.Games, nil
}

// GetMyGamesSOAPXML is the SOAP request template for  GetMyGames.
const GetMyGamesSOAPXML = `<?xml version="1.0" encoding="utf-8"?>
<soap:Envelope xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xmlns:xsd="http://www.w3.org/2001/XMLSchema" xmlns:soap="http://schemas.xmlsoap.org/soap/envelope/">
  <soap:Body>
    <GetMyGames xmlns="http://www.bennedik.com/webservices/XfccBasic">
      <username>%s</username>
      <password>%s</password>
    </GetMyGames>
  </soap:Body>
</soap:Envelope>
`
