package mymath

import "testing"

func TestEncodePerson2JsonFile(t *testing.T) {
	filename := "F:/golandWoke/file/人员.json"
	p := Person{Name: "吴铭", Age: 24, Rmb: 50, Sex: true, Hobby: []string{"踢球", "轮滑"}}
	ok := EncodePerson2JsonFile(&p, filename)
	if ok {
		pBack, _ := DecodeJsonFile2Person(filename)
		if p.Name == pBack.Name && p.Age == pBack.Age && p.Rmb == pBack.Rmb && p.Sex == pBack.Sex {
			ret := true
			if len(p.Hobby) == len(pBack.Hobby) {
				for i := range p.Hobby {
					if p.Hobby[i] != pBack.Hobby[i] {
						ret = false
						break
					}
				}
			}
			if ret {
				t.Log("编码成功。")
			} else {
				t.Fatal("编码失败，结果与预期不相符")
			}
		} else {
			t.Fatal("编码失败，结果与预期不相符")
		}
	}

}
func TestDecodeJsonFile2Person(t *testing.T) {
	filename := "F:/golandWoke/file/人员2.json"
	p := Person{Name: "吴铭", Age: 24, Rmb: 50, Sex: true, Hobby: []string{"踢球", "轮滑"}}
	ok := EncodePerson2JsonFile(&p, filename)
	if ok {
		pBack, _ := DecodeJsonFile2Person(filename)
		if p.Name == pBack.Name && p.Age == pBack.Age && p.Rmb == pBack.Rmb && p.Sex == pBack.Sex {
			ret := true
			if len(p.Hobby) == len(pBack.Hobby) {
				for i := range p.Hobby {
					if p.Hobby[i] != pBack.Hobby[i] {
						ret = false
						break
					}
				}
			}
			if ret {
				t.Log("解码成功。")
			} else {
				t.Fatal("解码失败，结果与预期不相符")
			}
		} else {
			t.Fatal("解码失败，结果与预期不相符")
		}
	}
}
