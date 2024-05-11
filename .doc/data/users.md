**users**

ユーザーテーブル。サービスの利用者情報。

**users**
| 物理名     |  主   | 型       | 必須  | 桁数  | 一意  |
| ---------- | :---: | -------- | :---: | :---: | :---: |
| id         |  〇   | int      |  〇   |       |  〇   |
| user_id    |       | varchar  |  〇   |  50   |  〇   |
| name       |       | varchar  |  〇   |  50   |       |
| created_at |       | datetime |       |       |       |
| updated_at |       | datetime |       |       |       |
| deleted_at |       | datetime |       |       |       |