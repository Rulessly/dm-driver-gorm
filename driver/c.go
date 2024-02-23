package driver

import (
	"io"
	"math"
)

type Dm_build_1282 struct {
	dm_build_1283 []byte
	dm_build_1284 int
}

func Dm_build_1285(dm_build_1286 int) *Dm_build_1282 {
	return &Dm_build_1282{make([]byte, 0, dm_build_1286), 0}
}

func Dm_build_1287(dm_build_1288 []byte) *Dm_build_1282 {
	return &Dm_build_1282{dm_build_1288, 0}
}

func (dm_build_1290 *Dm_build_1282) dm_build_1289(dm_build_1291 int) *Dm_build_1282 {

	dm_build_1292 := len(dm_build_1290.dm_build_1283)
	dm_build_1293 := cap(dm_build_1290.dm_build_1283)

	if dm_build_1292+dm_build_1291 <= dm_build_1293 {
		dm_build_1290.dm_build_1283 = dm_build_1290.dm_build_1283[:dm_build_1292+dm_build_1291]
	} else {

		var calCap = int64(math.Max(float64(2*dm_build_1293), float64(dm_build_1291+dm_build_1292)))

		nbuf := make([]byte, dm_build_1291+dm_build_1292, calCap)
		copy(nbuf, dm_build_1290.dm_build_1283)
		dm_build_1290.dm_build_1283 = nbuf
	}

	return dm_build_1290
}

func (dm_build_1295 *Dm_build_1282) Dm_build_1294() int {
	return len(dm_build_1295.dm_build_1283)
}

func (dm_build_1297 *Dm_build_1282) Dm_build_1296(dm_build_1298 int) *Dm_build_1282 {
	for i := dm_build_1298; i < len(dm_build_1297.dm_build_1283); i++ {
		dm_build_1297.dm_build_1283[i] = 0
	}
	dm_build_1297.dm_build_1283 = dm_build_1297.dm_build_1283[:dm_build_1298]
	return dm_build_1297
}

func (dm_build_1300 *Dm_build_1282) Dm_build_1299(dm_build_1301 int) *Dm_build_1282 {
	dm_build_1300.dm_build_1284 = dm_build_1301
	return dm_build_1300
}

func (dm_build_1303 *Dm_build_1282) Dm_build_1302() int {
	return dm_build_1303.dm_build_1284
}

func (dm_build_1305 *Dm_build_1282) Dm_build_1304(dm_build_1306 bool) int {
	return len(dm_build_1305.dm_build_1283) - dm_build_1305.dm_build_1284
}

func (dm_build_1308 *Dm_build_1282) Dm_build_1307(dm_build_1309 int, dm_build_1310 bool, dm_build_1311 bool) *Dm_build_1282 {

	if dm_build_1310 {
		if dm_build_1311 {
			dm_build_1308.dm_build_1289(dm_build_1309)
		} else {
			dm_build_1308.dm_build_1283 = dm_build_1308.dm_build_1283[:len(dm_build_1308.dm_build_1283)-dm_build_1309]
		}
	} else {
		if dm_build_1311 {
			dm_build_1308.dm_build_1284 += dm_build_1309
		} else {
			dm_build_1308.dm_build_1284 -= dm_build_1309
		}
	}

	return dm_build_1308
}

func (dm_build_1313 *Dm_build_1282) Dm_build_1312(dm_build_1314 io.Reader, dm_build_1315 int) (int, error) {
	dm_build_1316 := len(dm_build_1313.dm_build_1283)
	dm_build_1313.dm_build_1289(dm_build_1315)
	dm_build_1317 := 0
	for dm_build_1315 > 0 {
		n, err := dm_build_1314.Read(dm_build_1313.dm_build_1283[dm_build_1316+dm_build_1317:])
		if n > 0 && err == io.EOF {
			dm_build_1317 += n
			dm_build_1313.dm_build_1283 = dm_build_1313.dm_build_1283[:dm_build_1316+dm_build_1317]
			return dm_build_1317, nil
		} else if n > 0 && err == nil {
			dm_build_1315 -= n
			dm_build_1317 += n
		} else if n == 0 && err != nil {
			return -1, ECGO_COMMUNITION_ERROR.addDetailln(err.Error()).throw()
		}
	}

	return dm_build_1317, nil
}

func (dm_build_1319 *Dm_build_1282) Dm_build_1318(dm_build_1320 io.Writer) (*Dm_build_1282, error) {
	if _, err := dm_build_1320.Write(dm_build_1319.dm_build_1283); err != nil {
		return nil, ECGO_COMMUNITION_ERROR.addDetailln(err.Error()).throw()
	}
	return dm_build_1319, nil
}

func (dm_build_1322 *Dm_build_1282) Dm_build_1321(dm_build_1323 bool) int {
	dm_build_1324 := len(dm_build_1322.dm_build_1283)
	dm_build_1322.dm_build_1289(1)

	if dm_build_1323 {
		return copy(dm_build_1322.dm_build_1283[dm_build_1324:], []byte{1})
	} else {
		return copy(dm_build_1322.dm_build_1283[dm_build_1324:], []byte{0})
	}
}

func (dm_build_1326 *Dm_build_1282) Dm_build_1325(dm_build_1327 byte) int {
	dm_build_1328 := len(dm_build_1326.dm_build_1283)
	dm_build_1326.dm_build_1289(1)

	return copy(dm_build_1326.dm_build_1283[dm_build_1328:], Dm_build_923.Dm_build_1101(dm_build_1327))
}

func (dm_build_1330 *Dm_build_1282) Dm_build_1329(dm_build_1331 int8) int {
	dm_build_1332 := len(dm_build_1330.dm_build_1283)
	dm_build_1330.dm_build_1289(1)

	return copy(dm_build_1330.dm_build_1283[dm_build_1332:], Dm_build_923.Dm_build_1104(dm_build_1331))
}

func (dm_build_1334 *Dm_build_1282) Dm_build_1333(dm_build_1335 int16) int {
	dm_build_1336 := len(dm_build_1334.dm_build_1283)
	dm_build_1334.dm_build_1289(2)

	return copy(dm_build_1334.dm_build_1283[dm_build_1336:], Dm_build_923.Dm_build_1107(dm_build_1335))
}

func (dm_build_1338 *Dm_build_1282) Dm_build_1337(dm_build_1339 int32) int {
	dm_build_1340 := len(dm_build_1338.dm_build_1283)
	dm_build_1338.dm_build_1289(4)

	return copy(dm_build_1338.dm_build_1283[dm_build_1340:], Dm_build_923.Dm_build_1110(dm_build_1339))
}

func (dm_build_1342 *Dm_build_1282) Dm_build_1341(dm_build_1343 uint8) int {
	dm_build_1344 := len(dm_build_1342.dm_build_1283)
	dm_build_1342.dm_build_1289(1)

	return copy(dm_build_1342.dm_build_1283[dm_build_1344:], Dm_build_923.Dm_build_1122(dm_build_1343))
}

func (dm_build_1346 *Dm_build_1282) Dm_build_1345(dm_build_1347 uint16) int {
	dm_build_1348 := len(dm_build_1346.dm_build_1283)
	dm_build_1346.dm_build_1289(2)

	return copy(dm_build_1346.dm_build_1283[dm_build_1348:], Dm_build_923.Dm_build_1125(dm_build_1347))
}

func (dm_build_1350 *Dm_build_1282) Dm_build_1349(dm_build_1351 uint32) int {
	dm_build_1352 := len(dm_build_1350.dm_build_1283)
	dm_build_1350.dm_build_1289(4)

	return copy(dm_build_1350.dm_build_1283[dm_build_1352:], Dm_build_923.Dm_build_1128(dm_build_1351))
}

func (dm_build_1354 *Dm_build_1282) Dm_build_1353(dm_build_1355 uint64) int {
	dm_build_1356 := len(dm_build_1354.dm_build_1283)
	dm_build_1354.dm_build_1289(8)

	return copy(dm_build_1354.dm_build_1283[dm_build_1356:], Dm_build_923.Dm_build_1131(dm_build_1355))
}

func (dm_build_1358 *Dm_build_1282) Dm_build_1357(dm_build_1359 float32) int {
	dm_build_1360 := len(dm_build_1358.dm_build_1283)
	dm_build_1358.dm_build_1289(4)

	return copy(dm_build_1358.dm_build_1283[dm_build_1360:], Dm_build_923.Dm_build_1128(math.Float32bits(dm_build_1359)))
}

func (dm_build_1362 *Dm_build_1282) Dm_build_1361(dm_build_1363 float64) int {
	dm_build_1364 := len(dm_build_1362.dm_build_1283)
	dm_build_1362.dm_build_1289(8)

	return copy(dm_build_1362.dm_build_1283[dm_build_1364:], Dm_build_923.Dm_build_1131(math.Float64bits(dm_build_1363)))
}

func (dm_build_1366 *Dm_build_1282) Dm_build_1365(dm_build_1367 []byte) int {
	dm_build_1368 := len(dm_build_1366.dm_build_1283)
	dm_build_1366.dm_build_1289(len(dm_build_1367))
	return copy(dm_build_1366.dm_build_1283[dm_build_1368:], dm_build_1367)
}

func (dm_build_1370 *Dm_build_1282) Dm_build_1369(dm_build_1371 []byte) int {
	return dm_build_1370.Dm_build_1337(int32(len(dm_build_1371))) + dm_build_1370.Dm_build_1365(dm_build_1371)
}

func (dm_build_1373 *Dm_build_1282) Dm_build_1372(dm_build_1374 []byte) int {
	return dm_build_1373.Dm_build_1341(uint8(len(dm_build_1374))) + dm_build_1373.Dm_build_1365(dm_build_1374)
}

func (dm_build_1376 *Dm_build_1282) Dm_build_1375(dm_build_1377 []byte) int {
	return dm_build_1376.Dm_build_1345(uint16(len(dm_build_1377))) + dm_build_1376.Dm_build_1365(dm_build_1377)
}

func (dm_build_1379 *Dm_build_1282) Dm_build_1378(dm_build_1380 []byte) int {
	return dm_build_1379.Dm_build_1365(dm_build_1380) + dm_build_1379.Dm_build_1325(0)
}

func (dm_build_1382 *Dm_build_1282) Dm_build_1381(dm_build_1383 string, dm_build_1384 string, dm_build_1385 *DmConnection) int {
	dm_build_1386 := Dm_build_923.Dm_build_1139(dm_build_1383, dm_build_1384, dm_build_1385)
	return dm_build_1382.Dm_build_1369(dm_build_1386)
}

func (dm_build_1388 *Dm_build_1282) Dm_build_1387(dm_build_1389 string, dm_build_1390 string, dm_build_1391 *DmConnection) int {
	dm_build_1392 := Dm_build_923.Dm_build_1139(dm_build_1389, dm_build_1390, dm_build_1391)
	return dm_build_1388.Dm_build_1372(dm_build_1392)
}

func (dm_build_1394 *Dm_build_1282) Dm_build_1393(dm_build_1395 string, dm_build_1396 string, dm_build_1397 *DmConnection) int {
	dm_build_1398 := Dm_build_923.Dm_build_1139(dm_build_1395, dm_build_1396, dm_build_1397)
	return dm_build_1394.Dm_build_1375(dm_build_1398)
}

func (dm_build_1400 *Dm_build_1282) Dm_build_1399(dm_build_1401 string, dm_build_1402 string, dm_build_1403 *DmConnection) int {
	dm_build_1404 := Dm_build_923.Dm_build_1139(dm_build_1401, dm_build_1402, dm_build_1403)
	return dm_build_1400.Dm_build_1378(dm_build_1404)
}

func (dm_build_1406 *Dm_build_1282) Dm_build_1405() byte {
	dm_build_1407 := Dm_build_923.Dm_build_1016(dm_build_1406.dm_build_1283, dm_build_1406.dm_build_1284)
	dm_build_1406.dm_build_1284++
	return dm_build_1407
}

func (dm_build_1409 *Dm_build_1282) Dm_build_1408() int16 {
	dm_build_1410 := Dm_build_923.Dm_build_1020(dm_build_1409.dm_build_1283, dm_build_1409.dm_build_1284)
	dm_build_1409.dm_build_1284 += 2
	return dm_build_1410
}

func (dm_build_1412 *Dm_build_1282) Dm_build_1411() int32 {
	dm_build_1413 := Dm_build_923.Dm_build_1025(dm_build_1412.dm_build_1283, dm_build_1412.dm_build_1284)
	dm_build_1412.dm_build_1284 += 4
	return dm_build_1413
}

func (dm_build_1415 *Dm_build_1282) Dm_build_1414() int64 {
	dm_build_1416 := Dm_build_923.Dm_build_1030(dm_build_1415.dm_build_1283, dm_build_1415.dm_build_1284)
	dm_build_1415.dm_build_1284 += 8
	return dm_build_1416
}

func (dm_build_1418 *Dm_build_1282) Dm_build_1417() float32 {
	dm_build_1419 := Dm_build_923.Dm_build_1035(dm_build_1418.dm_build_1283, dm_build_1418.dm_build_1284)
	dm_build_1418.dm_build_1284 += 4
	return dm_build_1419
}

func (dm_build_1421 *Dm_build_1282) Dm_build_1420() float64 {
	dm_build_1422 := Dm_build_923.Dm_build_1039(dm_build_1421.dm_build_1283, dm_build_1421.dm_build_1284)
	dm_build_1421.dm_build_1284 += 8
	return dm_build_1422
}

func (dm_build_1424 *Dm_build_1282) Dm_build_1423() uint8 {
	dm_build_1425 := Dm_build_923.Dm_build_1043(dm_build_1424.dm_build_1283, dm_build_1424.dm_build_1284)
	dm_build_1424.dm_build_1284 += 1
	return dm_build_1425
}

func (dm_build_1427 *Dm_build_1282) Dm_build_1426() uint16 {
	dm_build_1428 := Dm_build_923.Dm_build_1047(dm_build_1427.dm_build_1283, dm_build_1427.dm_build_1284)
	dm_build_1427.dm_build_1284 += 2
	return dm_build_1428
}

func (dm_build_1430 *Dm_build_1282) Dm_build_1429() uint32 {
	dm_build_1431 := Dm_build_923.Dm_build_1052(dm_build_1430.dm_build_1283, dm_build_1430.dm_build_1284)
	dm_build_1430.dm_build_1284 += 4
	return dm_build_1431
}

func (dm_build_1433 *Dm_build_1282) Dm_build_1432(dm_build_1434 int) []byte {
	dm_build_1435 := Dm_build_923.Dm_build_1074(dm_build_1433.dm_build_1283, dm_build_1433.dm_build_1284, dm_build_1434)
	dm_build_1433.dm_build_1284 += dm_build_1434
	return dm_build_1435
}

func (dm_build_1437 *Dm_build_1282) Dm_build_1436() []byte {
	return dm_build_1437.Dm_build_1432(int(dm_build_1437.Dm_build_1411()))
}

func (dm_build_1439 *Dm_build_1282) Dm_build_1438() []byte {
	return dm_build_1439.Dm_build_1432(int(dm_build_1439.Dm_build_1405()))
}

func (dm_build_1441 *Dm_build_1282) Dm_build_1440() []byte {
	return dm_build_1441.Dm_build_1432(int(dm_build_1441.Dm_build_1408()))
}

func (dm_build_1443 *Dm_build_1282) Dm_build_1442(dm_build_1444 int) []byte {
	return dm_build_1443.Dm_build_1432(dm_build_1444)
}

func (dm_build_1446 *Dm_build_1282) Dm_build_1445() []byte {
	dm_build_1447 := 0
	for dm_build_1446.Dm_build_1405() != 0 {
		dm_build_1447++
	}
	dm_build_1446.Dm_build_1307(dm_build_1447, false, false)
	return dm_build_1446.Dm_build_1432(dm_build_1447)
}

func (dm_build_1449 *Dm_build_1282) Dm_build_1448(dm_build_1450 int, dm_build_1451 string, dm_build_1452 *DmConnection) string {
	return Dm_build_923.Dm_build_1175(dm_build_1449.Dm_build_1432(dm_build_1450), dm_build_1451, dm_build_1452)
}

func (dm_build_1454 *Dm_build_1282) Dm_build_1453(dm_build_1455 string, dm_build_1456 *DmConnection) string {
	return Dm_build_923.Dm_build_1175(dm_build_1454.Dm_build_1436(), dm_build_1455, dm_build_1456)
}

func (dm_build_1458 *Dm_build_1282) Dm_build_1457(dm_build_1459 string, dm_build_1460 *DmConnection) string {
	return Dm_build_923.Dm_build_1175(dm_build_1458.Dm_build_1438(), dm_build_1459, dm_build_1460)
}

func (dm_build_1462 *Dm_build_1282) Dm_build_1461(dm_build_1463 string, dm_build_1464 *DmConnection) string {
	return Dm_build_923.Dm_build_1175(dm_build_1462.Dm_build_1440(), dm_build_1463, dm_build_1464)
}

func (dm_build_1466 *Dm_build_1282) Dm_build_1465(dm_build_1467 string, dm_build_1468 *DmConnection) string {
	return Dm_build_923.Dm_build_1175(dm_build_1466.Dm_build_1445(), dm_build_1467, dm_build_1468)
}

func (dm_build_1470 *Dm_build_1282) Dm_build_1469(dm_build_1471 int, dm_build_1472 byte) int {
	return dm_build_1470.Dm_build_1505(dm_build_1471, Dm_build_923.Dm_build_1101(dm_build_1472))
}

func (dm_build_1474 *Dm_build_1282) Dm_build_1473(dm_build_1475 int, dm_build_1476 int16) int {
	return dm_build_1474.Dm_build_1505(dm_build_1475, Dm_build_923.Dm_build_1107(dm_build_1476))
}

func (dm_build_1478 *Dm_build_1282) Dm_build_1477(dm_build_1479 int, dm_build_1480 int32) int {
	return dm_build_1478.Dm_build_1505(dm_build_1479, Dm_build_923.Dm_build_1110(dm_build_1480))
}

func (dm_build_1482 *Dm_build_1282) Dm_build_1481(dm_build_1483 int, dm_build_1484 int64) int {
	return dm_build_1482.Dm_build_1505(dm_build_1483, Dm_build_923.Dm_build_1113(dm_build_1484))
}

func (dm_build_1486 *Dm_build_1282) Dm_build_1485(dm_build_1487 int, dm_build_1488 float32) int {
	return dm_build_1486.Dm_build_1505(dm_build_1487, Dm_build_923.Dm_build_1116(dm_build_1488))
}

func (dm_build_1490 *Dm_build_1282) Dm_build_1489(dm_build_1491 int, dm_build_1492 float64) int {
	return dm_build_1490.Dm_build_1505(dm_build_1491, Dm_build_923.Dm_build_1119(dm_build_1492))
}

func (dm_build_1494 *Dm_build_1282) Dm_build_1493(dm_build_1495 int, dm_build_1496 uint8) int {
	return dm_build_1494.Dm_build_1505(dm_build_1495, Dm_build_923.Dm_build_1122(dm_build_1496))
}

func (dm_build_1498 *Dm_build_1282) Dm_build_1497(dm_build_1499 int, dm_build_1500 uint16) int {
	return dm_build_1498.Dm_build_1505(dm_build_1499, Dm_build_923.Dm_build_1125(dm_build_1500))
}

func (dm_build_1502 *Dm_build_1282) Dm_build_1501(dm_build_1503 int, dm_build_1504 uint32) int {
	return dm_build_1502.Dm_build_1505(dm_build_1503, Dm_build_923.Dm_build_1128(dm_build_1504))
}

func (dm_build_1506 *Dm_build_1282) Dm_build_1505(dm_build_1507 int, dm_build_1508 []byte) int {
	return copy(dm_build_1506.dm_build_1283[dm_build_1507:], dm_build_1508)
}

func (dm_build_1510 *Dm_build_1282) Dm_build_1509(dm_build_1511 int, dm_build_1512 []byte) int {
	return dm_build_1510.Dm_build_1477(dm_build_1511, int32(len(dm_build_1512))) + dm_build_1510.Dm_build_1505(dm_build_1511+4, dm_build_1512)
}

func (dm_build_1514 *Dm_build_1282) Dm_build_1513(dm_build_1515 int, dm_build_1516 []byte) int {
	return dm_build_1514.Dm_build_1469(dm_build_1515, byte(len(dm_build_1516))) + dm_build_1514.Dm_build_1505(dm_build_1515+1, dm_build_1516)
}

func (dm_build_1518 *Dm_build_1282) Dm_build_1517(dm_build_1519 int, dm_build_1520 []byte) int {
	return dm_build_1518.Dm_build_1473(dm_build_1519, int16(len(dm_build_1520))) + dm_build_1518.Dm_build_1505(dm_build_1519+2, dm_build_1520)
}

func (dm_build_1522 *Dm_build_1282) Dm_build_1521(dm_build_1523 int, dm_build_1524 []byte) int {
	return dm_build_1522.Dm_build_1505(dm_build_1523, dm_build_1524) + dm_build_1522.Dm_build_1469(dm_build_1523+len(dm_build_1524), 0)
}

func (dm_build_1526 *Dm_build_1282) Dm_build_1525(dm_build_1527 int, dm_build_1528 string, dm_build_1529 string, dm_build_1530 *DmConnection) int {
	return dm_build_1526.Dm_build_1509(dm_build_1527, Dm_build_923.Dm_build_1139(dm_build_1528, dm_build_1529, dm_build_1530))
}

func (dm_build_1532 *Dm_build_1282) Dm_build_1531(dm_build_1533 int, dm_build_1534 string, dm_build_1535 string, dm_build_1536 *DmConnection) int {
	return dm_build_1532.Dm_build_1513(dm_build_1533, Dm_build_923.Dm_build_1139(dm_build_1534, dm_build_1535, dm_build_1536))
}

func (dm_build_1538 *Dm_build_1282) Dm_build_1537(dm_build_1539 int, dm_build_1540 string, dm_build_1541 string, dm_build_1542 *DmConnection) int {
	return dm_build_1538.Dm_build_1517(dm_build_1539, Dm_build_923.Dm_build_1139(dm_build_1540, dm_build_1541, dm_build_1542))
}

func (dm_build_1544 *Dm_build_1282) Dm_build_1543(dm_build_1545 int, dm_build_1546 string, dm_build_1547 string, dm_build_1548 *DmConnection) int {
	return dm_build_1544.Dm_build_1521(dm_build_1545, Dm_build_923.Dm_build_1139(dm_build_1546, dm_build_1547, dm_build_1548))
}

func (dm_build_1550 *Dm_build_1282) Dm_build_1549(dm_build_1551 int) byte {
	return Dm_build_923.Dm_build_1144(dm_build_1550.Dm_build_1576(dm_build_1551, 1))
}

func (dm_build_1553 *Dm_build_1282) Dm_build_1552(dm_build_1554 int) int16 {
	return Dm_build_923.Dm_build_1147(dm_build_1553.Dm_build_1576(dm_build_1554, 2))
}

func (dm_build_1556 *Dm_build_1282) Dm_build_1555(dm_build_1557 int) int32 {
	return Dm_build_923.Dm_build_1150(dm_build_1556.Dm_build_1576(dm_build_1557, 4))
}

func (dm_build_1559 *Dm_build_1282) Dm_build_1558(dm_build_1560 int) int64 {
	return Dm_build_923.Dm_build_1153(dm_build_1559.Dm_build_1576(dm_build_1560, 8))
}

func (dm_build_1562 *Dm_build_1282) Dm_build_1561(dm_build_1563 int) float32 {
	return Dm_build_923.Dm_build_1156(dm_build_1562.Dm_build_1576(dm_build_1563, 4))
}

func (dm_build_1565 *Dm_build_1282) Dm_build_1564(dm_build_1566 int) float64 {
	return Dm_build_923.Dm_build_1159(dm_build_1565.Dm_build_1576(dm_build_1566, 8))
}

func (dm_build_1568 *Dm_build_1282) Dm_build_1567(dm_build_1569 int) uint8 {
	return Dm_build_923.Dm_build_1162(dm_build_1568.Dm_build_1576(dm_build_1569, 1))
}

func (dm_build_1571 *Dm_build_1282) Dm_build_1570(dm_build_1572 int) uint16 {
	return Dm_build_923.Dm_build_1165(dm_build_1571.Dm_build_1576(dm_build_1572, 2))
}

func (dm_build_1574 *Dm_build_1282) Dm_build_1573(dm_build_1575 int) uint32 {
	return Dm_build_923.Dm_build_1168(dm_build_1574.Dm_build_1576(dm_build_1575, 4))
}

func (dm_build_1577 *Dm_build_1282) Dm_build_1576(dm_build_1578 int, dm_build_1579 int) []byte {
	return dm_build_1577.dm_build_1283[dm_build_1578 : dm_build_1578+dm_build_1579]
}

func (dm_build_1581 *Dm_build_1282) Dm_build_1580(dm_build_1582 int) []byte {
	dm_build_1583 := dm_build_1581.Dm_build_1555(dm_build_1582)
	return dm_build_1581.Dm_build_1576(dm_build_1582+4, int(dm_build_1583))
}

func (dm_build_1585 *Dm_build_1282) Dm_build_1584(dm_build_1586 int) []byte {
	dm_build_1587 := dm_build_1585.Dm_build_1549(dm_build_1586)
	return dm_build_1585.Dm_build_1576(dm_build_1586+1, int(dm_build_1587))
}

func (dm_build_1589 *Dm_build_1282) Dm_build_1588(dm_build_1590 int) []byte {
	dm_build_1591 := dm_build_1589.Dm_build_1552(dm_build_1590)
	return dm_build_1589.Dm_build_1576(dm_build_1590+2, int(dm_build_1591))
}

func (dm_build_1593 *Dm_build_1282) Dm_build_1592(dm_build_1594 int) []byte {
	dm_build_1595 := 0
	for dm_build_1593.Dm_build_1549(dm_build_1594) != 0 {
		dm_build_1594++
		dm_build_1595++
	}

	return dm_build_1593.Dm_build_1576(dm_build_1594-dm_build_1595, int(dm_build_1595))
}

func (dm_build_1597 *Dm_build_1282) Dm_build_1596(dm_build_1598 int, dm_build_1599 string, dm_build_1600 *DmConnection) string {
	return Dm_build_923.Dm_build_1175(dm_build_1597.Dm_build_1580(dm_build_1598), dm_build_1599, dm_build_1600)
}

func (dm_build_1602 *Dm_build_1282) Dm_build_1601(dm_build_1603 int, dm_build_1604 string, dm_build_1605 *DmConnection) string {
	return Dm_build_923.Dm_build_1175(dm_build_1602.Dm_build_1584(dm_build_1603), dm_build_1604, dm_build_1605)
}

func (dm_build_1607 *Dm_build_1282) Dm_build_1606(dm_build_1608 int, dm_build_1609 string, dm_build_1610 *DmConnection) string {
	return Dm_build_923.Dm_build_1175(dm_build_1607.Dm_build_1588(dm_build_1608), dm_build_1609, dm_build_1610)
}

func (dm_build_1612 *Dm_build_1282) Dm_build_1611(dm_build_1613 int, dm_build_1614 string, dm_build_1615 *DmConnection) string {
	return Dm_build_923.Dm_build_1175(dm_build_1612.Dm_build_1592(dm_build_1613), dm_build_1614, dm_build_1615)
}
