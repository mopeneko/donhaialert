## 必要な機能

- ログイン
- 登録 / 解除

## 内部処理

- ログイン
  - Hostに対するCredentialのチェック
    - 無い場合
      - 取得する
  - HostからCredentialを用いてログイン用URLを取得する
  - クライアントへログイン用URLを返す
  - コールバック先で得たアクセストークンを保存
- 登録 / 解除
  - アクセストークンの有効チェック
  - DBで登録 / 削除

## 使用技術
- Vue.js
- TypeSctipt
- Go
- Echo
- GORM
- MySQL
- h2o
- Docker
- Docker Compose

## テーブル
- credentials
  - gorm.Model
  - host
  - client_id
  - client_secret
- users
  - gorm.Model
  - toots_count
  - follows_count
  - followers_count
- access_tokens
  - gorm.Model
  - credential_id (relation: ->credentials)
  - access_token
  - user_id (relation: ->users)
