package jsonfile

type Site struct {
	Id      Slug   `json:"id"`
	Name    Name   `json:"name"`
	Domain  Domain `json:"domain"`
	Website Url    `json:"website"`
}
