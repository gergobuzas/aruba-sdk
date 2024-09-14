package container

import (
	"acloud/sdk"
)

func GetKaas(projectId, id string) {
	sdk.SendPayload("projects/"+projectId+"/providers/Aruba.Container/kaas/"+id, "GET", "")
}
