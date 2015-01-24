package bit

type BitInt32 uint32

// pos 以0为基
func (bitInt *BitInt32) SetFlag(pos uint32) {
	if pos > 32 {
		return
	}
	*bitInt = *bitInt | (1 << pos)
}
func (bitInt *BitInt32) GetFlag(pos uint32) uint32 {
	if pos > 32 {
		return 0
	}
	return uint32((*bitInt >> pos) & 1)
}
func (bitInt *BitInt32) IsPosTrue(pos uint32) bool {
	if pos > 32 {
		return false
	}
	return ((*bitInt >> pos) & 1) == 1
}
func (bitInt *BitInt32) UnSetFlag(pos uint32) {
	if pos > 32 {
		return
	}
	*bitInt = *bitInt & (^(1 << pos))
}
func (bitInt *BitInt32) IsAllZero() bool {
	return *bitInt == 0
}
func (bitInt *BitInt32) GetTrueLen(mostPos ...uint32) (cnt uint32) { //
	// see TestBitIntGetTrueLen
	// 返回共有多少位是1,可选参数mostPos ,指定不必检测所有的32位pos ,只需要检测指定个数 的位置即可
	// 如mostPos 为10 ， 则只检测0~9位的数字
	var total uint32 = 32
	var pos uint32 = 0
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
