package main

import (
	"log"
	"os"
	"io"
)

func LoggingSettings(logFile string){
	//os.OpenFileでファイルを開く。このとき、読み書き、新規作成、追記等の設定ができる
	logfile, _ := os.OpenFile(logFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	//os.MultiWriterでログの出力先を複数選択できる。この場合、コンソールとlogfileに出力している
	multiLogFile := io.MultiWriter(os.Stdout, logfile)
	//ログにフラグを仕込む。この場合、日付、時間、ファイルのパスを出力するようにしている
	log.SetFlags(log.Ldate|log.Ltime|log.Lshortfile)
	//ログを出力するio.Writerを設定する。この場合、MultiLogFileを設定している
	log.SetOutput(multiLogFile)
}

func main(){
	LoggingSettings("test.log")
	log.Println("log")
	log.Printf("%T %v", "test", "test")
	/*
	log.Fatalln("loglog")
	//fatalを使うと、exitするのでコードの処理が終了する
	//以下の処理も実行されない
	log.Fatalf("%T %v", "test", "test")
	*/
	_ , err := os.Open("nothing")
	if err !=nil {
		log.Fatalln("Exit", err)
	}
}