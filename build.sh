RUN_NAME=taimer-backend

export GO111MODULE=on
mkdir -p output/bin output/log
cp scripts/* output
cp .env* output

go mod tidy
CGO_ENABLED=1 go build -tags netgo -a -v -o ./output/bin/$RUN_NAME .