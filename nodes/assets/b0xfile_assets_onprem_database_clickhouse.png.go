// Code generaTed by fileb0x at "2023-09-21 09:26:25.666293644 -0500 CDT m=+1.303630259" from config file "b0x.yml" DO NOT EDIT.
// modified(2023-09-21 08:15:37.708143051 -0500 CDT)
// original path: ../../assets/onprem/database/clickhouse.png

package assets

import (
	"bytes"
	"compress/gzip"
	"io"

	"os"
)

// FileAssetsOnpremDatabaseClickhousePng is "assets/onprem/database/clickhouse.png"
var FileAssetsOnpremDatabaseClickhousePng = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xea\x0c\xf0\x73\xe7\xe5\x92\xe2\x62\x60\x60\xe0\xf5\xf4\x70\x09\x62\x60\x60\x64\x00\x61\x16\x66\x06\x06\x86\x75\x31\x5b\x43\x19\x18\x18\x98\x8b\x9d\x3c\x43\x38\x38\x38\x6e\x3f\xf4\x7f\xc0\xc0\xc0\xc0\x59\xe0\x11\x59\xcc\xc0\xe0\x79\x09\x84\x19\x57\xb8\x76\xfd\x60\x60\x60\x90\x2c\x71\x8d\x28\x09\xce\x4f\x2b\x29\x4f\x2c\x4a\x65\x28\x2f\x2f\xd7\xcb\xcc\xcb\x2e\x4e\x4e\x2c\x48\xd5\xcb\x2f\x4a\x9f\xfd\xce\x46\x8a\x81\x81\x41\x22\xc0\x27\xc4\x95\x81\x81\xe1\xff\x09\x86\xff\xe7\x18\xfe\x9f\x61\xf8\xcf\x80\x20\xf3\x6d\xba\x1c\x18\x18\x18\xd8\x4a\x82\xfc\x82\x19\x64\x2c\x43\x8b\x8a\xfa\x1f\xe8\xb2\x31\x30\x30\x1c\xf4\x74\x71\x0c\xa9\x98\xf3\xf6\xdc\x42\xde\x06\x06\x01\x96\x83\xfc\x9c\xa6\x09\x2f\x99\xbc\x39\x43\x1a\xd5\x39\x6e\x1d\x2f\xe2\xd0\x79\x94\xfe\xf6\xe9\xa4\x95\xda\xb7\x5e\xdd\x14\xb9\x2d\x79\x7b\xc2\x71\xcb\x63\xaa\xdb\x16\xef\x63\x62\x20\x19\xbc\x60\x97\xf8\xcc\x81\x5d\xea\x82\xc0\x07\x3e\xc2\x06\x3c\x28\x3f\x20\xb7\x16\x4d\xec\x2f\x83\xb9\x43\xe3\x4d\xb5\x4f\x27\x1d\x42\x40\x5c\x4f\x57\x3f\x97\x75\x4e\x09\x4d\x80\x00\x00\x00\xff\xff\xc2\x33\xf1\x40\x79\x01\x00\x00")

func init() {

	rb := bytes.NewReader(FileAssetsOnpremDatabaseClickhousePng)
	r, err := gzip.NewReader(rb)
	if err != nil {
		panic(err)
	}

	err = r.Close()
	if err != nil {
		panic(err)
	}

	f, err := FS.OpenFile(CTX, "assets/onprem/database/clickhouse.png", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0777)
	if err != nil {
		panic(err)
	}

	_, err = io.Copy(f, r)
	if err != nil {
		panic(err)
	}

	err = f.Close()
	if err != nil {
		panic(err)
	}
}
