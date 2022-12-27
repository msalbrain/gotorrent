package bencode

import (
	// "bytes"
	// "encoding/json"
	"errors"
	// "strings"
	"fmt"
	"io/ioutil"

	// "strings"
	"strconv"
	"unicode"
	// "io"
	// "os"
	// "iotuil"
)

var (
	N int
)

const dictval = 'd'
const listval = 'l'
const intval = 'i'
const end = 'e'
const delimiter = ':'

func GetFile(fn string) ([]byte, error) {

	file, err := ioutil.ReadFile(fn)

	if err != nil {
		return []byte(""), err

	}

	// f, err := os.Open(fn)
	// buf := make([]byte, 1000)
	// if err != nil {
	// 	return buf, err
	// }
	// defer f.Close()
	// for {
	// 	N, err = f.Read(buf)
	// 	if err == io.EOF {
	// 		break
	// 	}
	// 	if err != nil {
	// 		fmt.Println(err)
	// 		continue
	// 	}
	// }

	return file, nil
}

// d =

type Bencode struct {
	bArray []byte
	Decode map[string]interface{}
	index  int
}

func (b *Bencode) peek() byte {
	if len(b.bArray) > 0{
		return b.bArray[0]
	}else{
		return byte(0)
	}	
}

func (b *Bencode) consume() byte {
	if len(b.bArray) == 1{
		by := b.bArray[0]
		b.bArray = []byte("")
		return by
	}else if len(b.bArray) == 0{
		return byte(0)
	}else{
		ret := b.bArray[0]
		b.bArray = b.bArray[1:]
		return ret
	}
}

// half man karate
func (b *Bencode) Decoder() (interface{}, error) {

	// 
	check := b.peek()
	for {
		if unicode.IsDigit(rune(check)){
			return b.decodeString()
		}else if rune(check) == intval {
			return b.decodeInt()
		}else if rune(check) == listval {
			return b.decodeList()
		}else if rune(check) == dictval {
			return b.decodeDict()
		}else if rune(check) == end {
			return b.Decode ,nil
		}else {
			// fmt.Println(string(check))
			return nil, errors.New(string(check))
		}
	}
}

func (b *Bencode) decodeString() (string, error) {

	check := b.peek()

	str_num := ""
	int_num := 0
	for {
		if unicode.IsDigit(rune(check)) {
			str_num += string(b.consume())
			check = b.peek()
		} else {
			num, err := strconv.Atoi(str_num)
			if err != nil {
				return "", err
			}
			int_num = num
			break
		}

	}

	if rune(b.peek()) == ':' {
		b.consume()
	} else {
		return "", errors.New("its comming from decode string, I know i wasn't that good")
	}


	check = b.peek()
	ret_string := ""

	for {
		if len(ret_string) < int_num {
			ret_string += string(b.consume())
		}else{
			return ret_string, nil
		} 

	}

}

func (b *Bencode) decodeInt()  (int, error){
	var int_byte string
	
	check := b.peek()
	if rune(check) == 'i'{
		b.consume()
	}else{
		return 0, errors.New("this is comming from decode integer")
	}

	check = b.peek()
	for {
		if rune(check) == end {
			if string(int_byte) == ""{
				return 0, nil
			}
			num, err := strconv.Atoi(string(int_byte))
			if err != nil {
				return 0, err
			}
			
			b.consume()
			return num, nil
		}else if unicode.IsDigit(rune(check)) {
			int_byte += string(b.consume())
			check = b.peek()
		} else {
			return 0, errors.New("comming from int")
		}

		check = b.peek()
	}
}

func (b *Bencode) decodeList() ([]interface{}, error){
	// TODO: if decoding a array of int and index 0 is empty e.g.
	// ie it reurns [0] even if there is more after 
	
	b.consume()

	var ret_list []interface{}
	
	check := b.peek()
	for {
		if rune(check) == end {
			b.consume()
			return ret_list, nil
		}else {
			d, err := b.Decoder()
			if err != nil{
				return []interface{}{}, err
			}
			ret_list = append(ret_list, d)
		}
		check = b.peek()
	}
}

func (b *Bencode) decodeDict() (map[string]interface{}, error) {
	b.consume()
	
	ret_dict := make(map[string]interface{})

	check := b.peek()
	for {
		if rune(check) == end {
			b.consume()
			return ret_dict, nil
		}else {
			key, err := b.Decoder()
			if err != nil {
			return ret_dict, err
			}
			
			value, err := b.Decoder()
			if err != nil {
				return ret_dict, err
			}

			// if fmt.Sprintf("%v", key) == "pieces"{
			// 	value = []byte(fmt.Sprintf("%v", value))
			// }
			fmt.Println(fmt.Sprintf("%v", key))
			fmt.Print("\n\n\n\n\n")	
			
			ret_dict[fmt.Sprintf("%v", key)] = value
		}
	}
}
