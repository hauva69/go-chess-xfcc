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

// FewPiecesGame1 is something to be fixed
// FIXME remove links to the test data
// FIXME move these to a separate test file
// https://www.iccf.com/game?id=1086957[Event "FIN/AC/4 (FIN)"[Date "2019.04.10"]
// [Round "-"]
// [White "Mäkelä, Ari"]
// [Black "Rajala, Olavi"]
// [Result "1-0"]
// [WhiteElo "2181"]
// [BlackElo "1995"]
const FewPiecesGame1 = `1.Nf3 e6 2.c4 Nf6 3.Nc3 c5 4.g3 Nc6 5.Bg2 Qb6 6.O-O Be7 7.b3 d5 8.Na4 Qa6 9.Bb2 O-O 10.Bxf6 Bxf6 11.cxd5 exd5 12.Nxc5 Qa5 13.Rc1 Bb2 14.Rb1 Ba3 15.d4 Bf5 16.Nd3 Rac8 17.Nh4 Bxd3 18.Qxd3 Nb4 19.Qf3 g6 20.Bh3 Rc2 21.Ng2 Rd8 22.Ne3 Rc6 23.Bg2 Qc7 24.h4 b5 25.h5 f5 26.Nc4 dxc4 27.bxc4 Rxc4 28.hxg6 hxg6 29.Qxa3 a5 30.Qf3 Nc2 31.d5 Kg7 32.Qd3 Qc5 33.Rfc1 Nd4 34.Rxc4 Qxc4 35.Qxc4 bxc4 36.Kf1 Kf6 37.Rb6+ Kf7 38.Rb7+ Kf6 39.e3 Nc2 40.a4 c3 41.Rc7 Rb8 42.Rxc3 Rb1+ 43.Ke2 Nb4 44.f4 Rb2+ 45.Kf1 Rb1+ 46.Kf2 Rb2+ 47.Kg1 Ra2 48.e4 fxe4 49.Bxe4 Rxa4 50.Rc4 Ra2 51.Rd4 Nc2 52.Rd1 Ke7 53.d6+ Kd8 54.Bc6 Nb4 55.Bb5 Rc2 56.Re1 Nc6 57.Re6 Rc5 58.Bf1 a4 59.Rxg6 Rc1 60.Kf2 Kc8 61.Bh3+ Kb7 62.Rg7+ Kb6 63.Rg8 a3 64.Ra8 Rc2+ 65.Ke3 Rc3+ 66.Ke4 Rxg3 67.Be6 Rg6 68.Kf5 Rg3 69.d7 Rd3 70.Ke4 Nb4 71.Rb8+ Kc5 72.Rxb4 Rxd7 73.Bxd7 Kxb4 1-0`

// FewPiecesGame1b is something to be fixed
const FewPiecesGame1b = `
1.Nf3 {[%ccsnt 2019.04.02,21:27]} e6 {[%ccsnt 2019.04.04,07:43]} 
2.c4 {[%ccsnt 2019.04.04,07:59]} Nf6 {[%ccsnt 2019.04.05,08:38]} 
3.Nc3 {[%ccsnt 2019.04.06,14:12]} c5 {[%ccsnt 2019.04.06,19:57]} 
4.g3 {[%ccsnt 2019.04.07,11:01]} Nc6 {[%ccsnt 2019.04.07,19:10]} 
5.Bg2 {[%ccsnt 2019.04.08,20:32]} Qb6 {[%ccsnt 2019.04.09,20:44]} 
6.O-O {[%ccsnt 2019.04.11,00:42]} Be7 {[%ccsnt 2019.04.11,19:29]} 
7.b3 {[%ccsnt 2019.05.07,06:30]} d5 {[%ccsnt 2019.05.07,19:40]} 
8.Na4 {[%ccsnt 2019.05.20,21:15]} Qa6 {[%ccsnt 2019.05.21,19:51]} 
9.Bb2 {[%ccsnt 2019.05.22,03:30]} O-O {[%ccsnt 2019.05.23,19:57]} 
10.Bxf6 {[%ccsnt 2019.05.28,15:43]} Bxf6 {[%ccsnt 2019.05.28,19:20]} 
11.cxd5 {[%ccsnt 2019.05.29,03:53]} exd5 {[%ccsnt 2019.05.29,08:27]} 
12.Nxc5 {[%ccsnt 2019.05.29,13:57]} Qa5 {[%ccsnt 2019.05.29,21:12]} 
13.Rc1 {[%ccsnt 2019.05.30,04:13]} Bb2 {[%ccsnt 2019.05.30,18:50]} 
14.Rb1 {[%ccsnt 2019.06.01,07:47]} Ba3 {[%ccsnt 2019.06.01,10:20]} 
15.d4 {[%ccsnt 2019.06.02,15:49]} Bf5 {[%ccsnt 2019.06.03,13:27]} 
16.Nd3 {[%ccsnt 2019.06.05,17:43]} Rac8 {[%ccsnt 2019.06.06,03:11]} 
17.Nh4 {[%ccsnt 2019.06.08,11:22]} Bxd3 {[%ccsnt 2019.06.09,03:49]} 
18.Qxd3 {[%ccsnt 2019.06.09,07:32]} Nb4 {[%ccsnt 2019.06.09,11:08]} 
19.Qf3 {[%ccsnt 2019.06.09,15:37]} g6 {[%ccsnt 2019.06.10,08:36]} 
20.Bh3 {[%ccsnt 2019.06.16,12:25]} Rc2 {[%ccsnt 2019.06.17,00:41]} 
21.Ng2 {[%ccsnt 2019.06.21,12:38]} Rd8 {[%ccsnt 2019.06.22,12:15]} 
22.Ne3 {[%ccsnt 2019.07.12,20:34]} Rc6 {[%ccsnt 2019.07.13,12:10]} 
23.Bg2 {[%ccsnt 2019.07.15,18:23]} Qc7 {[%ccsnt 2019.07.22,15:06]} 
24.h4 {[%ccsnt 2019.07.30,04:31]} b5 {[%ccsnt 2019.07.31,05:26]} 
25.h5 {[%ccsnt 2019.07.31,08:11]} f5 {[%ccsnt 2019.07.31,13:15]} 
26.Nc4 {[%ccsnt 2019.08.14,06:41]} dxc4 {[%ccsnt 2019.08.14,09:36]} 
27.bxc4 {[%ccsnt 2019.08.22,08:59]} Rxc4 {[%ccsnt 2019.08.22,13:14]} 
28.hxg6 {[%ccsnt 2019.08.22,16:45]} hxg6 {[%ccsnt 2019.08.23,04:46]} 
29.Qxa3 {[%ccsnt 2019.08.24,04:02]} a5 {[%ccsnt 2019.08.24,08:52]} 
30.Qf3 {[%ccsnt 2019.08.30,07:31]} Nc2 {[%ccsnt 2019.08.30,13:43]} 
31.d5 {[%ccsnt 2019.08.31,00:57]} Kg7 {[%ccsnt 2019.09.01,07:49]} 
32.Qd3 {[%ccsnt 2019.09.05,06:11]} Qc5 {[%ccsnt 2019.09.05,13:21]} 
33.Rfc1 {[%ccsnt 2019.09.09,12:08]} Nd4 {[%ccsnt 2019.09.09,17:34]} 
34.Rxc4 {[%ccsnt 2019.09.10,03:49]} Qxc4 {[%ccsnt 2019.09.10,11:14]} 
35.Qxc4 {[%ccsnt 2019.09.11,17:32]} bxc4 {[%ccsnt 2019.09.12,10:16]} 
36.Kf1 {[%ccsnt 2019.09.12,19:32]} Kf6 {[%ccsnt 2019.09.13,13:09]} 
37.Rb6+ {[%ccsnt 2019.09.16,15:56]} Kf7 {[%ccsnt 2019.09.17,07:38]} 
38.Rb7+ {[%ccsnt 2019.09.19,04:43]} Kf6 {[%ccsnt 2019.09.19,07:34]} 
39.e3 {[%ccsnt 2019.09.21,05:30]} Nc2 {[%ccsnt 2019.09.21,16:47]} 
40.a4 {[%ccsnt 2019.09.23,12:07]} c3 {[%ccsnt 2019.09.23,14:39]} 
41.Rc7 {[%ccsnt 2019.09.25,05:29]} Rb8 {[%ccsnt 2019.09.25,10:47]} 
42.Rxc3 {[%ccsnt 2019.09.27,18:06]} Rb1+ {[%ccsnt 2019.09.28,07:09]} 
43.Ke2 {[%ccsnt 2019.09.30,14:31]} Nb4 {[%ccsnt 2019.10.01,07:15]} 
44.f4 {[%ccsnt 2019.10.02,08:11]} Rb2+ {[%ccsnt 2019.10.02,10:05]} 
45.Kf1 {[%ccsnt 2019.10.02,18:17]} Rb1+ {[%ccsnt 2019.10.03,07:39]} 
46.Kf2 {[%ccsnt 2019.10.03,08:20]} Rb2+ {[%ccsnt 2019.10.03,11:40]} 
47.Kg1 {[%ccsnt 2019.10.03,12:26]} Ra2 {[%ccsnt 2019.10.03,15:31]} 
48.e4 {[%ccsnt 2019.10.04,04:46]} fxe4 {[%ccsnt 2019.10.04,11:51]} 
49.Bxe4 {[%ccsnt 2019.10.04,13:05]} Rxa4 {[%ccsnt 2019.10.04,15:38]} 
50.Rc4 {[%ccsnt 2019.10.05,09:36]} Ra2 {[%ccsnt 2019.10.06,12:33]} 
51.Rd4 {[%ccsnt 2019.10.06,16:31]} Nc2 {[%ccsnt 2019.10.06,18:36]} 
52.Rd1 {[%ccsnt 2019.10.07,10:37]} Ke7 {[%ccsnt 2019.10.07,15:10]} 
53.d6+ {[%ccsnt 2019.10.08,08:54]} Kd8 {[%ccsnt 2019.10.08,09:35]} 
54.Bc6 {[%ccsnt 2019.10.08,17:31]} Nb4 {[%ccsnt 2019.10.08,19:09]} 
55.Bb5 {[%ccsnt 2019.10.09,06:02]} Rc2 {[%ccsnt 2019.10.09,09:16]} 
56.Re1 {[%ccsnt 2019.10.09,11:56]} Nc6 {[%ccsnt 2019.10.09,13:36]} 
57.Re6 {[%ccsnt 2019.10.10,08:54]} Rc5 {[%ccsnt 2019.10.10,14:31]} 
58.Bf1 {[%ccsnt 2019.10.11,05:57]} a4 {[%ccsnt 2019.10.11,15:05]} 
59.Rxg6 {[%ccsnt 2019.10.11,16:38]} Rc1 {[%ccsnt 2019.10.11,19:57]} 
60.Kf2 {[%ccsnt 2019.10.12,09:14]} Kc8 {[%ccsnt 2019.10.12,17:20]} 
61.Bh3+ {[%ccsnt 2019.10.14,09:38]} Kb7 {[%ccsnt 2019.10.14,09:50]} 
62.Rg7+ {[%ccsnt 2019.10.14,10:52]} Kb6 {[%ccsnt 2019.10.14,11:00]} 
63.Rg8 {[%ccsnt 2019.10.14,11:31]} a3 {[%ccsnt 2019.10.14,12:50]} 
64.Ra8 {[%ccsnt 2019.10.14,13:06]} Rc2+ {[%ccsnt 2019.10.14,14:25]} 
65.Ke3 {[%ccsnt 2019.10.14,15:41]} Rc3+ {[%ccsnt 2019.10.14,18:25]} 
66.Ke4 {[%ccsnt 2019.10.15,05:12]} Rxg3 {[%ccsnt 2019.10.15,09:44]} 
67.Be6 {[%ccsnt 2019.10.15,15:29]} Rg6 {[%ccsnt 2019.10.16,06:04]} 
68.Kf5 {[%ccsnt 2019.10.16,09:00]} Rg3 {[%ccsnt 2019.10.16,15:35]} 
69.d7 {[%ccsnt 2019.10.16,16:19]} Rd3 {[%ccsnt 2019.10.17,05:01]} 
70.Ke4 {[%ccsnt 2019.10.17,06:51]} Nb4 {[%ccsnt 2019.10.17,09:30]} 
71.Rb8+ {[%ccsnt 2019.10.17,09:54]} Kc5 {[%ccsnt 2019.10.17,19:37]} 
72.Rxb4 {[%ccsnt 2019.10.17,19:52]} Rxd7 {[%ccsnt 2019.10.18,05:58]} 
73.Bxd7 {[%ccsnt 2019.10.18,06:50]} Kxb4 {[%ccsnt 2019.10.18,09:40]} 
1-0`

// FewPiecesGame2 is something to be fixed
// https://www.iccf.com/game?id=1049523
// [Event "FIN/WS/47 (FIN)"]
// [Site "ICCF"]
// [Date "2018.10.01"]
// [Round "-"]
// [White "Hoogkamer, Michael"]
// [Black "Mäkelä, Ari"]
// [Result "0-1"]
// [WhiteElo "1975"]
const FewPiecesGame2 = `
1.d4 d5 2.c4 c6 3.cxd5 cxd5 4.Bf4 Nc6 5.e3 Qb6 6.Nc3 Nf6 7.Bd3 a6 8.Nge2 e6 9.O-O Be7 10.Na4 Qd8 11.h3 Nb4 12.Bb1 Nc6 13.Bd3 Nb4 14.Bb1 Nc6 15.a3 Na5 16.Qc2 Nc4 17.b3 Nd6 18.Rc1 Bd7 19.Nc5 Bc6 20.Qb2 O-O 21.Nd3 Qb6 22.a4 a5 23.g4 Rfc8 24.f3 Be8 25.Nc5 e5 26.dxe5 Nc4 27.Qc3 Nxe5 28.Bxe5 Bxc5 29.Bd4 Bxd4 30.Qxd4 Qxb3 31.Rxc8 Rxc8 32.g5 Nd7 33.Ba2 Qc2 34.Nf4 Qc7 35.Kg2 Kf8 36.Bxd5 Nb6 37.Be4 Rd8 38.Rc1 Qxc1 39.Qxd8 Nc4 40.Nd5 Qb2+ 41.Kg3 Qe5+ 42.Kg2 Qd6 43.Qxd6+ Nxd6 44.Bc2 b5 45.Nb6 bxa4 46.Nxa4 Nc4 47.Kf2 Bd7 48.h4 Bxa4 49.Bxa4 Nb6 50.Bb3 a4 51.Ba2 Nc8 52.Ke1 Nd6 53.Bb1 a3 54.Kd2 Nf5 55.h5 h6 56.f4 Ng3 57.Kc3 Nxh5 58.Bf5 Ng3 59.gxh6 gxh6 60.Bh3 Ne4+ 61.Kb3 h5 62.Kxa3 Nf2 63.Bf1 h4 64.Kb3 h3 65.Bxh3 Nxh3 66.Kc2 Ke7 0-1
`

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
