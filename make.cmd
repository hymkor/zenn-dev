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

"%~dp0\zennmkreadme\zennmkreadme.exe" >> readme.md

endlocal
exit /b
