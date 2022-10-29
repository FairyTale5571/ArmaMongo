class CfgPatches
{
	class RRPServer_Logger
	{
		units[]={};
		weapons[]={};
		requiredVersion=1;
		requiredAddons[]=
		{
			"RRPServer",
			"rimas_agent"
		};
	};
};
class CfgFunctions
{
	class RRPServer_Logger
	{
		class core
		{
			file="\RRPServer_Logger\bootstrap";
			class postInit
			{
				postInit=1;
			};
			class preInit
			{
				preInit=1;
			};
		};
	};
};