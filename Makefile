SHELL := /bin/zsh # Set Shell

# Go Commands
GODESKTOP=/Users/prad/Downloads/mobile-desktop/cmd/gomobile/gomobile
GOMOBILE=gomobile
GOCLEAN=$(GOMOBILE) clean
GOBIND=$(GOMOBILE) bind

# Plugin Directories
IOS_BUILDDIR=/Users/prad/Sonr/plugin/ios/Frameworks
IOS_ARTIFACT= $(IOS_BUILDDIR)/Core.framework
ANDROID_BUILDDIR=/Users/prad/Sonr/plugin/android/libs
ANDROID_ARTIFACT= $(ANDROID_BUILDDIR)/io.sonr.core.aar
MAC_BUILDDIR=/Users/prad/Sonr/core/build/Sonr.app/Contents/MacOS
MAC_ARTIFACT=$(MAC_BUILDDIR)/sonr_core
WIN_BUILDDIR=/Users/prad/Sonr/core/build
WIN_ARTIFACT=$(WIN_BUILDDIR)/sonr-core.exe

# Proto Directories
PB_PATH="/Users/prad/Sonr/core/internal/models"
CORE_PB_DIR="/Users/prad/Sonr/core/internal/models"
PLUGIN_PB_DIR="/Users/prad/Sonr/plugin/lib/models"

# Proto Build Commands
PB_BUILD_CORE="--go_out=$(CORE_PB_DIR)"
PB_BUILD_PLUGIN="--dart_out=$(PLUGIN_PB_DIR)"

all: Makefile
	@echo 'Sonr Core Module'
	@sed -n 's/^##//p' $<


## desktop     :   Builds Darwin and Windows Builds at Build Path
desktop: proto
	go clean -cache -x
	cd pkg && go build -o $(MAC_ARTIFACT)
	cd pkg && GOOS=windows GOARCH=amd64 go build -o $(WIN_ARTIFACT)
	@go mod tidy
	@cd /System/Library/Sounds && afplay Hero.aiff
	@echo ""
	@echo ""
	@echo "------------------------------------------------------------------"
	@echo "-------- ✅ ✅ ✅   FINISHED DESKTOP BUILD  ✅ ✅ ✅  --------------"
	@echo "------------------------------------------------------------------"

## └─ darwin   :   Compiles Desktop build of Sonr for MacOS
darwin:
	@echo ""
	@echo ""
	@echo "-----------------------------------------------------------"
	@echo "------------- 🖥  START DARWIN BUILD  🖥  -------------------"
	@echo "-----------------------------------------------------------"
	@go clean -cache
	cd pkg && go build -o $(MAC_ARTIFACT)
	@echo "Finished Building ➡ " && date
	@echo "--------------------------------------------------------------"
	@echo "------------- 🖥  COMPLETED DAWIN BULD  🖥  -------------------"
	@echo "--------------------------------------------------------------"
	@echo ""
	@cd $(MAC_BUILDDIR) && ./sonr_core

## └─ win      :   Compiles Desktop build of Sonr for Windows
win:
	@echo ""
	@echo ""
	@echo "-----------------------------------------------------------"
	@echo "------------- 🪟 START WINDOWS BUILD 🪟 --------------------"
	@echo "-----------------------------------------------------------"
	@go clean -cache
	cd pkg && GOOS=windows GOARCH=amd64 go build -ldflags -H=windowsgui -o $(WIN_ARTIFACT)
	@echo "Finished Binding ➡ " && date
	@echo "--------------------------------------------------------------"
	@echo "------------- 🪟 COMPLETED WINDOWS BULD 🪟 --------------------"
	@echo "--------------------------------------------------------------"
	@echo ""

## mobile      :   Builds Android and iOS Bind for Plugin Path
mobile: proto ios android
	@cd /Users/prad/Sonr/plugin && flutter clean
	@cd /Users/prad/Sonr/plugin/example && flutter clean
	@go mod tidy
	@cd /System/Library/Sounds && afplay Hero.aiff
	@echo ""
	@echo ""
	@echo "----------------------------------------------------------------"
	@echo "-------- ✅ ✅ ✅   FINISHED MOBILE BIND  ✅ ✅ ✅  --------------"
	@echo "----------------------------------------------------------------"


## └─ android  :   Builds Android Bind at Plugin Path
android:
	@echo ""
	@echo ""
	@echo "--------------------------------------------------------------"
	@echo "--------------- 🤖 BEGIN ANDROID BIND 🤖 ----------------------"
	@echo "--------------------------------------------------------------"
	cd bind && GODEBUG=asyncpreemptoff=1 gomobile bind -ldflags='-s -w' -target=android/arm64 -v -o $(ANDROID_ARTIFACT)
	@go mod tidy
	@cd /System/Library/Sounds && afplay Glass.aiff
	@echo "Finished Binding ➡ " && date
	@echo "--------------------------------------------------------------"
	@echo "------------- 🤖  COMPLETE ANDROID BIND 🤖  -------------------"
	@echo "--------------------------------------------------------------"
	@echo ""


## └─ ios      :   Builds iOS Bind at Plugin Path
ios:
	@echo ""
	@echo ""
	@echo "--------------------------------------------------------------"
	@echo "-------------- 📱 BEGIN IOS BIND 📱 ---------------------------"
	@echo "--------------------------------------------------------------"
	cd bind && GODEBUG=asyncpreemptoff=1 gomobile bind -ldflags='-s -w' -target=ios/arm64 -v -o $(IOS_ARTIFACT)
	@go mod tidy
	@cd /System/Library/Sounds && afplay Glass.aiff
	@echo "Finished Binding ➡ " && date
	@echo "--------------------------------------------------------------"
	@echo "-------------- 📱 COMPLETE IOS BIND 📱 ------------------------"
	@echo "--------------------------------------------------------------"
	@echo ""

## proto       :   Compiles Protobuf models for Core Library and Plugin
proto:
	@echo ""
	@echo ""
	@echo "--------------------------------------------------------------"
	@echo "------------- 🛸 START PROTOBUFS COMPILE 🛸 -------------------"
	@echo "--------------------------------------------------------------"
	@cd internal/models && protoc -I. --proto_path=$(PB_PATH) $(PB_BUILD_CORE) api.proto data.proto core.proto user.proto
	@cd internal/models && protoc -I. --proto_path=$(PB_PATH) $(PB_BUILD_PLUGIN) api.proto data.proto user.proto
	@echo "Finished Compiling ➡ " && date
	@echo "--------------------------------------------------------------"
	@echo "------------- 🛸 COMPILED ALL PROTOBUFS 🛸 --------------------"
	@echo "--------------------------------------------------------------"
	@echo ""

## run         :   Runs current Darwin Build
run:
	cd $(MAC_BUILDDIR) && ./sonr_core

## upgrade     :   Builds ALL supported devices
upgrade: proto mobile desktop

## [reset]     :   Cleans Gomobile, Removes Framworks from Plugin, and Initializes Gomobile
reset:
	cd bind && $(GOCLEAN)
	go mod tidy
	go clean -cache -x
	go clean -modcache -x
	rm -rf $(IOS_BUILDDIR)
	rm -rf $(ANDROID_BUILDDIR)
	mkdir -p $(IOS_BUILDDIR)
	mkdir -p $(ANDROID_BUILDDIR)
	cd bind && gomobile init
