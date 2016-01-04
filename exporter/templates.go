// Code generated by go-bindata.
// sources:
// exporter/templates/cat1/cat1.xml
// DO NOT EDIT!

package exporter

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

var _exporterTemplatesCat1Cat1Xml = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xec\x58\x6d\x73\xe2\x38\x12\xfe\x9e\x5f\xa1\x71\x15\x5b\xc9\x56\xb0\x79\x49\x26\xd9\x94\xe3\x2d\x06\xd8\x0b\x55\x09\xc9\x02\xb9\xbd\xfb\xb4\x25\x6c\x01\xba\xb5\x65\x4e\x96\x43\xb8\xad\xf9\xef\xd7\x7a\x31\xc8\xe6\x2d\x37\x35\xf7\x6d\x53\xa9\x60\x5a\xad\xee\x56\xf7\xd3\x8f\xda\xf1\x6b\x9f\xd0\x2c\x67\x21\xea\x62\xd1\x9c\x24\xcb\xf8\x7c\x85\x68\xea\xfe\xc6\xa9\x20\xfc\x12\x2d\xb1\xa0\x84\x09\xf4\x63\x92\x46\x24\xce\xdc\x11\x09\x53\x1e\x5d\xa2\x05\xc1\x11\xe1\x1b\xf1\x83\xfa\x7a\x81\x08\xe7\x29\x47\xb5\xe0\xcc\xaf\xd5\x10\x4d\x96\x29\x17\xc8\x99\x53\xb1\xc8\xa7\x6e\x98\x26\xde\x92\xa7\xff\x22\xa1\x08\xd7\x4b\x4e\xb2\xcc\x0b\x23\x2c\xd2\x34\xce\x3c\x6d\xc6\x41\xb5\x23\x5b\xa7\x29\x4f\x30\xf3\xf2\x9c\x46\xbb\x8a\x82\x26\x64\x57\x3a\x5d\x0b\x52\x58\xfd\xf9\x3d\x89\xd1\x1b\xe1\x19\x4d\xd9\xbd\xd3\x74\x1b\x0e\x22\x2c\x4c\x23\xca\xe6\xf7\x4e\x2e\x66\xf5\x5b\xe7\x67\x50\xeb\xc6\x94\xd1\x10\xc7\xbd\x34\xcc\x13\x79\x72\xd8\xc6\xb2\xbb\xf7\x8c\xde\x3b\x0b\x21\x96\x77\x9e\xb7\x5a\xad\xdc\x55\xdb\x4d\xf9\xdc\x6b\x35\x1a\x4d\xef\x1f\x4f\x8f\xe3\x70\x41\x12\x5c\xa7\x2c\x13\x98\x85\xc4\x39\xd3\xdb\xc0\x30\x67\x77\x8b\xf8\xa6\x0e\xca\x77\x6f\xed\x42\x7e\xf7\x96\x86\xd5\x35\x0f\x64\x9b\xf5\x2c\x12\x15\x05\x29\x71\x82\x33\x84\xfc\x4f\xf5\x3a\xfa\x75\xd4\xeb\x20\x9d\x74\x54\xaf\x2b\x31\x27\x38\x4e\xba\x90\x47\x04\x87\x22\xf7\xce\xeb\xd8\xf1\xd4\x82\x58\x2f\xc9\x20\x42\x3c\x4d\xc5\xbd\xd3\x72\x9b\x9f\xdd\xdb\xab\x86\xdb\x74\x9b\xcd\xf6\xed\x6d\x1b\x1e\xda\x90\x89\x77\x41\x98\xce\xcc\xcb\x73\xb7\xf7\xfb\x43\xaf\x01\x3f\x57\x0d\x63\x42\xba\x7c\x1d\xa3\x91\x74\x51\xb8\x9d\x10\x00\x0b\x16\x04\x81\x6d\x13\x82\x30\xa2\x63\xde\x1a\x6e\x0b\x7e\x5b\x52\x62\x19\x57\xe7\xb1\xb6\xff\xcf\x06\xaf\x94\x41\x64\x59\xec\x3d\xd5\xa7\x38\x23\xd1\xf7\xb1\xdd\xb2\x6d\x4f\x16\x34\x43\xf0\x2b\x16\x04\xcd\xe3\x74\x8a\xe3\x78\x8d\x72\x46\xff\x9d\x13\x44\x23\x00\x0d\x9d\x51\xc8\xd0\x0c\x5a\x41\x48\x55\x15\x41\x54\x00\xca\xf8\xa7\x85\x5f\xbf\x76\x8f\x24\xa8\xdd\x21\x59\x8d\x30\x8b\xd2\xe4\xfc\x02\x5a\xa8\x9a\x9d\xcd\x7e\x59\x4f\x55\xe3\xc2\x52\xb8\x2d\xfa\xf5\x75\xf3\xb6\x55\x07\x68\xcb\xaf\xe3\x75\x06\x87\xdc\x77\xb0\xcf\x32\x57\x5b\x95\x21\x4e\x60\xef\xe3\xf3\x60\xd8\x75\x50\x44\x33\xc8\xcb\x5a\xcb\x7e\xcd\x71\x4c\xc5\x1a\x3d\x11\x9c\xe5\x9c\x00\x02\x64\x63\x15\xc0\xa2\x22\x26\x81\x0a\x6e\xc0\x42\x79\xf0\xb0\xd0\xf0\x3d\xbd\xb8\x2f\x63\x9b\x83\x84\x00\x59\x01\x98\x43\xb2\x79\x8b\xc3\x90\xd9\x0c\x08\x82\xbe\x91\x89\x14\xbe\xe1\x38\x27\x3a\x45\x52\xc9\x1d\xa6\xab\xf3\x0b\xf7\x75\xd2\x95\x7f\x19\x7d\xb7\x33\x15\xa6\x6c\xa6\xb3\xaf\x62\xb6\x5a\x61\x78\x2a\x1f\xd7\x6e\xeb\xda\x58\x89\x31\x9b\xe7\x78\x4e\xac\xed\x84\x59\xb5\xe0\xea\x7c\x00\xab\x82\x19\x4d\xdc\x35\x88\x95\xa3\x69\x0e\xe1\xc3\x87\xe4\x1d\xf7\x8b\xfe\x52\x0b\xe4\xaa\x26\xce\x09\xe6\x73\x22\x14\xcd\xfe\xa0\x75\x37\x14\x7b\x61\x14\x0b\xf2\x1d\x0b\x0e\xd4\x74\xbe\xba\x34\x46\x5d\x23\xb8\x90\x8a\x67\xda\x25\x9d\x15\x44\xfc\xe9\x1e\x31\x1a\xa3\x3f\x41\x8e\x0e\x45\xa2\xd6\x24\x28\x7f\xbf\x44\x38\x17\x0b\x78\xba\xbb\x47\x1c\x0e\x4c\x8c\x19\xb7\xa3\xc4\x99\xb1\x83\x90\xfe\x5e\x8e\xf7\x07\xbd\xf7\x42\xa9\x7c\x55\x7f\x4f\xc7\xac\xd4\x6a\x81\xfa\x90\x69\x1c\x3f\x74\x1e\x1f\xd1\x02\xbf\x11\xd4\x74\xdd\x1f\x4d\x3c\x2e\x7a\xea\xfc\x13\x4d\x01\x23\xe4\x8d\x02\x98\x20\xc2\x25\x30\x76\xca\x5c\x13\xcf\x04\xf0\x63\x42\x4f\x67\x0a\x4d\x5d\xbb\x35\x28\xd3\x0d\x47\xde\x31\x84\x4c\x24\xe2\x70\x61\x0b\x0b\xf9\x8c\x05\x46\x59\x3e\x4d\x68\x26\xd9\x0e\xee\x03\xe8\x37\xee\x71\x32\xa7\x99\xe0\x6b\x57\x55\x53\xf9\x32\xb0\x05\x0c\xe4\x99\x80\x5b\x02\xb3\xfd\x1e\x35\xa6\x33\xe8\x16\x84\xf5\x73\x4c\xe6\x38\x56\x51\x4a\x2c\x86\x70\xc3\xf1\x22\x30\x65\xb9\x08\x0e\x5a\x5d\xa2\x09\x2e\x42\xd0\xd3\x5b\x35\xb8\x20\x63\x70\xf2\x39\x66\xf4\x3f\xaa\x3f\x5c\x83\x31\x13\x94\x4e\xdd\x46\xb4\x89\x2f\x30\x39\xf2\x31\x9c\x6d\xce\x48\xd4\xdd\x59\xa9\x99\x07\x54\xd4\x67\x04\xce\xc5\xf9\xc5\x46\xfc\x6c\xb9\xad\x54\xdd\x20\x64\x63\xd4\xb5\x75\x2f\x91\xb3\x39\x8a\xe5\xd8\x56\x71\xb6\x5e\x3e\x08\x17\x0b\x32\x10\xbb\x77\xe0\x58\xbe\x67\x65\xa0\x5c\xba\x7d\x95\x38\x50\x44\x40\x05\x78\x86\xaa\x68\xbc\xa1\xd5\x22\x55\x62\x21\xb5\x4d\x08\xaa\xc8\x76\x61\x8a\x8a\x6f\x21\x52\x01\xa0\xab\x89\x8f\xc4\x44\xf9\x49\x40\x4f\xa2\xdb\x24\xea\x78\x5d\x55\xf0\x1d\x3b\xf6\x4d\x2a\x76\xb5\x15\x1b\x97\xf9\xd2\xea\xe8\x8d\x05\x57\x91\x6a\xc1\x98\x87\x6d\xc9\x44\x63\x01\x84\x6f\xb1\xe0\xd8\xda\x54\x94\xa2\x0f\x86\xc5\x3a\xd8\x54\xcb\xb2\x36\xe8\x59\x06\x15\x59\x1d\x80\xdc\x86\x92\xe0\x42\xdc\x47\x47\xdb\xe0\x07\xd1\x96\x98\xf4\xcf\x20\xaa\x80\x94\x46\xb6\xe1\xaf\xbb\x4e\x70\x14\xc9\xc9\xf3\x84\xa7\x8e\xd6\x22\x55\x7f\x46\x5e\xe5\x43\x2d\x3d\xe1\x59\x00\x0a\x60\x8e\x3d\xe1\x79\xa2\xb5\xaa\x8e\x8d\xb8\xe2\xd8\x98\x3c\xe4\xf8\x5b\xda\xcc\x2a\xee\x8b\xea\x84\xc0\x32\xe8\x33\xe8\x80\xa0\x14\x18\xf2\xe7\x70\x57\xb3\xe0\x20\xe2\xb4\x15\xf7\x17\xca\x01\xfb\xb5\xc0\xf7\xb4\x7e\xc5\xc8\x0c\x27\x34\x5e\x9f\xb4\xf2\x88\x8d\x11\xa3\x6f\xc7\xe6\x95\x82\xdb\xd2\x45\x71\x8c\x0f\x40\xf1\x34\xfb\x95\xa3\x3a\xcc\x80\xff\x17\xe2\xb3\xbb\xcd\xf7\xf6\xf1\x83\x9e\x0a\xbe\x02\xe1\x64\x04\xfd\x59\xd8\xf0\xf5\x8d\x69\xbe\x54\x79\xe2\xd4\x5c\x65\x21\xa2\x63\xd9\x31\xcd\x0e\x4d\x6b\xbd\x3b\x74\xf5\x7b\x9d\x73\x78\xa2\xfe\xc9\x85\x41\xcb\xa6\x1a\x69\x64\xf8\x32\xb0\x45\x65\x9b\xbf\xe0\x3f\x08\x28\x1c\xb6\x79\xe5\x7e\xb6\xa9\x09\x9a\xd1\x02\x33\xdc\xea\x84\x08\xd3\xb7\x8f\x94\x91\xa0\xd5\x68\xa1\x2f\x39\x87\xd7\xbc\xb9\x00\x0a\x1f\x45\xae\xef\xed\x6a\x6d\x2d\x84\x32\xe7\x5f\x48\x04\x6d\x1c\xc1\x5d\x53\xe6\x3b\x78\xdd\x13\x24\x78\xea\x48\x13\xf2\x69\xbb\xb2\x4c\x41\x12\x4b\xfe\x0c\x1a\xcd\x9b\x76\xc3\xf7\x2c\x89\x65\x3d\xcd\x19\x0c\x1e\xc1\xeb\x18\x6c\x9b\xe7\x6d\xe1\xad\xb3\xf8\x05\x7f\xe4\x19\x14\xee\xb7\x17\xa7\xa8\x21\xc8\xef\xce\x6f\x6e\x9b\x17\xad\x9b\x66\xbd\x0d\x6f\x6d\x7b\x78\x5a\x17\x0e\x4e\xdc\x53\xb3\x90\xe5\x1f\xde\xa8\xf3\x19\x0e\x25\xd7\xf3\x27\xf9\x12\x2e\xe7\xfe\xc0\xd4\xd1\xf7\xf6\x2e\x23\x3b\x01\xe9\x4c\xac\x30\x27\xe5\x6d\xb6\x14\xed\x02\xb9\x12\x0e\x2a\x20\xbd\x8b\x33\x90\x6d\x9f\x2b\x93\xce\x91\x39\xe7\xd4\x2c\x62\x1d\x80\x1e\x79\xfd\x53\x60\xb5\x74\x99\x75\x48\xa0\x64\xe0\xa2\x1e\x09\xff\xa8\x30\xcf\x37\x55\x6a\x07\xb8\xdf\x03\xba\x27\xc0\x7b\x0c\xbe\x1f\x06\xf0\x71\x08\xef\x80\xd8\xfb\x58\x65\x0e\x0c\x7b\xa5\x51\xef\xc8\x88\xf4\x2d\x1c\x77\x74\xe4\x39\x30\xf0\x6c\xc1\x33\x0d\x1b\x4d\x7c\x1d\x41\x59\x71\xfb\xaa\x7e\xd5\xba\xfd\x5c\xbf\x6d\x85\x61\xfd\xaa\x4d\xa6\x8d\xab\xf0\xa7\x9b\x16\xbe\xf9\x8b\xa4\x8e\x93\x54\x75\xde\xd8\x9d\x36\xcc\xac\xf1\x40\x18\x5f\xef\x1b\x24\x8a\x31\x62\x4c\x60\x65\x77\x48\x28\x35\xea\xee\x80\x60\x94\x0e\xdc\xe2\xdf\x81\x32\xca\x44\x51\xea\x85\x63\x2d\xb0\xc5\xdc\xde\x4b\x5f\x5d\xf8\xe6\x3f\x02\xf2\x32\x7d\x19\x3d\xff\x7d\xd0\xeb\x8f\xc6\x68\xd4\x1f\xc2\x27\x7a\xe8\x8f\xfa\xc5\x7b\x2d\x94\x32\x59\xa6\x0c\x76\x17\xb0\x17\x3c\x57\xf4\x1e\x7d\x49\xa3\xb5\x7d\x2b\x3f\xf5\x3b\xe3\xd7\x51\xbf\x64\xa6\x7c\x6f\x8f\xfa\x2f\xcf\xa3\xc9\x60\xf8\x37\xf4\xd2\x19\x75\x9e\x8e\x68\xbe\x74\x26\x83\xfe\x70\x82\x7a\x9d\x49\xc7\x68\x8d\x4b\x6a\x0a\xdd\x95\x40\x24\xd6\x36\xc1\x9e\xf9\x5e\xf5\x3f\xb3\xc1\xd9\x7f\x03\x00\x00\xff\xff\x0f\xad\x75\x99\xbc\x16\x00\x00")

func exporterTemplatesCat1Cat1XmlBytes() ([]byte, error) {
	return bindataRead(
		_exporterTemplatesCat1Cat1Xml,
		"exporter/templates/cat1/cat1.xml",
	)
}

func exporterTemplatesCat1Cat1Xml() (*asset, error) {
	bytes, err := exporterTemplatesCat1Cat1XmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "exporter/templates/cat1/cat1.xml", size: 5820, mode: os.FileMode(420), modTime: time.Unix(1451935445, 0)}
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
	"exporter/templates/cat1/cat1.xml": exporterTemplatesCat1Cat1Xml,
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
	"exporter": &bintree{nil, map[string]*bintree{
		"templates": &bintree{nil, map[string]*bintree{
			"cat1": &bintree{nil, map[string]*bintree{
				"cat1.xml": &bintree{exporterTemplatesCat1Cat1Xml, map[string]*bintree{}},
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

