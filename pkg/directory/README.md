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
