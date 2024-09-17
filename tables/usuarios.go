package tables

import (
	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/modules/db"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/table"
	"github.com/GoAdminGroup/go-admin/template/types"
	"github.com/GoAdminGroup/go-admin/template/types/form"
)

func GetUsuariosTable(ctx *context.Context) table.Table {

	usuarios := table.NewDefaultTable(table.DefaultConfigWithDriverAndConnection("mysql", "mysql"))

	info := usuarios.GetInfo().SetFilterFormLayout(form.LayoutThreeCol)

	info.AddField("Id", "id", db.Int).
		FieldFilterable()
	info.AddField("Razon Social", "razon_social", db.Varchar).FieldFilterable(
		types.FilterType{
			Operator: types.FilterOperatorLike,
		},
	)

	info.AddField("Cuit_empresa", "cuit_empresa", db.Varchar).FieldFilterable(
		types.FilterType{
			Operator: types.FilterOperatorLike,
		},
	)

	info.AddField("Calle (domicilio legal)", "calle_domicilio_legal", db.Varchar)
	info.AddField("Numero (domicilio legal)", "numero_domicilio_legal", db.Varchar)
	info.AddField("Ciudad (domicilio legal)", "ciudad_domicilio_legal", db.Varchar)
	info.AddField("CP (domicilio legal)", "cp_domicilio_legal", db.Varchar)
	info.AddField("Calle (domicilio constituido)", "calle_domicilio_constituido", db.Varchar)
	info.AddField("Numero (domicilio constituido)", "numero_domicilio_constituido", db.Varchar)
	info.AddField("Ciudad (domicilio constituido)", "ciudad_domicilio_constituido", db.Varchar)
	info.AddField("CP (domicilio constituido)", "cp_domicilio_constituido", db.Varchar)

	info.SetTable("usuarios").SetTitle("Usuarios").SetDescription("Usuarios")

	formList := usuarios.GetForm()
	formList.AddField("Id", "id", db.Int, form.Default)
	formList.AddField("Razon Social", "razon_social", db.Varchar, form.Text)
	formList.AddField("CUIT", "cuit_empresa", db.Varchar, form.Text)
	formList.AddField("Calle (domicilio legal)", "calle_domicilio_legal", db.Varchar, form.Text)
	formList.AddField("Numero (domicilio legal)", "numero_domicilio_legal", db.Varchar, form.Text)
	formList.AddField("Ciudad (domicilio legal)", "ciudad_domicilio_legal", db.Varchar, form.Text)
	formList.AddField("CP (domicilio legal)", "cp_domicilio_legal", db.Varchar, form.Text)
	formList.AddField("Calle (domicilio constituido)", "calle_domicilio_constituido", db.Varchar, form.Text)
	formList.AddField("Numero (domicilio constituido)", "numero_domicilio_constituido", db.Varchar, form.Text)
	formList.AddField("Ciudad (domicilio constituido)", "ciudad_domicilio_constituido", db.Varchar, form.Text)
	formList.AddField("CP (domicilio constituido)", "cp_domicilio_constituido", db.Varchar, form.Text)

	formList.SetTable("usuarios").SetTitle("Usuarios").SetDescription("Usuarios")

	return usuarios
}
