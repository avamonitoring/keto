package storage

//import "fmt"

// A list of roles.
//
// swagger:ignore
type Roles []Role

// Role represents a group of users that share the same role. A role could be an administrator, a moderator, a regular
// user or some other sort of role.
//
// swagger:ignore
type Role struct {
	// ID is the role's unique id.
	ID string `json:"id"`

	// Description is the description of the role.
	Description string `json:"description"`

	// Members is who belongs to the role.
	Members []string `json:"members"`
}

func (r *Role) withMembers(members []string, flavor string) *Role {
  if r == nil || len(members) == 0 {
    return r
  }
  
	switch flavor {
		case "glob":
			if globMatch(members[0], r.Members) { 
				return r
			} 
		case "exact":
			if exactMatch(members[0], r.Members) {
				return r
			}
		case "regex":
			if regexMatch(members[0], r.Members) {
				return r
			}
		default:
			return nil
	}
  return nil
}

func (r *Role) withIDs(ids []string, flavor string) *Role {
  if r == nil || len(ids) == 0 {
    return r
  }
  
	switch flavor {
		case "glob":
      if globMatch(ids[0], []string{r.ID}) {
				return r
			}
		case "exact":
			if exactMatch(ids[0], []string{r.ID}) {
				return r
			}
		case "regex":
			if regexMatch(ids[0], []string{r.ID}) {
				return r
			}
		default:
			return nil
	}
  return nil
}
