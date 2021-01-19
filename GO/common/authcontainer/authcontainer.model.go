package authcontainer

// AuthContainer stores IDP information
type AuthContainer struct {
	IdpName string
	IdpData map[string]string
}

// New returns an initialized AuthContainer
func New(idpName string, idpData map[string]string) *AuthContainer {
	authContainer := new(AuthContainer)
	authContainer.IdpName = idpName
	authContainer.IdpData = idpData
	return authContainer
}

// GOOGLE is a constant which matches up with the IdpName field in an AuthContainer
const GOOGLE string = "google"

// GOOGLEID is a constant which matches up with the id for Google in the idps table
const GOOGLEID int = 1

// JUSTASKING is a constant which matches up with the IdpName field in an AuthContainer
const JUSTASKING string = "justasking"

// JUSTASKINGID is a constant which matches up with the id for Google in the idps table
const JUSTASKINGID int = 2
