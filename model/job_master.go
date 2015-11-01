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
	JobMaster{JobName:"私立探偵"},
	JobMaster{JobName:"医師"},
	JobMaster{JobName:"エンジニア"},
}

var JobSkillList = []JobSkillMaster{
	JobSkillMaster{JobID:1,SkillID:1, SkillType:0},
	JobSkillMaster{JobID:1,SkillID:8, SkillType:0},
	JobSkillMaster{JobID:1,SkillID:23,SkillType:0},
	JobSkillMaster{JobID:1,SkillID:33,SkillType:0},
	JobSkillMaster{JobID:1,SkillID:39,SkillType:0},
}