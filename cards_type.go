package sss1

///////////排序结构体/////////////////////////

type CardsSort []int

func (a CardsSort) Len() int { // 重写 Len() 方法
	return len(a)
}
func (a CardsSort) Swap(i, j int) { // 重写 Swap() 方法
	a[i], a[j] = a[j], a[i]
}
func (a CardsSort) Less(i, j int) bool { // 重写 Less() 方法， 从大到小排序
	var i1, j1 int
	if a[i] == 0x41 || a[i] == 0x42 {
		i1 = a[i]
	} else {
		i1 = a[i] % 16
	}

	if a[j] == 0x41 || a[j] == 0x42 {
		j1 = a[j]
	} else {
		j1 = a[j] % 16
	}
	if i1 == j1 {
		return a[i] < a[j]
	}
	return i1 < j1
}

///////////排序结构体/////////////////////////

//EndScore 结束积分
type EndScore struct {
	Tou   int
	Zhong int
	Wei   int
	TeShu int
}

type cardType struct {
	types int //类型
	count int //点数
}

type cardTypes struct {
	tou   cardType
	zhong cardType
	wei   cardType
}
