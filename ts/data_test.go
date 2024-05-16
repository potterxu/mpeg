package ts_test

import (
	"encoding/base64"
	"io"
	"strings"
)

const (
	nullPacketPID uint32 = 0x1fff
	nullPacket           = `
Rx//EP////////////////////////////////////////////////////////////////////////
//////////////////////////////////////////////////////////////////////////////
//////////////////////////////////////////////////////////////////////////////
////////////////8=
`

	dataPacketPID uint32 = 0x21
	dataPacket           = `
RwAhHbpBrFk0EitqYe9/6MqqfLuk5EU1IyVVIirpPhNJ4tZ4f3dPGu7p9dMqmCESJEGXqehE8V5F6o
oc96aqnt7qp7s54l3NdPZybreCc2qAAAABG1KLZpJLB0n71IVnDQYYadazHYM5BEugQ3OhKQaAk11M
PftXzvVe6hnhlLnuLk3T4C42RoJDfQ1sBKCkssDBMTyoAZgB98AP1kwEn/kebAM76qG1ks0BoCYslh
kx/0AnAMWIWJQFSxM=
`

	shortTsStream = `
RwAhHbpBrFk0EitqYe9/6MqqfLuk5EU1IyVVIirpPhNJ4tZ4f3dPGu7p9dMqmCESJEGXqehE8V5F6o
oc96aqnt7qp7s54l3NdPZybreCc2qAAAABG1KLZpJLB0n71IVnDQYYadazHYM5BEugQ3OhKQaAk11M
PftXzvVe6hnhlLnuLk3T4C42RoJDfQ1sBKCkssDBMTyoAZgB98AP1kwEn/kebAM76qG1ks0BoCYslh
kx/0AnAMWIWJQFSxNHADEW6ZoOyz5dOtCdXXTSgtpos7RldzE9UTUqiH4B3j6g04Op0FqEvSk+TZs1
dxYZnqMoVPwNqhU2N089CpwoH69ddYN8q66F1qnlzycrwacHU6C9CI7lnQAAAAEHMqoqcddUS8uTZ9
OmBwIm7mp0uFY13dd4AHlU6MNpqoJ2oRCaFT2AU4eTnkYsaAkw3Ekd1K9JzodROjPbE+hNMsTRSYWX
+52HbaSUgebBqDgLq+MiGGFsWxRQ8b8KgEcAIR7KHQAlAdAJklE3JDAIWy6JawcnskQUA2YmFF8nge
iEAxJmJgDsYvDx0GBixhTkaJpacQiYPA5tWpoygCgBlwC0mE1sHb+K3GIv9Z4m62olMGhEkh8HX+Rf
6pJaX7U2Oz5YxUJsk1j6GyximTdgoUyhkAvh9Er+B4+4Td/x60Cz0drsCgWPUllwQGuB6gwmIHh1oA
oGOSBo5JHol0CAY8XIN8tseQIy/xLJAHPt5JIYDV2AwQiEgb7oRwBBGyuuVQn6Y2YTs1YEJplyutvS
j89idGErND6OkQPBFQQCqCQWRA0EwAAAAQqCuE/NDrpdVTOWVpluyOW0ZSZSRwnZpz3qFPSOpZqqtF
dIVMdMdIUJXouSg9BQNQKg0g4NQa2QYMAAAAELgrvSql10mXTZbSdbKUpWlbCVk2jAOTLIJtJEBVV1
cJ1JekirSZUjYamPTKug8JoNCDabH0qAAAABDGK4TdImaTCY1LronaNLUtk2mUkyTJlHACEfmfHOgP
USv//fknz0S//G20Wimw2fDB8/eY2RqaBQ8lIf9MOm4ySiiYwKEQ9PPjDbklOUWwwP9E2gQXrzx3y9
xOi+tBF9m0NlDyaGfsB4UdQHiK//UF3Q8DVDWwEDaRVbDZ1n3YrNlvePms6NoakOtn3TRzbokSNIp6
kSKkXT0IkSJF03cmI91hhU+dNXNkIkSIq6ehEiKirp/RUiKp4uirp/qkSKhyZ/oqRIqf+f0SKqn9Ei
RIrQCA==
`

	adapationFieldPacket = `
RwAhNmcA//////////////////////////////////////////////////////////////////////
//////////////////////////////////////////////////////////////////TcHR1Jo5rLVT
2d3T66TlzUt9VUnJmpPqqqT4s0fP6pMnRS5rL6qqo68/KhJR++kZ+gt/CU0MqqSew09vQ1LDlScGSg
z8qwqC7u62hjn3AAA=
`

	adapationFieldDetailedPacket = `
R0kQNgpiCAIGKMgo7BCzAAAB4AAAhNANM0dhQs0TR2EKjYA0cQAAAAEJEAAAAAFnZAAorYQ1iH
QgULBBB1CGsQ6EChYIIOoQ1iHQgULBBB1CGsRzYYgqhDWI5sMQVQhrEc2GIKoQ/6aaTVoxUmyG
LJLIpDHLEEERJRwoxxYWHR4JBQsF0G0If9evk/Ef1+b4rwhw1gigoLKUB4BE/eA0IEBASAAAAw
AIAAADAZSgAAAAAWjvPLAAAAEGBgE=
`

	pcrDataPacket = `
R0kbOwcQKOmq1gA9AAABwAFagJAII0dN0UOAAoHDQ8zLArP2tjs5bcrafVPjQEVutJEgVavMy0WxGX
d4d5h11dGHWJAqJiD//YTAZWZmVVVLM4gAAACCyNbLzGHWmV1GYGYHmn456paZpZZH5opI3pKZZ6mL
RvBE0vMvXweD1p7dFO1Q8NKI3sOxCzIwtrA3atXsbK++EeW1g+QmCDO7nEu63qFMzopMj00tiV7Z0c
OwoZnmWIn02zRtM64=
`
)

func createReader(data string) func() io.Reader {
	return func() io.Reader { return base64.NewDecoder(base64.StdEncoding, strings.NewReader(data)) }
}

var (
	nullPacketReader              = createReader(nullPacket)
	dataPacketReader              = createReader(dataPacket)
	fivePacketReader              = createReader(shortTsStream)
	adaptationFieldReader         = createReader(adapationFieldPacket)
	detailedAdaptationFieldReader = createReader(adapationFieldDetailedPacket)
	pcrDataPacketReader           = createReader(pcrDataPacket)
)
