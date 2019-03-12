package aria2go

/*
// Linux Build Tags
// ----------------
#cgo linux CXXFLAGS: -I${SRCDIR}/library -std=c++11
#cgo linux LDFLAGS: -L${SRCDIR}/library -l aria2 -l:libcares.a -l:libexpat.a -l:libgmp.a -l:libsqlite3.a -l:libssh2.a -l:libz.a

// Windows Build Tags
// ----------------
#cgo windows CXXFLAGS: -I${SRCDIR}/library -std=c++11
#cgo windows LDFLAGS: -L${SRCDIR}/library -l:libaria2.dll.a -l:libcares.dll.a -l:libexpat.dll.a -l:libgmp.dll.a -l:libsqlite3.dll.a -l:libssh2.dll.a

// Darwin Build Tags
// ----------------
#cgo darwin CXXFLAGS: -I${SRCDIR}/library -std=c++11
#cgo darwin LDFLAGS: -framework FlutterEmbedder

*/
import "C"