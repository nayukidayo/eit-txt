package main

import (
	"bufio"
	"encoding/binary"
	"encoding/hex"
	"encoding/json"
	"flag"
	"log"
	"math/rand"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	var mode, txt string
	flag.StringVar(&mode, "mode", "", "模式 uell400 uell uref cirs cirs812")
	flag.StringVar(&txt, "txt", "", "文件路径")
	flag.Parse()

	f, err := os.Open(txt)
	if err != nil {
		log.Fatalln("文件路径错误")
	}
	defer f.Close()

	var arr any
	switch mode {
	case "uell400":
		arr = uell400(f, 208)
	case "uell":
		arr = uell(f, 208)
	case "uref":
		arr = uref(f, 208)
	case "cirs":
		arr = cirs(f, 208, 1024)
	case "cirs812":
		arr = cirs(f, 208, 812)
	default:
		log.Fatalln("未指定模式")
	}

	data, _ := json.Marshal(arr)

	reg := regexp.MustCompile(`\.txt$`)
	txt = reg.ReplaceAllLiteralString(txt, ".json")

	err = os.WriteFile(txt, data, 0666)
	if err != nil {
		log.Fatalln("写入JSON文件失败")
	}

	log.Println("完成")
}

func uell400(f *os.File, row int) [][]uint16 {
	defer func() {
		if r := recover(); r != nil {
			log.Fatalln("文件内容与模式不匹配")
		}
	}()
	var m [][]uint16
	g := make([]uint16, 0, row)
	p := regexp.MustCompile(`^line\d{2}|\s`)
	s := bufio.NewScanner(f)
	for s.Scan() {
		t := s.Text()
		if !strings.HasPrefix(t, "line") {
			continue
		}
		t = p.ReplaceAllLiteralString(t, "")
		h, err := hex.DecodeString(t)
		if err != nil {
			panic(err)
		}
		for i := 0; i < len(h); i += 2 {
			g = append(g, binary.LittleEndian.Uint16(h[i:i+2]))
			if len(g) == row {
				m = append(m, g)
				g = make([]uint16, 0, row)
			}
		}
	}
	if err := s.Err(); err != nil {
		panic(err)
	}
	return m
}

func uell(f *os.File, row int) []uint16 {
	m := uell400(f, row)
	n := rand.Intn(len(m))
	return m[n]
}

func uref(f *os.File, row int) []uint16 {
	defer func() {
		if r := recover(); r != nil {
			log.Fatalln("文件内容与模式不匹配")
		}
	}()
	g := make([]uint16, 0, row)
	s := bufio.NewScanner(f)
	for s.Scan() {
		n, err := strconv.ParseUint(s.Text(), 10, 16)
		if err != nil {
			panic(err)
		}
		g = append(g, uint16(n))
	}
	if err := s.Err(); err != nil {
		panic(err)
	}
	return g
}

func cirs(f *os.File, row, col int) [][]float64 {
	defer func() {
		if r := recover(); r != nil {
			log.Fatalln("文件内容与模式不匹配")
		}
	}()
	m := make([]float64, 0, row*col)
	s := bufio.NewScanner(f)
	for s.Scan() {
		n, err := strconv.ParseFloat(s.Text(), 64)
		if err != nil {
			panic(err)
		}
		m = append(m, n)
	}
	if err := s.Err(); err != nil {
		panic(err)
	}
	rows := make([][]float64, row)
	for i := range row {
		cols := make([]float64, col)
		for j := range col {
			cols[j] = m[j*row+i]
		}
		rows[i] = cols
	}
	return rows
}
