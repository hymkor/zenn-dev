pushd "%~dp0\zennmkreadme"
go build
popd
"%~dp0\zennmkreadme\zennmkreadme.exe" > readme.md
