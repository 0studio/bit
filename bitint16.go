package bit

type BitInt16 uint16

func (bitInt BitInt16) Len() int {
	return 16
}

// pos 以0为基
func (bitInt *BitInt16) SetFlag(pos uint16) {
	if pos > 16 {
		return
	}
	*bitInt = *bitInt | (1 << pos)
}
func (bitInt BitInt16) GetFlag(pos uint16) uint16 {
	if pos > 16 {
		return 0
	}
	return uint16((bitInt >> pos) & 1)
}
func (bitInt BitInt16) IsPosTrue(pos uint16) bool {
	if pos > 16 {
		return false
	}
	return ((bitInt >> pos) & 1) == 1
}
func (bitInt *BitInt16) UnSetFlag(pos uint16) {
	if pos > 16 {
		return
	}
	*bitInt = *bitInt & (^(1 << pos))
}
func (bitInt BitInt16) IsAllZero() bool {
	return bitInt == 0
}
func (bitInt BitInt16) GetTrueLen(mostPos ...uint16) (cnt uint16) { //
	// see TestBitIntGetTrueLen
	// 返回共有多少位是1,可选参数mostPos ,指定不必检测所有的32位pos ,只需要检测指定个数 的位置即可
	// 如mostPos 为10 ， 则只检测0~9位的数字
	var total uint16 = 16
	var pos uint16 = 0
	if len(mostPos) > 0 && mostPos[0] < total {
		total = mostPos[0]
	}
	for pos = 0; pos < total; pos++ {
		if bitInt.IsPosTrue(pos) {
			cnt++
		}
	}
	return
}
