package main

import (
	"bufio"
	"flag"
	"log"
	"os"
	"path/filepath"
)

type k27804 struct {
	head string
	kd   [5]kdata
}

type kdata struct {
	kCode  string
	kValue string
	kFlag  string
	kCmt   string
}

func failOnError(err error) {
	if err != nil {
		log.Fatal("Error:", err)
	}
}

func main() {
	flag.Parse()

	//ログファイル準備
	logfile, err := os.OpenFile("./log.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, os.ModePerm)
	failOnError(err)
	defer logfile.Close()

	log.SetOutput(logfile)
	log.Print("Start\r\n")

	// ファイルの読み込み準備
	infile, err := os.Open(flag.Arg(0))
	failOnError(err)
	defer infile.Close()

	// ファイルの書き込み準備
	writeFileDir, _ := filepath.Split(flag.Arg(0))
	writeFileDir = writeFileDir + "K27804iabc.DAT"
	outfile, err := os.Create(writeFileDir)
	failOnError(err)
	defer outfile.Close()

	var ik27804 k27804
	//ik27804の初期化
	ik27804.head = ""
	for i := 0; i < 5; i++ {
		ik27804.kd[i].kCode = ""
		ik27804.kd[i].kValue = ""
		ik27804.kd[i].kFlag = ""
		ik27804.kd[i].kCmt = ""
	}

	// ファイルの読み込み
	scanner := bufio.NewScanner(infile)
	for scanner.Scan() {
		line := scanner.Text()
		ik27804.head = line[:78]
		pos := 78
		for i := 0; i < 5; i++ {
			ik27804.kd[i].kCode = line[pos+(i*32) : pos+(i*32)+5]
			ik27804.kd[i].kFlag = line[pos+(i*32)+5 : pos+(i*32)+5+1]
			ik27804.kd[i].kValue = line[pos+(i*32)+5+1 : pos+(i*32)+5+1+19]
			ik27804.kd[i].kCmt = line[pos+(i*32)+5+1+19 : pos+(i*32)+5+1+19+7]

			switch ik27804.kd[i].kCode {
			case "22667":
				ik27804.kd[i].kCode = "22664"
			case "22668":
				ik27804.kd[i].kCode = "22665"
			case "22670":
				ik27804.kd[i].kCode = "03734"
			case "22671":
				ik27804.kd[i].kCode = "03735"
			case "22672":
				ik27804.kd[i].kCode = "03736"
			}
		}

		// ファイルの書き込み
		wline := ik27804.head
		for i := 0; i < 5; i++ {
			wline = wline + ik27804.kd[i].kCode
			wline = wline + ik27804.kd[i].kValue
			wline = wline + ik27804.kd[i].kFlag
			wline = wline + ik27804.kd[i].kCmt
		}
		_, err = outfile.WriteString(wline + "\r\n")
		failOnError(err)

		//ik27804の初期化
		ik27804.head = ""
		for i := 0; i < 5; i++ {
			ik27804.kd[i].kCode = ""
			ik27804.kd[i].kValue = ""
			ik27804.kd[i].kFlag = ""
			ik27804.kd[i].kCmt = ""
		}

	}

	log.Print("Finesh !\r\n")

}
