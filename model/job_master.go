package model

type JobMaster struct {
	ID        uint
	JobName   string
	JobSkills uint
}

type JobSkillMaster struct {
	ID      uint
	JobID   uint `sql:"index"`
	SkillID uint `sql:"index"`
	SkillType  int // 0:必須スキル 1:選択スキル
}


var JobList = []JobMaster{
	JobMaster{ID:1,JobName:"私立探偵"},
	JobMaster{ID:2,JobName:"医師"},
	JobMaster{ID:3,JobName:"エンジニア"},
}

var JobSkillList = []JobSkillMaster{
	JobSkillMaster{JobID:1,SkillID:1, SkillType:0},
	JobSkillMaster{JobID:1,SkillID:8, SkillType:0},
	JobSkillMaster{JobID:1,SkillID:23,SkillType:0},
	JobSkillMaster{JobID:1,SkillID:33,SkillType:0},
	JobSkillMaster{JobID:1,SkillID:39,SkillType:0},
	JobSkillMaster{JobID:1,SkillID:12,SkillType:1},
	JobSkillMaster{JobID:1,SkillID:19,SkillType:1},
	JobSkillMaster{JobID:1,SkillID:41,SkillType:1},
	JobSkillMaster{JobID:1,SkillID:59,SkillType:1},

	JobSkillMaster{JobID:2,SkillID:2, SkillType:0},
	JobSkillMaster{JobID:2,SkillID:4, SkillType:0},
	JobSkillMaster{JobID:2,SkillID:15,SkillType:0},
	JobSkillMaster{JobID:2,SkillID:22,SkillType:0},
	JobSkillMaster{JobID:2,SkillID:28,SkillType:0},
	JobSkillMaster{JobID:2,SkillID:29,SkillType:0},
	JobSkillMaster{JobID:2,SkillID:50,SkillType:0},
	JobSkillMaster{JobID:2,SkillID:46,SkillType:0},

	JobSkillMaster{JobID:3,SkillID:11,SkillType:0},
	JobSkillMaster{JobID:3,SkillID:17,SkillType:0},
	JobSkillMaster{JobID:3,SkillID:20,SkillType:0},
	JobSkillMaster{JobID:3,SkillID:34,SkillType:0},
	JobSkillMaster{JobID:3,SkillID:39,SkillType:0},
	JobSkillMaster{JobID:3,SkillID:43,SkillType:0},
	JobSkillMaster{JobID:3,SkillID:46,SkillType:0},
	JobSkillMaster{JobID:3,SkillID:7, SkillType:1},
	JobSkillMaster{JobID:3,SkillID:31,SkillType:1},
	JobSkillMaster{JobID:3,SkillID:35,SkillType:1},
}