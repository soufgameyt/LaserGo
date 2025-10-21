package datastream

type ByteStream struct {
	buffer    []byte
	offset    int
	bitOffset int
}

func NewByteStream() *ByteStream {
	return &ByteStream{
		buffer:    make([]byte, 0),
		offset:    0,
		bitOffset: 0,
	}
}

func (b *ByteStream) WriteInt(value int32) {
	b.ensureCapacity(4)
	b.buffer = append(b.buffer, byte(value>>24), byte(value>>16), byte(value>>8), byte(value))
	b.offset += 4
}

func (b *ByteStream) ReadInt() int32 {
	val := int32(b.buffer[b.offset])<<24 |
		int32(b.buffer[b.offset+1])<<16 |
		int32(b.buffer[b.offset+2])<<8 |
		int32(b.buffer[b.offset+3])
	b.offset += 4
	return val
}

func (b *ByteStream) WriteShort(value int16) {
	b.ensureCapacity(2)
	b.buffer = append(b.buffer, byte(value>>8), byte(value))
	b.offset += 2
}

func (b *ByteStream) ReadShort() int16 {
	val := int16(b.buffer[b.offset])<<8 | int16(b.buffer[b.offset+1])
	b.offset += 2
	return val
}

func (b *ByteStream) WriteByte(value byte) {
	b.ensureCapacity(1)
	b.buffer = append(b.buffer, value)
	b.offset++
}

func (b *ByteStream) ReadByte() byte {
	val := b.buffer[b.offset]
	b.offset++
	return val
}

func (b *ByteStream) WriteBytes(buf []byte) {
	if buf == nil {
		b.WriteInt(-1)
		return
	}
	b.WriteInt(int32(len(buf)))
	b.buffer = append(b.buffer, buf...)
	b.offset += len(buf)
}

func (b *ByteStream) ReadBytes() []byte {
	length := b.ReadInt()
	if length <= 0 {
		return nil
	}
	val := b.buffer[b.offset : b.offset+int(length)]
	b.offset += int(length)
	return val
}

func (b *ByteStream) WriteString(str string) {
	if str == "" {
		b.WriteInt(-1)
		return
	}
	data := []byte(str)
	b.WriteInt(int32(len(data)))
	b.buffer = append(b.buffer, data...)
	b.offset += len(data)
}

func (b *ByteStream) ReadString() string {
	length := b.ReadInt()
	if length <= 0 {
		return ""
	}
	val := string(b.buffer[b.offset : b.offset+int(length)])
	b.offset += int(length)
	return val
}

func (b *ByteStream) WriteBoolean(value bool) {
	if b.bitOffset == 0 {
		b.ensureCapacity(1)
		b.buffer = append(b.buffer, 0)
		b.offset++
	}
	if value {
		b.buffer[b.offset-1] |= 1 << b.bitOffset
	}
	b.bitOffset = (b.bitOffset + 1) & 7
}

func (b *ByteStream) ReadBoolean() bool {
	return b.ReadVInt() >= 1
}

func (b *ByteStream) WriteVInt(value int32) {
	bitOffset := 0
	temp := byte((value >> 25) & 0x40)
	flipped := value ^ (value >> 31)
	temp |= byte(value & 0x3F)
	value >>= 6
	flipped >>= 6

	if flipped == 0 {
		b.WriteByte(temp)
		return
	}

	b.WriteByte(temp | 0x80)

	for flipped != 0 {
		r := byte(0)
		if flipped != 0 {
			r = 0x80
		}
		b.WriteByte(byte(value&0x7F) | r)
		value >>= 7
		flipped >>= 7
		bitOffset++
	}
}

func (b *ByteStream) ReadVInt() int32 {
	result := int32(0)
	shift := uint(0)
	var a1, a2, s int32
	for {
		byteVal := int32(b.buffer[b.offset])
		b.offset++
		if shift == 0 {
			a1 = (byteVal & 0x40) >> 6
			a2 = (byteVal & 0x80) >> 7
			s = (byteVal << 1) & ^int32(0x181)
			byteVal = s | (a2 << 7) | a1
		}
		result |= (byteVal & 0x7F) << shift
		shift += 7
		if (byteVal & 0x80) == 0 {
			break
		}
	}
	return (result >> 1) ^ -(result & 1)
}

func (b *ByteStream) WriteLogicLong(high, low int32) {
	b.WriteVInt(high)
	b.WriteVInt(low)
}

func (b *ByteStream) ReadLogicLong() (high, low int32) {
	high = b.ReadVInt()
	low = b.ReadVInt()
	return
}

func (b *ByteStream) WriteLong(high, low int32) {
	b.WriteInt(high)
	b.WriteInt(low)
}

func (b *ByteStream) ReadLong() (high, low int32) {
	high = b.ReadInt()
	low = b.ReadInt()
	return
}

func (b *ByteStream) Reset() {
	b.buffer = make([]byte, 0)
	b.offset = 0
	b.bitOffset = 0
}

func (b *ByteStream) GetLength() int {
	return len(b.buffer)
}

func (b *ByteStream) GetBuffer() []byte {
	return b.buffer
}

func (b *ByteStream) ReplaceBuffer(buf []byte) {
	b.buffer = buf
	b.offset = 0
	b.bitOffset = 0
}

func (b *ByteStream) ensureCapacity(capacity int) {
	if len(b.buffer) < b.offset+capacity {
		tmp := make([]byte, capacity)
		b.buffer = append(b.buffer, tmp...)
	}
}

func (b *ByteStream) Skip(len int) {
	b.bitOffset += len
}
