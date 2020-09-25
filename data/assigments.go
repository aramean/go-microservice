package data

type Assigment struct {
	ID          int    `json:"id"`
	Methodology string `json:"methodology"`
	Roles       string `json:"roles"`
	Description string `json:"description"`
	Project     string `json:"project"`
}

func GetAssigments() []*Assigment {
	return assigmentList
}

var assigmentList = []*Assigment{
	&Assigment{
		ID:          1,
		Methodology: "Software Development",
		Roles:       "Front-end GIS Developer & DevOps",
		Description: ".NET and React development of a plugin that can measure carbon dioxide emissions over different selected areas for Hajk (GIS solution based on OpenLayers). Pipelining a CI/CD with Azure DevOps using YAML to deploy the application.",
		Project:     "Karta",
	},
	&Assigment{
		ID:          2,
		Methodology: "Software Development",
		Roles:       "Full-stack Developer & DevOps",
		Description: ".NET development, UI/UX design, and implementation of a webpage with MVC called Skolmuseet, listing old artifacts items, links, and articles. Pipelining a CI/CD with Azure DevOps using YAML to deploy the application.",
		Project:     "Skolmuseet",
	},
}
