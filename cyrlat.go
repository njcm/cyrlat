//Package cyrlat to translate the Cyrillic to the Latin alphabet for use in semantic URLs, file names, etc.
package cyrlat

import (
	"strings"
)

//const Modificators
const (
	_ = iota
	// Lower all letters mapped to their lower case
	Lower
	// Upper all letters mapped to their upper case
	Upper
	// UnSpace all spaces replaced underscore symbol
	UnSpace
)

var table = map[rune]string{
	1040: "A", 1041: "B", 1042: "V", 1043: "G", 1027: "G`", 1168: "G`", 1044: "D", 1045: "E",
	1025: "YO", 1028: "YE", 1046: "ZH", 1047: "Z", 1029: "Z", 1048: "I", 1049: "Y", 1032: "J",
	1030: "I", 1031: "YI", 1050: "K", 1036: "K", 1051: "L", 1033: "L", 1052: "M", 1053: "N",
	1034: "N", 1054: "O", 1055: "P", 1056: "R", 1057: "S", 1058: "T", 1059: "U", 1038: "U",
	1060: "F", 1061: "H", 1062: "TS", 1063: "CH", 1039: "DH", 1064: "SH", 1065: "SHH", 1066: "``",
	1067: "YI", 1068: "`", 1069: "E`", 1070: "YU", 1071: "YA", 1072: "a", 1073: "b", 1074: "v",
	1075: "g", 1107: "g", 1169: "g", 1076: "d", 1077: "e", 1105: "yo", 1108: "ye", 1078: "zh",
	1079: "z", 1109: "z", 1080: "i", 1081: "y", 1112: "j", 1110: "i", 1111: "yi", 1082: "k",
	1116: "k", 1083: "l", 1113: "l", 1084: "m", 1085: "n", 1114: "n", 1086: "o", 1087: "p",
	1088: "r", 1089: "s", 1090: "t", 1091: "u", 1118: "u", 1092: "f", 1093: "h", 1094: "ts",
	1095: "ch", 1119: "dh", 1096: "sh", 1097: "shh", 1100: "'", 1099: "yi", 1098: "\"", 1101: "e`",
	1102: "yu", 1103: "ya"}

// Latin returns a string in the Latin alphabet
func Latin(cyrillic string, modificators ...int) string {
	table[32] = " "
	var to int
	var lat string
	for _, md := range modificators {
		switch md {
		case UnSpace:
			table[32] = "_"
		default:
			to = md
		}
	}
	for _, rn := range cyrillic {
		if l := table[rn]; len(l) > 0 {
			lat += l
		} else {
			lat += string(rn)
		}
	}
	switch to {
	case Lower:
		lat = strings.ToLower(lat)
	case Upper:
		lat = strings.ToUpper(lat)
	}
	return lat
}
