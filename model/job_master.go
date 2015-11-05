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
//	JobMaster{ID:2,JobName:"医師"},
//	JobMaster{ID:3,JobName:"エンジニア"},
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
}