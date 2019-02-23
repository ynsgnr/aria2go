g++ -o libaria2go.dll aria2go.cpp -O3 -Wall -Wextra -fPIC -shared -l aria2 -DBUILD_DLL
go test -v