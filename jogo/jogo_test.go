package jogo

import (
	"testing"
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

func BenchmarkGet(t *testing.B) {

	t.ReportAllocs()
	exp, _, _ := Export(exampleJson)
	for i := 0; i < t.N; i++ {
		for _, v := range benchPaths {
			_, _ = exp.Get(v)
		}
	}
	t.N *= len(benchPaths)

}
