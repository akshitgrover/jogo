package jogo

import (
	"testing"
	"github.com/akshitgrover/jogo/jogo"
)

const (
	exampleJson = `{
	"widget": {
	  "debug": "on",
	  "window": {
		"title": "Sample Konfabulator Widget",
		"name": "main_window",
		"width": 500,
		"height": 500
	  },
	  "image": { 
		"src": "Images/Sun.png",
		"hOffset": 250,
		"vOffset": 250,
		"alignment": "center"
	  },
	  "text": {
		"data": "Click Here",
		"size": 36,
		"style": "bold",
		"vOffset": 100,
		"alignment": "center",
		"onMouseUp": "sun1.opacity = (sun1.opacity / 100) * 90;"
	  }
	}
  }`
)

var benchPaths = []string{
	"widget.window.name",
	"widget.image.hOffset",
	"widget.text.onMouseUp",
}

func BenchmarkJoGOGet(t *testing.B) {

	t.ReportAllocs()
	exp, _, _ := jogo.Export(exampleJson)
	for i := 0; i < t.N; i++ {
		for _, v := range benchPaths {
			_, _ = exp.Get(v)
		}
	}
	t.N *= len(benchPaths)

}

func TestInt(t *testing.T) {
	_, r, _ := jogo.Export(`7`)
	if r.Type != "NUMBER" {
		t.FailNow()
	}
	num := r.Int()
	if num != 7 {
		t.FailNow()
	}
}

func TestIntNested(t *testing.T) {
	exp, _, err := jogo.Export(`{"helloworld":{"number":7}}`)
	if err != nil {
		t.FailNow()
	}
	r, err := exp.Get("helloworld.number")
	if err != nil {
		t.FailNow()
	}
	if r.Type != "NUMBER" {
		t.FailNow()
	}
	num := r.Int()
	if num != 7 {
		t.FailNow()
	}
}

func TestStringNested(t *testing.T) {
	exp, _, err := jogo.Export(`{"hello":"world"}`)
	r, _ := exp.Get("hello")
	if err != nil {
		t.FailNow()
	}
	if r.Type != "STRING" {
		t.Log(r.Type)
		t.FailNow()
	}
	str := r.String()
	if str != "world" {
		t.FailNow()
	}
}
