cd ../cmd/bot
go build -o bot.exe
cd ../../deployment

docker-compose -f docker-compose.yaml down

docker-compose -f docker-compose.yaml up -d

cd ../cmd/bot
.\bot.exe