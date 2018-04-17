// Code generated by go-bindata.
// sources:
// resources/en.yml
// DO NOT EDIT!

package resources

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

var _resourcesEnYml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x84\x94\xdb\x92\x1a\x21\x10\x86\xef\xf7\x29\xba\x2a\xd7\xd9\x1a\xb3\x47\xe7\x36\x4f\x82\xd2\x93\x21\x41\xb0\xa0\xd1\x6c\x9e\x3e\xa5\xce\x56\x6d\x1f\x18\xbd\xb4\x3f\x7e\x9a\x0f\xa6\x31\x8d\x0f\x00\x1e\xf7\xe1\xe0\x62\x1d\xe1\x01\x00\xe0\x3b\x9c\x73\xf1\x23\x1c\x73\x48\x74\xfd\x07\xe0\x8c\xee\xcf\x08\x93\x8b\x15\x19\xe3\x73\x8c\xae\xd4\x3b\x14\xb6\x92\x6d\xe6\x1b\xfc\xcc\x2d\x7a\xd8\x21\xb8\xcf\x3e\xa0\x1e\x63\x20\xc2\x32\x42\x4e\x08\x2e\x79\x98\xc2\x09\x81\x30\xd1\x5c\x1f\xe1\xcb\x32\x17\x6b\x86\xdf\xad\xd2\x2d\x60\x0a\x31\x62\x59\x36\xa5\x73\x86\xb9\x25\x5f\xd0\x2f\x19\x13\x7d\x00\xcd\xb9\x55\x97\xfc\x23\x6b\xd0\x25\xcf\xda\xa3\xd2\xf0\xb2\xcf\x3e\xb7\x44\x58\xea\xc8\xe8\x7f\x58\xf2\x82\xa7\x76\xd8\x5d\x1a\x1d\x18\x90\xe7\xf5\x72\x42\x51\xdf\xb0\x3a\x9d\x65\xfc\x0f\x5e\x9f\x0b\xca\x84\x27\x46\x4c\xb9\x15\x01\x3c\x73\x20\x9c\x64\xc2\x0b\x03\x6a\xf8\x2b\xea\xaf\xbc\x8e\x27\x4c\x82\x78\xe3\x97\x1e\x7e\xcd\x24\x88\x77\x46\xa4\xa0\x44\x6c\xf9\x41\xd5\x16\x1b\x6e\x12\xa3\xd1\xc6\x46\xda\xc4\xa8\x0e\xbb\x91\x46\x43\x21\xd4\x49\xda\xaa\x45\x49\xb5\x93\x05\x29\xbd\x16\x64\x38\xb6\x30\x43\xb4\x41\x69\xd9\x16\xb5\x95\xb2\x12\x7d\xc8\xd7\x37\x68\x59\x92\x79\x1a\x84\x2a\x8d\x3c\x0f\xca\x93\x44\x5e\x06\x65\x49\x22\xaf\x83\xe1\x48\x42\x6f\x83\x36\x24\x99\xf7\x41\xfb\x91\xcc\xf6\xc2\x1c\x5a\xa4\x70\x8c\x41\xcd\x81\x65\xba\xa8\x27\x2a\x75\xdd\x26\x8e\xc6\x38\x77\x08\x31\x86\xac\x1f\xfc\xf5\xc7\xc8\xdd\x1a\x29\x61\x2a\x77\xe8\xdb\x02\x1f\x4e\xc1\xab\x13\x5e\x47\xee\xfa\x27\x78\x9b\xca\xeb\xcc\x22\xca\x88\xea\x70\x3a\x8f\x83\x7b\x4c\x72\xb4\x58\xc8\xbd\x98\xcf\xab\xb1\x1a\xeb\x91\x46\xa6\x79\x8f\x66\x66\x97\xb5\x52\x3b\xd7\xde\x0b\xee\xe2\x9d\xec\xde\x43\x59\xc9\xef\x2f\xe9\xef\x71\x59\xf3\x3f\x00\x00\xff\xff\x81\xc8\xe5\xe9\x67\x08\x00\x00")

func resourcesEnYmlBytes() ([]byte, error) {
	return bindataRead(
		_resourcesEnYml,
		"resources/en.yml",
	)
}

func resourcesEnYml() (*asset, error) {
	bytes, err := resourcesEnYmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "resources/en.yml", size: 2151, mode: os.FileMode(420), modTime: time.Unix(1523957071, 0)}
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
	"resources/en.yml": resourcesEnYml,
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
	"resources": &bintree{nil, map[string]*bintree{
		"en.yml": &bintree{resourcesEnYml, map[string]*bintree{}},
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

