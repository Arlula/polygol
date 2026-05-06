package polygol

import (
	"fmt"
	"log"
	"os"
	"path"
	"path/filepath"
	"strings"
	"testing"
)

const (
	endToEndDir = "testdata/end-to-end"
)

var (
	// USE ME TO RUN ONLY ONE TEST
	targetOnly = ""
	opOnly     = ""

	// USE ME TO SKIP TESTS
	targetsSkip = []string{}
	opsSkip     = []string{}
)

type testCase struct {
	Name          string
	OperationType string
	ResultPath    string
}

func TestEndToEnd(t *testing.T) {
	targets, err := os.ReadDir(endToEndDir)
	if err != nil {
		t.Fatal(err)
	}

	for _, target := range targets {

		if targetOnly != "" && target.Name() != targetOnly {
			continue
		}

		if contains(targetsSkip, target.Name()) {
			fmt.Printf("skipping target %s...\n", target.Name())
			continue
		}

		if !target.IsDir() {
			continue
		}

		targetDir := path.Join(endToEndDir, target.Name())
		argsPath := path.Join(targetDir, "args.geojson")

		args, err := loadGeoms(argsPath)
		if err != nil {
			t.Fatal(err)
		}

		files, err := os.ReadDir(targetDir)
		if err != nil {
			log.Fatal(err)
		}

		testCases := []testCase{}

		for _, f := range files {
			if f.Name() == "args.geojson" {
				continue
			}
			ext := filepath.Ext(f.Name())
			if ext != ".geojson" {
				continue
			}
			fn := f.Name()
			opType := strings.TrimSuffix(fn, ext)
			fp := filepath.Join(targetDir, fn)
			if opType != "all" {
				testCases = append(testCases, testCase{
					Name:          fmt.Sprintf("%s-%s", target.Name(), opType),
					OperationType: opType,
					ResultPath:    fp,
				})
			} else {
				testCases = []testCase{
					{
						Name:          target.Name() + "-union",
						OperationType: opUnion,
						ResultPath:    fp,
					},
					{
						Name:          target.Name() + "-intersection",
						OperationType: opIntersection,
						ResultPath:    fp,
					},
					{
						Name:          target.Name() + "-xor",
						OperationType: opXor,
						ResultPath:    fp,
					},
					{
						Name:          target.Name() + "-difference",
						OperationType: opDifference,
						ResultPath:    fp,
					},
				}
			}
		}

		for _, testCase := range testCases {

			t.Run(testCase.Name, func(t *testing.T) {

				// t.Parallel() // run all end-to-end tests in parallel

				if contains(opsSkip, testCase.OperationType) {
					fmt.Printf("skipping op type %s...\n", testCase.OperationType)
				}

				geoms, err := loadGeoms(testCase.ResultPath)
				if err != nil {
					t.Fatal(err)
				}

				expected := geoms[0]

				result, err := newOperation(testCase.OperationType).run(args[0], args[1:]...)
				resetPrecision()
				if err != nil {
					t.Error(err)
				}
				same := equalMultiPoly(expected, result)
				if !same {
					// d, _ := diff.Diff(expected, result)
					// t.Fatal(d)
					t.Fatal("resulting geometry does not match expectations")

				}
			})
		}
	}
}

func contains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}
	return false
}
