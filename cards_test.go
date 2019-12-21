package sss1

import (
	"fmt"
	"server/utils"
	"sort"
	"sync"
	"testing"
	"time"
)

func TestDelHuaSe(t *testing.T) {
	sss := NewSSS(4, 2)
	cards := []int{0x2c, 0x24, 0x14, 0x37, 0x7, 0x36, 0x15, 0x2, 0x2d, 0xc, 0xb, 0x3a, 0x39}
	sort.Sort(CardsSort(cards))
	cards1 := []int{14, 38, 52, 62, 54, 45, 12, 37, 22, 40, 3, 40, 55}
	sort.Sort(CardsSort(cards1))
	fmt.Printf("%0x \n", cards1)
	fmt.Printf("%0x \n", cards)

	//sort.Sort(CardsSort(cards))
	fmt.Println(sss.CardsCanOut(cards, true))
	//sss.Sort(cards)
	//DelHuaSe(cards)
	//fmt.Println(sss.GetSpecialCardsTypes(cards))
	//sort.Ints(cards)
	//cards = share.DelArrayMembers(cards, []int{4, 5, 6, 7, 8})
	//cards = share.DelArrayMembers(cards, []int{9, 10, 11, 12, 13})
	fmt.Printf("%X\n", cards)
}

func TestSSSInfo_GetCardsType2(t *testing.T) {
	//for i := 0; i < 10000; i++ {
	sss := NewSSS(4, 2)
	cards := []int{0x3C, 0x26, 0x06, 0x3B, 0x28, 0x08, 0x27, 0x05, 0x42, 0x2E, 0x1E, 0x0E, 0x33}
	if ok, err := sss.CardsCanOut(cards, false); !ok {
		t.Error(err)
	}
	fmt.Printf("%0x \n", cards)
	types, count := sss.GetCardsType(cards[0:3])
	fmt.Printf("type => %v, count =>  %0x \n", types, count)
	types, count = sss.GetCardsType(cards[3:8])
	fmt.Printf("type => %v, count =>  %0x \n", types, count)
	types, count = sss.GetCardsType(cards[8:13])
	fmt.Printf("type => %v, count =>  %0x \n", types, count)
	//}
}

func TestDelDianShu(t *testing.T) {
	sss := NewSSS(4, 2)
	cards := []int{0xb, 0x1b, 0x2b, 0x2a, 0x9, 0x19, 0x8, 0x17, 0x36, 0x5, 0x4, 0x23, 0x32}
	sss.Sort(cards)
	DelHuaSe(cards)
	cards = utils.DelArrayMembers(cards, []int{5, 6, 7, 8, 9})
	fmt.Printf("%X\n", cards)
	cards = utils.DelArrayMembers(cards, []int{0x7, 0x8, 0x9, 0xA, 0xB})
	fmt.Printf("%X\n", cards)
}

func TestSSSInfo_GetCardsType(t *testing.T) {
	sss := NewSSS(4, 2)
	types, conut := sss.GetCardsType([]int{0x41, 0x42, 3})
	fmt.Println(types)
	fmt.Printf("%X \n", conut)
}

func TestTest(t *testing.T) {
	sss := NewSSS(4, 2)
	cards := []int{0x2, 0x13, 0x24}
	DelHuaSe(cards)
	sort.Ints(cards)
	fmt.Println(sss.isShunZi(cards, 0))
}

func TestGetLenMaxCard(t *testing.T) {
	fmt.Println(GetLenMaxCard([]int{7, 0xa}, 0))
}

func TestSSSInfo_ComparisonCard(t *testing.T) {
	sss := NewSSS(4, 2)
	fmt.Println(sss.ComparisonCard([]int{61, 10, 25}, []int{13, 42, 24}, 0))
	fmt.Println(3 % 4)
}

func TestTestGetMaxcards(t *testing.T) {
	sss := NewSSS(6, 2)
	c := sss.GetMaxCards([]int{30, 60, 35, 61, 13, 11, 22, 4, 46, 39, 37, 37, 34})
	fmt.Printf("%0x\n", c)
	if ok, _ := sss.CardsCanOut(c, false); !ok {
		fmt.Println(c)
		panic(c)
	}
}

func TestGetMaxcard(t *testing.T) {
	n := 18
	wait := sync.WaitGroup{}
	wait.Add(n)
	for s := 0; s < n; s++ {
		go func() {
			for i := 0; i < 1e6; i++ {
				sss := NewSSS(6, 2)
				sss.Shuffle(0)
				for i := 0; i < 6; i++ {
					c := sss.GetMaxCards(sss.GetCards(13))
					fmt.Println(c)
					if ok, _ := sss.CardsCanOut(c, false); !ok {
						fmt.Println(c)
						panic(c)
					}
				}
			}
			wait.Done()
		}()
	}
	wait.Wait()
}

func BenchmarkGetMaxcard(b *testing.B) {
	sss := NewSSS(4, 2)
	sss.Shuffle(0)
	cards := sss.GetCards(13)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		sss.getMaxCard(cards)
	}

}

func TestNewSSS(t *testing.T) {

}

func TestSSSInfo_GetAllCards(t *testing.T) {
	a := []int{7, 8, 43, 46, 42, 46, 10, 50, 45, 61, 62, 61, 28}
	fmt.Printf("%0x", a)
}

func TestSSSInfo_Bipai(t *testing.T) {
	sss := NewSSS(4, 2)
	fmt.Println(sss.ComparisonCard([]int{0x9, 0xa, 0xb, 0xc, 0xd}, []int{0x42, 0x2, 0x3, 0x4, 0x5}, 2))
}

func TestSSSInfo_GetCards(t *testing.T) {
	sss := NewSSS(4, 2)
	_, conut := sss.isSanTiao1([]int{0x4, 0x4, 0x4, 1, 2}, 0)
	fmt.Printf("%0x \n", conut)
}
func TestSSSInfo_Shuffle(t *testing.T) {
	sss := NewSSS(4, 2)
	fmt.Println(sss.GetSpecialCardsTypes([]int{
		0x12, 0x23, 0x3e,
		0x6, 0x7, 0x8, 0x9, 0xa,
		0x7, 0x8, 0x9, 0xa, 0xb,
	}))
}

func TestIsTeshuPai(t *testing.T) {
	n := 12
	count := int(1e6)
	wait := sync.WaitGroup{}
	for k := 0; k < n; k++ {
		wait.Add(1)
		go func() {
			for i := 0; i < count/n; i++ {
				sss := NewSSS(6, 2)
				sss.Shuffle(0)
				for i := 0; i < 4; i++ {
					sss.GetSpecialCardsTypes(sss.GetCards(13))
				}
			}
			wait.Done()

		}()
	}
	wait.Wait()
	t.Log("运行完成")
}

func TestAnaysisToSiSanTiao(t *testing.T) {
	sss := NewSSS(4, 2)
	//2 4 5 14 65
	//5 6 7 9 65
	fmt.Println(sss.analysisToSanShunZi1([]int{2, 3, 4, 4, 5, 5, 6, 6, 8, 9, 10, 14, 65}, 0))
	//fmt.Println(sss.analysisToSanTongHua(, 1))
}

func TestHulu(t *testing.T) {
	sss := NewSSS(4, 2)
	isok, count := sss.isHuLu([]int{3, 3, 3, 2, 2}, 0)
	fmt.Printf("isok => %v, Count = %0x", isok, count)
}

func TestSss(t *testing.T) {
	num := 0
	ts1 := make(map[int]int)
	for i := 0; i < 100; i++ {
		func() {
			ts := make(map[int]int)
			s := 24
			chanNum := make(chan int, s)
			for k := 0; k < s; k++ {
				go func() {
					for {
						<-chanNum
						sss := NewSSS(6, 3)
						sss.Shuffle(0)
						r := utils.Rands.GetInt(6)
						for j := 0; j < 6; j++ {
							if j == r {
								cards := sss.GetCards(13)
								types, _, _ := sss.GetSpecialCardsTypes(cards)
								if types >= LIU_TONG {
									ts[types]++
								}
							}
						}
					}
				}()
			}
			for i := 0; i < 4200+s; i++ {
				select {
				case chanNum <- i:
				case <-time.After(time.Second * 5):
				}
			}
			fmt.Println(ts)
			//if ts[LIU_TONG] > 30 {
			ts1[ts[LIU_TONG]]++
			//num++
			//}
		}()
	}
	fmt.Println(num)
	fmt.Println(ts1)
}

func BenchmarkDelDianShu(b *testing.B) {
	sss := NewSSS(6, 3)
	sss.Shuffle(0)
	cards := sss.GetCards(13)
	for i := 0; i < b.N; i++ {
		sss.GetSpecialCardsTypes(cards)
	}
}

func TestSSSInfo_Shuffle2(t *testing.T) {
	s := 0
	typeMap := make(map[int]int)
	mchan := make(chan int, 24)
	for i := 0; i < 1e8; i++ {
		mchan <- 1
		go func() {
			sss := NewSSS(6, 3)
			sss.Shuffle(-1)
			for i1 := 0; i1 < 6; i1++ {
				types, _, _ := sss.GetSpecialCardsTypes(sss.GetCards(13))
				s++
				if types >= LIU_TONG {
					typeMap[types]++
				}
			}
			<-mchan
		}()
	}
	fmt.Println(s)
	fmt.Println(typeMap)
}

func BenchmarkSSSInfo_GetSpecialCardsTypes(b *testing.B) {
	sss := NewSSS(6, 3)
	sss.Shuffle(-1)
	//cards := []int{8, 51, 18, 60, 38, 24, 52, 58, 27, 13, 22, 3, 40}
	//515582 195265 271109 162963
	cards := sss.GetCards(13)
	//fmt.Println(cards)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		sss.GetSpecialCardsTypes(cards)
	}
}
