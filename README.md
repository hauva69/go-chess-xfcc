# go-chess-xfcc

Golang support for [XML web services for correspondence chess](https://www.bennedik.de/xfcc/).

## Configuration

Create a configuration file to you home directory:

```bash
mkdir -p "${HOME}/.config/go-chess-xfcc"
```

and a file named `go-chess-xfcc.xml` with contents:

```xml
<configuration>
	<user>REPLACE_WITH_YOUR_USERNAME</user>
	<password>REPLACE_WITH_YOUR_PASSWORD</password>
</configuration>
```

## ICCF

[International Correspondence Chess Federation](https://www.iccf.com/) supports XFCC. For practical examples see

- [https://www.iccf.com/xfccbasic.asmx](https://www.iccf.com/xfccbasic.asmx)

## Syzygy tables

- [Syzygy endgame tablebases](https://syzygy-tables.info/)

### Lichess API

- [Lichess API Introduction](https://lichess.org/api#section/Introduction)
  - [Tablebase server for lichess.org](https://github.com/niklasf/lila-tablebase)
  - [7-piece Syzygy tablebases are complete](https://lichess.org/blog/W3WeMyQAACQAdfAL/7-piece-syzygy-tablebases-are-complete)
