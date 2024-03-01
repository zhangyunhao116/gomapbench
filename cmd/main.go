package main

import (
	"bytes"
	"go/format"
	"log"
	"os"
	"text/template"

	_ "embed"
)

//go:embed bench.tpl
var templateCode string

// Inspired by sort/gen_sort_variants.go
type Variant struct {
	// Package is the package name.
	Package string

	// Name is the variant name: should be unique among variants.
	Name string

	// Path is the file path into which the generator will emit the code for this
	// variant.
	Path string

	// Imports is the imports needed for this package.
	Imports string

	genmap MapGen

	// Funcs is a map of functions used from within the template. The following
	// functions are expected to exist:
	Funcs template.FuncMap
}

func main() {
	genenrateVersion, ok := os.LookupEnv("BENCH_TYPE")
	if !ok || genenrateVersion == "" {
		genenrateVersion = "runtime"
	}

	impls := []*Variant{
		{
			Name:    "runtime",
			Package: "gomapbench",
			Path:    "bench_test.go",
			genmap:  &runtimeMap{},
			Imports: "\"strconv\"\n\"testing\"\n",
		},
		{
			Name:    "swiss0",
			Package: "gomapbench",
			Path:    "bench_test.go",
			genmap:  &swisstable0{},
			Imports: "\"strconv\"\n\"testing\"\n\"github.com/zhangyunhao116/xmap\"\n",
		},
		{
			Name:    "swiss1",
			Package: "gomapbench",
			Path:    "bench_test.go",
			genmap:  &swisstable1{},
			Imports: "\"strconv\"\n\"testing\"\n\"github.com/cockroachdb/swiss\"\n",
		},
	}
	for _, v := range impls {
		if v.Name == genenrateVersion {
			generate(v)
			return
		}
	}
	panic("invalid name:" + genenrateVersion)
}

// generate generates the code for variant `v` into a file named by `v.Path`.
func generate(v *Variant) {
	v.Funcs = template.FuncMap{
		"New":       v.genmap.New,
		"Store":     v.genmap.Store,
		"Load":      v.genmap.Load,
		"Delete":    v.genmap.Delete,
		"RangeAll":  v.genmap.RangeAll,
		"DeleteAll": v.genmap.DeleteAll,
	}

	// Parse templateCode anew for each variant because Parse requires Funcs to be
	// registered, and it helps type-check the funcs.
	tmpl, err := template.New("gen").Funcs(v.Funcs).Parse(templateCode)
	if err != nil {
		log.Fatal("template Parse:", err)
	}

	var out bytes.Buffer
	err = tmpl.Execute(&out, v)
	if err != nil {
		log.Fatal("template Execute:", err)
	}

	os.WriteFile(v.Path, out.Bytes(), 0644)

	formatted, err := format.Source(out.Bytes())
	if err != nil {
		println(string(out.Bytes()))
		log.Fatal("format:", err)
	}

	if err := os.WriteFile(v.Path, formatted, 0644); err != nil {
		log.Fatal("WriteFile:", err)
	}
}
