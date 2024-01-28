package types

type ServiceLocation struct {
	Id           int          `json:"id"`
	Organization Organization `json:"organization"`
}
