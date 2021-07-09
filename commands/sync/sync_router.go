package sync

import (
	"fmt"

	"github.com/iWinston/gf-cli/commands/sync/apifox"
)

func HelpRouter() {

}

func doSyncRouter(apiCollection *[]apifox.ApiItem) {
	for _, apiSystem := range *apiCollection {
		fmt.Printf("%+v", apiSystem)
	}
}
