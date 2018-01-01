package model

import ()

//Activity is used by Tasks to break down which activities are involved in a provided task
type Activity struct {
	Taskid       int64  `json:"taskid" db:"taskid"`             //`taskid` int(11) unsigned NOT NULL DEFAULT '0',
	Activityid   int64  `json:"activityid" db:"activityid"`     //`activityid` int(11) unsigned NOT NULL DEFAULT '0',
	Step         int64  `json:"step" db:"step"`                 //`step` int(11) unsigned NOT NULL DEFAULT '0',
	Activitytype int64  `json:"activitytype" db:"activitytype"` //`activitytype` tinyint(3) unsigned NOT NULL DEFAULT '0',
	Text1        string `json:"text1" db:"text1"`               //`text1` varchar(64) NOT NULL DEFAULT '',
	Text2        string `json:"text2" db:"text2"`               //`text2` varchar(64) NOT NULL DEFAULT '',
	Text3        string `json:"text3" db:"text3"`               //`text3` varchar(128) NOT NULL DEFAULT '',
	Goalid       int64  `json:"goalid" db:"goalid"`             //`goalid` int(11) unsigned NOT NULL DEFAULT '0',
	Goalmethod   int64  `json:"goalmethod" db:"goalmethod"`     //`goalmethod` int(10) unsigned NOT NULL DEFAULT '0',
	Goalcount    int64  `json:"goalcount" db:"goalcount"`       //`goalcount` int(11) DEFAULT '1',
	Delivertonpc int64  `json:"delivertonpc" db:"delivertonpc"` //`delivertonpc` int(11) unsigned NOT NULL DEFAULT '0',
	Zoneid       int64  `json:"zoneid" db:"zoneid"`             //`zoneid` int(11) NOT NULL DEFAULT '0',
	Optional     int64  `json:"optional" db:"optional"`         //`optional` tinyint(1) NOT NULL DEFAULT '0',
}
