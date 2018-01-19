echo "Install build-essential"
sudo apt-get install build-essential

echo "Install sqlite3"
sudo apt-get install sqlite3

echo "Install Golang sqlite3 Driver"
go get github.com/mattn/go-sqlite3
