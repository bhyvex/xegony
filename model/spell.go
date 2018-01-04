package model

import (
	"database/sql"
)

//Spell represents items inside everquest
type Spell struct {
	ID                  int64          `json:"id" db:"id"`                                   //`id` int(11) NOT NULL DEFAULT '0',
	Name                sql.NullString `json:"name" db:"name"`                               //`name` varchar(64) DEFAULT NULL,
	Player1             sql.NullString `json:"player1" db:"player_1"`                        //`player_1` varchar(64) DEFAULT 'BLUE_TRAIL',
	TeleportZone        sql.NullString `json:"teleportZone" db:"teleport_zone"`              //`teleport_zone` varchar(64) DEFAULT NULL,
	YouCast             sql.NullString `json:"youCast" db:"you_cast"`                        //`you_cast` varchar(120) DEFAULT NULL,
	OtherCasts          sql.NullString `json:"otherCasts" db:"other_casts"`                  //`other_casts` varchar(120) DEFAULT NULL,
	CastOnYou           sql.NullString `json:"castOnYou" db:"cast_on_you"`                   //`cast_on_you` varchar(120) DEFAULT NULL,
	CastOnOther         sql.NullString `json:"castOnOther" db:"cast_on_other"`               //`cast_on_other` varchar(120) DEFAULT NULL,
	SpellFades          sql.NullString `json:"spellFades" db:"spell_fades"`                  //`spell_fades` varchar(120) DEFAULT NULL,
	Range               int64          `json:"range" db:"range"`                             //`range` int(11) NOT NULL DEFAULT '100',
	Aoerange            int64          `json:"aoerange" db:"aoerange"`                       //`aoerange` int(11) NOT NULL DEFAULT '0',
	Pushback            int64          `json:"pushback" db:"pushback"`                       //`pushback` int(11) NOT NULL DEFAULT '0',
	Pushup              int64          `json:"pushup" db:"pushup"`                           //`pushup` int(11) NOT NULL DEFAULT '0',
	CastTime            int64          `json:"castTime" db:"cast_time"`                      //`cast_time` int(11) NOT NULL DEFAULT '0',
	RecoveryTime        int64          `json:"recoveryTime" db:"recovery_time"`              //`recovery_time` int(11) NOT NULL DEFAULT '0',
	RecastTime          int64          `json:"recastTime" db:"recast_time"`                  //`recast_time` int(11) NOT NULL DEFAULT '0',
	Buffdurationformula int64          `json:"buffdurationformula" db:"buffdurationformula"` //`buffdurationformula` int(11) NOT NULL DEFAULT '7',
	Buffduration        int64          `json:"buffduration" db:"buffduration"`               //`buffduration` int(11) NOT NULL DEFAULT '65',
	Aeduration          int64          `json:"AEDuration" db:"AEDuration"`                   //`AEDuration` int(11) NOT NULL DEFAULT '0',
	Mana                int64          `json:"mana" db:"mana"`                               //`mana` int(11) NOT NULL DEFAULT '0',
	EffectBaseValue1    int64          `json:"effectBaseValue1" db:"effect_base_value1"`     //`effect_base_value1` int(11) NOT NULL DEFAULT '100',
	EffectBaseValue2    int64          `json:"effectBaseValue2" db:"effect_base_value2"`     //`effect_base_value2` int(11) NOT NULL DEFAULT '0',
	EffectBaseValue3    int64          `json:"effectBaseValue3" db:"effect_base_value3"`     //`effect_base_value3` int(11) NOT NULL DEFAULT '0',
	EffectBaseValue4    int64          `json:"effectBaseValue4" db:"effect_base_value4"`     //`effect_base_value4` int(11) NOT NULL DEFAULT '0',
	EffectBaseValue5    int64          `json:"effectBaseValue5" db:"effect_base_value5"`     //`effect_base_value5` int(11) NOT NULL DEFAULT '0',
	EffectBaseValue6    int64          `json:"effectBaseValue6" db:"effect_base_value6"`     //`effect_base_value6` int(11) NOT NULL DEFAULT '0',
	EffectBaseValue7    int64          `json:"effectBaseValue7" db:"effect_base_value7"`     //`effect_base_value7` int(11) NOT NULL DEFAULT '0',
	EffectBaseValue8    int64          `json:"effectBaseValue8" db:"effect_base_value8"`     //`effect_base_value8` int(11) NOT NULL DEFAULT '0',
	EffectBaseValue9    int64          `json:"effectBaseValue9" db:"effect_base_value9"`     //`effect_base_value9` int(11) NOT NULL DEFAULT '0',
	EffectBaseValue10   int64          `json:"effectBaseValue10" db:"effect_base_value10"`   //`effect_base_value10` int(11) NOT NULL DEFAULT '0',
	EffectBaseValue11   int64          `json:"effectBaseValue11" db:"effect_base_value11"`   //`effect_base_value11` int(11) NOT NULL DEFAULT '0',
	EffectBaseValue12   int64          `json:"effectBaseValue12" db:"effect_base_value12"`   //`effect_base_value12` int(11) NOT NULL DEFAULT '0',
	EffectLimitValue1   int64          `json:"effectLimitValue1" db:"effect_limit_value1"`   //`effect_limit_value1` int(11) NOT NULL DEFAULT '0',
	EffectLimitValue2   int64          `json:"effectLimitValue2" db:"effect_limit_value2"`   //`effect_limit_value2` int(11) NOT NULL DEFAULT '0',
	EffectLimitValue3   int64          `json:"effectLimitValue3" db:"effect_limit_value3"`   //`effect_limit_value3` int(11) NOT NULL DEFAULT '0',
	EffectLimitValue4   int64          `json:"effectLimitValue4" db:"effect_limit_value4"`   //`effect_limit_value4` int(11) NOT NULL DEFAULT '0',
	EffectLimitValue5   int64          `json:"effectLimitValue5" db:"effect_limit_value5"`   //`effect_limit_value5` int(11) NOT NULL DEFAULT '0',
	EffectLimitValue6   int64          `json:"effectLimitValue6" db:"effect_limit_value6"`   //`effect_limit_value6` int(11) NOT NULL DEFAULT '0',
	EffectLimitValue7   int64          `json:"effectLimitValue7" db:"effect_limit_value7"`   //`effect_limit_value7` int(11) NOT NULL DEFAULT '0',
	EffectLimitValue8   int64          `json:"effectLimitValue8" db:"effect_limit_value8"`   //`effect_limit_value8` int(11) NOT NULL DEFAULT '0',
	EffectLimitValue9   int64          `json:"effectLimitValue9" db:"effect_limit_value9"`   //`effect_limit_value9` int(11) NOT NULL DEFAULT '0',
	EffectLimitValue10  int64          `json:"effectLimitValue10" db:"effect_limit_value10"` //`effect_limit_value10` int(11) NOT NULL DEFAULT '0',
	EffectLimitValue11  int64          `json:"effectLimitValue11" db:"effect_limit_value11"` //`effect_limit_value11` int(11) NOT NULL DEFAULT '0',
	EffectLimitValue12  int64          `json:"effectLimitValue12" db:"effect_limit_value12"` //`effect_limit_value12` int(11) NOT NULL DEFAULT '0',
	Max1                int64          `json:"max1" db:"max1"`                               //`max1` int(11) NOT NULL DEFAULT '0',
	Max2                int64          `json:"max2" db:"max2"`                               //`max2` int(11) NOT NULL DEFAULT '0',
	Max3                int64          `json:"max3" db:"max3"`                               //`max3` int(11) NOT NULL DEFAULT '0',
	Max4                int64          `json:"max4" db:"max4"`                               //`max4` int(11) NOT NULL DEFAULT '0',
	Max5                int64          `json:"max5" db:"max5"`                               //`max5` int(11) NOT NULL DEFAULT '0',
	Max6                int64          `json:"max6" db:"max6"`                               //`max6` int(11) NOT NULL DEFAULT '0',
	Max7                int64          `json:"max7" db:"max7"`                               //`max7` int(11) NOT NULL DEFAULT '0',
	Max8                int64          `json:"max8" db:"max8"`                               //`max8` int(11) NOT NULL DEFAULT '0',
	Max9                int64          `json:"max9" db:"max9"`                               //`max9` int(11) NOT NULL DEFAULT '0',
	Max10               int64          `json:"max10" db:"max10"`                             //`max10` int(11) NOT NULL DEFAULT '0',
	Max11               int64          `json:"max11" db:"max11"`                             //`max11` int(11) NOT NULL DEFAULT '0',
	Max12               int64          `json:"max12" db:"max12"`                             //`max12` int(11) NOT NULL DEFAULT '0',
	Icon                int64          `json:"icon" db:"icon"`                               //`icon` int(11) NOT NULL DEFAULT '0',
	Memicon             int64          `json:"memicon" db:"memicon"`                         //`memicon` int(11) NOT NULL DEFAULT '0',
	Components1         int64          `json:"components1" db:"components1"`                 //`components1` int(11) NOT NULL DEFAULT '-1',
	Components2         int64          `json:"components2" db:"components2"`                 //`components2` int(11) NOT NULL DEFAULT '-1',
	Components3         int64          `json:"components3" db:"components3"`                 //`components3` int(11) NOT NULL DEFAULT '-1',
	Components4         int64          `json:"components4" db:"components4"`                 //`components4` int(11) NOT NULL DEFAULT '-1',
	ComponentCounts1    int64          `json:"componentCounts1" db:"component_counts1"`      //`component_counts1` int(11) NOT NULL DEFAULT '1',
	ComponentCounts2    int64          `json:"componentCounts2" db:"component_counts2"`      //`component_counts2` int(11) NOT NULL DEFAULT '1',
	ComponentCounts3    int64          `json:"componentCounts3" db:"component_counts3"`      //`component_counts3` int(11) NOT NULL DEFAULT '1',
	ComponentCounts4    int64          `json:"componentCounts4" db:"component_counts4"`      //`component_counts4` int(11) NOT NULL DEFAULT '1',
	Noexpendreagent1    int64          `json:"NoexpendReagent1" db:"NoexpendReagent1"`       //`NoexpendReagent1` int(11) NOT NULL DEFAULT '-1',
	Noexpendreagent2    int64          `json:"NoexpendReagent2" db:"NoexpendReagent2"`       //`NoexpendReagent2` int(11) NOT NULL DEFAULT '-1',
	Noexpendreagent3    int64          `json:"NoexpendReagent3" db:"NoexpendReagent3"`       //`NoexpendReagent3` int(11) NOT NULL DEFAULT '-1',
	Noexpendreagent4    int64          `json:"NoexpendReagent4" db:"NoexpendReagent4"`       //`NoexpendReagent4` int(11) NOT NULL DEFAULT '-1',
	Formula1            int64          `json:"formula1" db:"formula1"`                       //`formula1` int(11) NOT NULL DEFAULT '100',
	Formula2            int64          `json:"formula2" db:"formula2"`                       //`formula2` int(11) NOT NULL DEFAULT '100',
	Formula3            int64          `json:"formula3" db:"formula3"`                       //`formula3` int(11) NOT NULL DEFAULT '100',
	Formula4            int64          `json:"formula4" db:"formula4"`                       //`formula4` int(11) NOT NULL DEFAULT '100',
	Formula5            int64          `json:"formula5" db:"formula5"`                       //`formula5` int(11) NOT NULL DEFAULT '100',
	Formula6            int64          `json:"formula6" db:"formula6"`                       //`formula6` int(11) NOT NULL DEFAULT '100',
	Formula7            int64          `json:"formula7" db:"formula7"`                       //`formula7` int(11) NOT NULL DEFAULT '100',
	Formula8            int64          `json:"formula8" db:"formula8"`                       //`formula8` int(11) NOT NULL DEFAULT '100',
	Formula9            int64          `json:"formula9" db:"formula9"`                       //`formula9` int(11) NOT NULL DEFAULT '100',
	Formula10           int64          `json:"formula10" db:"formula10"`                     //`formula10` int(11) NOT NULL DEFAULT '100',
	Formula11           int64          `json:"formula11" db:"formula11"`                     //`formula11` int(11) NOT NULL DEFAULT '100',
	Formula12           int64          `json:"formula12" db:"formula12"`                     //`formula12` int(11) NOT NULL DEFAULT '100',
	Lighttype           int64          `json:"LightType" db:"LightType"`                     //`LightType` int(11) NOT NULL DEFAULT '0',
	Goodeffect          int64          `json:"goodEffect" db:"goodEffect"`                   //`goodEffect` int(11) NOT NULL DEFAULT '0',
	Activated           int64          `json:"Activated" db:"Activated"`                     //`Activated` int(11) NOT NULL DEFAULT '0',
	Resisttype          int64          `json:"resisttype" db:"resisttype"`                   //`resisttype` int(11) NOT NULL DEFAULT '0',
	Effectid1           int64          `json:"effectid1" db:"effectid1"`                     //`effectid1` int(11) NOT NULL DEFAULT '254',
	Effectid2           int64          `json:"effectid2" db:"effectid2"`                     //`effectid2` int(11) NOT NULL DEFAULT '254',
	Effectid3           int64          `json:"effectid3" db:"effectid3"`                     //`effectid3` int(11) NOT NULL DEFAULT '254',
	Effectid4           int64          `json:"effectid4" db:"effectid4"`                     //`effectid4` int(11) NOT NULL DEFAULT '254',
	Effectid5           int64          `json:"effectid5" db:"effectid5"`                     //`effectid5` int(11) NOT NULL DEFAULT '254',
	Effectid6           int64          `json:"effectid6" db:"effectid6"`                     //`effectid6` int(11) NOT NULL DEFAULT '254',
	Effectid7           int64          `json:"effectid7" db:"effectid7"`                     //`effectid7` int(11) NOT NULL DEFAULT '254',
	Effectid8           int64          `json:"effectid8" db:"effectid8"`                     //`effectid8` int(11) NOT NULL DEFAULT '254',
	Effectid9           int64          `json:"effectid9" db:"effectid9"`                     //`effectid9` int(11) NOT NULL DEFAULT '254',
	Effectid10          int64          `json:"effectid10" db:"effectid10"`                   //`effectid10` int(11) NOT NULL DEFAULT '254',
	Effectid11          int64          `json:"effectid11" db:"effectid11"`                   //`effectid11` int(11) NOT NULL DEFAULT '254',
	Effectid12          int64          `json:"effectid12" db:"effectid12"`                   //`effectid12` int(11) NOT NULL DEFAULT '254',
	Targettype          int64          `json:"targettype" db:"targettype"`                   //`targettype` int(11) NOT NULL DEFAULT '2',
	Basediff            int64          `json:"basediff" db:"basediff"`                       //`basediff` int(11) NOT NULL DEFAULT '0',
	Skill               int64          `json:"skill" db:"skill"`                             //`skill` int(11) NOT NULL DEFAULT '98',
	Zonetype            int64          `json:"zonetype" db:"zonetype"`                       //`zonetype` int(11) NOT NULL DEFAULT '-1',
	Environmenttype     int64          `json:"EnvironmentType" db:"EnvironmentType"`         //`EnvironmentType` int(11) NOT NULL DEFAULT '0',
	Timeofday           int64          `json:"TimeOfDay" db:"TimeOfDay"`                     //`TimeOfDay` int(11) NOT NULL DEFAULT '0',
	Classes1            int64          `json:"classes1" db:"classes1"`                       //`classes1` int(11) NOT NULL DEFAULT '255',
	Classes2            int64          `json:"classes2" db:"classes2"`                       //`classes2` int(11) NOT NULL DEFAULT '255',
	Classes3            int64          `json:"classes3" db:"classes3"`                       //`classes3` int(11) NOT NULL DEFAULT '255',
	Classes4            int64          `json:"classes4" db:"classes4"`                       //`classes4` int(11) NOT NULL DEFAULT '255',
	Classes5            int64          `json:"classes5" db:"classes5"`                       //`classes5` int(11) NOT NULL DEFAULT '255',
	Classes6            int64          `json:"classes6" db:"classes6"`                       //`classes6` int(11) NOT NULL DEFAULT '255',
	Classes7            int64          `json:"classes7" db:"classes7"`                       //`classes7` int(11) NOT NULL DEFAULT '255',
	Classes8            int64          `json:"classes8" db:"classes8"`                       //`classes8` int(11) NOT NULL DEFAULT '255',
	Classes9            int64          `json:"classes9" db:"classes9"`                       //`classes9` int(11) NOT NULL DEFAULT '255',
	Classes10           int64          `json:"classes10" db:"classes10"`                     //`classes10` int(11) NOT NULL DEFAULT '255',
	Classes11           int64          `json:"classes11" db:"classes11"`                     //`classes11` int(11) NOT NULL DEFAULT '255',
	Classes12           int64          `json:"classes12" db:"classes12"`                     //`classes12` int(11) NOT NULL DEFAULT '255',
	Classes13           int64          `json:"classes13" db:"classes13"`                     //`classes13` int(11) NOT NULL DEFAULT '255',
	Classes14           int64          `json:"classes14" db:"classes14"`                     //`classes14` int(11) NOT NULL DEFAULT '255',
	Classes15           int64          `json:"classes15" db:"classes15"`                     //`classes15` int(11) NOT NULL DEFAULT '255',
	Classes16           int64          `json:"classes16" db:"classes16"`                     //`classes16` int(11) NOT NULL DEFAULT '255',
	Castinganim         int64          `json:"CastingAnim" db:"CastingAnim"`                 //`CastingAnim` int(11) NOT NULL DEFAULT '44',
	Targetanim          int64          `json:"TargetAnim" db:"TargetAnim"`                   //`TargetAnim` int(11) NOT NULL DEFAULT '13',
	Traveltype          int64          `json:"TravelType" db:"TravelType"`                   //`TravelType` int(11) NOT NULL DEFAULT '0',
	Spellaffectindex    int64          `json:"SpellAffectIndex" db:"SpellAffectIndex"`       //`SpellAffectIndex` int(11) NOT NULL DEFAULT '-1',
	DisallowSit         int64          `json:"disallowSit" db:"disallow_sit"`                //`disallow_sit` int(11) NOT NULL DEFAULT '0',
	Deities0            int64          `json:"deities0" db:"deities0"`                       //`deities0` int(11) NOT NULL DEFAULT '0',
	Deities1            int64          `json:"deities1" db:"deities1"`                       //`deities1` int(11) NOT NULL DEFAULT '0',
	Deities2            int64          `json:"deities2" db:"deities2"`                       //`deities2` int(11) NOT NULL DEFAULT '0',
	Deities3            int64          `json:"deities3" db:"deities3"`                       //`deities3` int(11) NOT NULL DEFAULT '0',
	Deities4            int64          `json:"deities4" db:"deities4"`                       //`deities4` int(11) NOT NULL DEFAULT '0',
	Deities5            int64          `json:"deities5" db:"deities5"`                       //`deities5` int(11) NOT NULL DEFAULT '0',
	Deities6            int64          `json:"deities6" db:"deities6"`                       //`deities6` int(11) NOT NULL DEFAULT '0',
	Deities7            int64          `json:"deities7" db:"deities7"`                       //`deities7` int(11) NOT NULL DEFAULT '0',
	Deities8            int64          `json:"deities8" db:"deities8"`                       //`deities8` int(11) NOT NULL DEFAULT '0',
	Deities9            int64          `json:"deities9" db:"deities9"`                       //`deities9` int(11) NOT NULL DEFAULT '0',
	Deities10           int64          `json:"deities10" db:"deities10"`                     //`deities10` int(11) NOT NULL DEFAULT '0',
	Deities11           int64          `json:"deities11" db:"deities11"`                     //`deities11` int(11) NOT NULL DEFAULT '0',
	Deities12           int64          `json:"deities12" db:"deities12"`                     //`deities12` int(12) NOT NULL DEFAULT '0',
	Deities13           int64          `json:"deities13" db:"deities13"`                     //`deities13` int(11) NOT NULL DEFAULT '0',
	Deities14           int64          `json:"deities14" db:"deities14"`                     //`deities14` int(11) NOT NULL DEFAULT '0',
	Deities15           int64          `json:"deities15" db:"deities15"`                     //`deities15` int(11) NOT NULL DEFAULT '0',
	Deities16           int64          `json:"deities16" db:"deities16"`                     //`deities16` int(11) NOT NULL DEFAULT '0',
	Field142            int64          `json:"field142" db:"field142"`                       //`field142` int(11) NOT NULL DEFAULT '100',
	Field143            int64          `json:"field143" db:"field143"`                       //`field143` int(11) NOT NULL DEFAULT '0',
	NewIcon             int64          `json:"newIcon" db:"new_icon"`                        //`new_icon` int(11) NOT NULL DEFAULT '161',
	Spellanim           int64          `json:"spellanim" db:"spellanim"`                     //`spellanim` int(11) NOT NULL DEFAULT '0',
	Uninterruptable     int64          `json:"uninterruptable" db:"uninterruptable"`         //`uninterruptable` int(11) NOT NULL DEFAULT '0',
	Resistdiff          int64          `json:"ResistDiff" db:"ResistDiff"`                   //`ResistDiff` int(11) NOT NULL DEFAULT '-150',
	DotStackingExempt   int64          `json:"dotStackingExempt" db:"dot_stacking_exempt"`   //`dot_stacking_exempt` int(11) NOT NULL DEFAULT '0',
	Deleteable          int64          `json:"deleteable" db:"deleteable"`                   //`deleteable` int(11) NOT NULL DEFAULT '0',
	Recourselink        int64          `json:"RecourseLink" db:"RecourseLink"`               //`RecourseLink` int(11) NOT NULL DEFAULT '0',
	NoPartialResist     int64          `json:"noPartialResist" db:"no_partial_resist"`       //`no_partial_resist` int(11) NOT NULL DEFAULT '0',
	Field152            int64          `json:"field152" db:"field152"`                       //`field152` int(11) NOT NULL DEFAULT '0',
	Field153            int64          `json:"field153" db:"field153"`                       //`field153` int(11) NOT NULL DEFAULT '0',
	ShortBuffBox        int64          `json:"shortBuffBox" db:"short_buff_box"`             //`short_buff_box` int(11) NOT NULL DEFAULT '-1',
	Descnum             int64          `json:"descnum" db:"descnum"`                         //`descnum` int(11) NOT NULL DEFAULT '0',
	Typedescnum         int64          `json:"typedescnum" db:"typedescnum"`                 //`typedescnum` int(11) DEFAULT NULL,
	Effectdescnum       int64          `json:"effectdescnum" db:"effectdescnum"`             //`effectdescnum` int(11) DEFAULT NULL,
	Effectdescnum2      int64          `json:"effectdescnum2" db:"effectdescnum2"`           //`effectdescnum2` int(11) NOT NULL DEFAULT '0',
	NpcNoLos            int64          `json:"npcNoLos" db:"npc_no_los"`                     //`npc_no_los` int(11) NOT NULL DEFAULT '0',
	Field160            int64          `json:"field160" db:"field160"`                       //`field160` int(11) NOT NULL DEFAULT '0',
	Reflectable         int64          `json:"reflectable" db:"reflectable"`                 //`reflectable` int(11) NOT NULL DEFAULT '0',
	Bonushate           int64          `json:"bonushate" db:"bonushate"`                     //`bonushate` int(11) NOT NULL DEFAULT '0',
	Field163            int64          `json:"field163" db:"field163"`                       //`field163` int(11) NOT NULL DEFAULT '100',
	Field164            int64          `json:"field164" db:"field164"`                       //`field164` int(11) NOT NULL DEFAULT '-150',
	LdonTrap            int64          `json:"ldonTrap" db:"ldon_trap"`                      //`ldon_trap` int(11) NOT NULL DEFAULT '0',
	Endurcost           int64          `json:"EndurCost" db:"EndurCost"`                     //`EndurCost` int(11) NOT NULL DEFAULT '0',
	Endurtimerindex     int64          `json:"EndurTimerIndex" db:"EndurTimerIndex"`         //`EndurTimerIndex` int(11) NOT NULL DEFAULT '0',
	Isdiscipline        int64          `json:"IsDiscipline" db:"IsDiscipline"`               //`IsDiscipline` int(11) NOT NULL DEFAULT '0',
	Field169            int64          `json:"field169" db:"field169"`                       //`field169` int(11) NOT NULL DEFAULT '0',
	Field170            int64          `json:"field170" db:"field170"`                       //`field170` int(11) NOT NULL DEFAULT '0',
	Field171            int64          `json:"field171" db:"field171"`                       //`field171` int(11) NOT NULL DEFAULT '0',
	Field172            int64          `json:"field172" db:"field172"`                       //`field172` int(11) NOT NULL DEFAULT '0',
	Hateadded           int64          `json:"HateAdded" db:"HateAdded"`                     //`HateAdded` int(11) NOT NULL DEFAULT '0',
	Endurupkeep         int64          `json:"EndurUpkeep" db:"EndurUpkeep"`                 //`EndurUpkeep` int(11) NOT NULL DEFAULT '0',
	Numhitstype         int64          `json:"numhitstype" db:"numhitstype"`                 //`numhitstype` int(11) NOT NULL DEFAULT '0',
	Numhits             int64          `json:"numhits" db:"numhits"`                         //`numhits` int(11) NOT NULL DEFAULT '0',
	Pvpresistbase       int64          `json:"pvpresistbase" db:"pvpresistbase"`             //`pvpresistbase` int(11) NOT NULL DEFAULT '-150',
	Pvpresistcalc       int64          `json:"pvpresistcalc" db:"pvpresistcalc"`             //`pvpresistcalc` int(11) NOT NULL DEFAULT '100',
	Pvpresistcap        int64          `json:"pvpresistcap" db:"pvpresistcap"`               //`pvpresistcap` int(11) NOT NULL DEFAULT '-150',
	SpellCategory       int64          `json:"spellCategory" db:"spell_category"`            //`spell_category` int(11) NOT NULL DEFAULT '-99',
	Field181            int64          `json:"field181" db:"field181"`                       //`field181` int(11) NOT NULL DEFAULT '7',
	Field182            int64          `json:"field182" db:"field182"`                       //`field182` int(11) NOT NULL DEFAULT '65',
	PcnpcOnlyFlag       int64          `json:"pcnpcOnlyFlag" db:"pcnpc_only_flag"`           //`pcnpc_only_flag` int(11) DEFAULT '0',
	CastNotStanding     int64          `json:"castNotStanding" db:"cast_not_standing"`       //`cast_not_standing` int(11) DEFAULT '0',
	CanMgb              int64          `json:"canMgb" db:"can_mgb"`                          //`can_mgb` int(11) NOT NULL DEFAULT '0',
	Nodispell           int64          `json:"nodispell" db:"nodispell"`                     //`nodispell` int(11) NOT NULL DEFAULT '-1',
	NpcCategory         int64          `json:"npcCategory" db:"npc_category"`                //`npc_category` int(11) NOT NULL DEFAULT '0',
	NpcUsefulness       int64          `json:"npcUsefulness" db:"npc_usefulness"`            //`npc_usefulness` int(11) NOT NULL DEFAULT '0',
	Minresist           int64          `json:"MinResist" db:"MinResist"`                     //`MinResist` int(11) NOT NULL DEFAULT '0',
	Maxresist           int64          `json:"MaxResist" db:"MaxResist"`                     //`MaxResist` int(11) NOT NULL DEFAULT '0',
	ViralTargets        int64          `json:"viralTargets" db:"viral_targets"`              //`viral_targets` int(11) NOT NULL DEFAULT '0',
	ViralTimer          int64          `json:"viralTimer" db:"viral_timer"`                  //`viral_timer` int(11) NOT NULL DEFAULT '0',
	Nimbuseffect        int64          `json:"nimbuseffect" db:"nimbuseffect"`               //`nimbuseffect` int(11) DEFAULT '0',
	Conestartangle      int64          `json:"ConeStartAngle" db:"ConeStartAngle"`           //`ConeStartAngle` int(11) NOT NULL DEFAULT '0',
	Conestopangle       int64          `json:"ConeStopAngle" db:"ConeStopAngle"`             //`ConeStopAngle` int(11) NOT NULL DEFAULT '0',
	Sneaking            int64          `json:"sneaking" db:"sneaking"`                       //`sneaking` int(11) NOT NULL DEFAULT '0',
	NotExtendable       int64          `json:"notExtendable" db:"not_extendable"`            //`not_extendable` int(11) NOT NULL DEFAULT '0',
	Field198            int64          `json:"field198" db:"field198"`                       //`field198` int(11) NOT NULL DEFAULT '0',
	Field199            int64          `json:"field199" db:"field199"`                       //`field199` int(11) NOT NULL DEFAULT '1',
	Suspendable         sql.NullInt64  `json:"suspendable" db:"suspendable"`                 //`suspendable` int(11) DEFAULT '0',
	ViralRange          int64          `json:"viralRange" db:"viral_range"`                  //`viral_range` int(11) NOT NULL DEFAULT '0',
	Songcap             sql.NullInt64  `json:"songcap" db:"songcap"`                         //`songcap` int(11) DEFAULT '0',
	Field203            sql.NullInt64  `json:"field203" db:"field203"`                       //`field203` int(11) DEFAULT '0',
	Field204            sql.NullInt64  `json:"field204" db:"field204"`                       //`field204` int(11) DEFAULT '0',
	NoBlock             int64          `json:"noBlock" db:"no_block"`                        //`no_block` int(11) NOT NULL DEFAULT '0',
	Field206            sql.NullInt64  `json:"field206" db:"field206"`                       //`field206` int(11) DEFAULT '-1',
	Spellgroup          sql.NullInt64  `json:"spellgroup" db:"spellgroup"`                   //`spellgroup` int(11) DEFAULT '0',
	Rank                int64          `json:"rank" db:"rank"`                               //`rank` int(11) NOT NULL DEFAULT '0',
	Field209            sql.NullInt64  `json:"field209" db:"field209"`                       //`field209` int(11) DEFAULT '0',
	Field210            sql.NullInt64  `json:"field210" db:"field210"`                       //`field210` int(11) DEFAULT '1',
	Castrestriction     int64          `json:"CastRestriction" db:"CastRestriction"`         //`CastRestriction` int(11) NOT NULL DEFAULT '0',
	Allowrest           sql.NullInt64  `json:"allowrest" db:"allowrest"`                     //`allowrest` int(11) DEFAULT '0',
	Incombat            int64          `json:"InCombat" db:"InCombat"`                       //`InCombat` int(11) NOT NULL DEFAULT '0',
	Outofcombat         int64          `json:"OutofCombat" db:"OutofCombat"`                 //`OutofCombat` int(11) NOT NULL DEFAULT '0',
	Field215            sql.NullInt64  `json:"field215" db:"field215"`                       //`field215` int(11) DEFAULT '0',
	Field216            sql.NullInt64  `json:"field216" db:"field216"`                       //`field216` int(11) DEFAULT '0',
	Field217            sql.NullInt64  `json:"field217" db:"field217"`                       //`field217` int(11) DEFAULT '0',
	Aemaxtargets        int64          `json:"aemaxtargets" db:"aemaxtargets"`               //`aemaxtargets` int(11) NOT NULL DEFAULT '0',
	Maxtargets          sql.NullInt64  `json:"maxtargets" db:"maxtargets"`                   //`maxtargets` int(11) DEFAULT '0',
	Field220            sql.NullInt64  `json:"field220" db:"field220"`                       //`field220` int(11) DEFAULT '0',
	Field221            sql.NullInt64  `json:"field221" db:"field221"`                       //`field221` int(11) DEFAULT '0',
	Field222            sql.NullInt64  `json:"field222" db:"field222"`                       //`field222` int(11) DEFAULT '0',
	Field223            sql.NullInt64  `json:"field223" db:"field223"`                       //`field223` int(11) DEFAULT '0',
	Persistdeath        sql.NullInt64  `json:"persistdeath" db:"persistdeath"`               //`persistdeath` int(11) DEFAULT '0',
	Field225            int64          `json:"field225" db:"field225"`                       //`field225` int(11) NOT NULL DEFAULT '0',
	Field226            int64          `json:"field226" db:"field226"`                       //`field226` int(11) NOT NULL DEFAULT '0',
	MinDist             float64        `json:"minDist" db:"min_dist"`                        //`min_dist` float NOT NULL DEFAULT '0',
	MinDistMod          float64        `json:"minDistMod" db:"min_dist_mod"`                 //`min_dist_mod` float NOT NULL DEFAULT '0',
	MaxDist             float64        `json:"maxDist" db:"max_dist"`                        //`max_dist` float NOT NULL DEFAULT '0',
	MaxDistMod          float64        `json:"maxDistMod" db:"max_dist_mod"`                 //`max_dist_mod` float NOT NULL DEFAULT '0',
	MinRange            int64          `json:"minRange" db:"min_range"`                      //`min_range` int(11) NOT NULL DEFAULT '0',
	Field232            int64          `json:"field232" db:"field232"`                       //`field232` int(11) NOT NULL DEFAULT '0',
	Field233            int64          `json:"field233" db:"field233"`                       //`field233` int(11) NOT NULL DEFAULT '0',
	Field234            int64          `json:"field234" db:"field234"`                       //`field234` int(11) NOT NULL DEFAULT '0',
	Field235            int64          `json:"field235" db:"field235"`                       //`field235` int(11) NOT NULL DEFAULT '0',
	Field236            int64          `json:"field236" db:"field236"`                       //`field236` int(11) NOT NULL DEFAULT '0',
}

//SkillName returns human readable version of Skill Name
func (s *Spell) SkillName() string {
	return SkillName(s.Skill)
}

//LowestLevel returns the lowest level a character can mem this
func (s *Spell) LowestLevel() int64 {
	lowestClass := int64(255)
	if s.Classes1 < lowestClass {
		lowestClass = s.Classes1
	}
	if s.Classes2 < lowestClass {
		lowestClass = s.Classes2
	}
	if s.Classes3 < lowestClass {
		lowestClass = s.Classes3
	}
	if s.Classes4 < lowestClass {
		lowestClass = s.Classes4
	}
	if s.Classes5 < lowestClass {
		lowestClass = s.Classes5
	}
	if s.Classes6 < lowestClass {
		lowestClass = s.Classes6
	}
	if s.Classes7 < lowestClass {
		lowestClass = s.Classes7
	}
	if s.Classes8 < lowestClass {
		lowestClass = s.Classes8
	}
	if s.Classes9 < lowestClass {
		lowestClass = s.Classes9
	}
	if s.Classes10 < lowestClass {
		lowestClass = s.Classes10
	}
	if s.Classes11 < lowestClass {
		lowestClass = s.Classes11
	}
	if s.Classes12 < lowestClass {
		lowestClass = s.Classes12
	}
	if s.Classes13 < lowestClass {
		lowestClass = s.Classes13
	}
	return lowestClass
}

//DescriptionName returns a summary of this spell's attributes
func (s *Spell) DescriptionName() string {
	return "Description"
}