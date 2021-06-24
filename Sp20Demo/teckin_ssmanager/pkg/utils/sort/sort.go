package sort

import (
	"errors"
	"sort"
)

//字符串列表正序排列
func StringListPositiveOrder(lstr []string) ([]string, error) {
	if len(lstr) <=0 {
		return []string{}, errors.New("传入字符串列表不能为空")
	}
	sort.Strings(lstr)
	return lstr, nil
}

//字符串列表倒序排列
func StringListReverseOrder(lstr []string) ([]string, error) {
	lstr2, err := StringListPositiveOrder(lstr)
	if err != nil{
		return []string{}, err
	}
	sort.Sort(sort.Reverse(sort.StringSlice(lstr2)))
	return lstr2, nil
}
