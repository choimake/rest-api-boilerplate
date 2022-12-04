# internal
外部に公開しないパッケージの置き場所。
APIの内部的な処理をここにおいている。

## パッケージ構成
| パッケージ   | 説明                |
|---------|-------------------|
| driver  | DBや外部API、FWの処理をおく |
| adapter | driverとusecaseを繋ぐ |
| usecase | ユースケース            |
| entity  | ビジネスロジック          |

## 参考
### パッケージ構成
- [CleanArchitectureでGolangらしいPackage構成を考える](https://qiita.com/inosy22/items/ce4a6ea7545c5cefd24b)
