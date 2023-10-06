package code

import (
	"fmt"
	"strings"
)

type AllergyIntoleranceCode string

const (
	AllergyIntoleranceCode105590001  AllergyIntoleranceCode = "105590001"
	AllergyIntoleranceCode102002     AllergyIntoleranceCode = "102002"
	AllergyIntoleranceCode120006     AllergyIntoleranceCode = "120006"
	AllergyIntoleranceCode125001     AllergyIntoleranceCode = "125001"
	AllergyIntoleranceCode126000     AllergyIntoleranceCode = "126000"
	AllergyIntoleranceCode130002     AllergyIntoleranceCode = "130002"
	AllergyIntoleranceCode131003     AllergyIntoleranceCode = "131003"
	AllergyIntoleranceCode159002     AllergyIntoleranceCode = "159002"
	AllergyIntoleranceCode164003     AllergyIntoleranceCode = "164003"
	AllergyIntoleranceCode178002     AllergyIntoleranceCode = "178002"
	AllergyIntoleranceCode186002     AllergyIntoleranceCode = "186002"
	AllergyIntoleranceCode187006     AllergyIntoleranceCode = "187006"
	AllergyIntoleranceCode200001     AllergyIntoleranceCode = "200001"
	AllergyIntoleranceCode217008     AllergyIntoleranceCode = "217008"
	AllergyIntoleranceCode231008     AllergyIntoleranceCode = "231008"
	AllergyIntoleranceCode238002     AllergyIntoleranceCode = "238002"
	AllergyIntoleranceCode261000     AllergyIntoleranceCode = "261000"
	AllergyIntoleranceCode296000     AllergyIntoleranceCode = "296000"
	AllergyIntoleranceCode322006     AllergyIntoleranceCode = "322006"
	AllergyIntoleranceCode327000     AllergyIntoleranceCode = "327000"
	AllergyIntoleranceCode329002     AllergyIntoleranceCode = "329002"
	AllergyIntoleranceCode336001     AllergyIntoleranceCode = "336001"
	AllergyIntoleranceCode340005     AllergyIntoleranceCode = "340005"
	AllergyIntoleranceCode363000     AllergyIntoleranceCode = "363000"
	AllergyIntoleranceCode370000     AllergyIntoleranceCode = "370000"
	AllergyIntoleranceCode371001     AllergyIntoleranceCode = "371001"
	AllergyIntoleranceCode377002     AllergyIntoleranceCode = "377002"
	AllergyIntoleranceCode392001     AllergyIntoleranceCode = "392001"
	AllergyIntoleranceCode395004     AllergyIntoleranceCode = "395004"
	AllergyIntoleranceCode412004     AllergyIntoleranceCode = "412004"
	AllergyIntoleranceCode424006     AllergyIntoleranceCode = "424006"
	AllergyIntoleranceCode425007     AllergyIntoleranceCode = "425007"
	AllergyIntoleranceCode432003     AllergyIntoleranceCode = "432003"
	AllergyIntoleranceCode438004     AllergyIntoleranceCode = "438004"
	AllergyIntoleranceCode462009     AllergyIntoleranceCode = "462009"
	AllergyIntoleranceCode472007     AllergyIntoleranceCode = "472007"
	AllergyIntoleranceCode476005     AllergyIntoleranceCode = "476005"
	AllergyIntoleranceCode498001     AllergyIntoleranceCode = "498001"
	AllergyIntoleranceCode501001     AllergyIntoleranceCode = "501001"
	AllergyIntoleranceCode505005     AllergyIntoleranceCode = "505005"
	AllergyIntoleranceCode506006     AllergyIntoleranceCode = "506006"
	AllergyIntoleranceCode515004     AllergyIntoleranceCode = "515004"
	AllergyIntoleranceCode519005     AllergyIntoleranceCode = "519005"
	AllergyIntoleranceCode521000     AllergyIntoleranceCode = "521000"
	AllergyIntoleranceCode529003     AllergyIntoleranceCode = "529003"
	AllergyIntoleranceCode538001     AllergyIntoleranceCode = "538001"
	AllergyIntoleranceCode566009     AllergyIntoleranceCode = "566009"
	AllergyIntoleranceCode576007     AllergyIntoleranceCode = "576007"
	AllergyIntoleranceCode578008     AllergyIntoleranceCode = "578008"
	AllergyIntoleranceCode584006     AllergyIntoleranceCode = "584006"
	AllergyIntoleranceCode585007     AllergyIntoleranceCode = "585007"
	AllergyIntoleranceCode591009     AllergyIntoleranceCode = "591009"
	AllergyIntoleranceCode593007     AllergyIntoleranceCode = "593007"
	AllergyIntoleranceCode594001     AllergyIntoleranceCode = "594001"
	AllergyIntoleranceCode597008     AllergyIntoleranceCode = "597008"
	AllergyIntoleranceCode604000     AllergyIntoleranceCode = "604000"
	AllergyIntoleranceCode611001     AllergyIntoleranceCode = "611001"
	AllergyIntoleranceCode620005     AllergyIntoleranceCode = "620005"
	AllergyIntoleranceCode648005     AllergyIntoleranceCode = "648005"
	AllergyIntoleranceCode662003     AllergyIntoleranceCode = "662003"
	AllergyIntoleranceCode668004     AllergyIntoleranceCode = "668004"
	AllergyIntoleranceCode683009     AllergyIntoleranceCode = "683009"
	AllergyIntoleranceCode686001     AllergyIntoleranceCode = "686001"
	AllergyIntoleranceCode693002     AllergyIntoleranceCode = "693002"
	AllergyIntoleranceCode698006     AllergyIntoleranceCode = "698006"
	AllergyIntoleranceCode699003     AllergyIntoleranceCode = "699003"
	AllergyIntoleranceCode704006     AllergyIntoleranceCode = "704006"
	AllergyIntoleranceCode732002     AllergyIntoleranceCode = "732002"
	AllergyIntoleranceCode735000     AllergyIntoleranceCode = "735000"
	AllergyIntoleranceCode747006     AllergyIntoleranceCode = "747006"
	AllergyIntoleranceCode773001     AllergyIntoleranceCode = "773001"
	AllergyIntoleranceCode785009     AllergyIntoleranceCode = "785009"
	AllergyIntoleranceCode804003     AllergyIntoleranceCode = "804003"
	AllergyIntoleranceCode819002     AllergyIntoleranceCode = "819002"
	AllergyIntoleranceCode850000     AllergyIntoleranceCode = "850000"
	AllergyIntoleranceCode859004     AllergyIntoleranceCode = "859004"
	AllergyIntoleranceCode860009     AllergyIntoleranceCode = "860009"
	AllergyIntoleranceCode873008     AllergyIntoleranceCode = "873008"
	AllergyIntoleranceCode876000     AllergyIntoleranceCode = "876000"
	AllergyIntoleranceCode877009     AllergyIntoleranceCode = "877009"
	AllergyIntoleranceCode889006     AllergyIntoleranceCode = "889006"
	AllergyIntoleranceCode896008     AllergyIntoleranceCode = "896008"
	AllergyIntoleranceCode905001     AllergyIntoleranceCode = "905001"
	AllergyIntoleranceCode923009     AllergyIntoleranceCode = "923009"
	AllergyIntoleranceCode925002     AllergyIntoleranceCode = "925002"
	AllergyIntoleranceCode963005     AllergyIntoleranceCode = "963005"
	AllergyIntoleranceCode974001     AllergyIntoleranceCode = "974001"
	AllergyIntoleranceCode979006     AllergyIntoleranceCode = "979006"
	AllergyIntoleranceCode993004     AllergyIntoleranceCode = "993004"
	AllergyIntoleranceCode1002007    AllergyIntoleranceCode = "1002007"
	AllergyIntoleranceCode1010008    AllergyIntoleranceCode = "1010008"
	AllergyIntoleranceCode1018001    AllergyIntoleranceCode = "1018001"
	AllergyIntoleranceCode1025008    AllergyIntoleranceCode = "1025008"
	AllergyIntoleranceCode1047008    AllergyIntoleranceCode = "1047008"
	AllergyIntoleranceCode1050006    AllergyIntoleranceCode = "1050006"
	AllergyIntoleranceCode1065007    AllergyIntoleranceCode = "1065007"
	AllergyIntoleranceCode1080001    AllergyIntoleranceCode = "1080001"
	AllergyIntoleranceCode1091008    AllergyIntoleranceCode = "1091008"
	AllergyIntoleranceCode1097007    AllergyIntoleranceCode = "1097007"
	AllergyIntoleranceCode1105007    AllergyIntoleranceCode = "1105007"
	AllergyIntoleranceCode1113008    AllergyIntoleranceCode = "1113008"
	AllergyIntoleranceCode1137008    AllergyIntoleranceCode = "1137008"
	AllergyIntoleranceCode1149009    AllergyIntoleranceCode = "1149009"
	AllergyIntoleranceCode1160000    AllergyIntoleranceCode = "1160000"
	AllergyIntoleranceCode1166006    AllergyIntoleranceCode = "1166006"
	AllergyIntoleranceCode1169004    AllergyIntoleranceCode = "1169004"
	AllergyIntoleranceCode1171004    AllergyIntoleranceCode = "1171004"
	AllergyIntoleranceCode1185009    AllergyIntoleranceCode = "1185009"
	AllergyIntoleranceCode1189003    AllergyIntoleranceCode = "1189003"
	AllergyIntoleranceCode1190007    AllergyIntoleranceCode = "1190007"
	AllergyIntoleranceCode1219001    AllergyIntoleranceCode = "1219001"
	AllergyIntoleranceCode1223009    AllergyIntoleranceCode = "1223009"
	AllergyIntoleranceCode1244009    AllergyIntoleranceCode = "1244009"
	AllergyIntoleranceCode1248007    AllergyIntoleranceCode = "1248007"
	AllergyIntoleranceCode1269009    AllergyIntoleranceCode = "1269009"
	AllergyIntoleranceCode1272002    AllergyIntoleranceCode = "1272002"
	AllergyIntoleranceCode1273007    AllergyIntoleranceCode = "1273007"
	AllergyIntoleranceCode1313002    AllergyIntoleranceCode = "1313002"
	AllergyIntoleranceCode1319003    AllergyIntoleranceCode = "1319003"
	AllergyIntoleranceCode1320009    AllergyIntoleranceCode = "1320009"
	AllergyIntoleranceCode1325004    AllergyIntoleranceCode = "1325004"
	AllergyIntoleranceCode1331001    AllergyIntoleranceCode = "1331001"
	AllergyIntoleranceCode1336006    AllergyIntoleranceCode = "1336006"
	AllergyIntoleranceCode1341003    AllergyIntoleranceCode = "1341003"
	AllergyIntoleranceCode1346008    AllergyIntoleranceCode = "1346008"
	AllergyIntoleranceCode1355006    AllergyIntoleranceCode = "1355006"
	AllergyIntoleranceCode1368003    AllergyIntoleranceCode = "1368003"
	AllergyIntoleranceCode1371006    AllergyIntoleranceCode = "1371006"
	AllergyIntoleranceCode1373009    AllergyIntoleranceCode = "1373009"
	AllergyIntoleranceCode1381005    AllergyIntoleranceCode = "1381005"
	AllergyIntoleranceCode1394007    AllergyIntoleranceCode = "1394007"
	AllergyIntoleranceCode1396009    AllergyIntoleranceCode = "1396009"
	AllergyIntoleranceCode1405004    AllergyIntoleranceCode = "1405004"
	AllergyIntoleranceCode1408002    AllergyIntoleranceCode = "1408002"
	AllergyIntoleranceCode1416006    AllergyIntoleranceCode = "1416006"
	AllergyIntoleranceCode1450002    AllergyIntoleranceCode = "1450002"
	AllergyIntoleranceCode1466000    AllergyIntoleranceCode = "1466000"
	AllergyIntoleranceCode1471007    AllergyIntoleranceCode = "1471007"
	AllergyIntoleranceCode1472000    AllergyIntoleranceCode = "1472000"
	AllergyIntoleranceCode1476002    AllergyIntoleranceCode = "1476002"
	AllergyIntoleranceCode1477006    AllergyIntoleranceCode = "1477006"
	AllergyIntoleranceCode1496005    AllergyIntoleranceCode = "1496005"
	AllergyIntoleranceCode1506001    AllergyIntoleranceCode = "1506001"
	AllergyIntoleranceCode1517000    AllergyIntoleranceCode = "1517000"
	AllergyIntoleranceCode1530004    AllergyIntoleranceCode = "1530004"
	AllergyIntoleranceCode1535009    AllergyIntoleranceCode = "1535009"
	AllergyIntoleranceCode1536005    AllergyIntoleranceCode = "1536005"
	AllergyIntoleranceCode1540001    AllergyIntoleranceCode = "1540001"
	AllergyIntoleranceCode1545006    AllergyIntoleranceCode = "1545006"
	AllergyIntoleranceCode1557002    AllergyIntoleranceCode = "1557002"
	AllergyIntoleranceCode1565004    AllergyIntoleranceCode = "1565004"
	AllergyIntoleranceCode1575001    AllergyIntoleranceCode = "1575001"
	AllergyIntoleranceCode1603001    AllergyIntoleranceCode = "1603001"
	AllergyIntoleranceCode1607000    AllergyIntoleranceCode = "1607000"
	AllergyIntoleranceCode1609002    AllergyIntoleranceCode = "1609002"
	AllergyIntoleranceCode1634002    AllergyIntoleranceCode = "1634002"
	AllergyIntoleranceCode1649005    AllergyIntoleranceCode = "1649005"
	AllergyIntoleranceCode1656004    AllergyIntoleranceCode = "1656004"
	AllergyIntoleranceCode1660001    AllergyIntoleranceCode = "1660001"
	AllergyIntoleranceCode1668008    AllergyIntoleranceCode = "1668008"
	AllergyIntoleranceCode1672007    AllergyIntoleranceCode = "1672007"
	AllergyIntoleranceCode1673002    AllergyIntoleranceCode = "1673002"
	AllergyIntoleranceCode1675009    AllergyIntoleranceCode = "1675009"
	AllergyIntoleranceCode1676005    AllergyIntoleranceCode = "1676005"
	AllergyIntoleranceCode1681001    AllergyIntoleranceCode = "1681001"
	AllergyIntoleranceCode1696002    AllergyIntoleranceCode = "1696002"
	AllergyIntoleranceCode1701009    AllergyIntoleranceCode = "1701009"
	AllergyIntoleranceCode1710001    AllergyIntoleranceCode = "1710001"
	AllergyIntoleranceCode1726000    AllergyIntoleranceCode = "1726000"
	AllergyIntoleranceCode1727009    AllergyIntoleranceCode = "1727009"
	AllergyIntoleranceCode1740004    AllergyIntoleranceCode = "1740004"
	AllergyIntoleranceCode1764003    AllergyIntoleranceCode = "1764003"
	AllergyIntoleranceCode1768000    AllergyIntoleranceCode = "1768000"
	AllergyIntoleranceCode1786002    AllergyIntoleranceCode = "1786002"
	AllergyIntoleranceCode1793003    AllergyIntoleranceCode = "1793003"
	AllergyIntoleranceCode1795005    AllergyIntoleranceCode = "1795005"
	AllergyIntoleranceCode1798007    AllergyIntoleranceCode = "1798007"
	AllergyIntoleranceCode1799004    AllergyIntoleranceCode = "1799004"
	AllergyIntoleranceCode1823002    AllergyIntoleranceCode = "1823002"
	AllergyIntoleranceCode1827001    AllergyIntoleranceCode = "1827001"
	AllergyIntoleranceCode1886008    AllergyIntoleranceCode = "1886008"
	AllergyIntoleranceCode1904005    AllergyIntoleranceCode = "1904005"
	AllergyIntoleranceCode1914001    AllergyIntoleranceCode = "1914001"
	AllergyIntoleranceCode1916004    AllergyIntoleranceCode = "1916004"
	AllergyIntoleranceCode1940007    AllergyIntoleranceCode = "1940007"
	AllergyIntoleranceCode1944003    AllergyIntoleranceCode = "1944003"
	AllergyIntoleranceCode1956002    AllergyIntoleranceCode = "1956002"
	AllergyIntoleranceCode1971003    AllergyIntoleranceCode = "1971003"
	AllergyIntoleranceCode1975007    AllergyIntoleranceCode = "1975007"
	AllergyIntoleranceCode1978009    AllergyIntoleranceCode = "1978009"
	AllergyIntoleranceCode1985008    AllergyIntoleranceCode = "1985008"
	AllergyIntoleranceCode1991005    AllergyIntoleranceCode = "1991005"
	AllergyIntoleranceCode2000001    AllergyIntoleranceCode = "2000001"
	AllergyIntoleranceCode2006007    AllergyIntoleranceCode = "2006007"
	AllergyIntoleranceCode2008008    AllergyIntoleranceCode = "2008008"
	AllergyIntoleranceCode2009000    AllergyIntoleranceCode = "2009000"
	AllergyIntoleranceCode2017008    AllergyIntoleranceCode = "2017008"
	AllergyIntoleranceCode2027002    AllergyIntoleranceCode = "2027002"
	AllergyIntoleranceCode2029004    AllergyIntoleranceCode = "2029004"
	AllergyIntoleranceCode2038002    AllergyIntoleranceCode = "2038002"
	AllergyIntoleranceCode2039005    AllergyIntoleranceCode = "2039005"
	AllergyIntoleranceCode2050008    AllergyIntoleranceCode = "2050008"
	AllergyIntoleranceCode2064008    AllergyIntoleranceCode = "2064008"
	AllergyIntoleranceCode2082006    AllergyIntoleranceCode = "2082006"
	AllergyIntoleranceCode2085008    AllergyIntoleranceCode = "2085008"
	AllergyIntoleranceCode2088005    AllergyIntoleranceCode = "2088005"
	AllergyIntoleranceCode2096000    AllergyIntoleranceCode = "2096000"
	AllergyIntoleranceCode2100004    AllergyIntoleranceCode = "2100004"
	AllergyIntoleranceCode2101000    AllergyIntoleranceCode = "2101000"
	AllergyIntoleranceCode2125008    AllergyIntoleranceCode = "2125008"
	AllergyIntoleranceCode2130007    AllergyIntoleranceCode = "2130007"
	AllergyIntoleranceCode2141009    AllergyIntoleranceCode = "2141009"
	AllergyIntoleranceCode2147008    AllergyIntoleranceCode = "2147008"
	AllergyIntoleranceCode2151005    AllergyIntoleranceCode = "2151005"
	AllergyIntoleranceCode2154002    AllergyIntoleranceCode = "2154002"
	AllergyIntoleranceCode2159007    AllergyIntoleranceCode = "2159007"
	AllergyIntoleranceCode2163000    AllergyIntoleranceCode = "2163000"
	AllergyIntoleranceCode2168009    AllergyIntoleranceCode = "2168009"
	AllergyIntoleranceCode2179004    AllergyIntoleranceCode = "2179004"
	AllergyIntoleranceCode2189000    AllergyIntoleranceCode = "2189000"
	AllergyIntoleranceCode2194000    AllergyIntoleranceCode = "2194000"
	AllergyIntoleranceCode2195004    AllergyIntoleranceCode = "2195004"
	AllergyIntoleranceCode2201007    AllergyIntoleranceCode = "2201007"
	AllergyIntoleranceCode2208001    AllergyIntoleranceCode = "2208001"
	AllergyIntoleranceCode2212007    AllergyIntoleranceCode = "2212007"
	AllergyIntoleranceCode2215009    AllergyIntoleranceCode = "2215009"
	AllergyIntoleranceCode2240002    AllergyIntoleranceCode = "2240002"
	AllergyIntoleranceCode2249001    AllergyIntoleranceCode = "2249001"
	AllergyIntoleranceCode2254005    AllergyIntoleranceCode = "2254005"
	AllergyIntoleranceCode2260005    AllergyIntoleranceCode = "2260005"
	AllergyIntoleranceCode2262002    AllergyIntoleranceCode = "2262002"
	AllergyIntoleranceCode2264001    AllergyIntoleranceCode = "2264001"
	AllergyIntoleranceCode2309006    AllergyIntoleranceCode = "2309006"
	AllergyIntoleranceCode2311002    AllergyIntoleranceCode = "2311002"
	AllergyIntoleranceCode2329007    AllergyIntoleranceCode = "2329007"
	AllergyIntoleranceCode2331003    AllergyIntoleranceCode = "2331003"
	AllergyIntoleranceCode2338009    AllergyIntoleranceCode = "2338009"
	AllergyIntoleranceCode2343002    AllergyIntoleranceCode = "2343002"
	AllergyIntoleranceCode2346005    AllergyIntoleranceCode = "2346005"
	AllergyIntoleranceCode2354007    AllergyIntoleranceCode = "2354007"
	AllergyIntoleranceCode2369008    AllergyIntoleranceCode = "2369008"
	AllergyIntoleranceCode2370009    AllergyIntoleranceCode = "2370009"
	AllergyIntoleranceCode2376003    AllergyIntoleranceCode = "2376003"
	AllergyIntoleranceCode2384004    AllergyIntoleranceCode = "2384004"
	AllergyIntoleranceCode2404002    AllergyIntoleranceCode = "2404002"
	AllergyIntoleranceCode2405001    AllergyIntoleranceCode = "2405001"
	AllergyIntoleranceCode2414006    AllergyIntoleranceCode = "2414006"
	AllergyIntoleranceCode2430003    AllergyIntoleranceCode = "2430003"
	AllergyIntoleranceCode2431004    AllergyIntoleranceCode = "2431004"
	AllergyIntoleranceCode2441001    AllergyIntoleranceCode = "2441001"
	AllergyIntoleranceCode2444009    AllergyIntoleranceCode = "2444009"
	AllergyIntoleranceCode2450004    AllergyIntoleranceCode = "2450004"
	AllergyIntoleranceCode2462000    AllergyIntoleranceCode = "2462000"
	AllergyIntoleranceCode2466002    AllergyIntoleranceCode = "2466002"
	AllergyIntoleranceCode2500009    AllergyIntoleranceCode = "2500009"
	AllergyIntoleranceCode2509005    AllergyIntoleranceCode = "2509005"
	AllergyIntoleranceCode2522002    AllergyIntoleranceCode = "2522002"
	AllergyIntoleranceCode2529006    AllergyIntoleranceCode = "2529006"
	AllergyIntoleranceCode2537003    AllergyIntoleranceCode = "2537003"
	AllergyIntoleranceCode2568004    AllergyIntoleranceCode = "2568004"
	AllergyIntoleranceCode2573005    AllergyIntoleranceCode = "2573005"
	AllergyIntoleranceCode2575003    AllergyIntoleranceCode = "2575003"
	AllergyIntoleranceCode2595009    AllergyIntoleranceCode = "2595009"
	AllergyIntoleranceCode2597001    AllergyIntoleranceCode = "2597001"
	AllergyIntoleranceCode2611008    AllergyIntoleranceCode = "2611008"
	AllergyIntoleranceCode2637006    AllergyIntoleranceCode = "2637006"
	AllergyIntoleranceCode2648004    AllergyIntoleranceCode = "2648004"
	AllergyIntoleranceCode2649007    AllergyIntoleranceCode = "2649007"
	AllergyIntoleranceCode2660003    AllergyIntoleranceCode = "2660003"
	AllergyIntoleranceCode2671002    AllergyIntoleranceCode = "2671002"
	AllergyIntoleranceCode2674005    AllergyIntoleranceCode = "2674005"
	AllergyIntoleranceCode2676007    AllergyIntoleranceCode = "2676007"
	AllergyIntoleranceCode2678008    AllergyIntoleranceCode = "2678008"
	AllergyIntoleranceCode2680002    AllergyIntoleranceCode = "2680002"
	AllergyIntoleranceCode2698003    AllergyIntoleranceCode = "2698003"
	AllergyIntoleranceCode2705002    AllergyIntoleranceCode = "2705002"
	AllergyIntoleranceCode2706001    AllergyIntoleranceCode = "2706001"
	AllergyIntoleranceCode2719002    AllergyIntoleranceCode = "2719002"
	AllergyIntoleranceCode2721007    AllergyIntoleranceCode = "2721007"
	AllergyIntoleranceCode2728001    AllergyIntoleranceCode = "2728001"
	AllergyIntoleranceCode2753003    AllergyIntoleranceCode = "2753003"
	AllergyIntoleranceCode2754009    AllergyIntoleranceCode = "2754009"
	AllergyIntoleranceCode2765004    AllergyIntoleranceCode = "2765004"
	AllergyIntoleranceCode2778004    AllergyIntoleranceCode = "2778004"
	AllergyIntoleranceCode2796008    AllergyIntoleranceCode = "2796008"
	AllergyIntoleranceCode2799001    AllergyIntoleranceCode = "2799001"
	AllergyIntoleranceCode2823004    AllergyIntoleranceCode = "2823004"
	AllergyIntoleranceCode2832002    AllergyIntoleranceCode = "2832002"
	AllergyIntoleranceCode2846002    AllergyIntoleranceCode = "2846002"
	AllergyIntoleranceCode2869004    AllergyIntoleranceCode = "2869004"
	AllergyIntoleranceCode2878005    AllergyIntoleranceCode = "2878005"
	AllergyIntoleranceCode2880004    AllergyIntoleranceCode = "2880004"
	AllergyIntoleranceCode2883002    AllergyIntoleranceCode = "2883002"
	AllergyIntoleranceCode2913009    AllergyIntoleranceCode = "2913009"
	AllergyIntoleranceCode2916001    AllergyIntoleranceCode = "2916001"
	AllergyIntoleranceCode2925007    AllergyIntoleranceCode = "2925007"
	AllergyIntoleranceCode2927004    AllergyIntoleranceCode = "2927004"
	AllergyIntoleranceCode2938004    AllergyIntoleranceCode = "2938004"
	AllergyIntoleranceCode2942001    AllergyIntoleranceCode = "2942001"
	AllergyIntoleranceCode2950005    AllergyIntoleranceCode = "2950005"
	AllergyIntoleranceCode2958003    AllergyIntoleranceCode = "2958003"
	AllergyIntoleranceCode2964005    AllergyIntoleranceCode = "2964005"
	AllergyIntoleranceCode2974008    AllergyIntoleranceCode = "2974008"
	AllergyIntoleranceCode2988007    AllergyIntoleranceCode = "2988007"
	AllergyIntoleranceCode2991007    AllergyIntoleranceCode = "2991007"
	AllergyIntoleranceCode2995003    AllergyIntoleranceCode = "2995003"
	AllergyIntoleranceCode3027009    AllergyIntoleranceCode = "3027009"
	AllergyIntoleranceCode3031003    AllergyIntoleranceCode = "3031003"
	AllergyIntoleranceCode3040004    AllergyIntoleranceCode = "3040004"
	AllergyIntoleranceCode3045009    AllergyIntoleranceCode = "3045009"
	AllergyIntoleranceCode3052006    AllergyIntoleranceCode = "3052006"
	AllergyIntoleranceCode3066001    AllergyIntoleranceCode = "3066001"
	AllergyIntoleranceCode3070009    AllergyIntoleranceCode = "3070009"
	AllergyIntoleranceCode3087006    AllergyIntoleranceCode = "3087006"
	AllergyIntoleranceCode3107005    AllergyIntoleranceCode = "3107005"
	AllergyIntoleranceCode3108000    AllergyIntoleranceCode = "3108000"
	AllergyIntoleranceCode3131000    AllergyIntoleranceCode = "3131000"
	AllergyIntoleranceCode3136005    AllergyIntoleranceCode = "3136005"
	AllergyIntoleranceCode3142009    AllergyIntoleranceCode = "3142009"
	AllergyIntoleranceCode3145006    AllergyIntoleranceCode = "3145006"
	AllergyIntoleranceCode3150000    AllergyIntoleranceCode = "3150000"
	AllergyIntoleranceCode3151001    AllergyIntoleranceCode = "3151001"
	AllergyIntoleranceCode3155005    AllergyIntoleranceCode = "3155005"
	AllergyIntoleranceCode3161008    AllergyIntoleranceCode = "3161008"
	AllergyIntoleranceCode3167007    AllergyIntoleranceCode = "3167007"
	AllergyIntoleranceCode3187008    AllergyIntoleranceCode = "3187008"
	AllergyIntoleranceCode3193000    AllergyIntoleranceCode = "3193000"
	AllergyIntoleranceCode3197004    AllergyIntoleranceCode = "3197004"
	AllergyIntoleranceCode3209002    AllergyIntoleranceCode = "3209002"
	AllergyIntoleranceCode3212004    AllergyIntoleranceCode = "3212004"
	AllergyIntoleranceCode3225007    AllergyIntoleranceCode = "3225007"
	AllergyIntoleranceCode3232003    AllergyIntoleranceCode = "3232003"
	AllergyIntoleranceCode3271000    AllergyIntoleranceCode = "3271000"
	AllergyIntoleranceCode3273002    AllergyIntoleranceCode = "3273002"
	AllergyIntoleranceCode3300001    AllergyIntoleranceCode = "3300001"
	AllergyIntoleranceCode3318003    AllergyIntoleranceCode = "3318003"
	AllergyIntoleranceCode3325005    AllergyIntoleranceCode = "3325005"
	AllergyIntoleranceCode3339005    AllergyIntoleranceCode = "3339005"
	AllergyIntoleranceCode3340007    AllergyIntoleranceCode = "3340007"
	AllergyIntoleranceCode3342004    AllergyIntoleranceCode = "3342004"
	AllergyIntoleranceCode3346001    AllergyIntoleranceCode = "3346001"
	AllergyIntoleranceCode3378009    AllergyIntoleranceCode = "3378009"
	AllergyIntoleranceCode3379001    AllergyIntoleranceCode = "3379001"
	AllergyIntoleranceCode3392003    AllergyIntoleranceCode = "3392003"
	AllergyIntoleranceCode3405005    AllergyIntoleranceCode = "3405005"
	AllergyIntoleranceCode3411008    AllergyIntoleranceCode = "3411008"
	AllergyIntoleranceCode3437006    AllergyIntoleranceCode = "3437006"
	AllergyIntoleranceCode3440006    AllergyIntoleranceCode = "3440006"
	AllergyIntoleranceCode3455002    AllergyIntoleranceCode = "3455002"
	AllergyIntoleranceCode3463001    AllergyIntoleranceCode = "3463001"
	AllergyIntoleranceCode3465008    AllergyIntoleranceCode = "3465008"
	AllergyIntoleranceCode3466009    AllergyIntoleranceCode = "3466009"
	AllergyIntoleranceCode3492002    AllergyIntoleranceCode = "3492002"
	AllergyIntoleranceCode3493007    AllergyIntoleranceCode = "3493007"
	AllergyIntoleranceCode3495000    AllergyIntoleranceCode = "3495000"
	AllergyIntoleranceCode3501003    AllergyIntoleranceCode = "3501003"
	AllergyIntoleranceCode3523004    AllergyIntoleranceCode = "3523004"
	AllergyIntoleranceCode3532002    AllergyIntoleranceCode = "3532002"
	AllergyIntoleranceCode3555004    AllergyIntoleranceCode = "3555004"
	AllergyIntoleranceCode3579002    AllergyIntoleranceCode = "3579002"
	AllergyIntoleranceCode3581000    AllergyIntoleranceCode = "3581000"
	AllergyIntoleranceCode3587001    AllergyIntoleranceCode = "3587001"
	AllergyIntoleranceCode3588006    AllergyIntoleranceCode = "3588006"
	AllergyIntoleranceCode3597005    AllergyIntoleranceCode = "3597005"
	AllergyIntoleranceCode3602003    AllergyIntoleranceCode = "3602003"
	AllergyIntoleranceCode3610002    AllergyIntoleranceCode = "3610002"
	AllergyIntoleranceCode3617004    AllergyIntoleranceCode = "3617004"
	AllergyIntoleranceCode3648007    AllergyIntoleranceCode = "3648007"
	AllergyIntoleranceCode3655009    AllergyIntoleranceCode = "3655009"
	AllergyIntoleranceCode3672002    AllergyIntoleranceCode = "3672002"
	AllergyIntoleranceCode3684000    AllergyIntoleranceCode = "3684000"
	AllergyIntoleranceCode3689005    AllergyIntoleranceCode = "3689005"
	AllergyIntoleranceCode3692009    AllergyIntoleranceCode = "3692009"
	AllergyIntoleranceCode3693004    AllergyIntoleranceCode = "3693004"
	AllergyIntoleranceCode3702007    AllergyIntoleranceCode = "3702007"
	AllergyIntoleranceCode3710008    AllergyIntoleranceCode = "3710008"
	AllergyIntoleranceCode3718001    AllergyIntoleranceCode = "3718001"
	AllergyIntoleranceCode3726009    AllergyIntoleranceCode = "3726009"
	AllergyIntoleranceCode3727000    AllergyIntoleranceCode = "3727000"
	AllergyIntoleranceCode3730007    AllergyIntoleranceCode = "3730007"
	AllergyIntoleranceCode3737005    AllergyIntoleranceCode = "3737005"
	AllergyIntoleranceCode3742002    AllergyIntoleranceCode = "3742002"
	AllergyIntoleranceCode3757009    AllergyIntoleranceCode = "3757009"
	AllergyIntoleranceCode3771001    AllergyIntoleranceCode = "3771001"
	AllergyIntoleranceCode3775005    AllergyIntoleranceCode = "3775005"
	AllergyIntoleranceCode3776006    AllergyIntoleranceCode = "3776006"
	AllergyIntoleranceCode3792001    AllergyIntoleranceCode = "3792001"
	AllergyIntoleranceCode3800009    AllergyIntoleranceCode = "3800009"
	AllergyIntoleranceCode3807007    AllergyIntoleranceCode = "3807007"
	AllergyIntoleranceCode3811001    AllergyIntoleranceCode = "3811001"
	AllergyIntoleranceCode3812008    AllergyIntoleranceCode = "3812008"
	AllergyIntoleranceCode3816006    AllergyIntoleranceCode = "3816006"
	AllergyIntoleranceCode3823007    AllergyIntoleranceCode = "3823007"
	AllergyIntoleranceCode3829006    AllergyIntoleranceCode = "3829006"
	AllergyIntoleranceCode3834005    AllergyIntoleranceCode = "3834005"
	AllergyIntoleranceCode3836007    AllergyIntoleranceCode = "3836007"
	AllergyIntoleranceCode3844007    AllergyIntoleranceCode = "3844007"
	AllergyIntoleranceCode3848005    AllergyIntoleranceCode = "3848005"
	AllergyIntoleranceCode3849002    AllergyIntoleranceCode = "3849002"
	AllergyIntoleranceCode3854006    AllergyIntoleranceCode = "3854006"
	AllergyIntoleranceCode3874004    AllergyIntoleranceCode = "3874004"
	AllergyIntoleranceCode3892007    AllergyIntoleranceCode = "3892007"
	AllergyIntoleranceCode3896005    AllergyIntoleranceCode = "3896005"
	AllergyIntoleranceCode3897001    AllergyIntoleranceCode = "3897001"
	AllergyIntoleranceCode3906002    AllergyIntoleranceCode = "3906002"
	AllergyIntoleranceCode3920009    AllergyIntoleranceCode = "3920009"
	AllergyIntoleranceCode3930000    AllergyIntoleranceCode = "3930000"
	AllergyIntoleranceCode3932008    AllergyIntoleranceCode = "3932008"
	AllergyIntoleranceCode3941003    AllergyIntoleranceCode = "3941003"
	AllergyIntoleranceCode3945007    AllergyIntoleranceCode = "3945007"
	AllergyIntoleranceCode3958008    AllergyIntoleranceCode = "3958008"
	AllergyIntoleranceCode3961009    AllergyIntoleranceCode = "3961009"
	AllergyIntoleranceCode3976001    AllergyIntoleranceCode = "3976001"
	AllergyIntoleranceCode3982003    AllergyIntoleranceCode = "3982003"
	AllergyIntoleranceCode3983008    AllergyIntoleranceCode = "3983008"
	AllergyIntoleranceCode3990003    AllergyIntoleranceCode = "3990003"
	AllergyIntoleranceCode3994007    AllergyIntoleranceCode = "3994007"
	AllergyIntoleranceCode4014000    AllergyIntoleranceCode = "4014000"
	AllergyIntoleranceCode4024008    AllergyIntoleranceCode = "4024008"
	AllergyIntoleranceCode4025009    AllergyIntoleranceCode = "4025009"
	AllergyIntoleranceCode4043008    AllergyIntoleranceCode = "4043008"
	AllergyIntoleranceCode4047009    AllergyIntoleranceCode = "4047009"
	AllergyIntoleranceCode4048004    AllergyIntoleranceCode = "4048004"
	AllergyIntoleranceCode4067000    AllergyIntoleranceCode = "4067000"
	AllergyIntoleranceCode4076007    AllergyIntoleranceCode = "4076007"
	AllergyIntoleranceCode4077003    AllergyIntoleranceCode = "4077003"
	AllergyIntoleranceCode4080002    AllergyIntoleranceCode = "4080002"
	AllergyIntoleranceCode4091009    AllergyIntoleranceCode = "4091009"
	AllergyIntoleranceCode4097008    AllergyIntoleranceCode = "4097008"
	AllergyIntoleranceCode4104007    AllergyIntoleranceCode = "4104007"
	AllergyIntoleranceCode4105008    AllergyIntoleranceCode = "4105008"
	AllergyIntoleranceCode4115002    AllergyIntoleranceCode = "4115002"
	AllergyIntoleranceCode4137009    AllergyIntoleranceCode = "4137009"
	AllergyIntoleranceCode4153007    AllergyIntoleranceCode = "4153007"
	AllergyIntoleranceCode4167003    AllergyIntoleranceCode = "4167003"
	AllergyIntoleranceCode4169000    AllergyIntoleranceCode = "4169000"
	AllergyIntoleranceCode4177001    AllergyIntoleranceCode = "4177001"
	AllergyIntoleranceCode4182008    AllergyIntoleranceCode = "4182008"
	AllergyIntoleranceCode4188007    AllergyIntoleranceCode = "4188007"
	AllergyIntoleranceCode4200007    AllergyIntoleranceCode = "4200007"
	AllergyIntoleranceCode4201006    AllergyIntoleranceCode = "4201006"
	AllergyIntoleranceCode4203009    AllergyIntoleranceCode = "4203009"
	AllergyIntoleranceCode4207005    AllergyIntoleranceCode = "4207005"
	AllergyIntoleranceCode4217000    AllergyIntoleranceCode = "4217000"
	AllergyIntoleranceCode4218005    AllergyIntoleranceCode = "4218005"
	AllergyIntoleranceCode4231000    AllergyIntoleranceCode = "4231000"
	AllergyIntoleranceCode4239003    AllergyIntoleranceCode = "4239003"
	AllergyIntoleranceCode4255005    AllergyIntoleranceCode = "4255005"
	AllergyIntoleranceCode4289006    AllergyIntoleranceCode = "4289006"
	AllergyIntoleranceCode4290002    AllergyIntoleranceCode = "4290002"
	AllergyIntoleranceCode4314009    AllergyIntoleranceCode = "4314009"
	AllergyIntoleranceCode4334005    AllergyIntoleranceCode = "4334005"
	AllergyIntoleranceCode4342006    AllergyIntoleranceCode = "4342006"
	AllergyIntoleranceCode4353000    AllergyIntoleranceCode = "4353000"
	AllergyIntoleranceCode4355007    AllergyIntoleranceCode = "4355007"
	AllergyIntoleranceCode4362003    AllergyIntoleranceCode = "4362003"
	AllergyIntoleranceCode4370008    AllergyIntoleranceCode = "4370008"
	AllergyIntoleranceCode4393002    AllergyIntoleranceCode = "4393002"
	AllergyIntoleranceCode4401009    AllergyIntoleranceCode = "4401009"
	AllergyIntoleranceCode4422003    AllergyIntoleranceCode = "4422003"
	AllergyIntoleranceCode4423008    AllergyIntoleranceCode = "4423008"
	AllergyIntoleranceCode4425001    AllergyIntoleranceCode = "4425001"
	AllergyIntoleranceCode4435007    AllergyIntoleranceCode = "4435007"
	AllergyIntoleranceCode4437004    AllergyIntoleranceCode = "4437004"
	AllergyIntoleranceCode4471008    AllergyIntoleranceCode = "4471008"
	AllergyIntoleranceCode4479005    AllergyIntoleranceCode = "4479005"
	AllergyIntoleranceCode4480008    AllergyIntoleranceCode = "4480008"
	AllergyIntoleranceCode4509009    AllergyIntoleranceCode = "4509009"
	AllergyIntoleranceCode4534009    AllergyIntoleranceCode = "4534009"
	AllergyIntoleranceCode4540002    AllergyIntoleranceCode = "4540002"
	AllergyIntoleranceCode4546008    AllergyIntoleranceCode = "4546008"
	AllergyIntoleranceCode4555006    AllergyIntoleranceCode = "4555006"
	AllergyIntoleranceCode4560005    AllergyIntoleranceCode = "4560005"
	AllergyIntoleranceCode4561009    AllergyIntoleranceCode = "4561009"
	AllergyIntoleranceCode4564001    AllergyIntoleranceCode = "4564001"
	AllergyIntoleranceCode4567008    AllergyIntoleranceCode = "4567008"
	AllergyIntoleranceCode4582003    AllergyIntoleranceCode = "4582003"
	AllergyIntoleranceCode4591004    AllergyIntoleranceCode = "4591004"
	AllergyIntoleranceCode4610008    AllergyIntoleranceCode = "4610008"
	AllergyIntoleranceCode4616002    AllergyIntoleranceCode = "4616002"
	AllergyIntoleranceCode4629002    AllergyIntoleranceCode = "4629002"
	AllergyIntoleranceCode4635002    AllergyIntoleranceCode = "4635002"
	AllergyIntoleranceCode4643007    AllergyIntoleranceCode = "4643007"
	AllergyIntoleranceCode4656000    AllergyIntoleranceCode = "4656000"
	AllergyIntoleranceCode4674009    AllergyIntoleranceCode = "4674009"
	AllergyIntoleranceCode4681002    AllergyIntoleranceCode = "4681002"
	AllergyIntoleranceCode4693006    AllergyIntoleranceCode = "4693006"
	AllergyIntoleranceCode4700006    AllergyIntoleranceCode = "4700006"
	AllergyIntoleranceCode4706000    AllergyIntoleranceCode = "4706000"
	AllergyIntoleranceCode4714006    AllergyIntoleranceCode = "4714006"
	AllergyIntoleranceCode4728000    AllergyIntoleranceCode = "4728000"
	AllergyIntoleranceCode4732006    AllergyIntoleranceCode = "4732006"
	AllergyIntoleranceCode4746006    AllergyIntoleranceCode = "4746006"
	AllergyIntoleranceCode4761007    AllergyIntoleranceCode = "4761007"
	AllergyIntoleranceCode4762000    AllergyIntoleranceCode = "4762000"
	AllergyIntoleranceCode4777008    AllergyIntoleranceCode = "4777008"
	AllergyIntoleranceCode4780009    AllergyIntoleranceCode = "4780009"
	AllergyIntoleranceCode4786003    AllergyIntoleranceCode = "4786003"
	AllergyIntoleranceCode4789005    AllergyIntoleranceCode = "4789005"
	AllergyIntoleranceCode4793004    AllergyIntoleranceCode = "4793004"
	AllergyIntoleranceCode4814001    AllergyIntoleranceCode = "4814001"
	AllergyIntoleranceCode4824009    AllergyIntoleranceCode = "4824009"
	AllergyIntoleranceCode4825005    AllergyIntoleranceCode = "4825005"
	AllergyIntoleranceCode4831008    AllergyIntoleranceCode = "4831008"
	AllergyIntoleranceCode4832001    AllergyIntoleranceCode = "4832001"
	AllergyIntoleranceCode4833006    AllergyIntoleranceCode = "4833006"
	AllergyIntoleranceCode4844003    AllergyIntoleranceCode = "4844003"
	AllergyIntoleranceCode4864008    AllergyIntoleranceCode = "4864008"
	AllergyIntoleranceCode4872005    AllergyIntoleranceCode = "4872005"
	AllergyIntoleranceCode4878009    AllergyIntoleranceCode = "4878009"
	AllergyIntoleranceCode4882006    AllergyIntoleranceCode = "4882006"
	AllergyIntoleranceCode4889002    AllergyIntoleranceCode = "4889002"
	AllergyIntoleranceCode4901003    AllergyIntoleranceCode = "4901003"
	AllergyIntoleranceCode4925006    AllergyIntoleranceCode = "4925006"
	AllergyIntoleranceCode4933007    AllergyIntoleranceCode = "4933007"
	AllergyIntoleranceCode4940008    AllergyIntoleranceCode = "4940008"
	AllergyIntoleranceCode4955004    AllergyIntoleranceCode = "4955004"
	AllergyIntoleranceCode4962008    AllergyIntoleranceCode = "4962008"
	AllergyIntoleranceCode4963003    AllergyIntoleranceCode = "4963003"
	AllergyIntoleranceCode4965005    AllergyIntoleranceCode = "4965005"
	AllergyIntoleranceCode4968007    AllergyIntoleranceCode = "4968007"
	AllergyIntoleranceCode4986005    AllergyIntoleranceCode = "4986005"
	AllergyIntoleranceCode5003005    AllergyIntoleranceCode = "5003005"
	AllergyIntoleranceCode5004004    AllergyIntoleranceCode = "5004004"
	AllergyIntoleranceCode5007006    AllergyIntoleranceCode = "5007006"
	AllergyIntoleranceCode5024000    AllergyIntoleranceCode = "5024000"
	AllergyIntoleranceCode5031001    AllergyIntoleranceCode = "5031001"
	AllergyIntoleranceCode5040002    AllergyIntoleranceCode = "5040002"
	AllergyIntoleranceCode5043000    AllergyIntoleranceCode = "5043000"
	AllergyIntoleranceCode5045007    AllergyIntoleranceCode = "5045007"
	AllergyIntoleranceCode5059000    AllergyIntoleranceCode = "5059000"
	AllergyIntoleranceCode5060005    AllergyIntoleranceCode = "5060005"
	AllergyIntoleranceCode5061009    AllergyIntoleranceCode = "5061009"
	AllergyIntoleranceCode5064001    AllergyIntoleranceCode = "5064001"
	AllergyIntoleranceCode5081005    AllergyIntoleranceCode = "5081005"
	AllergyIntoleranceCode5086000    AllergyIntoleranceCode = "5086000"
	AllergyIntoleranceCode5094007    AllergyIntoleranceCode = "5094007"
	AllergyIntoleranceCode5098005    AllergyIntoleranceCode = "5098005"
	AllergyIntoleranceCode5109006    AllergyIntoleranceCode = "5109006"
	AllergyIntoleranceCode5142007    AllergyIntoleranceCode = "5142007"
	AllergyIntoleranceCode5160007    AllergyIntoleranceCode = "5160007"
	AllergyIntoleranceCode5163009    AllergyIntoleranceCode = "5163009"
	AllergyIntoleranceCode5167005    AllergyIntoleranceCode = "5167005"
	AllergyIntoleranceCode5172001    AllergyIntoleranceCode = "5172001"
	AllergyIntoleranceCode5179005    AllergyIntoleranceCode = "5179005"
	AllergyIntoleranceCode5200001    AllergyIntoleranceCode = "5200001"
	AllergyIntoleranceCode5206007    AllergyIntoleranceCode = "5206007"
	AllergyIntoleranceCode5220000    AllergyIntoleranceCode = "5220000"
	AllergyIntoleranceCode5226006    AllergyIntoleranceCode = "5226006"
	AllergyIntoleranceCode5250008    AllergyIntoleranceCode = "5250008"
	AllergyIntoleranceCode5252000    AllergyIntoleranceCode = "5252000"
	AllergyIntoleranceCode5253005    AllergyIntoleranceCode = "5253005"
	AllergyIntoleranceCode5259009    AllergyIntoleranceCode = "5259009"
	AllergyIntoleranceCode5289002    AllergyIntoleranceCode = "5289002"
	AllergyIntoleranceCode5303002    AllergyIntoleranceCode = "5303002"
	AllergyIntoleranceCode5305009    AllergyIntoleranceCode = "5305009"
	AllergyIntoleranceCode5307001    AllergyIntoleranceCode = "5307001"
	AllergyIntoleranceCode5312000    AllergyIntoleranceCode = "5312000"
	AllergyIntoleranceCode5323001    AllergyIntoleranceCode = "5323001"
	AllergyIntoleranceCode5330007    AllergyIntoleranceCode = "5330007"
	AllergyIntoleranceCode5339008    AllergyIntoleranceCode = "5339008"
	AllergyIntoleranceCode5340005    AllergyIntoleranceCode = "5340005"
	AllergyIntoleranceCode5392001    AllergyIntoleranceCode = "5392001"
	AllergyIntoleranceCode5395004    AllergyIntoleranceCode = "5395004"
	AllergyIntoleranceCode5404007    AllergyIntoleranceCode = "5404007"
	AllergyIntoleranceCode5405008    AllergyIntoleranceCode = "5405008"
	AllergyIntoleranceCode5406009    AllergyIntoleranceCode = "5406009"
	AllergyIntoleranceCode5420002    AllergyIntoleranceCode = "5420002"
	AllergyIntoleranceCode5439007    AllergyIntoleranceCode = "5439007"
	AllergyIntoleranceCode5442001    AllergyIntoleranceCode = "5442001"
	AllergyIntoleranceCode5453007    AllergyIntoleranceCode = "5453007"
	AllergyIntoleranceCode5471000    AllergyIntoleranceCode = "5471000"
	AllergyIntoleranceCode5474008    AllergyIntoleranceCode = "5474008"
	AllergyIntoleranceCode5477001    AllergyIntoleranceCode = "5477001"
	AllergyIntoleranceCode5483003    AllergyIntoleranceCode = "5483003"
	AllergyIntoleranceCode5504009    AllergyIntoleranceCode = "5504009"
	AllergyIntoleranceCode5511008    AllergyIntoleranceCode = "5511008"
	AllergyIntoleranceCode5513006    AllergyIntoleranceCode = "5513006"
	AllergyIntoleranceCode5515004    AllergyIntoleranceCode = "5515004"
	AllergyIntoleranceCode5533005    AllergyIntoleranceCode = "5533005"
	AllergyIntoleranceCode5537006    AllergyIntoleranceCode = "5537006"
	AllergyIntoleranceCode5540006    AllergyIntoleranceCode = "5540006"
	AllergyIntoleranceCode5547009    AllergyIntoleranceCode = "5547009"
	AllergyIntoleranceCode5548004    AllergyIntoleranceCode = "5548004"
	AllergyIntoleranceCode5568005    AllergyIntoleranceCode = "5568005"
	AllergyIntoleranceCode5573004    AllergyIntoleranceCode = "5573004"
	AllergyIntoleranceCode5589001    AllergyIntoleranceCode = "5589001"
	AllergyIntoleranceCode5590005    AllergyIntoleranceCode = "5590005"
	AllergyIntoleranceCode5628003    AllergyIntoleranceCode = "5628003"
	AllergyIntoleranceCode5629006    AllergyIntoleranceCode = "5629006"
	AllergyIntoleranceCode5637003    AllergyIntoleranceCode = "5637003"
	AllergyIntoleranceCode5641004    AllergyIntoleranceCode = "5641004"
	AllergyIntoleranceCode5647000    AllergyIntoleranceCode = "5647000"
	AllergyIntoleranceCode5656008    AllergyIntoleranceCode = "5656008"
	AllergyIntoleranceCode5659001    AllergyIntoleranceCode = "5659001"
	AllergyIntoleranceCode5670008    AllergyIntoleranceCode = "5670008"
	AllergyIntoleranceCode5681006    AllergyIntoleranceCode = "5681006"
	AllergyIntoleranceCode5691000    AllergyIntoleranceCode = "5691000"
	AllergyIntoleranceCode5692007    AllergyIntoleranceCode = "5692007"
	AllergyIntoleranceCode5699003    AllergyIntoleranceCode = "5699003"
	AllergyIntoleranceCode5700002    AllergyIntoleranceCode = "5700002"
	AllergyIntoleranceCode5702005    AllergyIntoleranceCode = "5702005"
	AllergyIntoleranceCode5704006    AllergyIntoleranceCode = "5704006"
	AllergyIntoleranceCode5705007    AllergyIntoleranceCode = "5705007"
	AllergyIntoleranceCode5739006    AllergyIntoleranceCode = "5739006"
	AllergyIntoleranceCode5746002    AllergyIntoleranceCode = "5746002"
	AllergyIntoleranceCode5757007    AllergyIntoleranceCode = "5757007"
	AllergyIntoleranceCode5762008    AllergyIntoleranceCode = "5762008"
	AllergyIntoleranceCode5764009    AllergyIntoleranceCode = "5764009"
	AllergyIntoleranceCode5767002    AllergyIntoleranceCode = "5767002"
	AllergyIntoleranceCode5774007    AllergyIntoleranceCode = "5774007"
	AllergyIntoleranceCode5800007    AllergyIntoleranceCode = "5800007"
	AllergyIntoleranceCode5813001    AllergyIntoleranceCode = "5813001"
	AllergyIntoleranceCode5826002    AllergyIntoleranceCode = "5826002"
	AllergyIntoleranceCode5827006    AllergyIntoleranceCode = "5827006"
	AllergyIntoleranceCode5829009    AllergyIntoleranceCode = "5829009"
	AllergyIntoleranceCode5830004    AllergyIntoleranceCode = "5830004"
	AllergyIntoleranceCode5840001    AllergyIntoleranceCode = "5840001"
	AllergyIntoleranceCode5858007    AllergyIntoleranceCode = "5858007"
	AllergyIntoleranceCode5863006    AllergyIntoleranceCode = "5863006"
	AllergyIntoleranceCode5896008    AllergyIntoleranceCode = "5896008"
	AllergyIntoleranceCode5899001    AllergyIntoleranceCode = "5899001"
	AllergyIntoleranceCode5907009    AllergyIntoleranceCode = "5907009"
	AllergyIntoleranceCode5910002    AllergyIntoleranceCode = "5910002"
	AllergyIntoleranceCode5915007    AllergyIntoleranceCode = "5915007"
	AllergyIntoleranceCode5927005    AllergyIntoleranceCode = "5927005"
	AllergyIntoleranceCode5931004    AllergyIntoleranceCode = "5931004"
	AllergyIntoleranceCode5932006    AllergyIntoleranceCode = "5932006"
	AllergyIntoleranceCode5950004    AllergyIntoleranceCode = "5950004"
	AllergyIntoleranceCode5955009    AllergyIntoleranceCode = "5955009"
	AllergyIntoleranceCode5977008    AllergyIntoleranceCode = "5977008"
	AllergyIntoleranceCode5989005    AllergyIntoleranceCode = "5989005"
	AllergyIntoleranceCode5991002    AllergyIntoleranceCode = "5991002"
	AllergyIntoleranceCode6021003    AllergyIntoleranceCode = "6021003"
	AllergyIntoleranceCode6038004    AllergyIntoleranceCode = "6038004"
	AllergyIntoleranceCode6043006    AllergyIntoleranceCode = "6043006"
	AllergyIntoleranceCode6044000    AllergyIntoleranceCode = "6044000"
	AllergyIntoleranceCode6054001    AllergyIntoleranceCode = "6054001"
	AllergyIntoleranceCode6056004    AllergyIntoleranceCode = "6056004"
	AllergyIntoleranceCode6068008    AllergyIntoleranceCode = "6068008"
	AllergyIntoleranceCode6083003    AllergyIntoleranceCode = "6083003"
	AllergyIntoleranceCode6085005    AllergyIntoleranceCode = "6085005"
	AllergyIntoleranceCode6088007    AllergyIntoleranceCode = "6088007"
	AllergyIntoleranceCode6089004    AllergyIntoleranceCode = "6089004"
	AllergyIntoleranceCode6091007    AllergyIntoleranceCode = "6091007"
	AllergyIntoleranceCode6107003    AllergyIntoleranceCode = "6107003"
	AllergyIntoleranceCode6109000    AllergyIntoleranceCode = "6109000"
	AllergyIntoleranceCode6115000    AllergyIntoleranceCode = "6115000"
	AllergyIntoleranceCode6135004    AllergyIntoleranceCode = "6135004"
	AllergyIntoleranceCode6138002    AllergyIntoleranceCode = "6138002"
	AllergyIntoleranceCode6162007    AllergyIntoleranceCode = "6162007"
	AllergyIntoleranceCode6170002    AllergyIntoleranceCode = "6170002"
	AllergyIntoleranceCode6172005    AllergyIntoleranceCode = "6172005"
	AllergyIntoleranceCode6178009    AllergyIntoleranceCode = "6178009"
	AllergyIntoleranceCode6179001    AllergyIntoleranceCode = "6179001"
	AllergyIntoleranceCode6182006    AllergyIntoleranceCode = "6182006"
	AllergyIntoleranceCode6197009    AllergyIntoleranceCode = "6197009"
	AllergyIntoleranceCode6237004    AllergyIntoleranceCode = "6237004"
	AllergyIntoleranceCode6249003    AllergyIntoleranceCode = "6249003"
	AllergyIntoleranceCode6256009    AllergyIntoleranceCode = "6256009"
	AllergyIntoleranceCode6257000    AllergyIntoleranceCode = "6257000"
	AllergyIntoleranceCode6260007    AllergyIntoleranceCode = "6260007"
	AllergyIntoleranceCode6261006    AllergyIntoleranceCode = "6261006"
	AllergyIntoleranceCode6263009    AllergyIntoleranceCode = "6263009"
	AllergyIntoleranceCode6264003    AllergyIntoleranceCode = "6264003"
	AllergyIntoleranceCode6287006    AllergyIntoleranceCode = "6287006"
	AllergyIntoleranceCode6291001    AllergyIntoleranceCode = "6291001"
	AllergyIntoleranceCode6301006    AllergyIntoleranceCode = "6301006"
	AllergyIntoleranceCode6310003    AllergyIntoleranceCode = "6310003"
	AllergyIntoleranceCode6314007    AllergyIntoleranceCode = "6314007"
	AllergyIntoleranceCode6333002    AllergyIntoleranceCode = "6333002"
	AllergyIntoleranceCode6338006    AllergyIntoleranceCode = "6338006"
	AllergyIntoleranceCode6356006    AllergyIntoleranceCode = "6356006"
	AllergyIntoleranceCode6360009    AllergyIntoleranceCode = "6360009"
	AllergyIntoleranceCode6367007    AllergyIntoleranceCode = "6367007"
	AllergyIntoleranceCode6386004    AllergyIntoleranceCode = "6386004"
	AllergyIntoleranceCode6394006    AllergyIntoleranceCode = "6394006"
	AllergyIntoleranceCode6401007    AllergyIntoleranceCode = "6401007"
	AllergyIntoleranceCode6409009    AllergyIntoleranceCode = "6409009"
	AllergyIntoleranceCode6411000    AllergyIntoleranceCode = "6411000"
	AllergyIntoleranceCode6422001    AllergyIntoleranceCode = "6422001"
	AllergyIntoleranceCode6451002    AllergyIntoleranceCode = "6451002"
	AllergyIntoleranceCode6455006    AllergyIntoleranceCode = "6455006"
	AllergyIntoleranceCode6469006    AllergyIntoleranceCode = "6469006"
	AllergyIntoleranceCode6478000    AllergyIntoleranceCode = "6478000"
	AllergyIntoleranceCode6495008    AllergyIntoleranceCode = "6495008"
	AllergyIntoleranceCode6507009    AllergyIntoleranceCode = "6507009"
	AllergyIntoleranceCode6513000    AllergyIntoleranceCode = "6513000"
	AllergyIntoleranceCode6516008    AllergyIntoleranceCode = "6516008"
	AllergyIntoleranceCode6524003    AllergyIntoleranceCode = "6524003"
	AllergyIntoleranceCode6529008    AllergyIntoleranceCode = "6529008"
	AllergyIntoleranceCode6532006    AllergyIntoleranceCode = "6532006"
	AllergyIntoleranceCode6590001    AllergyIntoleranceCode = "6590001"
	AllergyIntoleranceCode6592009    AllergyIntoleranceCode = "6592009"
	AllergyIntoleranceCode6602005    AllergyIntoleranceCode = "6602005"
	AllergyIntoleranceCode6611005    AllergyIntoleranceCode = "6611005"
	AllergyIntoleranceCode6612003    AllergyIntoleranceCode = "6612003"
	AllergyIntoleranceCode6619007    AllergyIntoleranceCode = "6619007"
	AllergyIntoleranceCode6642000    AllergyIntoleranceCode = "6642000"
	AllergyIntoleranceCode6644004    AllergyIntoleranceCode = "6644004"
	AllergyIntoleranceCode6671004    AllergyIntoleranceCode = "6671004"
	AllergyIntoleranceCode6672006    AllergyIntoleranceCode = "6672006"
	AllergyIntoleranceCode6699008    AllergyIntoleranceCode = "6699008"
	AllergyIntoleranceCode6701008    AllergyIntoleranceCode = "6701008"
	AllergyIntoleranceCode6702001    AllergyIntoleranceCode = "6702001"
	AllergyIntoleranceCode6709005    AllergyIntoleranceCode = "6709005"
	AllergyIntoleranceCode6710000    AllergyIntoleranceCode = "6710000"
	AllergyIntoleranceCode6713003    AllergyIntoleranceCode = "6713003"
	AllergyIntoleranceCode6717002    AllergyIntoleranceCode = "6717002"
	AllergyIntoleranceCode6725000    AllergyIntoleranceCode = "6725000"
	AllergyIntoleranceCode6730001    AllergyIntoleranceCode = "6730001"
	AllergyIntoleranceCode6741004    AllergyIntoleranceCode = "6741004"
	AllergyIntoleranceCode6755007    AllergyIntoleranceCode = "6755007"
	AllergyIntoleranceCode6786001    AllergyIntoleranceCode = "6786001"
	AllergyIntoleranceCode6790004    AllergyIntoleranceCode = "6790004"
	AllergyIntoleranceCode6792007    AllergyIntoleranceCode = "6792007"
	AllergyIntoleranceCode6808006    AllergyIntoleranceCode = "6808006"
	AllergyIntoleranceCode6809003    AllergyIntoleranceCode = "6809003"
	AllergyIntoleranceCode6814004    AllergyIntoleranceCode = "6814004"
	AllergyIntoleranceCode6817006    AllergyIntoleranceCode = "6817006"
	AllergyIntoleranceCode6826009    AllergyIntoleranceCode = "6826009"
	AllergyIntoleranceCode6837005    AllergyIntoleranceCode = "6837005"
	AllergyIntoleranceCode6854002    AllergyIntoleranceCode = "6854002"
	AllergyIntoleranceCode6865007    AllergyIntoleranceCode = "6865007"
	AllergyIntoleranceCode6873003    AllergyIntoleranceCode = "6873003"
	AllergyIntoleranceCode6879004    AllergyIntoleranceCode = "6879004"
	AllergyIntoleranceCode6881002    AllergyIntoleranceCode = "6881002"
	AllergyIntoleranceCode6884005    AllergyIntoleranceCode = "6884005"
	AllergyIntoleranceCode6890009    AllergyIntoleranceCode = "6890009"
	AllergyIntoleranceCode6896003    AllergyIntoleranceCode = "6896003"
	AllergyIntoleranceCode6910009    AllergyIntoleranceCode = "6910009"
	AllergyIntoleranceCode6911008    AllergyIntoleranceCode = "6911008"
	AllergyIntoleranceCode6916003    AllergyIntoleranceCode = "6916003"
	AllergyIntoleranceCode6924008    AllergyIntoleranceCode = "6924008"
	AllergyIntoleranceCode6925009    AllergyIntoleranceCode = "6925009"
	AllergyIntoleranceCode6927001    AllergyIntoleranceCode = "6927001"
	AllergyIntoleranceCode6937006    AllergyIntoleranceCode = "6937006"
	AllergyIntoleranceCode6945001    AllergyIntoleranceCode = "6945001"
	AllergyIntoleranceCode6952004    AllergyIntoleranceCode = "6952004"
	AllergyIntoleranceCode6958000    AllergyIntoleranceCode = "6958000"
	AllergyIntoleranceCode6961004    AllergyIntoleranceCode = "6961004"
	AllergyIntoleranceCode6970001    AllergyIntoleranceCode = "6970001"
	AllergyIntoleranceCode6973004    AllergyIntoleranceCode = "6973004"
	AllergyIntoleranceCode6983000    AllergyIntoleranceCode = "6983000"
	AllergyIntoleranceCode6992002    AllergyIntoleranceCode = "6992002"
	AllergyIntoleranceCode6993007    AllergyIntoleranceCode = "6993007"
	AllergyIntoleranceCode6999006    AllergyIntoleranceCode = "6999006"
	AllergyIntoleranceCode7008002    AllergyIntoleranceCode = "7008002"
	AllergyIntoleranceCode7018007    AllergyIntoleranceCode = "7018007"
	AllergyIntoleranceCode7029006    AllergyIntoleranceCode = "7029006"
	AllergyIntoleranceCode7030001    AllergyIntoleranceCode = "7030001"
	AllergyIntoleranceCode7034005    AllergyIntoleranceCode = "7034005"
	AllergyIntoleranceCode7045008    AllergyIntoleranceCode = "7045008"
	AllergyIntoleranceCode7047000    AllergyIntoleranceCode = "7047000"
	AllergyIntoleranceCode7049002    AllergyIntoleranceCode = "7049002"
	AllergyIntoleranceCode7054006    AllergyIntoleranceCode = "7054006"
	AllergyIntoleranceCode7056008    AllergyIntoleranceCode = "7056008"
	AllergyIntoleranceCode7059001    AllergyIntoleranceCode = "7059001"
	AllergyIntoleranceCode7061005    AllergyIntoleranceCode = "7061005"
	AllergyIntoleranceCode7070008    AllergyIntoleranceCode = "7070008"
	AllergyIntoleranceCode7084003    AllergyIntoleranceCode = "7084003"
	AllergyIntoleranceCode7110002    AllergyIntoleranceCode = "7110002"
	AllergyIntoleranceCode7120007    AllergyIntoleranceCode = "7120007"
	AllergyIntoleranceCode7132006    AllergyIntoleranceCode = "7132006"
	AllergyIntoleranceCode7139002    AllergyIntoleranceCode = "7139002"
	AllergyIntoleranceCode7146006    AllergyIntoleranceCode = "7146006"
	AllergyIntoleranceCode7152007    AllergyIntoleranceCode = "7152007"
	AllergyIntoleranceCode7156005    AllergyIntoleranceCode = "7156005"
	AllergyIntoleranceCode7158006    AllergyIntoleranceCode = "7158006"
	AllergyIntoleranceCode7161007    AllergyIntoleranceCode = "7161007"
	AllergyIntoleranceCode7179006    AllergyIntoleranceCode = "7179006"
	AllergyIntoleranceCode7191002    AllergyIntoleranceCode = "7191002"
	AllergyIntoleranceCode7208009    AllergyIntoleranceCode = "7208009"
	AllergyIntoleranceCode7211005    AllergyIntoleranceCode = "7211005"
	AllergyIntoleranceCode7237008    AllergyIntoleranceCode = "7237008"
	AllergyIntoleranceCode7243005    AllergyIntoleranceCode = "7243005"
	AllergyIntoleranceCode7269004    AllergyIntoleranceCode = "7269004"
	AllergyIntoleranceCode7271004    AllergyIntoleranceCode = "7271004"
	AllergyIntoleranceCode7280004    AllergyIntoleranceCode = "7280004"
	AllergyIntoleranceCode7281000    AllergyIntoleranceCode = "7281000"
	AllergyIntoleranceCode7284008    AllergyIntoleranceCode = "7284008"
	AllergyIntoleranceCode7294003    AllergyIntoleranceCode = "7294003"
	AllergyIntoleranceCode7302008    AllergyIntoleranceCode = "7302008"
	AllergyIntoleranceCode7318002    AllergyIntoleranceCode = "7318002"
	AllergyIntoleranceCode7321000    AllergyIntoleranceCode = "7321000"
	AllergyIntoleranceCode7325009    AllergyIntoleranceCode = "7325009"
	AllergyIntoleranceCode7327001    AllergyIntoleranceCode = "7327001"
	AllergyIntoleranceCode7328006    AllergyIntoleranceCode = "7328006"
	AllergyIntoleranceCode7330008    AllergyIntoleranceCode = "7330008"
	AllergyIntoleranceCode7337006    AllergyIntoleranceCode = "7337006"
	AllergyIntoleranceCode7348004    AllergyIntoleranceCode = "7348004"
	AllergyIntoleranceCode7382005    AllergyIntoleranceCode = "7382005"
	AllergyIntoleranceCode7401000    AllergyIntoleranceCode = "7401000"
	AllergyIntoleranceCode7411007    AllergyIntoleranceCode = "7411007"
	AllergyIntoleranceCode7427000    AllergyIntoleranceCode = "7427000"
	AllergyIntoleranceCode7434003    AllergyIntoleranceCode = "7434003"
	AllergyIntoleranceCode7446004    AllergyIntoleranceCode = "7446004"
	AllergyIntoleranceCode7460002    AllergyIntoleranceCode = "7460002"
	AllergyIntoleranceCode7470000    AllergyIntoleranceCode = "7470000"
	AllergyIntoleranceCode7489000    AllergyIntoleranceCode = "7489000"
	AllergyIntoleranceCode7503004    AllergyIntoleranceCode = "7503004"
	AllergyIntoleranceCode7509000    AllergyIntoleranceCode = "7509000"
	AllergyIntoleranceCode7515000    AllergyIntoleranceCode = "7515000"
	AllergyIntoleranceCode7537007    AllergyIntoleranceCode = "7537007"
	AllergyIntoleranceCode7547005    AllergyIntoleranceCode = "7547005"
	AllergyIntoleranceCode7549008    AllergyIntoleranceCode = "7549008"
	AllergyIntoleranceCode7588005    AllergyIntoleranceCode = "7588005"
	AllergyIntoleranceCode7608003    AllergyIntoleranceCode = "7608003"
	AllergyIntoleranceCode7616007    AllergyIntoleranceCode = "7616007"
	AllergyIntoleranceCode7648006    AllergyIntoleranceCode = "7648006"
	AllergyIntoleranceCode7661006    AllergyIntoleranceCode = "7661006"
	AllergyIntoleranceCode7670009    AllergyIntoleranceCode = "7670009"
	AllergyIntoleranceCode7675004    AllergyIntoleranceCode = "7675004"
	AllergyIntoleranceCode7685003    AllergyIntoleranceCode = "7685003"
	AllergyIntoleranceCode7696006    AllergyIntoleranceCode = "7696006"
	AllergyIntoleranceCode7716001    AllergyIntoleranceCode = "7716001"
	AllergyIntoleranceCode7737009    AllergyIntoleranceCode = "7737009"
	AllergyIntoleranceCode7738004    AllergyIntoleranceCode = "7738004"
	AllergyIntoleranceCode7761002    AllergyIntoleranceCode = "7761002"
	AllergyIntoleranceCode7770004    AllergyIntoleranceCode = "7770004"
	AllergyIntoleranceCode7774008    AllergyIntoleranceCode = "7774008"
	AllergyIntoleranceCode7779003    AllergyIntoleranceCode = "7779003"
	AllergyIntoleranceCode7785005    AllergyIntoleranceCode = "7785005"
	AllergyIntoleranceCode7790008    AllergyIntoleranceCode = "7790008"
	AllergyIntoleranceCode7791007    AllergyIntoleranceCode = "7791007"
	AllergyIntoleranceCode7795003    AllergyIntoleranceCode = "7795003"
	AllergyIntoleranceCode7801007    AllergyIntoleranceCode = "7801007"
	AllergyIntoleranceCode7816005    AllergyIntoleranceCode = "7816005"
	AllergyIntoleranceCode7834009    AllergyIntoleranceCode = "7834009"
	AllergyIntoleranceCode7846008    AllergyIntoleranceCode = "7846008"
	AllergyIntoleranceCode7848009    AllergyIntoleranceCode = "7848009"
	AllergyIntoleranceCode7868003    AllergyIntoleranceCode = "7868003"
	AllergyIntoleranceCode7879008    AllergyIntoleranceCode = "7879008"
	AllergyIntoleranceCode7900007    AllergyIntoleranceCode = "7900007"
	AllergyIntoleranceCode7904003    AllergyIntoleranceCode = "7904003"
	AllergyIntoleranceCode7909008    AllergyIntoleranceCode = "7909008"
	AllergyIntoleranceCode7924004    AllergyIntoleranceCode = "7924004"
	AllergyIntoleranceCode7938006    AllergyIntoleranceCode = "7938006"
	AllergyIntoleranceCode7945006    AllergyIntoleranceCode = "7945006"
	AllergyIntoleranceCode7948008    AllergyIntoleranceCode = "7948008"
	AllergyIntoleranceCode7953003    AllergyIntoleranceCode = "7953003"
	AllergyIntoleranceCode7957002    AllergyIntoleranceCode = "7957002"
	AllergyIntoleranceCode7961008    AllergyIntoleranceCode = "7961008"
	AllergyIntoleranceCode7970006    AllergyIntoleranceCode = "7970006"
	AllergyIntoleranceCode7974002    AllergyIntoleranceCode = "7974002"
	AllergyIntoleranceCode7975001    AllergyIntoleranceCode = "7975001"
	AllergyIntoleranceCode7979007    AllergyIntoleranceCode = "7979007"
	AllergyIntoleranceCode7983007    AllergyIntoleranceCode = "7983007"
	AllergyIntoleranceCode7985000    AllergyIntoleranceCode = "7985000"
	AllergyIntoleranceCode7997004    AllergyIntoleranceCode = "7997004"
	AllergyIntoleranceCode8000007    AllergyIntoleranceCode = "8000007"
	AllergyIntoleranceCode8002004    AllergyIntoleranceCode = "8002004"
	AllergyIntoleranceCode8025003    AllergyIntoleranceCode = "8025003"
	AllergyIntoleranceCode8029009    AllergyIntoleranceCode = "8029009"
	AllergyIntoleranceCode8030004    AllergyIntoleranceCode = "8030004"
	AllergyIntoleranceCode8035009    AllergyIntoleranceCode = "8035009"
	AllergyIntoleranceCode8048008    AllergyIntoleranceCode = "8048008"
	AllergyIntoleranceCode8054009    AllergyIntoleranceCode = "8054009"
	AllergyIntoleranceCode8055005    AllergyIntoleranceCode = "8055005"
	AllergyIntoleranceCode8105004    AllergyIntoleranceCode = "8105004"
	AllergyIntoleranceCode8108002    AllergyIntoleranceCode = "8108002"
	AllergyIntoleranceCode8123007    AllergyIntoleranceCode = "8123007"
	AllergyIntoleranceCode8132009    AllergyIntoleranceCode = "8132009"
	AllergyIntoleranceCode8143001    AllergyIntoleranceCode = "8143001"
	AllergyIntoleranceCode8153000    AllergyIntoleranceCode = "8153000"
	AllergyIntoleranceCode8156008    AllergyIntoleranceCode = "8156008"
	AllergyIntoleranceCode8164002    AllergyIntoleranceCode = "8164002"
	AllergyIntoleranceCode8168004    AllergyIntoleranceCode = "8168004"
	AllergyIntoleranceCode8179009    AllergyIntoleranceCode = "8179009"
	AllergyIntoleranceCode8184003    AllergyIntoleranceCode = "8184003"
	AllergyIntoleranceCode8190004    AllergyIntoleranceCode = "8190004"
	AllergyIntoleranceCode8202008    AllergyIntoleranceCode = "8202008"
	AllergyIntoleranceCode8203003    AllergyIntoleranceCode = "8203003"
	AllergyIntoleranceCode8204009    AllergyIntoleranceCode = "8204009"
	AllergyIntoleranceCode8222007    AllergyIntoleranceCode = "8222007"
	AllergyIntoleranceCode8227001    AllergyIntoleranceCode = "8227001"
	AllergyIntoleranceCode8230008    AllergyIntoleranceCode = "8230008"
	AllergyIntoleranceCode8237006    AllergyIntoleranceCode = "8237006"
	AllergyIntoleranceCode8252004    AllergyIntoleranceCode = "8252004"
	AllergyIntoleranceCode8257005    AllergyIntoleranceCode = "8257005"
	AllergyIntoleranceCode8261004    AllergyIntoleranceCode = "8261004"
	AllergyIntoleranceCode8268005    AllergyIntoleranceCode = "8268005"
	AllergyIntoleranceCode8270001    AllergyIntoleranceCode = "8270001"
	AllergyIntoleranceCode8275006    AllergyIntoleranceCode = "8275006"
	AllergyIntoleranceCode8295000    AllergyIntoleranceCode = "8295000"
	AllergyIntoleranceCode8300003    AllergyIntoleranceCode = "8300003"
	AllergyIntoleranceCode8310007    AllergyIntoleranceCode = "8310007"
	AllergyIntoleranceCode8313009    AllergyIntoleranceCode = "8313009"
	AllergyIntoleranceCode8340009    AllergyIntoleranceCode = "8340009"
	AllergyIntoleranceCode8342001    AllergyIntoleranceCode = "8342001"
	AllergyIntoleranceCode8343006    AllergyIntoleranceCode = "8343006"
	AllergyIntoleranceCode8354001    AllergyIntoleranceCode = "8354001"
	AllergyIntoleranceCode8355000    AllergyIntoleranceCode = "8355000"
	AllergyIntoleranceCode8362009    AllergyIntoleranceCode = "8362009"
	AllergyIntoleranceCode8365006    AllergyIntoleranceCode = "8365006"
	AllergyIntoleranceCode8368008    AllergyIntoleranceCode = "8368008"
	AllergyIntoleranceCode8376005    AllergyIntoleranceCode = "8376005"
	AllergyIntoleranceCode8385005    AllergyIntoleranceCode = "8385005"
	AllergyIntoleranceCode8397006    AllergyIntoleranceCode = "8397006"
	AllergyIntoleranceCode8406008    AllergyIntoleranceCode = "8406008"
	AllergyIntoleranceCode8429000    AllergyIntoleranceCode = "8429000"
	AllergyIntoleranceCode8450009    AllergyIntoleranceCode = "8450009"
	AllergyIntoleranceCode8452001    AllergyIntoleranceCode = "8452001"
	AllergyIntoleranceCode8456003    AllergyIntoleranceCode = "8456003"
	AllergyIntoleranceCode8460000    AllergyIntoleranceCode = "8460000"
	AllergyIntoleranceCode8473001    AllergyIntoleranceCode = "8473001"
	AllergyIntoleranceCode8474007    AllergyIntoleranceCode = "8474007"
	AllergyIntoleranceCode8484008    AllergyIntoleranceCode = "8484008"
	AllergyIntoleranceCode8485009    AllergyIntoleranceCode = "8485009"
	AllergyIntoleranceCode8486005    AllergyIntoleranceCode = "8486005"
	AllergyIntoleranceCode8487001    AllergyIntoleranceCode = "8487001"
	AllergyIntoleranceCode8491006    AllergyIntoleranceCode = "8491006"
	AllergyIntoleranceCode8492004    AllergyIntoleranceCode = "8492004"
	AllergyIntoleranceCode8498000    AllergyIntoleranceCode = "8498000"
	AllergyIntoleranceCode8507001    AllergyIntoleranceCode = "8507001"
	AllergyIntoleranceCode8514004    AllergyIntoleranceCode = "8514004"
	AllergyIntoleranceCode8520003    AllergyIntoleranceCode = "8520003"
	AllergyIntoleranceCode8525008    AllergyIntoleranceCode = "8525008"
	AllergyIntoleranceCode8529002    AllergyIntoleranceCode = "8529002"
	AllergyIntoleranceCode8534003    AllergyIntoleranceCode = "8534003"
	AllergyIntoleranceCode8537005    AllergyIntoleranceCode = "8537005"
	AllergyIntoleranceCode8578007    AllergyIntoleranceCode = "8578007"
	AllergyIntoleranceCode8591008    AllergyIntoleranceCode = "8591008"
	AllergyIntoleranceCode8612007    AllergyIntoleranceCode = "8612007"
	AllergyIntoleranceCode8620009    AllergyIntoleranceCode = "8620009"
	AllergyIntoleranceCode8631001    AllergyIntoleranceCode = "8631001"
	AllergyIntoleranceCode8653004    AllergyIntoleranceCode = "8653004"
	AllergyIntoleranceCode8660005    AllergyIntoleranceCode = "8660005"
	AllergyIntoleranceCode8687009    AllergyIntoleranceCode = "8687009"
	AllergyIntoleranceCode8689007    AllergyIntoleranceCode = "8689007"
	AllergyIntoleranceCode8701002    AllergyIntoleranceCode = "8701002"
	AllergyIntoleranceCode8705006    AllergyIntoleranceCode = "8705006"
	AllergyIntoleranceCode8731008    AllergyIntoleranceCode = "8731008"
	AllergyIntoleranceCode8740007    AllergyIntoleranceCode = "8740007"
	AllergyIntoleranceCode8761000    AllergyIntoleranceCode = "8761000"
	AllergyIntoleranceCode8767001    AllergyIntoleranceCode = "8767001"
	AllergyIntoleranceCode8785008    AllergyIntoleranceCode = "8785008"
	AllergyIntoleranceCode8786009    AllergyIntoleranceCode = "8786009"
	AllergyIntoleranceCode8795001    AllergyIntoleranceCode = "8795001"
	AllergyIntoleranceCode8817004    AllergyIntoleranceCode = "8817004"
	AllergyIntoleranceCode8818009    AllergyIntoleranceCode = "8818009"
	AllergyIntoleranceCode8822004    AllergyIntoleranceCode = "8822004"
	AllergyIntoleranceCode8830003    AllergyIntoleranceCode = "8830003"
	AllergyIntoleranceCode8836009    AllergyIntoleranceCode = "8836009"
	AllergyIntoleranceCode8844009    AllergyIntoleranceCode = "8844009"
	AllergyIntoleranceCode8858006    AllergyIntoleranceCode = "8858006"
	AllergyIntoleranceCode8865003    AllergyIntoleranceCode = "8865003"
	AllergyIntoleranceCode8878003    AllergyIntoleranceCode = "8878003"
	AllergyIntoleranceCode8882001    AllergyIntoleranceCode = "8882001"
	AllergyIntoleranceCode8886003    AllergyIntoleranceCode = "8886003"
	AllergyIntoleranceCode8908003    AllergyIntoleranceCode = "8908003"
	AllergyIntoleranceCode8914005    AllergyIntoleranceCode = "8914005"
	AllergyIntoleranceCode8919000    AllergyIntoleranceCode = "8919000"
	AllergyIntoleranceCode8926000    AllergyIntoleranceCode = "8926000"
	AllergyIntoleranceCode8945009    AllergyIntoleranceCode = "8945009"
	AllergyIntoleranceCode8953001    AllergyIntoleranceCode = "8953001"
	AllergyIntoleranceCode8963009    AllergyIntoleranceCode = "8963009"
	AllergyIntoleranceCode8969008    AllergyIntoleranceCode = "8969008"
	AllergyIntoleranceCode8977007    AllergyIntoleranceCode = "8977007"
	AllergyIntoleranceCode8982000    AllergyIntoleranceCode = "8982000"
	AllergyIntoleranceCode8987006    AllergyIntoleranceCode = "8987006"
	AllergyIntoleranceCode8991001    AllergyIntoleranceCode = "8991001"
	AllergyIntoleranceCode9010006    AllergyIntoleranceCode = "9010006"
	AllergyIntoleranceCode9013008    AllergyIntoleranceCode = "9013008"
	AllergyIntoleranceCode9021002    AllergyIntoleranceCode = "9021002"
	AllergyIntoleranceCode9024005    AllergyIntoleranceCode = "9024005"
	AllergyIntoleranceCode9045003    AllergyIntoleranceCode = "9045003"
	AllergyIntoleranceCode9052001    AllergyIntoleranceCode = "9052001"
	AllergyIntoleranceCode9054000    AllergyIntoleranceCode = "9054000"
	AllergyIntoleranceCode9103003    AllergyIntoleranceCode = "9103003"
	AllergyIntoleranceCode9110009    AllergyIntoleranceCode = "9110009"
	AllergyIntoleranceCode9125009    AllergyIntoleranceCode = "9125009"
	AllergyIntoleranceCode9159008    AllergyIntoleranceCode = "9159008"
	AllergyIntoleranceCode9172009    AllergyIntoleranceCode = "9172009"
	AllergyIntoleranceCode9174005    AllergyIntoleranceCode = "9174005"
	AllergyIntoleranceCode9183000    AllergyIntoleranceCode = "9183000"
	AllergyIntoleranceCode9189001    AllergyIntoleranceCode = "9189001"
	AllergyIntoleranceCode9195000    AllergyIntoleranceCode = "9195000"
	AllergyIntoleranceCode9197008    AllergyIntoleranceCode = "9197008"
	AllergyIntoleranceCode9205004    AllergyIntoleranceCode = "9205004"
	AllergyIntoleranceCode9220005    AllergyIntoleranceCode = "9220005"
	AllergyIntoleranceCode9223007    AllergyIntoleranceCode = "9223007"
	AllergyIntoleranceCode9234005    AllergyIntoleranceCode = "9234005"
	AllergyIntoleranceCode9246009    AllergyIntoleranceCode = "9246009"
	AllergyIntoleranceCode9253000    AllergyIntoleranceCode = "9253000"
	AllergyIntoleranceCode9270008    AllergyIntoleranceCode = "9270008"
	AllergyIntoleranceCode9271007    AllergyIntoleranceCode = "9271007"
	AllergyIntoleranceCode9301005    AllergyIntoleranceCode = "9301005"
	AllergyIntoleranceCode9302003    AllergyIntoleranceCode = "9302003"
	AllergyIntoleranceCode9315007    AllergyIntoleranceCode = "9315007"
	AllergyIntoleranceCode9319001    AllergyIntoleranceCode = "9319001"
	AllergyIntoleranceCode9334007    AllergyIntoleranceCode = "9334007"
	AllergyIntoleranceCode9349004    AllergyIntoleranceCode = "9349004"
	AllergyIntoleranceCode9351000    AllergyIntoleranceCode = "9351000"
	AllergyIntoleranceCode9355009    AllergyIntoleranceCode = "9355009"
	AllergyIntoleranceCode9392009    AllergyIntoleranceCode = "9392009"
	AllergyIntoleranceCode9396007    AllergyIntoleranceCode = "9396007"
	AllergyIntoleranceCode9398008    AllergyIntoleranceCode = "9398008"
	AllergyIntoleranceCode9410003    AllergyIntoleranceCode = "9410003"
	AllergyIntoleranceCode9422000    AllergyIntoleranceCode = "9422000"
	AllergyIntoleranceCode418038007  AllergyIntoleranceCode = "418038007"
	AllergyIntoleranceCode25744000   AllergyIntoleranceCode = "25744000"
	AllergyIntoleranceCode54250004   AllergyIntoleranceCode = "54250004"
	AllergyIntoleranceCode59037007   AllergyIntoleranceCode = "59037007"
	AllergyIntoleranceCode61712006   AllergyIntoleranceCode = "61712006"
	AllergyIntoleranceCode72354005   AllergyIntoleranceCode = "72354005"
	AllergyIntoleranceCode91931000   AllergyIntoleranceCode = "91931000"
	AllergyIntoleranceCode91932007   AllergyIntoleranceCode = "91932007"
	AllergyIntoleranceCode91934008   AllergyIntoleranceCode = "91934008"
	AllergyIntoleranceCode91935009   AllergyIntoleranceCode = "91935009"
	AllergyIntoleranceCode91936005   AllergyIntoleranceCode = "91936005"
	AllergyIntoleranceCode91937001   AllergyIntoleranceCode = "91937001"
	AllergyIntoleranceCode91938006   AllergyIntoleranceCode = "91938006"
	AllergyIntoleranceCode91939003   AllergyIntoleranceCode = "91939003"
	AllergyIntoleranceCode91940001   AllergyIntoleranceCode = "91940001"
	AllergyIntoleranceCode190751001  AllergyIntoleranceCode = "190751001"
	AllergyIntoleranceCode213020009  AllergyIntoleranceCode = "213020009"
	AllergyIntoleranceCode232349006  AllergyIntoleranceCode = "232349006"
	AllergyIntoleranceCode232350006  AllergyIntoleranceCode = "232350006"
	AllergyIntoleranceCode235719002  AllergyIntoleranceCode = "235719002"
	AllergyIntoleranceCode293584003  AllergyIntoleranceCode = "293584003"
	AllergyIntoleranceCode293585002  AllergyIntoleranceCode = "293585002"
	AllergyIntoleranceCode293586001  AllergyIntoleranceCode = "293586001"
	AllergyIntoleranceCode293588000  AllergyIntoleranceCode = "293588000"
	AllergyIntoleranceCode293589008  AllergyIntoleranceCode = "293589008"
	AllergyIntoleranceCode293590004  AllergyIntoleranceCode = "293590004"
	AllergyIntoleranceCode293591000  AllergyIntoleranceCode = "293591000"
	AllergyIntoleranceCode293592007  AllergyIntoleranceCode = "293592007"
	AllergyIntoleranceCode293593002  AllergyIntoleranceCode = "293593002"
	AllergyIntoleranceCode293594008  AllergyIntoleranceCode = "293594008"
	AllergyIntoleranceCode293595009  AllergyIntoleranceCode = "293595009"
	AllergyIntoleranceCode293596005  AllergyIntoleranceCode = "293596005"
	AllergyIntoleranceCode293597001  AllergyIntoleranceCode = "293597001"
	AllergyIntoleranceCode293598006  AllergyIntoleranceCode = "293598006"
	AllergyIntoleranceCode293599003  AllergyIntoleranceCode = "293599003"
	AllergyIntoleranceCode293600000  AllergyIntoleranceCode = "293600000"
	AllergyIntoleranceCode293601001  AllergyIntoleranceCode = "293601001"
	AllergyIntoleranceCode293602008  AllergyIntoleranceCode = "293602008"
	AllergyIntoleranceCode293603003  AllergyIntoleranceCode = "293603003"
	AllergyIntoleranceCode293604009  AllergyIntoleranceCode = "293604009"
	AllergyIntoleranceCode293605005  AllergyIntoleranceCode = "293605005"
	AllergyIntoleranceCode293606006  AllergyIntoleranceCode = "293606006"
	AllergyIntoleranceCode293607002  AllergyIntoleranceCode = "293607002"
	AllergyIntoleranceCode293608007  AllergyIntoleranceCode = "293608007"
	AllergyIntoleranceCode293609004  AllergyIntoleranceCode = "293609004"
	AllergyIntoleranceCode293610009  AllergyIntoleranceCode = "293610009"
	AllergyIntoleranceCode293611008  AllergyIntoleranceCode = "293611008"
	AllergyIntoleranceCode293612001  AllergyIntoleranceCode = "293612001"
	AllergyIntoleranceCode293613006  AllergyIntoleranceCode = "293613006"
	AllergyIntoleranceCode293614000  AllergyIntoleranceCode = "293614000"
	AllergyIntoleranceCode293615004  AllergyIntoleranceCode = "293615004"
	AllergyIntoleranceCode293616003  AllergyIntoleranceCode = "293616003"
	AllergyIntoleranceCode293617007  AllergyIntoleranceCode = "293617007"
	AllergyIntoleranceCode293618002  AllergyIntoleranceCode = "293618002"
	AllergyIntoleranceCode293619005  AllergyIntoleranceCode = "293619005"
	AllergyIntoleranceCode293620004  AllergyIntoleranceCode = "293620004"
	AllergyIntoleranceCode293621000  AllergyIntoleranceCode = "293621000"
	AllergyIntoleranceCode293622007  AllergyIntoleranceCode = "293622007"
	AllergyIntoleranceCode293623002  AllergyIntoleranceCode = "293623002"
	AllergyIntoleranceCode293624008  AllergyIntoleranceCode = "293624008"
	AllergyIntoleranceCode293625009  AllergyIntoleranceCode = "293625009"
	AllergyIntoleranceCode293626005  AllergyIntoleranceCode = "293626005"
	AllergyIntoleranceCode293627001  AllergyIntoleranceCode = "293627001"
	AllergyIntoleranceCode293628006  AllergyIntoleranceCode = "293628006"
	AllergyIntoleranceCode293629003  AllergyIntoleranceCode = "293629003"
	AllergyIntoleranceCode293630008  AllergyIntoleranceCode = "293630008"
	AllergyIntoleranceCode293631007  AllergyIntoleranceCode = "293631007"
	AllergyIntoleranceCode293632000  AllergyIntoleranceCode = "293632000"
	AllergyIntoleranceCode293633005  AllergyIntoleranceCode = "293633005"
	AllergyIntoleranceCode293635003  AllergyIntoleranceCode = "293635003"
	AllergyIntoleranceCode293636002  AllergyIntoleranceCode = "293636002"
	AllergyIntoleranceCode293637006  AllergyIntoleranceCode = "293637006"
	AllergyIntoleranceCode293645001  AllergyIntoleranceCode = "293645001"
	AllergyIntoleranceCode293646000  AllergyIntoleranceCode = "293646000"
	AllergyIntoleranceCode293647009  AllergyIntoleranceCode = "293647009"
	AllergyIntoleranceCode293648004  AllergyIntoleranceCode = "293648004"
	AllergyIntoleranceCode293649007  AllergyIntoleranceCode = "293649007"
	AllergyIntoleranceCode293650007  AllergyIntoleranceCode = "293650007"
	AllergyIntoleranceCode293651006  AllergyIntoleranceCode = "293651006"
	AllergyIntoleranceCode293652004  AllergyIntoleranceCode = "293652004"
	AllergyIntoleranceCode293653009  AllergyIntoleranceCode = "293653009"
	AllergyIntoleranceCode293654003  AllergyIntoleranceCode = "293654003"
	AllergyIntoleranceCode293655002  AllergyIntoleranceCode = "293655002"
	AllergyIntoleranceCode293656001  AllergyIntoleranceCode = "293656001"
	AllergyIntoleranceCode293657005  AllergyIntoleranceCode = "293657005"
	AllergyIntoleranceCode293658000  AllergyIntoleranceCode = "293658000"
	AllergyIntoleranceCode293659008  AllergyIntoleranceCode = "293659008"
	AllergyIntoleranceCode293660003  AllergyIntoleranceCode = "293660003"
	AllergyIntoleranceCode293662006  AllergyIntoleranceCode = "293662006"
	AllergyIntoleranceCode293663001  AllergyIntoleranceCode = "293663001"
	AllergyIntoleranceCode293665008  AllergyIntoleranceCode = "293665008"
	AllergyIntoleranceCode293666009  AllergyIntoleranceCode = "293666009"
	AllergyIntoleranceCode293668005  AllergyIntoleranceCode = "293668005"
	AllergyIntoleranceCode293669002  AllergyIntoleranceCode = "293669002"
	AllergyIntoleranceCode293671002  AllergyIntoleranceCode = "293671002"
	AllergyIntoleranceCode293673004  AllergyIntoleranceCode = "293673004"
	AllergyIntoleranceCode293674005  AllergyIntoleranceCode = "293674005"
	AllergyIntoleranceCode293675006  AllergyIntoleranceCode = "293675006"
	AllergyIntoleranceCode293676007  AllergyIntoleranceCode = "293676007"
	AllergyIntoleranceCode293678008  AllergyIntoleranceCode = "293678008"
	AllergyIntoleranceCode293679000  AllergyIntoleranceCode = "293679000"
	AllergyIntoleranceCode293680002  AllergyIntoleranceCode = "293680002"
	AllergyIntoleranceCode293681003  AllergyIntoleranceCode = "293681003"
	AllergyIntoleranceCode293682005  AllergyIntoleranceCode = "293682005"
	AllergyIntoleranceCode293685007  AllergyIntoleranceCode = "293685007"
	AllergyIntoleranceCode293686008  AllergyIntoleranceCode = "293686008"
	AllergyIntoleranceCode293687004  AllergyIntoleranceCode = "293687004"
	AllergyIntoleranceCode293690005  AllergyIntoleranceCode = "293690005"
	AllergyIntoleranceCode293691009  AllergyIntoleranceCode = "293691009"
	AllergyIntoleranceCode293692002  AllergyIntoleranceCode = "293692002"
	AllergyIntoleranceCode293693007  AllergyIntoleranceCode = "293693007"
	AllergyIntoleranceCode293694001  AllergyIntoleranceCode = "293694001"
	AllergyIntoleranceCode293695000  AllergyIntoleranceCode = "293695000"
	AllergyIntoleranceCode293696004  AllergyIntoleranceCode = "293696004"
	AllergyIntoleranceCode293697008  AllergyIntoleranceCode = "293697008"
	AllergyIntoleranceCode293699006  AllergyIntoleranceCode = "293699006"
	AllergyIntoleranceCode293700007  AllergyIntoleranceCode = "293700007"
	AllergyIntoleranceCode293701006  AllergyIntoleranceCode = "293701006"
	AllergyIntoleranceCode293706001  AllergyIntoleranceCode = "293706001"
	AllergyIntoleranceCode293707005  AllergyIntoleranceCode = "293707005"
	AllergyIntoleranceCode293708000  AllergyIntoleranceCode = "293708000"
	AllergyIntoleranceCode293709008  AllergyIntoleranceCode = "293709008"
	AllergyIntoleranceCode293710003  AllergyIntoleranceCode = "293710003"
	AllergyIntoleranceCode293712006  AllergyIntoleranceCode = "293712006"
	AllergyIntoleranceCode293714007  AllergyIntoleranceCode = "293714007"
	AllergyIntoleranceCode293715008  AllergyIntoleranceCode = "293715008"
	AllergyIntoleranceCode293716009  AllergyIntoleranceCode = "293716009"
	AllergyIntoleranceCode293717000  AllergyIntoleranceCode = "293717000"
	AllergyIntoleranceCode293718005  AllergyIntoleranceCode = "293718005"
	AllergyIntoleranceCode293719002  AllergyIntoleranceCode = "293719002"
	AllergyIntoleranceCode293720008  AllergyIntoleranceCode = "293720008"
	AllergyIntoleranceCode293721007  AllergyIntoleranceCode = "293721007"
	AllergyIntoleranceCode293722000  AllergyIntoleranceCode = "293722000"
	AllergyIntoleranceCode293723005  AllergyIntoleranceCode = "293723005"
	AllergyIntoleranceCode293724004  AllergyIntoleranceCode = "293724004"
	AllergyIntoleranceCode293725003  AllergyIntoleranceCode = "293725003"
	AllergyIntoleranceCode293726002  AllergyIntoleranceCode = "293726002"
	AllergyIntoleranceCode293727006  AllergyIntoleranceCode = "293727006"
	AllergyIntoleranceCode293728001  AllergyIntoleranceCode = "293728001"
	AllergyIntoleranceCode293732007  AllergyIntoleranceCode = "293732007"
	AllergyIntoleranceCode293733002  AllergyIntoleranceCode = "293733002"
	AllergyIntoleranceCode293735009  AllergyIntoleranceCode = "293735009"
	AllergyIntoleranceCode293736005  AllergyIntoleranceCode = "293736005"
	AllergyIntoleranceCode293737001  AllergyIntoleranceCode = "293737001"
	AllergyIntoleranceCode293738006  AllergyIntoleranceCode = "293738006"
	AllergyIntoleranceCode293740001  AllergyIntoleranceCode = "293740001"
	AllergyIntoleranceCode293741002  AllergyIntoleranceCode = "293741002"
	AllergyIntoleranceCode293742009  AllergyIntoleranceCode = "293742009"
	AllergyIntoleranceCode293743004  AllergyIntoleranceCode = "293743004"
	AllergyIntoleranceCode293745006  AllergyIntoleranceCode = "293745006"
	AllergyIntoleranceCode293746007  AllergyIntoleranceCode = "293746007"
	AllergyIntoleranceCode293747003  AllergyIntoleranceCode = "293747003"
	AllergyIntoleranceCode293748008  AllergyIntoleranceCode = "293748008"
	AllergyIntoleranceCode293749000  AllergyIntoleranceCode = "293749000"
	AllergyIntoleranceCode293750000  AllergyIntoleranceCode = "293750000"
	AllergyIntoleranceCode293751001  AllergyIntoleranceCode = "293751001"
	AllergyIntoleranceCode293752008  AllergyIntoleranceCode = "293752008"
	AllergyIntoleranceCode293753003  AllergyIntoleranceCode = "293753003"
	AllergyIntoleranceCode293754009  AllergyIntoleranceCode = "293754009"
	AllergyIntoleranceCode293755005  AllergyIntoleranceCode = "293755005"
	AllergyIntoleranceCode293756006  AllergyIntoleranceCode = "293756006"
	AllergyIntoleranceCode293758007  AllergyIntoleranceCode = "293758007"
	AllergyIntoleranceCode293760009  AllergyIntoleranceCode = "293760009"
	AllergyIntoleranceCode293761008  AllergyIntoleranceCode = "293761008"
	AllergyIntoleranceCode293762001  AllergyIntoleranceCode = "293762001"
	AllergyIntoleranceCode293763006  AllergyIntoleranceCode = "293763006"
	AllergyIntoleranceCode293764000  AllergyIntoleranceCode = "293764000"
	AllergyIntoleranceCode293765004  AllergyIntoleranceCode = "293765004"
	AllergyIntoleranceCode293766003  AllergyIntoleranceCode = "293766003"
	AllergyIntoleranceCode293767007  AllergyIntoleranceCode = "293767007"
	AllergyIntoleranceCode293768002  AllergyIntoleranceCode = "293768002"
	AllergyIntoleranceCode293770006  AllergyIntoleranceCode = "293770006"
	AllergyIntoleranceCode293771005  AllergyIntoleranceCode = "293771005"
	AllergyIntoleranceCode293772003  AllergyIntoleranceCode = "293772003"
	AllergyIntoleranceCode293773008  AllergyIntoleranceCode = "293773008"
	AllergyIntoleranceCode293774002  AllergyIntoleranceCode = "293774002"
	AllergyIntoleranceCode293775001  AllergyIntoleranceCode = "293775001"
	AllergyIntoleranceCode293776000  AllergyIntoleranceCode = "293776000"
	AllergyIntoleranceCode293777009  AllergyIntoleranceCode = "293777009"
	AllergyIntoleranceCode293778004  AllergyIntoleranceCode = "293778004"
	AllergyIntoleranceCode293779007  AllergyIntoleranceCode = "293779007"
	AllergyIntoleranceCode293780005  AllergyIntoleranceCode = "293780005"
	AllergyIntoleranceCode293781009  AllergyIntoleranceCode = "293781009"
	AllergyIntoleranceCode293782002  AllergyIntoleranceCode = "293782002"
	AllergyIntoleranceCode293783007  AllergyIntoleranceCode = "293783007"
	AllergyIntoleranceCode293784001  AllergyIntoleranceCode = "293784001"
	AllergyIntoleranceCode293785000  AllergyIntoleranceCode = "293785000"
	AllergyIntoleranceCode293786004  AllergyIntoleranceCode = "293786004"
	AllergyIntoleranceCode293787008  AllergyIntoleranceCode = "293787008"
	AllergyIntoleranceCode293788003  AllergyIntoleranceCode = "293788003"
	AllergyIntoleranceCode293789006  AllergyIntoleranceCode = "293789006"
	AllergyIntoleranceCode293790002  AllergyIntoleranceCode = "293790002"
	AllergyIntoleranceCode293791003  AllergyIntoleranceCode = "293791003"
	AllergyIntoleranceCode293792005  AllergyIntoleranceCode = "293792005"
	AllergyIntoleranceCode293793000  AllergyIntoleranceCode = "293793000"
	AllergyIntoleranceCode293794006  AllergyIntoleranceCode = "293794006"
	AllergyIntoleranceCode293795007  AllergyIntoleranceCode = "293795007"
	AllergyIntoleranceCode293796008  AllergyIntoleranceCode = "293796008"
	AllergyIntoleranceCode293798009  AllergyIntoleranceCode = "293798009"
	AllergyIntoleranceCode293799001  AllergyIntoleranceCode = "293799001"
	AllergyIntoleranceCode293802005  AllergyIntoleranceCode = "293802005"
	AllergyIntoleranceCode293803000  AllergyIntoleranceCode = "293803000"
	AllergyIntoleranceCode293804006  AllergyIntoleranceCode = "293804006"
	AllergyIntoleranceCode293805007  AllergyIntoleranceCode = "293805007"
	AllergyIntoleranceCode293806008  AllergyIntoleranceCode = "293806008"
	AllergyIntoleranceCode293808009  AllergyIntoleranceCode = "293808009"
	AllergyIntoleranceCode293811005  AllergyIntoleranceCode = "293811005"
	AllergyIntoleranceCode293812003  AllergyIntoleranceCode = "293812003"
	AllergyIntoleranceCode293813008  AllergyIntoleranceCode = "293813008"
	AllergyIntoleranceCode293814002  AllergyIntoleranceCode = "293814002"
	AllergyIntoleranceCode293815001  AllergyIntoleranceCode = "293815001"
	AllergyIntoleranceCode293818004  AllergyIntoleranceCode = "293818004"
	AllergyIntoleranceCode293819007  AllergyIntoleranceCode = "293819007"
	AllergyIntoleranceCode293822009  AllergyIntoleranceCode = "293822009"
	AllergyIntoleranceCode293823004  AllergyIntoleranceCode = "293823004"
	AllergyIntoleranceCode293824005  AllergyIntoleranceCode = "293824005"
	AllergyIntoleranceCode293825006  AllergyIntoleranceCode = "293825006"
	AllergyIntoleranceCode293826007  AllergyIntoleranceCode = "293826007"
	AllergyIntoleranceCode293827003  AllergyIntoleranceCode = "293827003"
	AllergyIntoleranceCode293828008  AllergyIntoleranceCode = "293828008"
	AllergyIntoleranceCode293829000  AllergyIntoleranceCode = "293829000"
	AllergyIntoleranceCode293830005  AllergyIntoleranceCode = "293830005"
	AllergyIntoleranceCode293831009  AllergyIntoleranceCode = "293831009"
	AllergyIntoleranceCode293832002  AllergyIntoleranceCode = "293832002"
	AllergyIntoleranceCode293833007  AllergyIntoleranceCode = "293833007"
	AllergyIntoleranceCode293834001  AllergyIntoleranceCode = "293834001"
	AllergyIntoleranceCode293835000  AllergyIntoleranceCode = "293835000"
	AllergyIntoleranceCode293836004  AllergyIntoleranceCode = "293836004"
	AllergyIntoleranceCode293837008  AllergyIntoleranceCode = "293837008"
	AllergyIntoleranceCode293838003  AllergyIntoleranceCode = "293838003"
	AllergyIntoleranceCode293839006  AllergyIntoleranceCode = "293839006"
	AllergyIntoleranceCode293840008  AllergyIntoleranceCode = "293840008"
	AllergyIntoleranceCode293842000  AllergyIntoleranceCode = "293842000"
	AllergyIntoleranceCode293843005  AllergyIntoleranceCode = "293843005"
	AllergyIntoleranceCode293844004  AllergyIntoleranceCode = "293844004"
	AllergyIntoleranceCode293845003  AllergyIntoleranceCode = "293845003"
	AllergyIntoleranceCode293847006  AllergyIntoleranceCode = "293847006"
	AllergyIntoleranceCode293848001  AllergyIntoleranceCode = "293848001"
	AllergyIntoleranceCode293849009  AllergyIntoleranceCode = "293849009"
	AllergyIntoleranceCode293850009  AllergyIntoleranceCode = "293850009"
	AllergyIntoleranceCode293851008  AllergyIntoleranceCode = "293851008"
	AllergyIntoleranceCode293853006  AllergyIntoleranceCode = "293853006"
	AllergyIntoleranceCode293854000  AllergyIntoleranceCode = "293854000"
	AllergyIntoleranceCode293855004  AllergyIntoleranceCode = "293855004"
	AllergyIntoleranceCode293856003  AllergyIntoleranceCode = "293856003"
	AllergyIntoleranceCode293858002  AllergyIntoleranceCode = "293858002"
	AllergyIntoleranceCode293859005  AllergyIntoleranceCode = "293859005"
	AllergyIntoleranceCode293860000  AllergyIntoleranceCode = "293860000"
	AllergyIntoleranceCode293861001  AllergyIntoleranceCode = "293861001"
	AllergyIntoleranceCode293864009  AllergyIntoleranceCode = "293864009"
	AllergyIntoleranceCode293865005  AllergyIntoleranceCode = "293865005"
	AllergyIntoleranceCode293866006  AllergyIntoleranceCode = "293866006"
	AllergyIntoleranceCode293867002  AllergyIntoleranceCode = "293867002"
	AllergyIntoleranceCode293868007  AllergyIntoleranceCode = "293868007"
	AllergyIntoleranceCode293869004  AllergyIntoleranceCode = "293869004"
	AllergyIntoleranceCode293870003  AllergyIntoleranceCode = "293870003"
	AllergyIntoleranceCode293871004  AllergyIntoleranceCode = "293871004"
	AllergyIntoleranceCode293874007  AllergyIntoleranceCode = "293874007"
	AllergyIntoleranceCode293875008  AllergyIntoleranceCode = "293875008"
	AllergyIntoleranceCode293876009  AllergyIntoleranceCode = "293876009"
	AllergyIntoleranceCode293877000  AllergyIntoleranceCode = "293877000"
	AllergyIntoleranceCode293878005  AllergyIntoleranceCode = "293878005"
	AllergyIntoleranceCode293880004  AllergyIntoleranceCode = "293880004"
	AllergyIntoleranceCode293881000  AllergyIntoleranceCode = "293881000"
	AllergyIntoleranceCode293882007  AllergyIntoleranceCode = "293882007"
	AllergyIntoleranceCode293884008  AllergyIntoleranceCode = "293884008"
	AllergyIntoleranceCode293886005  AllergyIntoleranceCode = "293886005"
	AllergyIntoleranceCode293887001  AllergyIntoleranceCode = "293887001"
	AllergyIntoleranceCode293888006  AllergyIntoleranceCode = "293888006"
	AllergyIntoleranceCode293889003  AllergyIntoleranceCode = "293889003"
	AllergyIntoleranceCode293890007  AllergyIntoleranceCode = "293890007"
	AllergyIntoleranceCode293891006  AllergyIntoleranceCode = "293891006"
	AllergyIntoleranceCode293892004  AllergyIntoleranceCode = "293892004"
	AllergyIntoleranceCode293893009  AllergyIntoleranceCode = "293893009"
	AllergyIntoleranceCode293894003  AllergyIntoleranceCode = "293894003"
	AllergyIntoleranceCode293895002  AllergyIntoleranceCode = "293895002"
	AllergyIntoleranceCode293897005  AllergyIntoleranceCode = "293897005"
	AllergyIntoleranceCode293898000  AllergyIntoleranceCode = "293898000"
	AllergyIntoleranceCode293899008  AllergyIntoleranceCode = "293899008"
	AllergyIntoleranceCode293900003  AllergyIntoleranceCode = "293900003"
	AllergyIntoleranceCode293901004  AllergyIntoleranceCode = "293901004"
	AllergyIntoleranceCode293902006  AllergyIntoleranceCode = "293902006"
	AllergyIntoleranceCode293903001  AllergyIntoleranceCode = "293903001"
	AllergyIntoleranceCode293904007  AllergyIntoleranceCode = "293904007"
	AllergyIntoleranceCode293906009  AllergyIntoleranceCode = "293906009"
	AllergyIntoleranceCode293908005  AllergyIntoleranceCode = "293908005"
	AllergyIntoleranceCode293909002  AllergyIntoleranceCode = "293909002"
	AllergyIntoleranceCode293911006  AllergyIntoleranceCode = "293911006"
	AllergyIntoleranceCode293912004  AllergyIntoleranceCode = "293912004"
	AllergyIntoleranceCode293914003  AllergyIntoleranceCode = "293914003"
	AllergyIntoleranceCode293915002  AllergyIntoleranceCode = "293915002"
	AllergyIntoleranceCode293916001  AllergyIntoleranceCode = "293916001"
	AllergyIntoleranceCode293917005  AllergyIntoleranceCode = "293917005"
	AllergyIntoleranceCode293918000  AllergyIntoleranceCode = "293918000"
	AllergyIntoleranceCode293919008  AllergyIntoleranceCode = "293919008"
	AllergyIntoleranceCode293920002  AllergyIntoleranceCode = "293920002"
	AllergyIntoleranceCode293921003  AllergyIntoleranceCode = "293921003"
	AllergyIntoleranceCode293923000  AllergyIntoleranceCode = "293923000"
	AllergyIntoleranceCode293924006  AllergyIntoleranceCode = "293924006"
	AllergyIntoleranceCode293925007  AllergyIntoleranceCode = "293925007"
	AllergyIntoleranceCode293926008  AllergyIntoleranceCode = "293926008"
	AllergyIntoleranceCode293927004  AllergyIntoleranceCode = "293927004"
	AllergyIntoleranceCode293928009  AllergyIntoleranceCode = "293928009"
	AllergyIntoleranceCode293929001  AllergyIntoleranceCode = "293929001"
	AllergyIntoleranceCode293930006  AllergyIntoleranceCode = "293930006"
	AllergyIntoleranceCode293933008  AllergyIntoleranceCode = "293933008"
	AllergyIntoleranceCode293934002  AllergyIntoleranceCode = "293934002"
	AllergyIntoleranceCode293935001  AllergyIntoleranceCode = "293935001"
	AllergyIntoleranceCode293936000  AllergyIntoleranceCode = "293936000"
	AllergyIntoleranceCode293937009  AllergyIntoleranceCode = "293937009"
	AllergyIntoleranceCode293938004  AllergyIntoleranceCode = "293938004"
	AllergyIntoleranceCode293939007  AllergyIntoleranceCode = "293939007"
	AllergyIntoleranceCode293940009  AllergyIntoleranceCode = "293940009"
	AllergyIntoleranceCode293941008  AllergyIntoleranceCode = "293941008"
	AllergyIntoleranceCode293942001  AllergyIntoleranceCode = "293942001"
	AllergyIntoleranceCode293943006  AllergyIntoleranceCode = "293943006"
	AllergyIntoleranceCode293946003  AllergyIntoleranceCode = "293946003"
	AllergyIntoleranceCode293948002  AllergyIntoleranceCode = "293948002"
	AllergyIntoleranceCode293949005  AllergyIntoleranceCode = "293949005"
	AllergyIntoleranceCode293950005  AllergyIntoleranceCode = "293950005"
	AllergyIntoleranceCode293952002  AllergyIntoleranceCode = "293952002"
	AllergyIntoleranceCode293954001  AllergyIntoleranceCode = "293954001"
	AllergyIntoleranceCode293955000  AllergyIntoleranceCode = "293955000"
	AllergyIntoleranceCode293956004  AllergyIntoleranceCode = "293956004"
	AllergyIntoleranceCode293957008  AllergyIntoleranceCode = "293957008"
	AllergyIntoleranceCode293958003  AllergyIntoleranceCode = "293958003"
	AllergyIntoleranceCode293960001  AllergyIntoleranceCode = "293960001"
	AllergyIntoleranceCode293962009  AllergyIntoleranceCode = "293962009"
	AllergyIntoleranceCode293963004  AllergyIntoleranceCode = "293963004"
	AllergyIntoleranceCode293964005  AllergyIntoleranceCode = "293964005"
	AllergyIntoleranceCode293965006  AllergyIntoleranceCode = "293965006"
	AllergyIntoleranceCode293966007  AllergyIntoleranceCode = "293966007"
	AllergyIntoleranceCode293967003  AllergyIntoleranceCode = "293967003"
	AllergyIntoleranceCode293968008  AllergyIntoleranceCode = "293968008"
	AllergyIntoleranceCode293969000  AllergyIntoleranceCode = "293969000"
	AllergyIntoleranceCode293970004  AllergyIntoleranceCode = "293970004"
	AllergyIntoleranceCode293972007  AllergyIntoleranceCode = "293972007"
	AllergyIntoleranceCode293973002  AllergyIntoleranceCode = "293973002"
	AllergyIntoleranceCode293974008  AllergyIntoleranceCode = "293974008"
	AllergyIntoleranceCode293975009  AllergyIntoleranceCode = "293975009"
	AllergyIntoleranceCode293976005  AllergyIntoleranceCode = "293976005"
	AllergyIntoleranceCode293977001  AllergyIntoleranceCode = "293977001"
	AllergyIntoleranceCode293978006  AllergyIntoleranceCode = "293978006"
	AllergyIntoleranceCode293979003  AllergyIntoleranceCode = "293979003"
	AllergyIntoleranceCode293980000  AllergyIntoleranceCode = "293980000"
	AllergyIntoleranceCode293981001  AllergyIntoleranceCode = "293981001"
	AllergyIntoleranceCode293982008  AllergyIntoleranceCode = "293982008"
	AllergyIntoleranceCode293983003  AllergyIntoleranceCode = "293983003"
	AllergyIntoleranceCode293984009  AllergyIntoleranceCode = "293984009"
	AllergyIntoleranceCode293985005  AllergyIntoleranceCode = "293985005"
	AllergyIntoleranceCode293986006  AllergyIntoleranceCode = "293986006"
	AllergyIntoleranceCode293987002  AllergyIntoleranceCode = "293987002"
	AllergyIntoleranceCode293988007  AllergyIntoleranceCode = "293988007"
	AllergyIntoleranceCode293989004  AllergyIntoleranceCode = "293989004"
	AllergyIntoleranceCode293990008  AllergyIntoleranceCode = "293990008"
	AllergyIntoleranceCode293991007  AllergyIntoleranceCode = "293991007"
	AllergyIntoleranceCode293992000  AllergyIntoleranceCode = "293992000"
	AllergyIntoleranceCode293993005  AllergyIntoleranceCode = "293993005"
	AllergyIntoleranceCode293994004  AllergyIntoleranceCode = "293994004"
	AllergyIntoleranceCode293995003  AllergyIntoleranceCode = "293995003"
	AllergyIntoleranceCode293996002  AllergyIntoleranceCode = "293996002"
	AllergyIntoleranceCode293997006  AllergyIntoleranceCode = "293997006"
	AllergyIntoleranceCode293998001  AllergyIntoleranceCode = "293998001"
	AllergyIntoleranceCode293999009  AllergyIntoleranceCode = "293999009"
	AllergyIntoleranceCode294000006  AllergyIntoleranceCode = "294000006"
	AllergyIntoleranceCode294001005  AllergyIntoleranceCode = "294001005"
	AllergyIntoleranceCode294002003  AllergyIntoleranceCode = "294002003"
	AllergyIntoleranceCode294003008  AllergyIntoleranceCode = "294003008"
	AllergyIntoleranceCode294004002  AllergyIntoleranceCode = "294004002"
	AllergyIntoleranceCode294005001  AllergyIntoleranceCode = "294005001"
	AllergyIntoleranceCode294007009  AllergyIntoleranceCode = "294007009"
	AllergyIntoleranceCode294009007  AllergyIntoleranceCode = "294009007"
	AllergyIntoleranceCode294011003  AllergyIntoleranceCode = "294011003"
	AllergyIntoleranceCode294012005  AllergyIntoleranceCode = "294012005"
	AllergyIntoleranceCode294013000  AllergyIntoleranceCode = "294013000"
	AllergyIntoleranceCode294014006  AllergyIntoleranceCode = "294014006"
	AllergyIntoleranceCode294015007  AllergyIntoleranceCode = "294015007"
	AllergyIntoleranceCode294016008  AllergyIntoleranceCode = "294016008"
	AllergyIntoleranceCode294017004  AllergyIntoleranceCode = "294017004"
	AllergyIntoleranceCode294018009  AllergyIntoleranceCode = "294018009"
	AllergyIntoleranceCode294019001  AllergyIntoleranceCode = "294019001"
	AllergyIntoleranceCode294023009  AllergyIntoleranceCode = "294023009"
	AllergyIntoleranceCode294024003  AllergyIntoleranceCode = "294024003"
	AllergyIntoleranceCode294026001  AllergyIntoleranceCode = "294026001"
	AllergyIntoleranceCode294027005  AllergyIntoleranceCode = "294027005"
	AllergyIntoleranceCode294028000  AllergyIntoleranceCode = "294028000"
	AllergyIntoleranceCode294029008  AllergyIntoleranceCode = "294029008"
	AllergyIntoleranceCode294030003  AllergyIntoleranceCode = "294030003"
	AllergyIntoleranceCode294031004  AllergyIntoleranceCode = "294031004"
	AllergyIntoleranceCode294033001  AllergyIntoleranceCode = "294033001"
	AllergyIntoleranceCode294035008  AllergyIntoleranceCode = "294035008"
	AllergyIntoleranceCode294036009  AllergyIntoleranceCode = "294036009"
	AllergyIntoleranceCode294037000  AllergyIntoleranceCode = "294037000"
	AllergyIntoleranceCode294038005  AllergyIntoleranceCode = "294038005"
	AllergyIntoleranceCode294039002  AllergyIntoleranceCode = "294039002"
	AllergyIntoleranceCode294040000  AllergyIntoleranceCode = "294040000"
	AllergyIntoleranceCode294041001  AllergyIntoleranceCode = "294041001"
	AllergyIntoleranceCode294042008  AllergyIntoleranceCode = "294042008"
	AllergyIntoleranceCode294043003  AllergyIntoleranceCode = "294043003"
	AllergyIntoleranceCode294044009  AllergyIntoleranceCode = "294044009"
	AllergyIntoleranceCode294045005  AllergyIntoleranceCode = "294045005"
	AllergyIntoleranceCode294047002  AllergyIntoleranceCode = "294047002"
	AllergyIntoleranceCode294048007  AllergyIntoleranceCode = "294048007"
	AllergyIntoleranceCode294050004  AllergyIntoleranceCode = "294050004"
	AllergyIntoleranceCode294055009  AllergyIntoleranceCode = "294055009"
	AllergyIntoleranceCode294057001  AllergyIntoleranceCode = "294057001"
	AllergyIntoleranceCode294058006  AllergyIntoleranceCode = "294058006"
	AllergyIntoleranceCode294059003  AllergyIntoleranceCode = "294059003"
	AllergyIntoleranceCode294060008  AllergyIntoleranceCode = "294060008"
	AllergyIntoleranceCode294061007  AllergyIntoleranceCode = "294061007"
	AllergyIntoleranceCode294062000  AllergyIntoleranceCode = "294062000"
	AllergyIntoleranceCode294063005  AllergyIntoleranceCode = "294063005"
	AllergyIntoleranceCode294064004  AllergyIntoleranceCode = "294064004"
	AllergyIntoleranceCode294067006  AllergyIntoleranceCode = "294067006"
	AllergyIntoleranceCode294069009  AllergyIntoleranceCode = "294069009"
	AllergyIntoleranceCode294073007  AllergyIntoleranceCode = "294073007"
	AllergyIntoleranceCode294074001  AllergyIntoleranceCode = "294074001"
	AllergyIntoleranceCode294076004  AllergyIntoleranceCode = "294076004"
	AllergyIntoleranceCode294077008  AllergyIntoleranceCode = "294077008"
	AllergyIntoleranceCode294078003  AllergyIntoleranceCode = "294078003"
	AllergyIntoleranceCode294079006  AllergyIntoleranceCode = "294079006"
	AllergyIntoleranceCode294080009  AllergyIntoleranceCode = "294080009"
	AllergyIntoleranceCode294081008  AllergyIntoleranceCode = "294081008"
	AllergyIntoleranceCode294082001  AllergyIntoleranceCode = "294082001"
	AllergyIntoleranceCode294083006  AllergyIntoleranceCode = "294083006"
	AllergyIntoleranceCode294084000  AllergyIntoleranceCode = "294084000"
	AllergyIntoleranceCode294087007  AllergyIntoleranceCode = "294087007"
	AllergyIntoleranceCode294088002  AllergyIntoleranceCode = "294088002"
	AllergyIntoleranceCode294089005  AllergyIntoleranceCode = "294089005"
	AllergyIntoleranceCode294091002  AllergyIntoleranceCode = "294091002"
	AllergyIntoleranceCode294093004  AllergyIntoleranceCode = "294093004"
	AllergyIntoleranceCode294095006  AllergyIntoleranceCode = "294095006"
	AllergyIntoleranceCode294096007  AllergyIntoleranceCode = "294096007"
	AllergyIntoleranceCode294097003  AllergyIntoleranceCode = "294097003"
	AllergyIntoleranceCode294099000  AllergyIntoleranceCode = "294099000"
	AllergyIntoleranceCode294100008  AllergyIntoleranceCode = "294100008"
	AllergyIntoleranceCode294101007  AllergyIntoleranceCode = "294101007"
	AllergyIntoleranceCode294103005  AllergyIntoleranceCode = "294103005"
	AllergyIntoleranceCode294105003  AllergyIntoleranceCode = "294105003"
	AllergyIntoleranceCode294106002  AllergyIntoleranceCode = "294106002"
	AllergyIntoleranceCode294109009  AllergyIntoleranceCode = "294109009"
	AllergyIntoleranceCode294111000  AllergyIntoleranceCode = "294111000"
	AllergyIntoleranceCode294112007  AllergyIntoleranceCode = "294112007"
	AllergyIntoleranceCode294113002  AllergyIntoleranceCode = "294113002"
	AllergyIntoleranceCode294114008  AllergyIntoleranceCode = "294114008"
	AllergyIntoleranceCode294115009  AllergyIntoleranceCode = "294115009"
	AllergyIntoleranceCode294116005  AllergyIntoleranceCode = "294116005"
	AllergyIntoleranceCode294118006  AllergyIntoleranceCode = "294118006"
	AllergyIntoleranceCode294119003  AllergyIntoleranceCode = "294119003"
	AllergyIntoleranceCode294120009  AllergyIntoleranceCode = "294120009"
	AllergyIntoleranceCode294121008  AllergyIntoleranceCode = "294121008"
	AllergyIntoleranceCode294122001  AllergyIntoleranceCode = "294122001"
	AllergyIntoleranceCode294123006  AllergyIntoleranceCode = "294123006"
	AllergyIntoleranceCode294125004  AllergyIntoleranceCode = "294125004"
	AllergyIntoleranceCode294126003  AllergyIntoleranceCode = "294126003"
	AllergyIntoleranceCode294127007  AllergyIntoleranceCode = "294127007"
	AllergyIntoleranceCode294128002  AllergyIntoleranceCode = "294128002"
	AllergyIntoleranceCode294129005  AllergyIntoleranceCode = "294129005"
	AllergyIntoleranceCode294130000  AllergyIntoleranceCode = "294130000"
	AllergyIntoleranceCode294131001  AllergyIntoleranceCode = "294131001"
	AllergyIntoleranceCode294132008  AllergyIntoleranceCode = "294132008"
	AllergyIntoleranceCode294133003  AllergyIntoleranceCode = "294133003"
	AllergyIntoleranceCode294134009  AllergyIntoleranceCode = "294134009"
	AllergyIntoleranceCode294135005  AllergyIntoleranceCode = "294135005"
	AllergyIntoleranceCode294136006  AllergyIntoleranceCode = "294136006"
	AllergyIntoleranceCode294137002  AllergyIntoleranceCode = "294137002"
	AllergyIntoleranceCode294138007  AllergyIntoleranceCode = "294138007"
	AllergyIntoleranceCode294139004  AllergyIntoleranceCode = "294139004"
	AllergyIntoleranceCode294140002  AllergyIntoleranceCode = "294140002"
	AllergyIntoleranceCode294142005  AllergyIntoleranceCode = "294142005"
	AllergyIntoleranceCode294144006  AllergyIntoleranceCode = "294144006"
	AllergyIntoleranceCode294145007  AllergyIntoleranceCode = "294145007"
	AllergyIntoleranceCode294148009  AllergyIntoleranceCode = "294148009"
	AllergyIntoleranceCode294152009  AllergyIntoleranceCode = "294152009"
	AllergyIntoleranceCode294153004  AllergyIntoleranceCode = "294153004"
	AllergyIntoleranceCode294157003  AllergyIntoleranceCode = "294157003"
	AllergyIntoleranceCode294158008  AllergyIntoleranceCode = "294158008"
	AllergyIntoleranceCode294160005  AllergyIntoleranceCode = "294160005"
	AllergyIntoleranceCode294168003  AllergyIntoleranceCode = "294168003"
	AllergyIntoleranceCode294169006  AllergyIntoleranceCode = "294169006"
	AllergyIntoleranceCode294172004  AllergyIntoleranceCode = "294172004"
	AllergyIntoleranceCode294173009  AllergyIntoleranceCode = "294173009"
	AllergyIntoleranceCode294177005  AllergyIntoleranceCode = "294177005"
	AllergyIntoleranceCode294178000  AllergyIntoleranceCode = "294178000"
	AllergyIntoleranceCode294180006  AllergyIntoleranceCode = "294180006"
	AllergyIntoleranceCode294182003  AllergyIntoleranceCode = "294182003"
	AllergyIntoleranceCode294183008  AllergyIntoleranceCode = "294183008"
	AllergyIntoleranceCode294189007  AllergyIntoleranceCode = "294189007"
	AllergyIntoleranceCode294190003  AllergyIntoleranceCode = "294190003"
	AllergyIntoleranceCode294191004  AllergyIntoleranceCode = "294191004"
	AllergyIntoleranceCode294200004  AllergyIntoleranceCode = "294200004"
	AllergyIntoleranceCode294203002  AllergyIntoleranceCode = "294203002"
	AllergyIntoleranceCode294204008  AllergyIntoleranceCode = "294204008"
	AllergyIntoleranceCode294206005  AllergyIntoleranceCode = "294206005"
	AllergyIntoleranceCode294207001  AllergyIntoleranceCode = "294207001"
	AllergyIntoleranceCode294208006  AllergyIntoleranceCode = "294208006"
	AllergyIntoleranceCode294209003  AllergyIntoleranceCode = "294209003"
	AllergyIntoleranceCode294210008  AllergyIntoleranceCode = "294210008"
	AllergyIntoleranceCode294211007  AllergyIntoleranceCode = "294211007"
	AllergyIntoleranceCode294214004  AllergyIntoleranceCode = "294214004"
	AllergyIntoleranceCode294215003  AllergyIntoleranceCode = "294215003"
	AllergyIntoleranceCode294217006  AllergyIntoleranceCode = "294217006"
	AllergyIntoleranceCode294218001  AllergyIntoleranceCode = "294218001"
	AllergyIntoleranceCode294219009  AllergyIntoleranceCode = "294219009"
	AllergyIntoleranceCode294220003  AllergyIntoleranceCode = "294220003"
	AllergyIntoleranceCode294224007  AllergyIntoleranceCode = "294224007"
	AllergyIntoleranceCode294225008  AllergyIntoleranceCode = "294225008"
	AllergyIntoleranceCode294226009  AllergyIntoleranceCode = "294226009"
	AllergyIntoleranceCode294227000  AllergyIntoleranceCode = "294227000"
	AllergyIntoleranceCode294228005  AllergyIntoleranceCode = "294228005"
	AllergyIntoleranceCode294229002  AllergyIntoleranceCode = "294229002"
	AllergyIntoleranceCode294230007  AllergyIntoleranceCode = "294230007"
	AllergyIntoleranceCode294231006  AllergyIntoleranceCode = "294231006"
	AllergyIntoleranceCode294232004  AllergyIntoleranceCode = "294232004"
	AllergyIntoleranceCode294233009  AllergyIntoleranceCode = "294233009"
	AllergyIntoleranceCode294234003  AllergyIntoleranceCode = "294234003"
	AllergyIntoleranceCode294235002  AllergyIntoleranceCode = "294235002"
	AllergyIntoleranceCode294236001  AllergyIntoleranceCode = "294236001"
	AllergyIntoleranceCode294237005  AllergyIntoleranceCode = "294237005"
	AllergyIntoleranceCode294238000  AllergyIntoleranceCode = "294238000"
	AllergyIntoleranceCode294239008  AllergyIntoleranceCode = "294239008"
	AllergyIntoleranceCode294240005  AllergyIntoleranceCode = "294240005"
	AllergyIntoleranceCode294242002  AllergyIntoleranceCode = "294242002"
	AllergyIntoleranceCode294243007  AllergyIntoleranceCode = "294243007"
	AllergyIntoleranceCode294245000  AllergyIntoleranceCode = "294245000"
	AllergyIntoleranceCode294246004  AllergyIntoleranceCode = "294246004"
	AllergyIntoleranceCode294247008  AllergyIntoleranceCode = "294247008"
	AllergyIntoleranceCode294248003  AllergyIntoleranceCode = "294248003"
	AllergyIntoleranceCode294249006  AllergyIntoleranceCode = "294249006"
	AllergyIntoleranceCode294250006  AllergyIntoleranceCode = "294250006"
	AllergyIntoleranceCode294252003  AllergyIntoleranceCode = "294252003"
	AllergyIntoleranceCode294253008  AllergyIntoleranceCode = "294253008"
	AllergyIntoleranceCode294254002  AllergyIntoleranceCode = "294254002"
	AllergyIntoleranceCode294255001  AllergyIntoleranceCode = "294255001"
	AllergyIntoleranceCode294256000  AllergyIntoleranceCode = "294256000"
	AllergyIntoleranceCode294257009  AllergyIntoleranceCode = "294257009"
	AllergyIntoleranceCode294258004  AllergyIntoleranceCode = "294258004"
	AllergyIntoleranceCode294259007  AllergyIntoleranceCode = "294259007"
	AllergyIntoleranceCode294260002  AllergyIntoleranceCode = "294260002"
	AllergyIntoleranceCode294261003  AllergyIntoleranceCode = "294261003"
	AllergyIntoleranceCode294264006  AllergyIntoleranceCode = "294264006"
	AllergyIntoleranceCode294265007  AllergyIntoleranceCode = "294265007"
	AllergyIntoleranceCode294266008  AllergyIntoleranceCode = "294266008"
	AllergyIntoleranceCode294268009  AllergyIntoleranceCode = "294268009"
	AllergyIntoleranceCode294269001  AllergyIntoleranceCode = "294269001"
	AllergyIntoleranceCode294270000  AllergyIntoleranceCode = "294270000"
	AllergyIntoleranceCode294271001  AllergyIntoleranceCode = "294271001"
	AllergyIntoleranceCode294273003  AllergyIntoleranceCode = "294273003"
	AllergyIntoleranceCode294275005  AllergyIntoleranceCode = "294275005"
	AllergyIntoleranceCode294276006  AllergyIntoleranceCode = "294276006"
	AllergyIntoleranceCode294277002  AllergyIntoleranceCode = "294277002"
	AllergyIntoleranceCode294278007  AllergyIntoleranceCode = "294278007"
	AllergyIntoleranceCode294280001  AllergyIntoleranceCode = "294280001"
	AllergyIntoleranceCode294281002  AllergyIntoleranceCode = "294281002"
	AllergyIntoleranceCode294283004  AllergyIntoleranceCode = "294283004"
	AllergyIntoleranceCode294284005  AllergyIntoleranceCode = "294284005"
	AllergyIntoleranceCode294285006  AllergyIntoleranceCode = "294285006"
	AllergyIntoleranceCode294290009  AllergyIntoleranceCode = "294290009"
	AllergyIntoleranceCode294291008  AllergyIntoleranceCode = "294291008"
	AllergyIntoleranceCode294298002  AllergyIntoleranceCode = "294298002"
	AllergyIntoleranceCode294299005  AllergyIntoleranceCode = "294299005"
	AllergyIntoleranceCode294306008  AllergyIntoleranceCode = "294306008"
	AllergyIntoleranceCode294316000  AllergyIntoleranceCode = "294316000"
	AllergyIntoleranceCode294317009  AllergyIntoleranceCode = "294317009"
	AllergyIntoleranceCode294318004  AllergyIntoleranceCode = "294318004"
	AllergyIntoleranceCode294320001  AllergyIntoleranceCode = "294320001"
	AllergyIntoleranceCode294324005  AllergyIntoleranceCode = "294324005"
	AllergyIntoleranceCode294327003  AllergyIntoleranceCode = "294327003"
	AllergyIntoleranceCode294328008  AllergyIntoleranceCode = "294328008"
	AllergyIntoleranceCode294329000  AllergyIntoleranceCode = "294329000"
	AllergyIntoleranceCode294330005  AllergyIntoleranceCode = "294330005"
	AllergyIntoleranceCode294332002  AllergyIntoleranceCode = "294332002"
	AllergyIntoleranceCode294333007  AllergyIntoleranceCode = "294333007"
	AllergyIntoleranceCode294335000  AllergyIntoleranceCode = "294335000"
	AllergyIntoleranceCode294337008  AllergyIntoleranceCode = "294337008"
	AllergyIntoleranceCode294339006  AllergyIntoleranceCode = "294339006"
	AllergyIntoleranceCode294341007  AllergyIntoleranceCode = "294341007"
	AllergyIntoleranceCode294342000  AllergyIntoleranceCode = "294342000"
	AllergyIntoleranceCode294343005  AllergyIntoleranceCode = "294343005"
	AllergyIntoleranceCode294344004  AllergyIntoleranceCode = "294344004"
	AllergyIntoleranceCode294346002  AllergyIntoleranceCode = "294346002"
	AllergyIntoleranceCode294348001  AllergyIntoleranceCode = "294348001"
	AllergyIntoleranceCode294349009  AllergyIntoleranceCode = "294349009"
	AllergyIntoleranceCode294350009  AllergyIntoleranceCode = "294350009"
	AllergyIntoleranceCode294351008  AllergyIntoleranceCode = "294351008"
	AllergyIntoleranceCode294354000  AllergyIntoleranceCode = "294354000"
	AllergyIntoleranceCode294356003  AllergyIntoleranceCode = "294356003"
	AllergyIntoleranceCode294357007  AllergyIntoleranceCode = "294357007"
	AllergyIntoleranceCode294358002  AllergyIntoleranceCode = "294358002"
	AllergyIntoleranceCode294359005  AllergyIntoleranceCode = "294359005"
	AllergyIntoleranceCode294360000  AllergyIntoleranceCode = "294360000"
	AllergyIntoleranceCode294361001  AllergyIntoleranceCode = "294361001"
	AllergyIntoleranceCode294362008  AllergyIntoleranceCode = "294362008"
	AllergyIntoleranceCode294363003  AllergyIntoleranceCode = "294363003"
	AllergyIntoleranceCode294365005  AllergyIntoleranceCode = "294365005"
	AllergyIntoleranceCode294366006  AllergyIntoleranceCode = "294366006"
	AllergyIntoleranceCode294368007  AllergyIntoleranceCode = "294368007"
	AllergyIntoleranceCode294369004  AllergyIntoleranceCode = "294369004"
	AllergyIntoleranceCode294370003  AllergyIntoleranceCode = "294370003"
	AllergyIntoleranceCode294371004  AllergyIntoleranceCode = "294371004"
	AllergyIntoleranceCode294372006  AllergyIntoleranceCode = "294372006"
	AllergyIntoleranceCode294373001  AllergyIntoleranceCode = "294373001"
	AllergyIntoleranceCode294374007  AllergyIntoleranceCode = "294374007"
	AllergyIntoleranceCode294375008  AllergyIntoleranceCode = "294375008"
	AllergyIntoleranceCode294376009  AllergyIntoleranceCode = "294376009"
	AllergyIntoleranceCode294377000  AllergyIntoleranceCode = "294377000"
	AllergyIntoleranceCode294378005  AllergyIntoleranceCode = "294378005"
	AllergyIntoleranceCode294379002  AllergyIntoleranceCode = "294379002"
	AllergyIntoleranceCode294380004  AllergyIntoleranceCode = "294380004"
	AllergyIntoleranceCode294381000  AllergyIntoleranceCode = "294381000"
	AllergyIntoleranceCode294382007  AllergyIntoleranceCode = "294382007"
	AllergyIntoleranceCode294383002  AllergyIntoleranceCode = "294383002"
	AllergyIntoleranceCode294384008  AllergyIntoleranceCode = "294384008"
	AllergyIntoleranceCode294385009  AllergyIntoleranceCode = "294385009"
	AllergyIntoleranceCode294388006  AllergyIntoleranceCode = "294388006"
	AllergyIntoleranceCode294390007  AllergyIntoleranceCode = "294390007"
	AllergyIntoleranceCode294391006  AllergyIntoleranceCode = "294391006"
	AllergyIntoleranceCode294392004  AllergyIntoleranceCode = "294392004"
	AllergyIntoleranceCode294393009  AllergyIntoleranceCode = "294393009"
	AllergyIntoleranceCode294394003  AllergyIntoleranceCode = "294394003"
	AllergyIntoleranceCode294396001  AllergyIntoleranceCode = "294396001"
	AllergyIntoleranceCode294398000  AllergyIntoleranceCode = "294398000"
	AllergyIntoleranceCode294399008  AllergyIntoleranceCode = "294399008"
	AllergyIntoleranceCode294400001  AllergyIntoleranceCode = "294400001"
	AllergyIntoleranceCode294404005  AllergyIntoleranceCode = "294404005"
	AllergyIntoleranceCode294405006  AllergyIntoleranceCode = "294405006"
	AllergyIntoleranceCode294406007  AllergyIntoleranceCode = "294406007"
	AllergyIntoleranceCode294407003  AllergyIntoleranceCode = "294407003"
	AllergyIntoleranceCode294408008  AllergyIntoleranceCode = "294408008"
	AllergyIntoleranceCode294413007  AllergyIntoleranceCode = "294413007"
	AllergyIntoleranceCode294415000  AllergyIntoleranceCode = "294415000"
	AllergyIntoleranceCode294416004  AllergyIntoleranceCode = "294416004"
	AllergyIntoleranceCode294417008  AllergyIntoleranceCode = "294417008"
	AllergyIntoleranceCode294418003  AllergyIntoleranceCode = "294418003"
	AllergyIntoleranceCode294421001  AllergyIntoleranceCode = "294421001"
	AllergyIntoleranceCode294423003  AllergyIntoleranceCode = "294423003"
	AllergyIntoleranceCode294425005  AllergyIntoleranceCode = "294425005"
	AllergyIntoleranceCode294426006  AllergyIntoleranceCode = "294426006"
	AllergyIntoleranceCode294431008  AllergyIntoleranceCode = "294431008"
	AllergyIntoleranceCode294433006  AllergyIntoleranceCode = "294433006"
	AllergyIntoleranceCode294434000  AllergyIntoleranceCode = "294434000"
	AllergyIntoleranceCode294436003  AllergyIntoleranceCode = "294436003"
	AllergyIntoleranceCode294437007  AllergyIntoleranceCode = "294437007"
	AllergyIntoleranceCode294438002  AllergyIntoleranceCode = "294438002"
	AllergyIntoleranceCode294439005  AllergyIntoleranceCode = "294439005"
	AllergyIntoleranceCode294440007  AllergyIntoleranceCode = "294440007"
	AllergyIntoleranceCode294441006  AllergyIntoleranceCode = "294441006"
	AllergyIntoleranceCode294442004  AllergyIntoleranceCode = "294442004"
	AllergyIntoleranceCode294443009  AllergyIntoleranceCode = "294443009"
	AllergyIntoleranceCode294447005  AllergyIntoleranceCode = "294447005"
	AllergyIntoleranceCode294448000  AllergyIntoleranceCode = "294448000"
	AllergyIntoleranceCode294449008  AllergyIntoleranceCode = "294449008"
	AllergyIntoleranceCode294451007  AllergyIntoleranceCode = "294451007"
	AllergyIntoleranceCode294452000  AllergyIntoleranceCode = "294452000"
	AllergyIntoleranceCode294453005  AllergyIntoleranceCode = "294453005"
	AllergyIntoleranceCode294455003  AllergyIntoleranceCode = "294455003"
	AllergyIntoleranceCode294456002  AllergyIntoleranceCode = "294456002"
	AllergyIntoleranceCode294458001  AllergyIntoleranceCode = "294458001"
	AllergyIntoleranceCode294459009  AllergyIntoleranceCode = "294459009"
	AllergyIntoleranceCode294460004  AllergyIntoleranceCode = "294460004"
	AllergyIntoleranceCode294462007  AllergyIntoleranceCode = "294462007"
	AllergyIntoleranceCode294463002  AllergyIntoleranceCode = "294463002"
	AllergyIntoleranceCode294464008  AllergyIntoleranceCode = "294464008"
	AllergyIntoleranceCode294465009  AllergyIntoleranceCode = "294465009"
	AllergyIntoleranceCode294466005  AllergyIntoleranceCode = "294466005"
	AllergyIntoleranceCode294467001  AllergyIntoleranceCode = "294467001"
	AllergyIntoleranceCode294468006  AllergyIntoleranceCode = "294468006"
	AllergyIntoleranceCode294469003  AllergyIntoleranceCode = "294469003"
	AllergyIntoleranceCode294470002  AllergyIntoleranceCode = "294470002"
	AllergyIntoleranceCode294471003  AllergyIntoleranceCode = "294471003"
	AllergyIntoleranceCode294472005  AllergyIntoleranceCode = "294472005"
	AllergyIntoleranceCode294474006  AllergyIntoleranceCode = "294474006"
	AllergyIntoleranceCode294475007  AllergyIntoleranceCode = "294475007"
	AllergyIntoleranceCode294476008  AllergyIntoleranceCode = "294476008"
	AllergyIntoleranceCode294477004  AllergyIntoleranceCode = "294477004"
	AllergyIntoleranceCode294478009  AllergyIntoleranceCode = "294478009"
	AllergyIntoleranceCode294480003  AllergyIntoleranceCode = "294480003"
	AllergyIntoleranceCode294481004  AllergyIntoleranceCode = "294481004"
	AllergyIntoleranceCode294482006  AllergyIntoleranceCode = "294482006"
	AllergyIntoleranceCode294484007  AllergyIntoleranceCode = "294484007"
	AllergyIntoleranceCode294485008  AllergyIntoleranceCode = "294485008"
	AllergyIntoleranceCode294486009  AllergyIntoleranceCode = "294486009"
	AllergyIntoleranceCode294487000  AllergyIntoleranceCode = "294487000"
	AllergyIntoleranceCode294488005  AllergyIntoleranceCode = "294488005"
	AllergyIntoleranceCode294489002  AllergyIntoleranceCode = "294489002"
	AllergyIntoleranceCode294490006  AllergyIntoleranceCode = "294490006"
	AllergyIntoleranceCode294491005  AllergyIntoleranceCode = "294491005"
	AllergyIntoleranceCode294494002  AllergyIntoleranceCode = "294494002"
	AllergyIntoleranceCode294496000  AllergyIntoleranceCode = "294496000"
	AllergyIntoleranceCode294497009  AllergyIntoleranceCode = "294497009"
	AllergyIntoleranceCode294499007  AllergyIntoleranceCode = "294499007"
	AllergyIntoleranceCode294501004  AllergyIntoleranceCode = "294501004"
	AllergyIntoleranceCode294502006  AllergyIntoleranceCode = "294502006"
	AllergyIntoleranceCode294503001  AllergyIntoleranceCode = "294503001"
	AllergyIntoleranceCode294505008  AllergyIntoleranceCode = "294505008"
	AllergyIntoleranceCode294506009  AllergyIntoleranceCode = "294506009"
	AllergyIntoleranceCode294507000  AllergyIntoleranceCode = "294507000"
	AllergyIntoleranceCode294508005  AllergyIntoleranceCode = "294508005"
	AllergyIntoleranceCode294509002  AllergyIntoleranceCode = "294509002"
	AllergyIntoleranceCode294510007  AllergyIntoleranceCode = "294510007"
	AllergyIntoleranceCode294511006  AllergyIntoleranceCode = "294511006"
	AllergyIntoleranceCode294512004  AllergyIntoleranceCode = "294512004"
	AllergyIntoleranceCode294514003  AllergyIntoleranceCode = "294514003"
	AllergyIntoleranceCode294515002  AllergyIntoleranceCode = "294515002"
	AllergyIntoleranceCode294516001  AllergyIntoleranceCode = "294516001"
	AllergyIntoleranceCode294517005  AllergyIntoleranceCode = "294517005"
	AllergyIntoleranceCode294518000  AllergyIntoleranceCode = "294518000"
	AllergyIntoleranceCode294519008  AllergyIntoleranceCode = "294519008"
	AllergyIntoleranceCode294520002  AllergyIntoleranceCode = "294520002"
	AllergyIntoleranceCode294528009  AllergyIntoleranceCode = "294528009"
	AllergyIntoleranceCode294529001  AllergyIntoleranceCode = "294529001"
	AllergyIntoleranceCode294530006  AllergyIntoleranceCode = "294530006"
	AllergyIntoleranceCode294531005  AllergyIntoleranceCode = "294531005"
	AllergyIntoleranceCode294532003  AllergyIntoleranceCode = "294532003"
	AllergyIntoleranceCode294534002  AllergyIntoleranceCode = "294534002"
	AllergyIntoleranceCode294535001  AllergyIntoleranceCode = "294535001"
	AllergyIntoleranceCode294536000  AllergyIntoleranceCode = "294536000"
	AllergyIntoleranceCode294537009  AllergyIntoleranceCode = "294537009"
	AllergyIntoleranceCode294538004  AllergyIntoleranceCode = "294538004"
	AllergyIntoleranceCode294539007  AllergyIntoleranceCode = "294539007"
	AllergyIntoleranceCode294541008  AllergyIntoleranceCode = "294541008"
	AllergyIntoleranceCode294542001  AllergyIntoleranceCode = "294542001"
	AllergyIntoleranceCode294543006  AllergyIntoleranceCode = "294543006"
	AllergyIntoleranceCode294545004  AllergyIntoleranceCode = "294545004"
	AllergyIntoleranceCode294546003  AllergyIntoleranceCode = "294546003"
	AllergyIntoleranceCode294547007  AllergyIntoleranceCode = "294547007"
	AllergyIntoleranceCode294548002  AllergyIntoleranceCode = "294548002"
	AllergyIntoleranceCode294549005  AllergyIntoleranceCode = "294549005"
	AllergyIntoleranceCode294550005  AllergyIntoleranceCode = "294550005"
	AllergyIntoleranceCode294551009  AllergyIntoleranceCode = "294551009"
	AllergyIntoleranceCode294552002  AllergyIntoleranceCode = "294552002"
	AllergyIntoleranceCode294554001  AllergyIntoleranceCode = "294554001"
	AllergyIntoleranceCode294556004  AllergyIntoleranceCode = "294556004"
	AllergyIntoleranceCode294557008  AllergyIntoleranceCode = "294557008"
	AllergyIntoleranceCode294558003  AllergyIntoleranceCode = "294558003"
	AllergyIntoleranceCode294559006  AllergyIntoleranceCode = "294559006"
	AllergyIntoleranceCode294561002  AllergyIntoleranceCode = "294561002"
	AllergyIntoleranceCode294562009  AllergyIntoleranceCode = "294562009"
	AllergyIntoleranceCode294563004  AllergyIntoleranceCode = "294563004"
	AllergyIntoleranceCode294564005  AllergyIntoleranceCode = "294564005"
	AllergyIntoleranceCode294565006  AllergyIntoleranceCode = "294565006"
	AllergyIntoleranceCode294566007  AllergyIntoleranceCode = "294566007"
	AllergyIntoleranceCode294567003  AllergyIntoleranceCode = "294567003"
	AllergyIntoleranceCode294568008  AllergyIntoleranceCode = "294568008"
	AllergyIntoleranceCode294569000  AllergyIntoleranceCode = "294569000"
	AllergyIntoleranceCode294570004  AllergyIntoleranceCode = "294570004"
	AllergyIntoleranceCode294571000  AllergyIntoleranceCode = "294571000"
	AllergyIntoleranceCode294572007  AllergyIntoleranceCode = "294572007"
	AllergyIntoleranceCode294573002  AllergyIntoleranceCode = "294573002"
	AllergyIntoleranceCode294574008  AllergyIntoleranceCode = "294574008"
	AllergyIntoleranceCode294575009  AllergyIntoleranceCode = "294575009"
	AllergyIntoleranceCode294576005  AllergyIntoleranceCode = "294576005"
	AllergyIntoleranceCode294577001  AllergyIntoleranceCode = "294577001"
	AllergyIntoleranceCode294578006  AllergyIntoleranceCode = "294578006"
	AllergyIntoleranceCode294579003  AllergyIntoleranceCode = "294579003"
	AllergyIntoleranceCode294582008  AllergyIntoleranceCode = "294582008"
	AllergyIntoleranceCode294584009  AllergyIntoleranceCode = "294584009"
	AllergyIntoleranceCode294585005  AllergyIntoleranceCode = "294585005"
	AllergyIntoleranceCode294586006  AllergyIntoleranceCode = "294586006"
	AllergyIntoleranceCode294587002  AllergyIntoleranceCode = "294587002"
	AllergyIntoleranceCode294588007  AllergyIntoleranceCode = "294588007"
	AllergyIntoleranceCode294590008  AllergyIntoleranceCode = "294590008"
	AllergyIntoleranceCode294591007  AllergyIntoleranceCode = "294591007"
	AllergyIntoleranceCode294592000  AllergyIntoleranceCode = "294592000"
	AllergyIntoleranceCode294593005  AllergyIntoleranceCode = "294593005"
	AllergyIntoleranceCode294596002  AllergyIntoleranceCode = "294596002"
	AllergyIntoleranceCode294598001  AllergyIntoleranceCode = "294598001"
	AllergyIntoleranceCode294600007  AllergyIntoleranceCode = "294600007"
	AllergyIntoleranceCode294602004  AllergyIntoleranceCode = "294602004"
	AllergyIntoleranceCode294604003  AllergyIntoleranceCode = "294604003"
	AllergyIntoleranceCode294607005  AllergyIntoleranceCode = "294607005"
	AllergyIntoleranceCode294609008  AllergyIntoleranceCode = "294609008"
	AllergyIntoleranceCode294610003  AllergyIntoleranceCode = "294610003"
	AllergyIntoleranceCode294611004  AllergyIntoleranceCode = "294611004"
	AllergyIntoleranceCode294612006  AllergyIntoleranceCode = "294612006"
	AllergyIntoleranceCode294614007  AllergyIntoleranceCode = "294614007"
	AllergyIntoleranceCode294615008  AllergyIntoleranceCode = "294615008"
	AllergyIntoleranceCode294617000  AllergyIntoleranceCode = "294617000"
	AllergyIntoleranceCode294618005  AllergyIntoleranceCode = "294618005"
	AllergyIntoleranceCode294620008  AllergyIntoleranceCode = "294620008"
	AllergyIntoleranceCode294621007  AllergyIntoleranceCode = "294621007"
	AllergyIntoleranceCode294623005  AllergyIntoleranceCode = "294623005"
	AllergyIntoleranceCode294625003  AllergyIntoleranceCode = "294625003"
	AllergyIntoleranceCode294627006  AllergyIntoleranceCode = "294627006"
	AllergyIntoleranceCode294629009  AllergyIntoleranceCode = "294629009"
	AllergyIntoleranceCode294630004  AllergyIntoleranceCode = "294630004"
	AllergyIntoleranceCode294633002  AllergyIntoleranceCode = "294633002"
	AllergyIntoleranceCode294638006  AllergyIntoleranceCode = "294638006"
	AllergyIntoleranceCode294639003  AllergyIntoleranceCode = "294639003"
	AllergyIntoleranceCode294667007  AllergyIntoleranceCode = "294667007"
	AllergyIntoleranceCode294668002  AllergyIntoleranceCode = "294668002"
	AllergyIntoleranceCode294669005  AllergyIntoleranceCode = "294669005"
	AllergyIntoleranceCode294671005  AllergyIntoleranceCode = "294671005"
	AllergyIntoleranceCode294674002  AllergyIntoleranceCode = "294674002"
	AllergyIntoleranceCode294676000  AllergyIntoleranceCode = "294676000"
	AllergyIntoleranceCode294677009  AllergyIntoleranceCode = "294677009"
	AllergyIntoleranceCode294678004  AllergyIntoleranceCode = "294678004"
	AllergyIntoleranceCode294679007  AllergyIntoleranceCode = "294679007"
	AllergyIntoleranceCode294682002  AllergyIntoleranceCode = "294682002"
	AllergyIntoleranceCode294683007  AllergyIntoleranceCode = "294683007"
	AllergyIntoleranceCode294684001  AllergyIntoleranceCode = "294684001"
	AllergyIntoleranceCode294685000  AllergyIntoleranceCode = "294685000"
	AllergyIntoleranceCode294686004  AllergyIntoleranceCode = "294686004"
	AllergyIntoleranceCode294687008  AllergyIntoleranceCode = "294687008"
	AllergyIntoleranceCode294688003  AllergyIntoleranceCode = "294688003"
	AllergyIntoleranceCode294689006  AllergyIntoleranceCode = "294689006"
	AllergyIntoleranceCode294690002  AllergyIntoleranceCode = "294690002"
	AllergyIntoleranceCode294691003  AllergyIntoleranceCode = "294691003"
	AllergyIntoleranceCode294692005  AllergyIntoleranceCode = "294692005"
	AllergyIntoleranceCode294693000  AllergyIntoleranceCode = "294693000"
	AllergyIntoleranceCode294694006  AllergyIntoleranceCode = "294694006"
	AllergyIntoleranceCode294695007  AllergyIntoleranceCode = "294695007"
	AllergyIntoleranceCode294696008  AllergyIntoleranceCode = "294696008"
	AllergyIntoleranceCode294697004  AllergyIntoleranceCode = "294697004"
	AllergyIntoleranceCode294698009  AllergyIntoleranceCode = "294698009"
	AllergyIntoleranceCode294699001  AllergyIntoleranceCode = "294699001"
	AllergyIntoleranceCode294700000  AllergyIntoleranceCode = "294700000"
	AllergyIntoleranceCode294701001  AllergyIntoleranceCode = "294701001"
	AllergyIntoleranceCode294702008  AllergyIntoleranceCode = "294702008"
	AllergyIntoleranceCode294706006  AllergyIntoleranceCode = "294706006"
	AllergyIntoleranceCode294707002  AllergyIntoleranceCode = "294707002"
	AllergyIntoleranceCode294711008  AllergyIntoleranceCode = "294711008"
	AllergyIntoleranceCode294712001  AllergyIntoleranceCode = "294712001"
	AllergyIntoleranceCode294714000  AllergyIntoleranceCode = "294714000"
	AllergyIntoleranceCode294717007  AllergyIntoleranceCode = "294717007"
	AllergyIntoleranceCode294720004  AllergyIntoleranceCode = "294720004"
	AllergyIntoleranceCode294721000  AllergyIntoleranceCode = "294721000"
	AllergyIntoleranceCode294723002  AllergyIntoleranceCode = "294723002"
	AllergyIntoleranceCode294728006  AllergyIntoleranceCode = "294728006"
	AllergyIntoleranceCode294729003  AllergyIntoleranceCode = "294729003"
	AllergyIntoleranceCode294730008  AllergyIntoleranceCode = "294730008"
	AllergyIntoleranceCode294731007  AllergyIntoleranceCode = "294731007"
	AllergyIntoleranceCode294732000  AllergyIntoleranceCode = "294732000"
	AllergyIntoleranceCode294733005  AllergyIntoleranceCode = "294733005"
	AllergyIntoleranceCode294734004  AllergyIntoleranceCode = "294734004"
	AllergyIntoleranceCode294735003  AllergyIntoleranceCode = "294735003"
	AllergyIntoleranceCode294736002  AllergyIntoleranceCode = "294736002"
	AllergyIntoleranceCode294737006  AllergyIntoleranceCode = "294737006"
	AllergyIntoleranceCode294738001  AllergyIntoleranceCode = "294738001"
	AllergyIntoleranceCode294739009  AllergyIntoleranceCode = "294739009"
	AllergyIntoleranceCode294740006  AllergyIntoleranceCode = "294740006"
	AllergyIntoleranceCode294741005  AllergyIntoleranceCode = "294741005"
	AllergyIntoleranceCode294742003  AllergyIntoleranceCode = "294742003"
	AllergyIntoleranceCode294745001  AllergyIntoleranceCode = "294745001"
	AllergyIntoleranceCode294746000  AllergyIntoleranceCode = "294746000"
	AllergyIntoleranceCode294747009  AllergyIntoleranceCode = "294747009"
	AllergyIntoleranceCode294748004  AllergyIntoleranceCode = "294748004"
	AllergyIntoleranceCode294749007  AllergyIntoleranceCode = "294749007"
	AllergyIntoleranceCode294750007  AllergyIntoleranceCode = "294750007"
	AllergyIntoleranceCode294751006  AllergyIntoleranceCode = "294751006"
	AllergyIntoleranceCode294752004  AllergyIntoleranceCode = "294752004"
	AllergyIntoleranceCode294754003  AllergyIntoleranceCode = "294754003"
	AllergyIntoleranceCode294755002  AllergyIntoleranceCode = "294755002"
	AllergyIntoleranceCode294757005  AllergyIntoleranceCode = "294757005"
	AllergyIntoleranceCode294758000  AllergyIntoleranceCode = "294758000"
	AllergyIntoleranceCode294760003  AllergyIntoleranceCode = "294760003"
	AllergyIntoleranceCode294761004  AllergyIntoleranceCode = "294761004"
	AllergyIntoleranceCode294762006  AllergyIntoleranceCode = "294762006"
	AllergyIntoleranceCode294763001  AllergyIntoleranceCode = "294763001"
	AllergyIntoleranceCode294764007  AllergyIntoleranceCode = "294764007"
	AllergyIntoleranceCode294765008  AllergyIntoleranceCode = "294765008"
	AllergyIntoleranceCode294767000  AllergyIntoleranceCode = "294767000"
	AllergyIntoleranceCode294768005  AllergyIntoleranceCode = "294768005"
	AllergyIntoleranceCode294769002  AllergyIntoleranceCode = "294769002"
	AllergyIntoleranceCode294771002  AllergyIntoleranceCode = "294771002"
	AllergyIntoleranceCode294773004  AllergyIntoleranceCode = "294773004"
	AllergyIntoleranceCode294774005  AllergyIntoleranceCode = "294774005"
	AllergyIntoleranceCode294775006  AllergyIntoleranceCode = "294775006"
	AllergyIntoleranceCode294776007  AllergyIntoleranceCode = "294776007"
	AllergyIntoleranceCode294781003  AllergyIntoleranceCode = "294781003"
	AllergyIntoleranceCode294782005  AllergyIntoleranceCode = "294782005"
	AllergyIntoleranceCode294787004  AllergyIntoleranceCode = "294787004"
	AllergyIntoleranceCode294788009  AllergyIntoleranceCode = "294788009"
	AllergyIntoleranceCode294789001  AllergyIntoleranceCode = "294789001"
	AllergyIntoleranceCode294792002  AllergyIntoleranceCode = "294792002"
	AllergyIntoleranceCode294793007  AllergyIntoleranceCode = "294793007"
	AllergyIntoleranceCode294794001  AllergyIntoleranceCode = "294794001"
	AllergyIntoleranceCode294795000  AllergyIntoleranceCode = "294795000"
	AllergyIntoleranceCode294796004  AllergyIntoleranceCode = "294796004"
	AllergyIntoleranceCode294798003  AllergyIntoleranceCode = "294798003"
	AllergyIntoleranceCode294799006  AllergyIntoleranceCode = "294799006"
	AllergyIntoleranceCode294800005  AllergyIntoleranceCode = "294800005"
	AllergyIntoleranceCode294801009  AllergyIntoleranceCode = "294801009"
	AllergyIntoleranceCode294803007  AllergyIntoleranceCode = "294803007"
	AllergyIntoleranceCode294804001  AllergyIntoleranceCode = "294804001"
	AllergyIntoleranceCode294807008  AllergyIntoleranceCode = "294807008"
	AllergyIntoleranceCode294809006  AllergyIntoleranceCode = "294809006"
	AllergyIntoleranceCode294810001  AllergyIntoleranceCode = "294810001"
	AllergyIntoleranceCode294811002  AllergyIntoleranceCode = "294811002"
	AllergyIntoleranceCode294813004  AllergyIntoleranceCode = "294813004"
	AllergyIntoleranceCode294814005  AllergyIntoleranceCode = "294814005"
	AllergyIntoleranceCode294815006  AllergyIntoleranceCode = "294815006"
	AllergyIntoleranceCode294816007  AllergyIntoleranceCode = "294816007"
	AllergyIntoleranceCode294817003  AllergyIntoleranceCode = "294817003"
	AllergyIntoleranceCode294818008  AllergyIntoleranceCode = "294818008"
	AllergyIntoleranceCode294819000  AllergyIntoleranceCode = "294819000"
	AllergyIntoleranceCode294820006  AllergyIntoleranceCode = "294820006"
	AllergyIntoleranceCode294821005  AllergyIntoleranceCode = "294821005"
	AllergyIntoleranceCode294823008  AllergyIntoleranceCode = "294823008"
	AllergyIntoleranceCode294825001  AllergyIntoleranceCode = "294825001"
	AllergyIntoleranceCode294826000  AllergyIntoleranceCode = "294826000"
	AllergyIntoleranceCode294828004  AllergyIntoleranceCode = "294828004"
	AllergyIntoleranceCode294829007  AllergyIntoleranceCode = "294829007"
	AllergyIntoleranceCode294830002  AllergyIntoleranceCode = "294830002"
	AllergyIntoleranceCode294833000  AllergyIntoleranceCode = "294833000"
	AllergyIntoleranceCode294838009  AllergyIntoleranceCode = "294838009"
	AllergyIntoleranceCode294839001  AllergyIntoleranceCode = "294839001"
	AllergyIntoleranceCode294840004  AllergyIntoleranceCode = "294840004"
	AllergyIntoleranceCode294841000  AllergyIntoleranceCode = "294841000"
	AllergyIntoleranceCode294844008  AllergyIntoleranceCode = "294844008"
	AllergyIntoleranceCode294845009  AllergyIntoleranceCode = "294845009"
	AllergyIntoleranceCode294847001  AllergyIntoleranceCode = "294847001"
	AllergyIntoleranceCode294851004  AllergyIntoleranceCode = "294851004"
	AllergyIntoleranceCode294855008  AllergyIntoleranceCode = "294855008"
	AllergyIntoleranceCode294865002  AllergyIntoleranceCode = "294865002"
	AllergyIntoleranceCode294871008  AllergyIntoleranceCode = "294871008"
	AllergyIntoleranceCode294872001  AllergyIntoleranceCode = "294872001"
	AllergyIntoleranceCode294873006  AllergyIntoleranceCode = "294873006"
	AllergyIntoleranceCode294874000  AllergyIntoleranceCode = "294874000"
	AllergyIntoleranceCode294875004  AllergyIntoleranceCode = "294875004"
	AllergyIntoleranceCode294876003  AllergyIntoleranceCode = "294876003"
	AllergyIntoleranceCode294880008  AllergyIntoleranceCode = "294880008"
	AllergyIntoleranceCode294881007  AllergyIntoleranceCode = "294881007"
	AllergyIntoleranceCode294883005  AllergyIntoleranceCode = "294883005"
	AllergyIntoleranceCode294885003  AllergyIntoleranceCode = "294885003"
	AllergyIntoleranceCode294886002  AllergyIntoleranceCode = "294886002"
	AllergyIntoleranceCode294887006  AllergyIntoleranceCode = "294887006"
	AllergyIntoleranceCode294888001  AllergyIntoleranceCode = "294888001"
	AllergyIntoleranceCode294889009  AllergyIntoleranceCode = "294889009"
	AllergyIntoleranceCode294893003  AllergyIntoleranceCode = "294893003"
	AllergyIntoleranceCode294894009  AllergyIntoleranceCode = "294894009"
	AllergyIntoleranceCode294896006  AllergyIntoleranceCode = "294896006"
	AllergyIntoleranceCode294897002  AllergyIntoleranceCode = "294897002"
	AllergyIntoleranceCode294898007  AllergyIntoleranceCode = "294898007"
	AllergyIntoleranceCode294899004  AllergyIntoleranceCode = "294899004"
	AllergyIntoleranceCode294900009  AllergyIntoleranceCode = "294900009"
	AllergyIntoleranceCode294901008  AllergyIntoleranceCode = "294901008"
	AllergyIntoleranceCode294902001  AllergyIntoleranceCode = "294902001"
	AllergyIntoleranceCode294903006  AllergyIntoleranceCode = "294903006"
	AllergyIntoleranceCode294912008  AllergyIntoleranceCode = "294912008"
	AllergyIntoleranceCode294913003  AllergyIntoleranceCode = "294913003"
	AllergyIntoleranceCode294915005  AllergyIntoleranceCode = "294915005"
	AllergyIntoleranceCode294916006  AllergyIntoleranceCode = "294916006"
	AllergyIntoleranceCode294923007  AllergyIntoleranceCode = "294923007"
	AllergyIntoleranceCode294924001  AllergyIntoleranceCode = "294924001"
	AllergyIntoleranceCode294925000  AllergyIntoleranceCode = "294925000"
	AllergyIntoleranceCode294926004  AllergyIntoleranceCode = "294926004"
	AllergyIntoleranceCode294929006  AllergyIntoleranceCode = "294929006"
	AllergyIntoleranceCode294930001  AllergyIntoleranceCode = "294930001"
	AllergyIntoleranceCode294931002  AllergyIntoleranceCode = "294931002"
	AllergyIntoleranceCode294933004  AllergyIntoleranceCode = "294933004"
	AllergyIntoleranceCode294934005  AllergyIntoleranceCode = "294934005"
	AllergyIntoleranceCode294937003  AllergyIntoleranceCode = "294937003"
	AllergyIntoleranceCode294940003  AllergyIntoleranceCode = "294940003"
	AllergyIntoleranceCode294945008  AllergyIntoleranceCode = "294945008"
	AllergyIntoleranceCode294951003  AllergyIntoleranceCode = "294951003"
	AllergyIntoleranceCode294956008  AllergyIntoleranceCode = "294956008"
	AllergyIntoleranceCode294957004  AllergyIntoleranceCode = "294957004"
	AllergyIntoleranceCode294958009  AllergyIntoleranceCode = "294958009"
	AllergyIntoleranceCode294961005  AllergyIntoleranceCode = "294961005"
	AllergyIntoleranceCode294962003  AllergyIntoleranceCode = "294962003"
	AllergyIntoleranceCode294964002  AllergyIntoleranceCode = "294964002"
	AllergyIntoleranceCode294965001  AllergyIntoleranceCode = "294965001"
	AllergyIntoleranceCode294966000  AllergyIntoleranceCode = "294966000"
	AllergyIntoleranceCode294967009  AllergyIntoleranceCode = "294967009"
	AllergyIntoleranceCode294968004  AllergyIntoleranceCode = "294968004"
	AllergyIntoleranceCode294969007  AllergyIntoleranceCode = "294969007"
	AllergyIntoleranceCode294970008  AllergyIntoleranceCode = "294970008"
	AllergyIntoleranceCode294971007  AllergyIntoleranceCode = "294971007"
	AllergyIntoleranceCode294972000  AllergyIntoleranceCode = "294972000"
	AllergyIntoleranceCode294973005  AllergyIntoleranceCode = "294973005"
	AllergyIntoleranceCode294975003  AllergyIntoleranceCode = "294975003"
	AllergyIntoleranceCode294977006  AllergyIntoleranceCode = "294977006"
	AllergyIntoleranceCode294978001  AllergyIntoleranceCode = "294978001"
	AllergyIntoleranceCode294979009  AllergyIntoleranceCode = "294979009"
	AllergyIntoleranceCode294980007  AllergyIntoleranceCode = "294980007"
	AllergyIntoleranceCode294981006  AllergyIntoleranceCode = "294981006"
	AllergyIntoleranceCode294982004  AllergyIntoleranceCode = "294982004"
	AllergyIntoleranceCode294983009  AllergyIntoleranceCode = "294983009"
	AllergyIntoleranceCode294984003  AllergyIntoleranceCode = "294984003"
	AllergyIntoleranceCode294986001  AllergyIntoleranceCode = "294986001"
	AllergyIntoleranceCode294988000  AllergyIntoleranceCode = "294988000"
	AllergyIntoleranceCode294992007  AllergyIntoleranceCode = "294992007"
	AllergyIntoleranceCode294993002  AllergyIntoleranceCode = "294993002"
	AllergyIntoleranceCode294994008  AllergyIntoleranceCode = "294994008"
	AllergyIntoleranceCode294995009  AllergyIntoleranceCode = "294995009"
	AllergyIntoleranceCode294996005  AllergyIntoleranceCode = "294996005"
	AllergyIntoleranceCode294997001  AllergyIntoleranceCode = "294997001"
	AllergyIntoleranceCode294998006  AllergyIntoleranceCode = "294998006"
	AllergyIntoleranceCode295000003  AllergyIntoleranceCode = "295000003"
	AllergyIntoleranceCode295001004  AllergyIntoleranceCode = "295001004"
	AllergyIntoleranceCode295002006  AllergyIntoleranceCode = "295002006"
	AllergyIntoleranceCode295003001  AllergyIntoleranceCode = "295003001"
	AllergyIntoleranceCode295004007  AllergyIntoleranceCode = "295004007"
	AllergyIntoleranceCode295006009  AllergyIntoleranceCode = "295006009"
	AllergyIntoleranceCode295007000  AllergyIntoleranceCode = "295007000"
	AllergyIntoleranceCode295008005  AllergyIntoleranceCode = "295008005"
	AllergyIntoleranceCode295009002  AllergyIntoleranceCode = "295009002"
	AllergyIntoleranceCode295010007  AllergyIntoleranceCode = "295010007"
	AllergyIntoleranceCode295019008  AllergyIntoleranceCode = "295019008"
	AllergyIntoleranceCode267425008  AllergyIntoleranceCode = "267425008"
	AllergyIntoleranceCode29736007   AllergyIntoleranceCode = "29736007"
	AllergyIntoleranceCode29512005   AllergyIntoleranceCode = "29512005"
	AllergyIntoleranceCode52070001   AllergyIntoleranceCode = "52070001"
	AllergyIntoleranceCode237978005  AllergyIntoleranceCode = "237978005"
	AllergyIntoleranceCode340519003  AllergyIntoleranceCode = "340519003"
	AllergyIntoleranceCode190753003  AllergyIntoleranceCode = "190753003"
	AllergyIntoleranceCode413427002  AllergyIntoleranceCode = "413427002"
	AllergyIntoleranceCode716186003  AllergyIntoleranceCode = "716186003"
	AllergyIntoleranceCode409137002  AllergyIntoleranceCode = "409137002"
	AllergyIntoleranceCode428197003  AllergyIntoleranceCode = "428197003"
	AllergyIntoleranceCode428607008  AllergyIntoleranceCode = "428607008"
	AllergyIntoleranceCode429625007  AllergyIntoleranceCode = "429625007"
	AllergyIntoleranceCode716220001  AllergyIntoleranceCode = "716220001"
	AllergyIntoleranceCode1003774007 AllergyIntoleranceCode = "1003774007"
)

func AllAllergyIntoleranceCode() []AllergyIntoleranceCode {
	return []AllergyIntoleranceCode{
		AllergyIntoleranceCode105590001,
		AllergyIntoleranceCode102002,
		AllergyIntoleranceCode120006,
		AllergyIntoleranceCode125001,
		AllergyIntoleranceCode126000,
		AllergyIntoleranceCode130002,
		AllergyIntoleranceCode131003,
		AllergyIntoleranceCode159002,
		AllergyIntoleranceCode164003,
		AllergyIntoleranceCode178002,
		AllergyIntoleranceCode186002,
		AllergyIntoleranceCode187006,
		AllergyIntoleranceCode200001,
		AllergyIntoleranceCode217008,
		AllergyIntoleranceCode231008,
		AllergyIntoleranceCode238002,
		AllergyIntoleranceCode261000,
		AllergyIntoleranceCode296000,
		AllergyIntoleranceCode322006,
		AllergyIntoleranceCode327000,
		AllergyIntoleranceCode329002,
		AllergyIntoleranceCode336001,
		AllergyIntoleranceCode340005,
		AllergyIntoleranceCode363000,
		AllergyIntoleranceCode370000,
		AllergyIntoleranceCode371001,
		AllergyIntoleranceCode377002,
		AllergyIntoleranceCode392001,
		AllergyIntoleranceCode395004,
		AllergyIntoleranceCode412004,
		AllergyIntoleranceCode424006,
		AllergyIntoleranceCode425007,
		AllergyIntoleranceCode432003,
		AllergyIntoleranceCode438004,
		AllergyIntoleranceCode462009,
		AllergyIntoleranceCode472007,
		AllergyIntoleranceCode476005,
		AllergyIntoleranceCode498001,
		AllergyIntoleranceCode501001,
		AllergyIntoleranceCode505005,
		AllergyIntoleranceCode506006,
		AllergyIntoleranceCode515004,
		AllergyIntoleranceCode519005,
		AllergyIntoleranceCode521000,
		AllergyIntoleranceCode529003,
		AllergyIntoleranceCode538001,
		AllergyIntoleranceCode566009,
		AllergyIntoleranceCode576007,
		AllergyIntoleranceCode578008,
		AllergyIntoleranceCode584006,
		AllergyIntoleranceCode585007,
		AllergyIntoleranceCode591009,
		AllergyIntoleranceCode593007,
		AllergyIntoleranceCode594001,
		AllergyIntoleranceCode597008,
		AllergyIntoleranceCode604000,
		AllergyIntoleranceCode611001,
		AllergyIntoleranceCode620005,
		AllergyIntoleranceCode648005,
		AllergyIntoleranceCode662003,
		AllergyIntoleranceCode668004,
		AllergyIntoleranceCode683009,
		AllergyIntoleranceCode686001,
		AllergyIntoleranceCode693002,
		AllergyIntoleranceCode698006,
		AllergyIntoleranceCode699003,
		AllergyIntoleranceCode704006,
		AllergyIntoleranceCode732002,
		AllergyIntoleranceCode735000,
		AllergyIntoleranceCode747006,
		AllergyIntoleranceCode773001,
		AllergyIntoleranceCode785009,
		AllergyIntoleranceCode804003,
		AllergyIntoleranceCode819002,
		AllergyIntoleranceCode850000,
		AllergyIntoleranceCode859004,
		AllergyIntoleranceCode860009,
		AllergyIntoleranceCode873008,
		AllergyIntoleranceCode876000,
		AllergyIntoleranceCode877009,
		AllergyIntoleranceCode889006,
		AllergyIntoleranceCode896008,
		AllergyIntoleranceCode905001,
		AllergyIntoleranceCode923009,
		AllergyIntoleranceCode925002,
		AllergyIntoleranceCode963005,
		AllergyIntoleranceCode974001,
		AllergyIntoleranceCode979006,
		AllergyIntoleranceCode993004,
		AllergyIntoleranceCode1002007,
		AllergyIntoleranceCode1010008,
		AllergyIntoleranceCode1018001,
		AllergyIntoleranceCode1025008,
		AllergyIntoleranceCode1047008,
		AllergyIntoleranceCode1050006,
		AllergyIntoleranceCode1065007,
		AllergyIntoleranceCode1080001,
		AllergyIntoleranceCode1091008,
		AllergyIntoleranceCode1097007,
		AllergyIntoleranceCode1105007,
		AllergyIntoleranceCode1113008,
		AllergyIntoleranceCode1137008,
		AllergyIntoleranceCode1149009,
		AllergyIntoleranceCode1160000,
		AllergyIntoleranceCode1166006,
		AllergyIntoleranceCode1169004,
		AllergyIntoleranceCode1171004,
		AllergyIntoleranceCode1185009,
		AllergyIntoleranceCode1189003,
		AllergyIntoleranceCode1190007,
		AllergyIntoleranceCode1219001,
		AllergyIntoleranceCode1223009,
		AllergyIntoleranceCode1244009,
		AllergyIntoleranceCode1248007,
		AllergyIntoleranceCode1269009,
		AllergyIntoleranceCode1272002,
		AllergyIntoleranceCode1273007,
		AllergyIntoleranceCode1313002,
		AllergyIntoleranceCode1319003,
		AllergyIntoleranceCode1320009,
		AllergyIntoleranceCode1325004,
		AllergyIntoleranceCode1331001,
		AllergyIntoleranceCode1336006,
		AllergyIntoleranceCode1341003,
		AllergyIntoleranceCode1346008,
		AllergyIntoleranceCode1355006,
		AllergyIntoleranceCode1368003,
		AllergyIntoleranceCode1371006,
		AllergyIntoleranceCode1373009,
		AllergyIntoleranceCode1381005,
		AllergyIntoleranceCode1394007,
		AllergyIntoleranceCode1396009,
		AllergyIntoleranceCode1405004,
		AllergyIntoleranceCode1408002,
		AllergyIntoleranceCode1416006,
		AllergyIntoleranceCode1450002,
		AllergyIntoleranceCode1466000,
		AllergyIntoleranceCode1471007,
		AllergyIntoleranceCode1472000,
		AllergyIntoleranceCode1476002,
		AllergyIntoleranceCode1477006,
		AllergyIntoleranceCode1496005,
		AllergyIntoleranceCode1506001,
		AllergyIntoleranceCode1517000,
		AllergyIntoleranceCode1530004,
		AllergyIntoleranceCode1535009,
		AllergyIntoleranceCode1536005,
		AllergyIntoleranceCode1540001,
		AllergyIntoleranceCode1545006,
		AllergyIntoleranceCode1557002,
		AllergyIntoleranceCode1565004,
		AllergyIntoleranceCode1575001,
		AllergyIntoleranceCode1603001,
		AllergyIntoleranceCode1607000,
		AllergyIntoleranceCode1609002,
		AllergyIntoleranceCode1634002,
		AllergyIntoleranceCode1649005,
		AllergyIntoleranceCode1656004,
		AllergyIntoleranceCode1660001,
		AllergyIntoleranceCode1668008,
		AllergyIntoleranceCode1672007,
		AllergyIntoleranceCode1673002,
		AllergyIntoleranceCode1675009,
		AllergyIntoleranceCode1676005,
		AllergyIntoleranceCode1681001,
		AllergyIntoleranceCode1696002,
		AllergyIntoleranceCode1701009,
		AllergyIntoleranceCode1710001,
		AllergyIntoleranceCode1726000,
		AllergyIntoleranceCode1727009,
		AllergyIntoleranceCode1740004,
		AllergyIntoleranceCode1764003,
		AllergyIntoleranceCode1768000,
		AllergyIntoleranceCode1786002,
		AllergyIntoleranceCode1793003,
		AllergyIntoleranceCode1795005,
		AllergyIntoleranceCode1798007,
		AllergyIntoleranceCode1799004,
		AllergyIntoleranceCode1823002,
		AllergyIntoleranceCode1827001,
		AllergyIntoleranceCode1886008,
		AllergyIntoleranceCode1904005,
		AllergyIntoleranceCode1914001,
		AllergyIntoleranceCode1916004,
		AllergyIntoleranceCode1940007,
		AllergyIntoleranceCode1944003,
		AllergyIntoleranceCode1956002,
		AllergyIntoleranceCode1971003,
		AllergyIntoleranceCode1975007,
		AllergyIntoleranceCode1978009,
		AllergyIntoleranceCode1985008,
		AllergyIntoleranceCode1991005,
		AllergyIntoleranceCode2000001,
		AllergyIntoleranceCode2006007,
		AllergyIntoleranceCode2008008,
		AllergyIntoleranceCode2009000,
		AllergyIntoleranceCode2017008,
		AllergyIntoleranceCode2027002,
		AllergyIntoleranceCode2029004,
		AllergyIntoleranceCode2038002,
		AllergyIntoleranceCode2039005,
		AllergyIntoleranceCode2050008,
		AllergyIntoleranceCode2064008,
		AllergyIntoleranceCode2082006,
		AllergyIntoleranceCode2085008,
		AllergyIntoleranceCode2088005,
		AllergyIntoleranceCode2096000,
		AllergyIntoleranceCode2100004,
		AllergyIntoleranceCode2101000,
		AllergyIntoleranceCode2125008,
		AllergyIntoleranceCode2130007,
		AllergyIntoleranceCode2141009,
		AllergyIntoleranceCode2147008,
		AllergyIntoleranceCode2151005,
		AllergyIntoleranceCode2154002,
		AllergyIntoleranceCode2159007,
		AllergyIntoleranceCode2163000,
		AllergyIntoleranceCode2168009,
		AllergyIntoleranceCode2179004,
		AllergyIntoleranceCode2189000,
		AllergyIntoleranceCode2194000,
		AllergyIntoleranceCode2195004,
		AllergyIntoleranceCode2201007,
		AllergyIntoleranceCode2208001,
		AllergyIntoleranceCode2212007,
		AllergyIntoleranceCode2215009,
		AllergyIntoleranceCode2240002,
		AllergyIntoleranceCode2249001,
		AllergyIntoleranceCode2254005,
		AllergyIntoleranceCode2260005,
		AllergyIntoleranceCode2262002,
		AllergyIntoleranceCode2264001,
		AllergyIntoleranceCode2309006,
		AllergyIntoleranceCode2311002,
		AllergyIntoleranceCode2329007,
		AllergyIntoleranceCode2331003,
		AllergyIntoleranceCode2338009,
		AllergyIntoleranceCode2343002,
		AllergyIntoleranceCode2346005,
		AllergyIntoleranceCode2354007,
		AllergyIntoleranceCode2369008,
		AllergyIntoleranceCode2370009,
		AllergyIntoleranceCode2376003,
		AllergyIntoleranceCode2384004,
		AllergyIntoleranceCode2404002,
		AllergyIntoleranceCode2405001,
		AllergyIntoleranceCode2414006,
		AllergyIntoleranceCode2430003,
		AllergyIntoleranceCode2431004,
		AllergyIntoleranceCode2441001,
		AllergyIntoleranceCode2444009,
		AllergyIntoleranceCode2450004,
		AllergyIntoleranceCode2462000,
		AllergyIntoleranceCode2466002,
		AllergyIntoleranceCode2500009,
		AllergyIntoleranceCode2509005,
		AllergyIntoleranceCode2522002,
		AllergyIntoleranceCode2529006,
		AllergyIntoleranceCode2537003,
		AllergyIntoleranceCode2568004,
		AllergyIntoleranceCode2573005,
		AllergyIntoleranceCode2575003,
		AllergyIntoleranceCode2595009,
		AllergyIntoleranceCode2597001,
		AllergyIntoleranceCode2611008,
		AllergyIntoleranceCode2637006,
		AllergyIntoleranceCode2648004,
		AllergyIntoleranceCode2649007,
		AllergyIntoleranceCode2660003,
		AllergyIntoleranceCode2671002,
		AllergyIntoleranceCode2674005,
		AllergyIntoleranceCode2676007,
		AllergyIntoleranceCode2678008,
		AllergyIntoleranceCode2680002,
		AllergyIntoleranceCode2698003,
		AllergyIntoleranceCode2705002,
		AllergyIntoleranceCode2706001,
		AllergyIntoleranceCode2719002,
		AllergyIntoleranceCode2721007,
		AllergyIntoleranceCode2728001,
		AllergyIntoleranceCode2753003,
		AllergyIntoleranceCode2754009,
		AllergyIntoleranceCode2765004,
		AllergyIntoleranceCode2778004,
		AllergyIntoleranceCode2796008,
		AllergyIntoleranceCode2799001,
		AllergyIntoleranceCode2823004,
		AllergyIntoleranceCode2832002,
		AllergyIntoleranceCode2846002,
		AllergyIntoleranceCode2869004,
		AllergyIntoleranceCode2878005,
		AllergyIntoleranceCode2880004,
		AllergyIntoleranceCode2883002,
		AllergyIntoleranceCode2913009,
		AllergyIntoleranceCode2916001,
		AllergyIntoleranceCode2925007,
		AllergyIntoleranceCode2927004,
		AllergyIntoleranceCode2938004,
		AllergyIntoleranceCode2942001,
		AllergyIntoleranceCode2950005,
		AllergyIntoleranceCode2958003,
		AllergyIntoleranceCode2964005,
		AllergyIntoleranceCode2974008,
		AllergyIntoleranceCode2988007,
		AllergyIntoleranceCode2991007,
		AllergyIntoleranceCode2995003,
		AllergyIntoleranceCode3027009,
		AllergyIntoleranceCode3031003,
		AllergyIntoleranceCode3040004,
		AllergyIntoleranceCode3045009,
		AllergyIntoleranceCode3052006,
		AllergyIntoleranceCode3066001,
		AllergyIntoleranceCode3070009,
		AllergyIntoleranceCode3087006,
		AllergyIntoleranceCode3107005,
		AllergyIntoleranceCode3108000,
		AllergyIntoleranceCode3131000,
		AllergyIntoleranceCode3136005,
		AllergyIntoleranceCode3142009,
		AllergyIntoleranceCode3145006,
		AllergyIntoleranceCode3150000,
		AllergyIntoleranceCode3151001,
		AllergyIntoleranceCode3155005,
		AllergyIntoleranceCode3161008,
		AllergyIntoleranceCode3167007,
		AllergyIntoleranceCode3187008,
		AllergyIntoleranceCode3193000,
		AllergyIntoleranceCode3197004,
		AllergyIntoleranceCode3209002,
		AllergyIntoleranceCode3212004,
		AllergyIntoleranceCode3225007,
		AllergyIntoleranceCode3232003,
		AllergyIntoleranceCode3271000,
		AllergyIntoleranceCode3273002,
		AllergyIntoleranceCode3300001,
		AllergyIntoleranceCode3318003,
		AllergyIntoleranceCode3325005,
		AllergyIntoleranceCode3339005,
		AllergyIntoleranceCode3340007,
		AllergyIntoleranceCode3342004,
		AllergyIntoleranceCode3346001,
		AllergyIntoleranceCode3378009,
		AllergyIntoleranceCode3379001,
		AllergyIntoleranceCode3392003,
		AllergyIntoleranceCode3405005,
		AllergyIntoleranceCode3411008,
		AllergyIntoleranceCode3437006,
		AllergyIntoleranceCode3440006,
		AllergyIntoleranceCode3455002,
		AllergyIntoleranceCode3463001,
		AllergyIntoleranceCode3465008,
		AllergyIntoleranceCode3466009,
		AllergyIntoleranceCode3492002,
		AllergyIntoleranceCode3493007,
		AllergyIntoleranceCode3495000,
		AllergyIntoleranceCode3501003,
		AllergyIntoleranceCode3523004,
		AllergyIntoleranceCode3532002,
		AllergyIntoleranceCode3555004,
		AllergyIntoleranceCode3579002,
		AllergyIntoleranceCode3581000,
		AllergyIntoleranceCode3587001,
		AllergyIntoleranceCode3588006,
		AllergyIntoleranceCode3597005,
		AllergyIntoleranceCode3602003,
		AllergyIntoleranceCode3610002,
		AllergyIntoleranceCode3617004,
		AllergyIntoleranceCode3648007,
		AllergyIntoleranceCode3655009,
		AllergyIntoleranceCode3672002,
		AllergyIntoleranceCode3684000,
		AllergyIntoleranceCode3689005,
		AllergyIntoleranceCode3692009,
		AllergyIntoleranceCode3693004,
		AllergyIntoleranceCode3702007,
		AllergyIntoleranceCode3710008,
		AllergyIntoleranceCode3718001,
		AllergyIntoleranceCode3726009,
		AllergyIntoleranceCode3727000,
		AllergyIntoleranceCode3730007,
		AllergyIntoleranceCode3737005,
		AllergyIntoleranceCode3742002,
		AllergyIntoleranceCode3757009,
		AllergyIntoleranceCode3771001,
		AllergyIntoleranceCode3775005,
		AllergyIntoleranceCode3776006,
		AllergyIntoleranceCode3792001,
		AllergyIntoleranceCode3800009,
		AllergyIntoleranceCode3807007,
		AllergyIntoleranceCode3811001,
		AllergyIntoleranceCode3812008,
		AllergyIntoleranceCode3816006,
		AllergyIntoleranceCode3823007,
		AllergyIntoleranceCode3829006,
		AllergyIntoleranceCode3834005,
		AllergyIntoleranceCode3836007,
		AllergyIntoleranceCode3844007,
		AllergyIntoleranceCode3848005,
		AllergyIntoleranceCode3849002,
		AllergyIntoleranceCode3854006,
		AllergyIntoleranceCode3874004,
		AllergyIntoleranceCode3892007,
		AllergyIntoleranceCode3896005,
		AllergyIntoleranceCode3897001,
		AllergyIntoleranceCode3906002,
		AllergyIntoleranceCode3920009,
		AllergyIntoleranceCode3930000,
		AllergyIntoleranceCode3932008,
		AllergyIntoleranceCode3941003,
		AllergyIntoleranceCode3945007,
		AllergyIntoleranceCode3958008,
		AllergyIntoleranceCode3961009,
		AllergyIntoleranceCode3976001,
		AllergyIntoleranceCode3982003,
		AllergyIntoleranceCode3983008,
		AllergyIntoleranceCode3990003,
		AllergyIntoleranceCode3994007,
		AllergyIntoleranceCode4014000,
		AllergyIntoleranceCode4024008,
		AllergyIntoleranceCode4025009,
		AllergyIntoleranceCode4043008,
		AllergyIntoleranceCode4047009,
		AllergyIntoleranceCode4048004,
		AllergyIntoleranceCode4067000,
		AllergyIntoleranceCode4076007,
		AllergyIntoleranceCode4077003,
		AllergyIntoleranceCode4080002,
		AllergyIntoleranceCode4091009,
		AllergyIntoleranceCode4097008,
		AllergyIntoleranceCode4104007,
		AllergyIntoleranceCode4105008,
		AllergyIntoleranceCode4115002,
		AllergyIntoleranceCode4137009,
		AllergyIntoleranceCode4153007,
		AllergyIntoleranceCode4167003,
		AllergyIntoleranceCode4169000,
		AllergyIntoleranceCode4177001,
		AllergyIntoleranceCode4182008,
		AllergyIntoleranceCode4188007,
		AllergyIntoleranceCode4200007,
		AllergyIntoleranceCode4201006,
		AllergyIntoleranceCode4203009,
		AllergyIntoleranceCode4207005,
		AllergyIntoleranceCode4217000,
		AllergyIntoleranceCode4218005,
		AllergyIntoleranceCode4231000,
		AllergyIntoleranceCode4239003,
		AllergyIntoleranceCode4255005,
		AllergyIntoleranceCode4289006,
		AllergyIntoleranceCode4290002,
		AllergyIntoleranceCode4314009,
		AllergyIntoleranceCode4334005,
		AllergyIntoleranceCode4342006,
		AllergyIntoleranceCode4353000,
		AllergyIntoleranceCode4355007,
		AllergyIntoleranceCode4362003,
		AllergyIntoleranceCode4370008,
		AllergyIntoleranceCode4393002,
		AllergyIntoleranceCode4401009,
		AllergyIntoleranceCode4422003,
		AllergyIntoleranceCode4423008,
		AllergyIntoleranceCode4425001,
		AllergyIntoleranceCode4435007,
		AllergyIntoleranceCode4437004,
		AllergyIntoleranceCode4471008,
		AllergyIntoleranceCode4479005,
		AllergyIntoleranceCode4480008,
		AllergyIntoleranceCode4509009,
		AllergyIntoleranceCode4534009,
		AllergyIntoleranceCode4540002,
		AllergyIntoleranceCode4546008,
		AllergyIntoleranceCode4555006,
		AllergyIntoleranceCode4560005,
		AllergyIntoleranceCode4561009,
		AllergyIntoleranceCode4564001,
		AllergyIntoleranceCode4567008,
		AllergyIntoleranceCode4582003,
		AllergyIntoleranceCode4591004,
		AllergyIntoleranceCode4610008,
		AllergyIntoleranceCode4616002,
		AllergyIntoleranceCode4629002,
		AllergyIntoleranceCode4635002,
		AllergyIntoleranceCode4643007,
		AllergyIntoleranceCode4656000,
		AllergyIntoleranceCode4674009,
		AllergyIntoleranceCode4681002,
		AllergyIntoleranceCode4693006,
		AllergyIntoleranceCode4700006,
		AllergyIntoleranceCode4706000,
		AllergyIntoleranceCode4714006,
		AllergyIntoleranceCode4728000,
		AllergyIntoleranceCode4732006,
		AllergyIntoleranceCode4746006,
		AllergyIntoleranceCode4761007,
		AllergyIntoleranceCode4762000,
		AllergyIntoleranceCode4777008,
		AllergyIntoleranceCode4780009,
		AllergyIntoleranceCode4786003,
		AllergyIntoleranceCode4789005,
		AllergyIntoleranceCode4793004,
		AllergyIntoleranceCode4814001,
		AllergyIntoleranceCode4824009,
		AllergyIntoleranceCode4825005,
		AllergyIntoleranceCode4831008,
		AllergyIntoleranceCode4832001,
		AllergyIntoleranceCode4833006,
		AllergyIntoleranceCode4844003,
		AllergyIntoleranceCode4864008,
		AllergyIntoleranceCode4872005,
		AllergyIntoleranceCode4878009,
		AllergyIntoleranceCode4882006,
		AllergyIntoleranceCode4889002,
		AllergyIntoleranceCode4901003,
		AllergyIntoleranceCode4925006,
		AllergyIntoleranceCode4933007,
		AllergyIntoleranceCode4940008,
		AllergyIntoleranceCode4955004,
		AllergyIntoleranceCode4962008,
		AllergyIntoleranceCode4963003,
		AllergyIntoleranceCode4965005,
		AllergyIntoleranceCode4968007,
		AllergyIntoleranceCode4986005,
		AllergyIntoleranceCode5003005,
		AllergyIntoleranceCode5004004,
		AllergyIntoleranceCode5007006,
		AllergyIntoleranceCode5024000,
		AllergyIntoleranceCode5031001,
		AllergyIntoleranceCode5040002,
		AllergyIntoleranceCode5043000,
		AllergyIntoleranceCode5045007,
		AllergyIntoleranceCode5059000,
		AllergyIntoleranceCode5060005,
		AllergyIntoleranceCode5061009,
		AllergyIntoleranceCode5064001,
		AllergyIntoleranceCode5081005,
		AllergyIntoleranceCode5086000,
		AllergyIntoleranceCode5094007,
		AllergyIntoleranceCode5098005,
		AllergyIntoleranceCode5109006,
		AllergyIntoleranceCode5142007,
		AllergyIntoleranceCode5160007,
		AllergyIntoleranceCode5163009,
		AllergyIntoleranceCode5167005,
		AllergyIntoleranceCode5172001,
		AllergyIntoleranceCode5179005,
		AllergyIntoleranceCode5200001,
		AllergyIntoleranceCode5206007,
		AllergyIntoleranceCode5220000,
		AllergyIntoleranceCode5226006,
		AllergyIntoleranceCode5250008,
		AllergyIntoleranceCode5252000,
		AllergyIntoleranceCode5253005,
		AllergyIntoleranceCode5259009,
		AllergyIntoleranceCode5289002,
		AllergyIntoleranceCode5303002,
		AllergyIntoleranceCode5305009,
		AllergyIntoleranceCode5307001,
		AllergyIntoleranceCode5312000,
		AllergyIntoleranceCode5323001,
		AllergyIntoleranceCode5330007,
		AllergyIntoleranceCode5339008,
		AllergyIntoleranceCode5340005,
		AllergyIntoleranceCode5392001,
		AllergyIntoleranceCode5395004,
		AllergyIntoleranceCode5404007,
		AllergyIntoleranceCode5405008,
		AllergyIntoleranceCode5406009,
		AllergyIntoleranceCode5420002,
		AllergyIntoleranceCode5439007,
		AllergyIntoleranceCode5442001,
		AllergyIntoleranceCode5453007,
		AllergyIntoleranceCode5471000,
		AllergyIntoleranceCode5474008,
		AllergyIntoleranceCode5477001,
		AllergyIntoleranceCode5483003,
		AllergyIntoleranceCode5504009,
		AllergyIntoleranceCode5511008,
		AllergyIntoleranceCode5513006,
		AllergyIntoleranceCode5515004,
		AllergyIntoleranceCode5533005,
		AllergyIntoleranceCode5537006,
		AllergyIntoleranceCode5540006,
		AllergyIntoleranceCode5547009,
		AllergyIntoleranceCode5548004,
		AllergyIntoleranceCode5568005,
		AllergyIntoleranceCode5573004,
		AllergyIntoleranceCode5589001,
		AllergyIntoleranceCode5590005,
		AllergyIntoleranceCode5628003,
		AllergyIntoleranceCode5629006,
		AllergyIntoleranceCode5637003,
		AllergyIntoleranceCode5641004,
		AllergyIntoleranceCode5647000,
		AllergyIntoleranceCode5656008,
		AllergyIntoleranceCode5659001,
		AllergyIntoleranceCode5670008,
		AllergyIntoleranceCode5681006,
		AllergyIntoleranceCode5691000,
		AllergyIntoleranceCode5692007,
		AllergyIntoleranceCode5699003,
		AllergyIntoleranceCode5700002,
		AllergyIntoleranceCode5702005,
		AllergyIntoleranceCode5704006,
		AllergyIntoleranceCode5705007,
		AllergyIntoleranceCode5739006,
		AllergyIntoleranceCode5746002,
		AllergyIntoleranceCode5757007,
		AllergyIntoleranceCode5762008,
		AllergyIntoleranceCode5764009,
		AllergyIntoleranceCode5767002,
		AllergyIntoleranceCode5774007,
		AllergyIntoleranceCode5800007,
		AllergyIntoleranceCode5813001,
		AllergyIntoleranceCode5826002,
		AllergyIntoleranceCode5827006,
		AllergyIntoleranceCode5829009,
		AllergyIntoleranceCode5830004,
		AllergyIntoleranceCode5840001,
		AllergyIntoleranceCode5858007,
		AllergyIntoleranceCode5863006,
		AllergyIntoleranceCode5896008,
		AllergyIntoleranceCode5899001,
		AllergyIntoleranceCode5907009,
		AllergyIntoleranceCode5910002,
		AllergyIntoleranceCode5915007,
		AllergyIntoleranceCode5927005,
		AllergyIntoleranceCode5931004,
		AllergyIntoleranceCode5932006,
		AllergyIntoleranceCode5950004,
		AllergyIntoleranceCode5955009,
		AllergyIntoleranceCode5977008,
		AllergyIntoleranceCode5989005,
		AllergyIntoleranceCode5991002,
		AllergyIntoleranceCode6021003,
		AllergyIntoleranceCode6038004,
		AllergyIntoleranceCode6043006,
		AllergyIntoleranceCode6044000,
		AllergyIntoleranceCode6054001,
		AllergyIntoleranceCode6056004,
		AllergyIntoleranceCode6068008,
		AllergyIntoleranceCode6083003,
		AllergyIntoleranceCode6085005,
		AllergyIntoleranceCode6088007,
		AllergyIntoleranceCode6089004,
		AllergyIntoleranceCode6091007,
		AllergyIntoleranceCode6107003,
		AllergyIntoleranceCode6109000,
		AllergyIntoleranceCode6115000,
		AllergyIntoleranceCode6135004,
		AllergyIntoleranceCode6138002,
		AllergyIntoleranceCode6162007,
		AllergyIntoleranceCode6170002,
		AllergyIntoleranceCode6172005,
		AllergyIntoleranceCode6178009,
		AllergyIntoleranceCode6179001,
		AllergyIntoleranceCode6182006,
		AllergyIntoleranceCode6197009,
		AllergyIntoleranceCode6237004,
		AllergyIntoleranceCode6249003,
		AllergyIntoleranceCode6256009,
		AllergyIntoleranceCode6257000,
		AllergyIntoleranceCode6260007,
		AllergyIntoleranceCode6261006,
		AllergyIntoleranceCode6263009,
		AllergyIntoleranceCode6264003,
		AllergyIntoleranceCode6287006,
		AllergyIntoleranceCode6291001,
		AllergyIntoleranceCode6301006,
		AllergyIntoleranceCode6310003,
		AllergyIntoleranceCode6314007,
		AllergyIntoleranceCode6333002,
		AllergyIntoleranceCode6338006,
		AllergyIntoleranceCode6356006,
		AllergyIntoleranceCode6360009,
		AllergyIntoleranceCode6367007,
		AllergyIntoleranceCode6386004,
		AllergyIntoleranceCode6394006,
		AllergyIntoleranceCode6401007,
		AllergyIntoleranceCode6409009,
		AllergyIntoleranceCode6411000,
		AllergyIntoleranceCode6422001,
		AllergyIntoleranceCode6451002,
		AllergyIntoleranceCode6455006,
		AllergyIntoleranceCode6469006,
		AllergyIntoleranceCode6478000,
		AllergyIntoleranceCode6495008,
		AllergyIntoleranceCode6507009,
		AllergyIntoleranceCode6513000,
		AllergyIntoleranceCode6516008,
		AllergyIntoleranceCode6524003,
		AllergyIntoleranceCode6529008,
		AllergyIntoleranceCode6532006,
		AllergyIntoleranceCode6590001,
		AllergyIntoleranceCode6592009,
		AllergyIntoleranceCode6602005,
		AllergyIntoleranceCode6611005,
		AllergyIntoleranceCode6612003,
		AllergyIntoleranceCode6619007,
		AllergyIntoleranceCode6642000,
		AllergyIntoleranceCode6644004,
		AllergyIntoleranceCode6671004,
		AllergyIntoleranceCode6672006,
		AllergyIntoleranceCode6699008,
		AllergyIntoleranceCode6701008,
		AllergyIntoleranceCode6702001,
		AllergyIntoleranceCode6709005,
		AllergyIntoleranceCode6710000,
		AllergyIntoleranceCode6713003,
		AllergyIntoleranceCode6717002,
		AllergyIntoleranceCode6725000,
		AllergyIntoleranceCode6730001,
		AllergyIntoleranceCode6741004,
		AllergyIntoleranceCode6755007,
		AllergyIntoleranceCode6786001,
		AllergyIntoleranceCode6790004,
		AllergyIntoleranceCode6792007,
		AllergyIntoleranceCode6808006,
		AllergyIntoleranceCode6809003,
		AllergyIntoleranceCode6814004,
		AllergyIntoleranceCode6817006,
		AllergyIntoleranceCode6826009,
		AllergyIntoleranceCode6837005,
		AllergyIntoleranceCode6854002,
		AllergyIntoleranceCode6865007,
		AllergyIntoleranceCode6873003,
		AllergyIntoleranceCode6879004,
		AllergyIntoleranceCode6881002,
		AllergyIntoleranceCode6884005,
		AllergyIntoleranceCode6890009,
		AllergyIntoleranceCode6896003,
		AllergyIntoleranceCode6910009,
		AllergyIntoleranceCode6911008,
		AllergyIntoleranceCode6916003,
		AllergyIntoleranceCode6924008,
		AllergyIntoleranceCode6925009,
		AllergyIntoleranceCode6927001,
		AllergyIntoleranceCode6937006,
		AllergyIntoleranceCode6945001,
		AllergyIntoleranceCode6952004,
		AllergyIntoleranceCode6958000,
		AllergyIntoleranceCode6961004,
		AllergyIntoleranceCode6970001,
		AllergyIntoleranceCode6973004,
		AllergyIntoleranceCode6983000,
		AllergyIntoleranceCode6992002,
		AllergyIntoleranceCode6993007,
		AllergyIntoleranceCode6999006,
		AllergyIntoleranceCode7008002,
		AllergyIntoleranceCode7018007,
		AllergyIntoleranceCode7029006,
		AllergyIntoleranceCode7030001,
		AllergyIntoleranceCode7034005,
		AllergyIntoleranceCode7045008,
		AllergyIntoleranceCode7047000,
		AllergyIntoleranceCode7049002,
		AllergyIntoleranceCode7054006,
		AllergyIntoleranceCode7056008,
		AllergyIntoleranceCode7059001,
		AllergyIntoleranceCode7061005,
		AllergyIntoleranceCode7070008,
		AllergyIntoleranceCode7084003,
		AllergyIntoleranceCode7110002,
		AllergyIntoleranceCode7120007,
		AllergyIntoleranceCode7132006,
		AllergyIntoleranceCode7139002,
		AllergyIntoleranceCode7146006,
		AllergyIntoleranceCode7152007,
		AllergyIntoleranceCode7156005,
		AllergyIntoleranceCode7158006,
		AllergyIntoleranceCode7161007,
		AllergyIntoleranceCode7179006,
		AllergyIntoleranceCode7191002,
		AllergyIntoleranceCode7208009,
		AllergyIntoleranceCode7211005,
		AllergyIntoleranceCode7237008,
		AllergyIntoleranceCode7243005,
		AllergyIntoleranceCode7269004,
		AllergyIntoleranceCode7271004,
		AllergyIntoleranceCode7280004,
		AllergyIntoleranceCode7281000,
		AllergyIntoleranceCode7284008,
		AllergyIntoleranceCode7294003,
		AllergyIntoleranceCode7302008,
		AllergyIntoleranceCode7318002,
		AllergyIntoleranceCode7321000,
		AllergyIntoleranceCode7325009,
		AllergyIntoleranceCode7327001,
		AllergyIntoleranceCode7328006,
		AllergyIntoleranceCode7330008,
		AllergyIntoleranceCode7337006,
		AllergyIntoleranceCode7348004,
		AllergyIntoleranceCode7382005,
		AllergyIntoleranceCode7401000,
		AllergyIntoleranceCode7411007,
		AllergyIntoleranceCode7427000,
		AllergyIntoleranceCode7434003,
		AllergyIntoleranceCode7446004,
		AllergyIntoleranceCode7460002,
		AllergyIntoleranceCode7470000,
		AllergyIntoleranceCode7489000,
		AllergyIntoleranceCode7503004,
		AllergyIntoleranceCode7509000,
		AllergyIntoleranceCode7515000,
		AllergyIntoleranceCode7537007,
		AllergyIntoleranceCode7547005,
		AllergyIntoleranceCode7549008,
		AllergyIntoleranceCode7588005,
		AllergyIntoleranceCode7608003,
		AllergyIntoleranceCode7616007,
		AllergyIntoleranceCode7648006,
		AllergyIntoleranceCode7661006,
		AllergyIntoleranceCode7670009,
		AllergyIntoleranceCode7675004,
		AllergyIntoleranceCode7685003,
		AllergyIntoleranceCode7696006,
		AllergyIntoleranceCode7716001,
		AllergyIntoleranceCode7737009,
		AllergyIntoleranceCode7738004,
		AllergyIntoleranceCode7761002,
		AllergyIntoleranceCode7770004,
		AllergyIntoleranceCode7774008,
		AllergyIntoleranceCode7779003,
		AllergyIntoleranceCode7785005,
		AllergyIntoleranceCode7790008,
		AllergyIntoleranceCode7791007,
		AllergyIntoleranceCode7795003,
		AllergyIntoleranceCode7801007,
		AllergyIntoleranceCode7816005,
		AllergyIntoleranceCode7834009,
		AllergyIntoleranceCode7846008,
		AllergyIntoleranceCode7848009,
		AllergyIntoleranceCode7868003,
		AllergyIntoleranceCode7879008,
		AllergyIntoleranceCode7900007,
		AllergyIntoleranceCode7904003,
		AllergyIntoleranceCode7909008,
		AllergyIntoleranceCode7924004,
		AllergyIntoleranceCode7938006,
		AllergyIntoleranceCode7945006,
		AllergyIntoleranceCode7948008,
		AllergyIntoleranceCode7953003,
		AllergyIntoleranceCode7957002,
		AllergyIntoleranceCode7961008,
		AllergyIntoleranceCode7970006,
		AllergyIntoleranceCode7974002,
		AllergyIntoleranceCode7975001,
		AllergyIntoleranceCode7979007,
		AllergyIntoleranceCode7983007,
		AllergyIntoleranceCode7985000,
		AllergyIntoleranceCode7997004,
		AllergyIntoleranceCode8000007,
		AllergyIntoleranceCode8002004,
		AllergyIntoleranceCode8025003,
		AllergyIntoleranceCode8029009,
		AllergyIntoleranceCode8030004,
		AllergyIntoleranceCode8035009,
		AllergyIntoleranceCode8048008,
		AllergyIntoleranceCode8054009,
		AllergyIntoleranceCode8055005,
		AllergyIntoleranceCode8105004,
		AllergyIntoleranceCode8108002,
		AllergyIntoleranceCode8123007,
		AllergyIntoleranceCode8132009,
		AllergyIntoleranceCode8143001,
		AllergyIntoleranceCode8153000,
		AllergyIntoleranceCode8156008,
		AllergyIntoleranceCode8164002,
		AllergyIntoleranceCode8168004,
		AllergyIntoleranceCode8179009,
		AllergyIntoleranceCode8184003,
		AllergyIntoleranceCode8190004,
		AllergyIntoleranceCode8202008,
		AllergyIntoleranceCode8203003,
		AllergyIntoleranceCode8204009,
		AllergyIntoleranceCode8222007,
		AllergyIntoleranceCode8227001,
		AllergyIntoleranceCode8230008,
		AllergyIntoleranceCode8237006,
		AllergyIntoleranceCode8252004,
		AllergyIntoleranceCode8257005,
		AllergyIntoleranceCode8261004,
		AllergyIntoleranceCode8268005,
		AllergyIntoleranceCode8270001,
		AllergyIntoleranceCode8275006,
		AllergyIntoleranceCode8295000,
		AllergyIntoleranceCode8300003,
		AllergyIntoleranceCode8310007,
		AllergyIntoleranceCode8313009,
		AllergyIntoleranceCode8340009,
		AllergyIntoleranceCode8342001,
		AllergyIntoleranceCode8343006,
		AllergyIntoleranceCode8354001,
		AllergyIntoleranceCode8355000,
		AllergyIntoleranceCode8362009,
		AllergyIntoleranceCode8365006,
		AllergyIntoleranceCode8368008,
		AllergyIntoleranceCode8376005,
		AllergyIntoleranceCode8385005,
		AllergyIntoleranceCode8397006,
		AllergyIntoleranceCode8406008,
		AllergyIntoleranceCode8429000,
		AllergyIntoleranceCode8450009,
		AllergyIntoleranceCode8452001,
		AllergyIntoleranceCode8456003,
		AllergyIntoleranceCode8460000,
		AllergyIntoleranceCode8473001,
		AllergyIntoleranceCode8474007,
		AllergyIntoleranceCode8484008,
		AllergyIntoleranceCode8485009,
		AllergyIntoleranceCode8486005,
		AllergyIntoleranceCode8487001,
		AllergyIntoleranceCode8491006,
		AllergyIntoleranceCode8492004,
		AllergyIntoleranceCode8498000,
		AllergyIntoleranceCode8507001,
		AllergyIntoleranceCode8514004,
		AllergyIntoleranceCode8520003,
		AllergyIntoleranceCode8525008,
		AllergyIntoleranceCode8529002,
		AllergyIntoleranceCode8534003,
		AllergyIntoleranceCode8537005,
		AllergyIntoleranceCode8578007,
		AllergyIntoleranceCode8591008,
		AllergyIntoleranceCode8612007,
		AllergyIntoleranceCode8620009,
		AllergyIntoleranceCode8631001,
		AllergyIntoleranceCode8653004,
		AllergyIntoleranceCode8660005,
		AllergyIntoleranceCode8687009,
		AllergyIntoleranceCode8689007,
		AllergyIntoleranceCode8701002,
		AllergyIntoleranceCode8705006,
		AllergyIntoleranceCode8731008,
		AllergyIntoleranceCode8740007,
		AllergyIntoleranceCode8761000,
		AllergyIntoleranceCode8767001,
		AllergyIntoleranceCode8785008,
		AllergyIntoleranceCode8786009,
		AllergyIntoleranceCode8795001,
		AllergyIntoleranceCode8817004,
		AllergyIntoleranceCode8818009,
		AllergyIntoleranceCode8822004,
		AllergyIntoleranceCode8830003,
		AllergyIntoleranceCode8836009,
		AllergyIntoleranceCode8844009,
		AllergyIntoleranceCode8858006,
		AllergyIntoleranceCode8865003,
		AllergyIntoleranceCode8878003,
		AllergyIntoleranceCode8882001,
		AllergyIntoleranceCode8886003,
		AllergyIntoleranceCode8908003,
		AllergyIntoleranceCode8914005,
		AllergyIntoleranceCode8919000,
		AllergyIntoleranceCode8926000,
		AllergyIntoleranceCode8945009,
		AllergyIntoleranceCode8953001,
		AllergyIntoleranceCode8963009,
		AllergyIntoleranceCode8969008,
		AllergyIntoleranceCode8977007,
		AllergyIntoleranceCode8982000,
		AllergyIntoleranceCode8987006,
		AllergyIntoleranceCode8991001,
		AllergyIntoleranceCode9010006,
		AllergyIntoleranceCode9013008,
		AllergyIntoleranceCode9021002,
		AllergyIntoleranceCode9024005,
		AllergyIntoleranceCode9045003,
		AllergyIntoleranceCode9052001,
		AllergyIntoleranceCode9054000,
		AllergyIntoleranceCode9103003,
		AllergyIntoleranceCode9110009,
		AllergyIntoleranceCode9125009,
		AllergyIntoleranceCode9159008,
		AllergyIntoleranceCode9172009,
		AllergyIntoleranceCode9174005,
		AllergyIntoleranceCode9183000,
		AllergyIntoleranceCode9189001,
		AllergyIntoleranceCode9195000,
		AllergyIntoleranceCode9197008,
		AllergyIntoleranceCode9205004,
		AllergyIntoleranceCode9220005,
		AllergyIntoleranceCode9223007,
		AllergyIntoleranceCode9234005,
		AllergyIntoleranceCode9246009,
		AllergyIntoleranceCode9253000,
		AllergyIntoleranceCode9270008,
		AllergyIntoleranceCode9271007,
		AllergyIntoleranceCode9301005,
		AllergyIntoleranceCode9302003,
		AllergyIntoleranceCode9315007,
		AllergyIntoleranceCode9319001,
		AllergyIntoleranceCode9334007,
		AllergyIntoleranceCode9349004,
		AllergyIntoleranceCode9351000,
		AllergyIntoleranceCode9355009,
		AllergyIntoleranceCode9392009,
		AllergyIntoleranceCode9396007,
		AllergyIntoleranceCode9398008,
		AllergyIntoleranceCode9410003,
		AllergyIntoleranceCode9422000,
		AllergyIntoleranceCode418038007,
		AllergyIntoleranceCode25744000,
		AllergyIntoleranceCode54250004,
		AllergyIntoleranceCode59037007,
		AllergyIntoleranceCode61712006,
		AllergyIntoleranceCode72354005,
		AllergyIntoleranceCode91931000,
		AllergyIntoleranceCode91932007,
		AllergyIntoleranceCode91934008,
		AllergyIntoleranceCode91935009,
		AllergyIntoleranceCode91936005,
		AllergyIntoleranceCode91937001,
		AllergyIntoleranceCode91938006,
		AllergyIntoleranceCode91939003,
		AllergyIntoleranceCode91940001,
		AllergyIntoleranceCode190751001,
		AllergyIntoleranceCode213020009,
		AllergyIntoleranceCode232349006,
		AllergyIntoleranceCode232350006,
		AllergyIntoleranceCode235719002,
		AllergyIntoleranceCode293584003,
		AllergyIntoleranceCode293585002,
		AllergyIntoleranceCode293586001,
		AllergyIntoleranceCode293588000,
		AllergyIntoleranceCode293589008,
		AllergyIntoleranceCode293590004,
		AllergyIntoleranceCode293591000,
		AllergyIntoleranceCode293592007,
		AllergyIntoleranceCode293593002,
		AllergyIntoleranceCode293594008,
		AllergyIntoleranceCode293595009,
		AllergyIntoleranceCode293596005,
		AllergyIntoleranceCode293597001,
		AllergyIntoleranceCode293598006,
		AllergyIntoleranceCode293599003,
		AllergyIntoleranceCode293600000,
		AllergyIntoleranceCode293601001,
		AllergyIntoleranceCode293602008,
		AllergyIntoleranceCode293603003,
		AllergyIntoleranceCode293604009,
		AllergyIntoleranceCode293605005,
		AllergyIntoleranceCode293606006,
		AllergyIntoleranceCode293607002,
		AllergyIntoleranceCode293608007,
		AllergyIntoleranceCode293609004,
		AllergyIntoleranceCode293610009,
		AllergyIntoleranceCode293611008,
		AllergyIntoleranceCode293612001,
		AllergyIntoleranceCode293613006,
		AllergyIntoleranceCode293614000,
		AllergyIntoleranceCode293615004,
		AllergyIntoleranceCode293616003,
		AllergyIntoleranceCode293617007,
		AllergyIntoleranceCode293618002,
		AllergyIntoleranceCode293619005,
		AllergyIntoleranceCode293620004,
		AllergyIntoleranceCode293621000,
		AllergyIntoleranceCode293622007,
		AllergyIntoleranceCode293623002,
		AllergyIntoleranceCode293624008,
		AllergyIntoleranceCode293625009,
		AllergyIntoleranceCode293626005,
		AllergyIntoleranceCode293627001,
		AllergyIntoleranceCode293628006,
		AllergyIntoleranceCode293629003,
		AllergyIntoleranceCode293630008,
		AllergyIntoleranceCode293631007,
		AllergyIntoleranceCode293632000,
		AllergyIntoleranceCode293633005,
		AllergyIntoleranceCode293635003,
		AllergyIntoleranceCode293636002,
		AllergyIntoleranceCode293637006,
		AllergyIntoleranceCode293645001,
		AllergyIntoleranceCode293646000,
		AllergyIntoleranceCode293647009,
		AllergyIntoleranceCode293648004,
		AllergyIntoleranceCode293649007,
		AllergyIntoleranceCode293650007,
		AllergyIntoleranceCode293651006,
		AllergyIntoleranceCode293652004,
		AllergyIntoleranceCode293653009,
		AllergyIntoleranceCode293654003,
		AllergyIntoleranceCode293655002,
		AllergyIntoleranceCode293656001,
		AllergyIntoleranceCode293657005,
		AllergyIntoleranceCode293658000,
		AllergyIntoleranceCode293659008,
		AllergyIntoleranceCode293660003,
		AllergyIntoleranceCode293662006,
		AllergyIntoleranceCode293663001,
		AllergyIntoleranceCode293665008,
		AllergyIntoleranceCode293666009,
		AllergyIntoleranceCode293668005,
		AllergyIntoleranceCode293669002,
		AllergyIntoleranceCode293671002,
		AllergyIntoleranceCode293673004,
		AllergyIntoleranceCode293674005,
		AllergyIntoleranceCode293675006,
		AllergyIntoleranceCode293676007,
		AllergyIntoleranceCode293678008,
		AllergyIntoleranceCode293679000,
		AllergyIntoleranceCode293680002,
		AllergyIntoleranceCode293681003,
		AllergyIntoleranceCode293682005,
		AllergyIntoleranceCode293685007,
		AllergyIntoleranceCode293686008,
		AllergyIntoleranceCode293687004,
		AllergyIntoleranceCode293690005,
		AllergyIntoleranceCode293691009,
		AllergyIntoleranceCode293692002,
		AllergyIntoleranceCode293693007,
		AllergyIntoleranceCode293694001,
		AllergyIntoleranceCode293695000,
		AllergyIntoleranceCode293696004,
		AllergyIntoleranceCode293697008,
		AllergyIntoleranceCode293699006,
		AllergyIntoleranceCode293700007,
		AllergyIntoleranceCode293701006,
		AllergyIntoleranceCode293706001,
		AllergyIntoleranceCode293707005,
		AllergyIntoleranceCode293708000,
		AllergyIntoleranceCode293709008,
		AllergyIntoleranceCode293710003,
		AllergyIntoleranceCode293712006,
		AllergyIntoleranceCode293714007,
		AllergyIntoleranceCode293715008,
		AllergyIntoleranceCode293716009,
		AllergyIntoleranceCode293717000,
		AllergyIntoleranceCode293718005,
		AllergyIntoleranceCode293719002,
		AllergyIntoleranceCode293720008,
		AllergyIntoleranceCode293721007,
		AllergyIntoleranceCode293722000,
		AllergyIntoleranceCode293723005,
		AllergyIntoleranceCode293724004,
		AllergyIntoleranceCode293725003,
		AllergyIntoleranceCode293726002,
		AllergyIntoleranceCode293727006,
		AllergyIntoleranceCode293728001,
		AllergyIntoleranceCode293732007,
		AllergyIntoleranceCode293733002,
		AllergyIntoleranceCode293735009,
		AllergyIntoleranceCode293736005,
		AllergyIntoleranceCode293737001,
		AllergyIntoleranceCode293738006,
		AllergyIntoleranceCode293740001,
		AllergyIntoleranceCode293741002,
		AllergyIntoleranceCode293742009,
		AllergyIntoleranceCode293743004,
		AllergyIntoleranceCode293745006,
		AllergyIntoleranceCode293746007,
		AllergyIntoleranceCode293747003,
		AllergyIntoleranceCode293748008,
		AllergyIntoleranceCode293749000,
		AllergyIntoleranceCode293750000,
		AllergyIntoleranceCode293751001,
		AllergyIntoleranceCode293752008,
		AllergyIntoleranceCode293753003,
		AllergyIntoleranceCode293754009,
		AllergyIntoleranceCode293755005,
		AllergyIntoleranceCode293756006,
		AllergyIntoleranceCode293758007,
		AllergyIntoleranceCode293760009,
		AllergyIntoleranceCode293761008,
		AllergyIntoleranceCode293762001,
		AllergyIntoleranceCode293763006,
		AllergyIntoleranceCode293764000,
		AllergyIntoleranceCode293765004,
		AllergyIntoleranceCode293766003,
		AllergyIntoleranceCode293767007,
		AllergyIntoleranceCode293768002,
		AllergyIntoleranceCode293770006,
		AllergyIntoleranceCode293771005,
		AllergyIntoleranceCode293772003,
		AllergyIntoleranceCode293773008,
		AllergyIntoleranceCode293774002,
		AllergyIntoleranceCode293775001,
		AllergyIntoleranceCode293776000,
		AllergyIntoleranceCode293777009,
		AllergyIntoleranceCode293778004,
		AllergyIntoleranceCode293779007,
		AllergyIntoleranceCode293780005,
		AllergyIntoleranceCode293781009,
		AllergyIntoleranceCode293782002,
		AllergyIntoleranceCode293783007,
		AllergyIntoleranceCode293784001,
		AllergyIntoleranceCode293785000,
		AllergyIntoleranceCode293786004,
		AllergyIntoleranceCode293787008,
		AllergyIntoleranceCode293788003,
		AllergyIntoleranceCode293789006,
		AllergyIntoleranceCode293790002,
		AllergyIntoleranceCode293791003,
		AllergyIntoleranceCode293792005,
		AllergyIntoleranceCode293793000,
		AllergyIntoleranceCode293794006,
		AllergyIntoleranceCode293795007,
		AllergyIntoleranceCode293796008,
		AllergyIntoleranceCode293798009,
		AllergyIntoleranceCode293799001,
		AllergyIntoleranceCode293802005,
		AllergyIntoleranceCode293803000,
		AllergyIntoleranceCode293804006,
		AllergyIntoleranceCode293805007,
		AllergyIntoleranceCode293806008,
		AllergyIntoleranceCode293808009,
		AllergyIntoleranceCode293811005,
		AllergyIntoleranceCode293812003,
		AllergyIntoleranceCode293813008,
		AllergyIntoleranceCode293814002,
		AllergyIntoleranceCode293815001,
		AllergyIntoleranceCode293818004,
		AllergyIntoleranceCode293819007,
		AllergyIntoleranceCode293822009,
		AllergyIntoleranceCode293823004,
		AllergyIntoleranceCode293824005,
		AllergyIntoleranceCode293825006,
		AllergyIntoleranceCode293826007,
		AllergyIntoleranceCode293827003,
		AllergyIntoleranceCode293828008,
		AllergyIntoleranceCode293829000,
		AllergyIntoleranceCode293830005,
		AllergyIntoleranceCode293831009,
		AllergyIntoleranceCode293832002,
		AllergyIntoleranceCode293833007,
		AllergyIntoleranceCode293834001,
		AllergyIntoleranceCode293835000,
		AllergyIntoleranceCode293836004,
		AllergyIntoleranceCode293837008,
		AllergyIntoleranceCode293838003,
		AllergyIntoleranceCode293839006,
		AllergyIntoleranceCode293840008,
		AllergyIntoleranceCode293842000,
		AllergyIntoleranceCode293843005,
		AllergyIntoleranceCode293844004,
		AllergyIntoleranceCode293845003,
		AllergyIntoleranceCode293847006,
		AllergyIntoleranceCode293848001,
		AllergyIntoleranceCode293849009,
		AllergyIntoleranceCode293850009,
		AllergyIntoleranceCode293851008,
		AllergyIntoleranceCode293853006,
		AllergyIntoleranceCode293854000,
		AllergyIntoleranceCode293855004,
		AllergyIntoleranceCode293856003,
		AllergyIntoleranceCode293858002,
		AllergyIntoleranceCode293859005,
		AllergyIntoleranceCode293860000,
		AllergyIntoleranceCode293861001,
		AllergyIntoleranceCode293864009,
		AllergyIntoleranceCode293865005,
		AllergyIntoleranceCode293866006,
		AllergyIntoleranceCode293867002,
		AllergyIntoleranceCode293868007,
		AllergyIntoleranceCode293869004,
		AllergyIntoleranceCode293870003,
		AllergyIntoleranceCode293871004,
		AllergyIntoleranceCode293874007,
		AllergyIntoleranceCode293875008,
		AllergyIntoleranceCode293876009,
		AllergyIntoleranceCode293877000,
		AllergyIntoleranceCode293878005,
		AllergyIntoleranceCode293880004,
		AllergyIntoleranceCode293881000,
		AllergyIntoleranceCode293882007,
		AllergyIntoleranceCode293884008,
		AllergyIntoleranceCode293886005,
		AllergyIntoleranceCode293887001,
		AllergyIntoleranceCode293888006,
		AllergyIntoleranceCode293889003,
		AllergyIntoleranceCode293890007,
		AllergyIntoleranceCode293891006,
		AllergyIntoleranceCode293892004,
		AllergyIntoleranceCode293893009,
		AllergyIntoleranceCode293894003,
		AllergyIntoleranceCode293895002,
		AllergyIntoleranceCode293897005,
		AllergyIntoleranceCode293898000,
		AllergyIntoleranceCode293899008,
		AllergyIntoleranceCode293900003,
		AllergyIntoleranceCode293901004,
		AllergyIntoleranceCode293902006,
		AllergyIntoleranceCode293903001,
		AllergyIntoleranceCode293904007,
		AllergyIntoleranceCode293906009,
		AllergyIntoleranceCode293908005,
		AllergyIntoleranceCode293909002,
		AllergyIntoleranceCode293911006,
		AllergyIntoleranceCode293912004,
		AllergyIntoleranceCode293914003,
		AllergyIntoleranceCode293915002,
		AllergyIntoleranceCode293916001,
		AllergyIntoleranceCode293917005,
		AllergyIntoleranceCode293918000,
		AllergyIntoleranceCode293919008,
		AllergyIntoleranceCode293920002,
		AllergyIntoleranceCode293921003,
		AllergyIntoleranceCode293923000,
		AllergyIntoleranceCode293924006,
		AllergyIntoleranceCode293925007,
		AllergyIntoleranceCode293926008,
		AllergyIntoleranceCode293927004,
		AllergyIntoleranceCode293928009,
		AllergyIntoleranceCode293929001,
		AllergyIntoleranceCode293930006,
		AllergyIntoleranceCode293933008,
		AllergyIntoleranceCode293934002,
		AllergyIntoleranceCode293935001,
		AllergyIntoleranceCode293936000,
		AllergyIntoleranceCode293937009,
		AllergyIntoleranceCode293938004,
		AllergyIntoleranceCode293939007,
		AllergyIntoleranceCode293940009,
		AllergyIntoleranceCode293941008,
		AllergyIntoleranceCode293942001,
		AllergyIntoleranceCode293943006,
		AllergyIntoleranceCode293946003,
		AllergyIntoleranceCode293948002,
		AllergyIntoleranceCode293949005,
		AllergyIntoleranceCode293950005,
		AllergyIntoleranceCode293952002,
		AllergyIntoleranceCode293954001,
		AllergyIntoleranceCode293955000,
		AllergyIntoleranceCode293956004,
		AllergyIntoleranceCode293957008,
		AllergyIntoleranceCode293958003,
		AllergyIntoleranceCode293960001,
		AllergyIntoleranceCode293962009,
		AllergyIntoleranceCode293963004,
		AllergyIntoleranceCode293964005,
		AllergyIntoleranceCode293965006,
		AllergyIntoleranceCode293966007,
		AllergyIntoleranceCode293967003,
		AllergyIntoleranceCode293968008,
		AllergyIntoleranceCode293969000,
		AllergyIntoleranceCode293970004,
		AllergyIntoleranceCode293972007,
		AllergyIntoleranceCode293973002,
		AllergyIntoleranceCode293974008,
		AllergyIntoleranceCode293975009,
		AllergyIntoleranceCode293976005,
		AllergyIntoleranceCode293977001,
		AllergyIntoleranceCode293978006,
		AllergyIntoleranceCode293979003,
		AllergyIntoleranceCode293980000,
		AllergyIntoleranceCode293981001,
		AllergyIntoleranceCode293982008,
		AllergyIntoleranceCode293983003,
		AllergyIntoleranceCode293984009,
		AllergyIntoleranceCode293985005,
		AllergyIntoleranceCode293986006,
		AllergyIntoleranceCode293987002,
		AllergyIntoleranceCode293988007,
		AllergyIntoleranceCode293989004,
		AllergyIntoleranceCode293990008,
		AllergyIntoleranceCode293991007,
		AllergyIntoleranceCode293992000,
		AllergyIntoleranceCode293993005,
		AllergyIntoleranceCode293994004,
		AllergyIntoleranceCode293995003,
		AllergyIntoleranceCode293996002,
		AllergyIntoleranceCode293997006,
		AllergyIntoleranceCode293998001,
		AllergyIntoleranceCode293999009,
		AllergyIntoleranceCode294000006,
		AllergyIntoleranceCode294001005,
		AllergyIntoleranceCode294002003,
		AllergyIntoleranceCode294003008,
		AllergyIntoleranceCode294004002,
		AllergyIntoleranceCode294005001,
		AllergyIntoleranceCode294007009,
		AllergyIntoleranceCode294009007,
		AllergyIntoleranceCode294011003,
		AllergyIntoleranceCode294012005,
		AllergyIntoleranceCode294013000,
		AllergyIntoleranceCode294014006,
		AllergyIntoleranceCode294015007,
		AllergyIntoleranceCode294016008,
		AllergyIntoleranceCode294017004,
		AllergyIntoleranceCode294018009,
		AllergyIntoleranceCode294019001,
		AllergyIntoleranceCode294023009,
		AllergyIntoleranceCode294024003,
		AllergyIntoleranceCode294026001,
		AllergyIntoleranceCode294027005,
		AllergyIntoleranceCode294028000,
		AllergyIntoleranceCode294029008,
		AllergyIntoleranceCode294030003,
		AllergyIntoleranceCode294031004,
		AllergyIntoleranceCode294033001,
		AllergyIntoleranceCode294035008,
		AllergyIntoleranceCode294036009,
		AllergyIntoleranceCode294037000,
		AllergyIntoleranceCode294038005,
		AllergyIntoleranceCode294039002,
		AllergyIntoleranceCode294040000,
		AllergyIntoleranceCode294041001,
		AllergyIntoleranceCode294042008,
		AllergyIntoleranceCode294043003,
		AllergyIntoleranceCode294044009,
		AllergyIntoleranceCode294045005,
		AllergyIntoleranceCode294047002,
		AllergyIntoleranceCode294048007,
		AllergyIntoleranceCode294050004,
		AllergyIntoleranceCode294055009,
		AllergyIntoleranceCode294057001,
		AllergyIntoleranceCode294058006,
		AllergyIntoleranceCode294059003,
		AllergyIntoleranceCode294060008,
		AllergyIntoleranceCode294061007,
		AllergyIntoleranceCode294062000,
		AllergyIntoleranceCode294063005,
		AllergyIntoleranceCode294064004,
		AllergyIntoleranceCode294067006,
		AllergyIntoleranceCode294069009,
		AllergyIntoleranceCode294073007,
		AllergyIntoleranceCode294074001,
		AllergyIntoleranceCode294076004,
		AllergyIntoleranceCode294077008,
		AllergyIntoleranceCode294078003,
		AllergyIntoleranceCode294079006,
		AllergyIntoleranceCode294080009,
		AllergyIntoleranceCode294081008,
		AllergyIntoleranceCode294082001,
		AllergyIntoleranceCode294083006,
		AllergyIntoleranceCode294084000,
		AllergyIntoleranceCode294087007,
		AllergyIntoleranceCode294088002,
		AllergyIntoleranceCode294089005,
		AllergyIntoleranceCode294091002,
		AllergyIntoleranceCode294093004,
		AllergyIntoleranceCode294095006,
		AllergyIntoleranceCode294096007,
		AllergyIntoleranceCode294097003,
		AllergyIntoleranceCode294099000,
		AllergyIntoleranceCode294100008,
		AllergyIntoleranceCode294101007,
		AllergyIntoleranceCode294103005,
		AllergyIntoleranceCode294105003,
		AllergyIntoleranceCode294106002,
		AllergyIntoleranceCode294109009,
		AllergyIntoleranceCode294111000,
		AllergyIntoleranceCode294112007,
		AllergyIntoleranceCode294113002,
		AllergyIntoleranceCode294114008,
		AllergyIntoleranceCode294115009,
		AllergyIntoleranceCode294116005,
		AllergyIntoleranceCode294118006,
		AllergyIntoleranceCode294119003,
		AllergyIntoleranceCode294120009,
		AllergyIntoleranceCode294121008,
		AllergyIntoleranceCode294122001,
		AllergyIntoleranceCode294123006,
		AllergyIntoleranceCode294125004,
		AllergyIntoleranceCode294126003,
		AllergyIntoleranceCode294127007,
		AllergyIntoleranceCode294128002,
		AllergyIntoleranceCode294129005,
		AllergyIntoleranceCode294130000,
		AllergyIntoleranceCode294131001,
		AllergyIntoleranceCode294132008,
		AllergyIntoleranceCode294133003,
		AllergyIntoleranceCode294134009,
		AllergyIntoleranceCode294135005,
		AllergyIntoleranceCode294136006,
		AllergyIntoleranceCode294137002,
		AllergyIntoleranceCode294138007,
		AllergyIntoleranceCode294139004,
		AllergyIntoleranceCode294140002,
		AllergyIntoleranceCode294142005,
		AllergyIntoleranceCode294144006,
		AllergyIntoleranceCode294145007,
		AllergyIntoleranceCode294148009,
		AllergyIntoleranceCode294152009,
		AllergyIntoleranceCode294153004,
		AllergyIntoleranceCode294157003,
		AllergyIntoleranceCode294158008,
		AllergyIntoleranceCode294160005,
		AllergyIntoleranceCode294168003,
		AllergyIntoleranceCode294169006,
		AllergyIntoleranceCode294172004,
		AllergyIntoleranceCode294173009,
		AllergyIntoleranceCode294177005,
		AllergyIntoleranceCode294178000,
		AllergyIntoleranceCode294180006,
		AllergyIntoleranceCode294182003,
		AllergyIntoleranceCode294183008,
		AllergyIntoleranceCode294189007,
		AllergyIntoleranceCode294190003,
		AllergyIntoleranceCode294191004,
		AllergyIntoleranceCode294200004,
		AllergyIntoleranceCode294203002,
		AllergyIntoleranceCode294204008,
		AllergyIntoleranceCode294206005,
		AllergyIntoleranceCode294207001,
		AllergyIntoleranceCode294208006,
		AllergyIntoleranceCode294209003,
		AllergyIntoleranceCode294210008,
		AllergyIntoleranceCode294211007,
		AllergyIntoleranceCode294214004,
		AllergyIntoleranceCode294215003,
		AllergyIntoleranceCode294217006,
		AllergyIntoleranceCode294218001,
		AllergyIntoleranceCode294219009,
		AllergyIntoleranceCode294220003,
		AllergyIntoleranceCode294224007,
		AllergyIntoleranceCode294225008,
		AllergyIntoleranceCode294226009,
		AllergyIntoleranceCode294227000,
		AllergyIntoleranceCode294228005,
		AllergyIntoleranceCode294229002,
		AllergyIntoleranceCode294230007,
		AllergyIntoleranceCode294231006,
		AllergyIntoleranceCode294232004,
		AllergyIntoleranceCode294233009,
		AllergyIntoleranceCode294234003,
		AllergyIntoleranceCode294235002,
		AllergyIntoleranceCode294236001,
		AllergyIntoleranceCode294237005,
		AllergyIntoleranceCode294238000,
		AllergyIntoleranceCode294239008,
		AllergyIntoleranceCode294240005,
		AllergyIntoleranceCode294242002,
		AllergyIntoleranceCode294243007,
		AllergyIntoleranceCode294245000,
		AllergyIntoleranceCode294246004,
		AllergyIntoleranceCode294247008,
		AllergyIntoleranceCode294248003,
		AllergyIntoleranceCode294249006,
		AllergyIntoleranceCode294250006,
		AllergyIntoleranceCode294252003,
		AllergyIntoleranceCode294253008,
		AllergyIntoleranceCode294254002,
		AllergyIntoleranceCode294255001,
		AllergyIntoleranceCode294256000,
		AllergyIntoleranceCode294257009,
		AllergyIntoleranceCode294258004,
		AllergyIntoleranceCode294259007,
		AllergyIntoleranceCode294260002,
		AllergyIntoleranceCode294261003,
		AllergyIntoleranceCode294264006,
		AllergyIntoleranceCode294265007,
		AllergyIntoleranceCode294266008,
		AllergyIntoleranceCode294268009,
		AllergyIntoleranceCode294269001,
		AllergyIntoleranceCode294270000,
		AllergyIntoleranceCode294271001,
		AllergyIntoleranceCode294273003,
		AllergyIntoleranceCode294275005,
		AllergyIntoleranceCode294276006,
		AllergyIntoleranceCode294277002,
		AllergyIntoleranceCode294278007,
		AllergyIntoleranceCode294280001,
		AllergyIntoleranceCode294281002,
		AllergyIntoleranceCode294283004,
		AllergyIntoleranceCode294284005,
		AllergyIntoleranceCode294285006,
		AllergyIntoleranceCode294290009,
		AllergyIntoleranceCode294291008,
		AllergyIntoleranceCode294298002,
		AllergyIntoleranceCode294299005,
		AllergyIntoleranceCode294306008,
		AllergyIntoleranceCode294316000,
		AllergyIntoleranceCode294317009,
		AllergyIntoleranceCode294318004,
		AllergyIntoleranceCode294320001,
		AllergyIntoleranceCode294324005,
		AllergyIntoleranceCode294327003,
		AllergyIntoleranceCode294328008,
		AllergyIntoleranceCode294329000,
		AllergyIntoleranceCode294330005,
		AllergyIntoleranceCode294332002,
		AllergyIntoleranceCode294333007,
		AllergyIntoleranceCode294335000,
		AllergyIntoleranceCode294337008,
		AllergyIntoleranceCode294339006,
		AllergyIntoleranceCode294341007,
		AllergyIntoleranceCode294342000,
		AllergyIntoleranceCode294343005,
		AllergyIntoleranceCode294344004,
		AllergyIntoleranceCode294346002,
		AllergyIntoleranceCode294348001,
		AllergyIntoleranceCode294349009,
		AllergyIntoleranceCode294350009,
		AllergyIntoleranceCode294351008,
		AllergyIntoleranceCode294354000,
		AllergyIntoleranceCode294356003,
		AllergyIntoleranceCode294357007,
		AllergyIntoleranceCode294358002,
		AllergyIntoleranceCode294359005,
		AllergyIntoleranceCode294360000,
		AllergyIntoleranceCode294361001,
		AllergyIntoleranceCode294362008,
		AllergyIntoleranceCode294363003,
		AllergyIntoleranceCode294365005,
		AllergyIntoleranceCode294366006,
		AllergyIntoleranceCode294368007,
		AllergyIntoleranceCode294369004,
		AllergyIntoleranceCode294370003,
		AllergyIntoleranceCode294371004,
		AllergyIntoleranceCode294372006,
		AllergyIntoleranceCode294373001,
		AllergyIntoleranceCode294374007,
		AllergyIntoleranceCode294375008,
		AllergyIntoleranceCode294376009,
		AllergyIntoleranceCode294377000,
		AllergyIntoleranceCode294378005,
		AllergyIntoleranceCode294379002,
		AllergyIntoleranceCode294380004,
		AllergyIntoleranceCode294381000,
		AllergyIntoleranceCode294382007,
		AllergyIntoleranceCode294383002,
		AllergyIntoleranceCode294384008,
		AllergyIntoleranceCode294385009,
		AllergyIntoleranceCode294388006,
		AllergyIntoleranceCode294390007,
		AllergyIntoleranceCode294391006,
		AllergyIntoleranceCode294392004,
		AllergyIntoleranceCode294393009,
		AllergyIntoleranceCode294394003,
		AllergyIntoleranceCode294396001,
		AllergyIntoleranceCode294398000,
		AllergyIntoleranceCode294399008,
		AllergyIntoleranceCode294400001,
		AllergyIntoleranceCode294404005,
		AllergyIntoleranceCode294405006,
		AllergyIntoleranceCode294406007,
		AllergyIntoleranceCode294407003,
		AllergyIntoleranceCode294408008,
		AllergyIntoleranceCode294413007,
		AllergyIntoleranceCode294415000,
		AllergyIntoleranceCode294416004,
		AllergyIntoleranceCode294417008,
		AllergyIntoleranceCode294418003,
		AllergyIntoleranceCode294421001,
		AllergyIntoleranceCode294423003,
		AllergyIntoleranceCode294425005,
		AllergyIntoleranceCode294426006,
		AllergyIntoleranceCode294431008,
		AllergyIntoleranceCode294433006,
		AllergyIntoleranceCode294434000,
		AllergyIntoleranceCode294436003,
		AllergyIntoleranceCode294437007,
		AllergyIntoleranceCode294438002,
		AllergyIntoleranceCode294439005,
		AllergyIntoleranceCode294440007,
		AllergyIntoleranceCode294441006,
		AllergyIntoleranceCode294442004,
		AllergyIntoleranceCode294443009,
		AllergyIntoleranceCode294447005,
		AllergyIntoleranceCode294448000,
		AllergyIntoleranceCode294449008,
		AllergyIntoleranceCode294451007,
		AllergyIntoleranceCode294452000,
		AllergyIntoleranceCode294453005,
		AllergyIntoleranceCode294455003,
		AllergyIntoleranceCode294456002,
		AllergyIntoleranceCode294458001,
		AllergyIntoleranceCode294459009,
		AllergyIntoleranceCode294460004,
		AllergyIntoleranceCode294462007,
		AllergyIntoleranceCode294463002,
		AllergyIntoleranceCode294464008,
		AllergyIntoleranceCode294465009,
		AllergyIntoleranceCode294466005,
		AllergyIntoleranceCode294467001,
		AllergyIntoleranceCode294468006,
		AllergyIntoleranceCode294469003,
		AllergyIntoleranceCode294470002,
		AllergyIntoleranceCode294471003,
		AllergyIntoleranceCode294472005,
		AllergyIntoleranceCode294474006,
		AllergyIntoleranceCode294475007,
		AllergyIntoleranceCode294476008,
		AllergyIntoleranceCode294477004,
		AllergyIntoleranceCode294478009,
		AllergyIntoleranceCode294480003,
		AllergyIntoleranceCode294481004,
		AllergyIntoleranceCode294482006,
		AllergyIntoleranceCode294484007,
		AllergyIntoleranceCode294485008,
		AllergyIntoleranceCode294486009,
		AllergyIntoleranceCode294487000,
		AllergyIntoleranceCode294488005,
		AllergyIntoleranceCode294489002,
		AllergyIntoleranceCode294490006,
		AllergyIntoleranceCode294491005,
		AllergyIntoleranceCode294494002,
		AllergyIntoleranceCode294496000,
		AllergyIntoleranceCode294497009,
		AllergyIntoleranceCode294499007,
		AllergyIntoleranceCode294501004,
		AllergyIntoleranceCode294502006,
		AllergyIntoleranceCode294503001,
		AllergyIntoleranceCode294505008,
		AllergyIntoleranceCode294506009,
		AllergyIntoleranceCode294507000,
		AllergyIntoleranceCode294508005,
		AllergyIntoleranceCode294509002,
		AllergyIntoleranceCode294510007,
		AllergyIntoleranceCode294511006,
		AllergyIntoleranceCode294512004,
		AllergyIntoleranceCode294514003,
		AllergyIntoleranceCode294515002,
		AllergyIntoleranceCode294516001,
		AllergyIntoleranceCode294517005,
		AllergyIntoleranceCode294518000,
		AllergyIntoleranceCode294519008,
		AllergyIntoleranceCode294520002,
		AllergyIntoleranceCode294528009,
		AllergyIntoleranceCode294529001,
		AllergyIntoleranceCode294530006,
		AllergyIntoleranceCode294531005,
		AllergyIntoleranceCode294532003,
		AllergyIntoleranceCode294534002,
		AllergyIntoleranceCode294535001,
		AllergyIntoleranceCode294536000,
		AllergyIntoleranceCode294537009,
		AllergyIntoleranceCode294538004,
		AllergyIntoleranceCode294539007,
		AllergyIntoleranceCode294541008,
		AllergyIntoleranceCode294542001,
		AllergyIntoleranceCode294543006,
		AllergyIntoleranceCode294545004,
		AllergyIntoleranceCode294546003,
		AllergyIntoleranceCode294547007,
		AllergyIntoleranceCode294548002,
		AllergyIntoleranceCode294549005,
		AllergyIntoleranceCode294550005,
		AllergyIntoleranceCode294551009,
		AllergyIntoleranceCode294552002,
		AllergyIntoleranceCode294554001,
		AllergyIntoleranceCode294556004,
		AllergyIntoleranceCode294557008,
		AllergyIntoleranceCode294558003,
		AllergyIntoleranceCode294559006,
		AllergyIntoleranceCode294561002,
		AllergyIntoleranceCode294562009,
		AllergyIntoleranceCode294563004,
		AllergyIntoleranceCode294564005,
		AllergyIntoleranceCode294565006,
		AllergyIntoleranceCode294566007,
		AllergyIntoleranceCode294567003,
		AllergyIntoleranceCode294568008,
		AllergyIntoleranceCode294569000,
		AllergyIntoleranceCode294570004,
		AllergyIntoleranceCode294571000,
		AllergyIntoleranceCode294572007,
		AllergyIntoleranceCode294573002,
		AllergyIntoleranceCode294574008,
		AllergyIntoleranceCode294575009,
		AllergyIntoleranceCode294576005,
		AllergyIntoleranceCode294577001,
		AllergyIntoleranceCode294578006,
		AllergyIntoleranceCode294579003,
		AllergyIntoleranceCode294582008,
		AllergyIntoleranceCode294584009,
		AllergyIntoleranceCode294585005,
		AllergyIntoleranceCode294586006,
		AllergyIntoleranceCode294587002,
		AllergyIntoleranceCode294588007,
		AllergyIntoleranceCode294590008,
		AllergyIntoleranceCode294591007,
		AllergyIntoleranceCode294592000,
		AllergyIntoleranceCode294593005,
		AllergyIntoleranceCode294596002,
		AllergyIntoleranceCode294598001,
		AllergyIntoleranceCode294600007,
		AllergyIntoleranceCode294602004,
		AllergyIntoleranceCode294604003,
		AllergyIntoleranceCode294607005,
		AllergyIntoleranceCode294609008,
		AllergyIntoleranceCode294610003,
		AllergyIntoleranceCode294611004,
		AllergyIntoleranceCode294612006,
		AllergyIntoleranceCode294614007,
		AllergyIntoleranceCode294615008,
		AllergyIntoleranceCode294617000,
		AllergyIntoleranceCode294618005,
		AllergyIntoleranceCode294620008,
		AllergyIntoleranceCode294621007,
		AllergyIntoleranceCode294623005,
		AllergyIntoleranceCode294625003,
		AllergyIntoleranceCode294627006,
		AllergyIntoleranceCode294629009,
		AllergyIntoleranceCode294630004,
		AllergyIntoleranceCode294633002,
		AllergyIntoleranceCode294638006,
		AllergyIntoleranceCode294639003,
		AllergyIntoleranceCode294667007,
		AllergyIntoleranceCode294668002,
		AllergyIntoleranceCode294669005,
		AllergyIntoleranceCode294671005,
		AllergyIntoleranceCode294674002,
		AllergyIntoleranceCode294676000,
		AllergyIntoleranceCode294677009,
		AllergyIntoleranceCode294678004,
		AllergyIntoleranceCode294679007,
		AllergyIntoleranceCode294682002,
		AllergyIntoleranceCode294683007,
		AllergyIntoleranceCode294684001,
		AllergyIntoleranceCode294685000,
		AllergyIntoleranceCode294686004,
		AllergyIntoleranceCode294687008,
		AllergyIntoleranceCode294688003,
		AllergyIntoleranceCode294689006,
		AllergyIntoleranceCode294690002,
		AllergyIntoleranceCode294691003,
		AllergyIntoleranceCode294692005,
		AllergyIntoleranceCode294693000,
		AllergyIntoleranceCode294694006,
		AllergyIntoleranceCode294695007,
		AllergyIntoleranceCode294696008,
		AllergyIntoleranceCode294697004,
		AllergyIntoleranceCode294698009,
		AllergyIntoleranceCode294699001,
		AllergyIntoleranceCode294700000,
		AllergyIntoleranceCode294701001,
		AllergyIntoleranceCode294702008,
		AllergyIntoleranceCode294706006,
		AllergyIntoleranceCode294707002,
		AllergyIntoleranceCode294711008,
		AllergyIntoleranceCode294712001,
		AllergyIntoleranceCode294714000,
		AllergyIntoleranceCode294717007,
		AllergyIntoleranceCode294720004,
		AllergyIntoleranceCode294721000,
		AllergyIntoleranceCode294723002,
		AllergyIntoleranceCode294728006,
		AllergyIntoleranceCode294729003,
		AllergyIntoleranceCode294730008,
		AllergyIntoleranceCode294731007,
		AllergyIntoleranceCode294732000,
		AllergyIntoleranceCode294733005,
		AllergyIntoleranceCode294734004,
		AllergyIntoleranceCode294735003,
		AllergyIntoleranceCode294736002,
		AllergyIntoleranceCode294737006,
		AllergyIntoleranceCode294738001,
		AllergyIntoleranceCode294739009,
		AllergyIntoleranceCode294740006,
		AllergyIntoleranceCode294741005,
		AllergyIntoleranceCode294742003,
		AllergyIntoleranceCode294745001,
		AllergyIntoleranceCode294746000,
		AllergyIntoleranceCode294747009,
		AllergyIntoleranceCode294748004,
		AllergyIntoleranceCode294749007,
		AllergyIntoleranceCode294750007,
		AllergyIntoleranceCode294751006,
		AllergyIntoleranceCode294752004,
		AllergyIntoleranceCode294754003,
		AllergyIntoleranceCode294755002,
		AllergyIntoleranceCode294757005,
		AllergyIntoleranceCode294758000,
		AllergyIntoleranceCode294760003,
		AllergyIntoleranceCode294761004,
		AllergyIntoleranceCode294762006,
		AllergyIntoleranceCode294763001,
		AllergyIntoleranceCode294764007,
		AllergyIntoleranceCode294765008,
		AllergyIntoleranceCode294767000,
		AllergyIntoleranceCode294768005,
		AllergyIntoleranceCode294769002,
		AllergyIntoleranceCode294771002,
		AllergyIntoleranceCode294773004,
		AllergyIntoleranceCode294774005,
		AllergyIntoleranceCode294775006,
		AllergyIntoleranceCode294776007,
		AllergyIntoleranceCode294781003,
		AllergyIntoleranceCode294782005,
		AllergyIntoleranceCode294787004,
		AllergyIntoleranceCode294788009,
		AllergyIntoleranceCode294789001,
		AllergyIntoleranceCode294792002,
		AllergyIntoleranceCode294793007,
		AllergyIntoleranceCode294794001,
		AllergyIntoleranceCode294795000,
		AllergyIntoleranceCode294796004,
		AllergyIntoleranceCode294798003,
		AllergyIntoleranceCode294799006,
		AllergyIntoleranceCode294800005,
		AllergyIntoleranceCode294801009,
		AllergyIntoleranceCode294803007,
		AllergyIntoleranceCode294804001,
		AllergyIntoleranceCode294807008,
		AllergyIntoleranceCode294809006,
		AllergyIntoleranceCode294810001,
		AllergyIntoleranceCode294811002,
		AllergyIntoleranceCode294813004,
		AllergyIntoleranceCode294814005,
		AllergyIntoleranceCode294815006,
		AllergyIntoleranceCode294816007,
		AllergyIntoleranceCode294817003,
		AllergyIntoleranceCode294818008,
		AllergyIntoleranceCode294819000,
		AllergyIntoleranceCode294820006,
		AllergyIntoleranceCode294821005,
		AllergyIntoleranceCode294823008,
		AllergyIntoleranceCode294825001,
		AllergyIntoleranceCode294826000,
		AllergyIntoleranceCode294828004,
		AllergyIntoleranceCode294829007,
		AllergyIntoleranceCode294830002,
		AllergyIntoleranceCode294833000,
		AllergyIntoleranceCode294838009,
		AllergyIntoleranceCode294839001,
		AllergyIntoleranceCode294840004,
		AllergyIntoleranceCode294841000,
		AllergyIntoleranceCode294844008,
		AllergyIntoleranceCode294845009,
		AllergyIntoleranceCode294847001,
		AllergyIntoleranceCode294851004,
		AllergyIntoleranceCode294855008,
		AllergyIntoleranceCode294865002,
		AllergyIntoleranceCode294871008,
		AllergyIntoleranceCode294872001,
		AllergyIntoleranceCode294873006,
		AllergyIntoleranceCode294874000,
		AllergyIntoleranceCode294875004,
		AllergyIntoleranceCode294876003,
		AllergyIntoleranceCode294880008,
		AllergyIntoleranceCode294881007,
		AllergyIntoleranceCode294883005,
		AllergyIntoleranceCode294885003,
		AllergyIntoleranceCode294886002,
		AllergyIntoleranceCode294887006,
		AllergyIntoleranceCode294888001,
		AllergyIntoleranceCode294889009,
		AllergyIntoleranceCode294893003,
		AllergyIntoleranceCode294894009,
		AllergyIntoleranceCode294896006,
		AllergyIntoleranceCode294897002,
		AllergyIntoleranceCode294898007,
		AllergyIntoleranceCode294899004,
		AllergyIntoleranceCode294900009,
		AllergyIntoleranceCode294901008,
		AllergyIntoleranceCode294902001,
		AllergyIntoleranceCode294903006,
		AllergyIntoleranceCode294912008,
		AllergyIntoleranceCode294913003,
		AllergyIntoleranceCode294915005,
		AllergyIntoleranceCode294916006,
		AllergyIntoleranceCode294923007,
		AllergyIntoleranceCode294924001,
		AllergyIntoleranceCode294925000,
		AllergyIntoleranceCode294926004,
		AllergyIntoleranceCode294929006,
		AllergyIntoleranceCode294930001,
		AllergyIntoleranceCode294931002,
		AllergyIntoleranceCode294933004,
		AllergyIntoleranceCode294934005,
		AllergyIntoleranceCode294937003,
		AllergyIntoleranceCode294940003,
		AllergyIntoleranceCode294945008,
		AllergyIntoleranceCode294951003,
		AllergyIntoleranceCode294956008,
		AllergyIntoleranceCode294957004,
		AllergyIntoleranceCode294958009,
		AllergyIntoleranceCode294961005,
		AllergyIntoleranceCode294962003,
		AllergyIntoleranceCode294964002,
		AllergyIntoleranceCode294965001,
		AllergyIntoleranceCode294966000,
		AllergyIntoleranceCode294967009,
		AllergyIntoleranceCode294968004,
		AllergyIntoleranceCode294969007,
		AllergyIntoleranceCode294970008,
		AllergyIntoleranceCode294971007,
		AllergyIntoleranceCode294972000,
		AllergyIntoleranceCode294973005,
		AllergyIntoleranceCode294975003,
		AllergyIntoleranceCode294977006,
		AllergyIntoleranceCode294978001,
		AllergyIntoleranceCode294979009,
		AllergyIntoleranceCode294980007,
		AllergyIntoleranceCode294981006,
		AllergyIntoleranceCode294982004,
		AllergyIntoleranceCode294983009,
		AllergyIntoleranceCode294984003,
		AllergyIntoleranceCode294986001,
		AllergyIntoleranceCode294988000,
		AllergyIntoleranceCode294992007,
		AllergyIntoleranceCode294993002,
		AllergyIntoleranceCode294994008,
		AllergyIntoleranceCode294995009,
		AllergyIntoleranceCode294996005,
		AllergyIntoleranceCode294997001,
		AllergyIntoleranceCode294998006,
		AllergyIntoleranceCode295000003,
		AllergyIntoleranceCode295001004,
		AllergyIntoleranceCode295002006,
		AllergyIntoleranceCode295003001,
		AllergyIntoleranceCode295004007,
		AllergyIntoleranceCode295006009,
		AllergyIntoleranceCode295007000,
		AllergyIntoleranceCode295008005,
		AllergyIntoleranceCode295009002,
		AllergyIntoleranceCode295010007,
		AllergyIntoleranceCode295019008,
		AllergyIntoleranceCode267425008,
		AllergyIntoleranceCode29736007,
		AllergyIntoleranceCode29512005,
		AllergyIntoleranceCode52070001,
		AllergyIntoleranceCode237978005,
		AllergyIntoleranceCode340519003,
		AllergyIntoleranceCode190753003,
		AllergyIntoleranceCode413427002,
		AllergyIntoleranceCode716186003,
		AllergyIntoleranceCode409137002,
		AllergyIntoleranceCode428197003,
		AllergyIntoleranceCode428607008,
		AllergyIntoleranceCode429625007,
		AllergyIntoleranceCode716220001,
		AllergyIntoleranceCode1003774007,
	}
}
func FindAllergyIntoleranceCode(filter string) []AllergyIntoleranceCode {
	ret := make([]AllergyIntoleranceCode, 0)
	for _, at := range AllAllergyIntoleranceCode() {
		if strings.ToLower(at.String())[0:len(filter)] == strings.ToLower(filter) {
			ret = append(ret, at)
		}
	}
	return ret
}

func (cpt AllergyIntoleranceCode) ToString() {
	fmt.Println(cpt.String())
}

func (cpt AllergyIntoleranceCode) String() string {
	switch cpt {
	case AllergyIntoleranceCode105590001:
		return "Substance"
	case AllergyIntoleranceCode102002:
		return "Hemoglobin Okaloosa"
	case AllergyIntoleranceCode120006:
		return "Ornithine racemase"
	case AllergyIntoleranceCode125001:
		return "Ferrous (59-Fe) sulfate (substance)"
	case AllergyIntoleranceCode126000:
		return "Galactosyl-N-acetylglucosaminylgalactosylglucosylceramide alpha-galactosyltransferase"
	case AllergyIntoleranceCode130002:
		return "Hemoglobin Hopkins-II"
	case AllergyIntoleranceCode131003:
		return "Dolichyl-phosphate mannosyltransferase"
	case AllergyIntoleranceCode159002:
		return "Ferrocyanide salt"
	case AllergyIntoleranceCode164003:
		return "Phosphoenolpyruvate-protein phosphotransferase"
	case AllergyIntoleranceCode178002:
		return "Uridine diphosphate galactose"
	case AllergyIntoleranceCode186002:
		return "HLA-Cw9 antigen"
	case AllergyIntoleranceCode187006:
		return "Cyanocobalamin (57-Co) (substance)"
	case AllergyIntoleranceCode200001:
		return "Berberine"
	case AllergyIntoleranceCode217008:
		return "Blood group antigen IH"
	case AllergyIntoleranceCode231008:
		return "3-hydroxyisobutyrate dehydrogenase"
	case AllergyIntoleranceCode238002:
		return "Heptachlor"
	case AllergyIntoleranceCode261000:
		return "Codeine phosphate"
	case AllergyIntoleranceCode296000:
		return "Coumachlor"
	case AllergyIntoleranceCode322006:
		return "Octylphenoxy P.H. ethanol"
	case AllergyIntoleranceCode327000:
		return "Arsenic-76 (substance)"
	case AllergyIntoleranceCode329002:
		return "Antimony-127 (substance)"
	case AllergyIntoleranceCode336001:
		return "Fibrinogen Tokyo II"
	case AllergyIntoleranceCode340005:
		return "Enzyme variant"
	case AllergyIntoleranceCode363000:
		return "Fibrinogen San Juan"
	case AllergyIntoleranceCode370000:
		return "Beta>2S< glycoprotein (substance)"
	case AllergyIntoleranceCode371001:
		return "Acylcarnitine hydrolase"
	case AllergyIntoleranceCode377002:
		return "Sparteine"
	case AllergyIntoleranceCode392001:
		return "Gadolinium-151 (substance)"
	case AllergyIntoleranceCode395004:
		return "Immunoglobulin pentamer"
	case AllergyIntoleranceCode412004:
		return "Ribose-5-phosphate isomerase"
	case AllergyIntoleranceCode424006:
		return "Citramalyl-CoA lyase"
	case AllergyIntoleranceCode425007:
		return "Hemoglobin Nagoya"
	case AllergyIntoleranceCode432003:
		return "Carminic acid"
	case AllergyIntoleranceCode438004:
		return "2-hydroxyglutarate dehydrogenase (substance)"
	case AllergyIntoleranceCode462009:
		return "Urease (ATP-hydrolysing)"
	case AllergyIntoleranceCode472007:
		return "Vegetable textile fiber"
	case AllergyIntoleranceCode476005:
		return "Lymphocyte antigen CD1b"
	case AllergyIntoleranceCode498001:
		return "Nitrilase"
	case AllergyIntoleranceCode501001:
		return "Blood group antibody Sf^a^"
	case AllergyIntoleranceCode505005:
		return "Blood group antibody M'"
	case AllergyIntoleranceCode506006:
		return "3-oxosteroid delta^1^-dehydrogenase (substance)"
	case AllergyIntoleranceCode515004:
		return "Blood group antigen Giaigue"
	case AllergyIntoleranceCode519005:
		return "Free protein S"
	case AllergyIntoleranceCode521000:
		return "Mercury-197 (substance)"
	case AllergyIntoleranceCode529003:
		return "Guanosine"
	case AllergyIntoleranceCode538001:
		return "2,3-dihydroxybenzoate 3,4-dioxygenase (substance)"
	case AllergyIntoleranceCode566009:
		return "Acrosin"
	case AllergyIntoleranceCode576007:
		return "Blood group antibody Duck"
	case AllergyIntoleranceCode578008:
		return "Hemoglobin Jianghua"
	case AllergyIntoleranceCode584006:
		return "Blood group antibody Wr^b^"
	case AllergyIntoleranceCode585007:
		return "Substance P"
	case AllergyIntoleranceCode591009:
		return "2-oxoisovalerate dehydrogenase (acylating) (substance)"
	case AllergyIntoleranceCode593007:
		return "Blood group antibody Holmes"
	case AllergyIntoleranceCode594001:
		return "2-oxoglutarate synthase (substance)"
	case AllergyIntoleranceCode597008:
		return "Californium-247 (substance)"
	case AllergyIntoleranceCode604000:
		return "Plant sapogenin glycoside"
	case AllergyIntoleranceCode611001:
		return "Hippurate hydrolase"
	case AllergyIntoleranceCode620005:
		return "Trichlorophenol"
	case AllergyIntoleranceCode648005:
		return "Oil of calamus"
	case AllergyIntoleranceCode662003:
		return "Aeromonas proteolytica aminopeptidase"
	case AllergyIntoleranceCode668004:
		return "Osmium-185 (substance)"
	case AllergyIntoleranceCode683009:
		return "Mercuric acetate"
	case AllergyIntoleranceCode686001:
		return "Plastoquinol-plastocyanin reductase"
	case AllergyIntoleranceCode693002:
		return "Trichothecenes"
	case AllergyIntoleranceCode698006:
		return "Erythromycin lactobionate"
	case AllergyIntoleranceCode699003:
		return "Coal tar extract"
	case AllergyIntoleranceCode704006:
		return "Blood group antigen Rx"
	case AllergyIntoleranceCode732002:
		return "N-valeraldehyde"
	case AllergyIntoleranceCode735000:
		return "Blood group antigen Jobbins"
	case AllergyIntoleranceCode747006:
		return "Oxamniquine"
	case AllergyIntoleranceCode773001:
		return "Hemoglobin M-Iwate"
	case AllergyIntoleranceCode785009:
		return "Dextranase"
	case AllergyIntoleranceCode804003:
		return "Creosotic acid"
	case AllergyIntoleranceCode819002:
		return "Lytic antibody"
	case AllergyIntoleranceCode850000:
		return "Stizolobate synthase"
	case AllergyIntoleranceCode859004:
		return "Peptide-N^4^-(N-acetyl-b-glucosaminyl) asparagine amidase"
	case AllergyIntoleranceCode860009:
		return "Immunoglobulin, aggregated"
	case AllergyIntoleranceCode873008:
		return "Urethan"
	case AllergyIntoleranceCode876000:
		return "Blood group antigen D"
	case AllergyIntoleranceCode877009:
		return "Carboxypeptidase A"
	case AllergyIntoleranceCode889006:
		return "(acetyl-coenzyme A carboxylase) kinase (substance)"
	case AllergyIntoleranceCode896008:
		return "Ice"
	case AllergyIntoleranceCode905001:
		return "o-Dihydroxycoumarin O^7^-glucosyltransferase"
	case AllergyIntoleranceCode923009:
		return "Complement component C2"
	case AllergyIntoleranceCode925002:
		return "Sodium iodipamide"
	case AllergyIntoleranceCode963005:
		return "Pyridoxine 4-dehydrogenase"
	case AllergyIntoleranceCode974001:
		return "Adenosylmethionine decarboxylase"
	case AllergyIntoleranceCode979006:
		return "Carbamate kinase"
	case AllergyIntoleranceCode993004:
		return "Palladium compound"
	case AllergyIntoleranceCode1002007:
		return "Mannotetraose 2-alpha-N-acetylglucosaminyltransferase"
	case AllergyIntoleranceCode1010008:
		return "N-Acetylneuraminate monooxygenase"
	case AllergyIntoleranceCode1018001:
		return "Nornicotine"
	case AllergyIntoleranceCode1025008:
		return "Molybdenum-93 (substance)"
	case AllergyIntoleranceCode1047008:
		return "Guanine deaminase"
	case AllergyIntoleranceCode1050006:
		return "Melilotate 3-monooxygenase"
	case AllergyIntoleranceCode1065007:
		return "E. coli periplasmic proteinase"
	case AllergyIntoleranceCode1080001:
		return "Thallium-202 (substance)"
	case AllergyIntoleranceCode1091008:
		return "Coagulation factor inhibitor"
	case AllergyIntoleranceCode1097007:
		return "Blood group antigen M^A^"
	case AllergyIntoleranceCode1105007:
		return "Isochorismate synthase"
	case AllergyIntoleranceCode1113008:
		return "Pancreatic ribonuclease"
	case AllergyIntoleranceCode1137008:
		return "Uranium-240 (substance)"
	case AllergyIntoleranceCode1149009:
		return "Hemoglobin Barcelona"
	case AllergyIntoleranceCode1160000:
		return "Blood group antibody Lutheran"
	case AllergyIntoleranceCode1166006:
		return "Titanium"
	case AllergyIntoleranceCode1169004:
		return "Hemoglobin Gower-2"
	case AllergyIntoleranceCode1171004:
		return "Fibrinogen Kawaguchi"
	case AllergyIntoleranceCode1185009:
		return "Hemoglobin Roseau-Pointe à Pitre"
	case AllergyIntoleranceCode1189003:
		return "Hemoglobin F-M-Osaka"
	case AllergyIntoleranceCode1190007:
		return "Mephenoxalone"
	case AllergyIntoleranceCode1219001:
		return "Diethyl xanthogen disulfide"
	case AllergyIntoleranceCode1223009:
		return "Blood group antigen Marks"
	case AllergyIntoleranceCode1244009:
		return "Fibrinogen Madrid I"
	case AllergyIntoleranceCode1248007:
		return "Leucostoma neutral proteinase"
	case AllergyIntoleranceCode1269009:
		return "Amikacin sulfate"
	case AllergyIntoleranceCode1272002:
		return "Pteridine oxidase"
	case AllergyIntoleranceCode1273007:
		return "Blood group antibody Evelyn"
	case AllergyIntoleranceCode1313002:
		return "Nitrate reductase (cytochrome)"
	case AllergyIntoleranceCode1319003:
		return "Blood group antibody K18"
	case AllergyIntoleranceCode1320009:
		return "Hemoglobin Manitoba"
	case AllergyIntoleranceCode1325004:
		return "Metocurine iodide"
	case AllergyIntoleranceCode1331001:
		return "Methamidophos"
	case AllergyIntoleranceCode1336006:
		return "Deoxycortone"
	case AllergyIntoleranceCode1341003:
		return "Hemoglobin Ta-li"
	case AllergyIntoleranceCode1346008:
		return "Blue shade eosin"
	case AllergyIntoleranceCode1355006:
		return "Antihemophilic factor B Oxford 3 variant"
	case AllergyIntoleranceCode1368003:
		return "Iodine-131 (substance)"
	case AllergyIntoleranceCode1371006:
		return "Blood group antigen Big"
	case AllergyIntoleranceCode1373009:
		return "Zirconium-93 (substance)"
	case AllergyIntoleranceCode1381005:
		return "Iodine-126 (substance)"
	case AllergyIntoleranceCode1394007:
		return "Iron pentacarbonyl"
	case AllergyIntoleranceCode1396009:
		return "Actinium"
	case AllergyIntoleranceCode1405004:
		return "Blood group antibody M^e^"
	case AllergyIntoleranceCode1408002:
		return "Blood group antibody 1123K"
	case AllergyIntoleranceCode1416006:
		return "Radium compound"
	case AllergyIntoleranceCode1450002:
		return "Methylparafynol"
	case AllergyIntoleranceCode1466000:
		return "Cyclomaltodextrinase"
	case AllergyIntoleranceCode1471007:
		return "Elastin"
	case AllergyIntoleranceCode1472000:
		return "Adenosine-phosphate deaminase"
	case AllergyIntoleranceCode1476002:
		return "Codeine sulfate"
	case AllergyIntoleranceCode1477006:
		return "Hemoglobin Yatsushiro"
	case AllergyIntoleranceCode1496005:
		return "Proto-oncogene"
	case AllergyIntoleranceCode1506001:
		return "Blood group antigen Ch^a^"
	case AllergyIntoleranceCode1517000:
		return "HLA-B21 antigen"
	case AllergyIntoleranceCode1530004:
		return "6-carboxyhexanoate-CoA ligase"
	case AllergyIntoleranceCode1535009:
		return "Nitrogen fluoride"
	case AllergyIntoleranceCode1536005:
		return "Pargyline hydrochloride"
	case AllergyIntoleranceCode1540001:
		return "Tellurium radioisotope"
	case AllergyIntoleranceCode1545006:
		return "Uridine phosphorylase"
	case AllergyIntoleranceCode1557002:
		return "Talc"
	case AllergyIntoleranceCode1565004:
		return "Blood group antibody Buckalew"
	case AllergyIntoleranceCode1575001:
		return "Maltose tetrapalmitate"
	case AllergyIntoleranceCode1603001:
		return "Cobalt isotope"
	case AllergyIntoleranceCode1607000:
		return "Homoserine kinase"
	case AllergyIntoleranceCode1609002:
		return "N-octyl isosafrole sulfoxide"
	case AllergyIntoleranceCode1634002:
		return "Blood group antigen Ven"
	case AllergyIntoleranceCode1649005:
		return "Blood group antigen Sul"
	case AllergyIntoleranceCode1656004:
		return "Hemoglobin Shaare Zedek"
	case AllergyIntoleranceCode1660001:
		return "Plant seeds"
	case AllergyIntoleranceCode1668008:
		return "Ceforanide"
	case AllergyIntoleranceCode1672007:
		return "Ligase"
	case AllergyIntoleranceCode1673002:
		return "Xylenol"
	case AllergyIntoleranceCode1675009:
		return "Rubidium-86 (substance)"
	case AllergyIntoleranceCode1676005:
		return "Blood group antibody LW^ab^"
	case AllergyIntoleranceCode1681001:
		return "Blood group antibody BLe^b^"
	case AllergyIntoleranceCode1696002:
		return "12-HPETE"
	case AllergyIntoleranceCode1701009:
		return "Gold-191 (substance)"
	case AllergyIntoleranceCode1710001:
		return "Uric acid"
	case AllergyIntoleranceCode1726000:
		return "Diamond"
	case AllergyIntoleranceCode1727009:
		return "Deoxylimonate A-ring-lactonase"
	case AllergyIntoleranceCode1740004:
		return "Deoxy cytidine triphosphate"
	case AllergyIntoleranceCode1764003:
		return "Saccharopine dehydrogenase (NADP^+^,L-glutamate-forming)"
	case AllergyIntoleranceCode1768000:
		return "Sucrose phosphorylase"
	case AllergyIntoleranceCode1786002:
		return "Leucine-tRNA ligase"
	case AllergyIntoleranceCode1793003:
		return "Sodium trichloroacetate"
	case AllergyIntoleranceCode1795005:
		return "Glyodin"
	case AllergyIntoleranceCode1798007:
		return "Hemoglobin Hammersmith"
	case AllergyIntoleranceCode1799004:
		return "L-Lysine oxidase"
	case AllergyIntoleranceCode1823002:
		return "Hemoglobin Tochigi"
	case AllergyIntoleranceCode1827001:
		return "Ribonuclease T>1<"
	case AllergyIntoleranceCode1886008:
		return "Verdohemoglobin"
	case AllergyIntoleranceCode1904005:
		return "Galactoside 3-fucosyltransferase"
	case AllergyIntoleranceCode1914001:
		return "von Willebrand factor inhibitor"
	case AllergyIntoleranceCode1916004:
		return "Boroglycerin"
	case AllergyIntoleranceCode1940007:
		return "Immunoglobulin, GM>21< allotype"
	case AllergyIntoleranceCode1944003:
		return "Coagulation factor X Patient variant"
	case AllergyIntoleranceCode1956002:
		return "Buclizine hydrochloride"
	case AllergyIntoleranceCode1971003:
		return "Loxapine hydrochloride"
	case AllergyIntoleranceCode1975007:
		return "Blood group antibody Niemetz"
	case AllergyIntoleranceCode1978009:
		return "Site-specific methyltransferase (cytosine-specific)"
	case AllergyIntoleranceCode1985008:
		return "Vomitus"
	case AllergyIntoleranceCode1991005:
		return "Lignins"
	case AllergyIntoleranceCode2000001:
		return "Heavy nitrogen"
	case AllergyIntoleranceCode2006007:
		return "Inosine diphosphate"
	case AllergyIntoleranceCode2008008:
		return "Gallium-67 (substance)"
	case AllergyIntoleranceCode2009000:
		return "Cobalt carbonyl"
	case AllergyIntoleranceCode2017008:
		return "DNA topoisomerase"
	case AllergyIntoleranceCode2027002:
		return "Alternaria serine proteinase"
	case AllergyIntoleranceCode2029004:
		return "Fibrinogen Oslo II"
	case AllergyIntoleranceCode2038002:
		return "Blood group antibody Bg^b^"
	case AllergyIntoleranceCode2039005:
		return "sym-Norspermidine synthase"
	case AllergyIntoleranceCode2050008:
		return "Choloylglycine hydrolase"
	case AllergyIntoleranceCode2064008:
		return "L-Xylulokinase"
	case AllergyIntoleranceCode2082006:
		return "Lymphocyte antigen CD51"
	case AllergyIntoleranceCode2085008:
		return "Oncogene protein TCL"
	case AllergyIntoleranceCode2088005:
		return "Page blue G-90"
	case AllergyIntoleranceCode2096000:
		return "NAD^+^ ADP-ribosyltransferase"
	case AllergyIntoleranceCode2100004:
		return "Sulfonethylmethane"
	case AllergyIntoleranceCode2101000:
		return "Yeast proteinase B"
	case AllergyIntoleranceCode2125008:
		return "Betazole"
	case AllergyIntoleranceCode2130007:
		return "Cyclohexane-1,2-diol dehydrogenase"
	case AllergyIntoleranceCode2141009:
		return "Hydrogen"
	case AllergyIntoleranceCode2147008:
		return "Blood group antigen Paular"
	case AllergyIntoleranceCode2151005:
		return "Pyridoxamine-pyruvate aminotransferase"
	case AllergyIntoleranceCode2154002:
		return "Tagaturonate reductase"
	case AllergyIntoleranceCode2159007:
		return "Azorubin S"
	case AllergyIntoleranceCode2163000:
		return "Dicofol"
	case AllergyIntoleranceCode2168009:
		return "Bisphosphoglycerate mutase"
	case AllergyIntoleranceCode2179004:
		return "Malonate-semialdehyde dehydratase"
	case AllergyIntoleranceCode2189000:
		return "Hemoglobin F-Dammam"
	case AllergyIntoleranceCode2194000:
		return "Rhodium-101 (substance)"
	case AllergyIntoleranceCode2195004:
		return "Tocainide hydrochloride"
	case AllergyIntoleranceCode2201007:
		return "Bacteriopurpurin"
	case AllergyIntoleranceCode2208001:
		return "Phenylserine aldolase"
	case AllergyIntoleranceCode2212007:
		return "Fibrinogen Bethesda II"
	case AllergyIntoleranceCode2215009:
		return "Azuresin"
	case AllergyIntoleranceCode2240002:
		return "Guanidinobutyrase"
	case AllergyIntoleranceCode2249001:
		return "Gentamicin sulfate"
	case AllergyIntoleranceCode2254005:
		return "Orotic acid"
	case AllergyIntoleranceCode2260005:
		return "HLA-DRw18 antigen"
	case AllergyIntoleranceCode2262002:
		return "Cellulose polysulfatase"
	case AllergyIntoleranceCode2264001:
		return "Selenium isotope"
	case AllergyIntoleranceCode2309006:
		return "Gold"
	case AllergyIntoleranceCode2311002:
		return "Prostacyclin synthase"
	case AllergyIntoleranceCode2329007:
		return "Blood group antibody Vel"
	case AllergyIntoleranceCode2331003:
		return "Carbohydrate"
	case AllergyIntoleranceCode2338009:
		return "Plant roots"
	case AllergyIntoleranceCode2343002:
		return "Guthion"
	case AllergyIntoleranceCode2346005:
		return "Vascormone"
	case AllergyIntoleranceCode2354007:
		return "3'-nucleotidase"
	case AllergyIntoleranceCode2369008:
		return "Indole-3-acetate beta-glucosyltransferase"
	case AllergyIntoleranceCode2370009:
		return "UDP-N-acetylmuramate-alanine ligase"
	case AllergyIntoleranceCode2376003:
		return "Mercury compound"
	case AllergyIntoleranceCode2384004:
		return "Uranium-230 (substance)"
	case AllergyIntoleranceCode2404002:
		return "Blood group antibody St^a^"
	case AllergyIntoleranceCode2405001:
		return "Oxetanone"
	case AllergyIntoleranceCode2414006:
		return "Prolactin receptor"
	case AllergyIntoleranceCode2430003:
		return "Silicon radioisotope"
	case AllergyIntoleranceCode2431004:
		return "Blood group antibody Friedberg"
	case AllergyIntoleranceCode2441001:
		return "Mercury radioisotope"
	case AllergyIntoleranceCode2444009:
		return "HLA-Dw25 antigen"
	case AllergyIntoleranceCode2450004:
		return "Mannosamine"
	case AllergyIntoleranceCode2462000:
		return "Glucose dehydrogenase (NADP^+^)"
	case AllergyIntoleranceCode2466002:
		return "Chloride peroxidase"
	case AllergyIntoleranceCode2500009:
		return "Lymphocyte antigen CDw41b"
	case AllergyIntoleranceCode2509005:
		return "D-glutamic acid oxidase"
	case AllergyIntoleranceCode2522002:
		return "Extravascular blood"
	case AllergyIntoleranceCode2529006:
		return "Hemoglobin Wood"
	case AllergyIntoleranceCode2537003:
		return "Antituberculosis agent"
	case AllergyIntoleranceCode2568004:
		return "Blood group antigen McAuley"
	case AllergyIntoleranceCode2573005:
		return "Immunoglobulin, GM>13< allotype"
	case AllergyIntoleranceCode2575003:
		return "Zinc alpha>2< glycoprotein"
	case AllergyIntoleranceCode2595009:
		return "Tellurium-119m (substance)"
	case AllergyIntoleranceCode2597001:
		return "Alpha-1 globulin"
	case AllergyIntoleranceCode2611008:
		return "Blood group antibody La Fave"
	case AllergyIntoleranceCode2637006:
		return "Indium isotope"
	case AllergyIntoleranceCode2648004:
		return "Bile vomitus"
	case AllergyIntoleranceCode2649007:
		return "Azo dye"
	case AllergyIntoleranceCode2660003:
		return "Sodium dehydrocholate"
	case AllergyIntoleranceCode2671002:
		return "Dehydropantoate hydroxymethyltransferase"
	case AllergyIntoleranceCode2674005:
		return "Cesium-128 (substance)"
	case AllergyIntoleranceCode2676007:
		return "C3(H20)"
	case AllergyIntoleranceCode2678008:
		return "Hemoglobin New Mexico"
	case AllergyIntoleranceCode2680002:
		return "Anti-factor XIII"
	case AllergyIntoleranceCode2698003:
		return "Natural gas"
	case AllergyIntoleranceCode2705002:
		return "Arsenic-72 (substance)"
	case AllergyIntoleranceCode2706001:
		return "Blood group antigen Vennera"
	case AllergyIntoleranceCode2719002:
		return "Tartrate dehydratase"
	case AllergyIntoleranceCode2721007:
		return "Blood group antigen McC^f^"
	case AllergyIntoleranceCode2728001:
		return "Antigen in Lewis (Le) blood group system"
	case AllergyIntoleranceCode2753003:
		return "Blood group antibody M>1<"
	case AllergyIntoleranceCode2754009:
		return "Hemoglobin F-Kennestone"
	case AllergyIntoleranceCode2765004:
		return "Blood group antigen Sc3"
	case AllergyIntoleranceCode2778004:
		return "Pleural fluid"
	case AllergyIntoleranceCode2796008:
		return "Methantheline (substance)"
	case AllergyIntoleranceCode2799001:
		return "Methylbenzethonium chloride"
	case AllergyIntoleranceCode2823004:
		return "Hemoglobin Bristol"
	case AllergyIntoleranceCode2832002:
		return "Molybdenum compound"
	case AllergyIntoleranceCode2846002:
		return "Hemoglobin Saitama"
	case AllergyIntoleranceCode2869004:
		return "Ethanoic acid"
	case AllergyIntoleranceCode2878005:
		return "Isonipecaine hydrochloride"
	case AllergyIntoleranceCode2880004:
		return "Calcium sulfate"
	case AllergyIntoleranceCode2883002:
		return "Exopolygalacturonate lyase"
	case AllergyIntoleranceCode2913009:
		return "Immunoglobulin E, H chain"
	case AllergyIntoleranceCode2916001:
		return "Neon-22 (substance)"
	case AllergyIntoleranceCode2925007:
		return "Fluorometholone"
	case AllergyIntoleranceCode2927004:
		return "Rescinnamine"
	case AllergyIntoleranceCode2938004:
		return "Pyrazole"
	case AllergyIntoleranceCode2942001:
		return "Carbon (14-C) xylose (substance)"
	case AllergyIntoleranceCode2950005:
		return "Hemoglobin L-Persian Gulf"
	case AllergyIntoleranceCode2958003:
		return "Zinc caprylate"
	case AllergyIntoleranceCode2964005:
		return "Dimethoxyamphetamine"
	case AllergyIntoleranceCode2974008:
		return "Trichophyton schoenleinii collagenase"
	case AllergyIntoleranceCode2988007:
		return "HLA-Aw antigen"
	case AllergyIntoleranceCode2991007:
		return "Mecamylamine hydrochloride"
	case AllergyIntoleranceCode2995003:
		return "Arecoline"
	case AllergyIntoleranceCode3027009:
		return "Barium-133 (substance)"
	case AllergyIntoleranceCode3031003:
		return "Dihydroxyaluminum sodium carbonate"
	case AllergyIntoleranceCode3040004:
		return "Technetium (99m-Tc) disofenin (substance)"
	case AllergyIntoleranceCode3045009:
		return "Nitrochlorobenzene"
	case AllergyIntoleranceCode3052006:
		return "Ornithine-oxo-acid aminotransferase"
	case AllergyIntoleranceCode3066001:
		return "Triiodothyroacetic acid"
	case AllergyIntoleranceCode3070009:
		return "Aspartate-ammonia ligase"
	case AllergyIntoleranceCode3087006:
		return "Oil of male fern"
	case AllergyIntoleranceCode3107005:
		return "Hemoglobin Shuangfeng"
	case AllergyIntoleranceCode3108000:
		return "Aspergillus deoxyribonuclease K>1<"
	case AllergyIntoleranceCode3131000:
		return "Blood group antigen Middel"
	case AllergyIntoleranceCode3136005:
		return "Cefoperazone sodium"
	case AllergyIntoleranceCode3142009:
		return "Azacyclonol"
	case AllergyIntoleranceCode3145006:
		return "Penicillic acid"
	case AllergyIntoleranceCode3150000:
		return "Sialate O-acetylesterase"
	case AllergyIntoleranceCode3151001:
		return "Left upper lobe mucus"
	case AllergyIntoleranceCode3155005:
		return "3-phosphoglyceroyl-phosphate-polyphosphate phosphotransferase"
	case AllergyIntoleranceCode3161008:
		return "3-methyl histidine"
	case AllergyIntoleranceCode3167007:
		return "Hard coal"
	case AllergyIntoleranceCode3187008:
		return "Blood group antigen Nielsen"
	case AllergyIntoleranceCode3193000:
		return "Alpha-1,4-glucan-protein synthase (uridine diphosphate-forming)"
	case AllergyIntoleranceCode3197004:
		return "Inosine monophosphate"
	case AllergyIntoleranceCode3209002:
		return "Pancuronium sodium"
	case AllergyIntoleranceCode3212004:
		return "Manganese sulfate"
	case AllergyIntoleranceCode3225007:
		return "Fibrinogen Seattle I"
	case AllergyIntoleranceCode3232003:
		return "O-benzyl-parachlorophenol (substance)"
	case AllergyIntoleranceCode3271000:
		return "Hemoglobin Southampton"
	case AllergyIntoleranceCode3273002:
		return "Tyrosine-ester sulfotransferase"
	case AllergyIntoleranceCode3300001:
		return "Euphorbain"
	case AllergyIntoleranceCode3318003:
		return "Vaginal secretion"
	case AllergyIntoleranceCode3325005:
		return "Lipopolysaccharide"
	case AllergyIntoleranceCode3339005:
		return "(R)-20-Hydroxysteroid dehydrogenase"
	case AllergyIntoleranceCode3340007:
		return "Alpha-amylase (substance)"
	case AllergyIntoleranceCode3342004:
		return "Copper isotope"
	case AllergyIntoleranceCode3346001:
		return "Hemoglobin Brest"
	case AllergyIntoleranceCode3378009:
		return "Imipramine hydrochloride"
	case AllergyIntoleranceCode3379001:
		return "Thimerosal"
	case AllergyIntoleranceCode3392003:
		return "Aldehyde dehydrogenase (acceptor)"
	case AllergyIntoleranceCode3405005:
		return "2-hydroxy-3-oxoadipate synthase"
	case AllergyIntoleranceCode3411008:
		return "bis-(Dimethylthiocarbamyl) disulfide"
	case AllergyIntoleranceCode3437006:
		return "Hydroxymethylglutaryl-CoA hydrolase"
	case AllergyIntoleranceCode3440006:
		return "Biotin carboxylase"
	case AllergyIntoleranceCode3455002:
		return "Discontinued pesticide"
	case AllergyIntoleranceCode3463001:
		return "L-amino-acid dehydrogenase (substance)"
	case AllergyIntoleranceCode3465008:
		return "DNA topoisomerase (ATP-hydrolysing)"
	case AllergyIntoleranceCode3466009:
		return "Dimethylamine"
	case AllergyIntoleranceCode3492002:
		return "Galactinol-sucrose galactosyltransferase"
	case AllergyIntoleranceCode3493007:
		return "Smegma clitoridis"
	case AllergyIntoleranceCode3495000:
		return "Cystyl-aminopeptidase"
	case AllergyIntoleranceCode3501003:
		return "Isoxsuprine hydrochloride"
	case AllergyIntoleranceCode3523004:
		return "Hemoglobin Q-India"
	case AllergyIntoleranceCode3532002:
		return "Laryngeal mucus"
	case AllergyIntoleranceCode3555004:
		return "Blood group antigen Morrison"
	case AllergyIntoleranceCode3579002:
		return "Cesium-129 (substance)"
	case AllergyIntoleranceCode3581000:
		return "Glucose-6-phosphatase"
	case AllergyIntoleranceCode3587001:
		return "Malate dehydrogenase (decarboxylating)"
	case AllergyIntoleranceCode3588006:
		return "Complement enzyme"
	case AllergyIntoleranceCode3597005:
		return "Acebutolol hydrochloride"
	case AllergyIntoleranceCode3602003:
		return "Warm antibody"
	case AllergyIntoleranceCode3610002:
		return "Epoxide hydrolase"
	case AllergyIntoleranceCode3617004:
		return "Selenium-79 (substance)"
	case AllergyIntoleranceCode3648007:
		return "Glucocorticoid receptor"
	case AllergyIntoleranceCode3655009:
		return "Hemoglobin Constant Springs"
	case AllergyIntoleranceCode3672002:
		return "Fibrinogen Caracas"
	case AllergyIntoleranceCode3684000:
		return "Phenylacetic acid"
	case AllergyIntoleranceCode3689005:
		return "Hemoglobin Mizushi"
	case AllergyIntoleranceCode3692009:
		return "Sodium sulfite"
	case AllergyIntoleranceCode3693004:
		return "Fibrinogen Dusard"
	case AllergyIntoleranceCode3702007:
		return "CDPglycerol glycerophosphotransferase"
	case AllergyIntoleranceCode3710008:
		return "Prostaglandin synthase"
	case AllergyIntoleranceCode3718001:
		return "Cow's milk"
	case AllergyIntoleranceCode3726009:
		return "Valine-tRNA ligase"
	case AllergyIntoleranceCode3727000:
		return "Hemoglobin F-Port Royal"
	case AllergyIntoleranceCode3730007:
		return "Blood group antigen Tr^a^"
	case AllergyIntoleranceCode3737005:
		return "Nitrate reductase (NADH)"
	case AllergyIntoleranceCode3742002:
		return "Extracellular crystal"
	case AllergyIntoleranceCode3757009:
		return "Gossypol"
	case AllergyIntoleranceCode3771001:
		return "Neuromelanin"
	case AllergyIntoleranceCode3775005:
		return "Choline dehydrogenase"
	case AllergyIntoleranceCode3776006:
		return "Xanthine dehydrogenase"
	case AllergyIntoleranceCode3792001:
		return "Arachidonic acid"
	case AllergyIntoleranceCode3800009:
		return "Acetate kinase"
	case AllergyIntoleranceCode3807007:
		return "Blood group antigen c"
	case AllergyIntoleranceCode3811001:
		return "Magnesium-protoporphyrin methyltransferase"
	case AllergyIntoleranceCode3812008:
		return "Beryllium isotope"
	case AllergyIntoleranceCode3816006:
		return "Vanadium isotope"
	case AllergyIntoleranceCode3823007:
		return "Prochlorperazine edisylate"
	case AllergyIntoleranceCode3829006:
		return "Iron"
	case AllergyIntoleranceCode3834005:
		return "CMP-N-acetylneuraminate-(alpha-N-acetyl-neuraminyl-2,3-beta-galactosyl-1,3)-N-acetyl-galactosaminide alpha-2,6-sialyltransferase"
	case AllergyIntoleranceCode3836007:
		return "Glutaminase"
	case AllergyIntoleranceCode3844007:
		return "Protoaphin-aglucone dehydratase (cyclizing)"
	case AllergyIntoleranceCode3848005:
		return "Nitrotoluene"
	case AllergyIntoleranceCode3849002:
		return "Carbon black"
	case AllergyIntoleranceCode3854006:
		return "Bis-chloro methyl ether (substance)"
	case AllergyIntoleranceCode3874004:
		return "Hydrocodone bitartrate"
	case AllergyIntoleranceCode3892007:
		return "Thymidine"
	case AllergyIntoleranceCode3896005:
		return "p-Hydroxybenzoate ester"
	case AllergyIntoleranceCode3897001:
		return "Blood group antigen 'N'"
	case AllergyIntoleranceCode3906002:
		return "Rectified birch tar oil (substance)"
	case AllergyIntoleranceCode3920009:
		return "Hemoglobin Atago"
	case AllergyIntoleranceCode3930000:
		return "Manufactured gas"
	case AllergyIntoleranceCode3932008:
		return "Copper-64 (substance)"
	case AllergyIntoleranceCode3941003:
		return "Metronidazole hydrochloride"
	case AllergyIntoleranceCode3945007:
		return "Tin isotope"
	case AllergyIntoleranceCode3958008:
		return "Californium-245 (substance)"
	case AllergyIntoleranceCode3961009:
		return "Blood group antigen Ritherford"
	case AllergyIntoleranceCode3976001:
		return "Blood group antigen HEMPAS"
	case AllergyIntoleranceCode3982003:
		return "Oxaloacetate decarboxylase"
	case AllergyIntoleranceCode3983008:
		return "N,-N-dimethyltryptamine"
	case AllergyIntoleranceCode3990003:
		return "Alkaline phosphatase isoenzyme, bone fraction"
	case AllergyIntoleranceCode3994007:
		return "Hemoglobin Tampa"
	case AllergyIntoleranceCode4014000:
		return "Sulfisomidine"
	case AllergyIntoleranceCode4024008:
		return "Soft metal"
	case AllergyIntoleranceCode4025009:
		return "Captodiamine"
	case AllergyIntoleranceCode4043008:
		return "Etidocaine hydrochloride"
	case AllergyIntoleranceCode4047009:
		return "cis-1,2-Dihydrobenzene-1,2-diol dehydrogenase"
	case AllergyIntoleranceCode4048004:
		return "1,1,2,2-tetrachloro-1,2- difluoroethane (substance)"
	case AllergyIntoleranceCode4067000:
		return "Chorismate mutase"
	case AllergyIntoleranceCode4076007:
		return "Parathyroid hormone"
	case AllergyIntoleranceCode4077003:
		return "Dihydrolipoamide succinyltransferase"
	case AllergyIntoleranceCode4080002:
		return "Hemoglobin Grady, Dakar"
	case AllergyIntoleranceCode4091009:
		return "Enteropeptidase"
	case AllergyIntoleranceCode4097008:
		return "Apo-SAA complex"
	case AllergyIntoleranceCode4104007:
		return "Chondroitin sulfate"
	case AllergyIntoleranceCode4105008:
		return "Adenylate cyclase"
	case AllergyIntoleranceCode4115002:
		return "Blood group antibody Norlander"
	case AllergyIntoleranceCode4137009:
		return "sec-Butyl acetate"
	case AllergyIntoleranceCode4153007:
		return "Long-chain-enoyl-CoA hydratase"
	case AllergyIntoleranceCode4167003:
		return "Lymphocyte antigen CD31"
	case AllergyIntoleranceCode4169000:
		return "Blood group antibody Le^bH^"
	case AllergyIntoleranceCode4177001:
		return "Hemoglobin Long Island-Marseille"
	case AllergyIntoleranceCode4182008:
		return "CDPdiacylglycerol-serine O-phosphatidyl-transferase"
	case AllergyIntoleranceCode4188007:
		return "Fibrinogen Sydney II"
	case AllergyIntoleranceCode4200007:
		return "Neriifolin"
	case AllergyIntoleranceCode4201006:
		return "6-aminohexanoate-dimer hydrolase"
	case AllergyIntoleranceCode4203009:
		return "Imipramine pamoate"
	case AllergyIntoleranceCode4207005:
		return "Cortisone beta-reductase"
	case AllergyIntoleranceCode4217000:
		return "Fluorosilicate salt"
	case AllergyIntoleranceCode4218005:
		return "Immunoglobulin, GM>23< allotype"
	case AllergyIntoleranceCode4231000:
		return "Gallium isotope"
	case AllergyIntoleranceCode4239003:
		return "Glycerol dehydrogenase"
	case AllergyIntoleranceCode4255005:
		return "Americium-241 (substance)"
	case AllergyIntoleranceCode4289006:
		return "Keyhole-limpet hemocyanin"
	case AllergyIntoleranceCode4290002:
		return "Linamarin synthase"
	case AllergyIntoleranceCode4314009:
		return "Blood group antibody Allchurch"
	case AllergyIntoleranceCode4334005:
		return "Tar oil"
	case AllergyIntoleranceCode4342006:
		return "2-aminopyridine (substance)"
	case AllergyIntoleranceCode4353000:
		return "Dibutyl phthalate"
	case AllergyIntoleranceCode4355007:
		return "Coagulation factor IX San Dimas variant"
	case AllergyIntoleranceCode4362003:
		return "4-coumarate-CoA ligase"
	case AllergyIntoleranceCode4370008:
		return "Acetone"
	case AllergyIntoleranceCode4393002:
		return "Blood group antigen Fedor"
	case AllergyIntoleranceCode4401009:
		return "Blood group antibody H>T<"
	case AllergyIntoleranceCode4422003:
		return "Blood group antigen"
	case AllergyIntoleranceCode4423008:
		return "Fibrinogen New York II"
	case AllergyIntoleranceCode4425001:
		return "Blood group antibody Binge"
	case AllergyIntoleranceCode4435007:
		return "Sulfuryl fluoride"
	case AllergyIntoleranceCode4437004:
		return "Cesium-127 (substance)"
	case AllergyIntoleranceCode4471008:
		return "Californium-244 (substance)"
	case AllergyIntoleranceCode4479005:
		return "Hemoglobin Brockton"
	case AllergyIntoleranceCode4480008:
		return "Sulfaethidole"
	case AllergyIntoleranceCode4509009:
		return "Plant phenanthrene toxin"
	case AllergyIntoleranceCode4534009:
		return "Bismuth-208 (substance)"
	case AllergyIntoleranceCode4540002:
		return "ADP deaminase"
	case AllergyIntoleranceCode4546008:
		return "Aliphatic carboxylic acid, C14:0"
	case AllergyIntoleranceCode4555006:
		return "Blood group antibody Rils"
	case AllergyIntoleranceCode4560005:
		return "Hemoglobin Mizuho"
	case AllergyIntoleranceCode4561009:
		return "Arginine decarboxylase"
	case AllergyIntoleranceCode4564001:
		return "Blood group antibody Sisson"
	case AllergyIntoleranceCode4567008:
		return "Galactose-1-phosphate thymidylyltransferase"
	case AllergyIntoleranceCode4582003:
		return "Blood group antigen N^A^"
	case AllergyIntoleranceCode4591004:
		return "Blood group antigen Kam"
	case AllergyIntoleranceCode4610008:
		return "Senile cardiac protein"
	case AllergyIntoleranceCode4616002:
		return "Triclobisonium chloride"
	case AllergyIntoleranceCode4629002:
		return "Hypoglycin B"
	case AllergyIntoleranceCode4635002:
		return "Arterial blood"
	case AllergyIntoleranceCode4643007:
		return "Calf thymus ribonuclease H"
	case AllergyIntoleranceCode4656000:
		return "Alcian blue 8GX"
	case AllergyIntoleranceCode4674009:
		return "2,3-dihydroxybenzoate serine ligase (substance)"
	case AllergyIntoleranceCode4681002:
		return "Potassium permanganate"
	case AllergyIntoleranceCode4693006:
		return "Chromium (51-Cr) albumin (substance)"
	case AllergyIntoleranceCode4700006:
		return "Beef insulin"
	case AllergyIntoleranceCode4706000:
		return "Chlorine monoxide"
	case AllergyIntoleranceCode4714006:
		return "Osmium-183m (substance)"
	case AllergyIntoleranceCode4728000:
		return "Scopulariopsis proteinase"
	case AllergyIntoleranceCode4732006:
		return "Oncogene protein P55, V-MYC"
	case AllergyIntoleranceCode4746006:
		return "Hemoglobin Mito"
	case AllergyIntoleranceCode4761007:
		return "Lymphocyte antigen CD30"
	case AllergyIntoleranceCode4762000:
		return "Platelet antigen HPA-3b"
	case AllergyIntoleranceCode4777008:
		return "Fluroxene"
	case AllergyIntoleranceCode4780009:
		return "Secbutabarbital sodium"
	case AllergyIntoleranceCode4786003:
		return "Beta-1,4-mannosyl-glycoprotein beta-1,4-N-acetylglucosaminyltransferase (substance)"
	case AllergyIntoleranceCode4789005:
		return "Blood group antibody Bultar"
	case AllergyIntoleranceCode4793004:
		return "Azobenzene reductase"
	case AllergyIntoleranceCode4814001:
		return "Valethamate"
	case AllergyIntoleranceCode4824009:
		return "Amine oxidase (flavin-containing)"
	case AllergyIntoleranceCode4825005:
		return "Peptidyl-glycinamidase"
	case AllergyIntoleranceCode4831008:
		return "Arabinose-5-phosphate isomerase"
	case AllergyIntoleranceCode4832001:
		return "Technetium (99m-Tc) mebrofenin (substance)"
	case AllergyIntoleranceCode4833006:
		return "Glucan endo-1,3-alpha-glucosidase"
	case AllergyIntoleranceCode4844003:
		return "3,3' T>2<"
	case AllergyIntoleranceCode4864008:
		return "Adenylic acid"
	case AllergyIntoleranceCode4872005:
		return "Glucosulfone"
	case AllergyIntoleranceCode4878009:
		return "HLA-Dw3 antigen"
	case AllergyIntoleranceCode4882006:
		return "Ichthyoallyeinotoxin"
	case AllergyIntoleranceCode4889002:
		return "Xylulokinase"
	case AllergyIntoleranceCode4901003:
		return "Pyruvate oxidase (CoA-acetylating)"
	case AllergyIntoleranceCode4925006:
		return "Oncogene protein V-ABC"
	case AllergyIntoleranceCode4933007:
		return "Lymphocyte antigen CD15"
	case AllergyIntoleranceCode4940008:
		return "Tattoo dye"
	case AllergyIntoleranceCode4955004:
		return "Neoplastic structural gene"
	case AllergyIntoleranceCode4962008:
		return "Tree bark"
	case AllergyIntoleranceCode4963003:
		return "Neutral amino acid"
	case AllergyIntoleranceCode4965005:
		return "Glutathione reductase (NAD(P)H)"
	case AllergyIntoleranceCode4968007:
		return "Acumentin"
	case AllergyIntoleranceCode4986005:
		return "Magnesium borate"
	case AllergyIntoleranceCode5003005:
		return "Hemoglobin Swan River"
	case AllergyIntoleranceCode5004004:
		return "Blood group antibody Panzar"
	case AllergyIntoleranceCode5007006:
		return "Papain"
	case AllergyIntoleranceCode5024000:
		return "Fresh water"
	case AllergyIntoleranceCode5031001:
		return "3-3'dichlorobenzidine"
	case AllergyIntoleranceCode5040002:
		return "Cesium"
	case AllergyIntoleranceCode5043000:
		return "Erythrosin Y"
	case AllergyIntoleranceCode5045007:
		return "Oncogene protein TCL4"
	case AllergyIntoleranceCode5059000:
		return "Technetium-97 (substance)"
	case AllergyIntoleranceCode5060005:
		return "Cesium-132 (substance)"
	case AllergyIntoleranceCode5061009:
		return "Protein-methionine-S-oxide reductase"
	case AllergyIntoleranceCode5064001:
		return "Blood group antibody D 1276"
	case AllergyIntoleranceCode5081005:
		return "Blood group antigen hr^B^"
	case AllergyIntoleranceCode5086000:
		return "Gelsolin"
	case AllergyIntoleranceCode5094007:
		return "Blood group antigen Rios"
	case AllergyIntoleranceCode5098005:
		return "Fennel oil"
	case AllergyIntoleranceCode5109006:
		return "Methylated-DNA-protein-cysteine methyltransferase"
	case AllergyIntoleranceCode5142007:
		return "Coagulation factor II Houston variant"
	case AllergyIntoleranceCode5160007:
		return "Metallic compound"
	case AllergyIntoleranceCode5163009:
		return "Scombrotoxin"
	case AllergyIntoleranceCode5167005:
		return "Zinc chloride fumes"
	case AllergyIntoleranceCode5172001:
		return "Coagulation factor Xa"
	case AllergyIntoleranceCode5179005:
		return "Connective tissue fiber"
	case AllergyIntoleranceCode5200001:
		return "trans-Epoxysuccinate hydrolase"
	case AllergyIntoleranceCode5206007:
		return "Cyanate compound"
	case AllergyIntoleranceCode5220000:
		return "Bacitracin"
	case AllergyIntoleranceCode5226006:
		return "Flavone O^7^-beta-glucosyltransferase"
	case AllergyIntoleranceCode5250008:
		return "Thymus-independent antigen"
	case AllergyIntoleranceCode5252000:
		return "Hafnium radioisotope"
	case AllergyIntoleranceCode5253005:
		return "Hemoglobin Woodville"
	case AllergyIntoleranceCode5259009:
		return "Blood group antigen Braden"
	case AllergyIntoleranceCode5289002:
		return "Scilliroside"
	case AllergyIntoleranceCode5303002:
		return "Hemoglobin Hoshida"
	case AllergyIntoleranceCode5305009:
		return "Polynucleotide"
	case AllergyIntoleranceCode5307001:
		return "Blood group antigen Hamet"
	case AllergyIntoleranceCode5312000:
		return "Zinc-65 (substance)"
	case AllergyIntoleranceCode5323001:
		return "Uridine diphosphate glucuronic acid"
	case AllergyIntoleranceCode5330007:
		return "Actin-binding protein"
	case AllergyIntoleranceCode5339008:
		return "L-glycol dehydrogenase"
	case AllergyIntoleranceCode5340005:
		return "Blood group antigen Swietlik"
	case AllergyIntoleranceCode5392001:
		return "Propylene glycol monomethyl ether"
	case AllergyIntoleranceCode5395004:
		return "Pyridoxamine-phosphate oxidase"
	case AllergyIntoleranceCode5404007:
		return "Lymphocyte antigen CD45RA"
	case AllergyIntoleranceCode5405008:
		return "Cobalt-60 (substance)"
	case AllergyIntoleranceCode5406009:
		return "Beta-L-arabinosidase (substance)"
	case AllergyIntoleranceCode5420002:
		return "Accessory sinus mucus"
	case AllergyIntoleranceCode5439007:
		return "Blood group antibody Do^a^"
	case AllergyIntoleranceCode5442001:
		return "Page blue 83"
	case AllergyIntoleranceCode5453007:
		return "Iridium isotope"
	case AllergyIntoleranceCode5471000:
		return "Hemoglobin G-Coushatta"
	case AllergyIntoleranceCode5474008:
		return "Propionate-CoA ligase"
	case AllergyIntoleranceCode5477001:
		return "Ferric subsulfate"
	case AllergyIntoleranceCode5483003:
		return "Oxalate CoA-transferase"
	case AllergyIntoleranceCode5504009:
		return "Blood group antigen Fuerhart"
	case AllergyIntoleranceCode5511008:
		return "Inosinate nucleosidase"
	case AllergyIntoleranceCode5513006:
		return "Immunoglobulin A, H chain"
	case AllergyIntoleranceCode5515004:
		return "Rhodium fumes"
	case AllergyIntoleranceCode5533005:
		return "Blood group antibody Kp^a^"
	case AllergyIntoleranceCode5537006:
		return "Immunoglobulin D, H chain (substance)"
	case AllergyIntoleranceCode5540006:
		return "Calcium"
	case AllergyIntoleranceCode5547009:
		return "Plutonium-233 (substance)"
	case AllergyIntoleranceCode5548004:
		return "2-dehydro-3-deoxy-D-pentonate aldolase"
	case AllergyIntoleranceCode5568005:
		return "Hemoglobin Hijiyama"
	case AllergyIntoleranceCode5573004:
		return "Blood group antigen Oca"
	case AllergyIntoleranceCode5589001:
		return "Licodione O^2'^-methyltransferase"
	case AllergyIntoleranceCode5590005:
		return "Beryllium radioisotope"
	case AllergyIntoleranceCode5628003:
		return "Hemoglobin I-High Wycombe"
	case AllergyIntoleranceCode5629006:
		return "Cytidylic acid"
	case AllergyIntoleranceCode5637003:
		return "HLA-DQw6 antigen"
	case AllergyIntoleranceCode5641004:
		return "Valproate semisodium"
	case AllergyIntoleranceCode5647000:
		return "Griseofulvin ultramicrosize"
	case AllergyIntoleranceCode5656008:
		return "Antimony-116m (substance)"
	case AllergyIntoleranceCode5659001:
		return "Hemoglobin J-Tongariki"
	case AllergyIntoleranceCode5670008:
		return "Gold isotope"
	case AllergyIntoleranceCode5681006:
		return "Ceftizoxime sodium"
	case AllergyIntoleranceCode5691000:
		return "Absorbable gelatin sponge"
	case AllergyIntoleranceCode5692007:
		return "Cyanocobalamin (58-Co) (substance)"
	case AllergyIntoleranceCode5699003:
		return "Somatomedin C"
	case AllergyIntoleranceCode5700002:
		return "Blood group antibody Gomez"
	case AllergyIntoleranceCode5702005:
		return "Silver-106m (substance)"
	case AllergyIntoleranceCode5704006:
		return "Galactokinase"
	case AllergyIntoleranceCode5705007:
		return "3-hydroxypropion-aldehyde reductase"
	case AllergyIntoleranceCode5739006:
		return "Stramonium"
	case AllergyIntoleranceCode5746002:
		return "Antimony-118m (substance)"
	case AllergyIntoleranceCode5757007:
		return "HLA-Cw8 antigen"
	case AllergyIntoleranceCode5762008:
		return "Heterogeneous nuclear RNA"
	case AllergyIntoleranceCode5764009:
		return "Plutonium-242 (substance)"
	case AllergyIntoleranceCode5767002:
		return "Sulfamerazine"
	case AllergyIntoleranceCode5774007:
		return "White petrolatum"
	case AllergyIntoleranceCode5800007:
		return "tRNA (5-methylaminomethyl-2-thiouridylate)-methyltransferase"
	case AllergyIntoleranceCode5813001:
		return "Malate dehydrogenase"
	case AllergyIntoleranceCode5826002:
		return "Ethyl-4-bis-(hydroxypropyl)-1-aminobenzoate"
	case AllergyIntoleranceCode5827006:
		return "Crotonaldehyde"
	case AllergyIntoleranceCode5829009:
		return "Hemoglobin Vaasa"
	case AllergyIntoleranceCode5830004:
		return "Hemoglobin Bart"
	case AllergyIntoleranceCode5840001:
		return "Blood group antibody Wj"
	case AllergyIntoleranceCode5858007:
		return "Indium-110m (substance)"
	case AllergyIntoleranceCode5863006:
		return "Vitexin beta-glucosyltransferase"
	case AllergyIntoleranceCode5896008:
		return "Hellebrin"
	case AllergyIntoleranceCode5899001:
		return "Bacterial structural gene"
	case AllergyIntoleranceCode5907009:
		return "Quinidine polygalacturonate"
	case AllergyIntoleranceCode5910002:
		return "Oncogene protein PP60, V-SRC"
	case AllergyIntoleranceCode5915007:
		return "Blood group antigen Gladding"
	case AllergyIntoleranceCode5927005:
		return "Lactaldehyde dehydrogenase"
	case AllergyIntoleranceCode5931004:
		return "Technetium (99m-Tc) sulfur colloid (substance)"
	case AllergyIntoleranceCode5932006:
		return "Cysteine"
	case AllergyIntoleranceCode5950004:
		return "3',5'-cyclic-nucleotide phosphodiesterase"
	case AllergyIntoleranceCode5955009:
		return "Diethylene glycol"
	case AllergyIntoleranceCode5977008:
		return "Blood group antigen Bullock"
	case AllergyIntoleranceCode5989005:
		return "Immunoglobulin, GM>17< allotype"
	case AllergyIntoleranceCode5991002:
		return "D-fuconate dehydratase"
	case AllergyIntoleranceCode6021003:
		return "Yttrium-88 (substance)"
	case AllergyIntoleranceCode6038004:
		return "Oxygen radioisotope"
	case AllergyIntoleranceCode6043006:
		return "Bone cement"
	case AllergyIntoleranceCode6044000:
		return "Carbon disulfide"
	case AllergyIntoleranceCode6054001:
		return "Doxylamine succinate"
	case AllergyIntoleranceCode6056004:
		return "Blood group antibody Wk^a^"
	case AllergyIntoleranceCode6068008:
		return "Blood group antigen Mil"
	case AllergyIntoleranceCode6083003:
		return "Hydroxylysine"
	case AllergyIntoleranceCode6085005:
		return "Synovial fluid"
	case AllergyIntoleranceCode6088007:
		return "Benzfetamine hydrochloride"
	case AllergyIntoleranceCode6089004:
		return "Lochia alba"
	case AllergyIntoleranceCode6091007:
		return "Blood group antibody L Harris"
	case AllergyIntoleranceCode6107003:
		return "Asparagusate reductase (NADH)"
	case AllergyIntoleranceCode6109000:
		return "Aromatic-amino-acid aminotransferase"
	case AllergyIntoleranceCode6115000:
		return "Blood group antibody Anuszewska"
	case AllergyIntoleranceCode6135004:
		return "Blood group antigen Duck"
	case AllergyIntoleranceCode6138002:
		return "Blood group antigen Le Provost"
	case AllergyIntoleranceCode6162007:
		return "Meclocycline"
	case AllergyIntoleranceCode6170002:
		return "Heat labile antibody"
	case AllergyIntoleranceCode6172005:
		return "Fatty-acid methyltransferase"
	case AllergyIntoleranceCode6178009:
		return "Lymphocyte antigen CD63"
	case AllergyIntoleranceCode6179001:
		return "O-methyl-bufotenine"
	case AllergyIntoleranceCode6182006:
		return "Chloroacetone"
	case AllergyIntoleranceCode6197009:
		return "Blood group antigen Zd"
	case AllergyIntoleranceCode6237004:
		return "Bemegride"
	case AllergyIntoleranceCode6249003:
		return "Potassium metabisulfite"
	case AllergyIntoleranceCode6256009:
		return "Ribose isomerase"
	case AllergyIntoleranceCode6257000:
		return "Sodium (22-Na) chloride (substance)"
	case AllergyIntoleranceCode6260007:
		return "Protokylol"
	case AllergyIntoleranceCode6261006:
		return "Indoklon"
	case AllergyIntoleranceCode6263009:
		return "Plant residue"
	case AllergyIntoleranceCode6264003:
		return "Diazinon"
	case AllergyIntoleranceCode6287006:
		return "Methidathion"
	case AllergyIntoleranceCode6291001:
		return "Lysosomal alpha-N-acetylglucosaminidase"
	case AllergyIntoleranceCode6301006:
		return "Tantalum-178 (substance)"
	case AllergyIntoleranceCode6310003:
		return "Particulate antigen"
	case AllergyIntoleranceCode6314007:
		return "Phenol beta-glucosyltransferase"
	case AllergyIntoleranceCode6333002:
		return "Squill extract"
	case AllergyIntoleranceCode6338006:
		return "Imidazolonepropionase"
	case AllergyIntoleranceCode6356006:
		return "Chlorodiallylacetamide"
	case AllergyIntoleranceCode6360009:
		return "Kallidin II"
	case AllergyIntoleranceCode6367007:
		return "Technetium-95m (substance)"
	case AllergyIntoleranceCode6386004:
		return "N-Acetylneuraminate O^4^-acetyltransferase"
	case AllergyIntoleranceCode6394006:
		return "Phentermine hydrochloride"
	case AllergyIntoleranceCode6401007:
		return "Lichenase"
	case AllergyIntoleranceCode6409009:
		return "Morpholine"
	case AllergyIntoleranceCode6411000:
		return "Interleukin-12"
	case AllergyIntoleranceCode6422001:
		return "HLA-DRw14 antigen"
	case AllergyIntoleranceCode6451002:
		return "Chlorobenzilate"
	case AllergyIntoleranceCode6455006:
		return "Chloroprene"
	case AllergyIntoleranceCode6469006:
		return "1,2-didehydropipecolate reductase"
	case AllergyIntoleranceCode6478000:
		return "Phosphohexokinase"
	case AllergyIntoleranceCode6495008:
		return "Fibrinogen Montreal II"
	case AllergyIntoleranceCode6507009:
		return "Blood group antigen Much"
	case AllergyIntoleranceCode6513000:
		return "Flumethiazide"
	case AllergyIntoleranceCode6516008:
		return "Indium (111-In) ferric hydroxide (substance)"
	case AllergyIntoleranceCode6524003:
		return "Distilled spirits"
	case AllergyIntoleranceCode6529008:
		return "Blood group antigen Cl^a^"
	case AllergyIntoleranceCode6532006:
		return "Macrophage activating factor"
	case AllergyIntoleranceCode6590001:
		return "Galactosylceramidase"
	case AllergyIntoleranceCode6592009:
		return "HLA-Dw12 antigen"
	case AllergyIntoleranceCode6602005:
		return "Aminoacridine (substance)"
	case AllergyIntoleranceCode6611005:
		return "Diethylaminoethanol"
	case AllergyIntoleranceCode6612003:
		return "Chloramphenicol sodium succinate"
	case AllergyIntoleranceCode6619007:
		return "Bilirubin Y transport protein"
	case AllergyIntoleranceCode6642000:
		return "Opsonin"
	case AllergyIntoleranceCode6644004:
		return "Homoserine dehydrogenase"
	case AllergyIntoleranceCode6671004:
		return "Blood group antigen Caw"
	case AllergyIntoleranceCode6672006:
		return "Phosphoadenylate 3'-nucleotidase"
	case AllergyIntoleranceCode6699008:
		return "Titanium radioisotope"
	case AllergyIntoleranceCode6701008:
		return "Lissamine fast red B"
	case AllergyIntoleranceCode6702001:
		return "Ethyl mercaptoethyl diethyl thiophosphate"
	case AllergyIntoleranceCode6709005:
		return "Gentamicin 2\"-nucleotidyltransferase"
	case AllergyIntoleranceCode6710000:
		return "Nitric oxide"
	case AllergyIntoleranceCode6713003:
		return "Yttrium-91 (substance)"
	case AllergyIntoleranceCode6717002:
		return "Nifuroxime"
	case AllergyIntoleranceCode6725000:
		return "Methylthioninium chloride"
	case AllergyIntoleranceCode6730001:
		return "Uranium-234 (substance)"
	case AllergyIntoleranceCode6741004:
		return "Anti DNA antibody"
	case AllergyIntoleranceCode6755007:
		return "TL antigen"
	case AllergyIntoleranceCode6786001:
		return "Silver difluoride"
	case AllergyIntoleranceCode6790004:
		return "Aminopterin"
	case AllergyIntoleranceCode6792007:
		return "Veratrine"
	case AllergyIntoleranceCode6808006:
		return "Ferrous iron compound"
	case AllergyIntoleranceCode6809003:
		return "Phomopsin"
	case AllergyIntoleranceCode6814004:
		return "Discadenine synthase"
	case AllergyIntoleranceCode6817006:
		return "Oxidized glutathione"
	case AllergyIntoleranceCode6826009:
		return "Sterol hormone"
	case AllergyIntoleranceCode6837005:
		return "Dextropropoxyphene napsylate"
	case AllergyIntoleranceCode6854002:
		return "Platinum-188 (substance)"
	case AllergyIntoleranceCode6865007:
		return "Theophylline calcium salicylate"
	case AllergyIntoleranceCode6873003:
		return "Cefapirin sodium"
	case AllergyIntoleranceCode6879004:
		return "Mead acid"
	case AllergyIntoleranceCode6881002:
		return "Magnesium fumes"
	case AllergyIntoleranceCode6884005:
		return "(S)-3-Amino-2-methylpropionate aminotransferase"
	case AllergyIntoleranceCode6890009:
		return "3-deoxy-manno-octulosonate-8-phosphatase (substance)"
	case AllergyIntoleranceCode6896003:
		return "Thiopurine methyltransferase"
	case AllergyIntoleranceCode6910009:
		return "Sodium fluoride"
	case AllergyIntoleranceCode6911008:
		return "Deoxycytidylate methyltransferase"
	case AllergyIntoleranceCode6916003:
		return "Bowieine"
	case AllergyIntoleranceCode6924008:
		return "Exopolyphosphatase"
	case AllergyIntoleranceCode6925009:
		return "Leucine acetyltransferase"
	case AllergyIntoleranceCode6927001:
		return "Tin-121 (substance)"
	case AllergyIntoleranceCode6937006:
		return "Thymidylate synthase"
	case AllergyIntoleranceCode6945001:
		return "Blood group antigen Le^bH^"
	case AllergyIntoleranceCode6952004:
		return "Tin-121m (substance)"
	case AllergyIntoleranceCode6958000:
		return "Blood group antibody Frando"
	case AllergyIntoleranceCode6961004:
		return "Lysolecithin acylmutase"
	case AllergyIntoleranceCode6970001:
		return "4-hydroxyproline epimerase (substance)"
	case AllergyIntoleranceCode6973004:
		return "Chromium (51-Cr) chloride (substance)"
	case AllergyIntoleranceCode6983000:
		return "Acrylamide"
	case AllergyIntoleranceCode6992002:
		return "Triflupromazine hydrochloride"
	case AllergyIntoleranceCode6993007:
		return "Seminal fluid"
	case AllergyIntoleranceCode6999006:
		return "Ammonium compound"
	case AllergyIntoleranceCode7008002:
		return "Beta-carotene 15,15'-dioxygenase"
	case AllergyIntoleranceCode7018007:
		return "Malate-CoA ligase"
	case AllergyIntoleranceCode7029006:
		return "Blood group antigen Greenlee"
	case AllergyIntoleranceCode7030001:
		return "Globoside"
	case AllergyIntoleranceCode7034005:
		return "Diclofenac"
	case AllergyIntoleranceCode7045008:
		return "Lycorine"
	case AllergyIntoleranceCode7047000:
		return "Asphyxiant atmosphere"
	case AllergyIntoleranceCode7049002:
		return "Pyruvate carboxylase"
	case AllergyIntoleranceCode7054006:
		return "Hemoglobin Poissy"
	case AllergyIntoleranceCode7056008:
		return "3-propylmalate synthase"
	case AllergyIntoleranceCode7059001:
		return "N-Acylneuraminate-9-phosphatase"
	case AllergyIntoleranceCode7061005:
		return "Anthocyanidin O^3^-glucosyltransferase"
	case AllergyIntoleranceCode7070008:
		return "Convallamarin"
	case AllergyIntoleranceCode7084003:
		return "Fibrinogen Buenos Aires II"
	case AllergyIntoleranceCode7110002:
		return "Germanium-69 (substance)"
	case AllergyIntoleranceCode7120007:
		return "Antigen"
	case AllergyIntoleranceCode7132006:
		return "Gallium-73 (substance)"
	case AllergyIntoleranceCode7139002:
		return "Acid-CoA ligase (GDP-forming)"
	case AllergyIntoleranceCode7146006:
		return "Cyclohexene oxide"
	case AllergyIntoleranceCode7152007:
		return "Chlorthion"
	case AllergyIntoleranceCode7156005:
		return "Phosphorus isotope"
	case AllergyIntoleranceCode7158006:
		return "HLA-Dw19 antigen"
	case AllergyIntoleranceCode7161007:
		return "Complement component C2a"
	case AllergyIntoleranceCode7179006:
		return "Prekallikrein"
	case AllergyIntoleranceCode7191002:
		return "Methenyltetrahydrofolate cyclohydrolase"
	case AllergyIntoleranceCode7208009:
		return "Thiol oxidase"
	case AllergyIntoleranceCode7211005:
		return "Blood group antibody Haakestad"
	case AllergyIntoleranceCode7237008:
		return "Galactonate dehydratase"
	case AllergyIntoleranceCode7243005:
		return "Methyl isocyanate"
	case AllergyIntoleranceCode7269004:
		return "Thorium"
	case AllergyIntoleranceCode7271004:
		return "Mixed dust"
	case AllergyIntoleranceCode7280004:
		return "dTDP4-dehydrorhamnose reductase"
	case AllergyIntoleranceCode7281000:
		return "Technetium (99m-Tc) lidofenin (substance)"
	case AllergyIntoleranceCode7284008:
		return "Mercaptan compound"
	case AllergyIntoleranceCode7294003:
		return "Acetic acid tert-butyl ester"
	case AllergyIntoleranceCode7302008:
		return "Ambuphylline"
	case AllergyIntoleranceCode7318002:
		return "Bacteriochlorophyll"
	case AllergyIntoleranceCode7321000:
		return "Pyrimidine"
	case AllergyIntoleranceCode7325009:
		return "Calcium hydroxide"
	case AllergyIntoleranceCode7327001:
		return "Sulfurous acid"
	case AllergyIntoleranceCode7328006:
		return "Red petrolatum"
	case AllergyIntoleranceCode7330008:
		return "Shellac"
	case AllergyIntoleranceCode7337006:
		return "Blood group antibody Tr^a^"
	case AllergyIntoleranceCode7348004:
		return "Coagulation factor II"
	case AllergyIntoleranceCode7382005:
		return "Aminoalcohol ester"
	case AllergyIntoleranceCode7401000:
		return "Heme-hemopexin complex"
	case AllergyIntoleranceCode7411007:
		return "Blood group antibody HLA-B8"
	case AllergyIntoleranceCode7427000:
		return "Sepiapterin reductase"
	case AllergyIntoleranceCode7434003:
		return "Erythrosin B"
	case AllergyIntoleranceCode7446004:
		return "Ruthenium"
	case AllergyIntoleranceCode7460002:
		return "Tellurium-127 (substance)"
	case AllergyIntoleranceCode7470000:
		return "p-tert-butyltoluene (substance)"
	case AllergyIntoleranceCode7489000:
		return "Homocytotropic antibody"
	case AllergyIntoleranceCode7503004:
		return "Gallium-72 (substance)"
	case AllergyIntoleranceCode7509000:
		return "Mannitol hexanitrate"
	case AllergyIntoleranceCode7515000:
		return "Hepatotoxic mycotoxin"
	case AllergyIntoleranceCode7537007:
		return "Stizolobinate synthase"
	case AllergyIntoleranceCode7547005:
		return "Hemoglobin Lincoln Park"
	case AllergyIntoleranceCode7549008:
		return "Fibrinogen Bethesda I"
	case AllergyIntoleranceCode7588005:
		return "Blood group antibody Sk^a^"
	case AllergyIntoleranceCode7608003:
		return "Triethylene glycol"
	case AllergyIntoleranceCode7616007:
		return "Blood group antibody Pruitt"
	case AllergyIntoleranceCode7648006:
		return "HLA-Bw70 antigen"
	case AllergyIntoleranceCode7661006:
		return "Fish bone"
	case AllergyIntoleranceCode7670009:
		return "Aminobutyraldehyde dehydrogenase"
	case AllergyIntoleranceCode7675004:
		return "Blood group antigen Towey"
	case AllergyIntoleranceCode7685003:
		return "Blood group antibody Bg^c^"
	case AllergyIntoleranceCode7696006:
		return "Ferrovanadium dust"
	case AllergyIntoleranceCode7716001:
		return "Isovaleryl-CoA dehydrogenase"
	case AllergyIntoleranceCode7737009:
		return "Chlortetracycline hydrochloride"
	case AllergyIntoleranceCode7738004:
		return "HLA-B49 antigen"
	case AllergyIntoleranceCode7761002:
		return "Silver-111 (substance)"
	case AllergyIntoleranceCode7770004:
		return "Strontium-89 (substance)"
	case AllergyIntoleranceCode7774008:
		return "Neo-b-vitamin A1 (substance)"
	case AllergyIntoleranceCode7779003:
		return "Ruthenium-103 (substance)"
	case AllergyIntoleranceCode7785005:
		return "Sphingomyelin phosphodiesterase D"
	case AllergyIntoleranceCode7790008:
		return "1-Monoacylglycerol"
	case AllergyIntoleranceCode7791007:
		return "Soya protein"
	case AllergyIntoleranceCode7795003:
		return "Oxalate oxidase"
	case AllergyIntoleranceCode7801007:
		return "Tetrahydroxypteridine cycloisomerase"
	case AllergyIntoleranceCode7816005:
		return "Antazoline hydrochloride"
	case AllergyIntoleranceCode7834009:
		return "Acetyl digitoxin"
	case AllergyIntoleranceCode7846008:
		return "Sphingomyelin phosphodiesterase"
	case AllergyIntoleranceCode7848009:
		return "Monophosphatidylinositol phosphodiesterase"
	case AllergyIntoleranceCode7868003:
		return "Beta-cyclopiazonate dehydrogenase (substance)"
	case AllergyIntoleranceCode7879008:
		return "Radon-218 (substance)"
	case AllergyIntoleranceCode7900007:
		return "Hemoglobin Presbyterian"
	case AllergyIntoleranceCode7904003:
		return "Deanol"
	case AllergyIntoleranceCode7909008:
		return "Arginine carboxypeptidase"
	case AllergyIntoleranceCode7924004:
		return "Diflorasone"
	case AllergyIntoleranceCode7938006:
		return "D-arabitol dehydrogenase"
	case AllergyIntoleranceCode7945006:
		return "Orsellinate-depside hydrolase"
	case AllergyIntoleranceCode7948008:
		return "Reed-Sternberg antibody"
	case AllergyIntoleranceCode7953003:
		return "Thioneb"
	case AllergyIntoleranceCode7957002:
		return "Phosphatidate cytidylyltransferase"
	case AllergyIntoleranceCode7961008:
		return "Hemoglobin F-Shanghai"
	case AllergyIntoleranceCode7970006:
		return "Allograft"
	case AllergyIntoleranceCode7974002:
		return "Blood group antibody Dalman"
	case AllergyIntoleranceCode7975001:
		return "Amiphenazole"
	case AllergyIntoleranceCode7979007:
		return "3'-phosphoadenylylsulfate 3'-phosphatase (substance)"
	case AllergyIntoleranceCode7983007:
		return "Sodium rhodanide"
	case AllergyIntoleranceCode7985000:
		return "Sulfur isotope"
	case AllergyIntoleranceCode7997004:
		return "Butyl mercaptan"
	case AllergyIntoleranceCode8000007:
		return "Cucurbitacin delta^23^-reductase"
	case AllergyIntoleranceCode8002004:
		return "Blood group antibody Fleming"
	case AllergyIntoleranceCode8025003:
		return "Blood group antibody Gibson"
	case AllergyIntoleranceCode8029009:
		return "Allyl glycidyl ether"
	case AllergyIntoleranceCode8030004:
		return "Polyethylene glycol"
	case AllergyIntoleranceCode8035009:
		return "Cholestenol delta-isomerase"
	case AllergyIntoleranceCode8048008:
		return "Blood group antigen Th"
	case AllergyIntoleranceCode8054009:
		return "Orotate reductase (NADPH)"
	case AllergyIntoleranceCode8055005:
		return "Galactoside acetyltransferase"
	case AllergyIntoleranceCode8105004:
		return "Hemoglobin Leiden"
	case AllergyIntoleranceCode8108002:
		return "Undecaprenyl-diphosphatase"
	case AllergyIntoleranceCode8123007:
		return "Blood group antibody Schuppenhauer"
	case AllergyIntoleranceCode8132009:
		return "Magnesium acetylsalicylate"
	case AllergyIntoleranceCode8143001:
		return "Diosmin"
	case AllergyIntoleranceCode8153000:
		return "Homoproline"
	case AllergyIntoleranceCode8156008:
		return "Immunoglobulin, Fd fragment"
	case AllergyIntoleranceCode8164002:
		return "Lymphocyte antigen CD67"
	case AllergyIntoleranceCode8168004:
		return "Uracil-5-carboxylate decarboxylase"
	case AllergyIntoleranceCode8179009:
		return "Cevadilline"
	case AllergyIntoleranceCode8184003:
		return "Convallamarogenin"
	case AllergyIntoleranceCode8190004:
		return "Diaminopimelate epimerase"
	case AllergyIntoleranceCode8202008:
		return "Potassium-43 (substance)"
	case AllergyIntoleranceCode8203003:
		return "Human menopausal gonadotropin"
	case AllergyIntoleranceCode8204009:
		return "Polyester"
	case AllergyIntoleranceCode8222007:
		return "Coagulation factor II Padua variant"
	case AllergyIntoleranceCode8227001:
		return "Ruthenium-106 (substance)"
	case AllergyIntoleranceCode8230008:
		return "Streptococcal cysteine proteinase"
	case AllergyIntoleranceCode8237006:
		return "Strobane"
	case AllergyIntoleranceCode8252004:
		return "Chlorothiazide sodium"
	case AllergyIntoleranceCode8257005:
		return "Abnormal hemoglobin"
	case AllergyIntoleranceCode8261004:
		return "Potassium thiosulfate"
	case AllergyIntoleranceCode8268005:
		return "Blood group antibody Hildebrandt"
	case AllergyIntoleranceCode8270001:
		return "tRNA adenylyltransferase"
	case AllergyIntoleranceCode8275006:
		return "Methionine-S-oxide reductase"
	case AllergyIntoleranceCode8295000:
		return "Uromucoid protein"
	case AllergyIntoleranceCode8300003:
		return "Cyclohexanol"
	case AllergyIntoleranceCode8310007:
		return "Hemoglobin Madrid"
	case AllergyIntoleranceCode8313009:
		return "RNA-directed DNA polymerase"
	case AllergyIntoleranceCode8340009:
		return "Procollagen-lysine,2-oxoglutarate 5-dioxygenase"
	case AllergyIntoleranceCode8342001:
		return "Brilliant cresyl blue"
	case AllergyIntoleranceCode8343006:
		return "Blood group antibody Re^a^"
	case AllergyIntoleranceCode8354001:
		return "Manganese ethylene bis-dithiocarbamate"
	case AllergyIntoleranceCode8355000:
		return "Hafnium isotope"
	case AllergyIntoleranceCode8362009:
		return "Blood group antibody c"
	case AllergyIntoleranceCode8365006:
		return "Oil of pennyroyal-European"
	case AllergyIntoleranceCode8368008:
		return "Xylobiase"
	case AllergyIntoleranceCode8376005:
		return "Duffy blood group antibody"
	case AllergyIntoleranceCode8385005:
		return "Glucan 1,4-alpha-glucosidase"
	case AllergyIntoleranceCode8397006:
		return "Nicotine resin complex"
	case AllergyIntoleranceCode8406008:
		return "Nitroethane oxidase"
	case AllergyIntoleranceCode8429000:
		return "Brilliant orange"
	case AllergyIntoleranceCode8450009:
		return "Oil of lemon grass"
	case AllergyIntoleranceCode8452001:
		return "Blood group antigen Sisson"
	case AllergyIntoleranceCode8456003:
		return "Methyl ethyl ketone peroxide"
	case AllergyIntoleranceCode8460000:
		return "Blood group antibody Vg^a^"
	case AllergyIntoleranceCode8473001:
		return "Homocysteine methyltransferase"
	case AllergyIntoleranceCode8474007:
		return "Lead oleate"
	case AllergyIntoleranceCode8484008:
		return "Blood group antigen Mur"
	case AllergyIntoleranceCode8485009:
		return "Oncogene protein P210, BCR-ABL"
	case AllergyIntoleranceCode8486005:
		return "HLA-DRw15 antigen"
	case AllergyIntoleranceCode8487001:
		return "Vanadium-48 (substance)"
	case AllergyIntoleranceCode8491006:
		return "Complement inhibitor"
	case AllergyIntoleranceCode8492004:
		return "Allantoicase"
	case AllergyIntoleranceCode8498000:
		return "Short neurotoxin venom"
	case AllergyIntoleranceCode8507001:
		return "Cyclohexane"
	case AllergyIntoleranceCode8514004:
		return "Ornithine"
	case AllergyIntoleranceCode8520003:
		return "Hemoglobin Machida"
	case AllergyIntoleranceCode8525008:
		return "Osmium-183 (substance)"
	case AllergyIntoleranceCode8529002:
		return "Urinary protein of low molecular weight"
	case AllergyIntoleranceCode8534003:
		return "Tin-110 (substance)"
	case AllergyIntoleranceCode8537005:
		return "Solution"
	case AllergyIntoleranceCode8578007:
		return "Potassium cyanate"
	case AllergyIntoleranceCode8591008:
		return "Dichlorodifluoromethane"
	case AllergyIntoleranceCode8612007:
		return "Tumor necrosis factor"
	case AllergyIntoleranceCode8620009:
		return "Oncogene protein TCL6"
	case AllergyIntoleranceCode8631001:
		return "Potassium chloride"
	case AllergyIntoleranceCode8653004:
		return "Rubijervine"
	case AllergyIntoleranceCode8660005:
		return "Complement component C3c"
	case AllergyIntoleranceCode8687009:
		return "Gum arabic"
	case AllergyIntoleranceCode8689007:
		return "Kanamycin sulfate"
	case AllergyIntoleranceCode8701002:
		return "Sulfachlorpyridazine"
	case AllergyIntoleranceCode8705006:
		return "4-hydroxybenzoate decarboxylase (substance)"
	case AllergyIntoleranceCode8731008:
		return "Blood group antibody Austin"
	case AllergyIntoleranceCode8740007:
		return "C3(H20)Bb"
	case AllergyIntoleranceCode8761000:
		return "Adenylylsulfate kinase"
	case AllergyIntoleranceCode8767001:
		return "Santonin"
	case AllergyIntoleranceCode8785008:
		return "Chlorine dioxide"
	case AllergyIntoleranceCode8786009:
		return "Blood group antigen Wd^a^"
	case AllergyIntoleranceCode8795001:
		return "Hemoglobin F"
	case AllergyIntoleranceCode8817004:
		return "LH receptor site"
	case AllergyIntoleranceCode8818009:
		return "Blood group antibody Tri W"
	case AllergyIntoleranceCode8822004:
		return "Linoleic acid"
	case AllergyIntoleranceCode8830003:
		return "Nitrate reductase (NAD(P)H)"
	case AllergyIntoleranceCode8836009:
		return "Gallocyanine"
	case AllergyIntoleranceCode8844009:
		return "Hydroxybutyrate-dimer hydrolase"
	case AllergyIntoleranceCode8858006:
		return "Strontium (85-Sr) nitrate (substance)"
	case AllergyIntoleranceCode8865003:
		return "Natural graphite"
	case AllergyIntoleranceCode8878003:
		return "Blood group antigen Evelyn"
	case AllergyIntoleranceCode8882001:
		return "3-hydroxybenzoate 6-hydroxylase"
	case AllergyIntoleranceCode8886003:
		return "Flecainide acetate"
	case AllergyIntoleranceCode8908003:
		return "Blood group antibody I^T^"
	case AllergyIntoleranceCode8914005:
		return "Endolymph"
	case AllergyIntoleranceCode8919000:
		return "Biotin"
	case AllergyIntoleranceCode8926000:
		return "Azur B"
	case AllergyIntoleranceCode8945009:
		return "Phosphopantothenate-cysteine ligase"
	case AllergyIntoleranceCode8953001:
		return "2,3-dihydroxyindole 2,3-dioxygenase (substance)"
	case AllergyIntoleranceCode8963009:
		return "N-Acetylmuramoyl-L-alanine amidase"
	case AllergyIntoleranceCode8969008:
		return "Bulbourethral secretions"
	case AllergyIntoleranceCode8977007:
		return "Blood group antibody Tarplee"
	case AllergyIntoleranceCode8982000:
		return "Oleate hydratase"
	case AllergyIntoleranceCode8987006:
		return "Cycle-phase specific agent"
	case AllergyIntoleranceCode8991001:
		return "Ribulokinase"
	case AllergyIntoleranceCode9010006:
		return "Methyl blue"
	case AllergyIntoleranceCode9013008:
		return "Dephospho-CoA kinase"
	case AllergyIntoleranceCode9021002:
		return "Carbaryl"
	case AllergyIntoleranceCode9024005:
		return "Glucose-6-phosphate dehydrogenase"
	case AllergyIntoleranceCode9045003:
		return "Radon radioisotope"
	case AllergyIntoleranceCode9052001:
		return "Allspice oil"
	case AllergyIntoleranceCode9054000:
		return "Blood group antigen HLA-B15"
	case AllergyIntoleranceCode9103003:
		return "Retinol fatty-acyltransferase"
	case AllergyIntoleranceCode9110009:
		return "Mercuric compound"
	case AllergyIntoleranceCode9125009:
		return "Sempervirine"
	case AllergyIntoleranceCode9159008:
		return "Triacetate-lactonase"
	case AllergyIntoleranceCode9172009:
		return "Blood group antibody Alda"
	case AllergyIntoleranceCode9174005:
		return "Fibrinogen Poitiers"
	case AllergyIntoleranceCode9183000:
		return "Beta-N-acetylgalactosaminidase"
	case AllergyIntoleranceCode9189001:
		return "CMP-N-acetylneuraminate-lactosylceramide alpha-2,3-sialyltransferase"
	case AllergyIntoleranceCode9195000:
		return "Immunoglobulin gene INV allotype"
	case AllergyIntoleranceCode9197008:
		return "Apiose reductase"
	case AllergyIntoleranceCode9205004:
		return "Hemoglobin Tarrant"
	case AllergyIntoleranceCode9220005:
		return "Plant phenol oil"
	case AllergyIntoleranceCode9223007:
		return "Borneol dehydrogenase"
	case AllergyIntoleranceCode9234005:
		return "Chlorobutanol"
	case AllergyIntoleranceCode9246009:
		return "Tellurium-118 (substance)"
	case AllergyIntoleranceCode9253000:
		return "HLA-DRw16 antigen"
	case AllergyIntoleranceCode9270008:
		return "Catecholamine receptor"
	case AllergyIntoleranceCode9271007:
		return "Fibrinogen Pontoise"
	case AllergyIntoleranceCode9301005:
		return "Lens neutral proteinase"
	case AllergyIntoleranceCode9302003:
		return "Gentisate decarboxylase"
	case AllergyIntoleranceCode9315007:
		return "Spearmint oil"
	case AllergyIntoleranceCode9319001:
		return "Blood group antibody Vennera"
	case AllergyIntoleranceCode9334007:
		return "Isopropyl glycidyl ether"
	case AllergyIntoleranceCode9349004:
		return "Nitrobenzene"
	case AllergyIntoleranceCode9351000:
		return "Palladium-103 (substance)"
	case AllergyIntoleranceCode9355009:
		return "Hemoglobin F-Alexandra"
	case AllergyIntoleranceCode9392009:
		return "Blood group antibody Pollio"
	case AllergyIntoleranceCode9396007:
		return "Iron-60 (substance)"
	case AllergyIntoleranceCode9398008:
		return "Blood group antigen Pillsbury"
	case AllergyIntoleranceCode9410003:
		return "Bromoform"
	case AllergyIntoleranceCode9422000:
		return "High density lipoprotein"
	case AllergyIntoleranceCode418038007:
		return "Propensity to adverse reactions to substance"
	case AllergyIntoleranceCode25744000:
		return "Hereditary gastrogenic lactose intolerance"
	case AllergyIntoleranceCode54250004:
		return "Lactose intolerance in children without lactase deficiency"
	case AllergyIntoleranceCode59037007:
		return "Drug intolerance"
	case AllergyIntoleranceCode61712006:
		return "Transient gluten sensitivity"
	case AllergyIntoleranceCode72354005:
		return "Oral contraceptive intolerance"
	case AllergyIntoleranceCode91931000:
		return "Allergy to erythromycin"
	case AllergyIntoleranceCode91932007:
		return "Allergy to fruit"
	case AllergyIntoleranceCode91934008:
		return "Nut allergy"
	case AllergyIntoleranceCode91935009:
		return "Allergy to peanuts"
	case AllergyIntoleranceCode91936005:
		return "Allergy to penicillin"
	case AllergyIntoleranceCode91937001:
		return "Allergy to seafood"
	case AllergyIntoleranceCode91938006:
		return "Allergy to strawberries"
	case AllergyIntoleranceCode91939003:
		return "Allergy to sulfonamides"
	case AllergyIntoleranceCode91940001:
		return "Allergy to walnut"
	case AllergyIntoleranceCode190751001:
		return "Primary lactose intolerance"
	case AllergyIntoleranceCode213020009:
		return "Egg protein allergy"
	case AllergyIntoleranceCode232349006:
		return "House dust allergy"
	case AllergyIntoleranceCode232350006:
		return "House dust mite allergy"
	case AllergyIntoleranceCode235719002:
		return "Food intolerance"
	case AllergyIntoleranceCode293584003:
		return "Acetaminophen allergy"
	case AllergyIntoleranceCode293585002:
		return "Salicylate allergy"
	case AllergyIntoleranceCode293586001:
		return "Aspirin allergy"
	case AllergyIntoleranceCode293588000:
		return "Pentazocine allergy"
	case AllergyIntoleranceCode293589008:
		return "Phenazocine allergy"
	case AllergyIntoleranceCode293590004:
		return "Methadone analog allergy"
	case AllergyIntoleranceCode293591000:
		return "Dextromoramide allergy"
	case AllergyIntoleranceCode293592007:
		return "Dextropropoxyphene allergy"
	case AllergyIntoleranceCode293593002:
		return "Dipipanone allergy"
	case AllergyIntoleranceCode293594008:
		return "Methadone allergy"
	case AllergyIntoleranceCode293595009:
		return "Morphinan opioid allergy"
	case AllergyIntoleranceCode293596005:
		return "Buprenorphine allergy"
	case AllergyIntoleranceCode293597001:
		return "Codeine allergy"
	case AllergyIntoleranceCode293598006:
		return "Diamorphine allergy"
	case AllergyIntoleranceCode293599003:
		return "Dihydrocodeine allergy"
	case AllergyIntoleranceCode293600000:
		return "Nalbuphine allergy"
	case AllergyIntoleranceCode293601001:
		return "Morphine allergy"
	case AllergyIntoleranceCode293602008:
		return "Opium alkaloid allergy"
	case AllergyIntoleranceCode293603003:
		return "Pethidine analog allergy"
	case AllergyIntoleranceCode293604009:
		return "Alfentanil allergy"
	case AllergyIntoleranceCode293605005:
		return "Fentanyl allergy"
	case AllergyIntoleranceCode293606006:
		return "Pethidine allergy"
	case AllergyIntoleranceCode293607002:
		return "Phenoperidine allergy"
	case AllergyIntoleranceCode293608007:
		return "Meptazinol allergy"
	case AllergyIntoleranceCode293609004:
		return "Levorphanol allergy"
	case AllergyIntoleranceCode293610009:
		return "Non-steroidal anti-inflammatory drug allergy"
	case AllergyIntoleranceCode293611008:
		return "Acemetacin allergy"
	case AllergyIntoleranceCode293612001:
		return "Azapropazone allergy"
	case AllergyIntoleranceCode293613006:
		return "Diclofenac allergy"
	case AllergyIntoleranceCode293614000:
		return "Etodolac allergy"
	case AllergyIntoleranceCode293615004:
		return "Felbinac allergy"
	case AllergyIntoleranceCode293616003:
		return "Fenbufen allergy"
	case AllergyIntoleranceCode293617007:
		return "Fenoprofen allergy"
	case AllergyIntoleranceCode293618002:
		return "Flurbiprofen allergy"
	case AllergyIntoleranceCode293619005:
		return "Ibuprofen allergy"
	case AllergyIntoleranceCode293620004:
		return "Indomethacin allergy"
	case AllergyIntoleranceCode293621000:
		return "Ketoprofen allergy"
	case AllergyIntoleranceCode293622007:
		return "Ketorolac allergy"
	case AllergyIntoleranceCode293623002:
		return "Mefenamic acid allergy"
	case AllergyIntoleranceCode293624008:
		return "Nabumetone allergy"
	case AllergyIntoleranceCode293625009:
		return "Naproxen allergy"
	case AllergyIntoleranceCode293626005:
		return "Nefopam allergy"
	case AllergyIntoleranceCode293627001:
		return "Oxyphenbutazone allergy"
	case AllergyIntoleranceCode293628006:
		return "Phenylbutazone allergy"
	case AllergyIntoleranceCode293629003:
		return "Piroxicam allergy"
	case AllergyIntoleranceCode293630008:
		return "Sulindac allergy"
	case AllergyIntoleranceCode293631007:
		return "Tenoxicam allergy"
	case AllergyIntoleranceCode293632000:
		return "Tiaprofenic acid allergy"
	case AllergyIntoleranceCode293633005:
		return "Tolmetin allergy"
	case AllergyIntoleranceCode293635003:
		return "Tuberculin allergy"
	case AllergyIntoleranceCode293636002:
		return "Radiopharmaceutical allergy"
	case AllergyIntoleranceCode293637006:
		return "Contrast media allergy"
	case AllergyIntoleranceCode293645001:
		return "Bismuth chelate allergy"
	case AllergyIntoleranceCode293646000:
		return "Sucralfate allergy"
	case AllergyIntoleranceCode293647009:
		return "Liquorice allergy"
	case AllergyIntoleranceCode293648004:
		return "Misoprostol allergy"
	case AllergyIntoleranceCode293649007:
		return "H2 receptor antagonist allergy"
	case AllergyIntoleranceCode293650007:
		return "Cimetidine allergy"
	case AllergyIntoleranceCode293651006:
		return "Famotidine allergy"
	case AllergyIntoleranceCode293652004:
		return "Nizatidine allergy"
	case AllergyIntoleranceCode293653009:
		return "Ranitidine allergy"
	case AllergyIntoleranceCode293654003:
		return "Proton pump inhibitor allergy"
	case AllergyIntoleranceCode293655002:
		return "Omeprazole allergy"
	case AllergyIntoleranceCode293656001:
		return "Lansoprazole allergy"
	case AllergyIntoleranceCode293657005:
		return "Carbenoxolone allergy"
	case AllergyIntoleranceCode293658000:
		return "Pirenzepine allergy"
	case AllergyIntoleranceCode293659008:
		return "Pancreatin allergy"
	case AllergyIntoleranceCode293660003:
		return "5-aminosalicylic acid allergy"
	case AllergyIntoleranceCode293662006:
		return "Olsalazine allergy"
	case AllergyIntoleranceCode293663001:
		return "Sulfasalazine allergy"
	case AllergyIntoleranceCode293665008:
		return "Magnesium trisilicate allergy"
	case AllergyIntoleranceCode293666009:
		return "Aluminum hydroxide allergy"
	case AllergyIntoleranceCode293668005:
		return "Loperamide allergy"
	case AllergyIntoleranceCode293669002:
		return "Kaolin allergy"
	case AllergyIntoleranceCode293671002:
		return "Cisapride allergy"
	case AllergyIntoleranceCode293673004:
		return "Nabilone allergy"
	case AllergyIntoleranceCode293674005:
		return "Domperidone allergy"
	case AllergyIntoleranceCode293675006:
		return "Metoclopramide allergy"
	case AllergyIntoleranceCode293676007:
		return "5-HT3-receptor antagonist allergy"
	case AllergyIntoleranceCode293678008:
		return "Bisacodyl allergy"
	case AllergyIntoleranceCode293679000:
		return "Danthron allergy"
	case AllergyIntoleranceCode293680002:
		return "Sodium picosulfate allergy"
	case AllergyIntoleranceCode293681003:
		return "Lactulose allergy"
	case AllergyIntoleranceCode293682005:
		return "Magnesium sulfate allergy"
	case AllergyIntoleranceCode293685007:
		return "Cascara allergy"
	case AllergyIntoleranceCode293686008:
		return "Senna allergy"
	case AllergyIntoleranceCode293687004:
		return "Docusate allergy"
	case AllergyIntoleranceCode293690005:
		return "Peppermint oil allergy"
	case AllergyIntoleranceCode293691009:
		return "Alverine allergy"
	case AllergyIntoleranceCode293692002:
		return "Mebeverine allergy"
	case AllergyIntoleranceCode293693007:
		return "Dicyclomine allergy"
	case AllergyIntoleranceCode293694001:
		return "Mepenzolate allergy"
	case AllergyIntoleranceCode293695000:
		return "Pipenzolate allergy"
	case AllergyIntoleranceCode293696004:
		return "Poldine allergy"
	case AllergyIntoleranceCode293697008:
		return "Propantheline allergy"
	case AllergyIntoleranceCode293699006:
		return "Chenodeoxycholic acid allergy"
	case AllergyIntoleranceCode293700007:
		return "Dehydrocholic acid allergy"
	case AllergyIntoleranceCode293701006:
		return "Ursodeoxycholic acid allergy"
	case AllergyIntoleranceCode293706001:
		return "Etomidate allergy"
	case AllergyIntoleranceCode293707005:
		return "Ketamine allergy"
	case AllergyIntoleranceCode293708000:
		return "Propofol allergy"
	case AllergyIntoleranceCode293709008:
		return "Thiopentone allergy"
	case AllergyIntoleranceCode293710003:
		return "Methohexitone allergy"
	case AllergyIntoleranceCode293712006:
		return "Enflurane allergy"
	case AllergyIntoleranceCode293714007:
		return "Halothane allergy"
	case AllergyIntoleranceCode293715008:
		return "Isoflurane allergy"
	case AllergyIntoleranceCode293716009:
		return "Trichloroethylene allergy"
	case AllergyIntoleranceCode293717000:
		return "Desflurane allergy"
	case AllergyIntoleranceCode293718005:
		return "Local anesthetic drug allergy"
	case AllergyIntoleranceCode293719002:
		return "Bupivacaine allergy"
	case AllergyIntoleranceCode293720008:
		return "Cinchocaine allergy"
	case AllergyIntoleranceCode293721007:
		return "Prilocaine allergy"
	case AllergyIntoleranceCode293722000:
		return "Lignocaine allergy"
	case AllergyIntoleranceCode293723005:
		return "Cocaine allergy"
	case AllergyIntoleranceCode293724004:
		return "Benzocaine allergy"
	case AllergyIntoleranceCode293725003:
		return "Amethocaine allergy"
	case AllergyIntoleranceCode293726002:
		return "Oxybuprocaine allergy"
	case AllergyIntoleranceCode293727006:
		return "Procaine allergy"
	case AllergyIntoleranceCode293728001:
		return "Proxymetacaine allergy"
	case AllergyIntoleranceCode293732007:
		return "Amifostine allergy"
	case AllergyIntoleranceCode293733002:
		return "Aldesleukin allergy"
	case AllergyIntoleranceCode293735009:
		return "Molgramostim allergy"
	case AllergyIntoleranceCode293736005:
		return "Lenograstim allergy"
	case AllergyIntoleranceCode293737001:
		return "Filgrastim allergy"
	case AllergyIntoleranceCode293738006:
		return "Levamisole allergy"
	case AllergyIntoleranceCode293740001:
		return "Alkylating drug allergy"
	case AllergyIntoleranceCode293741002:
		return "Mitobronitol allergy"
	case AllergyIntoleranceCode293742009:
		return "Busulfan allergy"
	case AllergyIntoleranceCode293743004:
		return "Treosulfan allergy"
	case AllergyIntoleranceCode293745006:
		return "Thiotepa allergy"
	case AllergyIntoleranceCode293746007:
		return "Nitrogen mustard derivative allergy"
	case AllergyIntoleranceCode293747003:
		return "Chlorambucil allergy"
	case AllergyIntoleranceCode293748008:
		return "Cyclophosphamide allergy"
	case AllergyIntoleranceCode293749000:
		return "Ethoglucid allergy"
	case AllergyIntoleranceCode293750000:
		return "Ifosfamide allergy"
	case AllergyIntoleranceCode293751001:
		return "Melphalan allergy"
	case AllergyIntoleranceCode293752008:
		return "Estramustine allergy"
	case AllergyIntoleranceCode293753003:
		return "Mustine allergy"
	case AllergyIntoleranceCode293754009:
		return "Nitrosurea allergy"
	case AllergyIntoleranceCode293755005:
		return "Carmustine allergy"
	case AllergyIntoleranceCode293756006:
		return "Lomustine allergy"
	case AllergyIntoleranceCode293758007:
		return "Dacarbazine allergy"
	case AllergyIntoleranceCode293760009:
		return "Dactinomycin allergy"
	case AllergyIntoleranceCode293761008:
		return "Bleomycin allergy"
	case AllergyIntoleranceCode293762001:
		return "Mitomycin allergy"
	case AllergyIntoleranceCode293763006:
		return "Plicamycin allergy"
	case AllergyIntoleranceCode293764000:
		return "Aclarubicin allergy"
	case AllergyIntoleranceCode293765004:
		return "Mitozantrone allergy"
	case AllergyIntoleranceCode293766003:
		return "Doxorubicin allergy"
	case AllergyIntoleranceCode293767007:
		return "Epirubicin allergy"
	case AllergyIntoleranceCode293768002:
		return "Idarubicin allergy"
	case AllergyIntoleranceCode293770006:
		return "Mercuric oxide allergy"
	case AllergyIntoleranceCode293771005:
		return "Methotrexate allergy"
	case AllergyIntoleranceCode293772003:
		return "Mercaptopurine allergy"
	case AllergyIntoleranceCode293773008:
		return "Thioguanine allergy"
	case AllergyIntoleranceCode293774002:
		return "Pentostatin allergy"
	case AllergyIntoleranceCode293775001:
		return "Cytarabine allergy"
	case AllergyIntoleranceCode293776000:
		return "Fluorouracil allergy"
	case AllergyIntoleranceCode293777009:
		return "Etoposide allergy"
	case AllergyIntoleranceCode293778004:
		return "Amsacrine allergy"
	case AllergyIntoleranceCode293779007:
		return "Carboplatin allergy"
	case AllergyIntoleranceCode293780005:
		return "Cisplatin allergy"
	case AllergyIntoleranceCode293781009:
		return "Hydroxyurea allergy"
	case AllergyIntoleranceCode293782002:
		return "Procarbazine allergy"
	case AllergyIntoleranceCode293783007:
		return "Razoxane allergy"
	case AllergyIntoleranceCode293784001:
		return "Crisantaspase allergy"
	case AllergyIntoleranceCode293785000:
		return "Paclitaxel allergy"
	case AllergyIntoleranceCode293786004:
		return "Fludarabine allergy"
	case AllergyIntoleranceCode293787008:
		return "Aminoglutethimide allergy"
	case AllergyIntoleranceCode293788003:
		return "Estrogen antagonist allergy"
	case AllergyIntoleranceCode293789006:
		return "Trilostane allergy"
	case AllergyIntoleranceCode293790002:
		return "Tamoxifen allergy"
	case AllergyIntoleranceCode293791003:
		return "Formestane allergy"
	case AllergyIntoleranceCode293792005:
		return "Vinca alkaloid allergy"
	case AllergyIntoleranceCode293793000:
		return "Vinblastine allergy"
	case AllergyIntoleranceCode293794006:
		return "Vincristine allergy"
	case AllergyIntoleranceCode293795007:
		return "Vindesine allergy"
	case AllergyIntoleranceCode293796008:
		return "Dimethyl sulfoxide allergy"
	case AllergyIntoleranceCode293798009:
		return "Cyclosporin allergy"
	case AllergyIntoleranceCode293799001:
		return "Azathioprine allergy"
	case AllergyIntoleranceCode293802005:
		return "Mazindol allergy"
	case AllergyIntoleranceCode293803000:
		return "Phentermine allergy"
	case AllergyIntoleranceCode293804006:
		return "Dexfenfluramine allergy"
	case AllergyIntoleranceCode293805007:
		return "Diethylpropion allergy"
	case AllergyIntoleranceCode293806008:
		return "Fenfluramine allergy"
	case AllergyIntoleranceCode293808009:
		return "Levodopa allergy"
	case AllergyIntoleranceCode293811005:
		return "Amantadine allergy"
	case AllergyIntoleranceCode293812003:
		return "Apomorphine allergy"
	case AllergyIntoleranceCode293813008:
		return "Lysuride allergy"
	case AllergyIntoleranceCode293814002:
		return "Pergolide allergy"
	case AllergyIntoleranceCode293815001:
		return "Bromocriptine allergy"
	case AllergyIntoleranceCode293818004:
		return "Lithium carbonate allergy"
	case AllergyIntoleranceCode293819007:
		return "Lithium citrate allergy"
	case AllergyIntoleranceCode293822009:
		return "Butriptyline allergy"
	case AllergyIntoleranceCode293823004:
		return "Doxepin allergy"
	case AllergyIntoleranceCode293824005:
		return "Iprindole allergy"
	case AllergyIntoleranceCode293825006:
		return "Lofepramine allergy"
	case AllergyIntoleranceCode293826007:
		return "Nortriptyline allergy"
	case AllergyIntoleranceCode293827003:
		return "Trimipramine allergy"
	case AllergyIntoleranceCode293828008:
		return "Amoxapine allergy"
	case AllergyIntoleranceCode293829000:
		return "Amitriptyline allergy"
	case AllergyIntoleranceCode293830005:
		return "Clomipramine allergy"
	case AllergyIntoleranceCode293831009:
		return "Desipramine allergy"
	case AllergyIntoleranceCode293832002:
		return "Dothiepin allergy"
	case AllergyIntoleranceCode293833007:
		return "Imipramine allergy"
	case AllergyIntoleranceCode293834001:
		return "Protriptyline allergy"
	case AllergyIntoleranceCode293835000:
		return "Monoamine oxidase inhibitor allergy"
	case AllergyIntoleranceCode293836004:
		return "Phenelzine allergy"
	case AllergyIntoleranceCode293837008:
		return "Iproniazid allergy"
	case AllergyIntoleranceCode293838003:
		return "Isocarboxazid allergy"
	case AllergyIntoleranceCode293839006:
		return "Tranylcypromine allergy"
	case AllergyIntoleranceCode293840008:
		return "Moclobemide allergy"
	case AllergyIntoleranceCode293842000:
		return "Tryptophan allergy"
	case AllergyIntoleranceCode293843005:
		return "Venlafaxine allergy"
	case AllergyIntoleranceCode293844004:
		return "Selective serotonin re-uptake inhibitor allergy"
	case AllergyIntoleranceCode293845003:
		return "Sertraline allergy"
	case AllergyIntoleranceCode293847006:
		return "Paroxetine allergy"
	case AllergyIntoleranceCode293848001:
		return "Nefazodone allergy"
	case AllergyIntoleranceCode293849009:
		return "Citalopram allergy"
	case AllergyIntoleranceCode293850009:
		return "Fluoxetine allergy"
	case AllergyIntoleranceCode293851008:
		return "Fluvoxamine allergy"
	case AllergyIntoleranceCode293853006:
		return "Maprotiline allergy"
	case AllergyIntoleranceCode293854000:
		return "Mianserin allergy"
	case AllergyIntoleranceCode293855004:
		return "Trazodone allergy"
	case AllergyIntoleranceCode293856003:
		return "Viloxazine allergy"
	case AllergyIntoleranceCode293858002:
		return "Beclamide allergy"
	case AllergyIntoleranceCode293859005:
		return "Lamotrigine allergy"
	case AllergyIntoleranceCode293860000:
		return "Piracetam allergy"
	case AllergyIntoleranceCode293861001:
		return "Gabapentin allergy"
	case AllergyIntoleranceCode293864009:
		return "Methylphenobarbitone allergy"
	case AllergyIntoleranceCode293865005:
		return "Phenobarbitone allergy"
	case AllergyIntoleranceCode293866006:
		return "Primidone allergy"
	case AllergyIntoleranceCode293867002:
		return "Carbamazepine allergy"
	case AllergyIntoleranceCode293868007:
		return "Vigabatrin allergy"
	case AllergyIntoleranceCode293869004:
		return "Phenytoin allergy"
	case AllergyIntoleranceCode293870003:
		return "Ethosuximide allergy"
	case AllergyIntoleranceCode293871004:
		return "Clonazepam allergy"
	case AllergyIntoleranceCode293874007:
		return "Zopiclone allergy"
	case AllergyIntoleranceCode293875008:
		return "Zolpidem allergy"
	case AllergyIntoleranceCode293876009:
		return "Chlormezanone allergy"
	case AllergyIntoleranceCode293877000:
		return "Methyprylone allergy"
	case AllergyIntoleranceCode293878005:
		return "Paraldehyde allergy"
	case AllergyIntoleranceCode293880004:
		return "Amylobarbitone allergy"
	case AllergyIntoleranceCode293881000:
		return "Butobarbitone allergy"
	case AllergyIntoleranceCode293882007:
		return "Cyclobarbitone allergy"
	case AllergyIntoleranceCode293884008:
		return "Quinalbarbitone allergy"
	case AllergyIntoleranceCode293886005:
		return "Flunitrazepam allergy"
	case AllergyIntoleranceCode293887001:
		return "Flurazepam allergy"
	case AllergyIntoleranceCode293888006:
		return "Loprazolam allergy"
	case AllergyIntoleranceCode293889003:
		return "Lormetazepam allergy"
	case AllergyIntoleranceCode293890007:
		return "Nitrazepam allergy"
	case AllergyIntoleranceCode293891006:
		return "Triazolam allergy"
	case AllergyIntoleranceCode293892004:
		return "Alprazolam allergy"
	case AllergyIntoleranceCode293893009:
		return "Bromazepam allergy"
	case AllergyIntoleranceCode293894003:
		return "Chlordiazepoxide allergy"
	case AllergyIntoleranceCode293895002:
		return "Clobazam allergy"
	case AllergyIntoleranceCode293897005:
		return "Ketazolam allergy"
	case AllergyIntoleranceCode293898000:
		return "Medazepam allergy"
	case AllergyIntoleranceCode293899008:
		return "Oxazepam allergy"
	case AllergyIntoleranceCode293900003:
		return "Prazepam allergy"
	case AllergyIntoleranceCode293901004:
		return "Midazolam allergy"
	case AllergyIntoleranceCode293902006:
		return "Diazepam allergy"
	case AllergyIntoleranceCode293903001:
		return "Lorazepam allergy"
	case AllergyIntoleranceCode293904007:
		return "Temazepam allergy"
	case AllergyIntoleranceCode293906009:
		return "Meprobamate allergy"
	case AllergyIntoleranceCode293908005:
		return "Chloral hydrate allergy"
	case AllergyIntoleranceCode293909002:
		return "Dichloralphenazone allergy"
	case AllergyIntoleranceCode293911006:
		return "Buspirone allergy"
	case AllergyIntoleranceCode293912004:
		return "Chlormethiazole allergy"
	case AllergyIntoleranceCode293914003:
		return "Sulpiride allergy"
	case AllergyIntoleranceCode293915002:
		return "Loxapine allergy"
	case AllergyIntoleranceCode293916001:
		return "Clozapine allergy"
	case AllergyIntoleranceCode293917005:
		return "Risperidone allergy"
	case AllergyIntoleranceCode293918000:
		return "Tetrabenazine allergy"
	case AllergyIntoleranceCode293919008:
		return "Butyrophenone allergy"
	case AllergyIntoleranceCode293920002:
		return "Benperidol allergy"
	case AllergyIntoleranceCode293921003:
		return "Trifluperidol allergy"
	case AllergyIntoleranceCode293923000:
		return "Droperidol allergy"
	case AllergyIntoleranceCode293924006:
		return "Haloperidol allergy"
	case AllergyIntoleranceCode293925007:
		return "Diphenylbutylpiperidine allergy"
	case AllergyIntoleranceCode293926008:
		return "Pimozide allergy"
	case AllergyIntoleranceCode293927004:
		return "Fluspirilene allergy"
	case AllergyIntoleranceCode293928009:
		return "Phenothiazine allergy"
	case AllergyIntoleranceCode293929001:
		return "Methotrimeprazine allergy"
	case AllergyIntoleranceCode293930006:
		return "Pericyazine allergy"
	case AllergyIntoleranceCode293933008:
		return "Thiethylperazine allergy"
	case AllergyIntoleranceCode293934002:
		return "Fluphenazine allergy"
	case AllergyIntoleranceCode293935001:
		return "Chlorpromazine allergy"
	case AllergyIntoleranceCode293936000:
		return "Pipothiazine allergy"
	case AllergyIntoleranceCode293937009:
		return "Promazine allergy"
	case AllergyIntoleranceCode293938004:
		return "Thioridazine allergy"
	case AllergyIntoleranceCode293939007:
		return "Perphenazine allergy"
	case AllergyIntoleranceCode293940009:
		return "Prochlorperazine allergy"
	case AllergyIntoleranceCode293941008:
		return "Trifluoperazine allergy"
	case AllergyIntoleranceCode293942001:
		return "Thioxanthene allergy"
	case AllergyIntoleranceCode293943006:
		return "Chlorprothixene allergy"
	case AllergyIntoleranceCode293946003:
		return "Zuclopenthixol allergy"
	case AllergyIntoleranceCode293948002:
		return "Flupenthixol allergy"
	case AllergyIntoleranceCode293949005:
		return "Oxypertine allergy"
	case AllergyIntoleranceCode293950005:
		return "Remoxipride allergy"
	case AllergyIntoleranceCode293952002:
		return "Selegiline allergy"
	case AllergyIntoleranceCode293954001:
		return "Pemoline allergy"
	case AllergyIntoleranceCode293955000:
		return "Methylphenidate allergy"
	case AllergyIntoleranceCode293956004:
		return "Prolintane allergy"
	case AllergyIntoleranceCode293957008:
		return "Amphetamine group allergy"
	case AllergyIntoleranceCode293958003:
		return "Dexamphetamine allergy"
	case AllergyIntoleranceCode293960001:
		return "Disulfiram allergy"
	case AllergyIntoleranceCode293962009:
		return "Beta-adrenoceptor blocking drug allergy"
	case AllergyIntoleranceCode293963004:
		return "Cardioselective beta-blocker allergy"
	case AllergyIntoleranceCode293964005:
		return "Acebutolol allergy"
	case AllergyIntoleranceCode293965006:
		return "Atenolol allergy"
	case AllergyIntoleranceCode293966007:
		return "Betaxolol allergy"
	case AllergyIntoleranceCode293967003:
		return "Bisoprolol allergy"
	case AllergyIntoleranceCode293968008:
		return "Celiprolol allergy"
	case AllergyIntoleranceCode293969000:
		return "Esmolol allergy"
	case AllergyIntoleranceCode293970004:
		return "Metoprolol allergy"
	case AllergyIntoleranceCode293972007:
		return "Nadolol allergy"
	case AllergyIntoleranceCode293973002:
		return "Pindolol allergy"
	case AllergyIntoleranceCode293974008:
		return "Carvedilol allergy"
	case AllergyIntoleranceCode293975009:
		return "Metipranolol allergy"
	case AllergyIntoleranceCode293976005:
		return "Carteolol allergy"
	case AllergyIntoleranceCode293977001:
		return "Labetalol allergy"
	case AllergyIntoleranceCode293978006:
		return "Levobunolol allergy"
	case AllergyIntoleranceCode293979003:
		return "Oxprenolol allergy"
	case AllergyIntoleranceCode293980000:
		return "Penbutolol allergy"
	case AllergyIntoleranceCode293981001:
		return "Practolol allergy"
	case AllergyIntoleranceCode293982008:
		return "Propranolol allergy"
	case AllergyIntoleranceCode293983003:
		return "Sotalol allergy"
	case AllergyIntoleranceCode293984009:
		return "Timolol allergy"
	case AllergyIntoleranceCode293985005:
		return "Alpha-adrenoceptor blocking drug allergy"
	case AllergyIntoleranceCode293986006:
		return "Alfuzosin allergy"
	case AllergyIntoleranceCode293987002:
		return "Doxazosin allergy"
	case AllergyIntoleranceCode293988007:
		return "Indoramin allergy"
	case AllergyIntoleranceCode293989004:
		return "Phenoxybenzamine allergy"
	case AllergyIntoleranceCode293990008:
		return "Phentolamine allergy"
	case AllergyIntoleranceCode293991007:
		return "Prazosin allergy"
	case AllergyIntoleranceCode293992000:
		return "Terazosin allergy"
	case AllergyIntoleranceCode293993005:
		return "Nicotine allergy"
	case AllergyIntoleranceCode293994004:
		return "Calcium-channel blocker allergy"
	case AllergyIntoleranceCode293995003:
		return "Lidoflazine allergy"
	case AllergyIntoleranceCode293996002:
		return "Nifedipine allergy"
	case AllergyIntoleranceCode293997006:
		return "Prenylamine allergy"
	case AllergyIntoleranceCode293998001:
		return "Isradipine allergy"
	case AllergyIntoleranceCode293999009:
		return "Felodipine allergy"
	case AllergyIntoleranceCode294000006:
		return "Lacidipine allergy"
	case AllergyIntoleranceCode294001005:
		return "Nimodipine allergy"
	case AllergyIntoleranceCode294002003:
		return "Amlodipine allergy"
	case AllergyIntoleranceCode294003008:
		return "Diltiazem allergy"
	case AllergyIntoleranceCode294004002:
		return "Nicardipine allergy"
	case AllergyIntoleranceCode294005001:
		return "Verapamil allergy"
	case AllergyIntoleranceCode294007009:
		return "Pilocarpine allergy"
	case AllergyIntoleranceCode294009007:
		return "Methacholine allergy"
	case AllergyIntoleranceCode294011003:
		return "Physostigmine allergy"
	case AllergyIntoleranceCode294012005:
		return "Demecarium allergy"
	case AllergyIntoleranceCode294013000:
		return "Distigmine allergy"
	case AllergyIntoleranceCode294014006:
		return "Ecothiopate allergy"
	case AllergyIntoleranceCode294015007:
		return "Edrophonium allergy"
	case AllergyIntoleranceCode294016008:
		return "Pyridostigmine allergy"
	case AllergyIntoleranceCode294017004:
		return "Neostigmine allergy"
	case AllergyIntoleranceCode294018009:
		return "Bethanechol allergy"
	case AllergyIntoleranceCode294019001:
		return "Carbachol allergy"
	case AllergyIntoleranceCode294023009:
		return "Pseudoephedrine allergy"
	case AllergyIntoleranceCode294024003:
		return "Alpha-adrenoceptor agonist allergy"
	case AllergyIntoleranceCode294026001:
		return "Metaraminol allergy"
	case AllergyIntoleranceCode294027005:
		return "Methoxamine allergy"
	case AllergyIntoleranceCode294028000:
		return "Naphazoline allergy"
	case AllergyIntoleranceCode294029008:
		return "Norepinephrine allergy"
	case AllergyIntoleranceCode294030003:
		return "Phenylephrine allergy"
	case AllergyIntoleranceCode294031004:
		return "Xylometazoline allergy"
	case AllergyIntoleranceCode294033001:
		return "Selective beta-2 adrenoceptor stimulants allergy"
	case AllergyIntoleranceCode294035008:
		return "Pirbuterol allergy"
	case AllergyIntoleranceCode294036009:
		return "Salmeterol allergy"
	case AllergyIntoleranceCode294037000:
		return "Salbutamol allergy"
	case AllergyIntoleranceCode294038005:
		return "Bambuterol allergy"
	case AllergyIntoleranceCode294039002:
		return "Fenoterol allergy"
	case AllergyIntoleranceCode294040000:
		return "Orciprenaline allergy"
	case AllergyIntoleranceCode294041001:
		return "Reproterol allergy"
	case AllergyIntoleranceCode294042008:
		return "Rimiterol allergy"
	case AllergyIntoleranceCode294043003:
		return "Ritodrine allergy"
	case AllergyIntoleranceCode294044009:
		return "Terbutaline allergy"
	case AllergyIntoleranceCode294045005:
		return "Tulobuterol allergy"
	case AllergyIntoleranceCode294047002:
		return "Dobutamine allergy"
	case AllergyIntoleranceCode294048007:
		return "Dopexamine allergy"
	case AllergyIntoleranceCode294050004:
		return "Isoprenaline allergy"
	case AllergyIntoleranceCode294055009:
		return "Methyldopa allergy"
	case AllergyIntoleranceCode294057001:
		return "Apraclonidine allergy"
	case AllergyIntoleranceCode294058006:
		return "Clonidine allergy"
	case AllergyIntoleranceCode294059003:
		return "Lofexidine allergy"
	case AllergyIntoleranceCode294060008:
		return "Dipivefrine allergy"
	case AllergyIntoleranceCode294061007:
		return "Dopamine allergy"
	case AllergyIntoleranceCode294062000:
		return "Ephedrine allergy"
	case AllergyIntoleranceCode294063005:
		return "Oxymetazoline allergy"
	case AllergyIntoleranceCode294064004:
		return "Xamoterol allergy"
	case AllergyIntoleranceCode294067006:
		return "Belladonna alkaloids allergy"
	case AllergyIntoleranceCode294069009:
		return "Biperiden allergy"
	case AllergyIntoleranceCode294073007:
		return "Tropicamide allergy"
	case AllergyIntoleranceCode294074001:
		return "Hyoscine allergy"
	case AllergyIntoleranceCode294076004:
		return "Atropine allergy"
	case AllergyIntoleranceCode294077008:
		return "Benzhexol allergy"
	case AllergyIntoleranceCode294078003:
		return "Benztropine allergy"
	case AllergyIntoleranceCode294079006:
		return "Cyclopentolate allergy"
	case AllergyIntoleranceCode294080009:
		return "Glycopyrronium allergy"
	case AllergyIntoleranceCode294081008:
		return "Homatropine allergy"
	case AllergyIntoleranceCode294082001:
		return "Ipratropium allergy"
	case AllergyIntoleranceCode294083006:
		return "Methixene allergy"
	case AllergyIntoleranceCode294084000:
		return "Orphenadrine allergy"
	case AllergyIntoleranceCode294087007:
		return "Oxitropium allergy"
	case AllergyIntoleranceCode294088002:
		return "Oxybutynin allergy"
	case AllergyIntoleranceCode294089005:
		return "Procyclidine allergy"
	case AllergyIntoleranceCode294091002:
		return "Dornase alfa allergy"
	case AllergyIntoleranceCode294093004:
		return "Tyloxapol allergy"
	case AllergyIntoleranceCode294095006:
		return "Carbocisteine allergy"
	case AllergyIntoleranceCode294096007:
		return "Methylcysteine allergy"
	case AllergyIntoleranceCode294097003:
		return "Acetylcysteine allergy"
	case AllergyIntoleranceCode294099000:
		return "Nikethamide allergy"
	case AllergyIntoleranceCode294100008:
		return "Ethamivan allergy"
	case AllergyIntoleranceCode294101007:
		return "Doxapram allergy"
	case AllergyIntoleranceCode294103005:
		return "Beractant allergy"
	case AllergyIntoleranceCode294105003:
		return "Pumactant allergy"
	case AllergyIntoleranceCode294106002:
		return "Colfosceril allergy"
	case AllergyIntoleranceCode294109009:
		return "H1 antihistamine allergy"
	case AllergyIntoleranceCode294111000:
		return "Astemizole allergy"
	case AllergyIntoleranceCode294112007:
		return "Terfenadine allergy"
	case AllergyIntoleranceCode294113002:
		return "Acrivastine allergy"
	case AllergyIntoleranceCode294114008:
		return "Loratadine allergy"
	case AllergyIntoleranceCode294115009:
		return "Azelastine allergy"
	case AllergyIntoleranceCode294116005:
		return "Cetirizine allergy"
	case AllergyIntoleranceCode294118006:
		return "Clemastine allergy"
	case AllergyIntoleranceCode294119003:
		return "Mebhydrolin allergy"
	case AllergyIntoleranceCode294120009:
		return "Mequitazine allergy"
	case AllergyIntoleranceCode294121008:
		return "Oxatomide allergy"
	case AllergyIntoleranceCode294122001:
		return "Cyclizine allergy"
	case AllergyIntoleranceCode294123006:
		return "Dimenhydrinate allergy"
	case AllergyIntoleranceCode294125004:
		return "Antazoline allergy"
	case AllergyIntoleranceCode294126003:
		return "Promethazine allergy"
	case AllergyIntoleranceCode294127007:
		return "Azatadine allergy"
	case AllergyIntoleranceCode294128002:
		return "Brompheniramine allergy"
	case AllergyIntoleranceCode294129005:
		return "Chlorpheniramine allergy"
	case AllergyIntoleranceCode294130000:
		return "Cinnarizine allergy"
	case AllergyIntoleranceCode294131001:
		return "Cyproheptadine allergy"
	case AllergyIntoleranceCode294132008:
		return "Dimethindene allergy"
	case AllergyIntoleranceCode294133003:
		return "Diphenhydramine allergy"
	case AllergyIntoleranceCode294134009:
		return "Diphenylpyraline allergy"
	case AllergyIntoleranceCode294135005:
		return "Hydroxyzine allergy"
	case AllergyIntoleranceCode294136006:
		return "Mepyramine allergy"
	case AllergyIntoleranceCode294137002:
		return "Phenindamine allergy"
	case AllergyIntoleranceCode294138007:
		return "Pheniramine allergy"
	case AllergyIntoleranceCode294139004:
		return "Triprolidine allergy"
	case AllergyIntoleranceCode294140002:
		return "Trimeprazine allergy"
	case AllergyIntoleranceCode294142005:
		return "Nedocromil allergy"
	case AllergyIntoleranceCode294144006:
		return "Ketotifen allergy"
	case AllergyIntoleranceCode294145007:
		return "Lodoxamide allergy"
	case AllergyIntoleranceCode294148009:
		return "Isoaminile allergy"
	case AllergyIntoleranceCode294152009:
		return "Noscapine allergy"
	case AllergyIntoleranceCode294153004:
		return "Pholcodine allergy"
	case AllergyIntoleranceCode294157003:
		return "Xanthine allergy"
	case AllergyIntoleranceCode294158008:
		return "Aminophylline allergy"
	case AllergyIntoleranceCode294160005:
		return "Theophylline allergy"
	case AllergyIntoleranceCode294168003:
		return "Calamine allergy"
	case AllergyIntoleranceCode294169006:
		return "Coal tar allergy"
	case AllergyIntoleranceCode294172004:
		return "Bufexamac allergy"
	case AllergyIntoleranceCode294173009:
		return "Dithranol allergy"
	case AllergyIntoleranceCode294177005:
		return "Ichthammol allergy"
	case AllergyIntoleranceCode294178000:
		return "Calcipotriol allergy"
	case AllergyIntoleranceCode294180006:
		return "Azelaic acid allergy"
	case AllergyIntoleranceCode294182003:
		return "Podophyllum resin allergy"
	case AllergyIntoleranceCode294183008:
		return "Podophyllotoxin allergy"
	case AllergyIntoleranceCode294189007:
		return "Surgical tissue adhesive allergy"
	case AllergyIntoleranceCode294190003:
		return "Enbucrilate allergy"
	case AllergyIntoleranceCode294191004:
		return "Collodion allergy"
	case AllergyIntoleranceCode294200004:
		return "Crotamiton allergy"
	case AllergyIntoleranceCode294203002:
		return "Benzoyl peroxide allergy"
	case AllergyIntoleranceCode294204008:
		return "Silver nitrate allergy"
	case AllergyIntoleranceCode294206005:
		return "Gamolenic acid allergy"
	case AllergyIntoleranceCode294207001:
		return "Retinoid allergy"
	case AllergyIntoleranceCode294208006:
		return "Etretinate allergy"
	case AllergyIntoleranceCode294209003:
		return "Acitretin allergy"
	case AllergyIntoleranceCode294210008:
		return "Tretinoin allergy"
	case AllergyIntoleranceCode294211007:
		return "Isotretinoin allergy"
	case AllergyIntoleranceCode294214004:
		return "Colchicum alkaloid allergy"
	case AllergyIntoleranceCode294215003:
		return "Colchicine allergy"
	case AllergyIntoleranceCode294217006:
		return "Probenecid allergy"
	case AllergyIntoleranceCode294218001:
		return "Sulfinpyrazone allergy"
	case AllergyIntoleranceCode294219009:
		return "Xanthine oxidase inhibitor allergy"
	case AllergyIntoleranceCode294220003:
		return "Allopurinol allergy"
	case AllergyIntoleranceCode294224007:
		return "Suxamethonium allergy"
	case AllergyIntoleranceCode294225008:
		return "Non-depolarizing muscle relaxant allergy"
	case AllergyIntoleranceCode294226009:
		return "Mivacurium allergy"
	case AllergyIntoleranceCode294227000:
		return "Alcuronium allergy"
	case AllergyIntoleranceCode294228005:
		return "Atracurium allergy"
	case AllergyIntoleranceCode294229002:
		return "Gallamine allergy"
	case AllergyIntoleranceCode294230007:
		return "Pancuronium allergy"
	case AllergyIntoleranceCode294231006:
		return "Tubocurarine allergy"
	case AllergyIntoleranceCode294232004:
		return "Vecuronium allergy"
	case AllergyIntoleranceCode294233009:
		return "Rocuronium allergy"
	case AllergyIntoleranceCode294234003:
		return "Baclofen allergy"
	case AllergyIntoleranceCode294235002:
		return "Carisoprodol allergy"
	case AllergyIntoleranceCode294236001:
		return "Methocarbamol allergy"
	case AllergyIntoleranceCode294237005:
		return "Dantrolene allergy"
	case AllergyIntoleranceCode294238000:
		return "Gold allergy"
	case AllergyIntoleranceCode294239008:
		return "Sodium aurothiomalate allergy"
	case AllergyIntoleranceCode294240005:
		return "Auranofin allergy"
	case AllergyIntoleranceCode294242002:
		return "Papaverine allergy"
	case AllergyIntoleranceCode294243007:
		return "Flavoxate allergy"
	case AllergyIntoleranceCode294245000:
		return "Mifepristone allergy"
	case AllergyIntoleranceCode294246004:
		return "Non-ionic surfactant allergy"
	case AllergyIntoleranceCode294247008:
		return "Nonoxinol allergy"
	case AllergyIntoleranceCode294248003:
		return "Octoxinol allergy"
	case AllergyIntoleranceCode294249006:
		return "Prostaglandin allergy"
	case AllergyIntoleranceCode294250006:
		return "A series prostaglandin allergy"
	case AllergyIntoleranceCode294252003:
		return "E series prostaglandin allergy"
	case AllergyIntoleranceCode294253008:
		return "Dinoprostone allergy"
	case AllergyIntoleranceCode294254002:
		return "Gemeprost allergy"
	case AllergyIntoleranceCode294255001:
		return "Alprostadil allergy"
	case AllergyIntoleranceCode294256000:
		return "F series prostaglandin allergy"
	case AllergyIntoleranceCode294257009:
		return "Dinoprost allergy"
	case AllergyIntoleranceCode294258004:
		return "Carboprost allergy"
	case AllergyIntoleranceCode294259007:
		return "I series prostaglandin allergy"
	case AllergyIntoleranceCode294260002:
		return "Epoprostenol allergy"
	case AllergyIntoleranceCode294261003:
		return "Terpenes allergy"
	case AllergyIntoleranceCode294264006:
		return "Ipecacuanha allergy"
	case AllergyIntoleranceCode294265007:
		return "Charcoal-activated allergy"
	case AllergyIntoleranceCode294266008:
		return "Sodium nitrite allergy"
	case AllergyIntoleranceCode294268009:
		return "Digoxin-specific-antibody allergy"
	case AllergyIntoleranceCode294269001:
		return "Mesna allergy"
	case AllergyIntoleranceCode294270000:
		return "Benzodiazepine antagonist allergy"
	case AllergyIntoleranceCode294271001:
		return "Flumazenil allergy"
	case AllergyIntoleranceCode294273003:
		return "Pralidoxime allergy"
	case AllergyIntoleranceCode294275005:
		return "Opioid antagonist allergy"
	case AllergyIntoleranceCode294276006:
		return "Naltrexone allergy"
	case AllergyIntoleranceCode294277002:
		return "Naloxone allergy"
	case AllergyIntoleranceCode294278007:
		return "Protamine allergy"
	case AllergyIntoleranceCode294280001:
		return "Allergy to Fullers earth"
	case AllergyIntoleranceCode294281002:
		return "Allergy to bentonite"
	case AllergyIntoleranceCode294283004:
		return "Dimercaprol allergy"
	case AllergyIntoleranceCode294284005:
		return "Desferrioxamine allergy"
	case AllergyIntoleranceCode294285006:
		return "Edetate allergy"
	case AllergyIntoleranceCode294290009:
		return "Trientine allergy"
	case AllergyIntoleranceCode294291008:
		return "Penicillamine allergy"
	case AllergyIntoleranceCode294298002:
		return "Glycine allergy"
	case AllergyIntoleranceCode294299005:
		return "Dialysis fluid allergy"
	case AllergyIntoleranceCode294306008:
		return "Dimethyl-ether propane allergy"
	case AllergyIntoleranceCode294316000:
		return "Olive oil allergy"
	case AllergyIntoleranceCode294317009:
		return "Arachis oil allergy"
	case AllergyIntoleranceCode294318004:
		return "Castor oil allergy"
	case AllergyIntoleranceCode294320001:
		return "Glycerol allergy"
	case AllergyIntoleranceCode294324005:
		return "Paraffin allergy"
	case AllergyIntoleranceCode294327003:
		return "Liquid paraffin allergy"
	case AllergyIntoleranceCode294328008:
		return "Silicone allergy"
	case AllergyIntoleranceCode294329000:
		return "Dimethicone allergy"
	case AllergyIntoleranceCode294330005:
		return "Wool alcohol allergy"
	case AllergyIntoleranceCode294332002:
		return "Polyvinyl alcohol allergy"
	case AllergyIntoleranceCode294333007:
		return "Carbomer-940 allergy"
	case AllergyIntoleranceCode294335000:
		return "Hypromellose allergy"
	case AllergyIntoleranceCode294337008:
		return "Hydroxyethylcellulose allergy"
	case AllergyIntoleranceCode294339006:
		return "Carmellose allergy"
	case AllergyIntoleranceCode294341007:
		return "Flucytosine allergy"
	case AllergyIntoleranceCode294342000:
		return "Terbinafine allergy"
	case AllergyIntoleranceCode294343005:
		return "Nitrophenol allergy"
	case AllergyIntoleranceCode294344004:
		return "Tolnaftate allergy"
	case AllergyIntoleranceCode294346002:
		return "Amorolfine allergy"
	case AllergyIntoleranceCode294348001:
		return "Griseofulvin allergy"
	case AllergyIntoleranceCode294349009:
		return "Amphotericin allergy"
	case AllergyIntoleranceCode294350009:
		return "Natamycin allergy"
	case AllergyIntoleranceCode294351008:
		return "Nystatin allergy"
	case AllergyIntoleranceCode294354000:
		return "Allergy to undecenoate"
	case AllergyIntoleranceCode294356003:
		return "Clotrimazole allergy"
	case AllergyIntoleranceCode294357007:
		return "Fenticonazole allergy"
	case AllergyIntoleranceCode294358002:
		return "Tioconazole allergy"
	case AllergyIntoleranceCode294359005:
		return "Econazole allergy"
	case AllergyIntoleranceCode294360000:
		return "Isoconazole allergy"
	case AllergyIntoleranceCode294361001:
		return "Sulconazole allergy"
	case AllergyIntoleranceCode294362008:
		return "Ketoconazole allergy"
	case AllergyIntoleranceCode294363003:
		return "Miconazole allergy"
	case AllergyIntoleranceCode294365005:
		return "Fluconazole allergy"
	case AllergyIntoleranceCode294366006:
		return "Itraconazole allergy"
	case AllergyIntoleranceCode294368007:
		return "Inosine pranobex allergy"
	case AllergyIntoleranceCode294369004:
		return "Zidovudine allergy"
	case AllergyIntoleranceCode294370003:
		return "Ganciclovir allergy"
	case AllergyIntoleranceCode294371004:
		return "Famciclovir allergy"
	case AllergyIntoleranceCode294372006:
		return "Didanosine allergy"
	case AllergyIntoleranceCode294373001:
		return "Zalcitabine allergy"
	case AllergyIntoleranceCode294374007:
		return "Valaciclovir allergy"
	case AllergyIntoleranceCode294375008:
		return "Interferons allergy"
	case AllergyIntoleranceCode294376009:
		return "Human interferon gamma-1b allergy"
	case AllergyIntoleranceCode294377000:
		return "Interferon-A-2a allergy"
	case AllergyIntoleranceCode294378005:
		return "Interferon-A-2b allergy"
	case AllergyIntoleranceCode294379002:
		return "Interferon-A-N1 allergy"
	case AllergyIntoleranceCode294380004:
		return "Tribavirin allergy"
	case AllergyIntoleranceCode294381000:
		return "Trifluorothymidine allergy"
	case AllergyIntoleranceCode294382007:
		return "Foscarnet allergy"
	case AllergyIntoleranceCode294383002:
		return "Vidarabine allergy"
	case AllergyIntoleranceCode294384008:
		return "Aciclovir allergy"
	case AllergyIntoleranceCode294385009:
		return "Idoxuridine allergy"
	case AllergyIntoleranceCode294388006:
		return "Pyrimethamine allergy"
	case AllergyIntoleranceCode294390007:
		return "Amodiaquine allergy"
	case AllergyIntoleranceCode294391006:
		return "Primaquine allergy"
	case AllergyIntoleranceCode294392004:
		return "Mefloquine allergy"
	case AllergyIntoleranceCode294393009:
		return "Hydroxychloroquine allergy"
	case AllergyIntoleranceCode294394003:
		return "Chloroquine allergy"
	case AllergyIntoleranceCode294396001:
		return "Proguanil allergy"
	case AllergyIntoleranceCode294398000:
		return "Quinine allergy"
	case AllergyIntoleranceCode294399008:
		return "Halofantrine allergy"
	case AllergyIntoleranceCode294400001:
		return "Mepacrine allergy"
	case AllergyIntoleranceCode294404005:
		return "Acetic acid allergy"
	case AllergyIntoleranceCode294405006:
		return "Hydrargaphen allergy"
	case AllergyIntoleranceCode294406007:
		return "Polynoxylin allergy"
	case AllergyIntoleranceCode294407003:
		return "Hexetidine allergy"
	case AllergyIntoleranceCode294408008:
		return "Sodium perborate allergy"
	case AllergyIntoleranceCode294413007:
		return "Thymol allergy"
	case AllergyIntoleranceCode294415000:
		return "Chloroxylenol allergy"
	case AllergyIntoleranceCode294416004:
		return "Hexachlorophane allergy"
	case AllergyIntoleranceCode294417008:
		return "Triclosan allergy"
	case AllergyIntoleranceCode294418003:
		return "Phenol allergy"
	case AllergyIntoleranceCode294421001:
		return "Industrial methylated spirit allergy"
	case AllergyIntoleranceCode294423003:
		return "Glutaraldehyde allergy"
	case AllergyIntoleranceCode294425005:
		return "Noxythiolin allergy"
	case AllergyIntoleranceCode294426006:
		return "Formaldehyde allergy"
	case AllergyIntoleranceCode294431008:
		return "Chlorhexidine allergy"
	case AllergyIntoleranceCode294433006:
		return "Borate allergy"
	case AllergyIntoleranceCode294434000:
		return "Boric acid allergy"
	case AllergyIntoleranceCode294436003:
		return "Quaternary ammonium surfactant allergy"
	case AllergyIntoleranceCode294437007:
		return "Cetrimide allergy"
	case AllergyIntoleranceCode294438002:
		return "Benzalkonium allergy"
	case AllergyIntoleranceCode294439005:
		return "Domiphen allergy"
	case AllergyIntoleranceCode294440007:
		return "Quaternary pyridinium surfactant allergy"
	case AllergyIntoleranceCode294441006:
		return "Cetylpyridinium allergy"
	case AllergyIntoleranceCode294442004:
		return "Quaternary quinolinium surfactant allergy"
	case AllergyIntoleranceCode294443009:
		return "Dequalinium allergy"
	case AllergyIntoleranceCode294447005:
		return "Crystal violet allergy"
	case AllergyIntoleranceCode294448000:
		return "Brilliant green allergy"
	case AllergyIntoleranceCode294449008:
		return "Hydrogen peroxide allergy"
	case AllergyIntoleranceCode294451007:
		return "Piperazine allergy"
	case AllergyIntoleranceCode294452000:
		return "Pyrantel allergy"
	case AllergyIntoleranceCode294453005:
		return "Niclosamide allergy"
	case AllergyIntoleranceCode294455003:
		return "Bephenium allergy"
	case AllergyIntoleranceCode294456002:
		return "Diethylcarbamazine allergy"
	case AllergyIntoleranceCode294458001:
		return "Mebendazole allergy"
	case AllergyIntoleranceCode294459009:
		return "Albendazole allergy"
	case AllergyIntoleranceCode294460004:
		return "Thiabendazole allergy"
	case AllergyIntoleranceCode294462007:
		return "Aminoglycoside allergy"
	case AllergyIntoleranceCode294463002:
		return "Amikacin allergy"
	case AllergyIntoleranceCode294464008:
		return "Kanamycin allergy"
	case AllergyIntoleranceCode294465009:
		return "Netilmicin allergy"
	case AllergyIntoleranceCode294466005:
		return "Streptomycin allergy"
	case AllergyIntoleranceCode294467001:
		return "Framycetin allergy"
	case AllergyIntoleranceCode294468006:
		return "Neomycin allergy"
	case AllergyIntoleranceCode294469003:
		return "Gentamicin allergy"
	case AllergyIntoleranceCode294470002:
		return "Tobramycin allergy"
	case AllergyIntoleranceCode294471003:
		return "Clarithromycin allergy"
	case AllergyIntoleranceCode294472005:
		return "Azithromycin allergy"
	case AllergyIntoleranceCode294474006:
		return "Spectinomycin allergy"
	case AllergyIntoleranceCode294475007:
		return "Vancomycin allergy"
	case AllergyIntoleranceCode294476008:
		return "Teicoplanin allergy"
	case AllergyIntoleranceCode294477004:
		return "Trimethoprim allergy"
	case AllergyIntoleranceCode294478009:
		return "Nitrofurantoin allergy"
	case AllergyIntoleranceCode294480003:
		return "Mupirocin allergy"
	case AllergyIntoleranceCode294481004:
		return "Nitrofurazone allergy"
	case AllergyIntoleranceCode294482006:
		return "Fusidic acid allergy"
	case AllergyIntoleranceCode294484007:
		return "Acrosoxacin allergy"
	case AllergyIntoleranceCode294485008:
		return "Cinoxacin allergy"
	case AllergyIntoleranceCode294486009:
		return "Nalidixic acid allergy"
	case AllergyIntoleranceCode294487000:
		return "Ciprofloxacin allergy"
	case AllergyIntoleranceCode294488005:
		return "Enoxacin allergy"
	case AllergyIntoleranceCode294489002:
		return "Ofloxacin allergy"
	case AllergyIntoleranceCode294490006:
		return "Norfloxacin allergy"
	case AllergyIntoleranceCode294491005:
		return "Temafloxacin allergy"
	case AllergyIntoleranceCode294494002:
		return "Benethamine penicillin allergy"
	case AllergyIntoleranceCode294496000:
		return "Phenethicillin allergy"
	case AllergyIntoleranceCode294497009:
		return "Phenoxymethylpenicillin allergy"
	case AllergyIntoleranceCode294499007:
		return "Benzylpenicillin allergy"
	case AllergyIntoleranceCode294501004:
		return "Cloxacillin allergy"
	case AllergyIntoleranceCode294502006:
		return "Flucloxacillin allergy"
	case AllergyIntoleranceCode294503001:
		return "Methicillin allergy"
	case AllergyIntoleranceCode294505008:
		return "Amoxycillin allergy"
	case AllergyIntoleranceCode294506009:
		return "Ampicillin allergy"
	case AllergyIntoleranceCode294507000:
		return "Ciclacillin allergy"
	case AllergyIntoleranceCode294508005:
		return "Mezlocillin allergy"
	case AllergyIntoleranceCode294509002:
		return "Pivampicillin allergy"
	case AllergyIntoleranceCode294510007:
		return "Carbenicillin allergy"
	case AllergyIntoleranceCode294511006:
		return "Bacampicillin allergy"
	case AllergyIntoleranceCode294512004:
		return "Talampicillin allergy"
	case AllergyIntoleranceCode294514003:
		return "Temocillin allergy"
	case AllergyIntoleranceCode294515002:
		return "Piperacillin allergy"
	case AllergyIntoleranceCode294516001:
		return "Azlocillin allergy"
	case AllergyIntoleranceCode294517005:
		return "Ticarcillin allergy"
	case AllergyIntoleranceCode294518000:
		return "Carfecillin allergy"
	case AllergyIntoleranceCode294519008:
		return "Mecillinam allergy"
	case AllergyIntoleranceCode294520002:
		return "Pivmecillinam allergy"
	case AllergyIntoleranceCode294528009:
		return "Polymyxins allergy"
	case AllergyIntoleranceCode294529001:
		return "Colistin allergy"
	case AllergyIntoleranceCode294530006:
		return "Polymyxin B allergy"
	case AllergyIntoleranceCode294531005:
		return "Carbapenem allergy"
	case AllergyIntoleranceCode294532003:
		return "Cephalosporin allergy"
	case AllergyIntoleranceCode294534002:
		return "Cefadroxil allergy"
	case AllergyIntoleranceCode294535001:
		return "Cephalexin allergy"
	case AllergyIntoleranceCode294536000:
		return "Cephalothin allergy"
	case AllergyIntoleranceCode294537009:
		return "Cephazolin allergy"
	case AllergyIntoleranceCode294538004:
		return "Cephradine allergy"
	case AllergyIntoleranceCode294539007:
		return "Latamoxef allergy"
	case AllergyIntoleranceCode294541008:
		return "Cefaclor allergy"
	case AllergyIntoleranceCode294542001:
		return "Cefuroxime allergy"
	case AllergyIntoleranceCode294543006:
		return "Cephamandole allergy"
	case AllergyIntoleranceCode294545004:
		return "Cefotaxime allergy"
	case AllergyIntoleranceCode294546003:
		return "Ceftazidime allergy"
	case AllergyIntoleranceCode294547007:
		return "Ceftizoxime allergy"
	case AllergyIntoleranceCode294548002:
		return "Cefixime allergy"
	case AllergyIntoleranceCode294549005:
		return "Cefodizime allergy"
	case AllergyIntoleranceCode294550005:
		return "Cefpodoxime allergy"
	case AllergyIntoleranceCode294551009:
		return "Ceftriaxone allergy"
	case AllergyIntoleranceCode294552002:
		return "Ceftibuten allergy"
	case AllergyIntoleranceCode294554001:
		return "Cefsulodin allergy"
	case AllergyIntoleranceCode294556004:
		return "Cefpirome allergy"
	case AllergyIntoleranceCode294557008:
		return "Cephamycin allergy"
	case AllergyIntoleranceCode294558003:
		return "Cefoxitin allergy"
	case AllergyIntoleranceCode294559006:
		return "Fosfomycin allergy"
	case AllergyIntoleranceCode294561002:
		return "Clindamycin allergy"
	case AllergyIntoleranceCode294562009:
		return "Lincomycin allergy"
	case AllergyIntoleranceCode294563004:
		return "Mandelic acid allergy"
	case AllergyIntoleranceCode294564005:
		return "Monobactam allergy"
	case AllergyIntoleranceCode294565006:
		return "Aztreonam allergy"
	case AllergyIntoleranceCode294566007:
		return "Nitroimidazole allergy"
	case AllergyIntoleranceCode294567003:
		return "Metronidazole allergy"
	case AllergyIntoleranceCode294568008:
		return "Tinidazole allergy"
	case AllergyIntoleranceCode294569000:
		return "Nimorazole allergy"
	case AllergyIntoleranceCode294570004:
		return "Calcium sulfaloxate allergy"
	case AllergyIntoleranceCode294571000:
		return "Phthalylsulfathiazole allergy"
	case AllergyIntoleranceCode294572007:
		return "Sulfametopyrazine allergy"
	case AllergyIntoleranceCode294573002:
		return "Sulfadiazine allergy"
	case AllergyIntoleranceCode294574008:
		return "Sulfadimethoxine allergy"
	case AllergyIntoleranceCode294575009:
		return "Sulfadimidine allergy"
	case AllergyIntoleranceCode294576005:
		return "Sulfafurazole allergy"
	case AllergyIntoleranceCode294577001:
		return "Sulfaguanidine allergy"
	case AllergyIntoleranceCode294578006:
		return "Sulfaurea allergy"
	case AllergyIntoleranceCode294579003:
		return "Mafenide allergy"
	case AllergyIntoleranceCode294582008:
		return "Sulfacetamide allergy"
	case AllergyIntoleranceCode294584009:
		return "Clomocycline sodium allergy"
	case AllergyIntoleranceCode294585005:
		return "Doxycycline allergy"
	case AllergyIntoleranceCode294586006:
		return "Lymecycline allergy"
	case AllergyIntoleranceCode294587002:
		return "Minocycline allergy"
	case AllergyIntoleranceCode294588007:
		return "Oxytetracycline allergy"
	case AllergyIntoleranceCode294590008:
		return "Chlortetracycline allergy"
	case AllergyIntoleranceCode294591007:
		return "Demeclocycline allergy"
	case AllergyIntoleranceCode294592000:
		return "Tetracycline allergy"
	case AllergyIntoleranceCode294593005:
		return "Chloramphenicol allergy"
	case AllergyIntoleranceCode294596002:
		return "Atovaquone allergy"
	case AllergyIntoleranceCode294598001:
		return "Sodium stibogluconate allergy"
	case AllergyIntoleranceCode294600007:
		return "Pentamidine allergy"
	case AllergyIntoleranceCode294602004:
		return "Diloxanide allergy"
	case AllergyIntoleranceCode294604003:
		return "Clioquinol allergy"
	case AllergyIntoleranceCode294607005:
		return "Pyrazinamide allergy"
	case AllergyIntoleranceCode294609008:
		return "Capreomycin allergy"
	case AllergyIntoleranceCode294610003:
		return "Cycloserine allergy"
	case AllergyIntoleranceCode294611004:
		return "Rifampicin allergy"
	case AllergyIntoleranceCode294612006:
		return "Rifabutin allergy"
	case AllergyIntoleranceCode294614007:
		return "Isoniazid allergy"
	case AllergyIntoleranceCode294615008:
		return "Ethambutolol allergy"
	case AllergyIntoleranceCode294617000:
		return "Dapsone allergy"
	case AllergyIntoleranceCode294618005:
		return "Clofazimine allergy"
	case AllergyIntoleranceCode294620008:
		return "Benzyl benzoate allergy"
	case AllergyIntoleranceCode294621007:
		return "Monosulfiram allergy"
	case AllergyIntoleranceCode294623005:
		return "Carbaryl allergy"
	case AllergyIntoleranceCode294625003:
		return "Lindane allergy"
	case AllergyIntoleranceCode294627006:
		return "Malathion allergy"
	case AllergyIntoleranceCode294629009:
		return "Phenothrin allergy"
	case AllergyIntoleranceCode294630004:
		return "Permethrin allergy"
	case AllergyIntoleranceCode294633002:
		return "Human immunoglobulin allergy"
	case AllergyIntoleranceCode294638006:
		return "Tetanus immunoglobulin allergy"
	case AllergyIntoleranceCode294639003:
		return "Varicella-zoster immunoglobulin allergy"
	case AllergyIntoleranceCode294667007:
		return "Clostridium botulinum toxin allergy"
	case AllergyIntoleranceCode294668002:
		return "Botulism antitoxin allergy"
	case AllergyIntoleranceCode294669005:
		return "Diphtheria antitoxin allergy"
	case AllergyIntoleranceCode294671005:
		return "Glucagon allergy"
	case AllergyIntoleranceCode294674002:
		return "Carbimazole allergy"
	case AllergyIntoleranceCode294676000:
		return "Propylthiouracil allergy"
	case AllergyIntoleranceCode294677009:
		return "Corticosteroids allergy"
	case AllergyIntoleranceCode294678004:
		return "Betamethasone allergy"
	case AllergyIntoleranceCode294679007:
		return "Hydrocortisone allergy"
	case AllergyIntoleranceCode294682002:
		return "Prednisone allergy"
	case AllergyIntoleranceCode294683007:
		return "Fluorometholone allergy"
	case AllergyIntoleranceCode294684001:
		return "Flunisolide allergy"
	case AllergyIntoleranceCode294685000:
		return "Desonide allergy"
	case AllergyIntoleranceCode294686004:
		return "Desoxymethasone allergy"
	case AllergyIntoleranceCode294687008:
		return "Fluocinonide allergy"
	case AllergyIntoleranceCode294688003:
		return "Fluocortolone allergy"
	case AllergyIntoleranceCode294689006:
		return "Flurandrenolone allergy"
	case AllergyIntoleranceCode294690002:
		return "Halcinonide allergy"
	case AllergyIntoleranceCode294691003:
		return "Alclometasone allergy"
	case AllergyIntoleranceCode294692005:
		return "Beclomethasone allergy"
	case AllergyIntoleranceCode294693000:
		return "Clobetasol allergy"
	case AllergyIntoleranceCode294694006:
		return "Clobetasone allergy"
	case AllergyIntoleranceCode294695007:
		return "Cortisone allergy"
	case AllergyIntoleranceCode294696008:
		return "Diflucortolone allergy"
	case AllergyIntoleranceCode294697004:
		return "Fluclorolone allergy"
	case AllergyIntoleranceCode294698009:
		return "Fludrocortisone allergy"
	case AllergyIntoleranceCode294699001:
		return "Fluocinolone allergy"
	case AllergyIntoleranceCode294700000:
		return "Fluticasone allergy"
	case AllergyIntoleranceCode294701001:
		return "Mometasone allergy"
	case AllergyIntoleranceCode294702008:
		return "Dexamethasone allergy"
	case AllergyIntoleranceCode294706006:
		return "Methylprednisolone allergy"
	case AllergyIntoleranceCode294707002:
		return "Prednisolone allergy"
	case AllergyIntoleranceCode294711008:
		return "Triamcinolone allergy"
	case AllergyIntoleranceCode294712001:
		return "Budesonide allergy"
	case AllergyIntoleranceCode294714000:
		return "Insulin allergy"
	case AllergyIntoleranceCode294717007:
		return "Insulin zinc suspension allergy"
	case AllergyIntoleranceCode294720004:
		return "Isophane insulin allergy"
	case AllergyIntoleranceCode294721000:
		return "Protamine zinc insulin allergy"
	case AllergyIntoleranceCode294723002:
		return "Allergy to human insulin"
	case AllergyIntoleranceCode294728006:
		return "Sulfonylurea allergy"
	case AllergyIntoleranceCode294729003:
		return "Acetohexamide allergy"
	case AllergyIntoleranceCode294730008:
		return "Chlorpropamide allergy"
	case AllergyIntoleranceCode294731007:
		return "Glibenclamide allergy"
	case AllergyIntoleranceCode294732000:
		return "Glibornuride allergy"
	case AllergyIntoleranceCode294733005:
		return "Gliclazide allergy"
	case AllergyIntoleranceCode294734004:
		return "Glipizide allergy"
	case AllergyIntoleranceCode294735003:
		return "Gliquidone allergy"
	case AllergyIntoleranceCode294736002:
		return "Glymidine allergy"
	case AllergyIntoleranceCode294737006:
		return "Tolazamide allergy"
	case AllergyIntoleranceCode294738001:
		return "Tolbutamide allergy"
	case AllergyIntoleranceCode294739009:
		return "Biguanide allergy"
	case AllergyIntoleranceCode294740006:
		return "Metformin allergy"
	case AllergyIntoleranceCode294741005:
		return "Guar gum allergy"
	case AllergyIntoleranceCode294742003:
		return "Acarbose allergy"
	case AllergyIntoleranceCode294745001:
		return "Progestogen allergy"
	case AllergyIntoleranceCode294746000:
		return "Allylestrenol allergy"
	case AllergyIntoleranceCode294747009:
		return "Dydrogesterone allergy"
	case AllergyIntoleranceCode294748004:
		return "Progesterone allergy"
	case AllergyIntoleranceCode294749007:
		return "Gestronol allergy"
	case AllergyIntoleranceCode294750007:
		return "Hydroxyprogesterone allergy"
	case AllergyIntoleranceCode294751006:
		return "Megestrol allergy"
	case AllergyIntoleranceCode294752004:
		return "Norethisterone allergy"
	case AllergyIntoleranceCode294754003:
		return "Levonorgestrel allergy"
	case AllergyIntoleranceCode294755002:
		return "Medroxyprogesterone allergy"
	case AllergyIntoleranceCode294757005:
		return "Anabolic steroids allergy"
	case AllergyIntoleranceCode294758000:
		return "Tibolone allergy"
	case AllergyIntoleranceCode294760003:
		return "Oxymetholone allergy"
	case AllergyIntoleranceCode294761004:
		return "Nandrolone allergy"
	case AllergyIntoleranceCode294762006:
		return "Stanozolol allergy"
	case AllergyIntoleranceCode294763001:
		return "Cyclofenil allergy"
	case AllergyIntoleranceCode294764007:
		return "Danazol allergy"
	case AllergyIntoleranceCode294765008:
		return "Gestrinone allergy"
	case AllergyIntoleranceCode294767000:
		return "Finasteride allergy"
	case AllergyIntoleranceCode294768005:
		return "Flutamide allergy"
	case AllergyIntoleranceCode294769002:
		return "Bicalutamide allergy"
	case AllergyIntoleranceCode294771002:
		return "Cyproterone allergy"
	case AllergyIntoleranceCode294773004:
		return "Androgen allergy"
	case AllergyIntoleranceCode294774005:
		return "Mesterolone allergy"
	case AllergyIntoleranceCode294775006:
		return "Methyltestosterone allergy"
	case AllergyIntoleranceCode294776007:
		return "Testosterone allergy"
	case AllergyIntoleranceCode294781003:
		return "Estrogen allergy"
	case AllergyIntoleranceCode294782005:
		return "Estradiol allergy"
	case AllergyIntoleranceCode294787004:
		return "Quinestradol allergy"
	case AllergyIntoleranceCode294788009:
		return "Quinestrol allergy"
	case AllergyIntoleranceCode294789001:
		return "Dienestrol allergy"
	case AllergyIntoleranceCode294792002:
		return "Mestranol allergy"
	case AllergyIntoleranceCode294793007:
		return "Ethinylestradiol allergy"
	case AllergyIntoleranceCode294794001:
		return "Estriol allergy"
	case AllergyIntoleranceCode294795000:
		return "Conjugated estrogen allergy"
	case AllergyIntoleranceCode294796004:
		return "Stilbestrol allergy"
	case AllergyIntoleranceCode294798003:
		return "Clomiphene allergy"
	case AllergyIntoleranceCode294799006:
		return "Ergoline drug allergy"
	case AllergyIntoleranceCode294800005:
		return "Cabergoline allergy"
	case AllergyIntoleranceCode294801009:
		return "Quinagolide allergy"
	case AllergyIntoleranceCode294803007:
		return "Liothyronine allergy"
	case AllergyIntoleranceCode294804001:
		return "Thyroxine allergy"
	case AllergyIntoleranceCode294807008:
		return "Desmopressin allergy"
	case AllergyIntoleranceCode294809006:
		return "Terlipressin allergy"
	case AllergyIntoleranceCode294810001:
		return "Vasopressin allergy"
	case AllergyIntoleranceCode294811002:
		return "Corticotrophic hormone allergy"
	case AllergyIntoleranceCode294813004:
		return "Tetracosactrin allergy"
	case AllergyIntoleranceCode294814005:
		return "Gonad regulating hormone allergy"
	case AllergyIntoleranceCode294815006:
		return "Gonadorelin allergy"
	case AllergyIntoleranceCode294816007:
		return "Nafarelin allergy"
	case AllergyIntoleranceCode294817003:
		return "Buserelin allergy"
	case AllergyIntoleranceCode294818008:
		return "Goserelin allergy"
	case AllergyIntoleranceCode294819000:
		return "Triptorelin allergy"
	case AllergyIntoleranceCode294820006:
		return "Gonadotrophic hormone allergy"
	case AllergyIntoleranceCode294821005:
		return "Leuprorelin allergy"
	case AllergyIntoleranceCode294823008:
		return "Oxytocin allergy"
	case AllergyIntoleranceCode294825001:
		return "Somatrophic hormone allergy"
	case AllergyIntoleranceCode294826000:
		return "Octreotide allergy"
	case AllergyIntoleranceCode294828004:
		return "Thyrotrophic hormone allergy"
	case AllergyIntoleranceCode294829007:
		return "Protirelin allergy"
	case AllergyIntoleranceCode294830002:
		return "Thyrotrophin allergy"
	case AllergyIntoleranceCode294833000:
		return "Biphosphonates allergy"
	case AllergyIntoleranceCode294838009:
		return "Calcium regulating hormone allergy"
	case AllergyIntoleranceCode294839001:
		return "Calcitonin allergy"
	case AllergyIntoleranceCode294840004:
		return "Salcatonin allergy"
	case AllergyIntoleranceCode294841000:
		return "Calcitonin (pork) allergy"
	case AllergyIntoleranceCode294844008:
		return "Epoetin alfa allergy"
	case AllergyIntoleranceCode294845009:
		return "Epoetin beta allergy"
	case AllergyIntoleranceCode294847001:
		return "Gelatin allergy"
	case AllergyIntoleranceCode294851004:
		return "Perfluorochemical allergy"
	case AllergyIntoleranceCode294855008:
		return "Antithrombin III allergy"
	case AllergyIntoleranceCode294865002:
		return "Factor VIII by-passing fraction products allergy"
	case AllergyIntoleranceCode294871008:
		return "Ancrod allergy"
	case AllergyIntoleranceCode294872001:
		return "Heparin allergy"
	case AllergyIntoleranceCode294873006:
		return "Enoxaparin allergy"
	case AllergyIntoleranceCode294874000:
		return "Dalteparin allergy"
	case AllergyIntoleranceCode294875004:
		return "Tinzaparin allergy"
	case AllergyIntoleranceCode294876003:
		return "Heparinoid allergy"
	case AllergyIntoleranceCode294880008:
		return "Nicoumalone allergy"
	case AllergyIntoleranceCode294881007:
		return "Warfarin allergy"
	case AllergyIntoleranceCode294883005:
		return "Phenindione allergy"
	case AllergyIntoleranceCode294885003:
		return "Ethamsylate allergy"
	case AllergyIntoleranceCode294886002:
		return "Thromboplastin allergy"
	case AllergyIntoleranceCode294887006:
		return "Tranexamic acid allergy"
	case AllergyIntoleranceCode294888001:
		return "Collagen allergy"
	case AllergyIntoleranceCode294889009:
		return "Aprotinin allergy"
	case AllergyIntoleranceCode294893003:
		return "Polysaccharide iron complex allergy"
	case AllergyIntoleranceCode294894009:
		return "Sodium ironedetate allergy"
	case AllergyIntoleranceCode294896006:
		return "Iron sorbitol allergy"
	case AllergyIntoleranceCode294897002:
		return "Ferrous salt allergy"
	case AllergyIntoleranceCode294898007:
		return "Ferrous fumarate allergy"
	case AllergyIntoleranceCode294899004:
		return "Ferrous gluconate allergy"
	case AllergyIntoleranceCode294900009:
		return "Ferrous glycine sulfate allergy"
	case AllergyIntoleranceCode294901008:
		return "Ferrous succinate allergy"
	case AllergyIntoleranceCode294902001:
		return "Ferrous sulfate allergy"
	case AllergyIntoleranceCode294903006:
		return "Ferrous phosphate allergy"
	case AllergyIntoleranceCode294912008:
		return "L-Carnitine allergy"
	case AllergyIntoleranceCode294913003:
		return "Iodine compounds allergy"
	case AllergyIntoleranceCode294915005:
		return "Iodophore allergy"
	case AllergyIntoleranceCode294916006:
		return "Povidone iodine allergy"
	case AllergyIntoleranceCode294923007:
		return "Vitamin A allergy"
	case AllergyIntoleranceCode294924001:
		return "Allergy to vitamin D and/or vitamin D derivative"
	case AllergyIntoleranceCode294925000:
		return "Allergy to vitamin K and/or vitamin K derivative (finding)"
	case AllergyIntoleranceCode294926004:
		return "Alpha-tocopheryl allergy"
	case AllergyIntoleranceCode294929006:
		return "Nicotinic acid allergy"
	case AllergyIntoleranceCode294930001:
		return "Folic acid allergy"
	case AllergyIntoleranceCode294931002:
		return "Folinic acid allergy"
	case AllergyIntoleranceCode294933004:
		return "Hydroxocobalamin allergy"
	case AllergyIntoleranceCode294934005:
		return "Cyanocobalamin allergy"
	case AllergyIntoleranceCode294937003:
		return "Inositol allergy"
	case AllergyIntoleranceCode294940003:
		return "Ascorbic acid allergy"
	case AllergyIntoleranceCode294945008:
		return "Sodium polystyrene sulfonate allergy"
	case AllergyIntoleranceCode294951003:
		return "Fluoride allergy"
	case AllergyIntoleranceCode294956008:
		return "Gemfibrozil allergy"
	case AllergyIntoleranceCode294957004:
		return "Probucol allergy"
	case AllergyIntoleranceCode294958009:
		return "Acipimox allergy"
	case AllergyIntoleranceCode294961005:
		return "Colestipol allergy"
	case AllergyIntoleranceCode294962003:
		return "Cholestyramine allergy"
	case AllergyIntoleranceCode294964002:
		return "Bezafibrate allergy"
	case AllergyIntoleranceCode294965001:
		return "Clofibrate allergy"
	case AllergyIntoleranceCode294966000:
		return "Fenofibrate allergy"
	case AllergyIntoleranceCode294967009:
		return "Ciprofibrate allergy"
	case AllergyIntoleranceCode294968004:
		return "Fish oils allergy"
	case AllergyIntoleranceCode294969007:
		return "Omega 3-marine triglycerides allergy"
	case AllergyIntoleranceCode294970008:
		return "HMG COA reductase inhibitor allergy"
	case AllergyIntoleranceCode294971007:
		return "Simvastatin allergy"
	case AllergyIntoleranceCode294972000:
		return "Fluvastatin allergy"
	case AllergyIntoleranceCode294973005:
		return "Pravastatin allergy"
	case AllergyIntoleranceCode294975003:
		return "Adenosine allergy"
	case AllergyIntoleranceCode294977006:
		return "Disopyramide allergy"
	case AllergyIntoleranceCode294978001:
		return "Quinidine allergy"
	case AllergyIntoleranceCode294979009:
		return "Flecainide allergy"
	case AllergyIntoleranceCode294980007:
		return "Mexiletine allergy"
	case AllergyIntoleranceCode294981006:
		return "Moracizine allergy"
	case AllergyIntoleranceCode294982004:
		return "Procainamide allergy"
	case AllergyIntoleranceCode294983009:
		return "Propafenone allergy"
	case AllergyIntoleranceCode294984003:
		return "Tocainide allergy"
	case AllergyIntoleranceCode294986001:
		return "Bretylium allergy"
	case AllergyIntoleranceCode294988000:
		return "Amiodarone allergy"
	case AllergyIntoleranceCode294992007:
		return "Bendrofluazide allergy"
	case AllergyIntoleranceCode294993002:
		return "Chlorothiazide allergy"
	case AllergyIntoleranceCode294994008:
		return "Cyclopenthiazide allergy"
	case AllergyIntoleranceCode294995009:
		return "Hydrochlorothiazide allergy"
	case AllergyIntoleranceCode294996005:
		return "Hydroflumethiazide allergy"
	case AllergyIntoleranceCode294997001:
		return "Methyclothiazide allergy"
	case AllergyIntoleranceCode294998006:
		return "Polythiazide allergy"
	case AllergyIntoleranceCode295000003:
		return "Frusemide allergy"
	case AllergyIntoleranceCode295001004:
		return "Bumetanide allergy"
	case AllergyIntoleranceCode295002006:
		return "Ethacrynic acid allergy"
	case AllergyIntoleranceCode295003001:
		return "Piretanide allergy"
	case AllergyIntoleranceCode295004007:
		return "Torasemide allergy"
	case AllergyIntoleranceCode295006009:
		return "Triamterene allergy"
	case AllergyIntoleranceCode295007000:
		return "Aldosterone antagonists allergy"
	case AllergyIntoleranceCode295008005:
		return "Potassium canrenoate allergy"
	case AllergyIntoleranceCode295009002:
		return "Spironolactone allergy"
	case AllergyIntoleranceCode295010007:
		return "Amiloride allergy"
	case AllergyIntoleranceCode295019008:
		return "Mannitol allergy"
	case AllergyIntoleranceCode267425008:
		return "Lactose intolerance"
	case AllergyIntoleranceCode29736007:
		return "Syndrome of carbohydrate intolerance"
	case AllergyIntoleranceCode29512005:
		return "Lactase deficiency in diseases other than of the small intestine"
	case AllergyIntoleranceCode52070001:
		return "Acquired monosaccharide malabsorption"
	case AllergyIntoleranceCode237978005:
		return "Glycerol intolerance syndrome"
	case AllergyIntoleranceCode340519003:
		return "Lysine intolerance"
	case AllergyIntoleranceCode190753003:
		return "Sucrose intolerance"
	case AllergyIntoleranceCode413427002:
		return "Acquired fructose intolerance (disorder)"
	case AllergyIntoleranceCode716186003:
		return "No known allergy (situation)"
	case AllergyIntoleranceCode409137002:
		return "No known history of drug allergy"
	case AllergyIntoleranceCode428197003:
		return "No known insect allergy (situation)"
	case AllergyIntoleranceCode428607008:
		return "No known environmental allergy (situation)"
	case AllergyIntoleranceCode429625007:
		return "No known food allergy (situation)"
	case AllergyIntoleranceCode716220001:
		return "No known animal allergy (situation)"
	case AllergyIntoleranceCode1003774007:
		return "No known Hevea brasiliensis latex allergy (situation)"
	default:
		return "Unknown Allergy Intolerance Code"
	}
}
