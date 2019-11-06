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
