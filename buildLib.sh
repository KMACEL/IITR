echo "Welcome to Library Build Script"
echo "=======> cases"
cd cases
go fmt .
go build .
echo "=======> cases ======> operations"
cd operations
go fmt .
go build .
cd ..
echo "=======> cases ======> reports"
cd reports
go fmt .
go build .
cd ..
echo "=======> cases ======> tests"
cd tests
go fmt .
go build .
cd ..
cd ..
echo "=======> cases : OK"
echo "----------------------"

echo "=======> databasecenter"
cd databasecenter
go fmt .
go build .
cd ..
echo "=======> databasecenter : OK"
echo "----------------------"

echo "=======> errc"
cd errc
go fmt .
go build .
cd ..
echo "=======> errc : OK"
echo "----------------------"

echo "=======> logc"
cd logc
go fmt .
go build .
cd ..
echo "=======> logc : OK"
echo "----------------------"

echo "=======> mail"
cd errc
go fmt .
go build .
cd ..
echo "=======> mail : OK"
echo "----------------------"

echo "----------------------"
echo "=======> rest"
cd rest
go fmt .
go build .
echo "=======> rest ======> action"
cd action
go fmt .
go build .
cd ..
echo "=======> rest ======> adminarea"
cd adminarea
go fmt .
go build .
cd ..
echo "=======> rest ======> device"
cd device
go fmt .
go build .
cd ..
echo "=======> rest ======> download"
cd download
go fmt .
go build .
cd ..
echo "=======> rest ======> drom"
cd drom
go fmt .
go build .
cd ..
echo "=======> rest ======> profile"
cd profile
go fmt .
go build .
cd ..
echo "=======> rest ======> workgroup"
cd workgroup
go fmt .
go build .
cd ..
echo "=======> rest ======> workingset"
cd workingset
go fmt .
go build .
cd ..
cd ..
echo "=======> rest : OK"
echo "----------------------"

echo "=======> timop"
cd errc
go fmt .
go build .
cd ..
echo "=======> timop : OK"
echo "----------------------"

echo "=======> user"
cd user
go fmt .
go build .
cd ..
echo "=======> user : OK"
echo "----------------------"

echo "=======> writefile"
cd writefile
go fmt .
go build .
cd ..
echo "=======> writefile : OK"

go fmt .
go build .
echo "Build Finish"
