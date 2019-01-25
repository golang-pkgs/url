package url

import "net"

// Rule defines the validation conditions
type Rule struct {
	Allow           []string
	Deny            []string
	DenyLocal       bool
	DefaultAllowAll bool
}

// CreateRuleChecker is a high order function that receive a rule and returns a rule checker
// If DenyLocal, request from
// If the Deny rule contains string "*", Checker func always return false, otherwise
//     If the Allow rules contains string "*", Checker always return true, otherwise the return value is Checker by rules, deny rules take higher priority.
func CreateRuleChecker(rule Rule) func(ip string) bool {
	if rule.DenyLocal {
		rule.Deny = append(rule.Deny, "127.0.0.1/32")
	} else {
		rule.Allow = append(rule.Allow, "127.0.0.1/32")
	}

	for _, denyRule := range rule.Deny {
		if denyRule == "*" {
			return func(ip string) bool {
				return false
			}
		}
	}

	for _, allowRule := range rule.Allow {
		if allowRule == "*" {
			return func(ip string) bool {
				return true
			}
		}
	}

	return func(addr string) bool {
		if addr == "" {
			return false
		}
		ip, _, err := net.SplitHostPort(addr)

		if err != nil {
			return false
		}

		if ip == "" {
			ip = addr
		}

		for _, denyRule := range rule.Deny {
			if denyRule != "*" {
				_, ipNet, err := net.ParseCIDR(denyRule)
				if err != nil {
					continue
				}

				c := ipNet.Contains(net.ParseIP(ip))
				if c {
					return false
				}
			}
		}

		for _, allowRule := range rule.Allow {
			if allowRule != "*" {
				_, ipNet, err := net.ParseCIDR(allowRule)
				if err != nil {
					continue
				}

				c := ipNet.Contains(net.ParseIP(ip))
				if c {
					return true
				}
			}
		}

		return rule.DefaultAllowAll
	}
}

// AllowAllChecker always return true
var AllowAllChecker = CreateRuleChecker(Rule{
	Allow: []string{"*"},
})

// DenyAllChecker always return false
var DenyAllChecker = CreateRuleChecker(Rule{
	Deny: []string{"*"},
})

// OnlyAllowAClassNetChecker returns true if the host ip is an A class address
var OnlyAllowAClassNetChecker = CreateRuleChecker(Rule{
	Allow: []string{"10.0.0.0/8"},
})

// OnlyDenyAClassNetChecker returns false if the host ip is an A class address
var OnlyDenyAClassNetChecker = CreateRuleChecker(Rule{
	Deny:            []string{"10.0.0.0/8"},
	DefaultAllowAll: true,
})

// OnlyAllowBClassNetChecker returns true if the host ip is a B class address
var OnlyAllowBClassNetChecker = CreateRuleChecker(Rule{
	Allow: []string{"72.16.0.0/12"},
})

// OnlyDenyBClassNetChecker returns false if the host ip is a B class address
var OnlyDenyBClassNetChecker = CreateRuleChecker(Rule{
	Deny:            []string{"72.16.0.0/12"},
	DefaultAllowAll: true,
})

// OnlyAllowCClassNetChecker returns true if the host ip is a C class address
var OnlyAllowCClassNetChecker = CreateRuleChecker(Rule{
	Allow: []string{"	192.168.0.0/16"},
})

// OnlyDenyCClassNetChecker returns false if the host ip is a C class address
var OnlyDenyCClassNetChecker = CreateRuleChecker(Rule{
	Deny: []string{"	192.168.0.0/16"},
	DefaultAllowAll: true,
})

// OnlyAllowInternalChecker is an alias to OnlyAllowAClassNetChecker
var OnlyAllowInternalChecker = OnlyAllowAClassNetChecker

// OnlyDenyInternalChecker is an alias to OnlyDenyAClassNetChecker
var OnlyDenyInternalChecker = OnlyDenyAClassNetChecker
