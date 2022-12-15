package main

type Types interface {
	int |
		uint |
		byte |
		string |
		bool |
		float64 |
		map[any]any |
		[]any |
		*any
}
