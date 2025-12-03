package core

type Paswot struct{}

func NewPaswot() *Paswot {
	return &Paswot{}
}

type PaswotWithSalt struct {
	*Paswot
	Salt string
}

func NewPaswotWithSalt(salt string) *PaswotWithSalt {
	return &PaswotWithSalt{Paswot: &Paswot{}, Salt: salt}
}

type PaswotWithSaltAndPepper struct {
	*PaswotWithSalt
	Pepper string
}

func NewPaswotWithSaltAndPepper(salt, pepper string) *PaswotWithSaltAndPepper {
	return &PaswotWithSaltAndPepper{PaswotWithSalt: &PaswotWithSalt{Paswot: &Paswot{}, Salt: salt}, Pepper: pepper}
}
