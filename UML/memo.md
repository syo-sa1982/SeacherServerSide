## Controllerのメソッド数無駄に多すぎ問題

現在のURLを洗い出す

### user

* example.com/user/index
* example.com/user/add
* example.com/user/auth

### player

* example.com/player/joblist
    * 職業一覧
* example.com/player/list
    * キャラ一覧
* example.com/player/base_make
    * パラメーター生成
* example.com/player/generate
    * 生成パラメータをDBに格納
* example.com/player/skill_setting
    * 技能一覧
* example.com/player/skill_submit
    * 技能振り分け結果登録