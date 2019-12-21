package sss1

import (
	"shisanshui/utils"
	"sort"
)

//牌
type CardsType struct {
	HuaSe      int //花色
	CardsValue int //牌值
	Count      int //真实位置
}

type CardsNum struct {
	Num  int
	Card int
}

//五同
func (d *SSSInfo) isWuTong(NC []int, kingNum int) (bool, int) {
	if kingNum >= 4 {
		return true, NC[0]
	}
	cards := getNumCards(NC, kingNum)
	if len(cards) == 1 {
		return true, cards[0].Card
	}
	return false, 0
}

//铁支
func (d *SSSInfo) isTieZhi(NC []int, kingNum int) (ok bool, ret int) {
	card, num := GetLenMaxCard(NC, kingNum)
	//fmt.Printf("nc => %0x", NC)
	if num+kingNum >= 4 {
		for i := 0; i < len(NC)-kingNum; i++ {
			if NC[i] != card {
				ret += NC[i]
				break
			}
		}
		return true, ret + (card * 0x10)
	}
	return false, 0
}

//葫芦
func (d *SSSInfo) isHuLu(NC []int, kingNum int) (ok bool, count int) {
	cards := getNumCards(NC, kingNum)
	if len(cards) == 2 {
		if kingNum > 0 {
			cards[0].Num++
		}
		if cards[0].Num < cards[1].Num {
			cards[0], cards[1] = cards[1], cards[0]
		}
		getCount(&count, 2, cards, true)
		return true, count
	}
	return false, 0
}

/*
	同花比较特殊 3条 2对 对子 单张
*/
//同花 (小威重写)
func (d *SSSInfo) isTongHua1(NC []int, kingNum int) (ok bool, count int) {
	cardsNum := getNumCards(NC, kingNum)
	if len(cardsNum) == 5 { //五张散牌
		getCount(&count, 5, cardsNum, false)
		return true, count
	}
	//三条(因为两对和对子计算不同所有点数单独计算重写)
	if ok, _ := d.isSanTiao1(NC, kingNum); ok {
		//取出三张
		for i := range cardsNum {
			if cardsNum[i].Num+kingNum == 3 {
				count += cardsNum[i].Card * 0x10000000
				cardsNum = append(cardsNum[:i], cardsNum[i+1:]...)
				getCount(&count, 5, cardsNum, false)
				return true, count
			}

		}
	}
	//两对
	if ok, _ := d.isLiangDui(NC, kingNum); ok {
		//取出三张
		for k := 0; k < 2; k++ {
			for i := range cardsNum {
				if cardsNum[i].Num+kingNum == 2 {
					if k == 0 {
						count += cardsNum[i].Card * 0x1000000
					} else {
						count += cardsNum[i].Card * 0x100000
					}
					cardsNum = append(cardsNum[:i], cardsNum[i+1:]...)
					break
				}
			}
		}
		getCount(&count, 5, cardsNum, false)
		return true, count
	}
	//对子
	if ok, _ := d.isDuiZi(NC, kingNum); ok {
		//取出对子
		for i := range cardsNum {
			if cardsNum[i].Num+kingNum == 2 {
				count += cardsNum[i].Card * 0x100000
				cardsNum = append(cardsNum[:i], cardsNum[i+1:]...)
				getCount(&count, 5, cardsNum, false)
				return true, count
			}
		}
	}
	return false, count
}

//同花(主要用于同花计算)
func (d *SSSInfo) isTongHua2(NC []int, kingNum int) bool {
	hua := map[int]int{}
	for i := 0; i < len(NC)-kingNum; i++ {
		hua[NC[i]/16]++
	}
	if len(hua) == 1 {
		return true
	}
	return false
}

//顺子
func (d *SSSInfo) isShunZi(NC []int, kingNum int) (bool, int) {
	if len(NC) == kingNum {
		return true, 0xe
	}
	//不能有相等的牌
	for i := 1; i < len(NC)-kingNum; i++ {
		if NC[i-1] == NC[i] {
			return false, 0
		}
	}

	//特殊顺子A23或A2345
	if d.isDiShun(NC, kingNum) {
		return true, 0xe0000 - 1
	}
	tempGui := kingNum
	//最大牌与最小牌大小不能超过牌张数-1
	for i := 1; i < len(NC)-kingNum; i++ {
		if NC[i] == NC[i-1] {
			return false, 0
		}
		if NC[i]-NC[i-1] != 1 { //不是相邻的两个数
			tempGui -= NC[i] - NC[i-1] - 1
		}
	}
	if tempGui < 0 {
		return false, 0
	}
	maxCard := NC[len(NC)-1-kingNum]
	if maxCard+tempGui > 0xE {
		maxCard = 0xE
	} else {
		maxCard += tempGui
	}
	return true, maxCard * 0x10000
}

//地顺
func (d *SSSInfo) isDiShun(cards []int, kingNum int) bool {
	//地顺特别处理
	diShunKey := 0
	for i := 0; i < len(cards)-kingNum; i++ {
		num := 0
		if i == 0 {
			num = 0
		} else {
			num = 1
		}
		for i1 := 0; i1 < i; i1++ {
			num = num * 0x10
		}
		if num == 0 {
			diShunKey += cards[len(cards)-kingNum-i-1]
		} else {
			diShunKey += cards[len(cards)-kingNum-i-1] * num
		}
		if len(cards) == 3 {
			if _, ok := DI_SHUN_CARDS_3_MAP[kingNum][diShunKey]; ok {
				return true
			}
		} else {
			if _, ok := DI_SHUN_CARDS_MAP[kingNum][diShunKey]; ok {
				return true
			}
		}

	}
	return false
}

//三条(小威改)
func (d *SSSInfo) isSanTiao1(cards []int, kingNum int) (ret bool, count int) {
	cardsNums := getNumCards(cards, kingNum)
	for i := 0; i < len(cardsNums); i++ {
		if cardsNums[i].Num+kingNum == 3 {
			ret = true
			cardsNums[i].Num -= 3 - kingNum
			count = 0x10000 * cardsNums[i].Card
		}
	}
	getCount(&count, 4, cardsNums, true)
	return ret, count
}

//两对
func (d *SSSInfo) isLiangDui(NC []int, kingNum int) (ok bool, ret int) {
	if len(NC) != 5 {
		return false, 0
	}
	cards := getNumCards(NC, kingNum)
	if len(cards) != 3 {
		return false, 0
	} else {
		if kingNum > 0 {
			if cards[0].Num >= 2 {

			} else {
				cards[1].Num++
			}
		}
		for i := 1; i < len(cards); i++ {
			if cards[i].Num > cards[i-1].Num {
				cards[i], cards[i-1] = cards[i-1], cards[i]
			}
		}
		getCount(&ret, 5, cards, true)
		return true, ret
	}
}

//对子
func (d *SSSInfo) isDuiZi(NC []int, kingNum int) (ok bool, ret int) {
	cards := getNumCards(NC, kingNum)
	for i := range cards {
		if cards[i].Num+kingNum >= 2 {
			ret = cards[i].Card * 0x10000
			cards = append(cards[:i], cards[i+1:]...)
			getCount(&ret, 4, cards, true)
			return true, ret
		}
	}
	return false, 0
}

//取出所有相同大小的牌(去重)
func (d *SSSInfo) getOutAllSameSizeCards(cards []int, num int) (a []int, b []int) {
	for i := len(cards) - 1; i > num-2; {
		if i < 0 {
			break
		}

		if b != nil && cards[i] == b[len(b)-1] {
			continue
		}

		if cards[i] == cards[i-num+1] {
			b = append(b, cards[i:i+num]...)
			cards = append(cards[:i], cards[i+num:]...)

			i -= num
		} else {
			i--
		}
	}

	return cards, b
}

//////////////////////////////////////////////////////////////////////////

//analysisToQingLong 一条龙分析
func (this *SSSInfo) analysisToYiTiaoLong(cards []int, gui int) bool {
	cardsMap := make(map[int]int)
	for i := 0; i < len(cards)-gui; i++ {
		if cardsMap[cards[i]] != 0 {
			return false
		}
		cardsMap[cards[i]]++
	}
	return true
}

//analysisToLiuDuiBan 六对半
func (this *SSSInfo) analysisToLiuDuiBan(cards []int, gui int) (ok bool) {
	cardsMap := make(map[int]int)
	for i := 0; i < len(cards)-gui; i++ {
		cardsMap[cards[i]]++
	}
	dan := 0
	if len(cardsMap) > 7 {
		return false
	}
	//计算对子和单张数量
	for _, v := range cardsMap {
		if v%2 != 0 {
			dan++
		}
	}
	if dan-1 <= gui {
		return true
	}
	return false
}

//analysisToSanTongHua 三同花
func (d *SSSInfo) analysisToSanTongHua(cards []int, gui int) (isok bool) {
	cardsMap := make(map[int]int)
	for i := 0; i < len(cards)-gui; i++ {
		cardsMap[cards[i]/16]++
	}
	if len(cardsMap) == 4 {
		return false
	}
	if len(cardsMap) == 1 {
		return true
	}
	//填充
	for k, _ := range cardsMap {
		if cardsMap[k] == 3 || cardsMap[k] == 5 || cardsMap[k] == 8 || cardsMap[k] == 10 {
			continue
		}
		if cardsMap[k] < 3 && cardsMap[k] > 0 {
			gui += cardsMap[k] - 3
			cardsMap[k] = 3
		}
		if cardsMap[k] < 5 && cardsMap[k] > 3 {
			gui += cardsMap[k] - 5
			cardsMap[k] = 5
		}
		if cardsMap[k] < 8 && cardsMap[k] > 5 {
			gui += cardsMap[k] - 8
			cardsMap[k] = 8
		}
		if cardsMap[k] < 10 && cardsMap[k] > 8 {
			gui += cardsMap[k] - 10
			cardsMap[k] = 10
		}
	}
	if gui < 0 {
		return false
	}
	for _, v := range cardsMap {
		if len(cardsMap) == 3 {
			if v != 3 && v != 5 {
				return false
			}
		} else if len(cardsMap) == 2 {
			if v != 3 && v != 8 && v != 10 && v != 5 {
				return false
			}
		}
	}
	return true
}

//analysisToSanShunZi 三顺子
func (d *SSSInfo) analysisToSanShunZi(cards []int, gui int) (isok bool) {
	var shunZiAll [][]int
	//分析所有顺子
	for i := 0; i < len(cards); i++ {
		for i1 := i + 1; i1 < len(cards); i1++ {
			for i2 := i1 + 1; i2 < len(cards); i2++ {
				for i3 := i2 + 1; i3 < len(cards); i3++ {
					for i4 := i3 + 1; i4 < len(cards); i4++ {
						c := []int{cards[i], cards[i1], cards[i2], cards[i3], cards[i4]}
						DelHuaSe(c)
						sort.Ints(c)
						if ok, _ := d.isShunZi(c, d.GetGuiNum(c)); ok {
							shunZiAll = append(shunZiAll, c)
						}
					}
				}
			}
		}
	}
	//fmt.Println(len(shunZiAll))
	//fmt.Println(shunZiAll)
	//计算三顺子
	for i := 0; i < len(shunZiAll); i++ {
		for k := i + 1; k < len(shunZiAll); k++ {
			tempCards := make([]int, len(cards))
			copy(tempCards, cards)
			tempCards = utils.DelArrayMembers(tempCards, shunZiAll[i])
			tempCards = utils.DelArrayMembers(tempCards, shunZiAll[k])
			if len(tempCards) == 3 {
				sort.Ints(tempCards)
				if ok, _ := d.isShunZi(tempCards, d.GetGuiNum(tempCards)); ok {
					return true
				}
			}
		}
	}
	return
}

//analysisToSanShunZi 三顺子
func (d *SSSInfo) analysisToSanShunZi1(cards []int, gui int) (isok bool) {
	tempCards := make([]int, len(cards))
	copy(tempCards, cards)
	DelHuaSe(tempCards)
	sort.Ints(tempCards)
	var shunZiAll [][]int
	c := make([]int, 5) //复用性能提升巨大
	//分析所有顺子
	for i := 0; i < len(tempCards); i++ {
		for i1 := i + 1; i1 < len(tempCards); i1++ {
			for i2 := i1 + 1; i2 < len(tempCards); i2++ {
				for i3 := i2 + 1; i3 < len(tempCards); i3++ {
					for i4 := i3 + 1; i4 < len(tempCards); i4++ {
						c[0], c[1], c[2], c[3], c[4] = tempCards[i], tempCards[i1], tempCards[i2], tempCards[i3], tempCards[i4]
						if ok, _ := d.isShunZi(c, d.GetGuiNum(c)); ok {
							c_ := make([]int, len(c))
							copy(c_, c)
							shunZiAll = append(shunZiAll, c_)
						}
					}
				}
			}
		}
	}
	//fmt.Println(len(shunZiAll))
	//fmt.Println(shunZiAll)
	//tempCards_ := make([]int, len(cards))
	//计算三顺子
	for i := 0; i < len(shunZiAll); i++ {
		for k := i + 1; k < len(shunZiAll); k++ {
			if d.GetGuiNum(shunZiAll[i])+d.GetGuiNum(shunZiAll[k]) > d.maxGuiNum {
				continue
			}
			tempCards_ := make([]int, len(tempCards))
			copy(tempCards_, tempCards)
			tempCards_ = utils.DelArrayMembers(tempCards_, shunZiAll[i])
			tempCards_ = utils.DelArrayMembers(tempCards_, shunZiAll[k])
			if len(tempCards_) == 3 {
				sort.Ints(tempCards_)
				//fmt.Printf("%x2 \n ", tempCards_)
				if ok, _ := d.isShunZi(tempCards_, d.GetGuiNum(tempCards_)); ok {
					return true
				}
			}
		}
	}
	return
}

//analysisToSiSanTiao 四张三条
func (d *SSSInfo) analysisToSiSanTiao(cards []int, gui int) (card int, isok bool) {
	cardMap := make(map[int]int)
	for i := 0; i < len(cards)-gui; i++ {
		cardMap[cards[i]]++
	}
	if len(cardMap) > 5 || len(cardMap) < 3 { //这样一定不成功
		return 0, false
	}
	ok := false
	//取出单牌
	for i := range cardMap {
		if cardMap[i] == 1 {
			//删除一张的牌
			delete(cardMap, i)
			ok = true
			break
		}
	}
	//没有王就失败
	if ok == false && gui == 0 {
		return card, false
	}
	tempGui := gui
	s := 0
	for i := range cardMap {
		if cardMap[i] >= 3 {
			//s += cardMap[i] / 3 //防止出现两个完全相同的三个
			for cardMap[i] > 3 {
				s++
				cardMap[i] -= 3
			}
		}
		tempGui -= 3 - cardMap[i]
		if tempGui < 0 {
			return 0, false
		}
		s++
	}
	return 0, s == 4
}

//anaysisToLiuTong 六同(6张相同的牌算上王)
func (d *SSSInfo) anaysisToLiuTong(cards []int, gui int) (card int, isok bool) {
	card, lens := GetLenMaxCard(cards, gui)
	return card, lens+gui >= 6
}

//anaysisToLiuTong 七同(7张相同的牌算上王)
func (d *SSSInfo) anaysisToQiTong(cards []int, gui int) (card int, isok bool) {
	card, lens := GetLenMaxCard(cards, gui)
	return card, lens+gui >= 7
}

//anaysisToLiuTong 八同(8张相同的牌算上王)
func (d *SSSInfo) anaysisToBaTong(cards []int, gui int) (card int, isok bool) {
	card, lens := GetLenMaxCard(cards, gui)
	return card, lens+gui >= 8
}

//是否是同花
func isTongHua(cards []int, gui int) bool {
	hua := -1
	for i := 0; i < len(cards)-gui; i++ {
		if hua == -1 {
			hua = cards[i] / 16
		} else if hua != cards[i]/16 {
			return false
		}
	}
	return true
}

//GetGuiNum 获取癞子数量
func (d *SSSInfo) GetGuiNum(cards []int) (ret int) {
	for i := 0; i < len(cards); i++ {
		if cards[i] == 0x41 || cards[i] == 0x42 {
			ret++
		}
	}
	return
}

//Sort 排序
func (d *SSSInfo) Sort(cards []int) {
	sort.Sort(CardsSort(cards))
}

//getMaxCards 获取一个最大牌
func (d *SSSInfo) getMaxCard(cards []int) (ret []int) {
	var max []int
	for i := 0; i < len(cards); i++ {
		for i1 := i + 1; i1 < len(cards); i1++ {
			for i2 := i1 + 1; i2 < len(cards); i2++ {
				for i3 := i2 + 1; i3 < len(cards); i3++ {
					for i4 := i3 + 1; i4 < len(cards); i4++ {
						if len(max) == 0 {
							max = []int{cards[i], cards[i1], cards[i2], cards[i3], cards[i4]}
						} else {
							if d.ComparisonCard(max, []int{cards[i], cards[i1], cards[i2], cards[i3], cards[i4]}, 1) < 0 {
								max = []int{cards[i], cards[i1], cards[i2], cards[i3], cards[i4]}
							}
						}
					}
				}
			}
		}
	}
	return max
}

//getMaxCards 获取最大的牌组
func (d *SSSInfo) GetMaxCards(cards []int) (ret []int) {
	tempCards := make([]int, len(cards))
	copy(tempCards, cards)
	//得到尾道牌
	cardsWei := d.getMaxCard(cards)
	//删除尾道牌
	tempCards = utils.DelArrayMembers(tempCards, cardsWei)
	//得到中道的牌
	cardsZhong := d.getMaxCard(tempCards)
	///删除中道得到尾道的牌
	cardsShou := utils.DelArrayMembers(tempCards, cardsZhong)
	ret = append(ret, cardsShou...)
	ret = append(ret, cardsZhong...)
	ret = append(ret, cardsWei...)
	return ret
}

//GetLenMaxCard 获取一个张长度最大的牌
func GetLenMaxCard(cards []int, guiNum int) (card int, num int) {
	cardsMap := make(map[int]int)
	for i := 0; i < len(cards)-guiNum; i++ {
		cardsMap[cards[i]]++
	}
	max := -1
	for k, v := range cardsMap {
		if max == -1 {
			max = k
		} else {
			if cardsMap[max] < v {
				max = k
			} else if cardsMap[max] == v {
				if max < k {
					max = k
				}
			}
		}
	}
	return max, cardsMap[max]
}

//取出指定长度的牌指定牌
func GetDesignationNumCards(cards []int, num int) (ret []int) {
	cardsMap := make(map[int]int)
	for _, v := range cards {
		cardsMap[v]++
	}
	for k, v := range cardsMap {
		if v == num {
			ret = append(ret, k)
		}
	}
	//排序
	sort.Sort(sort.IntSlice(ret))
	return
}

//去一张最大的牌
func getMaxCard(cards CardsSort) int {
	return 0
}

//获取牌的数量
func getNumCards(cards []int, gui int) (ret []CardsNum) {
	cardsMap := make(map[int]int)
	for i := 0; i < len(cards)-gui; i++ {
		cardsMap[cards[i]]++
	}
	for k, v := range cardsMap {
		ret = append(ret, CardsNum{
			Num:  v,
			Card: k,
		})
	}
	//排序
	for i := 0; i < len(ret); i++ {
		for k := i + 1; k < len(ret); k++ {
			if ret[i].Card < ret[k].Card {
				ret[i], ret[k] = ret[k], ret[i]
			}
		}
	}
	return ret
}

//获取点数
func getCount(count *int, s int, cardsNum []CardsNum, dan bool) {
	for i := 0; i < len(cardsNum); i++ {
		if cardsNum[i].Num > 0 {
			for j := 0; j < cardsNum[i].Num; j++ {
				c := 0x10
				for k := 1; k < s-1; k++ {
					c *= 0x10
				}
				s--
				if s == 0 { //最后一张牌特殊处理
					c = 0
					*count += cardsNum[i].Card
				} else {
					*count += cardsNum[i].Card * c
				}
				if dan {
					break
				}
			}

		}
	}
}
