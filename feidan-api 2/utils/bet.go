package utils

type (
	PumInfo struct {
		Game     string
		Contents string
		Multiple int32
		Title    string
		State    int32
	}
)

var (
	gameMap = map[string]string{
		"168_xysc": "XYSSC",
		"168_xyft": "XYFT",
		"168_azsc": "AULUCKY5",
		"168_azpk": "AULUCKY10",
		"168_jssc": "SSCJSC",
		"168_jspk": "PK10JSC",
		"168_jsft": "LUCKYSB",
		"168_sgsc": "",
		"168_sgft": "",
	}

	pumMap = map[string]*PumInfo{
		// 第一球
		"1000000": {Game: "B1", Contents: "0"},
		"1000001": {Game: "B1", Contents: "1"},
		"1000002": {Game: "B1", Contents: "2"},
		"1000003": {Game: "B1", Contents: "3"},
		"1000004": {Game: "B1", Contents: "4"},
		"1000005": {Game: "B1", Contents: "5"},
		"1000006": {Game: "B1", Contents: "6"},
		"1000007": {Game: "B1", Contents: "7"},
		"1000008": {Game: "B1", Contents: "8"},
		"1000009": {Game: "B1", Contents: "9"},
		"1005001": {Game: "DX1", Contents: "D"},
		"1005002": {Game: "DX1", Contents: "X"},
		"1006001": {Game: "DS1", Contents: "D"},
		"1006002": {Game: "DS1", Contents: "S"},

		// 第二球
		"1001000": {Game: "B2", Contents: "0"},
		"1001001": {Game: "B2", Contents: "1"},
		"1001002": {Game: "B2", Contents: "2"},
		"1001003": {Game: "B2", Contents: "3"},
		"1001004": {Game: "B2", Contents: "4"},
		"1001005": {Game: "B2", Contents: "5"},
		"1001006": {Game: "B2", Contents: "6"},
		"1001007": {Game: "B2", Contents: "7"},
		"1001008": {Game: "B2", Contents: "8"},
		"1001009": {Game: "B2", Contents: "9"},
		"1007001": {Game: "DX2", Contents: "D"},
		"1007002": {Game: "DX2", Contents: "X"},
		"1008001": {Game: "DS2", Contents: "D"},
		"1008002": {Game: "DS2", Contents: "S"},

		// 第三球
		"1002000": {Game: "B3", Contents: "0"},
		"1002001": {Game: "B3", Contents: "1"},
		"1002002": {Game: "B3", Contents: "2"},
		"1002003": {Game: "B3", Contents: "3"},
		"1002004": {Game: "B3", Contents: "4"},
		"1002005": {Game: "B3", Contents: "5"},
		"1002006": {Game: "B3", Contents: "6"},
		"1002007": {Game: "B3", Contents: "7"},
		"1002008": {Game: "B3", Contents: "8"},
		"1002009": {Game: "B3", Contents: "9"},
		"1009001": {Game: "DX3", Contents: "D"},
		"1009002": {Game: "DX3", Contents: "X"},
		"1010001": {Game: "DS3", Contents: "D"},
		"1010002": {Game: "DS3", Contents: "S"},

		// 第四球
		"1003000": {Game: "B4", Contents: "0"},
		"1003001": {Game: "B4", Contents: "1"},
		"1003002": {Game: "B4", Contents: "2"},
		"1003003": {Game: "B4", Contents: "3"},
		"1003004": {Game: "B4", Contents: "4"},
		"1003005": {Game: "B4", Contents: "5"},
		"1003006": {Game: "B4", Contents: "6"},
		"1003007": {Game: "B4", Contents: "7"},
		"1003008": {Game: "B4", Contents: "8"},
		"1003009": {Game: "B4", Contents: "9"},
		"1011001": {Game: "DX4", Contents: "D"},
		"1011002": {Game: "DX4", Contents: "X"},
		"1012001": {Game: "DS4", Contents: "D"},
		"1012002": {Game: "DS4", Contents: "S"},

		// 第五球
		"1004000": {Game: "B5", Contents: "0"},
		"1004001": {Game: "B5", Contents: "1"},
		"1004002": {Game: "B5", Contents: "2"},
		"1004003": {Game: "B5", Contents: "3"},
		"1004004": {Game: "B5", Contents: "4"},
		"1004005": {Game: "B5", Contents: "5"},
		"1004006": {Game: "B5", Contents: "6"},
		"1004007": {Game: "B5", Contents: "7"},
		"1004008": {Game: "B5", Contents: "8"},
		"1004009": {Game: "B5", Contents: "9"},
		"1013001": {Game: "DX5", Contents: "D"},
		"1013002": {Game: "DX5", Contents: "X"},
		"1014001": {Game: "DS5", Contents: "D"},
		"1014002": {Game: "DS5", Contents: "S"},

		// 两面
		"1015001": {Game: "ZDX", Contents: "D"},
		"1015002": {Game: "ZDX", Contents: "X"},
		"1016001": {Game: "ZDS", Contents: "D"},
		"1016002": {Game: "ZDS", Contents: "S"},
		"1017001": {Game: "LH", Contents: "L"},
		"1017002": {Game: "LH", Contents: "H"},
		"1017003": {Game: "LH", Contents: "T"},

		"1018001": {Game: "TS1", Contents: "0"},
		"1019001": {Game: "TS1", Contents: "1"},
		"1020001": {Game: "TS1", Contents: "2"},
		"1021001": {Game: "TS1", Contents: "3"},
		"1022001": {Game: "TS1", Contents: "4"},

		"1023001": {Game: "TS2", Contents: "0"},
		"1024001": {Game: "TS2", Contents: "1"},
		"1025001": {Game: "TS2", Contents: "2"},
		"1026001": {Game: "TS2", Contents: "3"},
		"1027001": {Game: "TS2", Contents: "4"},

		"1028001": {Game: "TS3", Contents: "0"},
		"1029001": {Game: "TS3", Contents: "1"},
		"1030001": {Game: "TS3", Contents: "2"},
		"1031001": {Game: "TS3", Contents: "3"},
		"1032001": {Game: "TS3", Contents: "4"},

		// 二字定位
		"105000": {Game: "DW54", Multiple: 1, Title: "万千定位"},
		"105100": {Game: "DW53", Multiple: 1, Title: "万佰定位"},
		"105200": {Game: "DW52", Multiple: 1, Title: "万拾定位"},
		"105300": {Game: "DW51", Multiple: 1, Title: "万个定位"},
		"105400": {Game: "DW43", Multiple: 1, Title: "千佰定位"},
		"105500": {Game: "DW42", Multiple: 1, Title: "千拾定位"},
		"105600": {Game: "DW41", Multiple: 1, Title: "千个定位"},
		"105700": {Game: "DW32", Multiple: 1, Title: "佰拾定位"},
		"105800": {Game: "DW31", Multiple: 1, Title: "佰个定位"},
		"105900": {Game: "DW21", Multiple: 1, Title: "拾个定位"},

		// 三字定位
		"106000": {Game: "DW543", Multiple: 1, Title: "前三定位"},
		"106100": {Game: "DW432", Multiple: 1, Title: "中三定位"},
		"106200": {Game: "DW321", Multiple: 1, Title: "后三定位"},

		// 冠军
		"2010001": {Game: "DX1", Contents: "D"},
		"2010002": {Game: "DX1", Contents: "X"},
		"2011001": {Game: "DS1", Contents: "D"},
		"2011002": {Game: "DS1", Contents: "S"},
		"2012001": {Game: "LH1", Contents: "L"},
		"2012002": {Game: "LH1", Contents: "H"},

		// 亚军
		"2013001": {Game: "DX2", Contents: "D"},
		"2013002": {Game: "DX2", Contents: "X"},
		"2014001": {Game: "DS2", Contents: "D"},
		"2014002": {Game: "DS2", Contents: "S"},
		"2015001": {Game: "LH2", Contents: "L"},
		"2015002": {Game: "LH2", Contents: "H"},

		// 第三名
		"2016001": {Game: "DX3", Contents: "D"},
		"2016002": {Game: "DX3", Contents: "X"},
		"2017001": {Game: "DS3", Contents: "D"},
		"2017002": {Game: "DS3", Contents: "S"},
		"2018001": {Game: "LH3", Contents: "L"},
		"2018002": {Game: "LH3", Contents: "H"},

		// 第四名
		"2019001": {Game: "DX4", Contents: "D"},
		"2019002": {Game: "DX4", Contents: "X"},
		"2020001": {Game: "DS4", Contents: "D"},
		"2020002": {Game: "DS4", Contents: "S"},
		"2021001": {Game: "LH4", Contents: "L"},
		"2021002": {Game: "LH4", Contents: "H"},

		// 第五名
		"2022001": {Game: "DX5", Contents: "D"},
		"2022002": {Game: "DX5", Contents: "X"},
		"2023001": {Game: "DS5", Contents: "D"},
		"2023002": {Game: "DS5", Contents: "S"},
		"2024001": {Game: "LH5", Contents: "L"},
		"2024002": {Game: "LH5", Contents: "H"},

		// 第六名
		"2025001": {Game: "DX6", Contents: "D"},
		"2025002": {Game: "DX6", Contents: "X"},
		"2026001": {Game: "DS6", Contents: "D"},
		"2026002": {Game: "DS6", Contents: "S"},

		// 第七名
		"2027001": {Game: "DX7", Contents: "D"},
		"2027002": {Game: "DX7", Contents: "X"},
		"2028001": {Game: "DS7", Contents: "D"},
		"2028002": {Game: "DS7", Contents: "S"},

		// 第八名
		"2029001": {Game: "DX8", Contents: "D"},
		"2029002": {Game: "DX8", Contents: "X"},
		"2030001": {Game: "DS8", Contents: "D"},
		"2030002": {Game: "DS8", Contents: "S"},

		// 第九名
		"2031001": {Game: "DX9", Contents: "D"},
		"2031002": {Game: "DX9", Contents: "X"},
		"2032001": {Game: "DS9", Contents: "D"},
		"2032002": {Game: "DS9", Contents: "S"},

		// 第十个名
		"2033001": {Game: "DX10", Contents: "D"},
		"2033002": {Game: "DX10", Contents: "X"},
		"2034001": {Game: "DS10", Contents: "D"},
		"2034002": {Game: "DS10", Contents: "S"},

		// 1-10名
		// 冠军
		"2000001": {Game: "B1", Contents: "1"},
		"2000002": {Game: "B1", Contents: "2"},
		"2000003": {Game: "B1", Contents: "3"},
		"2000004": {Game: "B1", Contents: "4"},
		"2000005": {Game: "B1", Contents: "5"},
		"2000006": {Game: "B1", Contents: "6"},
		"2000007": {Game: "B1", Contents: "7"},
		"2000008": {Game: "B1", Contents: "8"},
		"2000009": {Game: "B1", Contents: "9"},
		"2000010": {Game: "B1", Contents: "10"},

		// 亚军
		"2001001": {Game: "B2", Contents: "1"},
		"2001002": {Game: "B2", Contents: "2"},
		"2001003": {Game: "B2", Contents: "3"},
		"2001004": {Game: "B2", Contents: "4"},
		"2001005": {Game: "B2", Contents: "5"},
		"2001006": {Game: "B2", Contents: "6"},
		"2001007": {Game: "B2", Contents: "7"},
		"2001008": {Game: "B2", Contents: "8"},
		"2001009": {Game: "B2", Contents: "9"},
		"2001010": {Game: "B2", Contents: "10"},

		"2002001": {Game: "B3", Contents: "1"},
		"2002002": {Game: "B3", Contents: "2"},
		"2002003": {Game: "B3", Contents: "3"},
		"2002004": {Game: "B3", Contents: "4"},
		"2002005": {Game: "B3", Contents: "5"},
		"2002006": {Game: "B3", Contents: "6"},
		"2002007": {Game: "B3", Contents: "7"},
		"2002008": {Game: "B3", Contents: "8"},
		"2002009": {Game: "B3", Contents: "9"},
		"2002010": {Game: "B3", Contents: "10"},

		"2003001": {Game: "B4", Contents: "1"},
		"2003002": {Game: "B4", Contents: "2"},
		"2003003": {Game: "B4", Contents: "3"},
		"2003004": {Game: "B4", Contents: "4"},
		"2003005": {Game: "B4", Contents: "5"},
		"2003006": {Game: "B4", Contents: "6"},
		"2003007": {Game: "B4", Contents: "7"},
		"2003008": {Game: "B4", Contents: "8"},
		"2003009": {Game: "B4", Contents: "9"},
		"2003010": {Game: "B4", Contents: "10"},

		"2004001": {Game: "B5", Contents: "1"},
		"2004002": {Game: "B5", Contents: "2"},
		"2004003": {Game: "B5", Contents: "3"},
		"2004004": {Game: "B5", Contents: "4"},
		"2004005": {Game: "B5", Contents: "5"},
		"2004006": {Game: "B5", Contents: "6"},
		"2004007": {Game: "B5", Contents: "7"},
		"2004008": {Game: "B5", Contents: "8"},
		"2004009": {Game: "B5", Contents: "9"},
		"2004010": {Game: "B5", Contents: "10"},

		"2005001": {Game: "B6", Contents: "1"},
		"2005002": {Game: "B6", Contents: "2"},
		"2005003": {Game: "B6", Contents: "3"},
		"2005004": {Game: "B6", Contents: "4"},
		"2005005": {Game: "B6", Contents: "5"},
		"2005006": {Game: "B6", Contents: "6"},
		"2005007": {Game: "B6", Contents: "7"},
		"2005008": {Game: "B6", Contents: "8"},
		"2005009": {Game: "B6", Contents: "9"},
		"2005010": {Game: "B6", Contents: "10"},

		"2006001": {Game: "B7", Contents: "1"},
		"2006002": {Game: "B7", Contents: "2"},
		"2006003": {Game: "B7", Contents: "3"},
		"2006004": {Game: "B7", Contents: "4"},
		"2006005": {Game: "B7", Contents: "5"},
		"2006006": {Game: "B7", Contents: "6"},
		"2006007": {Game: "B7", Contents: "7"},
		"2006008": {Game: "B7", Contents: "8"},
		"2006009": {Game: "B7", Contents: "9"},
		"2006010": {Game: "B7", Contents: "10"},

		"2007001": {Game: "B8", Contents: "1"},
		"2007002": {Game: "B8", Contents: "2"},
		"2007003": {Game: "B8", Contents: "3"},
		"2007004": {Game: "B8", Contents: "4"},
		"2007005": {Game: "B8", Contents: "5"},
		"2007006": {Game: "B8", Contents: "6"},
		"2007007": {Game: "B8", Contents: "7"},
		"2007008": {Game: "B8", Contents: "8"},
		"2007009": {Game: "B8", Contents: "9"},
		"2007010": {Game: "B8", Contents: "10"},

		"2008001": {Game: "B9", Contents: "1"},
		"2008002": {Game: "B9", Contents: "2"},
		"2008003": {Game: "B9", Contents: "3"},
		"2008004": {Game: "B9", Contents: "4"},
		"2008005": {Game: "B9", Contents: "5"},
		"2008006": {Game: "B9", Contents: "6"},
		"2008007": {Game: "B9", Contents: "7"},
		"2008008": {Game: "B9", Contents: "8"},
		"2008009": {Game: "B9", Contents: "9"},
		"2008010": {Game: "B9", Contents: "10"},

		"2009001": {Game: "B10", Contents: "1"},
		"2009002": {Game: "B10", Contents: "2"},
		"2009003": {Game: "B10", Contents: "3"},
		"2009004": {Game: "B10", Contents: "4"},
		"2009005": {Game: "B10", Contents: "5"},
		"2009006": {Game: "B10", Contents: "6"},
		"2009007": {Game: "B10", Contents: "7"},
		"2009008": {Game: "B10", Contents: "8"},
		"2009009": {Game: "B10", Contents: "9"},
		"2009010": {Game: "B10", Contents: "10"},

		// 冠亚和
		"2035003": {Game: "GYH", Contents: "3"},
		"2035004": {Game: "GYH", Contents: "4"},
		"2035005": {Game: "GYH", Contents: "5"},
		"2035006": {Game: "GYH", Contents: "6"},
		"2035007": {Game: "GYH", Contents: "7"},
		"2035008": {Game: "GYH", Contents: "8"},
		"2035009": {Game: "GYH", Contents: "9"},
		"2035010": {Game: "GYH", Contents: "10"},
		"2035011": {Game: "GYH", Contents: "11"},
		"2035012": {Game: "GYH", Contents: "12"},
		"2035013": {Game: "GYH", Contents: "13"},
		"2035014": {Game: "GYH", Contents: "14"},
		"2035015": {Game: "GYH", Contents: "15"},
		"2035016": {Game: "GYH", Contents: "16"},
		"2035017": {Game: "GYH", Contents: "17"},
		"2035018": {Game: "GYH", Contents: "18"},
		"2035019": {Game: "GYH", Contents: "19"},
		"2036001": {Game: "GDX", Contents: "D"},
		"2036002": {Game: "GDX", Contents: "X"},
		"2037001": {Game: "GDS", Contents: "D"},
		"2037002": {Game: "GDS", Contents: "S"},
	}
)

func GetGameName(game string) string {
	return gameMap[game]
}

func GetPumInfo(pum string) *PumInfo {
	return pumMap[pum]
}
