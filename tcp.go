package tcp

import ()

type Header struct {
	SourcePort      uint16
	DestinationPort uint16
	Sequence        uint32
	Acknowledgement uint32
	DataOffset      uint8
	//Unused Bytes Here..
	WindowSize    uint16
	Checksum      uint16
	UrgentPointer uint16
	Options       []byte
}

func ParseHeader(m []byte) (*Header, error) {
	h := &Header{}
	h.SourcePort = uint16(m[0])<<8 | uint16(m[1])
	h.DestinationPort = uint16(m[2])<<8 | uint16(m[3])
	h.Sequence = uint32(m[4])<<32 | uint32(m[5])<<16 | uint32(m[6])<<8 | uint32(m[7])
	h.Acknowledgement = uint32(m[8])<<32 | uint32(m[9])<<16 | uint32(m[10])<<8 | uint32(m[11])
	h.DataOffset = uint8(m[12] >> 4)
	h.WindowSize = uint16(m[14])<<8 | uint16(m[15])
	h.Checksum = uint16(m[16])<<8 | uint16(m[17])
	h.UrgentPointer = uint16(m[18])<<8 | uint16(m[19])

	return h, nil
}
