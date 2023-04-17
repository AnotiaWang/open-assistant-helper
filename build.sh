echo "Building for Linux.."
GOOS=linux GOARCH=amd64 go build -o oa-helper_linux_amd64 .
GOOD=linux GOARCH=arm64 go build -o oa-helper_linux_arm64 .
echo "Building for Windows.."
GOOS=windows GOARCH=amd64 go build -o oa-helper_windows_amd64.exe .
echo "Building for Mac.."
GOOS=darwin GOARCH=amd64 go build -o oa-helper_darwin_amd64 .
GOOS=darwin GOARCH=arm64 go build -o oa-helper_darwin_arm64 .
