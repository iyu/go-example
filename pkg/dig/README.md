DIG
====

`go.uber.org/dig`の検証

- 1つのconstructorを使って複数のinstanceを作り、別々にDIするには
  - fuga用のclientとhoge用のclient、オプションが違うだけで同じことをしたい
    - New時に必要とするinterface名をaliasで変えるだけでイケるか検証
      - 結果 → ~いけた~ いけないい
        - alias使ってるつもりがNamedTypeだった
        - NamedTypeの場合はいけるがaliasじゃいけない
        
