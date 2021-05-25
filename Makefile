install:
	cp build/bin/${BINNAME}_${GOOS}_${GOARCH} /usr/bin/${BINNAME}
	cp dist/debian/logo.png /usr/share/icons/hicolor/scalable/apps/squizit.png
	cp dist/debian/squizit.desktop /usr/share/applications