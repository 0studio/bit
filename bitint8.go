package bit

type BitInt8 uint8

// pos 以0为基
func (bitInt *BitInt8) SetFlag(pos uint8) {
	if pos > 8 {
		return
	}
	*bitInt = *bitInt | (1 << pos)
}
func (bitInt BitInt8) GetFlag(pos uint8) uint8 {
	if pos > 8 {
		return 0
	}
	return uint8((bitInt >> pos) & 1)
}
func (bitInt BitInt8) IsPosTrue(pos uint8) bool {
	if pos > 8 {
		return false
	}
	return ((bitInt >> pos) & 1) == 1
}
func (bitInt *BitInt8) UnSetFlag(pos uint8) {
	if pos > 8 {
		return
	}
	*bitInt = *bitInt & (^(1 << pos))
}
func (bitInt BitInt8) IsAllZero() bool {
	return bitInt == 0
}
func (bitInt BitInt8) GetTrueLen(mostPos ...uint8) (cnt uint8) { //
	// see TestBitIntGetTrueLen
	// 返回共有多少位是1,可选参数mostPos ,指定不必检测所有的32位pos ,只需要检测指定个数 的位置即可
	// 如mostPos 为10 ， 则只检测0~9位的数字
	var total uint8 = 8
	var pos uint8 = 0
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
