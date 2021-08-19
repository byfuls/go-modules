package bitFlags

type UByte uint8
type UShort uint16

const (
	/*_______________LOBYTE_________________*/
	Eat    UShort = 1 << iota /* 1 */
	Sleep                     /* 2 */
	Sing                      /* 4 */
	Play                      /* 8 */
	Call                      /* 16 */
	Sms                       /* 32 */
	Shower                    /* 64 */
	_                         /* 128 */
	/*_______________HIBYTE_________________*/
	_      /* 256 */
	Work   /* 512*/
	Health /* 1024 */
)

type BitStatus struct {
	val UShort
	hi  UByte
	lo  UByte
}

func (s *BitStatus) Set(flag UShort) {
	s.val = s.val | flag
	s.hi = UByte(s.val>>8) & 0xff
	s.lo = UByte(s.val) & 0xff
}

func (s *BitStatus) Clear(flag UShort) {
	s.val = s.val &^ flag
	s.hi = UByte(s.val>>8) & 0xff
	s.lo = UByte(s.val) & 0xff
}

func (s *BitStatus) Toggle(flag UShort) {
	s.val = s.val ^ flag
	s.hi = UByte(s.val>>8) & 0xff
	s.lo = UByte(s.val) & 0xff
}

func (s *BitStatus) Has(flag UShort) bool {
	return s.val&flag != 0
}

func (s *BitStatus) SetVal(val UShort) {
	s.val = val
	s.hi = UByte(s.val>>8) & 0xff
	s.lo = UByte(s.val) & 0xff
}

func (s *BitStatus) GetVal() UShort {
	return s.val
}

func (s *BitStatus) GetValHigh() UByte {
	return s.hi
}

func (s *BitStatus) GetValLow() UByte {
	return s.lo
}

func (s *BitStatus) AllClear() {
	s.val = 0
	s.hi = 0
	s.lo = 0
}
