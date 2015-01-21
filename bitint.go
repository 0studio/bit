package bit

type BitInt uint64

// pos 以0为基
func (bitInt *BitInt) SetFlag(pos uint64) {
	if pos > 63 {
		return
	}
	*bitInt = *bitInt | (1 << pos)
}
func (bitInt *BitInt) GetFlag(pos uint64) uint64 {
	if pos > 63 {
		return 0
	}
	return uint64((*bitInt >> pos) & 1)
}
func (bitInt *BitInt) IsPosTrue(pos uint64) bool {
	if pos > 63 {
		return false
	}
	return ((*bitInt >> pos) & 1) == 1
}
func (bitInt *BitInt) UnSetFlag(pos uint64) {
	if pos > 63 {
		return
	}
	*bitInt = *bitInt & (^(1 << pos))
}
func (bitInt *BitInt) IsAllZero() bool {
	return *bitInt == 0
}
func (bitInt *BitInt) GetTrueLen(mostPos ...uint64) (cnt uint64) { //
	// see TestBitIntGetTrueLen
	// 返回共有多少位是1,可选参数mostPos ,指定不必检测所有的63位pos ,只需要检测指定个数 的位置即可
	// 如mostPos 为10 ， 则只检测0~9位的数字
	var total uint64 = 64
	var pos uint64 = 0
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
