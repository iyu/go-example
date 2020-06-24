ディレクトリ構想について
====

[animal](./animal)と[car](./car)、どちらもコードだけ見たら同じ動きだけどパッケージの仕分け方が違う

どちらがよいのだろうか

```
pkg/directory/animal/
├── category
│   ├── animal.go
│   ├── cat.go
│   └── dog.go
└── service.go

pkg/directory/car/
├── category
│   ├── car.go
│   ├── sedan
│   │   └── sedan.go
│   └── wagon
│       └── wagon.go
└── service.go
```

- animal方式
  - importがディレクトリ階層の上から下への一方向でわかりやすい
  - categoryが増えてきたときにcategoryパッケージが肥大化する
    - registerで必要なものだけ選んでDIしたいのに、余計なコードまでロードされる
    - 同一パッケージだとtestが並行に走らないらしい
- car方式
  - パッケージが細かく分離できる
  - importがディレクトリ階層の上からも下からもされているのでわかりにくい
