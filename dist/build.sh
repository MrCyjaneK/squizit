#!/bin/bash
set -e

# This script should not be verbose.
# Simply telling what it is doing is enough.
GITVERSION="+git"$(date +%Y%m%d%H%M)"."$(git log -n 1 | tr " " "\n" | head -2 | tail -1 | head -c 7)
function ok {
    echo "OK"
}

root=$(dirname $0)
cd "$root"
root=$(pwd)
vcode="1.0.4-"$(cat ../VERSION_CODE | head -1)
echo "Building Squizit - version: $vcode";
cd ..
rm -rf webui/html
cd webui/frontend
npm install
npm run build
cd "$root"
cd ..
cp webui/frontend/dist webui/html -r
rm -rf build/
goprodbuilds=$(pwd)"/build/"
~/go/bin/packr2 clean
~/go/bin/packr2
goprod \
    -combo="linux/amd64;linux/arm;linux/386;linux/arm64;windows/amd64;windows/386;darwin/amd64" \
    -binname="squizit" \
    -tags="guibrowser" \
    -version="$vcode" \
    -ldflags="-s -w"
    
# Server builds
goprod \
    -combo="linux/arm;linux/386;linux/arm64;linux/amd64;windows/amd64;windows/386" \
    -binname="squizit-server" \
    -tags="nogui" \
    -version="$vcode" \
    -ldflags="-s -w"


echo "/ ubtouch builds - daemon (custom location)"
echo -n -e "|- bin/squizit_ubtouch_arm64......"
goprod \
    -combo="linux/arm;linux/arm64;linux/amd64" \
    -binname="squizit-ubtouch" \
    -version="$vcode" \
    -ldflags="-X main.dataDir=/home/phablet/.local/share/squizit.anon -X main.Port=15932" \
    -package=false \
    -tags="nogui" \
    -ldflags="-s -w"

if [[ "X$SKIPANDROID" == "X" ]];
then
    echo "/ android builds - daemon (custom location) + NDK"
    # Android Version
    AV=21
    # NDK downloaded in android studio -> tools? -> sdk manager
    NDKV=$(ls ~/Android/Sdk/ndk/* -d | tr "/" "\n" | tail -1)
    NDK=~/Android/Sdk/ndk/$NDKV/toolchains/llvm/prebuilt/linux-x86_64/bin
    goprod \
        -combo="android/arm;android/386;android/arm64;android/amd64" \
        -binname="squizit" \
        -version="$vcode" \
        -ldflags="-X main.dataDir=/data/data/x.x.squizit/ -X main.Port=15932" \
        -ndk="$NDK" \
        -tags="nogui" \
        -package=false \
        -ldflags="-s -w"
fi
echo "===== Packaging"
echo "/ Packaging for Ubuntu Touch"
mkdir -p $goprodbuilds/click
for arch in arm64 arm amd64
do
    echo -n -e "|- bin/squizit_$arch.click............" | head -c 34
    cd $goprodbuilds
    cp ../dist/ubtouch ubtouch-$arch -r
    cd $goprodbuilds/ubtouch-$arch/
    clickable clean
    cp "$goprodbuilds/bin/squizit-ubtouch_linux_$arch" $(find . -name libbin.so)
    chmod +x $(find . -name libbin.so)
    sed -i 's/BUILD_VERSION_CODE/'$vcode'/g' manifest.json.in
    archC=$arch
    if [[ "$arch" == "arm" ]];
    then
        archC="armhf"
    fi
    clickable build --arch=$archC
    cp build/*/app/*.click $goprodbuilds/click/squizit_$arch.click
    ok
done
echo "\_ DONE"
if [[ "X$SKIPANDROID" == "X" ]];
then
    echo "/ Packaging for android"
    mkdir -p $goprodbuilds/apk/
    for arch in arm64 arm amd64 386 all
    do
        echo -n -e "|- bin/squizit.android.$arch.apk.........." | head -c 34
        cd "$goprodbuilds"
        cp ../dist/android android-target-$arch -r
        cd android-target-$arch
        touch "$goprodbuilds/android-target-$arch/app/src/main/resources/lib/x86_64/libbin.so"
        touch "$goprodbuilds/android-target-$arch/app/src/main/resources/lib/x86/libbin.so"
        touch "$goprodbuilds/android-target-$arch/app/src/main/resources/lib/armeabi-v7a/libbin.so"
        touch "$goprodbuilds/android-target-$arch/app/src/main/resources/lib/arm64-v8a/libbin.so"
        touch "$goprodbuilds/android-target-$arch/app/src/main/jniLibs/x86_64/libbin.so"
        touch "$goprodbuilds/android-target-$arch/app/src/main/jniLibs/x86/libbin.so"
        touch "$goprodbuilds/android-target-$arch/app/src/main/jniLibs/armeabi-v7a/libbin.so"
        touch "$goprodbuilds/android-target-$arch/app/src/main/jniLibs/arm64-v8a/libbin.so"
        case $arch in
        "amd64")
            cp "$goprodbuilds/bin/squizit_android_amd64" "$goprodbuilds/android-target-$arch/app/src/main/resources/lib/x86_64/libbin.so"
            cp "$goprodbuilds/bin/squizit_android_amd64" "$goprodbuilds/android-target-$arch/app/src/main/jniLibs/x86_64/libbin.so"
            ;;
        "386")
            cp "$goprodbuilds/bin/squizit_android_386"   "$goprodbuilds/android-target-$arch/app/src/main/resources/lib/x86/libbin.so"
            cp "$goprodbuilds/bin/squizit_android_386"   "$goprodbuilds/android-target-$arch/app/src/main/jniLibs/x86/libbin.so"
            ;;
        "arm")
            cp "$goprodbuilds/bin/squizit_android_arm"   "$goprodbuilds/android-target-$arch/app/src/main/resources/lib/armeabi-v7a/libbin.so"
            cp "$goprodbuilds/bin/squizit_android_arm"   "$goprodbuilds/android-target-$arch/app/src/main/jniLibs/armeabi-v7a/libbin.so"
            ;;
        "arm64")
            cp "$goprodbuilds/bin/squizit_android_arm64" "$goprodbuilds/android-target-$arch/app/src/main/resources/lib/arm64-v8a/libbin.so"
            cp "$goprodbuilds/bin/squizit_android_arm64" "$goprodbuilds/android-target-$arch/app/src/main/jniLibs/arm64-v8a/libbin.so"
            ;;
        "all")
            cp "$goprodbuilds/bin/squizit_android_amd64" "$goprodbuilds/android-target-$arch/app/src/main/resources/lib/x86_64/libbin.so"
            cp "$goprodbuilds/bin/squizit_android_386"   "$goprodbuilds/android-target-$arch/app/src/main/resources/lib/x86/libbin.so"
            cp "$goprodbuilds/bin/squizit_android_arm"   "$goprodbuilds/android-target-$arch/app/src/main/resources/lib/armeabi-v7a/libbin.so"
            cp "$goprodbuilds/bin/squizit_android_arm64" "$goprodbuilds/android-target-$arch/app/src/main/resources/lib/arm64-v8a/libbin.so"
            cp "$goprodbuilds/bin/squizit_android_amd64" "$goprodbuilds/android-target-$arch/app/src/main/jniLibs/x86_64/libbin.so"
            cp "$goprodbuilds/bin/squizit_android_386"   "$goprodbuilds/android-target-$arch/app/src/main/jniLibs/x86/libbin.so"
            cp "$goprodbuilds/bin/squizit_android_arm"   "$goprodbuilds/android-target-$arch/app/src/main/jniLibs/armeabi-v7a/libbin.so"
            cp "$goprodbuilds/bin/squizit_android_arm64" "$goprodbuilds/android-target-$arch/app/src/main/jniLibs/arm64-v8a/libbin.so"
            ;;
        esac
        chmod +x $(find . -name libbin.so)
        rm $(find . -name .gitkeep) || true
        sed -i 's/BUILD_VERSION_CODE/'$vcode'/g' app/build.gradle
        ./gradlew build
        cp ./app/build/outputs/apk/debug/app-debug.apk "$goprodbuilds/apk/squizit.android.$arch.apk"
        ok
    done
    echo "\_ OK"
fi
echo "DONE! Everything is inside build/"

~/go/bin/packr2 clean
