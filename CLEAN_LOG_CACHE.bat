@echo off

echo Log Cache Tool v1.0 by pjarcher913
echo This tool was made to initialize and/or clean the log directory being used by pjarcher913's portfolio-api program.
echo ===========================================
echo:
echo !!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!
echo !!!!   PLEASE READ CAREFULLY:   !!!!
echo !!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!
echo:
echo Irreversible changes will be made to your files if you choose to continue using this tool any further.
echo:
echo If the "./logs" directory does not yet exist as a sub-directory beneath where this tool is stored, then the tool will create "./logs" and exit.
echo:
echo If "./logs" already exists, then all log files contained-within will be deleted and the directory will be refreshed.
echo:
echo ===========================================
echo:
echo Please close the portfolio-api program before using this tool any further.
echo:

pause

IF EXIST "logs" (
echo: && echo [CONSOLE]: Existing "./logs" located. Refreshing directory...
rd /s /q "logs"
) ELSE (
echo: && echo [CONSOLE]: "./logs" does not yet exist. Creating directory...
)
mkdir "logs"

echo:
echo ######################################
echo ###  Done! Press any key to exit.  ###
echo ######################################
echo:

pause