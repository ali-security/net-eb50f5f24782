// Standalone helper: build a proxy.golang.org-format module zip from a source dir.
// Usage: create_mod <module-path> <version> <src-dir> <out-zip>
package main

import (
	"fmt"
	"os"

	"golang.org/x/mod/module"
	"golang.org/x/mod/zip"
)

func main() {
	if len(os.Args) != 5 {
		fmt.Fprintln(os.Stderr, "usage: create_mod <module-path> <version> <src-dir> <out-zip>")
		os.Exit(2)
	}
	modPath, version, srcDir, outZip := os.Args[1], os.Args[2], os.Args[3], os.Args[4]

	f, err := os.Create(outZip)
	if err != nil {
		fmt.Fprintln(os.Stderr, "create out zip:", err)
		os.Exit(1)
	}
	defer f.Close()

	m := module.Version{Path: modPath, Version: version}
	if err := zip.CreateFromDir(f, m, srcDir); err != nil {
		fmt.Fprintln(os.Stderr, "CreateFromDir:", err)
		os.Exit(1)
	}
	fmt.Println("wrote", outZip)
}
