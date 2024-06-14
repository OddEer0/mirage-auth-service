package main

import (
	"encoding/json"
	"github.com/OddEer0/mirage-auth-service/scripts"
	"io/ioutil"
	"log"
	"os/exec"
	"path/filepath"
)

type GenI struct {
	BasePath string `json:"base_path"`
	Gens     []Gen  `json:"gens"`
}

type Gen struct {
	Pkg  string `json:"pkg"`
	Src  string `json:"src"`
	Dest string `json:"dest"`
}

func getGensData(root string) []Gen {
	mockgenJsonFile := filepath.Join(root, "scripts/mockgen_test/mockgen.json")
	content, err := ioutil.ReadFile(mockgenJsonFile)
	if err != nil {
		panic(err)
	}
	res := make([]GenI, 0, 25)
	err = json.Unmarshal(content, &res)
	if err != nil {
		panic(err)
	}
	result := make([]Gen, 0, len(res))
	for _, rg := range res {
		for _, gen := range rg.Gens {
			gen.Src = filepath.Join(rg.BasePath, gen.Src)
			result = append(result, gen)
		}
	}

	return result
}

// mockgen -source=internal/infrastructure/storage/postgres/iface.go -destination=internal/tests/mockgen/pg_mock/iface_mock.go -package=mockgenPostgres
func main() {
	root := scripts.GetAbsPathDir()
	mockgenOutputAbsDir := filepath.Join(root, "internal/tests/mockgen")
	gensData := getGensData(root)

	for _, gen := range gensData {
		cmd := exec.Command("mockgen", "-source="+filepath.Join(root, gen.Src), "-destination="+filepath.Join(mockgenOutputAbsDir, gen.Dest), "-package="+gen.Pkg)
		output, err := cmd.CombinedOutput()
		if err != nil {
			log.Fatalf("failed to run mockgen for %s: %v\nOutput: %s", gen.Src, err, string(output))
		}
		log.Printf("Mocks generated successfully for %s", gen.Src)
	}
}
