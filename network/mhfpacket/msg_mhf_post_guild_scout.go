package mhfpacket

import (
	"github.com/Andoryuuta/Erupe/network"
	"github.com/Andoryuuta/byteframe"
)

// MsgMhfPostGuildScout represents the MSG_MHF_POST_GUILD_SCOUT
type MsgMhfPostGuildScout struct {
	AckHandle uint32
	CharID    uint32
}

// Opcode returns the ID associated with this packet type.
func (m *MsgMhfPostGuildScout) Opcode() network.PacketID {
	return network.MSG_MHF_POST_GUILD_SCOUT
}

// Parse parses the packet from binary
func (m *MsgMhfPostGuildScout) Parse(bf *byteframe.ByteFrame) error {
	m.AckHandle = bf.ReadUint32()
	m.CharID = bf.ReadUint32()

	return nil
}

// Build builds a binary packet from the current data.
func (m *MsgMhfPostGuildScout) Build(bf *byteframe.ByteFrame) error {
	panic("Not implemented")
}