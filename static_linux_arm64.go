package gomf

/*
#cgo CPPFLAGS: -DDUCKDB_STATIC_BUILD
#cgo LDFLAGS: -lduckdb
#cgo darwin,arm64 LDFLAGS: -lc++ -L${SRCDIR}/linux-arm64/
#include <duckdb.h>
*/
import "C"
