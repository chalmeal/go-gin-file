**files**

ファイルテーブル。アップロードされたファイルの情報。

**files**
| 物理名     |  主   | 型       | 必須  | 桁数  | 一意  |
| ---------- | :---: | -------- | :---: | :---: | :---: |
| id         |  〇   | int      |  〇   |       |  〇   |
| file_id    |       | varchar  |  〇   |  100  |  〇   |
| user_id    |       | varchar  |  〇   |  50   |       |
| path       |       | varchar  |  〇   |  500  |       |
| created_at |       | datetime |       |       |       |
| updated_at |       | datetime |       |       |       |
| deleted_at |       | datetime |       |       |       |