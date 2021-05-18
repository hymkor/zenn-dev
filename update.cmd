@echo off
setlocal

rem **** update top readme.md ****

( echo Articles
  echo =========
  echo.) > readme.md

goawk "FNR==1{ fname=FILENAME ; gsub(/\\/,'\/',fname) } { gsub(/\x22/,'') } match($0,/title: /){ printf '* [%%s](%%s)\n',substr($0,RSTART+RLENGTH),fname}" articles\*.md >> readme.md

( echo.
  echo Books
  echo ======
  echo.) >> readme.md

for /D %%I in (books\*) do goawk -v "dir=%%I" "BEGIN{ gsub(/\\/,'\/',dir) } { gsub(/\x22/,'') } match($0,/title: /){ printf '* [%%s](%%s/readme.md)\n',substr($0,RSTART+RLENGTH),dir}" %%I\config.yaml >> readme.md

rem ***** update books/readme.md ****

for /D %%I in (books\*) do call :dir1 "%%I"
endlocal
exit /b

:dir1
    pushd "%~1"
    "%~dp0\zennmkreadme\zennmkreadme.exe" > readme.md
    popd
    exit /b
