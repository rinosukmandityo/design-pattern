package model

type maverick struct {
	gun
}

func newMaverick() IGun {
	return &maverick{
		gun: gun{
			name:  "Maverick",
			power: 5,
		},
	}
}
