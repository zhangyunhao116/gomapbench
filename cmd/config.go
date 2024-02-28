package main

import (
	"fmt"
)

type MapGen interface {
	New(k, v, cap string) string
	Store(self, k, v string) string
	Load(self, k string) string
	Delete(self, k string) string
	RangeAll(self, k, v, content string) string
	DeleteAll(self string) string
}

type runtimeMap struct {
}

func (r *runtimeMap) New(k, v, cap string) string {
	if cap == "0" {
		return fmt.Sprintf("make(map[%s]%s)", k, v)
	}
	return fmt.Sprintf("make(map[%s]%s, %s)", k, v, cap)
}

func (r *runtimeMap) Store(self, k, v string) string {
	return fmt.Sprintf("%s[%s] = %s", self, k, v)
}

func (r *runtimeMap) Load(self, k string) string {
	return fmt.Sprintf("%s[%s]", self, k)
}

func (r *runtimeMap) Delete(self, k string) string {
	return fmt.Sprintf("delete(%s,%s)", self, k)
}

func (r *runtimeMap) DeleteAll(self string) string {
	return fmt.Sprintf("for k := range %s { delete(%s,k) }", self, self)
}

func (r *runtimeMap) RangeAll(self, k, v, content string) string {
	return fmt.Sprintf("for %s,%s := range %s { %s}", k, v, self, content)
}

type swisstable0 struct {
}

func (r *swisstable0) New(k, v, cap string) string {
	return fmt.Sprintf("xmap.New[%s,%s](%s)", k, v, cap)
}

func (r *swisstable0) Store(self, k, v string) string {
	return fmt.Sprintf("%s.Store(%s, %s)", self, k, v)
}

func (r *swisstable0) Load(self, k string) string {
	return fmt.Sprintf("%s.Load(%s)", self, k)
}

func (r *swisstable0) Delete(self, k string) string {
	return fmt.Sprintf("%s.Delete(%s)", self, k)
}

func (r *swisstable0) DeleteAll(self string) string {
	return fmt.Sprintf("%s.Clear()", self)
}

func (r *swisstable0) RangeAll(self, k, v, content string) string {
	return fmt.Sprintf("%s.Range(func(%s, %s int) bool { %s \n return true })", self, k, v, content)
}
