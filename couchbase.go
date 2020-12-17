package main

import (
	"fmt"
	"time"

	gocb "github.com/couchbase/gocb/v2"
)


func main() {
	// クラスターへの接続
	cluster, err := gocb.Connect(
		"localhost",
		gocb.ClusterOptions{
			Username: "Administrator",
			Password: "password",
		})
	if err != nil {
		panic(err)
	}
	// バケットへの参照の取得
	bucket := cluster.Bucket("test")

	// バケットに接続され確実に利用可能になるまで待つ
	err = bucket.WaitUntilReady(5*time.Second, nil)
	if err != nil {
		panic(err)
	}

	// コレクションへの参照の取得（デフォルトコレクションを利用）
	collection := bucket.DefaultCollection()

	// ドキュメントの定義（ユーザ情報）
	docUser := map[string]string{"id": 1, "name": "田中", "type": "user"}
	//　ドキュメントIDの定義
	docId := "user_1"

	// ドキュメントをUpsertする
	upsertResult, err := collection.Upsert(docId, docUser, &gocb.UpsertOptions{})
	if err != nil {
		panic(err)
	}
	// UpsertResultからCAS値を出力
	fmt.Println(upsertResult.Cas())

	// ドキュメントをGetする
	getResult, err := collection.Get(docId, &gocb.GetOptions{})
	if err != nil {
		panic(err)
	}

	// GetResultから、Contentを取得
	var myContent interface{}
	if err := getResult.Content(&myContent); err != nil {
		panic(err)
	}
	// Contentを出力
	fmt.Println(myContent)

}
