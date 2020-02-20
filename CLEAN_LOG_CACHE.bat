@echo off

echo Clean Log Cache Tool v1.0 by pjarcher913
echo:
echo !!! This tool will delete all log files for the portfolio-api program. !!!
echo Please close the portfolio-api program before using this tool any further.
echo:

pause

IF EXIST "logs" (
rd /s /q "logs"
mkdir "logs"
) ELSE (
echo: && echo [ERROR]: "./logs" does not exist. Cancelling operation.
)

echo:
echo ######################################
echo ###  Done! Press any key to exit.  ###
echo ######################################
echo:

pause