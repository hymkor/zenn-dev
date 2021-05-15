@echo off
setlocal
type nul > readme.md
for /D %%I in (*) do call :dir1 "%%I"
endlocal
exit /b

:dir1
    set "DIR=%~1"
    pushd "%DIR%"
    call update.cmd
    popd
    for /F "tokens=2* delims=:" %%I in ('findstr title: "%DIR%\config.yaml"') do (
        echo * [%%I]^(%DIR%/readme.md^) >> readme.md
    )
    exit /b
