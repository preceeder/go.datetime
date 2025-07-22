package datetime

var constellations = []struct {
	startMonth, startDay int
	endMonth, endDay     int
}{
	{3, 21, 4, 19},   // Aries
	{4, 20, 5, 20},   // Taurus
	{5, 21, 6, 21},   // Gemini
	{6, 22, 7, 22},   // Cancer
	{7, 23, 8, 22},   // Leo
	{8, 23, 9, 22},   // Virgo
	{9, 23, 10, 23},  // Libra
	{10, 24, 11, 22}, // Scorpio
	{11, 23, 12, 21}, // Sagittarius
	{12, 22, 1, 19},  // Capricorn
	{1, 20, 2, 18},   // Aquarius
	{2, 19, 3, 20},   // Pisces
}

var constellationsSlices = []string{"白羊座", "金牛座", "双子座", "巨蟹座", "狮子座", "处女座", "天秤座", "天蝎座", "射手座", "摩羯座", "水瓶座", "双鱼座"}

func (t *DateTime) Constellation() string {
	index := -1
	_, month, day := t.time.Date()
	for i := 0; i < len(constellations); i++ {
		constellation := constellations[i]
		if int(month) == constellation.startMonth && day >= constellation.startDay {
			index = i
		}
		if int(month) == constellation.endMonth && day <= constellation.endDay {
			index = i
		}
	}
	return constellationsSlices[index]
}
