package tables

import (
	"fmt"

	"github.com/GoAdminGroup/example/models"
	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/modules/db"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/table"
	"github.com/GoAdminGroup/go-admin/template/types"
	"github.com/GoAdminGroup/go-admin/template/types/form"

	editType "github.com/GoAdminGroup/go-admin/template/types/table"
)

func GetPlantasTable(ctx *context.Context) table.Table {

	users, err := models.GetUsuarios()
	if err != nil {
		panic(0)
	}
	var userOptions []types.FieldOption
	userMap := make(map[string]string)
	for _, u := range users {
		userOptions = append(userOptions, types.FieldOption{Text: u.RazonSocial, Value: fmt.Sprintf("%d", u.ID)})
		userMap[fmt.Sprintf("%d", u.ID)] = u.RazonSocial
	}

	plantas := table.NewDefaultTable(table.DefaultConfigWithDriverAndConnection("mysql", "mysql"))

	info := plantas.GetInfo().HideFilterArea()

	info.AddField("Id", "id", db.Int).
		FieldFilterable()

	info.AddField("Usuario", "usuario_id", db.Int).
		FieldDisplay(func(model types.FieldModel) interface{} {
			if v, found := userMap[model.Value]; found {
				return v
			}

			return model.Value
		}).FieldEditAble(editType.Select).FieldEditOptions(
		userOptions,
	)

	info.AddField("Nombre", "nombre", db.Varchar)
	info.AddField("Calle_domicilio", "calle_domicilio", db.Varchar)
	info.AddField("Numero_domicilio", "numero_domicilio", db.Varchar)
	info.AddField("Ciudad_domicilio", "ciudad_domicilio", db.Varchar)
	info.AddField("Cp_domicilio", "cp_domicilio", db.Varchar)
	info.AddField("Partido", "partido", db.Varchar)
	info.AddField("Provincia", "provincia", db.Varchar)
	info.AddField("Rubro_general", "rubro_general", db.Varchar)
	info.AddField("Rubro_especifico", "rubro_especifico", db.Varchar)
	info.AddField("Contacto_id", "contacto_id", db.Int)

	info.SetTable("plantas").SetTitle("Plantas").SetDescription("Plantas")

	formList := plantas.GetForm()
	formList.AddField("Id", "id", db.Int, form.Default)

	formList.AddField("Usuario_id", "usuario_id", db.Int, form.SelectSingle).
		FieldOptions(userOptions)
	formList.AddField("Nombre", "nombre", db.Varchar, form.Text)
	formList.AddField("Calle_domicilio", "calle_domicilio", db.Varchar, form.Text)
	formList.AddField("Numero_domicilio", "numero_domicilio", db.Varchar, form.Text)
	formList.AddField("Ciudad_domicilio", "ciudad_domicilio", db.Varchar, form.Text)
	formList.AddField("Cp_domicilio", "cp_domicilio", db.Varchar, form.Text)
	formList.AddField("Partido", "partido", db.Varchar, form.Text)
	formList.AddField("Provincia", "provincia", db.Varchar, form.Text)
	formList.AddField("Rubro_general", "rubro_general", db.Varchar, form.Text)
	formList.AddField("Rubro_especifico", "rubro_especifico", db.Varchar, form.Text)
	formList.AddField("Contacto_id", "contacto_id", db.Int, form.Number)

	formList.SetTable("plantas").SetTitle("Plantas").SetDescription("Plantas")

	return plantas
}
