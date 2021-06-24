package strconv

import "strconv"

func S2I32(str string) int32 {
	if len(str) <= 0{
		return 0
	}
	res, _ := strconv.ParseInt(str,10,32)
	return int32(res)
}

func S2I64(str string) int64 {
	if len(str) <= 0{
		return 0
	}
	res, _ := strconv.ParseInt(str,10,64)
	return res
}

func S2I(str string) int {
	if len(str) <= 0{
		return 0
	}
	res, _ := strconv.Atoi(str)
	return res
}
