#!/bin/sh

APP="Pynote.app"
APPDIR=dist/mac/${APP}
mkdir -p $APPDIR/Contents/{MacOS,Resources}
go build -o $APPDIR/Contents/MacOS/pynote
cp config.json $APPDIR/Contents/MacOS/
cat > $APPDIR/Contents/Info.plist << EOF
<?xml version="1.0" encoding="UTF-8"?>
<!DOCTYPE plist PUBLIC "-//Apple//DTD PLIST 1.0//EN" "http://www.apple.com/DTDs/PropertyList-1.0.dtd">
<plist version="1.0">
<dict>
	<key>CFBundleExecutable</key>
	<string>pynote</string>
	<key>CFBundleIconFile</key>
	<string>icon.icns</string>
	<key>CFBundleIdentifier</key>
	<string>com.cabawa.pynote</string>
    <key>UIFileSharingEnabled</key>
    <true/>
    <key>LSSupportsOpeningDocumentsInPlace</key>
    <true/>
</dict>
</plist>
EOF
cp icons/icon.icns $APPDIR/Contents/Resources/icon.icns
find $APPDIR
