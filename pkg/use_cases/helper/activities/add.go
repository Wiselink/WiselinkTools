package activities

import (
	"context"
	"fmt"
	"log"

	"github.com/Wiselink/WiselinkTools/pkg/data/infrastructure/historyRepository"
	"github.com/Wiselink/WiselinkTools/pkg/domain/mongoModels"
	wisePkg "github.com/Wiselink/WiselinkTools/pkg/domain/response"
)

func AddActivityHelper(CompanyToken, ContactToken string, history mongoModels.Activities) wisePkg.Status {
	// CREACION DEL HISTORIAL
	hr := historyRepository.GetProvider()
	conn, client, err := hr.GetConeccion()
	if err != nil {
		fmt.Println("AddActivityHelper : ", err.Error())
		return wisePkg.DBInitError
	}
	defer client.Disconnect(context.Background())

	err = hr.AddActivity(conn, CompanyToken, ContactToken, history)
	if err != nil {
		log.Printf("AddActivityHelper : %v", err.Error())
		return wisePkg.DBInitError
	}
	return wisePkg.SuccessfulCreation
}
