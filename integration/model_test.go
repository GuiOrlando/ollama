//go:build integration

package integration

import (
	"bufio"
	"context"
	"encoding/json"
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

type models struct {
	Name      string `json:"name"`
	ModelType string `json:"type"`
}

func TestModels(t *testing.T) {
	// Skip if there aren't any models to load

	// TODO - have an env var to steer to an alternate location for the jsonl files

	matches, _ := filepath.Glob(filepath.Join("models", "*.jsonl"))
	if len(matches) == 0 {
		t.Skip("no models to test")
	}
	ctx, cancel := context.WithTimeout(context.Background(), 9*time.Minute)
	defer cancel()
	client, _, cleanup := InitServerConnection(ctx, t)
	defer cleanup()

	for _, filename := range matches {

		f, err := os.Open(filename)
		require.NoError(t, err)
		defer f.Close()

		scanner := bufio.NewScanner(f)
		for scanner.Scan() {
			var testModel models
			err = json.Unmarshal(scanner.Bytes(), &testModel)
			require.NoError(t, err)
			t.Run(testModel.Name, func(t *testing.T) {
				switch testModel.ModelType {
				case "chat":
					require.NoError(t, PullIfMissing(ctx, client, testModel.Name))
					req, resp := GenerateRequests() // TODO - refine this
					req[0].Model = testModel.Name
					DoGenerate(ctx, t, client, req[0], resp[0], 60*time.Second, 5*time.Second)
				default:
					t.Skip("model type not yet supported " + testModel.ModelType)
				}
			})
		}
	}
}
