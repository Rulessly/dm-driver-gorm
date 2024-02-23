package driver

import (
	"bytes"
	"compress/zlib"
	"github.com/golang/snappy"
)

func Compress(srcBuffer *Dm_build_1282, offset int, length int, compressID int) ([]byte, error) {
	if compressID == Dm_build_366 {
		return snappy.Encode(nil, srcBuffer.Dm_build_1576(offset, length)), nil
	}
	return GzlibCompress(srcBuffer, offset, length)
}

func UnCompress(srcBytes []byte, compressID int) ([]byte, error) {
	if compressID == Dm_build_366 {
		return snappy.Decode(nil, srcBytes)
	}
	return GzlibUncompress(srcBytes)
}

func GzlibCompress(srcBuffer *Dm_build_1282, offset int, length int) ([]byte, error) {
	var ret bytes.Buffer
	var w = zlib.NewWriter(&ret)
	w.Write(srcBuffer.Dm_build_1576(offset, length))
	w.Close()
	return ret.Bytes(), nil
}

func GzlibUncompress(srcBytes []byte) ([]byte, error) {
	var bytesBuf = new(bytes.Buffer)
	r, err := zlib.NewReader(bytes.NewReader(srcBytes))
	if err != nil {
		return nil, err
	}
	defer r.Close()
	_, err = bytesBuf.ReadFrom(r)
	if err != nil {
		return nil, err
	}
	return bytesBuf.Bytes(), nil
}
