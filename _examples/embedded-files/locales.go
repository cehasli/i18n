// Package main Code generated by go-bindata. (@generated) DO NOT EDIT.
// sources:
// ../basic/locales/el-GR/example.yml
// ../basic/locales/en-US/example.yml
// ../basic/locales/zh-CN/example.yml
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
		return nil, fmt.Errorf("read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	clErr := gz.Close()

	if err != nil {
		return nil, fmt.Errorf("read %q: %v", name, err)
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

// Name return file name
func (fi bindataFileInfo) Name() string {
	return fi.name
}

// Size return file size
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}

// Mode return file mode
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}

// ModTime return file modify time
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}

// IsDir return file whether a directory
func (fi bindataFileInfo) IsDir() bool {
	return fi.mode&os.ModeDir != 0
}

// Sys return file is sys mode
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _BasicLocalesElGrExampleYml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xca\xc8\xb4\x52\x50\x3a\x37\xf9\xdc\xd6\x73\x3b\xcf\xad\x51\x50\x2d\x56\xe2\xe5\xca\xcc\x2b\x29\xca\x07\x89\xce\x39\xb7\x55\xe1\xdc\xee\x73\x6b\xcf\xed\x3d\xb7\x55\xa1\xba\x5a\xcf\x2f\x31\x37\xb5\xb6\x56\xe1\xdc\xae\x73\x1b\x15\xce\x6d\x3d\xb7\xfe\xdc\x9e\x73\x1b\xcf\xed\x04\xc9\x38\xa6\x83\x24\xce\xb7\x9f\x6f\x3c\xb7\xff\xdc\xde\xf3\x7d\xe7\xf6\x2a\x01\x02\x00\x00\xff\xff\x0e\x50\xb0\xc8\x59\x00\x00\x00")

func BasicLocalesElGrExampleYmlBytes() ([]byte, error) {
	return bindataRead(
		_BasicLocalesElGrExampleYml,
		"../basic/locales/el-GR/example.yml",
	)
}

func BasicLocalesElGrExampleYml() (*asset, error) {
	bytes, err := BasicLocalesElGrExampleYmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "../basic/locales/el-GR/example.yml", size: 89, mode: os.FileMode(438), modTime: time.Unix(1575080794, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _BasicLocalesEnUsExampleYml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xca\xc8\xb4\x52\x50\xf2\xc8\x54\x50\x2d\x56\xe2\xe5\xca\xcc\x2b\x29\xca\xb7\x52\x50\xf2\xad\x54\xc8\x4b\xcc\x4d\x55\xc8\x2c\x56\xa8\xae\xd6\xf3\x4b\xcc\x4d\xad\xad\x55\x48\xcc\x4b\x51\xf0\x54\x48\xcc\x05\x09\x39\xa6\x83\x44\x2a\x53\x13\x8b\x8a\x15\xf2\x73\x52\x94\x00\x01\x00\x00\xff\xff\x23\xcc\x6d\x0c\x46\x00\x00\x00")

func BasicLocalesEnUsExampleYmlBytes() ([]byte, error) {
	return bindataRead(
		_BasicLocalesEnUsExampleYml,
		"../basic/locales/en-US/example.yml",
	)
}

func BasicLocalesEnUsExampleYml() (*asset, error) {
	bytes, err := BasicLocalesEnUsExampleYmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "../basic/locales/en-US/example.yml", size: 70, mode: os.FileMode(438), modTime: time.Unix(1575080796, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _BasicLocalesZhCnExampleYml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xca\xc8\xb4\x52\x50\x7a\xd6\xb4\xe2\xe9\xd2\xbd\x0a\xaa\xc5\x4a\xbc\x5c\x99\x79\x25\x45\xf9\x20\xb1\x8e\x89\x4f\xfb\x57\x2b\x54\x57\xeb\xf9\x25\xe6\xa6\xd6\xd6\xbe\xdf\xd3\xf3\x64\x77\xd7\xd3\x9d\x5b\x40\x42\x8e\xe9\xa9\xb5\xb5\x0a\x4f\x37\x35\x2a\x01\x02\x00\x00\xff\xff\xe8\xa2\x24\x0f\x40\x00\x00\x00")

func BasicLocalesZhCnExampleYmlBytes() ([]byte, error) {
	return bindataRead(
		_BasicLocalesZhCnExampleYml,
		"../basic/locales/zh-CN/example.yml",
	)
}

func BasicLocalesZhCnExampleYml() (*asset, error) {
	bytes, err := BasicLocalesZhCnExampleYmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "../basic/locales/zh-CN/example.yml", size: 64, mode: os.FileMode(438), modTime: time.Unix(1575080798, 0)}
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
	"../basic/locales/el-GR/example.yml": BasicLocalesElGrExampleYml,
	"../basic/locales/en-US/example.yml": BasicLocalesEnUsExampleYml,
	"../basic/locales/zh-CN/example.yml": BasicLocalesZhCnExampleYml,
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
	"..": {nil, map[string]*bintree{
		"basic": {nil, map[string]*bintree{
			"locales": {nil, map[string]*bintree{
				"el-GR": {nil, map[string]*bintree{
					"example.yml": {BasicLocalesElGrExampleYml, map[string]*bintree{}},
				}},
				"en-US": {nil, map[string]*bintree{
					"example.yml": {BasicLocalesEnUsExampleYml, map[string]*bintree{}},
				}},
				"zh-CN": {nil, map[string]*bintree{
					"example.yml": {BasicLocalesZhCnExampleYml, map[string]*bintree{}},
				}},
			}},
		}},
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
