// Code generated by go-bindata.
// sources:
// assets/10-eksclt.al2.conf
// assets/bootstrap.al2.sh
// DO NOT EDIT!

package nodebootstrap

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func bindataRead(data []byte, name string) ([]byte, error) {
	gz, err := gzip.NewReader(bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	clErr := gz.Close()

	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}
	if clErr != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

type asset struct {
	bytes []byte
	info  os.FileInfo
}

type bindataFileInfo struct {
	name    string
	size    int64
	mode    os.FileMode
	modTime time.Time
}

func (fi bindataFileInfo) Name() string {
	return fi.name
}
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}
func (fi bindataFileInfo) IsDir() bool {
	return false
}
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var __10EkscltAl2Conf = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x8c\x54\x4f\x4f\xfb\x46\x10\xbd\xfb\x53\xac\x04\x87\x56\x62\x6d\xc8\x2f\xe5\x80\xe4\x43\x44\x02\x42\x4d\x01\x11\x10\x95\xda\x2a\x9a\xec\x8e\xc3\x28\xf6\xac\x35\x3b\x4e\xf8\x23\xbe\x7b\xe5\x38\x46\x01\xaa\xea\x27\x5f\xbc\xfb\x66\xdf\xbe\x79\xf3\xec\x03\x83\xab\xe8\xb4\xb4\xb1\x46\x47\x05\x39\x13\x5f\xa2\x62\xe5\x8d\x97\x50\x5b\x62\xd3\x30\xa9\x29\x82\x98\x55\xb3\xc0\x12\xf5\x68\xbb\x18\x55\xf0\x1a\xd8\x4c\x89\x9b\x67\x33\x30\xbf\x8c\xa6\x83\x5f\x93\xe4\xaf\x19\xca\x9a\x1c\xfe\x93\x1c\x98\x69\x70\x50\x9a\x0a\x15\x3c\x28\x98\x1a\x04\x2a\x54\x94\x78\x66\xee\x26\x97\x57\x37\xd7\x47\x66\xf4\x38\x9b\x8f\x27\x17\xa3\x87\xe9\xfd\xbc\xdb\x4b\x26\xbc\x26\x09\x5c\x21\xeb\x05\x95\x98\x67\xa8\x2e\xeb\x24\x66\x3d\x57\x8a\xbc\x4e\x0e\xcc\x65\x19\x16\x50\x1a\x60\x6f\xa2\x82\x92\xfb\x74\xc7\x1f\xa3\x3f\xe7\xb7\x37\xe3\xd9\x91\x39\x9f\x3e\xcc\xee\x27\x77\xf3\xf1\xf5\xec\x7f\xe9\x77\xfd\xed\xd8\x3b\xf9\x1c\xd8\xfe\x07\xf9\xf5\xcd\x78\x32\xbf\xba\xfd\x29\xba\xb2\x25\xda\x92\x26\x93\x67\x74\x33\x05\xd1\x7c\xef\x35\x6b\xa2\x64\x0b\xe2\xfe\x80\xf9\x3b\x31\xc6\x5a\xf0\x5e\x30\xc6\xfc\x38\xdd\x3e\xbb\x5d\x0e\x1e\x2d\xd5\xf9\xe1\xdb\x4e\xc2\xfb\x0e\x70\x65\x13\x15\xc5\x7a\x8e\xf9\xe1\xdb\x5e\xcb\x7d\x41\x05\xcf\xb6\x0e\xbe\x45\x7b\x6b\x7a\x08\x1a\x7d\x42\x56\x72\xa0\x14\xd8\x6a\x58\x21\xdb\x0d\x2e\x9e\x42\x58\xed\x95\x04\xa1\xd7\xae\xa2\x0a\x1e\xf3\xc7\xcf\x05\x65\x19\x36\xb6\x16\x5a\x53\x89\x4b\xf4\xb9\x4a\x83\x3b\xac\x0e\xde\x12\x17\x02\xd6\x05\x56\x20\x46\xb1\x54\xc1\x12\xf3\xd3\xe3\xc1\xf0\xf8\xe4\x64\xf8\x63\xf8\xdb\x20\xf5\x2b\x49\xd1\x49\x7a\xf8\xf6\x3d\x17\xef\x29\x6c\x03\x07\x9b\x98\xba\x50\xb5\x1e\x67\x35\x34\x11\x2d\x54\xfe\x74\x78\xf6\x23\x3d\xf9\x30\x22\x34\xde\xd6\x12\xd6\xe4\x51\x72\xd8\xc4\xaf\x0e\x85\x0a\x88\xf3\xdd\xb2\x1b\x4f\x5f\xc2\x64\x17\xc4\xd6\x93\xe4\x59\xa8\x35\x73\x4c\xed\x68\xf6\x60\x17\xb8\xe8\xf0\x76\xd4\x2d\xce\xa8\xa9\xef\x2b\x3e\xfa\x93\x86\x95\x2a\xcc\x7d\x70\x2b\x94\x7e\x7a\xa8\x9b\x20\x2b\x5b\x97\xcd\xb2\x95\xc0\xd4\x9f\x5b\x4a\x68\x6a\xeb\x85\xd6\x28\x79\xb7\x2a\x7a\xe1\x82\x4b\xda\x2a\x6f\x87\xbf\xef\x6b\x1b\x98\x56\x0f\x2d\xbf\x05\xaf\xdb\x4e\x5f\xa0\xea\x7b\x2b\x10\xb4\x11\xb4\x4b\x50\x8c\xf9\x5d\x50\x50\xfc\xbd\x4b\x5c\xfb\xd1\xa2\x9c\xa3\x68\xfb\x03\x00\xfd\x74\x09\x70\xe0\x97\x2a\x34\x71\x9b\x81\xbc\x80\x32\xe2\x87\xa3\x84\xac\xd6\x81\x2d\xbe\x86\xdf\x41\xea\x44\x93\x7f\x03\x00\x00\xff\xff\xa8\xe6\x02\x07\x5e\x04\x00\x00")

func _10EkscltAl2ConfBytes() ([]byte, error) {
	return bindataRead(
		__10EkscltAl2Conf,
		"10-eksclt.al2.conf",
	)
}

func _10EkscltAl2Conf() (*asset, error) {
	bytes, err := _10EkscltAl2ConfBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "10-eksclt.al2.conf", size: 1118, mode: os.FileMode(420), modTime: time.Unix(1, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _bootstrapAl2Sh = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x4c\xcb\x31\x0e\xc2\x30\x0c\x46\xe1\x3d\xa7\x30\x85\x01\x86\x34\x27\x80\x09\x06\x16\xe0\x06\xc8\x49\x7f\x29\x55\x9d\x44\xaa\x5d\x24\x6e\xcf\x82\xaa\xae\xdf\xd3\xdb\xef\x42\x1c\x6b\xd0\x4c\x1e\x8b\x73\x48\xb9\x51\xf7\x78\x5e\x6f\xef\xfb\xeb\x7c\x38\xe6\xa6\x56\xb9\x80\xfc\x78\xea\xe8\x42\x01\x96\x02\x26\x4d\x26\x61\x5a\x22\x04\xd6\x4b\x4b\x2c\x3d\xea\xc7\x39\xfd\xaa\xa1\x24\x13\x1a\x18\xa5\x55\x3f\x43\x1a\x0f\x1b\x47\xe5\x28\xa0\xff\xbb\x09\x6a\x3c\xdb\xea\xbf\x00\x00\x00\xff\xff\xef\x6e\xff\x33\x97\x00\x00\x00")

func bootstrapAl2ShBytes() ([]byte, error) {
	return bindataRead(
		_bootstrapAl2Sh,
		"bootstrap.al2.sh",
	)
}

func bootstrapAl2Sh() (*asset, error) {
	bytes, err := bootstrapAl2ShBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "bootstrap.al2.sh", size: 151, mode: os.FileMode(420), modTime: time.Unix(1, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// MustAsset is like Asset but panics when Asset would return an error.
// It simplifies safe initialization of global variables.
func MustAsset(name string) []byte {
	a, err := Asset(name)
	if err != nil {
		panic("asset: Asset(" + name + "): " + err.Error())
	}

	return a
}

// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func AssetInfo(name string) (os.FileInfo, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, fmt.Errorf("AssetInfo %s not found", name)
}

// AssetNames returns the names of the assets.
func AssetNames() []string {
	names := make([]string, 0, len(_bindata))
	for name := range _bindata {
		names = append(names, name)
	}
	return names
}

// _bindata is a table, holding each asset generator, mapped to its name.
var _bindata = map[string]func() (*asset, error){
	"10-eksclt.al2.conf": _10EkscltAl2Conf,
	"bootstrap.al2.sh": bootstrapAl2Sh,
}

// AssetDir returns the file names below a certain
// directory embedded in the file by go-bindata.
// For example if you run go-bindata on data/... and data contains the
// following hierarchy:
//     data/
//       foo.txt
//       img/
//         a.png
//         b.png
// then AssetDir("data") would return []string{"foo.txt", "img"}
// AssetDir("data/img") would return []string{"a.png", "b.png"}
// AssetDir("foo.txt") and AssetDir("notexist") would return an error
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		cannonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(cannonicalName, "/")
		for _, p := range pathList {
			node = node.Children[p]
			if node == nil {
				return nil, fmt.Errorf("Asset %s not found", name)
			}
		}
	}
	if node.Func != nil {
		return nil, fmt.Errorf("Asset %s not found", name)
	}
	rv := make([]string, 0, len(node.Children))
	for childName := range node.Children {
		rv = append(rv, childName)
	}
	return rv, nil
}

type bintree struct {
	Func     func() (*asset, error)
	Children map[string]*bintree
}
var _bintree = &bintree{nil, map[string]*bintree{
	"10-eksclt.al2.conf": &bintree{_10EkscltAl2Conf, map[string]*bintree{}},
	"bootstrap.al2.sh": &bintree{bootstrapAl2Sh, map[string]*bintree{}},
}}

// RestoreAsset restores an asset under the given directory
func RestoreAsset(dir, name string) error {
	data, err := Asset(name)
	if err != nil {
		return err
	}
	info, err := AssetInfo(name)
	if err != nil {
		return err
	}
	err = os.MkdirAll(_filePath(dir, filepath.Dir(name)), os.FileMode(0755))
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(_filePath(dir, name), data, info.Mode())
	if err != nil {
		return err
	}
	err = os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
	if err != nil {
		return err
	}
	return nil
}

// RestoreAssets restores an asset under the given directory recursively
func RestoreAssets(dir, name string) error {
	children, err := AssetDir(name)
	// File
	if err != nil {
		return RestoreAsset(dir, name)
	}
	// Dir
	for _, child := range children {
		err = RestoreAssets(dir, filepath.Join(name, child))
		if err != nil {
			return err
		}
	}
	return nil
}

func _filePath(dir, name string) string {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(cannonicalName, "/")...)...)
}
