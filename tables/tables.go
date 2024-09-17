package tables

import "github.com/GoAdminGroup/go-admin/plugins/admin/modules/table"

// generators is a map of table models.
//
// The key of generators is the prefix of table info url.
// The corresponding value is the Form and TableName data.
//
// http://{{config.Domain}}:{{Port}}/{{config.Prefix}}/info/{{key}}
//
// example:
//
// "usuarios"   => http://localhost:9033/admin/info/usuarios
// "posts"   => http://localhost:9033/admin/info/posts
// "authors" => http://localhost:9033/admin/info/authors
//
var Generators = map[string]table.Generator{
	"posts":     GetPostsTable,
	"users":     GetUserTable,
	"authors":   GetAuthorsTable,
	"profile":   GetProfileTable,
	"usuarios":  GetUsuariosTable,
	"contactos": GetContactoTable,
	"plantas":   GetPlantasTable,
}
