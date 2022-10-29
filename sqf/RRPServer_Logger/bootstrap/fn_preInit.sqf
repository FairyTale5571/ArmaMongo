private ["_code", "_function", "_file"];
private _headerNoDebug = '
	private _fnc_scriptNameParent = if (isNil "_fnc_scriptName") then {"%1"} else {_fnc_scriptName};
	private _fnc_scriptName = "%1";
	scriptName _fnc_scriptName;
';
private _is_dev = serverName isEqualTo "#DEV";

private _path = "\RRPServer_Logger\";

//-- Server
[
	"RRPServer_logger_setup",
	"RRPServer_logger_write"
]apply{
	_function = _x;
	_code = if (_is_dev) then {compile (format[_headerNoDebug,_function] + preprocessFileLineNumbers format[_path + "code\%1.sqf",_function])
		} else {
			compileFinal (format[_headerNoDebug,_function] + preprocessFileLineNumbers format[_path + "code\%1.sqf",_function])
	};
	missionNamespace setVariable [_function, _code];
};

call RRPServer_logger_setup;

true
