# rest-api-boilerplate
GolangでRestAPIを作る際のボイラープレート。  
openapi定義からGolangのコードを生成する機能と、CleanArchitectureをベースにしたパッケージ構造を元に構成しています。

## 前提条件
- Golang 1.9

## 使い方
リポジトリをcloneします。
```
git clone https://github.com/choimake/rest-api-boilerplate.git
```

生成したコードを元に、APIの実装を行います。

## パッケージ構成
| パッケージ    | 説明                                |
|----------|-----------------------------------|
| cmd      | 実行ファイルの置き場所。                      |
| docs     | ドキュメント類の置き場所。                     |
| internal | 外部公開しないパッケージの置き場所。APIの内部処理はここに置く。 |

## 参考
### パッケージ構成
- [Standard Go Project Layout](https://github.com/golang-standards/project-layout)
