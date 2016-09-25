package setup

import (
	"io/ioutil"
	"path/filepath"

	"strings"

	"fmt"

	"errors"
	"gopkg.in/yaml.v2"
)

const (
	Equal int = iota
	NotEqual
	StartsWith
	NotStartsWith
	EndsWith
	NotEndsWith
	Contains
	NotContains
)

type filter struct {
	key   string
	value string
	fType int // filter type
}

func (c *filter) CheckString(s string) bool {
	switch c.fType {
	case Equal:
		return s == c.value
	case NotEqual:
		return s != c.value
	case StartsWith:
		return strings.HasPrefix(s, c.value)
	case NotStartsWith:
		return !strings.HasPrefix(s, c.value)
	case EndsWith:
		return strings.HasSuffix(s, c.value)
	case NotEndsWith:
		return !strings.HasSuffix(s, c.value)
	case Contains:
		return strings.Contains(s, c.value)
	case NotContains:
		return !strings.Contains(s, c.value)
	default:
		panic(fmt.Errorf("Unknown filter type: %v", c.fType))
	}
}

type fixturesSearch struct {
	fixtureName string
	data        []map[string]string
	filters     []filter
}

func (f *fixturesSearch) Filter(key string, t int, value string) *fixturesSearch {
	f.filters = append(f.filters, filter{
		key:   key,
		fType: t,
		value: value,
	})
	return f
}
func (f *fixturesSearch) checkFixture(fx map[string]string) bool {
	for _, f := range f.filters {
		val, ok := fx[f.key]
		if !ok {
			panic(fmt.Errorf("Can't filter by field %s. Fixture %v doesn't contains such key", f.key, fx))
		}
		if !f.CheckString(val) {
			return false
		}
	}
	return true
}
func (f *fixturesSearch) All() []map[string]string {
	result := []map[string]string{}
	for _, d := range f.data {
		if f.checkFixture(d) {
			result = append(result, d)
		}
	}
	return result
}
func (f *fixturesSearch) First() (map[string]string, error) {
	data := f.All()
	if len(data) == 0 {
		return nil, errors.New("No such fixture")
	}
	return data[0], nil
}
func (f *fixturesSearch) Count() int {
	data := f.All()
	return len(data)
}

func GetFixture(name string) *fixturesSearch {
	file, err := ioutil.ReadFile(filepath.Join(FixturesPath, name+".yml"))
	if err != nil {
		panic(err)
	}
	fixtures := []map[string]string{}
	err = yaml.Unmarshal(file, &fixtures)
	if err != nil {
		panic(err)
	}
	result := new(fixturesSearch)
	result.data = fixtures
	result.fixtureName = name
	return result
}
