@echo off
( echo Articles
  echo =========
  echo.) > readme.md

for %%I in (articles\*.md) do goawk -v "fname=%%I" "BEGIN{ gsub(/\\/,'\/',fname) } match($0,/title: /){ printf '* [%%s](%%s)\n',substr($0,RSTART+RLENGTH),fname}" %%I >> readme.md

( echo.
  echo Books
  echo ======
  echo.) >> readme.md

for /D %%I in (books\*) do goawk -v "dir=%%I" "BEGIN{ gsub(/\\/,'\/',dir)} match($0,/title: /){ printf '* [%%s](%%s/readme.md)\n',substr($0,RSTART+RLENGTH),dir}" %%I\config.yaml >> readme.md
