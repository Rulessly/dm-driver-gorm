package driver

import (
	"bytes"
	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/ianaindex"
	"golang.org/x/text/transform"
	"io"
	"io/ioutil"
	"math"
)

type dm_build_922 struct{}

var Dm_build_923 = &dm_build_922{}

func (Dm_build_925 *dm_build_922) Dm_build_924(dm_build_926 []byte, dm_build_927 int, dm_build_928 byte) int {
	dm_build_926[dm_build_927] = dm_build_928
	return 1
}

func (Dm_build_930 *dm_build_922) Dm_build_929(dm_build_931 []byte, dm_build_932 int, dm_build_933 int8) int {
	dm_build_931[dm_build_932] = byte(dm_build_933)
	return 1
}

func (Dm_build_935 *dm_build_922) Dm_build_934(dm_build_936 []byte, dm_build_937 int, dm_build_938 int16) int {
	dm_build_936[dm_build_937] = byte(dm_build_938)
	dm_build_937++
	dm_build_936[dm_build_937] = byte(dm_build_938 >> 8)
	return 2
}

func (Dm_build_940 *dm_build_922) Dm_build_939(dm_build_941 []byte, dm_build_942 int, dm_build_943 int32) int {
	dm_build_941[dm_build_942] = byte(dm_build_943)
	dm_build_942++
	dm_build_941[dm_build_942] = byte(dm_build_943 >> 8)
	dm_build_942++
	dm_build_941[dm_build_942] = byte(dm_build_943 >> 16)
	dm_build_942++
	dm_build_941[dm_build_942] = byte(dm_build_943 >> 24)
	dm_build_942++
	return 4
}

func (Dm_build_945 *dm_build_922) Dm_build_944(dm_build_946 []byte, dm_build_947 int, dm_build_948 int64) int {
	dm_build_946[dm_build_947] = byte(dm_build_948)
	dm_build_947++
	dm_build_946[dm_build_947] = byte(dm_build_948 >> 8)
	dm_build_947++
	dm_build_946[dm_build_947] = byte(dm_build_948 >> 16)
	dm_build_947++
	dm_build_946[dm_build_947] = byte(dm_build_948 >> 24)
	dm_build_947++
	dm_build_946[dm_build_947] = byte(dm_build_948 >> 32)
	dm_build_947++
	dm_build_946[dm_build_947] = byte(dm_build_948 >> 40)
	dm_build_947++
	dm_build_946[dm_build_947] = byte(dm_build_948 >> 48)
	dm_build_947++
	dm_build_946[dm_build_947] = byte(dm_build_948 >> 56)
	return 8
}

func (Dm_build_950 *dm_build_922) Dm_build_949(dm_build_951 []byte, dm_build_952 int, dm_build_953 float32) int {
	return Dm_build_950.Dm_build_969(dm_build_951, dm_build_952, math.Float32bits(dm_build_953))
}

func (Dm_build_955 *dm_build_922) Dm_build_954(dm_build_956 []byte, dm_build_957 int, dm_build_958 float64) int {
	return Dm_build_955.Dm_build_974(dm_build_956, dm_build_957, math.Float64bits(dm_build_958))
}

func (Dm_build_960 *dm_build_922) Dm_build_959(dm_build_961 []byte, dm_build_962 int, dm_build_963 uint8) int {
	dm_build_961[dm_build_962] = byte(dm_build_963)
	return 1
}

func (Dm_build_965 *dm_build_922) Dm_build_964(dm_build_966 []byte, dm_build_967 int, dm_build_968 uint16) int {
	dm_build_966[dm_build_967] = byte(dm_build_968)
	dm_build_967++
	dm_build_966[dm_build_967] = byte(dm_build_968 >> 8)
	return 2
}

func (Dm_build_970 *dm_build_922) Dm_build_969(dm_build_971 []byte, dm_build_972 int, dm_build_973 uint32) int {
	dm_build_971[dm_build_972] = byte(dm_build_973)
	dm_build_972++
	dm_build_971[dm_build_972] = byte(dm_build_973 >> 8)
	dm_build_972++
	dm_build_971[dm_build_972] = byte(dm_build_973 >> 16)
	dm_build_972++
	dm_build_971[dm_build_972] = byte(dm_build_973 >> 24)
	return 3
}

func (Dm_build_975 *dm_build_922) Dm_build_974(dm_build_976 []byte, dm_build_977 int, dm_build_978 uint64) int {
	dm_build_976[dm_build_977] = byte(dm_build_978)
	dm_build_977++
	dm_build_976[dm_build_977] = byte(dm_build_978 >> 8)
	dm_build_977++
	dm_build_976[dm_build_977] = byte(dm_build_978 >> 16)
	dm_build_977++
	dm_build_976[dm_build_977] = byte(dm_build_978 >> 24)
	dm_build_977++
	dm_build_976[dm_build_977] = byte(dm_build_978 >> 32)
	dm_build_977++
	dm_build_976[dm_build_977] = byte(dm_build_978 >> 40)
	dm_build_977++
	dm_build_976[dm_build_977] = byte(dm_build_978 >> 48)
	dm_build_977++
	dm_build_976[dm_build_977] = byte(dm_build_978 >> 56)
	return 3
}

func (Dm_build_980 *dm_build_922) Dm_build_979(dm_build_981 []byte, dm_build_982 int, dm_build_983 []byte, dm_build_984 int, dm_build_985 int) int {
	copy(dm_build_981[dm_build_982:dm_build_982+dm_build_985], dm_build_983[dm_build_984:dm_build_984+dm_build_985])
	return dm_build_985
}

func (Dm_build_987 *dm_build_922) Dm_build_986(dm_build_988 []byte, dm_build_989 int, dm_build_990 []byte, dm_build_991 int, dm_build_992 int) int {
	dm_build_989 += Dm_build_987.Dm_build_969(dm_build_988, dm_build_989, uint32(dm_build_992))
	return 4 + Dm_build_987.Dm_build_979(dm_build_988, dm_build_989, dm_build_990, dm_build_991, dm_build_992)
}

func (Dm_build_994 *dm_build_922) Dm_build_993(dm_build_995 []byte, dm_build_996 int, dm_build_997 []byte, dm_build_998 int, dm_build_999 int) int {
	dm_build_996 += Dm_build_994.Dm_build_964(dm_build_995, dm_build_996, uint16(dm_build_999))
	return 2 + Dm_build_994.Dm_build_979(dm_build_995, dm_build_996, dm_build_997, dm_build_998, dm_build_999)
}

func (Dm_build_1001 *dm_build_922) Dm_build_1000(dm_build_1002 []byte, dm_build_1003 int, dm_build_1004 string, dm_build_1005 string, dm_build_1006 *DmConnection) int {
	dm_build_1007 := Dm_build_1001.Dm_build_1139(dm_build_1004, dm_build_1005, dm_build_1006)
	dm_build_1003 += Dm_build_1001.Dm_build_969(dm_build_1002, dm_build_1003, uint32(len(dm_build_1007)))
	return 4 + Dm_build_1001.Dm_build_979(dm_build_1002, dm_build_1003, dm_build_1007, 0, len(dm_build_1007))
}

func (Dm_build_1009 *dm_build_922) Dm_build_1008(dm_build_1010 []byte, dm_build_1011 int, dm_build_1012 string, dm_build_1013 string, dm_build_1014 *DmConnection) int {
	dm_build_1015 := Dm_build_1009.Dm_build_1139(dm_build_1012, dm_build_1013, dm_build_1014)

	dm_build_1011 += Dm_build_1009.Dm_build_964(dm_build_1010, dm_build_1011, uint16(len(dm_build_1015)))
	return 2 + Dm_build_1009.Dm_build_979(dm_build_1010, dm_build_1011, dm_build_1015, 0, len(dm_build_1015))
}

func (Dm_build_1017 *dm_build_922) Dm_build_1016(dm_build_1018 []byte, dm_build_1019 int) byte {
	return dm_build_1018[dm_build_1019]
}

func (Dm_build_1021 *dm_build_922) Dm_build_1020(dm_build_1022 []byte, dm_build_1023 int) int16 {
	var dm_build_1024 int16
	dm_build_1024 = int16(dm_build_1022[dm_build_1023] & 0xff)
	dm_build_1023++
	dm_build_1024 |= int16(dm_build_1022[dm_build_1023]&0xff) << 8
	return dm_build_1024
}

func (Dm_build_1026 *dm_build_922) Dm_build_1025(dm_build_1027 []byte, dm_build_1028 int) int32 {
	var dm_build_1029 int32
	dm_build_1029 = int32(dm_build_1027[dm_build_1028] & 0xff)
	dm_build_1028++
	dm_build_1029 |= int32(dm_build_1027[dm_build_1028]&0xff) << 8
	dm_build_1028++
	dm_build_1029 |= int32(dm_build_1027[dm_build_1028]&0xff) << 16
	dm_build_1028++
	dm_build_1029 |= int32(dm_build_1027[dm_build_1028]&0xff) << 24
	return dm_build_1029
}

func (Dm_build_1031 *dm_build_922) Dm_build_1030(dm_build_1032 []byte, dm_build_1033 int) int64 {
	var dm_build_1034 int64
	dm_build_1034 = int64(dm_build_1032[dm_build_1033] & 0xff)
	dm_build_1033++
	dm_build_1034 |= int64(dm_build_1032[dm_build_1033]&0xff) << 8
	dm_build_1033++
	dm_build_1034 |= int64(dm_build_1032[dm_build_1033]&0xff) << 16
	dm_build_1033++
	dm_build_1034 |= int64(dm_build_1032[dm_build_1033]&0xff) << 24
	dm_build_1033++
	dm_build_1034 |= int64(dm_build_1032[dm_build_1033]&0xff) << 32
	dm_build_1033++
	dm_build_1034 |= int64(dm_build_1032[dm_build_1033]&0xff) << 40
	dm_build_1033++
	dm_build_1034 |= int64(dm_build_1032[dm_build_1033]&0xff) << 48
	dm_build_1033++
	dm_build_1034 |= int64(dm_build_1032[dm_build_1033]&0xff) << 56
	return dm_build_1034
}

func (Dm_build_1036 *dm_build_922) Dm_build_1035(dm_build_1037 []byte, dm_build_1038 int) float32 {
	return math.Float32frombits(Dm_build_1036.Dm_build_1052(dm_build_1037, dm_build_1038))
}

func (Dm_build_1040 *dm_build_922) Dm_build_1039(dm_build_1041 []byte, dm_build_1042 int) float64 {
	return math.Float64frombits(Dm_build_1040.Dm_build_1057(dm_build_1041, dm_build_1042))
}

func (Dm_build_1044 *dm_build_922) Dm_build_1043(dm_build_1045 []byte, dm_build_1046 int) uint8 {
	return uint8(dm_build_1045[dm_build_1046] & 0xff)
}

func (Dm_build_1048 *dm_build_922) Dm_build_1047(dm_build_1049 []byte, dm_build_1050 int) uint16 {
	var dm_build_1051 uint16
	dm_build_1051 = uint16(dm_build_1049[dm_build_1050] & 0xff)
	dm_build_1050++
	dm_build_1051 |= uint16(dm_build_1049[dm_build_1050]&0xff) << 8
	return dm_build_1051
}

func (Dm_build_1053 *dm_build_922) Dm_build_1052(dm_build_1054 []byte, dm_build_1055 int) uint32 {
	var dm_build_1056 uint32
	dm_build_1056 = uint32(dm_build_1054[dm_build_1055] & 0xff)
	dm_build_1055++
	dm_build_1056 |= uint32(dm_build_1054[dm_build_1055]&0xff) << 8
	dm_build_1055++
	dm_build_1056 |= uint32(dm_build_1054[dm_build_1055]&0xff) << 16
	dm_build_1055++
	dm_build_1056 |= uint32(dm_build_1054[dm_build_1055]&0xff) << 24
	return dm_build_1056
}

func (Dm_build_1058 *dm_build_922) Dm_build_1057(dm_build_1059 []byte, dm_build_1060 int) uint64 {
	var dm_build_1061 uint64
	dm_build_1061 = uint64(dm_build_1059[dm_build_1060] & 0xff)
	dm_build_1060++
	dm_build_1061 |= uint64(dm_build_1059[dm_build_1060]&0xff) << 8
	dm_build_1060++
	dm_build_1061 |= uint64(dm_build_1059[dm_build_1060]&0xff) << 16
	dm_build_1060++
	dm_build_1061 |= uint64(dm_build_1059[dm_build_1060]&0xff) << 24
	dm_build_1060++
	dm_build_1061 |= uint64(dm_build_1059[dm_build_1060]&0xff) << 32
	dm_build_1060++
	dm_build_1061 |= uint64(dm_build_1059[dm_build_1060]&0xff) << 40
	dm_build_1060++
	dm_build_1061 |= uint64(dm_build_1059[dm_build_1060]&0xff) << 48
	dm_build_1060++
	dm_build_1061 |= uint64(dm_build_1059[dm_build_1060]&0xff) << 56
	return dm_build_1061
}

func (Dm_build_1063 *dm_build_922) Dm_build_1062(dm_build_1064 []byte, dm_build_1065 int) []byte {
	dm_build_1066 := Dm_build_1063.Dm_build_1052(dm_build_1064, dm_build_1065)

	dm_build_1067 := make([]byte, dm_build_1066)
	copy(dm_build_1067[:int(dm_build_1066)], dm_build_1064[dm_build_1065+4:dm_build_1065+4+int(dm_build_1066)])
	return dm_build_1067
}

func (Dm_build_1069 *dm_build_922) Dm_build_1068(dm_build_1070 []byte, dm_build_1071 int) []byte {
	dm_build_1072 := Dm_build_1069.Dm_build_1047(dm_build_1070, dm_build_1071)

	dm_build_1073 := make([]byte, dm_build_1072)
	copy(dm_build_1073[:int(dm_build_1072)], dm_build_1070[dm_build_1071+2:dm_build_1071+2+int(dm_build_1072)])
	return dm_build_1073
}

func (Dm_build_1075 *dm_build_922) Dm_build_1074(dm_build_1076 []byte, dm_build_1077 int, dm_build_1078 int) []byte {

	dm_build_1079 := make([]byte, dm_build_1078)
	copy(dm_build_1079[:dm_build_1078], dm_build_1076[dm_build_1077:dm_build_1077+dm_build_1078])
	return dm_build_1079
}

func (Dm_build_1081 *dm_build_922) Dm_build_1080(dm_build_1082 []byte, dm_build_1083 int, dm_build_1084 int, dm_build_1085 string, dm_build_1086 *DmConnection) string {
	return Dm_build_1081.Dm_build_1175(dm_build_1082[dm_build_1083:dm_build_1083+dm_build_1084], dm_build_1085, dm_build_1086)
}

func (Dm_build_1088 *dm_build_922) Dm_build_1087(dm_build_1089 []byte, dm_build_1090 int, dm_build_1091 string, dm_build_1092 *DmConnection) string {
	dm_build_1093 := Dm_build_1088.Dm_build_1052(dm_build_1089, dm_build_1090)
	dm_build_1090 += 4
	return Dm_build_1088.Dm_build_1080(dm_build_1089, dm_build_1090, int(dm_build_1093), dm_build_1091, dm_build_1092)
}

func (Dm_build_1095 *dm_build_922) Dm_build_1094(dm_build_1096 []byte, dm_build_1097 int, dm_build_1098 string, dm_build_1099 *DmConnection) string {
	dm_build_1100 := Dm_build_1095.Dm_build_1047(dm_build_1096, dm_build_1097)
	dm_build_1097 += 2
	return Dm_build_1095.Dm_build_1080(dm_build_1096, dm_build_1097, int(dm_build_1100), dm_build_1098, dm_build_1099)
}

func (Dm_build_1102 *dm_build_922) Dm_build_1101(dm_build_1103 byte) []byte {
	return []byte{dm_build_1103}
}

func (Dm_build_1105 *dm_build_922) Dm_build_1104(dm_build_1106 int8) []byte {
	return []byte{byte(dm_build_1106)}
}

func (Dm_build_1108 *dm_build_922) Dm_build_1107(dm_build_1109 int16) []byte {
	return []byte{byte(dm_build_1109), byte(dm_build_1109 >> 8)}
}

func (Dm_build_1111 *dm_build_922) Dm_build_1110(dm_build_1112 int32) []byte {
	return []byte{byte(dm_build_1112), byte(dm_build_1112 >> 8), byte(dm_build_1112 >> 16), byte(dm_build_1112 >> 24)}
}

func (Dm_build_1114 *dm_build_922) Dm_build_1113(dm_build_1115 int64) []byte {
	return []byte{byte(dm_build_1115), byte(dm_build_1115 >> 8), byte(dm_build_1115 >> 16), byte(dm_build_1115 >> 24), byte(dm_build_1115 >> 32),
		byte(dm_build_1115 >> 40), byte(dm_build_1115 >> 48), byte(dm_build_1115 >> 56)}
}

func (Dm_build_1117 *dm_build_922) Dm_build_1116(dm_build_1118 float32) []byte {
	return Dm_build_1117.Dm_build_1128(math.Float32bits(dm_build_1118))
}

func (Dm_build_1120 *dm_build_922) Dm_build_1119(dm_build_1121 float64) []byte {
	return Dm_build_1120.Dm_build_1131(math.Float64bits(dm_build_1121))
}

func (Dm_build_1123 *dm_build_922) Dm_build_1122(dm_build_1124 uint8) []byte {
	return []byte{byte(dm_build_1124)}
}

func (Dm_build_1126 *dm_build_922) Dm_build_1125(dm_build_1127 uint16) []byte {
	return []byte{byte(dm_build_1127), byte(dm_build_1127 >> 8)}
}

func (Dm_build_1129 *dm_build_922) Dm_build_1128(dm_build_1130 uint32) []byte {
	return []byte{byte(dm_build_1130), byte(dm_build_1130 >> 8), byte(dm_build_1130 >> 16), byte(dm_build_1130 >> 24)}
}

func (Dm_build_1132 *dm_build_922) Dm_build_1131(dm_build_1133 uint64) []byte {
	return []byte{byte(dm_build_1133), byte(dm_build_1133 >> 8), byte(dm_build_1133 >> 16), byte(dm_build_1133 >> 24), byte(dm_build_1133 >> 32), byte(dm_build_1133 >> 40), byte(dm_build_1133 >> 48), byte(dm_build_1133 >> 56)}
}

func (Dm_build_1135 *dm_build_922) Dm_build_1134(dm_build_1136 []byte, dm_build_1137 string, dm_build_1138 *DmConnection) []byte {
	if dm_build_1137 == "UTF-8" {
		return dm_build_1136
	}

	if dm_build_1138 == nil {
		if e := dm_build_1180(dm_build_1137); e != nil {
			tmp, err := ioutil.ReadAll(
				transform.NewReader(bytes.NewReader(dm_build_1136), e.NewEncoder()),
			)
			if err != nil {
				panic("UTF8 To Charset error!")
			}

			return tmp
		}

		panic("Unsupported Charset!")
	}

	if dm_build_1138.encodeBuffer == nil {
		dm_build_1138.encodeBuffer = bytes.NewBuffer(nil)
		dm_build_1138.encode = dm_build_1180(dm_build_1138.getServerEncoding())
		dm_build_1138.transformReaderDst = make([]byte, 4096)
		dm_build_1138.transformReaderSrc = make([]byte, 4096)
	}

	if e := dm_build_1138.encode; e != nil {

		dm_build_1138.encodeBuffer.Reset()

		n, err := dm_build_1138.encodeBuffer.ReadFrom(
			Dm_build_1194(bytes.NewReader(dm_build_1136), e.NewEncoder(), dm_build_1138.transformReaderDst, dm_build_1138.transformReaderSrc),
		)
		if err != nil {
			panic("UTF8 To Charset error!")
		}
		var tmp = make([]byte, n)
		if _, err = dm_build_1138.encodeBuffer.Read(tmp); err != nil {
			panic("UTF8 To Charset error!")
		}
		return tmp
	}

	panic("Unsupported Charset!")
}

func (Dm_build_1140 *dm_build_922) Dm_build_1139(dm_build_1141 string, dm_build_1142 string, dm_build_1143 *DmConnection) []byte {
	return Dm_build_1140.Dm_build_1134([]byte(dm_build_1141), dm_build_1142, dm_build_1143)
}

func (Dm_build_1145 *dm_build_922) Dm_build_1144(dm_build_1146 []byte) byte {
	return Dm_build_1145.Dm_build_1016(dm_build_1146, 0)
}

func (Dm_build_1148 *dm_build_922) Dm_build_1147(dm_build_1149 []byte) int16 {
	return Dm_build_1148.Dm_build_1020(dm_build_1149, 0)
}

func (Dm_build_1151 *dm_build_922) Dm_build_1150(dm_build_1152 []byte) int32 {
	return Dm_build_1151.Dm_build_1025(dm_build_1152, 0)
}

func (Dm_build_1154 *dm_build_922) Dm_build_1153(dm_build_1155 []byte) int64 {
	return Dm_build_1154.Dm_build_1030(dm_build_1155, 0)
}

func (Dm_build_1157 *dm_build_922) Dm_build_1156(dm_build_1158 []byte) float32 {
	return Dm_build_1157.Dm_build_1035(dm_build_1158, 0)
}

func (Dm_build_1160 *dm_build_922) Dm_build_1159(dm_build_1161 []byte) float64 {
	return Dm_build_1160.Dm_build_1039(dm_build_1161, 0)
}

func (Dm_build_1163 *dm_build_922) Dm_build_1162(dm_build_1164 []byte) uint8 {
	return Dm_build_1163.Dm_build_1043(dm_build_1164, 0)
}

func (Dm_build_1166 *dm_build_922) Dm_build_1165(dm_build_1167 []byte) uint16 {
	return Dm_build_1166.Dm_build_1047(dm_build_1167, 0)
}

func (Dm_build_1169 *dm_build_922) Dm_build_1168(dm_build_1170 []byte) uint32 {
	return Dm_build_1169.Dm_build_1052(dm_build_1170, 0)
}

func (Dm_build_1172 *dm_build_922) Dm_build_1171(dm_build_1173 []byte, dm_build_1174 string) []byte {
	if dm_build_1174 == "UTF-8" {
		return dm_build_1173
	}

	if e := dm_build_1180(dm_build_1174); e != nil {

		tmp, err := ioutil.ReadAll(
			transform.NewReader(bytes.NewReader(dm_build_1173), e.NewDecoder()),
		)
		if err != nil {

			panic("Charset To UTF8 error!")
		}

		return tmp
	}

	panic("Unsupported Charset!")

}

func (Dm_build_1176 *dm_build_922) Dm_build_1175(dm_build_1177 []byte, dm_build_1178 string, dm_build_1179 *DmConnection) string {
	return string(Dm_build_1176.Dm_build_1171(dm_build_1177, dm_build_1178))
}

func dm_build_1180(dm_build_1181 string) encoding.Encoding {
	if e, err := ianaindex.MIB.Encoding(dm_build_1181); err == nil && e != nil {
		return e
	}
	return nil
}

type Dm_build_1182 struct {
	dm_build_1183 io.Reader
	dm_build_1184 transform.Transformer
	dm_build_1185 error

	dm_build_1186                []byte
	dm_build_1187, dm_build_1188 int

	dm_build_1189                []byte
	dm_build_1190, dm_build_1191 int

	dm_build_1192 bool
}

const dm_build_1193 = 4096

func Dm_build_1194(dm_build_1195 io.Reader, dm_build_1196 transform.Transformer, dm_build_1197 []byte, dm_build_1198 []byte) *Dm_build_1182 {
	dm_build_1196.Reset()
	return &Dm_build_1182{
		dm_build_1183: dm_build_1195,
		dm_build_1184: dm_build_1196,
		dm_build_1186: dm_build_1197,
		dm_build_1189: dm_build_1198,
	}
}

func (dm_build_1200 *Dm_build_1182) Read(dm_build_1201 []byte) (int, error) {
	dm_build_1202, dm_build_1203 := 0, error(nil)
	for {

		if dm_build_1200.dm_build_1187 != dm_build_1200.dm_build_1188 {
			dm_build_1202 = copy(dm_build_1201, dm_build_1200.dm_build_1186[dm_build_1200.dm_build_1187:dm_build_1200.dm_build_1188])
			dm_build_1200.dm_build_1187 += dm_build_1202
			if dm_build_1200.dm_build_1187 == dm_build_1200.dm_build_1188 && dm_build_1200.dm_build_1192 {
				return dm_build_1202, dm_build_1200.dm_build_1185
			}
			return dm_build_1202, nil
		} else if dm_build_1200.dm_build_1192 {
			return 0, dm_build_1200.dm_build_1185
		}

		if dm_build_1200.dm_build_1190 != dm_build_1200.dm_build_1191 || dm_build_1200.dm_build_1185 != nil {
			dm_build_1200.dm_build_1187 = 0
			dm_build_1200.dm_build_1188, dm_build_1202, dm_build_1203 = dm_build_1200.dm_build_1184.Transform(dm_build_1200.dm_build_1186, dm_build_1200.dm_build_1189[dm_build_1200.dm_build_1190:dm_build_1200.dm_build_1191], dm_build_1200.dm_build_1185 == io.EOF)
			dm_build_1200.dm_build_1190 += dm_build_1202

			switch {
			case dm_build_1203 == nil:
				if dm_build_1200.dm_build_1190 != dm_build_1200.dm_build_1191 {
					dm_build_1200.dm_build_1185 = nil
				}

				dm_build_1200.dm_build_1192 = dm_build_1200.dm_build_1185 != nil
				continue
			case dm_build_1203 == transform.ErrShortDst && (dm_build_1200.dm_build_1188 != 0 || dm_build_1202 != 0):

				continue
			case dm_build_1203 == transform.ErrShortSrc && dm_build_1200.dm_build_1191-dm_build_1200.dm_build_1190 != len(dm_build_1200.dm_build_1189) && dm_build_1200.dm_build_1185 == nil:

			default:
				dm_build_1200.dm_build_1192 = true

				if dm_build_1200.dm_build_1185 == nil || dm_build_1200.dm_build_1185 == io.EOF {
					dm_build_1200.dm_build_1185 = dm_build_1203
				}
				continue
			}
		}

		if dm_build_1200.dm_build_1190 != 0 {
			dm_build_1200.dm_build_1190, dm_build_1200.dm_build_1191 = 0, copy(dm_build_1200.dm_build_1189, dm_build_1200.dm_build_1189[dm_build_1200.dm_build_1190:dm_build_1200.dm_build_1191])
		}
		dm_build_1202, dm_build_1200.dm_build_1185 = dm_build_1200.dm_build_1183.Read(dm_build_1200.dm_build_1189[dm_build_1200.dm_build_1191:])
		dm_build_1200.dm_build_1191 += dm_build_1202
	}
}
