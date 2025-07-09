setup-hooks:
	git config core.hooksPath .githooks

ssh-do:
	ssh root@207.154.221.44

vendor:
	cd src-tauri && cargo vendor && cd ..

build:
	rm -f bin/warpnet-desktop-linux
	cd src-tauri && cargo clean && cargo tauri build && cd ..
	cp src-tauri/target/release/warpnet-desktop bin/warpnet-desktop-linux

prerequisites-linux:
	apt update && apt install libwebkit2gtk-4.1-dev build-essential curl wget file libxdo-dev libssl-dev libayatana-appindicator3-dev librsvg2-dev

