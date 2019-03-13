package aria2go

/*
// Linux Build Tags
// ----------------
#cgo linux CXXFLAGS: -I${SRCDIR}/library -std=c++11 
#cgo linux LDFLAGS: -L${SRCDIR}/library -l aria2

// Windows Build Tags
// ----------------
#cgo windows CXXFLAGS: -I${SRCDIR}/library -std=c++11
#cgo windows LDFLAGS: -L${SRCDIR}/library -l:libaria2.lib -l:libcares.lib -l:libexpat.lib -l:libgmp.lib -l:libsqlite3.lib -l:libssh2.lib -l:libz.lib -lws2_32 -lbcrypt -liphlpapi -lcrypt32 -lsecur32

// Darwin Build Tags
// ----------------
#cgo darwin CXXFLAGS: -I${SRCDIR}/library -std=c++11
#cgo darwin LDFLAGS: -framework FlutterEmbedder

*/
import "C"