package pprofile

var emptyProfileID = ProfileID([16]byte{})

type ProfileID [16]byte

func NewProfileIDEmpty() ProfileID { _ = "STUB: not implemented"; return *new(ProfileID) }

func (ms ProfileID) String() string { _ = "STUB: not implemented"; return "" }

func (ms ProfileID) IsEmpty() bool { _ = "STUB: not implemented"; return false }
