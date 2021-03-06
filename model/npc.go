package model

import (
	"database/sql"
)

// Npcs represents an array of Npc
// swagger:model
type Npcs []*Npc

//Npc represents Non player characters in everquest
// swagger:model
type Npc struct {
	Class            *Class            `json:"class,omitempty"`
	Race             *Race             `json:"race,omitempty"`
	CleanName        string            `json:"cleanName,omitempty"`
	Zone             *Zone             `json:"zone,omitempty"`
	Experience       int64             `json:"experience,omitempty"`
	SpecialAbilities map[string]string `json:"specialAbilitiesList,omitempty"`

	ID                   int64          `json:"ID,omitempty" db:"id"`                                     //`id` int(11) NOT NULL AUTO_INCREMENT,
	Name                 string         `json:"name,omitempty" db:"name"`                                 //`name` text NOT NULL,
	LastName             sql.NullString `json:"lastName,omitempty" db:"lastname"`                         //`lastname` varchar(32) DEFAULT NULL,
	Level                int64          `json:"level,omitempty" db:"level"`                               //`level` tinyint(2) unsigned NOT NULL DEFAULT '0',
	RaceID               int64          `json:"raceID,omitempty" db:"race"`                               //`race` smallint(5) unsigned NOT NULL DEFAULT '0',
	ClassID              int64          `json:"classID,omitempty" db:"class"`                             //`class` tinyint(2) unsigned NOT NULL DEFAULT '0',
	BodyTypeID           int64          `json:"bodyTypeID,omitempty" db:"bodytype"`                       //`bodytype` int(11) NOT NULL DEFAULT '1',
	Hitpoints            int64          `json:"hitpoints,omitempty" db:"hp"`                              //`hp` int(11) NOT NULL DEFAULT '0',
	Mana                 int64          `json:"mana,omitempty" db:"mana"`                                 //`mana` int(11) NOT NULL DEFAULT '0',
	Gender               int64          `json:"gender,omitempty" db:"gender"`                             //`gender` tinyint(2) unsigned NOT NULL DEFAULT '0',
	Texture              int64          `json:"texture,omitempty" db:"texture"`                           //`texture` tinyint(2) unsigned NOT NULL DEFAULT '0',
	HelmTexture          int64          `json:"helmTexture,omitempty" db:"helmtexture"`                   //`helmtexture` tinyint(2) unsigned NOT NULL DEFAULT '0',
	HerosForgeModel      int64          `json:"herosForgeModel,omitempty" db:"herosforgemodel"`           //`herosforgemodel` int(11) NOT NULL DEFAULT '0',
	Size                 float64        `json:"size,omitempty" db:"size"`                                 //`size` float NOT NULL DEFAULT '0',
	HpRegenRate          int64          `json:"hpRegenRate,omitempty" db:"hp_regen_rate"`                 //`hp_regen_rate` int(11) unsigned NOT NULL DEFAULT '0',
	ManaRegenRate        int64          `json:"manaRegenRate,omitempty" db:"mana_regen_rate"`             //`mana_regen_rate` int(11) unsigned NOT NULL DEFAULT '0',
	LootID               int64          `json:"lootID,omitempty" db:"loottable_id"`                       //`loottable_id` int(11) unsigned NOT NULL DEFAULT '0',
	MerchantID           int64          `json:"merchantID,omitempty" db:"merchant_id"`                    //`merchant_id` int(11) unsigned NOT NULL DEFAULT '0',
	AltCurrencyID        int64          `json:"altCurrencyID,omitempty" db:"alt_currency_id"`             //`alt_currency_id` int(11) unsigned NOT NULL DEFAULT '0',
	NpcSpellsID          int64          `json:"npcSpellsID,omitempty" db:"npc_spells_id"`                 //`npc_spells_id` int(11) unsigned NOT NULL DEFAULT '0',
	NpcSpellsEffectsID   int64          `json:"npcSpellsEffectsID,omitempty" db:"npc_spells_effects_id"`  //`npc_spells_effects_id` int(11) unsigned NOT NULL DEFAULT '0',
	NpcFactionID         int64          `json:"npcFactionID,omitempty" db:"npc_faction_id"`               //`npc_faction_id` int(11) NOT NULL DEFAULT '0',
	AdventureTemplateID  int64          `json:"adventureTemplateID,omitempty" db:"adventure_template_id"` //`adventure_template_id` int(10) unsigned NOT NULL DEFAULT '0',
	TrapTemplate         sql.NullInt64  `json:"trapTemplate,omitempty" db:"trap_template"`                //`trap_template` int(10) unsigned DEFAULT '0',
	MininumDamage        int64          `json:"minimumDamage,omitempty" db:"mindmg"`                      //`mindmg` int(10) unsigned NOT NULL DEFAULT '0',
	MaximumDamage        int64          `json:"maximumDamage,omitempty" db:"maxdmg"`                      //`maxdmg` int(10) unsigned NOT NULL DEFAULT '0',
	AttackCount          int64          `json:"attackCount,omitempty" db:"attack_count"`                  //`attack_count` smallint(6) NOT NULL DEFAULT '-1',
	NpcSpecialAttacks    string         `json:"npcSpecialAttacks,omitempty" db:"npcspecialattks"`         //`npcspecialattks` varchar(36) NOT NULL DEFAULT '',
	SpecialAbilitiesRaw  sql.NullString `json:"specialAbilitiesRaw,omitempty" db:"special_abilities"`     //`special_abilities` text,
	AggroRadius          int64          `json:"aggroRadius,omitempty" db:"aggroradius"`                   //`aggroradius` int(10) unsigned NOT NULL DEFAULT '0',
	AssistRadius         int64          `json:"assistRadius,omitempty" db:"assistradius"`                 //`assistradius` int(10) unsigned NOT NULL DEFAULT '0',
	Face                 int64          `json:"face,omitempty" db:"face"`                                 //`face` int(10) unsigned NOT NULL DEFAULT '1',
	LuclinHairStyle      int64          `json:"luclinHairStyle,omitempty" db:"luclin_hairstyle"`          //`luclin_hairstyle` int(10) unsigned NOT NULL DEFAULT '1',
	LuclinHairColor      int64          `json:"luclinHairColor,omitempty" db:"luclin_haircolor"`          //`luclin_haircolor` int(10) unsigned NOT NULL DEFAULT '1',
	LuclinEyeColor       int64          `json:"luclinEyeColor,omitempty" db:"luclin_eyecolor"`            //`luclin_eyecolor` int(10) unsigned NOT NULL DEFAULT '1',
	LuclinEyeColor2      int64          `json:"luclinEyeColor2,omitempty" db:"luclin_eyecolor2"`          //`luclin_eyecolor2` int(10) unsigned NOT NULL DEFAULT '1',
	LuclinBeardColor     int64          `json:"luclinBeardColor,omitempty" db:"luclin_beardcolor"`        //`luclin_beardcolor` int(10) unsigned NOT NULL DEFAULT '1',
	LuclinBeard          int64          `json:"luclinBeard,omitempty" db:"luclin_beard"`                  //`luclin_beard` int(10) unsigned NOT NULL DEFAULT '0',
	DrakkinHeritage      int64          `json:"drakkinHeritage,omitempty" db:"drakkin_heritage"`          //`drakkin_heritage` int(10) NOT NULL DEFAULT '0',
	DrakkinTattoo        int64          `json:"drakkinTattoo,omitempty" db:"drakkin_tattoo"`              //`drakkin_tattoo` int(10) NOT NULL DEFAULT '0',
	DrakkinDetails       int64          `json:"drakkinDetails,omitempty" db:"drakkin_details"`            //`drakkin_details` int(10) NOT NULL DEFAULT '0',
	ArmorTintID          int64          `json:"armorTintID,omitempty" db:"armortint_id"`                  //`armortint_id` int(10) unsigned NOT NULL DEFAULT '0',
	ArmorTintRed         int64          `json:"armorTintRed,omitempty" db:"armortint_red"`                //`armortint_red` tinyint(3) unsigned NOT NULL DEFAULT '0',
	ArmorTintGreen       int64          `json:"armorTintGreen,omitempty" db:"armortint_green"`            //`armortint_green` tinyint(3) unsigned NOT NULL DEFAULT '0',
	ArmorTintBlue        int64          `json:"armorTintBlue,omitempty" db:"armortint_blue"`              //`armortint_blue` tinyint(3) unsigned NOT NULL DEFAULT '0',
	DMeleeTexture1       int64          `json:"dMeleeTexture1,omitempty" db:"d_melee_texture1"`           //`d_melee_texture1` int(11) NOT NULL DEFAULT '0',
	DMeleeTexture2       int64          `json:"dMeleeTexture2,omitempty" db:"d_melee_texture2"`           //`d_melee_texture2` int(11) NOT NULL DEFAULT '0',
	AmmoIDFile           string         `json:"ammoIDFile,omitempty" db:"ammo_idfile"`                    //`ammo_idfile` varchar(30) NOT NULL DEFAULT 'IT10',
	PrimaryMeleeTypeID   int64          `json:"primaryMeleeTypeID,omitempty" db:"prim_melee_type"`        //`prim_melee_type` tinyint(4) unsigned NOT NULL DEFAULT '28',
	SecondaryMeleeTypeID int64          `json:"secondaryMeleeTypeID,omitempty" db:"sec_melee_type"`       //`sec_melee_type` tinyint(4) unsigned NOT NULL DEFAULT '28',
	RangedTypeID         int64          `json:"rangedTypeID,omitempty" db:"ranged_type"`                  //`ranged_type` tinyint(4) unsigned NOT NULL DEFAULT '7',
	Runspeed             float64        `json:"runspeed,omitempty" db:"runspeed"`                         //`runspeed` float NOT NULL DEFAULT '0',
	MagicResistance      int64          `json:"magicResistance,omitempty" db:"MR"`                        //`MR` smallint(5) NOT NULL DEFAULT '0',
	ColdResistance       int64          `json:"coldResistance,omitempty" db:"CR"`                         //`CR` smallint(5) NOT NULL DEFAULT '0',
	DiseaseResistance    int64          `json:"diseaseResistance,omitempty" db:"DR"`                      //`DR` smallint(5) NOT NULL DEFAULT '0',
	FireResistance       int64          `json:"fireResistance,omitempty" db:"FR"`                         //`FR` smallint(5) NOT NULL DEFAULT '0',
	PoisonResistance     int64          `json:"poisonResistance,omitempty" db:"PR"`                       //`PR` smallint(5) NOT NULL DEFAULT '0',
	CorruptionResistance int64          `json:"corruptionResistance,omitempty" db:"Corrup"`               //`Corrup` smallint(5) NOT NULL DEFAULT '0',
	PhysicalResistance   int64          `json:"physicalResistance,omitempty" db:"PhR"`                    //`PhR` smallint(5) unsigned NOT NULL DEFAULT '0',
	SeeInvisible         int64          `json:"seeInvisible,omitempty" db:"see_invis"`                    //`see_invis` smallint(4) NOT NULL DEFAULT '0',
	SeeInvisibleUndead   int64          `json:"seeInvisibleUndead,omitempty" db:"see_invis_undead"`       //`see_invis_undead` smallint(4) NOT NULL DEFAULT '0',
	QuestGlobal          int64          `json:"questGlobal,omitempty" db:"qglobal"`                       //`qglobal` int(2) unsigned NOT NULL DEFAULT '0',
	ArmorClass           int64          `json:"armorClass,omitempty" db:"AC"`                             //`AC` smallint(5) NOT NULL DEFAULT '0',
	NpcAggro             int64          `json:"npcAggro,omitempty" db:"npc_aggro"`                        //`npc_aggro` tinyint(4) NOT NULL DEFAULT '0',
	SpawnLimit           int64          `json:"spawnLimit,omitempty" db:"spawn_limit"`                    //`spawn_limit` tinyint(4) NOT NULL DEFAULT '0',
	AttackSpeed          float64        `json:"attackSpeed,omitempty" db:"attack_speed"`                  //`attack_speed` float NOT NULL DEFAULT '0',
	AttackDelay          int64          `json:"attackDelay,omitempty" db:"attack_delay"`                  //`attack_delay` tinyint(3) unsigned NOT NULL DEFAULT '30',
	Findable             int64          `json:"findable,omitempty" db:"findable"`                         //`findable` tinyint(4) NOT NULL DEFAULT '0',
	Strength             int64          `json:"strength,omitempty" db:"STR"`                              //`STR` mediumint(8) unsigned NOT NULL DEFAULT '75',
	Stamina              int64          `json:"stamina,omitempty" db:"STA"`                               //`STA` mediumint(8) unsigned NOT NULL DEFAULT '75',
	Dexterity            int64          `json:"dexterity,omitempty" db:"DEX"`                             //`DEX` mediumint(8) unsigned NOT NULL DEFAULT '75',
	Agility              int64          `json:"agility,omitempty" db:"AGI"`                               //`AGI` mediumint(8) unsigned NOT NULL DEFAULT '75',
	Intelligence         int64          `json:"intelligence,omitempty" db:"_INT"`                         //`_INT` mediumint(8) unsigned NOT NULL DEFAULT '80',
	Wisdom               int64          `json:"wisdom,omitempty" db:"WIS"`                                //`WIS` mediumint(8) unsigned NOT NULL DEFAULT '75',
	Charisma             int64          `json:"charisma,omitempty" db:"CHA"`                              //`CHA` mediumint(8) unsigned NOT NULL DEFAULT '75',
	SeeHide              int64          `json:"seeHide,omitempty" db:"see_hide"`                          //`see_hide` tinyint(4) NOT NULL DEFAULT '0',
	SeeImprovedHide      int64          `json:"seeImprovedHide,omitempty" db:"see_improved_hide"`         //`see_improved_hide` tinyint(4) NOT NULL DEFAULT '0',
	Trackable            int64          `json:"trackable,omitempty" db:"trackable"`                       //`trackable` tinyint(4) NOT NULL DEFAULT '1',
	IsBot                int64          `json:"isBot,omitempty" db:"isbot"`                               //`isbot` tinyint(4) NOT NULL DEFAULT '0',
	Exclude              int64          `json:"exclude,omitempty" db:"exclude"`                           //`exclude` tinyint(4) NOT NULL DEFAULT '1',
	Attack               int64          `json:"Attack,omitempty" db:"ATK"`                                //`ATK` mediumint(9) NOT NULL DEFAULT '0',
	Accuracy             int64          `json:"Accuracy,omitempty" db:"Accuracy"`                         //`Accuracy` mediumint(9) NOT NULL DEFAULT '0',
	Avoidance            int64          `json:"Avoidance,omitempty" db:"Avoidance"`                       //`Avoidance` mediumint(9) unsigned NOT NULL DEFAULT '0',
	SlowMitigation       int64          `json:"slowMitigation,omitempty" db:"slow_mitigation"`            //`slow_mitigation` smallint(4) NOT NULL DEFAULT '0',
	Version              int64          `json:"version,omitempty" db:"version"`                           //`version` smallint(5) unsigned NOT NULL DEFAULT '0',
	MaxLevel             int64          `json:"maxLevel,omitempty" db:"maxlevel"`                         //`maxlevel` tinyint(3) NOT NULL DEFAULT '0',
	ScaleRate            int64          `json:"scaleRate,omitempty" db:"scalerate"`                       //`scalerate` int(11) NOT NULL DEFAULT '100',
	PrivateCorpse        int64          `json:"privateCorpse,omitempty" db:"private_corpse"`              //`private_corpse` tinyint(3) unsigned NOT NULL DEFAULT '0',
	UniqueSpawnByName    int64          `json:"uniqueSpawnByName,omitempty" db:"unique_spawn_by_name"`    //`unique_spawn_by_name` tinyint(3) unsigned NOT NULL DEFAULT '0',
	Underwater           int64          `json:"underwater,omitempty" db:"underwater"`                     //`underwater` tinyint(3) unsigned NOT NULL DEFAULT '0',
	IsQuest              int64          `json:"isQuest,omitempty" db:"isquest"`                           //`isquest` tinyint(3) NOT NULL DEFAULT '0',
	EmoteID              int64          `json:"emoteID,omitempty" db:"emoteid"`                           //`emoteid` int(10) unsigned NOT NULL DEFAULT '0',
	SpellScale           float64        `json:"spellScale,omitempty" db:"spellscale"`                     //`spellscale` float NOT NULL DEFAULT '100',
	HealScale            float64        `json:"healScale,omitempty" db:"healscale"`                       //`healscale` float NOT NULL DEFAULT '100',
	NoTargetHotkey       int64          `json:"noTargetHotkey,omitempty" db:"no_target_hotkey"`           //`no_target_hotkey` tinyint(1) unsigned NOT NULL DEFAULT '0',
	RaidTarget           int64          `json:"raidTarget,omitempty" db:"raid_target"`                    //`raid_target` tinyint(1) unsigned NOT NULL DEFAULT '0',
	ArmTexture           int64          `json:"armTexture,omitempty" db:"armtexture"`                     //`armtexture` tinyint(2) NOT NULL DEFAULT '0',
	BracerTexture        int64          `json:"bracerTexture,omitempty" db:"bracertexture"`               //`bracertexture` tinyint(2) NOT NULL DEFAULT '0',
	HandTexture          int64          `json:"handTexture,omitempty" db:"handtexture"`                   //`handtexture` tinyint(2) NOT NULL DEFAULT '0',
	LegTexture           int64          `json:"legTexture,omitempty" db:"legtexture"`                     //`legtexture` tinyint(2) NOT NULL DEFAULT '0',
	FeetTexture          int64          `json:"feetTexture,omitempty" db:"feettexture"`                   //`feettexture` tinyint(2) NOT NULL DEFAULT '0',
	Light                int64          `json:"light,omitempty" db:"light"`                               //`light` tinyint(2) NOT NULL DEFAULT '0',
	WalkSpeed            int64          `json:"walkSpeed,omitempty" db:"walkspeed"`                       //`walkspeed` tinyint(2) NOT NULL DEFAULT '0',
	PeqID                int64          `json:"PeqID,omitempty" db:"peqid"`                               //`peqid` int(11) NOT NULL DEFAULT '0',
	Unique               int64          `json:"unique,omitempty" db:"unique_"`                            //`unique_` tinyint(2) NOT NULL DEFAULT '0',
	Fixed                int64          `json:"fixed,omitempty" db:"fixed"`                               //`fixed` tinyint(2) NOT NULL DEFAULT '0',
	IgnoreDespawn        int64          `json:"ignoreDespawn,omitempty" db:"ignore_despawn"`              //`ignore_despawn` tinyint(2) NOT NULL DEFAULT '0',
	ShowName             int64          `json:"showName,omitempty" db:"show_name"`                        //`show_name` tinyint(2) NOT NULL DEFAULT '1',
	Untargetable         int64          `json:"untargetable,omitempty" db:"untargetable"`                 //`untargetable` tinyint(2) NOT NULL DEFAULT '0',
}
