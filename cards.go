package sss1

import (
	"math"
	"shisanshui/utils/wrand.v1"
	"sync"
)

type SSSInfo struct {
	cards     []int //牌堆牌
	laizi     int   //癞子
	maxGuiNum int
}

//随机种子
//var rands *rand.Rand
var locker sync.Mutex

func NewSSS(CardsColors int, JiaDaXiaoWang int) *SSSInfo {
	sss := new(SSSInfo)
	cards := make([]int, len(CARDS))
	copy(cards, CARDS)
	sss.cards = cards[:13*CardsColors]
	switch JiaDaXiaoWang {
	case 1:
		sss.maxGuiNum = 0
	case 2:
		sss.cards = append(sss.cards, 0x41, 0x42)
		sss.maxGuiNum = 2
	case 3:
		sss.cards = append(sss.cards, 0x41, 0x42, 0x41, 0x42)
		sss.maxGuiNum = 4
	}
	return sss
}

//洗牌
func (d *SSSInfo) Shuffle(types int) {
	var tcatds []int
	switch types {
	case LIU_TONG, QI_TONG, BA_TONG: //六同洗牌
		num := 0
		if len(d.cards)/13+len(d.cards)%13 >= 6 && LIU_TONG == types { //六同条件
			num = 6
		} else if len(d.cards)/13+len(d.cards)%13 >= 7 && QI_TONG == types {
			num = 7
		} else if len(d.cards)/13+len(d.cards)%13 >= 8 && BA_TONG == types {
			num = 8
		} else {
			d.Shuffle1()
			return
		}
		card := wrand.GetInt(13) + 2
		tcatds = d.GetAssignCard(card)
		tcatds = append(tcatds, d.GetAssignCard(0x41)...)
		tcatds = append(tcatds, d.GetAssignCard(0x42)...)
		wrand.OutOfOrder(tcatds)
		d.Shuffle1()
		d.cards = append(tcatds[:num], d.cards...)
		d.cards = append(d.cards, tcatds[num:]...)
		tcatds = d.GetCards(13)
		if types1, _, _ := d.GetSpecialCardsTypes(tcatds); types1 != types {
			d.cards = append(d.cards, tcatds...)
			d.Shuffle(types)
			return
		}
		d.Shuffle1()
		d.cards = append(tcatds, d.cards...)
	default:
		d.Shuffle1()
		for i := 0; i < len(d.cards)/13; i++ {
			types, _, _ := d.GetSpecialCardsTypes1(d.cards[i*13 : i*13+13])
			if types >= LIU_TONG {
				if wrand.GetInt(100) < 50 {
					d.Shuffle1()
					break
				}
			}
		}
	}

}
func (d *SSSInfo) Shuffle1() {
	for i := 0; i < len(d.cards); i++ {
		j := wrand.GetInt(len(d.cards) - 1)
		d.cards[i%len(d.cards)], d.cards[j] = d.cards[j], d.cards[i%len(d.cards)]
	}
}

//获取指定牌
func (d *SSSInfo) GetAssignCard(card int) (ret []int) {
	for i := range d.cards {
		if d.cards[i] != 0x41 && d.cards[i] != 0x42 {
			if d.cards[i]%16 == card%16 && card < 0x41 {
				ret = append(ret, d.cards[i])
			}
		} else {
			if d.cards[i] == card {
				ret = append(ret, d.cards[i])
			}
		}
	}
	d.cards = DelCards(d.cards, card)
	return ret
}

//DelCards 删除牌组
func DelCards(s []int, d int) (ret []int) {
	ret = s
	for i := 0; i < len(s); i++ {
		if s[i] == 0x41 || s[i] == 0x42 {
			if s[i] == d {
				ret = append(s[:i], s[i+1:]...)
				return DelCards(ret, d)
			}
		} else {
			if s[i]%16 == d%16 && d < 0x41 {
				ret = append(s[:i], s[i+1:]...)
				return DelCards(ret, d)
			}
		}
	}
	return ret
}

//GetCards 获取牌(获取完后最好检测一次数组长度)
func (this *SSSInfo) GetCards(num int) []int {
	tempCards := make([]int, num)
	copy(tempCards, this.cards[:num])
	this.cards = append(this.cards[num:])
	return tempCards
}

//GetAllCards 获取全部牌组
func (this *SSSInfo) GetAllCards() []int {
	return this.cards
}

//DelDianShu 去除牌的点数
func DelDianShu(cards []int) {
	for i := 0; i < len(cards); i++ {
		if cards[i] < 0x41 {
			cards[i] = int(math.Ceil(float64(cards[i]) / 16))
		}
	}
}

//DelHuaSe 去除牌的花色
func DelHuaSe(cards []int) {
	for i := 0; i < len(cards); i++ {
		if cards[i] < 0x41 {
			cards[i] = cards[i] % 16
		}
	}
}
