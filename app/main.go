package main

import (
	"context"
	"crypto/sha256"
	"database/sql"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
	"os/signal"
	"runtime"
	"sync"
	"time"

	"github.com/google/uuid"
	_ "github.com/mattn/go-sqlite3"
	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

// Config はアプリケーション設定を保持する構造体
type Config struct {
	Port    int    `json:"port"`
	LogPath string `json:"log_path"`
}

// サーバーの初期化と設定を行う
func initServer(cfg Config) *http.Server {
	mux := http.NewServeMux()

	// メインハンドラー
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// ハッシュ値を生成（使用例として）
		hash := sha256.Sum256([]byte("sample data"))
		hashStr := base64.StdEncoding.EncodeToString(hash[:])

		// UUIDを生成
		id := uuid.New().String()

		// 現在時刻
		now := time.Now().Format(time.RFC3339)

		// 数値をフォーマット
		p := message.NewPrinter(language.Japanese)
		formattedNum := p.Sprintf("%d", 1000000)

		// レスポンスデータを構築
		data := map[string]interface{}{
			"message":    "sampleだよ",
			"hash":       hashStr,
			"uuid":       id,
			"time":       now,
			"go_version": runtime.Version(),
			"cpu_count":  runtime.NumCPU(),
			"os":         runtime.GOOS,
			"formatted":  formattedNum,
		}

		// JSONレスポンスを返す
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(data)
	})

	// プレーンテキストで「sampleだよ」を返すエンドポイント
	mux.HandleFunc("/plain", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		w.Write([]byte("sampleだよ"))
	})

	// ヘルスチェック
	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "OK")
	})

	// 環境情報
	mux.HandleFunc("/info", func(w http.ResponseWriter, r *http.Request) {
		info := map[string]string{
			"hostname":    os.Getenv("HOSTNAME"),
			"environment": os.Getenv("ENV"),
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(info)
	})

	// その他のルートを追加
	setupAdditionalRoutes(mux)

	// サーバー設定
	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", cfg.Port),
		Handler:      mux,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
	}

	return server
}

// その他のルートを設定する補助関数
func setupAdditionalRoutes(mux *http.ServeMux) {
	// データベース接続テスト
	mux.HandleFunc("/db-test", func(w http.ResponseWriter, r *http.Request) {
		db, err := sql.Open("sqlite3", ":memory:")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer db.Close()

		// テーブル作成
		_, err = db.Exec("CREATE TABLE IF NOT EXISTS test (id INTEGER PRIMARY KEY, name TEXT)")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		fmt.Fprintf(w, "Database connection successful")
	})

	// テンプレートレンダリングのテスト
	mux.HandleFunc("/template", func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.New("test").Parse("<html><body><h1>{{.Title}}</h1><p>{{.Message}}</p></body></html>"))
		data := struct {
			Title   string
			Message string
		}{
			Title:   "テンプレートテスト",
			Message: "sampleだよ",
		}

		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		tmpl.Execute(w, data)
	})
}

func main() {
	// 設定（通常はファイルから読み込むが、ここでは直接設定）
	config := Config{
		Port:    8080,
		LogPath: "app.log",
	}

	// ロガー設定
	logFile, err := os.OpenFile(config.LogPath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal("ログファイルを開けませんでした:", err)
	}
	defer logFile.Close()

	// 標準出力とファイルの両方にログを出力
	multiWriter := io.MultiWriter(os.Stdout, logFile)
	logger := log.New(multiWriter, "[APP] ", log.LstdFlags|log.Lshortfile)

	// サーバー初期化
	server := initServer(config)

	// 並行処理でサーバー起動
	var wg sync.WaitGroup
	wg.Add(1)

	// シャットダウンハンドラ
	_, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	// サーバー起動
	go func() {
		defer wg.Done()

		logger.Printf("サーバーを起動しています ポート: %d\n", config.Port)
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.Fatalf("サーバー起動エラー: %v", err)
		}
	}()

	// 現在の実行ディレクトリを表示
	if wd, err := os.Getwd(); err == nil {
		logger.Printf("作業ディレクトリ: %s\n", wd)
	}

	// 環境変数を表示
	logger.Printf("GO環境: %s\n", runtime.Version())
	logger.Printf("実行OS: %s\n", runtime.GOOS)
}
