readme.md:
	go fmt update.go
	go run update.go > readme.md

preview:
	npx zenn preview

new:
	npx zenn new:article

new-book:
	npx zenn new:book

init:
	scoop install nodejs
	npm init --yes
	npm install zenn-cli
	npx zenn init

update:
	npm install zenn-cli@latest

