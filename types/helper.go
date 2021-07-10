package roar

import "roar/util"

func CheckRlePairsBmps(bmp Bitmaps) int {
	_c := 0
	var i, j int
	for _, v := range bmp.Values {
		if v == 0 {
			continue
		}
		i = 0
	nextValue:
		for i < util.BmpRange {
			for ; i < util.BmpRange && (v&1<<i == 0); i++ {
			}
			if i == util.BmpRange {
				break nextValue
			}
			j = i
			for ; j < util.BmpRange && (v&1<<j == 1); j++ {
			}
			_c += 1
			i = j
		}
	}
	return _c
}

func CheckRLePairsSarr(ar Sarr) int {
	_c := 0
	for i := 1; i <= len(ar.Arr)-1; i++ {
		if ar.Arr[i-1] != ar.Arr[i]-1 {
			_c += 1
		}
	}
	_c += 1
	return _c
}

func BmpsVsRles(bmp Bitmaps) string {
	if util.BmpsLen > CheckRlePairsBmps(bmp) {
		return "rles"
	}
	return "bmps"
}

func BmpsVsSarr(bmp Bitmaps) string {
	if (bmp.NumElem()) < uint16(util.BmpsLen) {
		return "sarr"
	}
	return "bmps"
}

func RlesVsBmps(rle Rles) string {
	if len(rle.RlePairs) > util.BmpsLen { //len(rle.RlePairs)*2 * 16 > util.BmpsLen* 32
		return "bmps"
	}
	return "rles"
}

func RlesVsSarr(rle Rles) string {
	if int(rle.NumElem()) < len(rle.RlePairs)*2 { //int(rle.NumElem()) * 16 < len(rle.RlePairs) * 16 * 2
		return "sarr"
	}
	return "rles"
}

func SarrVsBmps(sarr Sarr) string {
	if len(sarr.Arr) < util.BmpsLen*2 { //len(sarr.Arr)*16 < util.BmpsLen*32
		return "sarr"
	}
	return "bmps"
}

func SarrVsRles(sarr Sarr) string {
	if len(sarr.Arr) > CheckRLePairsSarr(sarr)*2 {
		return "rles"
	}
	return "sarr"
}
