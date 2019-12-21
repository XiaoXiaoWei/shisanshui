package sss1

const (
	ZHUANG_JIA = 1
	XIAN_JIA   = 2
	HE_PAI     = 3
)

//牌
var CARDS = []int{
	0x32, 0x33, 0x34, 0x35, 0x36, 0x37, 0x38, 0x39, 0x3A, 0x3B, 0x3C, 0x3D, 0x3E,
	0x22, 0x23, 0x24, 0x25, 0x26, 0x27, 0x28, 0x29, 0x2A, 0x2B, 0x2C, 0x2D, 0x2E,
	0x12, 0x13, 0x14, 0x15, 0x16, 0x17, 0x18, 0x19, 0x1A, 0x1B, 0x1C, 0x1D, 0x1E,
	0x2, 0x3, 0x4, 0x5, 0x6, 0x7, 0x8, 0x9, 0x0A, 0x0B, 0x0C, 0x0D, 0x0E,

	0x32, 0x33, 0x34, 0x35, 0x36, 0x37, 0x38, 0x39, 0x3A, 0x3B, 0x3C, 0x3D, 0x3E,
	0x22, 0x23, 0x24, 0x25, 0x26, 0x27, 0x28, 0x29, 0x2A, 0x2B, 0x2C, 0x2D, 0x2E,
	0x12, 0x13, 0x14, 0x15, 0x16, 0x17, 0x18, 0x19, 0x1A, 0x1B, 0x1C, 0x1D, 0x1E,
	0x2, 0x3, 0x4, 0x5, 0x6, 0x7, 0x8, 0x9, 0x0A, 0x0B, 0x0C, 0x0D, 0x0E,
}

//牌类型["乌龙", "对子", "两对", "三条", "顺子", "同花", "葫芦", "铁支", "同花顺", "五同"]
const (
	WU_LONG       = 0  //乌龙
	DUI_ZI        = 1  //对子
	LIANG_DUI     = 2  //两对
	SAN_TIAO      = 3  //三条
	SHUANG_WANG   = 4  //双王
	SAN_WANG      = 5  //三王冲头
	SHUN_ZI       = 6  //顺子
	TONG_HUA      = 7  //同花
	HU_LU         = 8  //葫芦
	TIE_ZHI       = 9  //铁支
	TONG_HUA_SHUN = 10 //同花顺
	WU_TONG       = 11 //五同
)

const (
	WU           = iota //无
	LIU_DUI_BAN         //六对半
	SAN_SHUN_ZI         //三顺子
	SAN_TONG_HUA        //三同花
	SI_SAN_TIAO         //四三条
	LIU_TONG            //六同
	YI_TIAO_LONG        //一条龙
	QI_TONG             //七同
	BA_TONG             //八同
	QING_LONG           //清龙
)

//QuanLeiDaSpecialMap 全垒打特殊牌映射
var QuanLeiDaSpecialMap = map[int]int{
	LIU_TONG:     1, //六同
	YI_TIAO_LONG: 1, //一条龙
	QI_TONG:      1, //七同
	BA_TONG:      1, //八同
	QING_LONG:    1, //清龙
}

//SpecialCardsScoreMap 特殊牌型分数映射(六同开始 有通杀)
var SpecialCardsScoreMap = map[int]int{
	0:            0,  //无
	LIU_DUI_BAN:  3,  //六对半
	SAN_SHUN_ZI:  3,  //三顺子
	SAN_TONG_HUA: 3,  //三同花
	SI_SAN_TIAO:  18, //四三条
	LIU_TONG:     20, //六同
	YI_TIAO_LONG: 26, //一条龙
	QI_TONG:      40, //七同
	QING_LONG:    52, //清龙
	BA_TONG:      80, //八同
}

//SpecialCardWeightMap 特殊权重映射
var SpecialCardWeightMap = map[int]int{
	LIU_DUI_BAN:  1, //六对半
	SAN_SHUN_ZI:  1, //三顺子
	SAN_TONG_HUA: 1, //三同花
	SI_SAN_TIAO:  2, //四三条
	LIU_TONG:     3, //六同
	YI_TIAO_LONG: 4, //一条龙
	QI_TONG:      5, //七同
	QING_LONG:    6, //清龙
	BA_TONG:      7, //八同
}

//CardsTypeScoreMap 牌型分数 "乌龙", "对子", "两对", "三条", "顺子", "同花", "葫芦", "铁支", "同花顺", "五同"
var CardsTypeScoreMap = map[int][]int{
	WU_LONG:       {1, 1, 1},   //乌龙
	DUI_ZI:        {1, 1, 1},   //对子
	LIANG_DUI:     {1, 1, 1},   //两对
	SAN_TIAO:      {3, 1, 1},   //三条
	SHUANG_WANG:   {20, 1, 1},  //双王
	SAN_WANG:      {40, 1, 1},  //三王
	SHUN_ZI:       {1, 1, 1},   //顺子
	TONG_HUA:      {1, 1, 1},   //同花
	HU_LU:         {1, 2, 1},   //葫芦
	TIE_ZHI:       {1, 8, 4},   //铁支
	TONG_HUA_SHUN: {1, 10, 5},  //同花顺
	WU_TONG:       {1, 20, 10}, //五同
}

//地顺
var DI_SHUN_CARDS_MAP = map[int]map[int]int{
	4: {
		0x2: 1,
		0x3: 1,
		0x4: 1,
		0x5: 1,
		//	0xe: 1,
	},
	3: {
		0x23: 1,
		0x24: 1,
		0x25: 1,
		0x2e: 1,
		0x34: 1,
		0x35: 1,
		0x3e: 1,
		0x45: 1,
		0x4e: 1,
		0x5e: 1,
	},
	2: {
		0x234: 1,
		0x235: 1,
		0x23e: 1,
		0x245: 1,
		0x24e: 1,
		0x25e: 1,
		0x345: 1,
		0x34e: 1,
		0x35e: 1,
		0x45e: 1,
	},
	1: {
		0x2345: 1,
		0x234e: 1,
		0x235e: 1,
		0x245e: 1,
		0x345e: 1,
	},
	0: {
		0x2345E: 1,
	},
}

//天顺
var DI_SHUN_CARDS_3_MAP = map[int]map[int]int{
	0: {
		0x23E: 1,
	},
	1: {
		0x2E: 1,
		0x23: 1,
		0x3e: 1,
	},
	2: {
		0x2: 1,
		0x3: 1,
		0xE: 1,
	},
}
