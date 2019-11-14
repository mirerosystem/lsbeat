// Created by cgo -godefs - DO NOT EDIT
// cgo -godefs defs_freebsd.go

package ipv4

const (
	sysIP_OPTIONS     = 0x1
	sysIP_HDRINCL     = 0x2
	sysIP_TOS         = 0x3
	sysIP_TTL         = 0x4
	sysIP_RECVOPTS    = 0x5
	sysIP_RECVRETOPTS = 0x6
	sysIP_RECVDSTADDR = 0x7
	sysIP_SENDSRCADDR = 0x7
	sysIP_RETOPTS     = 0x8
	sysIP_RECVIF      = 0x14
	sysIP_ONESBCAST   = 0x17
	sysIP_BINDANY     = 0x18
	sysIP_RECVTTL     = 0x41
	sysIP_MINTTL      = 0x42
	sysIP_DONTFRAG    = 0x43
	sysIP_RECVTOS     = 0x44

	sysIP_MULTICAST_IF           = 0x9
	sysIP_MULTICAST_TTL          = 0xa
	sysIP_MULTICAST_LOOP         = 0xb
	sysIP_ADD_MEMBERSHIP         = 0xc
	sysIP_DROP_MEMBERSHIP        = 0xd
	sysIP_MULTICAST_VIF          = 0xe
	sysIP_ADD_SOURCE_MEMBERSHIP  = 0x46
	sysIP_DROP_SOURCE_MEMBERSHIP = 0x47
	sysIP_BLOCK_SOURCE           = 0x48
	sysIP_UNBLOCK_SOURCE         = 0x49
	sysMCAST_JOIN_GROUP          = 0x50
	sysMCAST_LEAVE_GROUP         = 0x51
	sysMCAST_JOIN_SOURCE_GROUP   = 0x52
	sysMCAST_LEAVE_SOURCE_GROUP  = 0x53
	sysMCAST_BLOCK_SOURCE        = 0x54
	sysMCAST_UNBLOCK_SOURCE      = 0x55

	sizeofSockaddrStorage = 0x80
	sizeofSockaddrInet    = 0x10

	sizeofIPMreq         = 0x8
	sizeofIPMreqn        = 0xc
	sizeofIPMreqSource   = 0xc
	sizeofGroupReq       = 0x88
	sizeofGroupSourceReq = 0x108
)

type sockaddrStorage struct {
	Len         uint8
	Family      uint8
	X__ss_pad1  [6]int8
	X__ss_align int64
	X__ss_pad2  [112]int8
}

type sockaddrInet struct {
	Len    uint8
	Family uint8
	Port   uint16
	Addr   [4]byte /* in_addr */
	Zero   [8]int8
}

type ipMreq struct {
	Multiaddr [4]byte /* in_addr */
	Interface [4]byte /* in_addr */
}

type ipMreqn struct {
	Multiaddr [4]byte /* in_addr */
	Address   [4]byte /* in_addr */
	Ifindex   int32
}

type ipMreqSource struct {
	Multiaddr  [4]byte /* in_addr */
	Sourceaddr [4]byte /* in_addr */
	Interface  [4]byte /* in_addr */
}

type groupReq struct {
	Interface uint32
	Pad_cgo_0 [4]byte
	Group     sockaddrStorage
}

type groupSourceReq struct {
	Interface uint32
	Pad_cgo_0 [4]byte
	Group     sockaddrStorage
	Source    sockaddrStorage
}
