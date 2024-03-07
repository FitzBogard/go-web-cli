package version

import (
	"encoding/json"
	"fmt"

	"go-web-cli/internal/pkg/version"

	"github.com/spf13/cobra"
)

// versionCmd represents the version command
var Cmd = &cobra.Command{
	Use: "version",
	Run: func(cmd *cobra.Command, args []string) {
		versionInfo := &version.PackageVersion{
			Version:     version.Version,
			FullCommit:  version.FullCommit,
			ReleaseDate: version.ReleaseDate,
			BuildDate:   version.BuildDate,
		}
		info, err := json.Marshal(versionInfo)
		if err != nil {
			panic(err)
		}
		fmt.Println(string(info))
	},
}
