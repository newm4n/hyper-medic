package code

import (
	"fmt"
	"strings"
)

type ApproachSiteCodes string

const (
	ApproachSiteCodes91723000 ApproachSiteCodes = "91723000"
	ApproachSiteCodes106004   ApproachSiteCodes = "106004"
	ApproachSiteCodes107008   ApproachSiteCodes = "107008"
	ApproachSiteCodes108003   ApproachSiteCodes = "108003"
	ApproachSiteCodes110001   ApproachSiteCodes = "110001"
	ApproachSiteCodes111002   ApproachSiteCodes = "111002"
	ApproachSiteCodes116007   ApproachSiteCodes = "116007"
	ApproachSiteCodes124002   ApproachSiteCodes = "124002"
	ApproachSiteCodes149003   ApproachSiteCodes = "149003"
	ApproachSiteCodes155008   ApproachSiteCodes = "155008"
	ApproachSiteCodes167005   ApproachSiteCodes = "167005"
	ApproachSiteCodes202009   ApproachSiteCodes = "202009"
	ApproachSiteCodes205006   ApproachSiteCodes = "205006"
	ApproachSiteCodes206007   ApproachSiteCodes = "206007"
	ApproachSiteCodes221001   ApproachSiteCodes = "221001"
	ApproachSiteCodes227002   ApproachSiteCodes = "227002"
	ApproachSiteCodes233006   ApproachSiteCodes = "233006"
	ApproachSiteCodes235004   ApproachSiteCodes = "235004"
	ApproachSiteCodes246001   ApproachSiteCodes = "246001"
	ApproachSiteCodes247005   ApproachSiteCodes = "247005"
	ApproachSiteCodes251007   ApproachSiteCodes = "251007"
	ApproachSiteCodes256002   ApproachSiteCodes = "256002"
	ApproachSiteCodes263002   ApproachSiteCodes = "263002"
	ApproachSiteCodes266005   ApproachSiteCodes = "266005"
	ApproachSiteCodes272005   ApproachSiteCodes = "272005"
	ApproachSiteCodes273000   ApproachSiteCodes = "273000"
	ApproachSiteCodes283001   ApproachSiteCodes = "283001"
	ApproachSiteCodes284007   ApproachSiteCodes = "284007"
	ApproachSiteCodes289002   ApproachSiteCodes = "289002"
	ApproachSiteCodes301000   ApproachSiteCodes = "301000"
	ApproachSiteCodes311007   ApproachSiteCodes = "311007"
	ApproachSiteCodes315003   ApproachSiteCodes = "315003"
	ApproachSiteCodes318001   ApproachSiteCodes = "318001"
	ApproachSiteCodes344001   ApproachSiteCodes = "344001"
	ApproachSiteCodes345000   ApproachSiteCodes = "345000"
	ApproachSiteCodes356000   ApproachSiteCodes = "356000"
	ApproachSiteCodes393006   ApproachSiteCodes = "393006"
	ApproachSiteCodes402006   ApproachSiteCodes = "402006"
	ApproachSiteCodes405008   ApproachSiteCodes = "405008"
	ApproachSiteCodes414003   ApproachSiteCodes = "414003"
	ApproachSiteCodes422005   ApproachSiteCodes = "422005"
	ApproachSiteCodes446003   ApproachSiteCodes = "446003"
	ApproachSiteCodes457008   ApproachSiteCodes = "457008"
	ApproachSiteCodes461002   ApproachSiteCodes = "461002"
	ApproachSiteCodes464005   ApproachSiteCodes = "464005"
	ApproachSiteCodes465006   ApproachSiteCodes = "465006"
	ApproachSiteCodes471000   ApproachSiteCodes = "471000"
	ApproachSiteCodes480000   ApproachSiteCodes = "480000"
	ApproachSiteCodes485005   ApproachSiteCodes = "485005"
	ApproachSiteCodes528006   ApproachSiteCodes = "528006"
	ApproachSiteCodes552004   ApproachSiteCodes = "552004"
	ApproachSiteCodes565008   ApproachSiteCodes = "565008"
	ApproachSiteCodes582005   ApproachSiteCodes = "582005"
	ApproachSiteCodes587004   ApproachSiteCodes = "587004"
	ApproachSiteCodes589001   ApproachSiteCodes = "589001"
	ApproachSiteCodes595000   ApproachSiteCodes = "595000"
	ApproachSiteCodes608002   ApproachSiteCodes = "608002"
	ApproachSiteCodes621009   ApproachSiteCodes = "621009"
	ApproachSiteCodes635006   ApproachSiteCodes = "635006"
	ApproachSiteCodes650002   ApproachSiteCodes = "650002"
	ApproachSiteCodes660006   ApproachSiteCodes = "660006"
	ApproachSiteCodes661005   ApproachSiteCodes = "661005"
	ApproachSiteCodes667009   ApproachSiteCodes = "667009"
	ApproachSiteCodes688000   ApproachSiteCodes = "688000"
	ApproachSiteCodes691000   ApproachSiteCodes = "691000"
	ApproachSiteCodes692007   ApproachSiteCodes = "692007"
	ApproachSiteCodes723004   ApproachSiteCodes = "723004"
	ApproachSiteCodes774007   ApproachSiteCodes = "774007"
	ApproachSiteCodes790007   ApproachSiteCodes = "790007"
	ApproachSiteCodes798000   ApproachSiteCodes = "798000"
	ApproachSiteCodes808000   ApproachSiteCodes = "808000"
	ApproachSiteCodes809008   ApproachSiteCodes = "809008"
	ApproachSiteCodes823005   ApproachSiteCodes = "823005"
	ApproachSiteCodes830004   ApproachSiteCodes = "830004"
	ApproachSiteCodes836005   ApproachSiteCodes = "836005"
	ApproachSiteCodes885000   ApproachSiteCodes = "885000"
	ApproachSiteCodes895007   ApproachSiteCodes = "895007"
	ApproachSiteCodes917004   ApproachSiteCodes = "917004"
	ApproachSiteCodes921006   ApproachSiteCodes = "921006"
	ApproachSiteCodes947002   ApproachSiteCodes = "947002"
	ApproachSiteCodes955009   ApproachSiteCodes = "955009"
	ApproachSiteCodes976004   ApproachSiteCodes = "976004"
	ApproachSiteCodes996007   ApproachSiteCodes = "996007"
	ApproachSiteCodes1005009  ApproachSiteCodes = "1005009"
	ApproachSiteCodes1012000  ApproachSiteCodes = "1012000"
	ApproachSiteCodes1015003  ApproachSiteCodes = "1015003"
	ApproachSiteCodes1028005  ApproachSiteCodes = "1028005"
	ApproachSiteCodes1030007  ApproachSiteCodes = "1030007"
	ApproachSiteCodes1063000  ApproachSiteCodes = "1063000"
	ApproachSiteCodes1075005  ApproachSiteCodes = "1075005"
	ApproachSiteCodes1076006  ApproachSiteCodes = "1076006"
	ApproachSiteCodes1086007  ApproachSiteCodes = "1086007"
	ApproachSiteCodes1087003  ApproachSiteCodes = "1087003"
	ApproachSiteCodes1092001  ApproachSiteCodes = "1092001"
	ApproachSiteCodes1099005  ApproachSiteCodes = "1099005"
	ApproachSiteCodes1101003  ApproachSiteCodes = "1101003"
	ApproachSiteCodes1106008  ApproachSiteCodes = "1106008"
	ApproachSiteCodes1110006  ApproachSiteCodes = "1110006"
	ApproachSiteCodes1122009  ApproachSiteCodes = "1122009"
	ApproachSiteCodes1136004  ApproachSiteCodes = "1136004"
	ApproachSiteCodes1159005  ApproachSiteCodes = "1159005"
	ApproachSiteCodes1172006  ApproachSiteCodes = "1172006"
	ApproachSiteCodes1173001  ApproachSiteCodes = "1173001"
	ApproachSiteCodes1174007  ApproachSiteCodes = "1174007"
	ApproachSiteCodes1193009  ApproachSiteCodes = "1193009"
	ApproachSiteCodes1216008  ApproachSiteCodes = "1216008"
	ApproachSiteCodes1231004  ApproachSiteCodes = "1231004"
	ApproachSiteCodes1236009  ApproachSiteCodes = "1236009"
	ApproachSiteCodes1243003  ApproachSiteCodes = "1243003"
	ApproachSiteCodes1246006  ApproachSiteCodes = "1246006"
	ApproachSiteCodes1263005  ApproachSiteCodes = "1263005"
	ApproachSiteCodes1277008  ApproachSiteCodes = "1277008"
	ApproachSiteCodes1307006  ApproachSiteCodes = "1307006"
	ApproachSiteCodes1311000  ApproachSiteCodes = "1311000"
	ApproachSiteCodes1350001  ApproachSiteCodes = "1350001"
	ApproachSiteCodes1353004  ApproachSiteCodes = "1353004"
	ApproachSiteCodes1403006  ApproachSiteCodes = "1403006"
	ApproachSiteCodes1425000  ApproachSiteCodes = "1425000"
	ApproachSiteCodes1439000  ApproachSiteCodes = "1439000"
	ApproachSiteCodes1441004  ApproachSiteCodes = "1441004"
	ApproachSiteCodes1456008  ApproachSiteCodes = "1456008"
	ApproachSiteCodes1467009  ApproachSiteCodes = "1467009"
	ApproachSiteCodes1484003  ApproachSiteCodes = "1484003"
	ApproachSiteCodes1490004  ApproachSiteCodes = "1490004"
	ApproachSiteCodes1502004  ApproachSiteCodes = "1502004"
	ApproachSiteCodes1511004  ApproachSiteCodes = "1511004"
	ApproachSiteCodes1516009  ApproachSiteCodes = "1516009"
	ApproachSiteCodes1527006  ApproachSiteCodes = "1527006"
	ApproachSiteCodes1537001  ApproachSiteCodes = "1537001"
	ApproachSiteCodes1541002  ApproachSiteCodes = "1541002"
	ApproachSiteCodes1562001  ApproachSiteCodes = "1562001"
	ApproachSiteCodes1580005  ApproachSiteCodes = "1580005"
	ApproachSiteCodes1581009  ApproachSiteCodes = "1581009"
	ApproachSiteCodes1584001  ApproachSiteCodes = "1584001"
	ApproachSiteCodes1600003  ApproachSiteCodes = "1600003"
	ApproachSiteCodes1605008  ApproachSiteCodes = "1605008"
	ApproachSiteCodes1610007  ApproachSiteCodes = "1610007"
	ApproachSiteCodes1611006  ApproachSiteCodes = "1611006"
	ApproachSiteCodes1617005  ApproachSiteCodes = "1617005"
	ApproachSiteCodes1620002  ApproachSiteCodes = "1620002"
	ApproachSiteCodes1626008  ApproachSiteCodes = "1626008"
	ApproachSiteCodes1627004  ApproachSiteCodes = "1627004"
	ApproachSiteCodes1630006  ApproachSiteCodes = "1630006"
	ApproachSiteCodes1631005  ApproachSiteCodes = "1631005"
	ApproachSiteCodes1650005  ApproachSiteCodes = "1650005"
	ApproachSiteCodes1655000  ApproachSiteCodes = "1655000"
	ApproachSiteCodes1659006  ApproachSiteCodes = "1659006"
	ApproachSiteCodes1684009  ApproachSiteCodes = "1684009"
	ApproachSiteCodes1706004  ApproachSiteCodes = "1706004"
	ApproachSiteCodes1707008  ApproachSiteCodes = "1707008"
	ApproachSiteCodes1711002  ApproachSiteCodes = "1711002"
	ApproachSiteCodes1716007  ApproachSiteCodes = "1716007"
	ApproachSiteCodes1721005  ApproachSiteCodes = "1721005"
	ApproachSiteCodes1729007  ApproachSiteCodes = "1729007"
	ApproachSiteCodes1732005  ApproachSiteCodes = "1732005"
	ApproachSiteCodes1765002  ApproachSiteCodes = "1765002"
	ApproachSiteCodes1780008  ApproachSiteCodes = "1780008"
	ApproachSiteCodes1781007  ApproachSiteCodes = "1781007"
	ApproachSiteCodes1797002  ApproachSiteCodes = "1797002"
	ApproachSiteCodes1818002  ApproachSiteCodes = "1818002"
	ApproachSiteCodes1825009  ApproachSiteCodes = "1825009"
	ApproachSiteCodes1832000  ApproachSiteCodes = "1832000"
	ApproachSiteCodes1840006  ApproachSiteCodes = "1840006"
	ApproachSiteCodes1849007  ApproachSiteCodes = "1849007"
	ApproachSiteCodes1853009  ApproachSiteCodes = "1853009"
	ApproachSiteCodes1874005  ApproachSiteCodes = "1874005"
	ApproachSiteCodes1895000  ApproachSiteCodes = "1895000"
	ApproachSiteCodes1902009  ApproachSiteCodes = "1902009"
	ApproachSiteCodes1910005  ApproachSiteCodes = "1910005"
	ApproachSiteCodes1918003  ApproachSiteCodes = "1918003"
	ApproachSiteCodes1927002  ApproachSiteCodes = "1927002"
	ApproachSiteCodes1992003  ApproachSiteCodes = "1992003"
	ApproachSiteCodes1997009  ApproachSiteCodes = "1997009"
	ApproachSiteCodes2010005  ApproachSiteCodes = "2010005"
	ApproachSiteCodes2020000  ApproachSiteCodes = "2020000"
	ApproachSiteCodes2031008  ApproachSiteCodes = "2031008"
	ApproachSiteCodes2033006  ApproachSiteCodes = "2033006"
	ApproachSiteCodes2044003  ApproachSiteCodes = "2044003"
	ApproachSiteCodes2048000  ApproachSiteCodes = "2048000"
	ApproachSiteCodes2049008  ApproachSiteCodes = "2049008"
	ApproachSiteCodes2059009  ApproachSiteCodes = "2059009"
	ApproachSiteCodes2071003  ApproachSiteCodes = "2071003"
	ApproachSiteCodes2076008  ApproachSiteCodes = "2076008"
	ApproachSiteCodes2083001  ApproachSiteCodes = "2083001"
	ApproachSiteCodes2095001  ApproachSiteCodes = "2095001"
	ApproachSiteCodes2123001  ApproachSiteCodes = "2123001"
	ApproachSiteCodes2150006  ApproachSiteCodes = "2150006"
	ApproachSiteCodes2156000  ApproachSiteCodes = "2156000"
	ApproachSiteCodes2160002  ApproachSiteCodes = "2160002"
	ApproachSiteCodes2175005  ApproachSiteCodes = "2175005"
	ApproachSiteCodes2182009  ApproachSiteCodes = "2182009"
	ApproachSiteCodes2192001  ApproachSiteCodes = "2192001"
	ApproachSiteCodes2205003  ApproachSiteCodes = "2205003"
	ApproachSiteCodes2209009  ApproachSiteCodes = "2209009"
	ApproachSiteCodes2236006  ApproachSiteCodes = "2236006"
	ApproachSiteCodes2246008  ApproachSiteCodes = "2246008"
	ApproachSiteCodes2255006  ApproachSiteCodes = "2255006"
	ApproachSiteCodes2292006  ApproachSiteCodes = "2292006"
	ApproachSiteCodes2302002  ApproachSiteCodes = "2302002"
	ApproachSiteCodes2305000  ApproachSiteCodes = "2305000"
	ApproachSiteCodes2306004  ApproachSiteCodes = "2306004"
	ApproachSiteCodes2327009  ApproachSiteCodes = "2327009"
	ApproachSiteCodes2330002  ApproachSiteCodes = "2330002"
	ApproachSiteCodes2332005  ApproachSiteCodes = "2332005"
	ApproachSiteCodes2334006  ApproachSiteCodes = "2334006"
	ApproachSiteCodes2349003  ApproachSiteCodes = "2349003"
	ApproachSiteCodes2372001  ApproachSiteCodes = "2372001"
	ApproachSiteCodes2383005  ApproachSiteCodes = "2383005"
	ApproachSiteCodes2389009  ApproachSiteCodes = "2389009"
	ApproachSiteCodes2395005  ApproachSiteCodes = "2395005"
	ApproachSiteCodes2397002  ApproachSiteCodes = "2397002"
	ApproachSiteCodes2400006  ApproachSiteCodes = "2400006"
	ApproachSiteCodes2402003  ApproachSiteCodes = "2402003"
	ApproachSiteCodes2421006  ApproachSiteCodes = "2421006"
	ApproachSiteCodes2433001  ApproachSiteCodes = "2433001"
	ApproachSiteCodes2436009  ApproachSiteCodes = "2436009"
	ApproachSiteCodes2453002  ApproachSiteCodes = "2453002"
	ApproachSiteCodes2454008  ApproachSiteCodes = "2454008"
	ApproachSiteCodes2484000  ApproachSiteCodes = "2484000"
	ApproachSiteCodes2489005  ApproachSiteCodes = "2489005"
	ApproachSiteCodes2499000  ApproachSiteCodes = "2499000"
	ApproachSiteCodes2502001  ApproachSiteCodes = "2502001"
	ApproachSiteCodes2504000  ApproachSiteCodes = "2504000"
	ApproachSiteCodes2510000  ApproachSiteCodes = "2510000"
	ApproachSiteCodes2539000  ApproachSiteCodes = "2539000"
	ApproachSiteCodes2543001  ApproachSiteCodes = "2543001"
	ApproachSiteCodes2550002  ApproachSiteCodes = "2550002"
	ApproachSiteCodes2577006  ApproachSiteCodes = "2577006"
	ApproachSiteCodes2579009  ApproachSiteCodes = "2579009"
	ApproachSiteCodes2592007  ApproachSiteCodes = "2592007"
	ApproachSiteCodes2600000  ApproachSiteCodes = "2600000"
	ApproachSiteCodes2620004  ApproachSiteCodes = "2620004"
	ApproachSiteCodes2639009  ApproachSiteCodes = "2639009"
	ApproachSiteCodes2653009  ApproachSiteCodes = "2653009"
	ApproachSiteCodes2666009  ApproachSiteCodes = "2666009"
	ApproachSiteCodes2672009  ApproachSiteCodes = "2672009"
	ApproachSiteCodes2675006  ApproachSiteCodes = "2675006"
	ApproachSiteCodes2681003  ApproachSiteCodes = "2681003"
	ApproachSiteCodes2682005  ApproachSiteCodes = "2682005"
	ApproachSiteCodes2686008  ApproachSiteCodes = "2686008"
	ApproachSiteCodes2687004  ApproachSiteCodes = "2687004"
	ApproachSiteCodes2695000  ApproachSiteCodes = "2695000"
	ApproachSiteCodes2703009  ApproachSiteCodes = "2703009"
	ApproachSiteCodes2712006  ApproachSiteCodes = "2712006"
	ApproachSiteCodes2718005  ApproachSiteCodes = "2718005"
	ApproachSiteCodes2726002  ApproachSiteCodes = "2726002"
	ApproachSiteCodes2730004  ApproachSiteCodes = "2730004"
	ApproachSiteCodes2739003  ApproachSiteCodes = "2739003"
	ApproachSiteCodes2741002  ApproachSiteCodes = "2741002"
	ApproachSiteCodes2746007  ApproachSiteCodes = "2746007"
	ApproachSiteCodes2748008  ApproachSiteCodes = "2748008"
	ApproachSiteCodes2759004  ApproachSiteCodes = "2759004"
	ApproachSiteCodes2771005  ApproachSiteCodes = "2771005"
	ApproachSiteCodes2789006  ApproachSiteCodes = "2789006"
	ApproachSiteCodes2792005  ApproachSiteCodes = "2792005"
	ApproachSiteCodes2803000  ApproachSiteCodes = "2803000"
	ApproachSiteCodes2810006  ApproachSiteCodes = "2810006"
	ApproachSiteCodes2812003  ApproachSiteCodes = "2812003"
	ApproachSiteCodes2824005  ApproachSiteCodes = "2824005"
	ApproachSiteCodes2826007  ApproachSiteCodes = "2826007"
	ApproachSiteCodes2830005  ApproachSiteCodes = "2830005"
	ApproachSiteCodes2839006  ApproachSiteCodes = "2839006"
	ApproachSiteCodes2841007  ApproachSiteCodes = "2841007"
	ApproachSiteCodes2845003  ApproachSiteCodes = "2845003"
	ApproachSiteCodes2848001  ApproachSiteCodes = "2848001"
	ApproachSiteCodes2855004  ApproachSiteCodes = "2855004"
	ApproachSiteCodes2861001  ApproachSiteCodes = "2861001"
	ApproachSiteCodes2894003  ApproachSiteCodes = "2894003"
	ApproachSiteCodes2905008  ApproachSiteCodes = "2905008"
	ApproachSiteCodes2909002  ApproachSiteCodes = "2909002"
	ApproachSiteCodes2920002  ApproachSiteCodes = "2920002"
	ApproachSiteCodes2922005  ApproachSiteCodes = "2922005"
	ApproachSiteCodes2923000  ApproachSiteCodes = "2923000"
	ApproachSiteCodes2969000  ApproachSiteCodes = "2969000"
	ApproachSiteCodes2979003  ApproachSiteCodes = "2979003"
	ApproachSiteCodes2986006  ApproachSiteCodes = "2986006"
	ApproachSiteCodes2998001  ApproachSiteCodes = "2998001"
	ApproachSiteCodes3003007  ApproachSiteCodes = "3003007"
	ApproachSiteCodes3008003  ApproachSiteCodes = "3008003"
	ApproachSiteCodes3028004  ApproachSiteCodes = "3028004"
	ApproachSiteCodes3039001  ApproachSiteCodes = "3039001"
	ApproachSiteCodes3042007  ApproachSiteCodes = "3042007"
	ApproachSiteCodes3054007  ApproachSiteCodes = "3054007"
	ApproachSiteCodes3055008  ApproachSiteCodes = "3055008"
	ApproachSiteCodes3056009  ApproachSiteCodes = "3056009"
	ApproachSiteCodes3057000  ApproachSiteCodes = "3057000"
	ApproachSiteCodes3058005  ApproachSiteCodes = "3058005"
	ApproachSiteCodes3062004  ApproachSiteCodes = "3062004"
	ApproachSiteCodes3068000  ApproachSiteCodes = "3068000"
	ApproachSiteCodes3081007  ApproachSiteCodes = "3081007"
	ApproachSiteCodes3093003  ApproachSiteCodes = "3093003"
	ApproachSiteCodes3100007  ApproachSiteCodes = "3100007"
	ApproachSiteCodes3113001  ApproachSiteCodes = "3113001"
	ApproachSiteCodes3117000  ApproachSiteCodes = "3117000"
	ApproachSiteCodes3118005  ApproachSiteCodes = "3118005"
	ApproachSiteCodes3120008  ApproachSiteCodes = "3120008"
	ApproachSiteCodes3134008  ApproachSiteCodes = "3134008"
	ApproachSiteCodes3138006  ApproachSiteCodes = "3138006"
	ApproachSiteCodes3153003  ApproachSiteCodes = "3153003"
	ApproachSiteCodes3156006  ApproachSiteCodes = "3156006"
	ApproachSiteCodes3159004  ApproachSiteCodes = "3159004"
	ApproachSiteCodes3169005  ApproachSiteCodes = "3169005"
	ApproachSiteCodes3178004  ApproachSiteCodes = "3178004"
	ApproachSiteCodes3194006  ApproachSiteCodes = "3194006"
	ApproachSiteCodes3198009  ApproachSiteCodes = "3198009"
	ApproachSiteCodes3215002  ApproachSiteCodes = "3215002"
	ApproachSiteCodes3222005  ApproachSiteCodes = "3222005"
	ApproachSiteCodes3227004  ApproachSiteCodes = "3227004"
	ApproachSiteCodes3236000  ApproachSiteCodes = "3236000"
	ApproachSiteCodes3243006  ApproachSiteCodes = "3243006"
	ApproachSiteCodes3255000  ApproachSiteCodes = "3255000"
	ApproachSiteCodes3262009  ApproachSiteCodes = "3262009"
	ApproachSiteCodes3279003  ApproachSiteCodes = "3279003"
	ApproachSiteCodes3295003  ApproachSiteCodes = "3295003"
	ApproachSiteCodes3301002  ApproachSiteCodes = "3301002"
	ApproachSiteCodes3302009  ApproachSiteCodes = "3302009"
	ApproachSiteCodes3315000  ApproachSiteCodes = "3315000"
	ApproachSiteCodes3332001  ApproachSiteCodes = "3332001"
	ApproachSiteCodes3336003  ApproachSiteCodes = "3336003"
	ApproachSiteCodes3341006  ApproachSiteCodes = "3341006"
	ApproachSiteCodes3350008  ApproachSiteCodes = "3350008"
	ApproachSiteCodes3362007  ApproachSiteCodes = "3362007"
	ApproachSiteCodes3366005  ApproachSiteCodes = "3366005"
	ApproachSiteCodes3370002  ApproachSiteCodes = "3370002"
	ApproachSiteCodes3374006  ApproachSiteCodes = "3374006"
	ApproachSiteCodes3377004  ApproachSiteCodes = "3377004"
	ApproachSiteCodes3382006  ApproachSiteCodes = "3382006"
	ApproachSiteCodes3383001  ApproachSiteCodes = "3383001"
	ApproachSiteCodes3394002  ApproachSiteCodes = "3394002"
	ApproachSiteCodes3395001  ApproachSiteCodes = "3395001"
	ApproachSiteCodes3396000  ApproachSiteCodes = "3396000"
	ApproachSiteCodes3400000  ApproachSiteCodes = "3400000"
	ApproachSiteCodes3409004  ApproachSiteCodes = "3409004"
	ApproachSiteCodes3417007  ApproachSiteCodes = "3417007"
	ApproachSiteCodes3438001  ApproachSiteCodes = "3438001"
	ApproachSiteCodes3444002  ApproachSiteCodes = "3444002"
	ApproachSiteCodes3447009  ApproachSiteCodes = "3447009"
	ApproachSiteCodes3460003  ApproachSiteCodes = "3460003"
	ApproachSiteCodes3462006  ApproachSiteCodes = "3462006"
	ApproachSiteCodes3471002  ApproachSiteCodes = "3471002"
	ApproachSiteCodes3478008  ApproachSiteCodes = "3478008"
	ApproachSiteCodes3481003  ApproachSiteCodes = "3481003"
	ApproachSiteCodes3488009  ApproachSiteCodes = "3488009"
	ApproachSiteCodes3490005  ApproachSiteCodes = "3490005"
	ApproachSiteCodes3524005  ApproachSiteCodes = "3524005"
	ApproachSiteCodes3538003  ApproachSiteCodes = "3538003"
	ApproachSiteCodes3541007  ApproachSiteCodes = "3541007"
	ApproachSiteCodes3553006  ApproachSiteCodes = "3553006"
	ApproachSiteCodes3556003  ApproachSiteCodes = "3556003"
	ApproachSiteCodes3563003  ApproachSiteCodes = "3563003"
	ApproachSiteCodes3572006  ApproachSiteCodes = "3572006"
	ApproachSiteCodes3578005  ApproachSiteCodes = "3578005"
	ApproachSiteCodes3582007  ApproachSiteCodes = "3582007"
	ApproachSiteCodes3584008  ApproachSiteCodes = "3584008"
	ApproachSiteCodes3594003  ApproachSiteCodes = "3594003"
	ApproachSiteCodes3608004  ApproachSiteCodes = "3608004"
	ApproachSiteCodes3609007  ApproachSiteCodes = "3609007"
	ApproachSiteCodes3646006  ApproachSiteCodes = "3646006"
	ApproachSiteCodes3663005  ApproachSiteCodes = "3663005"
	ApproachSiteCodes3665003  ApproachSiteCodes = "3665003"
	ApproachSiteCodes3670005  ApproachSiteCodes = "3670005"
	ApproachSiteCodes3711007  ApproachSiteCodes = "3711007"
	ApproachSiteCodes3743007  ApproachSiteCodes = "3743007"
	ApproachSiteCodes3761003  ApproachSiteCodes = "3761003"
	ApproachSiteCodes3766008  ApproachSiteCodes = "3766008"
	ApproachSiteCodes3785006  ApproachSiteCodes = "3785006"
	ApproachSiteCodes3788008  ApproachSiteCodes = "3788008"
	ApproachSiteCodes3789000  ApproachSiteCodes = "3789000"
	ApproachSiteCodes3810000  ApproachSiteCodes = "3810000"
	ApproachSiteCodes3838008  ApproachSiteCodes = "3838008"
	ApproachSiteCodes3860006  ApproachSiteCodes = "3860006"
	ApproachSiteCodes3865001  ApproachSiteCodes = "3865001"
	ApproachSiteCodes3867009  ApproachSiteCodes = "3867009"
	ApproachSiteCodes3876002  ApproachSiteCodes = "3876002"
	ApproachSiteCodes3877006  ApproachSiteCodes = "3877006"
	ApproachSiteCodes3910004  ApproachSiteCodes = "3910004"
	ApproachSiteCodes3916005  ApproachSiteCodes = "3916005"
	ApproachSiteCodes3924000  ApproachSiteCodes = "3924000"
	ApproachSiteCodes3931001  ApproachSiteCodes = "3931001"
	ApproachSiteCodes3935005  ApproachSiteCodes = "3935005"
	ApproachSiteCodes3937002  ApproachSiteCodes = "3937002"
	ApproachSiteCodes3954005  ApproachSiteCodes = "3954005"
	ApproachSiteCodes3956007  ApproachSiteCodes = "3956007"
	ApproachSiteCodes3959000  ApproachSiteCodes = "3959000"
	ApproachSiteCodes3960005  ApproachSiteCodes = "3960005"
	ApproachSiteCodes3964001  ApproachSiteCodes = "3964001"
	ApproachSiteCodes3966004  ApproachSiteCodes = "3966004"
	ApproachSiteCodes3977005  ApproachSiteCodes = "3977005"
	ApproachSiteCodes3984002  ApproachSiteCodes = "3984002"
	ApproachSiteCodes3989007  ApproachSiteCodes = "3989007"
	ApproachSiteCodes4015004  ApproachSiteCodes = "4015004"
	ApproachSiteCodes4019005  ApproachSiteCodes = "4019005"
	ApproachSiteCodes4029003  ApproachSiteCodes = "4029003"
	ApproachSiteCodes4061004  ApproachSiteCodes = "4061004"
	ApproachSiteCodes4066009  ApproachSiteCodes = "4066009"
	ApproachSiteCodes4072009  ApproachSiteCodes = "4072009"
	ApproachSiteCodes4081003  ApproachSiteCodes = "4081003"
	ApproachSiteCodes4093007  ApproachSiteCodes = "4093007"
	ApproachSiteCodes4111006  ApproachSiteCodes = "4111006"
	ApproachSiteCodes4117005  ApproachSiteCodes = "4117005"
	ApproachSiteCodes4121003  ApproachSiteCodes = "4121003"
	ApproachSiteCodes4146003  ApproachSiteCodes = "4146003"
	ApproachSiteCodes4148002  ApproachSiteCodes = "4148002"
	ApproachSiteCodes4150005  ApproachSiteCodes = "4150005"
	ApproachSiteCodes4158003  ApproachSiteCodes = "4158003"
	ApproachSiteCodes4159006  ApproachSiteCodes = "4159006"
	ApproachSiteCodes4180000  ApproachSiteCodes = "4180000"
	ApproachSiteCodes4193005  ApproachSiteCodes = "4193005"
	ApproachSiteCodes4205002  ApproachSiteCodes = "4205002"
	ApproachSiteCodes4212006  ApproachSiteCodes = "4212006"
	ApproachSiteCodes4215008  ApproachSiteCodes = "4215008"
	ApproachSiteCodes4247003  ApproachSiteCodes = "4247003"
	ApproachSiteCodes4258007  ApproachSiteCodes = "4258007"
	ApproachSiteCodes4276000  ApproachSiteCodes = "4276000"
	ApproachSiteCodes4281009  ApproachSiteCodes = "4281009"
	ApproachSiteCodes4295007  ApproachSiteCodes = "4295007"
	ApproachSiteCodes4303006  ApproachSiteCodes = "4303006"
	ApproachSiteCodes4312008  ApproachSiteCodes = "4312008"
	ApproachSiteCodes4317002  ApproachSiteCodes = "4317002"
	ApproachSiteCodes4328003  ApproachSiteCodes = "4328003"
	ApproachSiteCodes4335006  ApproachSiteCodes = "4335006"
	ApproachSiteCodes4352005  ApproachSiteCodes = "4352005"
	ApproachSiteCodes4358009  ApproachSiteCodes = "4358009"
	ApproachSiteCodes4360006  ApproachSiteCodes = "4360006"
	ApproachSiteCodes4369007  ApproachSiteCodes = "4369007"
	ApproachSiteCodes4371007  ApproachSiteCodes = "4371007"
	ApproachSiteCodes4375003  ApproachSiteCodes = "4375003"
	ApproachSiteCodes4377006  ApproachSiteCodes = "4377006"
	ApproachSiteCodes4394008  ApproachSiteCodes = "4394008"
	ApproachSiteCodes4402002  ApproachSiteCodes = "4402002"
	ApproachSiteCodes4419000  ApproachSiteCodes = "4419000"
	ApproachSiteCodes4421005  ApproachSiteCodes = "4421005"
	ApproachSiteCodes4430002  ApproachSiteCodes = "4430002"
	ApproachSiteCodes4432005  ApproachSiteCodes = "4432005"
	ApproachSiteCodes4442007  ApproachSiteCodes = "4442007"
	ApproachSiteCodes4486002  ApproachSiteCodes = "4486002"
	ApproachSiteCodes4524000  ApproachSiteCodes = "4524000"
	ApproachSiteCodes4527007  ApproachSiteCodes = "4527007"
	ApproachSiteCodes4537002  ApproachSiteCodes = "4537002"
	ApproachSiteCodes4548009  ApproachSiteCodes = "4548009"
	ApproachSiteCodes4549001  ApproachSiteCodes = "4549001"
	ApproachSiteCodes4566004  ApproachSiteCodes = "4566004"
	ApproachSiteCodes4573009  ApproachSiteCodes = "4573009"
	ApproachSiteCodes4578000  ApproachSiteCodes = "4578000"
	ApproachSiteCodes4588004  ApproachSiteCodes = "4588004"
	ApproachSiteCodes4596009  ApproachSiteCodes = "4596009"
	ApproachSiteCodes4603002  ApproachSiteCodes = "4603002"
	ApproachSiteCodes4606005  ApproachSiteCodes = "4606005"
	ApproachSiteCodes4621004  ApproachSiteCodes = "4621004"
	ApproachSiteCodes4624007  ApproachSiteCodes = "4624007"
	ApproachSiteCodes4647008  ApproachSiteCodes = "4647008"
	ApproachSiteCodes4651005  ApproachSiteCodes = "4651005"
	ApproachSiteCodes4658004  ApproachSiteCodes = "4658004"
	ApproachSiteCodes4677002  ApproachSiteCodes = "4677002"
	ApproachSiteCodes4703008  ApproachSiteCodes = "4703008"
	ApproachSiteCodes4717004  ApproachSiteCodes = "4717004"
	ApproachSiteCodes4718009  ApproachSiteCodes = "4718009"
	ApproachSiteCodes4743003  ApproachSiteCodes = "4743003"
	ApproachSiteCodes4755009  ApproachSiteCodes = "4755009"
	ApproachSiteCodes4759003  ApproachSiteCodes = "4759003"
	ApproachSiteCodes4766002  ApproachSiteCodes = "4766002"
	ApproachSiteCodes4768001  ApproachSiteCodes = "4768001"
	ApproachSiteCodes4774001  ApproachSiteCodes = "4774001"
	ApproachSiteCodes4775000  ApproachSiteCodes = "4775000"
	ApproachSiteCodes4799000  ApproachSiteCodes = "4799000"
	ApproachSiteCodes4810005  ApproachSiteCodes = "4810005"
	ApproachSiteCodes4812002  ApproachSiteCodes = "4812002"
	ApproachSiteCodes4828007  ApproachSiteCodes = "4828007"
	ApproachSiteCodes4840007  ApproachSiteCodes = "4840007"
	ApproachSiteCodes4843009  ApproachSiteCodes = "4843009"
	ApproachSiteCodes4861000  ApproachSiteCodes = "4861000"
	ApproachSiteCodes4866005  ApproachSiteCodes = "4866005"
	ApproachSiteCodes4870002  ApproachSiteCodes = "4870002"
	ApproachSiteCodes4871003  ApproachSiteCodes = "4871003"
	ApproachSiteCodes4881004  ApproachSiteCodes = "4881004"
	ApproachSiteCodes4888005  ApproachSiteCodes = "4888005"
	ApproachSiteCodes4897009  ApproachSiteCodes = "4897009"
	ApproachSiteCodes4905007  ApproachSiteCodes = "4905007"
	ApproachSiteCodes4906008  ApproachSiteCodes = "4906008"
	ApproachSiteCodes4924005  ApproachSiteCodes = "4924005"
	ApproachSiteCodes4942000  ApproachSiteCodes = "4942000"
	ApproachSiteCodes4954000  ApproachSiteCodes = "4954000"
	ApproachSiteCodes4956003  ApproachSiteCodes = "4956003"
	ApproachSiteCodes4958002  ApproachSiteCodes = "4958002"
	ApproachSiteCodes5001007  ApproachSiteCodes = "5001007"
	ApproachSiteCodes5023006  ApproachSiteCodes = "5023006"
	ApproachSiteCodes5026003  ApproachSiteCodes = "5026003"
	ApproachSiteCodes5046008  ApproachSiteCodes = "5046008"
	ApproachSiteCodes5068003  ApproachSiteCodes = "5068003"
	ApproachSiteCodes5069006  ApproachSiteCodes = "5069006"
	ApproachSiteCodes5076001  ApproachSiteCodes = "5076001"
	ApproachSiteCodes5115006  ApproachSiteCodes = "5115006"
	ApproachSiteCodes5122003  ApproachSiteCodes = "5122003"
	ApproachSiteCodes5128004  ApproachSiteCodes = "5128004"
	ApproachSiteCodes5192008  ApproachSiteCodes = "5192008"
	ApproachSiteCodes5194009  ApproachSiteCodes = "5194009"
	ApproachSiteCodes5195005  ApproachSiteCodes = "5195005"
	ApproachSiteCodes5204005  ApproachSiteCodes = "5204005"
	ApproachSiteCodes5213007  ApproachSiteCodes = "5213007"
	ApproachSiteCodes5225005  ApproachSiteCodes = "5225005"
	ApproachSiteCodes5228007  ApproachSiteCodes = "5228007"
	ApproachSiteCodes5229004  ApproachSiteCodes = "5229004"
	ApproachSiteCodes5261000  ApproachSiteCodes = "5261000"
	ApproachSiteCodes5272005  ApproachSiteCodes = "5272005"
	ApproachSiteCodes5279001  ApproachSiteCodes = "5279001"
	ApproachSiteCodes5296000  ApproachSiteCodes = "5296000"
	ApproachSiteCodes5324007  ApproachSiteCodes = "5324007"
	ApproachSiteCodes5329002  ApproachSiteCodes = "5329002"
	ApproachSiteCodes5336001  ApproachSiteCodes = "5336001"
	ApproachSiteCodes5347008  ApproachSiteCodes = "5347008"
	ApproachSiteCodes5362005  ApproachSiteCodes = "5362005"
	ApproachSiteCodes5366008  ApproachSiteCodes = "5366008"
	ApproachSiteCodes5379004  ApproachSiteCodes = "5379004"
	ApproachSiteCodes5382009  ApproachSiteCodes = "5382009"
	ApproachSiteCodes5394000  ApproachSiteCodes = "5394000"
	ApproachSiteCodes5398002  ApproachSiteCodes = "5398002"
	ApproachSiteCodes5403001  ApproachSiteCodes = "5403001"
	ApproachSiteCodes5421003  ApproachSiteCodes = "5421003"
	ApproachSiteCodes5427004  ApproachSiteCodes = "5427004"
	ApproachSiteCodes5458003  ApproachSiteCodes = "5458003"
	ApproachSiteCodes5459006  ApproachSiteCodes = "5459006"
	ApproachSiteCodes5491007  ApproachSiteCodes = "5491007"
	ApproachSiteCodes5493005  ApproachSiteCodes = "5493005"
	ApproachSiteCodes5498001  ApproachSiteCodes = "5498001"
	ApproachSiteCodes5520004  ApproachSiteCodes = "5520004"
	ApproachSiteCodes5538001  ApproachSiteCodes = "5538001"
	ApproachSiteCodes5544002  ApproachSiteCodes = "5544002"
	ApproachSiteCodes5560003  ApproachSiteCodes = "5560003"
	ApproachSiteCodes5564007  ApproachSiteCodes = "5564007"
	ApproachSiteCodes5574005  ApproachSiteCodes = "5574005"
	ApproachSiteCodes5580002  ApproachSiteCodes = "5580002"
	ApproachSiteCodes5597008  ApproachSiteCodes = "5597008"
	ApproachSiteCodes5611001  ApproachSiteCodes = "5611001"
	ApproachSiteCodes5625000  ApproachSiteCodes = "5625000"
	ApproachSiteCodes5627008  ApproachSiteCodes = "5627008"
	ApproachSiteCodes5633004  ApproachSiteCodes = "5633004"
	ApproachSiteCodes5643001  ApproachSiteCodes = "5643001"
	ApproachSiteCodes5644007  ApproachSiteCodes = "5644007"
	ApproachSiteCodes5653000  ApproachSiteCodes = "5653000"
	ApproachSiteCodes5665001  ApproachSiteCodes = "5665001"
	ApproachSiteCodes5668004  ApproachSiteCodes = "5668004"
	ApproachSiteCodes5682004  ApproachSiteCodes = "5682004"
	ApproachSiteCodes5696005  ApproachSiteCodes = "5696005"
	ApproachSiteCodes5697001  ApproachSiteCodes = "5697001"
	ApproachSiteCodes5709001  ApproachSiteCodes = "5709001"
	ApproachSiteCodes5713008  ApproachSiteCodes = "5713008"
	ApproachSiteCodes5717009  ApproachSiteCodes = "5717009"
	ApproachSiteCodes5718004  ApproachSiteCodes = "5718004"
	ApproachSiteCodes5727003  ApproachSiteCodes = "5727003"
	ApproachSiteCodes5742000  ApproachSiteCodes = "5742000"
	ApproachSiteCodes5751008  ApproachSiteCodes = "5751008"
	ApproachSiteCodes5769004  ApproachSiteCodes = "5769004"
	ApproachSiteCodes5780004  ApproachSiteCodes = "5780004"
	ApproachSiteCodes5798000  ApproachSiteCodes = "5798000"
	ApproachSiteCodes5802004  ApproachSiteCodes = "5802004"
	ApproachSiteCodes5814007  ApproachSiteCodes = "5814007"
	ApproachSiteCodes5815008  ApproachSiteCodes = "5815008"
	ApproachSiteCodes5816009  ApproachSiteCodes = "5816009"
	ApproachSiteCodes5825003  ApproachSiteCodes = "5825003"
	ApproachSiteCodes5828001  ApproachSiteCodes = "5828001"
	ApproachSiteCodes5847003  ApproachSiteCodes = "5847003"
	ApproachSiteCodes5854009  ApproachSiteCodes = "5854009"
	ApproachSiteCodes5868002  ApproachSiteCodes = "5868002"
	ApproachSiteCodes5872003  ApproachSiteCodes = "5872003"
	ApproachSiteCodes5881009  ApproachSiteCodes = "5881009"
	ApproachSiteCodes5882002  ApproachSiteCodes = "5882002"
	ApproachSiteCodes5889006  ApproachSiteCodes = "5889006"
	ApproachSiteCodes5890002  ApproachSiteCodes = "5890002"
	ApproachSiteCodes5893000  ApproachSiteCodes = "5893000"
	ApproachSiteCodes5898009  ApproachSiteCodes = "5898009"
	ApproachSiteCodes5923009  ApproachSiteCodes = "5923009"
	ApproachSiteCodes5926001  ApproachSiteCodes = "5926001"
	ApproachSiteCodes5928000  ApproachSiteCodes = "5928000"
	ApproachSiteCodes5942008  ApproachSiteCodes = "5942008"
	ApproachSiteCodes5943003  ApproachSiteCodes = "5943003"
	ApproachSiteCodes5944009  ApproachSiteCodes = "5944009"
	ApproachSiteCodes5948007  ApproachSiteCodes = "5948007"
	ApproachSiteCodes5951000  ApproachSiteCodes = "5951000"
	ApproachSiteCodes5953002  ApproachSiteCodes = "5953002"
	ApproachSiteCodes5976004  ApproachSiteCodes = "5976004"
	ApproachSiteCodes5979006  ApproachSiteCodes = "5979006"
	ApproachSiteCodes5996007  ApproachSiteCodes = "5996007"
	ApproachSiteCodes6001004  ApproachSiteCodes = "6001004"
	ApproachSiteCodes6004007  ApproachSiteCodes = "6004007"
	ApproachSiteCodes6006009  ApproachSiteCodes = "6006009"
	ApproachSiteCodes6009002  ApproachSiteCodes = "6009002"
	ApproachSiteCodes6014003  ApproachSiteCodes = "6014003"
	ApproachSiteCodes6023000  ApproachSiteCodes = "6023000"
	ApproachSiteCodes6032003  ApproachSiteCodes = "6032003"
	ApproachSiteCodes6046003  ApproachSiteCodes = "6046003"
	ApproachSiteCodes6050005  ApproachSiteCodes = "6050005"
	ApproachSiteCodes6059006  ApproachSiteCodes = "6059006"
	ApproachSiteCodes6062009  ApproachSiteCodes = "6062009"
	ApproachSiteCodes6073002  ApproachSiteCodes = "6073002"
	ApproachSiteCodes6074008  ApproachSiteCodes = "6074008"
	ApproachSiteCodes6076005  ApproachSiteCodes = "6076005"
	ApproachSiteCodes6104005  ApproachSiteCodes = "6104005"
	ApproachSiteCodes6105006  ApproachSiteCodes = "6105006"
	ApproachSiteCodes6110005  ApproachSiteCodes = "6110005"
	ApproachSiteCodes6216007  ApproachSiteCodes = "6216007"
	ApproachSiteCodes6217003  ApproachSiteCodes = "6217003"
	ApproachSiteCodes6229007  ApproachSiteCodes = "6229007"
	ApproachSiteCodes6253001  ApproachSiteCodes = "6253001"
	ApproachSiteCodes6268000  ApproachSiteCodes = "6268000"
	ApproachSiteCodes6269008  ApproachSiteCodes = "6269008"
	ApproachSiteCodes6279005  ApproachSiteCodes = "6279005"
	ApproachSiteCodes6317000  ApproachSiteCodes = "6317000"
	ApproachSiteCodes6325003  ApproachSiteCodes = "6325003"
	ApproachSiteCodes6326002  ApproachSiteCodes = "6326002"
	ApproachSiteCodes6335009  ApproachSiteCodes = "6335009"
	ApproachSiteCodes6359004  ApproachSiteCodes = "6359004"
	ApproachSiteCodes6371005  ApproachSiteCodes = "6371005"
	ApproachSiteCodes6375001  ApproachSiteCodes = "6375001"
	ApproachSiteCodes6392005  ApproachSiteCodes = "6392005"
	ApproachSiteCodes6404004  ApproachSiteCodes = "6404004"
	ApproachSiteCodes6413002  ApproachSiteCodes = "6413002"
	ApproachSiteCodes6417001  ApproachSiteCodes = "6417001"
	ApproachSiteCodes6423006  ApproachSiteCodes = "6423006"
	ApproachSiteCodes6424000  ApproachSiteCodes = "6424000"
	ApproachSiteCodes6445007  ApproachSiteCodes = "6445007"
	ApproachSiteCodes6448009  ApproachSiteCodes = "6448009"
	ApproachSiteCodes6450001  ApproachSiteCodes = "6450001"
	ApproachSiteCodes6472004  ApproachSiteCodes = "6472004"
	ApproachSiteCodes6504002  ApproachSiteCodes = "6504002"
	ApproachSiteCodes6511003  ApproachSiteCodes = "6511003"
	ApproachSiteCodes6530003  ApproachSiteCodes = "6530003"
	ApproachSiteCodes6533001  ApproachSiteCodes = "6533001"
	ApproachSiteCodes6538005  ApproachSiteCodes = "6538005"
	ApproachSiteCodes6541001  ApproachSiteCodes = "6541001"
	ApproachSiteCodes6544009  ApproachSiteCodes = "6544009"
	ApproachSiteCodes6550004  ApproachSiteCodes = "6550004"
	ApproachSiteCodes6551000  ApproachSiteCodes = "6551000"
	ApproachSiteCodes6553002  ApproachSiteCodes = "6553002"
	ApproachSiteCodes6564004  ApproachSiteCodes = "6564004"
	ApproachSiteCodes6566002  ApproachSiteCodes = "6566002"
	ApproachSiteCodes6572002  ApproachSiteCodes = "6572002"
	ApproachSiteCodes6598008  ApproachSiteCodes = "6598008"
	ApproachSiteCodes6606008  ApproachSiteCodes = "6606008"
	ApproachSiteCodes6608009  ApproachSiteCodes = "6608009"
	ApproachSiteCodes6620001  ApproachSiteCodes = "6620001"
	ApproachSiteCodes6623004  ApproachSiteCodes = "6623004"
	ApproachSiteCodes6633007  ApproachSiteCodes = "6633007"
	ApproachSiteCodes6643005  ApproachSiteCodes = "6643005"
	ApproachSiteCodes6646002  ApproachSiteCodes = "6646002"
	ApproachSiteCodes6649009  ApproachSiteCodes = "6649009"
	ApproachSiteCodes6651008  ApproachSiteCodes = "6651008"
	ApproachSiteCodes6684008  ApproachSiteCodes = "6684008"
	ApproachSiteCodes6685009  ApproachSiteCodes = "6685009"
	ApproachSiteCodes6711001  ApproachSiteCodes = "6711001"
	ApproachSiteCodes6720005  ApproachSiteCodes = "6720005"
	ApproachSiteCodes6731002  ApproachSiteCodes = "6731002"
	ApproachSiteCodes6739000  ApproachSiteCodes = "6739000"
	ApproachSiteCodes6742006  ApproachSiteCodes = "6742006"
	ApproachSiteCodes6750002  ApproachSiteCodes = "6750002"
	ApproachSiteCodes6757004  ApproachSiteCodes = "6757004"
	ApproachSiteCodes6787005  ApproachSiteCodes = "6787005"
	ApproachSiteCodes6789008  ApproachSiteCodes = "6789008"
	ApproachSiteCodes6799003  ApproachSiteCodes = "6799003"
	ApproachSiteCodes6805009  ApproachSiteCodes = "6805009"
	ApproachSiteCodes6820003  ApproachSiteCodes = "6820003"
	ApproachSiteCodes6828005  ApproachSiteCodes = "6828005"
	ApproachSiteCodes6829002  ApproachSiteCodes = "6829002"
	ApproachSiteCodes6834003  ApproachSiteCodes = "6834003"
	ApproachSiteCodes6841009  ApproachSiteCodes = "6841009"
	ApproachSiteCodes6844001  ApproachSiteCodes = "6844001"
	ApproachSiteCodes6850006  ApproachSiteCodes = "6850006"
	ApproachSiteCodes6864006  ApproachSiteCodes = "6864006"
	ApproachSiteCodes6866008  ApproachSiteCodes = "6866008"
	ApproachSiteCodes6871001  ApproachSiteCodes = "6871001"
	ApproachSiteCodes6894000  ApproachSiteCodes = "6894000"
	ApproachSiteCodes6902008  ApproachSiteCodes = "6902008"
	ApproachSiteCodes6905005  ApproachSiteCodes = "6905005"
	ApproachSiteCodes6912001  ApproachSiteCodes = "6912001"
	ApproachSiteCodes6914000  ApproachSiteCodes = "6914000"
	ApproachSiteCodes6921000  ApproachSiteCodes = "6921000"
	ApproachSiteCodes6930008  ApproachSiteCodes = "6930008"
	ApproachSiteCodes6944002  ApproachSiteCodes = "6944002"
	ApproachSiteCodes6969002  ApproachSiteCodes = "6969002"
	ApproachSiteCodes6975006  ApproachSiteCodes = "6975006"
	ApproachSiteCodes6981003  ApproachSiteCodes = "6981003"
	ApproachSiteCodes6987004  ApproachSiteCodes = "6987004"
	ApproachSiteCodes6989001  ApproachSiteCodes = "6989001"
	ApproachSiteCodes6991009  ApproachSiteCodes = "6991009"
	ApproachSiteCodes7035006  ApproachSiteCodes = "7035006"
	ApproachSiteCodes7050002  ApproachSiteCodes = "7050002"
	ApproachSiteCodes7067009  ApproachSiteCodes = "7067009"
	ApproachSiteCodes7076002  ApproachSiteCodes = "7076002"
	ApproachSiteCodes7083009  ApproachSiteCodes = "7083009"
	ApproachSiteCodes7091000  ApproachSiteCodes = "7091000"
	ApproachSiteCodes7099003  ApproachSiteCodes = "7099003"
	ApproachSiteCodes7117004  ApproachSiteCodes = "7117004"
	ApproachSiteCodes7148007  ApproachSiteCodes = "7148007"
	ApproachSiteCodes7149004  ApproachSiteCodes = "7149004"
	ApproachSiteCodes7154008  ApproachSiteCodes = "7154008"
	ApproachSiteCodes7160008  ApproachSiteCodes = "7160008"
	ApproachSiteCodes7167006  ApproachSiteCodes = "7167006"
	ApproachSiteCodes7173007  ApproachSiteCodes = "7173007"
	ApproachSiteCodes7188002  ApproachSiteCodes = "7188002"
	ApproachSiteCodes7192009  ApproachSiteCodes = "7192009"
	ApproachSiteCodes7227003  ApproachSiteCodes = "7227003"
	ApproachSiteCodes7234001  ApproachSiteCodes = "7234001"
	ApproachSiteCodes7242000  ApproachSiteCodes = "7242000"
	ApproachSiteCodes7295002  ApproachSiteCodes = "7295002"
	ApproachSiteCodes7296001  ApproachSiteCodes = "7296001"
	ApproachSiteCodes7311008  ApproachSiteCodes = "7311008"
	ApproachSiteCodes7344002  ApproachSiteCodes = "7344002"
	ApproachSiteCodes7345001  ApproachSiteCodes = "7345001"
	ApproachSiteCodes7362006  ApproachSiteCodes = "7362006"
	ApproachSiteCodes7376007  ApproachSiteCodes = "7376007"
	ApproachSiteCodes7378008  ApproachSiteCodes = "7378008"
	ApproachSiteCodes7384006  ApproachSiteCodes = "7384006"
	ApproachSiteCodes7404008  ApproachSiteCodes = "7404008"
	ApproachSiteCodes7435002  ApproachSiteCodes = "7435002"
	ApproachSiteCodes7471001  ApproachSiteCodes = "7471001"
	ApproachSiteCodes7477002  ApproachSiteCodes = "7477002"
	ApproachSiteCodes7480001  ApproachSiteCodes = "7480001"
	ApproachSiteCodes7494000  ApproachSiteCodes = "7494000"
	ApproachSiteCodes7498002  ApproachSiteCodes = "7498002"
	ApproachSiteCodes7507003  ApproachSiteCodes = "7507003"
	ApproachSiteCodes7524009  ApproachSiteCodes = "7524009"
	ApproachSiteCodes7532001  ApproachSiteCodes = "7532001"
	ApproachSiteCodes7554004  ApproachSiteCodes = "7554004"
	ApproachSiteCodes7566005  ApproachSiteCodes = "7566005"
	ApproachSiteCodes7569003  ApproachSiteCodes = "7569003"
	ApproachSiteCodes7591005  ApproachSiteCodes = "7591005"
	ApproachSiteCodes7597009  ApproachSiteCodes = "7597009"
	ApproachSiteCodes7605000  ApproachSiteCodes = "7605000"
	ApproachSiteCodes7610001  ApproachSiteCodes = "7610001"
	ApproachSiteCodes7629007  ApproachSiteCodes = "7629007"
	ApproachSiteCodes7651004  ApproachSiteCodes = "7651004"
	ApproachSiteCodes7652006  ApproachSiteCodes = "7652006"
	ApproachSiteCodes7657000  ApproachSiteCodes = "7657000"
	ApproachSiteCodes7658005  ApproachSiteCodes = "7658005"
	ApproachSiteCodes7697002  ApproachSiteCodes = "7697002"
	ApproachSiteCodes7712004  ApproachSiteCodes = "7712004"
	ApproachSiteCodes7726008  ApproachSiteCodes = "7726008"
	ApproachSiteCodes7736000  ApproachSiteCodes = "7736000"
	ApproachSiteCodes7742001  ApproachSiteCodes = "7742001"
	ApproachSiteCodes7748002  ApproachSiteCodes = "7748002"
	ApproachSiteCodes7755000  ApproachSiteCodes = "7755000"
	ApproachSiteCodes7756004  ApproachSiteCodes = "7756004"
	ApproachSiteCodes7764005  ApproachSiteCodes = "7764005"
	ApproachSiteCodes7769000  ApproachSiteCodes = "7769000"
	ApproachSiteCodes7783003  ApproachSiteCodes = "7783003"
	ApproachSiteCodes7820009  ApproachSiteCodes = "7820009"
	ApproachSiteCodes7829005  ApproachSiteCodes = "7829005"
	ApproachSiteCodes7832008  ApproachSiteCodes = "7832008"
	ApproachSiteCodes7835005  ApproachSiteCodes = "7835005"
	ApproachSiteCodes7840002  ApproachSiteCodes = "7840002"
	ApproachSiteCodes7841003  ApproachSiteCodes = "7841003"
	ApproachSiteCodes7844006  ApproachSiteCodes = "7844006"
	ApproachSiteCodes7851002  ApproachSiteCodes = "7851002"
	ApproachSiteCodes7854005  ApproachSiteCodes = "7854005"
	ApproachSiteCodes7872004  ApproachSiteCodes = "7872004"
	ApproachSiteCodes7874003  ApproachSiteCodes = "7874003"
	ApproachSiteCodes7880006  ApproachSiteCodes = "7880006"
	ApproachSiteCodes7884002  ApproachSiteCodes = "7884002"
	ApproachSiteCodes7885001  ApproachSiteCodes = "7885001"
	ApproachSiteCodes7892006  ApproachSiteCodes = "7892006"
	ApproachSiteCodes7896009  ApproachSiteCodes = "7896009"
	ApproachSiteCodes7911004  ApproachSiteCodes = "7911004"
	ApproachSiteCodes7925003  ApproachSiteCodes = "7925003"
	ApproachSiteCodes7936005  ApproachSiteCodes = "7936005"
	ApproachSiteCodes7944005  ApproachSiteCodes = "7944005"
	ApproachSiteCodes7954009  ApproachSiteCodes = "7954009"
	ApproachSiteCodes7967007  ApproachSiteCodes = "7967007"
	ApproachSiteCodes7986004  ApproachSiteCodes = "7986004"
	ApproachSiteCodes7991003  ApproachSiteCodes = "7991003"
	ApproachSiteCodes7999001  ApproachSiteCodes = "7999001"
	ApproachSiteCodes8001006  ApproachSiteCodes = "8001006"
	ApproachSiteCodes8012006  ApproachSiteCodes = "8012006"
	ApproachSiteCodes8017000  ApproachSiteCodes = "8017000"
	ApproachSiteCodes8024004  ApproachSiteCodes = "8024004"
	ApproachSiteCodes8039003  ApproachSiteCodes = "8039003"
	ApproachSiteCodes8040001  ApproachSiteCodes = "8040001"
	ApproachSiteCodes8045006  ApproachSiteCodes = "8045006"
	ApproachSiteCodes8057002  ApproachSiteCodes = "8057002"
	ApproachSiteCodes8059004  ApproachSiteCodes = "8059004"
	ApproachSiteCodes8067007  ApproachSiteCodes = "8067007"
	ApproachSiteCodes8068002  ApproachSiteCodes = "8068002"
	ApproachSiteCodes8079007  ApproachSiteCodes = "8079007"
	ApproachSiteCodes8091003  ApproachSiteCodes = "8091003"
	ApproachSiteCodes8100009  ApproachSiteCodes = "8100009"
	ApproachSiteCodes8111001  ApproachSiteCodes = "8111001"
	ApproachSiteCodes8112008  ApproachSiteCodes = "8112008"
	ApproachSiteCodes8119004  ApproachSiteCodes = "8119004"
	ApproachSiteCodes8128003  ApproachSiteCodes = "8128003"
	ApproachSiteCodes8133004  ApproachSiteCodes = "8133004"
	ApproachSiteCodes8157004  ApproachSiteCodes = "8157004"
	ApproachSiteCodes8158009  ApproachSiteCodes = "8158009"
	ApproachSiteCodes8159001  ApproachSiteCodes = "8159001"
	ApproachSiteCodes8160006  ApproachSiteCodes = "8160006"
	ApproachSiteCodes8161005  ApproachSiteCodes = "8161005"
	ApproachSiteCodes8165001  ApproachSiteCodes = "8165001"
	ApproachSiteCodes8205005  ApproachSiteCodes = "8205005"
	ApproachSiteCodes8225009  ApproachSiteCodes = "8225009"
	ApproachSiteCodes8242003  ApproachSiteCodes = "8242003"
	ApproachSiteCodes8251006  ApproachSiteCodes = "8251006"
	ApproachSiteCodes8264007  ApproachSiteCodes = "8264007"
	ApproachSiteCodes8265008  ApproachSiteCodes = "8265008"
	ApproachSiteCodes8266009  ApproachSiteCodes = "8266009"
	ApproachSiteCodes8289001  ApproachSiteCodes = "8289001"
	ApproachSiteCodes8292002  ApproachSiteCodes = "8292002"
	ApproachSiteCodes8314003  ApproachSiteCodes = "8314003"
	ApproachSiteCodes8334002  ApproachSiteCodes = "8334002"
	ApproachSiteCodes8356004  ApproachSiteCodes = "8356004"
	ApproachSiteCodes8369000  ApproachSiteCodes = "8369000"
	ApproachSiteCodes8373002  ApproachSiteCodes = "8373002"
	ApproachSiteCodes8387002  ApproachSiteCodes = "8387002"
	ApproachSiteCodes8389004  ApproachSiteCodes = "8389004"
	ApproachSiteCodes8412003  ApproachSiteCodes = "8412003"
	ApproachSiteCodes8415001  ApproachSiteCodes = "8415001"
	ApproachSiteCodes8454000  ApproachSiteCodes = "8454000"
	ApproachSiteCodes8464009  ApproachSiteCodes = "8464009"
	ApproachSiteCodes8482007  ApproachSiteCodes = "8482007"
	ApproachSiteCodes8483002  ApproachSiteCodes = "8483002"
	ApproachSiteCodes8496001  ApproachSiteCodes = "8496001"
	ApproachSiteCodes8523001  ApproachSiteCodes = "8523001"
	ApproachSiteCodes8546004  ApproachSiteCodes = "8546004"
	ApproachSiteCodes8556000  ApproachSiteCodes = "8556000"
	ApproachSiteCodes8559007  ApproachSiteCodes = "8559007"
	ApproachSiteCodes8560002  ApproachSiteCodes = "8560002"
	ApproachSiteCodes8580001  ApproachSiteCodes = "8580001"
	ApproachSiteCodes8595004  ApproachSiteCodes = "8595004"
	ApproachSiteCodes8598002  ApproachSiteCodes = "8598002"
	ApproachSiteCodes8600008  ApproachSiteCodes = "8600008"
	ApproachSiteCodes8603005  ApproachSiteCodes = "8603005"
	ApproachSiteCodes8604004  ApproachSiteCodes = "8604004"
	ApproachSiteCodes8608001  ApproachSiteCodes = "8608001"
	ApproachSiteCodes8617001  ApproachSiteCodes = "8617001"
	ApproachSiteCodes8623006  ApproachSiteCodes = "8623006"
	ApproachSiteCodes8629005  ApproachSiteCodes = "8629005"
	ApproachSiteCodes8640002  ApproachSiteCodes = "8640002"
	ApproachSiteCodes8668003  ApproachSiteCodes = "8668003"
	ApproachSiteCodes8671006  ApproachSiteCodes = "8671006"
	ApproachSiteCodes8677005  ApproachSiteCodes = "8677005"
	ApproachSiteCodes8688004  ApproachSiteCodes = "8688004"
	ApproachSiteCodes8695008  ApproachSiteCodes = "8695008"
	ApproachSiteCodes8710005  ApproachSiteCodes = "8710005"
	ApproachSiteCodes8711009  ApproachSiteCodes = "8711009"
	ApproachSiteCodes8714001  ApproachSiteCodes = "8714001"
	ApproachSiteCodes8752000  ApproachSiteCodes = "8752000"
	ApproachSiteCodes8775007  ApproachSiteCodes = "8775007"
	ApproachSiteCodes8784007  ApproachSiteCodes = "8784007"
	ApproachSiteCodes8810002  ApproachSiteCodes = "8810002"
	ApproachSiteCodes8814006  ApproachSiteCodes = "8814006"
	ApproachSiteCodes8815007  ApproachSiteCodes = "8815007"
	ApproachSiteCodes8820007  ApproachSiteCodes = "8820007"
	ApproachSiteCodes8821006  ApproachSiteCodes = "8821006"
	ApproachSiteCodes8827005  ApproachSiteCodes = "8827005"
	ApproachSiteCodes8839002  ApproachSiteCodes = "8839002"
	ApproachSiteCodes8845005  ApproachSiteCodes = "8845005"
	ApproachSiteCodes8850004  ApproachSiteCodes = "8850004"
	ApproachSiteCodes8854008  ApproachSiteCodes = "8854008"
	ApproachSiteCodes8862000  ApproachSiteCodes = "8862000"
	ApproachSiteCodes8873007  ApproachSiteCodes = "8873007"
	ApproachSiteCodes8887007  ApproachSiteCodes = "8887007"
	ApproachSiteCodes8892009  ApproachSiteCodes = "8892009"
	ApproachSiteCodes8894005  ApproachSiteCodes = "8894005"
	ApproachSiteCodes8897003  ApproachSiteCodes = "8897003"
	ApproachSiteCodes8907008  ApproachSiteCodes = "8907008"
	ApproachSiteCodes8910001  ApproachSiteCodes = "8910001"
	ApproachSiteCodes8911002  ApproachSiteCodes = "8911002"
	ApproachSiteCodes8928004  ApproachSiteCodes = "8928004"
	ApproachSiteCodes8931003  ApproachSiteCodes = "8931003"
	ApproachSiteCodes8935007  ApproachSiteCodes = "8935007"
	ApproachSiteCodes8942007  ApproachSiteCodes = "8942007"
	ApproachSiteCodes8965002  ApproachSiteCodes = "8965002"
	ApproachSiteCodes8966001  ApproachSiteCodes = "8966001"
	ApproachSiteCodes8983005  ApproachSiteCodes = "8983005"
	ApproachSiteCodes8988001  ApproachSiteCodes = "8988001"
	ApproachSiteCodes8993003  ApproachSiteCodes = "8993003"
	ApproachSiteCodes9000002  ApproachSiteCodes = "9000002"
	ApproachSiteCodes9003000  ApproachSiteCodes = "9003000"
	ApproachSiteCodes9018004  ApproachSiteCodes = "9018004"
	ApproachSiteCodes9040008  ApproachSiteCodes = "9040008"
	ApproachSiteCodes9055004  ApproachSiteCodes = "9055004"
	ApproachSiteCodes9073001  ApproachSiteCodes = "9073001"
	ApproachSiteCodes9081000  ApproachSiteCodes = "9081000"
	ApproachSiteCodes9086005  ApproachSiteCodes = "9086005"
	ApproachSiteCodes9089003  ApproachSiteCodes = "9089003"
	ApproachSiteCodes9108007  ApproachSiteCodes = "9108007"
	ApproachSiteCodes9127001  ApproachSiteCodes = "9127001"
	ApproachSiteCodes9156001  ApproachSiteCodes = "9156001"
	ApproachSiteCodes9185007  ApproachSiteCodes = "9185007"
	ApproachSiteCodes9186008  ApproachSiteCodes = "9186008"
	ApproachSiteCodes9188009  ApproachSiteCodes = "9188009"
	ApproachSiteCodes9208002  ApproachSiteCodes = "9208002"
	ApproachSiteCodes9212008  ApproachSiteCodes = "9212008"
	ApproachSiteCodes9229006  ApproachSiteCodes = "9229006"
	ApproachSiteCodes9231002  ApproachSiteCodes = "9231002"
	ApproachSiteCodes9240003  ApproachSiteCodes = "9240003"
	ApproachSiteCodes9242006  ApproachSiteCodes = "9242006"
	ApproachSiteCodes9258009  ApproachSiteCodes = "9258009"
	ApproachSiteCodes9261005  ApproachSiteCodes = "9261005"
	ApproachSiteCodes9262003  ApproachSiteCodes = "9262003"
	ApproachSiteCodes9284003  ApproachSiteCodes = "9284003"
	ApproachSiteCodes9290004  ApproachSiteCodes = "9290004"
	ApproachSiteCodes9305001  ApproachSiteCodes = "9305001"
	ApproachSiteCodes9317004  ApproachSiteCodes = "9317004"
	ApproachSiteCodes9320007  ApproachSiteCodes = "9320007"
	ApproachSiteCodes9321006  ApproachSiteCodes = "9321006"
	ApproachSiteCodes9325002  ApproachSiteCodes = "9325002"
	ApproachSiteCodes9332006  ApproachSiteCodes = "9332006"
	ApproachSiteCodes9348007  ApproachSiteCodes = "9348007"
	ApproachSiteCodes9379006  ApproachSiteCodes = "9379006"
	ApproachSiteCodes9380009  ApproachSiteCodes = "9380009"
	ApproachSiteCodes9384000  ApproachSiteCodes = "9384000"
	ApproachSiteCodes9390001  ApproachSiteCodes = "9390001"
	ApproachSiteCodes9432007  ApproachSiteCodes = "9432007"
	ApproachSiteCodes9438006  ApproachSiteCodes = "9438006"
	ApproachSiteCodes9454009  ApproachSiteCodes = "9454009"
	ApproachSiteCodes9455005  ApproachSiteCodes = "9455005"
	ApproachSiteCodes9475001  ApproachSiteCodes = "9475001"
	ApproachSiteCodes9481009  ApproachSiteCodes = "9481009"
	ApproachSiteCodes9490002  ApproachSiteCodes = "9490002"
	ApproachSiteCodes9498009  ApproachSiteCodes = "9498009"
	ApproachSiteCodes9502002  ApproachSiteCodes = "9502002"
	ApproachSiteCodes9512009  ApproachSiteCodes = "9512009"
	ApproachSiteCodes9535007  ApproachSiteCodes = "9535007"
	ApproachSiteCodes9558005  ApproachSiteCodes = "9558005"
	ApproachSiteCodes9566001  ApproachSiteCodes = "9566001"
	ApproachSiteCodes9568000  ApproachSiteCodes = "9568000"
	ApproachSiteCodes9596006  ApproachSiteCodes = "9596006"
	ApproachSiteCodes9609000  ApproachSiteCodes = "9609000"
	ApproachSiteCodes9625005  ApproachSiteCodes = "9625005"
	ApproachSiteCodes9642004  ApproachSiteCodes = "9642004"
	ApproachSiteCodes9646001  ApproachSiteCodes = "9646001"
	ApproachSiteCodes9654004  ApproachSiteCodes = "9654004"
	ApproachSiteCodes9659009  ApproachSiteCodes = "9659009"
	ApproachSiteCodes9662007  ApproachSiteCodes = "9662007"
	ApproachSiteCodes9668006  ApproachSiteCodes = "9668006"
	ApproachSiteCodes9677004  ApproachSiteCodes = "9677004"
	ApproachSiteCodes9683001  ApproachSiteCodes = "9683001"
	ApproachSiteCodes9684007  ApproachSiteCodes = "9684007"
	ApproachSiteCodes9708001  ApproachSiteCodes = "9708001"
	ApproachSiteCodes9736006  ApproachSiteCodes = "9736006"
	ApproachSiteCodes9743000  ApproachSiteCodes = "9743000"
	ApproachSiteCodes9758008  ApproachSiteCodes = "9758008"
	ApproachSiteCodes9770007  ApproachSiteCodes = "9770007"
	ApproachSiteCodes9775002  ApproachSiteCodes = "9775002"
	ApproachSiteCodes9779008  ApproachSiteCodes = "9779008"
	ApproachSiteCodes9783008  ApproachSiteCodes = "9783008"
	ApproachSiteCodes9791004  ApproachSiteCodes = "9791004"
	ApproachSiteCodes9796009  ApproachSiteCodes = "9796009"
	ApproachSiteCodes9813009  ApproachSiteCodes = "9813009"
	ApproachSiteCodes9825007  ApproachSiteCodes = "9825007"
	ApproachSiteCodes9837009  ApproachSiteCodes = "9837009"
	ApproachSiteCodes9840009  ApproachSiteCodes = "9840009"
	ApproachSiteCodes9841008  ApproachSiteCodes = "9841008"
	ApproachSiteCodes9846003  ApproachSiteCodes = "9846003"
	ApproachSiteCodes9847007  ApproachSiteCodes = "9847007"
	ApproachSiteCodes9849005  ApproachSiteCodes = "9849005"
	ApproachSiteCodes9870004  ApproachSiteCodes = "9870004"
	ApproachSiteCodes9875009  ApproachSiteCodes = "9875009"
	ApproachSiteCodes9878006  ApproachSiteCodes = "9878006"
	ApproachSiteCodes9880000  ApproachSiteCodes = "9880000"
	ApproachSiteCodes9881001  ApproachSiteCodes = "9881001"
	ApproachSiteCodes9891007  ApproachSiteCodes = "9891007"
	ApproachSiteCodes9898001  ApproachSiteCodes = "9898001"
	ApproachSiteCodes9951005  ApproachSiteCodes = "9951005"
	ApproachSiteCodes9968009  ApproachSiteCodes = "9968009"
	ApproachSiteCodes9970000  ApproachSiteCodes = "9970000"
	ApproachSiteCodes9976006  ApproachSiteCodes = "9976006"
	ApproachSiteCodes9994000  ApproachSiteCodes = "9994000"
	ApproachSiteCodes9999005  ApproachSiteCodes = "9999005"
	ApproachSiteCodes10013000 ApproachSiteCodes = "10013000"
	ApproachSiteCodes10024003 ApproachSiteCodes = "10024003"
	ApproachSiteCodes10025002 ApproachSiteCodes = "10025002"
	ApproachSiteCodes10026001 ApproachSiteCodes = "10026001"
	ApproachSiteCodes10036009 ApproachSiteCodes = "10036009"
	ApproachSiteCodes10042008 ApproachSiteCodes = "10042008"
	ApproachSiteCodes10047002 ApproachSiteCodes = "10047002"
	ApproachSiteCodes10052007 ApproachSiteCodes = "10052007"
	ApproachSiteCodes10056005 ApproachSiteCodes = "10056005"
	ApproachSiteCodes10062000 ApproachSiteCodes = "10062000"
	ApproachSiteCodes10119003 ApproachSiteCodes = "10119003"
	ApproachSiteCodes10124000 ApproachSiteCodes = "10124000"
	ApproachSiteCodes10134009 ApproachSiteCodes = "10134009"
	ApproachSiteCodes10141003 ApproachSiteCodes = "10141003"
	ApproachSiteCodes10145007 ApproachSiteCodes = "10145007"
	ApproachSiteCodes10148009 ApproachSiteCodes = "10148009"
	ApproachSiteCodes10149001 ApproachSiteCodes = "10149001"
	ApproachSiteCodes10151002 ApproachSiteCodes = "10151002"
	ApproachSiteCodes10167008 ApproachSiteCodes = "10167008"
	ApproachSiteCodes10176001 ApproachSiteCodes = "10176001"
	ApproachSiteCodes10200004 ApproachSiteCodes = "10200004"
	ApproachSiteCodes10208006 ApproachSiteCodes = "10208006"
	ApproachSiteCodes10209003 ApproachSiteCodes = "10209003"
	ApproachSiteCodes10245000 ApproachSiteCodes = "10245000"
	ApproachSiteCodes10271001 ApproachSiteCodes = "10271001"
	ApproachSiteCodes10293006 ApproachSiteCodes = "10293006"
	ApproachSiteCodes10296003 ApproachSiteCodes = "10296003"
	ApproachSiteCodes10299005 ApproachSiteCodes = "10299005"
	ApproachSiteCodes10328008 ApproachSiteCodes = "10328008"
	ApproachSiteCodes10339006 ApproachSiteCodes = "10339006"
	ApproachSiteCodes10410005 ApproachSiteCodes = "10410005"
	ApproachSiteCodes10415000 ApproachSiteCodes = "10415000"
	ApproachSiteCodes10417008 ApproachSiteCodes = "10417008"
	ApproachSiteCodes10418003 ApproachSiteCodes = "10418003"
)

func AllApproachSiteCodes() []ApproachSiteCodes {
	return []ApproachSiteCodes{
		ApproachSiteCodes91723000,
		ApproachSiteCodes106004,
		ApproachSiteCodes107008,
		ApproachSiteCodes108003,
		ApproachSiteCodes110001,
		ApproachSiteCodes111002,
		ApproachSiteCodes116007,
		ApproachSiteCodes124002,
		ApproachSiteCodes149003,
		ApproachSiteCodes155008,
		ApproachSiteCodes167005,
		ApproachSiteCodes202009,
		ApproachSiteCodes205006,
		ApproachSiteCodes206007,
		ApproachSiteCodes221001,
		ApproachSiteCodes227002,
		ApproachSiteCodes233006,
		ApproachSiteCodes235004,
		ApproachSiteCodes246001,
		ApproachSiteCodes247005,
		ApproachSiteCodes251007,
		ApproachSiteCodes256002,
		ApproachSiteCodes263002,
		ApproachSiteCodes266005,
		ApproachSiteCodes272005,
		ApproachSiteCodes273000,
		ApproachSiteCodes283001,
		ApproachSiteCodes284007,
		ApproachSiteCodes289002,
		ApproachSiteCodes301000,
		ApproachSiteCodes311007,
		ApproachSiteCodes315003,
		ApproachSiteCodes318001,
		ApproachSiteCodes344001,
		ApproachSiteCodes345000,
		ApproachSiteCodes356000,
		ApproachSiteCodes393006,
		ApproachSiteCodes402006,
		ApproachSiteCodes405008,
		ApproachSiteCodes414003,
		ApproachSiteCodes422005,
		ApproachSiteCodes446003,
		ApproachSiteCodes457008,
		ApproachSiteCodes461002,
		ApproachSiteCodes464005,
		ApproachSiteCodes465006,
		ApproachSiteCodes471000,
		ApproachSiteCodes480000,
		ApproachSiteCodes485005,
		ApproachSiteCodes528006,
		ApproachSiteCodes552004,
		ApproachSiteCodes565008,
		ApproachSiteCodes582005,
		ApproachSiteCodes587004,
		ApproachSiteCodes589001,
		ApproachSiteCodes595000,
		ApproachSiteCodes608002,
		ApproachSiteCodes621009,
		ApproachSiteCodes635006,
		ApproachSiteCodes650002,
		ApproachSiteCodes660006,
		ApproachSiteCodes661005,
		ApproachSiteCodes667009,
		ApproachSiteCodes688000,
		ApproachSiteCodes691000,
		ApproachSiteCodes692007,
		ApproachSiteCodes723004,
		ApproachSiteCodes774007,
		ApproachSiteCodes790007,
		ApproachSiteCodes798000,
		ApproachSiteCodes808000,
		ApproachSiteCodes809008,
		ApproachSiteCodes823005,
		ApproachSiteCodes830004,
		ApproachSiteCodes836005,
		ApproachSiteCodes885000,
		ApproachSiteCodes895007,
		ApproachSiteCodes917004,
		ApproachSiteCodes921006,
		ApproachSiteCodes947002,
		ApproachSiteCodes955009,
		ApproachSiteCodes976004,
		ApproachSiteCodes996007,
		ApproachSiteCodes1005009,
		ApproachSiteCodes1012000,
		ApproachSiteCodes1015003,
		ApproachSiteCodes1028005,
		ApproachSiteCodes1030007,
		ApproachSiteCodes1063000,
		ApproachSiteCodes1075005,
		ApproachSiteCodes1076006,
		ApproachSiteCodes1086007,
		ApproachSiteCodes1087003,
		ApproachSiteCodes1092001,
		ApproachSiteCodes1099005,
		ApproachSiteCodes1101003,
		ApproachSiteCodes1106008,
		ApproachSiteCodes1110006,
		ApproachSiteCodes1122009,
		ApproachSiteCodes1136004,
		ApproachSiteCodes1159005,
		ApproachSiteCodes1172006,
		ApproachSiteCodes1173001,
		ApproachSiteCodes1174007,
		ApproachSiteCodes1193009,
		ApproachSiteCodes1216008,
		ApproachSiteCodes1231004,
		ApproachSiteCodes1236009,
		ApproachSiteCodes1243003,
		ApproachSiteCodes1246006,
		ApproachSiteCodes1263005,
		ApproachSiteCodes1277008,
		ApproachSiteCodes1307006,
		ApproachSiteCodes1311000,
		ApproachSiteCodes1350001,
		ApproachSiteCodes1353004,
		ApproachSiteCodes1403006,
		ApproachSiteCodes1425000,
		ApproachSiteCodes1439000,
		ApproachSiteCodes1441004,
		ApproachSiteCodes1456008,
		ApproachSiteCodes1467009,
		ApproachSiteCodes1484003,
		ApproachSiteCodes1490004,
		ApproachSiteCodes1502004,
		ApproachSiteCodes1511004,
		ApproachSiteCodes1516009,
		ApproachSiteCodes1527006,
		ApproachSiteCodes1537001,
		ApproachSiteCodes1541002,
		ApproachSiteCodes1562001,
		ApproachSiteCodes1580005,
		ApproachSiteCodes1581009,
		ApproachSiteCodes1584001,
		ApproachSiteCodes1600003,
		ApproachSiteCodes1605008,
		ApproachSiteCodes1610007,
		ApproachSiteCodes1611006,
		ApproachSiteCodes1617005,
		ApproachSiteCodes1620002,
		ApproachSiteCodes1626008,
		ApproachSiteCodes1627004,
		ApproachSiteCodes1630006,
		ApproachSiteCodes1631005,
		ApproachSiteCodes1650005,
		ApproachSiteCodes1655000,
		ApproachSiteCodes1659006,
		ApproachSiteCodes1684009,
		ApproachSiteCodes1706004,
		ApproachSiteCodes1707008,
		ApproachSiteCodes1711002,
		ApproachSiteCodes1716007,
		ApproachSiteCodes1721005,
		ApproachSiteCodes1729007,
		ApproachSiteCodes1732005,
		ApproachSiteCodes1765002,
		ApproachSiteCodes1780008,
		ApproachSiteCodes1781007,
		ApproachSiteCodes1797002,
		ApproachSiteCodes1818002,
		ApproachSiteCodes1825009,
		ApproachSiteCodes1832000,
		ApproachSiteCodes1840006,
		ApproachSiteCodes1849007,
		ApproachSiteCodes1853009,
		ApproachSiteCodes1874005,
		ApproachSiteCodes1895000,
		ApproachSiteCodes1902009,
		ApproachSiteCodes1910005,
		ApproachSiteCodes1918003,
		ApproachSiteCodes1927002,
		ApproachSiteCodes1992003,
		ApproachSiteCodes1997009,
		ApproachSiteCodes2010005,
		ApproachSiteCodes2020000,
		ApproachSiteCodes2031008,
		ApproachSiteCodes2033006,
		ApproachSiteCodes2044003,
		ApproachSiteCodes2048000,
		ApproachSiteCodes2049008,
		ApproachSiteCodes2059009,
		ApproachSiteCodes2071003,
		ApproachSiteCodes2076008,
		ApproachSiteCodes2083001,
		ApproachSiteCodes2095001,
		ApproachSiteCodes2123001,
		ApproachSiteCodes2150006,
		ApproachSiteCodes2156000,
		ApproachSiteCodes2160002,
		ApproachSiteCodes2175005,
		ApproachSiteCodes2182009,
		ApproachSiteCodes2192001,
		ApproachSiteCodes2205003,
		ApproachSiteCodes2209009,
		ApproachSiteCodes2236006,
		ApproachSiteCodes2246008,
		ApproachSiteCodes2255006,
		ApproachSiteCodes2292006,
		ApproachSiteCodes2302002,
		ApproachSiteCodes2305000,
		ApproachSiteCodes2306004,
		ApproachSiteCodes2327009,
		ApproachSiteCodes2330002,
		ApproachSiteCodes2332005,
		ApproachSiteCodes2334006,
		ApproachSiteCodes2349003,
		ApproachSiteCodes2372001,
		ApproachSiteCodes2383005,
		ApproachSiteCodes2389009,
		ApproachSiteCodes2395005,
		ApproachSiteCodes2397002,
		ApproachSiteCodes2400006,
		ApproachSiteCodes2402003,
		ApproachSiteCodes2421006,
		ApproachSiteCodes2433001,
		ApproachSiteCodes2436009,
		ApproachSiteCodes2453002,
		ApproachSiteCodes2454008,
		ApproachSiteCodes2484000,
		ApproachSiteCodes2489005,
		ApproachSiteCodes2499000,
		ApproachSiteCodes2502001,
		ApproachSiteCodes2504000,
		ApproachSiteCodes2510000,
		ApproachSiteCodes2539000,
		ApproachSiteCodes2543001,
		ApproachSiteCodes2550002,
		ApproachSiteCodes2577006,
		ApproachSiteCodes2579009,
		ApproachSiteCodes2592007,
		ApproachSiteCodes2600000,
		ApproachSiteCodes2620004,
		ApproachSiteCodes2639009,
		ApproachSiteCodes2653009,
		ApproachSiteCodes2666009,
		ApproachSiteCodes2672009,
		ApproachSiteCodes2675006,
		ApproachSiteCodes2681003,
		ApproachSiteCodes2682005,
		ApproachSiteCodes2686008,
		ApproachSiteCodes2687004,
		ApproachSiteCodes2695000,
		ApproachSiteCodes2703009,
		ApproachSiteCodes2712006,
		ApproachSiteCodes2718005,
		ApproachSiteCodes2726002,
		ApproachSiteCodes2730004,
		ApproachSiteCodes2739003,
		ApproachSiteCodes2741002,
		ApproachSiteCodes2746007,
		ApproachSiteCodes2748008,
		ApproachSiteCodes2759004,
		ApproachSiteCodes2771005,
		ApproachSiteCodes2789006,
		ApproachSiteCodes2792005,
		ApproachSiteCodes2803000,
		ApproachSiteCodes2810006,
		ApproachSiteCodes2812003,
		ApproachSiteCodes2824005,
		ApproachSiteCodes2826007,
		ApproachSiteCodes2830005,
		ApproachSiteCodes2839006,
		ApproachSiteCodes2841007,
		ApproachSiteCodes2845003,
		ApproachSiteCodes2848001,
		ApproachSiteCodes2855004,
		ApproachSiteCodes2861001,
		ApproachSiteCodes2894003,
		ApproachSiteCodes2905008,
		ApproachSiteCodes2909002,
		ApproachSiteCodes2920002,
		ApproachSiteCodes2922005,
		ApproachSiteCodes2923000,
		ApproachSiteCodes2969000,
		ApproachSiteCodes2979003,
		ApproachSiteCodes2986006,
		ApproachSiteCodes2998001,
		ApproachSiteCodes3003007,
		ApproachSiteCodes3008003,
		ApproachSiteCodes3028004,
		ApproachSiteCodes3039001,
		ApproachSiteCodes3042007,
		ApproachSiteCodes3054007,
		ApproachSiteCodes3055008,
		ApproachSiteCodes3056009,
		ApproachSiteCodes3057000,
		ApproachSiteCodes3058005,
		ApproachSiteCodes3062004,
		ApproachSiteCodes3068000,
		ApproachSiteCodes3081007,
		ApproachSiteCodes3093003,
		ApproachSiteCodes3100007,
		ApproachSiteCodes3113001,
		ApproachSiteCodes3117000,
		ApproachSiteCodes3118005,
		ApproachSiteCodes3120008,
		ApproachSiteCodes3134008,
		ApproachSiteCodes3138006,
		ApproachSiteCodes3153003,
		ApproachSiteCodes3156006,
		ApproachSiteCodes3159004,
		ApproachSiteCodes3169005,
		ApproachSiteCodes3178004,
		ApproachSiteCodes3194006,
		ApproachSiteCodes3198009,
		ApproachSiteCodes3215002,
		ApproachSiteCodes3222005,
		ApproachSiteCodes3227004,
		ApproachSiteCodes3236000,
		ApproachSiteCodes3243006,
		ApproachSiteCodes3255000,
		ApproachSiteCodes3262009,
		ApproachSiteCodes3279003,
		ApproachSiteCodes3295003,
		ApproachSiteCodes3301002,
		ApproachSiteCodes3302009,
		ApproachSiteCodes3315000,
		ApproachSiteCodes3332001,
		ApproachSiteCodes3336003,
		ApproachSiteCodes3341006,
		ApproachSiteCodes3350008,
		ApproachSiteCodes3362007,
		ApproachSiteCodes3366005,
		ApproachSiteCodes3370002,
		ApproachSiteCodes3374006,
		ApproachSiteCodes3377004,
		ApproachSiteCodes3382006,
		ApproachSiteCodes3383001,
		ApproachSiteCodes3394002,
		ApproachSiteCodes3395001,
		ApproachSiteCodes3396000,
		ApproachSiteCodes3400000,
		ApproachSiteCodes3409004,
		ApproachSiteCodes3417007,
		ApproachSiteCodes3438001,
		ApproachSiteCodes3444002,
		ApproachSiteCodes3447009,
		ApproachSiteCodes3460003,
		ApproachSiteCodes3462006,
		ApproachSiteCodes3471002,
		ApproachSiteCodes3478008,
		ApproachSiteCodes3481003,
		ApproachSiteCodes3488009,
		ApproachSiteCodes3490005,
		ApproachSiteCodes3524005,
		ApproachSiteCodes3538003,
		ApproachSiteCodes3541007,
		ApproachSiteCodes3553006,
		ApproachSiteCodes3556003,
		ApproachSiteCodes3563003,
		ApproachSiteCodes3572006,
		ApproachSiteCodes3578005,
		ApproachSiteCodes3582007,
		ApproachSiteCodes3584008,
		ApproachSiteCodes3594003,
		ApproachSiteCodes3608004,
		ApproachSiteCodes3609007,
		ApproachSiteCodes3646006,
		ApproachSiteCodes3663005,
		ApproachSiteCodes3665003,
		ApproachSiteCodes3670005,
		ApproachSiteCodes3711007,
		ApproachSiteCodes3743007,
		ApproachSiteCodes3761003,
		ApproachSiteCodes3766008,
		ApproachSiteCodes3785006,
		ApproachSiteCodes3788008,
		ApproachSiteCodes3789000,
		ApproachSiteCodes3810000,
		ApproachSiteCodes3838008,
		ApproachSiteCodes3860006,
		ApproachSiteCodes3865001,
		ApproachSiteCodes3867009,
		ApproachSiteCodes3876002,
		ApproachSiteCodes3877006,
		ApproachSiteCodes3910004,
		ApproachSiteCodes3916005,
		ApproachSiteCodes3924000,
		ApproachSiteCodes3931001,
		ApproachSiteCodes3935005,
		ApproachSiteCodes3937002,
		ApproachSiteCodes3954005,
		ApproachSiteCodes3956007,
		ApproachSiteCodes3959000,
		ApproachSiteCodes3960005,
		ApproachSiteCodes3964001,
		ApproachSiteCodes3966004,
		ApproachSiteCodes3977005,
		ApproachSiteCodes3984002,
		ApproachSiteCodes3989007,
		ApproachSiteCodes4015004,
		ApproachSiteCodes4019005,
		ApproachSiteCodes4029003,
		ApproachSiteCodes4061004,
		ApproachSiteCodes4066009,
		ApproachSiteCodes4072009,
		ApproachSiteCodes4081003,
		ApproachSiteCodes4093007,
		ApproachSiteCodes4111006,
		ApproachSiteCodes4117005,
		ApproachSiteCodes4121003,
		ApproachSiteCodes4146003,
		ApproachSiteCodes4148002,
		ApproachSiteCodes4150005,
		ApproachSiteCodes4158003,
		ApproachSiteCodes4159006,
		ApproachSiteCodes4180000,
		ApproachSiteCodes4193005,
		ApproachSiteCodes4205002,
		ApproachSiteCodes4212006,
		ApproachSiteCodes4215008,
		ApproachSiteCodes4247003,
		ApproachSiteCodes4258007,
		ApproachSiteCodes4276000,
		ApproachSiteCodes4281009,
		ApproachSiteCodes4295007,
		ApproachSiteCodes4303006,
		ApproachSiteCodes4312008,
		ApproachSiteCodes4317002,
		ApproachSiteCodes4328003,
		ApproachSiteCodes4335006,
		ApproachSiteCodes4352005,
		ApproachSiteCodes4358009,
		ApproachSiteCodes4360006,
		ApproachSiteCodes4369007,
		ApproachSiteCodes4371007,
		ApproachSiteCodes4375003,
		ApproachSiteCodes4377006,
		ApproachSiteCodes4394008,
		ApproachSiteCodes4402002,
		ApproachSiteCodes4419000,
		ApproachSiteCodes4421005,
		ApproachSiteCodes4430002,
		ApproachSiteCodes4432005,
		ApproachSiteCodes4442007,
		ApproachSiteCodes4486002,
		ApproachSiteCodes4524000,
		ApproachSiteCodes4527007,
		ApproachSiteCodes4537002,
		ApproachSiteCodes4548009,
		ApproachSiteCodes4549001,
		ApproachSiteCodes4566004,
		ApproachSiteCodes4573009,
		ApproachSiteCodes4578000,
		ApproachSiteCodes4588004,
		ApproachSiteCodes4596009,
		ApproachSiteCodes4603002,
		ApproachSiteCodes4606005,
		ApproachSiteCodes4621004,
		ApproachSiteCodes4624007,
		ApproachSiteCodes4647008,
		ApproachSiteCodes4651005,
		ApproachSiteCodes4658004,
		ApproachSiteCodes4677002,
		ApproachSiteCodes4703008,
		ApproachSiteCodes4717004,
		ApproachSiteCodes4718009,
		ApproachSiteCodes4743003,
		ApproachSiteCodes4755009,
		ApproachSiteCodes4759003,
		ApproachSiteCodes4766002,
		ApproachSiteCodes4768001,
		ApproachSiteCodes4774001,
		ApproachSiteCodes4775000,
		ApproachSiteCodes4799000,
		ApproachSiteCodes4810005,
		ApproachSiteCodes4812002,
		ApproachSiteCodes4828007,
		ApproachSiteCodes4840007,
		ApproachSiteCodes4843009,
		ApproachSiteCodes4861000,
		ApproachSiteCodes4866005,
		ApproachSiteCodes4870002,
		ApproachSiteCodes4871003,
		ApproachSiteCodes4881004,
		ApproachSiteCodes4888005,
		ApproachSiteCodes4897009,
		ApproachSiteCodes4905007,
		ApproachSiteCodes4906008,
		ApproachSiteCodes4924005,
		ApproachSiteCodes4942000,
		ApproachSiteCodes4954000,
		ApproachSiteCodes4956003,
		ApproachSiteCodes4958002,
		ApproachSiteCodes5001007,
		ApproachSiteCodes5023006,
		ApproachSiteCodes5026003,
		ApproachSiteCodes5046008,
		ApproachSiteCodes5068003,
		ApproachSiteCodes5069006,
		ApproachSiteCodes5076001,
		ApproachSiteCodes5115006,
		ApproachSiteCodes5122003,
		ApproachSiteCodes5128004,
		ApproachSiteCodes5192008,
		ApproachSiteCodes5194009,
		ApproachSiteCodes5195005,
		ApproachSiteCodes5204005,
		ApproachSiteCodes5213007,
		ApproachSiteCodes5225005,
		ApproachSiteCodes5228007,
		ApproachSiteCodes5229004,
		ApproachSiteCodes5261000,
		ApproachSiteCodes5272005,
		ApproachSiteCodes5279001,
		ApproachSiteCodes5296000,
		ApproachSiteCodes5324007,
		ApproachSiteCodes5329002,
		ApproachSiteCodes5336001,
		ApproachSiteCodes5347008,
		ApproachSiteCodes5362005,
		ApproachSiteCodes5366008,
		ApproachSiteCodes5379004,
		ApproachSiteCodes5382009,
		ApproachSiteCodes5394000,
		ApproachSiteCodes5398002,
		ApproachSiteCodes5403001,
		ApproachSiteCodes5421003,
		ApproachSiteCodes5427004,
		ApproachSiteCodes5458003,
		ApproachSiteCodes5459006,
		ApproachSiteCodes5491007,
		ApproachSiteCodes5493005,
		ApproachSiteCodes5498001,
		ApproachSiteCodes5520004,
		ApproachSiteCodes5538001,
		ApproachSiteCodes5544002,
		ApproachSiteCodes5560003,
		ApproachSiteCodes5564007,
		ApproachSiteCodes5574005,
		ApproachSiteCodes5580002,
		ApproachSiteCodes5597008,
		ApproachSiteCodes5611001,
		ApproachSiteCodes5625000,
		ApproachSiteCodes5627008,
		ApproachSiteCodes5633004,
		ApproachSiteCodes5643001,
		ApproachSiteCodes5644007,
		ApproachSiteCodes5653000,
		ApproachSiteCodes5665001,
		ApproachSiteCodes5668004,
		ApproachSiteCodes5682004,
		ApproachSiteCodes5696005,
		ApproachSiteCodes5697001,
		ApproachSiteCodes5709001,
		ApproachSiteCodes5713008,
		ApproachSiteCodes5717009,
		ApproachSiteCodes5718004,
		ApproachSiteCodes5727003,
		ApproachSiteCodes5742000,
		ApproachSiteCodes5751008,
		ApproachSiteCodes5769004,
		ApproachSiteCodes5780004,
		ApproachSiteCodes5798000,
		ApproachSiteCodes5802004,
		ApproachSiteCodes5814007,
		ApproachSiteCodes5815008,
		ApproachSiteCodes5816009,
		ApproachSiteCodes5825003,
		ApproachSiteCodes5828001,
		ApproachSiteCodes5847003,
		ApproachSiteCodes5854009,
		ApproachSiteCodes5868002,
		ApproachSiteCodes5872003,
		ApproachSiteCodes5881009,
		ApproachSiteCodes5882002,
		ApproachSiteCodes5889006,
		ApproachSiteCodes5890002,
		ApproachSiteCodes5893000,
		ApproachSiteCodes5898009,
		ApproachSiteCodes5923009,
		ApproachSiteCodes5926001,
		ApproachSiteCodes5928000,
		ApproachSiteCodes5942008,
		ApproachSiteCodes5943003,
		ApproachSiteCodes5944009,
		ApproachSiteCodes5948007,
		ApproachSiteCodes5951000,
		ApproachSiteCodes5953002,
		ApproachSiteCodes5976004,
		ApproachSiteCodes5979006,
		ApproachSiteCodes5996007,
		ApproachSiteCodes6001004,
		ApproachSiteCodes6004007,
		ApproachSiteCodes6006009,
		ApproachSiteCodes6009002,
		ApproachSiteCodes6014003,
		ApproachSiteCodes6023000,
		ApproachSiteCodes6032003,
		ApproachSiteCodes6046003,
		ApproachSiteCodes6050005,
		ApproachSiteCodes6059006,
		ApproachSiteCodes6062009,
		ApproachSiteCodes6073002,
		ApproachSiteCodes6074008,
		ApproachSiteCodes6076005,
		ApproachSiteCodes6104005,
		ApproachSiteCodes6105006,
		ApproachSiteCodes6110005,
		ApproachSiteCodes6216007,
		ApproachSiteCodes6217003,
		ApproachSiteCodes6229007,
		ApproachSiteCodes6253001,
		ApproachSiteCodes6268000,
		ApproachSiteCodes6269008,
		ApproachSiteCodes6279005,
		ApproachSiteCodes6317000,
		ApproachSiteCodes6325003,
		ApproachSiteCodes6326002,
		ApproachSiteCodes6335009,
		ApproachSiteCodes6359004,
		ApproachSiteCodes6371005,
		ApproachSiteCodes6375001,
		ApproachSiteCodes6392005,
		ApproachSiteCodes6404004,
		ApproachSiteCodes6413002,
		ApproachSiteCodes6417001,
		ApproachSiteCodes6423006,
		ApproachSiteCodes6424000,
		ApproachSiteCodes6445007,
		ApproachSiteCodes6448009,
		ApproachSiteCodes6450001,
		ApproachSiteCodes6472004,
		ApproachSiteCodes6504002,
		ApproachSiteCodes6511003,
		ApproachSiteCodes6530003,
		ApproachSiteCodes6533001,
		ApproachSiteCodes6538005,
		ApproachSiteCodes6541001,
		ApproachSiteCodes6544009,
		ApproachSiteCodes6550004,
		ApproachSiteCodes6551000,
		ApproachSiteCodes6553002,
		ApproachSiteCodes6564004,
		ApproachSiteCodes6566002,
		ApproachSiteCodes6572002,
		ApproachSiteCodes6598008,
		ApproachSiteCodes6606008,
		ApproachSiteCodes6608009,
		ApproachSiteCodes6620001,
		ApproachSiteCodes6623004,
		ApproachSiteCodes6633007,
		ApproachSiteCodes6643005,
		ApproachSiteCodes6646002,
		ApproachSiteCodes6649009,
		ApproachSiteCodes6651008,
		ApproachSiteCodes6684008,
		ApproachSiteCodes6685009,
		ApproachSiteCodes6711001,
		ApproachSiteCodes6720005,
		ApproachSiteCodes6731002,
		ApproachSiteCodes6739000,
		ApproachSiteCodes6742006,
		ApproachSiteCodes6750002,
		ApproachSiteCodes6757004,
		ApproachSiteCodes6787005,
		ApproachSiteCodes6789008,
		ApproachSiteCodes6799003,
		ApproachSiteCodes6805009,
		ApproachSiteCodes6820003,
		ApproachSiteCodes6828005,
		ApproachSiteCodes6829002,
		ApproachSiteCodes6834003,
		ApproachSiteCodes6841009,
		ApproachSiteCodes6844001,
		ApproachSiteCodes6850006,
		ApproachSiteCodes6864006,
		ApproachSiteCodes6866008,
		ApproachSiteCodes6871001,
		ApproachSiteCodes6894000,
		ApproachSiteCodes6902008,
		ApproachSiteCodes6905005,
		ApproachSiteCodes6912001,
		ApproachSiteCodes6914000,
		ApproachSiteCodes6921000,
		ApproachSiteCodes6930008,
		ApproachSiteCodes6944002,
		ApproachSiteCodes6969002,
		ApproachSiteCodes6975006,
		ApproachSiteCodes6981003,
		ApproachSiteCodes6987004,
		ApproachSiteCodes6989001,
		ApproachSiteCodes6991009,
		ApproachSiteCodes7035006,
		ApproachSiteCodes7050002,
		ApproachSiteCodes7067009,
		ApproachSiteCodes7076002,
		ApproachSiteCodes7083009,
		ApproachSiteCodes7091000,
		ApproachSiteCodes7099003,
		ApproachSiteCodes7117004,
		ApproachSiteCodes7148007,
		ApproachSiteCodes7149004,
		ApproachSiteCodes7154008,
		ApproachSiteCodes7160008,
		ApproachSiteCodes7167006,
		ApproachSiteCodes7173007,
		ApproachSiteCodes7188002,
		ApproachSiteCodes7192009,
		ApproachSiteCodes7227003,
		ApproachSiteCodes7234001,
		ApproachSiteCodes7242000,
		ApproachSiteCodes7295002,
		ApproachSiteCodes7296001,
		ApproachSiteCodes7311008,
		ApproachSiteCodes7344002,
		ApproachSiteCodes7345001,
		ApproachSiteCodes7362006,
		ApproachSiteCodes7376007,
		ApproachSiteCodes7378008,
		ApproachSiteCodes7384006,
		ApproachSiteCodes7404008,
		ApproachSiteCodes7435002,
		ApproachSiteCodes7471001,
		ApproachSiteCodes7477002,
		ApproachSiteCodes7480001,
		ApproachSiteCodes7494000,
		ApproachSiteCodes7498002,
		ApproachSiteCodes7507003,
		ApproachSiteCodes7524009,
		ApproachSiteCodes7532001,
		ApproachSiteCodes7554004,
		ApproachSiteCodes7566005,
		ApproachSiteCodes7569003,
		ApproachSiteCodes7591005,
		ApproachSiteCodes7597009,
		ApproachSiteCodes7605000,
		ApproachSiteCodes7610001,
		ApproachSiteCodes7629007,
		ApproachSiteCodes7651004,
		ApproachSiteCodes7652006,
		ApproachSiteCodes7657000,
		ApproachSiteCodes7658005,
		ApproachSiteCodes7697002,
		ApproachSiteCodes7712004,
		ApproachSiteCodes7726008,
		ApproachSiteCodes7736000,
		ApproachSiteCodes7742001,
		ApproachSiteCodes7748002,
		ApproachSiteCodes7755000,
		ApproachSiteCodes7756004,
		ApproachSiteCodes7764005,
		ApproachSiteCodes7769000,
		ApproachSiteCodes7783003,
		ApproachSiteCodes7820009,
		ApproachSiteCodes7829005,
		ApproachSiteCodes7832008,
		ApproachSiteCodes7835005,
		ApproachSiteCodes7840002,
		ApproachSiteCodes7841003,
		ApproachSiteCodes7844006,
		ApproachSiteCodes7851002,
		ApproachSiteCodes7854005,
		ApproachSiteCodes7872004,
		ApproachSiteCodes7874003,
		ApproachSiteCodes7880006,
		ApproachSiteCodes7884002,
		ApproachSiteCodes7885001,
		ApproachSiteCodes7892006,
		ApproachSiteCodes7896009,
		ApproachSiteCodes7911004,
		ApproachSiteCodes7925003,
		ApproachSiteCodes7936005,
		ApproachSiteCodes7944005,
		ApproachSiteCodes7954009,
		ApproachSiteCodes7967007,
		ApproachSiteCodes7986004,
		ApproachSiteCodes7991003,
		ApproachSiteCodes7999001,
		ApproachSiteCodes8001006,
		ApproachSiteCodes8012006,
		ApproachSiteCodes8017000,
		ApproachSiteCodes8024004,
		ApproachSiteCodes8039003,
		ApproachSiteCodes8040001,
		ApproachSiteCodes8045006,
		ApproachSiteCodes8057002,
		ApproachSiteCodes8059004,
		ApproachSiteCodes8067007,
		ApproachSiteCodes8068002,
		ApproachSiteCodes8079007,
		ApproachSiteCodes8091003,
		ApproachSiteCodes8100009,
		ApproachSiteCodes8111001,
		ApproachSiteCodes8112008,
		ApproachSiteCodes8119004,
		ApproachSiteCodes8128003,
		ApproachSiteCodes8133004,
		ApproachSiteCodes8157004,
		ApproachSiteCodes8158009,
		ApproachSiteCodes8159001,
		ApproachSiteCodes8160006,
		ApproachSiteCodes8161005,
		ApproachSiteCodes8165001,
		ApproachSiteCodes8205005,
		ApproachSiteCodes8225009,
		ApproachSiteCodes8242003,
		ApproachSiteCodes8251006,
		ApproachSiteCodes8264007,
		ApproachSiteCodes8265008,
		ApproachSiteCodes8266009,
		ApproachSiteCodes8289001,
		ApproachSiteCodes8292002,
		ApproachSiteCodes8314003,
		ApproachSiteCodes8334002,
		ApproachSiteCodes8356004,
		ApproachSiteCodes8369000,
		ApproachSiteCodes8373002,
		ApproachSiteCodes8387002,
		ApproachSiteCodes8389004,
		ApproachSiteCodes8412003,
		ApproachSiteCodes8415001,
		ApproachSiteCodes8454000,
		ApproachSiteCodes8464009,
		ApproachSiteCodes8482007,
		ApproachSiteCodes8483002,
		ApproachSiteCodes8496001,
		ApproachSiteCodes8523001,
		ApproachSiteCodes8546004,
		ApproachSiteCodes8556000,
		ApproachSiteCodes8559007,
		ApproachSiteCodes8560002,
		ApproachSiteCodes8580001,
		ApproachSiteCodes8595004,
		ApproachSiteCodes8598002,
		ApproachSiteCodes8600008,
		ApproachSiteCodes8603005,
		ApproachSiteCodes8604004,
		ApproachSiteCodes8608001,
		ApproachSiteCodes8617001,
		ApproachSiteCodes8623006,
		ApproachSiteCodes8629005,
		ApproachSiteCodes8640002,
		ApproachSiteCodes8668003,
		ApproachSiteCodes8671006,
		ApproachSiteCodes8677005,
		ApproachSiteCodes8688004,
		ApproachSiteCodes8695008,
		ApproachSiteCodes8710005,
		ApproachSiteCodes8711009,
		ApproachSiteCodes8714001,
		ApproachSiteCodes8752000,
		ApproachSiteCodes8775007,
		ApproachSiteCodes8784007,
		ApproachSiteCodes8810002,
		ApproachSiteCodes8814006,
		ApproachSiteCodes8815007,
		ApproachSiteCodes8820007,
		ApproachSiteCodes8821006,
		ApproachSiteCodes8827005,
		ApproachSiteCodes8839002,
		ApproachSiteCodes8845005,
		ApproachSiteCodes8850004,
		ApproachSiteCodes8854008,
		ApproachSiteCodes8862000,
		ApproachSiteCodes8873007,
		ApproachSiteCodes8887007,
		ApproachSiteCodes8892009,
		ApproachSiteCodes8894005,
		ApproachSiteCodes8897003,
		ApproachSiteCodes8907008,
		ApproachSiteCodes8910001,
		ApproachSiteCodes8911002,
		ApproachSiteCodes8928004,
		ApproachSiteCodes8931003,
		ApproachSiteCodes8935007,
		ApproachSiteCodes8942007,
		ApproachSiteCodes8965002,
		ApproachSiteCodes8966001,
		ApproachSiteCodes8983005,
		ApproachSiteCodes8988001,
		ApproachSiteCodes8993003,
		ApproachSiteCodes9000002,
		ApproachSiteCodes9003000,
		ApproachSiteCodes9018004,
		ApproachSiteCodes9040008,
		ApproachSiteCodes9055004,
		ApproachSiteCodes9073001,
		ApproachSiteCodes9081000,
		ApproachSiteCodes9086005,
		ApproachSiteCodes9089003,
		ApproachSiteCodes9108007,
		ApproachSiteCodes9127001,
		ApproachSiteCodes9156001,
		ApproachSiteCodes9185007,
		ApproachSiteCodes9186008,
		ApproachSiteCodes9188009,
		ApproachSiteCodes9208002,
		ApproachSiteCodes9212008,
		ApproachSiteCodes9229006,
		ApproachSiteCodes9231002,
		ApproachSiteCodes9240003,
		ApproachSiteCodes9242006,
		ApproachSiteCodes9258009,
		ApproachSiteCodes9261005,
		ApproachSiteCodes9262003,
		ApproachSiteCodes9284003,
		ApproachSiteCodes9290004,
		ApproachSiteCodes9305001,
		ApproachSiteCodes9317004,
		ApproachSiteCodes9320007,
		ApproachSiteCodes9321006,
		ApproachSiteCodes9325002,
		ApproachSiteCodes9332006,
		ApproachSiteCodes9348007,
		ApproachSiteCodes9379006,
		ApproachSiteCodes9380009,
		ApproachSiteCodes9384000,
		ApproachSiteCodes9390001,
		ApproachSiteCodes9432007,
		ApproachSiteCodes9438006,
		ApproachSiteCodes9454009,
		ApproachSiteCodes9455005,
		ApproachSiteCodes9475001,
		ApproachSiteCodes9481009,
		ApproachSiteCodes9490002,
		ApproachSiteCodes9498009,
		ApproachSiteCodes9502002,
		ApproachSiteCodes9512009,
		ApproachSiteCodes9535007,
		ApproachSiteCodes9558005,
		ApproachSiteCodes9566001,
		ApproachSiteCodes9568000,
		ApproachSiteCodes9596006,
		ApproachSiteCodes9609000,
		ApproachSiteCodes9625005,
		ApproachSiteCodes9642004,
		ApproachSiteCodes9646001,
		ApproachSiteCodes9654004,
		ApproachSiteCodes9659009,
		ApproachSiteCodes9662007,
		ApproachSiteCodes9668006,
		ApproachSiteCodes9677004,
		ApproachSiteCodes9683001,
		ApproachSiteCodes9684007,
		ApproachSiteCodes9708001,
		ApproachSiteCodes9736006,
		ApproachSiteCodes9743000,
		ApproachSiteCodes9758008,
		ApproachSiteCodes9770007,
		ApproachSiteCodes9775002,
		ApproachSiteCodes9779008,
		ApproachSiteCodes9783008,
		ApproachSiteCodes9791004,
		ApproachSiteCodes9796009,
		ApproachSiteCodes9813009,
		ApproachSiteCodes9825007,
		ApproachSiteCodes9837009,
		ApproachSiteCodes9840009,
		ApproachSiteCodes9841008,
		ApproachSiteCodes9846003,
		ApproachSiteCodes9847007,
		ApproachSiteCodes9849005,
		ApproachSiteCodes9870004,
		ApproachSiteCodes9875009,
		ApproachSiteCodes9878006,
		ApproachSiteCodes9880000,
		ApproachSiteCodes9881001,
		ApproachSiteCodes9891007,
		ApproachSiteCodes9898001,
		ApproachSiteCodes9951005,
		ApproachSiteCodes9968009,
		ApproachSiteCodes9970000,
		ApproachSiteCodes9976006,
		ApproachSiteCodes9994000,
		ApproachSiteCodes9999005,
		ApproachSiteCodes10013000,
		ApproachSiteCodes10024003,
		ApproachSiteCodes10025002,
		ApproachSiteCodes10026001,
		ApproachSiteCodes10036009,
		ApproachSiteCodes10042008,
		ApproachSiteCodes10047002,
		ApproachSiteCodes10052007,
		ApproachSiteCodes10056005,
		ApproachSiteCodes10062000,
		ApproachSiteCodes10119003,
		ApproachSiteCodes10124000,
		ApproachSiteCodes10134009,
		ApproachSiteCodes10141003,
		ApproachSiteCodes10145007,
		ApproachSiteCodes10148009,
		ApproachSiteCodes10149001,
		ApproachSiteCodes10151002,
		ApproachSiteCodes10167008,
		ApproachSiteCodes10176001,
		ApproachSiteCodes10200004,
		ApproachSiteCodes10208006,
		ApproachSiteCodes10209003,
		ApproachSiteCodes10245000,
		ApproachSiteCodes10271001,
		ApproachSiteCodes10293006,
		ApproachSiteCodes10296003,
		ApproachSiteCodes10299005,
		ApproachSiteCodes10328008,
		ApproachSiteCodes10339006,
		ApproachSiteCodes10410005,
		ApproachSiteCodes10415000,
		ApproachSiteCodes10417008,
		ApproachSiteCodes10418003,
	}
}

func FindApproachSiteCodes(filter string) []ApproachSiteCodes {
	ret := make([]ApproachSiteCodes, 0)
	for _, at := range AllApproachSiteCodes() {
		if strings.ToLower(at.String())[0:len(filter)] == strings.ToLower(filter) {
			ret = append(ret, at)
		}
	}
	return ret
}

func (cpt ApproachSiteCodes) ToString() {
	fmt.Println(cpt.String())
}

func (cpt ApproachSiteCodes) String() string {
	switch cpt {
	case ApproachSiteCodes91723000:
		return "Anatomical structure"
	case ApproachSiteCodes106004:
		return "Posterior carpal region"
	case ApproachSiteCodes107008:
		return "Fetal part of placenta"
	case ApproachSiteCodes108003:
		return "Entire condylar emissary vein"
	case ApproachSiteCodes110001:
		return "Visceral layer of Bowman's capsule"
	case ApproachSiteCodes111002:
		return "Parathyroid gland"
	case ApproachSiteCodes116007:
		return "Subcutaneous tissue of medial surface of index finger"
	case ApproachSiteCodes124002:
		return "Coronoid process of mandible"
	case ApproachSiteCodes149003:
		return "Central pair of microtubules, cilium or flagellum, not bacterial"
	case ApproachSiteCodes155008:
		return "Deep circumflex artery of ilium"
	case ApproachSiteCodes167005:
		return "Supraclavicular part of brachial plexus"
	case ApproachSiteCodes202009:
		return "Anterior division of renal artery"
	case ApproachSiteCodes205006:
		return "Entire left commissure of aortic valve"
	case ApproachSiteCodes206007:
		return "Gluteus maximus muscle"
	case ApproachSiteCodes221001:
		return "Articular surface, phalanges, of fourth metacarpal bone"
	case ApproachSiteCodes227002:
		return "Canal of Hering"
	case ApproachSiteCodes233006:
		return "Hepatocolic ligament"
	case ApproachSiteCodes235004:
		return "Superior labial artery"
	case ApproachSiteCodes246001:
		return "Lateral vestibular nucleus"
	case ApproachSiteCodes247005:
		return "Mesotympanum"
	case ApproachSiteCodes251007:
		return "Pectoral region"
	case ApproachSiteCodes256002:
		return "Kupffer cell"
	case ApproachSiteCodes263002:
		return "Thoracic nerve"
	case ApproachSiteCodes266005:
		return "Right lower lobe of lung"
	case ApproachSiteCodes272005:
		return "Superior articular process of lumbar vertebra"
	case ApproachSiteCodes273000:
		return "Lateral myocardium"
	case ApproachSiteCodes283001:
		return "Central axillary lymph node"
	case ApproachSiteCodes284007:
		return "Entire flexor tendon and tendon sheath of fourth toe"
	case ApproachSiteCodes289002:
		return "Metacarpophalangeal joint of index finger"
	case ApproachSiteCodes301000:
		return "Fifth metatarsal bone"
	case ApproachSiteCodes311007:
		return "Plantar surface of great toe"
	case ApproachSiteCodes315003:
		return "Skin of umbilicus"
	case ApproachSiteCodes318001:
		return "Cardiac impression of liver"
	case ApproachSiteCodes344001:
		return "Ankle"
	case ApproachSiteCodes345000:
		return "Atrioventricular bundle"
	case ApproachSiteCodes356000:
		return "Reticular corium"
	case ApproachSiteCodes393006:
		return "Wall of urinary bladder"
	case ApproachSiteCodes402006:
		return "Dental branches of inferior alveolar artery"
	case ApproachSiteCodes405008:
		return "Entire posterior temporal diploic vein"
	case ApproachSiteCodes414003:
		return "Gastric fundus"
	case ApproachSiteCodes422005:
		return "Inferior surface of tongue"
	case ApproachSiteCodes446003:
		return "Trochanteric bursa"
	case ApproachSiteCodes457008:
		return "Collateral ligament"
	case ApproachSiteCodes461002:
		return "Lateral corticospinal tract"
	case ApproachSiteCodes464005:
		return "Basophilic normoblast"
	case ApproachSiteCodes465006:
		return "Ascending frontal gyrus"
	case ApproachSiteCodes471000:
		return "Flexor hallucis longus muscle"
	case ApproachSiteCodes480000:
		return "Cardiopulmonary circulatory system"
	case ApproachSiteCodes485005:
		return "Transverse colon"
	case ApproachSiteCodes528006:
		return "Costal surface of lung"
	case ApproachSiteCodes552004:
		return "Entire vagus nerve parasympathetic fibers to cardiac plexus"
	case ApproachSiteCodes565008:
		return "Intervertebral disc space of fifth lumbar vertebra"
	case ApproachSiteCodes582005:
		return "Head of phalanx of great toe"
	case ApproachSiteCodes587004:
		return "Capsule of proximal interphalangeal joint of third toe"
	case ApproachSiteCodes589001:
		return "Interventricular septum"
	case ApproachSiteCodes595000:
		return "Palpebral fissure"
	case ApproachSiteCodes608002:
		return "Subcutaneous tissue of philtrum"
	case ApproachSiteCodes621009:
		return "Multivesicular body, internal vesicles"
	case ApproachSiteCodes635006:
		return "Tuberosity of distal phalanx of little toe"
	case ApproachSiteCodes650002:
		return "Superior articular process of seventh thoracic vertebra"
	case ApproachSiteCodes660006:
		return "Tracheal mucous membrane"
	case ApproachSiteCodes661005:
		return "Jaw"
	case ApproachSiteCodes667009:
		return "Embryonic structure"
	case ApproachSiteCodes688000:
		return "Fetal hyaloid artery"
	case ApproachSiteCodes691000:
		return "Small intestine submucosa"
	case ApproachSiteCodes692007:
		return "Body of ischium"
	case ApproachSiteCodes723004:
		return "Dense intermediate filament bundles"
	case ApproachSiteCodes774007:
		return "Head and neck structure"
	case ApproachSiteCodes790007:
		return "Visceral aspect of liver"
	case ApproachSiteCodes798000:
		return "Entire deep temporal vein"
	case ApproachSiteCodes808000:
		return "Posterior intercostal arteries"
	case ApproachSiteCodes809008:
		return "Fetal chondrocranium"
	case ApproachSiteCodes823005:
		return "Posterior cervical spinal cord nerve root"
	case ApproachSiteCodes830004:
		return "Spinous process of fifth thoracic vertebra"
	case ApproachSiteCodes836005:
		return "Oral region of face"
	case ApproachSiteCodes885000:
		return "Lamina muscularis of colonic mucous membrane"
	case ApproachSiteCodes895007:
		return "Anterior cruciate ligament of knee joint"
	case ApproachSiteCodes917004:
		return "Superior laryngeal aperture"
	case ApproachSiteCodes921006:
		return "Thyrohyoid branch of ansa cervicalis"
	case ApproachSiteCodes947002:
		return "Crus of diaphragm"
	case ApproachSiteCodes955009:
		return "Bronchus"
	case ApproachSiteCodes976004:
		return "Ovarian vein"
	case ApproachSiteCodes996007:
		return "Meningeal branches of occipital artery"
	case ApproachSiteCodes1005009:
		return "Entire diaphragmatic lymph node"
	case ApproachSiteCodes1012000:
		return "Fibrous portion of pericardium"
	case ApproachSiteCodes1015003:
		return "Peritonsillar tissue"
	case ApproachSiteCodes1028005:
		return "Sebaceous gland"
	case ApproachSiteCodes1030007:
		return "Vesicular bursa of sternohyoid muscle"
	case ApproachSiteCodes1063000:
		return "Frontozygomatic suture of skull"
	case ApproachSiteCodes1075005:
		return "Promonocyte"
	case ApproachSiteCodes1076006:
		return "Entire subcutaneous prepatellar bursa"
	case ApproachSiteCodes1086007:
		return "Female"
	case ApproachSiteCodes1087003:
		return "Sternothyroid muscle"
	case ApproachSiteCodes1092001:
		return "Superior occipital gyrus"
	case ApproachSiteCodes1099005:
		return "Thymic cortex"
	case ApproachSiteCodes1101003:
		return "Cranial cavity"
	case ApproachSiteCodes1106008:
		return "Entire major calyx"
	case ApproachSiteCodes1110006:
		return "Tarsal gland"
	case ApproachSiteCodes1122009:
		return "Inferior longitudinal muscle of tongue"
	case ApproachSiteCodes1136004:
		return "Aortopulmonary septum"
	case ApproachSiteCodes1159005:
		return "Frenulum linguae"
	case ApproachSiteCodes1172006:
		return "Odontoid process of axis"
	case ApproachSiteCodes1173001:
		return "Mandibular nerve"
	case ApproachSiteCodes1174007:
		return "Chromosomes, group E"
	case ApproachSiteCodes1193009:
		return "Teres major muscle"
	case ApproachSiteCodes1216008:
		return "Synostosis"
	case ApproachSiteCodes1231004:
		return "Meninges"
	case ApproachSiteCodes1236009:
		return "Duodenal serosa"
	case ApproachSiteCodes1243003:
		return "Inferior articular process of sixth cervical vertebra"
	case ApproachSiteCodes1246006:
		return "Dorsal digital nerves of radial nerve"
	case ApproachSiteCodes1263005:
		return "Distinctive arrangement of microtubules"
	case ApproachSiteCodes1277008:
		return "Vertebral nerve"
	case ApproachSiteCodes1307006:
		return "Glottis"
	case ApproachSiteCodes1311000:
		return "Telogen hair"
	case ApproachSiteCodes1350001:
		return "Deep flexor tendon of index finger"
	case ApproachSiteCodes1353004:
		return "Gastric serosa"
	case ApproachSiteCodes1403006:
		return "Vastus lateralis muscle"
	case ApproachSiteCodes1425000:
		return "Posterior limb of stapes"
	case ApproachSiteCodes1439000:
		return "Paravesicular lymph node"
	case ApproachSiteCodes1441004:
		return "Ventricle of larynx"
	case ApproachSiteCodes1456008:
		return "Yellow fibrocartilage"
	case ApproachSiteCodes1467009:
		return "Parietal branch of superficial temporal artery"
	case ApproachSiteCodes1484003:
		return "Metatarsus"
	case ApproachSiteCodes1490004:
		return "Soft tissues of trunk"
	case ApproachSiteCodes1502004:
		return "Anterior cecal artery"
	case ApproachSiteCodes1511004:
		return "Ejaculatory duct"
	case ApproachSiteCodes1516009:
		return "Frontomental diameter of head"
	case ApproachSiteCodes1527006:
		return "Lamina of fourth thoracic vertebra"
	case ApproachSiteCodes1537001:
		return "Intervertebral disc of eleventh thoracic vertebra"
	case ApproachSiteCodes1541002:
		return "Coccygeal plexus"
	case ApproachSiteCodes1562001:
		return "Nucleus pulposus of intervertebral disc of third lumbar vertebra"
	case ApproachSiteCodes1580005:
		return "Ventral lateral nucleus of thalamus"
	case ApproachSiteCodes1581009:
		return "Ileal arteries"
	case ApproachSiteCodes1584001:
		return "Entire symphysis"
	case ApproachSiteCodes1600003:
		return "Splenius capitis muscle"
	case ApproachSiteCodes1605008:
		return "Histioblast"
	case ApproachSiteCodes1610007:
		return "Otoconia"
	case ApproachSiteCodes1611006:
		return "Entire paramammary lymph node"
	case ApproachSiteCodes1617005:
		return "Intrinsic larynx"
	case ApproachSiteCodes1620002:
		return "Metaphase nucleus"
	case ApproachSiteCodes1626008:
		return "Third thoracic vertebra"
	case ApproachSiteCodes1627004:
		return "Medial collateral ligament of knee joint"
	case ApproachSiteCodes1630006:
		return "Supraorbital vein"
	case ApproachSiteCodes1631005:
		return "Foregut"
	case ApproachSiteCodes1650005:
		return "Hilum of left lung"
	case ApproachSiteCodes1655000:
		return "Transverse peduncular tract nucleus"
	case ApproachSiteCodes1659006:
		return "Dorsomedial nucleus of thalamus"
	case ApproachSiteCodes1684009:
		return "Ligamentum teres of liver"
	case ApproachSiteCodes1706004:
		return "Thymic lobule"
	case ApproachSiteCodes1707008:
		return "Ventral nuclear group of thalamus"
	case ApproachSiteCodes1711002:
		return "Periorbital region"
	case ApproachSiteCodes1716007:
		return "Cupula ampullaris"
	case ApproachSiteCodes1721005:
		return "Right tonsil"
	case ApproachSiteCodes1729007:
		return "Central tegmental tract"
	case ApproachSiteCodes1732005:
		return "Thoracic duct"
	case ApproachSiteCodes1765002:
		return "Lymphatic of thorax"
	case ApproachSiteCodes1780008:
		return "Premelanosome"
	case ApproachSiteCodes1781007:
		return "Sacroiliac region"
	case ApproachSiteCodes1797002:
		return "Naris"
	case ApproachSiteCodes1818002:
		return "Greater circulus arteriosus of iris"
	case ApproachSiteCodes1825009:
		return "Entire root of nose"
	case ApproachSiteCodes1832000:
		return "Bulbar conjunctiva"
	case ApproachSiteCodes1840006:
		return "Arrector pili muscles"
	case ApproachSiteCodes1849007:
		return "Pharyngeal recess"
	case ApproachSiteCodes1853009:
		return "Suprahyoid muscles"
	case ApproachSiteCodes1874005:
		return "Entire promontory common iliac lymph node"
	case ApproachSiteCodes1895000:
		return "Entire musculophrenic vein"
	case ApproachSiteCodes1902009:
		return "Skin of external ear"
	case ApproachSiteCodes1910005:
		return "Entire ear"
	case ApproachSiteCodes1918003:
		return "Suprarenal aorta"
	case ApproachSiteCodes1927002:
		return "Entire left elbow region"
	case ApproachSiteCodes1992003:
		return "Porus acusticus internus"
	case ApproachSiteCodes1997009:
		return "Cingulum dentis"
	case ApproachSiteCodes2010005:
		return "Clavicular facet of scapula"
	case ApproachSiteCodes2020000:
		return "Superior thoracic artery"
	case ApproachSiteCodes2031008:
		return "Spinal cord ventral median fissure"
	case ApproachSiteCodes2033006:
		return "Right fallopian tube"
	case ApproachSiteCodes2044003:
		return "Vaginal nerves"
	case ApproachSiteCodes2048000:
		return "Lingual tonsil"
	case ApproachSiteCodes2049008:
		return "Chorionic villi"
	case ApproachSiteCodes2059009:
		return "Skin of ear lobule"
	case ApproachSiteCodes2071003:
		return "Reticular formation of spinal cord"
	case ApproachSiteCodes2076008:
		return "Head of phalanx of hand"
	case ApproachSiteCodes2083001:
		return "Nucleus ambiguus"
	case ApproachSiteCodes2095001:
		return "Accessory sinus"
	case ApproachSiteCodes2123001:
		return "Tuberomammillary nucleus"
	case ApproachSiteCodes2150006:
		return "Urinary tract transitional epithelial cell"
	case ApproachSiteCodes2156000:
		return "Glial cell"
	case ApproachSiteCodes2160002:
		return "Ligamentum arteriosum"
	case ApproachSiteCodes2175005:
		return "Pharyngeal cavity"
	case ApproachSiteCodes2182009:
		return "Endometrial zona basalis"
	case ApproachSiteCodes2192001:
		return "Clavicular part of pectoralis major muscle"
	case ApproachSiteCodes2205003:
		return "Lamina of fifth thoracic vertebra"
	case ApproachSiteCodes2209009:
		return "Cerebral basal surface"
	case ApproachSiteCodes2236006:
		return "Lesser osseous pelvis"
	case ApproachSiteCodes2246008:
		return "Type I hair cell"
	case ApproachSiteCodes2255006:
		return "Subserosa"
	case ApproachSiteCodes2292006:
		return "Nasopharyngeal gland"
	case ApproachSiteCodes2302002:
		return "Veins of the knee"
	case ApproachSiteCodes2305000:
		return "Spinous process of cervical vertebra"
	case ApproachSiteCodes2306004:
		return "Base of third metacarpal bone"
	case ApproachSiteCodes2327009:
		return "Entire salivary seromucous gland (body structure)"
	case ApproachSiteCodes2330002:
		return "Segmental bronchial branches"
	case ApproachSiteCodes2332005:
		return "Metencephalon of fetus"
	case ApproachSiteCodes2334006:
		return "Calyx"
	case ApproachSiteCodes2349003:
		return "Nasal suture of skull"
	case ApproachSiteCodes2372001:
		return "Medial surface of toe"
	case ApproachSiteCodes2383005:
		return "Papillary muscles of right ventricle"
	case ApproachSiteCodes2389009:
		return "Superior margin of adrenal gland"
	case ApproachSiteCodes2395005:
		return "Transverse facial artery"
	case ApproachSiteCodes2397002:
		return "First metatarsal facet of medial cuneiform bone"
	case ApproachSiteCodes2400006:
		return "Mandibular left first premolar tooth"
	case ApproachSiteCodes2402003:
		return "Dorsum of foot"
	case ApproachSiteCodes2421006:
		return "Submaxillary ganglion"
	case ApproachSiteCodes2433001:
		return "Digital tendon and tendon sheath of foot"
	case ApproachSiteCodes2436009:
		return "Tunica intima of vein"
	case ApproachSiteCodes2453002:
		return "Subcutaneous tissue of posterior surface of forearm"
	case ApproachSiteCodes2454008:
		return "Articular surface, third metacarpal, of second metacarpal bone"
	case ApproachSiteCodes2484000:
		return "Skin of frenulum of clitoris"
	case ApproachSiteCodes2489005:
		return "Medial check ligament of eye"
	case ApproachSiteCodes2499000:
		return "Entire cisterna pontis"
	case ApproachSiteCodes2502001:
		return "Membrane of lysosome"
	case ApproachSiteCodes2504000:
		return "Pancreatic plexus"
	case ApproachSiteCodes2510000:
		return "Femoral triangle"
	case ApproachSiteCodes2539000:
		return "Superior rectal artery"
	case ApproachSiteCodes2543001:
		return "Cuboid articular facet of fourth metatarsal bone"
	case ApproachSiteCodes2550002:
		return "Phalanx of thumb"
	case ApproachSiteCodes2577006:
		return "Gracilis muscle"
	case ApproachSiteCodes2579009:
		return "Plasmablast"
	case ApproachSiteCodes2592007:
		return "All extremities"
	case ApproachSiteCodes2600000:
		return "Flexor pollicis longus muscle tendon"
	case ApproachSiteCodes2620004:
		return "Intervertebral disc of third thoracic vertebra"
	case ApproachSiteCodes2639009:
		return "Neuroendocrine tissue"
	case ApproachSiteCodes2653009:
		return "Posterior thalamic radiation of internal capsule"
	case ApproachSiteCodes2666009:
		return "Semispinalis capitis muscle"
	case ApproachSiteCodes2672009:
		return "Anterior cutaneous branch of lumbosacral plexus"
	case ApproachSiteCodes2675006:
		return "Anterior ethmoidal artery"
	case ApproachSiteCodes2681003:
		return "Dorsal nerve of penis"
	case ApproachSiteCodes2682005:
		return "Mucous membrane of urinary bladder"
	case ApproachSiteCodes2686008:
		return "Medial olfactory gyrus"
	case ApproachSiteCodes2687004:
		return "Bowman's space"
	case ApproachSiteCodes2695000:
		return "Left maxillary sinus"
	case ApproachSiteCodes2703009:
		return "Entire calcarine artery"
	case ApproachSiteCodes2712006:
		return "Capsule of ankle joint"
	case ApproachSiteCodes2718005:
		return "Apical foramen of tooth"
	case ApproachSiteCodes2726002:
		return "Fold for stapes"
	case ApproachSiteCodes2730004:
		return "Entire vitelline vein of placenta"
	case ApproachSiteCodes2739003:
		return "Endometrium"
	case ApproachSiteCodes2741002:
		return "Medial occipitotemporal gyrus"
	case ApproachSiteCodes2746007:
		return "Circular layer of gastric muscularis"
	case ApproachSiteCodes2748008:
		return "Spinal cord"
	case ApproachSiteCodes2759004:
		return "Eccrine gland"
	case ApproachSiteCodes2771005:
		return "Lamina propria of ureter"
	case ApproachSiteCodes2789006:
		return "Apocrine gland"
	case ApproachSiteCodes2792005:
		return "Pars tensa of tympanic membrane"
	case ApproachSiteCodes2803000:
		return "Tendon sheath of popliteus muscle"
	case ApproachSiteCodes2810006:
		return "Cremasteric fascia"
	case ApproachSiteCodes2812003:
		return "Head of femur"
	case ApproachSiteCodes2824005:
		return "Spinous process of fourth thoracic vertebra"
	case ApproachSiteCodes2826007:
		return "Lamina of fourth lumbar vertebra"
	case ApproachSiteCodes2830005:
		return "Dorsal digital nerves of lateral hallux and medial second toe"
	case ApproachSiteCodes2839006:
		return "Perivesicular tissues of seminal vesicles"
	case ApproachSiteCodes2841007:
		return "Renal artery"
	case ApproachSiteCodes2845003:
		return "Respiratory epithelium"
	case ApproachSiteCodes2848001:
		return "Superficial epigastric artery"
	case ApproachSiteCodes2855004:
		return "Accessory cephalic vein"
	case ApproachSiteCodes2861001:
		return "Entire gland (organ)"
	case ApproachSiteCodes2894003:
		return "Posterior epiglottis"
	case ApproachSiteCodes2905008:
		return "Anterior ligament of uterus"
	case ApproachSiteCodes2909002:
		return "Posterior portion of diaphragmatic aspect of liver"
	case ApproachSiteCodes2920002:
		return "Facial nerve motor branch"
	case ApproachSiteCodes2922005:
		return "Posterior papillary muscle of left ventricle"
	case ApproachSiteCodes2923000:
		return "Subcutaneous tissue of supraorbital area"
	case ApproachSiteCodes2969000:
		return "Anatomical space structure (body structure)"
	case ApproachSiteCodes2979003:
		return "Medial cuneiform bone"
	case ApproachSiteCodes2986006:
		return "Talar facet of navicular bone of foot"
	case ApproachSiteCodes2998001:
		return "Entire right margin of uterus"
	case ApproachSiteCodes3003007:
		return "Anterior limb of internal capsule"
	case ApproachSiteCodes3008003:
		return "White fibrocartilage"
	case ApproachSiteCodes3028004:
		return "Transitional epithelial cell"
	case ApproachSiteCodes3039001:
		return "Subcutaneous tissue of thigh"
	case ApproachSiteCodes3042007:
		return "Glomerular urinary pole"
	case ApproachSiteCodes3054007:
		return "Articular surface, metacarpal, of phalanx of thumb"
	case ApproachSiteCodes3055008:
		return "Bone marrow of vertebral body"
	case ApproachSiteCodes3056009:
		return "Anteroventral nucleus of thalamus"
	case ApproachSiteCodes3057000:
		return "Nerve"
	case ApproachSiteCodes3058005:
		return "Peripheral nervous system"
	case ApproachSiteCodes3062004:
		return "Spinal arachnoid"
	case ApproachSiteCodes3068000:
		return "Seminal vesicle lumen"
	case ApproachSiteCodes3081007:
		return "Mitochondrion in division"
	case ApproachSiteCodes3093003:
		return "Tendinous arch of pelvic fascia"
	case ApproachSiteCodes3100007:
		return "Clinical crown of tooth"
	case ApproachSiteCodes3113001:
		return "Suprachoroidal space"
	case ApproachSiteCodes3117000:
		return "Dorsal surface of index finger"
	case ApproachSiteCodes3118005:
		return "Trabecula carnea of left ventricle"
	case ApproachSiteCodes3120008:
		return "Pleura"
	case ApproachSiteCodes3134008:
		return "Head of fourth metatarsal bone"
	case ApproachSiteCodes3138006:
		return "Bony tissue"
	case ApproachSiteCodes3153003:
		return "Tractus olivocochlearis"
	case ApproachSiteCodes3156006:
		return "Obturator artery"
	case ApproachSiteCodes3159004:
		return "Costocervical trunk"
	case ApproachSiteCodes3169005:
		return "Spinal nerve"
	case ApproachSiteCodes3178004:
		return "Raphe of soft palate"
	case ApproachSiteCodes3194006:
		return "Endocardium of right atrium"
	case ApproachSiteCodes3198009:
		return "Monostomatic sublingual gland"
	case ApproachSiteCodes3215002:
		return "Subcutaneous tissue of nuchal region"
	case ApproachSiteCodes3222005:
		return "All large arteries"
	case ApproachSiteCodes3227004:
		return "Left coronary artery main stem"
	case ApproachSiteCodes3236000:
		return "Posterior segment of right upper lobe of lung"
	case ApproachSiteCodes3243006:
		return "Parametrial lymph node"
	case ApproachSiteCodes3255000:
		return "Papillary area"
	case ApproachSiteCodes3262009:
		return "Root canal of tooth"
	case ApproachSiteCodes3279003:
		return "Pedicle of third cervical vertebra"
	case ApproachSiteCodes3295003:
		return "Ventral anterior nucleus"
	case ApproachSiteCodes3301002:
		return "Tectopontine fibers"
	case ApproachSiteCodes3302009:
		return "Splenic branches of splenic artery"
	case ApproachSiteCodes3315000:
		return "Vestibulospinal tract"
	case ApproachSiteCodes3332001:
		return "Occipitofrontal diameter of head"
	case ApproachSiteCodes3336003:
		return "Haversian canal"
	case ApproachSiteCodes3341006:
		return "Right lung"
	case ApproachSiteCodes3350008:
		return "Entire right commissure of pulmonic valve"
	case ApproachSiteCodes3362007:
		return "Intertragal incisure"
	case ApproachSiteCodes3366005:
		return "Anterior papillary muscle of left ventricle"
	case ApproachSiteCodes3370002:
		return "Supporting tissue of rectum"
	case ApproachSiteCodes3374006:
		return "Secondary spermatocyte"
	case ApproachSiteCodes3377004:
		return "Agger nasi"
	case ApproachSiteCodes3382006:
		return "Rima oris"
	case ApproachSiteCodes3383001:
		return "Nonsegmented basophil"
	case ApproachSiteCodes3394002:
		return "Suboccipitobregmatic diameter of head"
	case ApproachSiteCodes3395001:
		return "Superior palpebral arch"
	case ApproachSiteCodes3396000:
		return "Mesogastrium"
	case ApproachSiteCodes3400000:
		return "Cell of bone"
	case ApproachSiteCodes3409004:
		return "Lateral margin of forearm"
	case ApproachSiteCodes3417007:
		return "Rotator muscles"
	case ApproachSiteCodes3438001:
		return "Deep lymphatics of upper extremity"
	case ApproachSiteCodes3444002:
		return "Thalamostriate veins"
	case ApproachSiteCodes3447009:
		return "Penetrated oocyte"
	case ApproachSiteCodes3460003:
		return "Anterodorsal nucleus of thalamus"
	case ApproachSiteCodes3462006:
		return "Commissure of tricuspid valve"
	case ApproachSiteCodes3471002:
		return "Posterior midline of trunk"
	case ApproachSiteCodes3478008:
		return "Vastus medialis muscle"
	case ApproachSiteCodes3481003:
		return "Embryonic testis"
	case ApproachSiteCodes3488009:
		return "Annulate lamella, cisternal lumen"
	case ApproachSiteCodes3490005:
		return "Subcutaneous tissue of suboccipital region"
	case ApproachSiteCodes3524005:
		return "Lateral wall of mastoid antrum"
	case ApproachSiteCodes3538003:
		return "Capsule of proximal tibiofibular joint"
	case ApproachSiteCodes3541007:
		return "Dorsal metatarsal artery"
	case ApproachSiteCodes3553006:
		return "Thyroid capsule"
	case ApproachSiteCodes3556003:
		return "Dorsal nucleus of trapezoid body"
	case ApproachSiteCodes3563003:
		return "Muscularis of ureter"
	case ApproachSiteCodes3572006:
		return "Body of vertebra"
	case ApproachSiteCodes3578005:
		return "Body of gallbladder"
	case ApproachSiteCodes3582007:
		return "Gastrophrenic ligament"
	case ApproachSiteCodes3584008:
		return "Arch of tenth thoracic vertebra"
	case ApproachSiteCodes3594003:
		return "Straight part of longus colli muscle"
	case ApproachSiteCodes3608004:
		return "Ischiococcygeus muscle"
	case ApproachSiteCodes3609007:
		return "Occipital branch of posterior auricular artery"
	case ApproachSiteCodes3646006:
		return "Lamellipodium"
	case ApproachSiteCodes3663005:
		return "Tympanic ostium of Eustachian tube"
	case ApproachSiteCodes3665003:
		return "Pelvic wall"
	case ApproachSiteCodes3670005:
		return "Entire subpyloric lymph node"
	case ApproachSiteCodes3711007:
		return "Great vessel"
	case ApproachSiteCodes3743007:
		return "Lateral thoracic artery"
	case ApproachSiteCodes3761003:
		return "Nucleus pulposus of intervertebral disc of first thoracic vertebra"
	case ApproachSiteCodes3766008:
		return "Subcutaneous tissue of lower extremity"
	case ApproachSiteCodes3785006:
		return "Entire dorsal metacarpal ligament"
	case ApproachSiteCodes3788008:
		return "Superior segment of right lower lobe of lung"
	case ApproachSiteCodes3789000:
		return "Argentaffin cells"
	case ApproachSiteCodes3810000:
		return "Nasal septal cartilage"
	case ApproachSiteCodes3838008:
		return "Apex of coccyx"
	case ApproachSiteCodes3860006:
		return "Transplanted liver"
	case ApproachSiteCodes3865001:
		return "Interscapular region of back"
	case ApproachSiteCodes3867009:
		return "Dorsal surface of great toe"
	case ApproachSiteCodes3876002:
		return "Subcutaneous tissue of femoral region"
	case ApproachSiteCodes3877006:
		return "Common carotid plexus"
	case ApproachSiteCodes3910004:
		return "Subcutaneous tissue of lateral surface of fourth toe"
	case ApproachSiteCodes3916005:
		return "Occipital lymph node"
	case ApproachSiteCodes3924000:
		return "Pericardiophrenic artery"
	case ApproachSiteCodes3931001:
		return "Vestibular window"
	case ApproachSiteCodes3935005:
		return "Head of tenth rib"
	case ApproachSiteCodes3937002:
		return "Entorhinal cortex"
	case ApproachSiteCodes3954005:
		return "Lacrimal sac"
	case ApproachSiteCodes3956007:
		return "Fifth metatarsal articular facet of fourth metatarsal bone"
	case ApproachSiteCodes3959000:
		return "Rectus capitis muscle"
	case ApproachSiteCodes3960005:
		return "Olfactory tract"
	case ApproachSiteCodes3964001:
		return "Gyrus of brain"
	case ApproachSiteCodes3966004:
		return "Entire parietal branch of anterior cerebral artery"
	case ApproachSiteCodes3977005:
		return "Subcutaneous tissue of concha"
	case ApproachSiteCodes3984002:
		return "Deep veins of clitoris"
	case ApproachSiteCodes3989007:
		return "Medial nuclei of globus pallidus"
	case ApproachSiteCodes4015004:
		return "Chromosomes, group A"
	case ApproachSiteCodes4019005:
		return "Posterior commissure of labia majora"
	case ApproachSiteCodes4029003:
		return "Eosinophilic promyelocyte"
	case ApproachSiteCodes4061004:
		return "Lateral wall of orbit"
	case ApproachSiteCodes4066009:
		return "Capsule of proximal interphalangeal joint of index finger"
	case ApproachSiteCodes4072009:
		return "Fourth coccygeal vertebra"
	case ApproachSiteCodes4081003:
		return "Entire dorsal lingual vein"
	case ApproachSiteCodes4093007:
		return "Vagus nerve bronchial branches"
	case ApproachSiteCodes4111006:
		return "Macula of saccule"
	case ApproachSiteCodes4117005:
		return "Lumbosacral spinal cord central canal"
	case ApproachSiteCodes4121003:
		return "Superior frontal sulcus"
	case ApproachSiteCodes4146003:
		return "Artery of extremity"
	case ApproachSiteCodes4148002:
		return "Palmar surface of little finger"
	case ApproachSiteCodes4150005:
		return "Celiac plexus"
	case ApproachSiteCodes4158003:
		return "Abdominal aortic plexus"
	case ApproachSiteCodes4159006:
		return "Hyparterial bronchi"
	case ApproachSiteCodes4180000:
		return "Both lower extremities"
	case ApproachSiteCodes4193005:
		return "Entire extensor tendon and tendon sheath of fifth toe"
	case ApproachSiteCodes4205002:
		return "Türk cell"
	case ApproachSiteCodes4212006:
		return "Epithelial cell"
	case ApproachSiteCodes4215008:
		return "Head of second rib"
	case ApproachSiteCodes4247003:
		return "First metacarpal bone"
	case ApproachSiteCodes4258007:
		return "Posterior tibial veins"
	case ApproachSiteCodes4276000:
		return "Inferior articular process of seventh cervical vertebra"
	case ApproachSiteCodes4281009:
		return "Middle portion of ileum"
	case ApproachSiteCodes4295007:
		return "Paracortical area of lymph node"
	case ApproachSiteCodes4303006:
		return "Cartilage canal"
	case ApproachSiteCodes4312008:
		return "Anterior midline of abdomen"
	case ApproachSiteCodes4317002:
		return "Spinalis muscle"
	case ApproachSiteCodes4328003:
		return "Protoplasmic astrocyte"
	case ApproachSiteCodes4335006:
		return "Upper jaw"
	case ApproachSiteCodes4352005:
		return "Subchorionic space"
	case ApproachSiteCodes4358009:
		return "Lateral surface of little finger"
	case ApproachSiteCodes4360006:
		return "Stratum spinosum"
	case ApproachSiteCodes4369007:
		return "Small intestine mucous membrane"
	case ApproachSiteCodes4371007:
		return "Fourth metatarsal facet of lateral cuneiform bone"
	case ApproachSiteCodes4375003:
		return "Incisure of cartilaginous portion of auditory canal"
	case ApproachSiteCodes4377006:
		return "Parafascicular nucleus of thalamus"
	case ApproachSiteCodes4394008:
		return "Scala vestibuli"
	case ApproachSiteCodes4402002:
		return "Anterior articular surface for talus"
	case ApproachSiteCodes4419000:
		return "Tracheal submucosa"
	case ApproachSiteCodes4421005:
		return "Cell"
	case ApproachSiteCodes4430002:
		return "Clivus ossis sphenoidalis"
	case ApproachSiteCodes4432005:
		return "Ductus arteriosus"
	case ApproachSiteCodes4442007:
		return "Dental arch"
	case ApproachSiteCodes4486002:
		return "Accessory sinus gland"
	case ApproachSiteCodes4524000:
		return "Subclavian plexus"
	case ApproachSiteCodes4527007:
		return "Joint of lower extremity"
	case ApproachSiteCodes4537002:
		return "Internal medullary lamina of thalamus"
	case ApproachSiteCodes4548009:
		return "Lamellated granule, as in surfactant-secreting cell"
	case ApproachSiteCodes4549001:
		return "Entire vagus nerve parasympathetic fibers to liver, gallbladder, bile ducts and pancreas"
	case ApproachSiteCodes4566004:
		return "Tentorium cerebelli"
	case ApproachSiteCodes4573009:
		return "Desmosome"
	case ApproachSiteCodes4578000:
		return "Skin of posterior surface of thigh"
	case ApproachSiteCodes4588004:
		return "Middle trunk of brachial plexus"
	case ApproachSiteCodes4596009:
		return "Larynx"
	case ApproachSiteCodes4603002:
		return "Base of phalanx of foot"
	case ApproachSiteCodes4606005:
		return "Tubercle of eighth rib"
	case ApproachSiteCodes4621004:
		return "Lesser tuberosity of humerus"
	case ApproachSiteCodes4624007:
		return "Medullary cord"
	case ApproachSiteCodes4647008:
		return "Lipid droplet, homogeneous"
	case ApproachSiteCodes4651005:
		return "Tunica albuginea of corpus spongiosum"
	case ApproachSiteCodes4658004:
		return "Skin of nuchal region"
	case ApproachSiteCodes4677002:
		return "Basal lamina, inclusion in subepithelial location"
	case ApproachSiteCodes4703008:
		return "Cardinal vein"
	case ApproachSiteCodes4717004:
		return "Neutrophilic myelocyte"
	case ApproachSiteCodes4718009:
		return "Entire venous plexus of the foramen ovale basis cranii"
	case ApproachSiteCodes4743003:
		return "Ventral sacrococcygeal ligament"
	case ApproachSiteCodes4755009:
		return "Medial surface of great toe"
	case ApproachSiteCodes4759003:
		return "Gemellus muscle"
	case ApproachSiteCodes4766002:
		return "Supracardinal vein"
	case ApproachSiteCodes4768001:
		return "Perineal nerve"
	case ApproachSiteCodes4774001:
		return "Phrenic nerve pericardial branch"
	case ApproachSiteCodes4775000:
		return "Ventral posterior inferior nucleus"
	case ApproachSiteCodes4799000:
		return "Deiter's cell"
	case ApproachSiteCodes4810005:
		return "Uterine venous plexus"
	case ApproachSiteCodes4812002:
		return "Anterior tibial compartment"
	case ApproachSiteCodes4828007:
		return "Femoral canal"
	case ApproachSiteCodes4840007:
		return "Subcutaneous tissue of ring finger"
	case ApproachSiteCodes4843009:
		return "Membranous ampulla"
	case ApproachSiteCodes4861000:
		return "Tuberculum impar"
	case ApproachSiteCodes4866005:
		return "Constrictor pharyngeus muscle"
	case ApproachSiteCodes4870002:
		return "Dorsal tegmental nuclei of midbrain"
	case ApproachSiteCodes4871003:
		return "Spiral lamina of modiolus"
	case ApproachSiteCodes4881004:
		return "Entire sublingual vein"
	case ApproachSiteCodes4888005:
		return "Entire interlobular vein of kidney"
	case ApproachSiteCodes4897009:
		return "Cell membrane, prokaryotic"
	case ApproachSiteCodes4905007:
		return "Uterovaginal plexus"
	case ApproachSiteCodes4906008:
		return "Tympanic antrum"
	case ApproachSiteCodes4924005:
		return "Cerebellar gracile lobule"
	case ApproachSiteCodes4942000:
		return "Lymph node of lower extremity"
	case ApproachSiteCodes4954000:
		return "Radial notch of ulna"
	case ApproachSiteCodes4956003:
		return "Subcutaneous tissue of back"
	case ApproachSiteCodes4958002:
		return "Amygdaloid nucleus"
	case ApproachSiteCodes5001007:
		return "Superior temporal sulcus"
	case ApproachSiteCodes5023006:
		return "Yellow bone marrow"
	case ApproachSiteCodes5026003:
		return "Posterior surface of prostate"
	case ApproachSiteCodes5046008:
		return "Superficial dorsal veins of clitoris"
	case ApproachSiteCodes5068003:
		return "Obturator internus muscle ischial bursa"
	case ApproachSiteCodes5069006:
		return "Rugal column"
	case ApproachSiteCodes5076001:
		return "Infrasternal angle"
	case ApproachSiteCodes5115006:
		return "Posterior auricular vein"
	case ApproachSiteCodes5122003:
		return "Entire angle of first rib"
	case ApproachSiteCodes5128004:
		return "Suspensory ligament of lens"
	case ApproachSiteCodes5192008:
		return "Intervertebral foramen of twelfth thoracic vertebra"
	case ApproachSiteCodes5194009:
		return "Epithelium of lens"
	case ApproachSiteCodes5195005:
		return "Right external carotid artery"
	case ApproachSiteCodes5204005:
		return "Superior ileocecal recess"
	case ApproachSiteCodes5213007:
		return "Frontal veins"
	case ApproachSiteCodes5225005:
		return "Uterine ostium of fallopian tube"
	case ApproachSiteCodes5228007:
		return "Right cerebral hemisphere"
	case ApproachSiteCodes5229004:
		return "Mucosa of gallbladder"
	case ApproachSiteCodes5261000:
		return "Intervertebral disc of thoracic vertebra"
	case ApproachSiteCodes5272005:
		return "Skin of lateral portion of neck"
	case ApproachSiteCodes5279001:
		return "Foramen singulare"
	case ApproachSiteCodes5296000:
		return "Anterior mediastinal lymph node"
	case ApproachSiteCodes5324007:
		return "Superior pole of kidney"
	case ApproachSiteCodes5329002:
		return "Fourth cervical vertebra"
	case ApproachSiteCodes5336001:
		return "Inferior frontal gyrus"
	case ApproachSiteCodes5347008:
		return "Synaptic specialization, cytoplasmic"
	case ApproachSiteCodes5362005:
		return "Median arcuate ligament of diaphragm"
	case ApproachSiteCodes5366008:
		return "Hippocampus"
	case ApproachSiteCodes5379004:
		return "Small intestine muscularis propria"
	case ApproachSiteCodes5382009:
		return "Superior fascia of perineum"
	case ApproachSiteCodes5394000:
		return "Uterine paracervical lymph node"
	case ApproachSiteCodes5398002:
		return "Normal fat pad"
	case ApproachSiteCodes5403001:
		return "Articular process of third lumbar vertebra"
	case ApproachSiteCodes5421003:
		return "Sex chromosome Y"
	case ApproachSiteCodes5427004:
		return "Apocrine intraepidermal duct"
	case ApproachSiteCodes5458003:
		return "Deep artery of clitoris"
	case ApproachSiteCodes5459006:
		return "Cardiac incisure of stomach"
	case ApproachSiteCodes5491007:
		return "Lacrimal part of orbicularis oculi muscle"
	case ApproachSiteCodes5493005:
		return "Metacarpophalangeal joint of little finger"
	case ApproachSiteCodes5498001:
		return "Superior aberrant ductule of epididymis"
	case ApproachSiteCodes5520004:
		return "Subcutaneous tissue of chin"
	case ApproachSiteCodes5538001:
		return "Tegmental portion of pons"
	case ApproachSiteCodes5544002:
		return "Longitudinal layer of duodenal muscularis propria"
	case ApproachSiteCodes5560003:
		return "Alveolar ridge mucous membrane"
	case ApproachSiteCodes5564007:
		return "Singlet"
	case ApproachSiteCodes5574005:
		return "Seventh costal cartilage"
	case ApproachSiteCodes5580002:
		return "Tendon of supraspinatus muscle"
	case ApproachSiteCodes5597008:
		return "Retina of right eye"
	case ApproachSiteCodes5611001:
		return "Anulus fibrosus of intervertebral disc of fifth cervical vertebra"
	case ApproachSiteCodes5625000:
		return "Navicular facet of intermediate cuneiform bone"
	case ApproachSiteCodes5627008:
		return "Right visceral pleura"
	case ApproachSiteCodes5633004:
		return "Muscular portion of interventricular septum"
	case ApproachSiteCodes5643001:
		return "Canal of stomach"
	case ApproachSiteCodes5644007:
		return "Fractured membrane P face"
	case ApproachSiteCodes5653000:
		return "Entire inner surface of seventh rib"
	case ApproachSiteCodes5665001:
		return "Retina"
	case ApproachSiteCodes5668004:
		return "Lower digestive tract"
	case ApproachSiteCodes5682004:
		return "Subcutaneous tissue of upper extremity"
	case ApproachSiteCodes5696005:
		return "Entire articular part of tubercle of ninth rib"
	case ApproachSiteCodes5697001:
		return "Skin of lateral surface of finger"
	case ApproachSiteCodes5709001:
		return "Multifidus muscles"
	case ApproachSiteCodes5713008:
		return "Submandibular triangle"
	case ApproachSiteCodes5717009:
		return "Temporal fossa"
	case ApproachSiteCodes5718004:
		return "Tendon and tendon sheath of leg and ankle"
	case ApproachSiteCodes5727003:
		return "Anterior cervical lymph node"
	case ApproachSiteCodes5742000:
		return "Skin of forearm"
	case ApproachSiteCodes5751008:
		return "Subcutaneous tissue of anterior portion of neck"
	case ApproachSiteCodes5769004:
		return "Endocervical epithelium"
	case ApproachSiteCodes5780004:
		return "Paradidymis"
	case ApproachSiteCodes5798000:
		return "Diaphragm"
	case ApproachSiteCodes5802004:
		return "Medium sized neuron"
	case ApproachSiteCodes5814007:
		return "Entire angle of seventh rib"
	case ApproachSiteCodes5815008:
		return "Superior rectus muscle"
	case ApproachSiteCodes5816009:
		return "Duodenal fold"
	case ApproachSiteCodes5825003:
		return "Substantia propria of sclera"
	case ApproachSiteCodes5828001:
		return "Posterior cord of brachial plexus"
	case ApproachSiteCodes5847003:
		return "Superior articular process of seventh cervical vertebra"
	case ApproachSiteCodes5854009:
		return "Orbital plate of ethmoid bone"
	case ApproachSiteCodes5868002:
		return "Serosa of urinary bladder"
	case ApproachSiteCodes5872003:
		return "Subcutaneous tissue of lateral border of sole of foot"
	case ApproachSiteCodes5881009:
		return "Tuberosity of distal phalanx of hand"
	case ApproachSiteCodes5882002:
		return "Endothelial sieve plate"
	case ApproachSiteCodes5889006:
		return "Articular surface, third metacarpal, of fourth metacarpal bone"
	case ApproachSiteCodes5890002:
		return "Posterior cells of ethmoid sinus"
	case ApproachSiteCodes5893000:
		return "Superior recess of tympanic membrane"
	case ApproachSiteCodes5898009:
		return "Myotome"
	case ApproachSiteCodes5923009:
		return "Articular process of twelfth thoracic vertebra"
	case ApproachSiteCodes5926001:
		return "Bronchial lumen"
	case ApproachSiteCodes5928000:
		return "Great cardiac vein"
	case ApproachSiteCodes5942008:
		return "Tensor tympani muscle"
	case ApproachSiteCodes5943003:
		return "Vestibular vein"
	case ApproachSiteCodes5944009:
		return "Posterior palatine arch"
	case ApproachSiteCodes5948007:
		return "Capsule of distal interphalangeal joint of third toe"
	case ApproachSiteCodes5951000:
		return "Left wrist"
	case ApproachSiteCodes5953002:
		return "Eighth rib"
	case ApproachSiteCodes5976004:
		return "Subcutaneous tissue of eyelid"
	case ApproachSiteCodes5979006:
		return "Episcleral artery"
	case ApproachSiteCodes5996007:
		return "Chromosomes, group D"
	case ApproachSiteCodes6001004:
		return "Quadratus lumborum muscle"
	case ApproachSiteCodes6004007:
		return "Intervertebral disc of second thoracic vertebra"
	case ApproachSiteCodes6006009:
		return "Circular layer of duodenal muscularis propria"
	case ApproachSiteCodes6009002:
		return "Mesentery of ascending colon"
	case ApproachSiteCodes6014003:
		return "Penicilliary arteries"
	case ApproachSiteCodes6023000:
		return "Heterolysosome"
	case ApproachSiteCodes6032003:
		return "Columnar epithelial cell"
	case ApproachSiteCodes6046003:
		return "Entire outer surface of third rib"
	case ApproachSiteCodes6050005:
		return "Lacrimal vein"
	case ApproachSiteCodes6059006:
		return "Metacarpophalangeal joint of middle finger"
	case ApproachSiteCodes6062009:
		return "Deciduous mandibular right canine tooth"
	case ApproachSiteCodes6073002:
		return "Ligament of left superior vena cava"
	case ApproachSiteCodes6074008:
		return "Capsule of temporomandibular joint"
	case ApproachSiteCodes6076005:
		return "Gastrointestinal subserosa"
	case ApproachSiteCodes6104005:
		return "Subclavian nerve"
	case ApproachSiteCodes6105006:
		return "Body of fifth thoracic vertebra"
	case ApproachSiteCodes6110005:
		return "Facial nerve parasympathetic fibers"
	case ApproachSiteCodes6216007:
		return "Entire postcapillary venule"
	case ApproachSiteCodes6217003:
		return "Piriform recess"
	case ApproachSiteCodes6229007:
		return "Os lacrimale"
	case ApproachSiteCodes6253001:
		return "Sulcus terminalis cordis"
	case ApproachSiteCodes6268000:
		return "Accessory phrenic nerves"
	case ApproachSiteCodes6269008:
		return "Subcutaneous tissue of scalp"
	case ApproachSiteCodes6279005:
		return "Skin of dorsal surface of finger"
	case ApproachSiteCodes6317000:
		return "Posterior basal branch of left pulmonary artery"
	case ApproachSiteCodes6325003:
		return "Aryepiglottic muscle"
	case ApproachSiteCodes6326002:
		return "Structure of fetal atloid articulation"
	case ApproachSiteCodes6335009:
		return "Lymphoid follicle of stomach"
	case ApproachSiteCodes6359004:
		return "Hair medulla"
	case ApproachSiteCodes6371005:
		return "Lymphatics of thyroid gland"
	case ApproachSiteCodes6375001:
		return "Cavernous portion of urethra"
	case ApproachSiteCodes6392005:
		return "Coccygeal nerve"
	case ApproachSiteCodes6404004:
		return "Ligamentum nuchae"
	case ApproachSiteCodes6413002:
		return "Presymphysial lymph node"
	case ApproachSiteCodes6417001:
		return "Medial malleolus"
	case ApproachSiteCodes6423006:
		return "Supraspinatus muscle"
	case ApproachSiteCodes6424000:
		return "Structure of radiating portion of cortical lobule of kidney"
	case ApproachSiteCodes6445007:
		return "Mast cell"
	case ApproachSiteCodes6448009:
		return "Posterior vagal trunk"
	case ApproachSiteCodes6450001:
		return "Cytotrophoblast"
	case ApproachSiteCodes6472004:
		return "Medial aspect of ovary"
	case ApproachSiteCodes6504002:
		return "Glans clitoridis"
	case ApproachSiteCodes6511003:
		return "Distal portion of circumflex branch of left coronary artery"
	case ApproachSiteCodes6530003:
		return "Cardiac valve leaflet"
	case ApproachSiteCodes6533001:
		return "Colonic haustra"
	case ApproachSiteCodes6538005:
		return "Thyrocervical trunk"
	case ApproachSiteCodes6541001:
		return "Entire anterior commissure of mitral valve"
	case ApproachSiteCodes6544009:
		return "Gastrohepatic ligament"
	case ApproachSiteCodes6550004:
		return "Angular incisure of stomach"
	case ApproachSiteCodes6551000:
		return "Pollicis artery"
	case ApproachSiteCodes6553002:
		return "Inferior nasal turbinate"
	case ApproachSiteCodes6564004:
		return "Medial border of sole"
	case ApproachSiteCodes6566002:
		return "Cerebellar hemisphere"
	case ApproachSiteCodes6572002:
		return "Base of phalanx of middle finger"
	case ApproachSiteCodes6598008:
		return "Lingual nerve"
	case ApproachSiteCodes6606008:
		return "Structure of dorsal intercuneiform ligaments"
	case ApproachSiteCodes6608009:
		return "Sphenoparietal sinus"
	case ApproachSiteCodes6620001:
		return "Cuticle of nail"
	case ApproachSiteCodes6623004:
		return "Sternal muscle"
	case ApproachSiteCodes6633007:
		return "Entire right posterior cerebral artery"
	case ApproachSiteCodes6643005:
		return "Entire right anterior cerebral artery"
	case ApproachSiteCodes6646002:
		return "Anterior fossa of cranial cavity"
	case ApproachSiteCodes6649009:
		return "Uterine subserosa"
	case ApproachSiteCodes6651008:
		return "Central lobule of cerebellum"
	case ApproachSiteCodes6684008:
		return "Articular facet of head of fibula"
	case ApproachSiteCodes6685009:
		return "Right ankle"
	case ApproachSiteCodes6711001:
		return "Arch of second lumbar vertebra"
	case ApproachSiteCodes6720005:
		return "Femoral nerve lateral muscular branches"
	case ApproachSiteCodes6731002:
		return "Pleural recess"
	case ApproachSiteCodes6739000:
		return "Chorda tympani"
	case ApproachSiteCodes6742006:
		return "Entire callosomarginal branch of anterior cerebral artery"
	case ApproachSiteCodes6750002:
		return "Mitochondrial inclusion"
	case ApproachSiteCodes6757004:
		return "Right knee"
	case ApproachSiteCodes6787005:
		return "Structure of tendon and tendon sheath of hand"
	case ApproachSiteCodes6789008:
		return "Spermatozoa"
	case ApproachSiteCodes6799003:
		return "Macula of utricle"
	case ApproachSiteCodes6805009:
		return "Entire interstitial tissue of spleen"
	case ApproachSiteCodes6820003:
		return "Obturator nerve anterior branch"
	case ApproachSiteCodes6828005:
		return "Ligament of lumbosacral joint"
	case ApproachSiteCodes6829002:
		return "Pars ciliaris of retina"
	case ApproachSiteCodes6834003:
		return "Axial skeleton"
	case ApproachSiteCodes6841009:
		return "Corticomedullary junction of kidney"
	case ApproachSiteCodes6844001:
		return "Spore crystal"
	case ApproachSiteCodes6850006:
		return "Secondary foot process"
	case ApproachSiteCodes6864006:
		return "Leaf of epiglottis"
	case ApproachSiteCodes6866008:
		return "Habenular commissure"
	case ApproachSiteCodes6871001:
		return "Visceral pericardium"
	case ApproachSiteCodes6894000:
		return "Medial surface of arm"
	case ApproachSiteCodes6902008:
		return "Popliteal region"
	case ApproachSiteCodes6905005:
		return "Subcutaneous tissue of medial surface of third toe"
	case ApproachSiteCodes6912001:
		return "Lower alveolar ridge mucosa"
	case ApproachSiteCodes6914000:
		return "Perivascular space"
	case ApproachSiteCodes6921000:
		return "Right upper extremity"
	case ApproachSiteCodes6930008:
		return "Entire jugular arch"
	case ApproachSiteCodes6944002:
		return "Entire anterior labial vein"
	case ApproachSiteCodes6969002:
		return "Lymphocytic tissue"
	case ApproachSiteCodes6975006:
		return "Anterior myocardium"
	case ApproachSiteCodes6981003:
		return "Posterior hypothalamic nucleus"
	case ApproachSiteCodes6987004:
		return "Collateral sulcus"
	case ApproachSiteCodes6989001:
		return "Thoracolumbar region of back"
	case ApproachSiteCodes6991009:
		return "Subcutaneous tissue of jaw"
	case ApproachSiteCodes7035006:
		return "Bile duct mucous membrane"
	case ApproachSiteCodes7050002:
		return "Subcutaneous tissue of external genitalia"
	case ApproachSiteCodes7067009:
		return "Right colic artery"
	case ApproachSiteCodes7076002:
		return "Interstitial tissue of myocardium"
	case ApproachSiteCodes7083009:
		return "Middle phalanx of index finger"
	case ApproachSiteCodes7091000:
		return "Ventral posterolateral nucleus of thalamus"
	case ApproachSiteCodes7099003:
		return "Attachment plaque of desmosome or hemidesmosome"
	case ApproachSiteCodes7117004:
		return "Fetal implantation site"
	case ApproachSiteCodes7148007:
		return "Anulus fibrosus of intervertebral disc of thoracic vertebra"
	case ApproachSiteCodes7149004:
		return "False rib"
	case ApproachSiteCodes7154008:
		return "Trigeminal ganglion sensory root"
	case ApproachSiteCodes7160008:
		return "Base of metacarpal bone"
	case ApproachSiteCodes7167006:
		return "Paraduodenal recess"
	case ApproachSiteCodes7173007:
		return "Cauda equina"
	case ApproachSiteCodes7188002:
		return "Gustatory pore"
	case ApproachSiteCodes7192009:
		return "Isthmus tympani posticus"
	case ApproachSiteCodes7227003:
		return "Hypoglossal nerve intrinsic tongue muscle branch"
	case ApproachSiteCodes7234001:
		return "Entire inferior choroid vein"
	case ApproachSiteCodes7242000:
		return "Appendiceal muscularis propria"
	case ApproachSiteCodes7295002:
		return "Muscle of perineum"
	case ApproachSiteCodes7296001:
		return "Deep inguinal ring"
	case ApproachSiteCodes7311008:
		return "Anterior surface of arm"
	case ApproachSiteCodes7344002:
		return "Lingual gyrus"
	case ApproachSiteCodes7345001:
		return "Ciliary processes"
	case ApproachSiteCodes7362006:
		return "Lymphatic of head"
	case ApproachSiteCodes7376007:
		return "Entire left margin of uterus"
	case ApproachSiteCodes7378008:
		return "Paraventricular nucleus of thalamus"
	case ApproachSiteCodes7384006:
		return "Entire plantar calcaneocuboidal ligament"
	case ApproachSiteCodes7404008:
		return "Anterior semicircular duct"
	case ApproachSiteCodes7435002:
		return "Ovarian ligament"
	case ApproachSiteCodes7471001:
		return "Lateral surface of sublingual gland"
	case ApproachSiteCodes7477002:
		return "Lipid, crystalline"
	case ApproachSiteCodes7480001:
		return "Iliotibial tract"
	case ApproachSiteCodes7494000:
		return "Cerebellar lenticular nucleus"
	case ApproachSiteCodes7498002:
		return "Entire plantar tarsal ligament"
	case ApproachSiteCodes7507003:
		return "Entire anterior ligament of head of fibula"
	case ApproachSiteCodes7524009:
		return "Entire vasa vasorum"
	case ApproachSiteCodes7532001:
		return "Vagus nerve parasympathetic fibers"
	case ApproachSiteCodes7554004:
		return "Deep head of flexor pollicis brevis muscle"
	case ApproachSiteCodes7566005:
		return "Mitotic cell in anaphase"
	case ApproachSiteCodes7569003:
		return "Finger"
	case ApproachSiteCodes7591005:
		return "Intervertebral disc space of eleventh thoracic vertebra"
	case ApproachSiteCodes7597009:
		return "Subcutaneous tissue of vertex"
	case ApproachSiteCodes7605000:
		return "Connexon"
	case ApproachSiteCodes7610001:
		return "Tenth thoracic vertebra"
	case ApproachSiteCodes7629007:
		return "Thalamoolivary tract"
	case ApproachSiteCodes7651004:
		return "Intervenous tubercle of right atrium"
	case ApproachSiteCodes7652006:
		return "Frenulum labii"
	case ApproachSiteCodes7657000:
		return "Femoral artery"
	case ApproachSiteCodes7658005:
		return "Subtendinous bursa of triceps brachii muscle"
	case ApproachSiteCodes7697002:
		return "Pontine portion of medial longitudinal fasciculus"
	case ApproachSiteCodes7712004:
		return "Subdural space of spinal region"
	case ApproachSiteCodes7726008:
		return "Skin of medial surface of fifth toe"
	case ApproachSiteCodes7736000:
		return "Posterior choroidal artery"
	case ApproachSiteCodes7742001:
		return "Palatine duct"
	case ApproachSiteCodes7748002:
		return "Skin appendage"
	case ApproachSiteCodes7755000:
		return "Mesovarian margin of ovary"
	case ApproachSiteCodes7756004:
		return "Lamina of third thoracic vertebra"
	case ApproachSiteCodes7764005:
		return "Entire striate artery"
	case ApproachSiteCodes7769000:
		return "Right foot"
	case ApproachSiteCodes7783003:
		return "Sympathetic trunk spinal nerve branch"
	case ApproachSiteCodes7820009:
		return "Lateral posterior nucleus of thalamus"
	case ApproachSiteCodes7829005:
		return "Anterior surface of manubrium"
	case ApproachSiteCodes7832008:
		return "Abdominal aorta"
	case ApproachSiteCodes7835005:
		return "Posterior margin of nasal septum"
	case ApproachSiteCodes7840002:
		return "Subcutaneous tissue of submental area"
	case ApproachSiteCodes7841003:
		return "Macrocytic normochromic erythrocyte"
	case ApproachSiteCodes7844006:
		return "Sternoclavicular joint"
	case ApproachSiteCodes7851002:
		return "Intracranial subdural space"
	case ApproachSiteCodes7854005:
		return "Mandibular canal"
	case ApproachSiteCodes7872004:
		return "Myocardium of ventricle"
	case ApproachSiteCodes7874003:
		return "Scapular region of back"
	case ApproachSiteCodes7880006:
		return "Rhopheocytotic vesicle"
	case ApproachSiteCodes7884002:
		return "Corneal corpuscle"
	case ApproachSiteCodes7885001:
		return "Rotator cuff including muscles and tendons"
	case ApproachSiteCodes7892006:
		return "Submucosa of anal canal"
	case ApproachSiteCodes7896009:
		return "Occipital angle of parietal bone"
	case ApproachSiteCodes7911004:
		return "Olivocerebellar fibers"
	case ApproachSiteCodes7925003:
		return "Proximal phalanx of third toe"
	case ApproachSiteCodes7936005:
		return "Ligament of diaphragm"
	case ApproachSiteCodes7944005:
		return "Helper cell"
	case ApproachSiteCodes7954009:
		return "Lamina propria of ethmoid sinus"
	case ApproachSiteCodes7967007:
		return "Entire first left aortic arch"
	case ApproachSiteCodes7986004:
		return "Abdominopelvic portion of sympathetic nervous system"
	case ApproachSiteCodes7991003:
		return "Skin of glans penis"
	case ApproachSiteCodes7999001:
		return "Articulations of auditory ossicles"
	case ApproachSiteCodes8001006:
		return "Mucous membrane of tongue"
	case ApproachSiteCodes8012006:
		return "Anterior communicating artery"
	case ApproachSiteCodes8017000:
		return "Inflow tract of right ventricle"
	case ApproachSiteCodes8024004:
		return "Limitans nucleus"
	case ApproachSiteCodes8039003:
		return "Subcutaneous acromial bursa"
	case ApproachSiteCodes8040001:
		return "Superficial flexor tendon of little finger"
	case ApproachSiteCodes8045006:
		return "Membrane-coating granule, amorphous"
	case ApproachSiteCodes8057002:
		return "Lateral nuclei of globus pallidus"
	case ApproachSiteCodes8059004:
		return "Pancreatic veins"
	case ApproachSiteCodes8067007:
		return "Entire superficial circumflex iliac vein"
	case ApproachSiteCodes8068002:
		return "Stratum lemnisci of corpora quadrigemina"
	case ApproachSiteCodes8079007:
		return "Radial nerve"
	case ApproachSiteCodes8091003:
		return "Intervertebral disc space of twelfth thoracic vertebra"
	case ApproachSiteCodes8100009:
		return "Infundibulum of Fallopian tube"
	case ApproachSiteCodes8111001:
		return "Intranuclear crystal"
	case ApproachSiteCodes8112008:
		return "Hindgut"
	case ApproachSiteCodes8119004:
		return "Entire delphian lymph node"
	case ApproachSiteCodes8128003:
		return "Supraaortic valve area"
	case ApproachSiteCodes8133004:
		return "Superior anastomotic vein"
	case ApproachSiteCodes8157004:
		return "Vein of head"
	case ApproachSiteCodes8158009:
		return "Interlobar duct of pancreas"
	case ApproachSiteCodes8159001:
		return "Superior colliculus of corpora quadrigemina"
	case ApproachSiteCodes8160006:
		return "Entire lateral striate artery"
	case ApproachSiteCodes8161005:
		return "Infraorbital nerve"
	case ApproachSiteCodes8165001:
		return "Superior articular process of fifth thoracic vertebra"
	case ApproachSiteCodes8205005:
		return "Wrist"
	case ApproachSiteCodes8225009:
		return "Accessory atrioventricular bundle"
	case ApproachSiteCodes8242003:
		return "Apical branch of right pulmonary artery"
	case ApproachSiteCodes8251006:
		return "Osseous portion of Eustachian tube"
	case ApproachSiteCodes8264007:
		return "Tunica interna of eyeball"
	case ApproachSiteCodes8265008:
		return "Articular surface, metacarpal, of phalanx of hand"
	case ApproachSiteCodes8266009:
		return "Small intestine serosa"
	case ApproachSiteCodes8289001:
		return "Below knee region"
	case ApproachSiteCodes8292002:
		return "Interlobular arteries of liver"
	case ApproachSiteCodes8314003:
		return "Mastoid fontanel of skull"
	case ApproachSiteCodes8334002:
		return "Lumbar lymph node"
	case ApproachSiteCodes8356004:
		return "Colic lymph node"
	case ApproachSiteCodes8369000:
		return "Sphincter pupillae muscle"
	case ApproachSiteCodes8373002:
		return "Jugum of sphenoid bone"
	case ApproachSiteCodes8387002:
		return "Lamina of eighth thoracic vertebra"
	case ApproachSiteCodes8389004:
		return "Birth canal"
	case ApproachSiteCodes8412003:
		return "Iliac fossa"
	case ApproachSiteCodes8415001:
		return "Renal surface of adrenal gland"
	case ApproachSiteCodes8454000:
		return "Joint of lumbar vertebra"
	case ApproachSiteCodes8464009:
		return "Ligament of sacroiliac joint and pubic symphysis"
	case ApproachSiteCodes8482007:
		return "Sinoatrial node branch of right coronary artery"
	case ApproachSiteCodes8483002:
		return "Mesial surface of tooth"
	case ApproachSiteCodes8496001:
		return "Obliquus capitis muscle"
	case ApproachSiteCodes8523001:
		return "Inferior articular process of twelfth thoracic vertebra"
	case ApproachSiteCodes8546004:
		return "Posterior intercavernous sinus"
	case ApproachSiteCodes8556000:
		return "Lipid droplet"
	case ApproachSiteCodes8559007:
		return "Entire juxtaintestinal lymph node"
	case ApproachSiteCodes8560002:
		return "Interclavicular ligament"
	case ApproachSiteCodes8580001:
		return "Both feet"
	case ApproachSiteCodes8595004:
		return "Meissner's plexus"
	case ApproachSiteCodes8598002:
		return "Vestibulocochlear nerve"
	case ApproachSiteCodes8600008:
		return "Cricoid cartilage"
	case ApproachSiteCodes8603005:
		return "Adductor hallucis muscle"
	case ApproachSiteCodes8604004:
		return "Medulla oblongata fasciculus cuneatus"
	case ApproachSiteCodes8608001:
		return "Right margin of heart"
	case ApproachSiteCodes8617001:
		return "Zygomatic region of face"
	case ApproachSiteCodes8623006:
		return "Transplanted ureter"
	case ApproachSiteCodes8629005:
		return "Superior right pulmonary vein"
	case ApproachSiteCodes8640002:
		return "Choroidal branches of posterior spinal artery"
	case ApproachSiteCodes8668003:
		return "Glycogen vacuole"
	case ApproachSiteCodes8671006:
		return "All toes"
	case ApproachSiteCodes8677005:
		return "Body of right atrium"
	case ApproachSiteCodes8688004:
		return "Lateral olfactory gyrus"
	case ApproachSiteCodes8695008:
		return "Intervertebral foramen of second lumbar vertebra"
	case ApproachSiteCodes8710005:
		return "Entire minor sublingual ducts"
	case ApproachSiteCodes8711009:
		return "Periodontal tissues"
	case ApproachSiteCodes8714001:
		return "Subcutaneous tissue of interdigital space of hand"
	case ApproachSiteCodes8752000:
		return "Cavernous portion of internal carotid artery"
	case ApproachSiteCodes8775007:
		return "Tendinous arch"
	case ApproachSiteCodes8784007:
		return "Intranuclear body, granular with filamentous capsule"
	case ApproachSiteCodes8810002:
		return "Corticomedullary junction of adrenal gland"
	case ApproachSiteCodes8814006:
		return "Iliac tuberosity"
	case ApproachSiteCodes8815007:
		return "Thenar and hypothenar spaces"
	case ApproachSiteCodes8820007:
		return "Pedicle of eleventh thoracic vertebra"
	case ApproachSiteCodes8821006:
		return "Peroneal artery"
	case ApproachSiteCodes8827005:
		return "Shaft of phalanx of middle finger"
	case ApproachSiteCodes8839002:
		return "Agranular endoplasmic reticulum, connection with other organelle"
	case ApproachSiteCodes8845005:
		return "Entire subtendinous prepatellar bursa"
	case ApproachSiteCodes8850004:
		return "Proper fasciculus"
	case ApproachSiteCodes8854008:
		return "Crista galli"
	case ApproachSiteCodes8862000:
		return "Palmar surface of middle finger"
	case ApproachSiteCodes8873007:
		return "Mandibular right second premolar tooth"
	case ApproachSiteCodes8887007:
		return "Brachiocephalic vein"
	case ApproachSiteCodes8892009:
		return "Diaphragmatic surface of lung"
	case ApproachSiteCodes8894005:
		return "Entire gastric cardiac gland (body structure)"
	case ApproachSiteCodes8897003:
		return "Lateral glossoepiglottic fold"
	case ApproachSiteCodes8907008:
		return "Entire left ulnar artery"
	case ApproachSiteCodes8910001:
		return "Entire inferior transverse scapular ligament"
	case ApproachSiteCodes8911002:
		return "Entire endocardium of right ventricle"
	case ApproachSiteCodes8928004:
		return "Inguinal lymph node"
	case ApproachSiteCodes8931003:
		return "Coracoid process of scapula"
	case ApproachSiteCodes8935007:
		return "Cerebral meninges"
	case ApproachSiteCodes8942007:
		return "Trapezoid ligament"
	case ApproachSiteCodes8965002:
		return "Stratum zonale of corpora quadrigemina"
	case ApproachSiteCodes8966001:
		return "Left eye"
	case ApproachSiteCodes8983005:
		return "Joint of vertebral column"
	case ApproachSiteCodes8988001:
		return "Marginal part of orbicularis oris muscle"
	case ApproachSiteCodes8993003:
		return "Hepatic vein"
	case ApproachSiteCodes9000002:
		return "Cerebellar peduncle"
	case ApproachSiteCodes9003000:
		return "Left parietal lobe"
	case ApproachSiteCodes9018004:
		return "Entire middle colic vein"
	case ApproachSiteCodes9040008:
		return "Ascending colon"
	case ApproachSiteCodes9055004:
		return "Both forearms"
	case ApproachSiteCodes9073001:
		return "White matter of insula"
	case ApproachSiteCodes9081000:
		return "Splenic sinusoid"
	case ApproachSiteCodes9086005:
		return "Superior laryngeal vein"
	case ApproachSiteCodes9089003:
		return "Arch of foot"
	case ApproachSiteCodes9108007:
		return "Vein of the scala tympani"
	case ApproachSiteCodes9127001:
		return "Transverse folds of palate"
	case ApproachSiteCodes9156001:
		return "Embryo stage 1 structure"
	case ApproachSiteCodes9185007:
		return "Capsule of metatarsophalangeal joint of fifth toe"
	case ApproachSiteCodes9186008:
		return "Filaments of contractile apparatus"
	case ApproachSiteCodes9188009:
		return "Intervertebral disc of eighth thoracic vertebra"
	case ApproachSiteCodes9208002:
		return "Centriole"
	case ApproachSiteCodes9212008:
		return "Shaft of fifth metatarsal bone"
	case ApproachSiteCodes9229006:
		return "Structure of lumbar rotator muscle (body structure)"
	case ApproachSiteCodes9231002:
		return "Entire external pudendal vein"
	case ApproachSiteCodes9240003:
		return "Niemann-Pick cell"
	case ApproachSiteCodes9242006:
		return "Posterior segment of right lobe of liver"
	case ApproachSiteCodes9258009:
		return "Gravid uterus"
	case ApproachSiteCodes9261005:
		return "Tendon and tendon sheath of second toe"
	case ApproachSiteCodes9262003:
		return "Pelvic fascia"
	case ApproachSiteCodes9284003:
		return "Corpus cavernosum of penis"
	case ApproachSiteCodes9290004:
		return "Posterior intraoccipital synchondrosis"
	case ApproachSiteCodes9305001:
		return "Labial veins"
	case ApproachSiteCodes9317004:
		return "Merkel's tactile disc"
	case ApproachSiteCodes9320007:
		return "Subtendinous iliac bursa"
	case ApproachSiteCodes9321006:
		return "Tail of epididymis"
	case ApproachSiteCodes9325002:
		return "Interdental papilla of gingiva"
	case ApproachSiteCodes9332006:
		return "Lateral ligament of temporomandibular joint"
	case ApproachSiteCodes9348007:
		return "Skin of medial surface of middle finger"
	case ApproachSiteCodes9379006:
		return "All permanent teeth"
	case ApproachSiteCodes9380009:
		return "Pecten ani"
	case ApproachSiteCodes9384000:
		return "Lumbar vein"
	case ApproachSiteCodes9390001:
		return "Lymphatics of stomach"
	case ApproachSiteCodes9432007:
		return "Plantar surface of fourth toe"
	case ApproachSiteCodes9438006:
		return "Structure of deep cervical lymphatics"
	case ApproachSiteCodes9454009:
		return "Subclavian vein"
	case ApproachSiteCodes9455005:
		return "Structure of medial cartilaginous lamina of pharyngotympanic tube (body structure)"
	case ApproachSiteCodes9475001:
		return "Structure of amacrine cell of retina (body structure)"
	case ApproachSiteCodes9481009:
		return "Afferent glomerular arteriole"
	case ApproachSiteCodes9490002:
		return "Pulmonary ligament"
	case ApproachSiteCodes9498009:
		return "Head of metacarpal bone"
	case ApproachSiteCodes9502002:
		return "Coronal depression of tooth"
	case ApproachSiteCodes9512009:
		return "Calcaneocuboidal ligament"
	case ApproachSiteCodes9535007:
		return "Pyramid of medulla oblongata"
	case ApproachSiteCodes9558005:
		return "Entire facet for fifth costal cartilage of sternum"
	case ApproachSiteCodes9566001:
		return "Duodenal lumen"
	case ApproachSiteCodes9568000:
		return "Subcutaneous tissue of areola"
	case ApproachSiteCodes9596006:
		return "Deep branch of ulnar nerve"
	case ApproachSiteCodes9609000:
		return "Posterior process of nasal septal cartilage"
	case ApproachSiteCodes9625005:
		return "Lanugo hair"
	case ApproachSiteCodes9642004:
		return "Left superior vena cava"
	case ApproachSiteCodes9646001:
		return "Entire superior transverse scapular ligament"
	case ApproachSiteCodes9654004:
		return "Gastric mucous gland"
	case ApproachSiteCodes9659009:
		return "Infraclavicular lymph node"
	case ApproachSiteCodes9662007:
		return "Subcutaneous tissue of lower margin of nasal septum"
	case ApproachSiteCodes9668006:
		return "Ciliary muscle"
	case ApproachSiteCodes9677004:
		return "Head of second metatarsal bone"
	case ApproachSiteCodes9683001:
		return "Melanocyte"
	case ApproachSiteCodes9684007:
		return "Posterior scrotal branches of internal pudendal artery"
	case ApproachSiteCodes9708001:
		return "Iliac fascia"
	case ApproachSiteCodes9736006:
		return "Right wrist"
	case ApproachSiteCodes9743000:
		return "Entire tendon of index finger"
	case ApproachSiteCodes9758008:
		return "Submucosa of tonsil"
	case ApproachSiteCodes9770007:
		return "Genital tubercle"
	case ApproachSiteCodes9775002:
		return "Left carotid sinus"
	case ApproachSiteCodes9779008:
		return "Distinctive shape of mitochondrial cristae"
	case ApproachSiteCodes9783008:
		return "Superficial lymphatics of thorax"
	case ApproachSiteCodes9791004:
		return "Deep venous system of lower extremity"
	case ApproachSiteCodes9796009:
		return "Skeletal muscle fiber, type IIb"
	case ApproachSiteCodes9813009:
		return "Fascia of upper extremity"
	case ApproachSiteCodes9825007:
		return "Proximal phalanx of little toe"
	case ApproachSiteCodes9837009:
		return "Perforating branches of internal thoracic artery"
	case ApproachSiteCodes9840009:
		return "Biparietal diameter of head"
	case ApproachSiteCodes9841008:
		return "Interspinalis thoracis muscles"
	case ApproachSiteCodes9846003:
		return "Right kidney"
	case ApproachSiteCodes9847007:
		return "Hilum of adrenal gland"
	case ApproachSiteCodes9849005:
		return "Fornix of lacrimal sac"
	case ApproachSiteCodes9870004:
		return "Carunculae hymenales"
	case ApproachSiteCodes9875009:
		return "Thymus"
	case ApproachSiteCodes9878006:
		return "Entire appendicular vein"
	case ApproachSiteCodes9880000:
		return "Thyroid tubercle"
	case ApproachSiteCodes9881001:
		return "Peripheral nerve myelinated nerve fiber"
	case ApproachSiteCodes9891007:
		return "Transverse arytenoid muscle"
	case ApproachSiteCodes9898001:
		return "Paracentral lobule"
	case ApproachSiteCodes9951005:
		return "Posterior ethmoidal nerve"
	case ApproachSiteCodes9968009:
		return "Primary foot process"
	case ApproachSiteCodes9970000:
		return "Ileocecal ostium"
	case ApproachSiteCodes9976006:
		return "Rhomboideus cervicis muscle"
	case ApproachSiteCodes9994000:
		return "Superior articular process of sixth thoracic vertebra"
	case ApproachSiteCodes9999005:
		return "Duodenal ampulla"
	case ApproachSiteCodes10013000:
		return "Lateral meniscus of knee joint"
	case ApproachSiteCodes10024003:
		return "Base of lung"
	case ApproachSiteCodes10025002:
		return "Base of phalanx of index finger"
	case ApproachSiteCodes10026001:
		return "Ventral spinocerebellar tract of pons"
	case ApproachSiteCodes10036009:
		return "Nucleus pulposus of intervertebral disc of eighth thoracic vertebra"
	case ApproachSiteCodes10042008:
		return "Intervertebral foramen of fifth thoracic vertebra"
	case ApproachSiteCodes10047002:
		return "Transplanted lung"
	case ApproachSiteCodes10052007:
		return "Male"
	case ApproachSiteCodes10056005:
		return "Ophthalmic nerve"
	case ApproachSiteCodes10062000:
		return "Levator labii superioris muscle"
	case ApproachSiteCodes10119003:
		return "Deep volar arch of radial artery"
	case ApproachSiteCodes10124000:
		return "Deep dorsal sacrococcygeal ligament"
	case ApproachSiteCodes10134009:
		return "Medial surface of third toe"
	case ApproachSiteCodes10141003:
		return "Zygomatic nerve"
	case ApproachSiteCodes10145007:
		return "Vagus nerve pharyngeal plexus"
	case ApproachSiteCodes10148009:
		return "Entire costal groove of fifth rib"
	case ApproachSiteCodes10149001:
		return "Merkel's cell"
	case ApproachSiteCodes10151002:
		return "Middle part of cartilaginous portion of auditory canal"
	case ApproachSiteCodes10167008:
		return "Entire inner surface of eleventh rib"
	case ApproachSiteCodes10176001:
		return "Supraoptic region of hypothalamus"
	case ApproachSiteCodes10200004:
		return "Liver"
	case ApproachSiteCodes10208006:
		return "Gynecoid pelvis"
	case ApproachSiteCodes10209003:
		return "Parotid lymph node"
	case ApproachSiteCodes10245000:
		return "Abductor digiti minimi muscle of foot"
	case ApproachSiteCodes10271001:
		return "Entire anterior commissure of aortic valve"
	case ApproachSiteCodes10293006:
		return "Iliac artery"
	case ApproachSiteCodes10296003:
		return "Uropod"
	case ApproachSiteCodes10299005:
		return "Entire anteromedial perforating artery"
	case ApproachSiteCodes10328008:
		return "Capsule of calcaneocuboidal joint"
	case ApproachSiteCodes10339006:
		return "Anulus fibrosus of intervertebral disc of tenth thoracic vertebra"
	case ApproachSiteCodes10410005:
		return "Sweat gland"
	case ApproachSiteCodes10415000:
		return "Pigmented layer of ciliary body"
	case ApproachSiteCodes10417008:
		return "Lesser sciatic notch"
	case ApproachSiteCodes10418003:
		return "Pulmonic area"
	default:
		return "Unknown Approach Site Codes"
	}
}
