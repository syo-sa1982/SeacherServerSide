package model

type SkillMaster struct {
	ID         uint
	CategoryID int
	SkillName  string
	Value      int
}
type SkillCategory struct {
	ID           int
	CategoryName string
}


var SkillCategoryList = []SkillCategory{
	SkillCategory{CategoryName:"治癒"},
	SkillCategory{CategoryName:"戦闘"},
	SkillCategory{CategoryName:"探索"},
	SkillCategory{CategoryName:"火器"},
	SkillCategory{CategoryName:"特殊"},
	SkillCategory{CategoryName:"神話"},
}
var SkillMasterList = []SkillMaster{
	SkillMaster{CategoryID:3,SkillName:"言いくるめ",Value:5},
	SkillMaster{CategoryID:1,SkillName:"医学",Value:5},
	SkillMaster{CategoryID:3,SkillName:"運転",Value:20},
	SkillMaster{CategoryID:1,SkillName:"応急手当",Value:20},
	SkillMaster{CategoryID:2,SkillName:"オカルト",Value:5},
	SkillMaster{CategoryID:2,SkillName:"回避",Value:0},
	SkillMaster{CategoryID:3,SkillName:"科学",Value:1},
	SkillMaster{CategoryID:3,SkillName:"鍵開け",Value:1},
	SkillMaster{CategoryID:3,SkillName:"隠す",Value:15},
	SkillMaster{CategoryID:3,SkillName:"隠れる",Value:10},
	SkillMaster{CategoryID:3,SkillName:"機械修理",Value:20},
	SkillMaster{CategoryID:3,SkillName:"聞き耳",Value:25},
	SkillMaster{CategoryID:5,SkillName:"クトゥルフ神話",Value:0},
	SkillMaster{CategoryID:3,SkillName:"芸術",Value:5},
	SkillMaster{CategoryID:3,SkillName:"経理",Value:10},
	SkillMaster{CategoryID:3,SkillName:"考古学",Value:1},
	SkillMaster{CategoryID:3,SkillName:"コンピューター",Value:1},
	SkillMaster{CategoryID:3,SkillName:"忍び歩き",Value:10},
	SkillMaster{CategoryID:3,SkillName:"写真術",Value:10},
	SkillMaster{CategoryID:3,SkillName:"重機械操作",Value:1},
	SkillMaster{CategoryID:3,SkillName:"乗馬",Value:5},
	SkillMaster{CategoryID:3,SkillName:"信用",Value:15},
	SkillMaster{CategoryID:4,SkillName:"心理学",Value:5},
	SkillMaster{CategoryID:3,SkillName:"人類学",Value:1},
	SkillMaster{CategoryID:3,SkillName:"水泳",Value:25},
	SkillMaster{CategoryID:3,SkillName:"製作",Value:5},
	SkillMaster{CategoryID:1,SkillName:"精神分析",Value:1},
	SkillMaster{CategoryID:3,SkillName:"生物学",Value:1},
	SkillMaster{CategoryID:3,SkillName:"説得",Value:15},
	SkillMaster{CategoryID:3,SkillName:"操縦",Value:1},
	SkillMaster{CategoryID:3,SkillName:"地質学",Value:1},
	SkillMaster{CategoryID:3,SkillName:"跳躍",Value:25},
	SkillMaster{CategoryID:3,SkillName:"追跡",Value:10},
	SkillMaster{CategoryID:3,SkillName:"電器修理",Value:10},
	SkillMaster{CategoryID:3,SkillName:"電子工学",Value:1},
	SkillMaster{CategoryID:3,SkillName:"天文学",Value:1},
	SkillMaster{CategoryID:2,SkillName:"投擲",Value:25},
	SkillMaster{CategoryID:3,SkillName:"登攀",Value:40},
	SkillMaster{CategoryID:3,SkillName:"図書館",Value:25},
	SkillMaster{CategoryID:3,SkillName:"ナビゲート",Value:10},
	SkillMaster{CategoryID:3,SkillName:"値切り",Value:5},
	SkillMaster{CategoryID:3,SkillName:"博物学",Value:10},
	SkillMaster{CategoryID:3,SkillName:"物理学",Value:1},
	SkillMaster{CategoryID:3,SkillName:"変装",Value:1},
	SkillMaster{CategoryID:3,SkillName:"法律",Value:5},
	SkillMaster{CategoryID:3,SkillName:"外国語",Value:1},
	SkillMaster{CategoryID:3,SkillName:"母国語",Value:0},
	SkillMaster{CategoryID:2,SkillName:"格闘",Value:1},
	SkillMaster{CategoryID:3,SkillName:"目星",Value:25},
	SkillMaster{CategoryID:3,SkillName:"薬学",Value:1},
	SkillMaster{CategoryID:3,SkillName:"歴史",Value:20},
	SkillMaster{CategoryID:4,SkillName:"拳銃",Value:20},
	SkillMaster{CategoryID:4,SkillName:"サブマシンガン",Value:15},
	SkillMaster{CategoryID:4,SkillName:"ショットガン",Value:30},
	SkillMaster{CategoryID:4,SkillName:"マシンガン",Value:15},
	SkillMaster{CategoryID:4,SkillName:"ライフル",Value:25},
	SkillMaster{CategoryID:2,SkillName:"キック",Value:25},
	SkillMaster{CategoryID:2,SkillName:"組み付き",Value:25},
	SkillMaster{CategoryID:2,SkillName:"こぶし",Value:50},
	SkillMaster{CategoryID:2,SkillName:"頭突き",Value:10},
}
