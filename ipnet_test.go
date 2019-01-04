package url

import (
	"net"
)

type Rule struct {
	Allow           []string
	Deny            []string
	DefaultAllowAll bool
}

// If the Deny rule contains string "*", test func always return false, otherwise
//     If the Allow rules contains string "*", test always return true, otherwise the return value is test by rules, deny rules take higher priority.
func IpnetTest(rule Rule) func(ip string) bool {
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

	return func(ip string) bool {
		for _, denyRule := range rule.Deny {
			if denyRule != "*" {
				_, ipNet, err := net.ParseCIDR(denyRule)
				if err == nil {
					c := ipNet.Contains(net.ParseIP(ip))
					if c {
						return false
					}
				}
			}
		}

		for _, allowRule := range rule.Allow {
			if allowRule != "*" {
				_, ipNet, err := net.ParseCIDR(allowRule)
				if err == nil {
					c := ipNet.Contains(net.ParseIP(ip))
					if c {
						return true
					}
				}
			}
		}

		return rule.DefaultAllowAll
	}
}
