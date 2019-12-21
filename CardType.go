package sss1

import (
	"fmt"
	"sort"
)

//CardsCanOut 是否可以出牌
func (d *SSSInfo) CardsCanOut(cards []int, isSpecialCards bool) (bool, error) {
	if len(cards) != 13 {
		return false, fmt.Errorf("不是13张牌无法出牌: Cards => %0x", cards)
	}

	if isSpecialCards {
		types, _, _ := d.GetSpecialCardsTypes(cards)
		return types > 0, fmt.Errorf("不是特殊牌无法出牌: Cards => %0x", cards)
	}

	var cards1 = cards[0:3]
	var cards2 = cards[3:8]
	var cards3 = cards[8:13]

	cardsType1, count1 := d.GetCardsType(cards1)
	cardsType2, count2 := d.GetCardsType(cards2)
	cardsType3, count3 := d.GetCardsType(cards3)

	//中道大于尾道
	if cardsType2 > cardsType3 || (cardsType2 == cardsType3 && count2 > count3) {
		return false, fmt.Errorf("中道大于尾道: ZhongCards => %0x, WeiCards => %0x", cards2, cards3)
	}
	//首道大于中道
	if cardsType1 > cardsType2 || (cardsType1 == cardsType2 && count1 > count2) {
		return false, fmt.Errorf("首道大于中道: ShouCards => %0x, ZhongCards => %0x", cards1, cards2)
	}

	return true, nil
}

//GetCardsType 获取普通牌类型
func (d *SSSInfo) GetCardsType(cards []int) (types, s int) {
	//王数量
	kingNum := 0
	for i := 0; i < len(cards); i++ {
		if cards[i] == 0x41 || cards[i] == 0x42 {
			kingNum++
		}
	}

	NC := make([]int, len(cards))
	copy(NC, cards)
	sort.Sort(CardsSort(NC))
	//计算同花
	tonghua := d.isTongHua2(NC, kingNum)

	//点数
	DelHuaSe(NC)
	sort.Ints(NC)

	if len(NC) == 5 {
		//五同
		if cardsType, count := d.isWuTong(NC, kingNum); cardsType {
			return WU_TONG, count
		}

		//同花顺
		if cardsType, count := d.isShunZi(NC, kingNum); cardsType && tonghua {
			return TONG_HUA_SHUN, count
		}

		//铁支
		if cardsType, count := d.isTieZhi(NC, kingNum); cardsType {
			return TIE_ZHI, count
		}

		//葫芦
		if cardsType, count := d.isHuLu(NC, kingNum); cardsType {
			return HU_LU, count
		}

		//同花
		if cardsType, count := d.isTongHua1(NC, kingNum); cardsType && tonghua {
			return TONG_HUA, count
		}

		//顺子
		if cardsType, count := d.isShunZi(NC, kingNum); cardsType {
			return SHUN_ZI, count
		}
	}
	if len(NC) == 3 { //只在首道实现
		if kingNum == 3 {
			return SAN_WANG, NC[0]
		}
		if kingNum == 2 {
			return SHUANG_WANG, NC[0]
		}
	}
	//三条
	if cardsType, count := d.isSanTiao1(NC, kingNum); cardsType {
		return SAN_TIAO, count
	}

	//两对
	if cardsType, count := d.isLiangDui(NC, kingNum); cardsType {
		return LIANG_DUI, count
	}

	//对子
	if cardsType, count := d.isDuiZi(NC, kingNum); cardsType {
		return DUI_ZI, count
	}

	//乌龙
	var count = 0
	c := getNumCards(NC, kingNum)
	getCount(&count, 5, c, false)
	return 0, count
}

//GetSpecialCardsTypes 获取特殊牌类型
func (d *SSSInfo) GetSpecialCardsTypes(cards []int) (types, card, score int) {
	tempCards := make([]int, len(cards))
	copy(tempCards, cards)
	//获取鬼的数量
	gui := d.GetGuiNum(tempCards)
	//获取同花
	tongHua := isTongHua(tempCards, gui)
	//删除花色
	DelHuaSe(tempCards)
	//排序
	sort.Ints(tempCards)
	tcards := make([]int, len(cards))
	copy(tcards, cards)
	sort.Ints(tcards)
	if card_, ok := d.anaysisToBaTong(tempCards, gui); ok {
		//八同
		types = BA_TONG
		card = card_
	} else if d.analysisToYiTiaoLong(tempCards, gui) && tongHua {
		//清一条龙
		types = QING_LONG
	} else if card_, ok := d.anaysisToQiTong(tempCards, gui); ok {
		//七同
		types = QI_TONG
		card = card_
	} else if d.analysisToYiTiaoLong(tempCards, gui) {
		//一条龙
		types = YI_TIAO_LONG
	} else if card_, ok := d.anaysisToLiuTong(tempCards, gui); ok {
		//六同
		types = LIU_TONG
		card = card_
	} else if card_, ok := d.analysisToSiSanTiao(tempCards, gui); ok {
		//四三条
		types = SI_SAN_TIAO
		card = card_
	} else if d.analysisToSanTongHua(tcards, gui) {
		//三同花
		types = SAN_TONG_HUA
	} else if d.analysisToLiuDuiBan(tempCards, gui) {
		//六对半
		types = LIU_DUI_BAN
	} else if d.analysisToSanShunZi1(tempCards, gui) {
		//三顺子
		types = SAN_SHUN_ZI
	}
	score = SpecialCardsScoreMap[types]
	return
}

//GetSpecialCardsTypes 获取特殊牌类型
func (d *SSSInfo) GetSpecialCardsTypes1(cards []int) (types, card, score int) {
	tempCards := make([]int, len(cards))
	copy(tempCards, cards)
	//获取鬼的数量
	gui := d.GetGuiNum(tempCards)
	//获取同花
	tongHua := isTongHua(tempCards, gui)
	//删除花色
	DelHuaSe(tempCards)
	//排序
	sort.Ints(tempCards)
	tcards := make([]int, len(cards))
	copy(tcards, cards)
	sort.Ints(tcards)
	if card_, ok := d.anaysisToBaTong(tempCards, gui); ok {
		//八同
		types = BA_TONG
		card = card_
	} else if d.analysisToYiTiaoLong(tempCards, gui) && tongHua {
		//清一条龙
		types = QING_LONG
	} else if card_, ok := d.anaysisToQiTong(tempCards, gui); ok {
		//七同
		types = QI_TONG
		card = card_
	} else if d.analysisToYiTiaoLong(tempCards, gui) {
		//一条龙
		types = YI_TIAO_LONG
	} else if card_, ok := d.anaysisToLiuTong(tempCards, gui); ok {
		//六同
		types = LIU_TONG
		card = card_
	}
	score = SpecialCardsScoreMap[types]
	return
}

//ComparisonCards 比牌组
func (d *SSSInfo) ComparisonCards(zhuaCards []int, isZhuangSpecial bool, xianCards []int, isXianSpecial bool) (ret EndScore) {
	//如果两家特殊牌
	if isZhuangSpecial && isXianSpecial {
		zhuangTypes, zhuanCard, _ := d.GetSpecialCardsTypes(zhuaCards)
		xianTypes, xianCard, _ := d.GetSpecialCardsTypes(xianCards)
		if SpecialCardWeightMap[zhuangTypes] > SpecialCardWeightMap[xianTypes] { //庄家赢
			ret.TeShu = SpecialCardsScoreMap[zhuangTypes]
			return
		} else if SpecialCardWeightMap[zhuangTypes] == SpecialCardWeightMap[xianTypes] { //和牌
			if zhuangTypes == LIU_TONG || zhuangTypes == QI_TONG || zhuangTypes == BA_TONG || zhuangTypes == YI_TIAO_LONG {
				if zhuanCard > xianCard { //庄家赢
					ret.TeShu = SpecialCardsScoreMap[zhuangTypes]
					return
				}
				if zhuanCard < xianCard { //闲家赢
					ret.TeShu = -SpecialCardsScoreMap[xianTypes]
					return
				}
				//和牌
				return
			}
			return
		} else { //闲家赢
			ret.TeShu = -SpecialCardsScoreMap[xianTypes]
			return
		}
	}
	//有一家有特殊牌
	if isZhuangSpecial || isXianSpecial {
		if isZhuangSpecial { //庄家特殊牌
			zhuangTypes, _, _ := d.GetSpecialCardsTypes(zhuaCards)
			ret.TeShu = SpecialCardsScoreMap[zhuangTypes]
			return
		}
		//闲家特殊牌
		xianTypes, _, _ := d.GetSpecialCardsTypes(xianCards)
		ret.TeShu = -SpecialCardsScoreMap[xianTypes]
		return
	}
	//普通牌比大小
	ret.Tou = d.ComparisonCard(zhuaCards[0:3], xianCards[0:3], 0)
	ret.Zhong = d.ComparisonCard(zhuaCards[3:8], xianCards[3:8], 1)
	ret.Wei = d.ComparisonCard(zhuaCards[8:13], xianCards[8:13], 2)
	return
}

//ComparisonCard 比牌
func (d *SSSInfo) ComparisonCard(zhuangCards []int, xianCards []int, count int) (retScore int) {
	zhuanType, zhuanCount := d.GetCardsType(zhuangCards)
	xianType, xianCount := d.GetCardsType(xianCards)
	if zhuanType > xianType { //庄家赢
		return CardsTypeScoreMap[zhuanType][count]
	} else if zhuanType == xianType { //相等
		if zhuanCount > xianCount { //庄家赢
			return CardsTypeScoreMap[zhuanType][count]
		} else if zhuanCount == xianCount { //和牌
			return 0
		}
	}
	//闲家赢
	return -CardsTypeScoreMap[xianType][count]
}
