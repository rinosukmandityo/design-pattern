package model

type ak47 struct {
	gun
}

func newAk47() IGun {
	return &ak47{
		gun: gun{
			name:  "AK47",
			power: 4,
		},
	}
}
