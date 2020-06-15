package storage

// Policies is an array of policies.
//
// swagger:ignore
type Policies []Policy

// Policy specifies an ORY Access Policy document.
//
// swagger:ignore
type Policy struct {
	// ID is the unique identifier of the ORY Access Policy. It is used to query, update, and remove the ORY Access Policy.
	ID string `json:"id"`

	// Description is an optional, human-readable description.
	Description string `json:"description"`

	// Subjects is an array representing all the subjects this ORY Access Policy applies to.
	Subjects []string `json:"subjects"`

	// Resources is an array representing all the resources this ORY Access Policy applies to.
	Resources []string `json:"resources"`

	// Actions is an array representing all the actions this ORY Access Policy applies to.
	Actions []string `json:"actions"`

	// Effect is the effect of this ORY Access Policy. It can be "allow" or "deny".
	Effect string `json:"effect"`

	// Conditions represents a keyed object of conditions under which this ORY Access Policy is active.
	Conditions map[string]interface{} `json:"conditions"`
}

func (p *Policy) withSubjects(subjects []string, flavor string) *Policy {
  if p == nil || len(subjects) == 0 {
    return p
  }
  
	switch flavor {
		case "glob":
			if globMatch(subjects[0], p.Subjects) {
				return p
			}
		case "exact":
			if exactMatch(subjects[0], p.Subjects) {
				return p
			}
		case "regex":
			if regexMatch(subjects[0], p.Subjects) {
				return p
			}
		default:
			return nil
	}
  return nil
}

func (p *Policy) withResources(resources []string, flavor string) *Policy {
  if p == nil || len(resources) == 0 {
    return p
  }
  
	switch flavor {
		case "glob":
			if globMatch(resources[0], p.Resources) {
				return p
			}
		case "exact":
			if exactMatch(resources[0], p.Resources) {
				return p
			}
		case "regex":
			if regexMatch(resources[0], p.Resources) {
				return p
			}
		default:
			return nil
	}
  return nil
}

func (p *Policy) withActions(actions []string, flavor string) *Policy {
  if p == nil || len(actions) == 0 {
    return p
  }
  
	switch flavor {
		case "glob":
			if globMatch(actions[0], p.Actions) {
				return p
			}
		case "exact":
			if exactMatch(actions[0], p.Actions) {
				return p
			}
		case "regex":
			if regexMatch(actions[0], p.Actions) {
				return p
			}
		default:
			return nil
	}
  return nil
}

func (p *Policy) withIDs(ids []string, flavor string) *Policy {
  if p == nil || len(ids) == 0 {
    return p
  }
  
	switch flavor {
		case "glob":
			if globMatch(ids[0], []string{p.ID}) {
				return p
			}
		case "exact":
			if exactMatch(ids[0], []string{p.ID}) {
				return p
			}
		case "regex":
			if regexMatch(ids[0], []string{p.ID}) {
				return p
			}
		default:
			return nil
	}
  return nil
}
