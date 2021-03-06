// Code generated by go-bindata.
// sources:
// res/font.ini
// res/icon.ini
// res/lcd.ini
// res/res.ini
// DO NOT EDIT!

package main

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

var _resFontIni = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xbc\x56\x3b\x4e\xc4\x30\x14\xec\x23\xe5\x0e\xf4\x08\x34\x2f\x9b\x1f\x48\xa9\x38\xc6\x8a\x9b\x70\x09\xe8\x11\x9f\x86\x96\x9e\x02\x89\x13\xe4\x26\x5c\x81\x6c\x9c\xbc\x8c\x9c\x38\xfb\xe2\x48\xb8\x61\x14\x59\x1e\xe6\xf3\xec\xfd\xfd\xfa\x3e\xe2\x3e\x4d\xa4\x81\x9c\x16\x90\x26\x59\xd3\xfd\xe9\x96\x74\xf8\x40\x38\x27\x5c\x10\x2e\x09\x57\x74\x4e\xdd\x60\x58\x69\x92\x26\x47\x71\x34\xba\x35\x73\x58\x1c\x0d\x88\x06\x44\x03\xa2\x01\xd3\x28\xf6\x68\xb2\x55\x35\x1e\xcd\xb0\xa7\x53\x23\x32\x1c\x31\xaa\x41\x4f\x23\x6e\xcd\x69\x0e\xe7\x69\xfa\xef\x27\x1a\x77\xb8\xaa\x11\xd8\x4d\xcb\x7b\x1a\x21\xd3\x82\xd9\xe8\x7f\x1a\x61\x5a\xe1\x68\xf4\x88\x8c\x4c\x38\x0c\xdf\xb1\x9a\x8d\x49\x4d\x69\x68\x1a\x48\x0d\xe2\x9a\x56\xcd\xd4\x20\x5c\x81\x45\x35\x26\xd3\x6a\xf3\xdc\x60\x8f\x9a\x9b\x8d\x34\x91\xd9\xdc\xd2\x78\x82\x4c\xc3\x68\xda\x50\xe2\x9c\x70\x41\x7b\x4a\xfa\x5e\x11\xf6\x68\xda\x97\x40\xd5\xe2\x6e\x1b\x09\x85\xd3\x3e\x47\x4d\xce\x76\x9e\xa7\x49\x0f\x28\x1e\xc1\xc4\x23\xf6\x4e\x4b\x28\x9e\xf6\x7d\xa9\xd4\xc2\xf9\x80\xf2\xc1\x94\x8f\x70\x3e\xa0\x7c\xb0\x98\xcf\x9b\x7d\x46\x15\x17\x84\x8d\x75\x6b\x5f\x49\xcf\x86\x57\x47\xf7\x7b\xf7\xb4\x62\x9f\xe7\xc3\xdc\x83\xd0\xf8\x80\xf5\x04\x7b\xf0\xb8\x7e\x85\xb2\x6f\x7b\x7a\x70\x71\x6e\x4c\x95\x06\x14\x0f\x28\x1e\x90\x6d\x08\xd9\x76\xbd\x9f\x46\x5f\xd4\x8a\x70\x1d\xa6\xfc\xf9\x1c\x9b\xa7\x9c\x42\x0d\x17\x6e\xb8\xee\x89\x90\x76\xf7\x3f\x05\xbf\x1c\x1d\x14\x72\x70\x71\x5e\xbd\x97\x7b\xdb\xbc\x5e\x99\x83\x9a\xff\x40\xd8\x60\xda\x83\x59\x0d\x22\x6f\x9f\xbf\x00\x00\x00\xff\xff\xf2\xe8\xbc\x5c\x8a\x0a\x00\x00")

func resFontIniBytes() ([]byte, error) {
	return bindataRead(
		_resFontIni,
		"res/font.ini",
	)
}

func resFontIni() (*asset, error) {
	bytes, err := resFontIniBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "res/font.ini", size: 2698, mode: os.FileMode(438), modTime: time.Unix(1469216550, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _resIconIni = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xb4\x93\x4d\xae\x87\x20\x0c\xc4\xf7\x4d\x7a\x97\x8e\xdf\x9b\x9e\xc4\xbc\xfb\x5f\xe3\x21\x50\xb1\xd0\xff\xca\x58\x12\x1b\xc6\x02\x3f\x47\x7a\xca\x1f\x13\x54\x04\x8f\x10\x61\x9a\x54\x52\x6e\x81\xa4\xcd\xea\xa4\x24\x32\x2d\x8a\x2e\x98\x56\x95\x2e\x98\xb6\xeb\x8c\xbc\x4f\x79\x5e\xda\xae\x10\x3f\x98\x0e\xb5\xf7\x56\xcf\xc4\x74\xe2\x3d\xe6\xa0\x25\x4c\x4f\x5e\x31\x7b\xf4\x3d\xd0\x8e\x40\x4b\x98\x53\xc5\xb4\x2f\x28\x29\x63\xde\x33\xd4\xf2\xb9\x38\x77\xcd\x21\xd5\xb9\xc5\xea\x8a\x9e\xeb\xd6\x40\xdb\x82\xb5\x7b\x70\xc6\x11\xb0\x24\xcc\xf9\xb5\x9b\xa3\x73\x8f\x9f\x9e\x77\x13\xef\x26\xd0\xbb\x59\xab\x9c\x9b\x70\x6e\x2e\x9f\x60\xe2\xbe\x5d\xcd\xcd\x6a\xa4\x58\xb6\xbb\xf9\xd4\x9b\x9b\x6d\x7d\xc6\x5c\x3f\x69\xa1\xb1\x35\xb6\x40\xfb\xd5\x42\x7e\xfc\x07\x00\x00\xff\xff\xf4\x9f\xf8\x7e\xe6\x03\x00\x00")

func resIconIniBytes() ([]byte, error) {
	return bindataRead(
		_resIconIni,
		"res/icon.ini",
	)
}

func resIconIni() (*asset, error) {
	bytes, err := resIconIniBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "res/icon.ini", size: 998, mode: os.FileMode(438), modTime: time.Unix(1469220867, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _resLcdIni = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x9c\x93\x39\xce\xc2\x30\x14\x84\xfb\x48\xb9\xc3\xdf\xff\x02\xcd\xcb\xc2\x52\xa4\xa2\x04\x89\x03\x44\x1c\x0a\x7a\xc4\xd2\xd0\x72\x06\x4e\x90\x23\x91\xc4\x19\x87\x38\x36\x36\xa1\x1a\x99\x7c\x9a\x37\xe3\xe7\x12\x87\x38\x92\x42\xea\x1f\xe2\x28\x29\x04\x68\x44\x4a\x91\x51\xe4\xfc\x26\x8e\x4a\x69\x19\xa8\x3f\x12\x8a\x94\x22\xa3\xc8\x29\x6a\x26\x19\xfa\x68\xa6\x3b\x69\x7d\x60\xf8\xa4\x3e\x06\xe3\xd9\x32\xc5\x70\x36\x71\x33\xfd\x6c\xf9\xa8\x03\x78\x7d\x16\x3e\xc6\xd2\xdb\xd2\xd7\xb5\x65\xb6\x95\x8b\xf9\xe2\xb3\xf6\x31\x96\x3c\x1b\x57\x1e\x25\xac\xf7\xb3\x63\xd7\x08\x67\xf6\x13\xf6\x6d\xeb\xbe\x53\x18\x8c\xee\x6d\xce\x1d\x05\x77\x07\xdc\x51\xb0\x03\x90\x81\x62\xaa\x4b\xf8\x70\xda\xa8\x3a\xfb\x36\xce\x06\x9d\xe8\x04\x6f\x24\x75\xd2\x40\xf7\x21\x04\x61\x26\x61\x26\xe9\x32\x75\xa2\x81\x6e\xaa\x88\x80\x8b\x85\x2e\xbc\xba\x06\x8c\x37\x6a\xef\x11\x5c\x39\x7a\xe8\xe8\x7e\x47\xce\x22\x5e\x4f\x73\x3c\x31\x20\x8b\xd3\x94\x05\xff\x37\x13\x09\x0b\xd7\x0f\x49\x4c\x9f\x99\xab\x85\x8f\xc7\x67\x32\x7f\xbf\x35\xf7\x0e\x00\x00\xff\xff\xc9\xbb\x4b\x59\xb5\x05\x00\x00")

func resLcdIniBytes() ([]byte, error) {
	return bindataRead(
		_resLcdIni,
		"res/lcd.ini",
	)
}

func resLcdIni() (*asset, error) {
	bytes, err := resLcdIniBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "res/lcd.ini", size: 1461, mode: os.FileMode(438), modTime: time.Unix(1469219124, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _resResIni = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xf2\xcd\xcf\x4b\x49\xac\xb4\xbd\x30\xff\xc2\x5c\x5e\xae\x90\xd2\xd4\x62\x30\x6f\xd2\x85\x45\xbc\x5c\xe1\xa9\x29\x79\x50\xfe\xc2\x0b\x0b\x80\xb2\x19\xa5\x45\x10\xee\x72\x90\xb4\x5b\x51\x26\x54\x27\x90\x13\x9c\x58\x52\x5a\x04\x55\x3a\x11\xc8\x2d\xcd\x83\x9a\xb3\x10\x10\x00\x00\xff\xff\x1f\x03\xbe\x9f\x61\x00\x00\x00")

func resResIniBytes() ([]byte, error) {
	return bindataRead(
		_resResIni,
		"res/res.ini",
	)
}

func resResIni() (*asset, error) {
	bytes, err := resResIniBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "res/res.ini", size: 97, mode: os.FileMode(438), modTime: time.Unix(1469213540, 0)}
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
	"res/font.ini": resFontIni,
	"res/icon.ini": resIconIni,
	"res/lcd.ini": resLcdIni,
	"res/res.ini": resResIni,
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
	"res": &bintree{nil, map[string]*bintree{
		"font.ini": &bintree{resFontIni, map[string]*bintree{}},
		"icon.ini": &bintree{resIconIni, map[string]*bintree{}},
		"lcd.ini": &bintree{resLcdIni, map[string]*bintree{}},
		"res.ini": &bintree{resResIni, map[string]*bintree{}},
	}},
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

