package driver

import (
	"container/list"
	"io"
)

type Dm_build_1204 struct {
	dm_build_1205 *list.List
	dm_build_1206 *dm_build_1258
	dm_build_1207 int
}

func Dm_build_1208() *Dm_build_1204 {
	return &Dm_build_1204{
		dm_build_1205: list.New(),
		dm_build_1207: 0,
	}
}

func (dm_build_1210 *Dm_build_1204) Dm_build_1209() int {
	return dm_build_1210.dm_build_1207
}

func (dm_build_1212 *Dm_build_1204) Dm_build_1211(dm_build_1213 *Dm_build_1282, dm_build_1214 int) int {
	var dm_build_1215 = 0
	var dm_build_1216 = 0
	for dm_build_1215 < dm_build_1214 && dm_build_1212.dm_build_1206 != nil {
		dm_build_1216 = dm_build_1212.dm_build_1206.dm_build_1266(dm_build_1213, dm_build_1214-dm_build_1215)
		if dm_build_1212.dm_build_1206.dm_build_1261 == 0 {
			dm_build_1212.dm_build_1248()
		}
		dm_build_1215 += dm_build_1216
		dm_build_1212.dm_build_1207 -= dm_build_1216
	}
	return dm_build_1215
}

func (dm_build_1218 *Dm_build_1204) Dm_build_1217(dm_build_1219 []byte, dm_build_1220 int, dm_build_1221 int) int {
	var dm_build_1222 = 0
	var dm_build_1223 = 0
	for dm_build_1222 < dm_build_1221 && dm_build_1218.dm_build_1206 != nil {
		dm_build_1223 = dm_build_1218.dm_build_1206.dm_build_1270(dm_build_1219, dm_build_1220, dm_build_1221-dm_build_1222)
		if dm_build_1218.dm_build_1206.dm_build_1261 == 0 {
			dm_build_1218.dm_build_1248()
		}
		dm_build_1222 += dm_build_1223
		dm_build_1218.dm_build_1207 -= dm_build_1223
		dm_build_1220 += dm_build_1223
	}
	return dm_build_1222
}

func (dm_build_1225 *Dm_build_1204) Dm_build_1224(dm_build_1226 io.Writer, dm_build_1227 int) int {
	var dm_build_1228 = 0
	var dm_build_1229 = 0
	for dm_build_1228 < dm_build_1227 && dm_build_1225.dm_build_1206 != nil {
		dm_build_1229 = dm_build_1225.dm_build_1206.dm_build_1275(dm_build_1226, dm_build_1227-dm_build_1228)
		if dm_build_1225.dm_build_1206.dm_build_1261 == 0 {
			dm_build_1225.dm_build_1248()
		}
		dm_build_1228 += dm_build_1229
		dm_build_1225.dm_build_1207 -= dm_build_1229
	}
	return dm_build_1228
}

func (dm_build_1231 *Dm_build_1204) Dm_build_1230(dm_build_1232 []byte, dm_build_1233 int, dm_build_1234 int) {
	if dm_build_1234 == 0 {
		return
	}
	var dm_build_1235 = dm_build_1262(dm_build_1232, dm_build_1233, dm_build_1234)
	if dm_build_1231.dm_build_1206 == nil {
		dm_build_1231.dm_build_1206 = dm_build_1235
	} else {
		dm_build_1231.dm_build_1205.PushBack(dm_build_1235)
	}
	dm_build_1231.dm_build_1207 += dm_build_1234
}

func (dm_build_1237 *Dm_build_1204) dm_build_1236(dm_build_1238 int) byte {
	var dm_build_1239 = dm_build_1238
	var dm_build_1240 = dm_build_1237.dm_build_1206
	for dm_build_1239 > 0 && dm_build_1240 != nil {
		if dm_build_1240.dm_build_1261 == 0 {
			continue
		}
		if dm_build_1239 > dm_build_1240.dm_build_1261-1 {
			dm_build_1239 -= dm_build_1240.dm_build_1261
			dm_build_1240 = dm_build_1237.dm_build_1205.Front().Value.(*dm_build_1258)
		} else {
			break
		}
	}
	return dm_build_1240.dm_build_1279(dm_build_1239)
}
func (dm_build_1242 *Dm_build_1204) Dm_build_1241(dm_build_1243 *Dm_build_1204) {
	if dm_build_1243.dm_build_1207 == 0 {
		return
	}
	var dm_build_1244 = dm_build_1243.dm_build_1206
	for dm_build_1244 != nil {
		dm_build_1242.dm_build_1245(dm_build_1244)
		dm_build_1243.dm_build_1248()
		dm_build_1244 = dm_build_1243.dm_build_1206
	}
	dm_build_1243.dm_build_1207 = 0
}
func (dm_build_1246 *Dm_build_1204) dm_build_1245(dm_build_1247 *dm_build_1258) {
	if dm_build_1247.dm_build_1261 == 0 {
		return
	}
	if dm_build_1246.dm_build_1206 == nil {
		dm_build_1246.dm_build_1206 = dm_build_1247
	} else {
		dm_build_1246.dm_build_1205.PushBack(dm_build_1247)
	}
	dm_build_1246.dm_build_1207 += dm_build_1247.dm_build_1261
}

func (dm_build_1249 *Dm_build_1204) dm_build_1248() {
	var dm_build_1250 = dm_build_1249.dm_build_1205.Front()
	if dm_build_1250 == nil {
		dm_build_1249.dm_build_1206 = nil
	} else {
		dm_build_1249.dm_build_1206 = dm_build_1250.Value.(*dm_build_1258)
		dm_build_1249.dm_build_1205.Remove(dm_build_1250)
	}
}

func (dm_build_1252 *Dm_build_1204) Dm_build_1251() []byte {
	var dm_build_1253 = make([]byte, dm_build_1252.dm_build_1207)
	var dm_build_1254 = dm_build_1252.dm_build_1206
	var dm_build_1255 = 0
	var dm_build_1256 = len(dm_build_1253)
	var dm_build_1257 = 0
	for dm_build_1254 != nil {
		if dm_build_1254.dm_build_1261 > 0 {
			if dm_build_1256 > dm_build_1254.dm_build_1261 {
				dm_build_1257 = dm_build_1254.dm_build_1261
			} else {
				dm_build_1257 = dm_build_1256
			}
			copy(dm_build_1253[dm_build_1255:dm_build_1255+dm_build_1257], dm_build_1254.dm_build_1259[dm_build_1254.dm_build_1260:dm_build_1254.dm_build_1260+dm_build_1257])
			dm_build_1255 += dm_build_1257
			dm_build_1256 -= dm_build_1257
		}
		if dm_build_1252.dm_build_1205.Front() == nil {
			dm_build_1254 = nil
		} else {
			dm_build_1254 = dm_build_1252.dm_build_1205.Front().Value.(*dm_build_1258)
		}
	}
	return dm_build_1253
}

type dm_build_1258 struct {
	dm_build_1259 []byte
	dm_build_1260 int
	dm_build_1261 int
}

func dm_build_1262(dm_build_1263 []byte, dm_build_1264 int, dm_build_1265 int) *dm_build_1258 {
	return &dm_build_1258{
		dm_build_1263,
		dm_build_1264,
		dm_build_1265,
	}
}

func (dm_build_1267 *dm_build_1258) dm_build_1266(dm_build_1268 *Dm_build_1282, dm_build_1269 int) int {
	if dm_build_1267.dm_build_1261 <= dm_build_1269 {
		dm_build_1269 = dm_build_1267.dm_build_1261
	}
	dm_build_1268.Dm_build_1365(dm_build_1267.dm_build_1259[dm_build_1267.dm_build_1260 : dm_build_1267.dm_build_1260+dm_build_1269])
	dm_build_1267.dm_build_1260 += dm_build_1269
	dm_build_1267.dm_build_1261 -= dm_build_1269
	return dm_build_1269
}

func (dm_build_1271 *dm_build_1258) dm_build_1270(dm_build_1272 []byte, dm_build_1273 int, dm_build_1274 int) int {
	if dm_build_1271.dm_build_1261 <= dm_build_1274 {
		dm_build_1274 = dm_build_1271.dm_build_1261
	}
	copy(dm_build_1272[dm_build_1273:dm_build_1273+dm_build_1274], dm_build_1271.dm_build_1259[dm_build_1271.dm_build_1260:dm_build_1271.dm_build_1260+dm_build_1274])
	dm_build_1271.dm_build_1260 += dm_build_1274
	dm_build_1271.dm_build_1261 -= dm_build_1274
	return dm_build_1274
}

func (dm_build_1276 *dm_build_1258) dm_build_1275(dm_build_1277 io.Writer, dm_build_1278 int) int {
	if dm_build_1276.dm_build_1261 <= dm_build_1278 {
		dm_build_1278 = dm_build_1276.dm_build_1261
	}
	dm_build_1277.Write(dm_build_1276.dm_build_1259[dm_build_1276.dm_build_1260 : dm_build_1276.dm_build_1260+dm_build_1278])
	dm_build_1276.dm_build_1260 += dm_build_1278
	dm_build_1276.dm_build_1261 -= dm_build_1278
	return dm_build_1278
}
func (dm_build_1280 *dm_build_1258) dm_build_1279(dm_build_1281 int) byte {
	return dm_build_1280.dm_build_1259[dm_build_1280.dm_build_1260+dm_build_1281]
}
