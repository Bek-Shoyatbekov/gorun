package main

import ()

func ExcludeFiles(excluded, main []string) []string {
	var result []string
	for _, e := range main {
		has := contains(excluded, e)
		if !has {
			result = append(result, e)
		}
	}
	return result
}

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
