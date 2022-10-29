
params [
	"_collection",
	"_logType",
	"_uid",
	"_notParsedLog"
];


_toParse = _notParsedLog apply { _x joinString "^^"};
_toParse = _toParse joinString "~~";

"mongo_log" callExtension ["write",[_collection,_uid,_logType,_toParse]];
