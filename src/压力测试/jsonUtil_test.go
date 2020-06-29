package mymath

import "testing"

func BenchmarkEncodePerson2JsonFile(b *testing.B) {
	filename := "F:/golandWoke/file/人员.json"
	p := Person{Name: "吴铭", Age: 24, Rmb: 50, Sex: true, Hobby: []string{"踢球", "轮滑"}}
	b.Log("BenchmarkEncodePerson2JsonFile开始了")
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		EncodePerson2JsonFile(&p, filename)
	}
}

func BenchmarkDecodeJsonFile2Person(b *testing.B) {
	b.Log("BenchmarkDecodeJsonFile2Person开始了")
	filename := "F:/golandWoke/file/人员2.json"
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_, _ = DecodeJsonFile2Person(filename)
	}
}
