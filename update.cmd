@echo off
( echo Articles
  echo =========
  echo.) > readme.md

for %%I in (articles\*.md) do goawk "match($0,/title: /){ printf '* [%%s](%%s)\n',substr($0,RSTART+RLENGTH),FILENAME }" %%I >> readme.md

( echo.
  echo Books
  echo ======
  echo.) >> readme.md

for /D %%I in (books\*) do goawk -v "dir=%%I" "match($0,/title: /){ printf '* [%%s](%%s/readme.md)\n',substr($0,RSTART+RLENGTH),dir }" %%I\config.yaml >> readme.md
