# tilegen

大きな地図の画像から、タイル画像を生成するgolang製ツールです。

## 環境変数

`PROJECT_NAME`: 出力先のディレクトリ名
`IMAGE_NAME`: 入力画像のファイル名 (".png" は不要) (指定しない場合、PROJECT_NAME と同じになる)
`ZOOM_RANGE_MIN`: ズームレベルの最小値 (デフォルト: 0)
`ZOOM_RANGE_MAX`: ズームレベルの最大値 (デフォルト: 65535)
`TILE_SIZE`: タイル画像のサイズ (デフォルト: 256)

## 使用例

```
$ PROJECT_NAME=example IMAGE_NAME=example ZOOM_RANGE_MIN=0 ZOOM_RANGE_MAX=3 TILE_SIZE=256 go run main.go
```

```
$ PROJECT_NAME=example go run main.go
```

## ライセンス

MIT