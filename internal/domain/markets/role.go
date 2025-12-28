package markets

const (
	Admin  Role = "admin"
	Client Role = "client"
)

type Role string

// т.к. нету бизнес-логики с ролями не стал выносить в отдельный домен
// (используется как фильтр)

type allowedRoles []Role

func (r Role) isValid() bool {
	switch r {
	case Admin, Client:
		return true
	default:
		return false
	}
}
