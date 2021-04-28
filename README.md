heyTagos
===

heyTacos!にインスパイアを受けて作成しました。
無料で使いたかったので自作です。

本家
> https://www.heytaco.chat/


## 概要
- Slack :: slash command -> GoogleCloudPlatform :: Cloudfunction
  - GET:  /tagos-check [date, month, year] -> 任意の期間で送りたい人を元に集計(fireStoreにて集計)
  - POST: /tagos @送りたい人 メッセージ -> @送りたい人に対して投票データを追加(fireStoreにデータ作成)

## 使用技術
- golang.1.13
- slack slash command
- GoogleCloudPlatform
  - cloud function
  - fireStore

 2021/04/25 created by `okabe_yuya`

