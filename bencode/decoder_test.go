package bencode

import (
	// "errors"
	"fmt"
	"testing"
)

var G []byte

func TestFIler(t *testing.T){
	g, err := GetFile(string("../debian.torrent"))
	G = g
	if err != nil{
		t.Errorf("%v", err)
	}
	
	fmt.Println(string(g[2]))
}


func TestStringDecode(t *testing.T){
	d := Bencode{bArray: []byte("3:ape"), Decode: map[string]interface{}{}, index: 0}

	a, err := d.decodeString()

	if err != nil{
		t.Errorf("%v", err)
	}
	
	fmt.Println(a)


}




func TestIntegerDecode(t *testing.T){
	d := Bencode{bArray: []byte("ie"), Decode: map[string]interface{}{}, index: 0}

	a, err := d.decodeInt()

	if err != nil{
		t.Errorf("%v", err)
	}
	
	fmt.Println(a)


}



func TestListDecode(t *testing.T){
	d := Bencode{bArray: []byte("liei23ee"), Decode: map[string]interface{}{}, index: 0}

	a, err := d.decodeList()

	if err != nil{
		t.Errorf("%v", err)
	}
	
	fmt.Println(a)

}


func TestDictDecode(t *testing.T){
	d := Bencode{bArray: []byte("d3:cowi34e4:spam4:eggse"), Decode: map[string]interface{}{}, index: 0}

	a, err := d.decodeList()

	if err != nil{
		t.Errorf("%v", err)
	}
	
	fmt.Println(a)

}



func TestOverallDecode(t *testing.T){
	g, err := GetFile(string("../debian.torrent"))
	if err != nil{
		t.Errorf("%v", err)
	}

	d := Bencode{bArray: g, Decode: map[string]interface{}{}, index: 0}

	a, err := d.Decoder()

	if err != nil{
		t.Errorf("%v", err)
	}
	
	fmt.Println(a)
}





