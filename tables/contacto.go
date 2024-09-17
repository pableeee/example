package tables

import (
	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/modules/db"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/table"
	"github.com/GoAdminGroup/go-admin/template/types/form"
)

func GetContactoTable(ctx *context.Context) table.Table {

	contacto := table.NewDefaultTable(table.DefaultConfigWithDriverAndConnection("mysql", "mysql"))

	info := contacto.GetInfo().HideFilterArea()

	info.AddField("Id", "id", db.Int).
		FieldFilterable()
	info.AddField("Planta_id", "planta_id", db.Int)
	info.AddField("Nombre", "nombre", db.Varchar)
	info.AddField("Cargo", "cargo", db.Varchar)
	info.AddField("Cuit_personal", "cuit_personal", db.Varchar)
	info.AddField("Telefono", "telefono", db.Varchar)
	info.AddField("Mail", "mail", db.Varchar)

	info.SetTable("contacto").SetTitle("Contacto").SetDescription("Contacto")

	formList := contacto.GetForm()
	formList.AddField("Id", "id", db.Int, form.Default)
	formList.AddField("Planta_id", "planta_id", db.Int, form.Number)
	formList.AddField("Nombre", "nombre", db.Varchar, form.Text)
	formList.AddField("Cargo", "cargo", db.Varchar, form.Text)
	formList.AddField("Cuit_personal", "cuit_personal", db.Varchar, form.Text)
	formList.AddField("Telefono", "telefono", db.Varchar, form.Text)
	formList.AddField("Mail", "mail", db.Varchar, form.Text)

	formList.SetTable("contacto").SetTitle("Contacto").SetDescription("Contacto")

	return contacto
}
