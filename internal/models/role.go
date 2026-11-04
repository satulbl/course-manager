package models

// id ролей совпадают с roles.id в migrations/001_roles_users.sql:
// там роли вставляются в этом же порядке, и база выдаёт им id 1, 2, 3...
// В access-токене хранится именно id, права проверяются по этим константам.
const (
	RoleUser    int64 = 1
	RoleTeacher int64 = 2
	RoleAdmin   int64 = 3
)

// writeRoles — роли, которым можно создавать контент.
// Неизвестный id не попадает в map и получает отказ.
var writeRoles = map[int64]bool{
	RoleUser:    false,
	RoleAdmin:   true,
	RoleTeacher: true,
}
